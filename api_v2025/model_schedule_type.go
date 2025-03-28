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

// ScheduleType Enum representing the currently supported schedule types.  Additional values may be added in the future without notice. 
type ScheduleType string

// List of ScheduleType
const (
	SCHEDULETYPE_DAILY ScheduleType = "DAILY"
	SCHEDULETYPE_WEEKLY ScheduleType = "WEEKLY"
	SCHEDULETYPE_MONTHLY ScheduleType = "MONTHLY"
	SCHEDULETYPE_CALENDAR ScheduleType = "CALENDAR"
	SCHEDULETYPE_ANNUALLY ScheduleType = "ANNUALLY"
)

// All allowed values of ScheduleType enum
var AllowedScheduleTypeEnumValues = []ScheduleType{
	"DAILY",
	"WEEKLY",
	"MONTHLY",
	"CALENDAR",
	"ANNUALLY",
}

func (v *ScheduleType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ScheduleType(value)
	for _, existing := range AllowedScheduleTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ScheduleType", value)
}

// NewScheduleTypeFromValue returns a pointer to a valid ScheduleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewScheduleTypeFromValue(v string) (*ScheduleType, error) {
	ev := ScheduleType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ScheduleType: valid values are %v", v, AllowedScheduleTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ScheduleType) IsValid() bool {
	for _, existing := range AllowedScheduleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScheduleType value
func (v ScheduleType) Ptr() *ScheduleType {
	return &v
}

type NullableScheduleType struct {
	value *ScheduleType
	isSet bool
}

func (v NullableScheduleType) Get() *ScheduleType {
	return v.value
}

func (v *NullableScheduleType) Set(val *ScheduleType) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduleType) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduleType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduleType(val *ScheduleType) *NullableScheduleType {
	return &NullableScheduleType{value: val, isSet: true}
}

func (v NullableScheduleType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduleType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

