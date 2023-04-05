# \AccountsApi

All URIs are relative to *https://sailpoint.api.identitynow.com/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccount**](AccountsApi.md#CreateAccount) | **Post** /accounts | Create Account
[**DeleteAccount**](AccountsApi.md#DeleteAccount) | **Delete** /accounts/{id} | Delete Account
[**DisableAccount**](AccountsApi.md#DisableAccount) | **Post** /accounts/{id}/disable | Disable Account
[**DisableAccountForIdentity**](AccountsApi.md#DisableAccountForIdentity) | **Post** /identities-accounts/{id}/disable | Disable IDN Account for Identity
[**DisableAccountsForIdentities**](AccountsApi.md#DisableAccountsForIdentities) | **Post** /identities-accounts/disable | Disable IDN Accounts for Identities
[**EnableAccount**](AccountsApi.md#EnableAccount) | **Post** /accounts/{id}/enable | Enable Account
[**EnableAccountForIdentity**](AccountsApi.md#EnableAccountForIdentity) | **Post** /identities-accounts/{id}/enable | Enable IDN Account for Identity
[**EnableAccountsForIdentities**](AccountsApi.md#EnableAccountsForIdentities) | **Post** /identities-accounts/enable | Enable IDN Accounts for Identities
[**GetAccount**](AccountsApi.md#GetAccount) | **Get** /accounts/{id} | Account Details
[**GetAccountEntitlements**](AccountsApi.md#GetAccountEntitlements) | **Get** /accounts/{id}/entitlements | Account Entitlements
[**ListAccounts**](AccountsApi.md#ListAccounts) | **Get** /accounts | Accounts List
[**PutAccount**](AccountsApi.md#PutAccount) | **Put** /accounts/{id} | Update Account
[**ReloadAccount**](AccountsApi.md#ReloadAccount) | **Post** /accounts/{id}/reload | Reload Account
[**UnlockAccount**](AccountsApi.md#UnlockAccount) | **Post** /accounts/{id}/unlock | Unlock Account
[**UpdateAccount**](AccountsApi.md#UpdateAccount) | **Patch** /accounts/{id} | Update Account



## CreateAccount

> AccountsAsyncResult CreateAccount(ctx).AccountAttributes(accountAttributes).Execute()

Create Account



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
    accountAttributes := *openapiclient.NewAccountAttributes(map[string]interface{}({city=Austin, displayName=John Doe, userName=jdoe, sAMAccountName=jDoe, mail=john.doe@sailpoint.com})) // AccountAttributes | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.CreateAccount(context.Background()).AccountAttributes(accountAttributes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.CreateAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAccount`: AccountsAsyncResult
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.CreateAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountAttributes** | [**AccountAttributes**](AccountAttributes.md) |  | 

### Return type

[**AccountsAsyncResult**](AccountsAsyncResult.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAccount

> AccountsAsyncResult DeleteAccount(ctx, id).Execute()

Delete Account



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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The account ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.DeleteAccount(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.DeleteAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAccount`: AccountsAsyncResult
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.DeleteAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccountsAsyncResult**](AccountsAsyncResult.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisableAccount

> AccountsAsyncResult DisableAccount(ctx, id).AccountToggleRequest(accountToggleRequest).Execute()

Disable Account



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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The account id
    accountToggleRequest := *openapiclient.NewAccountToggleRequest() // AccountToggleRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.DisableAccount(context.Background(), id).AccountToggleRequest(accountToggleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.DisableAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DisableAccount`: AccountsAsyncResult
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.DisableAccount`: %v\n", resp)
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

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisableAccountForIdentity

> map[string]interface{} DisableAccountForIdentity(ctx, id).Execute()

Disable IDN Account for Identity



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
    id := "2c91808384203c2d018437e631158309" // string | The identity id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.DisableAccountForIdentity(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.DisableAccountForIdentity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DisableAccountForIdentity`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.DisableAccountForIdentity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The identity id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableAccountForIdentityRequest struct via the builder pattern


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


## DisableAccountsForIdentities

> []BulkIdentitiesAccountsResponse DisableAccountsForIdentities(ctx).IdentitiesAccountsBulkRequest(identitiesAccountsBulkRequest).Execute()

Disable IDN Accounts for Identities



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
    identitiesAccountsBulkRequest := *openapiclient.NewIdentitiesAccountsBulkRequest() // IdentitiesAccountsBulkRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.DisableAccountsForIdentities(context.Background()).IdentitiesAccountsBulkRequest(identitiesAccountsBulkRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.DisableAccountsForIdentities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DisableAccountsForIdentities`: []BulkIdentitiesAccountsResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.DisableAccountsForIdentities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDisableAccountsForIdentitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identitiesAccountsBulkRequest** | [**IdentitiesAccountsBulkRequest**](IdentitiesAccountsBulkRequest.md) |  | 

### Return type

[**[]BulkIdentitiesAccountsResponse**](BulkIdentitiesAccountsResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableAccount

> AccountsAsyncResult EnableAccount(ctx, id).AccountToggleRequest(accountToggleRequest).Execute()

Enable Account



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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The account id
    accountToggleRequest := *openapiclient.NewAccountToggleRequest() // AccountToggleRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.EnableAccount(context.Background(), id).AccountToggleRequest(accountToggleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.EnableAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableAccount`: AccountsAsyncResult
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.EnableAccount`: %v\n", resp)
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

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableAccountForIdentity

> map[string]interface{} EnableAccountForIdentity(ctx, id).Execute()

Enable IDN Account for Identity



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
    id := "2c91808384203c2d018437e631158309" // string | The identity id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.EnableAccountForIdentity(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.EnableAccountForIdentity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableAccountForIdentity`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.EnableAccountForIdentity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The identity id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableAccountForIdentityRequest struct via the builder pattern


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


## EnableAccountsForIdentities

> []BulkIdentitiesAccountsResponse EnableAccountsForIdentities(ctx).IdentitiesAccountsBulkRequest(identitiesAccountsBulkRequest).Execute()

Enable IDN Accounts for Identities



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
    identitiesAccountsBulkRequest := *openapiclient.NewIdentitiesAccountsBulkRequest() // IdentitiesAccountsBulkRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.EnableAccountsForIdentities(context.Background()).IdentitiesAccountsBulkRequest(identitiesAccountsBulkRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.EnableAccountsForIdentities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableAccountsForIdentities`: []BulkIdentitiesAccountsResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.EnableAccountsForIdentities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnableAccountsForIdentitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identitiesAccountsBulkRequest** | [**IdentitiesAccountsBulkRequest**](IdentitiesAccountsBulkRequest.md) |  | 

### Return type

[**[]BulkIdentitiesAccountsResponse**](BulkIdentitiesAccountsResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccount

> Account GetAccount(ctx, id).Execute()

Account Details



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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The account ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.GetAccount(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.GetAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccount`: Account
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.GetAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Account**](Account.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountEntitlements

> []Entitlement GetAccountEntitlements(ctx, id).Offset(offset).Limit(limit).Count(count).Execute()

Account Entitlements



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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The account id
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.GetAccountEntitlements(context.Background(), id).Offset(offset).Limit(limit).Count(count).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.GetAccountEntitlements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountEntitlements`: []Entitlement
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.GetAccountEntitlements`: %v\n", resp)
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

 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]

### Return type

[**[]Entitlement**](Entitlement.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccounts

> []ListAccounts200ResponseInner ListAccounts(ctx).DetailLevel(detailLevel).Limit(limit).Offset(offset).Count(count).Filters(filters).Execute()

Accounts List



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
    detailLevel := "FULL" // string | Determines whether Slim, or increased level of detail is provided for each account in the returned list. FULL is the default behavior. (optional)
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
    filters := "identityId eq "2c9180858082150f0180893dbaf44201"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq, in*  **identityId**: *eq*  **name**: *eq, in*  **nativeIdentity**: *eq, in*  **sourceId**: *eq, in*  **uncorrelated**: *eq* (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.ListAccounts(context.Background()).DetailLevel(detailLevel).Limit(limit).Offset(offset).Count(count).Filters(filters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.ListAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAccounts`: []ListAccounts200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.ListAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **detailLevel** | **string** | Determines whether Slim, or increased level of detail is provided for each account in the returned list. FULL is the default behavior. | 
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq, in*  **identityId**: *eq*  **name**: *eq, in*  **nativeIdentity**: *eq, in*  **sourceId**: *eq, in*  **uncorrelated**: *eq* | 

### Return type

[**[]ListAccounts200ResponseInner**](ListAccounts200ResponseInner.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAccount

> AccountsAsyncResult PutAccount(ctx, id).AccountAttributes(accountAttributes).Execute()

Update Account



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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The account ID
    accountAttributes := *openapiclient.NewAccountAttributes(map[string]interface{}({city=Austin, displayName=John Doe, userName=jdoe, sAMAccountName=jDoe, mail=john.doe@sailpoint.com})) // AccountAttributes | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.PutAccount(context.Background(), id).AccountAttributes(accountAttributes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.PutAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutAccount`: AccountsAsyncResult
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.PutAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountAttributes** | [**AccountAttributes**](AccountAttributes.md) |  | 

### Return type

[**AccountsAsyncResult**](AccountsAsyncResult.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReloadAccount

> AccountsAsyncResult ReloadAccount(ctx, id).Execute()

Reload Account



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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The account id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.ReloadAccount(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.ReloadAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReloadAccount`: AccountsAsyncResult
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.ReloadAccount`: %v\n", resp)
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

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnlockAccount

> AccountsAsyncResult UnlockAccount(ctx, id).AccountUnlockRequest(accountUnlockRequest).Execute()

Unlock Account



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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The account id
    accountUnlockRequest := *openapiclient.NewAccountUnlockRequest() // AccountUnlockRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.UnlockAccount(context.Background(), id).AccountUnlockRequest(accountUnlockRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.UnlockAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnlockAccount`: AccountsAsyncResult
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.UnlockAccount`: %v\n", resp)
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

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccount

> map[string]interface{} UpdateAccount(ctx, id).RequestBody(requestBody).Execute()

Update Account



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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The account ID
    requestBody := []map[string]interface{}{map[string]interface{}(123)} // []map[string]interface{} | A list of account update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.UpdateAccount(context.Background(), id).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.UpdateAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAccount`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.UpdateAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]map[string]interface{}** | A list of account update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard. | 

### Return type

**map[string]interface{}**

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

