/*
Identity Security Cloud V2025 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: v2025
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2025

import (
	"encoding/json"
)

// checks if the ApprovalConfigReminderConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApprovalConfigReminderConfig{}

// ApprovalConfigReminderConfig Configuration for reminders.
type ApprovalConfigReminderConfig struct {
	// Indicates if reminders are enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// Number of days until the first reminder.
	DaysUntilFirstReminder *int64 `json:"daysUntilFirstReminder,omitempty"`
	// Cron schedule for reminders.
	ReminderCronSchedule *string `json:"reminderCronSchedule,omitempty"`
	// Maximum number of reminders. Max is 20.
	MaxReminders *int64 `json:"maxReminders,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApprovalConfigReminderConfig ApprovalConfigReminderConfig

// NewApprovalConfigReminderConfig instantiates a new ApprovalConfigReminderConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApprovalConfigReminderConfig() *ApprovalConfigReminderConfig {
	this := ApprovalConfigReminderConfig{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// NewApprovalConfigReminderConfigWithDefaults instantiates a new ApprovalConfigReminderConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApprovalConfigReminderConfigWithDefaults() *ApprovalConfigReminderConfig {
	this := ApprovalConfigReminderConfig{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ApprovalConfigReminderConfig) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalConfigReminderConfig) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ApprovalConfigReminderConfig) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ApprovalConfigReminderConfig) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetDaysUntilFirstReminder returns the DaysUntilFirstReminder field value if set, zero value otherwise.
func (o *ApprovalConfigReminderConfig) GetDaysUntilFirstReminder() int64 {
	if o == nil || IsNil(o.DaysUntilFirstReminder) {
		var ret int64
		return ret
	}
	return *o.DaysUntilFirstReminder
}

// GetDaysUntilFirstReminderOk returns a tuple with the DaysUntilFirstReminder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalConfigReminderConfig) GetDaysUntilFirstReminderOk() (*int64, bool) {
	if o == nil || IsNil(o.DaysUntilFirstReminder) {
		return nil, false
	}
	return o.DaysUntilFirstReminder, true
}

// HasDaysUntilFirstReminder returns a boolean if a field has been set.
func (o *ApprovalConfigReminderConfig) HasDaysUntilFirstReminder() bool {
	if o != nil && !IsNil(o.DaysUntilFirstReminder) {
		return true
	}

	return false
}

// SetDaysUntilFirstReminder gets a reference to the given int64 and assigns it to the DaysUntilFirstReminder field.
func (o *ApprovalConfigReminderConfig) SetDaysUntilFirstReminder(v int64) {
	o.DaysUntilFirstReminder = &v
}

// GetReminderCronSchedule returns the ReminderCronSchedule field value if set, zero value otherwise.
func (o *ApprovalConfigReminderConfig) GetReminderCronSchedule() string {
	if o == nil || IsNil(o.ReminderCronSchedule) {
		var ret string
		return ret
	}
	return *o.ReminderCronSchedule
}

// GetReminderCronScheduleOk returns a tuple with the ReminderCronSchedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalConfigReminderConfig) GetReminderCronScheduleOk() (*string, bool) {
	if o == nil || IsNil(o.ReminderCronSchedule) {
		return nil, false
	}
	return o.ReminderCronSchedule, true
}

// HasReminderCronSchedule returns a boolean if a field has been set.
func (o *ApprovalConfigReminderConfig) HasReminderCronSchedule() bool {
	if o != nil && !IsNil(o.ReminderCronSchedule) {
		return true
	}

	return false
}

// SetReminderCronSchedule gets a reference to the given string and assigns it to the ReminderCronSchedule field.
func (o *ApprovalConfigReminderConfig) SetReminderCronSchedule(v string) {
	o.ReminderCronSchedule = &v
}

// GetMaxReminders returns the MaxReminders field value if set, zero value otherwise.
func (o *ApprovalConfigReminderConfig) GetMaxReminders() int64 {
	if o == nil || IsNil(o.MaxReminders) {
		var ret int64
		return ret
	}
	return *o.MaxReminders
}

// GetMaxRemindersOk returns a tuple with the MaxReminders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalConfigReminderConfig) GetMaxRemindersOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxReminders) {
		return nil, false
	}
	return o.MaxReminders, true
}

// HasMaxReminders returns a boolean if a field has been set.
func (o *ApprovalConfigReminderConfig) HasMaxReminders() bool {
	if o != nil && !IsNil(o.MaxReminders) {
		return true
	}

	return false
}

// SetMaxReminders gets a reference to the given int64 and assigns it to the MaxReminders field.
func (o *ApprovalConfigReminderConfig) SetMaxReminders(v int64) {
	o.MaxReminders = &v
}

func (o ApprovalConfigReminderConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApprovalConfigReminderConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.DaysUntilFirstReminder) {
		toSerialize["daysUntilFirstReminder"] = o.DaysUntilFirstReminder
	}
	if !IsNil(o.ReminderCronSchedule) {
		toSerialize["reminderCronSchedule"] = o.ReminderCronSchedule
	}
	if !IsNil(o.MaxReminders) {
		toSerialize["maxReminders"] = o.MaxReminders
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApprovalConfigReminderConfig) UnmarshalJSON(data []byte) (err error) {
	varApprovalConfigReminderConfig := _ApprovalConfigReminderConfig{}

	err = json.Unmarshal(data, &varApprovalConfigReminderConfig)

	if err != nil {
		return err
	}

	*o = ApprovalConfigReminderConfig(varApprovalConfigReminderConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "daysUntilFirstReminder")
		delete(additionalProperties, "reminderCronSchedule")
		delete(additionalProperties, "maxReminders")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApprovalConfigReminderConfig struct {
	value *ApprovalConfigReminderConfig
	isSet bool
}

func (v NullableApprovalConfigReminderConfig) Get() *ApprovalConfigReminderConfig {
	return v.value
}

func (v *NullableApprovalConfigReminderConfig) Set(val *ApprovalConfigReminderConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableApprovalConfigReminderConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableApprovalConfigReminderConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApprovalConfigReminderConfig(val *ApprovalConfigReminderConfig) *NullableApprovalConfigReminderConfig {
	return &NullableApprovalConfigReminderConfig{value: val, isSet: true}
}

func (v NullableApprovalConfigReminderConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApprovalConfigReminderConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


