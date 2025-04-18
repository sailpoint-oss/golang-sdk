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

// checks if the WorkgroupConnectionDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkgroupConnectionDto{}

// WorkgroupConnectionDto struct for WorkgroupConnectionDto
type WorkgroupConnectionDto struct {
	Object *WorkgroupConnectionDtoObject `json:"object,omitempty"`
	// Connection Type.
	ConnectionType *string `json:"connectionType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkgroupConnectionDto WorkgroupConnectionDto

// NewWorkgroupConnectionDto instantiates a new WorkgroupConnectionDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkgroupConnectionDto() *WorkgroupConnectionDto {
	this := WorkgroupConnectionDto{}
	return &this
}

// NewWorkgroupConnectionDtoWithDefaults instantiates a new WorkgroupConnectionDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkgroupConnectionDtoWithDefaults() *WorkgroupConnectionDto {
	this := WorkgroupConnectionDto{}
	return &this
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *WorkgroupConnectionDto) GetObject() WorkgroupConnectionDtoObject {
	if o == nil || IsNil(o.Object) {
		var ret WorkgroupConnectionDtoObject
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkgroupConnectionDto) GetObjectOk() (*WorkgroupConnectionDtoObject, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *WorkgroupConnectionDto) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given WorkgroupConnectionDtoObject and assigns it to the Object field.
func (o *WorkgroupConnectionDto) SetObject(v WorkgroupConnectionDtoObject) {
	o.Object = &v
}

// GetConnectionType returns the ConnectionType field value if set, zero value otherwise.
func (o *WorkgroupConnectionDto) GetConnectionType() string {
	if o == nil || IsNil(o.ConnectionType) {
		var ret string
		return ret
	}
	return *o.ConnectionType
}

// GetConnectionTypeOk returns a tuple with the ConnectionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkgroupConnectionDto) GetConnectionTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectionType) {
		return nil, false
	}
	return o.ConnectionType, true
}

// HasConnectionType returns a boolean if a field has been set.
func (o *WorkgroupConnectionDto) HasConnectionType() bool {
	if o != nil && !IsNil(o.ConnectionType) {
		return true
	}

	return false
}

// SetConnectionType gets a reference to the given string and assigns it to the ConnectionType field.
func (o *WorkgroupConnectionDto) SetConnectionType(v string) {
	o.ConnectionType = &v
}

func (o WorkgroupConnectionDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkgroupConnectionDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	if !IsNil(o.ConnectionType) {
		toSerialize["connectionType"] = o.ConnectionType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WorkgroupConnectionDto) UnmarshalJSON(data []byte) (err error) {
	varWorkgroupConnectionDto := _WorkgroupConnectionDto{}

	err = json.Unmarshal(data, &varWorkgroupConnectionDto)

	if err != nil {
		return err
	}

	*o = WorkgroupConnectionDto(varWorkgroupConnectionDto)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "object")
		delete(additionalProperties, "connectionType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkgroupConnectionDto struct {
	value *WorkgroupConnectionDto
	isSet bool
}

func (v NullableWorkgroupConnectionDto) Get() *WorkgroupConnectionDto {
	return v.value
}

func (v *NullableWorkgroupConnectionDto) Set(val *WorkgroupConnectionDto) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkgroupConnectionDto) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkgroupConnectionDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkgroupConnectionDto(val *WorkgroupConnectionDto) *NullableWorkgroupConnectionDto {
	return &NullableWorkgroupConnectionDto{value: val, isSet: true}
}

func (v NullableWorkgroupConnectionDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkgroupConnectionDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


