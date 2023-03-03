# ServiceDeskIntegrationTemplateDtoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The &#39;type&#39; property specifies the type of the Service Desk integration template. | [default to "Web Service SDIM"]
**Attributes** | **map[string]interface{}** | The &#39;attributes&#39; property value is a map of attributes available for integrations using this Service Desk integration template. | 
**ProvisioningConfig** | [**ProvisioningConfig**](ProvisioningConfig.md) |  | 

## Methods

### NewServiceDeskIntegrationTemplateDtoAllOf

`func NewServiceDeskIntegrationTemplateDtoAllOf(type_ string, attributes map[string]interface{}, provisioningConfig ProvisioningConfig, ) *ServiceDeskIntegrationTemplateDtoAllOf`

NewServiceDeskIntegrationTemplateDtoAllOf instantiates a new ServiceDeskIntegrationTemplateDtoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceDeskIntegrationTemplateDtoAllOfWithDefaults

`func NewServiceDeskIntegrationTemplateDtoAllOfWithDefaults() *ServiceDeskIntegrationTemplateDtoAllOf`

NewServiceDeskIntegrationTemplateDtoAllOfWithDefaults instantiates a new ServiceDeskIntegrationTemplateDtoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ServiceDeskIntegrationTemplateDtoAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceDeskIntegrationTemplateDtoAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServiceDeskIntegrationTemplateDtoAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *ServiceDeskIntegrationTemplateDtoAllOf) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ServiceDeskIntegrationTemplateDtoAllOf) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ServiceDeskIntegrationTemplateDtoAllOf) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.


### GetProvisioningConfig

`func (o *ServiceDeskIntegrationTemplateDtoAllOf) GetProvisioningConfig() ProvisioningConfig`

GetProvisioningConfig returns the ProvisioningConfig field if non-nil, zero value otherwise.

### GetProvisioningConfigOk

`func (o *ServiceDeskIntegrationTemplateDtoAllOf) GetProvisioningConfigOk() (*ProvisioningConfig, bool)`

GetProvisioningConfigOk returns a tuple with the ProvisioningConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningConfig

`func (o *ServiceDeskIntegrationTemplateDtoAllOf) SetProvisioningConfig(v ProvisioningConfig)`

SetProvisioningConfig sets ProvisioningConfig field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


