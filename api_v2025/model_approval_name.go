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

// checks if the ApprovalName type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApprovalName{}

// ApprovalName Approval Name Object
type ApprovalName struct {
	// Name of the approval
	Value *string `json:"value,omitempty"`
	// What locale the name of the approval is using
	Locale *string `json:"locale,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApprovalName ApprovalName

// NewApprovalName instantiates a new ApprovalName object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApprovalName() *ApprovalName {
	this := ApprovalName{}
	return &this
}

// NewApprovalNameWithDefaults instantiates a new ApprovalName object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApprovalNameWithDefaults() *ApprovalName {
	this := ApprovalName{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ApprovalName) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalName) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ApprovalName) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ApprovalName) SetValue(v string) {
	o.Value = &v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *ApprovalName) GetLocale() string {
	if o == nil || IsNil(o.Locale) {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalName) GetLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.Locale) {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *ApprovalName) HasLocale() bool {
	if o != nil && !IsNil(o.Locale) {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *ApprovalName) SetLocale(v string) {
	o.Locale = &v
}

func (o ApprovalName) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApprovalName) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Locale) {
		toSerialize["locale"] = o.Locale
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApprovalName) UnmarshalJSON(data []byte) (err error) {
	varApprovalName := _ApprovalName{}

	err = json.Unmarshal(data, &varApprovalName)

	if err != nil {
		return err
	}

	*o = ApprovalName(varApprovalName)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "locale")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApprovalName struct {
	value *ApprovalName
	isSet bool
}

func (v NullableApprovalName) Get() *ApprovalName {
	return v.value
}

func (v *NullableApprovalName) Set(val *ApprovalName) {
	v.value = val
	v.isSet = true
}

func (v NullableApprovalName) IsSet() bool {
	return v.isSet
}

func (v *NullableApprovalName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApprovalName(val *ApprovalName) *NullableApprovalName {
	return &NullableApprovalName{value: val, isSet: true}
}

func (v NullableApprovalName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApprovalName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


