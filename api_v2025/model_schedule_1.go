/*
Identity Security Cloud V2025 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: v2025
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2025

import (
	"encoding/json"
	"fmt"
)

// checks if the Schedule1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Schedule1{}

// Schedule1 struct for Schedule1
type Schedule1 struct {
	// The type of the Schedule.
	Type string `json:"type"`
	// The cron expression of the schedule.
	CronExpression string `json:"cronExpression"`
	AdditionalProperties map[string]interface{}
}

type _Schedule1 Schedule1

// NewSchedule1 instantiates a new Schedule1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchedule1(type_ string, cronExpression string) *Schedule1 {
	this := Schedule1{}
	this.Type = type_
	this.CronExpression = cronExpression
	return &this
}

// NewSchedule1WithDefaults instantiates a new Schedule1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchedule1WithDefaults() *Schedule1 {
	this := Schedule1{}
	return &this
}

// GetType returns the Type field value
func (o *Schedule1) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Schedule1) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Schedule1) SetType(v string) {
	o.Type = v
}

// GetCronExpression returns the CronExpression field value
func (o *Schedule1) GetCronExpression() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CronExpression
}

// GetCronExpressionOk returns a tuple with the CronExpression field value
// and a boolean to check if the value has been set.
func (o *Schedule1) GetCronExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CronExpression, true
}

// SetCronExpression sets field value
func (o *Schedule1) SetCronExpression(v string) {
	o.CronExpression = v
}

func (o Schedule1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Schedule1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["cronExpression"] = o.CronExpression

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Schedule1) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"cronExpression",
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

	varSchedule1 := _Schedule1{}

	err = json.Unmarshal(data, &varSchedule1)

	if err != nil {
		return err
	}

	*o = Schedule1(varSchedule1)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "cronExpression")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSchedule1 struct {
	value *Schedule1
	isSet bool
}

func (v NullableSchedule1) Get() *Schedule1 {
	return v.value
}

func (v *NullableSchedule1) Set(val *Schedule1) {
	v.value = val
	v.isSet = true
}

func (v NullableSchedule1) IsSet() bool {
	return v.isSet
}

func (v *NullableSchedule1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchedule1(val *Schedule1) *NullableSchedule1 {
	return &NullableSchedule1{value: val, isSet: true}
}

func (v NullableSchedule1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchedule1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


