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

// checks if the SubscriptionPatchRequestInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionPatchRequestInner{}

// SubscriptionPatchRequestInner A JSONPatch Operation as defined by [RFC 6902 - JSON Patch](https://tools.ietf.org/html/rfc6902)
type SubscriptionPatchRequestInner struct {
	// The operation to be performed
	Op string `json:"op"`
	// A string JSON Pointer representing the target path to an element to be affected by the operation
	Path string `json:"path"`
	Value *SubscriptionPatchRequestInnerValue `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SubscriptionPatchRequestInner SubscriptionPatchRequestInner

// NewSubscriptionPatchRequestInner instantiates a new SubscriptionPatchRequestInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionPatchRequestInner(op string, path string) *SubscriptionPatchRequestInner {
	this := SubscriptionPatchRequestInner{}
	this.Op = op
	this.Path = path
	return &this
}

// NewSubscriptionPatchRequestInnerWithDefaults instantiates a new SubscriptionPatchRequestInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionPatchRequestInnerWithDefaults() *SubscriptionPatchRequestInner {
	this := SubscriptionPatchRequestInner{}
	return &this
}

// GetOp returns the Op field value
func (o *SubscriptionPatchRequestInner) GetOp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Op
}

// GetOpOk returns a tuple with the Op field value
// and a boolean to check if the value has been set.
func (o *SubscriptionPatchRequestInner) GetOpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Op, true
}

// SetOp sets field value
func (o *SubscriptionPatchRequestInner) SetOp(v string) {
	o.Op = v
}

// GetPath returns the Path field value
func (o *SubscriptionPatchRequestInner) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *SubscriptionPatchRequestInner) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *SubscriptionPatchRequestInner) SetPath(v string) {
	o.Path = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SubscriptionPatchRequestInner) GetValue() SubscriptionPatchRequestInnerValue {
	if o == nil || IsNil(o.Value) {
		var ret SubscriptionPatchRequestInnerValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPatchRequestInner) GetValueOk() (*SubscriptionPatchRequestInnerValue, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SubscriptionPatchRequestInner) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given SubscriptionPatchRequestInnerValue and assigns it to the Value field.
func (o *SubscriptionPatchRequestInner) SetValue(v SubscriptionPatchRequestInnerValue) {
	o.Value = &v
}

func (o SubscriptionPatchRequestInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionPatchRequestInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["op"] = o.Op
	toSerialize["path"] = o.Path
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SubscriptionPatchRequestInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"op",
		"path",
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

	varSubscriptionPatchRequestInner := _SubscriptionPatchRequestInner{}

	err = json.Unmarshal(data, &varSubscriptionPatchRequestInner)

	if err != nil {
		return err
	}

	*o = SubscriptionPatchRequestInner(varSubscriptionPatchRequestInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "op")
		delete(additionalProperties, "path")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSubscriptionPatchRequestInner struct {
	value *SubscriptionPatchRequestInner
	isSet bool
}

func (v NullableSubscriptionPatchRequestInner) Get() *SubscriptionPatchRequestInner {
	return v.value
}

func (v *NullableSubscriptionPatchRequestInner) Set(val *SubscriptionPatchRequestInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionPatchRequestInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionPatchRequestInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionPatchRequestInner(val *SubscriptionPatchRequestInner) *NullableSubscriptionPatchRequestInner {
	return &NullableSubscriptionPatchRequestInner{value: val, isSet: true}
}

func (v NullableSubscriptionPatchRequestInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionPatchRequestInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


