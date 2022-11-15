# AccessProfileDocumentAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | Pointer to [**Reference1**](Reference1.md) |  | [optional] 
**Entitlements** | Pointer to [**[]BaseEntitlement**](BaseEntitlement.md) |  | [optional] 
**EntitlementCount** | Pointer to **int32** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAccessProfileDocumentAllOf

`func NewAccessProfileDocumentAllOf() *AccessProfileDocumentAllOf`

NewAccessProfileDocumentAllOf instantiates a new AccessProfileDocumentAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessProfileDocumentAllOfWithDefaults

`func NewAccessProfileDocumentAllOfWithDefaults() *AccessProfileDocumentAllOf`

NewAccessProfileDocumentAllOfWithDefaults instantiates a new AccessProfileDocumentAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *AccessProfileDocumentAllOf) GetSource() Reference1`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AccessProfileDocumentAllOf) GetSourceOk() (*Reference1, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AccessProfileDocumentAllOf) SetSource(v Reference1)`

SetSource sets Source field to given value.

### HasSource

`func (o *AccessProfileDocumentAllOf) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetEntitlements

`func (o *AccessProfileDocumentAllOf) GetEntitlements() []BaseEntitlement`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *AccessProfileDocumentAllOf) GetEntitlementsOk() (*[]BaseEntitlement, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *AccessProfileDocumentAllOf) SetEntitlements(v []BaseEntitlement)`

SetEntitlements sets Entitlements field to given value.

### HasEntitlements

`func (o *AccessProfileDocumentAllOf) HasEntitlements() bool`

HasEntitlements returns a boolean if a field has been set.

### GetEntitlementCount

`func (o *AccessProfileDocumentAllOf) GetEntitlementCount() int32`

GetEntitlementCount returns the EntitlementCount field if non-nil, zero value otherwise.

### GetEntitlementCountOk

`func (o *AccessProfileDocumentAllOf) GetEntitlementCountOk() (*int32, bool)`

GetEntitlementCountOk returns a tuple with the EntitlementCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementCount

`func (o *AccessProfileDocumentAllOf) SetEntitlementCount(v int32)`

SetEntitlementCount sets EntitlementCount field to given value.

### HasEntitlementCount

`func (o *AccessProfileDocumentAllOf) HasEntitlementCount() bool`

HasEntitlementCount returns a boolean if a field has been set.

### GetTags

`func (o *AccessProfileDocumentAllOf) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AccessProfileDocumentAllOf) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AccessProfileDocumentAllOf) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AccessProfileDocumentAllOf) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


