# IdentityDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alias** | Pointer to **string** | Alternate unique identifier for the identity | [optional] 
**EmailAddress** | Pointer to **string** | The email address of the identity | [optional] 
**ProcessingState** | Pointer to **NullableString** | The processing state of the identity | [optional] 
**IdentityStatus** | Pointer to **string** | The identity&#39;s status in the system | [optional] 
**ManagerRef** | Pointer to [**BaseReferenceDto**](BaseReferenceDto.md) |  | [optional] 
**IsManager** | Pointer to **bool** | Whether this identity is a manager of another identity | [optional] 
**LastRefresh** | Pointer to **time.Time** | The last time the identity was refreshed by the system | [optional] 
**Attributes** | Pointer to **map[string]interface{}** | A map with the identity attributes for the identity | [optional] 

## Methods

### NewIdentityDto

`func NewIdentityDto() *IdentityDto`

NewIdentityDto instantiates a new IdentityDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityDtoWithDefaults

`func NewIdentityDtoWithDefaults() *IdentityDto`

NewIdentityDtoWithDefaults instantiates a new IdentityDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlias

`func (o *IdentityDto) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *IdentityDto) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *IdentityDto) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *IdentityDto) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetEmailAddress

`func (o *IdentityDto) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *IdentityDto) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *IdentityDto) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *IdentityDto) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetProcessingState

`func (o *IdentityDto) GetProcessingState() string`

GetProcessingState returns the ProcessingState field if non-nil, zero value otherwise.

### GetProcessingStateOk

`func (o *IdentityDto) GetProcessingStateOk() (*string, bool)`

GetProcessingStateOk returns a tuple with the ProcessingState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingState

`func (o *IdentityDto) SetProcessingState(v string)`

SetProcessingState sets ProcessingState field to given value.

### HasProcessingState

`func (o *IdentityDto) HasProcessingState() bool`

HasProcessingState returns a boolean if a field has been set.

### SetProcessingStateNil

`func (o *IdentityDto) SetProcessingStateNil(b bool)`

 SetProcessingStateNil sets the value for ProcessingState to be an explicit nil

### UnsetProcessingState
`func (o *IdentityDto) UnsetProcessingState()`

UnsetProcessingState ensures that no value is present for ProcessingState, not even an explicit nil
### GetIdentityStatus

`func (o *IdentityDto) GetIdentityStatus() string`

GetIdentityStatus returns the IdentityStatus field if non-nil, zero value otherwise.

### GetIdentityStatusOk

`func (o *IdentityDto) GetIdentityStatusOk() (*string, bool)`

GetIdentityStatusOk returns a tuple with the IdentityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityStatus

`func (o *IdentityDto) SetIdentityStatus(v string)`

SetIdentityStatus sets IdentityStatus field to given value.

### HasIdentityStatus

`func (o *IdentityDto) HasIdentityStatus() bool`

HasIdentityStatus returns a boolean if a field has been set.

### GetManagerRef

`func (o *IdentityDto) GetManagerRef() BaseReferenceDto`

GetManagerRef returns the ManagerRef field if non-nil, zero value otherwise.

### GetManagerRefOk

`func (o *IdentityDto) GetManagerRefOk() (*BaseReferenceDto, bool)`

GetManagerRefOk returns a tuple with the ManagerRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerRef

`func (o *IdentityDto) SetManagerRef(v BaseReferenceDto)`

SetManagerRef sets ManagerRef field to given value.

### HasManagerRef

`func (o *IdentityDto) HasManagerRef() bool`

HasManagerRef returns a boolean if a field has been set.

### GetIsManager

`func (o *IdentityDto) GetIsManager() bool`

GetIsManager returns the IsManager field if non-nil, zero value otherwise.

### GetIsManagerOk

`func (o *IdentityDto) GetIsManagerOk() (*bool, bool)`

GetIsManagerOk returns a tuple with the IsManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManager

`func (o *IdentityDto) SetIsManager(v bool)`

SetIsManager sets IsManager field to given value.

### HasIsManager

`func (o *IdentityDto) HasIsManager() bool`

HasIsManager returns a boolean if a field has been set.

### GetLastRefresh

`func (o *IdentityDto) GetLastRefresh() time.Time`

GetLastRefresh returns the LastRefresh field if non-nil, zero value otherwise.

### GetLastRefreshOk

`func (o *IdentityDto) GetLastRefreshOk() (*time.Time, bool)`

GetLastRefreshOk returns a tuple with the LastRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRefresh

`func (o *IdentityDto) SetLastRefresh(v time.Time)`

SetLastRefresh sets LastRefresh field to given value.

### HasLastRefresh

`func (o *IdentityDto) HasLastRefresh() bool`

HasLastRefresh returns a boolean if a field has been set.

### GetAttributes

`func (o *IdentityDto) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *IdentityDto) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *IdentityDto) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *IdentityDto) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


