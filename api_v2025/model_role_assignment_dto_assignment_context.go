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

// checks if the RoleAssignmentDtoAssignmentContext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoleAssignmentDtoAssignmentContext{}

// RoleAssignmentDtoAssignmentContext struct for RoleAssignmentDtoAssignmentContext
type RoleAssignmentDtoAssignmentContext struct {
	Requested *AccessRequestContext `json:"requested,omitempty"`
	Matched []RoleMatchDto `json:"matched,omitempty"`
	// Date that the assignment will was evaluated
	ComputedDate *string `json:"computedDate,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RoleAssignmentDtoAssignmentContext RoleAssignmentDtoAssignmentContext

// NewRoleAssignmentDtoAssignmentContext instantiates a new RoleAssignmentDtoAssignmentContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleAssignmentDtoAssignmentContext() *RoleAssignmentDtoAssignmentContext {
	this := RoleAssignmentDtoAssignmentContext{}
	return &this
}

// NewRoleAssignmentDtoAssignmentContextWithDefaults instantiates a new RoleAssignmentDtoAssignmentContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleAssignmentDtoAssignmentContextWithDefaults() *RoleAssignmentDtoAssignmentContext {
	this := RoleAssignmentDtoAssignmentContext{}
	return &this
}

// GetRequested returns the Requested field value if set, zero value otherwise.
func (o *RoleAssignmentDtoAssignmentContext) GetRequested() AccessRequestContext {
	if o == nil || IsNil(o.Requested) {
		var ret AccessRequestContext
		return ret
	}
	return *o.Requested
}

// GetRequestedOk returns a tuple with the Requested field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleAssignmentDtoAssignmentContext) GetRequestedOk() (*AccessRequestContext, bool) {
	if o == nil || IsNil(o.Requested) {
		return nil, false
	}
	return o.Requested, true
}

// HasRequested returns a boolean if a field has been set.
func (o *RoleAssignmentDtoAssignmentContext) HasRequested() bool {
	if o != nil && !IsNil(o.Requested) {
		return true
	}

	return false
}

// SetRequested gets a reference to the given AccessRequestContext and assigns it to the Requested field.
func (o *RoleAssignmentDtoAssignmentContext) SetRequested(v AccessRequestContext) {
	o.Requested = &v
}

// GetMatched returns the Matched field value if set, zero value otherwise.
func (o *RoleAssignmentDtoAssignmentContext) GetMatched() []RoleMatchDto {
	if o == nil || IsNil(o.Matched) {
		var ret []RoleMatchDto
		return ret
	}
	return o.Matched
}

// GetMatchedOk returns a tuple with the Matched field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleAssignmentDtoAssignmentContext) GetMatchedOk() ([]RoleMatchDto, bool) {
	if o == nil || IsNil(o.Matched) {
		return nil, false
	}
	return o.Matched, true
}

// HasMatched returns a boolean if a field has been set.
func (o *RoleAssignmentDtoAssignmentContext) HasMatched() bool {
	if o != nil && !IsNil(o.Matched) {
		return true
	}

	return false
}

// SetMatched gets a reference to the given []RoleMatchDto and assigns it to the Matched field.
func (o *RoleAssignmentDtoAssignmentContext) SetMatched(v []RoleMatchDto) {
	o.Matched = v
}

// GetComputedDate returns the ComputedDate field value if set, zero value otherwise.
func (o *RoleAssignmentDtoAssignmentContext) GetComputedDate() string {
	if o == nil || IsNil(o.ComputedDate) {
		var ret string
		return ret
	}
	return *o.ComputedDate
}

// GetComputedDateOk returns a tuple with the ComputedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleAssignmentDtoAssignmentContext) GetComputedDateOk() (*string, bool) {
	if o == nil || IsNil(o.ComputedDate) {
		return nil, false
	}
	return o.ComputedDate, true
}

// HasComputedDate returns a boolean if a field has been set.
func (o *RoleAssignmentDtoAssignmentContext) HasComputedDate() bool {
	if o != nil && !IsNil(o.ComputedDate) {
		return true
	}

	return false
}

// SetComputedDate gets a reference to the given string and assigns it to the ComputedDate field.
func (o *RoleAssignmentDtoAssignmentContext) SetComputedDate(v string) {
	o.ComputedDate = &v
}

func (o RoleAssignmentDtoAssignmentContext) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoleAssignmentDtoAssignmentContext) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Requested) {
		toSerialize["requested"] = o.Requested
	}
	if !IsNil(o.Matched) {
		toSerialize["matched"] = o.Matched
	}
	if !IsNil(o.ComputedDate) {
		toSerialize["computedDate"] = o.ComputedDate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RoleAssignmentDtoAssignmentContext) UnmarshalJSON(data []byte) (err error) {
	varRoleAssignmentDtoAssignmentContext := _RoleAssignmentDtoAssignmentContext{}

	err = json.Unmarshal(data, &varRoleAssignmentDtoAssignmentContext)

	if err != nil {
		return err
	}

	*o = RoleAssignmentDtoAssignmentContext(varRoleAssignmentDtoAssignmentContext)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "requested")
		delete(additionalProperties, "matched")
		delete(additionalProperties, "computedDate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRoleAssignmentDtoAssignmentContext struct {
	value *RoleAssignmentDtoAssignmentContext
	isSet bool
}

func (v NullableRoleAssignmentDtoAssignmentContext) Get() *RoleAssignmentDtoAssignmentContext {
	return v.value
}

func (v *NullableRoleAssignmentDtoAssignmentContext) Set(val *RoleAssignmentDtoAssignmentContext) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleAssignmentDtoAssignmentContext) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleAssignmentDtoAssignmentContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleAssignmentDtoAssignmentContext(val *RoleAssignmentDtoAssignmentContext) *NullableRoleAssignmentDtoAssignmentContext {
	return &NullableRoleAssignmentDtoAssignmentContext{value: val, isSet: true}
}

func (v NullableRoleAssignmentDtoAssignmentContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleAssignmentDtoAssignmentContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


