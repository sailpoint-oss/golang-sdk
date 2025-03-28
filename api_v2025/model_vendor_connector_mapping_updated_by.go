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

// checks if the VendorConnectorMappingUpdatedBy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VendorConnectorMappingUpdatedBy{}

// VendorConnectorMappingUpdatedBy An object representing the nullable identifier of the user who last updated the mapping.
type VendorConnectorMappingUpdatedBy struct {
	// The identifier of the user who last updated the mapping, if available.
	String *string `json:"String,omitempty"`
	// A flag indicating if the 'String' field is set and valid.
	Valid *bool `json:"Valid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VendorConnectorMappingUpdatedBy VendorConnectorMappingUpdatedBy

// NewVendorConnectorMappingUpdatedBy instantiates a new VendorConnectorMappingUpdatedBy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVendorConnectorMappingUpdatedBy() *VendorConnectorMappingUpdatedBy {
	this := VendorConnectorMappingUpdatedBy{}
	var valid bool = false
	this.Valid = &valid
	return &this
}

// NewVendorConnectorMappingUpdatedByWithDefaults instantiates a new VendorConnectorMappingUpdatedBy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVendorConnectorMappingUpdatedByWithDefaults() *VendorConnectorMappingUpdatedBy {
	this := VendorConnectorMappingUpdatedBy{}
	var valid bool = false
	this.Valid = &valid
	return &this
}

// GetString returns the String field value if set, zero value otherwise.
func (o *VendorConnectorMappingUpdatedBy) GetString() string {
	if o == nil || IsNil(o.String) {
		var ret string
		return ret
	}
	return *o.String
}

// GetStringOk returns a tuple with the String field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VendorConnectorMappingUpdatedBy) GetStringOk() (*string, bool) {
	if o == nil || IsNil(o.String) {
		return nil, false
	}
	return o.String, true
}

// HasString returns a boolean if a field has been set.
func (o *VendorConnectorMappingUpdatedBy) HasString() bool {
	if o != nil && !IsNil(o.String) {
		return true
	}

	return false
}

// SetString gets a reference to the given string and assigns it to the String field.
func (o *VendorConnectorMappingUpdatedBy) SetString(v string) {
	o.String = &v
}

// GetValid returns the Valid field value if set, zero value otherwise.
func (o *VendorConnectorMappingUpdatedBy) GetValid() bool {
	if o == nil || IsNil(o.Valid) {
		var ret bool
		return ret
	}
	return *o.Valid
}

// GetValidOk returns a tuple with the Valid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VendorConnectorMappingUpdatedBy) GetValidOk() (*bool, bool) {
	if o == nil || IsNil(o.Valid) {
		return nil, false
	}
	return o.Valid, true
}

// HasValid returns a boolean if a field has been set.
func (o *VendorConnectorMappingUpdatedBy) HasValid() bool {
	if o != nil && !IsNil(o.Valid) {
		return true
	}

	return false
}

// SetValid gets a reference to the given bool and assigns it to the Valid field.
func (o *VendorConnectorMappingUpdatedBy) SetValid(v bool) {
	o.Valid = &v
}

func (o VendorConnectorMappingUpdatedBy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VendorConnectorMappingUpdatedBy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.String) {
		toSerialize["String"] = o.String
	}
	if !IsNil(o.Valid) {
		toSerialize["Valid"] = o.Valid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VendorConnectorMappingUpdatedBy) UnmarshalJSON(data []byte) (err error) {
	varVendorConnectorMappingUpdatedBy := _VendorConnectorMappingUpdatedBy{}

	err = json.Unmarshal(data, &varVendorConnectorMappingUpdatedBy)

	if err != nil {
		return err
	}

	*o = VendorConnectorMappingUpdatedBy(varVendorConnectorMappingUpdatedBy)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "String")
		delete(additionalProperties, "Valid")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVendorConnectorMappingUpdatedBy struct {
	value *VendorConnectorMappingUpdatedBy
	isSet bool
}

func (v NullableVendorConnectorMappingUpdatedBy) Get() *VendorConnectorMappingUpdatedBy {
	return v.value
}

func (v *NullableVendorConnectorMappingUpdatedBy) Set(val *VendorConnectorMappingUpdatedBy) {
	v.value = val
	v.isSet = true
}

func (v NullableVendorConnectorMappingUpdatedBy) IsSet() bool {
	return v.isSet
}

func (v *NullableVendorConnectorMappingUpdatedBy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVendorConnectorMappingUpdatedBy(val *VendorConnectorMappingUpdatedBy) *NullableVendorConnectorMappingUpdatedBy {
	return &NullableVendorConnectorMappingUpdatedBy{value: val, isSet: true}
}

func (v NullableVendorConnectorMappingUpdatedBy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVendorConnectorMappingUpdatedBy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


