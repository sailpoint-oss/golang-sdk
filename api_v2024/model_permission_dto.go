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

// checks if the PermissionDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PermissionDto{}

// PermissionDto Simplified DTO for the Permission objects stored in SailPoint's database. The data is aggregated from customer systems and is free-form, so its appearance can vary largely between different clients/customers.
type PermissionDto struct {
	// All the rights (e.g. actions) that this permission allows on the target
	Rights []string `json:"rights,omitempty"`
	// The target the permission would grants rights on.
	Target *string `json:"target,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PermissionDto PermissionDto

// NewPermissionDto instantiates a new PermissionDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissionDto() *PermissionDto {
	this := PermissionDto{}
	return &this
}

// NewPermissionDtoWithDefaults instantiates a new PermissionDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionDtoWithDefaults() *PermissionDto {
	this := PermissionDto{}
	return &this
}

// GetRights returns the Rights field value if set, zero value otherwise.
func (o *PermissionDto) GetRights() []string {
	if o == nil || IsNil(o.Rights) {
		var ret []string
		return ret
	}
	return o.Rights
}

// GetRightsOk returns a tuple with the Rights field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionDto) GetRightsOk() ([]string, bool) {
	if o == nil || IsNil(o.Rights) {
		return nil, false
	}
	return o.Rights, true
}

// HasRights returns a boolean if a field has been set.
func (o *PermissionDto) HasRights() bool {
	if o != nil && !IsNil(o.Rights) {
		return true
	}

	return false
}

// SetRights gets a reference to the given []string and assigns it to the Rights field.
func (o *PermissionDto) SetRights(v []string) {
	o.Rights = v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *PermissionDto) GetTarget() string {
	if o == nil || IsNil(o.Target) {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionDto) GetTargetOk() (*string, bool) {
	if o == nil || IsNil(o.Target) {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *PermissionDto) HasTarget() bool {
	if o != nil && !IsNil(o.Target) {
		return true
	}

	return false
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *PermissionDto) SetTarget(v string) {
	o.Target = &v
}

func (o PermissionDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PermissionDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Rights) {
		toSerialize["rights"] = o.Rights
	}
	if !IsNil(o.Target) {
		toSerialize["target"] = o.Target
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PermissionDto) UnmarshalJSON(data []byte) (err error) {
	varPermissionDto := _PermissionDto{}

	err = json.Unmarshal(data, &varPermissionDto)

	if err != nil {
		return err
	}

	*o = PermissionDto(varPermissionDto)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rights")
		delete(additionalProperties, "target")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePermissionDto struct {
	value *PermissionDto
	isSet bool
}

func (v NullablePermissionDto) Get() *PermissionDto {
	return v.value
}

func (v *NullablePermissionDto) Set(val *PermissionDto) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionDto) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionDto(val *PermissionDto) *NullablePermissionDto {
	return &NullablePermissionDto{value: val, isSet: true}
}

func (v NullablePermissionDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

