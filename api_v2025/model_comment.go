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

// checks if the Comment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Comment{}

// Comment struct for Comment
type Comment struct {
	// Id of the identity making the comment
	CommenterId *string `json:"commenterId,omitempty"`
	// Human-readable display name of the identity making the comment
	CommenterName *string `json:"commenterName,omitempty"`
	// Content of the comment
	Body *string `json:"body,omitempty"`
	// Date and time comment was made
	Date *SailPointTime `json:"date,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Comment Comment

// NewComment instantiates a new Comment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComment() *Comment {
	this := Comment{}
	return &this
}

// NewCommentWithDefaults instantiates a new Comment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommentWithDefaults() *Comment {
	this := Comment{}
	return &this
}

// GetCommenterId returns the CommenterId field value if set, zero value otherwise.
func (o *Comment) GetCommenterId() string {
	if o == nil || IsNil(o.CommenterId) {
		var ret string
		return ret
	}
	return *o.CommenterId
}

// GetCommenterIdOk returns a tuple with the CommenterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Comment) GetCommenterIdOk() (*string, bool) {
	if o == nil || IsNil(o.CommenterId) {
		return nil, false
	}
	return o.CommenterId, true
}

// HasCommenterId returns a boolean if a field has been set.
func (o *Comment) HasCommenterId() bool {
	if o != nil && !IsNil(o.CommenterId) {
		return true
	}

	return false
}

// SetCommenterId gets a reference to the given string and assigns it to the CommenterId field.
func (o *Comment) SetCommenterId(v string) {
	o.CommenterId = &v
}

// GetCommenterName returns the CommenterName field value if set, zero value otherwise.
func (o *Comment) GetCommenterName() string {
	if o == nil || IsNil(o.CommenterName) {
		var ret string
		return ret
	}
	return *o.CommenterName
}

// GetCommenterNameOk returns a tuple with the CommenterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Comment) GetCommenterNameOk() (*string, bool) {
	if o == nil || IsNil(o.CommenterName) {
		return nil, false
	}
	return o.CommenterName, true
}

// HasCommenterName returns a boolean if a field has been set.
func (o *Comment) HasCommenterName() bool {
	if o != nil && !IsNil(o.CommenterName) {
		return true
	}

	return false
}

// SetCommenterName gets a reference to the given string and assigns it to the CommenterName field.
func (o *Comment) SetCommenterName(v string) {
	o.CommenterName = &v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *Comment) GetBody() string {
	if o == nil || IsNil(o.Body) {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Comment) GetBodyOk() (*string, bool) {
	if o == nil || IsNil(o.Body) {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *Comment) HasBody() bool {
	if o != nil && !IsNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *Comment) SetBody(v string) {
	o.Body = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *Comment) GetDate() SailPointTime {
	if o == nil || IsNil(o.Date) {
		var ret SailPointTime
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Comment) GetDateOk() (*SailPointTime, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *Comment) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given SailPointTime and assigns it to the Date field.
func (o *Comment) SetDate(v SailPointTime) {
	o.Date = &v
}

func (o Comment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Comment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CommenterId) {
		toSerialize["commenterId"] = o.CommenterId
	}
	if !IsNil(o.CommenterName) {
		toSerialize["commenterName"] = o.CommenterName
	}
	if !IsNil(o.Body) {
		toSerialize["body"] = o.Body
	}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Comment) UnmarshalJSON(data []byte) (err error) {
	varComment := _Comment{}

	err = json.Unmarshal(data, &varComment)

	if err != nil {
		return err
	}

	*o = Comment(varComment)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "commenterId")
		delete(additionalProperties, "commenterName")
		delete(additionalProperties, "body")
		delete(additionalProperties, "date")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableComment struct {
	value *Comment
	isSet bool
}

func (v NullableComment) Get() *Comment {
	return v.value
}

func (v *NullableComment) Set(val *Comment) {
	v.value = val
	v.isSet = true
}

func (v NullableComment) IsSet() bool {
	return v.isSet
}

func (v *NullableComment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComment(val *Comment) *NullableComment {
	return &NullableComment{value: val, isSet: true}
}

func (v NullableComment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


