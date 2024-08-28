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

// checks if the ImportAccountsRequest1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImportAccountsRequest1{}

// ImportAccountsRequest1 This content type is provided for compatibility with services that don't support multipart/form-data, such as SailPoint Workflows. This content type does not support files, so it can only be used for direct connect sources.
type ImportAccountsRequest1 struct {
	// Use this flag to reprocess every account whether or not the data has changed.
	DisableOptimization *string `json:"disableOptimization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ImportAccountsRequest1 ImportAccountsRequest1

// NewImportAccountsRequest1 instantiates a new ImportAccountsRequest1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportAccountsRequest1() *ImportAccountsRequest1 {
	this := ImportAccountsRequest1{}
	return &this
}

// NewImportAccountsRequest1WithDefaults instantiates a new ImportAccountsRequest1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportAccountsRequest1WithDefaults() *ImportAccountsRequest1 {
	this := ImportAccountsRequest1{}
	return &this
}

// GetDisableOptimization returns the DisableOptimization field value if set, zero value otherwise.
func (o *ImportAccountsRequest1) GetDisableOptimization() string {
	if o == nil || IsNil(o.DisableOptimization) {
		var ret string
		return ret
	}
	return *o.DisableOptimization
}

// GetDisableOptimizationOk returns a tuple with the DisableOptimization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportAccountsRequest1) GetDisableOptimizationOk() (*string, bool) {
	if o == nil || IsNil(o.DisableOptimization) {
		return nil, false
	}
	return o.DisableOptimization, true
}

// HasDisableOptimization returns a boolean if a field has been set.
func (o *ImportAccountsRequest1) HasDisableOptimization() bool {
	if o != nil && !IsNil(o.DisableOptimization) {
		return true
	}

	return false
}

// SetDisableOptimization gets a reference to the given string and assigns it to the DisableOptimization field.
func (o *ImportAccountsRequest1) SetDisableOptimization(v string) {
	o.DisableOptimization = &v
}

func (o ImportAccountsRequest1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImportAccountsRequest1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DisableOptimization) {
		toSerialize["disableOptimization"] = o.DisableOptimization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ImportAccountsRequest1) UnmarshalJSON(data []byte) (err error) {
	varImportAccountsRequest1 := _ImportAccountsRequest1{}

	err = json.Unmarshal(data, &varImportAccountsRequest1)

	if err != nil {
		return err
	}

	*o = ImportAccountsRequest1(varImportAccountsRequest1)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "disableOptimization")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableImportAccountsRequest1 struct {
	value *ImportAccountsRequest1
	isSet bool
}

func (v NullableImportAccountsRequest1) Get() *ImportAccountsRequest1 {
	return v.value
}

func (v *NullableImportAccountsRequest1) Set(val *ImportAccountsRequest1) {
	v.value = val
	v.isSet = true
}

func (v NullableImportAccountsRequest1) IsSet() bool {
	return v.isSet
}

func (v *NullableImportAccountsRequest1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportAccountsRequest1(val *ImportAccountsRequest1) *NullableImportAccountsRequest1 {
	return &NullableImportAccountsRequest1{value: val, isSet: true}
}

func (v NullableImportAccountsRequest1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportAccountsRequest1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

