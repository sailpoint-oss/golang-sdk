# \SourcesAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProvisioningPolicy**](SourcesAPI.md#CreateProvisioningPolicy) | **Post** /sources/{sourceId}/provisioning-policies | Create Provisioning Policy
[**CreateSource**](SourcesAPI.md#CreateSource) | **Post** /sources | Creates a source in IdentityNow.
[**CreateSourceSchema**](SourcesAPI.md#CreateSourceSchema) | **Post** /sources/{sourceId}/schemas | Create Schema on a Source
[**DeleteProvisioningPolicy**](SourcesAPI.md#DeleteProvisioningPolicy) | **Delete** /sources/{sourceId}/provisioning-policies/{usageType} | Delete Provisioning Policy by UsageType
[**DeleteSource**](SourcesAPI.md#DeleteSource) | **Delete** /sources/{id} | Delete Source by ID
[**DeleteSourceSchema**](SourcesAPI.md#DeleteSourceSchema) | **Delete** /sources/{sourceId}/schemas/{schemaId} | Delete Source Schema by ID
[**GetAccountsSchema**](SourcesAPI.md#GetAccountsSchema) | **Get** /sources/{id}/schemas/accounts | Downloads source accounts schema template
[**GetEntitlementsSchema**](SourcesAPI.md#GetEntitlementsSchema) | **Get** /sources/{id}/schemas/entitlements | Downloads source entitlements schema template
[**GetProvisioningPolicy**](SourcesAPI.md#GetProvisioningPolicy) | **Get** /sources/{sourceId}/provisioning-policies/{usageType} | Get Provisioning Policy by UsageType
[**GetSource**](SourcesAPI.md#GetSource) | **Get** /sources/{id} | Get Source by ID
[**GetSourceHealth**](SourcesAPI.md#GetSourceHealth) | **Get** /sources/{sourceId}/source-health | Fetches source health by id
[**GetSourceSchema**](SourcesAPI.md#GetSourceSchema) | **Get** /sources/{sourceId}/schemas/{schemaId} | Get Source Schema by ID
[**ImportAccountsSchema**](SourcesAPI.md#ImportAccountsSchema) | **Post** /sources/{id}/schemas/accounts | Uploads source accounts schema template
[**ImportConnectorFile**](SourcesAPI.md#ImportConnectorFile) | **Post** /sources/{sourceId}/upload-connector-file | Upload connector file to source
[**ImportEntitlementsSchema**](SourcesAPI.md#ImportEntitlementsSchema) | **Post** /sources/{id}/schemas/entitlements | Uploads source entitlements schema template
[**ListProvisioningPolicies**](SourcesAPI.md#ListProvisioningPolicies) | **Get** /sources/{sourceId}/provisioning-policies | Lists ProvisioningPolicies
[**ListSourceSchemas**](SourcesAPI.md#ListSourceSchemas) | **Get** /sources/{sourceId}/schemas | List Schemas on a Source
[**ListSources**](SourcesAPI.md#ListSources) | **Get** /sources | Lists all sources in IdentityNow.
[**PutProvisioningPolicy**](SourcesAPI.md#PutProvisioningPolicy) | **Put** /sources/{sourceId}/provisioning-policies/{usageType} | Update Provisioning Policy by UsageType
[**PutSource**](SourcesAPI.md#PutSource) | **Put** /sources/{id} | Update Source (Full)
[**PutSourceSchema**](SourcesAPI.md#PutSourceSchema) | **Put** /sources/{sourceId}/schemas/{schemaId} | Update Source Schema (Full)
[**UpdateProvisioningPoliciesInBulk**](SourcesAPI.md#UpdateProvisioningPoliciesInBulk) | **Post** /sources/{sourceId}/provisioning-policies/bulk-update | Bulk Update Provisioning Policies
[**UpdateProvisioningPolicy**](SourcesAPI.md#UpdateProvisioningPolicy) | **Patch** /sources/{sourceId}/provisioning-policies/{usageType} | Partial update of Provisioning Policy
[**UpdateSource**](SourcesAPI.md#UpdateSource) | **Patch** /sources/{id} | Update Source (Partial)
[**UpdateSourceSchema**](SourcesAPI.md#UpdateSourceSchema) | **Patch** /sources/{sourceId}/schemas/{schemaId} | Update Source Schema (Partial)



## Create Provisioning Policy

> ProvisioningPolicyDto CreateProvisioningPolicy(ctx, sourceId).ProvisioningPolicyDto(provisioningPolicyDto).Execute()

This API generates a create policy/template based on field value transforms. This API is intended for use when setting up JDBC Provisioning type sources, but it will also work on other source types.
Transforms can be used in the provisioning policy to create a new attribute that you only need during provisioning.
Refer to [Transforms in Provisioning Policies](https://developer.sailpoint.com/idn/docs/transforms/guides/transforms-in-provisioning-policies) for more information.
A token with ORG_ADMIN authority is required to call this API.

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
    sourceId := "2c9180835d191a86015d28455b4a2329" // string | The Source id
    provisioningPolicyDto := *sailpoint.NewProvisioningPolicyDto("example provisioning policy for inactive identities") // ProvisioningPolicyDto | 

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.SourcesAPI.CreateProvisioningPolicy(context.Background(), sourceId).ProvisioningPolicyDto(provisioningPolicyDto).Execute()
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Creates a source in IdentityNow.

> Source CreateSource(ctx).Source(source).ProvisionAsCsv(provisionAsCsv).Execute()

This creates a specific source with a full source JSON representation. Any passwords are submitted as plain-text and encrypted upon receipt in IdentityNow.
A token with ORG_ADMIN, SOURCE_ADMIN, or SOURCE_SUBADMIN authority is required to call this API.

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
    source := *sailpoint.NewSource("My Source", *sailpoint.NewSourceOwner(), "active-directory") // Source | 
    provisionAsCsv := false // bool | If this parameter is `true`, it configures the source as a Delimited File (CSV) source. Setting this to `true` will automatically set the `type` of the source to `DelimitedFile`.  You must use this query parameter to create a Delimited File source as you would in the UI.  If you don't set this query parameter and you attempt to set the `type` attribute directly, the request won't correctly generate the source.   (optional)

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.SourcesAPI.CreateSource(context.Background()).Source(source).ProvisionAsCsv(provisionAsCsv).Execute()
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Create Schema on a Source

> Schema CreateSourceSchema(ctx, sourceId).Schema(schema).Execute()

Creates a new Schema on the specified Source in IdentityNow.


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
    sourceId := "2c9180835d191a86015d28455b4a2329" // string | The Source id.
    schema := *sailpoint.NewSchema() // Schema | 

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.SourcesAPI.CreateSourceSchema(context.Background(), sourceId).Schema(schema).Execute()
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
**sourceId** | **string** | The Source id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSourceSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **schema** | [**Schema**](Schema.md) |  | 

### Return type

[**Schema**](Schema.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete Provisioning Policy by UsageType

> DeleteProvisioningPolicy(ctx, sourceId, usageType).Execute()

Deletes the provisioning policy with the specified usage on an application.
A token with API, or ORG_ADMIN authority is required to call this API.

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
    sourceId := "2c9180835d191a86015d28455b4a2329" // string | The Source ID.
    usageType := sailpoint.UsageType("CREATE") // UsageType | The type of provisioning policy usage.  In IdentityNow, a source can support various provisioning operations. For example, when a joiner is added to a source, this may trigger both CREATE and UPDATE provisioning operations.  Each usage type is considered a provisioning policy.  A source can have any number of these provisioning policies defined.  These are the common usage types:  CREATE - This usage type relates to 'Create Account Profile', the provisioning template for the account to be created. For example, this would be used for a joiner on a source.   UPDATE - This usage type relates to 'Update Account Profile', the provisioning template for the 'Update' connector operations. For example, this would be used for an attribute sync on a source. ENABLE - This usage type relates to 'Enable Account Profile', the provisioning template for the account to be enabled. For example, this could be used for a joiner on a source once the joiner's account is created.  DISABLE - This usage type relates to 'Disable Account Profile', the provisioning template for the account to be disabled. For example, this could be used when a leaver is removed temporarily from a source.  You can use these four usage types for all your provisioning policy needs. 

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    r, err := apiClient.V3.SourcesAPI.DeleteProvisioningPolicy(context.Background(), sourceId, usageType).Execute()
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete Source by ID

> DeleteSource202Response DeleteSource(ctx, id).Execute()

This end-point deletes a specific source in IdentityNow.
A token with ORG_ADMIN, SOURCE_ADMIN, or SOURCE_SUBADMIN authority is required to call this API.
All of accounts on the source will be removed first, then the source will be deleted. Actual status of task execution can be retrieved via method GET `/task-status/{id}`

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
    id := "2c9180835d191a86015d28455b4a2329" // string | The Source id

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.SourcesAPI.DeleteSource(context.Background(), id).Execute()
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
**id** | **string** | The Source id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteSource202Response**](DeleteSource202Response.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete Source Schema by ID

> DeleteSourceSchema(ctx, sourceId, schemaId).Execute()



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
    sourceId := "2c9180835d191a86015d28455b4a2329" // string | The Source id.
    schemaId := "2c9180835d191a86015d28455b4a2329" // string | The Schema id.

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    r, err := apiClient.V3.SourcesAPI.DeleteSourceSchema(context.Background(), sourceId, schemaId).Execute()
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Downloads source accounts schema template

> GetAccountsSchema(ctx, id).Execute()

This API downloads the CSV schema that defines the account attributes on a source.
>**NOTE: This API is designated only for Delimited File sources.**

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
    id := "8c190e6787aa4ed9a90bd9d5344523fb" // string | The Source id

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    r, err := apiClient.V3.SourcesAPI.GetAccountsSchema(context.Background(), id).Execute()
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/csv, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Downloads source entitlements schema template

> GetEntitlementsSchema(ctx, id).SchemaName(schemaName).Execute()

This API downloads the CSV schema that defines the entitlement attributes on a source.

>**NOTE: This API is designated only for Delimited File sources.**

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
    id := "8c190e6787aa4ed9a90bd9d5344523fb" // string | The Source id
    schemaName := "?schemaName=group" // string | Name of entitlement schema (optional)

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    r, err := apiClient.V3.SourcesAPI.GetEntitlementsSchema(context.Background(), id).SchemaName(schemaName).Execute()
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/csv, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get Provisioning Policy by UsageType

> ProvisioningPolicyDto GetProvisioningPolicy(ctx, sourceId, usageType).Execute()

This end-point retrieves the ProvisioningPolicy with the specified usage on the specified Source in IdentityNow.
A token with API, ORG_ADMIN, SOURCE_ADMIN, or SOURCE_SUBADMIN authority is required to call this API.

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
    sourceId := "2c9180835d191a86015d28455b4a2329" // string | The Source ID.
    usageType := sailpoint.UsageType("CREATE") // UsageType | The type of provisioning policy usage.  In IdentityNow, a source can support various provisioning operations. For example, when a joiner is added to a source, this may trigger both CREATE and UPDATE provisioning operations.  Each usage type is considered a provisioning policy.  A source can have any number of these provisioning policies defined.  These are the common usage types:  CREATE - This usage type relates to 'Create Account Profile', the provisioning template for the account to be created. For example, this would be used for a joiner on a source.   UPDATE - This usage type relates to 'Update Account Profile', the provisioning template for the 'Update' connector operations. For example, this would be used for an attribute sync on a source. ENABLE - This usage type relates to 'Enable Account Profile', the provisioning template for the account to be enabled. For example, this could be used for a joiner on a source once the joiner's account is created.  DISABLE - This usage type relates to 'Disable Account Profile', the provisioning template for the account to be disabled. For example, this could be used when a leaver is removed temporarily from a source.  You can use these four usage types for all your provisioning policy needs. 

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.SourcesAPI.GetProvisioningPolicy(context.Background(), sourceId, usageType).Execute()
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get Source by ID

> Source GetSource(ctx, id).Execute()

This end-point gets a specific source in IdentityNow.
A token with ORG_ADMIN, SOURCE_ADMIN, or SOURCE_SUBADMIN authority is required to call this API.

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
    id := "2c9180835d191a86015d28455b4a2329" // string | The Source id

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.SourcesAPI.GetSource(context.Background(), id).Execute()
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
**id** | **string** | The Source id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Source**](Source.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Fetches source health by id

> SourceHealthDto GetSourceHealth(ctx, sourceId).Execute()

This endpoint fetches source health by source's id

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
    sourceId := "2c9180835d191a86015d28455b4a2329" // string | The Source id.

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.SourcesAPI.GetSourceHealth(context.Background(), sourceId).Execute()
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get Source Schema by ID

> Schema GetSourceSchema(ctx, sourceId, schemaId).Execute()

Get the Source Schema by ID in IdentityNow.


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
    sourceId := "2c9180835d191a86015d28455b4a2329" // string | The Source id.
    schemaId := "2c9180835d191a86015d28455b4a2329" // string | The Schema id.

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.SourcesAPI.GetSourceSchema(context.Background(), sourceId, schemaId).Execute()
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Uploads source accounts schema template

> Schema ImportAccountsSchema(ctx, id).File(file).Execute()

This API uploads a source schema template file to configure a source's account attributes.

To retrieve the file to modify and upload, log into Identity Now. 

Click **Admin** -> **Connections** -> **Sources** -> **`{SourceName}`** -> **Import Data** -> **Account Schema** -> **Options** -> **Download Schema**

>**NOTE: This API is designated only for Delimited File sources.**

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
    id := "8c190e6787aa4ed9a90bd9d5344523fb" // string | The Source id
    file := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.SourcesAPI.ImportAccountsSchema(context.Background(), id).File(file).Execute()
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Upload connector file to source

> Source ImportConnectorFile(ctx, sourceId).File(file).Execute()

This uploads a supplemental source connector file (like jdbc driver jars) to a source's S3 bucket. This also sends ETS and Audit events.
A token with ORG_ADMIN authority is required to call this API.

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
    sourceId := "2c9180835d191a86015d28455b4a2329" // string | The Source id.
    file := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.SourcesAPI.ImportConnectorFile(context.Background(), sourceId).File(file).Execute()
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Uploads source entitlements schema template

> Schema ImportEntitlementsSchema(ctx, id).SchemaName(schemaName).File(file).Execute()

This API uploads a source schema template file to configure a source's entitlement attributes.

To retrieve the file to modify and upload, log into Identity Now. 

Click **Admin** -> **Connections** -> **Sources** -> **`{SourceName}`** -> **Import Data** -> **Import Entitlements** -> **Download**

>**NOTE: This API is designated only for Delimited File sources.**

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
    id := "8c190e6787aa4ed9a90bd9d5344523fb" // string | The Source id
    schemaName := "?schemaName=group" // string | Name of entitlement schema (optional)
    file := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.SourcesAPI.ImportEntitlementsSchema(context.Background(), id).SchemaName(schemaName).File(file).Execute()
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Lists ProvisioningPolicies

> []ProvisioningPolicyDto ListProvisioningPolicies(ctx, sourceId).Execute()

This end-point lists all the ProvisioningPolicies in IdentityNow.
A token with API, or ORG_ADMIN authority is required to call this API.

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
    sourceId := "2c9180835d191a86015d28455b4a2329" // string | The Source id

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.SourcesAPI.ListProvisioningPolicies(context.Background(), sourceId).Execute()
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List Schemas on a Source

> []Schema ListSourceSchemas(ctx, sourceId).IncludeTypes(includeTypes).Execute()

Lists the Schemas that exist on the specified Source in IdentityNow.


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
    sourceId := "2c9180835d191a86015d28455b4a2329" // string | The Source ID.
    includeTypes := "group" // string | If set to 'group', then the account schema is filtered and only group schemas are returned. Only a value of 'group' is recognized. (optional)

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.SourcesAPI.ListSourceSchemas(context.Background(), sourceId).IncludeTypes(includeTypes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourcesAPI.ListSourceSchemas``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSourceSchemas`: []Schema
    fmt.Fprintf(os.Stdout, "Response from `SourcesAPI.ListSourceSchemas`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** | The Source ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSourceSchemasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeTypes** | **string** | If set to &#39;group&#39;, then the account schema is filtered and only group schemas are returned. Only a value of &#39;group&#39; is recognized. | 

### Return type

[**[]Schema**](Schema.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Lists all sources in IdentityNow.

> []Source ListSources(ctx).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).ForSubadmin(forSubadmin).Execute()

This end-point lists all the sources in IdentityNow.
A token with ORG_ADMIN, SOURCE_ADMIN, SOURCE_SUBADMIN, or ROLE_SUBADMIN authority is required to call this API.

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
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
    filters := "name eq "Employees"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq, in, ge, gt, le, lt, ne, isnull, sw*  **name**: *co, eq, in, sw, ge, gt, ne, isnull*  **type**: *eq, in, ge, gt, ne, isnull, sw*  **owner.id**: *eq, in, ge, gt, le, lt, ne, isnull, sw*  **features**: *ca, co*  **created**: *eq*  **modified**: *eq*  **managementWorkgroup.id**: *eq, ge, gt, in, le, lt, ne, isnull, sw*  **description**: *eq, sw*  **authoritative**: *eq, ne, isnull*  **healthy**: *isnull*  **status**: *eq, in, ge, gt, le, lt, ne, isnull, sw*  **connectionType**: *eq, ge, gt, in, le, lt, ne, isnull, sw*  **connectorName**: *eq, ge, gt, in, ne, isnull, sw* (optional)
    sorters := "name" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **type, created, modified, name, owner.name, healthy, status, id, description, owner.id, accountCorrelationConfig.id, accountCorrelationConfig.name, managerCorrelationRule.type, managerCorrelationRule.id, managerCorrelationRule.name, authoritative, managementWorkgroup.id, connectorName, connectionType** (optional)
    forSubadmin := "name" // string | Filter the returned list of sources for the identity specified by the parameter, which is the id of an identity with the role SOURCE_SUBADMIN. By convention, the value **me** indicates the identity id of the current user. Subadmins may only view Sources which they are able to administer; all other Sources will be filtered out when this parameter is set. If the current user is a SOURCE_SUBADMIN but fails to pass a valid value for this parameter, a 403 Forbidden is returned. (optional)

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.SourcesAPI.ListSources(context.Background()).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).ForSubadmin(forSubadmin).Execute()
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
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq, in, ge, gt, le, lt, ne, isnull, sw*  **name**: *co, eq, in, sw, ge, gt, ne, isnull*  **type**: *eq, in, ge, gt, ne, isnull, sw*  **owner.id**: *eq, in, ge, gt, le, lt, ne, isnull, sw*  **features**: *ca, co*  **created**: *eq*  **modified**: *eq*  **managementWorkgroup.id**: *eq, ge, gt, in, le, lt, ne, isnull, sw*  **description**: *eq, sw*  **authoritative**: *eq, ne, isnull*  **healthy**: *isnull*  **status**: *eq, in, ge, gt, le, lt, ne, isnull, sw*  **connectionType**: *eq, ge, gt, in, le, lt, ne, isnull, sw*  **connectorName**: *eq, ge, gt, in, ne, isnull, sw* | 
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **type, created, modified, name, owner.name, healthy, status, id, description, owner.id, accountCorrelationConfig.id, accountCorrelationConfig.name, managerCorrelationRule.type, managerCorrelationRule.id, managerCorrelationRule.name, authoritative, managementWorkgroup.id, connectorName, connectionType** | 
 **forSubadmin** | **string** | Filter the returned list of sources for the identity specified by the parameter, which is the id of an identity with the role SOURCE_SUBADMIN. By convention, the value **me** indicates the identity id of the current user. Subadmins may only view Sources which they are able to administer; all other Sources will be filtered out when this parameter is set. If the current user is a SOURCE_SUBADMIN but fails to pass a valid value for this parameter, a 403 Forbidden is returned. | 

### Return type

[**[]Source**](Source.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update Provisioning Policy by UsageType

> ProvisioningPolicyDto PutProvisioningPolicy(ctx, sourceId, usageType).ProvisioningPolicyDto(provisioningPolicyDto).Execute()

This end-point updates the provisioning policy with the specified usage on the specified source in IdentityNow.
Transforms can be used in the provisioning policy to create a new attribute that you only need during provisioning.
Refer to [Transforms in Provisioning Policies](https://developer.sailpoint.com/idn/docs/transforms/guides/transforms-in-provisioning-policies) for more information.
A token with API, ORG_ADMIN, SOURCE_ADMIN, or SOURCE_SUBADMIN authority is required to call this API.

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
    sourceId := "2c9180835d191a86015d28455b4a2329" // string | The Source ID.
    usageType := sailpoint.UsageType("CREATE") // UsageType | The type of provisioning policy usage.  In IdentityNow, a source can support various provisioning operations. For example, when a joiner is added to a source, this may trigger both CREATE and UPDATE provisioning operations.  Each usage type is considered a provisioning policy.  A source can have any number of these provisioning policies defined.  These are the common usage types:  CREATE - This usage type relates to 'Create Account Profile', the provisioning template for the account to be created. For example, this would be used for a joiner on a source.   UPDATE - This usage type relates to 'Update Account Profile', the provisioning template for the 'Update' connector operations. For example, this would be used for an attribute sync on a source. ENABLE - This usage type relates to 'Enable Account Profile', the provisioning template for the account to be enabled. For example, this could be used for a joiner on a source once the joiner's account is created.  DISABLE - This usage type relates to 'Disable Account Profile', the provisioning template for the account to be disabled. For example, this could be used when a leaver is removed temporarily from a source.  You can use these four usage types for all your provisioning policy needs. 
    provisioningPolicyDto := *sailpoint.NewProvisioningPolicyDto("example provisioning policy for inactive identities") // ProvisioningPolicyDto | 

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.SourcesAPI.PutProvisioningPolicy(context.Background(), sourceId, usageType).ProvisioningPolicyDto(provisioningPolicyDto).Execute()
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update Source (Full)

> Source PutSource(ctx, id).Source(source).Execute()

This API updates a source in IdentityNow, using a full object representation. In other words, the existing Source
configuration is completely replaced.

Some fields are immutable and cannot be changed, such as:

* id
* type
* authoritative
* connector
* connectorClass
* passwordPolicies

Attempts to modify these fields will result in a 400 error.

A token with ORG_ADMIN, SOURCE_ADMIN, or SOURCE_SUBADMIN authority is required to call this API.


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
    id := "2c9180835d191a86015d28455b4a2329" // string | The Source id
    source := *sailpoint.NewSource("My Source", *sailpoint.NewSourceOwner(), "active-directory") // Source | 

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.SourcesAPI.PutSource(context.Background(), id).Source(source).Execute()
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
**id** | **string** | The Source id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **source** | [**Source**](Source.md) |  | 

### Return type

[**Source**](Source.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update Source Schema (Full)

> Schema PutSourceSchema(ctx, sourceId, schemaId).Schema(schema).Execute()

This API will completely replace an existing Schema with the submitted payload. Some fields of the Schema cannot be updated. These fields are listed below.

* id
* name
* created
* modified

Any attempt to modify these fields will result in an error response with a status code of 400.

> `id` must remain in the request body, but it cannot be changed.  If `id` is omitted from the request body, the result will be a 400 error.


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
    sourceId := "2c9180835d191a86015d28455b4a2329" // string | The Source id.
    schemaId := "2c9180835d191a86015d28455b4a2329" // string | The Schema id.
    schema := *sailpoint.NewSchema() // Schema | 

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.SourcesAPI.PutSourceSchema(context.Background(), sourceId, schemaId).Schema(schema).Execute()
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Bulk Update Provisioning Policies

> []ProvisioningPolicyDto UpdateProvisioningPoliciesInBulk(ctx, sourceId).ProvisioningPolicyDto(provisioningPolicyDto).Execute()

This end-point updates a list of provisioning policies on the specified source in IdentityNow.
A token with API, or ORG_ADMIN authority is required to call this API.

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
    sourceId := "2c9180835d191a86015d28455b4a2329" // string | The Source id.
    provisioningPolicyDto := []sailpoint.ProvisioningPolicyDto{*sailpoint.NewProvisioningPolicyDto("example provisioning policy for inactive identities")} // []ProvisioningPolicyDto | 

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.SourcesAPI.UpdateProvisioningPoliciesInBulk(context.Background(), sourceId).ProvisioningPolicyDto(provisioningPolicyDto).Execute()
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Partial update of Provisioning Policy

> ProvisioningPolicyDto UpdateProvisioningPolicy(ctx, sourceId, usageType).JsonPatchOperation(jsonPatchOperation).Execute()

This API selectively updates an existing Provisioning Policy using a JSONPatch payload.
Transforms can be used in the provisioning policy to create a new attribute that you only need during provisioning.
Refer to [Transforms in Provisioning Policies](https://developer.sailpoint.com/idn/docs/transforms/guides/transforms-in-provisioning-policies) for more information.
A token with API, ORG_ADMIN, SOURCE_ADMIN, or SOURCE_SUBADMIN authority is required to call this API.

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
    sourceId := "2c9180835d191a86015d28455b4a2329" // string | The Source id.
    usageType := sailpoint.UsageType("CREATE") // UsageType | The type of provisioning policy usage.  In IdentityNow, a source can support various provisioning operations. For example, when a joiner is added to a source, this may trigger both CREATE and UPDATE provisioning operations.  Each usage type is considered a provisioning policy.  A source can have any number of these provisioning policies defined.  These are the common usage types:  CREATE - This usage type relates to 'Create Account Profile', the provisioning template for the account to be created. For example, this would be used for a joiner on a source.   UPDATE - This usage type relates to 'Update Account Profile', the provisioning template for the 'Update' connector operations. For example, this would be used for an attribute sync on a source. ENABLE - This usage type relates to 'Enable Account Profile', the provisioning template for the account to be enabled. For example, this could be used for a joiner on a source once the joiner's account is created.  DISABLE - This usage type relates to 'Disable Account Profile', the provisioning template for the account to be disabled. For example, this could be used when a leaver is removed temporarily from a source.  You can use these four usage types for all your provisioning policy needs. 
    jsonPatchOperation := []sailpoint.JsonPatchOperation{*sailpoint.NewJsonPatchOperation("replace", "/description")} // []JsonPatchOperation | The JSONPatch payload used to update the schema.

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.SourcesAPI.UpdateProvisioningPolicy(context.Background(), sourceId, usageType).JsonPatchOperation(jsonPatchOperation).Execute()
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update Source (Partial)

> Source UpdateSource(ctx, id).JsonPatchOperation(jsonPatchOperation).Execute()

This API partially updates a source in IdentityNow, using a list of patch operations according to the
[JSON Patch](https://tools.ietf.org/html/rfc6902) standard.

Some fields are immutable and cannot be changed, such as:

* id
* type
* authoritative
* created
* modified
* connector
* connectorClass
* passwordPolicies

Attempts to modify these fields will result in a 400 error.

A token with ORG_ADMIN, SOURCE_ADMIN, SOURCE_SUBADMIN, or API authority is required to call this API.


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
    id := "2c9180835d191a86015d28455b4a2329" // string | The Source id
    jsonPatchOperation := []sailpoint.JsonPatchOperation{*sailpoint.NewJsonPatchOperation("replace", "/description")} // []JsonPatchOperation | A list of account update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard. Any password changes are submitted as plain-text and encrypted upon receipt in IdentityNow.

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.SourcesAPI.UpdateSource(context.Background(), id).JsonPatchOperation(jsonPatchOperation).Execute()
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
**id** | **string** | The Source id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jsonPatchOperation** | [**[]JsonPatchOperation**](JsonPatchOperation.md) | A list of account update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard. Any password changes are submitted as plain-text and encrypted upon receipt in IdentityNow. | 

### Return type

[**Source**](Source.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update Source Schema (Partial)

> Schema UpdateSourceSchema(ctx, sourceId, schemaId).JsonPatchOperation(jsonPatchOperation).Execute()

Use this API to selectively update an existing Schema using a JSONPatch payload. 

The following schema fields are immutable and cannot be updated:

- id
- name
- created
- modified


To switch an account attribute to a group entitlement, you need to have the following in place:

- `isEntitlement: true`
- Must define a schema for the group and [add it to the source](https://developer.sailpoint.com/idn/api/v3/create-source-schema) before updating the `isGroup` flag.  For example, here is the `group` account attribute referencing a schema that defines the group:
```json
{
    "name": "groups",
    "type": "STRING",
    "schema": {
        "type": "CONNECTOR_SCHEMA",
        "id": "2c9180887671ff8c01767b4671fc7d60",
        "name": "group"
    },
    "description": "The groups, roles etc. that reference account group objects",
    "isMulti": true,
    "isEntitlement": true,
    "isGroup": true
}
```


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
    sourceId := "2c9180835d191a86015d28455b4a2329" // string | The Source id.
    schemaId := "2c9180835d191a86015d28455b4a2329" // string | The Schema id.
    jsonPatchOperation := []sailpoint.JsonPatchOperation{*sailpoint.NewJsonPatchOperation("replace", "/description")} // []JsonPatchOperation | The JSONPatch payload used to update the schema.

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.SourcesAPI.UpdateSourceSchema(context.Background(), sourceId, schemaId).JsonPatchOperation(jsonPatchOperation).Execute()
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

