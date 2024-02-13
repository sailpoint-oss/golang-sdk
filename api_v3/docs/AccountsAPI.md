# \AccountsAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccount**](AccountsAPI.md#CreateAccount) | **Post** /accounts | Create Account
[**DeleteAccount**](AccountsAPI.md#DeleteAccount) | **Delete** /accounts/{id} | Delete Account
[**DisableAccount**](AccountsAPI.md#DisableAccount) | **Post** /accounts/{id}/disable | Disable Account
[**EnableAccount**](AccountsAPI.md#EnableAccount) | **Post** /accounts/{id}/enable | Enable Account
[**GetAccount**](AccountsAPI.md#GetAccount) | **Get** /accounts/{id} | Account Details
[**GetAccountEntitlements**](AccountsAPI.md#GetAccountEntitlements) | **Get** /accounts/{id}/entitlements | Account Entitlements
[**ListAccounts**](AccountsAPI.md#ListAccounts) | **Get** /accounts | Accounts List
[**PutAccount**](AccountsAPI.md#PutAccount) | **Put** /accounts/{id} | Update Account
[**ReloadAccount**](AccountsAPI.md#ReloadAccount) | **Post** /accounts/{id}/reload | Reload Account
[**UnlockAccount**](AccountsAPI.md#UnlockAccount) | **Post** /accounts/{id}/unlock | Unlock Account
[**UpdateAccount**](AccountsAPI.md#UpdateAccount) | **Patch** /accounts/{id} | Update Account



## Create Account

> AccountsAsyncResult CreateAccount(ctx).AccountAttributesCreate(accountAttributesCreate).Execute()

This API submits an account creation task and returns the task ID.  
The `sourceId` where this account will be created must be included in the `attributes` object.
A token with ORG_ADMIN authority is required to call this API.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    accountAttributesCreate := *sailpoint.NewAccountAttributesCreate(*sailpoint.NewAccountAttributesCreateAttributes("34bfcbe116c9407464af37acbaf7a4dc")) // AccountAttributesCreate | 

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.AccountsAPI.CreateAccount(context.Background()).AccountAttributesCreate(accountAttributesCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.CreateAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAccount`: AccountsAsyncResult
    fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.CreateAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountAttributesCreate** | [**AccountAttributesCreate**](AccountAttributesCreate.md) |  | 

### Return type

[**AccountsAsyncResult**](AccountsAsyncResult.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete Account

> AccountsAsyncResult DeleteAccount(ctx, id).Execute()

Use this API to delete an account. 
This endpoint submits an account delete task and returns the task ID. 
A token with ORG_ADMIN authority is required to call this API.
>**NOTE: You can only delete accounts from sources of the "DelimitedFile" type.**

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    id := "ef38f94347e94562b5bb8424a56397d8" // string | Account ID.

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.AccountsAPI.DeleteAccount(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.DeleteAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAccount`: AccountsAsyncResult
    fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.DeleteAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Account ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccountsAsyncResult**](AccountsAsyncResult.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Disable Account

> AccountsAsyncResult DisableAccount(ctx, id).AccountToggleRequest(accountToggleRequest).Execute()

This API submits a task to disable the account and returns the task ID.  
A token with ORG_ADMIN authority is required to call this API.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The account id
    accountToggleRequest := *sailpoint.NewAccountToggleRequest() // AccountToggleRequest | 

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.AccountsAPI.DisableAccount(context.Background(), id).AccountToggleRequest(accountToggleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.DisableAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DisableAccount`: AccountsAsyncResult
    fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.DisableAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The account id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountToggleRequest** | [**AccountToggleRequest**](AccountToggleRequest.md) |  | 

### Return type

[**AccountsAsyncResult**](AccountsAsyncResult.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Enable Account

> AccountsAsyncResult EnableAccount(ctx, id).AccountToggleRequest(accountToggleRequest).Execute()

This API submits a task to enable account and returns the task ID.  
A token with ORG_ADMIN authority is required to call this API.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The account id
    accountToggleRequest := *sailpoint.NewAccountToggleRequest() // AccountToggleRequest | 

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.AccountsAPI.EnableAccount(context.Background(), id).AccountToggleRequest(accountToggleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.EnableAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableAccount`: AccountsAsyncResult
    fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.EnableAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The account id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountToggleRequest** | [**AccountToggleRequest**](AccountToggleRequest.md) |  | 

### Return type

[**AccountsAsyncResult**](AccountsAsyncResult.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Account Details

> Account GetAccount(ctx, id).Execute()

Use this API to return the details for a single account by its ID.  
A token with ORG_ADMIN authority is required to call this API.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    id := "ef38f94347e94562b5bb8424a56397d8" // string | Account ID.

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.AccountsAPI.GetAccount(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccount`: Account
    fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Account ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Account**](Account.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Account Entitlements

> []EntitlementDto GetAccountEntitlements(ctx, id).Limit(limit).Offset(offset).Count(count).Execute()

This API returns entitlements of the account.  
A token with ORG_ADMIN authority is required to call this API.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The account id
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.AccountsAPI.GetAccountEntitlements(context.Background(), id).Limit(limit).Offset(offset).Count(count).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetAccountEntitlements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountEntitlements`: []EntitlementDto
    fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetAccountEntitlements`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The account id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountEntitlementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]

### Return type

[**[]EntitlementDto**](EntitlementDto.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Accounts List

> []Account ListAccounts(ctx).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).Execute()

This returns a list of accounts.  
A token with ORG_ADMIN authority is required to call this API.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
    filters := "identityId eq "2c9180858082150f0180893dbaf44201"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq, in, sw*  **identityId**: *eq, in, sw*  **name**: *eq, in, sw*  **nativeIdentity**: *eq, in, sw*  **sourceId**: *eq, in, sw*  **uncorrelated**: *eq*  **identity.name**: *eq, in, sw* (optional)
    sorters := "id,name" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **id, name, created, modified, sourceId, identityId, nativeIdentity, uuid, manuallyCorrelated, identity.name** (optional)

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.AccountsAPI.ListAccounts(context.Background()).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.ListAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAccounts`: []Account
    fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.ListAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq, in, sw*  **identityId**: *eq, in, sw*  **name**: *eq, in, sw*  **nativeIdentity**: *eq, in, sw*  **sourceId**: *eq, in, sw*  **uncorrelated**: *eq*  **identity.name**: *eq, in, sw* | 
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **id, name, created, modified, sourceId, identityId, nativeIdentity, uuid, manuallyCorrelated, identity.name** | 

### Return type

[**[]Account**](Account.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update Account

> AccountsAsyncResult PutAccount(ctx, id).AccountAttributes(accountAttributes).Execute()

Use this API to update an account with a PUT request. 
This endpoint submits an account update task and returns the task ID. 
A token with ORG_ADMIN authority is required to call this API.
>**NOTE: You can only use this PUT endpoint to update accounts from sources of the "DelimitedFile" type.**

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    id := "ef38f94347e94562b5bb8424a56397d8" // string | Account ID.
    accountAttributes := *sailpoint.NewAccountAttributes(map[string]interface{}{"key": interface{}(123)}) // AccountAttributes | 

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.AccountsAPI.PutAccount(context.Background(), id).AccountAttributes(accountAttributes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.PutAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutAccount`: AccountsAsyncResult
    fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.PutAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Account ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountAttributes** | [**AccountAttributes**](AccountAttributes.md) |  | 

### Return type

[**AccountsAsyncResult**](AccountsAsyncResult.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Reload Account

> AccountsAsyncResult ReloadAccount(ctx, id).Execute()

This API asynchronously reloads the account directly from the connector and performs a one-time aggregation process.  
A token with ORG_ADMIN authority is required to call this API.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The account id

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.AccountsAPI.ReloadAccount(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.ReloadAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReloadAccount`: AccountsAsyncResult
    fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.ReloadAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The account id | 

### Other Parameters

Other parameters are passed through a pointer to a apiReloadAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccountsAsyncResult**](AccountsAsyncResult.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Unlock Account

> AccountsAsyncResult UnlockAccount(ctx, id).AccountUnlockRequest(accountUnlockRequest).Execute()

This API submits a task to unlock an account and returns the task ID.  
A token with ORG_ADMIN authority is required to call this API.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The account id
    accountUnlockRequest := *sailpoint.NewAccountUnlockRequest() // AccountUnlockRequest | 

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.AccountsAPI.UnlockAccount(context.Background(), id).AccountUnlockRequest(accountUnlockRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.UnlockAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnlockAccount`: AccountsAsyncResult
    fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.UnlockAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The account id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnlockAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountUnlockRequest** | [**AccountUnlockRequest**](AccountUnlockRequest.md) |  | 

### Return type

[**AccountsAsyncResult**](AccountsAsyncResult.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update Account

> map[string]interface{} UpdateAccount(ctx, id).JsonPatchOperation(jsonPatchOperation).Execute()

Use this API to update the account with a PATCH request.
This endpoint can only modify these fields:
* `identityId`
* `manuallyCorrelated`
The request must provide a JSONPatch payload.
A token with ORG_ADMIN authority is required to call this API.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    id := "ef38f94347e94562b5bb8424a56397d8" // string | Account ID.
    jsonPatchOperation := []sailpoint.JsonPatchOperation{*sailpoint.NewJsonPatchOperation("replace", "/description")} // []JsonPatchOperation | A list of account update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard.

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.AccountsAPI.UpdateAccount(context.Background(), id).JsonPatchOperation(jsonPatchOperation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.UpdateAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAccount`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.UpdateAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Account ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jsonPatchOperation** | [**[]JsonPatchOperation**](JsonPatchOperation.md) | A list of account update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard. | 

### Return type

**map[string]interface{}**

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

