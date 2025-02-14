/*
Identity Security Cloud V2024 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: v2024
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2024

import (
	"encoding/json"
	
)

// checks if the AccessRequestAdminItemStatusCancelledRequestDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessRequestAdminItemStatusCancelledRequestDetails{}

// AccessRequestAdminItemStatusCancelledRequestDetails struct for AccessRequestAdminItemStatusCancelledRequestDetails
type AccessRequestAdminItemStatusCancelledRequestDetails struct {
	// Comment made by the owner when cancelling the associated request.
	Comment *string `json:"comment,omitempty"`
	Owner *OwnerDto `json:"owner,omitempty"`
	// Date comment was added by the owner when cancelling the associated request.
	Modified *SailPointTime `json:"modified,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccessRequestAdminItemStatusCancelledRequestDetails AccessRequestAdminItemStatusCancelledRequestDetails

// NewAccessRequestAdminItemStatusCancelledRequestDetails instantiates a new AccessRequestAdminItemStatusCancelledRequestDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessRequestAdminItemStatusCancelledRequestDetails() *AccessRequestAdminItemStatusCancelledRequestDetails {
	this := AccessRequestAdminItemStatusCancelledRequestDetails{}
	return &this
}

// NewAccessRequestAdminItemStatusCancelledRequestDetailsWithDefaults instantiates a new AccessRequestAdminItemStatusCancelledRequestDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessRequestAdminItemStatusCancelledRequestDetailsWithDefaults() *AccessRequestAdminItemStatusCancelledRequestDetails {
	this := AccessRequestAdminItemStatusCancelledRequestDetails{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *AccessRequestAdminItemStatusCancelledRequestDetails) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequestAdminItemStatusCancelledRequestDetails) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *AccessRequestAdminItemStatusCancelledRequestDetails) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *AccessRequestAdminItemStatusCancelledRequestDetails) SetComment(v string) {
	o.Comment = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *AccessRequestAdminItemStatusCancelledRequestDetails) GetOwner() OwnerDto {
	if o == nil || IsNil(o.Owner) {
		var ret OwnerDto
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequestAdminItemStatusCancelledRequestDetails) GetOwnerOk() (*OwnerDto, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *AccessRequestAdminItemStatusCancelledRequestDetails) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given OwnerDto and assigns it to the Owner field.
func (o *AccessRequestAdminItemStatusCancelledRequestDetails) SetOwner(v OwnerDto) {
	o.Owner = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *AccessRequestAdminItemStatusCancelledRequestDetails) GetModified() SailPointTime {
	if o == nil || IsNil(o.Modified) {
		var ret SailPointTime
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequestAdminItemStatusCancelledRequestDetails) GetModifiedOk() (*SailPointTime, bool) {
	if o == nil || IsNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *AccessRequestAdminItemStatusCancelledRequestDetails) HasModified() bool {
	if o != nil && !IsNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given SailPointTime and assigns it to the Modified field.
func (o *AccessRequestAdminItemStatusCancelledRequestDetails) SetModified(v SailPointTime) {
	o.Modified = &v
}

func (o AccessRequestAdminItemStatusCancelledRequestDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessRequestAdminItemStatusCancelledRequestDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !IsNil(o.Modified) {
		toSerialize["modified"] = o.Modified
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccessRequestAdminItemStatusCancelledRequestDetails) UnmarshalJSON(data []byte) (err error) {
	varAccessRequestAdminItemStatusCancelledRequestDetails := _AccessRequestAdminItemStatusCancelledRequestDetails{}

	err = json.Unmarshal(data, &varAccessRequestAdminItemStatusCancelledRequestDetails)

	if err != nil {
		return err
	}

	*o = AccessRequestAdminItemStatusCancelledRequestDetails(varAccessRequestAdminItemStatusCancelledRequestDetails)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "comment")
		delete(additionalProperties, "owner")
		delete(additionalProperties, "modified")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccessRequestAdminItemStatusCancelledRequestDetails struct {
	value *AccessRequestAdminItemStatusCancelledRequestDetails
	isSet bool
}

func (v NullableAccessRequestAdminItemStatusCancelledRequestDetails) Get() *AccessRequestAdminItemStatusCancelledRequestDetails {
	return v.value
}

func (v *NullableAccessRequestAdminItemStatusCancelledRequestDetails) Set(val *AccessRequestAdminItemStatusCancelledRequestDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessRequestAdminItemStatusCancelledRequestDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessRequestAdminItemStatusCancelledRequestDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessRequestAdminItemStatusCancelledRequestDetails(val *AccessRequestAdminItemStatusCancelledRequestDetails) *NullableAccessRequestAdminItemStatusCancelledRequestDetails {
	return &NullableAccessRequestAdminItemStatusCancelledRequestDetails{value: val, isSet: true}
}

func (v NullableAccessRequestAdminItemStatusCancelledRequestDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessRequestAdminItemStatusCancelledRequestDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


