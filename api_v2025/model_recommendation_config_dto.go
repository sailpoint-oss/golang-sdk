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

// checks if the RecommendationConfigDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecommendationConfigDto{}

// RecommendationConfigDto struct for RecommendationConfigDto
type RecommendationConfigDto struct {
	// List of identity attributes to use for calculating certification recommendations
	RecommenderFeatures []string `json:"recommenderFeatures,omitempty"`
	// The percent value that the recommendation calculation must surpass to produce a YES recommendation
	PeerGroupPercentageThreshold *float32 `json:"peerGroupPercentageThreshold,omitempty"`
	// If true, rulesRecommenderConfig will be refreshed with new programatically selected attribute and threshold values on the next pipeline run
	RunAutoSelectOnce *bool `json:"runAutoSelectOnce,omitempty"`
	// If true, rulesRecommenderConfig will be refreshed with new programatically selected threshold values on the next pipeline run
	OnlyTuneThreshold *bool `json:"onlyTuneThreshold,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RecommendationConfigDto RecommendationConfigDto

// NewRecommendationConfigDto instantiates a new RecommendationConfigDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecommendationConfigDto() *RecommendationConfigDto {
	this := RecommendationConfigDto{}
	var runAutoSelectOnce bool = false
	this.RunAutoSelectOnce = &runAutoSelectOnce
	var onlyTuneThreshold bool = false
	this.OnlyTuneThreshold = &onlyTuneThreshold
	return &this
}

// NewRecommendationConfigDtoWithDefaults instantiates a new RecommendationConfigDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecommendationConfigDtoWithDefaults() *RecommendationConfigDto {
	this := RecommendationConfigDto{}
	var runAutoSelectOnce bool = false
	this.RunAutoSelectOnce = &runAutoSelectOnce
	var onlyTuneThreshold bool = false
	this.OnlyTuneThreshold = &onlyTuneThreshold
	return &this
}

// GetRecommenderFeatures returns the RecommenderFeatures field value if set, zero value otherwise.
func (o *RecommendationConfigDto) GetRecommenderFeatures() []string {
	if o == nil || IsNil(o.RecommenderFeatures) {
		var ret []string
		return ret
	}
	return o.RecommenderFeatures
}

// GetRecommenderFeaturesOk returns a tuple with the RecommenderFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationConfigDto) GetRecommenderFeaturesOk() ([]string, bool) {
	if o == nil || IsNil(o.RecommenderFeatures) {
		return nil, false
	}
	return o.RecommenderFeatures, true
}

// HasRecommenderFeatures returns a boolean if a field has been set.
func (o *RecommendationConfigDto) HasRecommenderFeatures() bool {
	if o != nil && !IsNil(o.RecommenderFeatures) {
		return true
	}

	return false
}

// SetRecommenderFeatures gets a reference to the given []string and assigns it to the RecommenderFeatures field.
func (o *RecommendationConfigDto) SetRecommenderFeatures(v []string) {
	o.RecommenderFeatures = v
}

// GetPeerGroupPercentageThreshold returns the PeerGroupPercentageThreshold field value if set, zero value otherwise.
func (o *RecommendationConfigDto) GetPeerGroupPercentageThreshold() float32 {
	if o == nil || IsNil(o.PeerGroupPercentageThreshold) {
		var ret float32
		return ret
	}
	return *o.PeerGroupPercentageThreshold
}

// GetPeerGroupPercentageThresholdOk returns a tuple with the PeerGroupPercentageThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationConfigDto) GetPeerGroupPercentageThresholdOk() (*float32, bool) {
	if o == nil || IsNil(o.PeerGroupPercentageThreshold) {
		return nil, false
	}
	return o.PeerGroupPercentageThreshold, true
}

// HasPeerGroupPercentageThreshold returns a boolean if a field has been set.
func (o *RecommendationConfigDto) HasPeerGroupPercentageThreshold() bool {
	if o != nil && !IsNil(o.PeerGroupPercentageThreshold) {
		return true
	}

	return false
}

// SetPeerGroupPercentageThreshold gets a reference to the given float32 and assigns it to the PeerGroupPercentageThreshold field.
func (o *RecommendationConfigDto) SetPeerGroupPercentageThreshold(v float32) {
	o.PeerGroupPercentageThreshold = &v
}

// GetRunAutoSelectOnce returns the RunAutoSelectOnce field value if set, zero value otherwise.
func (o *RecommendationConfigDto) GetRunAutoSelectOnce() bool {
	if o == nil || IsNil(o.RunAutoSelectOnce) {
		var ret bool
		return ret
	}
	return *o.RunAutoSelectOnce
}

// GetRunAutoSelectOnceOk returns a tuple with the RunAutoSelectOnce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationConfigDto) GetRunAutoSelectOnceOk() (*bool, bool) {
	if o == nil || IsNil(o.RunAutoSelectOnce) {
		return nil, false
	}
	return o.RunAutoSelectOnce, true
}

// HasRunAutoSelectOnce returns a boolean if a field has been set.
func (o *RecommendationConfigDto) HasRunAutoSelectOnce() bool {
	if o != nil && !IsNil(o.RunAutoSelectOnce) {
		return true
	}

	return false
}

// SetRunAutoSelectOnce gets a reference to the given bool and assigns it to the RunAutoSelectOnce field.
func (o *RecommendationConfigDto) SetRunAutoSelectOnce(v bool) {
	o.RunAutoSelectOnce = &v
}

// GetOnlyTuneThreshold returns the OnlyTuneThreshold field value if set, zero value otherwise.
func (o *RecommendationConfigDto) GetOnlyTuneThreshold() bool {
	if o == nil || IsNil(o.OnlyTuneThreshold) {
		var ret bool
		return ret
	}
	return *o.OnlyTuneThreshold
}

// GetOnlyTuneThresholdOk returns a tuple with the OnlyTuneThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationConfigDto) GetOnlyTuneThresholdOk() (*bool, bool) {
	if o == nil || IsNil(o.OnlyTuneThreshold) {
		return nil, false
	}
	return o.OnlyTuneThreshold, true
}

// HasOnlyTuneThreshold returns a boolean if a field has been set.
func (o *RecommendationConfigDto) HasOnlyTuneThreshold() bool {
	if o != nil && !IsNil(o.OnlyTuneThreshold) {
		return true
	}

	return false
}

// SetOnlyTuneThreshold gets a reference to the given bool and assigns it to the OnlyTuneThreshold field.
func (o *RecommendationConfigDto) SetOnlyTuneThreshold(v bool) {
	o.OnlyTuneThreshold = &v
}

func (o RecommendationConfigDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecommendationConfigDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RecommenderFeatures) {
		toSerialize["recommenderFeatures"] = o.RecommenderFeatures
	}
	if !IsNil(o.PeerGroupPercentageThreshold) {
		toSerialize["peerGroupPercentageThreshold"] = o.PeerGroupPercentageThreshold
	}
	if !IsNil(o.RunAutoSelectOnce) {
		toSerialize["runAutoSelectOnce"] = o.RunAutoSelectOnce
	}
	if !IsNil(o.OnlyTuneThreshold) {
		toSerialize["onlyTuneThreshold"] = o.OnlyTuneThreshold
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RecommendationConfigDto) UnmarshalJSON(data []byte) (err error) {
	varRecommendationConfigDto := _RecommendationConfigDto{}

	err = json.Unmarshal(data, &varRecommendationConfigDto)

	if err != nil {
		return err
	}

	*o = RecommendationConfigDto(varRecommendationConfigDto)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "recommenderFeatures")
		delete(additionalProperties, "peerGroupPercentageThreshold")
		delete(additionalProperties, "runAutoSelectOnce")
		delete(additionalProperties, "onlyTuneThreshold")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRecommendationConfigDto struct {
	value *RecommendationConfigDto
	isSet bool
}

func (v NullableRecommendationConfigDto) Get() *RecommendationConfigDto {
	return v.value
}

func (v *NullableRecommendationConfigDto) Set(val *RecommendationConfigDto) {
	v.value = val
	v.isSet = true
}

func (v NullableRecommendationConfigDto) IsSet() bool {
	return v.isSet
}

func (v *NullableRecommendationConfigDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecommendationConfigDto(val *RecommendationConfigDto) *NullableRecommendationConfigDto {
	return &NullableRecommendationConfigDto{value: val, isSet: true}
}

func (v NullableRecommendationConfigDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecommendationConfigDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


