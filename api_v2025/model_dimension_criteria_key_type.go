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

// DimensionCriteriaKeyType Indicates whether the associated criteria represents an expression on identity attributes.
type DimensionCriteriaKeyType string

// List of DimensionCriteriaKeyType
const (
	DIMENSIONCRITERIAKEYTYPE_IDENTITY DimensionCriteriaKeyType = "IDENTITY"
)

// All allowed values of DimensionCriteriaKeyType enum
var AllowedDimensionCriteriaKeyTypeEnumValues = []DimensionCriteriaKeyType{
	"IDENTITY",
}

func (v *DimensionCriteriaKeyType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DimensionCriteriaKeyType(value)
	for _, existing := range AllowedDimensionCriteriaKeyTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DimensionCriteriaKeyType", value)
}

// NewDimensionCriteriaKeyTypeFromValue returns a pointer to a valid DimensionCriteriaKeyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDimensionCriteriaKeyTypeFromValue(v string) (*DimensionCriteriaKeyType, error) {
	ev := DimensionCriteriaKeyType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DimensionCriteriaKeyType: valid values are %v", v, AllowedDimensionCriteriaKeyTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DimensionCriteriaKeyType) IsValid() bool {
	for _, existing := range AllowedDimensionCriteriaKeyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DimensionCriteriaKeyType value
func (v DimensionCriteriaKeyType) Ptr() *DimensionCriteriaKeyType {
	return &v
}

type NullableDimensionCriteriaKeyType struct {
	value *DimensionCriteriaKeyType
	isSet bool
}

func (v NullableDimensionCriteriaKeyType) Get() *DimensionCriteriaKeyType {
	return v.value
}

func (v *NullableDimensionCriteriaKeyType) Set(val *DimensionCriteriaKeyType) {
	v.value = val
	v.isSet = true
}

func (v NullableDimensionCriteriaKeyType) IsSet() bool {
	return v.isSet
}

func (v *NullableDimensionCriteriaKeyType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDimensionCriteriaKeyType(val *DimensionCriteriaKeyType) *NullableDimensionCriteriaKeyType {
	return &NullableDimensionCriteriaKeyType{value: val, isSet: true}
}

func (v NullableDimensionCriteriaKeyType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDimensionCriteriaKeyType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

