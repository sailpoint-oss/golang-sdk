---
id: nerm-attributes
title: Attributes
pagination_label: Attributes
sidebar_label: Attributes
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Attributes', 'NERMAttributes'] 
slug: /tools/sdk/go/nerm/methods/attributes
tags: ['SDK', 'Software Development Kit', 'Attributes', 'NERMAttributes']
---

# AttributesAPI
   
All URIs are relative to *https://acmeco.nonemployee.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create-attribute**](#create-attribute) | **Post** `/ne_attributes` | Create an attribute
[**delete-attribute-by-id**](#delete-attribute-by-id) | **Delete** `/ne_attributes/{id}` | Delete attribute by id
[**delete-attribute-by-uid**](#delete-attribute-by-uid) | **Delete** `/ne_attributes/{uid}` | Delete attribute
[**get-attribute-by-id**](#get-attribute-by-id) | **Get** `/ne_attributes/{id}` | Find attribute data by id
[**get-attribute-by-uid**](#get-attribute-by-uid) | **Get** `/ne_attributes/{uid}` | Find attribute data
[**get-attributes**](#get-attributes) | **Get** `/ne_attributes` | Get attribute data in bulk
[**update-attribute-by-id**](#update-attribute-by-id) | **Patch** `/ne_attributes/{id}` | Update attribute data by id
[**update-attribute-by-uid**](#update-attribute-by-uid) | **Patch** `/ne_attributes/{uid}` | Update attribute data


## create-attribute
Create an attribute
This endpoint can create an attribute

[API Spec](https://developer.sailpoint.com/docs/api/NERM/create-attribute)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAttributeRequest** | [**CreateAttributeRequest**](../models/create-attribute-request) |  | 

### Return type

[**CreateAttribute201Response**](../models/create-attribute201-response)

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
  "encoding/json"
    NERM "github.com/sailpoint-oss/golang-sdk/v2/api_nerm"
	sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    createattributerequest := []byte(``) // CreateAttributeRequest | 

    var createAttributeRequest NERM.CreateAttributeRequest
    if err := json.Unmarshal(createattributerequest, &createAttributeRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.AttributesAPI.CreateAttribute(context.Background()).CreateAttributeRequest(createAttributeRequest).Execute()
	  //resp, r, err := apiClient.NERM.AttributesAPI.CreateAttribute(context.Background()).CreateAttributeRequest(createAttributeRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `AttributesAPI.CreateAttribute``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAttribute`: CreateAttribute201Response
    fmt.Fprintf(os.Stdout, "Response from `AttributesAPI.CreateAttribute`: %v\n", resp)
}
```

[[Back to top]](#)

## delete-attribute-by-id
Delete attribute by id
Delete attribute by id

[API Spec](https://developer.sailpoint.com/docs/api/NERM/delete-attribute-by-id)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the object to retrieve, update, or delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAttributeByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateAttribute201Response**](../models/create-attribute201-response)

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
  
    
	sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    id := `1246d8b3-ac29-4015-8154-dea4434a73fa` // string | ID of the object to retrieve, update, or delete # string | ID of the object to retrieve, update, or delete

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.AttributesAPI.DeleteAttributeById(context.Background(), id).Execute()
	  //resp, r, err := apiClient.NERM.AttributesAPI.DeleteAttributeById(context.Background(), id).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `AttributesAPI.DeleteAttributeById``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAttributeById`: CreateAttribute201Response
    fmt.Fprintf(os.Stdout, "Response from `AttributesAPI.DeleteAttributeById`: %v\n", resp)
}
```

[[Back to top]](#)

## delete-attribute-by-uid
Delete attribute
Delete attribute by UID (user-specified identifier)

[API Spec](https://developer.sailpoint.com/docs/api/NERM/delete-attribute-by-uid)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | UID of the object to retrieve, update, or delete.  A UID or \&quot;specified identifier\&quot; is a string typically in \&quot;snake_case\&quot; format that provides a human-readable description of the record.  They are commonly used to ensure sandbox, qa, staging and production tenants have the identical configuration items loaded.  Every record has a UID assigned when persisted. When not specified the system assigns one by default.  A default value looks like a 32 character string of random hexadecimal characters. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAttributeByUidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateAttribute201Response**](../models/create-attribute201-response)

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
  
    
	sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    uid := `middle_initial_attribute` // string | UID of the object to retrieve, update, or delete.  A UID or \"specified identifier\" is a string typically in \"snake_case\" format that provides a human-readable description of the record.  They are commonly used to ensure sandbox, qa, staging and production tenants have the identical configuration items loaded.  Every record has a UID assigned when persisted. When not specified the system assigns one by default.  A default value looks like a 32 character string of random hexadecimal characters. (optional) # string | UID of the object to retrieve, update, or delete.  A UID or \"specified identifier\" is a string typically in \"snake_case\" format that provides a human-readable description of the record.  They are commonly used to ensure sandbox, qa, staging and production tenants have the identical configuration items loaded.  Every record has a UID assigned when persisted. When not specified the system assigns one by default.  A default value looks like a 32 character string of random hexadecimal characters. (optional)

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.AttributesAPI.DeleteAttributeByUid(context.Background(), uid).Execute()
	  //resp, r, err := apiClient.NERM.AttributesAPI.DeleteAttributeByUid(context.Background(), uid).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `AttributesAPI.DeleteAttributeByUid``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAttributeByUid`: CreateAttribute201Response
    fmt.Fprintf(os.Stdout, "Response from `AttributesAPI.DeleteAttributeByUid`: %v\n", resp)
}
```

[[Back to top]](#)

## get-attribute-by-id
Find attribute data by id
Info for a specific attribute

[API Spec](https://developer.sailpoint.com/docs/api/NERM/get-attribute-by-id)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the object to retrieve, update, or delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAttributeByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateAttribute201Response**](../models/create-attribute201-response)

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
  
    
	sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    id := `1246d8b3-ac29-4015-8154-dea4434a73fa` // string | ID of the object to retrieve, update, or delete # string | ID of the object to retrieve, update, or delete

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.AttributesAPI.GetAttributeById(context.Background(), id).Execute()
	  //resp, r, err := apiClient.NERM.AttributesAPI.GetAttributeById(context.Background(), id).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `AttributesAPI.GetAttributeById``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAttributeById`: CreateAttribute201Response
    fmt.Fprintf(os.Stdout, "Response from `AttributesAPI.GetAttributeById`: %v\n", resp)
}
```

[[Back to top]](#)

## get-attribute-by-uid
Find attribute data
Info for a specific attribute by UID (user-specified identifier)

[API Spec](https://developer.sailpoint.com/docs/api/NERM/get-attribute-by-uid)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | UID of the object to retrieve, update, or delete.  A UID or \&quot;specified identifier\&quot; is a string typically in \&quot;snake_case\&quot; format that provides a human-readable description of the record.  They are commonly used to ensure sandbox, qa, staging and production tenants have the identical configuration items loaded.  Every record has a UID assigned when persisted. When not specified the system assigns one by default.  A default value looks like a 32 character string of random hexadecimal characters. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAttributeByUidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateAttribute201Response**](../models/create-attribute201-response)

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
  
    
	sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    uid := `middle_initial_attribute` // string | UID of the object to retrieve, update, or delete.  A UID or \"specified identifier\" is a string typically in \"snake_case\" format that provides a human-readable description of the record.  They are commonly used to ensure sandbox, qa, staging and production tenants have the identical configuration items loaded.  Every record has a UID assigned when persisted. When not specified the system assigns one by default.  A default value looks like a 32 character string of random hexadecimal characters. (optional) # string | UID of the object to retrieve, update, or delete.  A UID or \"specified identifier\" is a string typically in \"snake_case\" format that provides a human-readable description of the record.  They are commonly used to ensure sandbox, qa, staging and production tenants have the identical configuration items loaded.  Every record has a UID assigned when persisted. When not specified the system assigns one by default.  A default value looks like a 32 character string of random hexadecimal characters. (optional)

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.AttributesAPI.GetAttributeByUid(context.Background(), uid).Execute()
	  //resp, r, err := apiClient.NERM.AttributesAPI.GetAttributeByUid(context.Background(), uid).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `AttributesAPI.GetAttributeByUid``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAttributeByUid`: CreateAttribute201Response
    fmt.Fprintf(os.Stdout, "Response from `AttributesAPI.GetAttributeByUid`: %v\n", resp)
}
```

[[Back to top]](#)

## get-attributes
Get attribute data in bulk
This endpoint can retrieve attribute data in bulk from Lifecycle or you can search for attributes using parameters

[API Spec](https://developer.sailpoint.com/docs/api/NERM/get-attributes)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The maximum number of items to return. | 
 **offset** | **int32** | The number of items to skip before starting to collect the result set. | 
 **order** | **string** | The field to order results by. | 
 **label** | **string** | The attribute label to filter by | 
 **dataType** | **string** | The attribute data type to filter by | 
 **metadata** | **bool** | Returns batching metadata in the response. This includes &#x60;total&#x60; as the total quantity, &#x60;next&#x60; as the path of the following query url, &#x60;limit&#x60; and &#x60;after_id&#x60; (if requested) with the next following id (null if it is the last \&quot;page\&quot;). | [default to false]

### Return type

[**GetAttributes200Response**](../models/get-attributes200-response)

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
  
    
	sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    limit := 5 // int32 | The maximum number of items to return. (optional) # int32 | The maximum number of items to return. (optional)
    offset := 5 // int32 | The number of items to skip before starting to collect the result set. (optional) # int32 | The number of items to skip before starting to collect the result set. (optional)
    order := `created_at` // string | The field to order results by. (optional) # string | The field to order results by. (optional)
    label := `birthday` // string | The attribute label to filter by (optional) # string | The attribute label to filter by (optional)
    dataType := `text field` // string | The attribute data type to filter by (optional) # string | The attribute data type to filter by (optional)
    metadata := true // bool | Returns batching metadata in the response. This includes `total` as the total quantity, `next` as the path of the following query url, `limit` and `after_id` (if requested) with the next following id (null if it is the last \"page\"). (optional) (default to false) # bool | Returns batching metadata in the response. This includes `total` as the total quantity, `next` as the path of the following query url, `limit` and `after_id` (if requested) with the next following id (null if it is the last \"page\"). (optional) (default to false)

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.AttributesAPI.GetAttributes(context.Background()).Execute()
	  //resp, r, err := apiClient.NERM.AttributesAPI.GetAttributes(context.Background()).Limit(limit).Offset(offset).Order(order).Label(label).DataType(dataType).Metadata(metadata).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `AttributesAPI.GetAttributes``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAttributes`: GetAttributes200Response
    fmt.Fprintf(os.Stdout, "Response from `AttributesAPI.GetAttributes`: %v\n", resp)
}
```

[[Back to top]](#)

## update-attribute-by-id
Update attribute data by id
Update info for a specific attribute

[API Spec](https://developer.sailpoint.com/docs/api/NERM/update-attribute-by-id)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the object to retrieve, update, or delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAttributeByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createAttributeRequest** | [**CreateAttributeRequest**](../models/create-attribute-request) |  | 

### Return type

[**CreateAttribute201Response**](../models/create-attribute201-response)

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
  "encoding/json"
    NERM "github.com/sailpoint-oss/golang-sdk/v2/api_nerm"
	sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    id := `1246d8b3-ac29-4015-8154-dea4434a73fa` // string | ID of the object to retrieve, update, or delete # string | ID of the object to retrieve, update, or delete
    createattributerequest := []byte(``) // CreateAttributeRequest | 

    var createAttributeRequest NERM.CreateAttributeRequest
    if err := json.Unmarshal(createattributerequest, &createAttributeRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.AttributesAPI.UpdateAttributeById(context.Background(), id).CreateAttributeRequest(createAttributeRequest).Execute()
	  //resp, r, err := apiClient.NERM.AttributesAPI.UpdateAttributeById(context.Background(), id).CreateAttributeRequest(createAttributeRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `AttributesAPI.UpdateAttributeById``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAttributeById`: CreateAttribute201Response
    fmt.Fprintf(os.Stdout, "Response from `AttributesAPI.UpdateAttributeById`: %v\n", resp)
}
```

[[Back to top]](#)

## update-attribute-by-uid
Update attribute data
Update info for a specific attribute by UID (user-specified identifier)

[API Spec](https://developer.sailpoint.com/docs/api/NERM/update-attribute-by-uid)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | UID of the object to retrieve, update, or delete.  A UID or \&quot;specified identifier\&quot; is a string typically in \&quot;snake_case\&quot; format that provides a human-readable description of the record.  They are commonly used to ensure sandbox, qa, staging and production tenants have the identical configuration items loaded.  Every record has a UID assigned when persisted. When not specified the system assigns one by default.  A default value looks like a 32 character string of random hexadecimal characters. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAttributeByUidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAttributeRequest** | [**CreateAttributeRequest**](../models/create-attribute-request) |  | 


### Return type

[**CreateAttribute201Response**](../models/create-attribute201-response)

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
  "encoding/json"
    NERM "github.com/sailpoint-oss/golang-sdk/v2/api_nerm"
	sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    createattributerequest := []byte(``) // CreateAttributeRequest | 
    uid := `middle_initial_attribute` // string | UID of the object to retrieve, update, or delete.  A UID or \"specified identifier\" is a string typically in \"snake_case\" format that provides a human-readable description of the record.  They are commonly used to ensure sandbox, qa, staging and production tenants have the identical configuration items loaded.  Every record has a UID assigned when persisted. When not specified the system assigns one by default.  A default value looks like a 32 character string of random hexadecimal characters. (optional) # string | UID of the object to retrieve, update, or delete.  A UID or \"specified identifier\" is a string typically in \"snake_case\" format that provides a human-readable description of the record.  They are commonly used to ensure sandbox, qa, staging and production tenants have the identical configuration items loaded.  Every record has a UID assigned when persisted. When not specified the system assigns one by default.  A default value looks like a 32 character string of random hexadecimal characters. (optional)

    var createAttributeRequest NERM.CreateAttributeRequest
    if err := json.Unmarshal(createattributerequest, &createAttributeRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.AttributesAPI.UpdateAttributeByUid(context.Background(), uid).CreateAttributeRequest(createAttributeRequest).Execute()
	  //resp, r, err := apiClient.NERM.AttributesAPI.UpdateAttributeByUid(context.Background(), uid).CreateAttributeRequest(createAttributeRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `AttributesAPI.UpdateAttributeByUid``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAttributeByUid`: CreateAttribute201Response
    fmt.Fprintf(os.Stdout, "Response from `AttributesAPI.UpdateAttributeByUid`: %v\n", resp)
}
```

[[Back to top]](#)

