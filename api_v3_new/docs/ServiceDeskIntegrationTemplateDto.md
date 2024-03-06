# ServiceDeskIntegrationTemplateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | System-generated unique ID of the Object | [optional] [readonly] 
**Name** | **string** | Name of the Object | 
**Created** | Pointer to **time.Time** | Creation date of the Object | [optional] [readonly] 
**Modified** | Pointer to **time.Time** | Last modification date of the Object | [optional] [readonly] 
**Type** | **string** | The &#39;type&#39; property specifies the type of the Service Desk integration template. | [default to "Web Service SDIM"]
**Attributes** | **map[string]interface{}** | The &#39;attributes&#39; property value is a map of attributes available for integrations using this Service Desk integration template. | 
**ProvisioningConfig** | [**ProvisioningConfig**](ProvisioningConfig.md) |  | 

## Methods

### NewServiceDeskIntegrationTemplateDto

`func NewServiceDeskIntegrationTemplateDto(name string, type_ string, attributes map[string]interface{}, provisioningConfig ProvisioningConfig, ) *ServiceDeskIntegrationTemplateDto`

NewServiceDeskIntegrationTemplateDto instantiates a new ServiceDeskIntegrationTemplateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceDeskIntegrationTemplateDtoWithDefaults

`func NewServiceDeskIntegrationTemplateDtoWithDefaults() *ServiceDeskIntegrationTemplateDto`

NewServiceDeskIntegrationTemplateDtoWithDefaults instantiates a new ServiceDeskIntegrationTemplateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServiceDeskIntegrationTemplateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceDeskIntegrationTemplateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceDeskIntegrationTemplateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceDeskIntegrationTemplateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ServiceDeskIntegrationTemplateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceDeskIntegrationTemplateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceDeskIntegrationTemplateDto) SetName(v string)`

SetName sets Name field to given value.


### GetCreated

`func (o *ServiceDeskIntegrationTemplateDto) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ServiceDeskIntegrationTemplateDto) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ServiceDeskIntegrationTemplateDto) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ServiceDeskIntegrationTemplateDto) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetModified

`func (o *ServiceDeskIntegrationTemplateDto) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *ServiceDeskIntegrationTemplateDto) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *ServiceDeskIntegrationTemplateDto) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *ServiceDeskIntegrationTemplateDto) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetType

`func (o *ServiceDeskIntegrationTemplateDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceDeskIntegrationTemplateDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServiceDeskIntegrationTemplateDto) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *ServiceDeskIntegrationTemplateDto) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ServiceDeskIntegrationTemplateDto) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ServiceDeskIntegrationTemplateDto) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.


### GetProvisioningConfig

`func (o *ServiceDeskIntegrationTemplateDto) GetProvisioningConfig() ProvisioningConfig`

GetProvisioningConfig returns the ProvisioningConfig field if non-nil, zero value otherwise.

### GetProvisioningConfigOk

`func (o *ServiceDeskIntegrationTemplateDto) GetProvisioningConfigOk() (*ProvisioningConfig, bool)`

GetProvisioningConfigOk returns a tuple with the ProvisioningConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningConfig

`func (o *ServiceDeskIntegrationTemplateDto) SetProvisioningConfig(v ProvisioningConfig)`

SetProvisioningConfig sets ProvisioningConfig field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


