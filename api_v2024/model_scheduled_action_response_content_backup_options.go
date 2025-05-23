/*
Identity Security Cloud V2024 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: v2024
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2024

import (
	"encoding/json"
)

// checks if the ScheduledActionResponseContentBackupOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScheduledActionResponseContentBackupOptions{}

// ScheduledActionResponseContentBackupOptions Options for BACKUP type jobs. Optional, applicable for BACKUP jobs only.
type ScheduledActionResponseContentBackupOptions struct {
	// Object types that are to be included in the backup.
	IncludeTypes []string `json:"includeTypes,omitempty"`
	// Map of objectType string to the options to be passed to the target service for that objectType.
	ObjectOptions *map[string]ScheduledActionResponseContentBackupOptionsObjectOptionsValue `json:"objectOptions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ScheduledActionResponseContentBackupOptions ScheduledActionResponseContentBackupOptions

// NewScheduledActionResponseContentBackupOptions instantiates a new ScheduledActionResponseContentBackupOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScheduledActionResponseContentBackupOptions() *ScheduledActionResponseContentBackupOptions {
	this := ScheduledActionResponseContentBackupOptions{}
	return &this
}

// NewScheduledActionResponseContentBackupOptionsWithDefaults instantiates a new ScheduledActionResponseContentBackupOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduledActionResponseContentBackupOptionsWithDefaults() *ScheduledActionResponseContentBackupOptions {
	this := ScheduledActionResponseContentBackupOptions{}
	return &this
}

// GetIncludeTypes returns the IncludeTypes field value if set, zero value otherwise.
func (o *ScheduledActionResponseContentBackupOptions) GetIncludeTypes() []string {
	if o == nil || IsNil(o.IncludeTypes) {
		var ret []string
		return ret
	}
	return o.IncludeTypes
}

// GetIncludeTypesOk returns a tuple with the IncludeTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledActionResponseContentBackupOptions) GetIncludeTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.IncludeTypes) {
		return nil, false
	}
	return o.IncludeTypes, true
}

// HasIncludeTypes returns a boolean if a field has been set.
func (o *ScheduledActionResponseContentBackupOptions) HasIncludeTypes() bool {
	if o != nil && !IsNil(o.IncludeTypes) {
		return true
	}

	return false
}

// SetIncludeTypes gets a reference to the given []string and assigns it to the IncludeTypes field.
func (o *ScheduledActionResponseContentBackupOptions) SetIncludeTypes(v []string) {
	o.IncludeTypes = v
}

// GetObjectOptions returns the ObjectOptions field value if set, zero value otherwise.
func (o *ScheduledActionResponseContentBackupOptions) GetObjectOptions() map[string]ScheduledActionResponseContentBackupOptionsObjectOptionsValue {
	if o == nil || IsNil(o.ObjectOptions) {
		var ret map[string]ScheduledActionResponseContentBackupOptionsObjectOptionsValue
		return ret
	}
	return *o.ObjectOptions
}

// GetObjectOptionsOk returns a tuple with the ObjectOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledActionResponseContentBackupOptions) GetObjectOptionsOk() (*map[string]ScheduledActionResponseContentBackupOptionsObjectOptionsValue, bool) {
	if o == nil || IsNil(o.ObjectOptions) {
		return nil, false
	}
	return o.ObjectOptions, true
}

// HasObjectOptions returns a boolean if a field has been set.
func (o *ScheduledActionResponseContentBackupOptions) HasObjectOptions() bool {
	if o != nil && !IsNil(o.ObjectOptions) {
		return true
	}

	return false
}

// SetObjectOptions gets a reference to the given map[string]ScheduledActionResponseContentBackupOptionsObjectOptionsValue and assigns it to the ObjectOptions field.
func (o *ScheduledActionResponseContentBackupOptions) SetObjectOptions(v map[string]ScheduledActionResponseContentBackupOptionsObjectOptionsValue) {
	o.ObjectOptions = &v
}

func (o ScheduledActionResponseContentBackupOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScheduledActionResponseContentBackupOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IncludeTypes) {
		toSerialize["includeTypes"] = o.IncludeTypes
	}
	if !IsNil(o.ObjectOptions) {
		toSerialize["objectOptions"] = o.ObjectOptions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ScheduledActionResponseContentBackupOptions) UnmarshalJSON(data []byte) (err error) {
	varScheduledActionResponseContentBackupOptions := _ScheduledActionResponseContentBackupOptions{}

	err = json.Unmarshal(data, &varScheduledActionResponseContentBackupOptions)

	if err != nil {
		return err
	}

	*o = ScheduledActionResponseContentBackupOptions(varScheduledActionResponseContentBackupOptions)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "includeTypes")
		delete(additionalProperties, "objectOptions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableScheduledActionResponseContentBackupOptions struct {
	value *ScheduledActionResponseContentBackupOptions
	isSet bool
}

func (v NullableScheduledActionResponseContentBackupOptions) Get() *ScheduledActionResponseContentBackupOptions {
	return v.value
}

func (v *NullableScheduledActionResponseContentBackupOptions) Set(val *ScheduledActionResponseContentBackupOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduledActionResponseContentBackupOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduledActionResponseContentBackupOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduledActionResponseContentBackupOptions(val *ScheduledActionResponseContentBackupOptions) *NullableScheduledActionResponseContentBackupOptions {
	return &NullableScheduledActionResponseContentBackupOptions{value: val, isSet: true}
}

func (v NullableScheduledActionResponseContentBackupOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduledActionResponseContentBackupOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


