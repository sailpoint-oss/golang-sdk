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

// checks if the BulkApproveAccessRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkApproveAccessRequest{}

// BulkApproveAccessRequest Request body payload for bulk approve access request endpoint.
type BulkApproveAccessRequest struct {
	// List of approval ids to approve the pending requests
	ApprovalIds []string `json:"approvalIds"`
	// Reason for approving the pending access request.
	Comment string `json:"comment"`
	AdditionalProperties map[string]interface{}
}

type _BulkApproveAccessRequest BulkApproveAccessRequest

// NewBulkApproveAccessRequest instantiates a new BulkApproveAccessRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkApproveAccessRequest(approvalIds []string, comment string) *BulkApproveAccessRequest {
	this := BulkApproveAccessRequest{}
	this.ApprovalIds = approvalIds
	this.Comment = comment
	return &this
}

// NewBulkApproveAccessRequestWithDefaults instantiates a new BulkApproveAccessRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkApproveAccessRequestWithDefaults() *BulkApproveAccessRequest {
	this := BulkApproveAccessRequest{}
	return &this
}

// GetApprovalIds returns the ApprovalIds field value
func (o *BulkApproveAccessRequest) GetApprovalIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ApprovalIds
}

// GetApprovalIdsOk returns a tuple with the ApprovalIds field value
// and a boolean to check if the value has been set.
func (o *BulkApproveAccessRequest) GetApprovalIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApprovalIds, true
}

// SetApprovalIds sets field value
func (o *BulkApproveAccessRequest) SetApprovalIds(v []string) {
	o.ApprovalIds = v
}

// GetComment returns the Comment field value
func (o *BulkApproveAccessRequest) GetComment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Comment
}

// GetCommentOk returns a tuple with the Comment field value
// and a boolean to check if the value has been set.
func (o *BulkApproveAccessRequest) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Comment, true
}

// SetComment sets field value
func (o *BulkApproveAccessRequest) SetComment(v string) {
	o.Comment = v
}

func (o BulkApproveAccessRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkApproveAccessRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["approvalIds"] = o.ApprovalIds
	toSerialize["comment"] = o.Comment

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BulkApproveAccessRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"approvalIds",
		"comment",
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

	varBulkApproveAccessRequest := _BulkApproveAccessRequest{}

	err = json.Unmarshal(data, &varBulkApproveAccessRequest)

	if err != nil {
		return err
	}

	*o = BulkApproveAccessRequest(varBulkApproveAccessRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "approvalIds")
		delete(additionalProperties, "comment")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBulkApproveAccessRequest struct {
	value *BulkApproveAccessRequest
	isSet bool
}

func (v NullableBulkApproveAccessRequest) Get() *BulkApproveAccessRequest {
	return v.value
}

func (v *NullableBulkApproveAccessRequest) Set(val *BulkApproveAccessRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkApproveAccessRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkApproveAccessRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkApproveAccessRequest(val *BulkApproveAccessRequest) *NullableBulkApproveAccessRequest {
	return &NullableBulkApproveAccessRequest{value: val, isSet: true}
}

func (v NullableBulkApproveAccessRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkApproveAccessRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

