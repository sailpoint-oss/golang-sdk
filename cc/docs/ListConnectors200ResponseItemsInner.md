# ListConnectors200ResponseItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationXml** | Pointer to **NullableString** |  | [optional] 
**ClassName** | Pointer to **NullableString** |  | [optional] 
**ConnectorMetadata** | Pointer to **map[string]interface{}** |  | [optional] 
**CorrelationConfigXml** | Pointer to **NullableString** |  | [optional] 
**DirectConnect** | Pointer to **bool** |  | [optional] 
**FileUpload** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**S3Location** | Pointer to **NullableString** |  | [optional] 
**Scope** | Pointer to **string** |  | [optional] 
**ScriptName** | Pointer to **string** |  | [optional] 
**SourceConfig** | Pointer to **NullableString** |  | [optional] 
**SourceConfigFrom** | Pointer to **NullableString** |  | [optional] 
**SourceConfigXml** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TranslationProperties** | Pointer to **map[string]interface{}** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UploadedFiles** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewListConnectors200ResponseItemsInner

`func NewListConnectors200ResponseItemsInner() *ListConnectors200ResponseItemsInner`

NewListConnectors200ResponseItemsInner instantiates a new ListConnectors200ResponseItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListConnectors200ResponseItemsInnerWithDefaults

`func NewListConnectors200ResponseItemsInnerWithDefaults() *ListConnectors200ResponseItemsInner`

NewListConnectors200ResponseItemsInnerWithDefaults instantiates a new ListConnectors200ResponseItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationXml

`func (o *ListConnectors200ResponseItemsInner) GetApplicationXml() string`

GetApplicationXml returns the ApplicationXml field if non-nil, zero value otherwise.

### GetApplicationXmlOk

`func (o *ListConnectors200ResponseItemsInner) GetApplicationXmlOk() (*string, bool)`

GetApplicationXmlOk returns a tuple with the ApplicationXml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationXml

`func (o *ListConnectors200ResponseItemsInner) SetApplicationXml(v string)`

SetApplicationXml sets ApplicationXml field to given value.

### HasApplicationXml

`func (o *ListConnectors200ResponseItemsInner) HasApplicationXml() bool`

HasApplicationXml returns a boolean if a field has been set.

### SetApplicationXmlNil

`func (o *ListConnectors200ResponseItemsInner) SetApplicationXmlNil(b bool)`

 SetApplicationXmlNil sets the value for ApplicationXml to be an explicit nil

### UnsetApplicationXml
`func (o *ListConnectors200ResponseItemsInner) UnsetApplicationXml()`

UnsetApplicationXml ensures that no value is present for ApplicationXml, not even an explicit nil
### GetClassName

`func (o *ListConnectors200ResponseItemsInner) GetClassName() string`

GetClassName returns the ClassName field if non-nil, zero value otherwise.

### GetClassNameOk

`func (o *ListConnectors200ResponseItemsInner) GetClassNameOk() (*string, bool)`

GetClassNameOk returns a tuple with the ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassName

`func (o *ListConnectors200ResponseItemsInner) SetClassName(v string)`

SetClassName sets ClassName field to given value.

### HasClassName

`func (o *ListConnectors200ResponseItemsInner) HasClassName() bool`

HasClassName returns a boolean if a field has been set.

### SetClassNameNil

`func (o *ListConnectors200ResponseItemsInner) SetClassNameNil(b bool)`

 SetClassNameNil sets the value for ClassName to be an explicit nil

### UnsetClassName
`func (o *ListConnectors200ResponseItemsInner) UnsetClassName()`

UnsetClassName ensures that no value is present for ClassName, not even an explicit nil
### GetConnectorMetadata

`func (o *ListConnectors200ResponseItemsInner) GetConnectorMetadata() map[string]interface{}`

GetConnectorMetadata returns the ConnectorMetadata field if non-nil, zero value otherwise.

### GetConnectorMetadataOk

`func (o *ListConnectors200ResponseItemsInner) GetConnectorMetadataOk() (*map[string]interface{}, bool)`

GetConnectorMetadataOk returns a tuple with the ConnectorMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorMetadata

`func (o *ListConnectors200ResponseItemsInner) SetConnectorMetadata(v map[string]interface{})`

SetConnectorMetadata sets ConnectorMetadata field to given value.

### HasConnectorMetadata

`func (o *ListConnectors200ResponseItemsInner) HasConnectorMetadata() bool`

HasConnectorMetadata returns a boolean if a field has been set.

### GetCorrelationConfigXml

`func (o *ListConnectors200ResponseItemsInner) GetCorrelationConfigXml() string`

GetCorrelationConfigXml returns the CorrelationConfigXml field if non-nil, zero value otherwise.

### GetCorrelationConfigXmlOk

`func (o *ListConnectors200ResponseItemsInner) GetCorrelationConfigXmlOk() (*string, bool)`

GetCorrelationConfigXmlOk returns a tuple with the CorrelationConfigXml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationConfigXml

