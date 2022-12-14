/*
SailPoint SaaS API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sailpointv2sdk

import (
	"encoding/json"
)

// CommentEto struct for CommentEto
type CommentEto struct {
	Comment *string `json:"comment,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CommentEto CommentEto

// NewCommentEto instantiates a new CommentEto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommentEto() *CommentEto {
	this := CommentEto{}
	return &this
}

// NewCommentEtoWithDefaults instantiates a new CommentEto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommentEtoWithDefaults() *CommentEto {
	this := CommentEto{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *CommentEto) GetComment() string {
	if o == nil || isNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommentEto) GetCommentOk() (*string, bool) {
	if o == nil || isNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *CommentEto) HasComment() bool {
	if o != nil && !isNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *CommentEto) SetComment(v string) {
	o.Comment = &v
}

func (o CommentEto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CommentEto) UnmarshalJSON(bytes []byte) (err error) {
	varCommentEto := _CommentEto{}

	if err = json.Unmarshal(bytes, &varCommentEto); err == nil {
		*o = CommentEto(varCommentEto)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "comment")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCommentEto struct {
	value *CommentEto
	isSet bool
}

func (v NullableCommentEto) Get() *CommentEto {
	return v.value
}

func (v *NullableCommentEto) Set(val *CommentEto) {
	v.value = val
	v.isSet = true
}

func (v NullableCommentEto) IsSet() bool {
	return v.isSet
}

func (v *NullableCommentEto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommentEto(val *CommentEto) *NullableCommentEto {
	return &NullableCommentEto{value: val, isSet: true}
}

func (v NullableCommentEto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommentEto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


