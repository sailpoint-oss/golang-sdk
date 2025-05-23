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

// checks if the SearchCriteriaFiltersValueRangeLower type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchCriteriaFiltersValueRangeLower{}

// SearchCriteriaFiltersValueRangeLower struct for SearchCriteriaFiltersValueRangeLower
type SearchCriteriaFiltersValueRangeLower struct {
	// The lower bound value.
	Value *string `json:"value,omitempty"`
	// Whether the lower bound is inclusive.
	Inclusive *bool `json:"inclusive,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SearchCriteriaFiltersValueRangeLower SearchCriteriaFiltersValueRangeLower

// NewSearchCriteriaFiltersValueRangeLower instantiates a new SearchCriteriaFiltersValueRangeLower object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchCriteriaFiltersValueRangeLower() *SearchCriteriaFiltersValueRangeLower {
	this := SearchCriteriaFiltersValueRangeLower{}
	var inclusive bool = false
	this.Inclusive = &inclusive
	return &this
}

// NewSearchCriteriaFiltersValueRangeLowerWithDefaults instantiates a new SearchCriteriaFiltersValueRangeLower object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchCriteriaFiltersValueRangeLowerWithDefaults() *SearchCriteriaFiltersValueRangeLower {
	this := SearchCriteriaFiltersValueRangeLower{}
	var inclusive bool = false
	this.Inclusive = &inclusive
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SearchCriteriaFiltersValueRangeLower) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCriteriaFiltersValueRangeLower) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SearchCriteriaFiltersValueRangeLower) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *SearchCriteriaFiltersValueRangeLower) SetValue(v string) {
	o.Value = &v
}

// GetInclusive returns the Inclusive field value if set, zero value otherwise.
func (o *SearchCriteriaFiltersValueRangeLower) GetInclusive() bool {
	if o == nil || IsNil(o.Inclusive) {
		var ret bool
		return ret
	}
	return *o.Inclusive
}

// GetInclusiveOk returns a tuple with the Inclusive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCriteriaFiltersValueRangeLower) GetInclusiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Inclusive) {
		return nil, false
	}
	return o.Inclusive, true
}

// HasInclusive returns a boolean if a field has been set.
func (o *SearchCriteriaFiltersValueRangeLower) HasInclusive() bool {
	if o != nil && !IsNil(o.Inclusive) {
		return true
	}

	return false
}

// SetInclusive gets a reference to the given bool and assigns it to the Inclusive field.
func (o *SearchCriteriaFiltersValueRangeLower) SetInclusive(v bool) {
	o.Inclusive = &v
}

func (o SearchCriteriaFiltersValueRangeLower) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchCriteriaFiltersValueRangeLower) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Inclusive) {
		toSerialize["inclusive"] = o.Inclusive
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SearchCriteriaFiltersValueRangeLower) UnmarshalJSON(data []byte) (err error) {
	varSearchCriteriaFiltersValueRangeLower := _SearchCriteriaFiltersValueRangeLower{}

	err = json.Unmarshal(data, &varSearchCriteriaFiltersValueRangeLower)

	if err != nil {
		return err
	}

	*o = SearchCriteriaFiltersValueRangeLower(varSearchCriteriaFiltersValueRangeLower)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "inclusive")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSearchCriteriaFiltersValueRangeLower struct {
	value *SearchCriteriaFiltersValueRangeLower
	isSet bool
}

func (v NullableSearchCriteriaFiltersValueRangeLower) Get() *SearchCriteriaFiltersValueRangeLower {
	return v.value
}

func (v *NullableSearchCriteriaFiltersValueRangeLower) Set(val *SearchCriteriaFiltersValueRangeLower) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchCriteriaFiltersValueRangeLower) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchCriteriaFiltersValueRangeLower) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchCriteriaFiltersValueRangeLower(val *SearchCriteriaFiltersValueRangeLower) *NullableSearchCriteriaFiltersValueRangeLower {
	return &NullableSearchCriteriaFiltersValueRangeLower{value: val, isSet: true}
}

func (v NullableSearchCriteriaFiltersValueRangeLower) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchCriteriaFiltersValueRangeLower) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


