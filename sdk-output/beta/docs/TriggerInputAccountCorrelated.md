# TriggerInputAccountCorrelated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identity** | [**TriggerInputAccountCorrelatedIdentity**](TriggerInputAccountCorrelatedIdentity.md) |  | 
**Source** | [**TriggerInputAccountCorrelatedSource**](TriggerInputAccountCorrelatedSource.md) |  | 
**Account** | [**TriggerInputAccountCorrelatedAccount**](TriggerInputAccountCorrelatedAccount.md) |  | 
**Attributes** | **map[string]interface{}** | The attributes associated with the account.  Attributes are unique per source. | 
**EntitlementCount** | Pointer to **int32** | The number of entitlements associated with this account. | [optional] 

## Methods

### NewTriggerInputAccountCorrelated

`func NewTriggerInputAccountCorrelated(identity TriggerInputAccountCorrelatedIdentity, source TriggerInputAccountCorrelatedSource, account TriggerInputAccountCorrelatedAccount, attributes map[string]interface{}, ) *TriggerInputAccountCorrelated`

NewTriggerInputAccountCorrelated instantiates a new TriggerInputAccountCorrelated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerInputAccountCorrelatedWithDefaults

`func NewTriggerInputAccountCorrelatedWithDefaults() *TriggerInputAccountCorrelated`

NewTriggerInputAccountCorrelatedWithDefaults instantiates a new TriggerInputAccountCorrelated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentity

`func (o *TriggerInputAccountCorrelated) GetIdentity() TriggerInputAccountCorrelatedIdentity`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *TriggerInputAccountCorrelated) GetIdentityOk() (*TriggerInputAccountCorrelatedIdentity, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *TriggerInputAccountCorrelated) SetIdentity(v TriggerInputAccountCorrelatedIdentity)`

SetIdentity sets Identity field to given value.


### GetSource

`func (o *TriggerInputAccountCorrelated) GetSource() TriggerInputAccountCorrelatedSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TriggerInputAccountCorrelated) GetSourceOk() (*TriggerInputAccountCorrelatedSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TriggerInputAccountCorrelated) SetSource(v TriggerInputAccountCorrelatedSource)`

SetSource sets Source field to given value.


### GetAccount

`func (o *TriggerInputAccountCorrelated) GetAccount() TriggerInputAccountCorrelatedAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *TriggerInputAccountCorrelated) GetAccountOk() (*TriggerInputAccountCorrelatedAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *TriggerInputAccountCorrelated) SetAccount(v TriggerInputAccountCorrelatedAccount)`

SetAccount sets Account field to given value.


### GetAttributes

`func (o *TriggerInputAccountCorrelated) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TriggerInputAccountCorrelated) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TriggerInputAccountCorrelated) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.


### GetEntitlementCount

`func (o *TriggerInputAccountCorrelated) GetEntitlementCount() int32`

GetEntitlementCount returns the EntitlementCount field if non-nil, zero value otherwise.

### GetEntitlementCountOk

`func (o *TriggerInputAccountCorrelated) GetEntitlementCountOk() (*int32, bool)`

GetEntitlementCountOk returns a tuple with the EntitlementCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementCount

`func (o *TriggerInputAccountCorrelated) SetEntitlementCount(v int32)`

SetEntitlementCount sets EntitlementCount field to given value.

### HasEntitlementCount

`func (o *TriggerInputAccountCorrelated) HasEntitlementCount() bool`

HasEntitlementCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


