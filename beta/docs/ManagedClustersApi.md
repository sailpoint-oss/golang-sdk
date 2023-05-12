# \ManagedClustersApi

All URIs are relative to *https://sailpoint.api.identitynow.com/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetClientLogConfiguration**](ManagedClustersApi.md#GetClientLogConfiguration) | **Get** /managed-clusters/{id}/log-config | Get managed cluster&#39;s log configuration
[**GetManagedCluster**](ManagedClustersApi.md#GetManagedCluster) | **Get** /managed-clusters/{id} | Get a specified ManagedCluster.
[**GetManagedClusters**](ManagedClustersApi.md#GetManagedClusters) | **Get** /managed-clusters | Retrieve all Managed Clusters.
[**PutClientLogConfiguration**](ManagedClustersApi.md#PutClientLogConfiguration) | **Put** /managed-clusters/{id}/log-config | Update managed cluster&#39;s log configuration



## GetClientLogConfiguration

> ClientLogConfiguration GetClientLogConfiguration(ctx, id).Execute()

Get managed cluster's log configuration



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "aClusterId" // string | ID of ManagedCluster to get log configuration for

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagedClustersApi.GetClientLogConfiguration(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedClustersApi.GetClientLogConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClientLogConfiguration`: ClientLogConfiguration
    fmt.Fprintf(os.Stdout, "Response from `ManagedClustersApi.GetClientLogConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of ManagedCluster to get log configuration for | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientLogConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClientLogConfiguration**](ClientLogConfiguration.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManagedCluster

> ManagedCluster GetManagedCluster(ctx, id).Execute()

Get a specified ManagedCluster.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "aClusterId" // string | ID of the ManagedCluster to get

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagedClustersApi.GetManagedCluster(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedClustersApi.GetManagedCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetManagedCluster`: ManagedCluster
    fmt.Fprintf(os.Stdout, "Response from `ManagedClustersApi.GetManagedCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the ManagedCluster to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetManagedClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ManagedCluster**](ManagedCluster.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManagedClusters

> []ManagedCluster GetManagedClusters(ctx).Offset(offset).Limit(limit).Count(count).Filters(filters).Execute()

Retrieve all Managed Clusters.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
    filters := "operational eq operation" // string | Filtering is supported for the following fields and operators:  **operational**: *eq* (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagedClustersApi.GetManagedClusters(context.Background()).Offset(offset).Limit(limit).Count(count).Filters(filters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedClustersApi.GetManagedClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetManagedClusters`: []ManagedCluster
    fmt.Fprintf(os.Stdout, "Response from `ManagedClustersApi.GetManagedClusters`: %v\n", resp)
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
 **filters** | **string** | Filtering is supported for the following fields and operators:  **operational**: *eq* | 

### Return type

[**[]ManagedCluster**](ManagedCluster.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutClientLogConfiguration

> ClientLogConfiguration PutClientLogConfiguration(ctx, id).ClientLogConfiguration(clientLogConfiguration).Execute()

Update managed cluster's log configuration



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "aClusterId" // string | ID of ManagedCluster to update log configuration for
    clientLogConfiguration := *openapiclient.NewClientLogConfiguration(int32(120), openapiclient.StandardLevel("false")) // ClientLogConfiguration | ClientLogConfiguration for given ManagedCluster

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagedClustersApi.PutClientLogConfiguration(context.Background(), id).ClientLogConfiguration(clientLogConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedClustersApi.PutClientLogConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutClientLogConfiguration`: ClientLogConfiguration
    fmt.Fprintf(os.Stdout, "Response from `ManagedClustersApi.PutClientLogConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of ManagedCluster to update log configuration for | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutClientLogConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientLogConfiguration** | [**ClientLogConfiguration**](ClientLogConfiguration.md) | ClientLogConfiguration for given ManagedCluster | 

### Return type

[**ClientLogConfiguration**](ClientLogConfiguration.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

