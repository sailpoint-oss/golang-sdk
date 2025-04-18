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

// checks if the NonEmployeeSourceLite type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NonEmployeeSourceLite{}

// NonEmployeeSourceLite struct for NonEmployeeSourceLite
type NonEmployeeSourceLite struct {
	// Non-Employee source id.
	Id *string `json:"id,omitempty"`
	// Source Id associated with this non-employee source.
	SourceId *string `json:"sourceId,omitempty"`
	// Source name associated with this non-employee source.
	Name *string `json:"name,omitempty"`
	// Source description associated with this non-employee source.
	Description *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NonEmployeeSourceLite NonEmployeeSourceLite

// NewNonEmployeeSourceLite instantiates a new NonEmployeeSourceLite object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNonEmployeeSourceLite() *NonEmployeeSourceLite {
	this := NonEmployeeSourceLite{}
	return &this
}

// NewNonEmployeeSourceLiteWithDefaults instantiates a new NonEmployeeSourceLite object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNonEmployeeSourceLiteWithDefaults() *NonEmployeeSourceLite {
	this := NonEmployeeSourceLite{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NonEmployeeSourceLite) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeSourceLite) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NonEmployeeSourceLite) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NonEmployeeSourceLite) SetId(v string) {
	o.Id = &v
}

// GetSourceId returns the SourceId field value if set, zero value otherwise.
func (o *NonEmployeeSourceLite) GetSourceId() string {
	if o == nil || IsNil(o.SourceId) {
		var ret string
		return ret
	}
	return *o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeSourceLite) GetSourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.SourceId) {
		return nil, false
	}
	return o.SourceId, true
}

// HasSourceId returns a boolean if a field has been set.
func (o *NonEmployeeSourceLite) HasSourceId() bool {
	if o != nil && !IsNil(o.SourceId) {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given string and assigns it to the SourceId field.
func (o *NonEmployeeSourceLite) SetSourceId(v string) {
	o.SourceId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NonEmployeeSourceLite) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeSourceLite) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NonEmployeeSourceLite) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NonEmployeeSourceLite) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NonEmployeeSourceLite) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeSourceLite) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NonEmployeeSourceLite) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NonEmployeeSourceLite) SetDescription(v string) {
	o.Description = &v
}

func (o NonEmployeeSourceLite) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NonEmployeeSourceLite) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.SourceId) {
		toSerialize["sourceId"] = o.SourceId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NonEmployeeSourceLite) UnmarshalJSON(data []byte) (err error) {
	varNonEmployeeSourceLite := _NonEmployeeSourceLite{}

	err = json.Unmarshal(data, &varNonEmployeeSourceLite)

	if err != nil {
		return err
	}

	*o = NonEmployeeSourceLite(varNonEmployeeSourceLite)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "sourceId")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNonEmployeeSourceLite struct {
	value *NonEmployeeSourceLite
	isSet bool
}

func (v NullableNonEmployeeSourceLite) Get() *NonEmployeeSourceLite {
	return v.value
}

func (v *NullableNonEmployeeSourceLite) Set(val *NonEmployeeSourceLite) {
	v.value = val
	v.isSet = true
}

func (v NullableNonEmployeeSourceLite) IsSet() bool {
	return v.isSet
}

func (v *NullableNonEmployeeSourceLite) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNonEmployeeSourceLite(val *NonEmployeeSourceLite) *NullableNonEmployeeSourceLite {
	return &NullableNonEmployeeSourceLite{value: val, isSet: true}
}

func (v NullableNonEmployeeSourceLite) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNonEmployeeSourceLite) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


