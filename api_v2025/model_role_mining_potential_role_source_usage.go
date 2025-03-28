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

// checks if the RoleMiningPotentialRoleSourceUsage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoleMiningPotentialRoleSourceUsage{}

// RoleMiningPotentialRoleSourceUsage struct for RoleMiningPotentialRoleSourceUsage
type RoleMiningPotentialRoleSourceUsage struct {
	// The identity ID
	Id *string `json:"id,omitempty"`
	// Display name for the identity
	DisplayName *string `json:"displayName,omitempty"`
	// Email address for the identity
	Email *string `json:"email,omitempty"`
	// The number of days there has been usage of the source by the identity.
	UsageCount *int32 `json:"usageCount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RoleMiningPotentialRoleSourceUsage RoleMiningPotentialRoleSourceUsage

// NewRoleMiningPotentialRoleSourceUsage instantiates a new RoleMiningPotentialRoleSourceUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleMiningPotentialRoleSourceUsage() *RoleMiningPotentialRoleSourceUsage {
	this := RoleMiningPotentialRoleSourceUsage{}
	return &this
}

// NewRoleMiningPotentialRoleSourceUsageWithDefaults instantiates a new RoleMiningPotentialRoleSourceUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleMiningPotentialRoleSourceUsageWithDefaults() *RoleMiningPotentialRoleSourceUsage {
	this := RoleMiningPotentialRoleSourceUsage{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RoleMiningPotentialRoleSourceUsage) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningPotentialRoleSourceUsage) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RoleMiningPotentialRoleSourceUsage) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RoleMiningPotentialRoleSourceUsage) SetId(v string) {
	o.Id = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *RoleMiningPotentialRoleSourceUsage) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningPotentialRoleSourceUsage) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *RoleMiningPotentialRoleSourceUsage) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *RoleMiningPotentialRoleSourceUsage) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *RoleMiningPotentialRoleSourceUsage) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningPotentialRoleSourceUsage) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *RoleMiningPotentialRoleSourceUsage) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *RoleMiningPotentialRoleSourceUsage) SetEmail(v string) {
	o.Email = &v
}

// GetUsageCount returns the UsageCount field value if set, zero value otherwise.
func (o *RoleMiningPotentialRoleSourceUsage) GetUsageCount() int32 {
	if o == nil || IsNil(o.UsageCount) {
		var ret int32
		return ret
	}
	return *o.UsageCount
}

// GetUsageCountOk returns a tuple with the UsageCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningPotentialRoleSourceUsage) GetUsageCountOk() (*int32, bool) {
	if o == nil || IsNil(o.UsageCount) {
		return nil, false
	}
	return o.UsageCount, true
}

// HasUsageCount returns a boolean if a field has been set.
func (o *RoleMiningPotentialRoleSourceUsage) HasUsageCount() bool {
	if o != nil && !IsNil(o.UsageCount) {
		return true
	}

	return false
}

// SetUsageCount gets a reference to the given int32 and assigns it to the UsageCount field.
func (o *RoleMiningPotentialRoleSourceUsage) SetUsageCount(v int32) {
	o.UsageCount = &v
}

func (o RoleMiningPotentialRoleSourceUsage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoleMiningPotentialRoleSourceUsage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.UsageCount) {
		toSerialize["usageCount"] = o.UsageCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RoleMiningPotentialRoleSourceUsage) UnmarshalJSON(data []byte) (err error) {
	varRoleMiningPotentialRoleSourceUsage := _RoleMiningPotentialRoleSourceUsage{}

	err = json.Unmarshal(data, &varRoleMiningPotentialRoleSourceUsage)

	if err != nil {
		return err
	}

	*o = RoleMiningPotentialRoleSourceUsage(varRoleMiningPotentialRoleSourceUsage)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "email")
		delete(additionalProperties, "usageCount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRoleMiningPotentialRoleSourceUsage struct {
	value *RoleMiningPotentialRoleSourceUsage
	isSet bool
}

func (v NullableRoleMiningPotentialRoleSourceUsage) Get() *RoleMiningPotentialRoleSourceUsage {
	return v.value
}

func (v *NullableRoleMiningPotentialRoleSourceUsage) Set(val *RoleMiningPotentialRoleSourceUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleMiningPotentialRoleSourceUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleMiningPotentialRoleSourceUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleMiningPotentialRoleSourceUsage(val *RoleMiningPotentialRoleSourceUsage) *NullableRoleMiningPotentialRoleSourceUsage {
	return &NullableRoleMiningPotentialRoleSourceUsage{value: val, isSet: true}
}

func (v NullableRoleMiningPotentialRoleSourceUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleMiningPotentialRoleSourceUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


