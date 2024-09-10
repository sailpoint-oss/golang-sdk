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

// checks if the SelectorAccountMatchConfigMatchExpression type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SelectorAccountMatchConfigMatchExpression{}

// SelectorAccountMatchConfigMatchExpression struct for SelectorAccountMatchConfigMatchExpression
type SelectorAccountMatchConfigMatchExpression struct {
	MatchTerms []MatchTerm `json:"matchTerms,omitempty"`
	// If it is AND operators for match terms
	And *bool `json:"and,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SelectorAccountMatchConfigMatchExpression SelectorAccountMatchConfigMatchExpression

// NewSelectorAccountMatchConfigMatchExpression instantiates a new SelectorAccountMatchConfigMatchExpression object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSelectorAccountMatchConfigMatchExpression() *SelectorAccountMatchConfigMatchExpression {
	this := SelectorAccountMatchConfigMatchExpression{}
	var and bool = true
	this.And = &and
	return &this
}

// NewSelectorAccountMatchConfigMatchExpressionWithDefaults instantiates a new SelectorAccountMatchConfigMatchExpression object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSelectorAccountMatchConfigMatchExpressionWithDefaults() *SelectorAccountMatchConfigMatchExpression {
	this := SelectorAccountMatchConfigMatchExpression{}
	var and bool = true
	this.And = &and
	return &this
}

// GetMatchTerms returns the MatchTerms field value if set, zero value otherwise.
func (o *SelectorAccountMatchConfigMatchExpression) GetMatchTerms() []MatchTerm {
	if o == nil || IsNil(o.MatchTerms) {
		var ret []MatchTerm
		return ret
	}
	return o.MatchTerms
}

// GetMatchTermsOk returns a tuple with the MatchTerms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelectorAccountMatchConfigMatchExpression) GetMatchTermsOk() ([]MatchTerm, bool) {
	if o == nil || IsNil(o.MatchTerms) {
		return nil, false
	}
	return o.MatchTerms, true
}

// HasMatchTerms returns a boolean if a field has been set.
func (o *SelectorAccountMatchConfigMatchExpression) HasMatchTerms() bool {
	if o != nil && !IsNil(o.MatchTerms) {
		return true
	}

	return false
}

// SetMatchTerms gets a reference to the given []MatchTerm and assigns it to the MatchTerms field.
func (o *SelectorAccountMatchConfigMatchExpression) SetMatchTerms(v []MatchTerm) {
	o.MatchTerms = v
}

// GetAnd returns the And field value if set, zero value otherwise.
func (o *SelectorAccountMatchConfigMatchExpression) GetAnd() bool {
	if o == nil || IsNil(o.And) {
		var ret bool
		return ret
	}
	return *o.And
}

// GetAndOk returns a tuple with the And field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelectorAccountMatchConfigMatchExpression) GetAndOk() (*bool, bool) {
	if o == nil || IsNil(o.And) {
		return nil, false
	}
	return o.And, true
}

// HasAnd returns a boolean if a field has been set.
func (o *SelectorAccountMatchConfigMatchExpression) HasAnd() bool {
	if o != nil && !IsNil(o.And) {
		return true
	}

	return false
}

// SetAnd gets a reference to the given bool and assigns it to the And field.
func (o *SelectorAccountMatchConfigMatchExpression) SetAnd(v bool) {
	o.And = &v
}

func (o SelectorAccountMatchConfigMatchExpression) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SelectorAccountMatchConfigMatchExpression) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MatchTerms) {
		toSerialize["matchTerms"] = o.MatchTerms
	}
	if !IsNil(o.And) {
		toSerialize["and"] = o.And
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SelectorAccountMatchConfigMatchExpression) UnmarshalJSON(data []byte) (err error) {
	varSelectorAccountMatchConfigMatchExpression := _SelectorAccountMatchConfigMatchExpression{}

	err = json.Unmarshal(data, &varSelectorAccountMatchConfigMatchExpression)

	if err != nil {
		return err
	}

	*o = SelectorAccountMatchConfigMatchExpression(varSelectorAccountMatchConfigMatchExpression)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "matchTerms")
		delete(additionalProperties, "and")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSelectorAccountMatchConfigMatchExpression struct {
	value *SelectorAccountMatchConfigMatchExpression
	isSet bool
}

func (v NullableSelectorAccountMatchConfigMatchExpression) Get() *SelectorAccountMatchConfigMatchExpression {
	return v.value
}

func (v *NullableSelectorAccountMatchConfigMatchExpression) Set(val *SelectorAccountMatchConfigMatchExpression) {
	v.value = val
	v.isSet = true
}

func (v NullableSelectorAccountMatchConfigMatchExpression) IsSet() bool {
	return v.isSet
}

func (v *NullableSelectorAccountMatchConfigMatchExpression) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelectorAccountMatchConfigMatchExpression(val *SelectorAccountMatchConfigMatchExpression) *NullableSelectorAccountMatchConfigMatchExpression {
	return &NullableSelectorAccountMatchConfigMatchExpression{value: val, isSet: true}
}

func (v NullableSelectorAccountMatchConfigMatchExpression) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelectorAccountMatchConfigMatchExpression) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

