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

// checks if the ManagedClientStatusAggResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagedClientStatusAggResponse{}

// ManagedClientStatusAggResponse Managed Client Status
type ManagedClientStatusAggResponse struct {
	// ManagedClientStatus body information
	Body map[string]interface{} `json:"body"`
	Status ManagedClientStatusEnum `json:"status"`
	Type NullableManagedClientType `json:"type"`
	// timestamp on the Client Status update
	Timestamp SailPointTime `json:"timestamp"`
	AdditionalProperties map[string]interface{}
}

type _ManagedClientStatusAggResponse ManagedClientStatusAggResponse

// NewManagedClientStatusAggResponse instantiates a new ManagedClientStatusAggResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagedClientStatusAggResponse(body map[string]interface{}, status ManagedClientStatusEnum, type_ NullableManagedClientType, timestamp SailPointTime) *ManagedClientStatusAggResponse {
	this := ManagedClientStatusAggResponse{}
	this.Body = body
	this.Status = status
	this.Type = type_
	this.Timestamp = timestamp
	return &this
}

// NewManagedClientStatusAggResponseWithDefaults instantiates a new ManagedClientStatusAggResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagedClientStatusAggResponseWithDefaults() *ManagedClientStatusAggResponse {
	this := ManagedClientStatusAggResponse{}
	return &this
}

// GetBody returns the Body field value
func (o *ManagedClientStatusAggResponse) GetBody() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *ManagedClientStatusAggResponse) GetBodyOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Body, true
}

// SetBody sets field value
func (o *ManagedClientStatusAggResponse) SetBody(v map[string]interface{}) {
	o.Body = v
}

// GetStatus returns the Status field value
func (o *ManagedClientStatusAggResponse) GetStatus() ManagedClientStatusEnum {
	if o == nil {
		var ret ManagedClientStatusEnum
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ManagedClientStatusAggResponse) GetStatusOk() (*ManagedClientStatusEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ManagedClientStatusAggResponse) SetStatus(v ManagedClientStatusEnum) {
	o.Status = v
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for ManagedClientType will be returned
func (o *ManagedClientStatusAggResponse) GetType() ManagedClientType {
	if o == nil || o.Type.Get() == nil {
		var ret ManagedClientType
		return ret
	}

	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManagedClientStatusAggResponse) GetTypeOk() (*ManagedClientType, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// SetType sets field value
func (o *ManagedClientStatusAggResponse) SetType(v ManagedClientType) {
	o.Type.Set(&v)
}

// GetTimestamp returns the Timestamp field value
func (o *ManagedClientStatusAggResponse) GetTimestamp() SailPointTime {
	if o == nil {
		var ret SailPointTime
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *ManagedClientStatusAggResponse) GetTimestampOk() (*SailPointTime, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *ManagedClientStatusAggResponse) SetTimestamp(v SailPointTime) {
	o.Timestamp = v
}

func (o ManagedClientStatusAggResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagedClientStatusAggResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["body"] = o.Body
	toSerialize["status"] = o.Status
	toSerialize["type"] = o.Type.Get()
	toSerialize["timestamp"] = o.Timestamp

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ManagedClientStatusAggResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"body",
		"status",
		"type",
		"timestamp",
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

	varManagedClientStatusAggResponse := _ManagedClientStatusAggResponse{}

	err = json.Unmarshal(data, &varManagedClientStatusAggResponse)

	if err != nil {
		return err
	}

	*o = ManagedClientStatusAggResponse(varManagedClientStatusAggResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "body")
		delete(additionalProperties, "status")
		delete(additionalProperties, "type")
		delete(additionalProperties, "timestamp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableManagedClientStatusAggResponse struct {
	value *ManagedClientStatusAggResponse
	isSet bool
}

func (v NullableManagedClientStatusAggResponse) Get() *ManagedClientStatusAggResponse {
	return v.value
}

func (v *NullableManagedClientStatusAggResponse) Set(val *ManagedClientStatusAggResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableManagedClientStatusAggResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableManagedClientStatusAggResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagedClientStatusAggResponse(val *ManagedClientStatusAggResponse) *NullableManagedClientStatusAggResponse {
	return &NullableManagedClientStatusAggResponse{value: val, isSet: true}
}

func (v NullableManagedClientStatusAggResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagedClientStatusAggResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


