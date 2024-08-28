/*
Identity Security Cloud V2024 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: v2024
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2024

import (
	"encoding/json"
	"fmt"
)

// checks if the TextQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TextQuery{}

// TextQuery Query parameters used to construct an Elasticsearch text query object.
type TextQuery struct {
	// Words or characters that specify a particular thing to be searched for.
	Terms []string `json:"terms"`
	// The fields to be searched.
	Fields []string `json:"fields"`
	// Indicates that at least one of the terms must be found in the specified fields;  otherwise, all terms must be found.
	MatchAny *bool `json:"matchAny,omitempty"`
	// Indicates that the terms can be located anywhere in the specified fields;  otherwise, the fields must begin with the terms.
	Contains *bool `json:"contains,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TextQuery TextQuery

// NewTextQuery instantiates a new TextQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTextQuery(terms []string, fields []string) *TextQuery {
	this := TextQuery{}
	this.Terms = terms
	this.Fields = fields
	var matchAny bool = false
	this.MatchAny = &matchAny
	var contains bool = false
	this.Contains = &contains
	return &this
}

// NewTextQueryWithDefaults instantiates a new TextQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTextQueryWithDefaults() *TextQuery {
	this := TextQuery{}
	var matchAny bool = false
	this.MatchAny = &matchAny
	var contains bool = false
	this.Contains = &contains
	return &this
}

// GetTerms returns the Terms field value
func (o *TextQuery) GetTerms() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Terms
}

// GetTermsOk returns a tuple with the Terms field value
// and a boolean to check if the value has been set.
func (o *TextQuery) GetTermsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Terms, true
}

// SetTerms sets field value
func (o *TextQuery) SetTerms(v []string) {
	o.Terms = v
}

// GetFields returns the Fields field value
func (o *TextQuery) GetFields() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *TextQuery) GetFieldsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Fields, true
}

// SetFields sets field value
func (o *TextQuery) SetFields(v []string) {
	o.Fields = v
}

// GetMatchAny returns the MatchAny field value if set, zero value otherwise.
func (o *TextQuery) GetMatchAny() bool {
	if o == nil || IsNil(o.MatchAny) {
		var ret bool
		return ret
	}
	return *o.MatchAny
}

// GetMatchAnyOk returns a tuple with the MatchAny field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextQuery) GetMatchAnyOk() (*bool, bool) {
	if o == nil || IsNil(o.MatchAny) {
		return nil, false
	}
	return o.MatchAny, true
}

// HasMatchAny returns a boolean if a field has been set.
func (o *TextQuery) HasMatchAny() bool {
	if o != nil && !IsNil(o.MatchAny) {
		return true
	}

	return false
}

// SetMatchAny gets a reference to the given bool and assigns it to the MatchAny field.
func (o *TextQuery) SetMatchAny(v bool) {
	o.MatchAny = &v
}

// GetContains returns the Contains field value if set, zero value otherwise.
func (o *TextQuery) GetContains() bool {
	if o == nil || IsNil(o.Contains) {
		var ret bool
		return ret
	}
	return *o.Contains
}

// GetContainsOk returns a tuple with the Contains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextQuery) GetContainsOk() (*bool, bool) {
	if o == nil || IsNil(o.Contains) {
		return nil, false
	}
	return o.Contains, true
}

// HasContains returns a boolean if a field has been set.
func (o *TextQuery) HasContains() bool {
	if o != nil && !IsNil(o.Contains) {
		return true
	}

	return false
}

// SetContains gets a reference to the given bool and assigns it to the Contains field.
func (o *TextQuery) SetContains(v bool) {
	o.Contains = &v
}

func (o TextQuery) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TextQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["terms"] = o.Terms
	toSerialize["fields"] = o.Fields
	if !IsNil(o.MatchAny) {
		toSerialize["matchAny"] = o.MatchAny
	}
	if !IsNil(o.Contains) {
		toSerialize["contains"] = o.Contains
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TextQuery) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"terms",
		"fields",
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

	varTextQuery := _TextQuery{}

	err = json.Unmarshal(data, &varTextQuery)

	if err != nil {
		return err
	}

	*o = TextQuery(varTextQuery)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "terms")
		delete(additionalProperties, "fields")
		delete(additionalProperties, "matchAny")
		delete(additionalProperties, "contains")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTextQuery struct {
	value *TextQuery
	isSet bool
}

func (v NullableTextQuery) Get() *TextQuery {
	return v.value
}

func (v *NullableTextQuery) Set(val *TextQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableTextQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableTextQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTextQuery(val *TextQuery) *NullableTextQuery {
	return &NullableTextQuery{value: val, isSet: true}
}

func (v NullableTextQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTextQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

