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

// checks if the CampaignAllOfMachineAccountCampaignInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignAllOfMachineAccountCampaignInfo{}

// CampaignAllOfMachineAccountCampaignInfo Must be set only if the campaign type is MACHINE_ACCOUNT.
type CampaignAllOfMachineAccountCampaignInfo struct {
	// The list of sources to be included in the campaign.
	SourceIds []string `json:"sourceIds,omitempty"`
	// The reviewer's type.
	ReviewerType *string `json:"reviewerType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CampaignAllOfMachineAccountCampaignInfo CampaignAllOfMachineAccountCampaignInfo

// NewCampaignAllOfMachineAccountCampaignInfo instantiates a new CampaignAllOfMachineAccountCampaignInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignAllOfMachineAccountCampaignInfo() *CampaignAllOfMachineAccountCampaignInfo {
	this := CampaignAllOfMachineAccountCampaignInfo{}
	return &this
}

// NewCampaignAllOfMachineAccountCampaignInfoWithDefaults instantiates a new CampaignAllOfMachineAccountCampaignInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignAllOfMachineAccountCampaignInfoWithDefaults() *CampaignAllOfMachineAccountCampaignInfo {
	this := CampaignAllOfMachineAccountCampaignInfo{}
	return &this
}

// GetSourceIds returns the SourceIds field value if set, zero value otherwise.
func (o *CampaignAllOfMachineAccountCampaignInfo) GetSourceIds() []string {
	if o == nil || IsNil(o.SourceIds) {
		var ret []string
		return ret
	}
	return o.SourceIds
}

// GetSourceIdsOk returns a tuple with the SourceIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignAllOfMachineAccountCampaignInfo) GetSourceIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.SourceIds) {
		return nil, false
	}
	return o.SourceIds, true
}

// HasSourceIds returns a boolean if a field has been set.
func (o *CampaignAllOfMachineAccountCampaignInfo) HasSourceIds() bool {
	if o != nil && !IsNil(o.SourceIds) {
		return true
	}

	return false
}

// SetSourceIds gets a reference to the given []string and assigns it to the SourceIds field.
func (o *CampaignAllOfMachineAccountCampaignInfo) SetSourceIds(v []string) {
	o.SourceIds = v
}

// GetReviewerType returns the ReviewerType field value if set, zero value otherwise.
func (o *CampaignAllOfMachineAccountCampaignInfo) GetReviewerType() string {
	if o == nil || IsNil(o.ReviewerType) {
		var ret string
		return ret
	}
	return *o.ReviewerType
}

// GetReviewerTypeOk returns a tuple with the ReviewerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignAllOfMachineAccountCampaignInfo) GetReviewerTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ReviewerType) {
		return nil, false
	}
	return o.ReviewerType, true
}

// HasReviewerType returns a boolean if a field has been set.
func (o *CampaignAllOfMachineAccountCampaignInfo) HasReviewerType() bool {
	if o != nil && !IsNil(o.ReviewerType) {
		return true
	}

	return false
}

// SetReviewerType gets a reference to the given string and assigns it to the ReviewerType field.
func (o *CampaignAllOfMachineAccountCampaignInfo) SetReviewerType(v string) {
	o.ReviewerType = &v
}

func (o CampaignAllOfMachineAccountCampaignInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CampaignAllOfMachineAccountCampaignInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SourceIds) {
		toSerialize["sourceIds"] = o.SourceIds
	}
	if !IsNil(o.ReviewerType) {
		toSerialize["reviewerType"] = o.ReviewerType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CampaignAllOfMachineAccountCampaignInfo) UnmarshalJSON(data []byte) (err error) {
	varCampaignAllOfMachineAccountCampaignInfo := _CampaignAllOfMachineAccountCampaignInfo{}

	err = json.Unmarshal(data, &varCampaignAllOfMachineAccountCampaignInfo)

	if err != nil {
		return err
	}

	*o = CampaignAllOfMachineAccountCampaignInfo(varCampaignAllOfMachineAccountCampaignInfo)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "sourceIds")
		delete(additionalProperties, "reviewerType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCampaignAllOfMachineAccountCampaignInfo struct {
	value *CampaignAllOfMachineAccountCampaignInfo
	isSet bool
}

func (v NullableCampaignAllOfMachineAccountCampaignInfo) Get() *CampaignAllOfMachineAccountCampaignInfo {
	return v.value
}

func (v *NullableCampaignAllOfMachineAccountCampaignInfo) Set(val *CampaignAllOfMachineAccountCampaignInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignAllOfMachineAccountCampaignInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignAllOfMachineAccountCampaignInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignAllOfMachineAccountCampaignInfo(val *CampaignAllOfMachineAccountCampaignInfo) *NullableCampaignAllOfMachineAccountCampaignInfo {
	return &NullableCampaignAllOfMachineAccountCampaignInfo{value: val, isSet: true}
}

func (v NullableCampaignAllOfMachineAccountCampaignInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignAllOfMachineAccountCampaignInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


