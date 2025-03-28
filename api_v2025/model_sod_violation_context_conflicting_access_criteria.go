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

// checks if the SodViolationContextConflictingAccessCriteria type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SodViolationContextConflictingAccessCriteria{}

// SodViolationContextConflictingAccessCriteria The object which contains the left and right hand side of the entitlements that got violated according to the policy.
type SodViolationContextConflictingAccessCriteria struct {
	LeftCriteria *SodViolationContextConflictingAccessCriteriaLeftCriteria `json:"leftCriteria,omitempty"`
	RightCriteria *SodViolationContextConflictingAccessCriteriaLeftCriteria `json:"rightCriteria,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SodViolationContextConflictingAccessCriteria SodViolationContextConflictingAccessCriteria

// NewSodViolationContextConflictingAccessCriteria instantiates a new SodViolationContextConflictingAccessCriteria object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSodViolationContextConflictingAccessCriteria() *SodViolationContextConflictingAccessCriteria {
	this := SodViolationContextConflictingAccessCriteria{}
	return &this
}

// NewSodViolationContextConflictingAccessCriteriaWithDefaults instantiates a new SodViolationContextConflictingAccessCriteria object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSodViolationContextConflictingAccessCriteriaWithDefaults() *SodViolationContextConflictingAccessCriteria {
	this := SodViolationContextConflictingAccessCriteria{}
	return &this
}

// GetLeftCriteria returns the LeftCriteria field value if set, zero value otherwise.
func (o *SodViolationContextConflictingAccessCriteria) GetLeftCriteria() SodViolationContextConflictingAccessCriteriaLeftCriteria {
	if o == nil || IsNil(o.LeftCriteria) {
		var ret SodViolationContextConflictingAccessCriteriaLeftCriteria
		return ret
	}
	return *o.LeftCriteria
}

// GetLeftCriteriaOk returns a tuple with the LeftCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SodViolationContextConflictingAccessCriteria) GetLeftCriteriaOk() (*SodViolationContextConflictingAccessCriteriaLeftCriteria, bool) {
	if o == nil || IsNil(o.LeftCriteria) {
		return nil, false
	}
	return o.LeftCriteria, true
}

// HasLeftCriteria returns a boolean if a field has been set.
func (o *SodViolationContextConflictingAccessCriteria) HasLeftCriteria() bool {
	if o != nil && !IsNil(o.LeftCriteria) {
		return true
	}

	return false
}

// SetLeftCriteria gets a reference to the given SodViolationContextConflictingAccessCriteriaLeftCriteria and assigns it to the LeftCriteria field.
func (o *SodViolationContextConflictingAccessCriteria) SetLeftCriteria(v SodViolationContextConflictingAccessCriteriaLeftCriteria) {
	o.LeftCriteria = &v
}

// GetRightCriteria returns the RightCriteria field value if set, zero value otherwise.
func (o *SodViolationContextConflictingAccessCriteria) GetRightCriteria() SodViolationContextConflictingAccessCriteriaLeftCriteria {
	if o == nil || IsNil(o.RightCriteria) {
		var ret SodViolationContextConflictingAccessCriteriaLeftCriteria
		return ret
	}
	return *o.RightCriteria
}

// GetRightCriteriaOk returns a tuple with the RightCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SodViolationContextConflictingAccessCriteria) GetRightCriteriaOk() (*SodViolationContextConflictingAccessCriteriaLeftCriteria, bool) {
	if o == nil || IsNil(o.RightCriteria) {
		return nil, false
	}
	return o.RightCriteria, true
}

// HasRightCriteria returns a boolean if a field has been set.
func (o *SodViolationContextConflictingAccessCriteria) HasRightCriteria() bool {
	if o != nil && !IsNil(o.RightCriteria) {
		return true
	}

	return false
}

// SetRightCriteria gets a reference to the given SodViolationContextConflictingAccessCriteriaLeftCriteria and assigns it to the RightCriteria field.
func (o *SodViolationContextConflictingAccessCriteria) SetRightCriteria(v SodViolationContextConflictingAccessCriteriaLeftCriteria) {
	o.RightCriteria = &v
}

func (o SodViolationContextConflictingAccessCriteria) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SodViolationContextConflictingAccessCriteria) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LeftCriteria) {
		toSerialize["leftCriteria"] = o.LeftCriteria
	}
	if !IsNil(o.RightCriteria) {
		toSerialize["rightCriteria"] = o.RightCriteria
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SodViolationContextConflictingAccessCriteria) UnmarshalJSON(data []byte) (err error) {
	varSodViolationContextConflictingAccessCriteria := _SodViolationContextConflictingAccessCriteria{}

	err = json.Unmarshal(data, &varSodViolationContextConflictingAccessCriteria)

	if err != nil {
		return err
	}

	*o = SodViolationContextConflictingAccessCriteria(varSodViolationContextConflictingAccessCriteria)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "leftCriteria")
		delete(additionalProperties, "rightCriteria")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSodViolationContextConflictingAccessCriteria struct {
	value *SodViolationContextConflictingAccessCriteria
	isSet bool
}

func (v NullableSodViolationContextConflictingAccessCriteria) Get() *SodViolationContextConflictingAccessCriteria {
	return v.value
}

func (v *NullableSodViolationContextConflictingAccessCriteria) Set(val *SodViolationContextConflictingAccessCriteria) {
	v.value = val
	v.isSet = true
}

func (v NullableSodViolationContextConflictingAccessCriteria) IsSet() bool {
	return v.isSet
}

func (v *NullableSodViolationContextConflictingAccessCriteria) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSodViolationContextConflictingAccessCriteria(val *SodViolationContextConflictingAccessCriteria) *NullableSodViolationContextConflictingAccessCriteria {
	return &NullableSodViolationContextConflictingAccessCriteria{value: val, isSet: true}
}

func (v NullableSodViolationContextConflictingAccessCriteria) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSodViolationContextConflictingAccessCriteria) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


