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

// checks if the PostExternalExecuteWorkflowRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostExternalExecuteWorkflowRequest{}

// PostExternalExecuteWorkflowRequest struct for PostExternalExecuteWorkflowRequest
type PostExternalExecuteWorkflowRequest struct {
	// The input for the workflow
	Input map[string]interface{} `json:"input,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PostExternalExecuteWorkflowRequest PostExternalExecuteWorkflowRequest

// NewPostExternalExecuteWorkflowRequest instantiates a new PostExternalExecuteWorkflowRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostExternalExecuteWorkflowRequest() *PostExternalExecuteWorkflowRequest {
	this := PostExternalExecuteWorkflowRequest{}
	return &this
}

// NewPostExternalExecuteWorkflowRequestWithDefaults instantiates a new PostExternalExecuteWorkflowRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostExternalExecuteWorkflowRequestWithDefaults() *PostExternalExecuteWorkflowRequest {
	this := PostExternalExecuteWorkflowRequest{}
	return &this
}

// GetInput returns the Input field value if set, zero value otherwise.
func (o *PostExternalExecuteWorkflowRequest) GetInput() map[string]interface{} {
	if o == nil || IsNil(o.Input) {
		var ret map[string]interface{}
		return ret
	}
	return o.Input
}

// GetInputOk returns a tuple with the Input field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostExternalExecuteWorkflowRequest) GetInputOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Input) {
		return map[string]interface{}{}, false
	}
	return o.Input, true
}

// HasInput returns a boolean if a field has been set.
func (o *PostExternalExecuteWorkflowRequest) HasInput() bool {
	if o != nil && !IsNil(o.Input) {
		return true
	}

	return false
}

// SetInput gets a reference to the given map[string]interface{} and assigns it to the Input field.
func (o *PostExternalExecuteWorkflowRequest) SetInput(v map[string]interface{}) {
	o.Input = v
}

func (o PostExternalExecuteWorkflowRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostExternalExecuteWorkflowRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Input) {
		toSerialize["input"] = o.Input
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PostExternalExecuteWorkflowRequest) UnmarshalJSON(data []byte) (err error) {
	varPostExternalExecuteWorkflowRequest := _PostExternalExecuteWorkflowRequest{}

	err = json.Unmarshal(data, &varPostExternalExecuteWorkflowRequest)

	if err != nil {
		return err
	}

	*o = PostExternalExecuteWorkflowRequest(varPostExternalExecuteWorkflowRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "input")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePostExternalExecuteWorkflowRequest struct {
	value *PostExternalExecuteWorkflowRequest
	isSet bool
}

func (v NullablePostExternalExecuteWorkflowRequest) Get() *PostExternalExecuteWorkflowRequest {
	return v.value
}

func (v *NullablePostExternalExecuteWorkflowRequest) Set(val *PostExternalExecuteWorkflowRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostExternalExecuteWorkflowRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostExternalExecuteWorkflowRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostExternalExecuteWorkflowRequest(val *PostExternalExecuteWorkflowRequest) *NullablePostExternalExecuteWorkflowRequest {
	return &NullablePostExternalExecuteWorkflowRequest{value: val, isSet: true}
}

func (v NullablePostExternalExecuteWorkflowRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostExternalExecuteWorkflowRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


