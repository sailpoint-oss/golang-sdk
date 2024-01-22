# \ScheduledSearchAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateScheduledSearch**](ScheduledSearchAPI.md#CreateScheduledSearch) | **Post** /scheduled-searches | Create a new scheduled search
[**DeleteScheduledSearch**](ScheduledSearchAPI.md#DeleteScheduledSearch) | **Delete** /scheduled-searches/{id} | Delete a Scheduled Search
[**GetScheduledSearch**](ScheduledSearchAPI.md#GetScheduledSearch) | **Get** /scheduled-searches/{id} | Get a Scheduled Search
[**ListScheduledSearch**](ScheduledSearchAPI.md#ListScheduledSearch) | **Get** /scheduled-searches | List scheduled searches
[**UnsubscribeScheduledSearch**](ScheduledSearchAPI.md#UnsubscribeScheduledSearch) | **Post** /scheduled-searches/{id}/unsubscribe | Unsubscribe a recipient from Scheduled Search
[**UpdateScheduledSearch**](ScheduledSearchAPI.md#UpdateScheduledSearch) | **Put** /scheduled-searches/{id} | Update an existing Scheduled Search



## CreateScheduledSearch

> ScheduledSearch CreateScheduledSearch(ctx).CreateScheduledSearchRequest(createScheduledSearchRequest).Execute()

Create a new scheduled search



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
    createScheduledSearchRequest := *openapiclient.NewCreateScheduledSearchRequest("554f1511-f0a1-4744-ab14-599514d3e57c", *openapiclient.NewSchedule1(openapiclient.ScheduleType("DAILY"), *openapiclient.NewSchedule1Hours(openapiclient.SelectorType("LIST"), []string{"Values_example"})), []openapiclient.SearchScheduleRecipientsInner{*openapiclient.NewSearchScheduleRecipientsInner("IDENTITY", "2c9180867624cbd7017642d8c8c81f67")}) // CreateScheduledSearchRequest | The scheduled search to persist.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScheduledSearchAPI.CreateScheduledSearch(context.Background()).CreateScheduledSearchRequest(createScheduledSearchRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduledSearchAPI.CreateScheduledSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateScheduledSearch`: ScheduledSearch
    fmt.Fprintf(os.Stdout, "Response from `ScheduledSearchAPI.CreateScheduledSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateScheduledSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createScheduledSearchRequest** | [**CreateScheduledSearchRequest**](CreateScheduledSearchRequest.md) | The scheduled search to persist. | 

### Return type

[**ScheduledSearch**](ScheduledSearch.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteScheduledSearch

> DeleteScheduledSearch(ctx, id).Execute()

Delete a Scheduled Search



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
    id := "2c91808568c529c60168cca6f90c1313" // string | ID of the requested document.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ScheduledSearchAPI.DeleteScheduledSearch(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduledSearchAPI.DeleteScheduledSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the requested document. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteScheduledSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetScheduledSearch

> ScheduledSearch GetScheduledSearch(ctx, id).Execute()

Get a Scheduled Search



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
    id := "2c91808568c529c60168cca6f90c1313" // string | ID of the requested document.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScheduledSearchAPI.GetScheduledSearch(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduledSearchAPI.GetScheduledSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetScheduledSearch`: ScheduledSearch
    fmt.Fprintf(os.Stdout, "Response from `ScheduledSearchAPI.GetScheduledSearch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the requested document. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetScheduledSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ScheduledSearch**](ScheduledSearch.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListScheduledSearch

> []ScheduledSearch ListScheduledSearch(ctx).Offset(offset).Limit(limit).Count(count).Filters(filters).Execute()

List scheduled searches



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
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
    filters := "savedSearchId eq "6cc0945d-9eeb-4948-9033-72d066e1153e"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **owner.id**: *eq*  **savedSearchId**: *eq* (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScheduledSearchAPI.ListScheduledSearch(context.Background()).Offset(offset).Limit(limit).Count(count).Filters(filters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduledSearchAPI.ListScheduledSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListScheduledSearch`: []ScheduledSearch
    fmt.Fprintf(os.Stdout, "Response from `ScheduledSearchAPI.ListScheduledSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListScheduledSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **owner.id**: *eq*  **savedSearchId**: *eq* | 

### Return type

[**[]ScheduledSearch**](ScheduledSearch.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnsubscribeScheduledSearch

> UnsubscribeScheduledSearch(ctx, id).TypedReference(typedReference).Execute()

Unsubscribe a recipient from Scheduled Search



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
    id := "2c91808568c529c60168cca6f90c1313" // string | ID of the requested document.
    typedReference := *openapiclient.NewTypedReference(openapiclient.DtoType("ACCOUNT_CORRELATION_CONFIG"), "2c91808568c529c60168cca6f90c1313") // TypedReference | The recipient to be removed from the scheduled search. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ScheduledSearchAPI.UnsubscribeScheduledSearch(context.Background(), id).TypedReference(typedReference).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduledSearchAPI.UnsubscribeScheduledSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the requested document. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnsubscribeScheduledSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **typedReference** | [**TypedReference**](TypedReference.md) | The recipient to be removed from the scheduled search.  | 

### Return type

 (empty response body)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateScheduledSearch

> ScheduledSearch UpdateScheduledSearch(ctx, id).ScheduledSearch(scheduledSearch).Execute()

Update an existing Scheduled Search



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
    id := "2c91808568c529c60168cca6f90c1313" // string | ID of the requested document.
    scheduledSearch := *openapiclient.NewScheduledSearch("554f1511-f0a1-4744-ab14-599514d3e57c", *openapiclient.NewSchedule1(openapiclient.ScheduleType("DAILY"), *openapiclient.NewSchedule1Hours(openapiclient.SelectorType("LIST"), []string{"Values_example"})), []openapiclient.SearchScheduleRecipientsInner{*openapiclient.NewSearchScheduleRecipientsInner("IDENTITY", "2c9180867624cbd7017642d8c8c81f67")}, "0de46054-fe90-434a-b84e-c6b3359d0c64", *openapiclient.NewScheduledSearchAllOfOwner("IDENTITY", "2c9180867624cbd7017642d8c8c81f67"), "2c9180867624cbd7017642d8c8c81f67") // ScheduledSearch | The scheduled search to persist.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScheduledSearchAPI.UpdateScheduledSearch(context.Background(), id).ScheduledSearch(scheduledSearch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduledSearchAPI.UpdateScheduledSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateScheduledSearch`: ScheduledSearch
    fmt.Fprintf(os.Stdout, "Response from `ScheduledSearchAPI.UpdateScheduledSearch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the requested document. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateScheduledSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scheduledSearch** | [**ScheduledSearch**](ScheduledSearch.md) | The scheduled search to persist. | 

### Return type

[**ScheduledSearch**](ScheduledSearch.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

