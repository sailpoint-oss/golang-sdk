# AccountDocumentAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Modified** | Pointer to **NullableTime** | A date-time in ISO-8601 format | [optional] 
**Attributes** | Pointer to **map[string]interface{}** | a map or dictionary of key/value pairs | [optional] 
**Identity** | Pointer to [**DisplayReference**](DisplayReference.md) |  | [optional] 
**Access** | Pointer to [**[]Entitlement1**](Entitlement1.md) |  | [optional] 
**EntitlementCount** | Pointer to **int32** | The number of entitlements assigned to the account | [optional] 
**Uncorrelated** | Pointer to **bool** | Indicates if the account is not correlated to an identity | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAccountDocumentAllOf

`func NewAccountDocumentAllOf() *AccountDocumentAllOf`

NewAccountDocumentAllOf instantiates a new AccountDocumentAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountDocumentAllOfWithDefaults

`func NewAccountDocumentAllOfWithDefaults() *AccountDocumentAllOf`

NewAccountDocumentAllOfWithDefaults instantiates a new AccountDocumentAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModified

`func (o *AccountDocumentAllOf) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *AccountDocumentAllOf) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *AccountDocumentAllOf) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *AccountDocumentAllOf) HasModified() bool`

HasModified returns a boolean if a field has been set.

### SetModifiedNil

`func (o *AccountDocumentAllOf) SetModifiedNil(b bool)`

 SetModifiedNil sets the value for Modified to be an explicit nil

### UnsetModified
`func (o *AccountDocumentAllOf) UnsetModified()`

UnsetModified ensures that no value is present for Modified, not even an explicit nil
### GetAttributes

`func (o *AccountDocumentAllOf) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AccountDocumentAllOf) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AccountDocumentAllOf) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *AccountDocumentAllOf) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetIdentity

`func (o *AccountDocumentAllOf) GetIdentity() DisplayReference`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *AccountDocumentAllOf) GetIdentityOk() (*DisplayReference, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *AccountDocumentAllOf) SetIdentity(v DisplayReference)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *AccountDocumentAllOf) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetAccess

`func (o *AccountDocumentAllOf) GetAccess() []Entitlement1`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *AccountDocumentAllOf) GetAccessOk() (*[]Entitlement1, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *AccountDocumentAllOf) SetAccess(v []Entitlement1)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *AccountDocumentAllOf) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetEntitlementCount

`func (o *AccountDocumentAllOf) GetEntitlementCount() int32`

GetEntitlementCount returns the EntitlementCount field if non-nil, zero value otherwise.

### GetEntitlementCountOk

`func (o *AccountDocumentAllOf) GetEntitlementCountOk() (*int32, bool)`

GetEntitlementCountOk returns a tuple with the EntitlementCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementCount

`func (o *AccountDocumentAllOf) SetEntitlementCount(v int32)`

SetEntitlementCount sets EntitlementCount field to given value.

### HasEntitlementCount

`func (o *AccountDocumentAllOf) HasEntitlementCount() bool`

HasEntitlementCount returns a boolean if a field has been set.

### GetUncorrelated

`func (o *AccountDocumentAllOf) GetUncorrelated() bool`

GetUncorrelated returns the Uncorrelated field if non-nil, zero value otherwise.

### GetUncorrelatedOk

`func (o *AccountDocumentAllOf) GetUncorrelatedOk() (*bool, bool)`

GetUncorrelatedOk returns a tuple with the Uncorrelated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncorrelated

`func (o *AccountDocumentAllOf) SetUncorrelated(v bool)`

SetUncorrelated sets Uncorrelated field to given value.

### HasUncorrelated

`func (o *AccountDocumentAllOf) HasUncorrelated() bool`

HasUncorrelated returns a boolean if a field has been set.

### GetTags

`func (o *AccountDocumentAllOf) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AccountDocumentAllOf) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AccountDocumentAllOf) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AccountDocumentAllOf) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


