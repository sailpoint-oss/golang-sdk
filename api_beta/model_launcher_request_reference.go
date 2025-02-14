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

// checks if the LauncherRequestReference type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LauncherRequestReference{}

// LauncherRequestReference struct for LauncherRequestReference
type LauncherRequestReference struct {
	// Type of Launcher reference
	Type string `json:"type"`
	// ID of Launcher reference
	Id string `json:"id"`
	AdditionalProperties map[string]interface{}
}

type _LauncherRequestReference LauncherRequestReference

// NewLauncherRequestReference instantiates a new LauncherRequestReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLauncherRequestReference(type_ string, id string) *LauncherRequestReference {
	this := LauncherRequestReference{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewLauncherRequestReferenceWithDefaults instantiates a new LauncherRequestReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLauncherRequestReferenceWithDefaults() *LauncherRequestReference {
	this := LauncherRequestReference{}
	return &this
}

// GetType returns the Type field value
func (o *LauncherRequestReference) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LauncherRequestReference) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *LauncherRequestReference) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *LauncherRequestReference) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LauncherRequestReference) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LauncherRequestReference) SetId(v string) {
	o.Id = v
}

func (o LauncherRequestReference) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LauncherRequestReference) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LauncherRequestReference) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"id",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varLauncherRequestReference := _LauncherRequestReference{}

	err = json.Unmarshal(data, &varLauncherRequestReference)

	if err != nil {
		return err
	}

	*o = LauncherRequestReference(varLauncherRequestReference)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLauncherRequestReference struct {
	value *LauncherRequestReference
	isSet bool
}

func (v NullableLauncherRequestReference) Get() *LauncherRequestReference {
	return v.value
}

func (v *NullableLauncherRequestReference) Set(val *LauncherRequestReference) {
	v.value = val
	v.isSet = true
}

func (v NullableLauncherRequestReference) IsSet() bool {
	return v.isSet
}

func (v *NullableLauncherRequestReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLauncherRequestReference(val *LauncherRequestReference) *NullableLauncherRequestReference {
	return &NullableLauncherRequestReference{value: val, isSet: true}
}

func (v NullableLauncherRequestReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLauncherRequestReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


