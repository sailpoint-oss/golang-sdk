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

// checks if the Entitlement1ManuallyUpdatedFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Entitlement1ManuallyUpdatedFields{}

// Entitlement1ManuallyUpdatedFields struct for Entitlement1ManuallyUpdatedFields
type Entitlement1ManuallyUpdatedFields struct {
	// True if the entitlements name was updated manually via entitlement import csv or patch endpoint.  False means that property value has not been change after first entitlement aggregation. Field refers to [Entitlement response schema](https://developer.sailpoint.com/idn/api/beta/get-entitlement) > `name` property.
	DISPLAY_NAME *bool `json:"DISPLAY_NAME,omitempty"`
	// True if the entitlement description was updated manually via entitlement import csv or patch endpoint.  False means that property value has not been change after first entitlement aggregation. Field refers to [Entitlement response schema](https://developer.sailpoint.com/idn/api/beta/get-entitlement) > `description` property.
	DESCRIPTION *bool `json:"DESCRIPTION,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Entitlement1ManuallyUpdatedFields Entitlement1ManuallyUpdatedFields

// NewEntitlement1ManuallyUpdatedFields instantiates a new Entitlement1ManuallyUpdatedFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlement1ManuallyUpdatedFields() *Entitlement1ManuallyUpdatedFields {
	this := Entitlement1ManuallyUpdatedFields{}
	var dISPLAYNAME bool = false
	this.DISPLAY_NAME = &dISPLAYNAME
	var dESCRIPTION bool = false
	this.DESCRIPTION = &dESCRIPTION
	return &this
}

// NewEntitlement1ManuallyUpdatedFieldsWithDefaults instantiates a new Entitlement1ManuallyUpdatedFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlement1ManuallyUpdatedFieldsWithDefaults() *Entitlement1ManuallyUpdatedFields {
	this := Entitlement1ManuallyUpdatedFields{}
	var dISPLAYNAME bool = false
	this.DISPLAY_NAME = &dISPLAYNAME
	var dESCRIPTION bool = false
	this.DESCRIPTION = &dESCRIPTION
	return &this
}

// GetDISPLAY_NAME returns the DISPLAY_NAME field value if set, zero value otherwise.
func (o *Entitlement1ManuallyUpdatedFields) GetDISPLAY_NAME() bool {
	if o == nil || IsNil(o.DISPLAY_NAME) {
		var ret bool
		return ret
	}
	return *o.DISPLAY_NAME
}

// GetDISPLAY_NAMEOk returns a tuple with the DISPLAY_NAME field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Entitlement1ManuallyUpdatedFields) GetDISPLAY_NAMEOk() (*bool, bool) {
	if o == nil || IsNil(o.DISPLAY_NAME) {
		return nil, false
	}
	return o.DISPLAY_NAME, true
}

// HasDISPLAY_NAME returns a boolean if a field has been set.
func (o *Entitlement1ManuallyUpdatedFields) HasDISPLAY_NAME() bool {
	if o != nil && !IsNil(o.DISPLAY_NAME) {
		return true
	}

	return false
}

// SetDISPLAY_NAME gets a reference to the given bool and assigns it to the DISPLAY_NAME field.
func (o *Entitlement1ManuallyUpdatedFields) SetDISPLAY_NAME(v bool) {
	o.DISPLAY_NAME = &v
}

// GetDESCRIPTION returns the DESCRIPTION field value if set, zero value otherwise.
func (o *Entitlement1ManuallyUpdatedFields) GetDESCRIPTION() bool {
	if o == nil || IsNil(o.DESCRIPTION) {
		var ret bool
		return ret
	}
	return *o.DESCRIPTION
}

// GetDESCRIPTIONOk returns a tuple with the DESCRIPTION field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Entitlement1ManuallyUpdatedFields) GetDESCRIPTIONOk() (*bool, bool) {
	if o == nil || IsNil(o.DESCRIPTION) {
		return nil, false
	}
	return o.DESCRIPTION, true
}

// HasDESCRIPTION returns a boolean if a field has been set.
func (o *Entitlement1ManuallyUpdatedFields) HasDESCRIPTION() bool {
	if o != nil && !IsNil(o.DESCRIPTION) {
		return true
	}

	return false
}

// SetDESCRIPTION gets a reference to the given bool and assigns it to the DESCRIPTION field.
func (o *Entitlement1ManuallyUpdatedFields) SetDESCRIPTION(v bool) {
	o.DESCRIPTION = &v
}

func (o Entitlement1ManuallyUpdatedFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Entitlement1ManuallyUpdatedFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DISPLAY_NAME) {
		toSerialize["DISPLAY_NAME"] = o.DISPLAY_NAME
	}
	if !IsNil(o.DESCRIPTION) {
		toSerialize["DESCRIPTION"] = o.DESCRIPTION
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Entitlement1ManuallyUpdatedFields) UnmarshalJSON(data []byte) (err error) {
	varEntitlement1ManuallyUpdatedFields := _Entitlement1ManuallyUpdatedFields{}

	err = json.Unmarshal(data, &varEntitlement1ManuallyUpdatedFields)

	if err != nil {
		return err
	}

	*o = Entitlement1ManuallyUpdatedFields(varEntitlement1ManuallyUpdatedFields)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "DISPLAY_NAME")
		delete(additionalProperties, "DESCRIPTION")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntitlement1ManuallyUpdatedFields struct {
	value *Entitlement1ManuallyUpdatedFields
	isSet bool
}

func (v NullableEntitlement1ManuallyUpdatedFields) Get() *Entitlement1ManuallyUpdatedFields {
	return v.value
}

func (v *NullableEntitlement1ManuallyUpdatedFields) Set(val *Entitlement1ManuallyUpdatedFields) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlement1ManuallyUpdatedFields) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlement1ManuallyUpdatedFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlement1ManuallyUpdatedFields(val *Entitlement1ManuallyUpdatedFields) *NullableEntitlement1ManuallyUpdatedFields {
	return &NullableEntitlement1ManuallyUpdatedFields{value: val, isSet: true}
}

func (v NullableEntitlement1ManuallyUpdatedFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlement1ManuallyUpdatedFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


