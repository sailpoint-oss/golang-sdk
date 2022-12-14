# \SearchApi

All URIs are relative to *https://sailpoint.api.identitynow.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EntitlementSearch**](SearchApi.md#EntitlementSearch) | **Get** /search/entitlements | Searches and retrieves the entitlements.
[**EntitlementSearchExport**](SearchApi.md#EntitlementSearchExport) | **Get** /search/entitlements/runExport | Runs csv results export job for a given search for entitlements.
[**EntitlementSearchExportPost**](SearchApi.md#EntitlementSearchExportPost) | **Post** /search/entitlements/runExport | Runs csv results export job for a given search for entitlements.
[**EntitlementSearchPost**](SearchApi.md#EntitlementSearchPost) | **Post** /search/entitlements | Searches and retrieves the entitlements.
[**ExportIdentitySearch**](SearchApi.md#ExportIdentitySearch) | **Get** /search/identities/runExport | Runs csv results export job for a given search for identities.
[**ExportSearchCsv**](SearchApi.md#ExportSearchCsv) | **Get** /search/runExport | Runs csv results export job for a given search query.
[**ExportSearchCsvPost**](SearchApi.md#ExportSearchCsvPost) | **Post** /search/runExport | Runs csv results export job for a given search query.
[**GetSearchIndexMapping**](SearchApi.md#GetSearchIndexMapping) | **Get** /search/{index}/mappings | Retrieves the mappings and operators for the search service for the given index path.
[**GetSearchMapping**](SearchApi.md#GetSearchMapping) | **Get** /search/mappings | Retrieves the mappings and operators for the search service.
[**RunEventSearch**](SearchApi.md#RunEventSearch) | **Get** /search/events | Searches and retrieves the events.
[**RunEventSearchPost**](SearchApi.md#RunEventSearchPost) | **Post** /search/events | Searches and retrieves the events.
[**RunIdentitySearch**](SearchApi.md#RunIdentitySearch) | **Post** /search/identities/runExport | Runs csv results export job for a given search for identities.
[**RunSearch**](SearchApi.md#RunSearch) | **Get** /search | Searches and retrieves the types specified (currently identity, entitlement, and event).
[**RunSearchPost**](SearchApi.md#RunSearchPost) | **Post** /search | Searches and retrieves the types specified (current identity, entitlement, and event).
[**SearchIdentities**](SearchApi.md#SearchIdentities) | **Get** /search/identities | Searches and retrieves the identities.
[**SearchIdentitiesPost**](SearchApi.md#SearchIdentitiesPost) | **Post** /search/identities | Searches and retrieves the identities.



## EntitlementSearch

> []SearchEntitlement EntitlementSearch(ctx).Query(query).Sort(sort).Offset(offset).Limit(limit).Fields(fields).Execute()

Searches and retrieves the entitlements.



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
    query := "query_example" // string | Query using the [Elastic Search Query String syntax](https://www.elastic.co/guide/en/elasticsearch/reference/5.2/query-dsl-query-string-query.html#query-string-syntax) from the Query DSL.
    sort := "sort_example" // string | One or more attributes on which to sort, each separated by a ','. Prefix with a minus sign (ex. -dateCreated) for descending sort. (optional)
    offset := int32(56) // int32 | Paging offset. (optional) (default to 0)
    limit := int32(56) // int32 | Paging limit. (optional) (default to 250)
    fields := []string{"Inner_example"} // []string | List of fields that the query should be restricted. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.EntitlementSearch(context.Background()).Query(query).Sort(sort).Offset(offset).Limit(limit).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.EntitlementSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EntitlementSearch`: []SearchEntitlement
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.EntitlementSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEntitlementSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | Query using the [Elastic Search Query String syntax](https://www.elastic.co/guide/en/elasticsearch/reference/5.2/query-dsl-query-string-query.html#query-string-syntax) from the Query DSL. | 
 **sort** | **string** | One or more attributes on which to sort, each separated by a &#39;,&#39;. Prefix with a minus sign (ex. -dateCreated) for descending sort. | 
 **offset** | **int32** | Paging offset. | [default to 0]
 **limit** | **int32** | Paging limit. | [default to 250]
 **fields** | **[]string** | List of fields that the query should be restricted. | 

### Return type

[**[]SearchEntitlement**](SearchEntitlement.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EntitlementSearchExport

> TaskResult EntitlementSearchExport(ctx).Query(query).QueryFields(queryFields).Sort(sort).Fields(fields).Execute()

Runs csv results export job for a given search for entitlements.



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
    query := "query_example" // string | Query using the [Elastic Search Query String syntax](https://www.elastic.co/guide/en/elasticsearch/reference/5.2/query-dsl-query-string-query.html#query-string-syntax) from the Query DSL.
    queryFields := []string{"Inner_example"} // []string | List of query fields that the query should be restricted. (optional)
    sort := "sort_example" // string | One or more attributes on which to sort, each separated by a ','. Prefix with a minus sign (ex. -dateCreated) for descending sort. (optional)
    fields := []string{"Inner_example"} // []string | List of field columns that should be included in the export. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.EntitlementSearchExport(context.Background()).Query(query).QueryFields(queryFields).Sort(sort).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.EntitlementSearchExport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EntitlementSearchExport`: TaskResult
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.EntitlementSearchExport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEntitlementSearchExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | Query using the [Elastic Search Query String syntax](https://www.elastic.co/guide/en/elasticsearch/reference/5.2/query-dsl-query-string-query.html#query-string-syntax) from the Query DSL. | 
 **queryFields** | **[]string** | List of query fields that the query should be restricted. | 
 **sort** | **string** | One or more attributes on which to sort, each separated by a &#39;,&#39;. Prefix with a minus sign (ex. -dateCreated) for descending sort. | 
 **fields** | **[]string** | List of field columns that should be included in the export. | 

### Return type

[**TaskResult**](TaskResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EntitlementSearchExportPost

> TaskResult EntitlementSearchExportPost(ctx).ExportSearchCsvPostRequest(exportSearchCsvPostRequest).Sort(sort).Fields(fields).Execute()

Runs csv results export job for a given search for entitlements.



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
    exportSearchCsvPostRequest := *openapiclient.NewExportSearchCsvPostRequest() // ExportSearchCsvPostRequest | Query object using the query portion of the [Elastic Search Query DSL JSON object](https://www.elastic.co/guide/en/elasticsearch/reference/5.2/query-dsl.html).
    sort := "sort_example" // string | One or more attributes on which to sort, each separated by a ','. Prefix with a minus sign (ex. -dateCreated) for descending sort. (optional)
    fields := []string{"Inner_example"} // []string | List of field columns that should be included in the export. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.EntitlementSearchExportPost(context.Background()).ExportSearchCsvPostRequest(exportSearchCsvPostRequest).Sort(sort).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.EntitlementSearchExportPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EntitlementSearchExportPost`: TaskResult
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.EntitlementSearchExportPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEntitlementSearchExportPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **exportSearchCsvPostRequest** | [**ExportSearchCsvPostRequest**](ExportSearchCsvPostRequest.md) | Query object using the query portion of the [Elastic Search Query DSL JSON object](https://www.elastic.co/guide/en/elasticsearch/reference/5.2/query-dsl.html). | 
 **sort** | **string** | One or more attributes on which to sort, each separated by a &#39;,&#39;. Prefix with a minus sign (ex. -dateCreated) for descending sort. | 
 **fields** | **[]string** | List of field columns that should be included in the export. | 

### Return type

[**TaskResult**](TaskResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EntitlementSearchPost

> []SearchEntitlement EntitlementSearchPost(ctx).Query(query).Sort(sort).Offset(offset).Limit(limit).Fields(fields).Execute()

Searches and retrieves the entitlements.



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
    query := "query_example" // string | Query using the [Elastic Search Query String syntax](https://www.elastic.co/guide/en/elasticsearch/reference/5.2/query-dsl-query-string-query.html#query-string-syntax) from the Query DSL.
    sort := "sort_example" // string | One or more attributes on which to sort, each separated by a ','. Prefix with a minus sign (ex. -dateCreated) for descending sort. (optional)
    offset := int32(56) // int32 | Paging offset. (optional) (default to 0)
    limit := int32(56) // int32 | Paging limit. (optional) (default to 250)
    fields := []string{"Inner_example"} // []string | List of fields that the query should be restricted. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.EntitlementSearchPost(context.Background()).Query(query).Sort(sort).Offset(offset).Limit(limit).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.EntitlementSearchPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EntitlementSearchPost`: []SearchEntitlement
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.EntitlementSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEntitlementSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | Query using the [Elastic Search Query String syntax](https://www.elastic.co/guide/en/elasticsearch/reference/5.2/query-dsl-query-string-query.html#query-string-syntax) from the Query DSL. | 
 **sort** | **string** | One or more attributes on which to sort, each separated by a &#39;,&#39;. Prefix with a minus sign (ex. -dateCreated) for descending sort. | 
 **offset** | **int32** | Paging offset. | [default to 0]
 **limit** | **int32** | Paging limit. | [default to 250]
 **fields** | **[]string** | List of fields that the query should be restricted. | 

### Return type

[**[]SearchEntitlement**](SearchEntitlement.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportIdentitySearch

> TaskResult ExportIdentitySearch(ctx).Query(query).QueryFields(queryFields).Sort(sort).Fields(fields).Execute()

Runs csv results export job for a given search for identities.



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
    query := "query_example" // string | Query using the [Elastic Search Query String syntax](https://www.elastic.co/guide/en/elasticsearch/reference/5.2/query-dsl-query-string-query.html#query-string-syntax) from the Query DSL.
    queryFields := []string{"Inner_example"} // []string | List of query fields that the query should be restricted. (optional)
    sort := "sort_example" // string | One or more attributes on which to sort, each separated by a ','. Prefix with a minus sign (ex. -dateCreated) for descending sort. (optional)
    fields := []string{"Inner_example"} // []string | List of field columns that should be included in the export. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.ExportIdentitySearch(context.Background()).Query(query).QueryFields(queryFields).Sort(sort).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.ExportIdentitySearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportIdentitySearch`: TaskResult
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.ExportIdentitySearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportIdentitySearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | Query using the [Elastic Search Query String syntax](https://www.elastic.co/guide/en/elasticsearch/reference/5.2/query-dsl-query-string-query.html#query-string-syntax) from the Query DSL. | 
 **queryFields** | **[]string** | List of query fields that the query should be restricted. | 
 **sort** | **string** | One or more attributes on which to sort, each separated by a &#39;,&#39;. Prefix with a minus sign (ex. -dateCreated) for descending sort. | 
 **fields** | **[]string** | List of field columns that should be included in the export. | 

### Return type

[**TaskResult**](TaskResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportSearchCsv

> TaskResult ExportSearchCsv(ctx).Query(query).Index(index).Type_(type_).QueryFields(queryFields).Sort(sort).Fields(fields).Execute()

Runs csv results export job for a given search query.



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
    query := "query_example" // string | Query using the [Elastic Search Query String syntax](https://www.elastic.co/guide/en/elasticsearch/reference/5.2/query-dsl-query-string-query.html#query-string-syntax) from the Query DSL.
    index := "index_example" // string | Index to be used (currently identities, entitlements, and events).
    type_ := "type__example" // string | Type to query (currently either identity, entitlement or event).
    queryFields := []string{"Inner_example"} // []string | List of query fields that the query should be restricted. (optional)
    sort := "sort_example" // string | One or more attributes on which to sort, each separated by a ','. Prefix with a minus sign (ex. -dateCreated) for descending sort. (optional)
    fields := []string{"Inner_example"} // []string | List of fields that the query should be restricted. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.ExportSearchCsv(context.Background()).Query(query).Index(index).Type_(type_).QueryFields(queryFields).Sort(sort).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.ExportSearchCsv``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportSearchCsv`: TaskResult
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.ExportSearchCsv`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportSearchCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | Query using the [Elastic Search Query String syntax](https://www.elastic.co/guide/en/elasticsearch/reference/5.2/query-dsl-query-string-query.html#query-string-syntax) from the Query DSL. | 
 **index** | **string** | Index to be used (currently identities, entitlements, and events). | 
 **type_** | **string** | Type to query (currently either identity, entitlement or event). | 
 **queryFields** | **[]string** | List of query fields that the query should be restricted. | 
 **sort** | **string** | One or more attributes on which to sort, each separated by a &#39;,&#39;. Prefix with a minus sign (ex. -dateCreated) for descending sort. | 
 **fields** | **[]string** | List of fields that the query should be restricted. | 

### Return type

[**TaskResult**](TaskResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportSearchCsvPost

> TaskResult ExportSearchCsvPost(ctx).Index(index).Type_(type_).ExportSearchCsvPostRequest(exportSearchCsvPostRequest).Sort(sort).Fields(fields).Execute()

Runs csv results export job for a given search query.



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
    index := "index_example" // string | Index to be used (currently identities, entitlements, and events).
    type_ := "type__example" // string | Type to query (currently either identity, entitlement or event).
    exportSearchCsvPostRequest := *openapiclient.NewExportSearchCsvPostRequest() // ExportSearchCsvPostRequest | Query object using the query portion of the [Elastic Search Query DSL JSON object](https://www.elastic.co/guide/en/elasticsearch/reference/5.2/query-dsl.html).
    sort := "sort_example" // string | One or more attributes on which to sort, each separated by a ','. Prefix with a minus sign (ex. -dateCreated) for descending sort. (optional)
    fields := []string{"Inner_example"} // []string | List of field columns that should be included in the export. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.ExportSearchCsvPost(context.Background()).Index(index).Type_(type_).ExportSearchCsvPostRequest(exportSearchCsvPostRequest).Sort(sort).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.ExportSearchCsvPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportSearchCsvPost`: TaskResult
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.ExportSearchCsvPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportSearchCsvPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **index** | **string** | Index to be used (currently identities, entitlements, and events). | 
 **type_** | **string** | Type to query (currently either identity, entitlement or event). | 
 **exportSearchCsvPostRequest** | [**ExportSearchCsvPostRequest**](ExportSearchCsvPostRequest.md) | Query object using the query portion of the [Elastic Search Query DSL JSON object](https://www.elastic.co/guide/en/elasticsearch/reference/5.2/query-dsl.html). | 
 **sort** | **string** | One or more attributes on which to sort, each separated by a &#39;,&#39;. Prefix with a minus sign (ex. -dateCreated) for descending sort. | 
 **fields** | **[]string** | List of field columns that should be included in the export. | 

### Return type

[**TaskResult**](TaskResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSearchIndexMapping

> Mapping GetSearchIndexMapping(ctx).Index(index).Execute()

Retrieves the mappings and operators for the search service for the given index path.



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
    index := "index_example" // string | Index to be used (currently identities, entitlements, and events).

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.GetSearchIndexMapping(context.Background()).Index(index).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.GetSearchIndexMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSearchIndexMapping`: Mapping
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.GetSearchIndexMapping`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSearchIndexMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **index** | **string** | Index to be used (currently identities, entitlements, and events). | 

### Return type

[**Mapping**](Mapping.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSearchMapping

> Mapping GetSearchMapping(ctx).Execute()

Retrieves the mappings and operators for the search service.



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
    resp, r, err := apiClient.SearchApi.GetSearchMapping(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.GetSearchMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSearchMapping`: Mapping
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.GetSearchMapping`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSearchMappingRequest struct via the builder pattern


### Return type

[**Mapping**](Mapping.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RunEventSearch

> []SearchEvent RunEventSearch(ctx).Query(query).Sort(sort).Offset(offset).Limit(limit).Fields(fields).Execute()

Searches and retrieves the events.



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
    query := "query_example" // string | Query using the [Elastic Search Query String syntax](https://www.elastic.co/guide/en/elasticsearch/reference/5.2/query-dsl-query-string-query.html#query-string-syntax) from the Query DSL.
    sort := "sort_example" // string | One or more attributes on which to sort, each separated by a ','. Prefix with a minus sign (ex. -dateCreated) for descending sort. (optional)
    offset := int32(56) // int32 | Paging offset. (optional) (default to 0)
    limit := int32(56) // int32 | Paging limit. (optional) (default to 250)
    fields := []string{"Inner_example"} // []string | List of fields that the query should be restricted. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.RunEventSearch(context.Background()).Query(query).Sort(sort).Offset(offset).Limit(limit).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.RunEventSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RunEventSearch`: []SearchEvent
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.RunEventSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRunEventSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | Query using the [Elastic Search Query String syntax](https://www.elastic.co/guide/en/elasticsearch/reference/5.2/query-dsl-query-string-query.html#query-string-syntax) from the Query DSL. | 
 **sort** | **string** | One or more attributes on which to sort, each separated by a &#39;,&#39;. Prefix with a minus sign (ex. -dateCreated) for descending sort. | 
 **offset** | **int32** | Paging offset. | [default to 0]
 **limit** | **int32** | Paging limit. | [default to 250]
 **fields** | **[]string** | List of fields that the query should be restricted. | 

### Return type

[**[]SearchEvent**](SearchEvent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RunEventSearchPost

> []SearchEvent RunEventSearchPost(ctx).Query(query).Sort(sort).Offset(offset).Limit(limit).Fields(fields).Execute()

Searches and retrieves the events.



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
    query := "query_example" // string | Query using the [Elastic Search Query String syntax](https://www.elastic.co/guide/en/elasticsearch/reference/5.2/query-dsl-query-string-query.html#query-string-syntax) from the Query DSL.
    sort := "sort_example" // string | One or more attributes on which to sort, each separated by a ','. Prefix with a minus sign (ex. -dateCreated) for descending sort. (optional)
    offset := int32(56) // int32 | Paging offset. (optional) (default to 0)
    limit := int32(56) // int32 | Paging limit. (optional) (default to 250)
    fields := []string{"Inner_example"} // []string | List of fields that the query should be restricted. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.RunEventSearchPost(context.Background()).Query(query).Sort(sort).Offset(offset).Limit(limit).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.RunEventSearchPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RunEventSearchPost`: []SearchEvent
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.RunEventSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRunEventSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | Query using the [Elastic Search Query String syntax](https://www.elastic.co/guide/en/elasticsearch/reference/5.2/query-dsl-query-string-query.html#query-string-syntax) from the Query DSL. | 
 **sort** | **string** | One or more attributes on which to sort, each separated by a &#39;,&#39;. Prefix with a minus sign (ex. -dateCreated) for descending sort. | 
 **offset** | **int32** | Paging offset. | [default to 0]
 **limit** | **int32** | Paging limit. | [default to 250]
 **fields** | **[]string** | List of fields that the query should be restricted. | 

### Return type

[**[]SearchEvent**](SearchEvent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RunIdentitySearch

> TaskResult RunIdentitySearch(ctx).ExportSearchCsvPostRequest(exportSearchCsvPostRequest).Sort(sort).Fields(fields).Execute()

Runs csv results export job for a given search for identities.



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
    exportSearchCsvPostRequest := *openapiclient.NewExportSearchCsvPostRequest() // ExportSearchCsvPostRequest | Query object using the query portion of the [Elastic Search Query DSL JSON object](https://www.elastic.co/guide/en/elasticsearch/reference/5.2/query-dsl.html).
    sort := "sort_example" // string | One or more attributes on which to sort, each separated by a ','. Prefix with a minus sign (ex. -dateCreated) for descending sort. (optional)
    fields := []string{"Inner_example"} // []string | List of field columns that should be included in the export. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.RunIdentitySearch(context.Background()).ExportSearchCsvPostRequest(exportSearchCsvPostRequest).Sort(sort).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.RunIdentitySearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RunIdentitySearch`: TaskResult
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.RunIdentitySearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRunIdentitySearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **exportSearchCsvPostRequest** | [**ExportSearchCsvPostRequest**](ExportSearchCsvPostRequest.md) | Query object using the query portion of the [Elastic Search Query DSL JSON object](https://www.elastic.co/guide/en/elasticsearch/reference/5.2/query-dsl.html). | 
 **sort** | **string** | One or more attributes on which to sort, each separated by a &#39;,&#39;. Prefix with a minus sign (ex. -dateCreated) for descending sort. | 
 **fields** | **[]string** | List of field columns that should be included in the export. | 

### Return type

[**TaskResult**](TaskResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RunSearch

> RunSearch200Response RunSearch(ctx).Query(query).Types(types).Sort(sort).Offset(offset).Limit(limit).Fields(fields).Execute()

Searches and retrieves the types specified (currently identity, entitlement, and event).



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
    query := "query_example" // string | Query using the [Elastic Search Query String syntax](https://www.elastic.co/guide/en/elasticsearch/reference/5.2/query-dsl-query-string-query.html#query-string-syntax) from the Query DSL.
    types := []string{"Inner_example"} // []string | List of types to query (currently identity, entitlement or event). (optional)
    sort := "sort_example" // string | One or more attributes on which to sort, each separated by a ','. Prefix with a minus sign (ex. -dateCreated) for descending sort. (optional)
    offset := int32(56) // int32 | Paging offset. (optional) (default to 0)
    limit := int32(56) // int32 | Paging limit. (optional) (default to 250)
    fields := []string{"Inner_example"} // []string | List of fields that the query should be restricted. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.RunSearch(context.Background()).Query(query).Types(types).Sort(sort).Offset(offset).Limit(limit).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.RunSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RunSearch`: RunSearch200Response
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.RunSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRunSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | Query using the [Elastic Search Query String syntax](https://www.elastic.co/guide/en/elasticsearch/reference/5.2/query-dsl-query-string-query.html#query-string-syntax) from the Query DSL. | 
 **types** | **[]string** | List of types to query (currently identity, entitlement or event). | 
 **sort** | **string** | One or more attributes on which to sort, each separated by a &#39;,&#39;. Prefix with a minus sign (ex. -dateCreated) for descending sort. | 
 **offset** | **int32** | Paging offset. | [default to 0]
 **limit** | **int32** | Paging limit. | [default to 250]
 **fields** | **[]string** | List of fields that the query should be restricted. | 

### Return type

[**RunSearch200Response**](RunSearch200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RunSearchPost

> RunSearch200Response RunSearchPost(ctx).Query(query).Types(types).Sort(sort).Offset(offset).Limit(limit).Fields(fields).Execute()

Searches and retrieves the types specified (current identity, entitlement, and event).



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
    query := "query_example" // string | Query using the [Elastic Search Query String syntax](https://www.elastic.co/guide/en/elasticsearch/reference/5.2/query-dsl-query-string-query.html#query-string-syntax) from the Query DSL.
    types := []string{"Inner_example"} // []string | List of types to query (currently identity, entitlement or event). (optional)
    sort := "sort_example" // string | One or more attributes on which to sort, each separated by a ','. Prefix with a minus sign (ex. -dateCreated) for descending sort. (optional)
    offset := int32(56) // int32 | Paging offset. (optional) (default to 0)
    limit := int32(56) // int32 | Paging limit. (optional) (default to 250)
    fields := []string{"Inner_example"} // []string | List of fields that the query should be restricted. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.RunSearchPost(context.Background()).Query(query).Types(types).Sort(sort).Offset(offset).Limit(limit).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.RunSearchPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RunSearchPost`: RunSearch200Response
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.RunSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRunSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | Query using the [Elastic Search Query String syntax](https://www.elastic.co/guide/en/elasticsearch/reference/5.2/query-dsl-query-string-query.html#query-string-syntax) from the Query DSL. | 
 **types** | **[]string** | List of types to query (currently identity, entitlement or event). | 
 **sort** | **string** | One or more attributes on which to sort, each separated by a &#39;,&#39;. Prefix with a minus sign (ex. -dateCreated) for descending sort. | 
 **offset** | **int32** | Paging offset. | [default to 0]
 **limit** | **int32** | Paging limit. | [default to 250]
 **fields** | **[]string** | List of fields that the query should be restricted. | 

### Return type

[**RunSearch200Response**](RunSearch200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchIdentities

> []SearchIdentity SearchIdentities(ctx).Query(query).Sort(sort).Offset(offset).Limit(limit).Fields(fields).Execute()

Searches and retrieves the identities.



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
    query := "query_example" // string | Query using the [Elastic Search Query String syntax](https://www.elastic.co/guide/en/elasticsearch/reference/5.2/query-dsl-query-string-query.html#query-string-syntax) from the Query DSL.
    sort := "sort_example" // string | One or more attributes on which to sort, each separated by a ','. Prefix with a minus sign (ex. -dateCreated) for descending sort. (optional)
    offset := int32(56) // int32 | Paging offset. (optional) (default to 0)
    limit := int32(56) // int32 | Paging limit. (optional) (default to 250)
    fields := []string{"Inner_example"} // []string | List of fields that the query should be restricted. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.SearchIdentities(context.Background()).Query(query).Sort(sort).Offset(offset).Limit(limit).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.SearchIdentities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchIdentities`: []SearchIdentity
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.SearchIdentities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchIdentitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | Query using the [Elastic Search Query String syntax](https://www.elastic.co/guide/en/elasticsearch/reference/5.2/query-dsl-query-string-query.html#query-string-syntax) from the Query DSL. | 
 **sort** | **string** | One or more attributes on which to sort, each separated by a &#39;,&#39;. Prefix with a minus sign (ex. -dateCreated) for descending sort. | 
 **offset** | **int32** | Paging offset. | [default to 0]
 **limit** | **int32** | Paging limit. | [default to 250]
 **fields** | **[]string** | List of fields that the query should be restricted. | 

### Return type

[**[]SearchIdentity**](SearchIdentity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchIdentitiesPost

> []SearchIdentity SearchIdentitiesPost(ctx).Query(query).Sort(sort).Offset(offset).Limit(limit).Fields(fields).Execute()

Searches and retrieves the identities.



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
    query := "query_example" // string | Query using the [Elastic Search Query String syntax](https://www.elastic.co/guide/en/elasticsearch/reference/5.2/query-dsl-query-string-query.html#query-string-syntax) from the Query DSL.
    sort := "sort_example" // string | One or more attributes on which to sort, each separated by a ','. Prefix with a minus sign (ex. -dateCreated) for descending sort. (optional)
    offset := int32(56) // int32 | Paging offset. (optional) (default to 0)
    limit := int32(56) // int32 | Paging limit. (optional) (default to 250)
    fields := []string{"Inner_example"} // []string | List of fields that the query should be restricted. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.SearchIdentitiesPost(context.Background()).Query(query).Sort(sort).Offset(offset).Limit(limit).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.SearchIdentitiesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchIdentitiesPost`: []SearchIdentity
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.SearchIdentitiesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchIdentitiesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | Query using the [Elastic Search Query String syntax](https://www.elastic.co/guide/en/elasticsearch/reference/5.2/query-dsl-query-string-query.html#query-string-syntax) from the Query DSL. | 
 **sort** | **string** | One or more attributes on which to sort, each separated by a &#39;,&#39;. Prefix with a minus sign (ex. -dateCreated) for descending sort. | 
 **offset** | **int32** | Paging offset. | [default to 0]
 **limit** | **int32** | Paging limit. | [default to 250]
 **fields** | **[]string** | List of fields that the query should be restricted. | 

### Return type

[**[]SearchIdentity**](SearchIdentity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

