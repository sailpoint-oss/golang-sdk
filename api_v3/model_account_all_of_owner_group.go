/*
Identity Security Cloud V3 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v3

import (
	"encoding/json"
)

// checks if the AccountAllOfOwnerGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountAllOfOwnerGroup{}

// AccountAllOfOwnerGroup The governance group who owns this account, typically used for non-human accounts
type AccountAllOfOwnerGroup struct {
	// The type of object being referenced
	Type *string `json:"type,omitempty"`
	// ID of the governance group
	Id *string `json:"id,omitempty"`
	// Human-readable display name of the governance group
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountAllOfOwnerGroup AccountAllOfOwnerGroup

// NewAccountAllOfOwnerGroup instantiates a new AccountAllOfOwnerGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountAllOfOwnerGroup() *AccountAllOfOwnerGroup {
	this := AccountAllOfOwnerGroup{}
	return &this
}

// NewAccountAllOfOwnerGroupWithDefaults instantiates a new AccountAllOfOwnerGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountAllOfOwnerGroupWithDefaults() *AccountAllOfOwnerGroup {
	this := AccountAllOfOwnerGroup{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AccountAllOfOwnerGroup) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountAllOfOwnerGroup) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AccountAllOfOwnerGroup) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AccountAllOfOwnerGroup) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AccountAllOfOwnerGroup) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountAllOfOwnerGroup) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AccountAllOfOwnerGroup) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AccountAllOfOwnerGroup) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AccountAllOfOwnerGroup) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountAllOfOwnerGroup) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AccountAllOfOwnerGroup) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AccountAllOfOwnerGroup) SetName(v string) {
	o.Name = &v
}

func (o AccountAllOfOwnerGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountAllOfOwnerGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountAllOfOwnerGroup) UnmarshalJSON(data []byte) (err error) {
	varAccountAllOfOwnerGroup := _AccountAllOfOwnerGroup{}

	err = json.Unmarshal(data, &varAccountAllOfOwnerGroup)

	if err != nil {
		return err
	}

	*o = AccountAllOfOwnerGroup(varAccountAllOfOwnerGroup)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountAllOfOwnerGroup struct {
	value *AccountAllOfOwnerGroup
	isSet bool
}

func (v NullableAccountAllOfOwnerGroup) Get() *AccountAllOfOwnerGroup {
	return v.value
}

func (v *NullableAccountAllOfOwnerGroup) Set(val *AccountAllOfOwnerGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountAllOfOwnerGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountAllOfOwnerGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountAllOfOwnerGroup(val *AccountAllOfOwnerGroup) *NullableAccountAllOfOwnerGroup {
	return &NullableAccountAllOfOwnerGroup{value: val, isSet: true}
}

func (v NullableAccountAllOfOwnerGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountAllOfOwnerGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

