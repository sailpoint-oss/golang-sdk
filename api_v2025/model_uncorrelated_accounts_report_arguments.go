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

// checks if the UncorrelatedAccountsReportArguments type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UncorrelatedAccountsReportArguments{}

// UncorrelatedAccountsReportArguments Arguments for Uncorrelated Accounts report (UNCORRELATED_ACCOUNTS)
type UncorrelatedAccountsReportArguments struct {
	// Output report file formats. These are formats for calling GET endpoint as query parameter 'fileFormat'.  In case report won't have this argument there will be ['CSV', 'PDF'] as default.
	SelectedFormats []string `json:"selectedFormats,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UncorrelatedAccountsReportArguments UncorrelatedAccountsReportArguments

// NewUncorrelatedAccountsReportArguments instantiates a new UncorrelatedAccountsReportArguments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUncorrelatedAccountsReportArguments() *UncorrelatedAccountsReportArguments {
	this := UncorrelatedAccountsReportArguments{}
	return &this
}

// NewUncorrelatedAccountsReportArgumentsWithDefaults instantiates a new UncorrelatedAccountsReportArguments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUncorrelatedAccountsReportArgumentsWithDefaults() *UncorrelatedAccountsReportArguments {
	this := UncorrelatedAccountsReportArguments{}
	return &this
}

// GetSelectedFormats returns the SelectedFormats field value if set, zero value otherwise.
func (o *UncorrelatedAccountsReportArguments) GetSelectedFormats() []string {
	if o == nil || IsNil(o.SelectedFormats) {
		var ret []string
		return ret
	}
	return o.SelectedFormats
}

// GetSelectedFormatsOk returns a tuple with the SelectedFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UncorrelatedAccountsReportArguments) GetSelectedFormatsOk() ([]string, bool) {
	if o == nil || IsNil(o.SelectedFormats) {
		return nil, false
	}
	return o.SelectedFormats, true
}

// HasSelectedFormats returns a boolean if a field has been set.
func (o *UncorrelatedAccountsReportArguments) HasSelectedFormats() bool {
	if o != nil && !IsNil(o.SelectedFormats) {
		return true
	}

	return false
}

// SetSelectedFormats gets a reference to the given []string and assigns it to the SelectedFormats field.
func (o *UncorrelatedAccountsReportArguments) SetSelectedFormats(v []string) {
	o.SelectedFormats = v
}

func (o UncorrelatedAccountsReportArguments) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UncorrelatedAccountsReportArguments) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SelectedFormats) {
		toSerialize["selectedFormats"] = o.SelectedFormats
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UncorrelatedAccountsReportArguments) UnmarshalJSON(data []byte) (err error) {
	varUncorrelatedAccountsReportArguments := _UncorrelatedAccountsReportArguments{}

	err = json.Unmarshal(data, &varUncorrelatedAccountsReportArguments)

	if err != nil {
		return err
	}

	*o = UncorrelatedAccountsReportArguments(varUncorrelatedAccountsReportArguments)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "selectedFormats")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUncorrelatedAccountsReportArguments struct {
	value *UncorrelatedAccountsReportArguments
	isSet bool
}

func (v NullableUncorrelatedAccountsReportArguments) Get() *UncorrelatedAccountsReportArguments {
	return v.value
}

func (v *NullableUncorrelatedAccountsReportArguments) Set(val *UncorrelatedAccountsReportArguments) {
	v.value = val
	v.isSet = true
}

func (v NullableUncorrelatedAccountsReportArguments) IsSet() bool {
	return v.isSet
}

func (v *NullableUncorrelatedAccountsReportArguments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUncorrelatedAccountsReportArguments(val *UncorrelatedAccountsReportArguments) *NullableUncorrelatedAccountsReportArguments {
	return &NullableUncorrelatedAccountsReportArguments{value: val, isSet: true}
}

func (v NullableUncorrelatedAccountsReportArguments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUncorrelatedAccountsReportArguments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


