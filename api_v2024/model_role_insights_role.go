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

// checks if the RoleInsightsRole type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoleInsightsRole{}

// RoleInsightsRole struct for RoleInsightsRole
type RoleInsightsRole struct {
	// Role name
	Name *string `json:"name,omitempty"`
	// Role id
	Id *string `json:"id,omitempty"`
	// Role description
	Description *string `json:"description,omitempty"`
	// Role owner name
	OwnerName *string `json:"ownerName,omitempty"`
	// Role owner id
	OwnerId *string `json:"ownerId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RoleInsightsRole RoleInsightsRole

// NewRoleInsightsRole instantiates a new RoleInsightsRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleInsightsRole() *RoleInsightsRole {
	this := RoleInsightsRole{}
	return &this
}

// NewRoleInsightsRoleWithDefaults instantiates a new RoleInsightsRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleInsightsRoleWithDefaults() *RoleInsightsRole {
	this := RoleInsightsRole{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RoleInsightsRole) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleInsightsRole) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RoleInsightsRole) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RoleInsightsRole) SetName(v string) {
	o.Name = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RoleInsightsRole) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleInsightsRole) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RoleInsightsRole) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RoleInsightsRole) SetId(v string) {
	o.Id = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RoleInsightsRole) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleInsightsRole) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RoleInsightsRole) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RoleInsightsRole) SetDescription(v string) {
	o.Description = &v
}

// GetOwnerName returns the OwnerName field value if set, zero value otherwise.
func (o *RoleInsightsRole) GetOwnerName() string {
	if o == nil || IsNil(o.OwnerName) {
		var ret string
		return ret
	}
	return *o.OwnerName
}

// GetOwnerNameOk returns a tuple with the OwnerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleInsightsRole) GetOwnerNameOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerName) {
		return nil, false
	}
	return o.OwnerName, true
}

// HasOwnerName returns a boolean if a field has been set.
func (o *RoleInsightsRole) HasOwnerName() bool {
	if o != nil && !IsNil(o.OwnerName) {
		return true
	}

	return false
}

// SetOwnerName gets a reference to the given string and assigns it to the OwnerName field.
func (o *RoleInsightsRole) SetOwnerName(v string) {
	o.OwnerName = &v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *RoleInsightsRole) GetOwnerId() string {
	if o == nil || IsNil(o.OwnerId) {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleInsightsRole) GetOwnerIdOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerId) {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *RoleInsightsRole) HasOwnerId() bool {
	if o != nil && !IsNil(o.OwnerId) {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *RoleInsightsRole) SetOwnerId(v string) {
	o.OwnerId = &v
}

func (o RoleInsightsRole) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoleInsightsRole) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.OwnerName) {
		toSerialize["ownerName"] = o.OwnerName
	}
	if !IsNil(o.OwnerId) {
		toSerialize["ownerId"] = o.OwnerId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RoleInsightsRole) UnmarshalJSON(data []byte) (err error) {
	varRoleInsightsRole := _RoleInsightsRole{}

	err = json.Unmarshal(data, &varRoleInsightsRole)

	if err != nil {
		return err
	}

	*o = RoleInsightsRole(varRoleInsightsRole)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "id")
		delete(additionalProperties, "description")
		delete(additionalProperties, "ownerName")
		delete(additionalProperties, "ownerId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRoleInsightsRole struct {
	value *RoleInsightsRole
	isSet bool
}

func (v NullableRoleInsightsRole) Get() *RoleInsightsRole {
	return v.value
}

func (v *NullableRoleInsightsRole) Set(val *RoleInsightsRole) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleInsightsRole) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleInsightsRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleInsightsRole(val *RoleInsightsRole) *NullableRoleInsightsRole {
	return &NullableRoleInsightsRole{value: val, isSet: true}
}

func (v NullableRoleInsightsRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleInsightsRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

