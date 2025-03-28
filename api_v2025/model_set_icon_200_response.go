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

// checks if the SetIcon200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetIcon200Response{}

// SetIcon200Response struct for SetIcon200Response
type SetIcon200Response struct {
	// url to file with icon
	Icon *string `json:"icon,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SetIcon200Response SetIcon200Response

// NewSetIcon200Response instantiates a new SetIcon200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetIcon200Response() *SetIcon200Response {
	this := SetIcon200Response{}
	return &this
}

// NewSetIcon200ResponseWithDefaults instantiates a new SetIcon200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetIcon200ResponseWithDefaults() *SetIcon200Response {
	this := SetIcon200Response{}
	return &this
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *SetIcon200Response) GetIcon() string {
	if o == nil || IsNil(o.Icon) {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetIcon200Response) GetIconOk() (*string, bool) {
	if o == nil || IsNil(o.Icon) {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *SetIcon200Response) HasIcon() bool {
	if o != nil && !IsNil(o.Icon) {
		return true
	}

	return false
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *SetIcon200Response) SetIcon(v string) {
	o.Icon = &v
}

func (o SetIcon200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetIcon200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Icon) {
		toSerialize["icon"] = o.Icon
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SetIcon200Response) UnmarshalJSON(data []byte) (err error) {
	varSetIcon200Response := _SetIcon200Response{}

	err = json.Unmarshal(data, &varSetIcon200Response)

	if err != nil {
		return err
	}

	*o = SetIcon200Response(varSetIcon200Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "icon")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSetIcon200Response struct {
	value *SetIcon200Response
	isSet bool
}

func (v NullableSetIcon200Response) Get() *SetIcon200Response {
	return v.value
}

func (v *NullableSetIcon200Response) Set(val *SetIcon200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableSetIcon200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableSetIcon200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetIcon200Response(val *SetIcon200Response) *NullableSetIcon200Response {
	return &NullableSetIcon200Response{value: val, isSet: true}
}

func (v NullableSetIcon200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetIcon200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


