/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

import (
	"encoding/json"
)

// checks if the CampaignFilterDetailsCriteriaListInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignFilterDetailsCriteriaListInner{}

// CampaignFilterDetailsCriteriaListInner struct for CampaignFilterDetailsCriteriaListInner
type CampaignFilterDetailsCriteriaListInner struct {
	Type CriteriaType `json:"type"`
	Operation Operation `json:"operation"`
	// The specified key from the Type of criteria.
	Property string `json:"property"`
	// The value for the specified key from the Type of Criteria
	Value string `json:"value"`
	AdditionalProperties map[string]interface{}
}

type _CampaignFilterDetailsCriteriaListInner CampaignFilterDetailsCriteriaListInner

// NewCampaignFilterDetailsCriteriaListInner instantiates a new CampaignFilterDetailsCriteriaListInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignFilterDetailsCriteriaListInner(type_ CriteriaType, operation Operation, property string, value string) *CampaignFilterDetailsCriteriaListInner {
	this := CampaignFilterDetailsCriteriaListInner{}
	this.Type = type_
	this.Operation = operation
	this.Property = property
	this.Value = value
	return &this
}

// NewCampaignFilterDetailsCriteriaListInnerWithDefaults instantiates a new CampaignFilterDetailsCriteriaListInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignFilterDetailsCriteriaListInnerWithDefaults() *CampaignFilterDetailsCriteriaListInner {
	this := CampaignFilterDetailsCriteriaListInner{}
	return &this
}

// GetType returns the Type field value
func (o *CampaignFilterDetailsCriteriaListInner) GetType() CriteriaType {
	if o == nil {
		var ret CriteriaType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CampaignFilterDetailsCriteriaListInner) GetTypeOk() (*CriteriaType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CampaignFilterDetailsCriteriaListInner) SetType(v CriteriaType) {
	o.Type = v
}

// GetOperation returns the Operation field value
func (o *CampaignFilterDetailsCriteriaListInner) GetOperation() Operation {
	if o == nil {
		var ret Operation
		return ret
	}

	return o.Operation
}

// GetOperationOk returns a tuple with the Operation field value
// and a boolean to check if the value has been set.
func (o *CampaignFilterDetailsCriteriaListInner) GetOperationOk() (*Operation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operation, true
}

// SetOperation sets field value
func (o *CampaignFilterDetailsCriteriaListInner) SetOperation(v Operation) {
	o.Operation = v
}

// GetProperty returns the Property field value
func (o *CampaignFilterDetailsCriteriaListInner) GetProperty() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Property
}

// GetPropertyOk returns a tuple with the Property field value
// and a boolean to check if the value has been set.
func (o *CampaignFilterDetailsCriteriaListInner) GetPropertyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Property, true
}

// SetProperty sets field value
func (o *CampaignFilterDetailsCriteriaListInner) SetProperty(v string) {
	o.Property = v
}

// GetValue returns the Value field value
func (o *CampaignFilterDetailsCriteriaListInner) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *CampaignFilterDetailsCriteriaListInner) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *CampaignFilterDetailsCriteriaListInner) SetValue(v string) {
	o.Value = v
}

func (o CampaignFilterDetailsCriteriaListInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CampaignFilterDetailsCriteriaListInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["operation"] = o.Operation
	toSerialize["property"] = o.Property
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CampaignFilterDetailsCriteriaListInner) UnmarshalJSON(bytes []byte) (err error) {
	varCampaignFilterDetailsCriteriaListInner := _CampaignFilterDetailsCriteriaListInner{}

	if err = json.Unmarshal(bytes, &varCampaignFilterDetailsCriteriaListInner); err == nil {
		*o = CampaignFilterDetailsCriteriaListInner(varCampaignFilterDetailsCriteriaListInner)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "operation")
		delete(additionalProperties, "property")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCampaignFilterDetailsCriteriaListInner struct {
	value *CampaignFilterDetailsCriteriaListInner
	isSet bool
}

func (v NullableCampaignFilterDetailsCriteriaListInner) Get() *CampaignFilterDetailsCriteriaListInner {
	return v.value
}

func (v *NullableCampaignFilterDetailsCriteriaListInner) Set(val *CampaignFilterDetailsCriteriaListInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignFilterDetailsCriteriaListInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignFilterDetailsCriteriaListInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignFilterDetailsCriteriaListInner(val *CampaignFilterDetailsCriteriaListInner) *NullableCampaignFilterDetailsCriteriaListInner {
	return &NullableCampaignFilterDetailsCriteriaListInner{value: val, isSet: true}
}

func (v NullableCampaignFilterDetailsCriteriaListInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignFilterDetailsCriteriaListInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

