/*
Identity Security Cloud Beta API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_beta

import (
	"encoding/json"
)

// checks if the BulkTaggedObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkTaggedObject{}

// BulkTaggedObject struct for BulkTaggedObject
type BulkTaggedObject struct {
	ObjectRefs []TaggedObjectDto `json:"objectRefs,omitempty"`
	// Label to be applied to object.
	Tags []string `json:"tags,omitempty"`
	// If APPEND, tags are appended to the list of tags for the object. A 400 error is returned if this would add duplicate tags to the object.  If MERGE, tags are merged with the existing tags. Duplicate tags are silently ignored.
	Operation *string `json:"operation,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BulkTaggedObject BulkTaggedObject

// NewBulkTaggedObject instantiates a new BulkTaggedObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkTaggedObject() *BulkTaggedObject {
	this := BulkTaggedObject{}
	var operation string = "APPEND"
	this.Operation = &operation
	return &this
}

// NewBulkTaggedObjectWithDefaults instantiates a new BulkTaggedObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkTaggedObjectWithDefaults() *BulkTaggedObject {
	this := BulkTaggedObject{}
	var operation string = "APPEND"
	this.Operation = &operation
	return &this
}

// GetObjectRefs returns the ObjectRefs field value if set, zero value otherwise.
func (o *BulkTaggedObject) GetObjectRefs() []TaggedObjectDto {
	if o == nil || IsNil(o.ObjectRefs) {
		var ret []TaggedObjectDto
		return ret
	}
	return o.ObjectRefs
}

// GetObjectRefsOk returns a tuple with the ObjectRefs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkTaggedObject) GetObjectRefsOk() ([]TaggedObjectDto, bool) {
	if o == nil || IsNil(o.ObjectRefs) {
		return nil, false
	}
	return o.ObjectRefs, true
}

// HasObjectRefs returns a boolean if a field has been set.
func (o *BulkTaggedObject) HasObjectRefs() bool {
	if o != nil && !IsNil(o.ObjectRefs) {
		return true
	}

	return false
}

// SetObjectRefs gets a reference to the given []TaggedObjectDto and assigns it to the ObjectRefs field.
func (o *BulkTaggedObject) SetObjectRefs(v []TaggedObjectDto) {
	o.ObjectRefs = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *BulkTaggedObject) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkTaggedObject) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *BulkTaggedObject) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *BulkTaggedObject) SetTags(v []string) {
	o.Tags = v
}

// GetOperation returns the Operation field value if set, zero value otherwise.
func (o *BulkTaggedObject) GetOperation() string {
	if o == nil || IsNil(o.Operation) {
		var ret string
		return ret
	}
	return *o.Operation
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkTaggedObject) GetOperationOk() (*string, bool) {
	if o == nil || IsNil(o.Operation) {
		return nil, false
	}
	return o.Operation, true
}

// HasOperation returns a boolean if a field has been set.
func (o *BulkTaggedObject) HasOperation() bool {
	if o != nil && !IsNil(o.Operation) {
		return true
	}

	return false
}

// SetOperation gets a reference to the given string and assigns it to the Operation field.
func (o *BulkTaggedObject) SetOperation(v string) {
	o.Operation = &v
}

func (o BulkTaggedObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkTaggedObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ObjectRefs) {
		toSerialize["objectRefs"] = o.ObjectRefs
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Operation) {
		toSerialize["operation"] = o.Operation
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BulkTaggedObject) UnmarshalJSON(data []byte) (err error) {
	varBulkTaggedObject := _BulkTaggedObject{}

	err = json.Unmarshal(data, &varBulkTaggedObject)

	if err != nil {
		return err
	}

	*o = BulkTaggedObject(varBulkTaggedObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "objectRefs")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "operation")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBulkTaggedObject struct {
	value *BulkTaggedObject
	isSet bool
}

func (v NullableBulkTaggedObject) Get() *BulkTaggedObject {
	return v.value
}

func (v *NullableBulkTaggedObject) Set(val *BulkTaggedObject) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkTaggedObject) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkTaggedObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkTaggedObject(val *BulkTaggedObject) *NullableBulkTaggedObject {
	return &NullableBulkTaggedObject{value: val, isSet: true}
}

func (v NullableBulkTaggedObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkTaggedObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


