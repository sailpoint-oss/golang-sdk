const fs = require("fs");
const path = require("path");
const getAllFiles = function (dirPath, arrayOfFiles) {
  files = fs.readdirSync(dirPath);
  arrayOfFiles = arrayOfFiles || [];
  files.forEach(function (file) {
    if (fs.statSync(dirPath + "/" + file).isDirectory()) {
      arrayOfFiles = getAllFiles(dirPath + "/" + file, arrayOfFiles);
    } else {
      arrayOfFiles.push(
        path.join(__dirname.replaceAll("sdk-resources", ""), dirPath, "/", file)
      );
    }
  });
  return arrayOfFiles;
};

const moveFile = function (sourcePath, targetPath) {
  // Create the target directory if it doesn't exist
  const directory = path.dirname(targetPath);
  if (!fs.existsSync(directory)) {
    fs.mkdirSync(directory, { recursive: true });
  }

  // Move the file
  fs.renameSync(sourcePath, targetPath);
  // console.log(`Moved: ${sourcePath} to ${targetPath}`);
};

const renameFileToIndices = function (filePath) {
  // Determine the new file path by changing the file's name to 'indices'
  let dirPath = path.dirname(filePath); // Gets the directory path of the current file

  const targetDir = path.join(dirPath, "Models"); // Appends 'models' to the directory path

  const newFilePath = path.join(targetDir, "indices.md"); // Constructs the full target path

  // Rename the file
  fs.renameSync(filePath, newFilePath);
  console.log(`Renamed: ${filePath} to ${newFilePath}`);
};