`func (o *ListConnectors200ResponseItemsInner) SetCorrelationConfigXml(v string)`

SetCorrelationConfigXml sets CorrelationConfigXml field to given value.

### HasCorrelationConfigXml

`func (o *ListConnectors200ResponseItemsInner) HasCorrelationConfigXml() bool`

HasCorrelationConfigXml returns a boolean if a field has been set.

### SetCorrelationConfigXmlNil

`func (o *ListConnectors200ResponseItemsInner) SetCorrelationConfigXmlNil(b bool)`

 SetCorrelationConfigXmlNil sets the value for CorrelationConfigXml to be an explicit nil

### UnsetCorrelationConfigXml
`func (o *ListConnectors200ResponseItemsInner) UnsetCorrelationConfigXml()`

UnsetCorrelationConfigXml ensures that no value is present for CorrelationConfigXml, not even an explicit nil
### GetDirectConnect

`func (o *ListConnectors200ResponseItemsInner) GetDirectConnect() bool`

GetDirectConnect returns the DirectConnect field if non-nil, zero value otherwise.

### GetDirectConnectOk

`func (o *ListConnectors200ResponseItemsInner) GetDirectConnectOk() (*bool, bool)`

GetDirectConnectOk returns a tuple with the DirectConnect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectConnect

`func (o *ListConnectors200ResponseItemsInner) SetDirectConnect(v bool)`

SetDirectConnect sets DirectConnect field to given value.

### HasDirectConnect

`func (o *ListConnectors200ResponseItemsInner) HasDirectConnect() bool`

HasDirectConnect returns a boolean if a field has been set.

### GetFileUpload

`func (o *ListConnectors200ResponseItemsInner) GetFileUpload() bool`

GetFileUpload returns the FileUpload field if non-nil, zero value otherwise.

### GetFileUploadOk

`func (o *ListConnectors200ResponseItemsInner) GetFileUploadOk() (*bool, bool)`

GetFileUploadOk returns a tuple with the FileUpload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUpload

`func (o *ListConnectors200ResponseItemsInner) SetFileUpload(v bool)`

SetFileUpload sets FileUpload field to given value.

### HasFileUpload

`func (o *ListConnectors200ResponseItemsInner) HasFileUpload() bool`

HasFileUpload returns a boolean if a field has been set.

### GetId

`func (o *ListConnectors200ResponseItemsInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListConnectors200ResponseItemsInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListConnectors200ResponseItemsInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ListConnectors200ResponseItemsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListConnectors200ResponseItemsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListConnectors200ResponseItemsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListConnectors200ResponseItemsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListConnectors200ResponseItemsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetS3Location

`func (o *ListConnectors200ResponseItemsInner) GetS3Location() string`

GetS3Location returns the S3Location field if non-nil, zero value otherwise.

### GetS3LocationOk

`func (o *ListConnectors200ResponseItemsInner) GetS3LocationOk() (*string, bool)`

GetS3LocationOk returns a tuple with the S3Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Location

`func (o *ListConnectors200ResponseItemsInner) SetS3Location(v string)`

SetS3Location sets S3Location field to given value.

### HasS3Location

`func (o *ListConnectors200ResponseItemsInner) HasS3Location() bool`

HasS3Location returns a boolean if a field has been set.

### SetS3LocationNil

`func (o *ListConnectors200ResponseItemsInner) SetS3LocationNil(b bool)`

 SetS3LocationNil sets the value for S3Location to be an explicit nil

### UnsetS3Location
`func (o *ListConnectors200ResponseItemsInner) UnsetS3Location()`

UnsetS3Location ensures that no value is present for S3Location, not even an explicit nil
### GetScope

`func (o *ListConnectors200ResponseItemsInner) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ListConnectors200ResponseItemsInner) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ListConnectors200ResponseItemsInner) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *ListConnectors200ResponseItemsInner) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetScriptName

`func (o *ListConnectors200ResponseItemsInner) GetScriptName() string`

GetScriptName returns the ScriptName field if non-nil, zero value otherwise.

### GetScriptNameOk

`func (o *ListConnectors200ResponseItemsInner) GetScriptNameOk() (*string, bool)`

GetScriptNameOk returns a tuple with the ScriptName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptName

`func (o *ListConnectors200ResponseItemsInner) SetScriptName(v string)`

SetScriptName sets ScriptName field to given value.

### HasScriptName

`func (o *ListConnectors200ResponseItemsInner) HasScriptName() bool`

HasScriptName returns a boolean if a field has been set.

### GetSourceConfig

`func (o *ListConnectors200ResponseItemsInner) GetSourceConfig() string`

GetSourceConfig returns the SourceConfig field if non-nil, zero value otherwise.

### GetSourceConfigOk

`func (o *ListConnectors200ResponseItemsInner) GetSourceConfigOk() (*string, bool)`

GetSourceConfigOk returns a tuple with the SourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceConfig

`func (o *ListConnectors200ResponseItemsInner) SetSourceConfig(v string)`

SetSourceConfig sets SourceConfig field to given value.

### HasSourceConfig

