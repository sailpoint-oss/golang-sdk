# \VendorConnectorMapppingAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v2024*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVendorConnectorMappings**](VendorConnectorMapppingAPI.md#GetVendorConnectorMappings) | **Get** /vendor-connector-mappings | List Vendor Connector Mappings



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
	resp, r, err := apiClient.VendorConnectorMapppingAPI.GetVendorConnectorMappings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VendorConnectorMapppingAPI.GetVendorConnectorMappings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVendorConnectorMappings`: []VendorConnectorMapping
	fmt.Fprintf(os.Stdout, "Response from `VendorConnectorMapppingAPI.GetVendorConnectorMappings`: %v\n", resp)
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

