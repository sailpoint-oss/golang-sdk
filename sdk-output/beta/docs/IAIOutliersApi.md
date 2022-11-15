# \IAIOutliersApi

All URIs are relative to *https://sailpoint.api.identitynow.com/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportOutliersZip**](IAIOutliersApi.md#ExportOutliersZip) | **Get** /outliers/export | IAI Identity Outliers Export
[**GetLatestOutlierSnapshots**](IAIOutliersApi.md#GetLatestOutlierSnapshots) | **Get** /outlier-summaries/latest | IAI Identity Outliers Latest Summary
[**GetOutlierSnapshots**](IAIOutliersApi.md#GetOutlierSnapshots) | **Get** /outlier-summaries | IAI Identity Outliers Summary
[**GetOutliers**](IAIOutliersApi.md#GetOutliers) | **Get** /outliers | IAI Get Identity Outliers
[**IgnoreOutliers**](IAIOutliersApi.md#IgnoreOutliers) | **Post** /outliers/ignore | IAI Identity Outliers Ignore
[**UnIgnoreOutliers**](IAIOutliersApi.md#UnIgnoreOutliers) | **Post** /outliers/unignore | IAI Identity Outliers Unignore



## ExportOutliersZip

> *os.File ExportOutliersZip(ctx).Type_(type_).Execute()

IAI Identity Outliers Export



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
    type_ := "type__example" // string | Type of the identity outliers snapshot to filter on (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IAIOutliersApi.ExportOutliersZip(context.Background()).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAIOutliersApi.ExportOutliersZip``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportOutliersZip`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `IAIOutliersApi.ExportOutliersZip`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportOutliersZipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | Type of the identity outliers snapshot to filter on | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/zip, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLatestOutlierSnapshots

> []LatestOutlierSummary GetLatestOutlierSnapshots(ctx).Type_(type_).Execute()

IAI Identity Outliers Latest Summary



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
    type_ := "type__example" // string | Type of the identity outliers snapshot to filter on (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IAIOutliersApi.GetLatestOutlierSnapshots(context.Background()).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAIOutliersApi.GetLatestOutlierSnapshots``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLatestOutlierSnapshots`: []LatestOutlierSummary
    fmt.Fprintf(os.Stdout, "Response from `IAIOutliersApi.GetLatestOutlierSnapshots`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLatestOutlierSnapshotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | Type of the identity outliers snapshot to filter on | 

### Return type

[**[]LatestOutlierSummary**](LatestOutlierSummary.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOutlierSnapshots

> []OutlierSummary GetOutlierSnapshots(ctx).Limit(limit).Offset(offset).Type_(type_).Filters(filters).Sorters(sorters).Execute()

IAI Identity Outliers Summary



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
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    type_ := "type__example" // string | Type of the identity outliers snapshot to filter on (optional)
    filters := "filters_example" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results) Filtering is supported for the following field and operators:  **snapshotDate**: *ge, le* (optional)
    sorters := "snapshotDate" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results) Sorting is supported for the following field: **snapshotDate** (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IAIOutliersApi.GetOutlierSnapshots(context.Background()).Limit(limit).Offset(offset).Type_(type_).Filters(filters).Sorters(sorters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAIOutliersApi.GetOutlierSnapshots``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOutlierSnapshots`: []OutlierSummary
    fmt.Fprintf(os.Stdout, "Response from `IAIOutliersApi.GetOutlierSnapshots`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOutlierSnapshotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **type_** | **string** | Type of the identity outliers snapshot to filter on | 
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results) Filtering is supported for the following field and operators:  **snapshotDate**: *ge, le* | 
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results) Sorting is supported for the following field: **snapshotDate** | 

### Return type

[**[]OutlierSummary**](OutlierSummary.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOutliers

> []Outlier GetOutliers(ctx).Limit(limit).Offset(offset).Count(count).Type_(type_).Filters(filters).Sorters(sorters).Execute()

IAI Get Identity Outliers



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
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
    type_ := "LOW_SIMILARITY" // string | Type of the identity outliers snapshot to filter on (optional)
    filters := "attributes.displayName sw "John" and certStatus eq "false"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following Entitlement fields and operators: **attributes**: *eq, sw, co, in*  **firstDetectionDate**: *ge, le*  **certStatus**: *eq*  **ignored**: *eq*  **score**: *ge, le* (optional)
    sorters := "attributes.displayName,firstDetectionDate,-score" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results) Sorting is supported for the following fields: **firstDetectionDate, attributes, score** (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IAIOutliersApi.GetOutliers(context.Background()).Limit(limit).Offset(offset).Count(count).Type_(type_).Filters(filters).Sorters(sorters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAIOutliersApi.GetOutliers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOutliers`: []Outlier
    fmt.Fprintf(os.Stdout, "Response from `IAIOutliersApi.GetOutliers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOutliersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **type_** | **string** | Type of the identity outliers snapshot to filter on | 
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following Entitlement fields and operators: **attributes**: *eq, sw, co, in*  **firstDetectionDate**: *ge, le*  **certStatus**: *eq*  **ignored**: *eq*  **score**: *ge, le* | 
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results) Sorting is supported for the following fields: **firstDetectionDate, attributes, score** | 

### Return type

[**[]Outlier**](Outlier.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IgnoreOutliers

> IgnoreOutliers(ctx).RequestBody(requestBody).Execute()

IAI Identity Outliers Ignore



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
    requestBody := []string{"Property_example"} // []string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IAIOutliersApi.IgnoreOutliers(context.Background()).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAIOutliersApi.IgnoreOutliers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIgnoreOutliersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]string** |  | 

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


## UnIgnoreOutliers

> UnIgnoreOutliers(ctx).RequestBody(requestBody).Execute()

IAI Identity Outliers Unignore



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
    requestBody := []string{"Property_example"} // []string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IAIOutliersApi.UnIgnoreOutliers(context.Background()).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAIOutliersApi.UnIgnoreOutliers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnIgnoreOutliersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]string** |  | 

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

