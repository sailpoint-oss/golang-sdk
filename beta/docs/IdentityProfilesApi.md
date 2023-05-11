# \IdentityProfilesApi

All URIs are relative to *https://sailpoint.api.identitynow.com/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIdentityProfile**](IdentityProfilesApi.md#CreateIdentityProfile) | **Post** /identity-profiles | Create an Identity Profile
[**DeleteIdentityProfile**](IdentityProfilesApi.md#DeleteIdentityProfile) | **Delete** /identity-profiles/{identity-profile-id} | Delete an Identity Profile
[**DeleteIdentityProfiles**](IdentityProfilesApi.md#DeleteIdentityProfiles) | **Post** /identity-profiles/bulk-delete | Delete Identity Profiles
[**ExportIdentityProfiles**](IdentityProfilesApi.md#ExportIdentityProfiles) | **Get** /identity-profiles/export | Export Identity Profiles
[**GenerateIdentityPreview**](IdentityProfilesApi.md#GenerateIdentityPreview) | **Post** /identity-profiles/identity-preview | Generate Identity Profile Preview
[**GetDefaultIdentityAttributeConfig**](IdentityProfilesApi.md#GetDefaultIdentityAttributeConfig) | **Get** /identity-profiles/{identity-profile-id}/default-identity-attribute-config | Default identity attribute config
[**GetIdentityProfile**](IdentityProfilesApi.md#GetIdentityProfile) | **Get** /identity-profiles/{identity-profile-id} | Gets a single Identity Profile
[**ImportIdentityProfiles**](IdentityProfilesApi.md#ImportIdentityProfiles) | **Post** /identity-profiles/import | Import Identity Profiles
[**ListIdentityProfiles**](IdentityProfilesApi.md#ListIdentityProfiles) | **Get** /identity-profiles | Identity Profiles list
[**RefreshIdentityProfile**](IdentityProfilesApi.md#RefreshIdentityProfile) | **Post** /identity-profiles/{identity-profile-id}/refresh-identities | Refreshes all identities under profile
[**UpdateIdentityProfile**](IdentityProfilesApi.md#UpdateIdentityProfile) | **Patch** /identity-profiles/{identity-profile-id} | Update the Identity Profile



## CreateIdentityProfile

> IdentityProfile CreateIdentityProfile(ctx).IdentityProfile(identityProfile).Execute()

Create an Identity Profile



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    identityProfile := *openapiclient.NewIdentityProfile("aName", *openapiclient.NewIdentityProfileAllOfAuthoritativeSource()) // IdentityProfile | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProfilesApi.CreateIdentityProfile(context.Background()).IdentityProfile(identityProfile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProfilesApi.CreateIdentityProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIdentityProfile`: IdentityProfile
    fmt.Fprintf(os.Stdout, "Response from `IdentityProfilesApi.CreateIdentityProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIdentityProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identityProfile** | [**IdentityProfile**](IdentityProfile.md) |  | 

### Return type

[**IdentityProfile**](IdentityProfile.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIdentityProfile

> BaseReferenceDto1 DeleteIdentityProfile(ctx, identityProfileId).Execute()

Delete an Identity Profile



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    identityProfileId := "ef38f94347e94562b5bb8424a56397d8" // string | The Identity Profile ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProfilesApi.DeleteIdentityProfile(context.Background(), identityProfileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProfilesApi.DeleteIdentityProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteIdentityProfile`: BaseReferenceDto1
    fmt.Fprintf(os.Stdout, "Response from `IdentityProfilesApi.DeleteIdentityProfile`: %v\n", resp)
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

[**BaseReferenceDto1**](BaseReferenceDto1.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIdentityProfiles

> BaseReferenceDto DeleteIdentityProfiles(ctx).RequestBody(requestBody).Execute()

Delete Identity Profiles



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    requestBody := []string{"Property_example"} // []string | Identity Profile bulk delete request body.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProfilesApi.DeleteIdentityProfiles(context.Background()).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProfilesApi.DeleteIdentityProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteIdentityProfiles`: BaseReferenceDto
    fmt.Fprintf(os.Stdout, "Response from `IdentityProfilesApi.DeleteIdentityProfiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdentityProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]string** | Identity Profile bulk delete request body. | 

### Return type

[**BaseReferenceDto**](BaseReferenceDto.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

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
    openapiclient "./openapi"
)

func main() {
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
    filters := "id eq 8c190e6787aa4ed9a90bd9d5344523fb" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq, ne*  **name**: *eq, ne*  **priority**: *eq, ne* (optional)
    sorters := "name,-priority" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **id**, **name**, **priority** (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProfilesApi.ExportIdentityProfiles(context.Background()).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProfilesApi.ExportIdentityProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportIdentityProfiles`: []IdentityProfileExportedObject
    fmt.Fprintf(os.Stdout, "Response from `IdentityProfilesApi.ExportIdentityProfiles`: %v\n", resp)
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
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **id**, **name**, **priority** | 

### Return type

[**[]IdentityProfileExportedObject**](IdentityProfileExportedObject.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateIdentityPreview

> IdentityPreviewResponse GenerateIdentityPreview(ctx).IdentityPreviewRequest(identityPreviewRequest).Execute()

Generate Identity Profile Preview



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    identityPreviewRequest := *openapiclient.NewIdentityPreviewRequest() // IdentityPreviewRequest | Identity Preview request body.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProfilesApi.GenerateIdentityPreview(context.Background()).IdentityPreviewRequest(identityPreviewRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProfilesApi.GenerateIdentityPreview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateIdentityPreview`: IdentityPreviewResponse
    fmt.Fprintf(os.Stdout, "Response from `IdentityProfilesApi.GenerateIdentityPreview`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateIdentityPreviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identityPreviewRequest** | [**IdentityPreviewRequest**](IdentityPreviewRequest.md) | Identity Preview request body. | 

### Return type

[**IdentityPreviewResponse**](IdentityPreviewResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDefaultIdentityAttributeConfig

> IdentityAttributeConfig GetDefaultIdentityAttributeConfig(ctx, identityProfileId).Execute()

Default identity attribute config



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    identityProfileId := "ef38f94347e94562b5bb8424a56397d8" // string | The Identity Profile ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProfilesApi.GetDefaultIdentityAttributeConfig(context.Background(), identityProfileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProfilesApi.GetDefaultIdentityAttributeConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDefaultIdentityAttributeConfig`: IdentityAttributeConfig
    fmt.Fprintf(os.Stdout, "Response from `IdentityProfilesApi.GetDefaultIdentityAttributeConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityProfileId** | **string** | The Identity Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefaultIdentityAttributeConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IdentityAttributeConfig**](IdentityAttributeConfig.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityProfile

> IdentityProfile GetIdentityProfile(ctx, identityProfileId).Execute()

Gets a single Identity Profile



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    identityProfileId := "ef38f94347e94562b5bb8424a56397d8" // string | The Identity Profile ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProfilesApi.GetIdentityProfile(context.Background(), identityProfileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProfilesApi.GetIdentityProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdentityProfile`: IdentityProfile
    fmt.Fprintf(os.Stdout, "Response from `IdentityProfilesApi.GetIdentityProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityProfileId** | **string** | The Identity Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IdentityProfile**](IdentityProfile.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

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
    openapiclient "./openapi"
)

func main() {
    identityProfileExportedObject := []openapiclient.IdentityProfileExportedObject{*openapiclient.NewIdentityProfileExportedObject()} // []IdentityProfileExportedObject | Previously exported Identity Profiles.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProfilesApi.ImportIdentityProfiles(context.Background()).IdentityProfileExportedObject(identityProfileExportedObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProfilesApi.ImportIdentityProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportIdentityProfiles`: ObjectImportResult
    fmt.Fprintf(os.Stdout, "Response from `IdentityProfilesApi.ImportIdentityProfiles`: %v\n", resp)
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

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIdentityProfiles

> []IdentityProfile ListIdentityProfiles(ctx).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).Execute()

Identity Profiles list



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
    filters := "id eq 8c190e6787aa4ed9a90bd9d5344523fb" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq, ne*  **name**: *eq, ne*  **priority**: *eq, ne* (optional)
    sorters := "name,-priority" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **id**, **name**, **priority** (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProfilesApi.ListIdentityProfiles(context.Background()).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProfilesApi.ListIdentityProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIdentityProfiles`: []IdentityProfile
    fmt.Fprintf(os.Stdout, "Response from `IdentityProfilesApi.ListIdentityProfiles`: %v\n", resp)
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
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq, ne*  **name**: *eq, ne*  **priority**: *eq, ne* | 
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **id**, **name**, **priority** | 

### Return type

[**[]IdentityProfile**](IdentityProfile.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefreshIdentityProfile

> map[string]interface{} RefreshIdentityProfile(ctx, identityProfileId).Execute()

Refreshes all identities under profile



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    identityProfileId := "ef38f94347e94562b5bb8424a56397d8" // string | The Identity Profile ID to be refreshed

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProfilesApi.RefreshIdentityProfile(context.Background(), identityProfileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProfilesApi.RefreshIdentityProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RefreshIdentityProfile`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `IdentityProfilesApi.RefreshIdentityProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityProfileId** | **string** | The Identity Profile ID to be refreshed | 

### Other Parameters

Other parameters are passed through a pointer to a apiRefreshIdentityProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIdentityProfile

> IdentityProfile UpdateIdentityProfile(ctx, identityProfileId).JsonPatchOperation(jsonPatchOperation).Execute()

Update the Identity Profile



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    identityProfileId := "ef38f94347e94562b5bb8424a56397d8" // string | The Identity Profile ID
    jsonPatchOperation := []openapiclient.JsonPatchOperation{*openapiclient.NewJsonPatchOperation("replace", "/description")} // []JsonPatchOperation | A list of Identity Profile update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProfilesApi.UpdateIdentityProfile(context.Background(), identityProfileId).JsonPatchOperation(jsonPatchOperation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProfilesApi.UpdateIdentityProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIdentityProfile`: IdentityProfile
    fmt.Fprintf(os.Stdout, "Response from `IdentityProfilesApi.UpdateIdentityProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityProfileId** | **string** | The Identity Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIdentityProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jsonPatchOperation** | [**[]JsonPatchOperation**](JsonPatchOperation.md) | A list of Identity Profile update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard. | 

### Return type

[**IdentityProfile**](IdentityProfile.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

