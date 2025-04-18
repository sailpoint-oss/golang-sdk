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

// ScopeType An enumeration of the types of scope choices
type ScopeType string

// List of ScopeType
const (
	SCOPETYPE_ENTITLEMENT ScopeType = "ENTITLEMENT"
	SCOPETYPE_CERTIFICATION ScopeType = "CERTIFICATION"
	SCOPETYPE_IDENTITY ScopeType = "IDENTITY"
	SCOPETYPE_ENTITLEMENTREQUEST ScopeType = "ENTITLEMENTREQUEST"
)

// All allowed values of ScopeType enum
var AllowedScopeTypeEnumValues = []ScopeType{
	"ENTITLEMENT",
	"CERTIFICATION",
	"IDENTITY",
	"ENTITLEMENTREQUEST",
}

func (v *ScopeType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ScopeType(value)
	for _, existing := range AllowedScopeTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ScopeType", value)
}

// NewScopeTypeFromValue returns a pointer to a valid ScopeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewScopeTypeFromValue(v string) (*ScopeType, error) {
	ev := ScopeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ScopeType: valid values are %v", v, AllowedScopeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ScopeType) IsValid() bool {
	for _, existing := range AllowedScopeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScopeType value
func (v ScopeType) Ptr() *ScopeType {
	return &v
}

type NullableScopeType struct {
	value *ScopeType
	isSet bool
}

func (v NullableScopeType) Get() *ScopeType {
	return v.value
}

func (v *NullableScopeType) Set(val *ScopeType) {
	v.value = val
	v.isSet = true
}

func (v NullableScopeType) IsSet() bool {
	return v.isSet
}

func (v *NullableScopeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScopeType(val *ScopeType) *NullableScopeType {
	return &NullableScopeType{value: val, isSet: true}
}

func (v NullableScopeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScopeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

