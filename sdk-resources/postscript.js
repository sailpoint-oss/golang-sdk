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
  
    if (file.includes("go.mod") || file.includes("go.sum")) {
      fs.unlinkSync(file)
      continue
    }
    

    let fileOut = [];
    let madeChange = false;
    let rawdata = fs.readFileSync(file).toString();
    let rawDataArra = rawdata.split("\n");

  
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

    if (file.includes("model_workgroup_dto.go")) {

    const sailPointTimeCode = `type SailPointTime struct {
    time.Time
}
    
func (m *SailPointTime) UnmarshalJSON(data []byte) error {
        if string(data) == "null" || string(data) == \`""\` {
            return nil;
        }
    
        // Strip the quotes from the data
        str := string(data);
        str = str[1 : len(str)-1];
    
        // Try parsing with seconds first
        formats := []string{
            time.RFC3339,                   // Format with seconds (e.g., "2012-04-23T18:25:43Z")
            "2006-01-02T15:04Z",            // Format without seconds (e.g., "2012-04-23T18:25Z")
            "2006-01-02T15:04:05.999999999", // Optional milliseconds and nanoseconds
        }
    
        var err error;
        for _, format := range formats {
            var t time.Time;
            t, err = time.Parse(format, str);
            if err == nil {
                *m = SailPointTime{t};
                return nil;
            }
        }
    
        return nil;
}

func (o *WorkgroupDto) UnmarshalJSON(data []byte) (err error) {`;
    

      for (const line of rawDataArra) {
        if (line.includes('Created *time.Time `json:"created,omitempty"`')) {
          fileOut.push(line.replace('Created *time.Time `json:"created,omitempty"`', 'Created *SailPointTime `json:"created,omitempty"`'));
          madeChange = true;
        } else if (line.includes('Modified *time.Time `json:"modified,omitempty"`')) {
          fileOut.push(line.replace('Modified *time.Time `json:"modified,omitempty"`', 'Modified *SailPointTime  `json:"modified,omitempty"`'));
          madeChange = true;
        } else if (line.includes('func (o *WorkgroupDto) GetCreated() time.Time {')) {
          fileOut.push(line.replace('func (o *WorkgroupDto) GetCreated() time.Time {', 'func (o *WorkgroupDto) GetCreated() SailPointTime {'));
          madeChange = true;
        } else if (line.includes('func (o *WorkgroupDto) GetCreatedOk() (*time.Time, bool) {')) {
          fileOut.push(line.replace('func (o *WorkgroupDto) GetCreatedOk() (*time.Time, bool) {', 'func (o *WorkgroupDto) GetCreatedOk() (*SailPointTime, bool) {'));
          madeChange = true;
        } else if (line.includes('func (o *WorkgroupDto) GetModified() time.Time {')) {
          fileOut.push(line.replace('func (o *WorkgroupDto) GetModified() time.Time {', 'func (o *WorkgroupDto) GetModified() SailPointTime {'));
          madeChange = true;
        } else if (line.includes('var ret time.Time')) {
          fileOut.push(line.replace('var ret time.Time', 'var ret SailPointTime'));
          madeChange = true;
        } else if (line.includes('func (o *WorkgroupDto) GetModifiedOk() (*time.Time, bool) {')) {
          fileOut.push(line.replace('func (o *WorkgroupDto) GetModifiedOk() (*time.Time, bool) {', 'func (o *WorkgroupDto) GetModifiedOk() (*SailPointTime, bool) {'));
          madeChange = true;
        } else if (line.includes('func (o *WorkgroupDto) SetModified(v time.Time) {')) {
          fileOut.push(line.replace('func (o *WorkgroupDto) SetModified(v time.Time) {', 'func (o *WorkgroupDto) SetModified(v SailPointTime) {'));
          madeChange = true;
        } else if (line.includes('func (o *WorkgroupDto) SetCreated(v time.Time) {')) {
          fileOut.push(line.replace('func (o *WorkgroupDto) SetCreated(v time.Time) {', 'func (o *WorkgroupDto) SetCreated(v SailPointTime) {'));
          madeChange = true;
        } else if (line.includes('func (o *WorkgroupDto) UnmarshalJSON(data []byte) (err error) {')) {
          fileOut.push(line.replace('func (o *WorkgroupDto) UnmarshalJSON(data []byte) (err error) {', sailPointTimeCode));
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
      fs.writeFileSync(file, rawDataArra.join("\n"));
    }
  }
  console.log(`fixed ${fixCheck} files`)
}




let myArray = [];
getAllFiles(process.argv[2], myArray);

fixFiles(myArray)

