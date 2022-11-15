# \SavedSearchApi

All URIs are relative to *https://sailpoint.api.identitynow.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SavedSearchCreate**](SavedSearchApi.md#SavedSearchCreate) | **Post** /saved-searches | Create a saved search
[**SavedSearchDelete**](SavedSearchApi.md#SavedSearchDelete) | **Delete** /saved-searches/{id} | Delete a document by ID
[**SavedSearchExecute**](SavedSearchApi.md#SavedSearchExecute) | **Post** /saved-searches/{id}/execute | Execute a saved search by ID
[**SavedSearchGet**](SavedSearchApi.md#SavedSearchGet) | **Get** /saved-searches/{id} | Return a saved search by ID
[**SavedSearchList**](SavedSearchApi.md#SavedSearchList) | **Get** /saved-searches | Return a list of Saved Searches
[**SavedSearchUpdate**](SavedSearchApi.md#SavedSearchUpdate) | **Put** /saved-searches/{id} | Updates an existing saved search 



## SavedSearchCreate

> SavedSearch SavedSearchCreate(ctx).SavedSearchCreateRequest(savedSearchCreateRequest).Execute()

Create a saved search



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
    savedSearchCreateRequest := *openapiclient.NewSavedSearchCreateRequest([]openapiclient.Index{openapiclient.Index("accessprofiles")}, "@accounts(disabled:true)") // SavedSearchCreateRequest | The saved search to persist.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SavedSearchApi.SavedSearchCreate(context.Background()).SavedSearchCreateRequest(savedSearchCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SavedSearchApi.SavedSearchCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SavedSearchCreate`: SavedSearch
    fmt.Fprintf(os.Stdout, "Response from `SavedSearchApi.SavedSearchCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSavedSearchCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **savedSearchCreateRequest** | [**SavedSearchCreateRequest**](SavedSearchCreateRequest.md) | The saved search to persist. | 

### Return type

[**SavedSearch**](SavedSearch.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SavedSearchDelete

> SavedSearchDelete(ctx, id).Execute()

Delete a document by ID



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
    resp, r, err := apiClient.SavedSearchApi.SavedSearchDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SavedSearchApi.SavedSearchDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSavedSearchDeleteRequest struct via the builder pattern


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


## SavedSearchExecute

> SavedSearchExecute(ctx, id).SearchArguments(searchArguments).Execute()

Execute a saved search by ID



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
    searchArguments := *openapiclient.NewSearchArguments() // SearchArguments | When saved search execution is triggered by a scheduled search, *scheduleId* will specify the ID of the triggering scheduled search.  If *scheduleId* is not specified (when execution is triggered by a UI test), the *owner* and *recipients* arguments must be provided. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SavedSearchApi.SavedSearchExecute(context.Background(), id).SearchArguments(searchArguments).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SavedSearchApi.SavedSearchExecute``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSavedSearchExecuteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **searchArguments** | [**SearchArguments**](SearchArguments.md) | When saved search execution is triggered by a scheduled search, *scheduleId* will specify the ID of the triggering scheduled search.  If *scheduleId* is not specified (when execution is triggered by a UI test), the *owner* and *recipients* arguments must be provided.  | 

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


## SavedSearchGet

> SavedSearch SavedSearchGet(ctx, id).Execute()

Return a saved search by ID



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
    resp, r, err := apiClient.SavedSearchApi.SavedSearchGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SavedSearchApi.SavedSearchGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SavedSearchGet`: SavedSearch
    fmt.Fprintf(os.Stdout, "Response from `SavedSearchApi.SavedSearchGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the requested document. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSavedSearchGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SavedSearch**](SavedSearch.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SavedSearchList

> []SavedSearch SavedSearchList(ctx).Offset(offset).Limit(limit).Count(count).Filters(filters).Execute()

Return a list of Saved Searches



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
    filters := "public eq true" // string | An expression used to constrain the result set using the filtering syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results).  Allowed filter properties: *owner.id*, *public*  Allowed filter operator: *eq*  **Example filters**:  ```owner.id eq \"0de46054-fe90-434a-b84e-c6b3359d0c64\"``` -- returns saved searches for the specified owner ID  ```public eq true``` -- returns all public saved searches  ```owner.id eq me or public eq true``` -- returns all of the current user's saved searches as well as all public saved searches belonging to other users in the current org  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SavedSearchApi.SavedSearchList(context.Background()).Offset(offset).Limit(limit).Count(count).Filters(filters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SavedSearchApi.SavedSearchList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SavedSearchList`: []SavedSearch
    fmt.Fprintf(os.Stdout, "Response from `SavedSearchApi.SavedSearchList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSavedSearchListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **filters** | **string** | An expression used to constrain the result set using the filtering syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results).  Allowed filter properties: *owner.id*, *public*  Allowed filter operator: *eq*  **Example filters**:  &#x60;&#x60;&#x60;owner.id eq \&quot;0de46054-fe90-434a-b84e-c6b3359d0c64\&quot;&#x60;&#x60;&#x60; -- returns saved searches for the specified owner ID  &#x60;&#x60;&#x60;public eq true&#x60;&#x60;&#x60; -- returns all public saved searches  &#x60;&#x60;&#x60;owner.id eq me or public eq true&#x60;&#x60;&#x60; -- returns all of the current user&#39;s saved searches as well as all public saved searches belonging to other users in the current org  | 

### Return type

[**[]SavedSearch**](SavedSearch.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SavedSearchUpdate

> SavedSearch SavedSearchUpdate(ctx, id).SavedSearch(savedSearch).Execute()

Updates an existing saved search 



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
    savedSearch := *openapiclient.NewSavedSearch([]openapiclient.Index{openapiclient.Index("accessprofiles")}, "@accounts(disabled:true)") // SavedSearch | The saved search to persist.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SavedSearchApi.SavedSearchUpdate(context.Background(), id).SavedSearch(savedSearch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SavedSearchApi.SavedSearchUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SavedSearchUpdate`: SavedSearch
    fmt.Fprintf(os.Stdout, "Response from `SavedSearchApi.SavedSearchUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the requested document. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSavedSearchUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **savedSearch** | [**SavedSearch**](SavedSearch.md) | The saved search to persist. | 

### Return type

[**SavedSearch**](SavedSearch.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

