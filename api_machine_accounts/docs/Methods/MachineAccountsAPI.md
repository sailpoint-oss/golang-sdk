---
id: v1-machine-accounts
title: MachineAccounts
pagination_label: MachineAccounts
sidebar_label: MachineAccounts
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'MachineAccounts', 'V1MachineAccounts'] 
slug: /tools/sdk/go/machineaccounts/methods/machine-accounts
tags: ['SDK', 'Software Development Kit', 'MachineAccounts', 'V1MachineAccounts']
---

# MachineAccountsAPI
   
All URIs are relative to *https://sailpoint.api.identitynow.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create-machine-account-subtype-v1**](#create-machine-account-subtype-v1) | **Post** `/sources/v1/{sourceId}/subtypes` | Create subtype
[**delete-machine-account-subtype-by-technical-name-v1**](#delete-machine-account-subtype-by-technical-name-v1) | **Delete** `/sources/v1/{sourceId}/subtypes/{technicalName}` | Delete subtype
[**get-machine-account-subtype-approval-config-v1**](#get-machine-account-subtype-approval-config-v1) | **Get** `/source-subtypes/v1/{subtypeId}/machine-config` | Machine Subtype Approval Config
[**get-machine-account-subtype-by-id-v1**](#get-machine-account-subtype-by-id-v1) | **Get** `/sources/v1/subtypes/{subtypeId}` | Retrieve subtype by subtype id
[**get-machine-account-subtype-by-technical-name-v1**](#get-machine-account-subtype-by-technical-name-v1) | **Get** `/sources/v1/{sourceId}/subtypes/{technicalName}` | Retrieve subtype by source and technicalName
[**get-machine-account-v1**](#get-machine-account-v1) | **Get** `/machine-accounts/v1/{id}` | Get machine account details
[**list-machine-account-subtypes-v1**](#list-machine-account-subtypes-v1) | **Get** `/sources/v1/{sourceId}/subtypes` | Retrieve all subtypes by source
[**list-machine-accounts-v1**](#list-machine-accounts-v1) | **Get** `/machine-accounts/v1` | List machine accounts
[**load-bulk-source-subtypes-v1**](#load-bulk-source-subtypes-v1) | **Post** `/source-subtypes/v1/bulk-retrieve` | Bulk Retrieve of Source Subtypes
[**patch-machine-account-subtype-by-technical-name-v1**](#patch-machine-account-subtype-by-technical-name-v1) | **Patch** `/sources/v1/{sourceId}/subtypes/{technicalName}` | Patch subtype
[**update-machine-account-subtype-approval-config-v1**](#update-machine-account-subtype-approval-config-v1) | **Patch** `/source-subtypes/v1/{subtypeId}/machine-config` | Machine Subtype Approval Config
[**update-machine-account-v1**](#update-machine-account-v1) | **Patch** `/machine-accounts/v1/{id}` | Update machine account details


## create-machine-account-subtype-v1
:::warning experimental 
This API is currently in an experimental state. The API is subject to change based on feedback and further testing. You must include the X-SailPoint-Experimental header and set it to `true` to use this endpoint.
:::
:::tip setting x-sailpoint-experimental header
 on the configuration object you can set the `x-sailpoint-experimental` header to `true' to enable all experimantl endpoints within the SDK.
 Example:
 ```go
   configuration = Configuration()
   configuration.Experimental = true
 ```
:::
Create subtype
Create a new machine account subtype for a source.

[API Spec](https://developer.sailpoint.com/docs/api/v1/create-machine-account-subtype-v1)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** | The ID of the source. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMachineAccountSubtypeV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **createMachineAccountSubtypeV1Request** | [**CreateMachineAccountSubtypeV1Request**](../models/create-machine-account-subtype-v1-request) |  | 

### Return type

[**Sourcesubtype**](../models/sourcesubtype)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
  "encoding/json"
    v1 "github.com/sailpoint-oss/golang-sdk/v3/api_machine_accounts"
	sailpoint "github.com/sailpoint-oss/golang-sdk/v3/api_machine_accounts"
)

func main() {
    sourceId := `6d0458373bec4b4b80460992b76016da` // string | The ID of the source. # string | The ID of the source.
    xSailPointExperimental := `true` // string | Use this header to enable this experimental API. (default to "true") # string | Use this header to enable this experimental API. (default to "true")
    createmachineaccountsubtypev1request := []byte(`{"technicalName":"foo","displayName":"Mr Foo","description":"fighters","type":"MACHINE"}`) // CreateMachineAccountSubtypeV1Request | 

    var createMachineAccountSubtypeV1Request v1.CreateMachineAccountSubtypeV1Request
    if err := json.Unmarshal(createmachineaccountsubtypev1request, &createMachineAccountSubtypeV1Request); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineAccounts.MachineAccountsAPI.CreateMachineAccountSubtypeV1(context.Background(), sourceId).XSailPointExperimental(xSailPointExperimental).CreateMachineAccountSubtypeV1Request(createMachineAccountSubtypeV1Request).Execute()
	  //resp, r, err := apiClient.MachineAccounts.MachineAccountsAPI.CreateMachineAccountSubtypeV1(context.Background(), sourceId).XSailPointExperimental(xSailPointExperimental).CreateMachineAccountSubtypeV1Request(createMachineAccountSubtypeV1Request).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `MachineAccountsAPI.CreateMachineAccountSubtypeV1``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMachineAccountSubtypeV1`: Sourcesubtype
    fmt.Fprintf(os.Stdout, "Response from `MachineAccountsAPI.CreateMachineAccountSubtypeV1`: %v\n", resp)
}
```

[[Back to top]](#)

## delete-machine-account-subtype-by-technical-name-v1
:::warning experimental 
This API is currently in an experimental state. The API is subject to change based on feedback and further testing. You must include the X-SailPoint-Experimental header and set it to `true` to use this endpoint.
:::
:::tip setting x-sailpoint-experimental header
 on the configuration object you can set the `x-sailpoint-experimental` header to `true' to enable all experimantl endpoints within the SDK.
 Example:
 ```go
   configuration = Configuration()
   configuration.Experimental = true
 ```
:::
Delete subtype
Delete a machine account subtype by source ID and technical name.

[API Spec](https://developer.sailpoint.com/docs/api/v1/delete-machine-account-subtype-by-technical-name-v1)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** | The ID of the source. | 
**technicalName** | **string** | The technical name of the subtype. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMachineAccountSubtypeByTechnicalNameV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
  
    
	sailpoint "github.com/sailpoint-oss/golang-sdk/v3/api_machine_accounts"
)

func main() {
    sourceId := `6d0458373bec4b4b80460992b76016da` // string | The ID of the source. # string | The ID of the source.
    technicalName := `foo` // string | The technical name of the subtype. # string | The technical name of the subtype.
    xSailPointExperimental := `true` // string | Use this header to enable this experimental API. (default to "true") # string | Use this header to enable this experimental API. (default to "true")

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    r, err := apiClient.MachineAccounts.MachineAccountsAPI.DeleteMachineAccountSubtypeByTechnicalNameV1(context.Background(), sourceId, technicalName).XSailPointExperimental(xSailPointExperimental).Execute()
	  //r, err := apiClient.MachineAccounts.MachineAccountsAPI.DeleteMachineAccountSubtypeByTechnicalNameV1(context.Background(), sourceId, technicalName).XSailPointExperimental(xSailPointExperimental).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `MachineAccountsAPI.DeleteMachineAccountSubtypeByTechnicalNameV1``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    
}
```

[[Back to top]](#)

## get-machine-account-subtype-approval-config-v1
:::warning experimental 
This API is currently in an experimental state. The API is subject to change based on feedback and further testing. You must include the X-SailPoint-Experimental header and set it to `true` to use this endpoint.
:::
:::tip setting x-sailpoint-experimental header
 on the configuration object you can set the `x-sailpoint-experimental` header to `true' to enable all experimantl endpoints within the SDK.
 Example:
 ```go
   configuration = Configuration()
   configuration.Experimental = true
 ```
:::
Machine Subtype Approval Config
This endpoint retrieves the approval configuration for machine account creation and deletion at the machine subtype level. By providing a specific subtypeId in the path, clients can fetch the approval rules and settings (such as required approvers and comments policy) that govern account creation and deletion for that particular machine subtype. The response includes a MachineAccountSubtypeConfigDto object detailing these configurations, enabling clients to understand or display the approval workflow required for creating and deleting machine accounts of the given subtype. Use this endpoint to get machine subtype level approval config for account creation and deletion.

[API Spec](https://developer.sailpoint.com/docs/api/v1/get-machine-account-subtype-approval-config-v1)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subtypeId** | **string** | machine subtype id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMachineAccountSubtypeApprovalConfigV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]


### Return type

[**Machineaccountsubtypeconfigdto**](../models/machineaccountsubtypeconfigdto)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
  
    
	sailpoint "github.com/sailpoint-oss/golang-sdk/v3/api_machine_accounts"
)

func main() {
    xSailPointExperimental := `true` // string | Use this header to enable this experimental API. (default to "true") # string | Use this header to enable this experimental API. (default to "true")
    subtypeId := `ef38f94347e94562b5bb8424a56498d8` // string | machine subtype id. # string | machine subtype id.

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineAccounts.MachineAccountsAPI.GetMachineAccountSubtypeApprovalConfigV1(context.Background(), subtypeId).XSailPointExperimental(xSailPointExperimental).Execute()
	  //resp, r, err := apiClient.MachineAccounts.MachineAccountsAPI.GetMachineAccountSubtypeApprovalConfigV1(context.Background(), subtypeId).XSailPointExperimental(xSailPointExperimental).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `MachineAccountsAPI.GetMachineAccountSubtypeApprovalConfigV1``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMachineAccountSubtypeApprovalConfigV1`: Machineaccountsubtypeconfigdto
    fmt.Fprintf(os.Stdout, "Response from `MachineAccountsAPI.GetMachineAccountSubtypeApprovalConfigV1`: %v\n", resp)
}
```

[[Back to top]](#)

## get-machine-account-subtype-by-id-v1
:::warning experimental 
This API is currently in an experimental state. The API is subject to change based on feedback and further testing. You must include the X-SailPoint-Experimental header and set it to `true` to use this endpoint.
:::
:::tip setting x-sailpoint-experimental header
 on the configuration object you can set the `x-sailpoint-experimental` header to `true' to enable all experimantl endpoints within the SDK.
 Example:
 ```go
   configuration = Configuration()
   configuration.Experimental = true
 ```
:::
Retrieve subtype by subtype id
Get a machine account subtype by its unique ID.

[API Spec](https://developer.sailpoint.com/docs/api/v1/get-machine-account-subtype-by-id-v1)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subtypeId** | **string** | The ID of the machine account subtype. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMachineAccountSubtypeByIdV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]

### Return type

[**Sourcesubtype**](../models/sourcesubtype)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
  
    
	sailpoint "github.com/sailpoint-oss/golang-sdk/v3/api_machine_accounts"
)

func main() {
    subtypeId := `43bdd144-4b17-4fce-a744-17c7fd3e717b` // string | The ID of the machine account subtype. # string | The ID of the machine account subtype.
    xSailPointExperimental := `true` // string | Use this header to enable this experimental API. (default to "true") # string | Use this header to enable this experimental API. (default to "true")

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineAccounts.MachineAccountsAPI.GetMachineAccountSubtypeByIdV1(context.Background(), subtypeId).XSailPointExperimental(xSailPointExperimental).Execute()
	  //resp, r, err := apiClient.MachineAccounts.MachineAccountsAPI.GetMachineAccountSubtypeByIdV1(context.Background(), subtypeId).XSailPointExperimental(xSailPointExperimental).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `MachineAccountsAPI.GetMachineAccountSubtypeByIdV1``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMachineAccountSubtypeByIdV1`: Sourcesubtype
    fmt.Fprintf(os.Stdout, "Response from `MachineAccountsAPI.GetMachineAccountSubtypeByIdV1`: %v\n", resp)
}
```

[[Back to top]](#)

## get-machine-account-subtype-by-technical-name-v1
:::warning experimental 
This API is currently in an experimental state. The API is subject to change based on feedback and further testing. You must include the X-SailPoint-Experimental header and set it to `true` to use this endpoint.
:::
:::tip setting x-sailpoint-experimental header
 on the configuration object you can set the `x-sailpoint-experimental` header to `true' to enable all experimantl endpoints within the SDK.
 Example:
 ```go
   configuration = Configuration()
   configuration.Experimental = true
 ```
:::
Retrieve subtype by source and technicalName
Get a machine account subtype by source ID and technical name.

[API Spec](https://developer.sailpoint.com/docs/api/v1/get-machine-account-subtype-by-technical-name-v1)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** | The ID of the source. | 
**technicalName** | **string** | The technical name of the subtype. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMachineAccountSubtypeByTechnicalNameV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]

### Return type

[**Sourcesubtype**](../models/sourcesubtype)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
  
    
	sailpoint "github.com/sailpoint-oss/golang-sdk/v3/api_machine_accounts"
)

func main() {
    sourceId := `6d0458373bec4b4b80460992b76016da` // string | The ID of the source. # string | The ID of the source.
    technicalName := `foo` // string | The technical name of the subtype. # string | The technical name of the subtype.
    xSailPointExperimental := `true` // string | Use this header to enable this experimental API. (default to "true") # string | Use this header to enable this experimental API. (default to "true")

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineAccounts.MachineAccountsAPI.GetMachineAccountSubtypeByTechnicalNameV1(context.Background(), sourceId, technicalName).XSailPointExperimental(xSailPointExperimental).Execute()
	  //resp, r, err := apiClient.MachineAccounts.MachineAccountsAPI.GetMachineAccountSubtypeByTechnicalNameV1(context.Background(), sourceId, technicalName).XSailPointExperimental(xSailPointExperimental).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `MachineAccountsAPI.GetMachineAccountSubtypeByTechnicalNameV1``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMachineAccountSubtypeByTechnicalNameV1`: Sourcesubtype
    fmt.Fprintf(os.Stdout, "Response from `MachineAccountsAPI.GetMachineAccountSubtypeByTechnicalNameV1`: %v\n", resp)
}
```

[[Back to top]](#)

## get-machine-account-v1
:::warning experimental 
This API is currently in an experimental state. The API is subject to change based on feedback and further testing. You must include the X-SailPoint-Experimental header and set it to `true` to use this endpoint.
:::
:::tip setting x-sailpoint-experimental header
 on the configuration object you can set the `x-sailpoint-experimental` header to `true' to enable all experimantl endpoints within the SDK.
 Example:
 ```go
   configuration = Configuration()
   configuration.Experimental = true
 ```
:::
Get machine account details
Use this API to return the details for a single machine account by its ID.  

[API Spec](https://developer.sailpoint.com/docs/api/v1/get-machine-account-v1)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Machine Account ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMachineAccountV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]

### Return type

[**Machineaccount**](../models/machineaccount)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
  
    
	sailpoint "github.com/sailpoint-oss/golang-sdk/v3/api_machine_accounts"
)

func main() {
    id := `ef38f94347e94562b5bb8424a56397d8` // string | Machine Account ID. # string | Machine Account ID.
    xSailPointExperimental := `true` // string | Use this header to enable this experimental API. (default to "true") # string | Use this header to enable this experimental API. (default to "true")

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineAccounts.MachineAccountsAPI.GetMachineAccountV1(context.Background(), id).XSailPointExperimental(xSailPointExperimental).Execute()
	  //resp, r, err := apiClient.MachineAccounts.MachineAccountsAPI.GetMachineAccountV1(context.Background(), id).XSailPointExperimental(xSailPointExperimental).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `MachineAccountsAPI.GetMachineAccountV1``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMachineAccountV1`: Machineaccount
    fmt.Fprintf(os.Stdout, "Response from `MachineAccountsAPI.GetMachineAccountV1`: %v\n", resp)
}
```

[[Back to top]](#)

## list-machine-account-subtypes-v1
:::warning experimental 
This API is currently in an experimental state. The API is subject to change based on feedback and further testing. You must include the X-SailPoint-Experimental header and set it to `true` to use this endpoint.
:::
:::tip setting x-sailpoint-experimental header
 on the configuration object you can set the `x-sailpoint-experimental` header to `true' to enable all experimantl endpoints within the SDK.
 Example:
 ```go
   configuration = Configuration()
   configuration.Experimental = true
 ```
:::
Retrieve all subtypes by source
Get all machine account subtypes for a given source.

[API Spec](https://developer.sailpoint.com/docs/api/v1/list-machine-account-subtypes-v1)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** | The ID of the source. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMachineAccountSubtypesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **displayName**: *eq, sw*  **technicalName**: *eq, sw* | 
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **displayName, technicalName** | 
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]

### Return type

[**[]Sourcesubtype**](../models/sourcesubtype)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
  
    
	sailpoint "github.com/sailpoint-oss/golang-sdk/v3/api_machine_accounts"
)

func main() {
    sourceId := `6d0458373bec4b4b80460992b76016da` // string | The ID of the source. # string | The ID of the source.
    xSailPointExperimental := `true` // string | Use this header to enable this experimental API. (default to "true") # string | Use this header to enable this experimental API. (default to "true")
    filters := `displayName eq "sail"` // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **displayName**: *eq, sw*  **technicalName**: *eq, sw* (optional) # string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **displayName**: *eq, sw*  **technicalName**: *eq, sw* (optional)
    sorters := `displayName` // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **displayName, technicalName** (optional) # string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **displayName, technicalName** (optional)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false) # bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
    limit := 250 // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250) # int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    offset := 0 // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0) # int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineAccounts.MachineAccountsAPI.ListMachineAccountSubtypesV1(context.Background(), sourceId).XSailPointExperimental(xSailPointExperimental).Execute()
	  //resp, r, err := apiClient.MachineAccounts.MachineAccountsAPI.ListMachineAccountSubtypesV1(context.Background(), sourceId).XSailPointExperimental(xSailPointExperimental).Filters(filters).Sorters(sorters).Count(count).Limit(limit).Offset(offset).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `MachineAccountsAPI.ListMachineAccountSubtypesV1``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMachineAccountSubtypesV1`: []Sourcesubtype
    fmt.Fprintf(os.Stdout, "Response from `MachineAccountsAPI.ListMachineAccountSubtypesV1`: %v\n", resp)
}
```

[[Back to top]](#)

## list-machine-accounts-v1
:::warning experimental 
This API is currently in an experimental state. The API is subject to change based on feedback and further testing. You must include the X-SailPoint-Experimental header and set it to `true` to use this endpoint.
:::
:::tip setting x-sailpoint-experimental header
 on the configuration object you can set the `x-sailpoint-experimental` header to `true' to enable all experimantl endpoints within the SDK.
 Example:
 ```go
   configuration = Configuration()
   configuration.Experimental = true
 ```
:::
List machine accounts
This returns a list of machine accounts.  

[API Spec](https://developer.sailpoint.com/docs/api/v1/list-machine-accounts-v1)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMachineAccountsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq, in*  **name**: *eq, in, sw*  **nativeIdentity**: *eq, in, sw*  **uuid**: *eq, in*  **description**: *eq, in, sw*  **machineIdentity.id**: *eq, in*  **machineIdentity.name**: *eq, in, sw*  **subtype.technicalName**: *eq, in, sw*  **subtype.displayName**: *eq, in, sw*  **accessType**: *eq, in, sw*  **environment**: *eq, in, sw*  **ownerIdentity**: *eq, in*  **ownerIdentity.id**: *eq, in*  **ownerIdentity.name**: *eq, in, sw*  **manuallyCorrelated**: *eq*  **enabled**: *eq*  **locked**: *eq*  **hasEntitlements**: *eq*  **attributes**: *eq*  **source.id**: *eq, in*  **source.name**: *eq, in, sw*  **created**: *eq, gt, lt, ge, le*  **modified**: *eq, gt, lt, ge, le* | 
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **id, name, nativeIdentity, ownerIdentity, uuid, description, machineIdentity.id, machineIdentity.name, subtype.technicalName, subtype.displayName, accessType, environment, manuallyCorrelated, enabled, locked, hasEntitlements, ownerIdentity.id, ownerIdentity.name, attributes, source.id, source.name, created, modified** | 

### Return type

[**[]Machineaccount**](../models/machineaccount)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
  
    
	sailpoint "github.com/sailpoint-oss/golang-sdk/v3/api_machine_accounts"
)

func main() {
    xSailPointExperimental := `true` // string | Use this header to enable this experimental API. (default to "true") # string | Use this header to enable this experimental API. (default to "true")
    limit := 250 // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250) # int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    offset := 0 // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0) # int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false) # bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
    filters := `hasEntitlements eq true` // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq, in*  **name**: *eq, in, sw*  **nativeIdentity**: *eq, in, sw*  **uuid**: *eq, in*  **description**: *eq, in, sw*  **machineIdentity.id**: *eq, in*  **machineIdentity.name**: *eq, in, sw*  **subtype.technicalName**: *eq, in, sw*  **subtype.displayName**: *eq, in, sw*  **accessType**: *eq, in, sw*  **environment**: *eq, in, sw*  **ownerIdentity**: *eq, in*  **ownerIdentity.id**: *eq, in*  **ownerIdentity.name**: *eq, in, sw*  **manuallyCorrelated**: *eq*  **enabled**: *eq*  **locked**: *eq*  **hasEntitlements**: *eq*  **attributes**: *eq*  **source.id**: *eq, in*  **source.name**: *eq, in, sw*  **created**: *eq, gt, lt, ge, le*  **modified**: *eq, gt, lt, ge, le* (optional) # string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq, in*  **name**: *eq, in, sw*  **nativeIdentity**: *eq, in, sw*  **uuid**: *eq, in*  **description**: *eq, in, sw*  **machineIdentity.id**: *eq, in*  **machineIdentity.name**: *eq, in, sw*  **subtype.technicalName**: *eq, in, sw*  **subtype.displayName**: *eq, in, sw*  **accessType**: *eq, in, sw*  **environment**: *eq, in, sw*  **ownerIdentity**: *eq, in*  **ownerIdentity.id**: *eq, in*  **ownerIdentity.name**: *eq, in, sw*  **manuallyCorrelated**: *eq*  **enabled**: *eq*  **locked**: *eq*  **hasEntitlements**: *eq*  **attributes**: *eq*  **source.id**: *eq, in*  **source.name**: *eq, in, sw*  **created**: *eq, gt, lt, ge, le*  **modified**: *eq, gt, lt, ge, le* (optional)
    sorters := `id,name` // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **id, name, nativeIdentity, ownerIdentity, uuid, description, machineIdentity.id, machineIdentity.name, subtype.technicalName, subtype.displayName, accessType, environment, manuallyCorrelated, enabled, locked, hasEntitlements, ownerIdentity.id, ownerIdentity.name, attributes, source.id, source.name, created, modified** (optional) # string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **id, name, nativeIdentity, ownerIdentity, uuid, description, machineIdentity.id, machineIdentity.name, subtype.technicalName, subtype.displayName, accessType, environment, manuallyCorrelated, enabled, locked, hasEntitlements, ownerIdentity.id, ownerIdentity.name, attributes, source.id, source.name, created, modified** (optional)

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineAccounts.MachineAccountsAPI.ListMachineAccountsV1(context.Background()).XSailPointExperimental(xSailPointExperimental).Execute()
	  //resp, r, err := apiClient.MachineAccounts.MachineAccountsAPI.ListMachineAccountsV1(context.Background()).XSailPointExperimental(xSailPointExperimental).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `MachineAccountsAPI.ListMachineAccountsV1``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMachineAccountsV1`: []Machineaccount
    fmt.Fprintf(os.Stdout, "Response from `MachineAccountsAPI.ListMachineAccountsV1`: %v\n", resp)
}
```

[[Back to top]](#)

## load-bulk-source-subtypes-v1
:::warning experimental 
This API is currently in an experimental state. The API is subject to change based on feedback and further testing. You must include the X-SailPoint-Experimental header and set it to `true` to use this endpoint.
:::
:::tip setting x-sailpoint-experimental header
 on the configuration object you can set the `x-sailpoint-experimental` header to `true' to enable all experimantl endpoints within the SDK.
 Example:
 ```go
   configuration = Configuration()
   configuration.Experimental = true
 ```
:::
Bulk Retrieve of Source Subtypes
This endpoint retrieves the subtypes for given subtypeIds.

[API Spec](https://developer.sailpoint.com/docs/api/v1/load-bulk-source-subtypes-v1)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLoadBulkSourceSubtypesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **requestBody** | **[]string** |  | 

### Return type

[**[]Sourcesubtypewithsource**](../models/sourcesubtypewithsource)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
  "encoding/json"
    v1 "github.com/sailpoint-oss/golang-sdk/v3/api_machine_accounts"
	sailpoint "github.com/sailpoint-oss/golang-sdk/v3/api_machine_accounts"
)

func main() {
    xSailPointExperimental := `true` // string | Use this header to enable this experimental API. (default to "true") # string | Use this header to enable this experimental API. (default to "true")
    requestbody := []byte(``) // []string | 

    var requestBody []v1.RequestBody
    if err := json.Unmarshal(requestbody, &requestBody); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineAccounts.MachineAccountsAPI.LoadBulkSourceSubtypesV1(context.Background()).XSailPointExperimental(xSailPointExperimental).RequestBody(requestBody).Execute()
	  //resp, r, err := apiClient.MachineAccounts.MachineAccountsAPI.LoadBulkSourceSubtypesV1(context.Background()).XSailPointExperimental(xSailPointExperimental).RequestBody(requestBody).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `MachineAccountsAPI.LoadBulkSourceSubtypesV1``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LoadBulkSourceSubtypesV1`: []Sourcesubtypewithsource
    fmt.Fprintf(os.Stdout, "Response from `MachineAccountsAPI.LoadBulkSourceSubtypesV1`: %v\n", resp)
}
```

[[Back to top]](#)

## patch-machine-account-subtype-by-technical-name-v1
:::warning experimental 
This API is currently in an experimental state. The API is subject to change based on feedback and further testing. You must include the X-SailPoint-Experimental header and set it to `true` to use this endpoint.
:::
:::tip setting x-sailpoint-experimental header
 on the configuration object you can set the `x-sailpoint-experimental` header to `true' to enable all experimantl endpoints within the SDK.
 Example:
 ```go
   configuration = Configuration()
   configuration.Experimental = true
 ```
:::
Patch subtype
Update fields of a machine account subtype by source ID and technical name.
Patchable fields include: `displayName`, `description`.

[API Spec](https://developer.sailpoint.com/docs/api/v1/patch-machine-account-subtype-by-technical-name-v1)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** | The ID of the source. | 
**technicalName** | **string** | The technical name of the subtype. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchMachineAccountSubtypeByTechnicalNameV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **requestBody** | **[]map[string]interface{}** | A JSON of updated values [JSON Patch](https://tools.ietf.org/html/rfc6902) standard. | 

### Return type

[**Sourcesubtype**](../models/sourcesubtype)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
  "encoding/json"
    v1 "github.com/sailpoint-oss/golang-sdk/v3/api_machine_accounts"
	sailpoint "github.com/sailpoint-oss/golang-sdk/v3/api_machine_accounts"
)

func main() {
    sourceId := `6d0458373bec4b4b80460992b76016da` // string | The ID of the source. # string | The ID of the source.
    technicalName := `foo` // string | The technical name of the subtype. # string | The technical name of the subtype.
    xSailPointExperimental := `true` // string | Use this header to enable this experimental API. (default to "true") # string | Use this header to enable this experimental API. (default to "true")
    requestbody := []byte(`[{"op":"replace","path":"/displayName","value":"Test New DisplayName"}]`) // []map[string]interface{} | A JSON of updated values [JSON Patch](https://tools.ietf.org/html/rfc6902) standard.

    var requestBody []v1.RequestBody
    if err := json.Unmarshal(requestbody, &requestBody); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineAccounts.MachineAccountsAPI.PatchMachineAccountSubtypeByTechnicalNameV1(context.Background(), sourceId, technicalName).XSailPointExperimental(xSailPointExperimental).RequestBody(requestBody).Execute()
	  //resp, r, err := apiClient.MachineAccounts.MachineAccountsAPI.PatchMachineAccountSubtypeByTechnicalNameV1(context.Background(), sourceId, technicalName).XSailPointExperimental(xSailPointExperimental).RequestBody(requestBody).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `MachineAccountsAPI.PatchMachineAccountSubtypeByTechnicalNameV1``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchMachineAccountSubtypeByTechnicalNameV1`: Sourcesubtype
    fmt.Fprintf(os.Stdout, "Response from `MachineAccountsAPI.PatchMachineAccountSubtypeByTechnicalNameV1`: %v\n", resp)
}
```

[[Back to top]](#)

## update-machine-account-subtype-approval-config-v1
:::warning experimental 
This API is currently in an experimental state. The API is subject to change based on feedback and further testing. You must include the X-SailPoint-Experimental header and set it to `true` to use this endpoint.
:::
:::tip setting x-sailpoint-experimental header
 on the configuration object you can set the `x-sailpoint-experimental` header to `true' to enable all experimantl endpoints within the SDK.
 Example:
 ```go
   configuration = Configuration()
   configuration.Experimental = true
 ```
:::
Machine Subtype Approval Config
Updates the approval configuration for machine account deletion at the specified machine subtype level. This endpoint allows clients to modify approval rules and settings (such as required approvers and comments policy) for account creation and deletion workflows associated with a given subtypeId. Use this to customize or enforce approval requirements for creating and deleting machine accounts of a particular subtype.

[API Spec](https://developer.sailpoint.com/docs/api/v1/update-machine-account-subtype-approval-config-v1)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subtypeId** | **string** | machine account subtype ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMachineAccountSubtypeApprovalConfigV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]

 **jsonpatchoperation** | [**[]Jsonpatchoperation**](../models/jsonpatchoperation) | The JSONPatch payload used to update the object. | 

### Return type

[**Machineaccountsubtypeconfigdto**](../models/machineaccountsubtypeconfigdto)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
  "encoding/json"
    v1 "github.com/sailpoint-oss/golang-sdk/v3/api_machine_accounts"
	sailpoint "github.com/sailpoint-oss/golang-sdk/v3/api_machine_accounts"
)

func main() {
    xSailPointExperimental := `true` // string | Use this header to enable this experimental API. (default to "true") # string | Use this header to enable this experimental API. (default to "true")
    subtypeId := `00eebcf881994e419d72e757fd30dc0e` // string | machine account subtype ID. # string | machine account subtype ID.
    jsonpatchoperation := []byte(``) // []Jsonpatchoperation | The JSONPatch payload used to update the object.

    var jsonpatchoperation []v1.Jsonpatchoperation
    if err := json.Unmarshal(jsonpatchoperation, &jsonpatchoperation); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineAccounts.MachineAccountsAPI.UpdateMachineAccountSubtypeApprovalConfigV1(context.Background(), subtypeId).XSailPointExperimental(xSailPointExperimental).Jsonpatchoperation(jsonpatchoperation).Execute()
	  //resp, r, err := apiClient.MachineAccounts.MachineAccountsAPI.UpdateMachineAccountSubtypeApprovalConfigV1(context.Background(), subtypeId).XSailPointExperimental(xSailPointExperimental).Jsonpatchoperation(jsonpatchoperation).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `MachineAccountsAPI.UpdateMachineAccountSubtypeApprovalConfigV1``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMachineAccountSubtypeApprovalConfigV1`: Machineaccountsubtypeconfigdto
    fmt.Fprintf(os.Stdout, "Response from `MachineAccountsAPI.UpdateMachineAccountSubtypeApprovalConfigV1`: %v\n", resp)
}
```

[[Back to top]](#)

## update-machine-account-v1
:::warning experimental 
This API is currently in an experimental state. The API is subject to change based on feedback and further testing. You must include the X-SailPoint-Experimental header and set it to `true` to use this endpoint.
:::
:::tip setting x-sailpoint-experimental header
 on the configuration object you can set the `x-sailpoint-experimental` header to `true' to enable all experimantl endpoints within the SDK.
 Example:
 ```go
   configuration = Configuration()
   configuration.Experimental = true
 ```
:::
Update machine account details
Use this API to update machine accounts details. 


[API Spec](https://developer.sailpoint.com/docs/api/v1/update-machine-account-v1)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Machine Account ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMachineAccountV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **requestBody** | **[]map[string]interface{}** | A JSON of updated values [JSON Patch](https://tools.ietf.org/html/rfc6902) standard. The following fields are patchable:           * description           * ownerIdentity           * subType           * accessType           * environment           * attributes           * classificationMethod           * manuallyEdited           * nativeIdentity           * uuid           * source           * manuallyCorrelated           * enabled           * locked           * hasEntitlements           * connectorAttributes | 

### Return type

[**Machineaccount**](../models/machineaccount)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
  "encoding/json"
    v1 "github.com/sailpoint-oss/golang-sdk/v3/api_machine_accounts"
	sailpoint "github.com/sailpoint-oss/golang-sdk/v3/api_machine_accounts"
)

func main() {
    id := `ef38f94347e94562b5bb8424a56397d8` // string | Machine Account ID. # string | Machine Account ID.
    xSailPointExperimental := `true` // string | Use this header to enable this experimental API. (default to "true") # string | Use this header to enable this experimental API. (default to "true")
    requestbody := []byte(`[{"op":"add","path":"/environment","value":"test"}]`) // []map[string]interface{} | A JSON of updated values [JSON Patch](https://tools.ietf.org/html/rfc6902) standard. The following fields are patchable:           * description           * ownerIdentity           * subType           * accessType           * environment           * attributes           * classificationMethod           * manuallyEdited           * nativeIdentity           * uuid           * source           * manuallyCorrelated           * enabled           * locked           * hasEntitlements           * connectorAttributes

    var requestBody []v1.RequestBody
    if err := json.Unmarshal(requestbody, &requestBody); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineAccounts.MachineAccountsAPI.UpdateMachineAccountV1(context.Background(), id).XSailPointExperimental(xSailPointExperimental).RequestBody(requestBody).Execute()
	  //resp, r, err := apiClient.MachineAccounts.MachineAccountsAPI.UpdateMachineAccountV1(context.Background(), id).XSailPointExperimental(xSailPointExperimental).RequestBody(requestBody).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `MachineAccountsAPI.UpdateMachineAccountV1``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMachineAccountV1`: Machineaccount
    fmt.Fprintf(os.Stdout, "Response from `MachineAccountsAPI.UpdateMachineAccountV1`: %v\n", resp)
}
```

[[Back to top]](#)

