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

// ApprovalStatus Enum representing the non-employee request approval status
type ApprovalStatus string

// List of ApprovalStatus
const (
	APPROVALSTATUS_APPROVED ApprovalStatus = "APPROVED"
	APPROVALSTATUS_REJECTED ApprovalStatus = "REJECTED"
	APPROVALSTATUS_PENDING ApprovalStatus = "PENDING"
	APPROVALSTATUS_NOT_READY ApprovalStatus = "NOT_READY"
	APPROVALSTATUS_CANCELLED ApprovalStatus = "CANCELLED"
)

// All allowed values of ApprovalStatus enum
var AllowedApprovalStatusEnumValues = []ApprovalStatus{
	"APPROVED",
	"REJECTED",
	"PENDING",
	"NOT_READY",
	"CANCELLED",
}

func (v *ApprovalStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ApprovalStatus(value)
	for _, existing := range AllowedApprovalStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ApprovalStatus", value)
}

// NewApprovalStatusFromValue returns a pointer to a valid ApprovalStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewApprovalStatusFromValue(v string) (*ApprovalStatus, error) {
	ev := ApprovalStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ApprovalStatus: valid values are %v", v, AllowedApprovalStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ApprovalStatus) IsValid() bool {
	for _, existing := range AllowedApprovalStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ApprovalStatus value
func (v ApprovalStatus) Ptr() *ApprovalStatus {
	return &v
}

type NullableApprovalStatus struct {
	value *ApprovalStatus
	isSet bool
}

func (v NullableApprovalStatus) Get() *ApprovalStatus {
	return v.value
}

func (v *NullableApprovalStatus) Set(val *ApprovalStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableApprovalStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableApprovalStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApprovalStatus(val *ApprovalStatus) *NullableApprovalStatus {
	return &NullableApprovalStatus{value: val, isSet: true}
}

func (v NullableApprovalStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApprovalStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

