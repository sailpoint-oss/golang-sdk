# \ManagedClusterTypesAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v2024*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateManagedClusterType**](ManagedClusterTypesAPI.md#CreateManagedClusterType) | **Post** /managed-cluster-types | Create new Managed Cluster Type
[**DeleteManagedClusterType**](ManagedClusterTypesAPI.md#DeleteManagedClusterType) | **Delete** /managed-cluster-types/{id} | Delete a Managed Cluster Type
[**GetManagedClusterType**](ManagedClusterTypesAPI.md#GetManagedClusterType) | **Get** /managed-cluster-types/{id} | Get a Managed Cluster Type
[**GetManagedClusterTypes**](ManagedClusterTypesAPI.md#GetManagedClusterTypes) | **Get** /managed-cluster-types | Get Managed Cluster Types
[**UpdateManagedClusterType**](ManagedClusterTypesAPI.md#UpdateManagedClusterType) | **Patch** /managed-cluster-types/{id} | Update a Managed Cluster Type



## CreateManagedClusterType

> ManagedClusterType CreateManagedClusterType(ctx).ManagedClusterType(managedClusterType).Execute()

Create new Managed Cluster Type



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
	managedClusterType := *openapiclient.NewManagedClusterType("idn", "megapod-useast1", "denali-cjh") // ManagedClusterType | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedClusterTypesAPI.CreateManagedClusterType(context.Background()).ManagedClusterType(managedClusterType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedClusterTypesAPI.CreateManagedClusterType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateManagedClusterType`: ManagedClusterType
	fmt.Fprintf(os.Stdout, "Response from `ManagedClusterTypesAPI.CreateManagedClusterType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateManagedClusterTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **managedClusterType** | [**ManagedClusterType**](ManagedClusterType.md) |  | 

### Return type

[**ManagedClusterType**](ManagedClusterType.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteManagedClusterType

> DeleteManagedClusterType(ctx, id).Execute()

Delete a Managed Cluster Type



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
	id := "aClusterTypeId" // string | The Managed Cluster Type ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ManagedClusterTypesAPI.DeleteManagedClusterType(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedClusterTypesAPI.DeleteManagedClusterType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Managed Cluster Type ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteManagedClusterTypeRequest struct via the builder pattern


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


## GetManagedClusterType

> ManagedClusterType GetManagedClusterType(ctx, id).Execute()

Get a Managed Cluster Type



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
	id := "aClusterTypeId" // string | The Managed Cluster Type ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedClusterTypesAPI.GetManagedClusterType(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedClusterTypesAPI.GetManagedClusterType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetManagedClusterType`: ManagedClusterType
	fmt.Fprintf(os.Stdout, "Response from `ManagedClusterTypesAPI.GetManagedClusterType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Managed Cluster Type ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetManagedClusterTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ManagedClusterType**](ManagedClusterType.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManagedClusterTypes

> []ManagedClusterType GetManagedClusterTypes(ctx).Type_(type_).Pod(pod).Org(org).Offset(offset).Limit(limit).Execute()

Get Managed Cluster Types



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
	type_ := "IDN" // string | Type descriptor (optional)
	pod := "megapod-useast1" // string | Pinned pod (or default) (optional)
	org := "denali-xyz" // string | Pinned org (or default) (optional)
	offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
	limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedClusterTypesAPI.GetManagedClusterTypes(context.Background()).Type_(type_).Pod(pod).Org(org).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedClusterTypesAPI.GetManagedClusterTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetManagedClusterTypes`: []ManagedClusterType
	fmt.Fprintf(os.Stdout, "Response from `ManagedClusterTypesAPI.GetManagedClusterTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetManagedClusterTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | Type descriptor | 
 **pod** | **string** | Pinned pod (or default) | 
 **org** | **string** | Pinned org (or default) | 
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]

### Return type

[**[]ManagedClusterType**](ManagedClusterType.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateManagedClusterType

> ManagedClusterType UpdateManagedClusterType(ctx, id).JsonPatch(jsonPatch).Execute()

Update a Managed Cluster Type



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
	id := "aClusterTypeId" // string | The Managed Cluster Type ID
	jsonPatch := *openapiclient.NewJsonPatch() // JsonPatch | The JSONPatch payload used to update the schema.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedClusterTypesAPI.UpdateManagedClusterType(context.Background(), id).JsonPatch(jsonPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedClusterTypesAPI.UpdateManagedClusterType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateManagedClusterType`: ManagedClusterType
	fmt.Fprintf(os.Stdout, "Response from `ManagedClusterTypesAPI.UpdateManagedClusterType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Managed Cluster Type ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateManagedClusterTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jsonPatch** | [**JsonPatch**](JsonPatch.md) | The JSONPatch payload used to update the schema. | 

### Return type

[**ManagedClusterType**](ManagedClusterType.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

