/*
Identity Security Cloud Beta API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_beta

import (
	"encoding/json"
)

// checks if the SelectorAccountMatchConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SelectorAccountMatchConfig{}

// SelectorAccountMatchConfig struct for SelectorAccountMatchConfig
type SelectorAccountMatchConfig struct {
	MatchExpression *SelectorAccountMatchConfigMatchExpression `json:"matchExpression,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SelectorAccountMatchConfig SelectorAccountMatchConfig

// NewSelectorAccountMatchConfig instantiates a new SelectorAccountMatchConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSelectorAccountMatchConfig() *SelectorAccountMatchConfig {
	this := SelectorAccountMatchConfig{}
	return &this
}

// NewSelectorAccountMatchConfigWithDefaults instantiates a new SelectorAccountMatchConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSelectorAccountMatchConfigWithDefaults() *SelectorAccountMatchConfig {
	this := SelectorAccountMatchConfig{}
	return &this
}

// GetMatchExpression returns the MatchExpression field value if set, zero value otherwise.
func (o *SelectorAccountMatchConfig) GetMatchExpression() SelectorAccountMatchConfigMatchExpression {
	if o == nil || IsNil(o.MatchExpression) {
		var ret SelectorAccountMatchConfigMatchExpression
		return ret
	}
	return *o.MatchExpression
}

// GetMatchExpressionOk returns a tuple with the MatchExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelectorAccountMatchConfig) GetMatchExpressionOk() (*SelectorAccountMatchConfigMatchExpression, bool) {
	if o == nil || IsNil(o.MatchExpression) {
		return nil, false
	}
	return o.MatchExpression, true
}

// HasMatchExpression returns a boolean if a field has been set.
func (o *SelectorAccountMatchConfig) HasMatchExpression() bool {
	if o != nil && !IsNil(o.MatchExpression) {
		return true
	}

	return false
}

// SetMatchExpression gets a reference to the given SelectorAccountMatchConfigMatchExpression and assigns it to the MatchExpression field.
func (o *SelectorAccountMatchConfig) SetMatchExpression(v SelectorAccountMatchConfigMatchExpression) {
	o.MatchExpression = &v
}

func (o SelectorAccountMatchConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SelectorAccountMatchConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MatchExpression) {
		toSerialize["matchExpression"] = o.MatchExpression
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SelectorAccountMatchConfig) UnmarshalJSON(data []byte) (err error) {
	varSelectorAccountMatchConfig := _SelectorAccountMatchConfig{}

	err = json.Unmarshal(data, &varSelectorAccountMatchConfig)

	if err != nil {
		return err
	}

	*o = SelectorAccountMatchConfig(varSelectorAccountMatchConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "matchExpression")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSelectorAccountMatchConfig struct {
	value *SelectorAccountMatchConfig
	isSet bool
}

func (v NullableSelectorAccountMatchConfig) Get() *SelectorAccountMatchConfig {
	return v.value
}

func (v *NullableSelectorAccountMatchConfig) Set(val *SelectorAccountMatchConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableSelectorAccountMatchConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableSelectorAccountMatchConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelectorAccountMatchConfig(val *SelectorAccountMatchConfig) *NullableSelectorAccountMatchConfig {
	return &NullableSelectorAccountMatchConfig{value: val, isSet: true}
}

func (v NullableSelectorAccountMatchConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelectorAccountMatchConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

