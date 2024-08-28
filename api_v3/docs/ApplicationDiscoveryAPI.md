# \ApplicationDiscoveryAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDiscoveredApplications**](ApplicationDiscoveryAPI.md#GetDiscoveredApplications) | **Get** /discovered-applications | Retrieve discovered applications for tenant
[**GetManualDiscoverApplicationsCsvTemplate**](ApplicationDiscoveryAPI.md#GetManualDiscoverApplicationsCsvTemplate) | **Get** /manual-discover-applications-template | CSV template download for discovery
[**GetVendorConnectorMappings**](ApplicationDiscoveryAPI.md#GetVendorConnectorMappings) | **Get** /vendor-connector-mappings | List vendor connector mappings
[**SendManualDiscoverApplicationsCsvTemplate**](ApplicationDiscoveryAPI.md#SendManualDiscoverApplicationsCsvTemplate) | **Post** /manual-discover-applications | CSV Upload to discover applications



## GetDiscoveredApplications

> []GetDiscoveredApplications200ResponseInner GetDiscoveredApplications(ctx).Limit(limit).Offset(offset).Detail(detail).Filter(filter).Sorters(sorters).Execute()

Retrieve discovered applications for tenant



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
	limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
	offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
	detail := "FULL" // string | Determines whether slim, or increased level of detail is provided for each discovered application in the returned list. SLIM is the default behavior. (optional)
	filter := "name eq "Okta" and description co "Okta" and discoverySource in ("csv", "Okta Saas")" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)       Filtering is supported for the following fields and operators:  **name**: *eq, sw, co*  **description**: *eq, sw, co*  **createdAtStart**: *eq, le, ge*  **createdAtEnd**: *eq, le, ge*  **discoveredAtStart**: *eq, le, ge*  **discoveredAtEnd**: *eq, le, ge*  **discoverySource**: *eq, in*  (optional)
	sorters := "name" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **name, description, discoveredAt, discoverySource** (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationDiscoveryAPI.GetDiscoveredApplications(context.Background()).Limit(limit).Offset(offset).Detail(detail).Filter(filter).Sorters(sorters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationDiscoveryAPI.GetDiscoveredApplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDiscoveredApplications`: []GetDiscoveredApplications200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ApplicationDiscoveryAPI.GetDiscoveredApplications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDiscoveredApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **detail** | **string** | Determines whether slim, or increased level of detail is provided for each discovered application in the returned list. SLIM is the default behavior. | 
 **filter** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)       Filtering is supported for the following fields and operators:  **name**: *eq, sw, co*  **description**: *eq, sw, co*  **createdAtStart**: *eq, le, ge*  **createdAtEnd**: *eq, le, ge*  **discoveredAtStart**: *eq, le, ge*  **discoveredAtEnd**: *eq, le, ge*  **discoverySource**: *eq, in*  | 
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **name, description, discoveredAt, discoverySource** | 

### Return type

[**[]GetDiscoveredApplications200ResponseInner**](GetDiscoveredApplications200ResponseInner.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManualDiscoverApplicationsCsvTemplate

> ManualDiscoverApplicationsTemplate GetManualDiscoverApplicationsCsvTemplate(ctx).Execute()

CSV template download for discovery



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationDiscoveryAPI.GetManualDiscoverApplicationsCsvTemplate(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationDiscoveryAPI.GetManualDiscoverApplicationsCsvTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetManualDiscoverApplicationsCsvTemplate`: ManualDiscoverApplicationsTemplate
	fmt.Fprintf(os.Stdout, "Response from `ApplicationDiscoveryAPI.GetManualDiscoverApplicationsCsvTemplate`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetManualDiscoverApplicationsCsvTemplateRequest struct via the builder pattern


### Return type

[**ManualDiscoverApplicationsTemplate**](ManualDiscoverApplicationsTemplate.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/csv, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVendorConnectorMappings

> []VendorConnectorMapping GetVendorConnectorMappings(ctx).Execute()

List vendor connector mappings



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationDiscoveryAPI.GetVendorConnectorMappings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationDiscoveryAPI.GetVendorConnectorMappings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVendorConnectorMappings`: []VendorConnectorMapping
	fmt.Fprintf(os.Stdout, "Response from `ApplicationDiscoveryAPI.GetVendorConnectorMappings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVendorConnectorMappingsRequest struct via the builder pattern


### Return type

[**[]VendorConnectorMapping**](VendorConnectorMapping.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendManualDiscoverApplicationsCsvTemplate

> SendManualDiscoverApplicationsCsvTemplate(ctx).File(file).Execute()

CSV Upload to discover applications



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
	file := os.NewFile(1234, "some_file") // *os.File | The CSV file to upload containing `application_name` and `description` columns. Each row represents an application to be discovered.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ApplicationDiscoveryAPI.SendManualDiscoverApplicationsCsvTemplate(context.Background()).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationDiscoveryAPI.SendManualDiscoverApplicationsCsvTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendManualDiscoverApplicationsCsvTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** | The CSV file to upload containing &#x60;application_name&#x60; and &#x60;description&#x60; columns. Each row represents an application to be discovered. | 

### Return type

 (empty response body)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

