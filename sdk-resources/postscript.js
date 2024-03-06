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

    // fix syntax error in model_patch
    if (file.includes("AccessProfilesAPI.md")) {
      for (const line of rawDataArra) {
        if (line.includes("//CreateAccessProfile")) {
          fileOut.push(
            line.replaceAll(
              "//CreateAccessProfile",
              ' id := "123" \n\t ownerName := "John Doe" \n\t ownerType := "Identity" \n\t sourceName := "Employee Database" \n\t sourceType := "SOURCE" \n accessProfile := v3.AccessProfile{ \n Name: "Employee-database-read-write", \n Owner: v3.OwnerReference{  \n Id:   &id, \n Name: &ownerName, \n Type: &ownerType, \n}, \n Source: v3.AccessProfileSourceRef{ \n Id:   &id, \n Name: &sourceName, \n Type: &sourceType, \n},}'
            )
          );
          madeChange = true;
        } else if (
          line.includes("(AccessProfilesAPI.md#CreateAccessProfile)")
        ) {
          fileOut.push(
            line.replace(
              "(AccessProfilesAPI.md#CreateAccessProfile)",
              "(AccessProfilesAPI.md#create-an-access-profile)"
            )
          );
          madeChange = true;
        } else if (
          line.includes("(AccessProfilesAPI.md#DeleteAccessProfile)")
        ) {
          fileOut.push(
            line.replace(
              "(AccessProfilesAPI.md#DeleteAccessProfile)",
              "(AccessProfilesAPI.md#delete-the-specified-access-profile)"
            )
          );
          madeChange = true;
        } else if (
          line.includes("(AccessProfilesAPI.md#DeleteAccessProfilesInBulk)")
        ) {
          fileOut.push(
            line.replace(
              "(AccessProfilesAPI.md#DeleteAccessProfilesInBulk)",
              "(AccessProfilesAPI.md#delete-access-profiles)"
            )
          );
          madeChange = true;
        } else if (line.includes("(AccessProfilesAPI.md#GetAccessProfile)")) {
          fileOut.push(
            line.replace(
              "(AccessProfilesAPI.md#GetAccessProfile)",
              "(AccessProfilesAPI.md#get-an-access-profile)"
            )
          );
          madeChange = true;
        } else if (
          line.includes("(AccessProfilesAPI.md#GetAccessProfileEntitlements)")
        ) {
          fileOut.push(
            line.replace(
              "(AccessProfilesAPI.md#GetAccessProfileEntitlements)",
              "(AccessProfilesAPI.md#list-access-profiles-entitlements)"
            )
          );
          madeChange = true;
        } else if (line.includes("(AccessProfilesAPI.md#ListAccessProfiles)")) {
          fileOut.push(
            line.replace(
              "(AccessProfilesAPI.md#ListAccessProfiles)",
              "(AccessProfilesAPI.md#list-access-profiles)"
            )
          );
          madeChange = true;
        } else if (line.includes("List Access Profiles")) {
          fileOut.push(
            line.replace("List Access Profiles", "list-access-profiles")
          );
          madeChange = true;
        } else if (line.includes("(AccessProfilesAPI.md#PatchAccessProfile)")) {
          fileOut.push(
            line.replace(
              "(AccessProfilesAPI.md#PatchAccessProfile)",
              "(AccessProfilesAPI.md#patch-a-specified-access-profile)"
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

    if (madeChange) {
      fixCheck += 1;
      fs.writeFileSync(file, rawDataArra.join("\n"));
    }
  }

  console.log(`fixed ${fixCheck} files`);
};

let myArray = [];
getAllFiles(process.argv[2], myArray);

fixFiles(myArray);
