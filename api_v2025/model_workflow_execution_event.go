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

// checks if the WorkflowExecutionEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowExecutionEvent{}

// WorkflowExecutionEvent struct for WorkflowExecutionEvent
type WorkflowExecutionEvent struct {
	// The type of event
	Type *string `json:"type,omitempty"`
	// The date-time when the event occurred
	Timestamp *SailPointTime `json:"timestamp,omitempty"`
	// Additional attributes associated with the event
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowExecutionEvent WorkflowExecutionEvent

// NewWorkflowExecutionEvent instantiates a new WorkflowExecutionEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowExecutionEvent() *WorkflowExecutionEvent {
	this := WorkflowExecutionEvent{}
	return &this
}

// NewWorkflowExecutionEventWithDefaults instantiates a new WorkflowExecutionEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowExecutionEventWithDefaults() *WorkflowExecutionEvent {
	this := WorkflowExecutionEvent{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WorkflowExecutionEvent) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowExecutionEvent) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WorkflowExecutionEvent) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *WorkflowExecutionEvent) SetType(v string) {
	o.Type = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *WorkflowExecutionEvent) GetTimestamp() SailPointTime {
	if o == nil || IsNil(o.Timestamp) {
		var ret SailPointTime
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowExecutionEvent) GetTimestampOk() (*SailPointTime, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *WorkflowExecutionEvent) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given SailPointTime and assigns it to the Timestamp field.
func (o *WorkflowExecutionEvent) SetTimestamp(v SailPointTime) {
	o.Timestamp = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *WorkflowExecutionEvent) GetAttributes() map[string]interface{} {
	if o == nil || IsNil(o.Attributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowExecutionEvent) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Attributes) {
		return map[string]interface{}{}, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *WorkflowExecutionEvent) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *WorkflowExecutionEvent) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

func (o WorkflowExecutionEvent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowExecutionEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WorkflowExecutionEvent) UnmarshalJSON(data []byte) (err error) {
	varWorkflowExecutionEvent := _WorkflowExecutionEvent{}

	err = json.Unmarshal(data, &varWorkflowExecutionEvent)

	if err != nil {
		return err
	}

	*o = WorkflowExecutionEvent(varWorkflowExecutionEvent)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "timestamp")
		delete(additionalProperties, "attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowExecutionEvent struct {
	value *WorkflowExecutionEvent
	isSet bool
}

func (v NullableWorkflowExecutionEvent) Get() *WorkflowExecutionEvent {
	return v.value
}

func (v *NullableWorkflowExecutionEvent) Set(val *WorkflowExecutionEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowExecutionEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowExecutionEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowExecutionEvent(val *WorkflowExecutionEvent) *NullableWorkflowExecutionEvent {
	return &NullableWorkflowExecutionEvent{value: val, isSet: true}
}

func (v NullableWorkflowExecutionEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowExecutionEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


