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

// AccountActivityItemOperation Represents an operation in an account activity item
type AccountActivityItemOperation string

// List of AccountActivityItemOperation
const (
	ACCOUNTACTIVITYITEMOPERATION_ADD AccountActivityItemOperation = "ADD"
	ACCOUNTACTIVITYITEMOPERATION_CREATE AccountActivityItemOperation = "CREATE"
	ACCOUNTACTIVITYITEMOPERATION_MODIFY AccountActivityItemOperation = "MODIFY"
	ACCOUNTACTIVITYITEMOPERATION_DELETE AccountActivityItemOperation = "DELETE"
	ACCOUNTACTIVITYITEMOPERATION_DISABLE AccountActivityItemOperation = "DISABLE"
	ACCOUNTACTIVITYITEMOPERATION_ENABLE AccountActivityItemOperation = "ENABLE"
	ACCOUNTACTIVITYITEMOPERATION_UNLOCK AccountActivityItemOperation = "UNLOCK"
	ACCOUNTACTIVITYITEMOPERATION_LOCK AccountActivityItemOperation = "LOCK"
	ACCOUNTACTIVITYITEMOPERATION_REMOVE AccountActivityItemOperation = "REMOVE"
	ACCOUNTACTIVITYITEMOPERATION_SET AccountActivityItemOperation = "SET"
)

// All allowed values of AccountActivityItemOperation enum
var AllowedAccountActivityItemOperationEnumValues = []AccountActivityItemOperation{
	"ADD",
	"CREATE",
	"MODIFY",
	"DELETE",
	"DISABLE",
	"ENABLE",
	"UNLOCK",
	"LOCK",
	"REMOVE",
	"SET",
}

func (v *AccountActivityItemOperation) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AccountActivityItemOperation(value)
	for _, existing := range AllowedAccountActivityItemOperationEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AccountActivityItemOperation", value)
}

// NewAccountActivityItemOperationFromValue returns a pointer to a valid AccountActivityItemOperation
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAccountActivityItemOperationFromValue(v string) (*AccountActivityItemOperation, error) {
	ev := AccountActivityItemOperation(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AccountActivityItemOperation: valid values are %v", v, AllowedAccountActivityItemOperationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AccountActivityItemOperation) IsValid() bool {
	for _, existing := range AllowedAccountActivityItemOperationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AccountActivityItemOperation value
func (v AccountActivityItemOperation) Ptr() *AccountActivityItemOperation {
	return &v
}

type NullableAccountActivityItemOperation struct {
	value *AccountActivityItemOperation
	isSet bool
}

func (v NullableAccountActivityItemOperation) Get() *AccountActivityItemOperation {
	return v.value
}

func (v *NullableAccountActivityItemOperation) Set(val *AccountActivityItemOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountActivityItemOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountActivityItemOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountActivityItemOperation(val *AccountActivityItemOperation) *NullableAccountActivityItemOperation {
	return &NullableAccountActivityItemOperation{value: val, isSet: true}
}

func (v NullableAccountActivityItemOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountActivityItemOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

