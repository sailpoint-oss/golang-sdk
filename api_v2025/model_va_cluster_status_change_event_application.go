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

// checks if the VAClusterStatusChangeEventApplication type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VAClusterStatusChangeEventApplication{}

// VAClusterStatusChangeEventApplication Details about the `CLUSTER` or `SOURCE` that initiated this event.
type VAClusterStatusChangeEventApplication struct {
	// The GUID of the application
	Id string `json:"id"`
	// The name of the application
	Name string `json:"name"`
	// Custom map of attributes for a source.  This will only be populated if type is `SOURCE` and the source has a proxy.
	Attributes map[string]interface{} `json:"attributes"`
	AdditionalProperties map[string]interface{}
}

type _VAClusterStatusChangeEventApplication VAClusterStatusChangeEventApplication

// NewVAClusterStatusChangeEventApplication instantiates a new VAClusterStatusChangeEventApplication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVAClusterStatusChangeEventApplication(id string, name string, attributes map[string]interface{}) *VAClusterStatusChangeEventApplication {
	this := VAClusterStatusChangeEventApplication{}
	this.Id = id
	this.Name = name
	this.Attributes = attributes
	return &this
}

// NewVAClusterStatusChangeEventApplicationWithDefaults instantiates a new VAClusterStatusChangeEventApplication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVAClusterStatusChangeEventApplicationWithDefaults() *VAClusterStatusChangeEventApplication {
	this := VAClusterStatusChangeEventApplication{}
	return &this
}

// GetId returns the Id field value
func (o *VAClusterStatusChangeEventApplication) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VAClusterStatusChangeEventApplication) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *VAClusterStatusChangeEventApplication) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *VAClusterStatusChangeEventApplication) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VAClusterStatusChangeEventApplication) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VAClusterStatusChangeEventApplication) SetName(v string) {
	o.Name = v
}

// GetAttributes returns the Attributes field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *VAClusterStatusChangeEventApplication) GetAttributes() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VAClusterStatusChangeEventApplication) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Attributes) {
		return map[string]interface{}{}, false
	}
	return o.Attributes, true
}

// SetAttributes sets field value
func (o *VAClusterStatusChangeEventApplication) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

func (o VAClusterStatusChangeEventApplication) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VAClusterStatusChangeEventApplication) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VAClusterStatusChangeEventApplication) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"attributes",
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

	varVAClusterStatusChangeEventApplication := _VAClusterStatusChangeEventApplication{}

	err = json.Unmarshal(data, &varVAClusterStatusChangeEventApplication)

	if err != nil {
		return err
	}

	*o = VAClusterStatusChangeEventApplication(varVAClusterStatusChangeEventApplication)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVAClusterStatusChangeEventApplication struct {
	value *VAClusterStatusChangeEventApplication
	isSet bool
}

func (v NullableVAClusterStatusChangeEventApplication) Get() *VAClusterStatusChangeEventApplication {
	return v.value
}

func (v *NullableVAClusterStatusChangeEventApplication) Set(val *VAClusterStatusChangeEventApplication) {
	v.value = val
	v.isSet = true
}

func (v NullableVAClusterStatusChangeEventApplication) IsSet() bool {
	return v.isSet
}

func (v *NullableVAClusterStatusChangeEventApplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVAClusterStatusChangeEventApplication(val *VAClusterStatusChangeEventApplication) *NullableVAClusterStatusChangeEventApplication {
	return &NullableVAClusterStatusChangeEventApplication{value: val, isSet: true}
}

func (v NullableVAClusterStatusChangeEventApplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVAClusterStatusChangeEventApplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


