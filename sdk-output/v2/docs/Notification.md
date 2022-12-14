# Notification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**ByEmail** | Pointer to **bool** |  | [optional] 
**Thresholds** | Pointer to [**NotificationThresholds**](NotificationThresholds.md) |  | [optional] 

## Methods

### NewNotification

`func NewNotification() *Notification`

NewNotification instantiates a new Notification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationWithDefaults

`func NewNotificationWithDefaults() *Notification`

NewNotificationWithDefaults instantiates a new Notification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Notification) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Notification) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Notification) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Notification) HasType() bool`

HasType returns a boolean if a field has been set.

### GetByEmail

`func (o *Notification) GetByEmail() bool`

GetByEmail returns the ByEmail field if non-nil, zero value otherwise.

### GetByEmailOk

`func (o *Notification) GetByEmailOk() (*bool, bool)`

GetByEmailOk returns a tuple with the ByEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByEmail

`func (o *Notification) SetByEmail(v bool)`

SetByEmail sets ByEmail field to given value.

### HasByEmail

`func (o *Notification) HasByEmail() bool`

HasByEmail returns a boolean if a field has been set.

### GetThresholds

`func (o *Notification) GetThresholds() NotificationThresholds`

GetThresholds returns the Thresholds field if non-nil, zero value otherwise.

### GetThresholdsOk

`func (o *Notification) GetThresholdsOk() (*NotificationThresholds, bool)`

GetThresholdsOk returns a tuple with the Thresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholds

`func (o *Notification) SetThresholds(v NotificationThresholds)`

SetThresholds sets Thresholds field to given value.

### HasThresholds

`func (o *Notification) HasThresholds() bool`

HasThresholds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


