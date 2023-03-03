# IdentityAttributesChanged

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identity** | [**TriggerInputIdentityAttributesChangedIdentity**](TriggerInputIdentityAttributesChangedIdentity.md) |  | 
**Changes** | [**[]TriggerInputIdentityAttributesChangedChangesInner**](TriggerInputIdentityAttributesChangedChangesInner.md) | A list of one or more identity attributes that changed on the identity. | 

## Methods

### NewIdentityAttributesChanged

`func NewIdentityAttributesChanged(identity TriggerInputIdentityAttributesChangedIdentity, changes []TriggerInputIdentityAttributesChangedChangesInner, ) *IdentityAttributesChanged`

NewIdentityAttributesChanged instantiates a new IdentityAttributesChanged object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityAttributesChangedWithDefaults

`func NewIdentityAttributesChangedWithDefaults() *IdentityAttributesChanged`

NewIdentityAttributesChangedWithDefaults instantiates a new IdentityAttributesChanged object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentity

`func (o *IdentityAttributesChanged) GetIdentity() TriggerInputIdentityAttributesChangedIdentity`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *IdentityAttributesChanged) GetIdentityOk() (*TriggerInputIdentityAttributesChangedIdentity, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *IdentityAttributesChanged) SetIdentity(v TriggerInputIdentityAttributesChangedIdentity)`

SetIdentity sets Identity field to given value.


### GetChanges

`func (o *IdentityAttributesChanged) GetChanges() []TriggerInputIdentityAttributesChangedChangesInner`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *IdentityAttributesChanged) GetChangesOk() (*[]TriggerInputIdentityAttributesChangedChangesInner, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *IdentityAttributesChanged) SetChanges(v []TriggerInputIdentityAttributesChangedChangesInner)`

SetChanges sets Changes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


