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

// checks if the EntitlementDocumentAllOfSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitlementDocumentAllOfSource{}

// EntitlementDocumentAllOfSource Entitlement's source.
type EntitlementDocumentAllOfSource struct {
	// ID of entitlement's source.
	Id *string `json:"id,omitempty"`
	// Display name of entitlement's source.
	Name *string `json:"name,omitempty"`
	// Type of object.
	Type *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EntitlementDocumentAllOfSource EntitlementDocumentAllOfSource

// NewEntitlementDocumentAllOfSource instantiates a new EntitlementDocumentAllOfSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementDocumentAllOfSource() *EntitlementDocumentAllOfSource {
	this := EntitlementDocumentAllOfSource{}
	return &this
}

// NewEntitlementDocumentAllOfSourceWithDefaults instantiates a new EntitlementDocumentAllOfSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementDocumentAllOfSourceWithDefaults() *EntitlementDocumentAllOfSource {
	this := EntitlementDocumentAllOfSource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EntitlementDocumentAllOfSource) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDocumentAllOfSource) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EntitlementDocumentAllOfSource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EntitlementDocumentAllOfSource) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EntitlementDocumentAllOfSource) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDocumentAllOfSource) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EntitlementDocumentAllOfSource) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EntitlementDocumentAllOfSource) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *EntitlementDocumentAllOfSource) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDocumentAllOfSource) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *EntitlementDocumentAllOfSource) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *EntitlementDocumentAllOfSource) SetType(v string) {
	o.Type = &v
}

func (o EntitlementDocumentAllOfSource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitlementDocumentAllOfSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EntitlementDocumentAllOfSource) UnmarshalJSON(data []byte) (err error) {
	varEntitlementDocumentAllOfSource := _EntitlementDocumentAllOfSource{}

	err = json.Unmarshal(data, &varEntitlementDocumentAllOfSource)

	if err != nil {
		return err
	}

	*o = EntitlementDocumentAllOfSource(varEntitlementDocumentAllOfSource)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntitlementDocumentAllOfSource struct {
	value *EntitlementDocumentAllOfSource
	isSet bool
}

func (v NullableEntitlementDocumentAllOfSource) Get() *EntitlementDocumentAllOfSource {
	return v.value
}

func (v *NullableEntitlementDocumentAllOfSource) Set(val *EntitlementDocumentAllOfSource) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementDocumentAllOfSource) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementDocumentAllOfSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementDocumentAllOfSource(val *EntitlementDocumentAllOfSource) *NullableEntitlementDocumentAllOfSource {
	return &NullableEntitlementDocumentAllOfSource{value: val, isSet: true}
}

func (v NullableEntitlementDocumentAllOfSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementDocumentAllOfSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


