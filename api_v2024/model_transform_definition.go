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

// checks if the TransformDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransformDefinition{}

// TransformDefinition struct for TransformDefinition
type TransformDefinition struct {
	// The type of the transform definition.
	Type *string `json:"type,omitempty"`
	// Arbitrary key-value pairs to store any metadata for the object
	Attributes *map[string]TransformDefinitionAttributesValue `json:"attributes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TransformDefinition TransformDefinition

// NewTransformDefinition instantiates a new TransformDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransformDefinition() *TransformDefinition {
	this := TransformDefinition{}
	return &this
}

// NewTransformDefinitionWithDefaults instantiates a new TransformDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransformDefinitionWithDefaults() *TransformDefinition {
	this := TransformDefinition{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TransformDefinition) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransformDefinition) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TransformDefinition) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TransformDefinition) SetType(v string) {
	o.Type = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *TransformDefinition) GetAttributes() map[string]TransformDefinitionAttributesValue {
	if o == nil || IsNil(o.Attributes) {
		var ret map[string]TransformDefinitionAttributesValue
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransformDefinition) GetAttributesOk() (*map[string]TransformDefinitionAttributesValue, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *TransformDefinition) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]TransformDefinitionAttributesValue and assigns it to the Attributes field.
func (o *TransformDefinition) SetAttributes(v map[string]TransformDefinitionAttributesValue) {
	o.Attributes = &v
}

func (o TransformDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransformDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TransformDefinition) UnmarshalJSON(data []byte) (err error) {
	varTransformDefinition := _TransformDefinition{}

	err = json.Unmarshal(data, &varTransformDefinition)

	if err != nil {
		return err
	}

	*o = TransformDefinition(varTransformDefinition)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransformDefinition struct {
	value *TransformDefinition
	isSet bool
}

func (v NullableTransformDefinition) Get() *TransformDefinition {
	return v.value
}

func (v *NullableTransformDefinition) Set(val *TransformDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableTransformDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableTransformDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransformDefinition(val *TransformDefinition) *NullableTransformDefinition {
	return &NullableTransformDefinition{value: val, isSet: true}
}

func (v NullableTransformDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransformDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

