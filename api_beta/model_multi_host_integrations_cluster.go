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

// checks if the MultiHostIntegrationsCluster type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MultiHostIntegrationsCluster{}

// MultiHostIntegrationsCluster Reference to the source's associated cluster.
type MultiHostIntegrationsCluster struct {
	// Type of object being referenced.
	Type string `json:"type"`
	// Cluster ID.
	Id string `json:"id"`
	// Cluster's human-readable display name.
	Name string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _MultiHostIntegrationsCluster MultiHostIntegrationsCluster

// NewMultiHostIntegrationsCluster instantiates a new MultiHostIntegrationsCluster object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultiHostIntegrationsCluster(type_ string, id string, name string) *MultiHostIntegrationsCluster {
	this := MultiHostIntegrationsCluster{}
	this.Type = type_
	this.Id = id
	this.Name = name
	return &this
}

// NewMultiHostIntegrationsClusterWithDefaults instantiates a new MultiHostIntegrationsCluster object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultiHostIntegrationsClusterWithDefaults() *MultiHostIntegrationsCluster {
	this := MultiHostIntegrationsCluster{}
	return &this
}

// GetType returns the Type field value
func (o *MultiHostIntegrationsCluster) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MultiHostIntegrationsCluster) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MultiHostIntegrationsCluster) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *MultiHostIntegrationsCluster) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MultiHostIntegrationsCluster) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MultiHostIntegrationsCluster) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *MultiHostIntegrationsCluster) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MultiHostIntegrationsCluster) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MultiHostIntegrationsCluster) SetName(v string) {
	o.Name = v
}

func (o MultiHostIntegrationsCluster) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MultiHostIntegrationsCluster) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MultiHostIntegrationsCluster) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"id",
		"name",
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

	varMultiHostIntegrationsCluster := _MultiHostIntegrationsCluster{}

	err = json.Unmarshal(data, &varMultiHostIntegrationsCluster)

	if err != nil {
		return err
	}

	*o = MultiHostIntegrationsCluster(varMultiHostIntegrationsCluster)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMultiHostIntegrationsCluster struct {
	value *MultiHostIntegrationsCluster
	isSet bool
}

func (v NullableMultiHostIntegrationsCluster) Get() *MultiHostIntegrationsCluster {
	return v.value
}

func (v *NullableMultiHostIntegrationsCluster) Set(val *MultiHostIntegrationsCluster) {
	v.value = val
	v.isSet = true
}

func (v NullableMultiHostIntegrationsCluster) IsSet() bool {
	return v.isSet
}

func (v *NullableMultiHostIntegrationsCluster) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultiHostIntegrationsCluster(val *MultiHostIntegrationsCluster) *NullableMultiHostIntegrationsCluster {
	return &NullableMultiHostIntegrationsCluster{value: val, isSet: true}
}

func (v NullableMultiHostIntegrationsCluster) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultiHostIntegrationsCluster) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


