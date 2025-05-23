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

// checks if the AuditDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuditDetails{}

// AuditDetails Audit details for the reassignment configuration of an identity
type AuditDetails struct {
	// Initial date and time when the record was created
	Created *SailPointTime `json:"created,omitempty"`
	CreatedBy *Identity1 `json:"createdBy,omitempty"`
	// Last modified date and time for the record
	Modified *SailPointTime `json:"modified,omitempty"`
	ModifiedBy *Identity1 `json:"modifiedBy,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuditDetails AuditDetails

// NewAuditDetails instantiates a new AuditDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuditDetails() *AuditDetails {
	this := AuditDetails{}
	return &this
}

// NewAuditDetailsWithDefaults instantiates a new AuditDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditDetailsWithDefaults() *AuditDetails {
	this := AuditDetails{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *AuditDetails) GetCreated() SailPointTime {
	if o == nil || IsNil(o.Created) {
		var ret SailPointTime
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditDetails) GetCreatedOk() (*SailPointTime, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *AuditDetails) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given SailPointTime and assigns it to the Created field.
func (o *AuditDetails) SetCreated(v SailPointTime) {
	o.Created = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *AuditDetails) GetCreatedBy() Identity1 {
	if o == nil || IsNil(o.CreatedBy) {
		var ret Identity1
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditDetails) GetCreatedByOk() (*Identity1, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *AuditDetails) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given Identity1 and assigns it to the CreatedBy field.
func (o *AuditDetails) SetCreatedBy(v Identity1) {
	o.CreatedBy = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *AuditDetails) GetModified() SailPointTime {
	if o == nil || IsNil(o.Modified) {
		var ret SailPointTime
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditDetails) GetModifiedOk() (*SailPointTime, bool) {
	if o == nil || IsNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *AuditDetails) HasModified() bool {
	if o != nil && !IsNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given SailPointTime and assigns it to the Modified field.
func (o *AuditDetails) SetModified(v SailPointTime) {
	o.Modified = &v
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise.
func (o *AuditDetails) GetModifiedBy() Identity1 {
	if o == nil || IsNil(o.ModifiedBy) {
		var ret Identity1
		return ret
	}
	return *o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditDetails) GetModifiedByOk() (*Identity1, bool) {
	if o == nil || IsNil(o.ModifiedBy) {
		return nil, false
	}
	return o.ModifiedBy, true
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *AuditDetails) HasModifiedBy() bool {
	if o != nil && !IsNil(o.ModifiedBy) {
		return true
	}

	return false
}

// SetModifiedBy gets a reference to the given Identity1 and assigns it to the ModifiedBy field.
func (o *AuditDetails) SetModifiedBy(v Identity1) {
	o.ModifiedBy = &v
}

func (o AuditDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuditDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if !IsNil(o.Modified) {
		toSerialize["modified"] = o.Modified
	}
	if !IsNil(o.ModifiedBy) {
		toSerialize["modifiedBy"] = o.ModifiedBy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AuditDetails) UnmarshalJSON(data []byte) (err error) {
	varAuditDetails := _AuditDetails{}

	err = json.Unmarshal(data, &varAuditDetails)

	if err != nil {
		return err
	}

	*o = AuditDetails(varAuditDetails)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "created")
		delete(additionalProperties, "createdBy")
		delete(additionalProperties, "modified")
		delete(additionalProperties, "modifiedBy")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAuditDetails struct {
	value *AuditDetails
	isSet bool
}

func (v NullableAuditDetails) Get() *AuditDetails {
	return v.value
}

func (v *NullableAuditDetails) Set(val *AuditDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableAuditDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableAuditDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuditDetails(val *AuditDetails) *NullableAuditDetails {
	return &NullableAuditDetails{value: val, isSet: true}
}

func (v NullableAuditDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuditDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


