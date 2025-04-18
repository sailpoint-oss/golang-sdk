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

// checks if the Filter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Filter{}

// Filter struct for Filter
type Filter struct {
	Type *FilterType `json:"type,omitempty"`
	Range *Range `json:"range,omitempty"`
	// The terms to be filtered.
	Terms []string `json:"terms,omitempty"`
	// Indicates if the filter excludes results.
	Exclude *bool `json:"exclude,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Filter Filter

// NewFilter instantiates a new Filter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilter() *Filter {
	this := Filter{}
	var exclude bool = false
	this.Exclude = &exclude
	return &this
}

// NewFilterWithDefaults instantiates a new Filter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilterWithDefaults() *Filter {
	this := Filter{}
	var exclude bool = false
	this.Exclude = &exclude
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Filter) GetType() FilterType {
	if o == nil || IsNil(o.Type) {
		var ret FilterType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Filter) GetTypeOk() (*FilterType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Filter) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given FilterType and assigns it to the Type field.
func (o *Filter) SetType(v FilterType) {
	o.Type = &v
}

// GetRange returns the Range field value if set, zero value otherwise.
func (o *Filter) GetRange() Range {
	if o == nil || IsNil(o.Range) {
		var ret Range
		return ret
	}
	return *o.Range
}

// GetRangeOk returns a tuple with the Range field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Filter) GetRangeOk() (*Range, bool) {
	if o == nil || IsNil(o.Range) {
		return nil, false
	}
	return o.Range, true
}

// HasRange returns a boolean if a field has been set.
func (o *Filter) HasRange() bool {
	if o != nil && !IsNil(o.Range) {
		return true
	}

	return false
}

// SetRange gets a reference to the given Range and assigns it to the Range field.
func (o *Filter) SetRange(v Range) {
	o.Range = &v
}

// GetTerms returns the Terms field value if set, zero value otherwise.
func (o *Filter) GetTerms() []string {
	if o == nil || IsNil(o.Terms) {
		var ret []string
		return ret
	}
	return o.Terms
}

// GetTermsOk returns a tuple with the Terms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Filter) GetTermsOk() ([]string, bool) {
	if o == nil || IsNil(o.Terms) {
		return nil, false
	}
	return o.Terms, true
}

// HasTerms returns a boolean if a field has been set.
func (o *Filter) HasTerms() bool {
	if o != nil && !IsNil(o.Terms) {
		return true
	}

	return false
}

// SetTerms gets a reference to the given []string and assigns it to the Terms field.
func (o *Filter) SetTerms(v []string) {
	o.Terms = v
}

// GetExclude returns the Exclude field value if set, zero value otherwise.
func (o *Filter) GetExclude() bool {
	if o == nil || IsNil(o.Exclude) {
		var ret bool
		return ret
	}
	return *o.Exclude
}

// GetExcludeOk returns a tuple with the Exclude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Filter) GetExcludeOk() (*bool, bool) {
	if o == nil || IsNil(o.Exclude) {
		return nil, false
	}
	return o.Exclude, true
}

// HasExclude returns a boolean if a field has been set.
func (o *Filter) HasExclude() bool {
	if o != nil && !IsNil(o.Exclude) {
		return true
	}

	return false
}

// SetExclude gets a reference to the given bool and assigns it to the Exclude field.
func (o *Filter) SetExclude(v bool) {
	o.Exclude = &v
}

func (o Filter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Filter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Range) {
		toSerialize["range"] = o.Range
	}
	if !IsNil(o.Terms) {
		toSerialize["terms"] = o.Terms
	}
	if !IsNil(o.Exclude) {
		toSerialize["exclude"] = o.Exclude
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Filter) UnmarshalJSON(data []byte) (err error) {
	varFilter := _Filter{}

	err = json.Unmarshal(data, &varFilter)

	if err != nil {
		return err
	}

	*o = Filter(varFilter)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "range")
		delete(additionalProperties, "terms")
		delete(additionalProperties, "exclude")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFilter struct {
	value *Filter
	isSet bool
}

func (v NullableFilter) Get() *Filter {
	return v.value
}

func (v *NullableFilter) Set(val *Filter) {
	v.value = val
	v.isSet = true
}

func (v NullableFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilter(val *Filter) *NullableFilter {
	return &NullableFilter{value: val, isSet: true}
}

func (v NullableFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


