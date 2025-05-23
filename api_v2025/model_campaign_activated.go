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

// checks if the CampaignActivated type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignActivated{}

// CampaignActivated struct for CampaignActivated
type CampaignActivated struct {
	Campaign CampaignActivatedCampaign `json:"campaign"`
	AdditionalProperties map[string]interface{}
}

type _CampaignActivated CampaignActivated

// NewCampaignActivated instantiates a new CampaignActivated object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignActivated(campaign CampaignActivatedCampaign) *CampaignActivated {
	this := CampaignActivated{}
	this.Campaign = campaign
	return &this
}

// NewCampaignActivatedWithDefaults instantiates a new CampaignActivated object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignActivatedWithDefaults() *CampaignActivated {
	this := CampaignActivated{}
	return &this
}

// GetCampaign returns the Campaign field value
func (o *CampaignActivated) GetCampaign() CampaignActivatedCampaign {
	if o == nil {
		var ret CampaignActivatedCampaign
		return ret
	}

	return o.Campaign
}

// GetCampaignOk returns a tuple with the Campaign field value
// and a boolean to check if the value has been set.
func (o *CampaignActivated) GetCampaignOk() (*CampaignActivatedCampaign, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Campaign, true
}

// SetCampaign sets field value
func (o *CampaignActivated) SetCampaign(v CampaignActivatedCampaign) {
	o.Campaign = v
}

func (o CampaignActivated) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CampaignActivated) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaign"] = o.Campaign

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CampaignActivated) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"campaign",
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

	varCampaignActivated := _CampaignActivated{}

	err = json.Unmarshal(data, &varCampaignActivated)

	if err != nil {
		return err
	}

	*o = CampaignActivated(varCampaignActivated)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "campaign")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCampaignActivated struct {
	value *CampaignActivated
	isSet bool
}

func (v NullableCampaignActivated) Get() *CampaignActivated {
	return v.value
}

func (v *NullableCampaignActivated) Set(val *CampaignActivated) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignActivated) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignActivated) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignActivated(val *CampaignActivated) *NullableCampaignActivated {
	return &NullableCampaignActivated{value: val, isSet: true}
}

func (v NullableCampaignActivated) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignActivated) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


