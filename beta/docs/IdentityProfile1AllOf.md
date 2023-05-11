# IdentityProfile1AllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | The description of the Identity Profile. | [optional] 
**Owner** | Pointer to [**NullableIdentityProfileAllOfOwner**](IdentityProfileAllOfOwner.md) |  | [optional] 
**Priority** | Pointer to **int64** | The priority for an Identity Profile. | [optional] 
**AuthoritativeSource** | [**IdentityProfile1AllOfAuthoritativeSource**](IdentityProfile1AllOfAuthoritativeSource.md) |  | 
**IdentityRefreshRequired** | Pointer to **bool** | True if a identity refresh is needed. Typically triggered when a change on the source has been made. | [optional] [default to false]
**IdentityCount** | Pointer to **int32** | The number of identities that belong to the Identity Profile. | [optional] 
**IdentityAttributeConfig** | Pointer to [**IdentityAttributeConfig1**](IdentityAttributeConfig1.md) |  | [optional] 
**IdentityExceptionReportReference** | Pointer to [**NullableIdentityExceptionReportReference1**](IdentityExceptionReportReference1.md) |  | [optional] 
**HasTimeBasedAttr** | Pointer to **bool** | Indicates the value of requiresPeriodicRefresh attribute for the Identity Profile. | [optional] [default to false]

## Methods

### NewIdentityProfile1AllOf

`func NewIdentityProfile1AllOf(authoritativeSource IdentityProfile1AllOfAuthoritativeSource, ) *IdentityProfile1AllOf`

NewIdentityProfile1AllOf instantiates a new IdentityProfile1AllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityProfile1AllOfWithDefaults

`func NewIdentityProfile1AllOfWithDefaults() *IdentityProfile1AllOf`

NewIdentityProfile1AllOfWithDefaults instantiates a new IdentityProfile1AllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *IdentityProfile1AllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IdentityProfile1AllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IdentityProfile1AllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IdentityProfile1AllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *IdentityProfile1AllOf) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *IdentityProfile1AllOf) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetOwner

`func (o *IdentityProfile1AllOf) GetOwner() IdentityProfileAllOfOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *IdentityProfile1AllOf) GetOwnerOk() (*IdentityProfileAllOfOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *IdentityProfile1AllOf) SetOwner(v IdentityProfileAllOfOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *IdentityProfile1AllOf) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *IdentityProfile1AllOf) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *IdentityProfile1AllOf) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetPriority

`func (o *IdentityProfile1AllOf) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *IdentityProfile1AllOf) GetPriorityOk() (*int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *IdentityProfile1AllOf) SetPriority(v int64)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *IdentityProfile1AllOf) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetAuthoritativeSource

`func (o *IdentityProfile1AllOf) GetAuthoritativeSource() IdentityProfile1AllOfAuthoritativeSource`

GetAuthoritativeSource returns the AuthoritativeSource field if non-nil, zero value otherwise.

### GetAuthoritativeSourceOk

`func (o *IdentityProfile1AllOf) GetAuthoritativeSourceOk() (*IdentityProfile1AllOfAuthoritativeSource, bool)`

GetAuthoritativeSourceOk returns a tuple with the AuthoritativeSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthoritativeSource

`func (o *IdentityProfile1AllOf) SetAuthoritativeSource(v IdentityProfile1AllOfAuthoritativeSource)`

SetAuthoritativeSource sets AuthoritativeSource field to given value.


### GetIdentityRefreshRequired

`func (o *IdentityProfile1AllOf) GetIdentityRefreshRequired() bool`

GetIdentityRefreshRequired returns the IdentityRefreshRequired field if non-nil, zero value otherwise.

### GetIdentityRefreshRequiredOk

`func (o *IdentityProfile1AllOf) GetIdentityRefreshRequiredOk() (*bool, bool)`

GetIdentityRefreshRequiredOk returns a tuple with the IdentityRefreshRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityRefreshRequired

`func (o *IdentityProfile1AllOf) SetIdentityRefreshRequired(v bool)`

SetIdentityRefreshRequired sets IdentityRefreshRequired field to given value.

### HasIdentityRefreshRequired

