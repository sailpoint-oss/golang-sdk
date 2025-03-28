# \DataSegmentationAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v2025*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDataSegment**](DataSegmentationAPI.md#CreateDataSegment) | **Post** /data-segments | Create Segment
[**DeleteDataSegment**](DataSegmentationAPI.md#DeleteDataSegment) | **Delete** /data-segments/{segmentId} | Delete Segment by ID
[**GetDataSegment**](DataSegmentationAPI.md#GetDataSegment) | **Get** /data-segments/{segmentId} | Get Segment by ID
[**GetDataSegmentIdentityMembership**](DataSegmentationAPI.md#GetDataSegmentIdentityMembership) | **Get** /data-segments/membership/{identityId} | Get SegmentMembership by Identity ID
[**GetDataSegmentationEnabledForUser**](DataSegmentationAPI.md#GetDataSegmentationEnabledForUser) | **Get** /data-segments/user-enabled/{identityId} | Is Segmentation enabled by Identity
[**ListDataSegments**](DataSegmentationAPI.md#ListDataSegments) | **Get** /data-segments | Get Segments
[**PatchDataSegment**](DataSegmentationAPI.md#PatchDataSegment) | **Patch** /data-segments/{segmentId} | Update Segment
[**PublishDataSegment**](DataSegmentationAPI.md#PublishDataSegment) | **Post** /data-segments/{segmentId} | Publish segment by ID



## CreateDataSegment

> DataSegment CreateDataSegment(ctx).DataSegment(dataSegment).Execute()

Create Segment



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
	dataSegment := *openapiclient.NewDataSegment() // DataSegment | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataSegmentationAPI.CreateDataSegment(context.Background()).DataSegment(dataSegment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSegmentationAPI.CreateDataSegment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDataSegment`: DataSegment
	fmt.Fprintf(os.Stdout, "Response from `DataSegmentationAPI.CreateDataSegment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDataSegmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataSegment** | [**DataSegment**](DataSegment.md) |  | 

### Return type

[**DataSegment**](DataSegment.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth), [applicationAuth](../README.md#applicationAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDataSegment

> DeleteDataSegment(ctx, id).XSailPointExperimental(xSailPointExperimental).Published(published).Execute()

Delete Segment by ID



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
	id := "ef38f943-47e9-4562-b5bb-8424a56397d8" // string | The segment ID to delete.
	xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")
	published := false // bool | This determines which version of the segment to delete (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DataSegmentationAPI.DeleteDataSegment(context.Background(), id).XSailPointExperimental(xSailPointExperimental).Published(published).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSegmentationAPI.DeleteDataSegment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The segment ID to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDataSegmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **published** | **bool** | This determines which version of the segment to delete | [default to false]

### Return type

 (empty response body)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth), [applicationAuth](../README.md#applicationAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataSegment

> DataSegment GetDataSegment(ctx, id).XSailPointExperimental(xSailPointExperimental).Execute()

Get Segment by ID



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
	id := "ef38f943-47e9-4562-b5bb-8424a56397d8" // string | The segment ID to retrieve.
	xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataSegmentationAPI.GetDataSegment(context.Background(), id).XSailPointExperimental(xSailPointExperimental).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSegmentationAPI.GetDataSegment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDataSegment`: DataSegment
	fmt.Fprintf(os.Stdout, "Response from `DataSegmentationAPI.GetDataSegment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The segment ID to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataSegmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]

### Return type

[**DataSegment**](DataSegment.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth), [applicationAuth](../README.md#applicationAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataSegmentIdentityMembership

> map[string]interface{} GetDataSegmentIdentityMembership(ctx, identityId).XSailPointExperimental(xSailPointExperimental).Execute()

Get SegmentMembership by Identity ID



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
	identityId := "ef38f943-47e9-4562-b5bb-8424a56397d8" // string | The identity ID to retrieve the segments they are in.
	xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataSegmentationAPI.GetDataSegmentIdentityMembership(context.Background(), identityId).XSailPointExperimental(xSailPointExperimental).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSegmentationAPI.GetDataSegmentIdentityMembership``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDataSegmentIdentityMembership`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DataSegmentationAPI.GetDataSegmentIdentityMembership`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityId** | **string** | The identity ID to retrieve the segments they are in. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataSegmentIdentityMembershipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]

### Return type

**map[string]interface{}**

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth), [applicationAuth](../README.md#applicationAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataSegmentationEnabledForUser

> bool GetDataSegmentationEnabledForUser(ctx, identityId).XSailPointExperimental(xSailPointExperimental).Execute()

Is Segmentation enabled by Identity



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
	identityId := "ef38f943-47e9-4562-b5bb-8424a56397d8" // string | The identity ID to retrieve if segmentation is enabled for the identity.
	xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataSegmentationAPI.GetDataSegmentationEnabledForUser(context.Background(), identityId).XSailPointExperimental(xSailPointExperimental).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSegmentationAPI.GetDataSegmentationEnabledForUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDataSegmentationEnabledForUser`: bool
	fmt.Fprintf(os.Stdout, "Response from `DataSegmentationAPI.GetDataSegmentationEnabledForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityId** | **string** | The identity ID to retrieve if segmentation is enabled for the identity. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataSegmentationEnabledForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]

### Return type

**bool**

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth), [applicationAuth](../README.md#applicationAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDataSegments

> []DataSegment ListDataSegments(ctx).XSailPointExperimental(xSailPointExperimental).Enabled(enabled).Unique(unique).Published(published).Limit(limit).Offset(offset).Count(count).Filters(filters).Execute()

Get Segments



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
	enabled := true // bool | This boolean indicates whether the segment is currently active. Inactive segments have no effect. (optional) (default to true)
	unique := false // bool | This returns only one record if set to true and that would be the published record if exists. (optional) (default to false)
	published := true // bool | This boolean indicates whether the segment is being applied to the accounts. If unpublished its being actively modified until published (optional) (default to true)
	limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
	offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
	count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
	filters := "name eq """ // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq, in, sw*  **name**: *eq, in, sw* (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataSegmentationAPI.ListDataSegments(context.Background()).XSailPointExperimental(xSailPointExperimental).Enabled(enabled).Unique(unique).Published(published).Limit(limit).Offset(offset).Count(count).Filters(filters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSegmentationAPI.ListDataSegments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDataSegments`: []DataSegment
	fmt.Fprintf(os.Stdout, "Response from `DataSegmentationAPI.ListDataSegments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDataSegmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **enabled** | **bool** | This boolean indicates whether the segment is currently active. Inactive segments have no effect. | [default to true]
 **unique** | **bool** | This returns only one record if set to true and that would be the published record if exists. | [default to false]
 **published** | **bool** | This boolean indicates whether the segment is being applied to the accounts. If unpublished its being actively modified until published | [default to true]
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq, in, sw*  **name**: *eq, in, sw* | 

### Return type

[**[]DataSegment**](DataSegment.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth), [applicationAuth](../README.md#applicationAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchDataSegment

> DataSegment PatchDataSegment(ctx, id).XSailPointExperimental(xSailPointExperimental).RequestBody(requestBody).Execute()

Update Segment



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
	id := "ef38f943-47e9-4562-b5bb-8424a56397d8" // string | The segment ID to modify.
	xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")
	requestBody := []map[string]interface{}{map[string]interface{}(123)} // []map[string]interface{} | A list of segment update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard.  The following fields are patchable: * name * description * membership * memberFilter * memberSelection * scopes * enabled 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataSegmentationAPI.PatchDataSegment(context.Background(), id).XSailPointExperimental(xSailPointExperimental).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSegmentationAPI.PatchDataSegment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchDataSegment`: DataSegment
	fmt.Fprintf(os.Stdout, "Response from `DataSegmentationAPI.PatchDataSegment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The segment ID to modify. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchDataSegmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **requestBody** | **[]map[string]interface{}** | A list of segment update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard.  The following fields are patchable: * name * description * membership * memberFilter * memberSelection * scopes * enabled  | 

### Return type

[**DataSegment**](DataSegment.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth), [applicationAuth](../README.md#applicationAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublishDataSegment

> PublishDataSegment(ctx).XSailPointExperimental(xSailPointExperimental).RequestBody(requestBody).PublishAll(publishAll).Execute()

Publish segment by ID



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
	requestBody := []string{"Property_example"} // []string | A list of segment ids that you wish to publish
	publishAll := true // bool | This flag decides whether you want to publish all unpublished or a list of specific segment ids (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DataSegmentationAPI.PublishDataSegment(context.Background()).XSailPointExperimental(xSailPointExperimental).RequestBody(requestBody).PublishAll(publishAll).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSegmentationAPI.PublishDataSegment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPublishDataSegmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **requestBody** | **[]string** | A list of segment ids that you wish to publish | 
 **publishAll** | **bool** | This flag decides whether you want to publish all unpublished or a list of specific segment ids | [default to true]

### Return type

 (empty response body)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth), [applicationAuth](../README.md#applicationAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

