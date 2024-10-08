/*
Identity Security Cloud Beta API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_beta

import (
	"encoding/json"
)

// checks if the SendTokenResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SendTokenResponse{}

// SendTokenResponse struct for SendTokenResponse
type SendTokenResponse struct {
	// The token request ID
	RequestId NullableString `json:"requestId,omitempty"`
	// Status of sending token
	Status *string `json:"status,omitempty"`
	// Error messages from token send request
	ErrorMessage NullableString `json:"errorMessage,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SendTokenResponse SendTokenResponse

// NewSendTokenResponse instantiates a new SendTokenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSendTokenResponse() *SendTokenResponse {
	this := SendTokenResponse{}
	return &this
}

// NewSendTokenResponseWithDefaults instantiates a new SendTokenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSendTokenResponseWithDefaults() *SendTokenResponse {
	this := SendTokenResponse{}
	return &this
}

// GetRequestId returns the RequestId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SendTokenResponse) GetRequestId() string {
	if o == nil || IsNil(o.RequestId.Get()) {
		var ret string
		return ret
	}
	return *o.RequestId.Get()
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SendTokenResponse) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestId.Get(), o.RequestId.IsSet()
}

// HasRequestId returns a boolean if a field has been set.
func (o *SendTokenResponse) HasRequestId() bool {
	if o != nil && o.RequestId.IsSet() {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given NullableString and assigns it to the RequestId field.
func (o *SendTokenResponse) SetRequestId(v string) {
	o.RequestId.Set(&v)
}
// SetRequestIdNil sets the value for RequestId to be an explicit nil
func (o *SendTokenResponse) SetRequestIdNil() {
	o.RequestId.Set(nil)
}

// UnsetRequestId ensures that no value is present for RequestId, not even an explicit nil
func (o *SendTokenResponse) UnsetRequestId() {
	o.RequestId.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SendTokenResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendTokenResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SendTokenResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SendTokenResponse) SetStatus(v string) {
	o.Status = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SendTokenResponse) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage.Get()) {
		var ret string
		return ret
	}
	return *o.ErrorMessage.Get()
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SendTokenResponse) GetErrorMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorMessage.Get(), o.ErrorMessage.IsSet()
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *SendTokenResponse) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage.IsSet() {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given NullableString and assigns it to the ErrorMessage field.
func (o *SendTokenResponse) SetErrorMessage(v string) {
	o.ErrorMessage.Set(&v)
}
// SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil
func (o *SendTokenResponse) SetErrorMessageNil() {
	o.ErrorMessage.Set(nil)
}

// UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
func (o *SendTokenResponse) UnsetErrorMessage() {
	o.ErrorMessage.Unset()
}

func (o SendTokenResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SendTokenResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.RequestId.IsSet() {
		toSerialize["requestId"] = o.RequestId.Get()
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.ErrorMessage.IsSet() {
		toSerialize["errorMessage"] = o.ErrorMessage.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SendTokenResponse) UnmarshalJSON(data []byte) (err error) {
	varSendTokenResponse := _SendTokenResponse{}

	err = json.Unmarshal(data, &varSendTokenResponse)

	if err != nil {
		return err
	}

	*o = SendTokenResponse(varSendTokenResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "requestId")
		delete(additionalProperties, "status")
		delete(additionalProperties, "errorMessage")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSendTokenResponse struct {
	value *SendTokenResponse
	isSet bool
}

func (v NullableSendTokenResponse) Get() *SendTokenResponse {
	return v.value
}

func (v *NullableSendTokenResponse) Set(val *SendTokenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSendTokenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSendTokenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSendTokenResponse(val *SendTokenResponse) *NullableSendTokenResponse {
	return &NullableSendTokenResponse{value: val, isSet: true}
}

func (v NullableSendTokenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSendTokenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


