# \IconsAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v2024*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIcon**](IconsAPI.md#DeleteIcon) | **Delete** /icons/{objectType}/{objectId} | Delete an icon
[**SetIcon**](IconsAPI.md#SetIcon) | **Put** /icons/{objectType}/{objectId} | Update an icon



## DeleteIcon

> DeleteIcon(ctx, objectType, objectId).XSailPointExperimental(xSailPointExperimental).Execute()

Delete an icon



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
	objectType := "application" // string | Object type. Available options ['application']
	objectId := "a291e870-48c3-4953-b656-fb5ce2a93169" // string | Object id.
	xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IconsAPI.DeleteIcon(context.Background(), objectType, objectId).XSailPointExperimental(xSailPointExperimental).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IconsAPI.DeleteIcon``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** | Object type. Available options [&#39;application&#39;] | 
**objectId** | **string** | Object id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIconRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]

### Return type

 (empty response body)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetIcon

> SetIcon200Response SetIcon(ctx, objectType, objectId).XSailPointExperimental(xSailPointExperimental).Image(image).Execute()

Update an icon



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
	objectType := "application" // string | Object type. Available options ['application']
	objectId := "a291e870-48c3-4953-b656-fb5ce2a93169" // string | Object id.
	xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")
	image := os.NewFile(1234, "some_file") // *os.File | file with icon. Allowed mime-types ['image/png', 'image/jpeg']

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IconsAPI.SetIcon(context.Background(), objectType, objectId).XSailPointExperimental(xSailPointExperimental).Image(image).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IconsAPI.SetIcon``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetIcon`: SetIcon200Response
	fmt.Fprintf(os.Stdout, "Response from `IconsAPI.SetIcon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** | Object type. Available options [&#39;application&#39;] | 
**objectId** | **string** | Object id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetIconRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **image** | ***os.File** | file with icon. Allowed mime-types [&#39;image/png&#39;, &#39;image/jpeg&#39;] | 

### Return type

[**SetIcon200Response**](SetIcon200Response.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

