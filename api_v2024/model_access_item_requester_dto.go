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

// checks if the AccessItemRequesterDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessItemRequesterDto{}

// AccessItemRequesterDto Access item requester's identity.
type AccessItemRequesterDto struct {
	// Access item requester's DTO type.
	Type *string `json:"type,omitempty"`
	// Access item requester's identity ID.
	Id *string `json:"id,omitempty"`
	// Access item owner's human-readable display name.
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccessItemRequesterDto AccessItemRequesterDto

// NewAccessItemRequesterDto instantiates a new AccessItemRequesterDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessItemRequesterDto() *AccessItemRequesterDto {
	this := AccessItemRequesterDto{}
	return &this
}

// NewAccessItemRequesterDtoWithDefaults instantiates a new AccessItemRequesterDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessItemRequesterDtoWithDefaults() *AccessItemRequesterDto {
	this := AccessItemRequesterDto{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AccessItemRequesterDto) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessItemRequesterDto) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AccessItemRequesterDto) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AccessItemRequesterDto) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AccessItemRequesterDto) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessItemRequesterDto) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AccessItemRequesterDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AccessItemRequesterDto) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AccessItemRequesterDto) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessItemRequesterDto) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AccessItemRequesterDto) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AccessItemRequesterDto) SetName(v string) {
	o.Name = &v
}

func (o AccessItemRequesterDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessItemRequesterDto) ToMap() (map[string]interface{}, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccessItemRequesterDto) UnmarshalJSON(data []byte) (err error) {
	varAccessItemRequesterDto := _AccessItemRequesterDto{}

	err = json.Unmarshal(data, &varAccessItemRequesterDto)

	if err != nil {
		return err
	}

	*o = AccessItemRequesterDto(varAccessItemRequesterDto)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccessItemRequesterDto struct {
	value *AccessItemRequesterDto
	isSet bool
}

func (v NullableAccessItemRequesterDto) Get() *AccessItemRequesterDto {
	return v.value
}

func (v *NullableAccessItemRequesterDto) Set(val *AccessItemRequesterDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessItemRequesterDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessItemRequesterDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessItemRequesterDto(val *AccessItemRequesterDto) *NullableAccessItemRequesterDto {
	return &NullableAccessItemRequesterDto{value: val, isSet: true}
}

func (v NullableAccessItemRequesterDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessItemRequesterDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

