# \ManagedClientsApi

All URIs are relative to *https://sailpoint.api.identitynow.com/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetManagedClientStatus**](ManagedClientsApi.md#GetManagedClientStatus) | **Get** /managed-clients/{id}/status | Specified Managed Client Status.
[**UpdateManagedClientStatus**](ManagedClientsApi.md#UpdateManagedClientStatus) | **Post** /managed-clients/{id}/status | Handle status request from client



## GetManagedClientStatus

> ManagedClientStatus GetManagedClientStatus(ctx, id).Type_(type_).Execute()

Specified Managed Client Status.



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
    id := "aClientId" // string | ID of the Managed Client Status to get
    type_ := openapiclient.ManagedClientType("CCG") // ManagedClientType | Type of the Managed Client Status to get

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagedClientsApi.GetManagedClientStatus(context.Background(), id).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedClientsApi.GetManagedClientStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetManagedClientStatus`: ManagedClientStatus
    fmt.Fprintf(os.Stdout, "Response from `ManagedClientsApi.GetManagedClientStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Managed Client Status to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetManagedClientStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**ManagedClientType**](ManagedClientType.md) | Type of the Managed Client Status to get | 

### Return type

[**ManagedClientStatus**](ManagedClientStatus.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateManagedClientStatus

> ManagedClientStatusAggResponse UpdateManagedClientStatus(ctx, id).ManagedClientStatus(managedClientStatus).Execute()

Handle status request from client



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    id := "aClientId" // string | ID of the Managed Client Status to update
    managedClientStatus := *openapiclient.NewManagedClientStatus(map[string]interface{}({alertKey=, id=5678, clusterId=1234, ccg_etag=ccg_etag123xyz456, ccg_pin=NONE, cookbook_etag=20210420125956-20210511144538, hostname=megapod-useast1-secret-hostname.sailpoint.com, internal_ip=127.0.0.1, lastSeen=1620843964604, sinceSeen=14708, sinceSeenMillis=14708, localDev=false, stacktrace=, state=null, status=NORMAL, uuid=null, product=idn, va_version=null, platform_version=2, os_version=2345.3.1, os_type=flatcar, hypervisor=unknown}), openapiclient.ManagedClientStatusEnum("NORMAL"), "TODO", time.Now()) // ManagedClientStatus | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagedClientsApi.UpdateManagedClientStatus(context.Background(), id).ManagedClientStatus(managedClientStatus).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedClientsApi.UpdateManagedClientStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateManagedClientStatus`: ManagedClientStatusAggResponse
    fmt.Fprintf(os.Stdout, "Response from `ManagedClientsApi.UpdateManagedClientStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Managed Client Status to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateManagedClientStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **managedClientStatus** | [**ManagedClientStatus**](ManagedClientStatus.md) |  | 

### Return type

[**ManagedClientStatusAggResponse**](ManagedClientStatusAggResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

