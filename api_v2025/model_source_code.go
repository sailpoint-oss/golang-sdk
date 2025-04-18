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

// checks if the SourceCode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SourceCode{}

// SourceCode SourceCode
type SourceCode struct {
	// the version of the code
	Version string `json:"version"`
	// The code
	Script string `json:"script"`
	AdditionalProperties map[string]interface{}
}

type _SourceCode SourceCode

// NewSourceCode instantiates a new SourceCode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceCode(version string, script string) *SourceCode {
	this := SourceCode{}
	this.Version = version
	this.Script = script
	return &this
}

// NewSourceCodeWithDefaults instantiates a new SourceCode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceCodeWithDefaults() *SourceCode {
	this := SourceCode{}
	return &this
}

// GetVersion returns the Version field value
func (o *SourceCode) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *SourceCode) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *SourceCode) SetVersion(v string) {
	o.Version = v
}

// GetScript returns the Script field value
func (o *SourceCode) GetScript() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Script
}

// GetScriptOk returns a tuple with the Script field value
// and a boolean to check if the value has been set.
func (o *SourceCode) GetScriptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Script, true
}

// SetScript sets field value
func (o *SourceCode) SetScript(v string) {
	o.Script = v
}

func (o SourceCode) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SourceCode) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["version"] = o.Version
	toSerialize["script"] = o.Script

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SourceCode) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"version",
		"script",
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

	varSourceCode := _SourceCode{}

	err = json.Unmarshal(data, &varSourceCode)

	if err != nil {
		return err
	}

	*o = SourceCode(varSourceCode)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "version")
		delete(additionalProperties, "script")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSourceCode struct {
	value *SourceCode
	isSet bool
}

func (v NullableSourceCode) Get() *SourceCode {
	return v.value
}

func (v *NullableSourceCode) Set(val *SourceCode) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceCode) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceCode(val *SourceCode) *NullableSourceCode {
	return &NullableSourceCode{value: val, isSet: true}
}

func (v NullableSourceCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


