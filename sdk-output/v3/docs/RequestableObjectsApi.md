# \RequestableObjectsApi

All URIs are relative to *https://sailpoint.api.identitynow.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListRequestableObjects**](RequestableObjectsApi.md#ListRequestableObjects) | **Get** /requestable-objects | Requestable Objects List



## ListRequestableObjects

> []RequestableObject ListRequestableObjects(ctx).IdentityId(identityId).Types(types).Term(term).Statuses(statuses).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).Execute()

Requestable Objects List



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
    identityId := "e7eab60924f64aa284175b9fa3309599" // string | If present, the value returns only requestable objects for the specified identity.  * Admin users can call this with any identity ID value.  * Non-admin users can only specify *me* or pass their own identity ID value.  * If absent, returns a list of all requestable objects for the tenant. Only admin users can make such a call. In this case, the available, pending, assigned accesses will not be annotated in the result. (optional)
    types := []openapiclient.RequestableObjectType{openapiclient.RequestableObjectType("ACCESS_PROFILE")} // []RequestableObjectType | Filters the results to the specified type/types, where each type is one of ROLE or ACCESS_PROFILE. If absent, all types are returned. Support for additional types may be added in the future without notice. (optional)
    term := "Finance Role" // string | It allows searching requestable access items with a partial match on the name or description. If term is provided, then the *filter* query parameter will be ignored. (optional)
    statuses := []openapiclient.RequestableObjectRequestStatus{openapiclient.RequestableObjectRequestStatus("AVAILABLE")} // []RequestableObjectRequestStatus | Filters the result to the specified status/statuses, where each status is one of AVAILABLE, ASSIGNED, or PENDING. It is an error to specify this parameter without also specifying an *identity-id* parameter. Additional statuses may be added in the future without notice. (optional)
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
    filters := "name sw "bob"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq, in*  **name**: *eq, in, sw*  (optional)
    sorters := "name" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **name**  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RequestableObjectsApi.ListRequestableObjects(context.Background()).IdentityId(identityId).Types(types).Term(term).Statuses(statuses).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RequestableObjectsApi.ListRequestableObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRequestableObjects`: []RequestableObject
    fmt.Fprintf(os.Stdout, "Response from `RequestableObjectsApi.ListRequestableObjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRequestableObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identityId** | **string** | If present, the value returns only requestable objects for the specified identity.  * Admin users can call this with any identity ID value.  * Non-admin users can only specify *me* or pass their own identity ID value.  * If absent, returns a list of all requestable objects for the tenant. Only admin users can make such a call. In this case, the available, pending, assigned accesses will not be annotated in the result. | 
 **types** | [**[]RequestableObjectType**](RequestableObjectType.md) | Filters the results to the specified type/types, where each type is one of ROLE or ACCESS_PROFILE. If absent, all types are returned. Support for additional types may be added in the future without notice. | 
 **term** | **string** | It allows searching requestable access items with a partial match on the name or description. If term is provided, then the *filter* query parameter will be ignored. | 
 **statuses** | [**[]RequestableObjectRequestStatus**](RequestableObjectRequestStatus.md) | Filters the result to the specified status/statuses, where each status is one of AVAILABLE, ASSIGNED, or PENDING. It is an error to specify this parameter without also specifying an *identity-id* parameter. Additional statuses may be added in the future without notice. | 
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq, in*  **name**: *eq, in, sw*  | 
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **name**  | 

### Return type

[**[]RequestableObject**](RequestableObject.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

