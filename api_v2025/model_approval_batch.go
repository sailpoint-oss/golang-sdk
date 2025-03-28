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

// checks if the ApprovalBatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApprovalBatch{}

// ApprovalBatch Batch properties if an approval is sent via batching.
type ApprovalBatch struct {
	// ID of the batch
	BatchId *string `json:"batchId,omitempty"`
	// How many approvals are going to be in this batch. Defaults to 1 if not provided.
	BatchSize *int64 `json:"batchSize,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApprovalBatch ApprovalBatch

// NewApprovalBatch instantiates a new ApprovalBatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApprovalBatch() *ApprovalBatch {
	this := ApprovalBatch{}
	return &this
}

// NewApprovalBatchWithDefaults instantiates a new ApprovalBatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApprovalBatchWithDefaults() *ApprovalBatch {
	this := ApprovalBatch{}
	return &this
}

// GetBatchId returns the BatchId field value if set, zero value otherwise.
func (o *ApprovalBatch) GetBatchId() string {
	if o == nil || IsNil(o.BatchId) {
		var ret string
		return ret
	}
	return *o.BatchId
}

// GetBatchIdOk returns a tuple with the BatchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalBatch) GetBatchIdOk() (*string, bool) {
	if o == nil || IsNil(o.BatchId) {
		return nil, false
	}
	return o.BatchId, true
}

// HasBatchId returns a boolean if a field has been set.
func (o *ApprovalBatch) HasBatchId() bool {
	if o != nil && !IsNil(o.BatchId) {
		return true
	}

	return false
}

// SetBatchId gets a reference to the given string and assigns it to the BatchId field.
func (o *ApprovalBatch) SetBatchId(v string) {
	o.BatchId = &v
}

// GetBatchSize returns the BatchSize field value if set, zero value otherwise.
func (o *ApprovalBatch) GetBatchSize() int64 {
	if o == nil || IsNil(o.BatchSize) {
		var ret int64
		return ret
	}
	return *o.BatchSize
}

// GetBatchSizeOk returns a tuple with the BatchSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalBatch) GetBatchSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.BatchSize) {
		return nil, false
	}
	return o.BatchSize, true
}

// HasBatchSize returns a boolean if a field has been set.
func (o *ApprovalBatch) HasBatchSize() bool {
	if o != nil && !IsNil(o.BatchSize) {
		return true
	}

	return false
}

// SetBatchSize gets a reference to the given int64 and assigns it to the BatchSize field.
func (o *ApprovalBatch) SetBatchSize(v int64) {
	o.BatchSize = &v
}

func (o ApprovalBatch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApprovalBatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BatchId) {
		toSerialize["batchId"] = o.BatchId
	}
	if !IsNil(o.BatchSize) {
		toSerialize["batchSize"] = o.BatchSize
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApprovalBatch) UnmarshalJSON(data []byte) (err error) {
	varApprovalBatch := _ApprovalBatch{}

	err = json.Unmarshal(data, &varApprovalBatch)

	if err != nil {
		return err
	}

	*o = ApprovalBatch(varApprovalBatch)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "batchId")
		delete(additionalProperties, "batchSize")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApprovalBatch struct {
	value *ApprovalBatch
	isSet bool
}

func (v NullableApprovalBatch) Get() *ApprovalBatch {
	return v.value
}

func (v *NullableApprovalBatch) Set(val *ApprovalBatch) {
	v.value = val
	v.isSet = true
}

func (v NullableApprovalBatch) IsSet() bool {
	return v.isSet
}

func (v *NullableApprovalBatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApprovalBatch(val *ApprovalBatch) *NullableApprovalBatch {
	return &NullableApprovalBatch{value: val, isSet: true}
}

func (v NullableApprovalBatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApprovalBatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


