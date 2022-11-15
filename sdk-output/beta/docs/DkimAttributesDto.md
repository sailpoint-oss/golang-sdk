# DkimAttributesDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | The identity or domain address | [optional] 
**DkimEnabled** | Pointer to **bool** | Whether or not DKIM has been enabled for this domain / identity | [optional] 
**DkimTokens** | Pointer to **[]string** | The tokens to be added to a DNS for verification | [optional] 
**DkimVerificationStatus** | Pointer to **string** | The current status if the domain /identity has been verified. Ie Success, Failed, Pending | [optional] 

## Methods

### NewDkimAttributesDto

`func NewDkimAttributesDto() *DkimAttributesDto`

NewDkimAttributesDto instantiates a new DkimAttributesDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDkimAttributesDtoWithDefaults

`func NewDkimAttributesDtoWithDefaults() *DkimAttributesDto`

NewDkimAttributesDtoWithDefaults instantiates a new DkimAttributesDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *DkimAttributesDto) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DkimAttributesDto) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DkimAttributesDto) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *DkimAttributesDto) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetDkimEnabled

`func (o *DkimAttributesDto) GetDkimEnabled() bool`

GetDkimEnabled returns the DkimEnabled field if non-nil, zero value otherwise.

### GetDkimEnabledOk

`func (o *DkimAttributesDto) GetDkimEnabledOk() (*bool, bool)`

GetDkimEnabledOk returns a tuple with the DkimEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDkimEnabled

`func (o *DkimAttributesDto) SetDkimEnabled(v bool)`

SetDkimEnabled sets DkimEnabled field to given value.

### HasDkimEnabled

`func (o *DkimAttributesDto) HasDkimEnabled() bool`

HasDkimEnabled returns a boolean if a field has been set.

### GetDkimTokens

`func (o *DkimAttributesDto) GetDkimTokens() []string`

GetDkimTokens returns the DkimTokens field if non-nil, zero value otherwise.

### GetDkimTokensOk

`func (o *DkimAttributesDto) GetDkimTokensOk() (*[]string, bool)`

GetDkimTokensOk returns a tuple with the DkimTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDkimTokens

`func (o *DkimAttributesDto) SetDkimTokens(v []string)`

SetDkimTokens sets DkimTokens field to given value.

### HasDkimTokens

`func (o *DkimAttributesDto) HasDkimTokens() bool`

HasDkimTokens returns a boolean if a field has been set.

### GetDkimVerificationStatus

`func (o *DkimAttributesDto) GetDkimVerificationStatus() string`

GetDkimVerificationStatus returns the DkimVerificationStatus field if non-nil, zero value otherwise.

### GetDkimVerificationStatusOk

`func (o *DkimAttributesDto) GetDkimVerificationStatusOk() (*string, bool)`

GetDkimVerificationStatusOk returns a tuple with the DkimVerificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDkimVerificationStatus

`func (o *DkimAttributesDto) SetDkimVerificationStatus(v string)`

SetDkimVerificationStatus sets DkimVerificationStatus field to given value.

### HasDkimVerificationStatus

`func (o *DkimAttributesDto) HasDkimVerificationStatus() bool`

HasDkimVerificationStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


