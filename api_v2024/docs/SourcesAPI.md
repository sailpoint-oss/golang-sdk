# \SourcesAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v2024*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProvisioningPolicy**](SourcesAPI.md#CreateProvisioningPolicy) | **Post** /sources/{sourceId}/provisioning-policies | Create Provisioning Policy
[**CreateSource**](SourcesAPI.md#CreateSource) | **Post** /sources | Creates a source in IdentityNow.
[**CreateSourceSchedule**](SourcesAPI.md#CreateSourceSchedule) | **Post** /sources/{sourceId}/schedules | Create Schedule on Source
[**CreateSourceSchema**](SourcesAPI.md#CreateSourceSchema) | **Post** /sources/{sourceId}/schemas | Create Schema on Source
[**DeleteAccountsAsync**](SourcesAPI.md#DeleteAccountsAsync) | **Post** /sources/{id}/remove-accounts | Remove All Accounts in a Source
[**DeleteNativeChangeDetectionConfig**](SourcesAPI.md#DeleteNativeChangeDetectionConfig) | **Delete** /sources/{sourceId}/native-change-detection-config | Delete Native Change Detection Configuration
[**DeleteProvisioningPolicy**](SourcesAPI.md#DeleteProvisioningPolicy) | **Delete** /sources/{sourceId}/provisioning-policies/{usageType} | Delete Provisioning Policy by UsageType
[**DeleteSource**](SourcesAPI.md#DeleteSource) | **Delete** /sources/{id} | Delete Source by ID
[**DeleteSourceSchedule**](SourcesAPI.md#DeleteSourceSchedule) | **Delete** /sources/{sourceId}/schedules/{scheduleType} | Delete Source Schedule by type.
[**DeleteSourceSchema**](SourcesAPI.md#DeleteSourceSchema) | **Delete** /sources/{sourceId}/schemas/{schemaId} | Delete Source Schema by ID
[**GetAccountsSchema**](SourcesAPI.md#GetAccountsSchema) | **Get** /sources/{id}/schemas/accounts | Downloads source accounts schema template
[**GetCorrelationConfig**](SourcesAPI.md#GetCorrelationConfig) | **Get** /sources/{id}/correlation-config | Get Source Correlation Configuration
[**GetEntitlementsSchema**](SourcesAPI.md#GetEntitlementsSchema) | **Get** /sources/{id}/schemas/entitlements | Downloads source entitlements schema template
[**GetNativeChangeDetectionConfig**](SourcesAPI.md#GetNativeChangeDetectionConfig) | **Get** /sources/{sourceId}/native-change-detection-config | Native Change Detection Configuration
[**GetProvisioningPolicy**](SourcesAPI.md#GetProvisioningPolicy) | **Get** /sources/{sourceId}/provisioning-policies/{usageType} | Get Provisioning Policy by UsageType
[**GetSource**](SourcesAPI.md#GetSource) | **Get** /sources/{id} | Get Source by ID
[**GetSourceAttrSyncConfig**](SourcesAPI.md#GetSourceAttrSyncConfig) | **Get** /sources/{id}/attribute-sync-config | Attribute Sync Config
[**GetSourceConfig**](SourcesAPI.md#GetSourceConfig) | **Get** /sources/{id}/connectors/source-config | Gets source config with language translations
[**GetSourceEntitlementRequestConfig**](SourcesAPI.md#GetSourceEntitlementRequestConfig) | **Get** /sources/{id}/entitlement-request-config | Get Source Entitlement Request Configuration
[**GetSourceHealth**](SourcesAPI.md#GetSourceHealth) | **Get** /sources/{sourceId}/source-health | Fetches source health by id
[**GetSourceSchedule**](SourcesAPI.md#GetSourceSchedule) | **Get** /sources/{sourceId}/schedules/{scheduleType} | Get Source Schedule by Type
[**GetSourceSchedules**](SourcesAPI.md#GetSourceSchedules) | **Get** /sources/{sourceId}/schedules | List Schedules on Source
[**GetSourceSchema**](SourcesAPI.md#GetSourceSchema) | **Get** /sources/{sourceId}/schemas/{schemaId} | Get Source Schema by ID
[**GetSourceSchemas**](SourcesAPI.md#GetSourceSchemas) | **Get** /sources/{sourceId}/schemas | List Schemas on Source
[**ImportAccounts**](SourcesAPI.md#ImportAccounts) | **Post** /sources/{id}/load-accounts | Account Aggregation
[**ImportAccountsSchema**](SourcesAPI.md#ImportAccountsSchema) | **Post** /sources/{id}/schemas/accounts | Uploads source accounts schema template
[**ImportConnectorFile**](SourcesAPI.md#ImportConnectorFile) | **Post** /sources/{sourceId}/upload-connector-file | Upload connector file to source
[**ImportEntitlementsSchema**](SourcesAPI.md#ImportEntitlementsSchema) | **Post** /sources/{id}/schemas/entitlements | Uploads source entitlements schema template
[**ImportUncorrelatedAccounts**](SourcesAPI.md#ImportUncorrelatedAccounts) | **Post** /sources/{id}/load-uncorrelated-accounts | Process Uncorrelated Accounts
[**ListProvisioningPolicies**](SourcesAPI.md#ListProvisioningPolicies) | **Get** /sources/{sourceId}/provisioning-policies | Lists ProvisioningPolicies
[**ListSources**](SourcesAPI.md#ListSources) | **Get** /sources | Lists all sources in IdentityNow.
[**PeekResourceObjects**](SourcesAPI.md#PeekResourceObjects) | **Post** /sources/{sourceId}/connector/peek-resource-objects | Peek source connector&#39;s resource objects
[**PingCluster**](SourcesAPI.md#PingCluster) | **Post** /sources/{sourceId}/connector/ping-cluster | Ping cluster for source connector
[**PutCorrelationConfig**](SourcesAPI.md#PutCorrelationConfig) | **Put** /sources/{id}/correlation-config | Update Source Correlation Configuration
[**PutNativeChangeDetectionConfig**](SourcesAPI.md#PutNativeChangeDetectionConfig) | **Put** /sources/{sourceId}/native-change-detection-config | Update Native Change Detection Configuration
[**PutProvisioningPolicy**](SourcesAPI.md#PutProvisioningPolicy) | **Put** /sources/{sourceId}/provisioning-policies/{usageType} | Update Provisioning Policy by UsageType
[**PutSource**](SourcesAPI.md#PutSource) | **Put** /sources/{id} | Update Source (Full)
[**PutSourceAttrSyncConfig**](SourcesAPI.md#PutSourceAttrSyncConfig) | **Put** /sources/{id}/attribute-sync-config | Update Attribute Sync Config
[**PutSourceSchema**](SourcesAPI.md#PutSourceSchema) | **Put** /sources/{sourceId}/schemas/{schemaId} | Update Source Schema (Full)
[**SyncAttributesForSource**](SourcesAPI.md#SyncAttributesForSource) | **Post** /sources/{id}/synchronize-attributes | Synchronize single source attributes.
[**TestSourceConfiguration**](SourcesAPI.md#TestSourceConfiguration) | **Post** /sources/{sourceId}/connector/test-configuration | Test configuration for source connector
[**TestSourceConnection**](SourcesAPI.md#TestSourceConnection) | **Post** /sources/{sourceId}/connector/check-connection | Check connection for source connector.
[**UpdatePasswordPolicyHolders**](SourcesAPI.md#UpdatePasswordPolicyHolders) | **Patch** /sources/{sourceId}/password-policies | Update Password Policy
[**UpdateProvisioningPoliciesInBulk**](SourcesAPI.md#UpdateProvisioningPoliciesInBulk) | **Post** /sources/{sourceId}/provisioning-policies/bulk-update | Bulk Update Provisioning Policies
[**UpdateProvisioningPolicy**](SourcesAPI.md#UpdateProvisioningPolicy) | **Patch** /sources/{sourceId}/provisioning-policies/{usageType} | Partial update of Provisioning Policy
[**UpdateSource**](SourcesAPI.md#UpdateSource) | **Patch** /sources/{id} | Update Source (Partial)
[**UpdateSourceEntitlementRequestConfig**](SourcesAPI.md#UpdateSourceEntitlementRequestConfig) | **Put** /sources/{id}/entitlement-request-config | Update Source Entitlement Request Configuration
[**UpdateSourceSchedule**](SourcesAPI.md#UpdateSourceSchedule) | **Patch** /sources/{sourceId}/schedules/{scheduleType} | Update Source Schedule (Partial)
[**UpdateSourceSchema**](SourcesAPI.md#UpdateSourceSchema) | **Patch** /sources/{sourceId}/schemas/{schemaId} | Update Source Schema (Partial)



## CreateProvisioningPolicy

> ProvisioningPolicyDto CreateProvisioningPolicy(ctx, sourceId).ProvisioningPolicyDto(provisioningPolicyDto).Execute()

Create Provisioning Policy



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
	sourceId := "2c9180835d191a86015d28455b4a2329" // string | The Source id
	provisioningPolicyDto := *openapiclient.NewProvisioningPolicyDto("example provisioning policy for inactive identities") // ProvisioningPolicyDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourcesAPI.CreateProvisioningPolicy(context.Background(), sourceId).ProvisioningPolicyDto(provisioningPolicyDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourcesAPI.CreateProvisioningPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateProvisioningPolicy`: ProvisioningPolicyDto
	fmt.Fprintf(os.Stdout, "Response from `SourcesAPI.CreateProvisioningPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** | The Source id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProvisioningPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **provisioningPolicyDto** | [**ProvisioningPolicyDto**](ProvisioningPolicyDto.md) |  | 

### Return type

[**ProvisioningPolicyDto**](ProvisioningPolicyDto.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSource

> Source CreateSource(ctx).Source(source).ProvisionAsCsv(provisionAsCsv).Execute()

Creates a source in IdentityNow.



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
	source := *openapiclient.NewSource("My Source", *openapiclient.NewSourceOwner(), "active-directory") // Source | 
	provisionAsCsv := false // bool | If this parameter is `true`, it configures the source as a Delimited File (CSV) source. Setting this to `true` will automatically set the `type` of the source to `DelimitedFile`.  You must use this query parameter to create a Delimited File source as you would in the UI.  If you don't set this query parameter and you attempt to set the `type` attribute directly, the request won't correctly generate the source.   (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourcesAPI.CreateSource(context.Background()).Source(source).ProvisionAsCsv(provisionAsCsv).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourcesAPI.CreateSource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSource`: Source
	fmt.Fprintf(os.Stdout, "Response from `SourcesAPI.CreateSource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **source** | [**Source**](Source.md) |  | 
 **provisionAsCsv** | **bool** | If this parameter is &#x60;true&#x60;, it configures the source as a Delimited File (CSV) source. Setting this to &#x60;true&#x60; will automatically set the &#x60;type&#x60; of the source to &#x60;DelimitedFile&#x60;.  You must use this query parameter to create a Delimited File source as you would in the UI.  If you don&#39;t set this query parameter and you attempt to set the &#x60;type&#x60; attribute directly, the request won&#39;t correctly generate the source.   | 

### Return type

[**Source**](Source.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSourceSchedule

> Schedule1 CreateSourceSchedule(ctx, sourceId).Schedule1(schedule1).Execute()

Create Schedule on Source



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
	sourceId := "2c9180835d191a86015d28455b4a2329" // string | Source ID.
	schedule1 := *openapiclient.NewSchedule1("ACCOUNT_AGGREGATION", "0 0 5,13,21 * * ?") // Schedule1 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourcesAPI.CreateSourceSchedule(context.Background(), sourceId).Schedule1(schedule1).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourcesAPI.CreateSourceSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSourceSchedule`: Schedule1
	fmt.Fprintf(os.Stdout, "Response from `SourcesAPI.CreateSourceSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** | Source ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSourceScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **schedule1** | [**Schedule1**](Schedule1.md) |  | 

### Return type

[**Schedule1**](Schedule1.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSourceSchema

> Schema CreateSourceSchema(ctx, sourceId).Schema(schema).Execute()

Create Schema on Source



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
	sourceId := "2c9180835d191a86015d28455b4a2329" // string | Source ID.
	schema := *openapiclient.NewSchema() // Schema | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourcesAPI.CreateSourceSchema(context.Background(), sourceId).Schema(schema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourcesAPI.CreateSourceSchema``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSourceSchema`: Schema
	fmt.Fprintf(os.Stdout, "Response from `SourcesAPI.CreateSourceSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** | Source ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSourceSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **schema** | [**Schema**](Schema.md) |  | 

### Return type

[**Schema**](Schema.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAccountsAsync

> TaskResultDto DeleteAccountsAsync(ctx, id).XSailPointExperimental(xSailPointExperimental).Execute()

Remove All Accounts in a Source



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
	id := "ebbf35756e1140699ce52b233121384a" // string | The source id
	xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourcesAPI.DeleteAccountsAsync(context.Background(), id).XSailPointExperimental(xSailPointExperimental).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourcesAPI.DeleteAccountsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAccountsAsync`: TaskResultDto
	fmt.Fprintf(os.Stdout, "Response from `SourcesAPI.DeleteAccountsAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The source id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccountsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]

### Return type

[**TaskResultDto**](TaskResultDto.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNativeChangeDetectionConfig

> DeleteNativeChangeDetectionConfig(ctx, id).XSailPointExperimental(xSailPointExperimental).Execute()

Delete Native Change Detection Configuration



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
	id := "2c9180835d191a86015d28455b4a2329" // string | The source id
	xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SourcesAPI.DeleteNativeChangeDetectionConfig(context.Background(), id).XSailPointExperimental(xSailPointExperimental).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourcesAPI.DeleteNativeChangeDetectionConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The source id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNativeChangeDetectionConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]

### Return type

 (empty response body)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth), [applicationAuth](../README.md#applicationAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProvisioningPolicy

> DeleteProvisioningPolicy(ctx, sourceId, usageType).Execute()

Delete Provisioning Policy by UsageType



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
	sourceId := "2c9180835d191a86015d28455b4a2329" // string | The Source ID.
	usageType := openapiclient.UsageType("CREATE") // UsageType | The type of provisioning policy usage.  In IdentityNow, a source can support various provisioning operations. For example, when a joiner is added to a source, this may trigger both CREATE and UPDATE provisioning operations.  Each usage type is considered a provisioning policy.  A source can have any number of these provisioning policies defined.  These are the common usage types:  CREATE - This usage type relates to 'Create Account Profile', the provisioning template for the account to be created. For example, this would be used for a joiner on a source.   UPDATE - This usage type relates to 'Update Account Profile', the provisioning template for the 'Update' connector operations. For example, this would be used for an attribute sync on a source. ENABLE - This usage type relates to 'Enable Account Profile', the provisioning template for the account to be enabled. For example, this could be used for a joiner on a source once the joiner's account is created.  DISABLE - This usage type relates to 'Disable Account Profile', the provisioning template for the account to be disabled. For example, this could be used when a leaver is removed temporarily from a source.  You can use these four usage types for all your provisioning policy needs. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SourcesAPI.DeleteProvisioningPolicy(context.Background(), sourceId, usageType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourcesAPI.DeleteProvisioningPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** | The Source ID. | 
**usageType** | [**UsageType**](.md) | The type of provisioning policy usage.  In IdentityNow, a source can support various provisioning operations. For example, when a joiner is added to a source, this may trigger both CREATE and UPDATE provisioning operations.  Each usage type is considered a provisioning policy.  A source can have any number of these provisioning policies defined.  These are the common usage types:  CREATE - This usage type relates to &#39;Create Account Profile&#39;, the provisioning template for the account to be created. For example, this would be used for a joiner on a source.   UPDATE - This usage type relates to &#39;Update Account Profile&#39;, the provisioning template for the &#39;Update&#39; connector operations. For example, this would be used for an attribute sync on a source. ENABLE - This usage type relates to &#39;Enable Account Profile&#39;, the provisioning template for the account to be enabled. For example, this could be used for a joiner on a source once the joiner&#39;s account is created.  DISABLE - This usage type relates to &#39;Disable Account Profile&#39;, the provisioning template for the account to be disabled. For example, this could be used when a leaver is removed temporarily from a source.  You can use these four usage types for all your provisioning policy needs.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProvisioningPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth), [applicationAuth](../README.md#applicationAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSource

> DeleteSource202Response DeleteSource(ctx, id).Execute()

Delete Source by ID



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
	id := "2c9180835d191a86015d28455b4a2329" // string | Source ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourcesAPI.DeleteSource(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourcesAPI.DeleteSource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSource`: DeleteSource202Response
	fmt.Fprintf(os.Stdout, "Response from `SourcesAPI.DeleteSource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Source ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteSource202Response**](DeleteSource202Response.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSourceSchedule

> DeleteSourceSchedule(ctx, sourceId, scheduleType).Execute()

Delete Source Schedule by type.

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
	sourceId := "2c9180835d191a86015d28455b4a2329" // string | The Source id.
	scheduleType := "ACCOUNT_AGGREGATION" // string | The Schedule type.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SourcesAPI.DeleteSourceSchedule(context.Background(), sourceId, scheduleType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourcesAPI.DeleteSourceSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** | The Source id. | 
**scheduleType** | **string** | The Schedule type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSourceScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## DeleteSourceSchema

> DeleteSourceSchema(ctx, sourceId, schemaId).Execute()

Delete Source Schema by ID

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
	sourceId := "2c9180835d191a86015d28455b4a2329" // string | The Source id.
	schemaId := "2c9180835d191a86015d28455b4a2329" // string | The Schema id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SourcesAPI.DeleteSourceSchema(context.Background(), sourceId, schemaId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourcesAPI.DeleteSourceSchema``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** | The Source id. | 
**schemaId** | **string** | The Schema id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSourceSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## GetAccountsSchema

> GetAccountsSchema(ctx, id).Execute()

Downloads source accounts schema template



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
	id := "8c190e6787aa4ed9a90bd9d5344523fb" // string | The Source id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SourcesAPI.GetAccountsSchema(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourcesAPI.GetAccountsSchema``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Source id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountsSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/csv, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCorrelationConfig

> CorrelationConfig GetCorrelationConfig(ctx, id).Execute()

Get Source Correlation Configuration



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
	id := "2c9180835d191a86015d28455b4a2329" // string | The source id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourcesAPI.GetCorrelationConfig(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourcesAPI.GetCorrelationConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCorrelationConfig`: CorrelationConfig
	fmt.Fprintf(os.Stdout, "Response from `SourcesAPI.GetCorrelationConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The source id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCorrelationConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CorrelationConfig**](CorrelationConfig.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEntitlementsSchema

> GetEntitlementsSchema(ctx, id).SchemaName(schemaName).Execute()

Downloads source entitlements schema template



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
	id := "8c190e6787aa4ed9a90bd9d5344523fb" // string | The Source id
	schemaName := "?schemaName=group" // string | Name of entitlement schema (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SourcesAPI.GetEntitlementsSchema(context.Background(), id).SchemaName(schemaName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourcesAPI.GetEntitlementsSchema``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Source id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEntitlementsSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **schemaName** | **string** | Name of entitlement schema | 

### Return type

 (empty response body)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/csv, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNativeChangeDetectionConfig

> NativeChangeDetectionConfig GetNativeChangeDetectionConfig(ctx, id).XSailPointExperimental(xSailPointExperimental).Execute()

Native Change Detection Configuration



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
	id := "2c9180835d191a86015d28455b4a2329" // string | The source id
	xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourcesAPI.GetNativeChangeDetectionConfig(context.Background(), id).XSailPointExperimental(xSailPointExperimental).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourcesAPI.GetNativeChangeDetectionConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNativeChangeDetectionConfig`: NativeChangeDetectionConfig
	fmt.Fprintf(os.Stdout, "Response from `SourcesAPI.GetNativeChangeDetectionConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The source id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNativeChangeDetectionConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]

### Return type

[**NativeChangeDetectionConfig**](NativeChangeDetectionConfig.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProvisioningPolicy

> ProvisioningPolicyDto GetProvisioningPolicy(ctx, sourceId, usageType).Execute()

Get Provisioning Policy by UsageType



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
	sourceId := "2c9180835d191a86015d28455b4a2329" // string | The Source ID.
	usageType := openapiclient.UsageType("CREATE") // UsageType | The type of provisioning policy usage.  In IdentityNow, a source can support various provisioning operations. For example, when a joiner is added to a source, this may trigger both CREATE and UPDATE provisioning operations.  Each usage type is considered a provisioning policy.  A source can have any number of these provisioning policies defined.  These are the common usage types:  CREATE - This usage type relates to 'Create Account Profile', the provisioning template for the account to be created. For example, this would be used for a joiner on a source.   UPDATE - This usage type relates to 'Update Account Profile', the provisioning template for the 'Update' connector operations. For example, this would be used for an attribute sync on a source. ENABLE - This usage type relates to 'Enable Account Profile', the provisioning template for the account to be enabled. For example, this could be used for a joiner on a source once the joiner's account is created.  DISABLE - This usage type relates to 'Disable Account Profile', the provisioning template for the account to be disabled. For example, this could be used when a leaver is removed temporarily from a source.  You can use these four usage types for all your provisioning policy needs. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourcesAPI.GetProvisioningPolicy(context.Background(), sourceId, usageType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourcesAPI.GetProvisioningPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProvisioningPolicy`: ProvisioningPolicyDto
	fmt.Fprintf(os.Stdout, "Response from `SourcesAPI.GetProvisioningPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** | The Source ID. | 
**usageType** | [**UsageType**](.md) | The type of provisioning policy usage.  In IdentityNow, a source can support various provisioning operations. For example, when a joiner is added to a source, this may trigger both CREATE and UPDATE provisioning operations.  Each usage type is considered a provisioning policy.  A source can have any number of these provisioning policies defined.  These are the common usage types:  CREATE - This usage type relates to &#39;Create Account Profile&#39;, the provisioning template for the account to be created. For example, this would be used for a joiner on a source.   UPDATE - This usage type relates to &#39;Update Account Profile&#39;, the provisioning template for the &#39;Update&#39; connector operations. For example, this would be used for an attribute sync on a source. ENABLE - This usage type relates to &#39;Enable Account Profile&#39;, the provisioning template for the account to be enabled. For example, this could be used for a joiner on a source once the joiner&#39;s account is created.  DISABLE - This usage type relates to &#39;Disable Account Profile&#39;, the provisioning template for the account to be disabled. For example, this could be used when a leaver is removed temporarily from a source.  You can use these four usage types for all your provisioning policy needs.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProvisioningPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ProvisioningPolicyDto**](ProvisioningPolicyDto.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth), [applicationAuth](../README.md#applicationAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSource

> Source GetSource(ctx, id).Execute()

Get Source by ID



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
	id := "2c9180835d191a86015d28455b4a2329" // string | Source ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourcesAPI.GetSource(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourcesAPI.GetSource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSource`: Source
	fmt.Fprintf(os.Stdout, "Response from `SourcesAPI.GetSource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Source ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Source**](Source.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSourceAttrSyncConfig

> AttrSyncSourceConfig GetSourceAttrSyncConfig(ctx, id).XSailPointExperimental(xSailPointExperimental).Execute()

Attribute Sync Config



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
	id := "2c9180835d191a86015d28455b4a2329" // string | The source id
	xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourcesAPI.GetSourceAttrSyncConfig(context.Background(), id).XSailPointExperimental(xSailPointExperimental).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourcesAPI.GetSourceAttrSyncConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSourceAttrSyncConfig`: AttrSyncSourceConfig
	fmt.Fprintf(os.Stdout, "Response from `SourcesAPI.GetSourceAttrSyncConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The source id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSourceAttrSyncConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]

### Return type

[**AttrSyncSourceConfig**](AttrSyncSourceConfig.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSourceConfig

> ConnectorDetail GetSourceConfig(ctx, id).XSailPointExperimental(xSailPointExperimental).Locale(locale).Execute()

Gets source config with language translations



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
	id := "id_example" // string | The Source id
	xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")
	locale := "locale_example" // string | The locale to apply to the config. If no viable locale is given, it will default to \"en\" (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourcesAPI.GetSourceConfig(context.Background(), id).XSailPointExperimental(xSailPointExperimental).Locale(locale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourcesAPI.GetSourceConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSourceConfig`: ConnectorDetail
	fmt.Fprintf(os.Stdout, "Response from `SourcesAPI.GetSourceConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Source id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSourceConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **locale** | **string** | The locale to apply to the config. If no viable locale is given, it will default to \&quot;en\&quot; | 

### Return type

[**ConnectorDetail**](ConnectorDetail.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSourceEntitlementRequestConfig

> SourceEntitlementRequestConfig GetSourceEntitlementRequestConfig(ctx).XSailPointExperimental(xSailPointExperimental).Execute()

Get Source Entitlement Request Configuration



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourcesAPI.GetSourceEntitlementRequestConfig(context.Background()).XSailPointExperimental(xSailPointExperimental).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourcesAPI.GetSourceEntitlementRequestConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSourceEntitlementRequestConfig`: SourceEntitlementRequestConfig
	fmt.Fprintf(os.Stdout, "Response from `SourcesAPI.GetSourceEntitlementRequestConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSourceEntitlementRequestConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]

### Return type

[**SourceEntitlementRequestConfig**](SourceEntitlementRequestConfig.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSourceHealth

> SourceHealthDto GetSourceHealth(ctx, sourceId).Execute()

Fetches source health by id



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
	sourceId := "2c9180835d191a86015d28455b4a2329" // string | The Source id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourcesAPI.GetSourceHealth(context.Background(), sourceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourcesAPI.GetSourceHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSourceHealth`: SourceHealthDto
	fmt.Fprintf(os.Stdout, "Response from `SourcesAPI.GetSourceHealth`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** | The Source id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSourceHealthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SourceHealthDto**](SourceHealthDto.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSourceSchedule

> Schedule1 GetSourceSchedule(ctx, sourceId, scheduleType).Execute()

Get Source Schedule by Type



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
	sourceId := "2c9180835d191a86015d28455b4a2329" // string | The Source id.
	scheduleType := "ACCOUNT_AGGREGATION" // string | The Schedule type.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourcesAPI.GetSourceSchedule(context.Background(), sourceId, scheduleType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourcesAPI.GetSourceSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSourceSchedule`: Schedule1
	fmt.Fprintf(os.Stdout, "Response from `SourcesAPI.GetSourceSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** | The Source id. | 
**scheduleType** | **string** | The Schedule type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSourceScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Schedule1**](Schedule1.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSourceSchedules

> []Schedule1 GetSourceSchedules(ctx, sourceId).Execute()

List Schedules on Source



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
	sourceId := "2c9180835d191a86015d28455b4a2329" // string | Source ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourcesAPI.GetSourceSchedules(context.Background(), sourceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourcesAPI.GetSourceSchedules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSourceSchedules`: []Schedule1
	fmt.Fprintf(os.Stdout, "Response from `SourcesAPI.GetSourceSchedules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** | Source ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSourceSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Schedule1**](Schedule1.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSourceSchema

> Schema GetSourceSchema(ctx, sourceId, schemaId).Execute()

Get Source Schema by ID



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
	sourceId := "2c9180835d191a86015d28455b4a2329" // string | The Source id.
	schemaId := "2c9180835d191a86015d28455b4a2329" // string | The Schema id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourcesAPI.GetSourceSchema(context.Background(), sourceId, schemaId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourcesAPI.GetSourceSchema``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSourceSchema`: Schema
	fmt.Fprintf(os.Stdout, "Response from `SourcesAPI.GetSourceSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** | The Source id. | 
**schemaId** | **string** | The Schema id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSourceSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Schema**](Schema.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSourceSchemas

> []Schema GetSourceSchemas(ctx, sourceId).IncludeTypes(includeTypes).IncludeNames(includeNames).Execute()

List Schemas on Source



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
	sourceId := "2c9180835d191a86015d28455b4a2329" // string | Source ID.
	includeTypes := "group" // string | If set to 'group', then the account schema is filtered and only group schemas are returned. Only a value of 'group' is recognized presently.  Note: The API will check whether include-types is group or not, if not, it will list schemas based on include-names, if include-names is not provided, it will list all schemas. (optional)
	includeNames := "account" // string | A comma-separated list of schema names to filter result. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourcesAPI.GetSourceSchemas(context.Background(), sourceId).IncludeTypes(includeTypes).IncludeNames(includeNames).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourcesAPI.GetSourceSchemas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSourceSchemas`: []Schema
	fmt.Fprintf(os.Stdout, "Response from `SourcesAPI.GetSourceSchemas`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** | Source ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSourceSchemasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeTypes** | **string** | If set to &#39;group&#39;, then the account schema is filtered and only group schemas are returned. Only a value of &#39;group&#39; is recognized presently.  Note: The API will check whether include-types is group or not, if not, it will list schemas based on include-names, if include-names is not provided, it will list all schemas. | 
 **includeNames** | **string** | A comma-separated list of schema names to filter result. | 

### Return type

[**[]Schema**](Schema.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportAccounts

> LoadAccountsTask ImportAccounts(ctx, id).XSailPointExperimental(xSailPointExperimental).File(file).DisableOptimization(disableOptimization).Execute()

Account Aggregation



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
	id := "ef38f94347e94562b5bb8424a56397d8" // string | Source Id
	xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")
	file := os.NewFile(1234, "some_file") // *os.File | The CSV file containing the source accounts to aggregate. (optional)
	disableOptimization := "disableOptimization_example" // string | Use this flag to reprocess every account whether or not the data has changed. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourcesAPI.ImportAccounts(context.Background(), id).XSailPointExperimental(xSailPointExperimental).File(file).DisableOptimization(disableOptimization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourcesAPI.ImportAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportAccounts`: LoadAccountsTask
	fmt.Fprintf(os.Stdout, "Response from `SourcesAPI.ImportAccounts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Source Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **file** | ***os.File** | The CSV file containing the source accounts to aggregate. | 
 **disableOptimization** | **string** | Use this flag to reprocess every account whether or not the data has changed. | 

### Return type

[**LoadAccountsTask**](LoadAccountsTask.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportAccountsSchema

> Schema ImportAccountsSchema(ctx, id).File(file).Execute()

Uploads source accounts schema template



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
	id := "8c190e6787aa4ed9a90bd9d5344523fb" // string | The Source id
	file := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourcesAPI.ImportAccountsSchema(context.Background(), id).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourcesAPI.ImportAccountsSchema``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportAccountsSchema`: Schema
	fmt.Fprintf(os.Stdout, "Response from `SourcesAPI.ImportAccountsSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Source id | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportAccountsSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | ***os.File** |  | 

### Return type

[**Schema**](Schema.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportConnectorFile

> Source ImportConnectorFile(ctx, sourceId).File(file).Execute()

Upload connector file to source



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
	sourceId := "2c9180835d191a86015d28455b4a2329" // string | The Source id.
	file := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourcesAPI.ImportConnectorFile(context.Background(), sourceId).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourcesAPI.ImportConnectorFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportConnectorFile`: Source
	fmt.Fprintf(os.Stdout, "Response from `SourcesAPI.ImportConnectorFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** | The Source id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportConnectorFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | ***os.File** |  | 

### Return type

[**Source**](Source.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportEntitlementsSchema

> Schema ImportEntitlementsSchema(ctx, id).SchemaName(schemaName).File(file).Execute()

Uploads source entitlements schema template



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
	id := "8c190e6787aa4ed9a90bd9d5344523fb" // string | The Source id
	schemaName := "?schemaName=group" // string | Name of entitlement schema (optional)
	file := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourcesAPI.ImportEntitlementsSchema(context.Background(), id).SchemaName(schemaName).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourcesAPI.ImportEntitlementsSchema``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportEntitlementsSchema`: Schema
	fmt.Fprintf(os.Stdout, "Response from `SourcesAPI.ImportEntitlementsSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Source id | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportEntitlementsSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **schemaName** | **string** | Name of entitlement schema | 
 **file** | ***os.File** |  | 

### Return type

[**Schema**](Schema.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportUncorrelatedAccounts

> LoadUncorrelatedAccountsTask ImportUncorrelatedAccounts(ctx, id).XSailPointExperimental(xSailPointExperimental).File(file).Execute()

Process Uncorrelated Accounts



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
	id := "75dbec1ebe154d5785da27b95e1dd5d7" // string | Source Id
	xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")
	file := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourcesAPI.ImportUncorrelatedAccounts(context.Background(), id).XSailPointExperimental(xSailPointExperimental).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourcesAPI.ImportUncorrelatedAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportUncorrelatedAccounts`: LoadUncorrelatedAccountsTask
	fmt.Fprintf(os.Stdout, "Response from `SourcesAPI.ImportUncorrelatedAccounts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Source Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportUncorrelatedAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **file** | ***os.File** |  | 

### Return type

[**LoadUncorrelatedAccountsTask**](LoadUncorrelatedAccountsTask.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProvisioningPolicies

> []ProvisioningPolicyDto ListProvisioningPolicies(ctx, sourceId).Execute()

Lists ProvisioningPolicies



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
	sourceId := "2c9180835d191a86015d28455b4a2329" // string | The Source id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourcesAPI.ListProvisioningPolicies(context.Background(), sourceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourcesAPI.ListProvisioningPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListProvisioningPolicies`: []ProvisioningPolicyDto
	fmt.Fprintf(os.Stdout, "Response from `SourcesAPI.ListProvisioningPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** | The Source id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListProvisioningPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ProvisioningPolicyDto**](ProvisioningPolicyDto.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth), [applicationAuth](../README.md#applicationAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSources

> []Source ListSources(ctx).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).ForSubadmin(forSubadmin).IncludeIDNSource(includeIDNSource).Execute()

Lists all sources in IdentityNow.



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
	limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
	offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
	count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
	filters := "name eq "Employees"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq, in, ge, gt, le, lt, ne, isnull, sw*  **name**: *co, eq, in, sw, ge, gt, ne, isnull*  **type**: *eq, in, ge, gt, ne, isnull, sw*  **owner.id**: *eq, in, ge, gt, le, lt, ne, isnull, sw*  **features**: *ca, co*  **created**: *eq, ge, gt, in, le, lt, ne, isnull, sw*  **modified**: *eq, ge, gt, in, le, lt, ne, isnull, sw*  **managementWorkgroup.id**: *eq, ge, gt, in, le, lt, ne, isnull, sw*  **description**: *eq, sw*  **authoritative**: *eq, ne, isnull*  **healthy**: *isnull*  **status**: *eq, in, ge, gt, le, lt, ne, isnull, sw*  **connectionType**: *eq, ge, gt, in, le, lt, ne, isnull, sw*  **connectorName**: *eq, ge, gt, in, ne, isnull, sw*  **category**: *co, eq, ge, gt, in, le, lt, ne, sw* (optional)
	sorters := "name" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **type, created, modified, name, owner.name, healthy, status, id, description, owner.id, accountCorrelationConfig.id, accountCorrelationConfig.name, managerCorrelationRule.type, managerCorrelationRule.id, managerCorrelationRule.name, authoritative, managementWorkgroup.id, connectorName, connectionType** (optional)
	forSubadmin := "name" // string | Filter the returned list of sources for the identity specified by the parameter, which is the id of an identity with the role SOURCE_SUBADMIN. By convention, the value **me** indicates the identity id of the current user. Subadmins may only view Sources which they are able to administer; all other Sources will be filtered out when this parameter is set. If the current user is a SOURCE_SUBADMIN but fails to pass a valid value for this parameter, a 403 Forbidden is returned. (optional)
	includeIDNSource := true // bool | Include the IdentityNow source in the response. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourcesAPI.ListSources(context.Background()).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).ForSubadmin(forSubadmin).IncludeIDNSource(includeIDNSource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourcesAPI.ListSources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSources`: []Source
	fmt.Fprintf(os.Stdout, "Response from `SourcesAPI.ListSources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq, in, ge, gt, le, lt, ne, isnull, sw*  **name**: *co, eq, in, sw, ge, gt, ne, isnull*  **type**: *eq, in, ge, gt, ne, isnull, sw*  **owner.id**: *eq, in, ge, gt, le, lt, ne, isnull, sw*  **features**: *ca, co*  **created**: *eq, ge, gt, in, le, lt, ne, isnull, sw*  **modified**: *eq, ge, gt, in, le, lt, ne, isnull, sw*  **managementWorkgroup.id**: *eq, ge, gt, in, le, lt, ne, isnull, sw*  **description**: *eq, sw*  **authoritative**: *eq, ne, isnull*  **healthy**: *isnull*  **status**: *eq, in, ge, gt, le, lt, ne, isnull, sw*  **connectionType**: *eq, ge, gt, in, le, lt, ne, isnull, sw*  **connectorName**: *eq, ge, gt, in, ne, isnull, sw*  **category**: *co, eq, ge, gt, in, le, lt, ne, sw* | 
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **type, created, modified, name, owner.name, healthy, status, id, description, owner.id, accountCorrelationConfig.id, accountCorrelationConfig.name, managerCorrelationRule.type, managerCorrelationRule.id, managerCorrelationRule.name, authoritative, managementWorkgroup.id, connectorName, connectionType** | 
 **forSubadmin** | **string** | Filter the returned list of sources for the identity specified by the parameter, which is the id of an identity with the role SOURCE_SUBADMIN. By convention, the value **me** indicates the identity id of the current user. Subadmins may only view Sources which they are able to administer; all other Sources will be filtered out when this parameter is set. If the current user is a SOURCE_SUBADMIN but fails to pass a valid value for this parameter, a 403 Forbidden is returned. | 
 **includeIDNSource** | **bool** | Include the IdentityNow source in the response. | [default to false]

### Return type

[**[]Source**](Source.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth), [applicationAuth](../README.md#applicationAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PeekResourceObjects

> ResourceObjectsResponse PeekResourceObjects(ctx, sourceId).XSailPointExperimental(xSailPointExperimental).ResourceObjectsRequest(resourceObjectsRequest).Execute()

Peek source connector's resource objects



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
	sourceId := "cef3ee201db947c5912551015ba0c679" // string | The ID of the Source
	xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")
	resourceObjectsRequest := *openapiclient.NewResourceObjectsRequest() // ResourceObjectsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourcesAPI.PeekResourceObjects(context.Background(), sourceId).XSailPointExperimental(xSailPointExperimental).ResourceObjectsRequest(resourceObjectsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourcesAPI.PeekResourceObjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PeekResourceObjects`: ResourceObjectsResponse
	fmt.Fprintf(os.Stdout, "Response from `SourcesAPI.PeekResourceObjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** | The ID of the Source | 

### Other Parameters

Other parameters are passed through a pointer to a apiPeekResourceObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **resourceObjectsRequest** | [**ResourceObjectsRequest**](ResourceObjectsRequest.md) |  | 

### Return type

[**ResourceObjectsResponse**](ResourceObjectsResponse.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PingCluster

> StatusResponse PingCluster(ctx, sourceId).XSailPointExperimental(xSailPointExperimental).Execute()

Ping cluster for source connector



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
	sourceId := "cef3ee201db947c5912551015ba0c679" // string | The ID of the Source
	xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourcesAPI.PingCluster(context.Background(), sourceId).XSailPointExperimental(xSailPointExperimental).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourcesAPI.PingCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PingCluster`: StatusResponse
	fmt.Fprintf(os.Stdout, "Response from `SourcesAPI.PingCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** | The ID of the Source | 

### Other Parameters

Other parameters are passed through a pointer to a apiPingClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]

### Return type

[**StatusResponse**](StatusResponse.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCorrelationConfig

> CorrelationConfig PutCorrelationConfig(ctx, id).CorrelationConfig(correlationConfig).Execute()

Update Source Correlation Configuration



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
	id := "2c9180835d191a86015d28455b4a2329" // string | The source id
	correlationConfig := *openapiclient.NewCorrelationConfig() // CorrelationConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourcesAPI.PutCorrelationConfig(context.Background(), id).CorrelationConfig(correlationConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourcesAPI.PutCorrelationConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutCorrelationConfig`: CorrelationConfig
	fmt.Fprintf(os.Stdout, "Response from `SourcesAPI.PutCorrelationConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The source id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCorrelationConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **correlationConfig** | [**CorrelationConfig**](CorrelationConfig.md) |  | 

### Return type

[**CorrelationConfig**](CorrelationConfig.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutNativeChangeDetectionConfig

> NativeChangeDetectionConfig PutNativeChangeDetectionConfig(ctx, id).XSailPointExperimental(xSailPointExperimental).NativeChangeDetectionConfig(nativeChangeDetectionConfig).Execute()

Update Native Change Detection Configuration



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
	id := "2c9180835d191a86015d28455b4a2329" // string | The source id
	xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")
	nativeChangeDetectionConfig := *openapiclient.NewNativeChangeDetectionConfig() // NativeChangeDetectionConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourcesAPI.PutNativeChangeDetectionConfig(context.Background(), id).XSailPointExperimental(xSailPointExperimental).NativeChangeDetectionConfig(nativeChangeDetectionConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourcesAPI.PutNativeChangeDetectionConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutNativeChangeDetectionConfig`: NativeChangeDetectionConfig
	fmt.Fprintf(os.Stdout, "Response from `SourcesAPI.PutNativeChangeDetectionConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The source id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutNativeChangeDetectionConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **nativeChangeDetectionConfig** | [**NativeChangeDetectionConfig**](NativeChangeDetectionConfig.md) |  | 

### Return type

[**NativeChangeDetectionConfig**](NativeChangeDetectionConfig.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutProvisioningPolicy

> ProvisioningPolicyDto PutProvisioningPolicy(ctx, sourceId, usageType).ProvisioningPolicyDto(provisioningPolicyDto).Execute()

Update Provisioning Policy by UsageType



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
	sourceId := "2c9180835d191a86015d28455b4a2329" // string | The Source ID.
	usageType := openapiclient.UsageType("CREATE") // UsageType | The type of provisioning policy usage.  In IdentityNow, a source can support various provisioning operations. For example, when a joiner is added to a source, this may trigger both CREATE and UPDATE provisioning operations.  Each usage type is considered a provisioning policy.  A source can have any number of these provisioning policies defined.  These are the common usage types:  CREATE - This usage type relates to 'Create Account Profile', the provisioning template for the account to be created. For example, this would be used for a joiner on a source.   UPDATE - This usage type relates to 'Update Account Profile', the provisioning template for the 'Update' connector operations. For example, this would be used for an attribute sync on a source. ENABLE - This usage type relates to 'Enable Account Profile', the provisioning template for the account to be enabled. For example, this could be used for a joiner on a source once the joiner's account is created.  DISABLE - This usage type relates to 'Disable Account Profile', the provisioning template for the account to be disabled. For example, this could be used when a leaver is removed temporarily from a source.  You can use these four usage types for all your provisioning policy needs. 
	provisioningPolicyDto := *openapiclient.NewProvisioningPolicyDto("example provisioning policy for inactive identities") // ProvisioningPolicyDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourcesAPI.PutProvisioningPolicy(context.Background(), sourceId, usageType).ProvisioningPolicyDto(provisioningPolicyDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourcesAPI.PutProvisioningPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutProvisioningPolicy`: ProvisioningPolicyDto
	fmt.Fprintf(os.Stdout, "Response from `SourcesAPI.PutProvisioningPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** | The Source ID. | 
**usageType** | [**UsageType**](.md) | The type of provisioning policy usage.  In IdentityNow, a source can support various provisioning operations. For example, when a joiner is added to a source, this may trigger both CREATE and UPDATE provisioning operations.  Each usage type is considered a provisioning policy.  A source can have any number of these provisioning policies defined.  These are the common usage types:  CREATE - This usage type relates to &#39;Create Account Profile&#39;, the provisioning template for the account to be created. For example, this would be used for a joiner on a source.   UPDATE - This usage type relates to &#39;Update Account Profile&#39;, the provisioning template for the &#39;Update&#39; connector operations. For example, this would be used for an attribute sync on a source. ENABLE - This usage type relates to &#39;Enable Account Profile&#39;, the provisioning template for the account to be enabled. For example, this could be used for a joiner on a source once the joiner&#39;s account is created.  DISABLE - This usage type relates to &#39;Disable Account Profile&#39;, the provisioning template for the account to be disabled. For example, this could be used when a leaver is removed temporarily from a source.  You can use these four usage types for all your provisioning policy needs.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutProvisioningPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **provisioningPolicyDto** | [**ProvisioningPolicyDto**](ProvisioningPolicyDto.md) |  | 

### Return type

[**ProvisioningPolicyDto**](ProvisioningPolicyDto.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth), [applicationAuth](../README.md#applicationAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSource

> Source PutSource(ctx, id).Source(source).Execute()

Update Source (Full)



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
	id := "2c9180835d191a86015d28455b4a2329" // string | Source ID.
	source := *openapiclient.NewSource("My Source", *openapiclient.NewSourceOwner(), "active-directory") // Source | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourcesAPI.PutSource(context.Background(), id).Source(source).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourcesAPI.PutSource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutSource`: Source
	fmt.Fprintf(os.Stdout, "Response from `SourcesAPI.PutSource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Source ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **source** | [**Source**](Source.md) |  | 

### Return type

[**Source**](Source.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSourceAttrSyncConfig

> AttrSyncSourceConfig PutSourceAttrSyncConfig(ctx, id).XSailPointExperimental(xSailPointExperimental).AttrSyncSourceConfig(attrSyncSourceConfig).Execute()

Update Attribute Sync Config



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
	id := "2c9180835d191a86015d28455b4a2329" // string | The source id
	xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")
	attrSyncSourceConfig := *openapiclient.NewAttrSyncSourceConfig(*openapiclient.NewAttrSyncSource(), []openapiclient.AttrSyncSourceAttributeConfig{*openapiclient.NewAttrSyncSourceAttributeConfig("email", "Email", true, "mail")}) // AttrSyncSourceConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourcesAPI.PutSourceAttrSyncConfig(context.Background(), id).XSailPointExperimental(xSailPointExperimental).AttrSyncSourceConfig(attrSyncSourceConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourcesAPI.PutSourceAttrSyncConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutSourceAttrSyncConfig`: AttrSyncSourceConfig
	fmt.Fprintf(os.Stdout, "Response from `SourcesAPI.PutSourceAttrSyncConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The source id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSourceAttrSyncConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **attrSyncSourceConfig** | [**AttrSyncSourceConfig**](AttrSyncSourceConfig.md) |  | 

### Return type

[**AttrSyncSourceConfig**](AttrSyncSourceConfig.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSourceSchema

> Schema PutSourceSchema(ctx, sourceId, schemaId).Schema(schema).Execute()

Update Source Schema (Full)



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
	sourceId := "2c9180835d191a86015d28455b4a2329" // string | The Source id.
	schemaId := "2c9180835d191a86015d28455b4a2329" // string | The Schema id.
	schema := *openapiclient.NewSchema() // Schema | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourcesAPI.PutSourceSchema(context.Background(), sourceId, schemaId).Schema(schema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourcesAPI.PutSourceSchema``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutSourceSchema`: Schema
	fmt.Fprintf(os.Stdout, "Response from `SourcesAPI.PutSourceSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** | The Source id. | 
**schemaId** | **string** | The Schema id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSourceSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **schema** | [**Schema**](Schema.md) |  | 

### Return type

[**Schema**](Schema.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncAttributesForSource

> SourceSyncJob SyncAttributesForSource(ctx, id).XSailPointExperimental(xSailPointExperimental).Execute()

Synchronize single source attributes.



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
	id := "id_example" // string | The Source id
	xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourcesAPI.SyncAttributesForSource(context.Background(), id).XSailPointExperimental(xSailPointExperimental).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourcesAPI.SyncAttributesForSource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SyncAttributesForSource`: SourceSyncJob
	fmt.Fprintf(os.Stdout, "Response from `SourcesAPI.SyncAttributesForSource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Source id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSyncAttributesForSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]

### Return type

[**SourceSyncJob**](SourceSyncJob.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestSourceConfiguration

> StatusResponse TestSourceConfiguration(ctx, sourceId).XSailPointExperimental(xSailPointExperimental).Execute()

Test configuration for source connector



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
	sourceId := "cef3ee201db947c5912551015ba0c679" // string | The ID of the Source
	xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourcesAPI.TestSourceConfiguration(context.Background(), sourceId).XSailPointExperimental(xSailPointExperimental).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourcesAPI.TestSourceConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestSourceConfiguration`: StatusResponse
	fmt.Fprintf(os.Stdout, "Response from `SourcesAPI.TestSourceConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** | The ID of the Source | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestSourceConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]

### Return type

[**StatusResponse**](StatusResponse.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestSourceConnection

> StatusResponse TestSourceConnection(ctx, sourceId).XSailPointExperimental(xSailPointExperimental).Execute()

Check connection for source connector.



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
	sourceId := "cef3ee201db947c5912551015ba0c679" // string | The ID of the Source.
	xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourcesAPI.TestSourceConnection(context.Background(), sourceId).XSailPointExperimental(xSailPointExperimental).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourcesAPI.TestSourceConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestSourceConnection`: StatusResponse
	fmt.Fprintf(os.Stdout, "Response from `SourcesAPI.TestSourceConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** | The ID of the Source. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestSourceConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]

### Return type

[**StatusResponse**](StatusResponse.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePasswordPolicyHolders

> []PasswordPolicyHoldersDtoInner UpdatePasswordPolicyHolders(ctx, sourceId).PasswordPolicyHoldersDtoInner(passwordPolicyHoldersDtoInner).Execute()

Update Password Policy



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
	sourceId := "8c190e6787aa4ed9a90bd9d5344523fb" // string | The Source id
	passwordPolicyHoldersDtoInner := []openapiclient.PasswordPolicyHoldersDtoInner{*openapiclient.NewPasswordPolicyHoldersDtoInner()} // []PasswordPolicyHoldersDtoInner | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourcesAPI.UpdatePasswordPolicyHolders(context.Background(), sourceId).PasswordPolicyHoldersDtoInner(passwordPolicyHoldersDtoInner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourcesAPI.UpdatePasswordPolicyHolders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePasswordPolicyHolders`: []PasswordPolicyHoldersDtoInner
	fmt.Fprintf(os.Stdout, "Response from `SourcesAPI.UpdatePasswordPolicyHolders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** | The Source id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePasswordPolicyHoldersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **passwordPolicyHoldersDtoInner** | [**[]PasswordPolicyHoldersDtoInner**](PasswordPolicyHoldersDtoInner.md) |  | 

### Return type

[**[]PasswordPolicyHoldersDtoInner**](PasswordPolicyHoldersDtoInner.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProvisioningPoliciesInBulk

> []ProvisioningPolicyDto UpdateProvisioningPoliciesInBulk(ctx, sourceId).ProvisioningPolicyDto(provisioningPolicyDto).Execute()

Bulk Update Provisioning Policies



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
	sourceId := "2c9180835d191a86015d28455b4a2329" // string | The Source id.
	provisioningPolicyDto := []openapiclient.ProvisioningPolicyDto{*openapiclient.NewProvisioningPolicyDto("example provisioning policy for inactive identities")} // []ProvisioningPolicyDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourcesAPI.UpdateProvisioningPoliciesInBulk(context.Background(), sourceId).ProvisioningPolicyDto(provisioningPolicyDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourcesAPI.UpdateProvisioningPoliciesInBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateProvisioningPoliciesInBulk`: []ProvisioningPolicyDto
	fmt.Fprintf(os.Stdout, "Response from `SourcesAPI.UpdateProvisioningPoliciesInBulk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** | The Source id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProvisioningPoliciesInBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **provisioningPolicyDto** | [**[]ProvisioningPolicyDto**](ProvisioningPolicyDto.md) |  | 

### Return type

[**[]ProvisioningPolicyDto**](ProvisioningPolicyDto.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth), [applicationAuth](../README.md#applicationAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProvisioningPolicy

> ProvisioningPolicyDto UpdateProvisioningPolicy(ctx, sourceId, usageType).JsonPatchOperation(jsonPatchOperation).Execute()

Partial update of Provisioning Policy



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
	sourceId := "2c9180835d191a86015d28455b4a2329" // string | The Source id.
	usageType := openapiclient.UsageType("CREATE") // UsageType | The type of provisioning policy usage.  In IdentityNow, a source can support various provisioning operations. For example, when a joiner is added to a source, this may trigger both CREATE and UPDATE provisioning operations.  Each usage type is considered a provisioning policy.  A source can have any number of these provisioning policies defined.  These are the common usage types:  CREATE - This usage type relates to 'Create Account Profile', the provisioning template for the account to be created. For example, this would be used for a joiner on a source.   UPDATE - This usage type relates to 'Update Account Profile', the provisioning template for the 'Update' connector operations. For example, this would be used for an attribute sync on a source. ENABLE - This usage type relates to 'Enable Account Profile', the provisioning template for the account to be enabled. For example, this could be used for a joiner on a source once the joiner's account is created.  DISABLE - This usage type relates to 'Disable Account Profile', the provisioning template for the account to be disabled. For example, this could be used when a leaver is removed temporarily from a source.  You can use these four usage types for all your provisioning policy needs. 
	jsonPatchOperation := []openapiclient.JsonPatchOperation{*openapiclient.NewJsonPatchOperation("replace", "/description")} // []JsonPatchOperation | The JSONPatch payload used to update the schema.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourcesAPI.UpdateProvisioningPolicy(context.Background(), sourceId, usageType).JsonPatchOperation(jsonPatchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourcesAPI.UpdateProvisioningPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateProvisioningPolicy`: ProvisioningPolicyDto
	fmt.Fprintf(os.Stdout, "Response from `SourcesAPI.UpdateProvisioningPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** | The Source id. | 
**usageType** | [**UsageType**](.md) | The type of provisioning policy usage.  In IdentityNow, a source can support various provisioning operations. For example, when a joiner is added to a source, this may trigger both CREATE and UPDATE provisioning operations.  Each usage type is considered a provisioning policy.  A source can have any number of these provisioning policies defined.  These are the common usage types:  CREATE - This usage type relates to &#39;Create Account Profile&#39;, the provisioning template for the account to be created. For example, this would be used for a joiner on a source.   UPDATE - This usage type relates to &#39;Update Account Profile&#39;, the provisioning template for the &#39;Update&#39; connector operations. For example, this would be used for an attribute sync on a source. ENABLE - This usage type relates to &#39;Enable Account Profile&#39;, the provisioning template for the account to be enabled. For example, this could be used for a joiner on a source once the joiner&#39;s account is created.  DISABLE - This usage type relates to &#39;Disable Account Profile&#39;, the provisioning template for the account to be disabled. For example, this could be used when a leaver is removed temporarily from a source.  You can use these four usage types for all your provisioning policy needs.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProvisioningPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **jsonPatchOperation** | [**[]JsonPatchOperation**](JsonPatchOperation.md) | The JSONPatch payload used to update the schema. | 

### Return type

[**ProvisioningPolicyDto**](ProvisioningPolicyDto.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth), [applicationAuth](../README.md#applicationAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSource

> Source UpdateSource(ctx, id).JsonPatchOperation(jsonPatchOperation).Execute()

Update Source (Partial)



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
	id := "2c9180835d191a86015d28455b4a2329" // string | Source ID.
	jsonPatchOperation := []openapiclient.JsonPatchOperation{*openapiclient.NewJsonPatchOperation("replace", "/description")} // []JsonPatchOperation | A list of account update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard. Any password changes are submitted as plain-text and encrypted upon receipt in Identity Security Cloud (ISC).

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourcesAPI.UpdateSource(context.Background(), id).JsonPatchOperation(jsonPatchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourcesAPI.UpdateSource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSource`: Source
	fmt.Fprintf(os.Stdout, "Response from `SourcesAPI.UpdateSource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Source ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jsonPatchOperation** | [**[]JsonPatchOperation**](JsonPatchOperation.md) | A list of account update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard. Any password changes are submitted as plain-text and encrypted upon receipt in Identity Security Cloud (ISC). | 

### Return type

[**Source**](Source.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSourceEntitlementRequestConfig

> SourceEntitlementRequestConfig UpdateSourceEntitlementRequestConfig(ctx).XSailPointExperimental(xSailPointExperimental).SourceEntitlementRequestConfig(sourceEntitlementRequestConfig).Execute()

Update Source Entitlement Request Configuration



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
	sourceEntitlementRequestConfig := *openapiclient.NewSourceEntitlementRequestConfig() // SourceEntitlementRequestConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourcesAPI.UpdateSourceEntitlementRequestConfig(context.Background()).XSailPointExperimental(xSailPointExperimental).SourceEntitlementRequestConfig(sourceEntitlementRequestConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourcesAPI.UpdateSourceEntitlementRequestConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSourceEntitlementRequestConfig`: SourceEntitlementRequestConfig
	fmt.Fprintf(os.Stdout, "Response from `SourcesAPI.UpdateSourceEntitlementRequestConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSourceEntitlementRequestConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **sourceEntitlementRequestConfig** | [**SourceEntitlementRequestConfig**](SourceEntitlementRequestConfig.md) |  | 

### Return type

[**SourceEntitlementRequestConfig**](SourceEntitlementRequestConfig.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSourceSchedule

> Schedule1 UpdateSourceSchedule(ctx, sourceId, scheduleType).JsonPatchOperation(jsonPatchOperation).Execute()

Update Source Schedule (Partial)



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
	sourceId := "2c9180835d191a86015d28455b4a2329" // string | The Source id.
	scheduleType := "ACCOUNT_AGGREGATION" // string | The Schedule type.
	jsonPatchOperation := []openapiclient.JsonPatchOperation{*openapiclient.NewJsonPatchOperation("replace", "/description")} // []JsonPatchOperation | The JSONPatch payload used to update the schedule.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourcesAPI.UpdateSourceSchedule(context.Background(), sourceId, scheduleType).JsonPatchOperation(jsonPatchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourcesAPI.UpdateSourceSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSourceSchedule`: Schedule1
	fmt.Fprintf(os.Stdout, "Response from `SourcesAPI.UpdateSourceSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** | The Source id. | 
**scheduleType** | **string** | The Schedule type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSourceScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **jsonPatchOperation** | [**[]JsonPatchOperation**](JsonPatchOperation.md) | The JSONPatch payload used to update the schedule. | 

### Return type

[**Schedule1**](Schedule1.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSourceSchema

> Schema UpdateSourceSchema(ctx, sourceId, schemaId).JsonPatchOperation(jsonPatchOperation).Execute()

Update Source Schema (Partial)



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
	sourceId := "2c9180835d191a86015d28455b4a2329" // string | The Source id.
	schemaId := "2c9180835d191a86015d28455b4a2329" // string | The Schema id.
	jsonPatchOperation := []openapiclient.JsonPatchOperation{*openapiclient.NewJsonPatchOperation("replace", "/description")} // []JsonPatchOperation | The JSONPatch payload used to update the schema.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourcesAPI.UpdateSourceSchema(context.Background(), sourceId, schemaId).JsonPatchOperation(jsonPatchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourcesAPI.UpdateSourceSchema``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSourceSchema`: Schema
	fmt.Fprintf(os.Stdout, "Response from `SourcesAPI.UpdateSourceSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** | The Source id. | 
**schemaId** | **string** | The Schema id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSourceSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **jsonPatchOperation** | [**[]JsonPatchOperation**](JsonPatchOperation.md) | The JSONPatch payload used to update the schema. | 

### Return type

[**Schema**](Schema.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