`func (o *IdentityProfile1AllOf) HasIdentityRefreshRequired() bool`

HasIdentityRefreshRequired returns a boolean if a field has been set.

### GetIdentityCount

`func (o *IdentityProfile1AllOf) GetIdentityCount() int32`

GetIdentityCount returns the IdentityCount field if non-nil, zero value otherwise.

### GetIdentityCountOk

`func (o *IdentityProfile1AllOf) GetIdentityCountOk() (*int32, bool)`

GetIdentityCountOk returns a tuple with the IdentityCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityCount

`func (o *IdentityProfile1AllOf) SetIdentityCount(v int32)`

SetIdentityCount sets IdentityCount field to given value.

### HasIdentityCount

`func (o *IdentityProfile1AllOf) HasIdentityCount() bool`

HasIdentityCount returns a boolean if a field has been set.

### GetIdentityAttributeConfig

`func (o *IdentityProfile1AllOf) GetIdentityAttributeConfig() IdentityAttributeConfig1`

GetIdentityAttributeConfig returns the IdentityAttributeConfig field if non-nil, zero value otherwise.

### GetIdentityAttributeConfigOk

`func (o *IdentityProfile1AllOf) GetIdentityAttributeConfigOk() (*IdentityAttributeConfig1, bool)`

GetIdentityAttributeConfigOk returns a tuple with the IdentityAttributeConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityAttributeConfig

`func (o *IdentityProfile1AllOf) SetIdentityAttributeConfig(v IdentityAttributeConfig1)`

SetIdentityAttributeConfig sets IdentityAttributeConfig field to given value.

### HasIdentityAttributeConfig

`func (o *IdentityProfile1AllOf) HasIdentityAttributeConfig() bool`

HasIdentityAttributeConfig returns a boolean if a field has been set.

### GetIdentityExceptionReportReference

`func (o *IdentityProfile1AllOf) GetIdentityExceptionReportReference() IdentityExceptionReportReference1`

GetIdentityExceptionReportReference returns the IdentityExceptionReportReference field if non-nil, zero value otherwise.

### GetIdentityExceptionReportReferenceOk

`func (o *IdentityProfile1AllOf) GetIdentityExceptionReportReferenceOk() (*IdentityExceptionReportReference1, bool)`

GetIdentityExceptionReportReferenceOk returns a tuple with the IdentityExceptionReportReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityExceptionReportReference

`func (o *IdentityProfile1AllOf) SetIdentityExceptionReportReference(v IdentityExceptionReportReference1)`

SetIdentityExceptionReportReference sets IdentityExceptionReportReference field to given value.

### HasIdentityExceptionReportReference

`func (o *IdentityProfile1AllOf) HasIdentityExceptionReportReference() bool`

HasIdentityExceptionReportReference returns a boolean if a field has been set.

### SetIdentityExceptionReportReferenceNil

`func (o *IdentityProfile1AllOf) SetIdentityExceptionReportReferenceNil(b bool)`

 SetIdentityExceptionReportReferenceNil sets the value for IdentityExceptionReportReference to be an explicit nil

### UnsetIdentityExceptionReportReference
`func (o *IdentityProfile1AllOf) UnsetIdentityExceptionReportReference()`

UnsetIdentityExceptionReportReference ensures that no value is present for IdentityExceptionReportReference, not even an explicit nil
### GetHasTimeBasedAttr

`func (o *IdentityProfile1AllOf) GetHasTimeBasedAttr() bool`

GetHasTimeBasedAttr returns the HasTimeBasedAttr field if non-nil, zero value otherwise.

### GetHasTimeBasedAttrOk

`func (o *IdentityProfile1AllOf) GetHasTimeBasedAttrOk() (*bool, bool)`

GetHasTimeBasedAttrOk returns a tuple with the HasTimeBasedAttr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasTimeBasedAttr

`func (o *IdentityProfile1AllOf) SetHasTimeBasedAttr(v bool)`

SetHasTimeBasedAttr sets HasTimeBasedAttr field to given value.

### HasHasTimeBasedAttr

`func (o *IdentityProfile1AllOf) HasHasTimeBasedAttr() bool`

HasHasTimeBasedAttr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


