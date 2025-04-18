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

// checks if the PasswordDigitTokenReset type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PasswordDigitTokenReset{}

// PasswordDigitTokenReset struct for PasswordDigitTokenReset
type PasswordDigitTokenReset struct {
	// The uid of the user requested for digit token
	UserId string `json:"userId"`
	// The length of digit token. It should be from 6 to 18, inclusive. The default value is 6.
	Length *int32 `json:"length,omitempty"`
	// The time to live for the digit token in minutes. The default value is 5 minutes.
	DurationMinutes *int32 `json:"durationMinutes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PasswordDigitTokenReset PasswordDigitTokenReset

// NewPasswordDigitTokenReset instantiates a new PasswordDigitTokenReset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordDigitTokenReset(userId string) *PasswordDigitTokenReset {
	this := PasswordDigitTokenReset{}
	this.UserId = userId
	return &this
}

// NewPasswordDigitTokenResetWithDefaults instantiates a new PasswordDigitTokenReset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordDigitTokenResetWithDefaults() *PasswordDigitTokenReset {
	this := PasswordDigitTokenReset{}
	return &this
}

// GetUserId returns the UserId field value
func (o *PasswordDigitTokenReset) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *PasswordDigitTokenReset) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *PasswordDigitTokenReset) SetUserId(v string) {
	o.UserId = v
}

// GetLength returns the Length field value if set, zero value otherwise.
func (o *PasswordDigitTokenReset) GetLength() int32 {
	if o == nil || IsNil(o.Length) {
		var ret int32
		return ret
	}
	return *o.Length
}

// GetLengthOk returns a tuple with the Length field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordDigitTokenReset) GetLengthOk() (*int32, bool) {
	if o == nil || IsNil(o.Length) {
		return nil, false
	}
	return o.Length, true
}

// HasLength returns a boolean if a field has been set.
func (o *PasswordDigitTokenReset) HasLength() bool {
	if o != nil && !IsNil(o.Length) {
		return true
	}

	return false
}

// SetLength gets a reference to the given int32 and assigns it to the Length field.
func (o *PasswordDigitTokenReset) SetLength(v int32) {
	o.Length = &v
}

// GetDurationMinutes returns the DurationMinutes field value if set, zero value otherwise.
func (o *PasswordDigitTokenReset) GetDurationMinutes() int32 {
	if o == nil || IsNil(o.DurationMinutes) {
		var ret int32
		return ret
	}
	return *o.DurationMinutes
}

// GetDurationMinutesOk returns a tuple with the DurationMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordDigitTokenReset) GetDurationMinutesOk() (*int32, bool) {
	if o == nil || IsNil(o.DurationMinutes) {
		return nil, false
	}
	return o.DurationMinutes, true
}

// HasDurationMinutes returns a boolean if a field has been set.
func (o *PasswordDigitTokenReset) HasDurationMinutes() bool {
	if o != nil && !IsNil(o.DurationMinutes) {
		return true
	}

	return false
}

// SetDurationMinutes gets a reference to the given int32 and assigns it to the DurationMinutes field.
func (o *PasswordDigitTokenReset) SetDurationMinutes(v int32) {
	o.DurationMinutes = &v
}

func (o PasswordDigitTokenReset) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PasswordDigitTokenReset) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["userId"] = o.UserId
	if !IsNil(o.Length) {
		toSerialize["length"] = o.Length
	}
	if !IsNil(o.DurationMinutes) {
		toSerialize["durationMinutes"] = o.DurationMinutes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PasswordDigitTokenReset) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"userId",
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

	varPasswordDigitTokenReset := _PasswordDigitTokenReset{}

	err = json.Unmarshal(data, &varPasswordDigitTokenReset)

	if err != nil {
		return err
	}

	*o = PasswordDigitTokenReset(varPasswordDigitTokenReset)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "userId")
		delete(additionalProperties, "length")
		delete(additionalProperties, "durationMinutes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePasswordDigitTokenReset struct {
	value *PasswordDigitTokenReset
	isSet bool
}

func (v NullablePasswordDigitTokenReset) Get() *PasswordDigitTokenReset {
	return v.value
}

func (v *NullablePasswordDigitTokenReset) Set(val *PasswordDigitTokenReset) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordDigitTokenReset) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordDigitTokenReset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordDigitTokenReset(val *PasswordDigitTokenReset) *NullablePasswordDigitTokenReset {
	return &NullablePasswordDigitTokenReset{value: val, isSet: true}
}

func (v NullablePasswordDigitTokenReset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordDigitTokenReset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


