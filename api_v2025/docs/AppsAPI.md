# \AppsAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v2025*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSourceApp**](AppsAPI.md#CreateSourceApp) | **Post** /source-apps | Create source app
[**DeleteAccessProfilesFromSourceAppByBulk**](AppsAPI.md#DeleteAccessProfilesFromSourceAppByBulk) | **Post** /source-apps/{id}/access-profiles/bulk-remove | Bulk remove access profiles from the specified source app
[**DeleteSourceApp**](AppsAPI.md#DeleteSourceApp) | **Delete** /source-apps/{id} | Delete source app by ID
[**GetSourceApp**](AppsAPI.md#GetSourceApp) | **Get** /source-apps/{id} | Get source app by ID
[**ListAccessProfilesForSourceApp**](AppsAPI.md#ListAccessProfilesForSourceApp) | **Get** /source-apps/{id}/access-profiles | List access profiles for the specified source app
[**ListAllSourceApp**](AppsAPI.md#ListAllSourceApp) | **Get** /source-apps/all | List all source apps
[**ListAllUserApps**](AppsAPI.md#ListAllUserApps) | **Get** /user-apps/all | List all user apps
[**ListAssignedSourceApp**](AppsAPI.md#ListAssignedSourceApp) | **Get** /source-apps/assigned | List assigned source apps
[**ListAvailableAccountsForUserApp**](AppsAPI.md#ListAvailableAccountsForUserApp) | **Get** /user-apps/{id}/available-accounts | List available accounts for user app
[**ListAvailableSourceApps**](AppsAPI.md#ListAvailableSourceApps) | **Get** /source-apps | List available source apps
[**ListOwnedUserApps**](AppsAPI.md#ListOwnedUserApps) | **Get** /user-apps | List owned user apps
[**PatchSourceApp**](AppsAPI.md#PatchSourceApp) | **Patch** /source-apps/{id} | Patch source app by ID
[**PatchUserApp**](AppsAPI.md#PatchUserApp) | **Patch** /user-apps/{id} | Patch user app by ID
[**UpdateSourceAppsInBulk**](AppsAPI.md#UpdateSourceAppsInBulk) | **Post** /source-apps/bulk-update | Bulk update source apps



## CreateSourceApp

> SourceApp CreateSourceApp(ctx).XSailPointExperimental(xSailPointExperimental).SourceAppCreateDto(sourceAppCreateDto).Execute()

Create source app



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
	xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")
	sourceAppCreateDto := *openapiclient.NewSourceAppCreateDto("my app", "the source app for engineers", *openapiclient.NewSourceAppCreateDtoAccountSource("2c9180827ca885d7017ca8ce28a000eb")) // SourceAppCreateDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.CreateSourceApp(context.Background()).XSailPointExperimental(xSailPointExperimental).SourceAppCreateDto(sourceAppCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.CreateSourceApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSourceApp`: SourceApp
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.CreateSourceApp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSourceAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **sourceAppCreateDto** | [**SourceAppCreateDto**](SourceAppCreateDto.md) |  | 

### Return type

[**SourceApp**](SourceApp.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAccessProfilesFromSourceAppByBulk

> []AccessProfileDetails DeleteAccessProfilesFromSourceAppByBulk(ctx, id).XSailPointExperimental(xSailPointExperimental).RequestBody(requestBody).Limit(limit).Execute()

Bulk remove access profiles from the specified source app



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
	id := "2c91808a7813090a017814121e121518" // string | ID of the source app
	xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")
	requestBody := []string{"Property_example"} // []string | 
	limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.DeleteAccessProfilesFromSourceAppByBulk(context.Background(), id).XSailPointExperimental(xSailPointExperimental).RequestBody(requestBody).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.DeleteAccessProfilesFromSourceAppByBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAccessProfilesFromSourceAppByBulk`: []AccessProfileDetails
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.DeleteAccessProfilesFromSourceAppByBulk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the source app | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccessProfilesFromSourceAppByBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **requestBody** | **[]string** |  | 
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]

### Return type

[**[]AccessProfileDetails**](AccessProfileDetails.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSourceApp

> SourceApp DeleteSourceApp(ctx, id).XSailPointExperimental(xSailPointExperimental).Execute()

Delete source app by ID



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
	id := "2c9180835d191a86015d28455b4a2329" // string | source app ID.
	xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.DeleteSourceApp(context.Background(), id).XSailPointExperimental(xSailPointExperimental).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.DeleteSourceApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSourceApp`: SourceApp
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.DeleteSourceApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | source app ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSourceAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]

### Return type

[**SourceApp**](SourceApp.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSourceApp

> SourceApp GetSourceApp(ctx, id).XSailPointExperimental(xSailPointExperimental).Execute()

Get source app by ID



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
	id := "2c91808a7813090a017814121e121518" // string | ID of the source app
	xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.GetSourceApp(context.Background(), id).XSailPointExperimental(xSailPointExperimental).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.GetSourceApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSourceApp`: SourceApp
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.GetSourceApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the source app | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSourceAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]

### Return type

[**SourceApp**](SourceApp.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccessProfilesForSourceApp

> []AccessProfileDetails ListAccessProfilesForSourceApp(ctx, id).XSailPointExperimental(xSailPointExperimental).Limit(limit).Offset(offset).Filters(filters).Execute()

List access profiles for the specified source app



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
	id := "2c91808a7813090a017814121e121518" // string | ID of the source app
	xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")
	limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
	offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
	filters := "name eq "developer access profile"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq, in*  **name**: *eq, in*  **created**: *gt, lt, ge, le*  **modified**: *gt, lt, ge, le* (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.ListAccessProfilesForSourceApp(context.Background(), id).XSailPointExperimental(xSailPointExperimental).Limit(limit).Offset(offset).Filters(filters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.ListAccessProfilesForSourceApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAccessProfilesForSourceApp`: []AccessProfileDetails
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.ListAccessProfilesForSourceApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the source app | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAccessProfilesForSourceAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq, in*  **name**: *eq, in*  **created**: *gt, lt, ge, le*  **modified**: *gt, lt, ge, le* | 

### Return type

[**[]AccessProfileDetails**](AccessProfileDetails.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllSourceApp

> []SourceApp ListAllSourceApp(ctx).XSailPointExperimental(xSailPointExperimental).Limit(limit).Count(count).Offset(offset).Sorters(sorters).Filters(filters).Execute()

List all source apps



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
	xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")
	limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
	count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
	offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
	sorters := "name,-modified" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **id, name, created, modified, owner.id, accountSource.id** (optional)
	filters := "enabled eq true" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq, in*  **name**: *eq, in, co, sw*  **created**: *gt, lt, ge, le*  **modified**: *gt, lt, ge, le*  **owner.id**: *eq, in*  **accountSource.id**: *eq, in*  **enabled**: *eq* (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.ListAllSourceApp(context.Background()).XSailPointExperimental(xSailPointExperimental).Limit(limit).Count(count).Offset(offset).Sorters(sorters).Filters(filters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.ListAllSourceApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAllSourceApp`: []SourceApp
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.ListAllSourceApp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAllSourceAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **id, name, created, modified, owner.id, accountSource.id** | 
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq, in*  **name**: *eq, in, co, sw*  **created**: *gt, lt, ge, le*  **modified**: *gt, lt, ge, le*  **owner.id**: *eq, in*  **accountSource.id**: *eq, in*  **enabled**: *eq* | 

### Return type

[**[]SourceApp**](SourceApp.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllUserApps

> []UserApp ListAllUserApps(ctx).Filters(filters).XSailPointExperimental(xSailPointExperimental).Limit(limit).Count(count).Offset(offset).Execute()

List all user apps



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
	filters := "name eq "user app name"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq*  **ownerId**: *eq*  **ownerName**: *eq, sw*  **ownerAlias**: *eq, sw*  **accountId**: *eq*  **sourceAppId**: *eq*
	xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")
	limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
	count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
	offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.ListAllUserApps(context.Background()).Filters(filters).XSailPointExperimental(xSailPointExperimental).Limit(limit).Count(count).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.ListAllUserApps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAllUserApps`: []UserApp
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.ListAllUserApps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAllUserAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq*  **ownerId**: *eq*  **ownerName**: *eq, sw*  **ownerAlias**: *eq, sw*  **accountId**: *eq*  **sourceAppId**: *eq* | 
 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]

### Return type

[**[]UserApp**](UserApp.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAssignedSourceApp

> []SourceApp ListAssignedSourceApp(ctx).XSailPointExperimental(xSailPointExperimental).Limit(limit).Count(count).Offset(offset).Sorters(sorters).Filters(filters).Execute()

List assigned source apps



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
	xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")
	limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
	count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
	offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
	sorters := "name,-modified" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **id, name, created, modified, accountSource.id** (optional)
	filters := "name eq "source app name"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq, in*  **name**: *eq, in, co, sw*  **created**: *gt, lt, ge, le*  **modified**: *gt, lt, ge, le*  **accountSource.id**: *eq, in* (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.ListAssignedSourceApp(context.Background()).XSailPointExperimental(xSailPointExperimental).Limit(limit).Count(count).Offset(offset).Sorters(sorters).Filters(filters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.ListAssignedSourceApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAssignedSourceApp`: []SourceApp
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.ListAssignedSourceApp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAssignedSourceAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **id, name, created, modified, accountSource.id** | 
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq, in*  **name**: *eq, in, co, sw*  **created**: *gt, lt, ge, le*  **modified**: *gt, lt, ge, le*  **accountSource.id**: *eq, in* | 

### Return type

[**[]SourceApp**](SourceApp.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAvailableAccountsForUserApp

> []AppAccountDetails ListAvailableAccountsForUserApp(ctx, id).XSailPointExperimental(xSailPointExperimental).Limit(limit).Count(count).Execute()

List available accounts for user app



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
	id := "2c91808a7813090a017814121e121518" // string | ID of the user app
	xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")
	limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
	count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.ListAvailableAccountsForUserApp(context.Background(), id).XSailPointExperimental(xSailPointExperimental).Limit(limit).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.ListAvailableAccountsForUserApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAvailableAccountsForUserApp`: []AppAccountDetails
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.ListAvailableAccountsForUserApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the user app | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAvailableAccountsForUserAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]

### Return type

[**[]AppAccountDetails**](AppAccountDetails.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAvailableSourceApps

> []SourceApp ListAvailableSourceApps(ctx).XSailPointExperimental(xSailPointExperimental).Limit(limit).Count(count).Offset(offset).Sorters(sorters).Filters(filters).Execute()

List available source apps



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
	xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")
	limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
	count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
	offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
	sorters := "name,-modified" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **id, name, created, modified, owner.id, accountSource.id** (optional)
	filters := "name eq "source app name"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq, in*  **name**: *eq, in, co, sw*  **created**: *gt, lt, ge, le*  **modified**: *gt, lt, ge, le*  **accountSource.id**: *eq, in* (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.ListAvailableSourceApps(context.Background()).XSailPointExperimental(xSailPointExperimental).Limit(limit).Count(count).Offset(offset).Sorters(sorters).Filters(filters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.ListAvailableSourceApps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAvailableSourceApps`: []SourceApp
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.ListAvailableSourceApps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAvailableSourceAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **id, name, created, modified, owner.id, accountSource.id** | 
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq, in*  **name**: *eq, in, co, sw*  **created**: *gt, lt, ge, le*  **modified**: *gt, lt, ge, le*  **accountSource.id**: *eq, in* | 

### Return type

[**[]SourceApp**](SourceApp.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOwnedUserApps

> []UserApp ListOwnedUserApps(ctx).XSailPointExperimental(xSailPointExperimental).Limit(limit).Count(count).Offset(offset).Filters(filters).Execute()

List owned user apps



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
	xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")
	limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
	count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
	offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
	filters := "name eq "user app name"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq*  **ownerName**: *eq, sw*  **ownerAlias**: *eq, sw*  **accountId**: *eq*  **sourceAppId**: *eq* (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.ListOwnedUserApps(context.Background()).XSailPointExperimental(xSailPointExperimental).Limit(limit).Count(count).Offset(offset).Filters(filters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.ListOwnedUserApps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOwnedUserApps`: []UserApp
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.ListOwnedUserApps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOwnedUserAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq*  **ownerName**: *eq, sw*  **ownerAlias**: *eq, sw*  **accountId**: *eq*  **sourceAppId**: *eq* | 

### Return type

[**[]UserApp**](UserApp.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchSourceApp

> SourceAppPatchDto PatchSourceApp(ctx, id).XSailPointExperimental(xSailPointExperimental).JsonPatchOperation(jsonPatchOperation).Execute()

Patch source app by ID



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
	id := "2c91808a7813090a017814121e121518" // string | ID of the source app to patch
	xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")
	jsonPatchOperation := []openapiclient.JsonPatchOperation{*openapiclient.NewJsonPatchOperation("replace", "/description")} // []JsonPatchOperation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.PatchSourceApp(context.Background(), id).XSailPointExperimental(xSailPointExperimental).JsonPatchOperation(jsonPatchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.PatchSourceApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchSourceApp`: SourceAppPatchDto
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.PatchSourceApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the source app to patch | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSourceAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **jsonPatchOperation** | [**[]JsonPatchOperation**](JsonPatchOperation.md) |  | 

### Return type

[**SourceAppPatchDto**](SourceAppPatchDto.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchUserApp

> UserApp PatchUserApp(ctx, id).XSailPointExperimental(xSailPointExperimental).JsonPatchOperation(jsonPatchOperation).Execute()

Patch user app by ID



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
	id := "2c91808a7813090a017814121e121518" // string | ID of the user app to patch
	xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")
	jsonPatchOperation := []openapiclient.JsonPatchOperation{*openapiclient.NewJsonPatchOperation("replace", "/description")} // []JsonPatchOperation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.PatchUserApp(context.Background(), id).XSailPointExperimental(xSailPointExperimental).JsonPatchOperation(jsonPatchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.PatchUserApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchUserApp`: UserApp
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.PatchUserApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the user app to patch | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchUserAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **jsonPatchOperation** | [**[]JsonPatchOperation**](JsonPatchOperation.md) |  | 

### Return type

[**UserApp**](UserApp.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSourceAppsInBulk

> UpdateSourceAppsInBulk(ctx).XSailPointExperimental(xSailPointExperimental).SourceAppBulkUpdateRequest(sourceAppBulkUpdateRequest).Execute()

Bulk update source apps



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
	xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")
	sourceAppBulkUpdateRequest := *openapiclient.NewSourceAppBulkUpdateRequest([]string{"AppIds_example"}, []openapiclient.JsonPatchOperation{*openapiclient.NewJsonPatchOperation("replace", "/description")}) // SourceAppBulkUpdateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AppsAPI.UpdateSourceAppsInBulk(context.Background()).XSailPointExperimental(xSailPointExperimental).SourceAppBulkUpdateRequest(sourceAppBulkUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.UpdateSourceAppsInBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSourceAppsInBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **sourceAppBulkUpdateRequest** | [**SourceAppBulkUpdateRequest**](SourceAppBulkUpdateRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

