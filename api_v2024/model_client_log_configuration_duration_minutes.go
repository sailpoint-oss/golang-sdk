/*
Identity Security Cloud V2024 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: v2024
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2024

import (
	"encoding/json"
	"fmt"
)

// checks if the ClientLogConfigurationDurationMinutes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientLogConfigurationDurationMinutes{}

// ClientLogConfigurationDurationMinutes Client Runtime Logging Configuration
type ClientLogConfigurationDurationMinutes struct {
	// Log configuration's client ID
	ClientId *string `json:"clientId,omitempty"`
	// Duration in minutes for log configuration to remain in effect before resetting to defaults.
	DurationMinutes *int32 `json:"durationMinutes,omitempty"`
	RootLevel StandardLevel `json:"rootLevel"`
	// Mapping of identifiers to Standard Log Level values
	LogLevels *map[string]StandardLevel `json:"logLevels,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ClientLogConfigurationDurationMinutes ClientLogConfigurationDurationMinutes

// NewClientLogConfigurationDurationMinutes instantiates a new ClientLogConfigurationDurationMinutes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientLogConfigurationDurationMinutes(rootLevel StandardLevel) *ClientLogConfigurationDurationMinutes {
	this := ClientLogConfigurationDurationMinutes{}
	var durationMinutes int32 = 240
	this.DurationMinutes = &durationMinutes
	this.RootLevel = rootLevel
	return &this
}

// NewClientLogConfigurationDurationMinutesWithDefaults instantiates a new ClientLogConfigurationDurationMinutes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientLogConfigurationDurationMinutesWithDefaults() *ClientLogConfigurationDurationMinutes {
	this := ClientLogConfigurationDurationMinutes{}
	var durationMinutes int32 = 240
	this.DurationMinutes = &durationMinutes
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *ClientLogConfigurationDurationMinutes) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientLogConfigurationDurationMinutes) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *ClientLogConfigurationDurationMinutes) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *ClientLogConfigurationDurationMinutes) SetClientId(v string) {
	o.ClientId = &v
}

// GetDurationMinutes returns the DurationMinutes field value if set, zero value otherwise.
func (o *ClientLogConfigurationDurationMinutes) GetDurationMinutes() int32 {
	if o == nil || IsNil(o.DurationMinutes) {
		var ret int32
		return ret
	}
	return *o.DurationMinutes
}

// GetDurationMinutesOk returns a tuple with the DurationMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientLogConfigurationDurationMinutes) GetDurationMinutesOk() (*int32, bool) {
	if o == nil || IsNil(o.DurationMinutes) {
		return nil, false
	}
	return o.DurationMinutes, true
}

// HasDurationMinutes returns a boolean if a field has been set.
func (o *ClientLogConfigurationDurationMinutes) HasDurationMinutes() bool {
	if o != nil && !IsNil(o.DurationMinutes) {
		return true
	}

	return false
}

// SetDurationMinutes gets a reference to the given int32 and assigns it to the DurationMinutes field.
func (o *ClientLogConfigurationDurationMinutes) SetDurationMinutes(v int32) {
	o.DurationMinutes = &v
}

// GetRootLevel returns the RootLevel field value
func (o *ClientLogConfigurationDurationMinutes) GetRootLevel() StandardLevel {
	if o == nil {
		var ret StandardLevel
		return ret
	}

	return o.RootLevel
}

// GetRootLevelOk returns a tuple with the RootLevel field value
// and a boolean to check if the value has been set.
func (o *ClientLogConfigurationDurationMinutes) GetRootLevelOk() (*StandardLevel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RootLevel, true
}

// SetRootLevel sets field value
func (o *ClientLogConfigurationDurationMinutes) SetRootLevel(v StandardLevel) {
	o.RootLevel = v
}

// GetLogLevels returns the LogLevels field value if set, zero value otherwise.
func (o *ClientLogConfigurationDurationMinutes) GetLogLevels() map[string]StandardLevel {
	if o == nil || IsNil(o.LogLevels) {
		var ret map[string]StandardLevel
		return ret
	}
	return *o.LogLevels
}

// GetLogLevelsOk returns a tuple with the LogLevels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientLogConfigurationDurationMinutes) GetLogLevelsOk() (*map[string]StandardLevel, bool) {
	if o == nil || IsNil(o.LogLevels) {
		return nil, false
	}
	return o.LogLevels, true
}

// HasLogLevels returns a boolean if a field has been set.
func (o *ClientLogConfigurationDurationMinutes) HasLogLevels() bool {
	if o != nil && !IsNil(o.LogLevels) {
		return true
	}

	return false
}

// SetLogLevels gets a reference to the given map[string]StandardLevel and assigns it to the LogLevels field.
func (o *ClientLogConfigurationDurationMinutes) SetLogLevels(v map[string]StandardLevel) {
	o.LogLevels = &v
}

func (o ClientLogConfigurationDurationMinutes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientLogConfigurationDurationMinutes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClientId) {
		toSerialize["clientId"] = o.ClientId
	}
	if !IsNil(o.DurationMinutes) {
		toSerialize["durationMinutes"] = o.DurationMinutes
	}
	toSerialize["rootLevel"] = o.RootLevel
	if !IsNil(o.LogLevels) {
		toSerialize["logLevels"] = o.LogLevels
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ClientLogConfigurationDurationMinutes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"rootLevel",
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

	varClientLogConfigurationDurationMinutes := _ClientLogConfigurationDurationMinutes{}

	err = json.Unmarshal(data, &varClientLogConfigurationDurationMinutes)

	if err != nil {
		return err
	}

	*o = ClientLogConfigurationDurationMinutes(varClientLogConfigurationDurationMinutes)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "clientId")
		delete(additionalProperties, "durationMinutes")
		delete(additionalProperties, "rootLevel")
		delete(additionalProperties, "logLevels")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableClientLogConfigurationDurationMinutes struct {
	value *ClientLogConfigurationDurationMinutes
	isSet bool
}

func (v NullableClientLogConfigurationDurationMinutes) Get() *ClientLogConfigurationDurationMinutes {
	return v.value
}

func (v *NullableClientLogConfigurationDurationMinutes) Set(val *ClientLogConfigurationDurationMinutes) {
	v.value = val
	v.isSet = true
}

func (v NullableClientLogConfigurationDurationMinutes) IsSet() bool {
	return v.isSet
}

func (v *NullableClientLogConfigurationDurationMinutes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientLogConfigurationDurationMinutes(val *ClientLogConfigurationDurationMinutes) *NullableClientLogConfigurationDurationMinutes {
	return &NullableClientLogConfigurationDurationMinutes{value: val, isSet: true}
}

func (v NullableClientLogConfigurationDurationMinutes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientLogConfigurationDurationMinutes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

