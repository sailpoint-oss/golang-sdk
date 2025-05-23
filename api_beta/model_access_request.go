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

// checks if the AccessRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessRequest{}

// AccessRequest struct for AccessRequest
type AccessRequest struct {
	// A list of Identity IDs for whom the Access is requested. If it's a Revoke request, there can only be one Identity ID.
	RequestedFor []string `json:"requestedFor"`
	RequestType NullableAccessRequestType `json:"requestType,omitempty"`
	RequestedItems []AccessRequestItem `json:"requestedItems"`
	// Arbitrary key-value pairs. They will never be processed by the IdentityNow system but will be returned on associated APIs such as /account-activities.
	ClientMetadata *map[string]string `json:"clientMetadata,omitempty"`
	// Additional submit data structure with requestedFor containing requestedItems allowing distinction for each request item and Identity. * Can only be used when 'requestedFor' and 'requestedItems' are not separately provided * Adds ability to specify which account the user wants the access on, in case they have multiple accounts on a source * Allows the ability to request items with different remove dates * Also allows different combinations of request items and identities in the same request * Only for use in GRANT_ACCESS type requests 
	RequestedForWithRequestedItems []RequestedForDtoRef `json:"requestedForWithRequestedItems,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccessRequest AccessRequest

// NewAccessRequest instantiates a new AccessRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessRequest(requestedFor []string, requestedItems []AccessRequestItem) *AccessRequest {
	this := AccessRequest{}
	this.RequestedFor = requestedFor
	this.RequestedItems = requestedItems
	return &this
}

// NewAccessRequestWithDefaults instantiates a new AccessRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessRequestWithDefaults() *AccessRequest {
	this := AccessRequest{}
	return &this
}

// GetRequestedFor returns the RequestedFor field value
func (o *AccessRequest) GetRequestedFor() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RequestedFor
}

// GetRequestedForOk returns a tuple with the RequestedFor field value
// and a boolean to check if the value has been set.
func (o *AccessRequest) GetRequestedForOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestedFor, true
}

// SetRequestedFor sets field value
func (o *AccessRequest) SetRequestedFor(v []string) {
	o.RequestedFor = v
}

// GetRequestType returns the RequestType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessRequest) GetRequestType() AccessRequestType {
	if o == nil || IsNil(o.RequestType.Get()) {
		var ret AccessRequestType
		return ret
	}
	return *o.RequestType.Get()
}

// GetRequestTypeOk returns a tuple with the RequestType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessRequest) GetRequestTypeOk() (*AccessRequestType, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestType.Get(), o.RequestType.IsSet()
}

// HasRequestType returns a boolean if a field has been set.
func (o *AccessRequest) HasRequestType() bool {
	if o != nil && o.RequestType.IsSet() {
		return true
	}

	return false
}

// SetRequestType gets a reference to the given NullableAccessRequestType and assigns it to the RequestType field.
func (o *AccessRequest) SetRequestType(v AccessRequestType) {
	o.RequestType.Set(&v)
}
// SetRequestTypeNil sets the value for RequestType to be an explicit nil
func (o *AccessRequest) SetRequestTypeNil() {
	o.RequestType.Set(nil)
}

// UnsetRequestType ensures that no value is present for RequestType, not even an explicit nil
func (o *AccessRequest) UnsetRequestType() {
	o.RequestType.Unset()
}

// GetRequestedItems returns the RequestedItems field value
func (o *AccessRequest) GetRequestedItems() []AccessRequestItem {
	if o == nil {
		var ret []AccessRequestItem
		return ret
	}

	return o.RequestedItems
}

// GetRequestedItemsOk returns a tuple with the RequestedItems field value
// and a boolean to check if the value has been set.
func (o *AccessRequest) GetRequestedItemsOk() ([]AccessRequestItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestedItems, true
}

// SetRequestedItems sets field value
func (o *AccessRequest) SetRequestedItems(v []AccessRequestItem) {
	o.RequestedItems = v
}

// GetClientMetadata returns the ClientMetadata field value if set, zero value otherwise.
func (o *AccessRequest) GetClientMetadata() map[string]string {
	if o == nil || IsNil(o.ClientMetadata) {
		var ret map[string]string
		return ret
	}
	return *o.ClientMetadata
}

// GetClientMetadataOk returns a tuple with the ClientMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequest) GetClientMetadataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.ClientMetadata) {
		return nil, false
	}
	return o.ClientMetadata, true
}

// HasClientMetadata returns a boolean if a field has been set.
func (o *AccessRequest) HasClientMetadata() bool {
	if o != nil && !IsNil(o.ClientMetadata) {
		return true
	}

	return false
}

// SetClientMetadata gets a reference to the given map[string]string and assigns it to the ClientMetadata field.
func (o *AccessRequest) SetClientMetadata(v map[string]string) {
	o.ClientMetadata = &v
}

// GetRequestedForWithRequestedItems returns the RequestedForWithRequestedItems field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessRequest) GetRequestedForWithRequestedItems() []RequestedForDtoRef {
	if o == nil {
		var ret []RequestedForDtoRef
		return ret
	}
	return o.RequestedForWithRequestedItems
}

// GetRequestedForWithRequestedItemsOk returns a tuple with the RequestedForWithRequestedItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessRequest) GetRequestedForWithRequestedItemsOk() ([]RequestedForDtoRef, bool) {
	if o == nil || IsNil(o.RequestedForWithRequestedItems) {
		return nil, false
	}
	return o.RequestedForWithRequestedItems, true
}

// HasRequestedForWithRequestedItems returns a boolean if a field has been set.
func (o *AccessRequest) HasRequestedForWithRequestedItems() bool {
	if o != nil && !IsNil(o.RequestedForWithRequestedItems) {
		return true
	}

	return false
}

// SetRequestedForWithRequestedItems gets a reference to the given []RequestedForDtoRef and assigns it to the RequestedForWithRequestedItems field.
func (o *AccessRequest) SetRequestedForWithRequestedItems(v []RequestedForDtoRef) {
	o.RequestedForWithRequestedItems = v
}

func (o AccessRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["requestedFor"] = o.RequestedFor
	if o.RequestType.IsSet() {
		toSerialize["requestType"] = o.RequestType.Get()
	}
	toSerialize["requestedItems"] = o.RequestedItems
	if !IsNil(o.ClientMetadata) {
		toSerialize["clientMetadata"] = o.ClientMetadata
	}
	if o.RequestedForWithRequestedItems != nil {
		toSerialize["requestedForWithRequestedItems"] = o.RequestedForWithRequestedItems
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccessRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"requestedFor",
		"requestedItems",
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

	varAccessRequest := _AccessRequest{}

	err = json.Unmarshal(data, &varAccessRequest)

	if err != nil {
		return err
	}

	*o = AccessRequest(varAccessRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "requestedFor")
		delete(additionalProperties, "requestType")
		delete(additionalProperties, "requestedItems")
		delete(additionalProperties, "clientMetadata")
		delete(additionalProperties, "requestedForWithRequestedItems")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccessRequest struct {
	value *AccessRequest
	isSet bool
}

func (v NullableAccessRequest) Get() *AccessRequest {
	return v.value
}

func (v *NullableAccessRequest) Set(val *AccessRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessRequest(val *AccessRequest) *NullableAccessRequest {
	return &NullableAccessRequest{value: val, isSet: true}
}

func (v NullableAccessRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


