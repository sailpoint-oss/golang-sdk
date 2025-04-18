/*
Identity Security Cloud V3 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v3

import (
	"encoding/json"
	"fmt"
)

// checks if the SendTokenRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SendTokenRequest{}

// SendTokenRequest struct for SendTokenRequest
type SendTokenRequest struct {
	// User alias from table spt_identity field named 'name'
	UserAlias string `json:"userAlias"`
	// Token delivery type
	DeliveryType string `json:"deliveryType"`
	AdditionalProperties map[string]interface{}
}

type _SendTokenRequest SendTokenRequest

// NewSendTokenRequest instantiates a new SendTokenRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSendTokenRequest(userAlias string, deliveryType string) *SendTokenRequest {
	this := SendTokenRequest{}
	this.UserAlias = userAlias
	this.DeliveryType = deliveryType
	return &this
}

// NewSendTokenRequestWithDefaults instantiates a new SendTokenRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSendTokenRequestWithDefaults() *SendTokenRequest {
	this := SendTokenRequest{}
	return &this
}

// GetUserAlias returns the UserAlias field value
func (o *SendTokenRequest) GetUserAlias() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserAlias
}

// GetUserAliasOk returns a tuple with the UserAlias field value
// and a boolean to check if the value has been set.
func (o *SendTokenRequest) GetUserAliasOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserAlias, true
}

// SetUserAlias sets field value
func (o *SendTokenRequest) SetUserAlias(v string) {
	o.UserAlias = v
}

// GetDeliveryType returns the DeliveryType field value
func (o *SendTokenRequest) GetDeliveryType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeliveryType
}

// GetDeliveryTypeOk returns a tuple with the DeliveryType field value
// and a boolean to check if the value has been set.
func (o *SendTokenRequest) GetDeliveryTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeliveryType, true
}

// SetDeliveryType sets field value
func (o *SendTokenRequest) SetDeliveryType(v string) {
	o.DeliveryType = v
}

func (o SendTokenRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SendTokenRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["userAlias"] = o.UserAlias
	toSerialize["deliveryType"] = o.DeliveryType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SendTokenRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"userAlias",
		"deliveryType",
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

	varSendTokenRequest := _SendTokenRequest{}

	err = json.Unmarshal(data, &varSendTokenRequest)

	if err != nil {
		return err
	}

	*o = SendTokenRequest(varSendTokenRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "userAlias")
		delete(additionalProperties, "deliveryType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSendTokenRequest struct {
	value *SendTokenRequest
	isSet bool
}

func (v NullableSendTokenRequest) Get() *SendTokenRequest {
	return v.value
}

func (v *NullableSendTokenRequest) Set(val *SendTokenRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSendTokenRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSendTokenRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSendTokenRequest(val *SendTokenRequest) *NullableSendTokenRequest {
	return &NullableSendTokenRequest{value: val, isSet: true}
}

func (v NullableSendTokenRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSendTokenRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


