# \PasswordPoliciesAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v2024*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePasswordPolicy**](PasswordPoliciesAPI.md#CreatePasswordPolicy) | **Post** /password-policies | Create Password Policy
[**DeletePasswordPolicy**](PasswordPoliciesAPI.md#DeletePasswordPolicy) | **Delete** /password-policies/{id} | Delete Password Policy by ID
[**GetPasswordPolicyById**](PasswordPoliciesAPI.md#GetPasswordPolicyById) | **Get** /password-policies/{id} | Get Password Policy by ID
[**ListPasswordPolicies**](PasswordPoliciesAPI.md#ListPasswordPolicies) | **Get** /password-policies | List Password Policies
[**SetPasswordPolicy**](PasswordPoliciesAPI.md#SetPasswordPolicy) | **Put** /password-policies/{id} | Update Password Policy by ID



## CreatePasswordPolicy

> PasswordPolicyV3Dto CreatePasswordPolicy(ctx).PasswordPolicyV3Dto(passwordPolicyV3Dto).Execute()

Create Password Policy



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
	passwordPolicyV3Dto := *openapiclient.NewPasswordPolicyV3Dto() // PasswordPolicyV3Dto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PasswordPoliciesAPI.CreatePasswordPolicy(context.Background()).PasswordPolicyV3Dto(passwordPolicyV3Dto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PasswordPoliciesAPI.CreatePasswordPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePasswordPolicy`: PasswordPolicyV3Dto
	fmt.Fprintf(os.Stdout, "Response from `PasswordPoliciesAPI.CreatePasswordPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePasswordPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **passwordPolicyV3Dto** | [**PasswordPolicyV3Dto**](PasswordPolicyV3Dto.md) |  | 

### Return type

[**PasswordPolicyV3Dto**](PasswordPolicyV3Dto.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePasswordPolicy

> DeletePasswordPolicy(ctx, id).Execute()

Delete Password Policy by ID



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
	id := "ff808081838d9e9d01838da6a03e0002" // string | The ID of password policy to delete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PasswordPoliciesAPI.DeletePasswordPolicy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PasswordPoliciesAPI.DeletePasswordPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of password policy to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePasswordPolicyRequest struct via the builder pattern


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


## GetPasswordPolicyById

> PasswordPolicyV3Dto GetPasswordPolicyById(ctx, id).Execute()

Get Password Policy by ID



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
	id := "ff808081838d9e9d01838da6a03e0005" // string | The ID of password policy to retrieve.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PasswordPoliciesAPI.GetPasswordPolicyById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PasswordPoliciesAPI.GetPasswordPolicyById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPasswordPolicyById`: PasswordPolicyV3Dto
	fmt.Fprintf(os.Stdout, "Response from `PasswordPoliciesAPI.GetPasswordPolicyById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of password policy to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPasswordPolicyByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PasswordPolicyV3Dto**](PasswordPolicyV3Dto.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPasswordPolicies

> []PasswordPolicyV3Dto ListPasswordPolicies(ctx).Limit(limit).Offset(offset).Count(count).Execute()

List Password Policies



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
	count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PasswordPoliciesAPI.ListPasswordPolicies(context.Background()).Limit(limit).Offset(offset).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PasswordPoliciesAPI.ListPasswordPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPasswordPolicies`: []PasswordPolicyV3Dto
	fmt.Fprintf(os.Stdout, "Response from `PasswordPoliciesAPI.ListPasswordPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPasswordPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]

### Return type

[**[]PasswordPolicyV3Dto**](PasswordPolicyV3Dto.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetPasswordPolicy

> PasswordPolicyV3Dto SetPasswordPolicy(ctx, id).PasswordPolicyV3Dto(passwordPolicyV3Dto).Execute()

Update Password Policy by ID



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
	id := "ff808081838d9e9d01838da6a03e0007" // string | The ID of password policy to update.
	passwordPolicyV3Dto := *openapiclient.NewPasswordPolicyV3Dto() // PasswordPolicyV3Dto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PasswordPoliciesAPI.SetPasswordPolicy(context.Background(), id).PasswordPolicyV3Dto(passwordPolicyV3Dto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PasswordPoliciesAPI.SetPasswordPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetPasswordPolicy`: PasswordPolicyV3Dto
	fmt.Fprintf(os.Stdout, "Response from `PasswordPoliciesAPI.SetPasswordPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of password policy to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetPasswordPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **passwordPolicyV3Dto** | [**PasswordPolicyV3Dto**](PasswordPolicyV3Dto.md) |  | 

### Return type

[**PasswordPolicyV3Dto**](PasswordPolicyV3Dto.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

