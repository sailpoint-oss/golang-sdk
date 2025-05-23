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

// checks if the AccountItemRef type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountItemRef{}

// AccountItemRef struct for AccountItemRef
type AccountItemRef struct {
	// The uuid for the account, available under the 'objectguid' attribute
	AccountUuid NullableString `json:"accountUuid,omitempty"`
	// The 'distinguishedName' attribute for the account
	NativeIdentity *string `json:"nativeIdentity,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountItemRef AccountItemRef

// NewAccountItemRef instantiates a new AccountItemRef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountItemRef() *AccountItemRef {
	this := AccountItemRef{}
	return &this
}

// NewAccountItemRefWithDefaults instantiates a new AccountItemRef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountItemRefWithDefaults() *AccountItemRef {
	this := AccountItemRef{}
	return &this
}

// GetAccountUuid returns the AccountUuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountItemRef) GetAccountUuid() string {
	if o == nil || IsNil(o.AccountUuid.Get()) {
		var ret string
		return ret
	}
	return *o.AccountUuid.Get()
}

// GetAccountUuidOk returns a tuple with the AccountUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountItemRef) GetAccountUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountUuid.Get(), o.AccountUuid.IsSet()
}

// HasAccountUuid returns a boolean if a field has been set.
func (o *AccountItemRef) HasAccountUuid() bool {
	if o != nil && o.AccountUuid.IsSet() {
		return true
	}

	return false
}

// SetAccountUuid gets a reference to the given NullableString and assigns it to the AccountUuid field.
func (o *AccountItemRef) SetAccountUuid(v string) {
	o.AccountUuid.Set(&v)
}
// SetAccountUuidNil sets the value for AccountUuid to be an explicit nil
func (o *AccountItemRef) SetAccountUuidNil() {
	o.AccountUuid.Set(nil)
}

// UnsetAccountUuid ensures that no value is present for AccountUuid, not even an explicit nil
func (o *AccountItemRef) UnsetAccountUuid() {
	o.AccountUuid.Unset()
}

// GetNativeIdentity returns the NativeIdentity field value if set, zero value otherwise.
func (o *AccountItemRef) GetNativeIdentity() string {
	if o == nil || IsNil(o.NativeIdentity) {
		var ret string
		return ret
	}
	return *o.NativeIdentity
}

// GetNativeIdentityOk returns a tuple with the NativeIdentity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountItemRef) GetNativeIdentityOk() (*string, bool) {
	if o == nil || IsNil(o.NativeIdentity) {
		return nil, false
	}
	return o.NativeIdentity, true
}

// HasNativeIdentity returns a boolean if a field has been set.
func (o *AccountItemRef) HasNativeIdentity() bool {
	if o != nil && !IsNil(o.NativeIdentity) {
		return true
	}

	return false
}

// SetNativeIdentity gets a reference to the given string and assigns it to the NativeIdentity field.
func (o *AccountItemRef) SetNativeIdentity(v string) {
	o.NativeIdentity = &v
}

func (o AccountItemRef) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountItemRef) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountUuid.IsSet() {
		toSerialize["accountUuid"] = o.AccountUuid.Get()
	}
	if !IsNil(o.NativeIdentity) {
		toSerialize["nativeIdentity"] = o.NativeIdentity
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountItemRef) UnmarshalJSON(data []byte) (err error) {
	varAccountItemRef := _AccountItemRef{}

	err = json.Unmarshal(data, &varAccountItemRef)

	if err != nil {
		return err
	}

	*o = AccountItemRef(varAccountItemRef)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "accountUuid")
		delete(additionalProperties, "nativeIdentity")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountItemRef struct {
	value *AccountItemRef
	isSet bool
}

func (v NullableAccountItemRef) Get() *AccountItemRef {
	return v.value
}

func (v *NullableAccountItemRef) Set(val *AccountItemRef) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountItemRef) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountItemRef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountItemRef(val *AccountItemRef) *NullableAccountItemRef {
	return &NullableAccountItemRef{value: val, isSet: true}
}

func (v NullableAccountItemRef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountItemRef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


