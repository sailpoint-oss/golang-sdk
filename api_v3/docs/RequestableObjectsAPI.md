# \RequestableObjectsAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListRequestableObjects**](RequestableObjectsAPI.md#ListRequestableObjects) | **Get** /requestable-objects | Requestable Objects List



## Requestable Objects List


This endpoint returns a list of acccess items that that can be requested through the Access Request endpoints. Access items are marked with AVAILABLE, PENDING or ASSIGNED with respect to the identity provided using *identity-id* query param.
Any authenticated token can call this endpoint to see their requestable access items. A token with ORG_ADMIN authority is required to call this endpoint to return a list of all of the requestable access items for the org or for another identity.

### Parameters 
Param Type | Name | Data Type | Required  | Description
------------- | ------------- | ------------- | ------------- | ------------- 
  Query | identityId | **string** |   (optional) | If present, the value returns only requestable objects for the specified identity.  * Admin users can call this with any identity ID value.  * Non-admin users can only specify *me* or pass their own identity ID value.  * If absent, returns a list of all requestable objects for the tenant. Only admin users can make such a call. In this case, the available, pending, assigned accesses will not be annotated in the result.
  Query | types | [**[]RequestableObjectType**](RequestableObjectType.md) |   (optional) | Filters the results to the specified type/types, where each type is one of ROLE or ACCESS_PROFILE. If absent, all types are returned. Support for additional types may be added in the future without notice.
  Query | term | **string** |   (optional) | It allows searching requestable access items with a partial match on the name or description. If term is provided, then the *filter* query parameter will be ignored.
  Query | statuses | [**[]RequestableObjectRequestStatus**](RequestableObjectRequestStatus.md) |   (optional) | Filters the result to the specified status/statuses, where each status is one of AVAILABLE, ASSIGNED, or PENDING. It is an error to specify this parameter without also specifying an *identity-id* parameter. Additional statuses may be added in the future without notice.
  Query | limit | **int32** |   (optional) (default to 250) | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information.
  Query | offset | **int32** |   (optional) (default to 0) | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information.
  Query | count | **bool** |   (optional) (default to false) | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information.
  Query | filters | **string** |   (optional) | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq, in*  **name**: *eq, in, sw* 
  Query | sorters | **string** |   (optional) | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **name** 


### Return type

[**[]RequestableObject**](RequestableObject.md)

### Responses
Code | Description  | Data Type
------------- | ------------- | -------------
200 | List of requestable objects | []RequestableObject
400 | Client Error - Returned if the request body is invalid. | ErrorResponseDto
401 | Unauthorized - Returned if there is no authorization header, or if the JWT token is expired. | ListAccessProfiles401Response
403 | Forbidden - Returned if the user you are running as, doesn&#39;t have access to this end-point. | ErrorResponseDto
429 | Too Many Requests - Returned in response to too many requests in a given period of time - rate limited. The Retry-After header in the response includes how long to wait before trying again. | ListAccessProfiles429Response
500 | Internal Server Error - Returned if there is an unexpected error. | ErrorResponseDto


### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    v3 "github.com/sailpoint-oss/golang-sdk/v2/api_v3"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {

//ListRequestableObjects

    //identityId := "e7eab60924f64aa284175b9fa3309599"
    //types := []sailpoint.RequestableObjectType{sailpoint.RequestableObjectType("ACCESS_PROFILE")}
    //term := "Finance Role"
    //statuses := []sailpoint.RequestableObjectRequestStatus{sailpoint.RequestableObjectRequestStatus("AVAILABLE")}
    //limit := int32(250)
    //offset := int32(0)
    //count := true
    //filters := "name sw "bob""
    //sorters := "name"



    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.RequestableObjectsAPI.ListRequestableObjects(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RequestableObjectsAPI.ListRequestableObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRequestableObjects`: []RequestableObject
    fmt.Fprintf(os.Stdout, "Response from `RequestableObjectsAPI.ListRequestableObjects`: %v\n", resp)
}
```



