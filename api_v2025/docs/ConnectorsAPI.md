# \ConnectorsAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v2025*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomConnector**](ConnectorsAPI.md#CreateCustomConnector) | **Post** /connectors | Create Custom Connector
[**DeleteCustomConnector**](ConnectorsAPI.md#DeleteCustomConnector) | **Delete** /connectors/{scriptName} | Delete Connector by Script Name
[**GetConnector**](ConnectorsAPI.md#GetConnector) | **Get** /connectors/{scriptName} | Get Connector by Script Name
[**GetConnectorCorrelationConfig**](ConnectorsAPI.md#GetConnectorCorrelationConfig) | **Get** /connectors/{scriptName}/correlation-config | Get Connector Correlation Configuration
[**GetConnectorList**](ConnectorsAPI.md#GetConnectorList) | **Get** /connectors | Get Connector List
[**GetConnectorSourceConfig**](ConnectorsAPI.md#GetConnectorSourceConfig) | **Get** /connectors/{scriptName}/source-config | Get Connector Source Configuration
[**GetConnectorSourceTemplate**](ConnectorsAPI.md#GetConnectorSourceTemplate) | **Get** /connectors/{scriptName}/source-template | Get Connector Source Template
[**GetConnectorTranslations**](ConnectorsAPI.md#GetConnectorTranslations) | **Get** /connectors/{scriptName}/translations/{locale} | Get Connector Translations
[**PutConnectorCorrelationConfig**](ConnectorsAPI.md#PutConnectorCorrelationConfig) | **Put** /connectors/{scriptName}/correlation-config | Update Connector Correlation Configuration
[**PutConnectorSourceConfig**](ConnectorsAPI.md#PutConnectorSourceConfig) | **Put** /connectors/{scriptName}/source-config | Update Connector Source Configuration
[**PutConnectorSourceTemplate**](ConnectorsAPI.md#PutConnectorSourceTemplate) | **Put** /connectors/{scriptName}/source-template | Update Connector Source Template
[**PutConnectorTranslations**](ConnectorsAPI.md#PutConnectorTranslations) | **Put** /connectors/{scriptName}/translations/{locale} | Update Connector Translations
[**UpdateConnector**](ConnectorsAPI.md#UpdateConnector) | **Patch** /connectors/{scriptName} | Update Connector by Script Name



## CreateCustomConnector

> V3ConnectorDto CreateCustomConnector(ctx).V3CreateConnectorDto(v3CreateConnectorDto).Execute()

Create Custom Connector



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
	v3CreateConnectorDto := *openapiclient.NewV3CreateConnectorDto("custom connector", "sailpoint.connector.OpenConnectorAdapter") // V3CreateConnectorDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectorsAPI.CreateCustomConnector(context.Background()).V3CreateConnectorDto(v3CreateConnectorDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorsAPI.CreateCustomConnector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCustomConnector`: V3ConnectorDto
	fmt.Fprintf(os.Stdout, "Response from `ConnectorsAPI.CreateCustomConnector`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v3CreateConnectorDto** | [**V3CreateConnectorDto**](V3CreateConnectorDto.md) |  | 

### Return type

[**V3ConnectorDto**](V3ConnectorDto.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomConnector

> DeleteCustomConnector(ctx, scriptName).Execute()

Delete Connector by Script Name



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
	scriptName := "aScriptName" // string | The scriptName value of the connector. ScriptName is the unique id generated at connector creation.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConnectorsAPI.DeleteCustomConnector(context.Background(), scriptName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorsAPI.DeleteCustomConnector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scriptName** | **string** | The scriptName value of the connector. ScriptName is the unique id generated at connector creation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomConnectorRequest struct via the builder pattern


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


## GetConnector

> ConnectorDetail GetConnector(ctx, scriptName).Locale(locale).Execute()

Get Connector by Script Name



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
	scriptName := "aScriptName" // string | The scriptName value of the connector. ScriptName is the unique id generated at connector creation.
	locale := "de" // string | The locale to apply to the config. If no viable locale is given, it will default to \"en\" (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectorsAPI.GetConnector(context.Background(), scriptName).Locale(locale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorsAPI.GetConnector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnector`: ConnectorDetail
	fmt.Fprintf(os.Stdout, "Response from `ConnectorsAPI.GetConnector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scriptName** | **string** | The scriptName value of the connector. ScriptName is the unique id generated at connector creation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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


## GetConnectorCorrelationConfig

> string GetConnectorCorrelationConfig(ctx, scriptName).Execute()

Get Connector Correlation Configuration



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
	scriptName := "aScriptName" // string | The scriptName value of the connector. Scriptname is the unique id generated at connector creation.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectorsAPI.GetConnectorCorrelationConfig(context.Background(), scriptName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorsAPI.GetConnectorCorrelationConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnectorCorrelationConfig`: string
	fmt.Fprintf(os.Stdout, "Response from `ConnectorsAPI.GetConnectorCorrelationConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scriptName** | **string** | The scriptName value of the connector. Scriptname is the unique id generated at connector creation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectorCorrelationConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectorList

> []V3ConnectorDto GetConnectorList(ctx).Filters(filters).Limit(limit).Offset(offset).Count(count).Locale(locale).Execute()

Get Connector List



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
	filters := "directConnect eq "true"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **name**: *sw, co*  **type**: *sw, co, eq*  **directConnect**: *eq*  **category**: *eq*  **features**: *ca*  **labels**: *ca* (optional)
	limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
	offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
	count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
	locale := "de" // string | The locale to apply to the config. If no viable locale is given, it will default to \"en\" (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectorsAPI.GetConnectorList(context.Background()).Filters(filters).Limit(limit).Offset(offset).Count(count).Locale(locale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorsAPI.GetConnectorList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnectorList`: []V3ConnectorDto
	fmt.Fprintf(os.Stdout, "Response from `ConnectorsAPI.GetConnectorList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectorListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **name**: *sw, co*  **type**: *sw, co, eq*  **directConnect**: *eq*  **category**: *eq*  **features**: *ca*  **labels**: *ca* | 
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **locale** | **string** | The locale to apply to the config. If no viable locale is given, it will default to \&quot;en\&quot; | 

### Return type

[**[]V3ConnectorDto**](V3ConnectorDto.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectorSourceConfig

> string GetConnectorSourceConfig(ctx, scriptName).Execute()

Get Connector Source Configuration



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
	scriptName := "aScriptName" // string | The scriptName value of the connector. ScriptName is the unique id generated at connector creation.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectorsAPI.GetConnectorSourceConfig(context.Background(), scriptName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorsAPI.GetConnectorSourceConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnectorSourceConfig`: string
	fmt.Fprintf(os.Stdout, "Response from `ConnectorsAPI.GetConnectorSourceConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scriptName** | **string** | The scriptName value of the connector. ScriptName is the unique id generated at connector creation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectorSourceConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectorSourceTemplate

> string GetConnectorSourceTemplate(ctx, scriptName).Execute()

Get Connector Source Template



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
	scriptName := "aScriptName" // string | The scriptName value of the connector. ScriptName is the unique id generated at connector creation.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectorsAPI.GetConnectorSourceTemplate(context.Background(), scriptName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorsAPI.GetConnectorSourceTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnectorSourceTemplate`: string
	fmt.Fprintf(os.Stdout, "Response from `ConnectorsAPI.GetConnectorSourceTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scriptName** | **string** | The scriptName value of the connector. ScriptName is the unique id generated at connector creation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectorSourceTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectorTranslations

> string GetConnectorTranslations(ctx, scriptName, locale).Execute()

Get Connector Translations



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
	scriptName := "aScriptName" // string | The scriptName value of the connector. Scriptname is the unique id generated at connector creation.
	locale := "de" // string | The locale to apply to the config. If no viable locale is given, it will default to \"en\"

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectorsAPI.GetConnectorTranslations(context.Background(), scriptName, locale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorsAPI.GetConnectorTranslations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnectorTranslations`: string
	fmt.Fprintf(os.Stdout, "Response from `ConnectorsAPI.GetConnectorTranslations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scriptName** | **string** | The scriptName value of the connector. Scriptname is the unique id generated at connector creation. | 
**locale** | **string** | The locale to apply to the config. If no viable locale is given, it will default to \&quot;en\&quot; | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectorTranslationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutConnectorCorrelationConfig

> UpdateDetail PutConnectorCorrelationConfig(ctx, scriptName).File(file).Execute()

Update Connector Correlation Configuration



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
	scriptName := "aScriptName" // string | The scriptName value of the connector. Scriptname is the unique id generated at connector creation.
	file := os.NewFile(1234, "some_file") // *os.File | connector correlation config xml file

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectorsAPI.PutConnectorCorrelationConfig(context.Background(), scriptName).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorsAPI.PutConnectorCorrelationConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutConnectorCorrelationConfig`: UpdateDetail
	fmt.Fprintf(os.Stdout, "Response from `ConnectorsAPI.PutConnectorCorrelationConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scriptName** | **string** | The scriptName value of the connector. Scriptname is the unique id generated at connector creation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutConnectorCorrelationConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | ***os.File** | connector correlation config xml file | 

### Return type

[**UpdateDetail**](UpdateDetail.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutConnectorSourceConfig

> UpdateDetail PutConnectorSourceConfig(ctx, scriptName).File(file).Execute()

Update Connector Source Configuration



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
	scriptName := "aScriptName" // string | The scriptName value of the connector. ScriptName is the unique id generated at connector creation.
	file := os.NewFile(1234, "some_file") // *os.File | connector source config xml file

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectorsAPI.PutConnectorSourceConfig(context.Background(), scriptName).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorsAPI.PutConnectorSourceConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutConnectorSourceConfig`: UpdateDetail
	fmt.Fprintf(os.Stdout, "Response from `ConnectorsAPI.PutConnectorSourceConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scriptName** | **string** | The scriptName value of the connector. ScriptName is the unique id generated at connector creation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutConnectorSourceConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | ***os.File** | connector source config xml file | 

### Return type

[**UpdateDetail**](UpdateDetail.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutConnectorSourceTemplate

> UpdateDetail PutConnectorSourceTemplate(ctx, scriptName).File(file).Execute()

Update Connector Source Template



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
	scriptName := "aScriptName" // string | The scriptName value of the connector. ScriptName is the unique id generated at connector creation.
	file := os.NewFile(1234, "some_file") // *os.File | connector source template xml file

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectorsAPI.PutConnectorSourceTemplate(context.Background(), scriptName).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorsAPI.PutConnectorSourceTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutConnectorSourceTemplate`: UpdateDetail
	fmt.Fprintf(os.Stdout, "Response from `ConnectorsAPI.PutConnectorSourceTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scriptName** | **string** | The scriptName value of the connector. ScriptName is the unique id generated at connector creation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutConnectorSourceTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | ***os.File** | connector source template xml file | 

### Return type

[**UpdateDetail**](UpdateDetail.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutConnectorTranslations

> UpdateDetail PutConnectorTranslations(ctx, scriptName, locale).Execute()

Update Connector Translations



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
	scriptName := "aScriptName" // string | The scriptName value of the connector. Scriptname is the unique id generated at connector creation.
	locale := "de" // string | The locale to apply to the config. If no viable locale is given, it will default to \"en\"

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectorsAPI.PutConnectorTranslations(context.Background(), scriptName, locale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorsAPI.PutConnectorTranslations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutConnectorTranslations`: UpdateDetail
	fmt.Fprintf(os.Stdout, "Response from `ConnectorsAPI.PutConnectorTranslations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scriptName** | **string** | The scriptName value of the connector. Scriptname is the unique id generated at connector creation. | 
**locale** | **string** | The locale to apply to the config. If no viable locale is given, it will default to \&quot;en\&quot; | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutConnectorTranslationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UpdateDetail**](UpdateDetail.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConnector

> ConnectorDetail UpdateConnector(ctx, scriptName).JsonPatchOperation(jsonPatchOperation).Execute()

Update Connector by Script Name



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
	scriptName := "aScriptName" // string | The scriptName value of the connector. ScriptName is the unique id generated at connector creation.
	jsonPatchOperation := []openapiclient.JsonPatchOperation{*openapiclient.NewJsonPatchOperation("replace", "/description")} // []JsonPatchOperation | A list of connector detail update operations 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectorsAPI.UpdateConnector(context.Background(), scriptName).JsonPatchOperation(jsonPatchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorsAPI.UpdateConnector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateConnector`: ConnectorDetail
	fmt.Fprintf(os.Stdout, "Response from `ConnectorsAPI.UpdateConnector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scriptName** | **string** | The scriptName value of the connector. ScriptName is the unique id generated at connector creation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jsonPatchOperation** | [**[]JsonPatchOperation**](JsonPatchOperation.md) | A list of connector detail update operations  | 

### Return type

[**ConnectorDetail**](ConnectorDetail.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

