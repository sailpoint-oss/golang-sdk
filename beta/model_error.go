/*
IdentityNow Beta API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package beta

import (
	"encoding/json"
)

// checks if the Error type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Error{}

// Error struct for Error
type Error struct {
	// DetailCode is the text of the status code returned
	DetailCode *string `json:"detailCode,omitempty"`
	Messages []ErrorMessage `json:"messages,omitempty"`
	// TrackingID is the request tracking unique identifier
	TrackingId *string `json:"trackingId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Error Error

// NewError instantiates a new Error object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewError() *Error {
	this := Error{}
	return &this
}

// NewErrorWithDefaults instantiates a new Error object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorWithDefaults() *Error {
	this := Error{}
	return &this
}

// GetDetailCode returns the DetailCode field value if set, zero value otherwise.
func (o *Error) GetDetailCode() string {
	if o == nil || isNil(o.DetailCode) {
		var ret string
		return ret
	}
	return *o.DetailCode
}

// GetDetailCodeOk returns a tuple with the DetailCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Error) GetDetailCodeOk() (*string, bool) {
	if o == nil || isNil(o.DetailCode) {
		return nil, false
	}
	return o.DetailCode, true
}

// HasDetailCode returns a boolean if a field has been set.
func (o *Error) HasDetailCode() bool {
	if o != nil && !isNil(o.DetailCode) {
		return true
	}

	return false
}

// SetDetailCode gets a reference to the given string and assigns it to the DetailCode field.
func (o *Error) SetDetailCode(v string) {
	o.DetailCode = &v
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *Error) GetMessages() []ErrorMessage {
	if o == nil || isNil(o.Messages) {
		var ret []ErrorMessage
		return ret
	}
	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Error) GetMessagesOk() ([]ErrorMessage, bool) {
	if o == nil || isNil(o.Messages) {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *Error) HasMessages() bool {
	if o != nil && !isNil(o.Messages) {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []ErrorMessage and assigns it to the Messages field.
func (o *Error) SetMessages(v []ErrorMessage) {
	o.Messages = v
}

// GetTrackingId returns the TrackingId field value if set, zero value otherwise.
func (o *Error) GetTrackingId() string {
	if o == nil || isNil(o.TrackingId) {
		var ret string
		return ret
	}
	return *o.TrackingId
}

// GetTrackingIdOk returns a tuple with the TrackingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Error) GetTrackingIdOk() (*string, bool) {
	if o == nil || isNil(o.TrackingId) {
		return nil, false
	}
	return o.TrackingId, true
}

// HasTrackingId returns a boolean if a field has been set.
func (o *Error) HasTrackingId() bool {
	if o != nil && !isNil(o.TrackingId) {
		return true
	}

	return false
}

// SetTrackingId gets a reference to the given string and assigns it to the TrackingId field.
func (o *Error) SetTrackingId(v string) {
	o.TrackingId = &v
}

func (o Error) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Error) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DetailCode) {
		toSerialize["detailCode"] = o.DetailCode
	}
	if !isNil(o.Messages) {
		toSerialize["messages"] = o.Messages
	}
	if !isNil(o.TrackingId) {
		toSerialize["trackingId"] = o.TrackingId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Error) UnmarshalJSON(bytes []byte) (err error) {
	varError := _Error{}

	if err = json.Unmarshal(bytes, &varError); err == nil {
		*o = Error(varError)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "detailCode")
		delete(additionalProperties, "messages")
		delete(additionalProperties, "trackingId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableError struct {
	value *Error
	isSet bool
}

func (v NullableError) Get() *Error {
	return v.value
}

func (v *NullableError) Set(val *Error) {
	v.value = val
	v.isSet = true
}

func (v NullableError) IsSet() bool {
	return v.isSet
}

func (v *NullableError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableError(val *Error) *NullableError {
	return &NullableError{value: val, isSet: true}
}

func (v NullableError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

