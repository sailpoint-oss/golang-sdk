# \BrandingAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBrandingItem**](BrandingAPI.md#CreateBrandingItem) | **Post** /brandings | Create a branding item
[**DeleteBranding**](BrandingAPI.md#DeleteBranding) | **Delete** /brandings/{name} | Delete a branding item
[**GetBranding**](BrandingAPI.md#GetBranding) | **Get** /brandings/{name} | Get a branding item
[**GetBrandingList**](BrandingAPI.md#GetBrandingList) | **Get** /brandings | List of branding items
[**SetBrandingItem**](BrandingAPI.md#SetBrandingItem) | **Put** /brandings/{name} | Update a branding item



## CreateBrandingItem

> BrandingItem CreateBrandingItem(ctx).Name(name).ProductName(productName).ActionButtonColor(actionButtonColor).ActiveLinkColor(activeLinkColor).NavigationColor(navigationColor).EmailFromAddress(emailFromAddress).LoginInformationalMessage(loginInformationalMessage).FileStandard(fileStandard).Execute()

Create a branding item



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
    name := "name_example" // string | name of branding item
    productName := "productName_example" // string | product name
    actionButtonColor := "actionButtonColor_example" // string | hex value of color for action button (optional)
    activeLinkColor := "activeLinkColor_example" // string | hex value of color for link (optional)
    navigationColor := "navigationColor_example" // string | hex value of color for navigation bar (optional)
    emailFromAddress := "emailFromAddress_example" // string | email from address (optional)
    loginInformationalMessage := "loginInformationalMessage_example" // string | login information message (optional)
    fileStandard := os.NewFile(1234, "some_file") // *os.File | png file with logo (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BrandingAPI.CreateBrandingItem(context.Background()).Name(name).ProductName(productName).ActionButtonColor(actionButtonColor).ActiveLinkColor(activeLinkColor).NavigationColor(navigationColor).EmailFromAddress(emailFromAddress).LoginInformationalMessage(loginInformationalMessage).FileStandard(fileStandard).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingAPI.CreateBrandingItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBrandingItem`: BrandingItem
    fmt.Fprintf(os.Stdout, "Response from `BrandingAPI.CreateBrandingItem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBrandingItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | name of branding item | 
 **productName** | **string** | product name | 
 **actionButtonColor** | **string** | hex value of color for action button | 
 **activeLinkColor** | **string** | hex value of color for link | 
 **navigationColor** | **string** | hex value of color for navigation bar | 
 **emailFromAddress** | **string** | email from address | 
 **loginInformationalMessage** | **string** | login information message | 
 **fileStandard** | ***os.File** | png file with logo | 

### Return type

[**BrandingItem**](BrandingItem.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBranding

> DeleteBranding(ctx, name).Execute()

Delete a branding item



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
    name := "default" // string | The name of the branding item to be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BrandingAPI.DeleteBranding(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingAPI.DeleteBranding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the branding item to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBrandingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetBranding

> BrandingItem GetBranding(ctx, name).Execute()

Get a branding item



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
    name := "default" // string | The name of the branding item to be retrieved

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BrandingAPI.GetBranding(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingAPI.GetBranding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBranding`: BrandingItem
    fmt.Fprintf(os.Stdout, "Response from `BrandingAPI.GetBranding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the branding item to be retrieved | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBrandingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BrandingItem**](BrandingItem.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBrandingList

> []BrandingItem GetBrandingList(ctx).Execute()

List of branding items



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
    resp, r, err := apiClient.BrandingAPI.GetBrandingList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingAPI.GetBrandingList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBrandingList`: []BrandingItem
    fmt.Fprintf(os.Stdout, "Response from `BrandingAPI.GetBrandingList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBrandingListRequest struct via the builder pattern


### Return type

[**[]BrandingItem**](BrandingItem.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetBrandingItem

> BrandingItem SetBrandingItem(ctx, name).Name2(name2).ProductName(productName).ActionButtonColor(actionButtonColor).ActiveLinkColor(activeLinkColor).NavigationColor(navigationColor).EmailFromAddress(emailFromAddress).LoginInformationalMessage(loginInformationalMessage).FileStandard(fileStandard).Execute()

Update a branding item



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
    name := "default" // string | The name of the branding item to be retrieved
    name2 := "name_example" // string | name of branding item
    productName := "productName_example" // string | product name
    actionButtonColor := "actionButtonColor_example" // string | hex value of color for action button (optional)
    activeLinkColor := "activeLinkColor_example" // string | hex value of color for link (optional)
    navigationColor := "navigationColor_example" // string | hex value of color for navigation bar (optional)
    emailFromAddress := "emailFromAddress_example" // string | email from address (optional)
    loginInformationalMessage := "loginInformationalMessage_example" // string | login information message (optional)
    fileStandard := os.NewFile(1234, "some_file") // *os.File | png file with logo (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BrandingAPI.SetBrandingItem(context.Background(), name).Name2(name2).ProductName(productName).ActionButtonColor(actionButtonColor).ActiveLinkColor(activeLinkColor).NavigationColor(navigationColor).EmailFromAddress(emailFromAddress).LoginInformationalMessage(loginInformationalMessage).FileStandard(fileStandard).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingAPI.SetBrandingItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetBrandingItem`: BrandingItem
    fmt.Fprintf(os.Stdout, "Response from `BrandingAPI.SetBrandingItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the branding item to be retrieved | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetBrandingItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name2** | **string** | name of branding item | 
 **productName** | **string** | product name | 
 **actionButtonColor** | **string** | hex value of color for action button | 
 **activeLinkColor** | **string** | hex value of color for link | 
 **navigationColor** | **string** | hex value of color for navigation bar | 
 **emailFromAddress** | **string** | email from address | 
 **loginInformationalMessage** | **string** | login information message | 
 **fileStandard** | ***os.File** | png file with logo | 

### Return type

[**BrandingItem**](BrandingItem.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

