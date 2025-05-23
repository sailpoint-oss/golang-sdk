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

// checks if the LoadAccountsTask type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoadAccountsTask{}

// LoadAccountsTask struct for LoadAccountsTask
type LoadAccountsTask struct {
	// The status of the result
	Success *bool `json:"success,omitempty"`
	Task *LoadAccountsTaskTask `json:"task,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LoadAccountsTask LoadAccountsTask

// NewLoadAccountsTask instantiates a new LoadAccountsTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadAccountsTask() *LoadAccountsTask {
	this := LoadAccountsTask{}
	var success bool = true
	this.Success = &success
	return &this
}

// NewLoadAccountsTaskWithDefaults instantiates a new LoadAccountsTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadAccountsTaskWithDefaults() *LoadAccountsTask {
	this := LoadAccountsTask{}
	var success bool = true
	this.Success = &success
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *LoadAccountsTask) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadAccountsTask) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *LoadAccountsTask) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *LoadAccountsTask) SetSuccess(v bool) {
	o.Success = &v
}

// GetTask returns the Task field value if set, zero value otherwise.
func (o *LoadAccountsTask) GetTask() LoadAccountsTaskTask {
	if o == nil || IsNil(o.Task) {
		var ret LoadAccountsTaskTask
		return ret
	}
	return *o.Task
}

// GetTaskOk returns a tuple with the Task field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadAccountsTask) GetTaskOk() (*LoadAccountsTaskTask, bool) {
	if o == nil || IsNil(o.Task) {
		return nil, false
	}
	return o.Task, true
}

// HasTask returns a boolean if a field has been set.
func (o *LoadAccountsTask) HasTask() bool {
	if o != nil && !IsNil(o.Task) {
		return true
	}

	return false
}

// SetTask gets a reference to the given LoadAccountsTaskTask and assigns it to the Task field.
func (o *LoadAccountsTask) SetTask(v LoadAccountsTaskTask) {
	o.Task = &v
}

func (o LoadAccountsTask) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoadAccountsTask) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Task) {
		toSerialize["task"] = o.Task
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LoadAccountsTask) UnmarshalJSON(data []byte) (err error) {
	varLoadAccountsTask := _LoadAccountsTask{}

	err = json.Unmarshal(data, &varLoadAccountsTask)

	if err != nil {
		return err
	}

	*o = LoadAccountsTask(varLoadAccountsTask)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "success")
		delete(additionalProperties, "task")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLoadAccountsTask struct {
	value *LoadAccountsTask
	isSet bool
}

func (v NullableLoadAccountsTask) Get() *LoadAccountsTask {
	return v.value
}

func (v *NullableLoadAccountsTask) Set(val *LoadAccountsTask) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadAccountsTask) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadAccountsTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadAccountsTask(val *LoadAccountsTask) *NullableLoadAccountsTask {
	return &NullableLoadAccountsTask{value: val, isSet: true}
}

func (v NullableLoadAccountsTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadAccountsTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


