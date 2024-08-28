/*
Identity Security Cloud V3 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v3

import (
	"encoding/json"
)

// checks if the SessionConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SessionConfiguration{}

// SessionConfiguration struct for SessionConfiguration
type SessionConfiguration struct {
	// The maximum time in minutes a session can be idle.
	MaxIdleTime *int32 `json:"maxIdleTime,omitempty"`
	// Denotes if 'remember me' is enabled.
	RememberMe *bool `json:"rememberMe,omitempty"`
	// The maximum allowable session time in minutes.
	MaxSessionTime *int32 `json:"maxSessionTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SessionConfiguration SessionConfiguration

// NewSessionConfiguration instantiates a new SessionConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionConfiguration() *SessionConfiguration {
	this := SessionConfiguration{}
	var rememberMe bool = false
	this.RememberMe = &rememberMe
	return &this
}

// NewSessionConfigurationWithDefaults instantiates a new SessionConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionConfigurationWithDefaults() *SessionConfiguration {
	this := SessionConfiguration{}
	var rememberMe bool = false
	this.RememberMe = &rememberMe
	return &this
}

// GetMaxIdleTime returns the MaxIdleTime field value if set, zero value otherwise.
func (o *SessionConfiguration) GetMaxIdleTime() int32 {
	if o == nil || IsNil(o.MaxIdleTime) {
		var ret int32
		return ret
	}
	return *o.MaxIdleTime
}

// GetMaxIdleTimeOk returns a tuple with the MaxIdleTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionConfiguration) GetMaxIdleTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxIdleTime) {
		return nil, false
	}
	return o.MaxIdleTime, true
}

// HasMaxIdleTime returns a boolean if a field has been set.
func (o *SessionConfiguration) HasMaxIdleTime() bool {
	if o != nil && !IsNil(o.MaxIdleTime) {
		return true
	}

	return false
}

// SetMaxIdleTime gets a reference to the given int32 and assigns it to the MaxIdleTime field.
func (o *SessionConfiguration) SetMaxIdleTime(v int32) {
	o.MaxIdleTime = &v
}

// GetRememberMe returns the RememberMe field value if set, zero value otherwise.
func (o *SessionConfiguration) GetRememberMe() bool {
	if o == nil || IsNil(o.RememberMe) {
		var ret bool
		return ret
	}
	return *o.RememberMe
}

// GetRememberMeOk returns a tuple with the RememberMe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionConfiguration) GetRememberMeOk() (*bool, bool) {
	if o == nil || IsNil(o.RememberMe) {
		return nil, false
	}
	return o.RememberMe, true
}

// HasRememberMe returns a boolean if a field has been set.
func (o *SessionConfiguration) HasRememberMe() bool {
	if o != nil && !IsNil(o.RememberMe) {
		return true
	}

	return false
}

// SetRememberMe gets a reference to the given bool and assigns it to the RememberMe field.
func (o *SessionConfiguration) SetRememberMe(v bool) {
	o.RememberMe = &v
}

// GetMaxSessionTime returns the MaxSessionTime field value if set, zero value otherwise.
func (o *SessionConfiguration) GetMaxSessionTime() int32 {
	if o == nil || IsNil(o.MaxSessionTime) {
		var ret int32
		return ret
	}
	return *o.MaxSessionTime
}

// GetMaxSessionTimeOk returns a tuple with the MaxSessionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionConfiguration) GetMaxSessionTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxSessionTime) {
		return nil, false
	}
	return o.MaxSessionTime, true
}

// HasMaxSessionTime returns a boolean if a field has been set.
func (o *SessionConfiguration) HasMaxSessionTime() bool {
	if o != nil && !IsNil(o.MaxSessionTime) {
		return true
	}

	return false
}

// SetMaxSessionTime gets a reference to the given int32 and assigns it to the MaxSessionTime field.
func (o *SessionConfiguration) SetMaxSessionTime(v int32) {
	o.MaxSessionTime = &v
}

func (o SessionConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SessionConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MaxIdleTime) {
		toSerialize["maxIdleTime"] = o.MaxIdleTime
	}
	if !IsNil(o.RememberMe) {
		toSerialize["rememberMe"] = o.RememberMe
	}
	if !IsNil(o.MaxSessionTime) {
		toSerialize["maxSessionTime"] = o.MaxSessionTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SessionConfiguration) UnmarshalJSON(data []byte) (err error) {
	varSessionConfiguration := _SessionConfiguration{}

	err = json.Unmarshal(data, &varSessionConfiguration)

	if err != nil {
		return err
	}

	*o = SessionConfiguration(varSessionConfiguration)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "maxIdleTime")
		delete(additionalProperties, "rememberMe")
		delete(additionalProperties, "maxSessionTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSessionConfiguration struct {
	value *SessionConfiguration
	isSet bool
}

func (v NullableSessionConfiguration) Get() *SessionConfiguration {
	return v.value
}

func (v *NullableSessionConfiguration) Set(val *SessionConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionConfiguration(val *SessionConfiguration) *NullableSessionConfiguration {
	return &NullableSessionConfiguration{value: val, isSet: true}
}

func (v NullableSessionConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

