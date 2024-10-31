# \VendorConnectorMappingsAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVendorConnectorMapping**](VendorConnectorMappingsAPI.md#CreateVendorConnectorMapping) | **Post** /vendor-connector-mappings | Create Vendor Connector Mapping
[**DeleteVendorConnectorMapping**](VendorConnectorMappingsAPI.md#DeleteVendorConnectorMapping) | **Delete** /vendor-connector-mappings | Delete Vendor Connector Mapping
[**GetVendorConnectorMappings**](VendorConnectorMappingsAPI.md#GetVendorConnectorMappings) | **Get** /vendor-connector-mappings | List Vendor Connector Mappings



## CreateVendorConnectorMapping

> VendorConnectorMapping CreateVendorConnectorMapping(ctx).VendorConnectorMapping(vendorConnectorMapping).Execute()

Create Vendor Connector Mapping



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
	vendorConnectorMapping := *openapiclient.NewVendorConnectorMapping() // VendorConnectorMapping | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VendorConnectorMappingsAPI.CreateVendorConnectorMapping(context.Background()).VendorConnectorMapping(vendorConnectorMapping).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VendorConnectorMappingsAPI.CreateVendorConnectorMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVendorConnectorMapping`: VendorConnectorMapping
	fmt.Fprintf(os.Stdout, "Response from `VendorConnectorMappingsAPI.CreateVendorConnectorMapping`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVendorConnectorMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vendorConnectorMapping** | [**VendorConnectorMapping**](VendorConnectorMapping.md) |  | 

### Return type

[**VendorConnectorMapping**](VendorConnectorMapping.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVendorConnectorMapping

> DeleteVendorConnectorMapping200Response DeleteVendorConnectorMapping(ctx).VendorConnectorMapping(vendorConnectorMapping).Execute()

Delete Vendor Connector Mapping



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
	vendorConnectorMapping := *openapiclient.NewVendorConnectorMapping() // VendorConnectorMapping | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VendorConnectorMappingsAPI.DeleteVendorConnectorMapping(context.Background()).VendorConnectorMapping(vendorConnectorMapping).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VendorConnectorMappingsAPI.DeleteVendorConnectorMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteVendorConnectorMapping`: DeleteVendorConnectorMapping200Response
	fmt.Fprintf(os.Stdout, "Response from `VendorConnectorMappingsAPI.DeleteVendorConnectorMapping`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVendorConnectorMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vendorConnectorMapping** | [**VendorConnectorMapping**](VendorConnectorMapping.md) |  | 

### Return type

[**DeleteVendorConnectorMapping200Response**](DeleteVendorConnectorMapping200Response.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVendorConnectorMappings

> []VendorConnectorMapping GetVendorConnectorMappings(ctx).Execute()

List Vendor Connector Mappings



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
	resp, r, err := apiClient.VendorConnectorMappingsAPI.GetVendorConnectorMappings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VendorConnectorMappingsAPI.GetVendorConnectorMappings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVendorConnectorMappings`: []VendorConnectorMapping
	fmt.Fprintf(os.Stdout, "Response from `VendorConnectorMappingsAPI.GetVendorConnectorMappings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVendorConnectorMappingsRequest struct via the builder pattern


### Return type

[**[]VendorConnectorMapping**](VendorConnectorMapping.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

