# \ScheduledSearchApi

All URIs are relative to *https://sailpoint.api.identitynow.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ScheduledSearchCreate**](ScheduledSearchApi.md#ScheduledSearchCreate) | **Post** /scheduled-searches | Create a new scheduled search
[**ScheduledSearchDelete**](ScheduledSearchApi.md#ScheduledSearchDelete) | **Delete** /scheduled-searches/{id} | Delete a Scheduled Search by ID
[**ScheduledSearchGet**](ScheduledSearchApi.md#ScheduledSearchGet) | **Get** /scheduled-searches/{id} | Return a Scheduled Search by ID
[**ScheduledSearchList**](ScheduledSearchApi.md#ScheduledSearchList) | **Get** /scheduled-searches | Return a list of scheduled searches
[**ScheduledSearchUnsubscribe**](ScheduledSearchApi.md#ScheduledSearchUnsubscribe) | **Post** /scheduled-searches/{id}/unsubscribe | Unsubscribe a recipient from Scheduled Search
[**ScheduledSearchUpdate**](ScheduledSearchApi.md#ScheduledSearchUpdate) | **Put** /scheduled-searches/{id} | Update an existing Scheduled Search



## ScheduledSearchCreate

> ScheduledSearch ScheduledSearchCreate(ctx).ScheduledSearchCreateRequest(scheduledSearchCreateRequest).Execute()

Create a new scheduled search



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
    scheduledSearchCreateRequest := *openapiclient.NewScheduledSearchCreateRequest("554f1511-f0a1-4744-ab14-599514d3e57c", *openapiclient.NewSchedule1(openapiclient.ScheduleType("DAILY"), "TODO"), []openapiclient.TypedReference{*openapiclient.NewTypedReference(openapiclient.DtoType("ACCOUNT_CORRELATION_CONFIG"), "2c91808568c529c60168cca6f90c1313")}) // ScheduledSearchCreateRequest | The scheduled search to persist.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScheduledSearchApi.ScheduledSearchCreate(context.Background()).ScheduledSearchCreateRequest(scheduledSearchCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduledSearchApi.ScheduledSearchCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ScheduledSearchCreate`: ScheduledSearch
    fmt.Fprintf(os.Stdout, "Response from `ScheduledSearchApi.ScheduledSearchCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiScheduledSearchCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scheduledSearchCreateRequest** | [**ScheduledSearchCreateRequest**](ScheduledSearchCreateRequest.md) | The scheduled search to persist. | 

### Return type

[**ScheduledSearch**](ScheduledSearch.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScheduledSearchDelete

> ScheduledSearchDelete(ctx, id).Execute()

Delete a Scheduled Search by ID



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
    id := "2c91808568c529c60168cca6f90c1313" // string | ID of the requested document.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScheduledSearchApi.ScheduledSearchDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduledSearchApi.ScheduledSearchDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiScheduledSearchDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScheduledSearchGet

> ScheduledSearch ScheduledSearchGet(ctx, id).Execute()

Return a Scheduled Search by ID



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
    id := "2c91808568c529c60168cca6f90c1313" // string | ID of the requested document.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScheduledSearchApi.ScheduledSearchGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduledSearchApi.ScheduledSearchGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ScheduledSearchGet`: ScheduledSearch
    fmt.Fprintf(os.Stdout, "Response from `ScheduledSearchApi.ScheduledSearchGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the requested document. | 

### Other Parameters

Other parameters are passed through a pointer to a apiScheduledSearchGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ScheduledSearch**](ScheduledSearch.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScheduledSearchList

> []ScheduledSearch ScheduledSearchList(ctx).Offset(offset).Limit(limit).Count(count).Filters(filters).Execute()

Return a list of scheduled searches



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
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
    filters := "savedSearchId eq "6cc0945d-9eeb-4948-9033-72d066e1153e"" // string | An expression used to constrain the result set using the filtering syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results).  Allowed filter properties: *owner.id*, *savedSearchId*  Allowed filter operator: *eq*  **Example filters**:  ```owner.id eq \"0de46054-fe90-434a-b84e-c6b3359d0c64\"``` -- returns scheduled searches for the specified owner ID  ```savedSearchId eq \"6cc0945d-9eeb-4948-9033-72d066e1153e\"``` -- returns scheduled searches that reference the specified saved search  ```owner.id eq me or savedSearchId eq \"6cc0945d-9eeb-4948-9033-72d066e1153e\"``` -- returns all of the current user's scheduled searches as well as all scheduled searches that reference the specified saved search  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScheduledSearchApi.ScheduledSearchList(context.Background()).Offset(offset).Limit(limit).Count(count).Filters(filters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduledSearchApi.ScheduledSearchList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ScheduledSearchList`: []ScheduledSearch
    fmt.Fprintf(os.Stdout, "Response from `ScheduledSearchApi.ScheduledSearchList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiScheduledSearchListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **filters** | **string** | An expression used to constrain the result set using the filtering syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results).  Allowed filter properties: *owner.id*, *savedSearchId*  Allowed filter operator: *eq*  **Example filters**:  &#x60;&#x60;&#x60;owner.id eq \&quot;0de46054-fe90-434a-b84e-c6b3359d0c64\&quot;&#x60;&#x60;&#x60; -- returns scheduled searches for the specified owner ID  &#x60;&#x60;&#x60;savedSearchId eq \&quot;6cc0945d-9eeb-4948-9033-72d066e1153e\&quot;&#x60;&#x60;&#x60; -- returns scheduled searches that reference the specified saved search  &#x60;&#x60;&#x60;owner.id eq me or savedSearchId eq \&quot;6cc0945d-9eeb-4948-9033-72d066e1153e\&quot;&#x60;&#x60;&#x60; -- returns all of the current user&#39;s scheduled searches as well as all scheduled searches that reference the specified saved search  | 

### Return type

[**[]ScheduledSearch**](ScheduledSearch.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScheduledSearchUnsubscribe

> ScheduledSearchUnsubscribe(ctx, id).TypedReference(typedReference).Execute()

Unsubscribe a recipient from Scheduled Search



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
    id := "2c91808568c529c60168cca6f90c1313" // string | ID of the requested document.
    typedReference := *openapiclient.NewTypedReference(openapiclient.DtoType("ACCOUNT_CORRELATION_CONFIG"), "2c91808568c529c60168cca6f90c1313") // TypedReference | The recipient to be removed from the scheduled search. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScheduledSearchApi.ScheduledSearchUnsubscribe(context.Background(), id).TypedReference(typedReference).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduledSearchApi.ScheduledSearchUnsubscribe``: %v\n", err)
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

Other parameters are passed through a pointer to a apiScheduledSearchUnsubscribeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **typedReference** | [**TypedReference**](TypedReference.md) | The recipient to be removed from the scheduled search.  | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScheduledSearchUpdate

> ScheduledSearch ScheduledSearchUpdate(ctx, id).ScheduledSearch(scheduledSearch).Execute()

Update an existing Scheduled Search



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
    id := "2c91808568c529c60168cca6f90c1313" // string | ID of the requested document.
    scheduledSearch := *openapiclient.NewScheduledSearch("554f1511-f0a1-4744-ab14-599514d3e57c", *openapiclient.NewSchedule1(openapiclient.ScheduleType("DAILY"), "TODO"), []openapiclient.TypedReference{*openapiclient.NewTypedReference(openapiclient.DtoType("ACCOUNT_CORRELATION_CONFIG"), "2c91808568c529c60168cca6f90c1313")}) // ScheduledSearch | The scheduled search to persist.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScheduledSearchApi.ScheduledSearchUpdate(context.Background(), id).ScheduledSearch(scheduledSearch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduledSearchApi.ScheduledSearchUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ScheduledSearchUpdate`: ScheduledSearch
    fmt.Fprintf(os.Stdout, "Response from `ScheduledSearchApi.ScheduledSearchUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the requested document. | 

### Other Parameters

Other parameters are passed through a pointer to a apiScheduledSearchUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scheduledSearch** | [**ScheduledSearch**](ScheduledSearch.md) | The scheduled search to persist. | 

### Return type

[**ScheduledSearch**](ScheduledSearch.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

