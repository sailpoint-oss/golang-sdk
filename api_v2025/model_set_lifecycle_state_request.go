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

// checks if the SetLifecycleStateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetLifecycleStateRequest{}

// SetLifecycleStateRequest struct for SetLifecycleStateRequest
type SetLifecycleStateRequest struct {
	// ID of the lifecycle state to set.
	LifecycleStateId *string `json:"lifecycleStateId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SetLifecycleStateRequest SetLifecycleStateRequest

// NewSetLifecycleStateRequest instantiates a new SetLifecycleStateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetLifecycleStateRequest() *SetLifecycleStateRequest {
	this := SetLifecycleStateRequest{}
	return &this
}

// NewSetLifecycleStateRequestWithDefaults instantiates a new SetLifecycleStateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetLifecycleStateRequestWithDefaults() *SetLifecycleStateRequest {
	this := SetLifecycleStateRequest{}
	return &this
}

// GetLifecycleStateId returns the LifecycleStateId field value if set, zero value otherwise.
func (o *SetLifecycleStateRequest) GetLifecycleStateId() string {
	if o == nil || IsNil(o.LifecycleStateId) {
		var ret string
		return ret
	}
	return *o.LifecycleStateId
}

// GetLifecycleStateIdOk returns a tuple with the LifecycleStateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetLifecycleStateRequest) GetLifecycleStateIdOk() (*string, bool) {
	if o == nil || IsNil(o.LifecycleStateId) {
		return nil, false
	}
	return o.LifecycleStateId, true
}

// HasLifecycleStateId returns a boolean if a field has been set.
func (o *SetLifecycleStateRequest) HasLifecycleStateId() bool {
	if o != nil && !IsNil(o.LifecycleStateId) {
		return true
	}

	return false
}

// SetLifecycleStateId gets a reference to the given string and assigns it to the LifecycleStateId field.
func (o *SetLifecycleStateRequest) SetLifecycleStateId(v string) {
	o.LifecycleStateId = &v
}

func (o SetLifecycleStateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetLifecycleStateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LifecycleStateId) {
		toSerialize["lifecycleStateId"] = o.LifecycleStateId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SetLifecycleStateRequest) UnmarshalJSON(data []byte) (err error) {
	varSetLifecycleStateRequest := _SetLifecycleStateRequest{}

	err = json.Unmarshal(data, &varSetLifecycleStateRequest)

	if err != nil {
		return err
	}

	*o = SetLifecycleStateRequest(varSetLifecycleStateRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "lifecycleStateId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSetLifecycleStateRequest struct {
	value *SetLifecycleStateRequest
	isSet bool
}

func (v NullableSetLifecycleStateRequest) Get() *SetLifecycleStateRequest {
	return v.value
}

func (v *NullableSetLifecycleStateRequest) Set(val *SetLifecycleStateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSetLifecycleStateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSetLifecycleStateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetLifecycleStateRequest(val *SetLifecycleStateRequest) *NullableSetLifecycleStateRequest {
	return &NullableSetLifecycleStateRequest{value: val, isSet: true}
}

func (v NullableSetLifecycleStateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetLifecycleStateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