const fixFiles = function (myArray) {
  let fixCheck = 0;
  for (const file of myArray) {
    if (file.includes("go.mod") || file.includes("go.sum")) {
      fs.unlinkSync(file);
      continue;
    }

    let fileOut = [];
    let madeChange = false;
    let rawdata = fs.readFileSync(file).toString();
    let rawDataArra = rawdata.split("\n");

    // change the poor naming for variables

    if (
      rawdata.includes("map[string]interface{}var") ||
      rawdata.includes("ArrayOf*string") ||
      file.includes("model_json_patch_operation_value.go")
    ) {
      for (const line of rawDataArra) {
        if (line.includes("ginterface{}")) {
          fileOut.push(line.replaceAll("ginterface{}", "ginterface"));
          madeChange = true;
        } else if (line.includes("map[string]interface{}var")) {
          fileOut.push(line.replaceAll("map[string]interface{}var", "mapvar"));
          madeChange = true;
        } else if (line.includes("ArrayOf*string")) {
          fileOut.push(line.replaceAll("ArrayOf*string", "ArrayOfstring"));
          madeChange = true;
        } else {
          fileOut.push(line);
        }
      }
      rawDataArra = fileOut.slice();
      fileOut = [];
    }

    //fix duplicate type names in beta api
    if (
      file.includes("api_iai_peer_group_strategies.go") ||
      file.includes("api_iai_recommendations.go")
    ) {
      for (const line of rawDataArra) {
        if (line.includes("ApiGetOutliersRequest")) {
          fileOut.push(
            line.replaceAll("ApiGetOutliersRequest", "ApiGetOutliersRequest2")
          );
          madeChange = true;
        } else if (line.includes("ApiGetMessageCatalogsRequest")) {
          fileOut.push(
            line.replaceAll(
              "ApiGetMessageCatalogsRequest",
              "ApiGetMessageCatalogsRequest2"
            )
          );
          madeChange = true;
        } else {
          fileOut.push(line);
        }
      }
      rawDataArra = fileOut.slice();
      fileOut = [];
    }

    // fix syntax error in model_patch
    if (file.includes("model_patch_potential_role_request_inner.go")) {
      for (const line of rawDataArra) {
        if (line.includes("this.Op = op")) {
          fileOut.push(line.replaceAll("this.Op = op", "this.Op = &op"));
          madeChange = true;
        } else {
          fileOut.push(line);
        }
      }
      rawDataArra = fileOut.slice();
      fileOut = [];
    }

    // fix syntax error in model_subscription_patch_request_inner_value
    if (file.includes("model_subscription_patch_request_inner_value.go")) {
      for (const line of rawDataArra) {
        if (
          line.includes("[]SubscriptionPatchRequestInnerValueAnyOfInnervar")
        ) {
          fileOut.push(
            line.replaceAll(
              "[]SubscriptionPatchRequestInnerValueAnyOfInnervar",
              "SubscriptionPatchRequestInnerValueAnyOfInnervar"
            )
          );
          madeChange = true;
        } else {
          fileOut.push(line);
        }
      }
      rawDataArra = fileOut.slice();
      fileOut = [];
    }

    //adjust the document type naming to fix the duplicate type errors
    if (
      file.includes("model_event_document.go") ||
      file.includes("model_event.go")
    ) {
      for (const line of rawDataArra) {
        if (line.includes("Type DocumentType")) {
          fileOut.push(
            line.replace("Type DocumentType", "DocumentType DocumentType")
          );
          madeChange = true;
        } else if (line.includes("this.Type = type_")) {
          fileOut.push(
            line.replace("this.Type = type_", "this.DocumentType = type_")
          );
          madeChange = true;
        } else if (
          line.includes("func (o *EventDocument) GetType() DocumentType {")
        ) {
          fileOut.push(
            line.replace(
              "func (o *EventDocument) GetType() DocumentType {",
              "func (o *EventDocument) GetDocumentType() DocumentType {"
            )
          );
          madeChange = true;
        } else if (line.includes("func (o *Event) GetType() DocumentType {")) {
          fileOut.push(
            line.replace(
              "func (o *Event) GetType() DocumentType {",
              "func (o *Event) GetDocumentType() DocumentType {"
            )
          );
          madeChange = true;
        } else if (
          line.includes("return o.Type") &&
          !line.includes("return o.Type, true")
        ) {
          fileOut.push(line.replace("return o.Type", "return o.DocumentType"));
          madeChange = true;
        } else if (
          line.includes(
            "func (o *EventDocument) GetTypeOk() (*DocumentType, bool) {"
          )
        ) {
          fileOut.push(
            line.replace(
              "func (o *EventDocument) GetTypeOk() (*DocumentType, bool) {",
              "func (o *EventDocument) GetDocumentTypeOk() (*DocumentType, bool) {"
            )
          );
          madeChange = true;
        } else if (
          line.includes("func (o *Event) GetTypeOk() (*DocumentType, bool) {")
        ) {
          fileOut.push(
            line.replace(
              "func (o *Event) GetTypeOk() (*DocumentType, bool) {",
              "func (o *Event) GetDocumentTypeOk() (*DocumentType, bool) {"
            )
          );
          madeChange = true;
        } else if (line.includes("	return &o.Type, true")) {
          fileOut.push(
            line.replace("	return &o.Type, true", "	return &o.DocumentType, true")
          );
          madeChange = true;
        } else if (
          line.includes("func (o *EventDocument) SetType(v DocumentType) {")
        ) {
          fileOut.push(
            line.replace(
              "func (o *EventDocument) SetType(v DocumentType) {",
              "func (o *EventDocument) SetDocumentType(v DocumentType) {"
            )
          );
          madeChange = true;
        } else if (line.includes("func (o *Event) SetType(v DocumentType) {")) {
          fileOut.push(
            line.replace(
              "func (o *Event) SetType(v DocumentType) {",
              "func (o *Event) SetDocumentType(v DocumentType) {"
            )
          );
          madeChange = true;
        } else if (line.includes("o.Type = v")) {
          fileOut.push(line.replace("o.Type = v", "o.DocumentType = v"));
          madeChange = true;
        } else {
          fileOut.push(line);
        }
      }
      rawDataArra = fileOut.slice();
      fileOut = [];
    }

    for (const line of rawDataArra) {
      if (line.includes("**Indices** | Pointer to [**[]Index**](Index)")) {
        fileOut.push(
          line.replaceAll(
            "**Indices** | Pointer to [**[]Index**](Index)",
            "**Indices** | Pointer to [**[]Index**](Indices)"
          )
        );
        madeChange = true;
      } else if (line.includes("[***os.File**](../models/os-file)")) {
        fileOut.push(
          line.replaceAll(
            "[***os.File**](../models/os-file)",
            "[***os.File**](https://pkg.go.dev/os)"
          )
        );
        madeChange = true;
      } else if (line.includes("[**map[string][]Column**](array)")){
        fileOut.push(
          line.replaceAll(
            "[**map[string][]Column**](array)",
            "[**map[string][]Column**]"
          )
        );
        madeChange = true;
      } else {
        fileOut.push(line);
      }
    }

    rawDataArra = fileOut.slice();
    fileOut = [];

    if (madeChange) {
      fixCheck += 1;
      fs.writeFileSync(file, rawDataArra.join("\n"));
    }

    if (file.endsWith("API.md")) {
      let dirPath = path.dirname(file); // Gets the directory path of the current file

      const targetDir = path.join(dirPath, "Methods"); // Appends 'methods' to the directory path

      const targetPath = path.join(targetDir, path.basename(file)); // Constructs the full target path

      // Check if the file exists before trying to move it
      if (fs.existsSync(file)) {
        moveFile(file, targetPath);
        fixCheck++;
      } else {
        console.error(`File not found: ${file}`);
      }
    } else if (
      file.includes(`\\golang-sdk\\api_v3\\docs`) ||
      file.includes(`\\golang-sdk\\api_beta\\docs`) ||
      file.includes(`/golang-sdk/api_v3/docs`) ||
      file.includes(`/golang-sdk/api_beta/docs`)
    ) {
      if (file.includes("Index.md")) {
        if (fs.existsSync(file)) {
          renameFileToIndices(file);
          fixCheck++;
          continue; // Skip further processing for this file
        } else {
          console.error(`File not found: ${file}`);
          continue;
        }
      } else {
        let dirPath = path.dirname(file); // Gets the directory path of the current file

        const targetDir = path.join(dirPath, "Models"); // Appends 'models' to the directory path

        const targetPath = path.join(targetDir, path.basename(file)); // Constructs the full target path

        // Check if the file exists before trying to move it
        if (fs.existsSync(file)) {
          moveFile(file, targetPath);
          fixCheck++;
        } else {
          console.error(`File not found: ${file}`);
        }
      }
    }
  }

  console.log(`fixed ${fixCheck} files`);
};

let myArray = [];
getAllFiles(process.argv[2], myArray);

fixFiles(myArray);
