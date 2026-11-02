---
id: default
title: Default
pagination_label: Default
sidebar_label: Default
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Default', 'Default'] 
slug: /tools/sdk/go//methods/default
tags: ['SDK', 'Software Development Kit', 'Default', 'Default']
---

# DefaultAPI
   
All URIs are relative to *https://sailpoint.api.identitynow.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**generic-delete**](#generic-delete) | **Delete** `/{path}` | Generic DELETE request
[**generic-get**](#generic-get) | **Get** `/{path}` | Generic GET request
[**generic-patch**](#generic-patch) | **Patch** `/{path}` | Generic PATCH request
[**generic-post**](#generic-post) | **Post** `/{path}` | Generic POST request
[**generic-put**](#generic-put) | **Put** `/{path}` | Generic PUT request


## generic-delete
:::warning experimental 
This API is currently in an experimental state. The API is subject to change based on feedback and further testing. You must include the X-SailPoint-Experimental header and set it to `true` to use this endpoint.
:::
:::tip setting x-sailpoint-experimental header
 on the configuration object you can set the `x-sailpoint-experimental` header to `true' to enable all experimantl endpoints within the SDK.
 Example:
 ```go
   configuration = Configuration()
   configuration.Experimental = true
 ```
:::
Generic DELETE request
Sends a DELETE request to any versioned API path.

[API Spec](https://developer.sailpoint.com/docs/api/generic-delete)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | API path relative to the tenant base URL, including the API version, for example &#x60;accounts/v1&#x60;, &#x60;sources/v1/2c9180835d2e5168015d32f890ca1581&#x60; or &#x60;sources/v2/2c9180835d2e5168015d32f890ca1581/provisioning-policies&#x60;. Use whichever version the endpoint has. A leading slash is optional. Slashes separate path segments, and the SDK URL-encodes each segment, so pass the segments unencoded. Anything after a &#x60;?&#x60; is sent as query parameters, for example &#x60;accounts/v1?detailLevel&#x3D;SLIM&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenericDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSailPointExperimental** | **string** | Set this header to &#x60;true&#x60; to call an experimental API. The SDK configuration must also enable experimental APIs. | 

### Return type

**interface{}**

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
  
    
	sailpoint "github.com/sailpoint-oss/golang-sdk/v3"
)

func main() {
    path := `accounts/v1` // string | API path relative to the tenant base URL, including the API version, for example `accounts/v1`, `sources/v1/2c9180835d2e5168015d32f890ca1581` or `sources/v2/2c9180835d2e5168015d32f890ca1581/provisioning-policies`. Use whichever version the endpoint has. A leading slash is optional. Slashes separate path segments, and the SDK URL-encodes each segment, so pass the segments unencoded. Anything after a `?` is sent as query parameters, for example `accounts/v1?detailLevel=SLIM`. # string | API path relative to the tenant base URL, including the API version, for example `accounts/v1`, `sources/v1/2c9180835d2e5168015d32f890ca1581` or `sources/v2/2c9180835d2e5168015d32f890ca1581/provisioning-policies`. Use whichever version the endpoint has. A leading slash is optional. Slashes separate path segments, and the SDK URL-encodes each segment, so pass the segments unencoded. Anything after a `?` is sent as query parameters, for example `accounts/v1?detailLevel=SLIM`.
    xSailPointExperimental := `true` // string | Set this header to `true` to call an experimental API. The SDK configuration must also enable experimental APIs. (optional) # string | Set this header to `true` to call an experimental API. The SDK configuration must also enable experimental APIs. (optional)

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GenericDelete(context.Background(), path).Execute()
	  //resp, r, err := apiClient.DefaultAPI.GenericDelete(context.Background(), path).XSailPointExperimental(xSailPointExperimental).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GenericDelete``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenericDelete`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GenericDelete`: %v\n", resp)
}
```

[[Back to top]](#)

## generic-get
:::warning experimental 
This API is currently in an experimental state. The API is subject to change based on feedback and further testing. You must include the X-SailPoint-Experimental header and set it to `true` to use this endpoint.
:::
:::tip setting x-sailpoint-experimental header
 on the configuration object you can set the `x-sailpoint-experimental` header to `true' to enable all experimantl endpoints within the SDK.
 Example:
 ```go
   configuration = Configuration()
   configuration.Experimental = true
 ```
:::
Generic GET request
Sends a GET request to any versioned API path, for example `sources/v1/{id}/schemas` or `sources/v2/{sourceId}/provisioning-policies`.

[API Spec](https://developer.sailpoint.com/docs/api/generic-get)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | API path relative to the tenant base URL, including the API version, for example &#x60;accounts/v1&#x60;, &#x60;sources/v1/2c9180835d2e5168015d32f890ca1581&#x60; or &#x60;sources/v2/2c9180835d2e5168015d32f890ca1581/provisioning-policies&#x60;. Use whichever version the endpoint has. A leading slash is optional. Slashes separate path segments, and the SDK URL-encodes each segment, so pass the segments unencoded. Anything after a &#x60;?&#x60; is sent as query parameters, for example &#x60;accounts/v1?detailLevel&#x3D;SLIM&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenericGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Max number of results to return. See [Standard Collection Parameters](https://developer.sailpoint.com/docs/standard-collection-parameters) for more information. | 
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. | 
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored. | 
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/docs/standard-collection-parameters#filtering-results). | 
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/docs/standard-collection-parameters#sorting-results). | 
 **xSailPointExperimental** | **string** | Set this header to &#x60;true&#x60; to call an experimental API. The SDK configuration must also enable experimental APIs. | 

### Return type

**interface{}**

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
  
    
	sailpoint "github.com/sailpoint-oss/golang-sdk/v3"
)

func main() {
    path := `accounts/v1` // string | API path relative to the tenant base URL, including the API version, for example `accounts/v1`, `sources/v1/2c9180835d2e5168015d32f890ca1581` or `sources/v2/2c9180835d2e5168015d32f890ca1581/provisioning-policies`. Use whichever version the endpoint has. A leading slash is optional. Slashes separate path segments, and the SDK URL-encodes each segment, so pass the segments unencoded. Anything after a `?` is sent as query parameters, for example `accounts/v1?detailLevel=SLIM`. # string | API path relative to the tenant base URL, including the API version, for example `accounts/v1`, `sources/v1/2c9180835d2e5168015d32f890ca1581` or `sources/v2/2c9180835d2e5168015d32f890ca1581/provisioning-policies`. Use whichever version the endpoint has. A leading slash is optional. Slashes separate path segments, and the SDK URL-encodes each segment, so pass the segments unencoded. Anything after a `?` is sent as query parameters, for example `accounts/v1?detailLevel=SLIM`.
    limit := 250 // int32 | Max number of results to return. See [Standard Collection Parameters](https://developer.sailpoint.com/docs/standard-collection-parameters) for more information. (optional) # int32 | Max number of results to return. See [Standard Collection Parameters](https://developer.sailpoint.com/docs/standard-collection-parameters) for more information. (optional)
    offset := 0 // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. (optional) # int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. (optional)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored. (optional) # bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored. (optional)
    filters := `name eq "Example"` // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/docs/standard-collection-parameters#filtering-results). (optional) # string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/docs/standard-collection-parameters#filtering-results). (optional)
    sorters := `name,-created` // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/docs/standard-collection-parameters#sorting-results). (optional) # string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/docs/standard-collection-parameters#sorting-results). (optional)
    xSailPointExperimental := `true` // string | Set this header to `true` to call an experimental API. The SDK configuration must also enable experimental APIs. (optional) # string | Set this header to `true` to call an experimental API. The SDK configuration must also enable experimental APIs. (optional)

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GenericGet(context.Background(), path).Execute()
	  //resp, r, err := apiClient.DefaultAPI.GenericGet(context.Background(), path).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).XSailPointExperimental(xSailPointExperimental).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GenericGet``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenericGet`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GenericGet`: %v\n", resp)
}
```

[[Back to top]](#)

## generic-patch
:::warning experimental 
This API is currently in an experimental state. The API is subject to change based on feedback and further testing. You must include the X-SailPoint-Experimental header and set it to `true` to use this endpoint.
:::
:::tip setting x-sailpoint-experimental header
 on the configuration object you can set the `x-sailpoint-experimental` header to `true' to enable all experimantl endpoints within the SDK.
 Example:
 ```go
   configuration = Configuration()
   configuration.Experimental = true
 ```
:::
Generic PATCH request
Sends a PATCH request to any versioned API path. The body is a JSON Patch document (RFC 6902): a list of operations such as `[{"op": "replace", "path": "/description", "value": "New description"}]`.

[API Spec](https://developer.sailpoint.com/docs/api/generic-patch)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | API path relative to the tenant base URL, including the API version, for example &#x60;accounts/v1&#x60;, &#x60;sources/v1/2c9180835d2e5168015d32f890ca1581&#x60; or &#x60;sources/v2/2c9180835d2e5168015d32f890ca1581/provisioning-policies&#x60;. Use whichever version the endpoint has. A leading slash is optional. Slashes separate path segments, and the SDK URL-encodes each segment, so pass the segments unencoded. Anything after a &#x60;?&#x60; is sent as query parameters, for example &#x60;accounts/v1?detailLevel&#x3D;SLIM&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenericPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSailPointExperimental** | **string** | Set this header to &#x60;true&#x60; to call an experimental API. The SDK configuration must also enable experimental APIs. | 
 **requestBody** | **interface{}** |  | 

### Return type

**interface{}**

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
  
    
	sailpoint "github.com/sailpoint-oss/golang-sdk/v3"
)

func main() {
    path := `accounts/v1` // string | API path relative to the tenant base URL, including the API version, for example `accounts/v1`, `sources/v1/2c9180835d2e5168015d32f890ca1581` or `sources/v2/2c9180835d2e5168015d32f890ca1581/provisioning-policies`. Use whichever version the endpoint has. A leading slash is optional. Slashes separate path segments, and the SDK URL-encodes each segment, so pass the segments unencoded. Anything after a `?` is sent as query parameters, for example `accounts/v1?detailLevel=SLIM`. # string | API path relative to the tenant base URL, including the API version, for example `accounts/v1`, `sources/v1/2c9180835d2e5168015d32f890ca1581` or `sources/v2/2c9180835d2e5168015d32f890ca1581/provisioning-policies`. Use whichever version the endpoint has. A leading slash is optional. Slashes separate path segments, and the SDK URL-encodes each segment, so pass the segments unencoded. Anything after a `?` is sent as query parameters, for example `accounts/v1?detailLevel=SLIM`.
    xSailPointExperimental := `true` // string | Set this header to `true` to call an experimental API. The SDK configuration must also enable experimental APIs. (optional) # string | Set this header to `true` to call an experimental API. The SDK configuration must also enable experimental APIs. (optional)
    requestBody :=  // interface{} |  (optional) # interface{} |  (optional)

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GenericPatch(context.Background(), path).Execute()
	  //resp, r, err := apiClient.DefaultAPI.GenericPatch(context.Background(), path).XSailPointExperimental(xSailPointExperimental).RequestBody(requestBody).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GenericPatch``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenericPatch`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GenericPatch`: %v\n", resp)
}
```

[[Back to top]](#)

## generic-post
:::warning experimental 
This API is currently in an experimental state. The API is subject to change based on feedback and further testing. You must include the X-SailPoint-Experimental header and set it to `true` to use this endpoint.
:::
:::tip setting x-sailpoint-experimental header
 on the configuration object you can set the `x-sailpoint-experimental` header to `true' to enable all experimantl endpoints within the SDK.
 Example:
 ```go
   configuration = Configuration()
   configuration.Experimental = true
 ```
:::
Generic POST request
Sends a POST request with a JSON body to any versioned API path.

[API Spec](https://developer.sailpoint.com/docs/api/generic-post)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | API path relative to the tenant base URL, including the API version, for example &#x60;accounts/v1&#x60;, &#x60;sources/v1/2c9180835d2e5168015d32f890ca1581&#x60; or &#x60;sources/v2/2c9180835d2e5168015d32f890ca1581/provisioning-policies&#x60;. Use whichever version the endpoint has. A leading slash is optional. Slashes separate path segments, and the SDK URL-encodes each segment, so pass the segments unencoded. Anything after a &#x60;?&#x60; is sent as query parameters, for example &#x60;accounts/v1?detailLevel&#x3D;SLIM&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenericPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSailPointExperimental** | **string** | Set this header to &#x60;true&#x60; to call an experimental API. The SDK configuration must also enable experimental APIs. | 
 **requestBody** | **interface{}** |  | 

### Return type

**interface{}**

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
  
    
	sailpoint "github.com/sailpoint-oss/golang-sdk/v3"
)

func main() {
    path := `accounts/v1` // string | API path relative to the tenant base URL, including the API version, for example `accounts/v1`, `sources/v1/2c9180835d2e5168015d32f890ca1581` or `sources/v2/2c9180835d2e5168015d32f890ca1581/provisioning-policies`. Use whichever version the endpoint has. A leading slash is optional. Slashes separate path segments, and the SDK URL-encodes each segment, so pass the segments unencoded. Anything after a `?` is sent as query parameters, for example `accounts/v1?detailLevel=SLIM`. # string | API path relative to the tenant base URL, including the API version, for example `accounts/v1`, `sources/v1/2c9180835d2e5168015d32f890ca1581` or `sources/v2/2c9180835d2e5168015d32f890ca1581/provisioning-policies`. Use whichever version the endpoint has. A leading slash is optional. Slashes separate path segments, and the SDK URL-encodes each segment, so pass the segments unencoded. Anything after a `?` is sent as query parameters, for example `accounts/v1?detailLevel=SLIM`.
    xSailPointExperimental := `true` // string | Set this header to `true` to call an experimental API. The SDK configuration must also enable experimental APIs. (optional) # string | Set this header to `true` to call an experimental API. The SDK configuration must also enable experimental APIs. (optional)
    requestBody :=  // interface{} |  (optional) # interface{} |  (optional)

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GenericPost(context.Background(), path).Execute()
	  //resp, r, err := apiClient.DefaultAPI.GenericPost(context.Background(), path).XSailPointExperimental(xSailPointExperimental).RequestBody(requestBody).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GenericPost``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenericPost`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GenericPost`: %v\n", resp)
}
```

[[Back to top]](#)

## generic-put
:::warning experimental 
This API is currently in an experimental state. The API is subject to change based on feedback and further testing. You must include the X-SailPoint-Experimental header and set it to `true` to use this endpoint.
:::
:::tip setting x-sailpoint-experimental header
 on the configuration object you can set the `x-sailpoint-experimental` header to `true' to enable all experimantl endpoints within the SDK.
 Example:
 ```go
   configuration = Configuration()
   configuration.Experimental = true
 ```
:::
Generic PUT request
Sends a PUT request with a JSON body to any versioned API path.

[API Spec](https://developer.sailpoint.com/docs/api/generic-put)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | API path relative to the tenant base URL, including the API version, for example &#x60;accounts/v1&#x60;, &#x60;sources/v1/2c9180835d2e5168015d32f890ca1581&#x60; or &#x60;sources/v2/2c9180835d2e5168015d32f890ca1581/provisioning-policies&#x60;. Use whichever version the endpoint has. A leading slash is optional. Slashes separate path segments, and the SDK URL-encodes each segment, so pass the segments unencoded. Anything after a &#x60;?&#x60; is sent as query parameters, for example &#x60;accounts/v1?detailLevel&#x3D;SLIM&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenericPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSailPointExperimental** | **string** | Set this header to &#x60;true&#x60; to call an experimental API. The SDK configuration must also enable experimental APIs. | 
 **requestBody** | **interface{}** |  | 

### Return type

**interface{}**

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
  
    
	sailpoint "github.com/sailpoint-oss/golang-sdk/v3"
)

func main() {
    path := `accounts/v1` // string | API path relative to the tenant base URL, including the API version, for example `accounts/v1`, `sources/v1/2c9180835d2e5168015d32f890ca1581` or `sources/v2/2c9180835d2e5168015d32f890ca1581/provisioning-policies`. Use whichever version the endpoint has. A leading slash is optional. Slashes separate path segments, and the SDK URL-encodes each segment, so pass the segments unencoded. Anything after a `?` is sent as query parameters, for example `accounts/v1?detailLevel=SLIM`. # string | API path relative to the tenant base URL, including the API version, for example `accounts/v1`, `sources/v1/2c9180835d2e5168015d32f890ca1581` or `sources/v2/2c9180835d2e5168015d32f890ca1581/provisioning-policies`. Use whichever version the endpoint has. A leading slash is optional. Slashes separate path segments, and the SDK URL-encodes each segment, so pass the segments unencoded. Anything after a `?` is sent as query parameters, for example `accounts/v1?detailLevel=SLIM`.
    xSailPointExperimental := `true` // string | Set this header to `true` to call an experimental API. The SDK configuration must also enable experimental APIs. (optional) # string | Set this header to `true` to call an experimental API. The SDK configuration must also enable experimental APIs. (optional)
    requestBody :=  // interface{} |  (optional) # interface{} |  (optional)

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GenericPut(context.Background(), path).Execute()
	  //resp, r, err := apiClient.DefaultAPI.GenericPut(context.Background(), path).XSailPointExperimental(xSailPointExperimental).RequestBody(requestBody).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GenericPut``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenericPut`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GenericPut`: %v\n", resp)
}
```

[[Back to top]](#)

