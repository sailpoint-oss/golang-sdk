# \IdentityAttributesAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIdentityAttribute**](IdentityAttributesAPI.md#CreateIdentityAttribute) | **Post** /identity-attributes | Create Identity Attribute
[**DeleteIdentityAttribute**](IdentityAttributesAPI.md#DeleteIdentityAttribute) | **Delete** /identity-attributes/{name} | Delete Identity Attribute
[**DeleteIdentityAttributesInBulk**](IdentityAttributesAPI.md#DeleteIdentityAttributesInBulk) | **Delete** /identity-attributes/bulk-delete | Bulk delete Identity Attributes
[**GetIdentityAttribute**](IdentityAttributesAPI.md#GetIdentityAttribute) | **Get** /identity-attributes/{name} | Get Identity Attribute
[**ListIdentityAttributes**](IdentityAttributesAPI.md#ListIdentityAttributes) | **Get** /identity-attributes | List Identity Attributes
[**PutIdentityAttribute**](IdentityAttributesAPI.md#PutIdentityAttribute) | **Put** /identity-attributes/{name} | Update Identity Attribute



## CreateIdentityAttribute

> IdentityAttribute CreateIdentityAttribute(ctx).IdentityAttribute(identityAttribute).Execute()

Create Identity Attribute



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
	identityAttribute := *openapiclient.NewIdentityAttribute("costCenter") // IdentityAttribute | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityAttributesAPI.CreateIdentityAttribute(context.Background()).IdentityAttribute(identityAttribute).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityAttributesAPI.CreateIdentityAttribute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateIdentityAttribute`: IdentityAttribute
	fmt.Fprintf(os.Stdout, "Response from `IdentityAttributesAPI.CreateIdentityAttribute`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIdentityAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identityAttribute** | [**IdentityAttribute**](IdentityAttribute.md) |  | 

### Return type

[**IdentityAttribute**](IdentityAttribute.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIdentityAttribute

> DeleteIdentityAttribute(ctx, name).Execute()

Delete Identity Attribute



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
	name := "displayName" // string | The attribute's technical name.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IdentityAttributesAPI.DeleteIdentityAttribute(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityAttributesAPI.DeleteIdentityAttribute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The attribute&#39;s technical name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdentityAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIdentityAttributesInBulk

> DeleteIdentityAttributesInBulk(ctx).IdentityAttributeNames(identityAttributeNames).Execute()

Bulk delete Identity Attributes



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
	identityAttributeNames := *openapiclient.NewIdentityAttributeNames() // IdentityAttributeNames | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IdentityAttributesAPI.DeleteIdentityAttributesInBulk(context.Background()).IdentityAttributeNames(identityAttributeNames).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityAttributesAPI.DeleteIdentityAttributesInBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdentityAttributesInBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identityAttributeNames** | [**IdentityAttributeNames**](IdentityAttributeNames.md) |  | 

### Return type

 (empty response body)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityAttribute

> IdentityAttribute GetIdentityAttribute(ctx, name).Execute()

Get Identity Attribute



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
	name := "displayName" // string | The attribute's technical name.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityAttributesAPI.GetIdentityAttribute(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityAttributesAPI.GetIdentityAttribute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIdentityAttribute`: IdentityAttribute
	fmt.Fprintf(os.Stdout, "Response from `IdentityAttributesAPI.GetIdentityAttribute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The attribute&#39;s technical name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IdentityAttribute**](IdentityAttribute.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIdentityAttributes

> []IdentityAttribute ListIdentityAttributes(ctx).IncludeSystem(includeSystem).IncludeSilent(includeSilent).SearchableOnly(searchableOnly).Count(count).Execute()

List Identity Attributes



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
	includeSystem := false // bool | Include 'system' attributes in the response. (optional) (default to false)
	includeSilent := false // bool | Include 'silent' attributes in the response. (optional) (default to false)
	searchableOnly := false // bool | Include only 'searchable' attributes in the response. (optional) (default to false)
	count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityAttributesAPI.ListIdentityAttributes(context.Background()).IncludeSystem(includeSystem).IncludeSilent(includeSilent).SearchableOnly(searchableOnly).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityAttributesAPI.ListIdentityAttributes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListIdentityAttributes`: []IdentityAttribute
	fmt.Fprintf(os.Stdout, "Response from `IdentityAttributesAPI.ListIdentityAttributes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListIdentityAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeSystem** | **bool** | Include &#39;system&#39; attributes in the response. | [default to false]
 **includeSilent** | **bool** | Include &#39;silent&#39; attributes in the response. | [default to false]
 **searchableOnly** | **bool** | Include only &#39;searchable&#39; attributes in the response. | [default to false]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]

### Return type

[**[]IdentityAttribute**](IdentityAttribute.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutIdentityAttribute

> IdentityAttribute PutIdentityAttribute(ctx, name).IdentityAttribute(identityAttribute).Execute()

Update Identity Attribute



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
	name := "displayName" // string | The attribute's technical name.
	identityAttribute := *openapiclient.NewIdentityAttribute("costCenter") // IdentityAttribute | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityAttributesAPI.PutIdentityAttribute(context.Background(), name).IdentityAttribute(identityAttribute).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityAttributesAPI.PutIdentityAttribute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutIdentityAttribute`: IdentityAttribute
	fmt.Fprintf(os.Stdout, "Response from `IdentityAttributesAPI.PutIdentityAttribute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The attribute&#39;s technical name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutIdentityAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **identityAttribute** | [**IdentityAttribute**](IdentityAttribute.md) |  | 

### Return type

[**IdentityAttribute**](IdentityAttribute.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

