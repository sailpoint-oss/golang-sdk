# \UIMetadataAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTenantUiMetadata**](UIMetadataAPI.md#GetTenantUiMetadata) | **Get** /ui-metadata/tenant | Get a tenant UI metadata
[**SetTenantUiMetadata**](UIMetadataAPI.md#SetTenantUiMetadata) | **Put** /ui-metadata/tenant | Update tenant UI metadata



## GetTenantUiMetadata

> TenantUiMetadataItemResponse GetTenantUiMetadata(ctx).Execute()

Get a tenant UI metadata



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
	resp, r, err := apiClient.UIMetadataAPI.GetTenantUiMetadata(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UIMetadataAPI.GetTenantUiMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTenantUiMetadata`: TenantUiMetadataItemResponse
	fmt.Fprintf(os.Stdout, "Response from `UIMetadataAPI.GetTenantUiMetadata`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantUiMetadataRequest struct via the builder pattern


### Return type

[**TenantUiMetadataItemResponse**](TenantUiMetadataItemResponse.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetTenantUiMetadata

> TenantUiMetadataItemResponse SetTenantUiMetadata(ctx).TenantUiMetadataItemUpdateRequest(tenantUiMetadataItemUpdateRequest).Execute()

Update tenant UI metadata



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
	tenantUiMetadataItemUpdateRequest := *openapiclient.NewTenantUiMetadataItemUpdateRequest() // TenantUiMetadataItemUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UIMetadataAPI.SetTenantUiMetadata(context.Background()).TenantUiMetadataItemUpdateRequest(tenantUiMetadataItemUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UIMetadataAPI.SetTenantUiMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetTenantUiMetadata`: TenantUiMetadataItemResponse
	fmt.Fprintf(os.Stdout, "Response from `UIMetadataAPI.SetTenantUiMetadata`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetTenantUiMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantUiMetadataItemUpdateRequest** | [**TenantUiMetadataItemUpdateRequest**](TenantUiMetadataItemUpdateRequest.md) |  | 

### Return type

[**TenantUiMetadataItemResponse**](TenantUiMetadataItemResponse.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

