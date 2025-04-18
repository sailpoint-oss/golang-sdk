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

// checks if the Column type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Column{}

// Column struct for Column
type Column struct {
	// The name of the field. 
	Field string `json:"field"`
	// The value of the header. 
	Header *string `json:"header,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Column Column

// NewColumn instantiates a new Column object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewColumn(field string) *Column {
	this := Column{}
	this.Field = field
	return &this
}

// NewColumnWithDefaults instantiates a new Column object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewColumnWithDefaults() *Column {
	this := Column{}
	return &this
}

// GetField returns the Field field value
func (o *Column) GetField() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Field
}

// GetFieldOk returns a tuple with the Field field value
// and a boolean to check if the value has been set.
func (o *Column) GetFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Field, true
}

// SetField sets field value
func (o *Column) SetField(v string) {
	o.Field = v
}

// GetHeader returns the Header field value if set, zero value otherwise.
func (o *Column) GetHeader() string {
	if o == nil || IsNil(o.Header) {
		var ret string
		return ret
	}
	return *o.Header
}

// GetHeaderOk returns a tuple with the Header field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Column) GetHeaderOk() (*string, bool) {
	if o == nil || IsNil(o.Header) {
		return nil, false
	}
	return o.Header, true
}

// HasHeader returns a boolean if a field has been set.
func (o *Column) HasHeader() bool {
	if o != nil && !IsNil(o.Header) {
		return true
	}

	return false
}

// SetHeader gets a reference to the given string and assigns it to the Header field.
func (o *Column) SetHeader(v string) {
	o.Header = &v
}

func (o Column) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Column) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["field"] = o.Field
	if !IsNil(o.Header) {
		toSerialize["header"] = o.Header
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Column) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"field",
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

	varColumn := _Column{}

	err = json.Unmarshal(data, &varColumn)

	if err != nil {
		return err
	}

	*o = Column(varColumn)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "field")
		delete(additionalProperties, "header")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableColumn struct {
	value *Column
	isSet bool
}

func (v NullableColumn) Get() *Column {
	return v.value
}

func (v *NullableColumn) Set(val *Column) {
	v.value = val
	v.isSet = true
}

func (v NullableColumn) IsSet() bool {
	return v.isSet
}

func (v *NullableColumn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableColumn(val *Column) *NullableColumn {
	return &NullableColumn{value: val, isSet: true}
}

func (v NullableColumn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableColumn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


