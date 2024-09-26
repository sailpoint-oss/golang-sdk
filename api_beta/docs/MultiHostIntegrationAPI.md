# \MultiHostIntegrationAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMultiHostIntegration**](MultiHostIntegrationAPI.md#CreateMultiHostIntegration) | **Post** /multihosts | Create Multi-Host Integration
[**CreateSourcesWithinMultiHost**](MultiHostIntegrationAPI.md#CreateSourcesWithinMultiHost) | **Post** /multihosts/{id} | Create Sources Within Multi-Host Integration
[**DeleteMultiHost**](MultiHostIntegrationAPI.md#DeleteMultiHost) | **Delete** /multihosts/{id} | Delete Multi-Host Integration
[**GetAcctAggregationGroups**](MultiHostIntegrationAPI.md#GetAcctAggregationGroups) | **Get** /multihosts/{multihostId}/acctAggregationGroups | Get Account Aggregation Groups Within Multi-Host Integration ID
[**GetEntitlementAggregationGroups**](MultiHostIntegrationAPI.md#GetEntitlementAggregationGroups) | **Get** /multihosts/{multiHostId}/entitlementAggregationGroups | Get Entitlement Aggregation Groups Within Multi-Host Integration ID
[**GetMultiHostIntegrations**](MultiHostIntegrationAPI.md#GetMultiHostIntegrations) | **Get** /multihosts/{id} | Get Multi-Host Integration By ID
[**GetMultiHostIntegrationsList**](MultiHostIntegrationAPI.md#GetMultiHostIntegrationsList) | **Get** /multihosts | List All Existing Multi-Host Integrations
[**GetMultiHostSourceCreationErrors**](MultiHostIntegrationAPI.md#GetMultiHostSourceCreationErrors) | **Get** /multihosts/{multiHostId}/sources/errors | List Multi-Host Source Creation Errors
[**GetMultihostIntegrationTypes**](MultiHostIntegrationAPI.md#GetMultihostIntegrationTypes) | **Get** /multihosts/types | List Multi-Host Integration Types
[**GetSourcesWithinMultiHost**](MultiHostIntegrationAPI.md#GetSourcesWithinMultiHost) | **Get** /multihosts/{id}/sources | List Sources Within Multi-Host Integration
[**TestConnectionMultiHostSources**](MultiHostIntegrationAPI.md#TestConnectionMultiHostSources) | **Post** /multihosts/{multihost_id}/sources/testConnection | Test Configuration For Multi-Host Integration
[**TestSourceConnectionMultihost**](MultiHostIntegrationAPI.md#TestSourceConnectionMultihost) | **Get** /multihosts/{multihost_id}/sources/{sourceId}/testConnection | Test Configuration For Multi-Host Integration&#39;s Single Source
[**UpdateMultiHostSources**](MultiHostIntegrationAPI.md#UpdateMultiHostSources) | **Patch** /multihosts/{id} | Update Multi-Host Integration



## CreateMultiHostIntegration

> MultiHostIntegrations CreateMultiHostIntegration(ctx).MultiHostIntegrationsCreate(multiHostIntegrationsCreate).Execute()

Create Multi-Host Integration



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
	multiHostIntegrationsCreate := *openapiclient.NewMultiHostIntegrationsCreate("My Multi-Host Integration", "This is the Multi-Host Integration.", *openapiclient.NewMultiHostIntegrationsOwner(), "multihost-microsoft-sql-server") // MultiHostIntegrationsCreate | The specifics of the Multi-Host Integration to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultiHostIntegrationAPI.CreateMultiHostIntegration(context.Background()).MultiHostIntegrationsCreate(multiHostIntegrationsCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiHostIntegrationAPI.CreateMultiHostIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMultiHostIntegration`: MultiHostIntegrations
	fmt.Fprintf(os.Stdout, "Response from `MultiHostIntegrationAPI.CreateMultiHostIntegration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMultiHostIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **multiHostIntegrationsCreate** | [**MultiHostIntegrationsCreate**](MultiHostIntegrationsCreate.md) | The specifics of the Multi-Host Integration to create | 

### Return type

[**MultiHostIntegrations**](MultiHostIntegrations.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSourcesWithinMultiHost

> CreateSourcesWithinMultiHost(ctx, id).MultiHostIntegrationsCreateSources(multiHostIntegrationsCreateSources).Execute()

Create Sources Within Multi-Host Integration



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
	id := "2c91808568c529c60168cca6f90c1326" // string | ID of the Multi-Host Integration.
	multiHostIntegrationsCreateSources := []openapiclient.MultiHostIntegrationsCreateSources{*openapiclient.NewMultiHostIntegrationsCreateSources("My Source")} // []MultiHostIntegrationsCreateSources | The specifics of the sources to create within Multi-Host Integration.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MultiHostIntegrationAPI.CreateSourcesWithinMultiHost(context.Background(), id).MultiHostIntegrationsCreateSources(multiHostIntegrationsCreateSources).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiHostIntegrationAPI.CreateSourcesWithinMultiHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Multi-Host Integration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSourcesWithinMultiHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **multiHostIntegrationsCreateSources** | [**[]MultiHostIntegrationsCreateSources**](MultiHostIntegrationsCreateSources.md) | The specifics of the sources to create within Multi-Host Integration. | 

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


## DeleteMultiHost

> DeleteMultiHost(ctx, id).Execute()

Delete Multi-Host Integration



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
	id := "2c91808568c529c60168cca6f90c1326" // string | ID of Multi-Host Integration to delete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MultiHostIntegrationAPI.DeleteMultiHost(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiHostIntegrationAPI.DeleteMultiHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Multi-Host Integration to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMultiHostRequest struct via the builder pattern


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


## GetAcctAggregationGroups

> MultiHostIntegrationsAggScheduleUpdate GetAcctAggregationGroups(ctx, multiHostId).Execute()

Get Account Aggregation Groups Within Multi-Host Integration ID



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
	multiHostId := "aMultiHostId" // string | ID of the Multi-Host Integration to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultiHostIntegrationAPI.GetAcctAggregationGroups(context.Background(), multiHostId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiHostIntegrationAPI.GetAcctAggregationGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAcctAggregationGroups`: MultiHostIntegrationsAggScheduleUpdate
	fmt.Fprintf(os.Stdout, "Response from `MultiHostIntegrationAPI.GetAcctAggregationGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**multiHostId** | **string** | ID of the Multi-Host Integration to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAcctAggregationGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MultiHostIntegrationsAggScheduleUpdate**](MultiHostIntegrationsAggScheduleUpdate.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEntitlementAggregationGroups

> MultiHostIntegrationsAggScheduleUpdate GetEntitlementAggregationGroups(ctx, multiHostId).Execute()

Get Entitlement Aggregation Groups Within Multi-Host Integration ID



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
	multiHostId := "aMultiHostId" // string | ID of the Multi-Host Integration to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultiHostIntegrationAPI.GetEntitlementAggregationGroups(context.Background(), multiHostId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiHostIntegrationAPI.GetEntitlementAggregationGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEntitlementAggregationGroups`: MultiHostIntegrationsAggScheduleUpdate
	fmt.Fprintf(os.Stdout, "Response from `MultiHostIntegrationAPI.GetEntitlementAggregationGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**multiHostId** | **string** | ID of the Multi-Host Integration to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEntitlementAggregationGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MultiHostIntegrationsAggScheduleUpdate**](MultiHostIntegrationsAggScheduleUpdate.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMultiHostIntegrations

> MultiHostIntegrations GetMultiHostIntegrations(ctx, id).Execute()

Get Multi-Host Integration By ID



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
	id := "2c91808568c529c60168cca6f90c1326" // string | ID of the Multi-Host Integration.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultiHostIntegrationAPI.GetMultiHostIntegrations(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiHostIntegrationAPI.GetMultiHostIntegrations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMultiHostIntegrations`: MultiHostIntegrations
	fmt.Fprintf(os.Stdout, "Response from `MultiHostIntegrationAPI.GetMultiHostIntegrations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Multi-Host Integration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMultiHostIntegrationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MultiHostIntegrations**](MultiHostIntegrations.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMultiHostIntegrationsList

> []MultiHostIntegrations GetMultiHostIntegrationsList(ctx).Offset(offset).Limit(limit).Sorters(sorters).Filters(filters).Count(count).ForSubadmin(forSubadmin).Execute()

List All Existing Multi-Host Integrations



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
	sorters := "name" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **name** (optional)
	filters := "id eq 2c91808b6ef1d43e016efba0ce470904" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **type**: *in*  **forSubAdminId**: *in* (optional)
	count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
	forSubadmin := "5168015d32f890ca15812c9180835d2e" // string | If provided, filters the returned list according to what is visible to the indicated ROLE_SUBADMIN Identity or SOURCE_SUBADMIN identity.  The value of the parameter is either an Identity ID, or the special value **me**, which is shorthand for the calling Identity's ID.  A 400 Bad Request error is returned if the **for-subadmin** parameter is specified for an Identity that is not a subadmin. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultiHostIntegrationAPI.GetMultiHostIntegrationsList(context.Background()).Offset(offset).Limit(limit).Sorters(sorters).Filters(filters).Count(count).ForSubadmin(forSubadmin).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiHostIntegrationAPI.GetMultiHostIntegrationsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMultiHostIntegrationsList`: []MultiHostIntegrations
	fmt.Fprintf(os.Stdout, "Response from `MultiHostIntegrationAPI.GetMultiHostIntegrationsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMultiHostIntegrationsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **name** | 
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **type**: *in*  **forSubAdminId**: *in* | 
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **forSubadmin** | **string** | If provided, filters the returned list according to what is visible to the indicated ROLE_SUBADMIN Identity or SOURCE_SUBADMIN identity.  The value of the parameter is either an Identity ID, or the special value **me**, which is shorthand for the calling Identity&#39;s ID.  A 400 Bad Request error is returned if the **for-subadmin** parameter is specified for an Identity that is not a subadmin. | 

### Return type

[**[]MultiHostIntegrations**](MultiHostIntegrations.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMultiHostSourceCreationErrors

> []SourceCreationErrors GetMultiHostSourceCreationErrors(ctx, multiHostId).Execute()

List Multi-Host Source Creation Errors



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
	multiHostId := "004091cb79b04636b88662afa50a4440" // string | ID of the Multi-Host Integration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultiHostIntegrationAPI.GetMultiHostSourceCreationErrors(context.Background(), multiHostId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiHostIntegrationAPI.GetMultiHostSourceCreationErrors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMultiHostSourceCreationErrors`: []SourceCreationErrors
	fmt.Fprintf(os.Stdout, "Response from `MultiHostIntegrationAPI.GetMultiHostSourceCreationErrors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**multiHostId** | **string** | ID of the Multi-Host Integration | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMultiHostSourceCreationErrorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]SourceCreationErrors**](SourceCreationErrors.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMultihostIntegrationTypes

> []MultiHostIntegrationTemplateType GetMultihostIntegrationTypes(ctx).Execute()

List Multi-Host Integration Types



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
	resp, r, err := apiClient.MultiHostIntegrationAPI.GetMultihostIntegrationTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiHostIntegrationAPI.GetMultihostIntegrationTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMultihostIntegrationTypes`: []MultiHostIntegrationTemplateType
	fmt.Fprintf(os.Stdout, "Response from `MultiHostIntegrationAPI.GetMultihostIntegrationTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMultihostIntegrationTypesRequest struct via the builder pattern


### Return type

[**[]MultiHostIntegrationTemplateType**](MultiHostIntegrationTemplateType.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSourcesWithinMultiHost

> []MultiHostSources GetSourcesWithinMultiHost(ctx).Offset(offset).Limit(limit).Sorters(sorters).Filters(filters).Count(count).Execute()

List Sources Within Multi-Host Integration



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
	sorters := "name" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **name** (optional)
	filters := "id eq 2c91808b6ef1d43e016efba0ce470904" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *in* (optional)
	count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultiHostIntegrationAPI.GetSourcesWithinMultiHost(context.Background()).Offset(offset).Limit(limit).Sorters(sorters).Filters(filters).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiHostIntegrationAPI.GetSourcesWithinMultiHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSourcesWithinMultiHost`: []MultiHostSources
	fmt.Fprintf(os.Stdout, "Response from `MultiHostIntegrationAPI.GetSourcesWithinMultiHost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSourcesWithinMultiHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **name** | 
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *in* | 
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]

### Return type

[**[]MultiHostSources**](MultiHostSources.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestConnectionMultiHostSources

> TestConnectionMultiHostSources(ctx, multihostId).Execute()

Test Configuration For Multi-Host Integration



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
	multihostId := "2c91808568c529c60168cca6f90c1324" // string | ID of the Multi-Host Integration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MultiHostIntegrationAPI.TestConnectionMultiHostSources(context.Background(), multihostId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiHostIntegrationAPI.TestConnectionMultiHostSources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**multihostId** | **string** | ID of the Multi-Host Integration | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestConnectionMultiHostSourcesRequest struct via the builder pattern


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


## TestSourceConnectionMultihost

> TestSourceConnectionMultihost200Response TestSourceConnectionMultihost(ctx, multihostId, sourceId).Execute()

Test Configuration For Multi-Host Integration's Single Source



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
	multihostId := "2c91808568c529c60168cca6f90c1326" // string | ID of the Multi-Host Integration
	sourceId := "2c91808568c529f60168cca6f90c1324" // string | ID of the source within the Multi-Host Integration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultiHostIntegrationAPI.TestSourceConnectionMultihost(context.Background(), multihostId, sourceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiHostIntegrationAPI.TestSourceConnectionMultihost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestSourceConnectionMultihost`: TestSourceConnectionMultihost200Response
	fmt.Fprintf(os.Stdout, "Response from `MultiHostIntegrationAPI.TestSourceConnectionMultihost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**multihostId** | **string** | ID of the Multi-Host Integration | 
**sourceId** | **string** | ID of the source within the Multi-Host Integration | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestSourceConnectionMultihostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TestSourceConnectionMultihost200Response**](TestSourceConnectionMultihost200Response.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMultiHostSources

> UpdateMultiHostSources(ctx, multihostId).UpdateMultiHostSourcesRequestInner(updateMultiHostSourcesRequestInner).Execute()

Update Multi-Host Integration



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
	multihostId := "anId" // string | ID of the Multi-Host Integration to update.
	updateMultiHostSourcesRequestInner := []openapiclient.UpdateMultiHostSourcesRequestInner{*openapiclient.NewUpdateMultiHostSourcesRequestInner("replace", "/description")} // []UpdateMultiHostSourcesRequestInner | This endpoint allows you to update a Multi-Host Integration. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MultiHostIntegrationAPI.UpdateMultiHostSources(context.Background(), multihostId).UpdateMultiHostSourcesRequestInner(updateMultiHostSourcesRequestInner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiHostIntegrationAPI.UpdateMultiHostSources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**multihostId** | **string** | ID of the Multi-Host Integration to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMultiHostSourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateMultiHostSourcesRequestInner** | [**[]UpdateMultiHostSourcesRequestInner**](UpdateMultiHostSourcesRequestInner.md) | This endpoint allows you to update a Multi-Host Integration.  | 

### Return type

 (empty response body)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

