# ApprovalReminderAndEscalationConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DaysUntilEscalation** | Pointer to **int32** | Number of days to wait before the first reminder. If no reminders are configured, then this is the number of days to wait before escalation. | [optional] 
**DaysBetweenReminders** | Pointer to **int32** | Number of days to wait between reminder notifications. | [optional] 
**MaxReminders** | Pointer to **int32** | Maximum number of reminder notification to send to the reviewer before approval escalation. | [optional] 
**FallbackApproverRef** | Pointer to [**NullableIdentityReferenceWithNameAndEmail**](IdentityReferenceWithNameAndEmail.md) |  | [optional] 

## Methods

### NewApprovalReminderAndEscalationConfig

`func NewApprovalReminderAndEscalationConfig() *ApprovalReminderAndEscalationConfig`

NewApprovalReminderAndEscalationConfig instantiates a new ApprovalReminderAndEscalationConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalReminderAndEscalationConfigWithDefaults

`func NewApprovalReminderAndEscalationConfigWithDefaults() *ApprovalReminderAndEscalationConfig`

NewApprovalReminderAndEscalationConfigWithDefaults instantiates a new ApprovalReminderAndEscalationConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDaysUntilEscalation

`func (o *ApprovalReminderAndEscalationConfig) GetDaysUntilEscalation() int32`

GetDaysUntilEscalation returns the DaysUntilEscalation field if non-nil, zero value otherwise.

### GetDaysUntilEscalationOk

`func (o *ApprovalReminderAndEscalationConfig) GetDaysUntilEscalationOk() (*int32, bool)`

GetDaysUntilEscalationOk returns a tuple with the DaysUntilEscalation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysUntilEscalation

`func (o *ApprovalReminderAndEscalationConfig) SetDaysUntilEscalation(v int32)`

SetDaysUntilEscalation sets DaysUntilEscalation field to given value.

### HasDaysUntilEscalation

`func (o *ApprovalReminderAndEscalationConfig) HasDaysUntilEscalation() bool`

HasDaysUntilEscalation returns a boolean if a field has been set.

### GetDaysBetweenReminders

`func (o *ApprovalReminderAndEscalationConfig) GetDaysBetweenReminders() int32`

GetDaysBetweenReminders returns the DaysBetweenReminders field if non-nil, zero value otherwise.

### GetDaysBetweenRemindersOk

`func (o *ApprovalReminderAndEscalationConfig) GetDaysBetweenRemindersOk() (*int32, bool)`

GetDaysBetweenRemindersOk returns a tuple with the DaysBetweenReminders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysBetweenReminders

`func (o *ApprovalReminderAndEscalationConfig) SetDaysBetweenReminders(v int32)`

SetDaysBetweenReminders sets DaysBetweenReminders field to given value.

### HasDaysBetweenReminders

`func (o *ApprovalReminderAndEscalationConfig) HasDaysBetweenReminders() bool`

HasDaysBetweenReminders returns a boolean if a field has been set.

### GetMaxReminders

`func (o *ApprovalReminderAndEscalationConfig) GetMaxReminders() int32`

GetMaxReminders returns the MaxReminders field if non-nil, zero value otherwise.

### GetMaxRemindersOk

`func (o *ApprovalReminderAndEscalationConfig) GetMaxRemindersOk() (*int32, bool)`

GetMaxRemindersOk returns a tuple with the MaxReminders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxReminders

`func (o *ApprovalReminderAndEscalationConfig) SetMaxReminders(v int32)`

SetMaxReminders sets MaxReminders field to given value.

### HasMaxReminders

`func (o *ApprovalReminderAndEscalationConfig) HasMaxReminders() bool`

HasMaxReminders returns a boolean if a field has been set.

### GetFallbackApproverRef

`func (o *ApprovalReminderAndEscalationConfig) GetFallbackApproverRef() IdentityReferenceWithNameAndEmail`

GetFallbackApproverRef returns the FallbackApproverRef field if non-nil, zero value otherwise.

### GetFallbackApproverRefOk

`func (o *ApprovalReminderAndEscalationConfig) GetFallbackApproverRefOk() (*IdentityReferenceWithNameAndEmail, bool)`

GetFallbackApproverRefOk returns a tuple with the FallbackApproverRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackApproverRef

`func (o *ApprovalReminderAndEscalationConfig) SetFallbackApproverRef(v IdentityReferenceWithNameAndEmail)`

SetFallbackApproverRef sets FallbackApproverRef field to given value.

### HasFallbackApproverRef

`func (o *ApprovalReminderAndEscalationConfig) HasFallbackApproverRef() bool`

HasFallbackApproverRef returns a boolean if a field has been set.

### SetFallbackApproverRefNil

`func (o *ApprovalReminderAndEscalationConfig) SetFallbackApproverRefNil(b bool)`

 SetFallbackApproverRefNil sets the value for FallbackApproverRef to be an explicit nil

### UnsetFallbackApproverRef
`func (o *ApprovalReminderAndEscalationConfig) UnsetFallbackApproverRef()`

UnsetFallbackApproverRef ensures that no value is present for FallbackApproverRef, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


