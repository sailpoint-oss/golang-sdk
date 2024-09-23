/*
Identity Security Cloud V2024 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: v2024
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2024

import (
	"encoding/json"
	"time"
)

// checks if the ApprovalComment2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApprovalComment2{}

// ApprovalComment2 struct for ApprovalComment2
type ApprovalComment2 struct {
	// The comment text
	Comment *string `json:"comment,omitempty"`
	// The name of the commenter
	Commenter *string `json:"commenter,omitempty"`
	// A date-time in ISO-8601 format
	Date NullableTime `json:"date,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApprovalComment2 ApprovalComment2

// NewApprovalComment2 instantiates a new ApprovalComment2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApprovalComment2() *ApprovalComment2 {
	this := ApprovalComment2{}
	return &this
}

// NewApprovalComment2WithDefaults instantiates a new ApprovalComment2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApprovalComment2WithDefaults() *ApprovalComment2 {
	this := ApprovalComment2{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *ApprovalComment2) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalComment2) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *ApprovalComment2) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *ApprovalComment2) SetComment(v string) {
	o.Comment = &v
}

// GetCommenter returns the Commenter field value if set, zero value otherwise.
func (o *ApprovalComment2) GetCommenter() string {
	if o == nil || IsNil(o.Commenter) {
		var ret string
		return ret
	}
	return *o.Commenter
}

// GetCommenterOk returns a tuple with the Commenter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalComment2) GetCommenterOk() (*string, bool) {
	if o == nil || IsNil(o.Commenter) {
		return nil, false
	}
	return o.Commenter, true
}

// HasCommenter returns a boolean if a field has been set.
func (o *ApprovalComment2) HasCommenter() bool {
	if o != nil && !IsNil(o.Commenter) {
		return true
	}

	return false
}

// SetCommenter gets a reference to the given string and assigns it to the Commenter field.
func (o *ApprovalComment2) SetCommenter(v string) {
	o.Commenter = &v
}

// GetDate returns the Date field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApprovalComment2) GetDate() time.Time {
	if o == nil || IsNil(o.Date.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Date.Get()
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApprovalComment2) GetDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Date.Get(), o.Date.IsSet()
}

// HasDate returns a boolean if a field has been set.
func (o *ApprovalComment2) HasDate() bool {
	if o != nil && o.Date.IsSet() {
		return true
	}

	return false
}

// SetDate gets a reference to the given NullableTime and assigns it to the Date field.
func (o *ApprovalComment2) SetDate(v time.Time) {
	o.Date.Set(&v)
}
// SetDateNil sets the value for Date to be an explicit nil
func (o *ApprovalComment2) SetDateNil() {
	o.Date.Set(nil)
}

// UnsetDate ensures that no value is present for Date, not even an explicit nil
func (o *ApprovalComment2) UnsetDate() {
	o.Date.Unset()
}

func (o ApprovalComment2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApprovalComment2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.Commenter) {
		toSerialize["commenter"] = o.Commenter
	}
	if o.Date.IsSet() {
		toSerialize["date"] = o.Date.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApprovalComment2) UnmarshalJSON(data []byte) (err error) {
	varApprovalComment2 := _ApprovalComment2{}

	err = json.Unmarshal(data, &varApprovalComment2)

	if err != nil {
		return err
	}

	*o = ApprovalComment2(varApprovalComment2)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "comment")
		delete(additionalProperties, "commenter")
		delete(additionalProperties, "date")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApprovalComment2 struct {
	value *ApprovalComment2
	isSet bool
}

func (v NullableApprovalComment2) Get() *ApprovalComment2 {
	return v.value
}

func (v *NullableApprovalComment2) Set(val *ApprovalComment2) {
	v.value = val
	v.isSet = true
}

func (v NullableApprovalComment2) IsSet() bool {
	return v.isSet
}

func (v *NullableApprovalComment2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApprovalComment2(val *ApprovalComment2) *NullableApprovalComment2 {
	return &NullableApprovalComment2{value: val, isSet: true}
}

func (v NullableApprovalComment2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApprovalComment2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

