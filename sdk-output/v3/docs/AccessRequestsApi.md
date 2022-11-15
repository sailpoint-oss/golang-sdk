# \AccessRequestsApi

All URIs are relative to *https://sailpoint.api.identitynow.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelAccessRequest**](AccessRequestsApi.md#CancelAccessRequest) | **Post** /access-requests/cancel | Cancel Access Request
[**CreateAccessRequest**](AccessRequestsApi.md#CreateAccessRequest) | **Post** /access-requests | Submit an Access Request
[**GetAccessRequestConfig**](AccessRequestsApi.md#GetAccessRequestConfig) | **Get** /access-request-config | Get Access Request Configuration
[**ListAccessRequestStatus**](AccessRequestsApi.md#ListAccessRequestStatus) | **Get** /access-request-status | Access Request Status
[**UpdateAccessRequestConfig**](AccessRequestsApi.md#UpdateAccessRequestConfig) | **Put** /access-request-config | Update Access Request Configuration



## CancelAccessRequest

> map[string]interface{} CancelAccessRequest(ctx).CancelAccessRequest(cancelAccessRequest).Execute()

Cancel Access Request



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
    cancelAccessRequest := *openapiclient.NewCancelAccessRequest("2c9180835d2e5168015d32f890ca1581", "I requested this role by mistake.") // CancelAccessRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessRequestsApi.CancelAccessRequest(context.Background()).CancelAccessRequest(cancelAccessRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessRequestsApi.CancelAccessRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelAccessRequest`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AccessRequestsApi.CancelAccessRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCancelAccessRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cancelAccessRequest** | [**CancelAccessRequest**](CancelAccessRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAccessRequest

> map[string]interface{} CreateAccessRequest(ctx).AccessRequest(accessRequest).Execute()

Submit an Access Request



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
    accessRequest := *openapiclient.NewAccessRequest([]string{"2c918084660f45d6016617daa9210584"}, []openapiclient.AccessRequestItem{*openapiclient.NewAccessRequestItem("ACCESS_PROFILE", "2c9180835d2e5168015d32f890ca1581")}) // AccessRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessRequestsApi.CreateAccessRequest(context.Background()).AccessRequest(accessRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessRequestsApi.CreateAccessRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAccessRequest`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AccessRequestsApi.CreateAccessRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccessRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessRequest** | [**AccessRequest**](AccessRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccessRequestConfig

> AccessRequestConfig GetAccessRequestConfig(ctx).Execute()

Get Access Request Configuration



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessRequestsApi.GetAccessRequestConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessRequestsApi.GetAccessRequestConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccessRequestConfig`: AccessRequestConfig
    fmt.Fprintf(os.Stdout, "Response from `AccessRequestsApi.GetAccessRequestConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessRequestConfigRequest struct via the builder pattern


### Return type

[**AccessRequestConfig**](AccessRequestConfig.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccessRequestStatus

> []RequestedItemStatus ListAccessRequestStatus(ctx).RequestedFor(requestedFor).RequestedBy(requestedBy).RegardingIdentity(regardingIdentity).Count(count).Limit(limit).Offset(offset).Filters(filters).Sorters(sorters).Execute()

Access Request Status



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
    requestedFor := "2c9180877b2b6ea4017b2c545f971429" // string | Filter the results by the identity for which the requests were made. *me* indicates the current user. Mutually exclusive with *regarding-identity*. (optional)
    requestedBy := "2c9180877b2b6ea4017b2c545f971429" // string | Filter the results by the identity that made the requests. *me* indicates the current user. Mutually exclusive with *regarding-identity*. (optional)
    regardingIdentity := "2c9180877b2b6ea4017b2c545f971429" // string | Filter the results by the specified identity which is either the requester or target of the requests. *me* indicates the current user. Mutually exclusive with *requested-for* and *requested-by*. (optional)
    count := false // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored. (optional) (default to false)
    limit := int32(100) // int32 | Max number of results to return. (optional) (default to 250)
    offset := int32(10) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. Defaults to 0 if not specified. (optional)
    filters := "accountActivityItemId eq "2c918086771c86df0177401efcdf54c0"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **accountActivityItemId**: *eq, in* (optional)
    sorters := "created" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **created, modified, accountActivityItemId** (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessRequestsApi.ListAccessRequestStatus(context.Background()).RequestedFor(requestedFor).RequestedBy(requestedBy).RegardingIdentity(regardingIdentity).Count(count).Limit(limit).Offset(offset).Filters(filters).Sorters(sorters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessRequestsApi.ListAccessRequestStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAccessRequestStatus`: []RequestedItemStatus
    fmt.Fprintf(os.Stdout, "Response from `AccessRequestsApi.ListAccessRequestStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAccessRequestStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestedFor** | **string** | Filter the results by the identity for which the requests were made. *me* indicates the current user. Mutually exclusive with *regarding-identity*. | 
 **requestedBy** | **string** | Filter the results by the identity that made the requests. *me* indicates the current user. Mutually exclusive with *regarding-identity*. | 
 **regardingIdentity** | **string** | Filter the results by the specified identity which is either the requester or target of the requests. *me* indicates the current user. Mutually exclusive with *requested-for* and *requested-by*. | 
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored. | [default to false]
 **limit** | **int32** | Max number of results to return. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. Defaults to 0 if not specified. | 
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **accountActivityItemId**: *eq, in* | 
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **created, modified, accountActivityItemId** | 

### Return type

[**[]RequestedItemStatus**](RequestedItemStatus.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccessRequestConfig

> AccessRequestConfig UpdateAccessRequestConfig(ctx).AccessRequestConfig(accessRequestConfig).Execute()

Update Access Request Configuration



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
    accessRequestConfig := *openapiclient.NewAccessRequestConfig() // AccessRequestConfig | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessRequestsApi.UpdateAccessRequestConfig(context.Background()).AccessRequestConfig(accessRequestConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessRequestsApi.UpdateAccessRequestConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAccessRequestConfig`: AccessRequestConfig
    fmt.Fprintf(os.Stdout, "Response from `AccessRequestsApi.UpdateAccessRequestConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccessRequestConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessRequestConfig** | [**AccessRequestConfig**](AccessRequestConfig.md) |  | 

### Return type

[**AccessRequestConfig**](AccessRequestConfig.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

