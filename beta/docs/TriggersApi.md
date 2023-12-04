# \TriggersAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CompleteTriggerInvocation**](TriggersAPI.md#CompleteTriggerInvocation) | **Post** /trigger-invocations/{id}/complete | Complete Trigger Invocation
[**CreateSubscription**](TriggersAPI.md#CreateSubscription) | **Post** /trigger-subscriptions | Create a Subscription
[**DeleteSubscription**](TriggersAPI.md#DeleteSubscription) | **Delete** /trigger-subscriptions/{id} | Delete a Subscription
[**ListSubscriptions**](TriggersAPI.md#ListSubscriptions) | **Get** /trigger-subscriptions | List Subscriptions
[**ListTriggerInvocationStatus**](TriggersAPI.md#ListTriggerInvocationStatus) | **Get** /trigger-invocations/status | List Latest Invocation Statuses
[**ListTriggers**](TriggersAPI.md#ListTriggers) | **Get** /triggers | List Triggers
[**PatchSubscription**](TriggersAPI.md#PatchSubscription) | **Patch** /trigger-subscriptions/{id} | Patch a Subscription
[**StartTestTriggerInvocation**](TriggersAPI.md#StartTestTriggerInvocation) | **Post** /trigger-invocations/test | Start a Test Invocation
[**TestSubscriptionFilter**](TriggersAPI.md#TestSubscriptionFilter) | **Post** /trigger-subscriptions/validate-filter | Validate a Subscription Filter
[**UpdateSubscription**](TriggersAPI.md#UpdateSubscription) | **Put** /trigger-subscriptions/{id} | Update a Subscription



## CompleteTriggerInvocation

> CompleteTriggerInvocation(ctx, id).CompleteInvocation(completeInvocation).Execute()

Complete Trigger Invocation



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/sailpoint-oss/golang-sdk"
)

func main() {
    id := "0f11f2a4-7c94-4bf3-a2bd-742580fe3bde" // string | The ID of the invocation to complete.
    completeInvocation := *openapiclient.NewCompleteInvocation("0f11f2a4-7c94-4bf3-a2bd-742580fe3bde", map[string]interface{}({approved=false})) // CompleteInvocation | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TriggersAPI.CompleteTriggerInvocation(context.Background(), id).CompleteInvocation(completeInvocation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TriggersAPI.CompleteTriggerInvocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the invocation to complete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompleteTriggerInvocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **completeInvocation** | [**CompleteInvocation**](CompleteInvocation.md) |  | 

### Return type

 (empty response body)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSubscription

> Subscription CreateSubscription(ctx).SubscriptionPostRequest(subscriptionPostRequest).Execute()

Create a Subscription



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/sailpoint-oss/golang-sdk"
)

func main() {
    subscriptionPostRequest := *openapiclient.NewSubscriptionPostRequest("Access request subscription", "idn:access-requested", openapiclient.SubscriptionType("HTTP")) // SubscriptionPostRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TriggersAPI.CreateSubscription(context.Background()).SubscriptionPostRequest(subscriptionPostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TriggersAPI.CreateSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSubscription`: Subscription
    fmt.Fprintf(os.Stdout, "Response from `TriggersAPI.CreateSubscription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionPostRequest** | [**SubscriptionPostRequest**](SubscriptionPostRequest.md) |  | 

### Return type

[**Subscription**](Subscription.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSubscription

> DeleteSubscription(ctx, id).Execute()

Delete a Subscription



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/sailpoint-oss/golang-sdk"
)

func main() {
    id := "0f11f2a4-7c94-4bf3-a2bd-742580fe3bde" // string | Subscription ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TriggersAPI.DeleteSubscription(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TriggersAPI.DeleteSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSubscriptions

> []Subscription ListSubscriptions(ctx).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).Execute()

List Subscriptions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/sailpoint-oss/golang-sdk"
)

func main() {
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
    filters := "id eq "12cff757-c0c0-413b-8ad7-2a47956d1e89"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq*  **triggerId**: *eq*  **type**: *eq, le* (optional)
    sorters := "triggerName" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **triggerId, triggerName** (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TriggersAPI.ListSubscriptions(context.Background()).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TriggersAPI.ListSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSubscriptions`: []Subscription
    fmt.Fprintf(os.Stdout, "Response from `TriggersAPI.ListSubscriptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq*  **triggerId**: *eq*  **type**: *eq, le* | 
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **triggerId, triggerName** | 

### Return type

[**[]Subscription**](Subscription.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTriggerInvocationStatus

> []InvocationStatus ListTriggerInvocationStatus(ctx).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).Execute()

List Latest Invocation Statuses



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/sailpoint-oss/golang-sdk"
)

func main() {
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
    filters := "triggerId eq "idn:access-request-dynamic-approver"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **triggerId**: *eq*  **subscriptionId**: *eq* (optional)
    sorters := "created" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **triggerId, subscriptionName, created, completed** (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TriggersAPI.ListTriggerInvocationStatus(context.Background()).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TriggersAPI.ListTriggerInvocationStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTriggerInvocationStatus`: []InvocationStatus
    fmt.Fprintf(os.Stdout, "Response from `TriggersAPI.ListTriggerInvocationStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTriggerInvocationStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **triggerId**: *eq*  **subscriptionId**: *eq* | 
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **triggerId, subscriptionName, created, completed** | 

### Return type

[**[]InvocationStatus**](InvocationStatus.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTriggers

> []Trigger ListTriggers(ctx).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).Execute()

List Triggers



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/sailpoint-oss/golang-sdk"
)

func main() {
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
    filters := "id eq "idn:access-request-post-approval"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq, ge, le* (optional)
    sorters := "name" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **id, name** (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TriggersAPI.ListTriggers(context.Background()).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TriggersAPI.ListTriggers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTriggers`: []Trigger
    fmt.Fprintf(os.Stdout, "Response from `TriggersAPI.ListTriggers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTriggersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq, ge, le* | 
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **id, name** | 

### Return type

[**[]Trigger**](Trigger.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchSubscription

> Subscription PatchSubscription(ctx, id).SubscriptionPatchRequestInner(subscriptionPatchRequestInner).Execute()

Patch a Subscription



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/sailpoint-oss/golang-sdk"
)

func main() {
    id := "0f11f2a4-7c94-4bf3-a2bd-742580fe3bde" // string | ID of the Subscription to patch
    subscriptionPatchRequestInner := []openapiclient.SubscriptionPatchRequestInner{*openapiclient.NewSubscriptionPatchRequestInner("replace", "/description")} // []SubscriptionPatchRequestInner | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TriggersAPI.PatchSubscription(context.Background(), id).SubscriptionPatchRequestInner(subscriptionPatchRequestInner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TriggersAPI.PatchSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchSubscription`: Subscription
    fmt.Fprintf(os.Stdout, "Response from `TriggersAPI.PatchSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Subscription to patch | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subscriptionPatchRequestInner** | [**[]SubscriptionPatchRequestInner**](SubscriptionPatchRequestInner.md) |  | 

### Return type

[**Subscription**](Subscription.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartTestTriggerInvocation

> []Invocation StartTestTriggerInvocation(ctx).TestInvocation(testInvocation).Execute()

Start a Test Invocation



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/sailpoint-oss/golang-sdk"
)

func main() {
    testInvocation := *openapiclient.NewTestInvocation("idn:access-request-post-approval", map[string]interface{}({workflowId=1234})) // TestInvocation | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TriggersAPI.StartTestTriggerInvocation(context.Background()).TestInvocation(testInvocation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TriggersAPI.StartTestTriggerInvocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartTestTriggerInvocation`: []Invocation
    fmt.Fprintf(os.Stdout, "Response from `TriggersAPI.StartTestTriggerInvocation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStartTestTriggerInvocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **testInvocation** | [**TestInvocation**](TestInvocation.md) |  | 

### Return type

[**[]Invocation**](Invocation.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestSubscriptionFilter

> ValidateFilterOutputDto TestSubscriptionFilter(ctx).ValidateFilterInputDto(validateFilterInputDto).Execute()

Validate a Subscription Filter



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/sailpoint-oss/golang-sdk"
)

func main() {
    validateFilterInputDto := *openapiclient.NewValidateFilterInputDto(map[string]interface{}({identityId=201327fda1c44704ac01181e963d463c}), "$[?($.identityId == "201327fda1c44704ac01181e963d463c")]") // ValidateFilterInputDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TriggersAPI.TestSubscriptionFilter(context.Background()).ValidateFilterInputDto(validateFilterInputDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TriggersAPI.TestSubscriptionFilter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestSubscriptionFilter`: ValidateFilterOutputDto
    fmt.Fprintf(os.Stdout, "Response from `TriggersAPI.TestSubscriptionFilter`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestSubscriptionFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **validateFilterInputDto** | [**ValidateFilterInputDto**](ValidateFilterInputDto.md) |  | 

### Return type

[**ValidateFilterOutputDto**](ValidateFilterOutputDto.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSubscription

> Subscription UpdateSubscription(ctx, id).SubscriptionPutRequest(subscriptionPutRequest).Execute()

Update a Subscription



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/sailpoint-oss/golang-sdk"
)

func main() {
    id := "0f11f2a4-7c94-4bf3-a2bd-742580fe3bde" // string | Subscription ID
    subscriptionPutRequest := *openapiclient.NewSubscriptionPutRequest() // SubscriptionPutRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TriggersAPI.UpdateSubscription(context.Background(), id).SubscriptionPutRequest(subscriptionPutRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TriggersAPI.UpdateSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSubscription`: Subscription
    fmt.Fprintf(os.Stdout, "Response from `TriggersAPI.UpdateSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subscriptionPutRequest** | [**SubscriptionPutRequest**](SubscriptionPutRequest.md) |  | 

### Return type

[**Subscription**](Subscription.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

