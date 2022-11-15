# TriggerInputAccountUncorrelated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identity** | [**TriggerInputAccountUncorrelatedIdentity**](TriggerInputAccountUncorrelatedIdentity.md) |  | 
**Source** | [**TriggerInputAccountUncorrelatedSource**](TriggerInputAccountUncorrelatedSource.md) |  | 
**Account** | [**TriggerInputAccountUncorrelatedAccount**](TriggerInputAccountUncorrelatedAccount.md) |  | 
**EntitlementCount** | Pointer to **int32** | The number of entitlements associated with this account. | [optional] 

## Methods

### NewTriggerInputAccountUncorrelated

`func NewTriggerInputAccountUncorrelated(identity TriggerInputAccountUncorrelatedIdentity, source TriggerInputAccountUncorrelatedSource, account TriggerInputAccountUncorrelatedAccount, ) *TriggerInputAccountUncorrelated`

NewTriggerInputAccountUncorrelated instantiates a new TriggerInputAccountUncorrelated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerInputAccountUncorrelatedWithDefaults

`func NewTriggerInputAccountUncorrelatedWithDefaults() *TriggerInputAccountUncorrelated`

NewTriggerInputAccountUncorrelatedWithDefaults instantiates a new TriggerInputAccountUncorrelated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentity

`func (o *TriggerInputAccountUncorrelated) GetIdentity() TriggerInputAccountUncorrelatedIdentity`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *TriggerInputAccountUncorrelated) GetIdentityOk() (*TriggerInputAccountUncorrelatedIdentity, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *TriggerInputAccountUncorrelated) SetIdentity(v TriggerInputAccountUncorrelatedIdentity)`

SetIdentity sets Identity field to given value.


### GetSource

`func (o *TriggerInputAccountUncorrelated) GetSource() TriggerInputAccountUncorrelatedSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TriggerInputAccountUncorrelated) GetSourceOk() (*TriggerInputAccountUncorrelatedSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TriggerInputAccountUncorrelated) SetSource(v TriggerInputAccountUncorrelatedSource)`

SetSource sets Source field to given value.


### GetAccount

`func (o *TriggerInputAccountUncorrelated) GetAccount() TriggerInputAccountUncorrelatedAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *TriggerInputAccountUncorrelated) GetAccountOk() (*TriggerInputAccountUncorrelatedAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *TriggerInputAccountUncorrelated) SetAccount(v TriggerInputAccountUncorrelatedAccount)`

SetAccount sets Account field to given value.


### GetEntitlementCount

`func (o *TriggerInputAccountUncorrelated) GetEntitlementCount() int32`

GetEntitlementCount returns the EntitlementCount field if non-nil, zero value otherwise.

### GetEntitlementCountOk

`func (o *TriggerInputAccountUncorrelated) GetEntitlementCountOk() (*int32, bool)`

GetEntitlementCountOk returns a tuple with the EntitlementCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementCount

`func (o *TriggerInputAccountUncorrelated) SetEntitlementCount(v int32)`

SetEntitlementCount sets EntitlementCount field to given value.

### HasEntitlementCount

`func (o *TriggerInputAccountUncorrelated) HasEntitlementCount() bool`

HasEntitlementCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


