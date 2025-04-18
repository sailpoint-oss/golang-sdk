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

// checks if the AccountUsage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountUsage{}

// AccountUsage struct for AccountUsage
type AccountUsage struct {
	// The first day of the month for which activity is aggregated.
	Date *string `json:"date,omitempty"`
	// The number of days within the month that the account was active in a source.
	Count *int64 `json:"count,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountUsage AccountUsage

// NewAccountUsage instantiates a new AccountUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountUsage() *AccountUsage {
	this := AccountUsage{}
	return &this
}

// NewAccountUsageWithDefaults instantiates a new AccountUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountUsageWithDefaults() *AccountUsage {
	this := AccountUsage{}
	return &this
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *AccountUsage) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUsage) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *AccountUsage) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *AccountUsage) SetDate(v string) {
	o.Date = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *AccountUsage) GetCount() int64 {
	if o == nil || IsNil(o.Count) {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUsage) GetCountOk() (*int64, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *AccountUsage) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *AccountUsage) SetCount(v int64) {
	o.Count = &v
}

func (o AccountUsage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountUsage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountUsage) UnmarshalJSON(data []byte) (err error) {
	varAccountUsage := _AccountUsage{}

	err = json.Unmarshal(data, &varAccountUsage)

	if err != nil {
		return err
	}

	*o = AccountUsage(varAccountUsage)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "date")
		delete(additionalProperties, "count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountUsage struct {
	value *AccountUsage
	isSet bool
}

func (v NullableAccountUsage) Get() *AccountUsage {
	return v.value
}

func (v *NullableAccountUsage) Set(val *AccountUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountUsage(val *AccountUsage) *NullableAccountUsage {
	return &NullableAccountUsage{value: val, isSet: true}
}

func (v NullableAccountUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


