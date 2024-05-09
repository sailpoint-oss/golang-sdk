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

// checks if the ApprovalComment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApprovalComment{}

// ApprovalComment Comments Object
type ApprovalComment struct {
	Author *ApprovalIdentity `json:"author,omitempty"`
	// Comment to be left on an approval
	Comment *string `json:"comment,omitempty"`
	// Date the comment was created
	CreatedDate *string `json:"createdDate,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApprovalComment ApprovalComment

// NewApprovalComment instantiates a new ApprovalComment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApprovalComment() *ApprovalComment {
	this := ApprovalComment{}
	return &this
}

// NewApprovalCommentWithDefaults instantiates a new ApprovalComment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApprovalCommentWithDefaults() *ApprovalComment {
	this := ApprovalComment{}
	return &this
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *ApprovalComment) GetAuthor() ApprovalIdentity {
	if o == nil || isNil(o.Author) {
		var ret ApprovalIdentity
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalComment) GetAuthorOk() (*ApprovalIdentity, bool) {
	if o == nil || isNil(o.Author) {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *ApprovalComment) HasAuthor() bool {
	if o != nil && !isNil(o.Author) {
		return true
	}

	return false
}

// SetAuthor gets a reference to the given ApprovalIdentity and assigns it to the Author field.
func (o *ApprovalComment) SetAuthor(v ApprovalIdentity) {
	o.Author = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *ApprovalComment) GetComment() string {
	if o == nil || isNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalComment) GetCommentOk() (*string, bool) {
	if o == nil || isNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *ApprovalComment) HasComment() bool {
	if o != nil && !isNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *ApprovalComment) SetComment(v string) {
	o.Comment = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *ApprovalComment) GetCreatedDate() string {
	if o == nil || isNil(o.CreatedDate) {
		var ret string
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalComment) GetCreatedDateOk() (*string, bool) {
	if o == nil || isNil(o.CreatedDate) {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *ApprovalComment) HasCreatedDate() bool {
	if o != nil && !isNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given string and assigns it to the CreatedDate field.
func (o *ApprovalComment) SetCreatedDate(v string) {
	o.CreatedDate = &v
}

func (o ApprovalComment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApprovalComment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Author) {
		toSerialize["author"] = o.Author
	}
	if !isNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !isNil(o.CreatedDate) {
		toSerialize["createdDate"] = o.CreatedDate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApprovalComment) UnmarshalJSON(bytes []byte) (err error) {
	varApprovalComment := _ApprovalComment{}

	if err = json.Unmarshal(bytes, &varApprovalComment); err == nil {
	*o = ApprovalComment(varApprovalComment)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "author")
		delete(additionalProperties, "comment")
		delete(additionalProperties, "createdDate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApprovalComment struct {
	value *ApprovalComment
	isSet bool
}

func (v NullableApprovalComment) Get() *ApprovalComment {
	return v.value
}

func (v *NullableApprovalComment) Set(val *ApprovalComment) {
	v.value = val
	v.isSet = true
}

func (v NullableApprovalComment) IsSet() bool {
	return v.isSet
}

func (v *NullableApprovalComment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApprovalComment(val *ApprovalComment) *NullableApprovalComment {
	return &NullableApprovalComment{value: val, isSet: true}
}

func (v NullableApprovalComment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApprovalComment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

