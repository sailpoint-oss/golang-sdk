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

// checks if the OutlierFeatureTranslation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OutlierFeatureTranslation{}

// OutlierFeatureTranslation struct for OutlierFeatureTranslation
type OutlierFeatureTranslation struct {
	DisplayName *TranslationMessage `json:"displayName,omitempty"`
	Description *TranslationMessage `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OutlierFeatureTranslation OutlierFeatureTranslation

// NewOutlierFeatureTranslation instantiates a new OutlierFeatureTranslation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutlierFeatureTranslation() *OutlierFeatureTranslation {
	this := OutlierFeatureTranslation{}
	return &this
}

// NewOutlierFeatureTranslationWithDefaults instantiates a new OutlierFeatureTranslation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutlierFeatureTranslationWithDefaults() *OutlierFeatureTranslation {
	this := OutlierFeatureTranslation{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *OutlierFeatureTranslation) GetDisplayName() TranslationMessage {
	if o == nil || IsNil(o.DisplayName) {
		var ret TranslationMessage
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutlierFeatureTranslation) GetDisplayNameOk() (*TranslationMessage, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *OutlierFeatureTranslation) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given TranslationMessage and assigns it to the DisplayName field.
func (o *OutlierFeatureTranslation) SetDisplayName(v TranslationMessage) {
	o.DisplayName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OutlierFeatureTranslation) GetDescription() TranslationMessage {
	if o == nil || IsNil(o.Description) {
		var ret TranslationMessage
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutlierFeatureTranslation) GetDescriptionOk() (*TranslationMessage, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OutlierFeatureTranslation) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given TranslationMessage and assigns it to the Description field.
func (o *OutlierFeatureTranslation) SetDescription(v TranslationMessage) {
	o.Description = &v
}

func (o OutlierFeatureTranslation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OutlierFeatureTranslation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OutlierFeatureTranslation) UnmarshalJSON(data []byte) (err error) {
	varOutlierFeatureTranslation := _OutlierFeatureTranslation{}

	err = json.Unmarshal(data, &varOutlierFeatureTranslation)

	if err != nil {
		return err
	}

	*o = OutlierFeatureTranslation(varOutlierFeatureTranslation)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOutlierFeatureTranslation struct {
	value *OutlierFeatureTranslation
	isSet bool
}

func (v NullableOutlierFeatureTranslation) Get() *OutlierFeatureTranslation {
	return v.value
}

func (v *NullableOutlierFeatureTranslation) Set(val *OutlierFeatureTranslation) {
	v.value = val
	v.isSet = true
}

func (v NullableOutlierFeatureTranslation) IsSet() bool {
	return v.isSet
}

func (v *NullableOutlierFeatureTranslation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutlierFeatureTranslation(val *OutlierFeatureTranslation) *NullableOutlierFeatureTranslation {
	return &NullableOutlierFeatureTranslation{value: val, isSet: true}
}

func (v NullableOutlierFeatureTranslation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutlierFeatureTranslation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


