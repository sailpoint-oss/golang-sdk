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

// checks if the UpdateAccessProfilesInBulk412Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateAccessProfilesInBulk412Response{}

// UpdateAccessProfilesInBulk412Response struct for UpdateAccessProfilesInBulk412Response
type UpdateAccessProfilesInBulk412Response struct {
	// A message describing the error
	Message map[string]interface{} `json:"message,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateAccessProfilesInBulk412Response UpdateAccessProfilesInBulk412Response

// NewUpdateAccessProfilesInBulk412Response instantiates a new UpdateAccessProfilesInBulk412Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAccessProfilesInBulk412Response() *UpdateAccessProfilesInBulk412Response {
	this := UpdateAccessProfilesInBulk412Response{}
	return &this
}

// NewUpdateAccessProfilesInBulk412ResponseWithDefaults instantiates a new UpdateAccessProfilesInBulk412Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAccessProfilesInBulk412ResponseWithDefaults() *UpdateAccessProfilesInBulk412Response {
	this := UpdateAccessProfilesInBulk412Response{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *UpdateAccessProfilesInBulk412Response) GetMessage() map[string]interface{} {
	if o == nil || IsNil(o.Message) {
		var ret map[string]interface{}
		return ret
	}
	return o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccessProfilesInBulk412Response) GetMessageOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Message) {
		return map[string]interface{}{}, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *UpdateAccessProfilesInBulk412Response) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given map[string]interface{} and assigns it to the Message field.
func (o *UpdateAccessProfilesInBulk412Response) SetMessage(v map[string]interface{}) {
	o.Message = v
}

func (o UpdateAccessProfilesInBulk412Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateAccessProfilesInBulk412Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateAccessProfilesInBulk412Response) UnmarshalJSON(data []byte) (err error) {
	varUpdateAccessProfilesInBulk412Response := _UpdateAccessProfilesInBulk412Response{}

	err = json.Unmarshal(data, &varUpdateAccessProfilesInBulk412Response)

	if err != nil {
		return err
	}

	*o = UpdateAccessProfilesInBulk412Response(varUpdateAccessProfilesInBulk412Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "message")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateAccessProfilesInBulk412Response struct {
	value *UpdateAccessProfilesInBulk412Response
	isSet bool
}

func (v NullableUpdateAccessProfilesInBulk412Response) Get() *UpdateAccessProfilesInBulk412Response {
	return v.value
}

func (v *NullableUpdateAccessProfilesInBulk412Response) Set(val *UpdateAccessProfilesInBulk412Response) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAccessProfilesInBulk412Response) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAccessProfilesInBulk412Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAccessProfilesInBulk412Response(val *UpdateAccessProfilesInBulk412Response) *NullableUpdateAccessProfilesInBulk412Response {
	return &NullableUpdateAccessProfilesInBulk412Response{value: val, isSet: true}
}

func (v NullableUpdateAccessProfilesInBulk412Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAccessProfilesInBulk412Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


