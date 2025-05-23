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

// checks if the ConditionEffect type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConditionEffect{}

// ConditionEffect Effect produced by a condition.
type ConditionEffect struct {
	// Type of effect to perform when the conditions are evaluated for this logic block. HIDE ConditionEffectTypeHide  Disables validations. SHOW ConditionEffectTypeShow  Enables validations. DISABLE ConditionEffectTypeDisable  Disables validations. ENABLE ConditionEffectTypeEnable  Enables validations. REQUIRE ConditionEffectTypeRequire OPTIONAL ConditionEffectTypeOptional SUBMIT_MESSAGE ConditionEffectTypeSubmitMessage SUBMIT_NOTIFICATION ConditionEffectTypeSubmitNotification SET_DEFAULT_VALUE ConditionEffectTypeSetDefaultValue  This value is ignored on purpose.
	EffectType *string `json:"effectType,omitempty"`
	Config *ConditionEffectConfig `json:"config,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConditionEffect ConditionEffect

// NewConditionEffect instantiates a new ConditionEffect object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConditionEffect() *ConditionEffect {
	this := ConditionEffect{}
	return &this
}

// NewConditionEffectWithDefaults instantiates a new ConditionEffect object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConditionEffectWithDefaults() *ConditionEffect {
	this := ConditionEffect{}
	return &this
}

// GetEffectType returns the EffectType field value if set, zero value otherwise.
func (o *ConditionEffect) GetEffectType() string {
	if o == nil || IsNil(o.EffectType) {
		var ret string
		return ret
	}
	return *o.EffectType
}

// GetEffectTypeOk returns a tuple with the EffectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionEffect) GetEffectTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EffectType) {
		return nil, false
	}
	return o.EffectType, true
}

// HasEffectType returns a boolean if a field has been set.
func (o *ConditionEffect) HasEffectType() bool {
	if o != nil && !IsNil(o.EffectType) {
		return true
	}

	return false
}

// SetEffectType gets a reference to the given string and assigns it to the EffectType field.
func (o *ConditionEffect) SetEffectType(v string) {
	o.EffectType = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *ConditionEffect) GetConfig() ConditionEffectConfig {
	if o == nil || IsNil(o.Config) {
		var ret ConditionEffectConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionEffect) GetConfigOk() (*ConditionEffectConfig, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *ConditionEffect) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given ConditionEffectConfig and assigns it to the Config field.
func (o *ConditionEffect) SetConfig(v ConditionEffectConfig) {
	o.Config = &v
}

func (o ConditionEffect) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConditionEffect) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EffectType) {
		toSerialize["effectType"] = o.EffectType
	}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConditionEffect) UnmarshalJSON(data []byte) (err error) {
	varConditionEffect := _ConditionEffect{}

	err = json.Unmarshal(data, &varConditionEffect)

	if err != nil {
		return err
	}

	*o = ConditionEffect(varConditionEffect)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "effectType")
		delete(additionalProperties, "config")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConditionEffect struct {
	value *ConditionEffect
	isSet bool
}

func (v NullableConditionEffect) Get() *ConditionEffect {
	return v.value
}

func (v *NullableConditionEffect) Set(val *ConditionEffect) {
	v.value = val
	v.isSet = true
}

func (v NullableConditionEffect) IsSet() bool {
	return v.isSet
}

func (v *NullableConditionEffect) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConditionEffect(val *ConditionEffect) *NullableConditionEffect {
	return &NullableConditionEffect{value: val, isSet: true}
}

func (v NullableConditionEffect) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConditionEffect) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


