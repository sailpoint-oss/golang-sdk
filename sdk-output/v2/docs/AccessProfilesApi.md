# \AccessProfilesApi

All URIs are relative to *https://sailpoint.api.identitynow.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccessProfile**](AccessProfilesApi.md#CreateAccessProfile) | **Post** /access-profiles | Creates a new access profile.
[**DeleteAccessProfile**](AccessProfilesApi.md#DeleteAccessProfile) | **Delete** /access-profiles/{accessProfileId} | Deletes an access profile.
[**GetAccessProfile**](AccessProfilesApi.md#GetAccessProfile) | **Get** /access-profiles/{accessProfileId} | Retrieves an access profile.
[**ListAccessProfileEntitlements**](AccessProfilesApi.md#ListAccessProfileEntitlements) | **Get** /access-profiles/{accessProfileId}/entitlements | Lists all entitlements of an access profile.
[**ListAccessProfiles**](AccessProfilesApi.md#ListAccessProfiles) | **Get** /access-profiles | Lists the access profiles.
[**PatchAccessProfile**](AccessProfilesApi.md#PatchAccessProfile) | **Patch** /access-profiles/{accessProfileId} | Updates one or more access profile attributes.
[**PutAccessProfile**](AccessProfilesApi.md#PutAccessProfile) | **Put** /access-profiles/{accessProfileId} | Updates an existing access profile.



## CreateAccessProfile

> map[string]interface{} CreateAccessProfile(ctx).AccessProfile(accessProfile).Execute()

Creates a new access profile.



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
    accessProfile := *openapiclient.NewAccessProfile() // AccessProfile | Attribute values for the new access profile.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessProfilesApi.CreateAccessProfile(context.Background()).AccessProfile(accessProfile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessProfilesApi.CreateAccessProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAccessProfile`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AccessProfilesApi.CreateAccessProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccessProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessProfile** | [**AccessProfile**](AccessProfile.md) | Attribute values for the new access profile. | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAccessProfile

> DeleteAccessProfile(ctx, accessProfileId).Execute()

Deletes an access profile.



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
    accessProfileId := "accessProfileId_example" // string | ID of an access profile.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessProfilesApi.DeleteAccessProfile(context.Background(), accessProfileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessProfilesApi.DeleteAccessProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessProfileId** | **string** | ID of an access profile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccessProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccessProfile

> AccessProfile GetAccessProfile(ctx, accessProfileId).Execute()

Retrieves an access profile.



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
    accessProfileId := "accessProfileId_example" // string | ID of an access profile.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessProfilesApi.GetAccessProfile(context.Background(), accessProfileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessProfilesApi.GetAccessProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccessProfile`: AccessProfile
    fmt.Fprintf(os.Stdout, "Response from `AccessProfilesApi.GetAccessProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessProfileId** | **string** | ID of an access profile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccessProfile**](AccessProfile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccessProfileEntitlements

> Entitlement ListAccessProfileEntitlements(ctx, accessProfileId).Execute()

Lists all entitlements of an access profile.



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
    accessProfileId := "accessProfileId_example" // string | ID of an access profile.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessProfilesApi.ListAccessProfileEntitlements(context.Background(), accessProfileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessProfilesApi.ListAccessProfileEntitlements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAccessProfileEntitlements`: Entitlement
    fmt.Fprintf(os.Stdout, "Response from `AccessProfilesApi.ListAccessProfileEntitlements`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessProfileId** | **string** | ID of an access profile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAccessProfileEntitlementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Entitlement**](Entitlement.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccessProfiles

> []AccessProfile ListAccessProfiles(ctx).Filters(filters).Sort(sort).Offset(offset).Limit(limit).Execute()

Lists the access profiles.



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
    filters := "filters_example" // string | Search filters. Example: 'property EQ \"value\"' (optional)
    sort := "sort_example" // string | One or more attributes on which to sort, each separated by a ','. Prefix with a minus sign (ex. -dateCreated) for descending sort. (optional)
    offset := int32(56) // int32 | Paging offset. (optional) (default to 0)
    limit := int32(56) // int32 | Paging limit. (optional) (default to 250)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessProfilesApi.ListAccessProfiles(context.Background()).Filters(filters).Sort(sort).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessProfilesApi.ListAccessProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAccessProfiles`: []AccessProfile
    fmt.Fprintf(os.Stdout, "Response from `AccessProfilesApi.ListAccessProfiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAccessProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filters** | **string** | Search filters. Example: &#39;property EQ \&quot;value\&quot;&#39; | 
 **sort** | **string** | One or more attributes on which to sort, each separated by a &#39;,&#39;. Prefix with a minus sign (ex. -dateCreated) for descending sort. | 
 **offset** | **int32** | Paging offset. | [default to 0]
 **limit** | **int32** | Paging limit. | [default to 250]

### Return type

[**[]AccessProfile**](AccessProfile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAccessProfile

> AccessProfile PatchAccessProfile(ctx, accessProfileId).AccessProfile(accessProfile).Execute()

Updates one or more access profile attributes.



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
    accessProfileId := "accessProfileId_example" // string | ID of an access profile.
    accessProfile := *openapiclient.NewAccessProfile() // AccessProfile | Access profile attributes to update.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessProfilesApi.PatchAccessProfile(context.Background(), accessProfileId).AccessProfile(accessProfile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessProfilesApi.PatchAccessProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchAccessProfile`: AccessProfile
    fmt.Fprintf(os.Stdout, "Response from `AccessProfilesApi.PatchAccessProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessProfileId** | **string** | ID of an access profile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAccessProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessProfile** | [**AccessProfile**](AccessProfile.md) | Access profile attributes to update. | 

### Return type

[**AccessProfile**](AccessProfile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAccessProfile

> AccessProfile PutAccessProfile(ctx, accessProfileId).AccessProfileCreateEto(accessProfileCreateEto).Execute()

Updates an existing access profile.



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
    accessProfileId := "accessProfileId_example" // string | ID of an access profile.
    accessProfileCreateEto := *openapiclient.NewAccessProfileCreateEto("Name_example", "Description_example", "SourceId_example", "OwnerId_example", []string{"Entitlements_example"}) // AccessProfileCreateEto | Access profile attributes to update.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessProfilesApi.PutAccessProfile(context.Background(), accessProfileId).AccessProfileCreateEto(accessProfileCreateEto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessProfilesApi.PutAccessProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutAccessProfile`: AccessProfile
    fmt.Fprintf(os.Stdout, "Response from `AccessProfilesApi.PutAccessProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessProfileId** | **string** | ID of an access profile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutAccessProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessProfileCreateEto** | [**AccessProfileCreateEto**](AccessProfileCreateEto.md) | Access profile attributes to update. | 

### Return type

[**AccessProfile**](AccessProfile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

