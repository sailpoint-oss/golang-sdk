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

// checks if the ConfigObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigObject{}

// ConfigObject Config export and import format for individual object configurations.
type ConfigObject struct {
	// Current version of configuration object.
	Version *int32 `json:"version,omitempty"`
	Self *SelfImportExportDto `json:"self,omitempty"`
	// Object details. Format dependant on the object type.
	Object map[string]interface{} `json:"object,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConfigObject ConfigObject

// NewConfigObject instantiates a new ConfigObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigObject() *ConfigObject {
	this := ConfigObject{}
	return &this
}

// NewConfigObjectWithDefaults instantiates a new ConfigObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigObjectWithDefaults() *ConfigObject {
	this := ConfigObject{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ConfigObject) GetVersion() int32 {
	if o == nil || IsNil(o.Version) {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigObject) GetVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ConfigObject) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *ConfigObject) SetVersion(v int32) {
	o.Version = &v
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *ConfigObject) GetSelf() SelfImportExportDto {
	if o == nil || IsNil(o.Self) {
		var ret SelfImportExportDto
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigObject) GetSelfOk() (*SelfImportExportDto, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *ConfigObject) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given SelfImportExportDto and assigns it to the Self field.
func (o *ConfigObject) SetSelf(v SelfImportExportDto) {
	o.Self = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *ConfigObject) GetObject() map[string]interface{} {
	if o == nil || IsNil(o.Object) {
		var ret map[string]interface{}
		return ret
	}
	return o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigObject) GetObjectOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Object) {
		return map[string]interface{}{}, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *ConfigObject) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given map[string]interface{} and assigns it to the Object field.
func (o *ConfigObject) SetObject(v map[string]interface{}) {
	o.Object = v
}

func (o ConfigObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConfigObject) UnmarshalJSON(data []byte) (err error) {
	varConfigObject := _ConfigObject{}

	err = json.Unmarshal(data, &varConfigObject)

	if err != nil {
		return err
	}

	*o = ConfigObject(varConfigObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "version")
		delete(additionalProperties, "self")
		delete(additionalProperties, "object")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConfigObject struct {
	value *ConfigObject
	isSet bool
}

func (v NullableConfigObject) Get() *ConfigObject {
	return v.value
}

func (v *NullableConfigObject) Set(val *ConfigObject) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigObject) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigObject(val *ConfigObject) *NullableConfigObject {
	return &NullableConfigObject{value: val, isSet: true}
}

func (v NullableConfigObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


