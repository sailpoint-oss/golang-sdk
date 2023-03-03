# AccountAttributesChanged

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identity** | [**TriggerInputAccountAttributesChangedIdentity**](TriggerInputAccountAttributesChangedIdentity.md) |  | 
**Source** | [**TriggerInputAccountAttributesChangedSource**](TriggerInputAccountAttributesChangedSource.md) |  | 
**Account** | [**TriggerInputAccountAttributesChangedAccount**](TriggerInputAccountAttributesChangedAccount.md) |  | 
**Changes** | [**[]TriggerInputAccountAttributesChangedChangesInner**](TriggerInputAccountAttributesChangedChangesInner.md) | A list of attributes that changed. | 

## Methods

### NewAccountAttributesChanged

`func NewAccountAttributesChanged(identity TriggerInputAccountAttributesChangedIdentity, source TriggerInputAccountAttributesChangedSource, account TriggerInputAccountAttributesChangedAccount, changes []TriggerInputAccountAttributesChangedChangesInner, ) *AccountAttributesChanged`

NewAccountAttributesChanged instantiates a new AccountAttributesChanged object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountAttributesChangedWithDefaults

`func NewAccountAttributesChangedWithDefaults() *AccountAttributesChanged`

NewAccountAttributesChangedWithDefaults instantiates a new AccountAttributesChanged object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentity

`func (o *AccountAttributesChanged) GetIdentity() TriggerInputAccountAttributesChangedIdentity`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *AccountAttributesChanged) GetIdentityOk() (*TriggerInputAccountAttributesChangedIdentity, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *AccountAttributesChanged) SetIdentity(v TriggerInputAccountAttributesChangedIdentity)`

SetIdentity sets Identity field to given value.


### GetSource

`func (o *AccountAttributesChanged) GetSource() TriggerInputAccountAttributesChangedSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AccountAttributesChanged) GetSourceOk() (*TriggerInputAccountAttributesChangedSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AccountAttributesChanged) SetSource(v TriggerInputAccountAttributesChangedSource)`

SetSource sets Source field to given value.


### GetAccount

`func (o *AccountAttributesChanged) GetAccount() TriggerInputAccountAttributesChangedAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AccountAttributesChanged) GetAccountOk() (*TriggerInputAccountAttributesChangedAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AccountAttributesChanged) SetAccount(v TriggerInputAccountAttributesChangedAccount)`

SetAccount sets Account field to given value.


### GetChanges

`func (o *AccountAttributesChanged) GetChanges() []TriggerInputAccountAttributesChangedChangesInner`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *AccountAttributesChanged) GetChangesOk() (*[]TriggerInputAccountAttributesChangedChangesInner, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *AccountAttributesChanged) SetChanges(v []TriggerInputAccountAttributesChangedChangesInner)`

SetChanges sets Changes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


