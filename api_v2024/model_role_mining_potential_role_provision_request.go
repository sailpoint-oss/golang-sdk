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

// checks if the RoleMiningPotentialRoleProvisionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoleMiningPotentialRoleProvisionRequest{}

// RoleMiningPotentialRoleProvisionRequest struct for RoleMiningPotentialRoleProvisionRequest
type RoleMiningPotentialRoleProvisionRequest struct {
	// Name of the new role being created
	RoleName *string `json:"roleName,omitempty"`
	// Short description of the new role being created
	RoleDescription *string `json:"roleDescription,omitempty"`
	// ID of the identity that will own this role
	OwnerId *string `json:"ownerId,omitempty"`
	// When true, create access requests for the identities associated with the potential role
	IncludeIdentities *bool `json:"includeIdentities,omitempty"`
	// When true, assign entitlements directly to the role; otherwise, create access profiles containing the entitlements
	DirectlyAssignedEntitlements *bool `json:"directlyAssignedEntitlements,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RoleMiningPotentialRoleProvisionRequest RoleMiningPotentialRoleProvisionRequest

// NewRoleMiningPotentialRoleProvisionRequest instantiates a new RoleMiningPotentialRoleProvisionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleMiningPotentialRoleProvisionRequest() *RoleMiningPotentialRoleProvisionRequest {
	this := RoleMiningPotentialRoleProvisionRequest{}
	var includeIdentities bool = false
	this.IncludeIdentities = &includeIdentities
	var directlyAssignedEntitlements bool = false
	this.DirectlyAssignedEntitlements = &directlyAssignedEntitlements
	return &this
}

// NewRoleMiningPotentialRoleProvisionRequestWithDefaults instantiates a new RoleMiningPotentialRoleProvisionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleMiningPotentialRoleProvisionRequestWithDefaults() *RoleMiningPotentialRoleProvisionRequest {
	this := RoleMiningPotentialRoleProvisionRequest{}
	var includeIdentities bool = false
	this.IncludeIdentities = &includeIdentities
	var directlyAssignedEntitlements bool = false
	this.DirectlyAssignedEntitlements = &directlyAssignedEntitlements
	return &this
}

// GetRoleName returns the RoleName field value if set, zero value otherwise.
func (o *RoleMiningPotentialRoleProvisionRequest) GetRoleName() string {
	if o == nil || IsNil(o.RoleName) {
		var ret string
		return ret
	}
	return *o.RoleName
}

// GetRoleNameOk returns a tuple with the RoleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningPotentialRoleProvisionRequest) GetRoleNameOk() (*string, bool) {
	if o == nil || IsNil(o.RoleName) {
		return nil, false
	}
	return o.RoleName, true
}

// HasRoleName returns a boolean if a field has been set.
func (o *RoleMiningPotentialRoleProvisionRequest) HasRoleName() bool {
	if o != nil && !IsNil(o.RoleName) {
		return true
	}

	return false
}

// SetRoleName gets a reference to the given string and assigns it to the RoleName field.
func (o *RoleMiningPotentialRoleProvisionRequest) SetRoleName(v string) {
	o.RoleName = &v
}

// GetRoleDescription returns the RoleDescription field value if set, zero value otherwise.
func (o *RoleMiningPotentialRoleProvisionRequest) GetRoleDescription() string {
	if o == nil || IsNil(o.RoleDescription) {
		var ret string
		return ret
	}
	return *o.RoleDescription
}

// GetRoleDescriptionOk returns a tuple with the RoleDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningPotentialRoleProvisionRequest) GetRoleDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.RoleDescription) {
		return nil, false
	}
	return o.RoleDescription, true
}

// HasRoleDescription returns a boolean if a field has been set.
func (o *RoleMiningPotentialRoleProvisionRequest) HasRoleDescription() bool {
	if o != nil && !IsNil(o.RoleDescription) {
		return true
	}

	return false
}

// SetRoleDescription gets a reference to the given string and assigns it to the RoleDescription field.
func (o *RoleMiningPotentialRoleProvisionRequest) SetRoleDescription(v string) {
	o.RoleDescription = &v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *RoleMiningPotentialRoleProvisionRequest) GetOwnerId() string {
	if o == nil || IsNil(o.OwnerId) {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningPotentialRoleProvisionRequest) GetOwnerIdOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerId) {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *RoleMiningPotentialRoleProvisionRequest) HasOwnerId() bool {
	if o != nil && !IsNil(o.OwnerId) {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *RoleMiningPotentialRoleProvisionRequest) SetOwnerId(v string) {
	o.OwnerId = &v
}

// GetIncludeIdentities returns the IncludeIdentities field value if set, zero value otherwise.
func (o *RoleMiningPotentialRoleProvisionRequest) GetIncludeIdentities() bool {
	if o == nil || IsNil(o.IncludeIdentities) {
		var ret bool
		return ret
	}
	return *o.IncludeIdentities
}

// GetIncludeIdentitiesOk returns a tuple with the IncludeIdentities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningPotentialRoleProvisionRequest) GetIncludeIdentitiesOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeIdentities) {
		return nil, false
	}
	return o.IncludeIdentities, true
}

// HasIncludeIdentities returns a boolean if a field has been set.
func (o *RoleMiningPotentialRoleProvisionRequest) HasIncludeIdentities() bool {
	if o != nil && !IsNil(o.IncludeIdentities) {
		return true
	}

	return false
}

// SetIncludeIdentities gets a reference to the given bool and assigns it to the IncludeIdentities field.
func (o *RoleMiningPotentialRoleProvisionRequest) SetIncludeIdentities(v bool) {
	o.IncludeIdentities = &v
}

// GetDirectlyAssignedEntitlements returns the DirectlyAssignedEntitlements field value if set, zero value otherwise.
func (o *RoleMiningPotentialRoleProvisionRequest) GetDirectlyAssignedEntitlements() bool {
	if o == nil || IsNil(o.DirectlyAssignedEntitlements) {
		var ret bool
		return ret
	}
	return *o.DirectlyAssignedEntitlements
}

// GetDirectlyAssignedEntitlementsOk returns a tuple with the DirectlyAssignedEntitlements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningPotentialRoleProvisionRequest) GetDirectlyAssignedEntitlementsOk() (*bool, bool) {
	if o == nil || IsNil(o.DirectlyAssignedEntitlements) {
		return nil, false
	}
	return o.DirectlyAssignedEntitlements, true
}

// HasDirectlyAssignedEntitlements returns a boolean if a field has been set.
func (o *RoleMiningPotentialRoleProvisionRequest) HasDirectlyAssignedEntitlements() bool {
	if o != nil && !IsNil(o.DirectlyAssignedEntitlements) {
		return true
	}

	return false
}

// SetDirectlyAssignedEntitlements gets a reference to the given bool and assigns it to the DirectlyAssignedEntitlements field.
func (o *RoleMiningPotentialRoleProvisionRequest) SetDirectlyAssignedEntitlements(v bool) {
	o.DirectlyAssignedEntitlements = &v
}

func (o RoleMiningPotentialRoleProvisionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoleMiningPotentialRoleProvisionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RoleName) {
		toSerialize["roleName"] = o.RoleName
	}
	if !IsNil(o.RoleDescription) {
		toSerialize["roleDescription"] = o.RoleDescription
	}
	if !IsNil(o.OwnerId) {
		toSerialize["ownerId"] = o.OwnerId
	}
	if !IsNil(o.IncludeIdentities) {
		toSerialize["includeIdentities"] = o.IncludeIdentities
	}
	if !IsNil(o.DirectlyAssignedEntitlements) {
		toSerialize["directlyAssignedEntitlements"] = o.DirectlyAssignedEntitlements
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RoleMiningPotentialRoleProvisionRequest) UnmarshalJSON(data []byte) (err error) {
	varRoleMiningPotentialRoleProvisionRequest := _RoleMiningPotentialRoleProvisionRequest{}

	err = json.Unmarshal(data, &varRoleMiningPotentialRoleProvisionRequest)

	if err != nil {
		return err
	}

	*o = RoleMiningPotentialRoleProvisionRequest(varRoleMiningPotentialRoleProvisionRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "roleName")
		delete(additionalProperties, "roleDescription")
		delete(additionalProperties, "ownerId")
		delete(additionalProperties, "includeIdentities")
		delete(additionalProperties, "directlyAssignedEntitlements")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRoleMiningPotentialRoleProvisionRequest struct {
	value *RoleMiningPotentialRoleProvisionRequest
	isSet bool
}

func (v NullableRoleMiningPotentialRoleProvisionRequest) Get() *RoleMiningPotentialRoleProvisionRequest {
	return v.value
}

func (v *NullableRoleMiningPotentialRoleProvisionRequest) Set(val *RoleMiningPotentialRoleProvisionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleMiningPotentialRoleProvisionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleMiningPotentialRoleProvisionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleMiningPotentialRoleProvisionRequest(val *RoleMiningPotentialRoleProvisionRequest) *NullableRoleMiningPotentialRoleProvisionRequest {
	return &NullableRoleMiningPotentialRoleProvisionRequest{value: val, isSet: true}
}

func (v NullableRoleMiningPotentialRoleProvisionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleMiningPotentialRoleProvisionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

