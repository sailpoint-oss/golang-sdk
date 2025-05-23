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

// ClientType Type of an API Client indicating public or confidentials use
type ClientType string

// List of ClientType
const (
	CLIENTTYPE_CONFIDENTIAL ClientType = "CONFIDENTIAL"
	CLIENTTYPE_PUBLIC ClientType = "PUBLIC"
)

// All allowed values of ClientType enum
var AllowedClientTypeEnumValues = []ClientType{
	"CONFIDENTIAL",
	"PUBLIC",
}

func (v *ClientType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ClientType(value)
	for _, existing := range AllowedClientTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ClientType", value)
}

// NewClientTypeFromValue returns a pointer to a valid ClientType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClientTypeFromValue(v string) (*ClientType, error) {
	ev := ClientType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClientType: valid values are %v", v, AllowedClientTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClientType) IsValid() bool {
	for _, existing := range AllowedClientTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ClientType value
func (v ClientType) Ptr() *ClientType {
	return &v
}

type NullableClientType struct {
	value *ClientType
	isSet bool
}

func (v NullableClientType) Get() *ClientType {
	return v.value
}

func (v *NullableClientType) Set(val *ClientType) {
	v.value = val
	v.isSet = true
}

func (v NullableClientType) IsSet() bool {
	return v.isSet
}

func (v *NullableClientType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientType(val *ClientType) *NullableClientType {
	return &NullableClientType{value: val, isSet: true}
}

func (v NullableClientType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

