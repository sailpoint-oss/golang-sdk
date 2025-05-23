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

// checks if the ProcessIdentitiesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProcessIdentitiesRequest{}

// ProcessIdentitiesRequest struct for ProcessIdentitiesRequest
type ProcessIdentitiesRequest struct {
	// List of up to 250 identity IDs to process.
	IdentityIds []string `json:"identityIds,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProcessIdentitiesRequest ProcessIdentitiesRequest

// NewProcessIdentitiesRequest instantiates a new ProcessIdentitiesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessIdentitiesRequest() *ProcessIdentitiesRequest {
	this := ProcessIdentitiesRequest{}
	return &this
}

// NewProcessIdentitiesRequestWithDefaults instantiates a new ProcessIdentitiesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessIdentitiesRequestWithDefaults() *ProcessIdentitiesRequest {
	this := ProcessIdentitiesRequest{}
	return &this
}

// GetIdentityIds returns the IdentityIds field value if set, zero value otherwise.
func (o *ProcessIdentitiesRequest) GetIdentityIds() []string {
	if o == nil || IsNil(o.IdentityIds) {
		var ret []string
		return ret
	}
	return o.IdentityIds
}

// GetIdentityIdsOk returns a tuple with the IdentityIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessIdentitiesRequest) GetIdentityIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.IdentityIds) {
		return nil, false
	}
	return o.IdentityIds, true
}

// HasIdentityIds returns a boolean if a field has been set.
func (o *ProcessIdentitiesRequest) HasIdentityIds() bool {
	if o != nil && !IsNil(o.IdentityIds) {
		return true
	}

	return false
}

// SetIdentityIds gets a reference to the given []string and assigns it to the IdentityIds field.
func (o *ProcessIdentitiesRequest) SetIdentityIds(v []string) {
	o.IdentityIds = v
}

func (o ProcessIdentitiesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProcessIdentitiesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IdentityIds) {
		toSerialize["identityIds"] = o.IdentityIds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProcessIdentitiesRequest) UnmarshalJSON(data []byte) (err error) {
	varProcessIdentitiesRequest := _ProcessIdentitiesRequest{}

	err = json.Unmarshal(data, &varProcessIdentitiesRequest)

	if err != nil {
		return err
	}

	*o = ProcessIdentitiesRequest(varProcessIdentitiesRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "identityIds")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProcessIdentitiesRequest struct {
	value *ProcessIdentitiesRequest
	isSet bool
}

func (v NullableProcessIdentitiesRequest) Get() *ProcessIdentitiesRequest {
	return v.value
}

func (v *NullableProcessIdentitiesRequest) Set(val *ProcessIdentitiesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessIdentitiesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessIdentitiesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessIdentitiesRequest(val *ProcessIdentitiesRequest) *NullableProcessIdentitiesRequest {
	return &NullableProcessIdentitiesRequest{value: val, isSet: true}
}

func (v NullableProcessIdentitiesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessIdentitiesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


