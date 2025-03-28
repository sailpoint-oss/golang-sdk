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

// checks if the TaggedObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaggedObject{}

// TaggedObject Tagged object.
type TaggedObject struct {
	ObjectRef *TaggedObjectDto `json:"objectRef,omitempty"`
	// Labels to be applied to an Object
	Tags []string `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TaggedObject TaggedObject

// NewTaggedObject instantiates a new TaggedObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaggedObject() *TaggedObject {
	this := TaggedObject{}
	return &this
}

// NewTaggedObjectWithDefaults instantiates a new TaggedObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaggedObjectWithDefaults() *TaggedObject {
	this := TaggedObject{}
	return &this
}

// GetObjectRef returns the ObjectRef field value if set, zero value otherwise.
func (o *TaggedObject) GetObjectRef() TaggedObjectDto {
	if o == nil || IsNil(o.ObjectRef) {
		var ret TaggedObjectDto
		return ret
	}
	return *o.ObjectRef
}

// GetObjectRefOk returns a tuple with the ObjectRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaggedObject) GetObjectRefOk() (*TaggedObjectDto, bool) {
	if o == nil || IsNil(o.ObjectRef) {
		return nil, false
	}
	return o.ObjectRef, true
}

// HasObjectRef returns a boolean if a field has been set.
func (o *TaggedObject) HasObjectRef() bool {
	if o != nil && !IsNil(o.ObjectRef) {
		return true
	}

	return false
}

// SetObjectRef gets a reference to the given TaggedObjectDto and assigns it to the ObjectRef field.
func (o *TaggedObject) SetObjectRef(v TaggedObjectDto) {
	o.ObjectRef = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *TaggedObject) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaggedObject) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *TaggedObject) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *TaggedObject) SetTags(v []string) {
	o.Tags = v
}

func (o TaggedObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaggedObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ObjectRef) {
		toSerialize["objectRef"] = o.ObjectRef
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TaggedObject) UnmarshalJSON(data []byte) (err error) {
	varTaggedObject := _TaggedObject{}

	err = json.Unmarshal(data, &varTaggedObject)

	if err != nil {
		return err
	}

	*o = TaggedObject(varTaggedObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "objectRef")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTaggedObject struct {
	value *TaggedObject
	isSet bool
}

func (v NullableTaggedObject) Get() *TaggedObject {
	return v.value
}

func (v *NullableTaggedObject) Set(val *TaggedObject) {
	v.value = val
	v.isSet = true
}

func (v NullableTaggedObject) IsSet() bool {
	return v.isSet
}

func (v *NullableTaggedObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaggedObject(val *TaggedObject) *NullableTaggedObject {
	return &NullableTaggedObject{value: val, isSet: true}
}

func (v NullableTaggedObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaggedObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


