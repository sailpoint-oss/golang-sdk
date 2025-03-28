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

// checks if the SourceBeforeProvisioningRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SourceBeforeProvisioningRule{}

// SourceBeforeProvisioningRule Rule that runs on the CCG and allows for customization of provisioning plans before the API calls the connector. 
type SourceBeforeProvisioningRule struct {
	// Type of object being referenced.
	Type *string `json:"type,omitempty"`
	// Rule ID.
	Id *string `json:"id,omitempty"`
	// Rule's human-readable display name.
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SourceBeforeProvisioningRule SourceBeforeProvisioningRule

// NewSourceBeforeProvisioningRule instantiates a new SourceBeforeProvisioningRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceBeforeProvisioningRule() *SourceBeforeProvisioningRule {
	this := SourceBeforeProvisioningRule{}
	return &this
}

// NewSourceBeforeProvisioningRuleWithDefaults instantiates a new SourceBeforeProvisioningRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceBeforeProvisioningRuleWithDefaults() *SourceBeforeProvisioningRule {
	this := SourceBeforeProvisioningRule{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SourceBeforeProvisioningRule) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceBeforeProvisioningRule) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SourceBeforeProvisioningRule) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SourceBeforeProvisioningRule) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SourceBeforeProvisioningRule) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceBeforeProvisioningRule) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SourceBeforeProvisioningRule) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SourceBeforeProvisioningRule) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SourceBeforeProvisioningRule) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceBeforeProvisioningRule) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SourceBeforeProvisioningRule) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SourceBeforeProvisioningRule) SetName(v string) {
	o.Name = &v
}

func (o SourceBeforeProvisioningRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SourceBeforeProvisioningRule) ToMap() (map[string]interface{}, error) {
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

func (o *SourceBeforeProvisioningRule) UnmarshalJSON(data []byte) (err error) {
	varSourceBeforeProvisioningRule := _SourceBeforeProvisioningRule{}

	err = json.Unmarshal(data, &varSourceBeforeProvisioningRule)

	if err != nil {
		return err
	}

	*o = SourceBeforeProvisioningRule(varSourceBeforeProvisioningRule)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSourceBeforeProvisioningRule struct {
	value *SourceBeforeProvisioningRule
	isSet bool
}

func (v NullableSourceBeforeProvisioningRule) Get() *SourceBeforeProvisioningRule {
	return v.value
}

func (v *NullableSourceBeforeProvisioningRule) Set(val *SourceBeforeProvisioningRule) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceBeforeProvisioningRule) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceBeforeProvisioningRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceBeforeProvisioningRule(val *SourceBeforeProvisioningRule) *NullableSourceBeforeProvisioningRule {
	return &NullableSourceBeforeProvisioningRule{value: val, isSet: true}
}

func (v NullableSourceBeforeProvisioningRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceBeforeProvisioningRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


