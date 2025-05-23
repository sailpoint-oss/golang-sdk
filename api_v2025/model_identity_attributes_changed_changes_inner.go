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

// checks if the IdentityAttributesChangedChangesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentityAttributesChangedChangesInner{}

// IdentityAttributesChangedChangesInner struct for IdentityAttributesChangedChangesInner
type IdentityAttributesChangedChangesInner struct {
	// The name of the identity attribute that changed.
	Attribute string `json:"attribute"`
	OldValue NullableIdentityAttributesChangedChangesInnerOldValue `json:"oldValue,omitempty"`
	NewValue *IdentityAttributesChangedChangesInnerNewValue `json:"newValue,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdentityAttributesChangedChangesInner IdentityAttributesChangedChangesInner

// NewIdentityAttributesChangedChangesInner instantiates a new IdentityAttributesChangedChangesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityAttributesChangedChangesInner(attribute string) *IdentityAttributesChangedChangesInner {
	this := IdentityAttributesChangedChangesInner{}
	this.Attribute = attribute
	return &this
}

// NewIdentityAttributesChangedChangesInnerWithDefaults instantiates a new IdentityAttributesChangedChangesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityAttributesChangedChangesInnerWithDefaults() *IdentityAttributesChangedChangesInner {
	this := IdentityAttributesChangedChangesInner{}
	return &this
}

// GetAttribute returns the Attribute field value
func (o *IdentityAttributesChangedChangesInner) GetAttribute() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Attribute
}

// GetAttributeOk returns a tuple with the Attribute field value
// and a boolean to check if the value has been set.
func (o *IdentityAttributesChangedChangesInner) GetAttributeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attribute, true
}

// SetAttribute sets field value
func (o *IdentityAttributesChangedChangesInner) SetAttribute(v string) {
	o.Attribute = v
}

// GetOldValue returns the OldValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityAttributesChangedChangesInner) GetOldValue() IdentityAttributesChangedChangesInnerOldValue {
	if o == nil || IsNil(o.OldValue.Get()) {
		var ret IdentityAttributesChangedChangesInnerOldValue
		return ret
	}
	return *o.OldValue.Get()
}

// GetOldValueOk returns a tuple with the OldValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityAttributesChangedChangesInner) GetOldValueOk() (*IdentityAttributesChangedChangesInnerOldValue, bool) {
	if o == nil {
		return nil, false
	}
	return o.OldValue.Get(), o.OldValue.IsSet()
}

// HasOldValue returns a boolean if a field has been set.
func (o *IdentityAttributesChangedChangesInner) HasOldValue() bool {
	if o != nil && o.OldValue.IsSet() {
		return true
	}

	return false
}

// SetOldValue gets a reference to the given NullableIdentityAttributesChangedChangesInnerOldValue and assigns it to the OldValue field.
func (o *IdentityAttributesChangedChangesInner) SetOldValue(v IdentityAttributesChangedChangesInnerOldValue) {
	o.OldValue.Set(&v)
}
// SetOldValueNil sets the value for OldValue to be an explicit nil
func (o *IdentityAttributesChangedChangesInner) SetOldValueNil() {
	o.OldValue.Set(nil)
}

// UnsetOldValue ensures that no value is present for OldValue, not even an explicit nil
func (o *IdentityAttributesChangedChangesInner) UnsetOldValue() {
	o.OldValue.Unset()
}

// GetNewValue returns the NewValue field value if set, zero value otherwise.
func (o *IdentityAttributesChangedChangesInner) GetNewValue() IdentityAttributesChangedChangesInnerNewValue {
	if o == nil || IsNil(o.NewValue) {
		var ret IdentityAttributesChangedChangesInnerNewValue
		return ret
	}
	return *o.NewValue
}

// GetNewValueOk returns a tuple with the NewValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityAttributesChangedChangesInner) GetNewValueOk() (*IdentityAttributesChangedChangesInnerNewValue, bool) {
	if o == nil || IsNil(o.NewValue) {
		return nil, false
	}
	return o.NewValue, true
}

// HasNewValue returns a boolean if a field has been set.
func (o *IdentityAttributesChangedChangesInner) HasNewValue() bool {
	if o != nil && !IsNil(o.NewValue) {
		return true
	}

	return false
}

// SetNewValue gets a reference to the given IdentityAttributesChangedChangesInnerNewValue and assigns it to the NewValue field.
func (o *IdentityAttributesChangedChangesInner) SetNewValue(v IdentityAttributesChangedChangesInnerNewValue) {
	o.NewValue = &v
}

func (o IdentityAttributesChangedChangesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentityAttributesChangedChangesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["attribute"] = o.Attribute
	if o.OldValue.IsSet() {
		toSerialize["oldValue"] = o.OldValue.Get()
	}
	if !IsNil(o.NewValue) {
		toSerialize["newValue"] = o.NewValue
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IdentityAttributesChangedChangesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"attribute",
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

	varIdentityAttributesChangedChangesInner := _IdentityAttributesChangedChangesInner{}

	err = json.Unmarshal(data, &varIdentityAttributesChangedChangesInner)

	if err != nil {
		return err
	}

	*o = IdentityAttributesChangedChangesInner(varIdentityAttributesChangedChangesInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "attribute")
		delete(additionalProperties, "oldValue")
		delete(additionalProperties, "newValue")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentityAttributesChangedChangesInner struct {
	value *IdentityAttributesChangedChangesInner
	isSet bool
}

func (v NullableIdentityAttributesChangedChangesInner) Get() *IdentityAttributesChangedChangesInner {
	return v.value
}

func (v *NullableIdentityAttributesChangedChangesInner) Set(val *IdentityAttributesChangedChangesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityAttributesChangedChangesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityAttributesChangedChangesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityAttributesChangedChangesInner(val *IdentityAttributesChangedChangesInner) *NullableIdentityAttributesChangedChangesInner {
	return &NullableIdentityAttributesChangedChangesInner{value: val, isSet: true}
}

func (v NullableIdentityAttributesChangedChangesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityAttributesChangedChangesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


