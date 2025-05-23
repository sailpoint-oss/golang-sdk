/*
Identity Security Cloud Beta API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_beta

import (
	"encoding/json"
	"fmt"
)

// checks if the AccountUncorrelatedIdentity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountUncorrelatedIdentity{}

// AccountUncorrelatedIdentity Identity the account is uncorrelated with.
type AccountUncorrelatedIdentity struct {
	// DTO type of the identity the account is uncorrelated with.
	Type string `json:"type"`
	// ID of the identity the account is uncorrelated with.
	Id string `json:"id"`
	// Display name of the identity the account is uncorrelated with.
	Name string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _AccountUncorrelatedIdentity AccountUncorrelatedIdentity

// NewAccountUncorrelatedIdentity instantiates a new AccountUncorrelatedIdentity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountUncorrelatedIdentity(type_ string, id string, name string) *AccountUncorrelatedIdentity {
	this := AccountUncorrelatedIdentity{}
	this.Type = type_
	this.Id = id
	this.Name = name
	return &this
}

// NewAccountUncorrelatedIdentityWithDefaults instantiates a new AccountUncorrelatedIdentity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountUncorrelatedIdentityWithDefaults() *AccountUncorrelatedIdentity {
	this := AccountUncorrelatedIdentity{}
	return &this
}

// GetType returns the Type field value
func (o *AccountUncorrelatedIdentity) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AccountUncorrelatedIdentity) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AccountUncorrelatedIdentity) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AccountUncorrelatedIdentity) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AccountUncorrelatedIdentity) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AccountUncorrelatedIdentity) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *AccountUncorrelatedIdentity) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AccountUncorrelatedIdentity) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AccountUncorrelatedIdentity) SetName(v string) {
	o.Name = v
}

func (o AccountUncorrelatedIdentity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountUncorrelatedIdentity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountUncorrelatedIdentity) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"id",
		"name",
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

	varAccountUncorrelatedIdentity := _AccountUncorrelatedIdentity{}

	err = json.Unmarshal(data, &varAccountUncorrelatedIdentity)

	if err != nil {
		return err
	}

	*o = AccountUncorrelatedIdentity(varAccountUncorrelatedIdentity)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountUncorrelatedIdentity struct {
	value *AccountUncorrelatedIdentity
	isSet bool
}

func (v NullableAccountUncorrelatedIdentity) Get() *AccountUncorrelatedIdentity {
	return v.value
}

func (v *NullableAccountUncorrelatedIdentity) Set(val *AccountUncorrelatedIdentity) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountUncorrelatedIdentity) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountUncorrelatedIdentity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountUncorrelatedIdentity(val *AccountUncorrelatedIdentity) *NullableAccountUncorrelatedIdentity {
	return &NullableAccountUncorrelatedIdentity{value: val, isSet: true}
}

func (v NullableAccountUncorrelatedIdentity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountUncorrelatedIdentity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


