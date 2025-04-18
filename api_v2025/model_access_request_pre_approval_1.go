/*
Identity Security Cloud V2025 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: v2025
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2025

import (
	"encoding/json"
	"fmt"
)

// checks if the AccessRequestPreApproval1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessRequestPreApproval1{}

// AccessRequestPreApproval1 struct for AccessRequestPreApproval1
type AccessRequestPreApproval1 struct {
	// Whether or not to approve the access request.
	Approved bool `json:"approved"`
	// A comment about the decision to approve or deny the request.
	Comment string `json:"comment"`
	// The name of the entity that approved or denied the request.
	Approver string `json:"approver"`
	AdditionalProperties map[string]interface{}
}

type _AccessRequestPreApproval1 AccessRequestPreApproval1

// NewAccessRequestPreApproval1 instantiates a new AccessRequestPreApproval1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessRequestPreApproval1(approved bool, comment string, approver string) *AccessRequestPreApproval1 {
	this := AccessRequestPreApproval1{}
	this.Approved = approved
	this.Comment = comment
	this.Approver = approver
	return &this
}

// NewAccessRequestPreApproval1WithDefaults instantiates a new AccessRequestPreApproval1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessRequestPreApproval1WithDefaults() *AccessRequestPreApproval1 {
	this := AccessRequestPreApproval1{}
	return &this
}

// GetApproved returns the Approved field value
func (o *AccessRequestPreApproval1) GetApproved() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Approved
}

// GetApprovedOk returns a tuple with the Approved field value
// and a boolean to check if the value has been set.
func (o *AccessRequestPreApproval1) GetApprovedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Approved, true
}

// SetApproved sets field value
func (o *AccessRequestPreApproval1) SetApproved(v bool) {
	o.Approved = v
}

// GetComment returns the Comment field value
func (o *AccessRequestPreApproval1) GetComment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Comment
}

// GetCommentOk returns a tuple with the Comment field value
// and a boolean to check if the value has been set.
func (o *AccessRequestPreApproval1) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Comment, true
}

// SetComment sets field value
func (o *AccessRequestPreApproval1) SetComment(v string) {
	o.Comment = v
}

// GetApprover returns the Approver field value
func (o *AccessRequestPreApproval1) GetApprover() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Approver
}

// GetApproverOk returns a tuple with the Approver field value
// and a boolean to check if the value has been set.
func (o *AccessRequestPreApproval1) GetApproverOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Approver, true
}

// SetApprover sets field value
func (o *AccessRequestPreApproval1) SetApprover(v string) {
	o.Approver = v
}

func (o AccessRequestPreApproval1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessRequestPreApproval1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["approved"] = o.Approved
	toSerialize["comment"] = o.Comment
	toSerialize["approver"] = o.Approver

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccessRequestPreApproval1) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"approved",
		"comment",
		"approver",
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

	varAccessRequestPreApproval1 := _AccessRequestPreApproval1{}

	err = json.Unmarshal(data, &varAccessRequestPreApproval1)

	if err != nil {
		return err
	}

	*o = AccessRequestPreApproval1(varAccessRequestPreApproval1)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "approved")
		delete(additionalProperties, "comment")
		delete(additionalProperties, "approver")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccessRequestPreApproval1 struct {
	value *AccessRequestPreApproval1
	isSet bool
}

func (v NullableAccessRequestPreApproval1) Get() *AccessRequestPreApproval1 {
	return v.value
}

func (v *NullableAccessRequestPreApproval1) Set(val *AccessRequestPreApproval1) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessRequestPreApproval1) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessRequestPreApproval1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessRequestPreApproval1(val *AccessRequestPreApproval1) *NullableAccessRequestPreApproval1 {
	return &NullableAccessRequestPreApproval1{value: val, isSet: true}
}

func (v NullableAccessRequestPreApproval1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessRequestPreApproval1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


