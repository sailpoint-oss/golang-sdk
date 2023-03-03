# LifecycleStateAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Whether the lifecycle state is enabled or disabled. | [optional] 
**TechnicalName** | **string** | The technical name for lifecycle state. This is for internal use. | 
**Description** | Pointer to **string** | Lifecycle state description. | [optional] 
**IdentityCount** | Pointer to **int32** | Number of identities that have the lifecycle state. | [optional] [readonly] 
**EmailNotificationOption** | Pointer to [**EmailNotificationOption**](EmailNotificationOption.md) |  | [optional] 
**AccountActions** | Pointer to [**[]AccountAction**](AccountAction.md) |  | [optional] 
**AccessProfileIds** | Pointer to **[]string** | List of unique access-profile IDs that are associated with the lifecycle state. | [optional] 

## Methods

### NewLifecycleStateAllOf

`func NewLifecycleStateAllOf(technicalName string, ) *LifecycleStateAllOf`

NewLifecycleStateAllOf instantiates a new LifecycleStateAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLifecycleStateAllOfWithDefaults

`func NewLifecycleStateAllOfWithDefaults() *LifecycleStateAllOf`

NewLifecycleStateAllOfWithDefaults instantiates a new LifecycleStateAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *LifecycleStateAllOf) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *LifecycleStateAllOf) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *LifecycleStateAllOf) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *LifecycleStateAllOf) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTechnicalName

`func (o *LifecycleStateAllOf) GetTechnicalName() string`

GetTechnicalName returns the TechnicalName field if non-nil, zero value otherwise.

### GetTechnicalNameOk

`func (o *LifecycleStateAllOf) GetTechnicalNameOk() (*string, bool)`

GetTechnicalNameOk returns a tuple with the TechnicalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnicalName

`func (o *LifecycleStateAllOf) SetTechnicalName(v string)`

SetTechnicalName sets TechnicalName field to given value.


### GetDescription

`func (o *LifecycleStateAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LifecycleStateAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LifecycleStateAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LifecycleStateAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIdentityCount

`func (o *LifecycleStateAllOf) GetIdentityCount() int32`

GetIdentityCount returns the IdentityCount field if non-nil, zero value otherwise.

### GetIdentityCountOk

`func (o *LifecycleStateAllOf) GetIdentityCountOk() (*int32, bool)`

GetIdentityCountOk returns a tuple with the IdentityCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityCount

`func (o *LifecycleStateAllOf) SetIdentityCount(v int32)`

SetIdentityCount sets IdentityCount field to given value.

### HasIdentityCount

`func (o *LifecycleStateAllOf) HasIdentityCount() bool`

HasIdentityCount returns a boolean if a field has been set.

### GetEmailNotificationOption

`func (o *LifecycleStateAllOf) GetEmailNotificationOption() EmailNotificationOption`

GetEmailNotificationOption returns the EmailNotificationOption field if non-nil, zero value otherwise.

### GetEmailNotificationOptionOk

`func (o *LifecycleStateAllOf) GetEmailNotificationOptionOk() (*EmailNotificationOption, bool)`

GetEmailNotificationOptionOk returns a tuple with the EmailNotificationOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailNotificationOption

`func (o *LifecycleStateAllOf) SetEmailNotificationOption(v EmailNotificationOption)`

SetEmailNotificationOption sets EmailNotificationOption field to given value.

### HasEmailNotificationOption

`func (o *LifecycleStateAllOf) HasEmailNotificationOption() bool`

HasEmailNotificationOption returns a boolean if a field has been set.

### GetAccountActions

`func (o *LifecycleStateAllOf) GetAccountActions() []AccountAction`

GetAccountActions returns the AccountActions field if non-nil, zero value otherwise.

### GetAccountActionsOk

`func (o *LifecycleStateAllOf) GetAccountActionsOk() (*[]AccountAction, bool)`

GetAccountActionsOk returns a tuple with the AccountActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountActions

`func (o *LifecycleStateAllOf) SetAccountActions(v []AccountAction)`

SetAccountActions sets AccountActions field to given value.

### HasAccountActions

`func (o *LifecycleStateAllOf) HasAccountActions() bool`

HasAccountActions returns a boolean if a field has been set.

### GetAccessProfileIds

`func (o *LifecycleStateAllOf) GetAccessProfileIds() []string`

GetAccessProfileIds returns the AccessProfileIds field if non-nil, zero value otherwise.

### GetAccessProfileIdsOk

`func (o *LifecycleStateAllOf) GetAccessProfileIdsOk() (*[]string, bool)`

GetAccessProfileIdsOk returns a tuple with the AccessProfileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessProfileIds

`func (o *LifecycleStateAllOf) SetAccessProfileIds(v []string)`

SetAccessProfileIds sets AccessProfileIds field to given value.

### HasAccessProfileIds

`func (o *LifecycleStateAllOf) HasAccessProfileIds() bool`

HasAccessProfileIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


