# \IAIMessageCatalogsAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMessageCatalogs**](IAIMessageCatalogsAPI.md#GetMessageCatalogs) | **Get** /translation-catalogs/{catalog-id} | Get Message catalogs



## GetMessageCatalogs

> []MessageCatalogDto GetMessageCatalogs(ctx, catalogId).Execute()

Get Message catalogs



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
	catalogId := "recommender" // string | The ID of the message catalog.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IAIMessageCatalogsAPI.GetMessageCatalogs(context.Background(), catalogId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IAIMessageCatalogsAPI.GetMessageCatalogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMessageCatalogs`: []MessageCatalogDto
	fmt.Fprintf(os.Stdout, "Response from `IAIMessageCatalogsAPI.GetMessageCatalogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**catalogId** | **string** | The ID of the message catalog. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMessageCatalogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]MessageCatalogDto**](MessageCatalogDto.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

