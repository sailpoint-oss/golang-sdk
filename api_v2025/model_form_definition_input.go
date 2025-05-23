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

// checks if the FormDefinitionInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FormDefinitionInput{}

// FormDefinitionInput struct for FormDefinitionInput
type FormDefinitionInput struct {
	// Unique identifier for the form input.
	Id *string `json:"id,omitempty"`
	// FormDefinitionInputType value. STRING FormDefinitionInputTypeString
	Type *string `json:"type,omitempty"`
	// Name for the form input.
	Label *string `json:"label,omitempty"`
	// Form input's description.
	Description *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FormDefinitionInput FormDefinitionInput

// NewFormDefinitionInput instantiates a new FormDefinitionInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormDefinitionInput() *FormDefinitionInput {
	this := FormDefinitionInput{}
	return &this
}

// NewFormDefinitionInputWithDefaults instantiates a new FormDefinitionInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormDefinitionInputWithDefaults() *FormDefinitionInput {
	this := FormDefinitionInput{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FormDefinitionInput) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDefinitionInput) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FormDefinitionInput) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FormDefinitionInput) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FormDefinitionInput) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDefinitionInput) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FormDefinitionInput) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *FormDefinitionInput) SetType(v string) {
	o.Type = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *FormDefinitionInput) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDefinitionInput) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *FormDefinitionInput) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *FormDefinitionInput) SetLabel(v string) {
	o.Label = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FormDefinitionInput) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDefinitionInput) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FormDefinitionInput) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FormDefinitionInput) SetDescription(v string) {
	o.Description = &v
}

func (o FormDefinitionInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FormDefinitionInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FormDefinitionInput) UnmarshalJSON(data []byte) (err error) {
	varFormDefinitionInput := _FormDefinitionInput{}

	err = json.Unmarshal(data, &varFormDefinitionInput)

	if err != nil {
		return err
	}

	*o = FormDefinitionInput(varFormDefinitionInput)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "label")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFormDefinitionInput struct {
	value *FormDefinitionInput
	isSet bool
}

func (v NullableFormDefinitionInput) Get() *FormDefinitionInput {
	return v.value
}

func (v *NullableFormDefinitionInput) Set(val *FormDefinitionInput) {
	v.value = val
	v.isSet = true
}

func (v NullableFormDefinitionInput) IsSet() bool {
	return v.isSet
}

func (v *NullableFormDefinitionInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormDefinitionInput(val *FormDefinitionInput) *NullableFormDefinitionInput {
	return &NullableFormDefinitionInput{value: val, isSet: true}
}

func (v NullableFormDefinitionInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormDefinitionInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


