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

// checks if the AccountAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountAction{}

// AccountAction Object for specifying Actions to be performed on a specified list of sources' account.
type AccountAction struct {
	// Describes if action will be enabled or disabled
	Action *string `json:"action,omitempty"`
	// A unique list of specific source IDs to apply the action to. The sources must have the ENABLE feature or flat file source. Required if allSources is not true. Must not be provided if allSources is true. Cannot be used together with excludeSourceIds See \"/sources\" endpoint for source features.
	SourceIds []string `json:"sourceIds,omitempty"`
	// A list of source IDs to exclude from the action. Cannot be used together with sourceIds.
	ExcludeSourceIds []string `json:"excludeSourceIds,omitempty"`
	// If true, the action applies to all available sources. If true, sourceIds must not be provided. If false or not set, sourceIds is required.
	AllSources *bool `json:"allSources,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountAction AccountAction

// NewAccountAction instantiates a new AccountAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountAction() *AccountAction {
	this := AccountAction{}
	var allSources bool = false
	this.AllSources = &allSources
	return &this
}

// NewAccountActionWithDefaults instantiates a new AccountAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountActionWithDefaults() *AccountAction {
	this := AccountAction{}
	var allSources bool = false
	this.AllSources = &allSources
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *AccountAction) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountAction) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *AccountAction) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *AccountAction) SetAction(v string) {
	o.Action = &v
}

// GetSourceIds returns the SourceIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountAction) GetSourceIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SourceIds
}

// GetSourceIdsOk returns a tuple with the SourceIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountAction) GetSourceIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.SourceIds) {
		return nil, false
	}
	return o.SourceIds, true
}

// HasSourceIds returns a boolean if a field has been set.
func (o *AccountAction) HasSourceIds() bool {
	if o != nil && !IsNil(o.SourceIds) {
		return true
	}

	return false
}

// SetSourceIds gets a reference to the given []string and assigns it to the SourceIds field.
func (o *AccountAction) SetSourceIds(v []string) {
	o.SourceIds = v
}

// GetExcludeSourceIds returns the ExcludeSourceIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountAction) GetExcludeSourceIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExcludeSourceIds
}

// GetExcludeSourceIdsOk returns a tuple with the ExcludeSourceIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountAction) GetExcludeSourceIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ExcludeSourceIds) {
		return nil, false
	}
	return o.ExcludeSourceIds, true
}

// HasExcludeSourceIds returns a boolean if a field has been set.
func (o *AccountAction) HasExcludeSourceIds() bool {
	if o != nil && !IsNil(o.ExcludeSourceIds) {
		return true
	}

	return false
}

// SetExcludeSourceIds gets a reference to the given []string and assigns it to the ExcludeSourceIds field.
func (o *AccountAction) SetExcludeSourceIds(v []string) {
	o.ExcludeSourceIds = v
}

// GetAllSources returns the AllSources field value if set, zero value otherwise.
func (o *AccountAction) GetAllSources() bool {
	if o == nil || IsNil(o.AllSources) {
		var ret bool
		return ret
	}
	return *o.AllSources
}

// GetAllSourcesOk returns a tuple with the AllSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountAction) GetAllSourcesOk() (*bool, bool) {
	if o == nil || IsNil(o.AllSources) {
		return nil, false
	}
	return o.AllSources, true
}

// HasAllSources returns a boolean if a field has been set.
func (o *AccountAction) HasAllSources() bool {
	if o != nil && !IsNil(o.AllSources) {
		return true
	}

	return false
}

// SetAllSources gets a reference to the given bool and assigns it to the AllSources field.
func (o *AccountAction) SetAllSources(v bool) {
	o.AllSources = &v
}

func (o AccountAction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if o.SourceIds != nil {
		toSerialize["sourceIds"] = o.SourceIds
	}
	if o.ExcludeSourceIds != nil {
		toSerialize["excludeSourceIds"] = o.ExcludeSourceIds
	}
	if !IsNil(o.AllSources) {
		toSerialize["allSources"] = o.AllSources
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountAction) UnmarshalJSON(data []byte) (err error) {
	varAccountAction := _AccountAction{}

	err = json.Unmarshal(data, &varAccountAction)

	if err != nil {
		return err
	}

	*o = AccountAction(varAccountAction)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "action")
		delete(additionalProperties, "sourceIds")
		delete(additionalProperties, "excludeSourceIds")
		delete(additionalProperties, "allSources")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountAction struct {
	value *AccountAction
	isSet bool
}

func (v NullableAccountAction) Get() *AccountAction {
	return v.value
}

func (v *NullableAccountAction) Set(val *AccountAction) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountAction) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountAction(val *AccountAction) *NullableAccountAction {
	return &NullableAccountAction{value: val, isSet: true}
}

func (v NullableAccountAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


