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

// checks if the MetricResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetricResponse{}

// MetricResponse struct for MetricResponse
type MetricResponse struct {
	// the name of metric
	Name *string `json:"name,omitempty"`
	// the value associated to the metric
	Value *float32 `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MetricResponse MetricResponse

// NewMetricResponse instantiates a new MetricResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricResponse() *MetricResponse {
	this := MetricResponse{}
	return &this
}

// NewMetricResponseWithDefaults instantiates a new MetricResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricResponseWithDefaults() *MetricResponse {
	this := MetricResponse{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MetricResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MetricResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MetricResponse) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *MetricResponse) GetValue() float32 {
	if o == nil || IsNil(o.Value) {
		var ret float32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricResponse) GetValueOk() (*float32, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *MetricResponse) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given float32 and assigns it to the Value field.
func (o *MetricResponse) SetValue(v float32) {
	o.Value = &v
}

func (o MetricResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetricResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MetricResponse) UnmarshalJSON(data []byte) (err error) {
	varMetricResponse := _MetricResponse{}

	err = json.Unmarshal(data, &varMetricResponse)

	if err != nil {
		return err
	}

	*o = MetricResponse(varMetricResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMetricResponse struct {
	value *MetricResponse
	isSet bool
}

func (v NullableMetricResponse) Get() *MetricResponse {
	return v.value
}

func (v *NullableMetricResponse) Set(val *MetricResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricResponse(val *MetricResponse) *NullableMetricResponse {
	return &NullableMetricResponse{value: val, isSet: true}
}

func (v NullableMetricResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


