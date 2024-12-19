# \GlobalTenantSecuritySettingsAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAuthOrgNetworkConfig**](GlobalTenantSecuritySettingsAPI.md#CreateAuthOrgNetworkConfig) | **Post** /auth-org/network-config | Create security network configuration.
[**GetAuthOrgLockoutConfig**](GlobalTenantSecuritySettingsAPI.md#GetAuthOrgLockoutConfig) | **Get** /auth-org/lockout-config | Get Auth Org Lockout Configuration.
[**GetAuthOrgNetworkConfig**](GlobalTenantSecuritySettingsAPI.md#GetAuthOrgNetworkConfig) | **Get** /auth-org/network-config | Get security network configuration.
[**GetAuthOrgServiceProviderConfig**](GlobalTenantSecuritySettingsAPI.md#GetAuthOrgServiceProviderConfig) | **Get** /auth-org/service-provider-config | Get Service Provider Configuration.
[**GetAuthOrgSessionConfig**](GlobalTenantSecuritySettingsAPI.md#GetAuthOrgSessionConfig) | **Get** /auth-org/session-config | Get Auth Org Session Configuration.
[**PatchAuthOrgLockoutConfig**](GlobalTenantSecuritySettingsAPI.md#PatchAuthOrgLockoutConfig) | **Patch** /auth-org/lockout-config | Update Auth Org Lockout Configuration
[**PatchAuthOrgNetworkConfig**](GlobalTenantSecuritySettingsAPI.md#PatchAuthOrgNetworkConfig) | **Patch** /auth-org/network-config | Update security network configuration.
[**PatchAuthOrgServiceProviderConfig**](GlobalTenantSecuritySettingsAPI.md#PatchAuthOrgServiceProviderConfig) | **Patch** /auth-org/service-provider-config | Update Service Provider Configuration
[**PatchAuthOrgSessionConfig**](GlobalTenantSecuritySettingsAPI.md#PatchAuthOrgSessionConfig) | **Patch** /auth-org/session-config | Update Auth Org Session Configuration



## CreateAuthOrgNetworkConfig

> NetworkConfiguration CreateAuthOrgNetworkConfig(ctx).NetworkConfiguration(networkConfiguration).Execute()

Create security network configuration.



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
	networkConfiguration := *openapiclient.NewNetworkConfiguration() // NetworkConfiguration | Network configuration creation request body.   The following constraints ensure the request body conforms to certain logical guidelines, which are:   1. Each string element in the range array must be a valid ip address or ip subnet mask.   2. Each string element in the geolocation array must be 2 characters, and they can only be uppercase letters.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlobalTenantSecuritySettingsAPI.CreateAuthOrgNetworkConfig(context.Background()).NetworkConfiguration(networkConfiguration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalTenantSecuritySettingsAPI.CreateAuthOrgNetworkConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAuthOrgNetworkConfig`: NetworkConfiguration
	fmt.Fprintf(os.Stdout, "Response from `GlobalTenantSecuritySettingsAPI.CreateAuthOrgNetworkConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAuthOrgNetworkConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **networkConfiguration** | [**NetworkConfiguration**](NetworkConfiguration.md) | Network configuration creation request body.   The following constraints ensure the request body conforms to certain logical guidelines, which are:   1. Each string element in the range array must be a valid ip address or ip subnet mask.   2. Each string element in the geolocation array must be 2 characters, and they can only be uppercase letters. | 

### Return type

[**NetworkConfiguration**](NetworkConfiguration.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthOrgLockoutConfig

> LockoutConfiguration GetAuthOrgLockoutConfig(ctx).Execute()

Get Auth Org Lockout Configuration.



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
	resp, r, err := apiClient.GlobalTenantSecuritySettingsAPI.GetAuthOrgLockoutConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalTenantSecuritySettingsAPI.GetAuthOrgLockoutConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuthOrgLockoutConfig`: LockoutConfiguration
	fmt.Fprintf(os.Stdout, "Response from `GlobalTenantSecuritySettingsAPI.GetAuthOrgLockoutConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthOrgLockoutConfigRequest struct via the builder pattern


### Return type

[**LockoutConfiguration**](LockoutConfiguration.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthOrgNetworkConfig

> NetworkConfiguration GetAuthOrgNetworkConfig(ctx).Execute()

Get security network configuration.



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
	resp, r, err := apiClient.GlobalTenantSecuritySettingsAPI.GetAuthOrgNetworkConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalTenantSecuritySettingsAPI.GetAuthOrgNetworkConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuthOrgNetworkConfig`: NetworkConfiguration
	fmt.Fprintf(os.Stdout, "Response from `GlobalTenantSecuritySettingsAPI.GetAuthOrgNetworkConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthOrgNetworkConfigRequest struct via the builder pattern


### Return type

[**NetworkConfiguration**](NetworkConfiguration.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthOrgServiceProviderConfig

> ServiceProviderConfiguration GetAuthOrgServiceProviderConfig(ctx).Execute()

Get Service Provider Configuration.



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
	resp, r, err := apiClient.GlobalTenantSecuritySettingsAPI.GetAuthOrgServiceProviderConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalTenantSecuritySettingsAPI.GetAuthOrgServiceProviderConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuthOrgServiceProviderConfig`: ServiceProviderConfiguration
	fmt.Fprintf(os.Stdout, "Response from `GlobalTenantSecuritySettingsAPI.GetAuthOrgServiceProviderConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthOrgServiceProviderConfigRequest struct via the builder pattern


### Return type

[**ServiceProviderConfiguration**](ServiceProviderConfiguration.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthOrgSessionConfig

> SessionConfiguration GetAuthOrgSessionConfig(ctx).Execute()

Get Auth Org Session Configuration.



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
	resp, r, err := apiClient.GlobalTenantSecuritySettingsAPI.GetAuthOrgSessionConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalTenantSecuritySettingsAPI.GetAuthOrgSessionConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuthOrgSessionConfig`: SessionConfiguration
	fmt.Fprintf(os.Stdout, "Response from `GlobalTenantSecuritySettingsAPI.GetAuthOrgSessionConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthOrgSessionConfigRequest struct via the builder pattern


### Return type

[**SessionConfiguration**](SessionConfiguration.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAuthOrgLockoutConfig

> LockoutConfiguration PatchAuthOrgLockoutConfig(ctx).JsonPatchOperation(jsonPatchOperation).Execute()

Update Auth Org Lockout Configuration



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
	jsonPatchOperation := []openapiclient.JsonPatchOperation{*openapiclient.NewJsonPatchOperation("replace", "/description")} // []JsonPatchOperation | A list of auth org lockout configuration update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard. Ensures that the patched Lockout Config conforms to certain logical guidelines, which are:   `1. maximumAttempts >= 1 && maximumAttempts <= 15   2. lockoutDuration >= 5 && lockoutDuration <= 60   3. lockoutWindow >= 5 && lockoutDuration <= 60`

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlobalTenantSecuritySettingsAPI.PatchAuthOrgLockoutConfig(context.Background()).JsonPatchOperation(jsonPatchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalTenantSecuritySettingsAPI.PatchAuthOrgLockoutConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAuthOrgLockoutConfig`: LockoutConfiguration
	fmt.Fprintf(os.Stdout, "Response from `GlobalTenantSecuritySettingsAPI.PatchAuthOrgLockoutConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPatchAuthOrgLockoutConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jsonPatchOperation** | [**[]JsonPatchOperation**](JsonPatchOperation.md) | A list of auth org lockout configuration update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard. Ensures that the patched Lockout Config conforms to certain logical guidelines, which are:   &#x60;1. maximumAttempts &gt;&#x3D; 1 &amp;&amp; maximumAttempts &lt;&#x3D; 15   2. lockoutDuration &gt;&#x3D; 5 &amp;&amp; lockoutDuration &lt;&#x3D; 60   3. lockoutWindow &gt;&#x3D; 5 &amp;&amp; lockoutDuration &lt;&#x3D; 60&#x60; | 

### Return type

[**LockoutConfiguration**](LockoutConfiguration.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAuthOrgNetworkConfig

> NetworkConfiguration PatchAuthOrgNetworkConfig(ctx).JsonPatchOperation(jsonPatchOperation).Execute()

Update security network configuration.



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
	jsonPatchOperation := []openapiclient.JsonPatchOperation{*openapiclient.NewJsonPatchOperation("replace", "/description")} // []JsonPatchOperation | A list of auth org network configuration update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard. Ensures that the patched Network Config conforms to certain logical guidelines, which are:   1. Each string element in the range array must be a valid ip address or ip subnet mask.   2. Each string element in the geolocation array must be 2 characters, and they can only be uppercase letters.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlobalTenantSecuritySettingsAPI.PatchAuthOrgNetworkConfig(context.Background()).JsonPatchOperation(jsonPatchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalTenantSecuritySettingsAPI.PatchAuthOrgNetworkConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAuthOrgNetworkConfig`: NetworkConfiguration
	fmt.Fprintf(os.Stdout, "Response from `GlobalTenantSecuritySettingsAPI.PatchAuthOrgNetworkConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPatchAuthOrgNetworkConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jsonPatchOperation** | [**[]JsonPatchOperation**](JsonPatchOperation.md) | A list of auth org network configuration update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard. Ensures that the patched Network Config conforms to certain logical guidelines, which are:   1. Each string element in the range array must be a valid ip address or ip subnet mask.   2. Each string element in the geolocation array must be 2 characters, and they can only be uppercase letters. | 

### Return type

[**NetworkConfiguration**](NetworkConfiguration.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAuthOrgServiceProviderConfig

> ServiceProviderConfiguration PatchAuthOrgServiceProviderConfig(ctx).JsonPatchOperation(jsonPatchOperation).Execute()

Update Service Provider Configuration



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
	jsonPatchOperation := []openapiclient.JsonPatchOperation{*openapiclient.NewJsonPatchOperation("replace", "/description")} // []JsonPatchOperation | A list of auth org service provider configuration update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard. Note: /federationProtocolDetails/0 is IdpDetails /federationProtocolDetails/1 is SpDetails Ensures that the patched ServiceProviderConfig conforms to certain logical guidelines, which are:   1. Do not add or remove any elements in the federation protocol details in the service provider configuration.   2. Do not modify, add, or delete the service provider details element in the federation protocol details.   3. If this is the first time the patched ServiceProviderConfig enables Remote IDP sign-in, it must also include IDPDetails.   4. If the patch enables Remote IDP sign in, the entityID in the IDPDetails cannot be null. IDPDetails must include an entityID.   5. Any JIT configuration update must be valid.  Just in time configuration update must be valid when enabled. This includes:   - A Source ID   - Source attribute mappings   - Source attribute maps have all the required key values (firstName, lastName, email)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlobalTenantSecuritySettingsAPI.PatchAuthOrgServiceProviderConfig(context.Background()).JsonPatchOperation(jsonPatchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalTenantSecuritySettingsAPI.PatchAuthOrgServiceProviderConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAuthOrgServiceProviderConfig`: ServiceProviderConfiguration
	fmt.Fprintf(os.Stdout, "Response from `GlobalTenantSecuritySettingsAPI.PatchAuthOrgServiceProviderConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPatchAuthOrgServiceProviderConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jsonPatchOperation** | [**[]JsonPatchOperation**](JsonPatchOperation.md) | A list of auth org service provider configuration update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard. Note: /federationProtocolDetails/0 is IdpDetails /federationProtocolDetails/1 is SpDetails Ensures that the patched ServiceProviderConfig conforms to certain logical guidelines, which are:   1. Do not add or remove any elements in the federation protocol details in the service provider configuration.   2. Do not modify, add, or delete the service provider details element in the federation protocol details.   3. If this is the first time the patched ServiceProviderConfig enables Remote IDP sign-in, it must also include IDPDetails.   4. If the patch enables Remote IDP sign in, the entityID in the IDPDetails cannot be null. IDPDetails must include an entityID.   5. Any JIT configuration update must be valid.  Just in time configuration update must be valid when enabled. This includes:   - A Source ID   - Source attribute mappings   - Source attribute maps have all the required key values (firstName, lastName, email) | 

### Return type

[**ServiceProviderConfiguration**](ServiceProviderConfiguration.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAuthOrgSessionConfig

> SessionConfiguration PatchAuthOrgSessionConfig(ctx).JsonPatchOperation(jsonPatchOperation).Execute()

Update Auth Org Session Configuration



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
	jsonPatchOperation := []openapiclient.JsonPatchOperation{*openapiclient.NewJsonPatchOperation("replace", "/description")} // []JsonPatchOperation | A list of auth org session configuration update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard.  Ensures that the patched Session Config conforms to certain logical guidelines, which are:   `1. maxSessionTime >= 1 && maxSessionTime <= 10080 (1 week)   2. maxIdleTime >= 1 && maxIdleTime <= 1440 (1 day)   3. maxSessionTime must have a greater duration than maxIdleTime.` 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlobalTenantSecuritySettingsAPI.PatchAuthOrgSessionConfig(context.Background()).JsonPatchOperation(jsonPatchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalTenantSecuritySettingsAPI.PatchAuthOrgSessionConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAuthOrgSessionConfig`: SessionConfiguration
	fmt.Fprintf(os.Stdout, "Response from `GlobalTenantSecuritySettingsAPI.PatchAuthOrgSessionConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPatchAuthOrgSessionConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jsonPatchOperation** | [**[]JsonPatchOperation**](JsonPatchOperation.md) | A list of auth org session configuration update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard.  Ensures that the patched Session Config conforms to certain logical guidelines, which are:   &#x60;1. maxSessionTime &gt;&#x3D; 1 &amp;&amp; maxSessionTime &lt;&#x3D; 10080 (1 week)   2. maxIdleTime &gt;&#x3D; 1 &amp;&amp; maxIdleTime &lt;&#x3D; 1440 (1 day)   3. maxSessionTime must have a greater duration than maxIdleTime.&#x60;  | 

### Return type

[**SessionConfiguration**](SessionConfiguration.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

