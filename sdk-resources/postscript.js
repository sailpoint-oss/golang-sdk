const fs = require("fs").promises;
const path = require("path");

const getAllFiles = async function (dirPath, arrayOfFiles = []) {
  const files = await fs.readdir(dirPath);

  for (const file of files) {
    const fullPath = path.join(dirPath, file);
    const stat = await fs.stat(fullPath);

    if (stat.isDirectory()) {
      arrayOfFiles = await getAllFiles(fullPath, arrayOfFiles);
    } else {
      arrayOfFiles.push(path.join(__dirname.replace(/sdk-resources/g, ''), dirPath, file));
    }
  }

  return arrayOfFiles;
};


const mergeCodeExampleFiles = async (srcDir) => {
  try {
    // Check if the source directory exists asynchronously
    await fs.access(srcDir);

    const destFile = path.join(srcDir, 'go_code_examples_overlay.yaml');
    const files = await fs.readdir(srcDir);

    // Use writeFile instead of createWriteStream to simplify handling in async/await context
    let allContent = '';

    for (const file of files) {
      const filePath = path.join(srcDir, file);
      const stat = await fs.stat(filePath);

      if (stat.isFile()) {
        const fileContent = await fs.readFile(filePath, 'utf8');
        allContent += fileContent;  // Concatenate file contents
        await fs.unlink(filePath);  // Delete the original file after merging
      }
    }

    // Write all content into the destination file
    await fs.writeFile(destFile, allContent);

    console.log(`All files have been merged into ${destFile}`);
  } catch (error) {
    console.error(`Error processing the directory: ${error.message}`);
  }
};

const renameFileToIndices = async (filePath) => {
  try {
    const newFilePath = path.join(path.dirname(filePath), "Indices.md");
    await fs.rename(filePath, newFilePath);
  } catch (error) {
    console.error(`Failed to rename file: ${error.message}`);
  }
};

const createDir = async (srcDir, dirName) => {
  const newDir = path.join(srcDir, dirName);
  try {
    await fs.mkdir(newDir, { recursive: true });
  } catch (err) {
    console.error(`Error creating directory: ${err.message}`);
  }
  return newDir;
};

const moveFiles = async (srcPath, destPath, filename = null) => {
  try {
    if (!await fs.stat(destPath)) {
      await fs.mkdir(destPath, { recursive: true });
    }

    if (filename) {
      const filePath = path.join(srcPath, filename);
      if (await fs.stat(filePath)) {
        await fs.rename(filePath, path.join(destPath, filename));
      } else {
        console.error(`File ${filename} does not exist in the source path.`);
      }
    } else {
      const files = await fs.readdir(srcPath);
      for (const file of files) {
        const filePath = path.join(srcPath, file);
        if ((await fs.stat(filePath)).isFile()) {
          await fs.rename(filePath, path.join(destPath, file));
        }
      }
    }
  } catch (error) {
    console.error(`Error moving files: ${error.message}`);
  }
};

const processDirectory = async (srcDir) => {
  const files = await fs.readdir(srcDir);
  const examplesDir = await createDir(srcDir, 'Examples');
  const methodsDir = await createDir(srcDir, 'Methods');
  const modelsDir = await createDir(srcDir, 'Models');

  for (const file of files) {
    const currentPath = path.join(srcDir, file);
    const destExamplePath = path.join(examplesDir, file);
    const destMethodPath = path.join(methodsDir, file);
    const destModelPath = path.join(modelsDir, file);

    try {
      const stat = await fs.stat(currentPath);
      if (stat.isDirectory()) {
        if (file.includes('developerSite_code_examples')) {
          await moveFiles(currentPath, examplesDir);
        } else if (file.includes('API.md')) {
          await moveFiles(currentPath, methodsDir);
        } else if (file.includes('index.md') || file.includes('Methods') || file.includes('Models')) {
          if (file.includes('Methods')) {
            await moveFiles(currentPath, methodsDir);
          } else {
            await moveFiles(currentPath, modelsDir);
          }
        } else {
          await moveFiles(currentPath, modelsDir);
        }
      } else {
        if (file.includes('developerSite_code_examples')) {
          await fs.rename(currentPath, destExamplePath);
        } else if (file.includes('API.md')) {
          await fs.rename(currentPath, destMethodPath);
        } else {
          await fs.rename(currentPath, destModelPath);
        }
      }
    } catch (err) {
      console.error(`Error processing file: ${file}, ${err.message}`);
    }
  }
};

