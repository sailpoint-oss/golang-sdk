/*
Identity Security Cloud V2025 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: v2025
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2025

import (
	"encoding/json"
	"fmt"
)

// checks if the AccountUncorrelatedAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountUncorrelatedAccount{}

// AccountUncorrelatedAccount Uncorrelated account.
type AccountUncorrelatedAccount struct {
	// Uncorrelated account's DTO type.
	Type map[string]interface{} `json:"type"`
	// Uncorrelated account's ID.
	Id string `json:"id"`
	// Uncorrelated account's display name.
	Name string `json:"name"`
	// Unique ID of the account on the source.
	NativeIdentity string `json:"nativeIdentity"`
	// The source's unique identifier for the account. UUID is generated by the source system.
	Uuid NullableString `json:"uuid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountUncorrelatedAccount AccountUncorrelatedAccount

// NewAccountUncorrelatedAccount instantiates a new AccountUncorrelatedAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountUncorrelatedAccount(type_ map[string]interface{}, id string, name string, nativeIdentity string) *AccountUncorrelatedAccount {
	this := AccountUncorrelatedAccount{}
	this.Type = type_
	this.Id = id
	this.Name = name
	this.NativeIdentity = nativeIdentity
	return &this
}

// NewAccountUncorrelatedAccountWithDefaults instantiates a new AccountUncorrelatedAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountUncorrelatedAccountWithDefaults() *AccountUncorrelatedAccount {
	this := AccountUncorrelatedAccount{}
	return &this
}

// GetType returns the Type field value
func (o *AccountUncorrelatedAccount) GetType() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AccountUncorrelatedAccount) GetTypeOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Type, true
}

// SetType sets field value
func (o *AccountUncorrelatedAccount) SetType(v map[string]interface{}) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AccountUncorrelatedAccount) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AccountUncorrelatedAccount) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AccountUncorrelatedAccount) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *AccountUncorrelatedAccount) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AccountUncorrelatedAccount) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AccountUncorrelatedAccount) SetName(v string) {
	o.Name = v
}

// GetNativeIdentity returns the NativeIdentity field value
func (o *AccountUncorrelatedAccount) GetNativeIdentity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NativeIdentity
}

// GetNativeIdentityOk returns a tuple with the NativeIdentity field value
// and a boolean to check if the value has been set.
func (o *AccountUncorrelatedAccount) GetNativeIdentityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NativeIdentity, true
}

// SetNativeIdentity sets field value
func (o *AccountUncorrelatedAccount) SetNativeIdentity(v string) {
	o.NativeIdentity = v
}

// GetUuid returns the Uuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountUncorrelatedAccount) GetUuid() string {
	if o == nil || IsNil(o.Uuid.Get()) {
		var ret string
		return ret
	}
	return *o.Uuid.Get()
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountUncorrelatedAccount) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uuid.Get(), o.Uuid.IsSet()
}

// HasUuid returns a boolean if a field has been set.
func (o *AccountUncorrelatedAccount) HasUuid() bool {
	if o != nil && o.Uuid.IsSet() {
		return true
	}

	return false
}

// SetUuid gets a reference to the given NullableString and assigns it to the Uuid field.
func (o *AccountUncorrelatedAccount) SetUuid(v string) {
	o.Uuid.Set(&v)
}
// SetUuidNil sets the value for Uuid to be an explicit nil
func (o *AccountUncorrelatedAccount) SetUuidNil() {
	o.Uuid.Set(nil)
}

// UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
func (o *AccountUncorrelatedAccount) UnsetUuid() {
	o.Uuid.Unset()
}

func (o AccountUncorrelatedAccount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountUncorrelatedAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["nativeIdentity"] = o.NativeIdentity
	if o.Uuid.IsSet() {
		toSerialize["uuid"] = o.Uuid.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountUncorrelatedAccount) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"id",
		"name",
		"nativeIdentity",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAccountUncorrelatedAccount := _AccountUncorrelatedAccount{}

	err = json.Unmarshal(data, &varAccountUncorrelatedAccount)

	if err != nil {
		return err
	}

	*o = AccountUncorrelatedAccount(varAccountUncorrelatedAccount)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "nativeIdentity")
		delete(additionalProperties, "uuid")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountUncorrelatedAccount struct {
	value *AccountUncorrelatedAccount
	isSet bool
}

func (v NullableAccountUncorrelatedAccount) Get() *AccountUncorrelatedAccount {
	return v.value
}

func (v *NullableAccountUncorrelatedAccount) Set(val *AccountUncorrelatedAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountUncorrelatedAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountUncorrelatedAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountUncorrelatedAccount(val *AccountUncorrelatedAccount) *NullableAccountUncorrelatedAccount {
	return &NullableAccountUncorrelatedAccount{value: val, isSet: true}
}

func (v NullableAccountUncorrelatedAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountUncorrelatedAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


