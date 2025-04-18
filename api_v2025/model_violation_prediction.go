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

// checks if the ViolationPrediction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ViolationPrediction{}

// ViolationPrediction An object containing a listing of the SOD violation reasons detected by this check.
type ViolationPrediction struct {
	// List of Violation Contexts
	ViolationContexts []ViolationContext `json:"violationContexts,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ViolationPrediction ViolationPrediction

// NewViolationPrediction instantiates a new ViolationPrediction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViolationPrediction() *ViolationPrediction {
	this := ViolationPrediction{}
	return &this
}

// NewViolationPredictionWithDefaults instantiates a new ViolationPrediction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViolationPredictionWithDefaults() *ViolationPrediction {
	this := ViolationPrediction{}
	return &this
}

// GetViolationContexts returns the ViolationContexts field value if set, zero value otherwise.
func (o *ViolationPrediction) GetViolationContexts() []ViolationContext {
	if o == nil || IsNil(o.ViolationContexts) {
		var ret []ViolationContext
		return ret
	}
	return o.ViolationContexts
}

// GetViolationContextsOk returns a tuple with the ViolationContexts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViolationPrediction) GetViolationContextsOk() ([]ViolationContext, bool) {
	if o == nil || IsNil(o.ViolationContexts) {
		return nil, false
	}
	return o.ViolationContexts, true
}

// HasViolationContexts returns a boolean if a field has been set.
func (o *ViolationPrediction) HasViolationContexts() bool {
	if o != nil && !IsNil(o.ViolationContexts) {
		return true
	}

	return false
}

// SetViolationContexts gets a reference to the given []ViolationContext and assigns it to the ViolationContexts field.
func (o *ViolationPrediction) SetViolationContexts(v []ViolationContext) {
	o.ViolationContexts = v
}

func (o ViolationPrediction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ViolationPrediction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ViolationContexts) {
		toSerialize["violationContexts"] = o.ViolationContexts
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ViolationPrediction) UnmarshalJSON(data []byte) (err error) {
	varViolationPrediction := _ViolationPrediction{}

	err = json.Unmarshal(data, &varViolationPrediction)

	if err != nil {
		return err
	}

	*o = ViolationPrediction(varViolationPrediction)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "violationContexts")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableViolationPrediction struct {
	value *ViolationPrediction
	isSet bool
}

func (v NullableViolationPrediction) Get() *ViolationPrediction {
	return v.value
}

func (v *NullableViolationPrediction) Set(val *ViolationPrediction) {
	v.value = val
	v.isSet = true
}

func (v NullableViolationPrediction) IsSet() bool {
	return v.isSet
}

func (v *NullableViolationPrediction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViolationPrediction(val *ViolationPrediction) *NullableViolationPrediction {
	return &NullableViolationPrediction{value: val, isSet: true}
}

func (v NullableViolationPrediction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableViolationPrediction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


