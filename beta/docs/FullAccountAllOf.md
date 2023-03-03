# FullAccountAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authoritative** | Pointer to **bool** | Whether this account belongs to an authoritative source | [optional] 
**SystemAccount** | Pointer to **bool** | Whether this account is for the IdentityNow source | [optional] 
**Uncorrelated** | Pointer to **bool** | True if this account is not correlated to an identity | [optional] 
**Features** | Pointer to **string** | A string list containing the owning source&#39;s features | [optional] 

## Methods

### NewFullAccountAllOf

`func NewFullAccountAllOf() *FullAccountAllOf`

NewFullAccountAllOf instantiates a new FullAccountAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFullAccountAllOfWithDefaults

`func NewFullAccountAllOfWithDefaults() *FullAccountAllOf`

NewFullAccountAllOfWithDefaults instantiates a new FullAccountAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthoritative

`func (o *FullAccountAllOf) GetAuthoritative() bool`

GetAuthoritative returns the Authoritative field if non-nil, zero value otherwise.

### GetAuthoritativeOk

`func (o *FullAccountAllOf) GetAuthoritativeOk() (*bool, bool)`

GetAuthoritativeOk returns a tuple with the Authoritative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthoritative

`func (o *FullAccountAllOf) SetAuthoritative(v bool)`

SetAuthoritative sets Authoritative field to given value.

### HasAuthoritative

`func (o *FullAccountAllOf) HasAuthoritative() bool`

HasAuthoritative returns a boolean if a field has been set.

### GetSystemAccount

`func (o *FullAccountAllOf) GetSystemAccount() bool`

GetSystemAccount returns the SystemAccount field if non-nil, zero value otherwise.

### GetSystemAccountOk

`func (o *FullAccountAllOf) GetSystemAccountOk() (*bool, bool)`

GetSystemAccountOk returns a tuple with the SystemAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemAccount

`func (o *FullAccountAllOf) SetSystemAccount(v bool)`

SetSystemAccount sets SystemAccount field to given value.

### HasSystemAccount

`func (o *FullAccountAllOf) HasSystemAccount() bool`

HasSystemAccount returns a boolean if a field has been set.

### GetUncorrelated

`func (o *FullAccountAllOf) GetUncorrelated() bool`

GetUncorrelated returns the Uncorrelated field if non-nil, zero value otherwise.

### GetUncorrelatedOk

`func (o *FullAccountAllOf) GetUncorrelatedOk() (*bool, bool)`

GetUncorrelatedOk returns a tuple with the Uncorrelated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncorrelated

`func (o *FullAccountAllOf) SetUncorrelated(v bool)`

SetUncorrelated sets Uncorrelated field to given value.

### HasUncorrelated

`func (o *FullAccountAllOf) HasUncorrelated() bool`

HasUncorrelated returns a boolean if a field has been set.

### GetFeatures

`func (o *FullAccountAllOf) GetFeatures() string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *FullAccountAllOf) GetFeaturesOk() (*string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *FullAccountAllOf) SetFeatures(v string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *FullAccountAllOf) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


