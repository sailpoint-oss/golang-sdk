/*
SailPoint SaaS API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sailpointv2sdk

import (
	"encoding/json"
)

// SearchEventAttributesSSO struct for SearchEventAttributesSSO
type SearchEventAttributesSSO struct {
	ContextId *string `json:"contextId,omitempty"`
	IpAddress *string `json:"ipAddress,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SearchEventAttributesSSO SearchEventAttributesSSO

// NewSearchEventAttributesSSO instantiates a new SearchEventAttributesSSO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchEventAttributesSSO() *SearchEventAttributesSSO {
	this := SearchEventAttributesSSO{}
	return &this
}

// NewSearchEventAttributesSSOWithDefaults instantiates a new SearchEventAttributesSSO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchEventAttributesSSOWithDefaults() *SearchEventAttributesSSO {
	this := SearchEventAttributesSSO{}
	return &this
}

// GetContextId returns the ContextId field value if set, zero value otherwise.
func (o *SearchEventAttributesSSO) GetContextId() string {
	if o == nil || isNil(o.ContextId) {
		var ret string
		return ret
	}
	return *o.ContextId
}

// GetContextIdOk returns a tuple with the ContextId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchEventAttributesSSO) GetContextIdOk() (*string, bool) {
	if o == nil || isNil(o.ContextId) {
		return nil, false
	}
	return o.ContextId, true
}

// HasContextId returns a boolean if a field has been set.
func (o *SearchEventAttributesSSO) HasContextId() bool {
	if o != nil && !isNil(o.ContextId) {
		return true
	}

	return false
}

// SetContextId gets a reference to the given string and assigns it to the ContextId field.
func (o *SearchEventAttributesSSO) SetContextId(v string) {
	o.ContextId = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *SearchEventAttributesSSO) GetIpAddress() string {
	if o == nil || isNil(o.IpAddress) {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchEventAttributesSSO) GetIpAddressOk() (*string, bool) {
	if o == nil || isNil(o.IpAddress) {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *SearchEventAttributesSSO) HasIpAddress() bool {
	if o != nil && !isNil(o.IpAddress) {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *SearchEventAttributesSSO) SetIpAddress(v string) {
	o.IpAddress = &v
}

func (o SearchEventAttributesSSO) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ContextId) {
		toSerialize["contextId"] = o.ContextId
	}
	if !isNil(o.IpAddress) {
		toSerialize["ipAddress"] = o.IpAddress
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SearchEventAttributesSSO) UnmarshalJSON(bytes []byte) (err error) {
	varSearchEventAttributesSSO := _SearchEventAttributesSSO{}

	if err = json.Unmarshal(bytes, &varSearchEventAttributesSSO); err == nil {
		*o = SearchEventAttributesSSO(varSearchEventAttributesSSO)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "contextId")
		delete(additionalProperties, "ipAddress")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSearchEventAttributesSSO struct {
	value *SearchEventAttributesSSO
	isSet bool
}

func (v NullableSearchEventAttributesSSO) Get() *SearchEventAttributesSSO {
	return v.value
}

func (v *NullableSearchEventAttributesSSO) Set(val *SearchEventAttributesSSO) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchEventAttributesSSO) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchEventAttributesSSO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchEventAttributesSSO(val *SearchEventAttributesSSO) *NullableSearchEventAttributesSSO {
	return &NullableSearchEventAttributesSSO{value: val, isSet: true}
}

func (v NullableSearchEventAttributesSSO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchEventAttributesSSO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