`func (o *ListConnectors200ResponseItemsInner) HasSourceConfig() bool`

HasSourceConfig returns a boolean if a field has been set.

### SetSourceConfigNil

`func (o *ListConnectors200ResponseItemsInner) SetSourceConfigNil(b bool)`

 SetSourceConfigNil sets the value for SourceConfig to be an explicit nil

### UnsetSourceConfig
`func (o *ListConnectors200ResponseItemsInner) UnsetSourceConfig()`

UnsetSourceConfig ensures that no value is present for SourceConfig, not even an explicit nil
### GetSourceConfigFrom

`func (o *ListConnectors200ResponseItemsInner) GetSourceConfigFrom() string`

GetSourceConfigFrom returns the SourceConfigFrom field if non-nil, zero value otherwise.

### GetSourceConfigFromOk

`func (o *ListConnectors200ResponseItemsInner) GetSourceConfigFromOk() (*string, bool)`

GetSourceConfigFromOk returns a tuple with the SourceConfigFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceConfigFrom

`func (o *ListConnectors200ResponseItemsInner) SetSourceConfigFrom(v string)`

SetSourceConfigFrom sets SourceConfigFrom field to given value.

### HasSourceConfigFrom

`func (o *ListConnectors200ResponseItemsInner) HasSourceConfigFrom() bool`

HasSourceConfigFrom returns a boolean if a field has been set.

### SetSourceConfigFromNil

`func (o *ListConnectors200ResponseItemsInner) SetSourceConfigFromNil(b bool)`

 SetSourceConfigFromNil sets the value for SourceConfigFrom to be an explicit nil

### UnsetSourceConfigFrom
`func (o *ListConnectors200ResponseItemsInner) UnsetSourceConfigFrom()`

UnsetSourceConfigFrom ensures that no value is present for SourceConfigFrom, not even an explicit nil
### GetSourceConfigXml

`func (o *ListConnectors200ResponseItemsInner) GetSourceConfigXml() string`

GetSourceConfigXml returns the SourceConfigXml field if non-nil, zero value otherwise.

### GetSourceConfigXmlOk

`func (o *ListConnectors200ResponseItemsInner) GetSourceConfigXmlOk() (*string, bool)`

GetSourceConfigXmlOk returns a tuple with the SourceConfigXml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceConfigXml

`func (o *ListConnectors200ResponseItemsInner) SetSourceConfigXml(v string)`

SetSourceConfigXml sets SourceConfigXml field to given value.

### HasSourceConfigXml

`func (o *ListConnectors200ResponseItemsInner) HasSourceConfigXml() bool`

HasSourceConfigXml returns a boolean if a field has been set.

### SetSourceConfigXmlNil

`func (o *ListConnectors200ResponseItemsInner) SetSourceConfigXmlNil(b bool)`

 SetSourceConfigXmlNil sets the value for SourceConfigXml to be an explicit nil

### UnsetSourceConfigXml
`func (o *ListConnectors200ResponseItemsInner) UnsetSourceConfigXml()`

UnsetSourceConfigXml ensures that no value is present for SourceConfigXml, not even an explicit nil
### GetStatus

`func (o *ListConnectors200ResponseItemsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListConnectors200ResponseItemsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListConnectors200ResponseItemsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListConnectors200ResponseItemsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTranslationProperties

`func (o *ListConnectors200ResponseItemsInner) GetTranslationProperties() map[string]interface{}`

GetTranslationProperties returns the TranslationProperties field if non-nil, zero value otherwise.

### GetTranslationPropertiesOk

`func (o *ListConnectors200ResponseItemsInner) GetTranslationPropertiesOk() (*map[string]interface{}, bool)`

GetTranslationPropertiesOk returns a tuple with the TranslationProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslationProperties

`func (o *ListConnectors200ResponseItemsInner) SetTranslationProperties(v map[string]interface{})`

SetTranslationProperties sets TranslationProperties field to given value.

### HasTranslationProperties

`func (o *ListConnectors200ResponseItemsInner) HasTranslationProperties() bool`

HasTranslationProperties returns a boolean if a field has been set.

### GetType

`func (o *ListConnectors200ResponseItemsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListConnectors200ResponseItemsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListConnectors200ResponseItemsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ListConnectors200ResponseItemsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUploadedFiles

`func (o *ListConnectors200ResponseItemsInner) GetUploadedFiles() []map[string]interface{}`

GetUploadedFiles returns the UploadedFiles field if non-nil, zero value otherwise.

### GetUploadedFilesOk

`func (o *ListConnectors200ResponseItemsInner) GetUploadedFilesOk() (*[]map[string]interface{}, bool)`

GetUploadedFilesOk returns a tuple with the UploadedFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadedFiles

`func (o *ListConnectors200ResponseItemsInner) SetUploadedFiles(v []map[string]interface{})`

SetUploadedFiles sets UploadedFiles field to given value.

### HasUploadedFiles

`func (o *ListConnectors200ResponseItemsInner) HasUploadedFiles() bool`

HasUploadedFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


