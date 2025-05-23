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

// checks if the SodExemptCriteria type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SodExemptCriteria{}

// SodExemptCriteria Details of the Entitlement criteria
type SodExemptCriteria struct {
	// If the entitlement already belonged to the user or not.
	Existing *bool `json:"existing,omitempty"`
	Type *DtoType `json:"type,omitempty"`
	// Entitlement ID
	Id *string `json:"id,omitempty"`
	// Entitlement name
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SodExemptCriteria SodExemptCriteria

// NewSodExemptCriteria instantiates a new SodExemptCriteria object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSodExemptCriteria() *SodExemptCriteria {
	this := SodExemptCriteria{}
	var existing bool = false
	this.Existing = &existing
	return &this
}

// NewSodExemptCriteriaWithDefaults instantiates a new SodExemptCriteria object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSodExemptCriteriaWithDefaults() *SodExemptCriteria {
	this := SodExemptCriteria{}
	var existing bool = false
	this.Existing = &existing
	return &this
}

// GetExisting returns the Existing field value if set, zero value otherwise.
func (o *SodExemptCriteria) GetExisting() bool {
	if o == nil || IsNil(o.Existing) {
		var ret bool
		return ret
	}
	return *o.Existing
}

// GetExistingOk returns a tuple with the Existing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SodExemptCriteria) GetExistingOk() (*bool, bool) {
	if o == nil || IsNil(o.Existing) {
		return nil, false
	}
	return o.Existing, true
}

// HasExisting returns a boolean if a field has been set.
func (o *SodExemptCriteria) HasExisting() bool {
	if o != nil && !IsNil(o.Existing) {
		return true
	}

	return false
}

// SetExisting gets a reference to the given bool and assigns it to the Existing field.
func (o *SodExemptCriteria) SetExisting(v bool) {
	o.Existing = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SodExemptCriteria) GetType() DtoType {
	if o == nil || IsNil(o.Type) {
		var ret DtoType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SodExemptCriteria) GetTypeOk() (*DtoType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SodExemptCriteria) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given DtoType and assigns it to the Type field.
func (o *SodExemptCriteria) SetType(v DtoType) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SodExemptCriteria) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SodExemptCriteria) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SodExemptCriteria) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SodExemptCriteria) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SodExemptCriteria) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SodExemptCriteria) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SodExemptCriteria) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SodExemptCriteria) SetName(v string) {
	o.Name = &v
}

func (o SodExemptCriteria) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SodExemptCriteria) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Existing) {
		toSerialize["existing"] = o.Existing
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SodExemptCriteria) UnmarshalJSON(data []byte) (err error) {
	varSodExemptCriteria := _SodExemptCriteria{}

	err = json.Unmarshal(data, &varSodExemptCriteria)

	if err != nil {
		return err
	}

	*o = SodExemptCriteria(varSodExemptCriteria)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "existing")
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSodExemptCriteria struct {
	value *SodExemptCriteria
	isSet bool
}

func (v NullableSodExemptCriteria) Get() *SodExemptCriteria {
	return v.value
}

func (v *NullableSodExemptCriteria) Set(val *SodExemptCriteria) {
	v.value = val
	v.isSet = true
}

func (v NullableSodExemptCriteria) IsSet() bool {
	return v.isSet
}

func (v *NullableSodExemptCriteria) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSodExemptCriteria(val *SodExemptCriteria) *NullableSodExemptCriteria {
	return &NullableSodExemptCriteria{value: val, isSet: true}
}

func (v NullableSodExemptCriteria) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSodExemptCriteria) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


