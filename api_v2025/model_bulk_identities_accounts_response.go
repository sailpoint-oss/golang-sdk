/*
Identity Security Cloud V2025 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: v2025
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2025

import (
	"encoding/json"
)

// checks if the BulkIdentitiesAccountsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkIdentitiesAccountsResponse{}

// BulkIdentitiesAccountsResponse Bulk response object.
type BulkIdentitiesAccountsResponse struct {
	// Identifier of bulk request item.
	Id *string `json:"id,omitempty"`
	// Response status value.
	StatusCode *int32 `json:"statusCode,omitempty"`
	// Status containing additional context information about failures.
	Message *string `json:"message,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BulkIdentitiesAccountsResponse BulkIdentitiesAccountsResponse

// NewBulkIdentitiesAccountsResponse instantiates a new BulkIdentitiesAccountsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkIdentitiesAccountsResponse() *BulkIdentitiesAccountsResponse {
	this := BulkIdentitiesAccountsResponse{}
	return &this
}

// NewBulkIdentitiesAccountsResponseWithDefaults instantiates a new BulkIdentitiesAccountsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkIdentitiesAccountsResponseWithDefaults() *BulkIdentitiesAccountsResponse {
	this := BulkIdentitiesAccountsResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BulkIdentitiesAccountsResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkIdentitiesAccountsResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BulkIdentitiesAccountsResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BulkIdentitiesAccountsResponse) SetId(v string) {
	o.Id = &v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *BulkIdentitiesAccountsResponse) GetStatusCode() int32 {
	if o == nil || IsNil(o.StatusCode) {
		var ret int32
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkIdentitiesAccountsResponse) GetStatusCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.StatusCode) {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *BulkIdentitiesAccountsResponse) HasStatusCode() bool {
	if o != nil && !IsNil(o.StatusCode) {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given int32 and assigns it to the StatusCode field.
func (o *BulkIdentitiesAccountsResponse) SetStatusCode(v int32) {
	o.StatusCode = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *BulkIdentitiesAccountsResponse) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkIdentitiesAccountsResponse) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *BulkIdentitiesAccountsResponse) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *BulkIdentitiesAccountsResponse) SetMessage(v string) {
	o.Message = &v
}

func (o BulkIdentitiesAccountsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkIdentitiesAccountsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.StatusCode) {
		toSerialize["statusCode"] = o.StatusCode
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BulkIdentitiesAccountsResponse) UnmarshalJSON(data []byte) (err error) {
	varBulkIdentitiesAccountsResponse := _BulkIdentitiesAccountsResponse{}

	err = json.Unmarshal(data, &varBulkIdentitiesAccountsResponse)

	if err != nil {
		return err
	}

	*o = BulkIdentitiesAccountsResponse(varBulkIdentitiesAccountsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "statusCode")
		delete(additionalProperties, "message")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBulkIdentitiesAccountsResponse struct {
	value *BulkIdentitiesAccountsResponse
	isSet bool
}

func (v NullableBulkIdentitiesAccountsResponse) Get() *BulkIdentitiesAccountsResponse {
	return v.value
}

func (v *NullableBulkIdentitiesAccountsResponse) Set(val *BulkIdentitiesAccountsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkIdentitiesAccountsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkIdentitiesAccountsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkIdentitiesAccountsResponse(val *BulkIdentitiesAccountsResponse) *NullableBulkIdentitiesAccountsResponse {
	return &NullableBulkIdentitiesAccountsResponse{value: val, isSet: true}
}

func (v NullableBulkIdentitiesAccountsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkIdentitiesAccountsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


