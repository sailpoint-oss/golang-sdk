/*
Identity Security Cloud V2024 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: v2024
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2024

import (
	"encoding/json"
)

// checks if the DimensionCriteriaLevel3 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DimensionCriteriaLevel3{}

// DimensionCriteriaLevel3 Defines STANDARD type Dimension membership
type DimensionCriteriaLevel3 struct {
	Operation *DimensionCriteriaOperation `json:"operation,omitempty"`
	Key NullableDimensionCriteriaKey `json:"key,omitempty"`
	// String value to test the Identity attribute specified in the key w/r/t the specified operation. If this criteria is a leaf node, that is, if the operation is one of EQUALS, this field is required. Otherwise, specifying it is an error.
	StringValue *string `json:"stringValue,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DimensionCriteriaLevel3 DimensionCriteriaLevel3

// NewDimensionCriteriaLevel3 instantiates a new DimensionCriteriaLevel3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDimensionCriteriaLevel3() *DimensionCriteriaLevel3 {
	this := DimensionCriteriaLevel3{}
	return &this
}

// NewDimensionCriteriaLevel3WithDefaults instantiates a new DimensionCriteriaLevel3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDimensionCriteriaLevel3WithDefaults() *DimensionCriteriaLevel3 {
	this := DimensionCriteriaLevel3{}
	return &this
}

// GetOperation returns the Operation field value if set, zero value otherwise.
func (o *DimensionCriteriaLevel3) GetOperation() DimensionCriteriaOperation {
	if o == nil || IsNil(o.Operation) {
		var ret DimensionCriteriaOperation
		return ret
	}
	return *o.Operation
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DimensionCriteriaLevel3) GetOperationOk() (*DimensionCriteriaOperation, bool) {
	if o == nil || IsNil(o.Operation) {
		return nil, false
	}
	return o.Operation, true
}

// HasOperation returns a boolean if a field has been set.
func (o *DimensionCriteriaLevel3) HasOperation() bool {
	if o != nil && !IsNil(o.Operation) {
		return true
	}

	return false
}

// SetOperation gets a reference to the given DimensionCriteriaOperation and assigns it to the Operation field.
func (o *DimensionCriteriaLevel3) SetOperation(v DimensionCriteriaOperation) {
	o.Operation = &v
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DimensionCriteriaLevel3) GetKey() DimensionCriteriaKey {
	if o == nil || IsNil(o.Key.Get()) {
		var ret DimensionCriteriaKey
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DimensionCriteriaLevel3) GetKeyOk() (*DimensionCriteriaKey, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *DimensionCriteriaLevel3) HasKey() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableDimensionCriteriaKey and assigns it to the Key field.
func (o *DimensionCriteriaLevel3) SetKey(v DimensionCriteriaKey) {
	o.Key.Set(&v)
}
// SetKeyNil sets the value for Key to be an explicit nil
func (o *DimensionCriteriaLevel3) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *DimensionCriteriaLevel3) UnsetKey() {
	o.Key.Unset()
}

// GetStringValue returns the StringValue field value if set, zero value otherwise.
func (o *DimensionCriteriaLevel3) GetStringValue() string {
	if o == nil || IsNil(o.StringValue) {
		var ret string
		return ret
	}
	return *o.StringValue
}

// GetStringValueOk returns a tuple with the StringValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DimensionCriteriaLevel3) GetStringValueOk() (*string, bool) {
	if o == nil || IsNil(o.StringValue) {
		return nil, false
	}
	return o.StringValue, true
}

// HasStringValue returns a boolean if a field has been set.
func (o *DimensionCriteriaLevel3) HasStringValue() bool {
	if o != nil && !IsNil(o.StringValue) {
		return true
	}

	return false
}

// SetStringValue gets a reference to the given string and assigns it to the StringValue field.
func (o *DimensionCriteriaLevel3) SetStringValue(v string) {
	o.StringValue = &v
}

func (o DimensionCriteriaLevel3) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DimensionCriteriaLevel3) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Operation) {
		toSerialize["operation"] = o.Operation
	}
	if o.Key.IsSet() {
		toSerialize["key"] = o.Key.Get()
	}
	if !IsNil(o.StringValue) {
		toSerialize["stringValue"] = o.StringValue
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DimensionCriteriaLevel3) UnmarshalJSON(data []byte) (err error) {
	varDimensionCriteriaLevel3 := _DimensionCriteriaLevel3{}

	err = json.Unmarshal(data, &varDimensionCriteriaLevel3)

	if err != nil {
		return err
	}

	*o = DimensionCriteriaLevel3(varDimensionCriteriaLevel3)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "operation")
		delete(additionalProperties, "key")
		delete(additionalProperties, "stringValue")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDimensionCriteriaLevel3 struct {
	value *DimensionCriteriaLevel3
	isSet bool
}

func (v NullableDimensionCriteriaLevel3) Get() *DimensionCriteriaLevel3 {
	return v.value
}

func (v *NullableDimensionCriteriaLevel3) Set(val *DimensionCriteriaLevel3) {
	v.value = val
	v.isSet = true
}

func (v NullableDimensionCriteriaLevel3) IsSet() bool {
	return v.isSet
}

func (v *NullableDimensionCriteriaLevel3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDimensionCriteriaLevel3(val *DimensionCriteriaLevel3) *NullableDimensionCriteriaLevel3 {
	return &NullableDimensionCriteriaLevel3{value: val, isSet: true}
}

func (v NullableDimensionCriteriaLevel3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDimensionCriteriaLevel3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

