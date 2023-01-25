# GetOrgSettings200ResponseSystemNotificationConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Notifications** | Pointer to [**[]GetOrgSettings200ResponseSystemNotificationConfigNotificationsInner**](GetOrgSettings200ResponseSystemNotificationConfigNotificationsInner.md) |  | [optional] 
**RecipientType** | Pointer to **string** |  | [optional] 

## Methods

### NewGetOrgSettings200ResponseSystemNotificationConfig

`func NewGetOrgSettings200ResponseSystemNotificationConfig() *GetOrgSettings200ResponseSystemNotificationConfig`

NewGetOrgSettings200ResponseSystemNotificationConfig instantiates a new GetOrgSettings200ResponseSystemNotificationConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrgSettings200ResponseSystemNotificationConfigWithDefaults

`func NewGetOrgSettings200ResponseSystemNotificationConfigWithDefaults() *GetOrgSettings200ResponseSystemNotificationConfig`

NewGetOrgSettings200ResponseSystemNotificationConfigWithDefaults instantiates a new GetOrgSettings200ResponseSystemNotificationConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotifications

`func (o *GetOrgSettings200ResponseSystemNotificationConfig) GetNotifications() []GetOrgSettings200ResponseSystemNotificationConfigNotificationsInner`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *GetOrgSettings200ResponseSystemNotificationConfig) GetNotificationsOk() (*[]GetOrgSettings200ResponseSystemNotificationConfigNotificationsInner, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *GetOrgSettings200ResponseSystemNotificationConfig) SetNotifications(v []GetOrgSettings200ResponseSystemNotificationConfigNotificationsInner)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *GetOrgSettings200ResponseSystemNotificationConfig) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetRecipientType

`func (o *GetOrgSettings200ResponseSystemNotificationConfig) GetRecipientType() string`

GetRecipientType returns the RecipientType field if non-nil, zero value otherwise.

### GetRecipientTypeOk

`func (o *GetOrgSettings200ResponseSystemNotificationConfig) GetRecipientTypeOk() (*string, bool)`

GetRecipientTypeOk returns a tuple with the RecipientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientType

`func (o *GetOrgSettings200ResponseSystemNotificationConfig) SetRecipientType(v string)`

SetRecipientType sets RecipientType field to given value.

### HasRecipientType

`func (o *GetOrgSettings200ResponseSystemNotificationConfig) HasRecipientType() bool`

HasRecipientType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


