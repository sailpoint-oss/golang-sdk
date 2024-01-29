/*
IdentityNow Beta API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_beta

import (
	"encoding/json"
	"time"
)

// checks if the CommonAccessResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommonAccessResponse{}

// CommonAccessResponse struct for CommonAccessResponse
type CommonAccessResponse struct {
	Access *CommonAccessItemAccess `json:"access,omitempty"`
	// CONFIRMED or DENIED
	Status *string `json:"status,omitempty"`
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	// true if user has confirmed or denied status
	ReviewedByUser *bool `json:"reviewedByUser,omitempty"`
	LastReviewed *time.Time `json:"lastReviewed,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CommonAccessResponse CommonAccessResponse

// NewCommonAccessResponse instantiates a new CommonAccessResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonAccessResponse() *CommonAccessResponse {
	this := CommonAccessResponse{}
	return &this
}

// NewCommonAccessResponseWithDefaults instantiates a new CommonAccessResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonAccessResponseWithDefaults() *CommonAccessResponse {
	this := CommonAccessResponse{}
	return &this
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *CommonAccessResponse) GetAccess() CommonAccessItemAccess {
	if o == nil || isNil(o.Access) {
		var ret CommonAccessItemAccess
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonAccessResponse) GetAccessOk() (*CommonAccessItemAccess, bool) {
	if o == nil || isNil(o.Access) {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *CommonAccessResponse) HasAccess() bool {
	if o != nil && !isNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given CommonAccessItemAccess and assigns it to the Access field.
func (o *CommonAccessResponse) SetAccess(v CommonAccessItemAccess) {
	o.Access = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CommonAccessResponse) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonAccessResponse) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CommonAccessResponse) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CommonAccessResponse) SetStatus(v string) {
	o.Status = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *CommonAccessResponse) GetLastUpdated() time.Time {
	if o == nil || isNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonAccessResponse) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || isNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *CommonAccessResponse) HasLastUpdated() bool {
	if o != nil && !isNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *CommonAccessResponse) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetReviewedByUser returns the ReviewedByUser field value if set, zero value otherwise.
func (o *CommonAccessResponse) GetReviewedByUser() bool {
	if o == nil || isNil(o.ReviewedByUser) {
		var ret bool
		return ret
	}
	return *o.ReviewedByUser
}

// GetReviewedByUserOk returns a tuple with the ReviewedByUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonAccessResponse) GetReviewedByUserOk() (*bool, bool) {
	if o == nil || isNil(o.ReviewedByUser) {
		return nil, false
	}
	return o.ReviewedByUser, true
}

// HasReviewedByUser returns a boolean if a field has been set.
func (o *CommonAccessResponse) HasReviewedByUser() bool {
	if o != nil && !isNil(o.ReviewedByUser) {
		return true
	}

	return false
}

// SetReviewedByUser gets a reference to the given bool and assigns it to the ReviewedByUser field.
func (o *CommonAccessResponse) SetReviewedByUser(v bool) {
	o.ReviewedByUser = &v
}

// GetLastReviewed returns the LastReviewed field value if set, zero value otherwise.
func (o *CommonAccessResponse) GetLastReviewed() time.Time {
	if o == nil || isNil(o.LastReviewed) {
		var ret time.Time
		return ret
	}
	return *o.LastReviewed
}

// GetLastReviewedOk returns a tuple with the LastReviewed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonAccessResponse) GetLastReviewedOk() (*time.Time, bool) {
	if o == nil || isNil(o.LastReviewed) {
		return nil, false
	}
	return o.LastReviewed, true
}

// HasLastReviewed returns a boolean if a field has been set.
func (o *CommonAccessResponse) HasLastReviewed() bool {
	if o != nil && !isNil(o.LastReviewed) {
		return true
	}

	return false
}

// SetLastReviewed gets a reference to the given time.Time and assigns it to the LastReviewed field.
func (o *CommonAccessResponse) SetLastReviewed(v time.Time) {
	o.LastReviewed = &v
}

func (o CommonAccessResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommonAccessResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Access) {
		toSerialize["access"] = o.Access
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	// skip: lastUpdated is readOnly
	if !isNil(o.ReviewedByUser) {
		toSerialize["reviewedByUser"] = o.ReviewedByUser
	}
	// skip: lastReviewed is readOnly

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CommonAccessResponse) UnmarshalJSON(bytes []byte) (err error) {
	varCommonAccessResponse := _CommonAccessResponse{}

	if err = json.Unmarshal(bytes, &varCommonAccessResponse); err == nil {
	*o = CommonAccessResponse(varCommonAccessResponse)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "access")
		delete(additionalProperties, "status")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "reviewedByUser")
		delete(additionalProperties, "lastReviewed")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCommonAccessResponse struct {
	value *CommonAccessResponse
	isSet bool
}

func (v NullableCommonAccessResponse) Get() *CommonAccessResponse {
	return v.value
}

func (v *NullableCommonAccessResponse) Set(val *CommonAccessResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonAccessResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonAccessResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonAccessResponse(val *CommonAccessResponse) *NullableCommonAccessResponse {
	return &NullableCommonAccessResponse{value: val, isSet: true}
}

func (v NullableCommonAccessResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonAccessResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

