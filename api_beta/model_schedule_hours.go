/*
Identity Security Cloud Beta API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_beta

import (
	"encoding/json"
	"fmt"
)

// checks if the ScheduleHours type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScheduleHours{}

// ScheduleHours Specifies which hour(s) a schedule is active for. Examples:  Every three hours starting from 8AM, inclusive: * type LIST * values \"8\" * interval 3  During business hours: * type RANGE * values \"9\", \"5\"  At 5AM, noon, and 5PM: * type LIST * values \"5\", \"12\", \"17\" 
type ScheduleHours struct {
	// Enum type to specify hours value
	Type string `json:"type"`
	// Values of the days based on the enum type mentioned above
	Values []string `json:"values"`
	// Interval between the cert generations
	Interval *int64 `json:"interval,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ScheduleHours ScheduleHours

// NewScheduleHours instantiates a new ScheduleHours object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScheduleHours(type_ string, values []string) *ScheduleHours {
	this := ScheduleHours{}
	this.Type = type_
	this.Values = values
	return &this
}

// NewScheduleHoursWithDefaults instantiates a new ScheduleHours object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduleHoursWithDefaults() *ScheduleHours {
	this := ScheduleHours{}
	return &this
}

// GetType returns the Type field value
func (o *ScheduleHours) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ScheduleHours) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ScheduleHours) SetType(v string) {
	o.Type = v
}

// GetValues returns the Values field value
func (o *ScheduleHours) GetValues() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *ScheduleHours) GetValuesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *ScheduleHours) SetValues(v []string) {
	o.Values = v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *ScheduleHours) GetInterval() int64 {
	if o == nil || IsNil(o.Interval) {
		var ret int64
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleHours) GetIntervalOk() (*int64, bool) {
	if o == nil || IsNil(o.Interval) {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *ScheduleHours) HasInterval() bool {
	if o != nil && !IsNil(o.Interval) {
		return true
	}

	return false
}

// SetInterval gets a reference to the given int64 and assigns it to the Interval field.
func (o *ScheduleHours) SetInterval(v int64) {
	o.Interval = &v
}

func (o ScheduleHours) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScheduleHours) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["values"] = o.Values
	if !IsNil(o.Interval) {
		toSerialize["interval"] = o.Interval
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ScheduleHours) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"values",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varScheduleHours := _ScheduleHours{}

	err = json.Unmarshal(data, &varScheduleHours)

	if err != nil {
		return err
	}

	*o = ScheduleHours(varScheduleHours)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "values")
		delete(additionalProperties, "interval")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableScheduleHours struct {
	value *ScheduleHours
	isSet bool
}

func (v NullableScheduleHours) Get() *ScheduleHours {
	return v.value
}

func (v *NullableScheduleHours) Set(val *ScheduleHours) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduleHours) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduleHours) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduleHours(val *ScheduleHours) *NullableScheduleHours {
	return &NullableScheduleHours{value: val, isSet: true}
}

func (v NullableScheduleHours) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduleHours) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


