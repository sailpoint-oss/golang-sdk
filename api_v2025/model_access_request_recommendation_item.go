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

// checks if the AccessRequestRecommendationItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessRequestRecommendationItem{}

// AccessRequestRecommendationItem struct for AccessRequestRecommendationItem
type AccessRequestRecommendationItem struct {
	// ID of access item being recommended.
	Id *string `json:"id,omitempty"`
	Type *AccessRequestRecommendationItemType `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccessRequestRecommendationItem AccessRequestRecommendationItem

// NewAccessRequestRecommendationItem instantiates a new AccessRequestRecommendationItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessRequestRecommendationItem() *AccessRequestRecommendationItem {
	this := AccessRequestRecommendationItem{}
	return &this
}

// NewAccessRequestRecommendationItemWithDefaults instantiates a new AccessRequestRecommendationItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessRequestRecommendationItemWithDefaults() *AccessRequestRecommendationItem {
	this := AccessRequestRecommendationItem{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AccessRequestRecommendationItem) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequestRecommendationItem) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AccessRequestRecommendationItem) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AccessRequestRecommendationItem) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AccessRequestRecommendationItem) GetType() AccessRequestRecommendationItemType {
	if o == nil || IsNil(o.Type) {
		var ret AccessRequestRecommendationItemType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequestRecommendationItem) GetTypeOk() (*AccessRequestRecommendationItemType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AccessRequestRecommendationItem) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given AccessRequestRecommendationItemType and assigns it to the Type field.
func (o *AccessRequestRecommendationItem) SetType(v AccessRequestRecommendationItemType) {
	o.Type = &v
}

func (o AccessRequestRecommendationItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessRequestRecommendationItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccessRequestRecommendationItem) UnmarshalJSON(data []byte) (err error) {
	varAccessRequestRecommendationItem := _AccessRequestRecommendationItem{}

	err = json.Unmarshal(data, &varAccessRequestRecommendationItem)

	if err != nil {
		return err
	}

	*o = AccessRequestRecommendationItem(varAccessRequestRecommendationItem)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccessRequestRecommendationItem struct {
	value *AccessRequestRecommendationItem
	isSet bool
}

func (v NullableAccessRequestRecommendationItem) Get() *AccessRequestRecommendationItem {
	return v.value
}

func (v *NullableAccessRequestRecommendationItem) Set(val *AccessRequestRecommendationItem) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessRequestRecommendationItem) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessRequestRecommendationItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessRequestRecommendationItem(val *AccessRequestRecommendationItem) *NullableAccessRequestRecommendationItem {
	return &NullableAccessRequestRecommendationItem{value: val, isSet: true}
}

func (v NullableAccessRequestRecommendationItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessRequestRecommendationItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


