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

// checks if the SodViolationContext1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SodViolationContext1{}

// SodViolationContext1 The contextual information of the violated criteria.
type SodViolationContext1 struct {
	Policy *SodPolicyDto1 `json:"policy,omitempty"`
	ConflictingAccessCriteria *SodViolationContext1ConflictingAccessCriteria `json:"conflictingAccessCriteria,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SodViolationContext1 SodViolationContext1

// NewSodViolationContext1 instantiates a new SodViolationContext1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSodViolationContext1() *SodViolationContext1 {
	this := SodViolationContext1{}
	return &this
}

// NewSodViolationContext1WithDefaults instantiates a new SodViolationContext1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSodViolationContext1WithDefaults() *SodViolationContext1 {
	this := SodViolationContext1{}
	return &this
}

// GetPolicy returns the Policy field value if set, zero value otherwise.
func (o *SodViolationContext1) GetPolicy() SodPolicyDto1 {
	if o == nil || IsNil(o.Policy) {
		var ret SodPolicyDto1
		return ret
	}
	return *o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SodViolationContext1) GetPolicyOk() (*SodPolicyDto1, bool) {
	if o == nil || IsNil(o.Policy) {
		return nil, false
	}
	return o.Policy, true
}

// HasPolicy returns a boolean if a field has been set.
func (o *SodViolationContext1) HasPolicy() bool {
	if o != nil && !IsNil(o.Policy) {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given SodPolicyDto1 and assigns it to the Policy field.
func (o *SodViolationContext1) SetPolicy(v SodPolicyDto1) {
	o.Policy = &v
}

// GetConflictingAccessCriteria returns the ConflictingAccessCriteria field value if set, zero value otherwise.
func (o *SodViolationContext1) GetConflictingAccessCriteria() SodViolationContext1ConflictingAccessCriteria {
	if o == nil || IsNil(o.ConflictingAccessCriteria) {
		var ret SodViolationContext1ConflictingAccessCriteria
		return ret
	}
	return *o.ConflictingAccessCriteria
}

// GetConflictingAccessCriteriaOk returns a tuple with the ConflictingAccessCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SodViolationContext1) GetConflictingAccessCriteriaOk() (*SodViolationContext1ConflictingAccessCriteria, bool) {
	if o == nil || IsNil(o.ConflictingAccessCriteria) {
		return nil, false
	}
	return o.ConflictingAccessCriteria, true
}

// HasConflictingAccessCriteria returns a boolean if a field has been set.
func (o *SodViolationContext1) HasConflictingAccessCriteria() bool {
	if o != nil && !IsNil(o.ConflictingAccessCriteria) {
		return true
	}

	return false
}

// SetConflictingAccessCriteria gets a reference to the given SodViolationContext1ConflictingAccessCriteria and assigns it to the ConflictingAccessCriteria field.
func (o *SodViolationContext1) SetConflictingAccessCriteria(v SodViolationContext1ConflictingAccessCriteria) {
	o.ConflictingAccessCriteria = &v
}

func (o SodViolationContext1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SodViolationContext1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Policy) {
		toSerialize["policy"] = o.Policy
	}
	if !IsNil(o.ConflictingAccessCriteria) {
		toSerialize["conflictingAccessCriteria"] = o.ConflictingAccessCriteria
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SodViolationContext1) UnmarshalJSON(data []byte) (err error) {
	varSodViolationContext1 := _SodViolationContext1{}

	err = json.Unmarshal(data, &varSodViolationContext1)

	if err != nil {
		return err
	}

	*o = SodViolationContext1(varSodViolationContext1)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "policy")
		delete(additionalProperties, "conflictingAccessCriteria")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSodViolationContext1 struct {
	value *SodViolationContext1
	isSet bool
}

func (v NullableSodViolationContext1) Get() *SodViolationContext1 {
	return v.value
}

func (v *NullableSodViolationContext1) Set(val *SodViolationContext1) {
	v.value = val
	v.isSet = true
}

func (v NullableSodViolationContext1) IsSet() bool {
	return v.isSet
}

func (v *NullableSodViolationContext1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSodViolationContext1(val *SodViolationContext1) *NullableSodViolationContext1 {
	return &NullableSodViolationContext1{value: val, isSet: true}
}

func (v NullableSodViolationContext1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSodViolationContext1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