const fixFiles = async function (myArray) {
  let fixCheck = 0;
  for (const file of myArray) {
    let fileOut = [];
    let madeChange = false;
    let rawdata;

    try {
      rawdata = await fs.readFile(file, 'utf-8');
    } catch (err) {
      console.error(`Failed to read file: ${file}, ${err.message}`);
      continue;
    }

    let rawDataArra = rawdata.split("\n");

    const applyReplacements = (patterns) => {
      for (const line of rawDataArra) {
        let changed = false;
        let updatedLine = line;
  
        for (const [search, replace] of patterns) {
          if (line.includes(search)) {
            updatedLine = updatedLine.replaceAll(search, replace);
            madeChange = true;
            changed = true;
          }
        }
  
        fileOut.push(changed ? updatedLine : line);
      }
      rawDataArra = fileOut.slice();
      fileOut = [];
    };

    if (file.includes("go.mod") || file.includes("go.sum")) {
      fs.unlink(file)
      continue
    }

    if (file.includes("api_default.go")) {
      let updatedData = rawdata.replace(
        /localVarPath\s*:=\s*localBasePath\s*\+\s*"\/\{path\}"\s*\n\s*localVarPath\s*=\s*strings\.Replace\(localVarPath,\s*"\{"\s*\+\s*"path"\s*\+\s*"\}",\s*url\.PathEscape\(parameterValueToString\(r\.path,\s*"path"\)\),\s*-1\)/g,
        `// Added this code segment to allow for paths to be replaced fully
    localVarPath := strings.TrimSuffix(localBasePath, "/{path}")
    
    if r.path != "" {
        // Ensure the path starts with a forward slash
        if !strings.HasPrefix(r.path, "/") {
            localVarPath += "/"
        }
        localVarPath += r.path
    }
    
    segments := strings.Split(localVarPath, "/")
    for i, segment := range segments {
        segments[i] = url.PathEscape(segment)
    }
    localVarPath = strings.Join(segments, "/")
    // Done adding path replacement`
      );

      if (rawdata !== updatedData) {
        fs.writeFile(file, updatedData, 'utf8');
        console.log(`Updated ${file}`);
      } else {
        console.log(`No changes needed in ${file}`);
      }
      continue;
    }

    // change the poor naming for variables

    if (rawdata.includes("map[string]interface{}var") || rawdata.includes("ArrayOf*string") || file.includes("model_json_patch_operation_value.go")) {
      for (const line of rawDataArra) {
        if (line.includes("ginterface{}")) {
          fileOut.push(line.replaceAll("ginterface{}", "ginterface"));
          madeChange = true;
        } else if (line.includes("map[string]interface{}var")) {
          fileOut.push(line.replaceAll("map[string]interface{}var", "mapvar"));
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

    if (rawdata.includes("ArrayOfMapmapOfStringinterface{}") || rawdata.includes("MapmapOfStringinterface{}") || file.includes("model_workflow_library_action_example_output.go")) {
      for (const line of rawDataArra) {
        if (line.includes("ArrayOfMapmapOfStringinterface{}")) {
          fileOut.push(line.replaceAll("ArrayOfMapmapOfStringinterface{}", "ArrayOfMapmapOfStringinterface"));
          madeChange = true;
        } else if (line.includes("MapmapOfStringinterface{}")) {
          fileOut.push(line.replaceAll("MapmapOfStringinterface{}", "MapmapOfStringinterface"));
          madeChange = true;
        } else {
          fileOut.push(line);
        }
      }
      rawDataArra = fileOut.slice();
      fileOut = [];
    }



    //fix duplicate type names in beta api
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
        if (line.includes("[]SubscriptionPatchRequestInnerValueAnyOfInnervar")) {
          fileOut.push(line.replaceAll("[]SubscriptionPatchRequestInnerValueAnyOfInnervar", "SubscriptionPatchRequestInnerValueAnyOfInnervar"));
          madeChange = true;
        } else {
          fileOut.push(line);
        }
      }
      rawDataArra = fileOut.slice();
      fileOut = [];
    }


    //adjust the document type naming to fix the duplicate type errors
    if (file.includes("model_event_document.go") || file.includes("model_event.go")) {
      for (const line of rawDataArra) {
        if (line.includes("Type DocumentType")) {
          fileOut.push(line.replace("Type DocumentType", "DocumentType DocumentType"));
          madeChange = true;
        }
        else if (line.includes("this.Type = type_")) {
          fileOut.push(line.replace("this.Type = type_", "this.DocumentType = type_"));
          madeChange = true;
        }
        else if (line.includes("func (o *EventDocument) GetType() DocumentType {")) {
          fileOut.push(line.replace("func (o *EventDocument) GetType() DocumentType {", "func (o *EventDocument) GetDocumentType() DocumentType {"));
          madeChange = true;
        }
        else if (line.includes("func (o *Event) GetType() DocumentType {")) {
          fileOut.push(line.replace("func (o *Event) GetType() DocumentType {", "func (o *Event) GetDocumentType() DocumentType {"));
          madeChange = true;
        }
        else if (line.includes('return o.Type') && !line.includes('return o.Type, true')) {
          fileOut.push(line.replace("return o.Type", "return o.DocumentType"));
          madeChange = true;
        }

        else if (line.includes("func (o *EventDocument) GetTypeOk() (*DocumentType, bool) {")) {
          fileOut.push(line.replace("func (o *EventDocument) GetTypeOk() (*DocumentType, bool) {", "func (o *EventDocument) GetDocumentTypeOk() (*DocumentType, bool) {"));
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

        else if (line.includes("func (o *EventDocument) SetType(v DocumentType) {")) {
          fileOut.push(line.replace("func (o *EventDocument) SetType(v DocumentType) {", "func (o *EventDocument) SetDocumentType(v DocumentType) {"));
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

    if (file.includes(".md")) {
      applyReplacements([
        ["(../models/os-file)", "(https://pkg.go.dev/os)"],
        ["(list-error-message-dto1)", "(error-message-dto)"],
        ["(list-error-message-dto)", "(error-message-dto)"],
        ["(list)", "(https://docs.python.org/3/tutorial/datastructures.html#more-on-lists)"]
      ]);
    }

    if (file.includes("model_") || file.includes(".md")) {
      console.log(`Processing file: ${file}`);
      for (const line of rawDataArra) {
        if (line.includes("\"time\"")) {
          console.log(`Found "time" in line: ${line}`);
          fileOut.push(line.replace("\"time\"", ""));
          madeChange = true;
        } else if (line.includes('time.Time')) {
          fileOut.push(line.replace(/time\.Time/g, 'SailPointTime'));
          madeChange = true;
        } else {
          fileOut.push(line);
        }
      }

      rawDataArra = fileOut.slice();
      fileOut = [];
    }

    if (madeChange) {
      fixCheck += 1;
      fs.writeFile(file, rawDataArra.join("\n"));
    }
  }
  console.log(`fixed ${fixCheck} files`)
}

const main = async () => {
  let myArray = [];

  await processDirectory(path.join(process.argv[2], '/docs'));
  await renameFileToIndices(path.join(process.argv[2], '/docs/Models/Index.md'));
  await getAllFiles(process.argv[2], myArray);
  await fixFiles(myArray);
  await moveFiles(process.argv[2], path.join(process.argv[2], '/docs/Models'), "Index.md");
  await mergeCodeExampleFiles(path.join(process.argv[2], 'docs/Examples'));
};

main().catch((error) => {
  console.error(`Error in script: ${error.message}`);
});

