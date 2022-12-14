# OrgSystemNotificationConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Notifications** | Pointer to [**[]Notification**](Notification.md) |  | [optional] 
**RecipientType** | Pointer to **string** |  | [optional] 

## Methods

### NewOrgSystemNotificationConfig

`func NewOrgSystemNotificationConfig() *OrgSystemNotificationConfig`

NewOrgSystemNotificationConfig instantiates a new OrgSystemNotificationConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgSystemNotificationConfigWithDefaults

`func NewOrgSystemNotificationConfigWithDefaults() *OrgSystemNotificationConfig`

NewOrgSystemNotificationConfigWithDefaults instantiates a new OrgSystemNotificationConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotifications

`func (o *OrgSystemNotificationConfig) GetNotifications() []Notification`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *OrgSystemNotificationConfig) GetNotificationsOk() (*[]Notification, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *OrgSystemNotificationConfig) SetNotifications(v []Notification)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *OrgSystemNotificationConfig) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetRecipientType

`func (o *OrgSystemNotificationConfig) GetRecipientType() string`

GetRecipientType returns the RecipientType field if non-nil, zero value otherwise.

### GetRecipientTypeOk

`func (o *OrgSystemNotificationConfig) GetRecipientTypeOk() (*string, bool)`

GetRecipientTypeOk returns a tuple with the RecipientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientType

`func (o *OrgSystemNotificationConfig) SetRecipientType(v string)`

SetRecipientType sets RecipientType field to given value.

### HasRecipientType

`func (o *OrgSystemNotificationConfig) HasRecipientType() bool`

HasRecipientType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


