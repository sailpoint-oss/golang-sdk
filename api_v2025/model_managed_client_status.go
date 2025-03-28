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

// checks if the ManagedClientStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagedClientStatus{}

// ManagedClientStatus Managed Client Status
type ManagedClientStatus struct {
	// ManagedClientStatus body information
	Body map[string]interface{} `json:"body"`
	Status ManagedClientStatusCode `json:"status"`
	Type NullableManagedClientType `json:"type"`
	// timestamp on the Client Status update
	Timestamp SailPointTime `json:"timestamp"`
	AdditionalProperties map[string]interface{}
}

type _ManagedClientStatus ManagedClientStatus

// NewManagedClientStatus instantiates a new ManagedClientStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagedClientStatus(body map[string]interface{}, status ManagedClientStatusCode, type_ NullableManagedClientType, timestamp SailPointTime) *ManagedClientStatus {
	this := ManagedClientStatus{}
	this.Body = body
	this.Status = status
	this.Type = type_
	this.Timestamp = timestamp
	return &this
}

// NewManagedClientStatusWithDefaults instantiates a new ManagedClientStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagedClientStatusWithDefaults() *ManagedClientStatus {
	this := ManagedClientStatus{}
	return &this
}

// GetBody returns the Body field value
func (o *ManagedClientStatus) GetBody() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *ManagedClientStatus) GetBodyOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Body, true
}

// SetBody sets field value
func (o *ManagedClientStatus) SetBody(v map[string]interface{}) {
	o.Body = v
}

// GetStatus returns the Status field value
func (o *ManagedClientStatus) GetStatus() ManagedClientStatusCode {
	if o == nil {
		var ret ManagedClientStatusCode
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ManagedClientStatus) GetStatusOk() (*ManagedClientStatusCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ManagedClientStatus) SetStatus(v ManagedClientStatusCode) {
	o.Status = v
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for ManagedClientType will be returned
func (o *ManagedClientStatus) GetType() ManagedClientType {
	if o == nil || o.Type.Get() == nil {
		var ret ManagedClientType
		return ret
	}

	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManagedClientStatus) GetTypeOk() (*ManagedClientType, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// SetType sets field value
func (o *ManagedClientStatus) SetType(v ManagedClientType) {
	o.Type.Set(&v)
}

// GetTimestamp returns the Timestamp field value
func (o *ManagedClientStatus) GetTimestamp() SailPointTime {
	if o == nil {
		var ret SailPointTime
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *ManagedClientStatus) GetTimestampOk() (*SailPointTime, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *ManagedClientStatus) SetTimestamp(v SailPointTime) {
	o.Timestamp = v
}

func (o ManagedClientStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagedClientStatus) ToMap() (map[string]interface{}, error) {
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

func (o *ManagedClientStatus) UnmarshalJSON(data []byte) (err error) {
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

	varManagedClientStatus := _ManagedClientStatus{}

	err = json.Unmarshal(data, &varManagedClientStatus)

	if err != nil {
		return err
	}

	*o = ManagedClientStatus(varManagedClientStatus)

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

type NullableManagedClientStatus struct {
	value *ManagedClientStatus
	isSet bool
}

func (v NullableManagedClientStatus) Get() *ManagedClientStatus {
	return v.value
}

func (v *NullableManagedClientStatus) Set(val *ManagedClientStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableManagedClientStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableManagedClientStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagedClientStatus(val *ManagedClientStatus) *NullableManagedClientStatus {
	return &NullableManagedClientStatus{value: val, isSet: true}
}

func (v NullableManagedClientStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagedClientStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


