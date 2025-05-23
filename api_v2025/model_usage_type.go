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

// UsageType The type of provisioning policy usage.  In IdentityNow, a source can support various provisioning operations. For example, when a joiner is added to a source, this may trigger both CREATE and UPDATE provisioning operations.  Each usage type is considered a provisioning policy.  A source can have any number of these provisioning policies defined.  These are the common usage types:  CREATE - This usage type relates to 'Create Account Profile', the provisioning template for the account to be created. For example, this would be used for a joiner on a source.   UPDATE - This usage type relates to 'Update Account Profile', the provisioning template for the 'Update' connector operations. For example, this would be used for an attribute sync on a source. ENABLE - This usage type relates to 'Enable Account Profile', the provisioning template for the account to be enabled. For example, this could be used for a joiner on a source once the joiner's account is created.  DISABLE - This usage type relates to 'Disable Account Profile', the provisioning template for the account to be disabled. For example, this could be used when a leaver is removed temporarily from a source.  You can use these four usage types for all your provisioning policy needs. 
type UsageType string

// List of UsageType
const (
	USAGETYPE_CREATE UsageType = "CREATE"
	USAGETYPE_UPDATE UsageType = "UPDATE"
	USAGETYPE_ENABLE UsageType = "ENABLE"
	USAGETYPE_DISABLE UsageType = "DISABLE"
	USAGETYPE_DELETE UsageType = "DELETE"
	USAGETYPE_ASSIGN UsageType = "ASSIGN"
	USAGETYPE_UNASSIGN UsageType = "UNASSIGN"
	USAGETYPE_CREATE_GROUP UsageType = "CREATE_GROUP"
	USAGETYPE_UPDATE_GROUP UsageType = "UPDATE_GROUP"
	USAGETYPE_DELETE_GROUP UsageType = "DELETE_GROUP"
	USAGETYPE_REGISTER UsageType = "REGISTER"
	USAGETYPE_CREATE_IDENTITY UsageType = "CREATE_IDENTITY"
	USAGETYPE_UPDATE_IDENTITY UsageType = "UPDATE_IDENTITY"
	USAGETYPE_EDIT_GROUP UsageType = "EDIT_GROUP"
	USAGETYPE_UNLOCK UsageType = "UNLOCK"
	USAGETYPE_CHANGE_PASSWORD UsageType = "CHANGE_PASSWORD"
)

// All allowed values of UsageType enum
var AllowedUsageTypeEnumValues = []UsageType{
	"CREATE",
	"UPDATE",
	"ENABLE",
	"DISABLE",
	"DELETE",
	"ASSIGN",
	"UNASSIGN",
	"CREATE_GROUP",
	"UPDATE_GROUP",
	"DELETE_GROUP",
	"REGISTER",
	"CREATE_IDENTITY",
	"UPDATE_IDENTITY",
	"EDIT_GROUP",
	"UNLOCK",
	"CHANGE_PASSWORD",
}

func (v *UsageType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UsageType(value)
	for _, existing := range AllowedUsageTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UsageType", value)
}

// NewUsageTypeFromValue returns a pointer to a valid UsageType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUsageTypeFromValue(v string) (*UsageType, error) {
	ev := UsageType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UsageType: valid values are %v", v, AllowedUsageTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UsageType) IsValid() bool {
	for _, existing := range AllowedUsageTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UsageType value
func (v UsageType) Ptr() *UsageType {
	return &v
}

type NullableUsageType struct {
	value *UsageType
	isSet bool
}

func (v NullableUsageType) Get() *UsageType {
	return v.value
}

func (v *NullableUsageType) Set(val *UsageType) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageType) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageType(val *UsageType) *NullableUsageType {
	return &NullableUsageType{value: val, isSet: true}
}

func (v NullableUsageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

