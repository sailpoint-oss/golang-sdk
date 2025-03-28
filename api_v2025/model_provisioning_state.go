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

// ProvisioningState Provisioning state of an account activity item
type ProvisioningState string

// List of ProvisioningState
const (
	PROVISIONINGSTATE_PENDING ProvisioningState = "PENDING"
	PROVISIONINGSTATE_FINISHED ProvisioningState = "FINISHED"
	PROVISIONINGSTATE_UNVERIFIABLE ProvisioningState = "UNVERIFIABLE"
	PROVISIONINGSTATE_COMMITED ProvisioningState = "COMMITED"
	PROVISIONINGSTATE_FAILED ProvisioningState = "FAILED"
	PROVISIONINGSTATE_RETRY ProvisioningState = "RETRY"
)

// All allowed values of ProvisioningState enum
var AllowedProvisioningStateEnumValues = []ProvisioningState{
	"PENDING",
	"FINISHED",
	"UNVERIFIABLE",
	"COMMITED",
	"FAILED",
	"RETRY",
}

func (v *ProvisioningState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProvisioningState(value)
	for _, existing := range AllowedProvisioningStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProvisioningState", value)
}

// NewProvisioningStateFromValue returns a pointer to a valid ProvisioningState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProvisioningStateFromValue(v string) (*ProvisioningState, error) {
	ev := ProvisioningState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProvisioningState: valid values are %v", v, AllowedProvisioningStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProvisioningState) IsValid() bool {
	for _, existing := range AllowedProvisioningStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProvisioningState value
func (v ProvisioningState) Ptr() *ProvisioningState {
	return &v
}

type NullableProvisioningState struct {
	value *ProvisioningState
	isSet bool
}

func (v NullableProvisioningState) Get() *ProvisioningState {
	return v.value
}

func (v *NullableProvisioningState) Set(val *ProvisioningState) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisioningState) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisioningState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisioningState(val *ProvisioningState) *NullableProvisioningState {
	return &NullableProvisioningState{value: val, isSet: true}
}

func (v NullableProvisioningState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisioningState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

