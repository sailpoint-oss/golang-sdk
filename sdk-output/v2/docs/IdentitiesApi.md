# \IdentitiesApi

All URIs are relative to *https://sailpoint.api.identitynow.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIdentity**](IdentitiesApi.md#CreateIdentity) | **Post** /identities | Creates a new identity.
[**CreateLauncher**](IdentitiesApi.md#CreateLauncher) | **Post** /identities/{identityIdOrAlias}/launchers | Creates a new launcher.
[**DeleteIdentity**](IdentitiesApi.md#DeleteIdentity) | **Delete** /identities/{identityIdOrAlias} | Deletes an identity.
[**DeleteLauncher**](IdentitiesApi.md#DeleteLauncher) | **Delete** /identities/{identityIdOrAlias}/launchers/{launcherId} | Deletes a launcher.
[**GetIdentity**](IdentitiesApi.md#GetIdentity) | **Get** /identities/{identityIdOrAlias} | Retrieves the identity by ID or alias.
[**ListApprovals**](IdentitiesApi.md#ListApprovals) | **Get** /identities/{identityIdOrAlias}/approvals | Lists the approvals.
[**ListApps**](IdentitiesApi.md#ListApps) | **Get** /identities/{identityIdOrAlias}/apps | Lists available apps.
[**ListIdentities**](IdentitiesApi.md#ListIdentities) | **Get** /identities | Retrieves the identities.
[**ListLaunchers**](IdentitiesApi.md#ListLaunchers) | **Get** /identities/{identityIdOrAlias}/launchers | Lists the launchers.
[**LockIdentities**](IdentitiesApi.md#LockIdentities) | **Post** /identities/bulk-lock | Locks one or more identities.
[**UpdateIdentity**](IdentitiesApi.md#UpdateIdentity) | **Patch** /identities/{identityIdOrAlias} | Updates one or more identity attributes.



## CreateIdentity

> CreateIdentity(ctx).SourceId(sourceId).DynamicSchemaEto(dynamicSchemaEto).Execute()

Creates a new identity.



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
    sourceId := "sourceId_example" // string | ID of a flat-file source into which the new account will be created.
    dynamicSchemaEto := *openapiclient.NewDynamicSchemaEto() // DynamicSchemaEto | Attribute values for the new identity. The schema and required attributes are dictated by the source.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentitiesApi.CreateIdentity(context.Background()).SourceId(sourceId).DynamicSchemaEto(dynamicSchemaEto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentitiesApi.CreateIdentity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIdentityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sourceId** | **string** | ID of a flat-file source into which the new account will be created. | 
 **dynamicSchemaEto** | [**DynamicSchemaEto**](DynamicSchemaEto.md) | Attribute values for the new identity. The schema and required attributes are dictated by the source. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLauncher

> Launcher CreateLauncher(ctx, identityIdOrAlias, accessProfileId).AppId(appId).Context(context).Execute()

Creates a new launcher.



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
    identityIdOrAlias := "identityIdOrAlias_example" // string | ID or alias of an identity.
    appId := "appId_example" // string | ID of the app to create a launcher for.
    accessProfileId := "accessProfileId_example" // string | ID of an access profile.
    context := "context_example" // string | Context of launcher links to include. Specifying a mobile context will provide links to mobile resources for the launcher if available. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentitiesApi.CreateLauncher(context.Background(), identityIdOrAlias, accessProfileId).AppId(appId).Context(context).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentitiesApi.CreateLauncher``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLauncher`: Launcher
    fmt.Fprintf(os.Stdout, "Response from `IdentitiesApi.CreateLauncher`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityIdOrAlias** | **string** | ID or alias of an identity. | 
**accessProfileId** | **string** | ID of an access profile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLauncherRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appId** | **string** | ID of the app to create a launcher for. | 

 **context** | **string** | Context of launcher links to include. Specifying a mobile context will provide links to mobile resources for the launcher if available. | 

### Return type

[**Launcher**](Launcher.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIdentity

> DeleteIdentity(ctx, identityIdOrAlias).Execute()

Deletes an identity.



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
    identityIdOrAlias := "identityIdOrAlias_example" // string | ID or alias of an identity.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentitiesApi.DeleteIdentity(context.Background(), identityIdOrAlias).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentitiesApi.DeleteIdentity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityIdOrAlias** | **string** | ID or alias of an identity. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdentityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLauncher

> DeleteLauncher(ctx, identityIdOrAlias, launcherId).Execute()

Deletes a launcher.



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
    identityIdOrAlias := "identityIdOrAlias_example" // string | ID or alias of an identity.
    launcherId := "launcherId_example" // string | ID of a launcher.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentitiesApi.DeleteLauncher(context.Background(), identityIdOrAlias, launcherId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentitiesApi.DeleteLauncher``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityIdOrAlias** | **string** | ID or alias of an identity. | 
**launcherId** | **string** | ID of a launcher. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLauncherRequest struct via the builder pattern


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


## GetIdentity

> IdentityV2 GetIdentity(ctx, identityIdOrAlias).Execute()

Retrieves the identity by ID or alias.



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
    identityIdOrAlias := "identityIdOrAlias_example" // string | ID or alias of an identity.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentitiesApi.GetIdentity(context.Background(), identityIdOrAlias).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentitiesApi.GetIdentity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdentity`: IdentityV2
    fmt.Fprintf(os.Stdout, "Response from `IdentitiesApi.GetIdentity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityIdOrAlias** | **string** | ID or alias of an identity. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IdentityV2**](IdentityV2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApprovals

> []Approval ListApprovals(ctx, identityIdOrAlias).Sort(sort).Offset(offset).Limit(limit).Execute()

Lists the approvals.



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
    identityIdOrAlias := "identityIdOrAlias_example" // string | ID or alias of an identity.
    sort := "sort_example" // string | One or more attributes on which to sort, each separated by a ','. Prefix with a minus sign (ex. -dateCreated) for descending sort. (optional)
    offset := int32(56) // int32 | Paging offset. (optional) (default to 0)
    limit := int32(56) // int32 | Paging limit. (optional) (default to 250)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentitiesApi.ListApprovals(context.Background(), identityIdOrAlias).Sort(sort).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentitiesApi.ListApprovals``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApprovals`: []Approval
    fmt.Fprintf(os.Stdout, "Response from `IdentitiesApi.ListApprovals`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityIdOrAlias** | **string** | ID or alias of an identity. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListApprovalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sort** | **string** | One or more attributes on which to sort, each separated by a &#39;,&#39;. Prefix with a minus sign (ex. -dateCreated) for descending sort. | 
 **offset** | **int32** | Paging offset. | [default to 0]
 **limit** | **int32** | Paging limit. | [default to 250]

### Return type

[**[]Approval**](Approval.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApps

> []App ListApps(ctx, identityIdOrAlias).Filters(filters).Sort(sort).Offset(offset).Limit(limit).Execute()

Lists available apps.



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
    identityIdOrAlias := "identityIdOrAlias_example" // string | ID or alias of an identity.
    filters := "filters_example" // string | Search filters. Supported operator is 'EQ'. Supported attribute for filtering is ‘description’. Example: 'description EQ \"Some description.\"'. (optional)
    sort := "sort_example" // string | One or more attributes on which to sort, each separated by a ','. Prefix with a minus sign (ex. -dateCreated) for descending sort. (optional)
    offset := int32(56) // int32 | Paging offset. (optional) (default to 0)
    limit := int32(56) // int32 | Paging limit. (optional) (default to 250)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentitiesApi.ListApps(context.Background(), identityIdOrAlias).Filters(filters).Sort(sort).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentitiesApi.ListApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApps`: []App
    fmt.Fprintf(os.Stdout, "Response from `IdentitiesApi.ListApps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityIdOrAlias** | **string** | ID or alias of an identity. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filters** | **string** | Search filters. Supported operator is &#39;EQ&#39;. Supported attribute for filtering is ‘description’. Example: &#39;description EQ \&quot;Some description.\&quot;&#39;. | 
 **sort** | **string** | One or more attributes on which to sort, each separated by a &#39;,&#39;. Prefix with a minus sign (ex. -dateCreated) for descending sort. | 
 **offset** | **int32** | Paging offset. | [default to 0]
 **limit** | **int32** | Paging limit. | [default to 250]

### Return type

[**[]App**](App.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIdentities

> []IdentityV2 ListIdentities(ctx).Filters(filters).Sort(sort).Offset(offset).Limit(limit).Execute()

Retrieves the identities.



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
    filters := "filters_example" // string | Search filters. Supported operator is 'EQ'. Supported attribute for filtering is ‘permissions’. Example: 'permissions EQ \"admin\"'. (optional)
    sort := "sort_example" // string | One or more attributes on which to sort, each separated by a ','. Prefix with a minus sign (ex. -dateCreated) for descending sort. (optional)
    offset := int32(56) // int32 | Paging offset. (optional) (default to 0)
    limit := int32(56) // int32 | Paging limit. (optional) (default to 250)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentitiesApi.ListIdentities(context.Background()).Filters(filters).Sort(sort).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentitiesApi.ListIdentities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIdentities`: []IdentityV2
    fmt.Fprintf(os.Stdout, "Response from `IdentitiesApi.ListIdentities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListIdentitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filters** | **string** | Search filters. Supported operator is &#39;EQ&#39;. Supported attribute for filtering is ‘permissions’. Example: &#39;permissions EQ \&quot;admin\&quot;&#39;. | 
 **sort** | **string** | One or more attributes on which to sort, each separated by a &#39;,&#39;. Prefix with a minus sign (ex. -dateCreated) for descending sort. | 
 **offset** | **int32** | Paging offset. | [default to 0]
 **limit** | **int32** | Paging limit. | [default to 250]

### Return type

[**[]IdentityV2**](IdentityV2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLaunchers

> []Launcher ListLaunchers(ctx, identityIdOrAlias).Context(context).Sort(sort).Execute()

Lists the launchers.



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
    identityIdOrAlias := "identityIdOrAlias_example" // string | ID or alias of an identity.
    context := "context_example" // string | Context of launcher links to include. Specifying a mobile context will provide links to mobile resources for the launcher if available. (optional)
    sort := "sort_example" // string | One or more attributes on which to sort, each separated by a ','. Prefix with a minus sign (ex. -dateCreated) for descending sort. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentitiesApi.ListLaunchers(context.Background(), identityIdOrAlias).Context(context).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentitiesApi.ListLaunchers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLaunchers`: []Launcher
    fmt.Fprintf(os.Stdout, "Response from `IdentitiesApi.ListLaunchers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityIdOrAlias** | **string** | ID or alias of an identity. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLaunchersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **context** | **string** | Context of launcher links to include. Specifying a mobile context will provide links to mobile resources for the launcher if available. | 
 **sort** | **string** | One or more attributes on which to sort, each separated by a &#39;,&#39;. Prefix with a minus sign (ex. -dateCreated) for descending sort. | 

### Return type

[**[]Launcher**](Launcher.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LockIdentities

> MultiStatusObject LockIdentities(ctx).RequestBody(requestBody).Execute()

Locks one or more identities.



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
    requestBody := []string{"Property_example"} // []string | Array of one or more IDs.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentitiesApi.LockIdentities(context.Background()).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentitiesApi.LockIdentities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LockIdentities`: MultiStatusObject
    fmt.Fprintf(os.Stdout, "Response from `IdentitiesApi.LockIdentities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLockIdentitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]string** | Array of one or more IDs. | 

### Return type

[**MultiStatusObject**](MultiStatusObject.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIdentity

> IdentityV2 UpdateIdentity(ctx, identityIdOrAlias).IdentityEto(identityEto).Execute()

Updates one or more identity attributes.



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
    identityIdOrAlias := "identityIdOrAlias_example" // string | ID or alias of an identity.
    identityEto := *openapiclient.NewIdentityEto() // IdentityEto | Identity attributes to update.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentitiesApi.UpdateIdentity(context.Background(), identityIdOrAlias).IdentityEto(identityEto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentitiesApi.UpdateIdentity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIdentity`: IdentityV2
    fmt.Fprintf(os.Stdout, "Response from `IdentitiesApi.UpdateIdentity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityIdOrAlias** | **string** | ID or alias of an identity. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIdentityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **identityEto** | [**IdentityEto**](IdentityEto.md) | Identity attributes to update. | 

### Return type

[**IdentityV2**](IdentityV2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

