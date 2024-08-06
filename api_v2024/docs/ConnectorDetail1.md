# ConnectorDetail1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The connector name | [optional] 
**SourceConfigXml** | Pointer to **string** | XML representation of the source config data | [optional] 
**SourceConfig** | Pointer to **string** | JSON representation of the source config data | [optional] 
**DirectConnect** | Pointer to **bool** | true if the source is a direct connect source | [optional] 
**FileUpload** | Pointer to **bool** | Connector config&#39;s file upload attribute, false if not there | [optional] 
**UploadedFiles** | Pointer to **string** | List of uploaded file strings for the connector | [optional] 
**ConnectorMetadata** | Pointer to **map[string]interface{}** | Object containing metadata pertinent to the UI to be used | [optional] 

## Methods

### NewConnectorDetail1

`func NewConnectorDetail1() *ConnectorDetail1`

NewConnectorDetail1 instantiates a new ConnectorDetail1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorDetail1WithDefaults

`func NewConnectorDetail1WithDefaults() *ConnectorDetail1`

NewConnectorDetail1WithDefaults instantiates a new ConnectorDetail1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ConnectorDetail1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectorDetail1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectorDetail1) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConnectorDetail1) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSourceConfigXml

`func (o *ConnectorDetail1) GetSourceConfigXml() string`

GetSourceConfigXml returns the SourceConfigXml field if non-nil, zero value otherwise.

### GetSourceConfigXmlOk

`func (o *ConnectorDetail1) GetSourceConfigXmlOk() (*string, bool)`

GetSourceConfigXmlOk returns a tuple with the SourceConfigXml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceConfigXml

`func (o *ConnectorDetail1) SetSourceConfigXml(v string)`

SetSourceConfigXml sets SourceConfigXml field to given value.

### HasSourceConfigXml

`func (o *ConnectorDetail1) HasSourceConfigXml() bool`

HasSourceConfigXml returns a boolean if a field has been set.

### GetSourceConfig

`func (o *ConnectorDetail1) GetSourceConfig() string`

GetSourceConfig returns the SourceConfig field if non-nil, zero value otherwise.

### GetSourceConfigOk

`func (o *ConnectorDetail1) GetSourceConfigOk() (*string, bool)`

GetSourceConfigOk returns a tuple with the SourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceConfig

`func (o *ConnectorDetail1) SetSourceConfig(v string)`

SetSourceConfig sets SourceConfig field to given value.

### HasSourceConfig

`func (o *ConnectorDetail1) HasSourceConfig() bool`

HasSourceConfig returns a boolean if a field has been set.

### GetDirectConnect

`func (o *ConnectorDetail1) GetDirectConnect() bool`

GetDirectConnect returns the DirectConnect field if non-nil, zero value otherwise.

### GetDirectConnectOk

`func (o *ConnectorDetail1) GetDirectConnectOk() (*bool, bool)`

GetDirectConnectOk returns a tuple with the DirectConnect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectConnect

`func (o *ConnectorDetail1) SetDirectConnect(v bool)`

SetDirectConnect sets DirectConnect field to given value.

### HasDirectConnect

`func (o *ConnectorDetail1) HasDirectConnect() bool`

HasDirectConnect returns a boolean if a field has been set.

### GetFileUpload

`func (o *ConnectorDetail1) GetFileUpload() bool`

GetFileUpload returns the FileUpload field if non-nil, zero value otherwise.

### GetFileUploadOk

`func (o *ConnectorDetail1) GetFileUploadOk() (*bool, bool)`

GetFileUploadOk returns a tuple with the FileUpload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUpload

`func (o *ConnectorDetail1) SetFileUpload(v bool)`

SetFileUpload sets FileUpload field to given value.

### HasFileUpload

`func (o *ConnectorDetail1) HasFileUpload() bool`

HasFileUpload returns a boolean if a field has been set.

### GetUploadedFiles

`func (o *ConnectorDetail1) GetUploadedFiles() string`

GetUploadedFiles returns the UploadedFiles field if non-nil, zero value otherwise.

### GetUploadedFilesOk

`func (o *ConnectorDetail1) GetUploadedFilesOk() (*string, bool)`

GetUploadedFilesOk returns a tuple with the UploadedFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadedFiles

`func (o *ConnectorDetail1) SetUploadedFiles(v string)`

SetUploadedFiles sets UploadedFiles field to given value.

### HasUploadedFiles

`func (o *ConnectorDetail1) HasUploadedFiles() bool`

HasUploadedFiles returns a boolean if a field has been set.

### GetConnectorMetadata

`func (o *ConnectorDetail1) GetConnectorMetadata() map[string]interface{}`

GetConnectorMetadata returns the ConnectorMetadata field if non-nil, zero value otherwise.

### GetConnectorMetadataOk

`func (o *ConnectorDetail1) GetConnectorMetadataOk() (*map[string]interface{}, bool)`

GetConnectorMetadataOk returns a tuple with the ConnectorMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorMetadata

`func (o *ConnectorDetail1) SetConnectorMetadata(v map[string]interface{})`

SetConnectorMetadata sets ConnectorMetadata field to given value.

### HasConnectorMetadata

`func (o *ConnectorDetail1) HasConnectorMetadata() bool`

HasConnectorMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


