# \ManagedClustersAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v2024*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateManagedCluster**](ManagedClustersAPI.md#CreateManagedCluster) | **Post** /managed-clusters | Create Create Managed Cluster
[**DeleteManagedCluster**](ManagedClustersAPI.md#DeleteManagedCluster) | **Delete** /managed-clusters/{id} | Delete Managed Cluster
[**GetClientLogConfiguration**](ManagedClustersAPI.md#GetClientLogConfiguration) | **Get** /managed-clusters/{id}/log-config | Get Managed Cluster Log Configuration
[**GetManagedCluster**](ManagedClustersAPI.md#GetManagedCluster) | **Get** /managed-clusters/{id} | Get Managed Cluster
[**GetManagedClusters**](ManagedClustersAPI.md#GetManagedClusters) | **Get** /managed-clusters | Get Managed Clusters
[**PutClientLogConfiguration**](ManagedClustersAPI.md#PutClientLogConfiguration) | **Put** /managed-clusters/{id}/log-config | Update Managed Cluster Log Configuration
[**Update**](ManagedClustersAPI.md#Update) | **Post** /managed-clusters/{id}/manualUpgrade | Trigger Manual Upgrade for Managed Cluster
[**UpdateManagedCluster**](ManagedClustersAPI.md#UpdateManagedCluster) | **Patch** /managed-clusters/{id} | Update Managed Cluster



## CreateManagedCluster

> ManagedCluster CreateManagedCluster(ctx).ManagedClusterRequest(managedClusterRequest).Execute()

Create Create Managed Cluster



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
	managedClusterRequest := *openapiclient.NewManagedClusterRequest("Managed Cluster Name") // ManagedClusterRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedClustersAPI.CreateManagedCluster(context.Background()).ManagedClusterRequest(managedClusterRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedClustersAPI.CreateManagedCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateManagedCluster`: ManagedCluster
	fmt.Fprintf(os.Stdout, "Response from `ManagedClustersAPI.CreateManagedCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateManagedClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **managedClusterRequest** | [**ManagedClusterRequest**](ManagedClusterRequest.md) |  | 

### Return type

[**ManagedCluster**](ManagedCluster.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteManagedCluster

> DeleteManagedCluster(ctx, id).RemoveClients(removeClients).Execute()

Delete Managed Cluster



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
	id := "2c9180897de347a2017de8859e8c5039" // string | Managed cluster ID.
	removeClients := false // bool | Flag to determine the need to delete a cluster with clients. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ManagedClustersAPI.DeleteManagedCluster(context.Background(), id).RemoveClients(removeClients).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedClustersAPI.DeleteManagedCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Managed cluster ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteManagedClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **removeClients** | **bool** | Flag to determine the need to delete a cluster with clients. | [default to false]

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


## GetClientLogConfiguration

> ClientLogConfiguration GetClientLogConfiguration(ctx, id).Execute()

Get Managed Cluster Log Configuration



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
	id := "2b838de9-db9b-abcf-e646-d4f274ad4238" // string | ID of managed cluster to get log configuration for.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedClustersAPI.GetClientLogConfiguration(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedClustersAPI.GetClientLogConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientLogConfiguration`: ClientLogConfiguration
	fmt.Fprintf(os.Stdout, "Response from `ManagedClustersAPI.GetClientLogConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of managed cluster to get log configuration for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientLogConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClientLogConfiguration**](ClientLogConfiguration.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManagedCluster

> ManagedCluster GetManagedCluster(ctx, id).Execute()

Get Managed Cluster



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
	id := "2c9180897de347a2017de8859e8c5039" // string | Managed cluster ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedClustersAPI.GetManagedCluster(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedClustersAPI.GetManagedCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetManagedCluster`: ManagedCluster
	fmt.Fprintf(os.Stdout, "Response from `ManagedClustersAPI.GetManagedCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Managed cluster ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetManagedClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ManagedCluster**](ManagedCluster.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManagedClusters

> []ManagedCluster GetManagedClusters(ctx).Offset(offset).Limit(limit).Count(count).Filters(filters).Execute()

Get Managed Clusters



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
	offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
	limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
	count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
	filters := "operational eq "operation"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **operational**: *eq* (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedClustersAPI.GetManagedClusters(context.Background()).Offset(offset).Limit(limit).Count(count).Filters(filters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedClustersAPI.GetManagedClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetManagedClusters`: []ManagedCluster
	fmt.Fprintf(os.Stdout, "Response from `ManagedClustersAPI.GetManagedClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetManagedClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **operational**: *eq* | 

### Return type

[**[]ManagedCluster**](ManagedCluster.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutClientLogConfiguration

> ClientLogConfiguration PutClientLogConfiguration(ctx, id).PutClientLogConfigurationRequest(putClientLogConfigurationRequest).Execute()

Update Managed Cluster Log Configuration



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
	id := "2b838de9-db9b-abcf-e646-d4f274ad4238" // string | ID of the managed cluster to update the log configuration for.
	putClientLogConfigurationRequest := openapiclient.putClientLogConfiguration_request{ClientLogConfigurationDurationMinutes: openapiclient.NewClientLogConfigurationDurationMinutes(openapiclient.StandardLevel("false"))} // PutClientLogConfigurationRequest | Client log configuration for the given managed cluster.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedClustersAPI.PutClientLogConfiguration(context.Background(), id).PutClientLogConfigurationRequest(putClientLogConfigurationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedClustersAPI.PutClientLogConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutClientLogConfiguration`: ClientLogConfiguration
	fmt.Fprintf(os.Stdout, "Response from `ManagedClustersAPI.PutClientLogConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the managed cluster to update the log configuration for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutClientLogConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putClientLogConfigurationRequest** | [**PutClientLogConfigurationRequest**](PutClientLogConfigurationRequest.md) | Client log configuration for the given managed cluster. | 

### Return type

[**ClientLogConfiguration**](ClientLogConfiguration.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> ClusterManualUpgrade Update(ctx, id).Execute()

Trigger Manual Upgrade for Managed Cluster



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
	id := "2b838de9-db9b-abcf-e646-d4f274ad4238" // string | ID of managed cluster to trigger manual upgrade.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedClustersAPI.Update(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedClustersAPI.Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update`: ClusterManualUpgrade
	fmt.Fprintf(os.Stdout, "Response from `ManagedClustersAPI.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of managed cluster to trigger manual upgrade. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClusterManualUpgrade**](ClusterManualUpgrade.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateManagedCluster

> ManagedCluster UpdateManagedCluster(ctx, id).JsonPatchOperation(jsonPatchOperation).Execute()

Update Managed Cluster



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
	id := "2c9180897de347a2017de8859e8c5039" // string | Managed cluster ID.
	jsonPatchOperation := []openapiclient.JsonPatchOperation{*openapiclient.NewJsonPatchOperation("replace", "/description")} // []JsonPatchOperation | JSONPatch payload used to update the object.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedClustersAPI.UpdateManagedCluster(context.Background(), id).JsonPatchOperation(jsonPatchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedClustersAPI.UpdateManagedCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateManagedCluster`: ManagedCluster
	fmt.Fprintf(os.Stdout, "Response from `ManagedClustersAPI.UpdateManagedCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Managed cluster ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateManagedClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jsonPatchOperation** | [**[]JsonPatchOperation**](JsonPatchOperation.md) | JSONPatch payload used to update the object. | 

### Return type

[**ManagedCluster**](ManagedCluster.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

