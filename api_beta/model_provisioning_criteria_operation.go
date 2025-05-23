/*
Identity Security Cloud Beta API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_beta

import (
	"encoding/json"
	"fmt"
)

// ProvisioningCriteriaOperation Supported operations on `ProvisioningCriteria`.
type ProvisioningCriteriaOperation string

// List of ProvisioningCriteriaOperation
const (
	PROVISIONINGCRITERIAOPERATION_EQUALS ProvisioningCriteriaOperation = "EQUALS"
	PROVISIONINGCRITERIAOPERATION_NOT_EQUALS ProvisioningCriteriaOperation = "NOT_EQUALS"
	PROVISIONINGCRITERIAOPERATION_CONTAINS ProvisioningCriteriaOperation = "CONTAINS"
	PROVISIONINGCRITERIAOPERATION_HAS ProvisioningCriteriaOperation = "HAS"
	PROVISIONINGCRITERIAOPERATION_AND ProvisioningCriteriaOperation = "AND"
	PROVISIONINGCRITERIAOPERATION_OR ProvisioningCriteriaOperation = "OR"
)

// All allowed values of ProvisioningCriteriaOperation enum
var AllowedProvisioningCriteriaOperationEnumValues = []ProvisioningCriteriaOperation{
	"EQUALS",
	"NOT_EQUALS",
	"CONTAINS",
	"HAS",
	"AND",
	"OR",
}

func (v *ProvisioningCriteriaOperation) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProvisioningCriteriaOperation(value)
	for _, existing := range AllowedProvisioningCriteriaOperationEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProvisioningCriteriaOperation", value)
}

// NewProvisioningCriteriaOperationFromValue returns a pointer to a valid ProvisioningCriteriaOperation
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProvisioningCriteriaOperationFromValue(v string) (*ProvisioningCriteriaOperation, error) {
	ev := ProvisioningCriteriaOperation(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProvisioningCriteriaOperation: valid values are %v", v, AllowedProvisioningCriteriaOperationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProvisioningCriteriaOperation) IsValid() bool {
	for _, existing := range AllowedProvisioningCriteriaOperationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProvisioningCriteriaOperation value
func (v ProvisioningCriteriaOperation) Ptr() *ProvisioningCriteriaOperation {
	return &v
}

type NullableProvisioningCriteriaOperation struct {
	value *ProvisioningCriteriaOperation
	isSet bool
}

func (v NullableProvisioningCriteriaOperation) Get() *ProvisioningCriteriaOperation {
	return v.value
}

func (v *NullableProvisioningCriteriaOperation) Set(val *ProvisioningCriteriaOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisioningCriteriaOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisioningCriteriaOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisioningCriteriaOperation(val *ProvisioningCriteriaOperation) *NullableProvisioningCriteriaOperation {
	return &NullableProvisioningCriteriaOperation{value: val, isSet: true}
}

func (v NullableProvisioningCriteriaOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisioningCriteriaOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

