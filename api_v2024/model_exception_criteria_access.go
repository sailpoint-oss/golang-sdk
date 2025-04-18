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

// checks if the ExceptionCriteriaAccess type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExceptionCriteriaAccess{}

// ExceptionCriteriaAccess Access reference with addition of boolean existing flag to indicate whether the access was extant
type ExceptionCriteriaAccess struct {
	Type *DtoType `json:"type,omitempty"`
	// ID of the object to which this reference applies
	Id *string `json:"id,omitempty"`
	// Human-readable display name of the object to which this reference applies
	Name *string `json:"name,omitempty"`
	// Whether the subject identity already had that access or not
	Existing *bool `json:"existing,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ExceptionCriteriaAccess ExceptionCriteriaAccess

// NewExceptionCriteriaAccess instantiates a new ExceptionCriteriaAccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExceptionCriteriaAccess() *ExceptionCriteriaAccess {
	this := ExceptionCriteriaAccess{}
	var existing bool = false
	this.Existing = &existing
	return &this
}

// NewExceptionCriteriaAccessWithDefaults instantiates a new ExceptionCriteriaAccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExceptionCriteriaAccessWithDefaults() *ExceptionCriteriaAccess {
	this := ExceptionCriteriaAccess{}
	var existing bool = false
	this.Existing = &existing
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ExceptionCriteriaAccess) GetType() DtoType {
	if o == nil || IsNil(o.Type) {
		var ret DtoType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExceptionCriteriaAccess) GetTypeOk() (*DtoType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ExceptionCriteriaAccess) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given DtoType and assigns it to the Type field.
func (o *ExceptionCriteriaAccess) SetType(v DtoType) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ExceptionCriteriaAccess) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExceptionCriteriaAccess) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ExceptionCriteriaAccess) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ExceptionCriteriaAccess) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ExceptionCriteriaAccess) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExceptionCriteriaAccess) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ExceptionCriteriaAccess) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ExceptionCriteriaAccess) SetName(v string) {
	o.Name = &v
}

// GetExisting returns the Existing field value if set, zero value otherwise.
func (o *ExceptionCriteriaAccess) GetExisting() bool {
	if o == nil || IsNil(o.Existing) {
		var ret bool
		return ret
	}
	return *o.Existing
}

// GetExistingOk returns a tuple with the Existing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExceptionCriteriaAccess) GetExistingOk() (*bool, bool) {
	if o == nil || IsNil(o.Existing) {
		return nil, false
	}
	return o.Existing, true
}

// HasExisting returns a boolean if a field has been set.
func (o *ExceptionCriteriaAccess) HasExisting() bool {
	if o != nil && !IsNil(o.Existing) {
		return true
	}

	return false
}

// SetExisting gets a reference to the given bool and assigns it to the Existing field.
func (o *ExceptionCriteriaAccess) SetExisting(v bool) {
	o.Existing = &v
}

func (o ExceptionCriteriaAccess) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExceptionCriteriaAccess) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Existing) {
		toSerialize["existing"] = o.Existing
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ExceptionCriteriaAccess) UnmarshalJSON(data []byte) (err error) {
	varExceptionCriteriaAccess := _ExceptionCriteriaAccess{}

	err = json.Unmarshal(data, &varExceptionCriteriaAccess)

	if err != nil {
		return err
	}

	*o = ExceptionCriteriaAccess(varExceptionCriteriaAccess)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "existing")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExceptionCriteriaAccess struct {
	value *ExceptionCriteriaAccess
	isSet bool
}

func (v NullableExceptionCriteriaAccess) Get() *ExceptionCriteriaAccess {
	return v.value
}

func (v *NullableExceptionCriteriaAccess) Set(val *ExceptionCriteriaAccess) {
	v.value = val
	v.isSet = true
}

func (v NullableExceptionCriteriaAccess) IsSet() bool {
	return v.isSet
}

func (v *NullableExceptionCriteriaAccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExceptionCriteriaAccess(val *ExceptionCriteriaAccess) *NullableExceptionCriteriaAccess {
	return &NullableExceptionCriteriaAccess{value: val, isSet: true}
}

func (v NullableExceptionCriteriaAccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExceptionCriteriaAccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


