/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sailpointsdk

import (
	"encoding/json"
)

// DateFormat struct for DateFormat
type DateFormat struct {
	InputFormat *DateFormatInputFormat `json:"inputFormat,omitempty"`
	OutputFormat *DateFormatOutputFormat `json:"outputFormat,omitempty"`
	// A value that indicates whether the transform logic should be re-evaluated every evening as part of the identity refresh process
	RequiresPeriodicRefresh *bool `json:"requiresPeriodicRefresh,omitempty"`
	// This is an optional attribute that can explicitly define the input data which will be fed into the transform logic. If input is not provided, the transform will take its input from the source and attribute combination configured via the UI.
	Input map[string]interface{} `json:"input,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DateFormat DateFormat

// NewDateFormat instantiates a new DateFormat object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDateFormat() *DateFormat {
	this := DateFormat{}
	var requiresPeriodicRefresh bool = false
	this.RequiresPeriodicRefresh = &requiresPeriodicRefresh
	return &this
}

// NewDateFormatWithDefaults instantiates a new DateFormat object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDateFormatWithDefaults() *DateFormat {
	this := DateFormat{}
	var requiresPeriodicRefresh bool = false
	this.RequiresPeriodicRefresh = &requiresPeriodicRefresh
	return &this
}

// GetInputFormat returns the InputFormat field value if set, zero value otherwise.
func (o *DateFormat) GetInputFormat() DateFormatInputFormat {
	if o == nil || isNil(o.InputFormat) {
		var ret DateFormatInputFormat
		return ret
	}
	return *o.InputFormat
}

// GetInputFormatOk returns a tuple with the InputFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateFormat) GetInputFormatOk() (*DateFormatInputFormat, bool) {
	if o == nil || isNil(o.InputFormat) {
		return nil, false
	}
	return o.InputFormat, true
}

// HasInputFormat returns a boolean if a field has been set.
func (o *DateFormat) HasInputFormat() bool {
	if o != nil && !isNil(o.InputFormat) {
		return true
	}

	return false
}

// SetInputFormat gets a reference to the given DateFormatInputFormat and assigns it to the InputFormat field.
func (o *DateFormat) SetInputFormat(v DateFormatInputFormat) {
	o.InputFormat = &v
}

// GetOutputFormat returns the OutputFormat field value if set, zero value otherwise.
func (o *DateFormat) GetOutputFormat() DateFormatOutputFormat {
	if o == nil || isNil(o.OutputFormat) {
		var ret DateFormatOutputFormat
		return ret
	}
	return *o.OutputFormat
}

// GetOutputFormatOk returns a tuple with the OutputFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateFormat) GetOutputFormatOk() (*DateFormatOutputFormat, bool) {
	if o == nil || isNil(o.OutputFormat) {
		return nil, false
	}
	return o.OutputFormat, true
}

// HasOutputFormat returns a boolean if a field has been set.
func (o *DateFormat) HasOutputFormat() bool {
	if o != nil && !isNil(o.OutputFormat) {
		return true
	}

	return false
}

// SetOutputFormat gets a reference to the given DateFormatOutputFormat and assigns it to the OutputFormat field.
func (o *DateFormat) SetOutputFormat(v DateFormatOutputFormat) {
	o.OutputFormat = &v
}

// GetRequiresPeriodicRefresh returns the RequiresPeriodicRefresh field value if set, zero value otherwise.
func (o *DateFormat) GetRequiresPeriodicRefresh() bool {
	if o == nil || isNil(o.RequiresPeriodicRefresh) {
		var ret bool
		return ret
	}
	return *o.RequiresPeriodicRefresh
}

// GetRequiresPeriodicRefreshOk returns a tuple with the RequiresPeriodicRefresh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateFormat) GetRequiresPeriodicRefreshOk() (*bool, bool) {
	if o == nil || isNil(o.RequiresPeriodicRefresh) {
		return nil, false
	}
	return o.RequiresPeriodicRefresh, true
}

// HasRequiresPeriodicRefresh returns a boolean if a field has been set.
func (o *DateFormat) HasRequiresPeriodicRefresh() bool {
	if o != nil && !isNil(o.RequiresPeriodicRefresh) {
		return true
	}

	return false
}

// SetRequiresPeriodicRefresh gets a reference to the given bool and assigns it to the RequiresPeriodicRefresh field.
func (o *DateFormat) SetRequiresPeriodicRefresh(v bool) {
	o.RequiresPeriodicRefresh = &v
}

// GetInput returns the Input field value if set, zero value otherwise.
func (o *DateFormat) GetInput() map[string]interface{} {
	if o == nil || isNil(o.Input) {
		var ret map[string]interface{}
		return ret
	}
	return o.Input
}

// GetInputOk returns a tuple with the Input field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateFormat) GetInputOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Input) {
		return map[string]interface{}{}, false
	}
	return o.Input, true
}

// HasInput returns a boolean if a field has been set.
func (o *DateFormat) HasInput() bool {
	if o != nil && !isNil(o.Input) {
		return true
	}

	return false
}

// SetInput gets a reference to the given map[string]interface{} and assigns it to the Input field.
func (o *DateFormat) SetInput(v map[string]interface{}) {
	o.Input = v
}

func (o DateFormat) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.InputFormat) {
		toSerialize["inputFormat"] = o.InputFormat
	}
	if !isNil(o.OutputFormat) {
		toSerialize["outputFormat"] = o.OutputFormat
	}
	if !isNil(o.RequiresPeriodicRefresh) {
		toSerialize["requiresPeriodicRefresh"] = o.RequiresPeriodicRefresh
	}
	if !isNil(o.Input) {
		toSerialize["input"] = o.Input
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DateFormat) UnmarshalJSON(bytes []byte) (err error) {
	varDateFormat := _DateFormat{}

	if err = json.Unmarshal(bytes, &varDateFormat); err == nil {
		*o = DateFormat(varDateFormat)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "inputFormat")
		delete(additionalProperties, "outputFormat")
		delete(additionalProperties, "requiresPeriodicRefresh")
		delete(additionalProperties, "input")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDateFormat struct {
	value *DateFormat
	isSet bool
}

func (v NullableDateFormat) Get() *DateFormat {
	return v.value
}

func (v *NullableDateFormat) Set(val *DateFormat) {
	v.value = val
	v.isSet = true
}

func (v NullableDateFormat) IsSet() bool {
	return v.isSet
}

func (v *NullableDateFormat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDateFormat(val *DateFormat) *NullableDateFormat {
	return &NullableDateFormat{value: val, isSet: true}
}

func (v NullableDateFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDateFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


