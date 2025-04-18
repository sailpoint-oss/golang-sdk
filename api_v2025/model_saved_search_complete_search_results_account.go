/*
Identity Security Cloud V2025 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: v2025
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2025

import (
	"encoding/json"
	"fmt"
)

// checks if the SavedSearchCompleteSearchResultsAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SavedSearchCompleteSearchResultsAccount{}

// SavedSearchCompleteSearchResultsAccount A table of accounts that match the search criteria.
type SavedSearchCompleteSearchResultsAccount struct {
	// The number of rows in the table.
	Count string `json:"count"`
	// The type of object represented in the table.
	Noun string `json:"noun"`
	// A sample of the data in the table.
	Preview [][]string `json:"preview"`
	AdditionalProperties map[string]interface{}
}

type _SavedSearchCompleteSearchResultsAccount SavedSearchCompleteSearchResultsAccount

// NewSavedSearchCompleteSearchResultsAccount instantiates a new SavedSearchCompleteSearchResultsAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSavedSearchCompleteSearchResultsAccount(count string, noun string, preview [][]string) *SavedSearchCompleteSearchResultsAccount {
	this := SavedSearchCompleteSearchResultsAccount{}
	this.Count = count
	this.Noun = noun
	this.Preview = preview
	return &this
}

// NewSavedSearchCompleteSearchResultsAccountWithDefaults instantiates a new SavedSearchCompleteSearchResultsAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSavedSearchCompleteSearchResultsAccountWithDefaults() *SavedSearchCompleteSearchResultsAccount {
	this := SavedSearchCompleteSearchResultsAccount{}
	return &this
}

// GetCount returns the Count field value
func (o *SavedSearchCompleteSearchResultsAccount) GetCount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *SavedSearchCompleteSearchResultsAccount) GetCountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *SavedSearchCompleteSearchResultsAccount) SetCount(v string) {
	o.Count = v
}

// GetNoun returns the Noun field value
func (o *SavedSearchCompleteSearchResultsAccount) GetNoun() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Noun
}

// GetNounOk returns a tuple with the Noun field value
// and a boolean to check if the value has been set.
func (o *SavedSearchCompleteSearchResultsAccount) GetNounOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Noun, true
}

// SetNoun sets field value
func (o *SavedSearchCompleteSearchResultsAccount) SetNoun(v string) {
	o.Noun = v
}

// GetPreview returns the Preview field value
func (o *SavedSearchCompleteSearchResultsAccount) GetPreview() [][]string {
	if o == nil {
		var ret [][]string
		return ret
	}

	return o.Preview
}

// GetPreviewOk returns a tuple with the Preview field value
// and a boolean to check if the value has been set.
func (o *SavedSearchCompleteSearchResultsAccount) GetPreviewOk() ([][]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Preview, true
}

// SetPreview sets field value
func (o *SavedSearchCompleteSearchResultsAccount) SetPreview(v [][]string) {
	o.Preview = v
}

func (o SavedSearchCompleteSearchResultsAccount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SavedSearchCompleteSearchResultsAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["count"] = o.Count
	toSerialize["noun"] = o.Noun
	toSerialize["preview"] = o.Preview

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SavedSearchCompleteSearchResultsAccount) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"count",
		"noun",
		"preview",
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

	varSavedSearchCompleteSearchResultsAccount := _SavedSearchCompleteSearchResultsAccount{}

	err = json.Unmarshal(data, &varSavedSearchCompleteSearchResultsAccount)

	if err != nil {
		return err
	}

	*o = SavedSearchCompleteSearchResultsAccount(varSavedSearchCompleteSearchResultsAccount)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "count")
		delete(additionalProperties, "noun")
		delete(additionalProperties, "preview")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSavedSearchCompleteSearchResultsAccount struct {
	value *SavedSearchCompleteSearchResultsAccount
	isSet bool
}

func (v NullableSavedSearchCompleteSearchResultsAccount) Get() *SavedSearchCompleteSearchResultsAccount {
	return v.value
}

func (v *NullableSavedSearchCompleteSearchResultsAccount) Set(val *SavedSearchCompleteSearchResultsAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableSavedSearchCompleteSearchResultsAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableSavedSearchCompleteSearchResultsAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSavedSearchCompleteSearchResultsAccount(val *SavedSearchCompleteSearchResultsAccount) *NullableSavedSearchCompleteSearchResultsAccount {
	return &NullableSavedSearchCompleteSearchResultsAccount{value: val, isSet: true}
}

func (v NullableSavedSearchCompleteSearchResultsAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSavedSearchCompleteSearchResultsAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


