# \BrandingAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBrandingItem**](BrandingAPI.md#CreateBrandingItem) | **Post** /brandings | Create a branding item
[**DeleteBranding**](BrandingAPI.md#DeleteBranding) | **Delete** /brandings/{name} | Delete a branding item
[**GetBranding**](BrandingAPI.md#GetBranding) | **Get** /brandings/{name} | Get a branding item
[**GetBrandingList**](BrandingAPI.md#GetBrandingList) | **Get** /brandings | List of branding items
[**SetBrandingItem**](BrandingAPI.md#SetBrandingItem) | **Put** /brandings/{name} | Update a branding item



## Create a branding item


This API endpoint creates a branding item.
A token with API, ORG_ADMIN authority is required to call this API.

### Parameters 
Param Type | Name | Data Type | Required  | Description
------------- | ------------- | ------------- | ------------- | ------------- 
   | name | **string** | True  | name of branding item
   | productName | **string** | True  | product name
   | actionButtonColor | **string** |   (optional) | hex value of color for action button
   | activeLinkColor | **string** |   (optional) | hex value of color for link
   | navigationColor | **string** |   (optional) | hex value of color for navigation bar
   | emailFromAddress | **string** |   (optional) | email from address
   | loginInformationalMessage | **string** |   (optional) | login information message
   | fileStandard | ***os.File** |   (optional) | png file with logo


### Return type

[**BrandingItem**](BrandingItem.md)

### Responses
Code | Description  | Data Type
------------- | ------------- | -------------
201 | Branding item created | BrandingItem
400 | Client Error - Returned if the request body is invalid. | ErrorResponseDto
401 | Unauthorized - Returned if there is no authorization header, or if the JWT token is expired. | ListAccessProfiles401Response
403 | Forbidden - Returned if the user you are running as, doesn&#39;t have access to this end-point. | ErrorResponseDto
429 | Too Many Requests - Returned in response to too many requests in a given period of time - rate limited. The Retry-After header in the response includes how long to wait before trying again. | ListAccessProfiles429Response
500 | Internal Server Error - Returned if there is an unexpected error. | ErrorResponseDto


### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    v3 "github.com/sailpoint-oss/golang-sdk/v2/api_v3"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {

//CreateBrandingItem

    name := "name_example"
    productName := "productName_example"
    //actionButtonColor := "actionButtonColor_example"
    //activeLinkColor := "activeLinkColor_example"
    //navigationColor := "navigationColor_example"
    //emailFromAddress := "emailFromAddress_example"
    //loginInformationalMessage := "loginInformationalMessage_example"
    //fileStandard := os.NewFile(1234, "some_file")



    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.BrandingAPI.CreateBrandingItem(context.Background()).Name(name).ProductName(productName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingAPI.CreateBrandingItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBrandingItem`: BrandingItem
    fmt.Fprintf(os.Stdout, "Response from `BrandingAPI.CreateBrandingItem`: %v\n", resp)
}
```




## Delete a branding item


This API endpoint delete information for an existing branding item by name.
A token with API, ORG_ADMIN authority is required to call this API.

### Parameters 
Param Type | Name | Data Type | Required  | Description
------------- | ------------- | ------------- | ------------- | ------------- 
Path   | name | **string** | True  | The name of the branding item to be deleted


### Return type

 (empty response body)

### Responses
Code | Description  | Data Type
------------- | ------------- | -------------
204 | No content - indicates the request was successful but there is no content to be returned in the response. | 
400 | Client Error - Returned if the request body is invalid. | ErrorResponseDto
401 | Unauthorized - Returned if there is no authorization header, or if the JWT token is expired. | ListAccessProfiles401Response
403 | Forbidden - Returned if the user you are running as, doesn&#39;t have access to this end-point. | ErrorResponseDto
404 | Not Found - returned if the request URL refers to a resource or object that does not exist | ErrorResponseDto
429 | Too Many Requests - Returned in response to too many requests in a given period of time - rate limited. The Retry-After header in the response includes how long to wait before trying again. | ListAccessProfiles429Response
500 | Internal Server Error - Returned if there is an unexpected error. | ErrorResponseDto


### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    v3 "github.com/sailpoint-oss/golang-sdk/v2/api_v3"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {

//DeleteBranding

    name := "default"



    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    r, err := apiClient.V3.BrandingAPI.DeleteBranding(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingAPI.DeleteBranding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```




## Get a branding item


This API endpoint retrieves information for an existing branding item by name.
A token with API, ORG_ADMIN authority is required to call this API.

### Parameters 
Param Type | Name | Data Type | Required  | Description
------------- | ------------- | ------------- | ------------- | ------------- 
Path   | name | **string** | True  | The name of the branding item to be retrieved


### Return type

[**BrandingItem**](BrandingItem.md)

### Responses
Code | Description  | Data Type
------------- | ------------- | -------------
200 | A branding item object | BrandingItem
400 | Client Error - Returned if the request body is invalid. | ErrorResponseDto
401 | Unauthorized - Returned if there is no authorization header, or if the JWT token is expired. | ListAccessProfiles401Response
403 | Forbidden - Returned if the user you are running as, doesn&#39;t have access to this end-point. | ErrorResponseDto
404 | Not Found - returned if the request URL refers to a resource or object that does not exist | ErrorResponseDto
429 | Too Many Requests - Returned in response to too many requests in a given period of time - rate limited. The Retry-After header in the response includes how long to wait before trying again. | ListAccessProfiles429Response
500 | Internal Server Error - Returned if there is an unexpected error. | ErrorResponseDto


### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    v3 "github.com/sailpoint-oss/golang-sdk/v2/api_v3"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {

//GetBranding

    name := "default"



    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.BrandingAPI.GetBranding(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingAPI.GetBranding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBranding`: BrandingItem
    fmt.Fprintf(os.Stdout, "Response from `BrandingAPI.GetBranding`: %v\n", resp)
}
```




## List of branding items


This API endpoint returns a list of branding items.

A token with API, ORG_ADMIN authority is required to call this API.

### Parameters 
Param Type | Name | Data Type | Required  | Description
------------- | ------------- | ------------- | ------------- | ------------- 


### Return type

[**[]BrandingItem**](BrandingItem.md)

### Responses
Code | Description  | Data Type
------------- | ------------- | -------------
200 | A list of branding items. | []BrandingItem
400 | Client Error - Returned if the request body is invalid. | ErrorResponseDto
401 | Unauthorized - Returned if there is no authorization header, or if the JWT token is expired. | ListAccessProfiles401Response
403 | Forbidden - Returned if the user you are running as, doesn&#39;t have access to this end-point. | ErrorResponseDto
429 | Too Many Requests - Returned in response to too many requests in a given period of time - rate limited. The Retry-After header in the response includes how long to wait before trying again. | ListAccessProfiles429Response
500 | Internal Server Error - Returned if there is an unexpected error. | ErrorResponseDto


### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    v3 "github.com/sailpoint-oss/golang-sdk/v2/api_v3"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {

//GetBrandingList




    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.BrandingAPI.GetBrandingList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingAPI.GetBrandingList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBrandingList`: []BrandingItem
    fmt.Fprintf(os.Stdout, "Response from `BrandingAPI.GetBrandingList`: %v\n", resp)
}
```




## Update a branding item


This API endpoint updates information for an existing branding item.
A token with API, ORG_ADMIN authority is required to call this API.

### Parameters 
Param Type | Name | Data Type | Required  | Description
------------- | ------------- | ------------- | ------------- | ------------- 
Path   | name | **string** | True  | The name of the branding item to be retrieved
   | name2 | **string** | True  | name of branding item
   | productName | **string** | True  | product name
   | actionButtonColor | **string** |   (optional) | hex value of color for action button
   | activeLinkColor | **string** |   (optional) | hex value of color for link
   | navigationColor | **string** |   (optional) | hex value of color for navigation bar
   | emailFromAddress | **string** |   (optional) | email from address
   | loginInformationalMessage | **string** |   (optional) | login information message
   | fileStandard | ***os.File** |   (optional) | png file with logo


### Return type

[**BrandingItem**](BrandingItem.md)

### Responses
Code | Description  | Data Type
------------- | ------------- | -------------
200 | Branding item updated | BrandingItem
400 | Client Error - Returned if the request body is invalid. | ErrorResponseDto
401 | Unauthorized - Returned if there is no authorization header, or if the JWT token is expired. | ListAccessProfiles401Response
403 | Forbidden - Returned if the user you are running as, doesn&#39;t have access to this end-point. | ErrorResponseDto
404 | Not Found - returned if the request URL refers to a resource or object that does not exist | ErrorResponseDto
429 | Too Many Requests - Returned in response to too many requests in a given period of time - rate limited. The Retry-After header in the response includes how long to wait before trying again. | ListAccessProfiles429Response
500 | Internal Server Error - Returned if there is an unexpected error. | ErrorResponseDto


### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    v3 "github.com/sailpoint-oss/golang-sdk/v2/api_v3"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {

//SetBrandingItem

    name := "default"
    name2 := "name_example"
    productName := "productName_example"
    //actionButtonColor := "actionButtonColor_example"
    //activeLinkColor := "activeLinkColor_example"
    //navigationColor := "navigationColor_example"
    //emailFromAddress := "emailFromAddress_example"
    //loginInformationalMessage := "loginInformationalMessage_example"
    //fileStandard := os.NewFile(1234, "some_file")



    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.BrandingAPI.SetBrandingItem(context.Background(), name).Name2(name2).ProductName(productName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingAPI.SetBrandingItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetBrandingItem`: BrandingItem
    fmt.Fprintf(os.Stdout, "Response from `BrandingAPI.SetBrandingItem`: %v\n", resp)
}
```



