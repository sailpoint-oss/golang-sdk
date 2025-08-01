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

// checks if the ApprovalApproveRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApprovalApproveRequest{}

// ApprovalApproveRequest Approval Approve Request
type ApprovalApproveRequest struct {
	// Additional attributes as key-value pairs that are not part of the standard schema but can be included for custom data.
	AdditionalAttributes *map[string]string `json:"additionalAttributes,omitempty"`
	// Comment associated with the request.
	Comment *string `json:"comment,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApprovalApproveRequest ApprovalApproveRequest

// NewApprovalApproveRequest instantiates a new ApprovalApproveRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApprovalApproveRequest() *ApprovalApproveRequest {
	this := ApprovalApproveRequest{}
	return &this
}

// NewApprovalApproveRequestWithDefaults instantiates a new ApprovalApproveRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApprovalApproveRequestWithDefaults() *ApprovalApproveRequest {
	this := ApprovalApproveRequest{}
	return &this
}

// GetAdditionalAttributes returns the AdditionalAttributes field value if set, zero value otherwise.
func (o *ApprovalApproveRequest) GetAdditionalAttributes() map[string]string {
	if o == nil || IsNil(o.AdditionalAttributes) {
		var ret map[string]string
		return ret
	}
	return *o.AdditionalAttributes
}

// GetAdditionalAttributesOk returns a tuple with the AdditionalAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalApproveRequest) GetAdditionalAttributesOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.AdditionalAttributes) {
		return nil, false
	}
	return o.AdditionalAttributes, true
}

// HasAdditionalAttributes returns a boolean if a field has been set.
func (o *ApprovalApproveRequest) HasAdditionalAttributes() bool {
	if o != nil && !IsNil(o.AdditionalAttributes) {
		return true
	}

	return false
}

// SetAdditionalAttributes gets a reference to the given map[string]string and assigns it to the AdditionalAttributes field.
func (o *ApprovalApproveRequest) SetAdditionalAttributes(v map[string]string) {
	o.AdditionalAttributes = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *ApprovalApproveRequest) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalApproveRequest) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *ApprovalApproveRequest) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *ApprovalApproveRequest) SetComment(v string) {
	o.Comment = &v
}

func (o ApprovalApproveRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApprovalApproveRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdditionalAttributes) {
		toSerialize["additionalAttributes"] = o.AdditionalAttributes
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApprovalApproveRequest) UnmarshalJSON(data []byte) (err error) {
	varApprovalApproveRequest := _ApprovalApproveRequest{}

	err = json.Unmarshal(data, &varApprovalApproveRequest)

	if err != nil {
		return err
	}

	*o = ApprovalApproveRequest(varApprovalApproveRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "additionalAttributes")
		delete(additionalProperties, "comment")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApprovalApproveRequest struct {
	value *ApprovalApproveRequest
	isSet bool
}

func (v NullableApprovalApproveRequest) Get() *ApprovalApproveRequest {
	return v.value
}

func (v *NullableApprovalApproveRequest) Set(val *ApprovalApproveRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApprovalApproveRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApprovalApproveRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApprovalApproveRequest(val *ApprovalApproveRequest) *NullableApprovalApproveRequest {
	return &NullableApprovalApproveRequest{value: val, isSet: true}
}

func (v NullableApprovalApproveRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApprovalApproveRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


