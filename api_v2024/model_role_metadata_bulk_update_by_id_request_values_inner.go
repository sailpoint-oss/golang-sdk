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

// checks if the RoleMetadataBulkUpdateByIdRequestValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoleMetadataBulkUpdateByIdRequestValuesInner{}

// RoleMetadataBulkUpdateByIdRequestValuesInner struct for RoleMetadataBulkUpdateByIdRequestValuesInner
type RoleMetadataBulkUpdateByIdRequestValuesInner struct {
	// the key of metadata attribute
	Attribute string `json:"attribute"`
	// the values of attribute to be updated
	Values []string `json:"values"`
	AdditionalProperties map[string]interface{}
}

type _RoleMetadataBulkUpdateByIdRequestValuesInner RoleMetadataBulkUpdateByIdRequestValuesInner

// NewRoleMetadataBulkUpdateByIdRequestValuesInner instantiates a new RoleMetadataBulkUpdateByIdRequestValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleMetadataBulkUpdateByIdRequestValuesInner(attribute string, values []string) *RoleMetadataBulkUpdateByIdRequestValuesInner {
	this := RoleMetadataBulkUpdateByIdRequestValuesInner{}
	this.Attribute = attribute
	this.Values = values
	return &this
}

// NewRoleMetadataBulkUpdateByIdRequestValuesInnerWithDefaults instantiates a new RoleMetadataBulkUpdateByIdRequestValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleMetadataBulkUpdateByIdRequestValuesInnerWithDefaults() *RoleMetadataBulkUpdateByIdRequestValuesInner {
	this := RoleMetadataBulkUpdateByIdRequestValuesInner{}
	return &this
}

// GetAttribute returns the Attribute field value
func (o *RoleMetadataBulkUpdateByIdRequestValuesInner) GetAttribute() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Attribute
}

// GetAttributeOk returns a tuple with the Attribute field value
// and a boolean to check if the value has been set.
func (o *RoleMetadataBulkUpdateByIdRequestValuesInner) GetAttributeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attribute, true
}

// SetAttribute sets field value
func (o *RoleMetadataBulkUpdateByIdRequestValuesInner) SetAttribute(v string) {
	o.Attribute = v
}

// GetValues returns the Values field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *RoleMetadataBulkUpdateByIdRequestValuesInner) GetValues() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoleMetadataBulkUpdateByIdRequestValuesInner) GetValuesOk() ([]string, bool) {
	if o == nil || IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *RoleMetadataBulkUpdateByIdRequestValuesInner) SetValues(v []string) {
	o.Values = v
}

func (o RoleMetadataBulkUpdateByIdRequestValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoleMetadataBulkUpdateByIdRequestValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["attribute"] = o.Attribute
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RoleMetadataBulkUpdateByIdRequestValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"attribute",
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

	varRoleMetadataBulkUpdateByIdRequestValuesInner := _RoleMetadataBulkUpdateByIdRequestValuesInner{}

	err = json.Unmarshal(data, &varRoleMetadataBulkUpdateByIdRequestValuesInner)

	if err != nil {
		return err
	}

	*o = RoleMetadataBulkUpdateByIdRequestValuesInner(varRoleMetadataBulkUpdateByIdRequestValuesInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "attribute")
		delete(additionalProperties, "values")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRoleMetadataBulkUpdateByIdRequestValuesInner struct {
	value *RoleMetadataBulkUpdateByIdRequestValuesInner
	isSet bool
}

func (v NullableRoleMetadataBulkUpdateByIdRequestValuesInner) Get() *RoleMetadataBulkUpdateByIdRequestValuesInner {
	return v.value
}

func (v *NullableRoleMetadataBulkUpdateByIdRequestValuesInner) Set(val *RoleMetadataBulkUpdateByIdRequestValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleMetadataBulkUpdateByIdRequestValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleMetadataBulkUpdateByIdRequestValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleMetadataBulkUpdateByIdRequestValuesInner(val *RoleMetadataBulkUpdateByIdRequestValuesInner) *NullableRoleMetadataBulkUpdateByIdRequestValuesInner {
	return &NullableRoleMetadataBulkUpdateByIdRequestValuesInner{value: val, isSet: true}
}

func (v NullableRoleMetadataBulkUpdateByIdRequestValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleMetadataBulkUpdateByIdRequestValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

