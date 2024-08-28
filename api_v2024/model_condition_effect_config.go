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

// checks if the ConditionEffectConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConditionEffectConfig{}

// ConditionEffectConfig Arbitrary map containing a configuration based on the EffectType.
type ConditionEffectConfig struct {
	// Effect type's label.
	DefaultValueLabel *string `json:"defaultValueLabel,omitempty"`
	// Element's identifier.
	Element *string `json:"element,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConditionEffectConfig ConditionEffectConfig

// NewConditionEffectConfig instantiates a new ConditionEffectConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConditionEffectConfig() *ConditionEffectConfig {
	this := ConditionEffectConfig{}
	return &this
}

// NewConditionEffectConfigWithDefaults instantiates a new ConditionEffectConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConditionEffectConfigWithDefaults() *ConditionEffectConfig {
	this := ConditionEffectConfig{}
	return &this
}

// GetDefaultValueLabel returns the DefaultValueLabel field value if set, zero value otherwise.
func (o *ConditionEffectConfig) GetDefaultValueLabel() string {
	if o == nil || IsNil(o.DefaultValueLabel) {
		var ret string
		return ret
	}
	return *o.DefaultValueLabel
}

// GetDefaultValueLabelOk returns a tuple with the DefaultValueLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionEffectConfig) GetDefaultValueLabelOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultValueLabel) {
		return nil, false
	}
	return o.DefaultValueLabel, true
}

// HasDefaultValueLabel returns a boolean if a field has been set.
func (o *ConditionEffectConfig) HasDefaultValueLabel() bool {
	if o != nil && !IsNil(o.DefaultValueLabel) {
		return true
	}

	return false
}

// SetDefaultValueLabel gets a reference to the given string and assigns it to the DefaultValueLabel field.
func (o *ConditionEffectConfig) SetDefaultValueLabel(v string) {
	o.DefaultValueLabel = &v
}

// GetElement returns the Element field value if set, zero value otherwise.
func (o *ConditionEffectConfig) GetElement() string {
	if o == nil || IsNil(o.Element) {
		var ret string
		return ret
	}
	return *o.Element
}

// GetElementOk returns a tuple with the Element field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionEffectConfig) GetElementOk() (*string, bool) {
	if o == nil || IsNil(o.Element) {
		return nil, false
	}
	return o.Element, true
}

// HasElement returns a boolean if a field has been set.
func (o *ConditionEffectConfig) HasElement() bool {
	if o != nil && !IsNil(o.Element) {
		return true
	}

	return false
}

// SetElement gets a reference to the given string and assigns it to the Element field.
func (o *ConditionEffectConfig) SetElement(v string) {
	o.Element = &v
}

func (o ConditionEffectConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConditionEffectConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DefaultValueLabel) {
		toSerialize["defaultValueLabel"] = o.DefaultValueLabel
	}
	if !IsNil(o.Element) {
		toSerialize["element"] = o.Element
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConditionEffectConfig) UnmarshalJSON(data []byte) (err error) {
	varConditionEffectConfig := _ConditionEffectConfig{}

	err = json.Unmarshal(data, &varConditionEffectConfig)

	if err != nil {
		return err
	}

	*o = ConditionEffectConfig(varConditionEffectConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "defaultValueLabel")
		delete(additionalProperties, "element")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConditionEffectConfig struct {
	value *ConditionEffectConfig
	isSet bool
}

func (v NullableConditionEffectConfig) Get() *ConditionEffectConfig {
	return v.value
}

func (v *NullableConditionEffectConfig) Set(val *ConditionEffectConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableConditionEffectConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableConditionEffectConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConditionEffectConfig(val *ConditionEffectConfig) *NullableConditionEffectConfig {
	return &NullableConditionEffectConfig{value: val, isSet: true}
}

func (v NullableConditionEffectConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConditionEffectConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

