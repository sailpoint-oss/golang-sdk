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

// checks if the AccessRequestPhases type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessRequestPhases{}

// AccessRequestPhases Provides additional details about this access request phase.
type AccessRequestPhases struct {
	// The time that this phase started.
	Started *SailPointTime `json:"started,omitempty"`
	// The time that this phase finished.
	Finished NullableTime `json:"finished,omitempty"`
	// The name of this phase.
	Name *string `json:"name,omitempty"`
	// The state of this phase.
	State *string `json:"state,omitempty"`
	// The state of this phase.
	Result NullableString `json:"result,omitempty"`
	// A reference to another object on the RequestedItemStatus that contains more details about the phase. Note that for the Provisioning phase, this will be empty if there are no manual work items.
	PhaseReference NullableString `json:"phaseReference,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccessRequestPhases AccessRequestPhases

// NewAccessRequestPhases instantiates a new AccessRequestPhases object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessRequestPhases() *AccessRequestPhases {
	this := AccessRequestPhases{}
	return &this
}

// NewAccessRequestPhasesWithDefaults instantiates a new AccessRequestPhases object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessRequestPhasesWithDefaults() *AccessRequestPhases {
	this := AccessRequestPhases{}
	return &this
}

// GetStarted returns the Started field value if set, zero value otherwise.
func (o *AccessRequestPhases) GetStarted() SailPointTime {
	if o == nil || IsNil(o.Started) {
		var ret SailPointTime
		return ret
	}
	return *o.Started
}

// GetStartedOk returns a tuple with the Started field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequestPhases) GetStartedOk() (*SailPointTime, bool) {
	if o == nil || IsNil(o.Started) {
		return nil, false
	}
	return o.Started, true
}

// HasStarted returns a boolean if a field has been set.
func (o *AccessRequestPhases) HasStarted() bool {
	if o != nil && !IsNil(o.Started) {
		return true
	}

	return false
}

// SetStarted gets a reference to the given SailPointTime and assigns it to the Started field.
func (o *AccessRequestPhases) SetStarted(v SailPointTime) {
	o.Started = &v
}

// GetFinished returns the Finished field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessRequestPhases) GetFinished() SailPointTime {
	if o == nil || IsNil(o.Finished.Get()) {
		var ret SailPointTime
		return ret
	}
	return *o.Finished.Get()
}

// GetFinishedOk returns a tuple with the Finished field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessRequestPhases) GetFinishedOk() (*SailPointTime, bool) {
	if o == nil {
		return nil, false
	}
	return o.Finished.Get(), o.Finished.IsSet()
}

// HasFinished returns a boolean if a field has been set.
func (o *AccessRequestPhases) HasFinished() bool {
	if o != nil && o.Finished.IsSet() {
		return true
	}

	return false
}

// SetFinished gets a reference to the given NullableTime and assigns it to the Finished field.
func (o *AccessRequestPhases) SetFinished(v SailPointTime) {
	o.Finished.Set(&v)
}
// SetFinishedNil sets the value for Finished to be an explicit nil
func (o *AccessRequestPhases) SetFinishedNil() {
	o.Finished.Set(nil)
}

// UnsetFinished ensures that no value is present for Finished, not even an explicit nil
func (o *AccessRequestPhases) UnsetFinished() {
	o.Finished.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AccessRequestPhases) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequestPhases) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AccessRequestPhases) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AccessRequestPhases) SetName(v string) {
	o.Name = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *AccessRequestPhases) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequestPhases) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *AccessRequestPhases) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *AccessRequestPhases) SetState(v string) {
	o.State = &v
}

// GetResult returns the Result field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessRequestPhases) GetResult() string {
	if o == nil || IsNil(o.Result.Get()) {
		var ret string
		return ret
	}
	return *o.Result.Get()
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessRequestPhases) GetResultOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Result.Get(), o.Result.IsSet()
}

// HasResult returns a boolean if a field has been set.
func (o *AccessRequestPhases) HasResult() bool {
	if o != nil && o.Result.IsSet() {
		return true
	}

	return false
}

// SetResult gets a reference to the given NullableString and assigns it to the Result field.
func (o *AccessRequestPhases) SetResult(v string) {
	o.Result.Set(&v)
}
// SetResultNil sets the value for Result to be an explicit nil
func (o *AccessRequestPhases) SetResultNil() {
	o.Result.Set(nil)
}

// UnsetResult ensures that no value is present for Result, not even an explicit nil
func (o *AccessRequestPhases) UnsetResult() {
	o.Result.Unset()
}

// GetPhaseReference returns the PhaseReference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessRequestPhases) GetPhaseReference() string {
	if o == nil || IsNil(o.PhaseReference.Get()) {
		var ret string
		return ret
	}
	return *o.PhaseReference.Get()
}

// GetPhaseReferenceOk returns a tuple with the PhaseReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessRequestPhases) GetPhaseReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PhaseReference.Get(), o.PhaseReference.IsSet()
}

// HasPhaseReference returns a boolean if a field has been set.
func (o *AccessRequestPhases) HasPhaseReference() bool {
	if o != nil && o.PhaseReference.IsSet() {
		return true
	}

	return false
}

// SetPhaseReference gets a reference to the given NullableString and assigns it to the PhaseReference field.
func (o *AccessRequestPhases) SetPhaseReference(v string) {
	o.PhaseReference.Set(&v)
}
// SetPhaseReferenceNil sets the value for PhaseReference to be an explicit nil
func (o *AccessRequestPhases) SetPhaseReferenceNil() {
	o.PhaseReference.Set(nil)
}

// UnsetPhaseReference ensures that no value is present for PhaseReference, not even an explicit nil
func (o *AccessRequestPhases) UnsetPhaseReference() {
	o.PhaseReference.Unset()
}

func (o AccessRequestPhases) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessRequestPhases) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Started) {
		toSerialize["started"] = o.Started
	}
	if o.Finished.IsSet() {
		toSerialize["finished"] = o.Finished.Get()
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if o.Result.IsSet() {
		toSerialize["result"] = o.Result.Get()
	}
	if o.PhaseReference.IsSet() {
		toSerialize["phaseReference"] = o.PhaseReference.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccessRequestPhases) UnmarshalJSON(data []byte) (err error) {
	varAccessRequestPhases := _AccessRequestPhases{}

	err = json.Unmarshal(data, &varAccessRequestPhases)

	if err != nil {
		return err
	}

	*o = AccessRequestPhases(varAccessRequestPhases)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "started")
		delete(additionalProperties, "finished")
		delete(additionalProperties, "name")
		delete(additionalProperties, "state")
		delete(additionalProperties, "result")
		delete(additionalProperties, "phaseReference")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccessRequestPhases struct {
	value *AccessRequestPhases
	isSet bool
}

func (v NullableAccessRequestPhases) Get() *AccessRequestPhases {
	return v.value
}

func (v *NullableAccessRequestPhases) Set(val *AccessRequestPhases) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessRequestPhases) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessRequestPhases) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessRequestPhases(val *AccessRequestPhases) *NullableAccessRequestPhases {
	return &NullableAccessRequestPhases{value: val, isSet: true}
}

func (v NullableAccessRequestPhases) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessRequestPhases) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


