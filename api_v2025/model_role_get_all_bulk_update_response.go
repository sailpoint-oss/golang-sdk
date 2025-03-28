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

// checks if the RoleGetAllBulkUpdateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoleGetAllBulkUpdateResponse{}

// RoleGetAllBulkUpdateResponse struct for RoleGetAllBulkUpdateResponse
type RoleGetAllBulkUpdateResponse struct {
	// ID of the task which is executing the bulk update. This also used in to the bulk-update/_** API to track status.
	Id *string `json:"id,omitempty"`
	// Type of the bulk update object.
	Type *string `json:"type,omitempty"`
	// The status of the bulk update request, only list unfinished request's status, the status could also checked by getBulkUpdateStatus API
	Status *string `json:"status,omitempty"`
	// Time when the bulk update request was created
	Created *SailPointTime `json:"created,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RoleGetAllBulkUpdateResponse RoleGetAllBulkUpdateResponse

// NewRoleGetAllBulkUpdateResponse instantiates a new RoleGetAllBulkUpdateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleGetAllBulkUpdateResponse() *RoleGetAllBulkUpdateResponse {
	this := RoleGetAllBulkUpdateResponse{}
	return &this
}

// NewRoleGetAllBulkUpdateResponseWithDefaults instantiates a new RoleGetAllBulkUpdateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleGetAllBulkUpdateResponseWithDefaults() *RoleGetAllBulkUpdateResponse {
	this := RoleGetAllBulkUpdateResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RoleGetAllBulkUpdateResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleGetAllBulkUpdateResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RoleGetAllBulkUpdateResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RoleGetAllBulkUpdateResponse) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RoleGetAllBulkUpdateResponse) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleGetAllBulkUpdateResponse) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RoleGetAllBulkUpdateResponse) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RoleGetAllBulkUpdateResponse) SetType(v string) {
	o.Type = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *RoleGetAllBulkUpdateResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleGetAllBulkUpdateResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *RoleGetAllBulkUpdateResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *RoleGetAllBulkUpdateResponse) SetStatus(v string) {
	o.Status = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *RoleGetAllBulkUpdateResponse) GetCreated() SailPointTime {
	if o == nil || IsNil(o.Created) {
		var ret SailPointTime
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleGetAllBulkUpdateResponse) GetCreatedOk() (*SailPointTime, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *RoleGetAllBulkUpdateResponse) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given SailPointTime and assigns it to the Created field.
func (o *RoleGetAllBulkUpdateResponse) SetCreated(v SailPointTime) {
	o.Created = &v
}

func (o RoleGetAllBulkUpdateResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoleGetAllBulkUpdateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RoleGetAllBulkUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	varRoleGetAllBulkUpdateResponse := _RoleGetAllBulkUpdateResponse{}

	err = json.Unmarshal(data, &varRoleGetAllBulkUpdateResponse)

	if err != nil {
		return err
	}

	*o = RoleGetAllBulkUpdateResponse(varRoleGetAllBulkUpdateResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "status")
		delete(additionalProperties, "created")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRoleGetAllBulkUpdateResponse struct {
	value *RoleGetAllBulkUpdateResponse
	isSet bool
}

func (v NullableRoleGetAllBulkUpdateResponse) Get() *RoleGetAllBulkUpdateResponse {
	return v.value
}

func (v *NullableRoleGetAllBulkUpdateResponse) Set(val *RoleGetAllBulkUpdateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleGetAllBulkUpdateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleGetAllBulkUpdateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleGetAllBulkUpdateResponse(val *RoleGetAllBulkUpdateResponse) *NullableRoleGetAllBulkUpdateResponse {
	return &NullableRoleGetAllBulkUpdateResponse{value: val, isSet: true}
}

func (v NullableRoleGetAllBulkUpdateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleGetAllBulkUpdateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


