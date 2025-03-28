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

// checks if the ClientLogConfigurationExpiration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientLogConfigurationExpiration{}

// ClientLogConfigurationExpiration Client Runtime Logging Configuration
type ClientLogConfigurationExpiration struct {
	// Log configuration's client ID
	ClientId *string `json:"clientId,omitempty"`
	// Expiration date-time of the log configuration request.  Can be no greater than 24 hours from current date-time.
	Expiration *SailPointTime `json:"expiration,omitempty"`
	RootLevel StandardLevel `json:"rootLevel"`
	// Mapping of identifiers to Standard Log Level values
	LogLevels *map[string]StandardLevel `json:"logLevels,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ClientLogConfigurationExpiration ClientLogConfigurationExpiration

// NewClientLogConfigurationExpiration instantiates a new ClientLogConfigurationExpiration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientLogConfigurationExpiration(rootLevel StandardLevel) *ClientLogConfigurationExpiration {
	this := ClientLogConfigurationExpiration{}
	this.RootLevel = rootLevel
	return &this
}

// NewClientLogConfigurationExpirationWithDefaults instantiates a new ClientLogConfigurationExpiration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientLogConfigurationExpirationWithDefaults() *ClientLogConfigurationExpiration {
	this := ClientLogConfigurationExpiration{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *ClientLogConfigurationExpiration) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientLogConfigurationExpiration) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *ClientLogConfigurationExpiration) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *ClientLogConfigurationExpiration) SetClientId(v string) {
	o.ClientId = &v
}

// GetExpiration returns the Expiration field value if set, zero value otherwise.
func (o *ClientLogConfigurationExpiration) GetExpiration() SailPointTime {
	if o == nil || IsNil(o.Expiration) {
		var ret SailPointTime
		return ret
	}
	return *o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientLogConfigurationExpiration) GetExpirationOk() (*SailPointTime, bool) {
	if o == nil || IsNil(o.Expiration) {
		return nil, false
	}
	return o.Expiration, true
}

// HasExpiration returns a boolean if a field has been set.
func (o *ClientLogConfigurationExpiration) HasExpiration() bool {
	if o != nil && !IsNil(o.Expiration) {
		return true
	}

	return false
}

// SetExpiration gets a reference to the given SailPointTime and assigns it to the Expiration field.
func (o *ClientLogConfigurationExpiration) SetExpiration(v SailPointTime) {
	o.Expiration = &v
}

// GetRootLevel returns the RootLevel field value
func (o *ClientLogConfigurationExpiration) GetRootLevel() StandardLevel {
	if o == nil {
		var ret StandardLevel
		return ret
	}

	return o.RootLevel
}

// GetRootLevelOk returns a tuple with the RootLevel field value
// and a boolean to check if the value has been set.
func (o *ClientLogConfigurationExpiration) GetRootLevelOk() (*StandardLevel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RootLevel, true
}

// SetRootLevel sets field value
func (o *ClientLogConfigurationExpiration) SetRootLevel(v StandardLevel) {
	o.RootLevel = v
}

// GetLogLevels returns the LogLevels field value if set, zero value otherwise.
func (o *ClientLogConfigurationExpiration) GetLogLevels() map[string]StandardLevel {
	if o == nil || IsNil(o.LogLevels) {
		var ret map[string]StandardLevel
		return ret
	}
	return *o.LogLevels
}

// GetLogLevelsOk returns a tuple with the LogLevels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientLogConfigurationExpiration) GetLogLevelsOk() (*map[string]StandardLevel, bool) {
	if o == nil || IsNil(o.LogLevels) {
		return nil, false
	}
	return o.LogLevels, true
}

// HasLogLevels returns a boolean if a field has been set.
func (o *ClientLogConfigurationExpiration) HasLogLevels() bool {
	if o != nil && !IsNil(o.LogLevels) {
		return true
	}

	return false
}

// SetLogLevels gets a reference to the given map[string]StandardLevel and assigns it to the LogLevels field.
func (o *ClientLogConfigurationExpiration) SetLogLevels(v map[string]StandardLevel) {
	o.LogLevels = &v
}

func (o ClientLogConfigurationExpiration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientLogConfigurationExpiration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClientId) {
		toSerialize["clientId"] = o.ClientId
	}
	if !IsNil(o.Expiration) {
		toSerialize["expiration"] = o.Expiration
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

func (o *ClientLogConfigurationExpiration) UnmarshalJSON(data []byte) (err error) {
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

	varClientLogConfigurationExpiration := _ClientLogConfigurationExpiration{}

	err = json.Unmarshal(data, &varClientLogConfigurationExpiration)

	if err != nil {
		return err
	}

	*o = ClientLogConfigurationExpiration(varClientLogConfigurationExpiration)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "clientId")
		delete(additionalProperties, "expiration")
		delete(additionalProperties, "rootLevel")
		delete(additionalProperties, "logLevels")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableClientLogConfigurationExpiration struct {
	value *ClientLogConfigurationExpiration
	isSet bool
}

func (v NullableClientLogConfigurationExpiration) Get() *ClientLogConfigurationExpiration {
	return v.value
}

func (v *NullableClientLogConfigurationExpiration) Set(val *ClientLogConfigurationExpiration) {
	v.value = val
	v.isSet = true
}

func (v NullableClientLogConfigurationExpiration) IsSet() bool {
	return v.isSet
}

func (v *NullableClientLogConfigurationExpiration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientLogConfigurationExpiration(val *ClientLogConfigurationExpiration) *NullableClientLogConfigurationExpiration {
	return &NullableClientLogConfigurationExpiration{value: val, isSet: true}
}

func (v NullableClientLogConfigurationExpiration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientLogConfigurationExpiration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


