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

// checks if the ManagerCorrelationMapping type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagerCorrelationMapping{}

// ManagerCorrelationMapping struct for ManagerCorrelationMapping
type ManagerCorrelationMapping struct {
	// Name of the attribute to use for manager correlation. The value found on the account attribute will be used to lookup the manager's identity.
	AccountAttributeName *string `json:"accountAttributeName,omitempty"`
	// Name of the identity attribute to search when trying to find a manager using the value from the accountAttribute.
	IdentityAttributeName *string `json:"identityAttributeName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ManagerCorrelationMapping ManagerCorrelationMapping

// NewManagerCorrelationMapping instantiates a new ManagerCorrelationMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagerCorrelationMapping() *ManagerCorrelationMapping {
	this := ManagerCorrelationMapping{}
	return &this
}

// NewManagerCorrelationMappingWithDefaults instantiates a new ManagerCorrelationMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagerCorrelationMappingWithDefaults() *ManagerCorrelationMapping {
	this := ManagerCorrelationMapping{}
	return &this
}

// GetAccountAttributeName returns the AccountAttributeName field value if set, zero value otherwise.
func (o *ManagerCorrelationMapping) GetAccountAttributeName() string {
	if o == nil || IsNil(o.AccountAttributeName) {
		var ret string
		return ret
	}
	return *o.AccountAttributeName
}

// GetAccountAttributeNameOk returns a tuple with the AccountAttributeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagerCorrelationMapping) GetAccountAttributeNameOk() (*string, bool) {
	if o == nil || IsNil(o.AccountAttributeName) {
		return nil, false
	}
	return o.AccountAttributeName, true
}

// HasAccountAttributeName returns a boolean if a field has been set.
func (o *ManagerCorrelationMapping) HasAccountAttributeName() bool {
	if o != nil && !IsNil(o.AccountAttributeName) {
		return true
	}

	return false
}

// SetAccountAttributeName gets a reference to the given string and assigns it to the AccountAttributeName field.
func (o *ManagerCorrelationMapping) SetAccountAttributeName(v string) {
	o.AccountAttributeName = &v
}

// GetIdentityAttributeName returns the IdentityAttributeName field value if set, zero value otherwise.
func (o *ManagerCorrelationMapping) GetIdentityAttributeName() string {
	if o == nil || IsNil(o.IdentityAttributeName) {
		var ret string
		return ret
	}
	return *o.IdentityAttributeName
}

// GetIdentityAttributeNameOk returns a tuple with the IdentityAttributeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagerCorrelationMapping) GetIdentityAttributeNameOk() (*string, bool) {
	if o == nil || IsNil(o.IdentityAttributeName) {
		return nil, false
	}
	return o.IdentityAttributeName, true
}

// HasIdentityAttributeName returns a boolean if a field has been set.
func (o *ManagerCorrelationMapping) HasIdentityAttributeName() bool {
	if o != nil && !IsNil(o.IdentityAttributeName) {
		return true
	}

	return false
}

// SetIdentityAttributeName gets a reference to the given string and assigns it to the IdentityAttributeName field.
func (o *ManagerCorrelationMapping) SetIdentityAttributeName(v string) {
	o.IdentityAttributeName = &v
}

func (o ManagerCorrelationMapping) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagerCorrelationMapping) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountAttributeName) {
		toSerialize["accountAttributeName"] = o.AccountAttributeName
	}
	if !IsNil(o.IdentityAttributeName) {
		toSerialize["identityAttributeName"] = o.IdentityAttributeName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ManagerCorrelationMapping) UnmarshalJSON(data []byte) (err error) {
	varManagerCorrelationMapping := _ManagerCorrelationMapping{}

	err = json.Unmarshal(data, &varManagerCorrelationMapping)

	if err != nil {
		return err
	}

	*o = ManagerCorrelationMapping(varManagerCorrelationMapping)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "accountAttributeName")
		delete(additionalProperties, "identityAttributeName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableManagerCorrelationMapping struct {
	value *ManagerCorrelationMapping
	isSet bool
}

func (v NullableManagerCorrelationMapping) Get() *ManagerCorrelationMapping {
	return v.value
}

func (v *NullableManagerCorrelationMapping) Set(val *ManagerCorrelationMapping) {
	v.value = val
	v.isSet = true
}

func (v NullableManagerCorrelationMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableManagerCorrelationMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagerCorrelationMapping(val *ManagerCorrelationMapping) *NullableManagerCorrelationMapping {
	return &NullableManagerCorrelationMapping{value: val, isSet: true}
}

func (v NullableManagerCorrelationMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagerCorrelationMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


