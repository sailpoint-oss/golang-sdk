# TriggerInputIdentityCreated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identity** | [**TriggerInputIdentityCreatedIdentity**](TriggerInputIdentityCreatedIdentity.md) |  | 
**Attributes** | **map[string]interface{}** | The attributes assigned to the identity.  Attributes are determined by the identity profile. | 

## Methods

### NewTriggerInputIdentityCreated

`func NewTriggerInputIdentityCreated(identity TriggerInputIdentityCreatedIdentity, attributes map[string]interface{}, ) *TriggerInputIdentityCreated`

NewTriggerInputIdentityCreated instantiates a new TriggerInputIdentityCreated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerInputIdentityCreatedWithDefaults

`func NewTriggerInputIdentityCreatedWithDefaults() *TriggerInputIdentityCreated`

NewTriggerInputIdentityCreatedWithDefaults instantiates a new TriggerInputIdentityCreated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentity

`func (o *TriggerInputIdentityCreated) GetIdentity() TriggerInputIdentityCreatedIdentity`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *TriggerInputIdentityCreated) GetIdentityOk() (*TriggerInputIdentityCreatedIdentity, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *TriggerInputIdentityCreated) SetIdentity(v TriggerInputIdentityCreatedIdentity)`

SetIdentity sets Identity field to given value.


### GetAttributes

`func (o *TriggerInputIdentityCreated) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TriggerInputIdentityCreated) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TriggerInputIdentityCreated) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


