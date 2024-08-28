/*
Identity Security Cloud V2024 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: v2024
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2024

import (
	"encoding/json"
	"fmt"
)

// checks if the AttrSyncSourceConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AttrSyncSourceConfig{}

// AttrSyncSourceConfig Specification of attribute sync configuration for a source
type AttrSyncSourceConfig struct {
	Source AttrSyncSource `json:"source"`
	// Attribute synchronization configuration for specific identity attributes in the context of a source
	Attributes []AttrSyncSourceAttributeConfig `json:"attributes"`
	AdditionalProperties map[string]interface{}
}

type _AttrSyncSourceConfig AttrSyncSourceConfig

// NewAttrSyncSourceConfig instantiates a new AttrSyncSourceConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttrSyncSourceConfig(source AttrSyncSource, attributes []AttrSyncSourceAttributeConfig) *AttrSyncSourceConfig {
	this := AttrSyncSourceConfig{}
	this.Source = source
	this.Attributes = attributes
	return &this
}

// NewAttrSyncSourceConfigWithDefaults instantiates a new AttrSyncSourceConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttrSyncSourceConfigWithDefaults() *AttrSyncSourceConfig {
	this := AttrSyncSourceConfig{}
	return &this
}

// GetSource returns the Source field value
func (o *AttrSyncSourceConfig) GetSource() AttrSyncSource {
	if o == nil {
		var ret AttrSyncSource
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *AttrSyncSourceConfig) GetSourceOk() (*AttrSyncSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *AttrSyncSourceConfig) SetSource(v AttrSyncSource) {
	o.Source = v
}

// GetAttributes returns the Attributes field value
func (o *AttrSyncSourceConfig) GetAttributes() []AttrSyncSourceAttributeConfig {
	if o == nil {
		var ret []AttrSyncSourceAttributeConfig
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AttrSyncSourceConfig) GetAttributesOk() ([]AttrSyncSourceAttributeConfig, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attributes, true
}

// SetAttributes sets field value
func (o *AttrSyncSourceConfig) SetAttributes(v []AttrSyncSourceAttributeConfig) {
	o.Attributes = v
}

func (o AttrSyncSourceConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttrSyncSourceConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["source"] = o.Source
	toSerialize["attributes"] = o.Attributes

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AttrSyncSourceConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"source",
		"attributes",
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

	varAttrSyncSourceConfig := _AttrSyncSourceConfig{}

	err = json.Unmarshal(data, &varAttrSyncSourceConfig)

	if err != nil {
		return err
	}

	*o = AttrSyncSourceConfig(varAttrSyncSourceConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "source")
		delete(additionalProperties, "attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAttrSyncSourceConfig struct {
	value *AttrSyncSourceConfig
	isSet bool
}

func (v NullableAttrSyncSourceConfig) Get() *AttrSyncSourceConfig {
	return v.value
}

func (v *NullableAttrSyncSourceConfig) Set(val *AttrSyncSourceConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableAttrSyncSourceConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableAttrSyncSourceConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttrSyncSourceConfig(val *AttrSyncSourceConfig) *NullableAttrSyncSourceConfig {
	return &NullableAttrSyncSourceConfig{value: val, isSet: true}
}

func (v NullableAttrSyncSourceConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttrSyncSourceConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

