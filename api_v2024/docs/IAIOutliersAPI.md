# \IAIOutliersAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v2024*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportOutliersZip**](IAIOutliersAPI.md#ExportOutliersZip) | **Get** /outliers/export | IAI Identity Outliers Export
[**GetIdentityOutlierSnapshots**](IAIOutliersAPI.md#GetIdentityOutlierSnapshots) | **Get** /outlier-summaries | IAI Identity Outliers Summary
[**GetIdentityOutliers**](IAIOutliersAPI.md#GetIdentityOutliers) | **Get** /outliers | IAI Get Identity Outliers
[**GetLatestIdentityOutlierSnapshots**](IAIOutliersAPI.md#GetLatestIdentityOutlierSnapshots) | **Get** /outlier-summaries/latest | IAI Identity Outliers Latest Summary
[**GetOutlierContributingFeatureSummary**](IAIOutliersAPI.md#GetOutlierContributingFeatureSummary) | **Get** /outlier-feature-summaries/{outlierFeatureId} | Get identity outlier contibuting feature summary
[**GetPeerGroupOutliersContributingFeatures**](IAIOutliersAPI.md#GetPeerGroupOutliersContributingFeatures) | **Get** /outliers/{outlierId}/contributing-features | Get identity outlier&#39;s contibuting features
[**IgnoreIdentityOutliers**](IAIOutliersAPI.md#IgnoreIdentityOutliers) | **Post** /outliers/ignore | IAI Identity Outliers Ignore
[**ListOutliersContributingFeatureAccessItems**](IAIOutliersAPI.md#ListOutliersContributingFeatureAccessItems) | **Get** /outliers/{outlierId}/feature-details/{contributingFeatureName}/access-items | Gets a list of access items associated with each identity outlier contributing feature
[**UnIgnoreIdentityOutliers**](IAIOutliersAPI.md#UnIgnoreIdentityOutliers) | **Post** /outliers/unignore | IAI Identity Outliers Unignore



## ExportOutliersZip

> *os.File ExportOutliersZip(ctx).XSailPointExperimental(xSailPointExperimental).Type_(type_).Execute()

IAI Identity Outliers Export



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
	xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")
	type_ := "type__example" // string | Type of the identity outliers snapshot to filter on (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IAIOutliersAPI.ExportOutliersZip(context.Background()).XSailPointExperimental(xSailPointExperimental).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IAIOutliersAPI.ExportOutliersZip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportOutliersZip`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `IAIOutliersAPI.ExportOutliersZip`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportOutliersZipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **type_** | **string** | Type of the identity outliers snapshot to filter on | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/zip, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityOutlierSnapshots

> []OutlierSummary GetIdentityOutlierSnapshots(ctx).XSailPointExperimental(xSailPointExperimental).Limit(limit).Offset(offset).Type_(type_).Filters(filters).Sorters(sorters).Execute()

IAI Identity Outliers Summary



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
	xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")
	limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
	offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
	type_ := "type__example" // string | Type of the identity outliers snapshot to filter on (optional)
	filters := "snapshotDate ge "2022-02-07T20:13:29.356648026Z"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **snapshotDate**: *ge, le* (optional)
	sorters := "snapshotDate" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **snapshotDate** (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IAIOutliersAPI.GetIdentityOutlierSnapshots(context.Background()).XSailPointExperimental(xSailPointExperimental).Limit(limit).Offset(offset).Type_(type_).Filters(filters).Sorters(sorters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IAIOutliersAPI.GetIdentityOutlierSnapshots``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIdentityOutlierSnapshots`: []OutlierSummary
	fmt.Fprintf(os.Stdout, "Response from `IAIOutliersAPI.GetIdentityOutlierSnapshots`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityOutlierSnapshotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **type_** | **string** | Type of the identity outliers snapshot to filter on | 
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **snapshotDate**: *ge, le* | 
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **snapshotDate** | 

### Return type

[**[]OutlierSummary**](OutlierSummary.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityOutliers

> []Outlier GetIdentityOutliers(ctx).XSailPointExperimental(xSailPointExperimental).Limit(limit).Offset(offset).Count(count).Type_(type_).Filters(filters).Sorters(sorters).Execute()

IAI Get Identity Outliers



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
	xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")
	limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
	offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
	count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
	type_ := "LOW_SIMILARITY" // string | Type of the identity outliers snapshot to filter on (optional)
	filters := "attributes.displayName sw "John" and certStatus eq "false"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **attributes**: *eq, sw, co, in*  **firstDetectionDate**: *ge, le*  **certStatus**: *eq*  **ignored**: *eq*  **score**: *ge, le* (optional)
	sorters := "attributes.displayName,firstDetectionDate,-score" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **firstDetectionDate, attributes, score** (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IAIOutliersAPI.GetIdentityOutliers(context.Background()).XSailPointExperimental(xSailPointExperimental).Limit(limit).Offset(offset).Count(count).Type_(type_).Filters(filters).Sorters(sorters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IAIOutliersAPI.GetIdentityOutliers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIdentityOutliers`: []Outlier
	fmt.Fprintf(os.Stdout, "Response from `IAIOutliersAPI.GetIdentityOutliers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityOutliersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **type_** | **string** | Type of the identity outliers snapshot to filter on | 
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **attributes**: *eq, sw, co, in*  **firstDetectionDate**: *ge, le*  **certStatus**: *eq*  **ignored**: *eq*  **score**: *ge, le* | 
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **firstDetectionDate, attributes, score** | 

### Return type

[**[]Outlier**](Outlier.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLatestIdentityOutlierSnapshots

> []LatestOutlierSummary GetLatestIdentityOutlierSnapshots(ctx).XSailPointExperimental(xSailPointExperimental).Type_(type_).Execute()

IAI Identity Outliers Latest Summary



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
	xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")
	type_ := "type__example" // string | Type of the identity outliers snapshot to filter on (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IAIOutliersAPI.GetLatestIdentityOutlierSnapshots(context.Background()).XSailPointExperimental(xSailPointExperimental).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IAIOutliersAPI.GetLatestIdentityOutlierSnapshots``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLatestIdentityOutlierSnapshots`: []LatestOutlierSummary
	fmt.Fprintf(os.Stdout, "Response from `IAIOutliersAPI.GetLatestIdentityOutlierSnapshots`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLatestIdentityOutlierSnapshotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **type_** | **string** | Type of the identity outliers snapshot to filter on | 

### Return type

[**[]LatestOutlierSummary**](LatestOutlierSummary.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOutlierContributingFeatureSummary

> OutlierFeatureSummary GetOutlierContributingFeatureSummary(ctx, outlierFeatureId).XSailPointExperimental(xSailPointExperimental).Execute()

Get identity outlier contibuting feature summary



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
	outlierFeatureId := "04654b66-7561-4090-94f9-abee0722a1af" // string | Contributing feature id
	xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IAIOutliersAPI.GetOutlierContributingFeatureSummary(context.Background(), outlierFeatureId).XSailPointExperimental(xSailPointExperimental).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IAIOutliersAPI.GetOutlierContributingFeatureSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOutlierContributingFeatureSummary`: OutlierFeatureSummary
	fmt.Fprintf(os.Stdout, "Response from `IAIOutliersAPI.GetOutlierContributingFeatureSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**outlierFeatureId** | **string** | Contributing feature id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOutlierContributingFeatureSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]

### Return type

[**OutlierFeatureSummary**](OutlierFeatureSummary.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPeerGroupOutliersContributingFeatures

> []OutlierContributingFeature GetPeerGroupOutliersContributingFeatures(ctx, outlierId).XSailPointExperimental(xSailPointExperimental).Limit(limit).Offset(offset).Count(count).IncludeTranslationMessages(includeTranslationMessages).Sorters(sorters).Execute()

Get identity outlier's contibuting features



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
	outlierId := "2c918085842e69ae018432d22ccb212f" // string | The outlier id
	xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")
	limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
	offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
	count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
	includeTranslationMessages := "include-translation-messages=" // string | Whether or not to include translation messages object in returned response (optional)
	sorters := "importance" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **importance** (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IAIOutliersAPI.GetPeerGroupOutliersContributingFeatures(context.Background(), outlierId).XSailPointExperimental(xSailPointExperimental).Limit(limit).Offset(offset).Count(count).IncludeTranslationMessages(includeTranslationMessages).Sorters(sorters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IAIOutliersAPI.GetPeerGroupOutliersContributingFeatures``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPeerGroupOutliersContributingFeatures`: []OutlierContributingFeature
	fmt.Fprintf(os.Stdout, "Response from `IAIOutliersAPI.GetPeerGroupOutliersContributingFeatures`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**outlierId** | **string** | The outlier id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPeerGroupOutliersContributingFeaturesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **includeTranslationMessages** | **string** | Whether or not to include translation messages object in returned response | 
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **importance** | 

### Return type

[**[]OutlierContributingFeature**](OutlierContributingFeature.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IgnoreIdentityOutliers

> IgnoreIdentityOutliers(ctx).XSailPointExperimental(xSailPointExperimental).RequestBody(requestBody).Execute()

IAI Identity Outliers Ignore



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
	xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")
	requestBody := []string{"Property_example"} // []string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IAIOutliersAPI.IgnoreIdentityOutliers(context.Background()).XSailPointExperimental(xSailPointExperimental).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IAIOutliersAPI.IgnoreIdentityOutliers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIgnoreIdentityOutliersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **requestBody** | **[]string** |  | 

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


## ListOutliersContributingFeatureAccessItems

> []OutliersContributingFeatureAccessItems ListOutliersContributingFeatureAccessItems(ctx, outlierId, contributingFeatureName).XSailPointExperimental(xSailPointExperimental).Limit(limit).Offset(offset).Count(count).AccessType(accessType).Sorters(sorters).Execute()

Gets a list of access items associated with each identity outlier contributing feature



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
	outlierId := "2c918085842e69ae018432d22ccb212f" // string | The outlier id
	contributingFeatureName := "entitlement_count" // string | The name of contributing feature
	xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")
	limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
	offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
	count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
	accessType := "ENTITLEMENT" // string | The type of access item for the identity outlier contributing feature. If not provided, it returns all (optional)
	sorters := "displayName" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **displayName** (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IAIOutliersAPI.ListOutliersContributingFeatureAccessItems(context.Background(), outlierId, contributingFeatureName).XSailPointExperimental(xSailPointExperimental).Limit(limit).Offset(offset).Count(count).AccessType(accessType).Sorters(sorters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IAIOutliersAPI.ListOutliersContributingFeatureAccessItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOutliersContributingFeatureAccessItems`: []OutliersContributingFeatureAccessItems
	fmt.Fprintf(os.Stdout, "Response from `IAIOutliersAPI.ListOutliersContributingFeatureAccessItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**outlierId** | **string** | The outlier id | 
**contributingFeatureName** | **string** | The name of contributing feature | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOutliersContributingFeatureAccessItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **accessType** | **string** | The type of access item for the identity outlier contributing feature. If not provided, it returns all | 
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **displayName** | 

### Return type

[**[]OutliersContributingFeatureAccessItems**](OutliersContributingFeatureAccessItems.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnIgnoreIdentityOutliers

> UnIgnoreIdentityOutliers(ctx).XSailPointExperimental(xSailPointExperimental).RequestBody(requestBody).Execute()

IAI Identity Outliers Unignore



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
	xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")
	requestBody := []string{"Property_example"} // []string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IAIOutliersAPI.UnIgnoreIdentityOutliers(context.Background()).XSailPointExperimental(xSailPointExperimental).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IAIOutliersAPI.UnIgnoreIdentityOutliers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnIgnoreIdentityOutliersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **requestBody** | **[]string** |  | 

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

