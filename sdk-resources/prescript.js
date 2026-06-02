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
  let fixCheck = 0;
  for (const file of myArray) {
    let fileOut = [];
    let madeChange = false;
    let rawdata = fs.readFileSync(file).toString();
    let rawDataArra = rawdata.split("\n");

  
    // remove the complex transform schema (year-versioned and apis/ paths)
    if (file.includes(path.join("schemas","Transform.yaml")) ||
        file.includes(path.join("transforms","schemas","transform.yaml"))) {
      for (let line of rawDataArra) {
        if (line.includes('oneOf')) {
          line = line.replaceAll("oneOf:", "type: object")
          madeChange = true;
        }
        if (line.includes('- $ref:')) {
          
        } else {
          fileOut.push(line);
        }
        
      }
      rawDataArra = fileOut.slice();
      fileOut = [];
    }

    // remove the complex workflow trigger schema (year-versioned and apis/ paths)
    if (file.includes(path.join("schemas","WorkflowTrigger.yaml")) ||
        file.includes(path.join("schemas","workflows","WorkflowTrigger.yaml")) ||
        file.includes(path.join("workflows","schemas","workflowtrigger.yaml"))) {
      for (let line of rawDataArra) {
        if (line.includes('anyOf')) {
          line = line.replaceAll("anyOf:", "type: object")
          madeChange = true;
        }
        if (line.includes('- $ref:')) {
          
        } else {
          fileOut.push(line);
        }
        
      }
      rawDataArra = fileOut.slice();
      fileOut = [];
    }

    // remove the complex account schema (year-versioned paths only — not needed in apis/ structure)
    if (file.includes(path.join("paths","accounts.yaml"))) {
      
      for (let line of rawDataArra) {
        if (line.includes('oneOf')) {
          line = line.replaceAll("oneOf:", "$ref: '../schemas/SlimAccount.yaml'")
          madeChange = true;
        }
        if (line.includes("- $ref: '../schemas/SlimAccount.yaml'") || line.includes("- $ref: '../schemas/FullAccount.yaml'")) {
          
        } else {
          fileOut.push(line);
        }
        
      }
      rawDataArra = fileOut.slice();
      fileOut = [];
    }

    // remove the complex search schema (year-versioned and apis/ paths)
    if (file.includes(path.join("documents", "SearchDocument.yaml")) ||
        file.includes(path.join("search","schemas","searchdocument.yaml"))) {
          fileOut.push("type: object");
          rawDataArra = fileOut.slice();
          fileOut = [];
          madeChange = true;
    }

    // remove the complex search schema (year-versioned and apis/ paths)
    if (file.includes(path.join("documents", "SearchDocuments.yaml")) ||
        file.includes(path.join("search","schemas","searchdocuments.yaml"))) {
      fileOut.push("type: object");
      rawDataArra = fileOut.slice();
      fileOut = [];
      madeChange = true;
    }
  
    if (madeChange) {
      fixCheck += 1;
      fs.writeFileSync(file, rawDataArra.join("\n"));
    }
  }
  console.log(`fixed ${fixCheck} files`)
}




let myArray = [];
getAllFiles(process.argv[2], myArray);

fixFiles(myArray)

