# \IdentityHistoryApi

All URIs are relative to *https://sailpoint.api.identitynow.com/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CompareIdentitySnapshots**](IdentityHistoryApi.md#CompareIdentitySnapshots) | **Get** /historical-identities/{id}/compare | Gets a difference of count for each access item types for the given identity between 2 snapshots
[**CompareIdentitySnapshotsAccessType**](IdentityHistoryApi.md#CompareIdentitySnapshotsAccessType) | **Get** /historical-identities/{id}/compare/{access-type} | Gets a list of differences of specific accessType for the given identity between 2 snapshots
[**GetHistoricalIdentity**](IdentityHistoryApi.md#GetHistoricalIdentity) | **Get** /historical-identities/{id} | Get latest snapshot of identity
[**GetHistoricalIdentityEvents**](IdentityHistoryApi.md#GetHistoricalIdentityEvents) | **Get** /historical-identities/{id}/events | Lists all events for the given identity
[**GetIdentitySnapshot**](IdentityHistoryApi.md#GetIdentitySnapshot) | **Get** /historical-identities/{id}/snapshots/{date} | Gets an identity snapshot at a given date
[**GetIdentitySnapshotSummary**](IdentityHistoryApi.md#GetIdentitySnapshotSummary) | **Get** /historical-identities/{id}/snapshot-summary | Gets the summary for the event count for a specific identity
[**GetIdentityStartDate**](IdentityHistoryApi.md#GetIdentityStartDate) | **Get** /historical-identities/{id}/start-date | Gets the start date of the identity
[**ListHistoricalIdentities**](IdentityHistoryApi.md#ListHistoricalIdentities) | **Get** /historical-identities | Lists all the identities
[**ListIdentityAccessItems**](IdentityHistoryApi.md#ListIdentityAccessItems) | **Get** /historical-identities/{id}/access-items | Gets a list of access items for the identity filtered by item type
[**ListIdentitySnapshotAccessItems**](IdentityHistoryApi.md#ListIdentitySnapshotAccessItems) | **Get** /historical-identities/{id}/snapshots/{date}/access-items | Gets the list of identity access items at a given date filterd by item type
[**ListIdentitySnapshots**](IdentityHistoryApi.md#ListIdentitySnapshots) | **Get** /historical-identities/{id}/snapshots | Lists all the snapshots for the identity



## CompareIdentitySnapshots

> []IdentityCompareResponse CompareIdentitySnapshots(ctx, id).Snapshot1(snapshot1).Snapshot2(snapshot2).AccessItemTypes(accessItemTypes).Limit(limit).Offset(offset).Count(count).Execute()

Gets a difference of count for each access item types for the given identity between 2 snapshots



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
    id := "8c190e6787aa4ed9a90bd9d5344523fb" // string | The identity id
    snapshot1 := "2007-03-01T13:00:00Z" // string | The snapshot 1 of identity (optional)
    snapshot2 := "2008-03-01T13:00:00Z" // string | The snapshot 2 of identity (optional)
    accessItemTypes := []string{"Inner_example"} // []string | An optional list of access item types (app, account, entitlement, etc...) to return.   If null or empty, all access items types are returned  (optional)
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityHistoryApi.CompareIdentitySnapshots(context.Background(), id).Snapshot1(snapshot1).Snapshot2(snapshot2).AccessItemTypes(accessItemTypes).Limit(limit).Offset(offset).Count(count).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityHistoryApi.CompareIdentitySnapshots``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CompareIdentitySnapshots`: []IdentityCompareResponse
    fmt.Fprintf(os.Stdout, "Response from `IdentityHistoryApi.CompareIdentitySnapshots`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The identity id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompareIdentitySnapshotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **snapshot1** | **string** | The snapshot 1 of identity | 
 **snapshot2** | **string** | The snapshot 2 of identity | 
 **accessItemTypes** | **[]string** | An optional list of access item types (app, account, entitlement, etc...) to return.   If null or empty, all access items types are returned  | 
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]

### Return type

[**[]IdentityCompareResponse**](IdentityCompareResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompareIdentitySnapshotsAccessType

> []AccessItemDiff CompareIdentitySnapshotsAccessType(ctx, id, accessType).AccessAssociated(accessAssociated).Snapshot1(snapshot1).Snapshot2(snapshot2).Limit(limit).Offset(offset).Count(count).Execute()

Gets a list of differences of specific accessType for the given identity between 2 snapshots



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
    id := "8c190e6787aa4ed9a90bd9d5344523fb" // string | The identity id
    accessType := "role" // string | The specific type which needs to be compared
    accessAssociated := false // bool | Indicates if added or removed access needs to be returned. true - added, false - removed, null - both added & removed (optional)
    snapshot1 := "2008-03-01T13:00:00Z" // string | The snapshot 1 of identity (optional)
    snapshot2 := "2009-03-01T13:00:00Z" // string | The snapshot 2 of identity (optional)
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityHistoryApi.CompareIdentitySnapshotsAccessType(context.Background(), id, accessType).AccessAssociated(accessAssociated).Snapshot1(snapshot1).Snapshot2(snapshot2).Limit(limit).Offset(offset).Count(count).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityHistoryApi.CompareIdentitySnapshotsAccessType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CompareIdentitySnapshotsAccessType`: []AccessItemDiff
    fmt.Fprintf(os.Stdout, "Response from `IdentityHistoryApi.CompareIdentitySnapshotsAccessType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The identity id | 
**accessType** | **string** | The specific type which needs to be compared | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompareIdentitySnapshotsAccessTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accessAssociated** | **bool** | Indicates if added or removed access needs to be returned. true - added, false - removed, null - both added &amp; removed | 
 **snapshot1** | **string** | The snapshot 1 of identity | 
 **snapshot2** | **string** | The snapshot 2 of identity | 
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]

### Return type

[**[]AccessItemDiff**](AccessItemDiff.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHistoricalIdentity

> IdentityHistoryResponse GetHistoricalIdentity(ctx, id).Execute()

Get latest snapshot of identity



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
    id := "8c190e6787aa4ed9a90bd9d5344523fb" // string | The identity id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityHistoryApi.GetHistoricalIdentity(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityHistoryApi.GetHistoricalIdentity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHistoricalIdentity`: IdentityHistoryResponse
    fmt.Fprintf(os.Stdout, "Response from `IdentityHistoryApi.GetHistoricalIdentity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The identity id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHistoricalIdentityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IdentityHistoryResponse**](IdentityHistoryResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHistoricalIdentityEvents

> []GetHistoricalIdentityEvents200ResponseInner GetHistoricalIdentityEvents(ctx, id).From(from).EventTypes(eventTypes).AccessItemTypes(accessItemTypes).Limit(limit).Offset(offset).Count(count).Execute()

Lists all events for the given identity



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
    id := "8c190e6787aa4ed9a90bd9d5344523fb" // string | The identity id
    from := "2007-03-01T13:00:00Z" // string | The optional instant from which to return the access events (optional)
    eventTypes := []string{"Inner_example"} // []string | An optional list of event types to return.  If null or empty, all events are returned (optional)
    accessItemTypes := []string{"Inner_example"} // []string | An optional list of access item types (app, account, entitlement, etc...) to return.   If null or empty, all access items types are returned (optional)
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityHistoryApi.GetHistoricalIdentityEvents(context.Background(), id).From(from).EventTypes(eventTypes).AccessItemTypes(accessItemTypes).Limit(limit).Offset(offset).Count(count).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityHistoryApi.GetHistoricalIdentityEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHistoricalIdentityEvents`: []GetHistoricalIdentityEvents200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `IdentityHistoryApi.GetHistoricalIdentityEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The identity id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHistoricalIdentityEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **string** | The optional instant from which to return the access events | 
 **eventTypes** | **[]string** | An optional list of event types to return.  If null or empty, all events are returned | 
 **accessItemTypes** | **[]string** | An optional list of access item types (app, account, entitlement, etc...) to return.   If null or empty, all access items types are returned | 
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]

### Return type

[**[]GetHistoricalIdentityEvents200ResponseInner**](GetHistoricalIdentityEvents200ResponseInner.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentitySnapshot

> IdentityHistoryResponse GetIdentitySnapshot(ctx, id, date).Execute()

Gets an identity snapshot at a given date



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
    id := "8c190e6787aa4ed9a90bd9d5344523fb" // string | The identity id
    date := "2007-03-01T13:00:00Z" // string | The specified date

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityHistoryApi.GetIdentitySnapshot(context.Background(), id, date).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityHistoryApi.GetIdentitySnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdentitySnapshot`: IdentityHistoryResponse
    fmt.Fprintf(os.Stdout, "Response from `IdentityHistoryApi.GetIdentitySnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The identity id | 
**date** | **string** | The specified date | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentitySnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IdentityHistoryResponse**](IdentityHistoryResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentitySnapshotSummary

> []MetricResponse GetIdentitySnapshotSummary(ctx, id).Before(before).Interval(interval).TimeZone(timeZone).Limit(limit).Offset(offset).Count(count).Execute()

Gets the summary for the event count for a specific identity



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
    id := "8c190e6787aa4ed9a90bd9d5344523fb" // string | The identity id
    before := "2007-03-01T13:00:00Z" // string | The date before which snapshot summary is required (optional)
    interval := "interval_example" // string | The interval indicating day or month. Defaults to month if not specified (optional)
    timeZone := "UTC" // string | The time zone. Defaults to UTC if not provided (optional)
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityHistoryApi.GetIdentitySnapshotSummary(context.Background(), id).Before(before).Interval(interval).TimeZone(timeZone).Limit(limit).Offset(offset).Count(count).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityHistoryApi.GetIdentitySnapshotSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdentitySnapshotSummary`: []MetricResponse
    fmt.Fprintf(os.Stdout, "Response from `IdentityHistoryApi.GetIdentitySnapshotSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The identity id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentitySnapshotSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **before** | **string** | The date before which snapshot summary is required | 
 **interval** | **string** | The interval indicating day or month. Defaults to month if not specified | 
 **timeZone** | **string** | The time zone. Defaults to UTC if not provided | 
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]

### Return type

[**[]MetricResponse**](MetricResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityStartDate

> string GetIdentityStartDate(ctx, id).Execute()

Gets the start date of the identity



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
    id := "8c190e6787aa4ed9a90bd9d5344523fb" // string | The identity id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityHistoryApi.GetIdentityStartDate(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityHistoryApi.GetIdentityStartDate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdentityStartDate`: string
    fmt.Fprintf(os.Stdout, "Response from `IdentityHistoryApi.GetIdentityStartDate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The identity id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityStartDateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHistoricalIdentities

> []IdentityListItem ListHistoricalIdentities(ctx).StartsWithQuery(startsWithQuery).IsDeleted(isDeleted).IsActive(isActive).Limit(limit).Offset(offset).Execute()

Lists all the identities



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
    startsWithQuery := "Ada" // string | This param is used for starts-with search for first, last and display name of the identity (optional)
    isDeleted := true // bool | Indicates if we want to only list down deleted identities or not. (optional)
    isActive := true // bool | Indicates if we want to only list active or inactive identities. (optional)
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityHistoryApi.ListHistoricalIdentities(context.Background()).StartsWithQuery(startsWithQuery).IsDeleted(isDeleted).IsActive(isActive).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityHistoryApi.ListHistoricalIdentities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListHistoricalIdentities`: []IdentityListItem
    fmt.Fprintf(os.Stdout, "Response from `IdentityHistoryApi.ListHistoricalIdentities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListHistoricalIdentitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startsWithQuery** | **string** | This param is used for starts-with search for first, last and display name of the identity | 
 **isDeleted** | **bool** | Indicates if we want to only list down deleted identities or not. | 
 **isActive** | **bool** | Indicates if we want to only list active or inactive identities. | 
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]

### Return type

[**[]IdentityListItem**](IdentityListItem.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIdentityAccessItems

> []ListIdentityAccessItems200ResponseInner ListIdentityAccessItems(ctx, id).Type_(type_).Execute()

Gets a list of access items for the identity filtered by item type



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
    id := "8c190e6787aa4ed9a90bd9d5344523fb" // string | The identity id
    type_ := "account" // string | The type of access item for the identity. If not provided, it defaults to account (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityHistoryApi.ListIdentityAccessItems(context.Background(), id).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityHistoryApi.ListIdentityAccessItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIdentityAccessItems`: []ListIdentityAccessItems200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `IdentityHistoryApi.ListIdentityAccessItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The identity id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListIdentityAccessItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | The type of access item for the identity. If not provided, it defaults to account | 

### Return type

[**[]ListIdentityAccessItems200ResponseInner**](ListIdentityAccessItems200ResponseInner.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIdentitySnapshotAccessItems

> []ListIdentityAccessItems200ResponseInner ListIdentitySnapshotAccessItems(ctx, id, date).Type_(type_).Execute()

Gets the list of identity access items at a given date filterd by item type



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
    id := "8c190e6787aa4ed9a90bd9d5344523fb" // string | The identity id
    date := "2007-03-01T13:00:00Z" // string | The specified date
    type_ := "account" // string | The access item type (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityHistoryApi.ListIdentitySnapshotAccessItems(context.Background(), id, date).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityHistoryApi.ListIdentitySnapshotAccessItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIdentitySnapshotAccessItems`: []ListIdentityAccessItems200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `IdentityHistoryApi.ListIdentitySnapshotAccessItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The identity id | 
**date** | **string** | The specified date | 

### Other Parameters

Other parameters are passed through a pointer to a apiListIdentitySnapshotAccessItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **type_** | **string** | The access item type | 

### Return type

[**[]ListIdentityAccessItems200ResponseInner**](ListIdentityAccessItems200ResponseInner.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIdentitySnapshots

> []IdentitySnapshotSummaryResponse ListIdentitySnapshots(ctx, id).Start(start).Interval(interval).Limit(limit).Offset(offset).Count(count).Execute()

Lists all the snapshots for the identity



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
    id := "8c190e6787aa4ed9a90bd9d5344523fb" // string | The identity id
    start := "2007-03-01T13:00:00Z" // string | The specified start date (optional)
    interval := "interval_example" // string | The interval indicating the range in day or month for the specified interval-name (optional)
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityHistoryApi.ListIdentitySnapshots(context.Background(), id).Start(start).Interval(interval).Limit(limit).Offset(offset).Count(count).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityHistoryApi.ListIdentitySnapshots``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIdentitySnapshots`: []IdentitySnapshotSummaryResponse
    fmt.Fprintf(os.Stdout, "Response from `IdentityHistoryApi.ListIdentitySnapshots`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The identity id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListIdentitySnapshotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **string** | The specified start date | 
 **interval** | **string** | The interval indicating the range in day or month for the specified interval-name | 
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]

### Return type

[**[]IdentitySnapshotSummaryResponse**](IdentitySnapshotSummaryResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

