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

// checks if the CampaignActivatedCampaignCampaignOwner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignActivatedCampaignCampaignOwner{}

// CampaignActivatedCampaignCampaignOwner Details of the identity that owns the campaign.
type CampaignActivatedCampaignCampaignOwner struct {
	// The unique ID of the identity.
	Id string `json:"id"`
	// The human friendly name of the identity.
	DisplayName string `json:"displayName"`
	// The primary email address of the identity.
	Email string `json:"email"`
	AdditionalProperties map[string]interface{}
}

type _CampaignActivatedCampaignCampaignOwner CampaignActivatedCampaignCampaignOwner

// NewCampaignActivatedCampaignCampaignOwner instantiates a new CampaignActivatedCampaignCampaignOwner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignActivatedCampaignCampaignOwner(id string, displayName string, email string) *CampaignActivatedCampaignCampaignOwner {
	this := CampaignActivatedCampaignCampaignOwner{}
	this.Id = id
	this.DisplayName = displayName
	this.Email = email
	return &this
}

// NewCampaignActivatedCampaignCampaignOwnerWithDefaults instantiates a new CampaignActivatedCampaignCampaignOwner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignActivatedCampaignCampaignOwnerWithDefaults() *CampaignActivatedCampaignCampaignOwner {
	this := CampaignActivatedCampaignCampaignOwner{}
	return &this
}

// GetId returns the Id field value
func (o *CampaignActivatedCampaignCampaignOwner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CampaignActivatedCampaignCampaignOwner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CampaignActivatedCampaignCampaignOwner) SetId(v string) {
	o.Id = v
}

// GetDisplayName returns the DisplayName field value
func (o *CampaignActivatedCampaignCampaignOwner) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *CampaignActivatedCampaignCampaignOwner) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *CampaignActivatedCampaignCampaignOwner) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetEmail returns the Email field value
func (o *CampaignActivatedCampaignCampaignOwner) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *CampaignActivatedCampaignCampaignOwner) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *CampaignActivatedCampaignCampaignOwner) SetEmail(v string) {
	o.Email = v
}

func (o CampaignActivatedCampaignCampaignOwner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CampaignActivatedCampaignCampaignOwner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["displayName"] = o.DisplayName
	toSerialize["email"] = o.Email

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CampaignActivatedCampaignCampaignOwner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"displayName",
		"email",
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

	varCampaignActivatedCampaignCampaignOwner := _CampaignActivatedCampaignCampaignOwner{}

	err = json.Unmarshal(data, &varCampaignActivatedCampaignCampaignOwner)

	if err != nil {
		return err
	}

	*o = CampaignActivatedCampaignCampaignOwner(varCampaignActivatedCampaignCampaignOwner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "email")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCampaignActivatedCampaignCampaignOwner struct {
	value *CampaignActivatedCampaignCampaignOwner
	isSet bool
}

func (v NullableCampaignActivatedCampaignCampaignOwner) Get() *CampaignActivatedCampaignCampaignOwner {
	return v.value
}

func (v *NullableCampaignActivatedCampaignCampaignOwner) Set(val *CampaignActivatedCampaignCampaignOwner) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignActivatedCampaignCampaignOwner) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignActivatedCampaignCampaignOwner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignActivatedCampaignCampaignOwner(val *CampaignActivatedCampaignCampaignOwner) *NullableCampaignActivatedCampaignCampaignOwner {
	return &NullableCampaignActivatedCampaignCampaignOwner{value: val, isSet: true}
}

func (v NullableCampaignActivatedCampaignCampaignOwner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignActivatedCampaignCampaignOwner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


