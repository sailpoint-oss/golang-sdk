# \AccountActivitiesApi

All URIs are relative to *https://sailpoint.api.identitynow.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccountActivity**](AccountActivitiesApi.md#GetAccountActivity) | **Get** /account-activities/{id} | Get an Account Activity
[**ListAccountActivities**](AccountActivitiesApi.md#ListAccountActivities) | **Get** /account-activities | List Account Activities



## GetAccountActivity

> AccountActivity GetAccountActivity(ctx, id).Execute()

Get an Account Activity



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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The account activity id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountActivitiesApi.GetAccountActivity(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountActivitiesApi.GetAccountActivity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountActivity`: AccountActivity
    fmt.Fprintf(os.Stdout, "Response from `AccountActivitiesApi.GetAccountActivity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The account activity id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountActivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccountActivity**](AccountActivity.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccountActivities

> []AccountActivity ListAccountActivities(ctx).RequestedFor(requestedFor).RequestedBy(requestedBy).RegardingIdentity(regardingIdentity).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).Execute()

List Account Activities



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
    requestedFor := "2c91808568c529c60168cca6f90c1313" // string | The identity that the activity was requested for. *me* indicates the current user. Mutually exclusive with *regarding-identity*. (optional)
    requestedBy := "2c91808568c529c60168cca6f90c1313" // string | The identity that requested the activity. *me* indicates the current user. Mutually exclusive with *regarding-identity*. (optional)
    regardingIdentity := "2c91808568c529c60168cca6f90c1313" // string | The specified identity will be either the requester or target of the account activity. *me* indicates the current user. Mutually exclusive with *requested-for* and *requested-by*. (optional)
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
    filters := "type eq "Identity Refresh"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results).  Filtering is supported for the following fields and operators:  **type**: *eq, in* (See the `type` property in the response schema for possible values)  **created**: *gt, lt, ge, le*  **modified**: *gt, lt, ge, le* (optional)
    sorters := "created" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results) Sorting is supported for the following fields: **type, created, modified** (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountActivitiesApi.ListAccountActivities(context.Background()).RequestedFor(requestedFor).RequestedBy(requestedBy).RegardingIdentity(regardingIdentity).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountActivitiesApi.ListAccountActivities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAccountActivities`: []AccountActivity
    fmt.Fprintf(os.Stdout, "Response from `AccountActivitiesApi.ListAccountActivities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAccountActivitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestedFor** | **string** | The identity that the activity was requested for. *me* indicates the current user. Mutually exclusive with *regarding-identity*. | 
 **requestedBy** | **string** | The identity that requested the activity. *me* indicates the current user. Mutually exclusive with *regarding-identity*. | 
 **regardingIdentity** | **string** | The specified identity will be either the requester or target of the account activity. *me* indicates the current user. Mutually exclusive with *requested-for* and *requested-by*. | 
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results).  Filtering is supported for the following fields and operators:  **type**: *eq, in* (See the &#x60;type&#x60; property in the response schema for possible values)  **created**: *gt, lt, ge, le*  **modified**: *gt, lt, ge, le* | 
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results) Sorting is supported for the following fields: **type, created, modified** | 

### Return type

[**[]AccountActivity**](AccountActivity.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

