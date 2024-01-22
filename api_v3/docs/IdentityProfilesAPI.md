# \IdentityProfilesAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIdentityProfile**](IdentityProfilesAPI.md#DeleteIdentityProfile) | **Delete** /identity-profiles/{identity-profile-id} | Delete an Identity Profile
[**DeleteIdentityProfiles**](IdentityProfilesAPI.md#DeleteIdentityProfiles) | **Post** /identity-profiles/bulk-delete | Delete Identity Profiles
[**ExportIdentityProfiles**](IdentityProfilesAPI.md#ExportIdentityProfiles) | **Get** /identity-profiles/export | Export Identity Profiles
[**GetDefaultIdentityAttributeConfig**](IdentityProfilesAPI.md#GetDefaultIdentityAttributeConfig) | **Get** /identity-profiles/{identity-profile-id}/default-identity-attribute-config | Get default Identity Attribute Config
[**GetIdentityProfile**](IdentityProfilesAPI.md#GetIdentityProfile) | **Get** /identity-profiles/{identity-profile-id} | Get single Identity Profile
[**ImportIdentityProfiles**](IdentityProfilesAPI.md#ImportIdentityProfiles) | **Post** /identity-profiles/import | Import Identity Profiles
[**ListIdentityProfiles**](IdentityProfilesAPI.md#ListIdentityProfiles) | **Get** /identity-profiles | Identity Profiles List
[**SyncIdentityProfile**](IdentityProfilesAPI.md#SyncIdentityProfile) | **Post** /identity-profiles/{identity-profile-id}/process-identities | Process identities under profile



## DeleteIdentityProfile

> TaskResultSimplified DeleteIdentityProfile(ctx, identityProfileId).Execute()

Delete an Identity Profile



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/sailpoint-oss/golang-sdk"
)

func main() {
    identityProfileId := "ef38f94347e94562b5bb8424a56397d8" // string | The Identity Profile ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProfilesAPI.DeleteIdentityProfile(context.Background(), identityProfileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProfilesAPI.DeleteIdentityProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteIdentityProfile`: TaskResultSimplified
    fmt.Fprintf(os.Stdout, "Response from `IdentityProfilesAPI.DeleteIdentityProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityProfileId** | **string** | The Identity Profile ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdentityProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TaskResultSimplified**](TaskResultSimplified.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIdentityProfiles

> TaskResultSimplified DeleteIdentityProfiles(ctx).RequestBody(requestBody).Execute()

Delete Identity Profiles



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/sailpoint-oss/golang-sdk"
)

func main() {
    requestBody := []string{"Property_example"} // []string | Identity Profile bulk delete request body.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProfilesAPI.DeleteIdentityProfiles(context.Background()).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProfilesAPI.DeleteIdentityProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteIdentityProfiles`: TaskResultSimplified
    fmt.Fprintf(os.Stdout, "Response from `IdentityProfilesAPI.DeleteIdentityProfiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdentityProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]string** | Identity Profile bulk delete request body. | 

### Return type

[**TaskResultSimplified**](TaskResultSimplified.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportIdentityProfiles

> []IdentityProfileExportedObject ExportIdentityProfiles(ctx).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).Execute()

Export Identity Profiles



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/sailpoint-oss/golang-sdk"
)

func main() {
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
    filters := "id eq "ef38f94347e94562b5bb8424a56397d8"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq, ne*  **name**: *eq, ne*  **priority**: *eq, ne* (optional)
    sorters := "id,name" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **id, name, priority** (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProfilesAPI.ExportIdentityProfiles(context.Background()).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProfilesAPI.ExportIdentityProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportIdentityProfiles`: []IdentityProfileExportedObject
    fmt.Fprintf(os.Stdout, "Response from `IdentityProfilesAPI.ExportIdentityProfiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportIdentityProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq, ne*  **name**: *eq, ne*  **priority**: *eq, ne* | 
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **id, name, priority** | 

### Return type

[**[]IdentityProfileExportedObject**](IdentityProfileExportedObject.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDefaultIdentityAttributeConfig

> IdentityAttributeConfig GetDefaultIdentityAttributeConfig(ctx, identityProfileId).Execute()

Get default Identity Attribute Config



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/sailpoint-oss/golang-sdk"
)

func main() {
    identityProfileId := "2b838de9-db9b-abcf-e646-d4f274ad4238" // string | The Identity Profile ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProfilesAPI.GetDefaultIdentityAttributeConfig(context.Background(), identityProfileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProfilesAPI.GetDefaultIdentityAttributeConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDefaultIdentityAttributeConfig`: IdentityAttributeConfig
    fmt.Fprintf(os.Stdout, "Response from `IdentityProfilesAPI.GetDefaultIdentityAttributeConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityProfileId** | **string** | The Identity Profile ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefaultIdentityAttributeConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IdentityAttributeConfig**](IdentityAttributeConfig.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityProfile

> IdentityProfile GetIdentityProfile(ctx, identityProfileId).Execute()

Get single Identity Profile



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/sailpoint-oss/golang-sdk"
)

func main() {
    identityProfileId := "2b838de9-db9b-abcf-e646-d4f274ad4238" // string | The Identity Profile ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProfilesAPI.GetIdentityProfile(context.Background(), identityProfileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProfilesAPI.GetIdentityProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdentityProfile`: IdentityProfile
    fmt.Fprintf(os.Stdout, "Response from `IdentityProfilesAPI.GetIdentityProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityProfileId** | **string** | The Identity Profile ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IdentityProfile**](IdentityProfile.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportIdentityProfiles

> ObjectImportResult ImportIdentityProfiles(ctx).IdentityProfileExportedObject(identityProfileExportedObject).Execute()

Import Identity Profiles



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/sailpoint-oss/golang-sdk"
)

func main() {
    identityProfileExportedObject := []openapiclient.IdentityProfileExportedObject{*openapiclient.NewIdentityProfileExportedObject()} // []IdentityProfileExportedObject | Previously exported Identity Profiles.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProfilesAPI.ImportIdentityProfiles(context.Background()).IdentityProfileExportedObject(identityProfileExportedObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProfilesAPI.ImportIdentityProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportIdentityProfiles`: ObjectImportResult
    fmt.Fprintf(os.Stdout, "Response from `IdentityProfilesAPI.ImportIdentityProfiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportIdentityProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identityProfileExportedObject** | [**[]IdentityProfileExportedObject**](IdentityProfileExportedObject.md) | Previously exported Identity Profiles. | 

### Return type

[**ObjectImportResult**](ObjectImportResult.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIdentityProfiles

> []IdentityProfile ListIdentityProfiles(ctx).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).Execute()

Identity Profiles List



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/sailpoint-oss/golang-sdk"
)

func main() {
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
    filters := "id eq "ef38f94347e94562b5bb8424a56397d8"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq, ne, ge, gt, in, le, lt, isnull, sw*  **name**: *eq, ne, in, le, lt, isnull, sw*  **priority**: *eq, ne* (optional)
    sorters := "id,name" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **id, name, priority, created, modified, owner.id, owner.name** (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProfilesAPI.ListIdentityProfiles(context.Background()).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProfilesAPI.ListIdentityProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIdentityProfiles`: []IdentityProfile
    fmt.Fprintf(os.Stdout, "Response from `IdentityProfilesAPI.ListIdentityProfiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListIdentityProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq, ne, ge, gt, in, le, lt, isnull, sw*  **name**: *eq, ne, in, le, lt, isnull, sw*  **priority**: *eq, ne* | 
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **id, name, priority, created, modified, owner.id, owner.name** | 

### Return type

[**[]IdentityProfile**](IdentityProfile.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncIdentityProfile

> map[string]interface{} SyncIdentityProfile(ctx, identityProfileId).Execute()

Process identities under profile



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/sailpoint-oss/golang-sdk"
)

func main() {
    identityProfileId := "ef38f94347e94562b5bb8424a56397d8" // string | The Identity Profile ID to be processed

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProfilesAPI.SyncIdentityProfile(context.Background(), identityProfileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProfilesAPI.SyncIdentityProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SyncIdentityProfile`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `IdentityProfilesAPI.SyncIdentityProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityProfileId** | **string** | The Identity Profile ID to be processed | 

### Other Parameters

Other parameters are passed through a pointer to a apiSyncIdentityProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

