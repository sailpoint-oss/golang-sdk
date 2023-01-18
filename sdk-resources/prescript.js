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
  
  
  
  
  
    if (madeChange) {
      fs.writeFileSync(file, rawDataArra.join("\n"));
    }
  }
}




let myArray = [];
getAllFiles(process.argv[2], myArray);

fixFiles(myArray)

