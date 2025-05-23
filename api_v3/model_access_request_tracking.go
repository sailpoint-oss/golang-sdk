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

// checks if the AccessRequestTracking type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessRequestTracking{}

// AccessRequestTracking struct for AccessRequestTracking
type AccessRequestTracking struct {
	// The identity id in which the access request is for.
	RequestedFor *string `json:"requestedFor,omitempty"`
	// The details of the item requested.
	RequestedItemsDetails []RequestedItemDetails `json:"requestedItemsDetails,omitempty"`
	// a hash representation of the access requested, useful for longer term tracking client side.
	AttributesHash *int32 `json:"attributesHash,omitempty"`
	// a list of access request identifiers, generally only one will be populated, but high volume requested may result in multiple ids.
	AccessRequestIds []string `json:"accessRequestIds,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccessRequestTracking AccessRequestTracking

// NewAccessRequestTracking instantiates a new AccessRequestTracking object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessRequestTracking() *AccessRequestTracking {
	this := AccessRequestTracking{}
	return &this
}

// NewAccessRequestTrackingWithDefaults instantiates a new AccessRequestTracking object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessRequestTrackingWithDefaults() *AccessRequestTracking {
	this := AccessRequestTracking{}
	return &this
}

// GetRequestedFor returns the RequestedFor field value if set, zero value otherwise.
func (o *AccessRequestTracking) GetRequestedFor() string {
	if o == nil || IsNil(o.RequestedFor) {
		var ret string
		return ret
	}
	return *o.RequestedFor
}

// GetRequestedForOk returns a tuple with the RequestedFor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequestTracking) GetRequestedForOk() (*string, bool) {
	if o == nil || IsNil(o.RequestedFor) {
		return nil, false
	}
	return o.RequestedFor, true
}

// HasRequestedFor returns a boolean if a field has been set.
func (o *AccessRequestTracking) HasRequestedFor() bool {
	if o != nil && !IsNil(o.RequestedFor) {
		return true
	}

	return false
}

// SetRequestedFor gets a reference to the given string and assigns it to the RequestedFor field.
func (o *AccessRequestTracking) SetRequestedFor(v string) {
	o.RequestedFor = &v
}

// GetRequestedItemsDetails returns the RequestedItemsDetails field value if set, zero value otherwise.
func (o *AccessRequestTracking) GetRequestedItemsDetails() []RequestedItemDetails {
	if o == nil || IsNil(o.RequestedItemsDetails) {
		var ret []RequestedItemDetails
		return ret
	}
	return o.RequestedItemsDetails
}

// GetRequestedItemsDetailsOk returns a tuple with the RequestedItemsDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequestTracking) GetRequestedItemsDetailsOk() ([]RequestedItemDetails, bool) {
	if o == nil || IsNil(o.RequestedItemsDetails) {
		return nil, false
	}
	return o.RequestedItemsDetails, true
}

// HasRequestedItemsDetails returns a boolean if a field has been set.
func (o *AccessRequestTracking) HasRequestedItemsDetails() bool {
	if o != nil && !IsNil(o.RequestedItemsDetails) {
		return true
	}

	return false
}

// SetRequestedItemsDetails gets a reference to the given []RequestedItemDetails and assigns it to the RequestedItemsDetails field.
func (o *AccessRequestTracking) SetRequestedItemsDetails(v []RequestedItemDetails) {
	o.RequestedItemsDetails = v
}

// GetAttributesHash returns the AttributesHash field value if set, zero value otherwise.
func (o *AccessRequestTracking) GetAttributesHash() int32 {
	if o == nil || IsNil(o.AttributesHash) {
		var ret int32
		return ret
	}
	return *o.AttributesHash
}

// GetAttributesHashOk returns a tuple with the AttributesHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequestTracking) GetAttributesHashOk() (*int32, bool) {
	if o == nil || IsNil(o.AttributesHash) {
		return nil, false
	}
	return o.AttributesHash, true
}

// HasAttributesHash returns a boolean if a field has been set.
func (o *AccessRequestTracking) HasAttributesHash() bool {
	if o != nil && !IsNil(o.AttributesHash) {
		return true
	}

	return false
}

// SetAttributesHash gets a reference to the given int32 and assigns it to the AttributesHash field.
func (o *AccessRequestTracking) SetAttributesHash(v int32) {
	o.AttributesHash = &v
}

// GetAccessRequestIds returns the AccessRequestIds field value if set, zero value otherwise.
func (o *AccessRequestTracking) GetAccessRequestIds() []string {
	if o == nil || IsNil(o.AccessRequestIds) {
		var ret []string
		return ret
	}
	return o.AccessRequestIds
}

// GetAccessRequestIdsOk returns a tuple with the AccessRequestIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequestTracking) GetAccessRequestIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.AccessRequestIds) {
		return nil, false
	}
	return o.AccessRequestIds, true
}

// HasAccessRequestIds returns a boolean if a field has been set.
func (o *AccessRequestTracking) HasAccessRequestIds() bool {
	if o != nil && !IsNil(o.AccessRequestIds) {
		return true
	}

	return false
}

// SetAccessRequestIds gets a reference to the given []string and assigns it to the AccessRequestIds field.
func (o *AccessRequestTracking) SetAccessRequestIds(v []string) {
	o.AccessRequestIds = v
}

func (o AccessRequestTracking) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessRequestTracking) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RequestedFor) {
		toSerialize["requestedFor"] = o.RequestedFor
	}
	if !IsNil(o.RequestedItemsDetails) {
		toSerialize["requestedItemsDetails"] = o.RequestedItemsDetails
	}
	if !IsNil(o.AttributesHash) {
		toSerialize["attributesHash"] = o.AttributesHash
	}
	if !IsNil(o.AccessRequestIds) {
		toSerialize["accessRequestIds"] = o.AccessRequestIds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccessRequestTracking) UnmarshalJSON(data []byte) (err error) {
	varAccessRequestTracking := _AccessRequestTracking{}

	err = json.Unmarshal(data, &varAccessRequestTracking)

	if err != nil {
		return err
	}

	*o = AccessRequestTracking(varAccessRequestTracking)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "requestedFor")
		delete(additionalProperties, "requestedItemsDetails")
		delete(additionalProperties, "attributesHash")
		delete(additionalProperties, "accessRequestIds")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccessRequestTracking struct {
	value *AccessRequestTracking
	isSet bool
}

func (v NullableAccessRequestTracking) Get() *AccessRequestTracking {
	return v.value
}

func (v *NullableAccessRequestTracking) Set(val *AccessRequestTracking) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessRequestTracking) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessRequestTracking) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessRequestTracking(val *AccessRequestTracking) *NullableAccessRequestTracking {
	return &NullableAccessRequestTracking{value: val, isSet: true}
}

func (v NullableAccessRequestTracking) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessRequestTracking) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


