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

// checks if the RecommenderCalculationsIdentityAttributesValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecommenderCalculationsIdentityAttributesValue{}

// RecommenderCalculationsIdentityAttributesValue struct for RecommenderCalculationsIdentityAttributesValue
type RecommenderCalculationsIdentityAttributesValue struct {
	Value *string `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RecommenderCalculationsIdentityAttributesValue RecommenderCalculationsIdentityAttributesValue

// NewRecommenderCalculationsIdentityAttributesValue instantiates a new RecommenderCalculationsIdentityAttributesValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecommenderCalculationsIdentityAttributesValue() *RecommenderCalculationsIdentityAttributesValue {
	this := RecommenderCalculationsIdentityAttributesValue{}
	return &this
}

// NewRecommenderCalculationsIdentityAttributesValueWithDefaults instantiates a new RecommenderCalculationsIdentityAttributesValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecommenderCalculationsIdentityAttributesValueWithDefaults() *RecommenderCalculationsIdentityAttributesValue {
	this := RecommenderCalculationsIdentityAttributesValue{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *RecommenderCalculationsIdentityAttributesValue) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommenderCalculationsIdentityAttributesValue) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *RecommenderCalculationsIdentityAttributesValue) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *RecommenderCalculationsIdentityAttributesValue) SetValue(v string) {
	o.Value = &v
}

func (o RecommenderCalculationsIdentityAttributesValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecommenderCalculationsIdentityAttributesValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RecommenderCalculationsIdentityAttributesValue) UnmarshalJSON(data []byte) (err error) {
	varRecommenderCalculationsIdentityAttributesValue := _RecommenderCalculationsIdentityAttributesValue{}

	err = json.Unmarshal(data, &varRecommenderCalculationsIdentityAttributesValue)

	if err != nil {
		return err
	}

	*o = RecommenderCalculationsIdentityAttributesValue(varRecommenderCalculationsIdentityAttributesValue)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRecommenderCalculationsIdentityAttributesValue struct {
	value *RecommenderCalculationsIdentityAttributesValue
	isSet bool
}

func (v NullableRecommenderCalculationsIdentityAttributesValue) Get() *RecommenderCalculationsIdentityAttributesValue {
	return v.value
}

func (v *NullableRecommenderCalculationsIdentityAttributesValue) Set(val *RecommenderCalculationsIdentityAttributesValue) {
	v.value = val
	v.isSet = true
}

func (v NullableRecommenderCalculationsIdentityAttributesValue) IsSet() bool {
	return v.isSet
}

func (v *NullableRecommenderCalculationsIdentityAttributesValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecommenderCalculationsIdentityAttributesValue(val *RecommenderCalculationsIdentityAttributesValue) *NullableRecommenderCalculationsIdentityAttributesValue {
	return &NullableRecommenderCalculationsIdentityAttributesValue{value: val, isSet: true}
}

func (v NullableRecommenderCalculationsIdentityAttributesValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecommenderCalculationsIdentityAttributesValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


