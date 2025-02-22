/*
Identity Security Cloud Beta API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_beta

import (
	"encoding/json"
	
	"fmt"
)

// checks if the Tag type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Tag{}

// Tag struct for Tag
type Tag struct {
	// Tag id
	Id string `json:"id"`
	// Name of the tag.
	Name string `json:"name"`
	// Date the tag was created.
	Created SailPointTime `json:"created"`
	// Date the tag was last modified.
	Modified SailPointTime `json:"modified"`
	TagCategoryRefs []TagTagCategoryRefsInner `json:"tagCategoryRefs"`
	AdditionalProperties map[string]interface{}
}

type _Tag Tag

// NewTag instantiates a new Tag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTag(id string, name string, created SailPointTime, modified SailPointTime, tagCategoryRefs []TagTagCategoryRefsInner) *Tag {
	this := Tag{}
	this.Id = id
	this.Name = name
	this.Created = created
	this.Modified = modified
	this.TagCategoryRefs = tagCategoryRefs
	return &this
}

// NewTagWithDefaults instantiates a new Tag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagWithDefaults() *Tag {
	this := Tag{}
	return &this
}

// GetId returns the Id field value
func (o *Tag) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Tag) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Tag) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Tag) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Tag) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Tag) SetName(v string) {
	o.Name = v
}

// GetCreated returns the Created field value
func (o *Tag) GetCreated() SailPointTime {
	if o == nil {
		var ret SailPointTime
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *Tag) GetCreatedOk() (*SailPointTime, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *Tag) SetCreated(v SailPointTime) {
	o.Created = v
}

// GetModified returns the Modified field value
func (o *Tag) GetModified() SailPointTime {
	if o == nil {
		var ret SailPointTime
		return ret
	}

	return o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value
// and a boolean to check if the value has been set.
func (o *Tag) GetModifiedOk() (*SailPointTime, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Modified, true
}

// SetModified sets field value
func (o *Tag) SetModified(v SailPointTime) {
	o.Modified = v
}

// GetTagCategoryRefs returns the TagCategoryRefs field value
func (o *Tag) GetTagCategoryRefs() []TagTagCategoryRefsInner {
	if o == nil {
		var ret []TagTagCategoryRefsInner
		return ret
	}

	return o.TagCategoryRefs
}

// GetTagCategoryRefsOk returns a tuple with the TagCategoryRefs field value
// and a boolean to check if the value has been set.
func (o *Tag) GetTagCategoryRefsOk() ([]TagTagCategoryRefsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.TagCategoryRefs, true
}

// SetTagCategoryRefs sets field value
func (o *Tag) SetTagCategoryRefs(v []TagTagCategoryRefsInner) {
	o.TagCategoryRefs = v
}

func (o Tag) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Tag) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["created"] = o.Created
	toSerialize["modified"] = o.Modified
	toSerialize["tagCategoryRefs"] = o.TagCategoryRefs

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Tag) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"created",
		"modified",
		"tagCategoryRefs",
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

	varTag := _Tag{}

	err = json.Unmarshal(data, &varTag)

	if err != nil {
		return err
	}

	*o = Tag(varTag)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "created")
		delete(additionalProperties, "modified")
		delete(additionalProperties, "tagCategoryRefs")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTag struct {
	value *Tag
	isSet bool
}

func (v NullableTag) Get() *Tag {
	return v.value
}

func (v *NullableTag) Set(val *Tag) {
	v.value = val
	v.isSet = true
}

func (v NullableTag) IsSet() bool {
	return v.isSet
}

func (v *NullableTag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTag(val *Tag) *NullableTag {
	return &NullableTag{value: val, isSet: true}
}

func (v NullableTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


