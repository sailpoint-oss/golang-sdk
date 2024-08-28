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

// checks if the KbaAnswerResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KbaAnswerResponseItem{}

// KbaAnswerResponseItem struct for KbaAnswerResponseItem
type KbaAnswerResponseItem struct {
	// Question Id
	Id string `json:"id"`
	// Question description
	Question string `json:"question"`
	// Denotes whether the KBA question has an answer configured for the current user
	HasAnswer bool `json:"hasAnswer"`
	AdditionalProperties map[string]interface{}
}

type _KbaAnswerResponseItem KbaAnswerResponseItem

// NewKbaAnswerResponseItem instantiates a new KbaAnswerResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKbaAnswerResponseItem(id string, question string, hasAnswer bool) *KbaAnswerResponseItem {
	this := KbaAnswerResponseItem{}
	this.Id = id
	this.Question = question
	this.HasAnswer = hasAnswer
	return &this
}

// NewKbaAnswerResponseItemWithDefaults instantiates a new KbaAnswerResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKbaAnswerResponseItemWithDefaults() *KbaAnswerResponseItem {
	this := KbaAnswerResponseItem{}
	return &this
}

// GetId returns the Id field value
func (o *KbaAnswerResponseItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *KbaAnswerResponseItem) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *KbaAnswerResponseItem) SetId(v string) {
	o.Id = v
}

// GetQuestion returns the Question field value
func (o *KbaAnswerResponseItem) GetQuestion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Question
}

// GetQuestionOk returns a tuple with the Question field value
// and a boolean to check if the value has been set.
func (o *KbaAnswerResponseItem) GetQuestionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Question, true
}

// SetQuestion sets field value
func (o *KbaAnswerResponseItem) SetQuestion(v string) {
	o.Question = v
}

// GetHasAnswer returns the HasAnswer field value
func (o *KbaAnswerResponseItem) GetHasAnswer() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasAnswer
}

// GetHasAnswerOk returns a tuple with the HasAnswer field value
// and a boolean to check if the value has been set.
func (o *KbaAnswerResponseItem) GetHasAnswerOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasAnswer, true
}

// SetHasAnswer sets field value
func (o *KbaAnswerResponseItem) SetHasAnswer(v bool) {
	o.HasAnswer = v
}

func (o KbaAnswerResponseItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KbaAnswerResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["question"] = o.Question
	toSerialize["hasAnswer"] = o.HasAnswer

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *KbaAnswerResponseItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"question",
		"hasAnswer",
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

	varKbaAnswerResponseItem := _KbaAnswerResponseItem{}

	err = json.Unmarshal(data, &varKbaAnswerResponseItem)

	if err != nil {
		return err
	}

	*o = KbaAnswerResponseItem(varKbaAnswerResponseItem)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "question")
		delete(additionalProperties, "hasAnswer")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKbaAnswerResponseItem struct {
	value *KbaAnswerResponseItem
	isSet bool
}

func (v NullableKbaAnswerResponseItem) Get() *KbaAnswerResponseItem {
	return v.value
}

func (v *NullableKbaAnswerResponseItem) Set(val *KbaAnswerResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableKbaAnswerResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableKbaAnswerResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKbaAnswerResponseItem(val *KbaAnswerResponseItem) *NullableKbaAnswerResponseItem {
	return &NullableKbaAnswerResponseItem{value: val, isSet: true}
}

func (v NullableKbaAnswerResponseItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKbaAnswerResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

