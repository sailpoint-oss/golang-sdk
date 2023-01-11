const fs = require("fs");
const path = require("path");
const getAllFiles = function (dirPath, arrayOfFiles) {
  files = fs.readdirSync(dirPath);
  arrayOfFiles = arrayOfFiles || [];
  files.forEach(function (file) {
    if (fs.statSync(dirPath + "/" + file).isDirectory()) {
      arrayOfFiles = getAllFiles(dirPath + "/" + file, arrayOfFiles);
    } else {
      arrayOfFiles.push(path.join(__dirname.replaceAll('sdk-resources',''), dirPath, "/", file));
    }
  });
  return arrayOfFiles;
};


const fixFiles = function (myArray) {
  for (const file of myArray) {
    let fileOut = [];
    let madeChange = false;
    let rawdata = fs.readFileSync(file).toString();
    let rawDataArra = rawdata.split("\n");
  
    if (file.includes("model_transform_definition_attributes_value.go")) {
      console.log("Found it");
    }
  
    // add the time import when it is missing in files
    if (
      (rawdata.includes("NullableTime") || rawdata.includes("*time.Time")) &&
      !rawdata.includes('"time"')
    ) {
      for (const line of rawDataArra) {
        if (line.includes('"encoding/json"')) {
          fileOut.push('	"time"');
          madeChange = true;
        }
        fileOut.push(line);
      }
      rawDataArra = fileOut.slice();
      fileOut = [];
    }
  
    // change the poor naming for variables
    if (rawdata.includes("map[string]interface{}var") || rawdata.includes("[]JsonPatchOperationValueAnyOfInnervar") || rawdata.includes("ArrayOf*string")) {
      for (const line of rawDataArra) {
        if (line.includes("map[string]interface{}var")) {
          fileOut.push(line.replaceAll("map[string]interface{}var", "mapvar"));
          madeChange = true;
        }
        else if (line.includes("[]JsonPatchOperationValueAnyOfInnervar")) {
          fileOut.push(line.replaceAll("[]JsonPatchOperationValueAnyOfInnervar", "jsonPatchOperationValueAnyOfInnervar"));
          madeChange = true;
        }
        else if (line.includes("[]JsonPatchOperationValueAnyOfInnerBetaBetavar")) {
          fileOut.push(line.replaceAll("[]JsonPatchOperationValueAnyOfInnerBetaBetavar", "jsonPatchOperationValueAnyOfInnerBetaBetavar"));
          madeChange = true;
        }
        else if (line.includes("ArrayOf*string")) {
          fileOut.push(line.replaceAll("ArrayOf*string", "ArrayOfstring"));
          madeChange = true;
        } else {
          fileOut.push(line);
        }
      }
      rawDataArra = fileOut.slice();
      fileOut = [];
    }
  
  
  
    // fix duplicate type names
    if (file.includes("api_iai_peer_group_strategies.go") || file.includes("api_iai_recommendations.go")) {
      for (const line of rawDataArra) {
        if (line.includes("ApiGetOutliersRequest")) {
          fileOut.push(line.replaceAll("ApiGetOutliersRequest", "ApiGetOutliersRequest2"));
          madeChange = true;
        }
        else if (line.includes("ApiGetMessageCatalogsRequest")) {
          fileOut.push(line.replaceAll("ApiGetMessageCatalogsRequest", "ApiGetMessageCatalogsRequest2"));
          madeChange = true;
        } else {
          fileOut.push(line);
        }
      }
      rawDataArra = fileOut.slice();
      fileOut = [];
    }
  
  
    // adjust the document type naming to fix the duplicate type errors
    if (file.includes("model_event.go")) {
      for (const line of rawDataArra) {
        if (line.includes("Type DocumentType")) {
          fileOut.push(line.replace("Type DocumentType", "DocumentType DocumentType"));
          madeChange = true;
        }
        else if (line.includes("this.Type = type_")) {
          fileOut.push(line.replace("this.Type = type_", "this.DocumentType = type_"));
          madeChange = true;
        }
        else if (line.includes("func (o *Event) GetType() DocumentType {")) {
          fileOut.push(line.replace("func (o *Event) GetType() DocumentType {", "func (o *Event) GetDocumentType() DocumentType {"));
          madeChange = true;
        }
        else if (line === "\treturn o.Type") {
          fileOut.push(line.replace("return o.Type", "return o.DocumentType"));
          madeChange = true;
        }
        else if (line.includes("func (o *Event) GetTypeOk() (*DocumentType, bool) {")) {
          fileOut.push(line.replace("func (o *Event) GetTypeOk() (*DocumentType, bool) {", "func (o *Event) GetDocumentTypeOk() (*DocumentType, bool) {"));
          madeChange = true;
        }
        else if (line.includes("	return &o.Type, true")) {
          fileOut.push(line.replace("	return &o.Type, true", "	return &o.DocumentType, true"));
          madeChange = true;
        }
        else if (line.includes("func (o *Event) SetType(v DocumentType) {")) {
          fileOut.push(line.replace("func (o *Event) SetType(v DocumentType) {", "func (o *Event) SetDocumentType(v DocumentType) {"));
          madeChange = true;
        }
        else if (line.includes("o.Type = v")) {
          fileOut.push(line.replace("o.Type = v", "o.DocumentType = v"));
          madeChange = true;
        } else {
          fileOut.push(line);
        }
      }
      rawDataArra = fileOut.slice();
      fileOut = [];
    }
  
    if (madeChange) {
      fs.writeFileSync(file, rawDataArra.join("\n"));
    }
  }
}




let myArray = [];
getAllFiles(process.argv[2], myArray);

fixFiles(myArray)

