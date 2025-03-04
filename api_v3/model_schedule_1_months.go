/*
Identity Security Cloud V3 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v3

import (
	"encoding/json"
	"fmt"
)

// checks if the Schedule1Months type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Schedule1Months{}

// Schedule1Months struct for Schedule1Months
type Schedule1Months struct {
	Type SelectorType `json:"type"`
	// The selected values. 
	Values []string `json:"values"`
	// The selected interval for RANGE selectors. 
	Interval NullableInt32 `json:"interval,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Schedule1Months Schedule1Months

// NewSchedule1Months instantiates a new Schedule1Months object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchedule1Months(type_ SelectorType, values []string) *Schedule1Months {
	this := Schedule1Months{}
	this.Type = type_
	this.Values = values
	return &this
}

// NewSchedule1MonthsWithDefaults instantiates a new Schedule1Months object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchedule1MonthsWithDefaults() *Schedule1Months {
	this := Schedule1Months{}
	return &this
}

// GetType returns the Type field value
func (o *Schedule1Months) GetType() SelectorType {
	if o == nil {
		var ret SelectorType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Schedule1Months) GetTypeOk() (*SelectorType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Schedule1Months) SetType(v SelectorType) {
	o.Type = v
}

// GetValues returns the Values field value
func (o *Schedule1Months) GetValues() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *Schedule1Months) GetValuesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *Schedule1Months) SetValues(v []string) {
	o.Values = v
}

// GetInterval returns the Interval field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Schedule1Months) GetInterval() int32 {
	if o == nil || IsNil(o.Interval.Get()) {
		var ret int32
		return ret
	}
	return *o.Interval.Get()
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Schedule1Months) GetIntervalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Interval.Get(), o.Interval.IsSet()
}

// HasInterval returns a boolean if a field has been set.
func (o *Schedule1Months) HasInterval() bool {
	if o != nil && o.Interval.IsSet() {
		return true
	}

	return false
}

// SetInterval gets a reference to the given NullableInt32 and assigns it to the Interval field.
func (o *Schedule1Months) SetInterval(v int32) {
	o.Interval.Set(&v)
}
// SetIntervalNil sets the value for Interval to be an explicit nil
func (o *Schedule1Months) SetIntervalNil() {
	o.Interval.Set(nil)
}

// UnsetInterval ensures that no value is present for Interval, not even an explicit nil
func (o *Schedule1Months) UnsetInterval() {
	o.Interval.Unset()
}

func (o Schedule1Months) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Schedule1Months) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["values"] = o.Values
	if o.Interval.IsSet() {
		toSerialize["interval"] = o.Interval.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Schedule1Months) UnmarshalJSON(data []byte) (err error) {
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

	varSchedule1Months := _Schedule1Months{}

	err = json.Unmarshal(data, &varSchedule1Months)

	if err != nil {
		return err
	}

	*o = Schedule1Months(varSchedule1Months)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "values")
		delete(additionalProperties, "interval")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSchedule1Months struct {
	value *Schedule1Months
	isSet bool
}

func (v NullableSchedule1Months) Get() *Schedule1Months {
	return v.value
}

func (v *NullableSchedule1Months) Set(val *Schedule1Months) {
	v.value = val
	v.isSet = true
}

func (v NullableSchedule1Months) IsSet() bool {
	return v.isSet
}

func (v *NullableSchedule1Months) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchedule1Months(val *Schedule1Months) *NullableSchedule1Months {
	return &NullableSchedule1Months{value: val, isSet: true}
}

func (v NullableSchedule1Months) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchedule1Months) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


