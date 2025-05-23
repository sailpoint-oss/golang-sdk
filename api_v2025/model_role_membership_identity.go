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

// checks if the RoleMembershipIdentity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoleMembershipIdentity{}

// RoleMembershipIdentity A reference to an Identity in an IDENTITY_LIST role membership criteria.
type RoleMembershipIdentity struct {
	Type *DtoType `json:"type,omitempty"`
	// Identity id
	Id *string `json:"id,omitempty"`
	// Human-readable display name of the Identity.
	Name NullableString `json:"name,omitempty"`
	// User name of the Identity
	AliasName NullableString `json:"aliasName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RoleMembershipIdentity RoleMembershipIdentity

// NewRoleMembershipIdentity instantiates a new RoleMembershipIdentity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleMembershipIdentity() *RoleMembershipIdentity {
	this := RoleMembershipIdentity{}
	return &this
}

// NewRoleMembershipIdentityWithDefaults instantiates a new RoleMembershipIdentity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleMembershipIdentityWithDefaults() *RoleMembershipIdentity {
	this := RoleMembershipIdentity{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RoleMembershipIdentity) GetType() DtoType {
	if o == nil || IsNil(o.Type) {
		var ret DtoType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMembershipIdentity) GetTypeOk() (*DtoType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RoleMembershipIdentity) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given DtoType and assigns it to the Type field.
func (o *RoleMembershipIdentity) SetType(v DtoType) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RoleMembershipIdentity) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMembershipIdentity) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RoleMembershipIdentity) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RoleMembershipIdentity) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RoleMembershipIdentity) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoleMembershipIdentity) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *RoleMembershipIdentity) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *RoleMembershipIdentity) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *RoleMembershipIdentity) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *RoleMembershipIdentity) UnsetName() {
	o.Name.Unset()
}

// GetAliasName returns the AliasName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RoleMembershipIdentity) GetAliasName() string {
	if o == nil || IsNil(o.AliasName.Get()) {
		var ret string
		return ret
	}
	return *o.AliasName.Get()
}

// GetAliasNameOk returns a tuple with the AliasName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoleMembershipIdentity) GetAliasNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AliasName.Get(), o.AliasName.IsSet()
}

// HasAliasName returns a boolean if a field has been set.
func (o *RoleMembershipIdentity) HasAliasName() bool {
	if o != nil && o.AliasName.IsSet() {
		return true
	}

	return false
}

// SetAliasName gets a reference to the given NullableString and assigns it to the AliasName field.
func (o *RoleMembershipIdentity) SetAliasName(v string) {
	o.AliasName.Set(&v)
}
// SetAliasNameNil sets the value for AliasName to be an explicit nil
func (o *RoleMembershipIdentity) SetAliasNameNil() {
	o.AliasName.Set(nil)
}

// UnsetAliasName ensures that no value is present for AliasName, not even an explicit nil
func (o *RoleMembershipIdentity) UnsetAliasName() {
	o.AliasName.Unset()
}

func (o RoleMembershipIdentity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoleMembershipIdentity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.AliasName.IsSet() {
		toSerialize["aliasName"] = o.AliasName.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RoleMembershipIdentity) UnmarshalJSON(data []byte) (err error) {
	varRoleMembershipIdentity := _RoleMembershipIdentity{}

	err = json.Unmarshal(data, &varRoleMembershipIdentity)

	if err != nil {
		return err
	}

	*o = RoleMembershipIdentity(varRoleMembershipIdentity)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "aliasName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRoleMembershipIdentity struct {
	value *RoleMembershipIdentity
	isSet bool
}

func (v NullableRoleMembershipIdentity) Get() *RoleMembershipIdentity {
	return v.value
}

func (v *NullableRoleMembershipIdentity) Set(val *RoleMembershipIdentity) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleMembershipIdentity) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleMembershipIdentity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleMembershipIdentity(val *RoleMembershipIdentity) *NullableRoleMembershipIdentity {
	return &NullableRoleMembershipIdentity{value: val, isSet: true}
}

func (v NullableRoleMembershipIdentity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleMembershipIdentity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


