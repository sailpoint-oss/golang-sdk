/*
IdentityNow Beta API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_beta

import (
	"encoding/json"
	"fmt"
)

// checks if the KbaAnswerRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KbaAnswerRequest{}

// KbaAnswerRequest struct for KbaAnswerRequest
type KbaAnswerRequest struct {
	// Kba answers
	Answers []KbaAnswerRequestItem `json:"answers"`
	AdditionalProperties map[string]interface{}
}

type _KbaAnswerRequest KbaAnswerRequest

// NewKbaAnswerRequest instantiates a new KbaAnswerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKbaAnswerRequest(answers []KbaAnswerRequestItem) *KbaAnswerRequest {
	this := KbaAnswerRequest{}
	this.Answers = answers
	return &this
}

// NewKbaAnswerRequestWithDefaults instantiates a new KbaAnswerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKbaAnswerRequestWithDefaults() *KbaAnswerRequest {
	this := KbaAnswerRequest{}
	return &this
}

// GetAnswers returns the Answers field value
func (o *KbaAnswerRequest) GetAnswers() []KbaAnswerRequestItem {
	if o == nil {
		var ret []KbaAnswerRequestItem
		return ret
	}

	return o.Answers
}

// GetAnswersOk returns a tuple with the Answers field value
// and a boolean to check if the value has been set.
func (o *KbaAnswerRequest) GetAnswersOk() ([]KbaAnswerRequestItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Answers, true
}

// SetAnswers sets field value
func (o *KbaAnswerRequest) SetAnswers(v []KbaAnswerRequestItem) {
	o.Answers = v
}

func (o KbaAnswerRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KbaAnswerRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["answers"] = o.Answers

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *KbaAnswerRequest) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"answers",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varKbaAnswerRequest := _KbaAnswerRequest{}

	if err = json.Unmarshal(bytes, &varKbaAnswerRequest); err == nil {
	*o = KbaAnswerRequest(varKbaAnswerRequest)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "answers")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKbaAnswerRequest struct {
	value *KbaAnswerRequest
	isSet bool
}

func (v NullableKbaAnswerRequest) Get() *KbaAnswerRequest {
	return v.value
}

func (v *NullableKbaAnswerRequest) Set(val *KbaAnswerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableKbaAnswerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableKbaAnswerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKbaAnswerRequest(val *KbaAnswerRequest) *NullableKbaAnswerRequest {
	return &NullableKbaAnswerRequest{value: val, isSet: true}
}

func (v NullableKbaAnswerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKbaAnswerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

