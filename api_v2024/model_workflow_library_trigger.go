/*
Identity Security Cloud V2024 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: v2024
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2024

import (
	"encoding/json"
	"time"
)

// checks if the WorkflowLibraryTrigger type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowLibraryTrigger{}

// WorkflowLibraryTrigger struct for WorkflowLibraryTrigger
type WorkflowLibraryTrigger struct {
	// Trigger ID. This is a static namespaced ID for the trigger.
	Id *string `json:"id,omitempty"`
	// Trigger type
	Type *string `json:"type,omitempty"`
	Deprecated *bool `json:"deprecated,omitempty"`
	DeprecatedBy *time.Time `json:"deprecatedBy,omitempty"`
	IsSimulationEnabled *bool `json:"isSimulationEnabled,omitempty"`
	// Example output schema
	OutputSchema map[string]interface{} `json:"outputSchema,omitempty"`
	// Trigger Name
	Name *string `json:"name,omitempty"`
	// Trigger Description
	Description *string `json:"description,omitempty"`
	// Determines whether the dynamic output schema is returned in place of the action's output schema. The dynamic schema lists non-static properties, like properties of a workflow form where each form has different fields. These will be provided dynamically based on available form fields.
	IsDynamicSchema *bool `json:"isDynamicSchema,omitempty"`
	// Example trigger payload if applicable
	InputExample map[string]interface{} `json:"inputExample,omitempty"`
	// One or more inputs that the trigger accepts
	FormFields []WorkflowLibraryFormFields `json:"formFields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowLibraryTrigger WorkflowLibraryTrigger

// NewWorkflowLibraryTrigger instantiates a new WorkflowLibraryTrigger object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowLibraryTrigger() *WorkflowLibraryTrigger {
	this := WorkflowLibraryTrigger{}
	var isDynamicSchema bool = false
	this.IsDynamicSchema = &isDynamicSchema
	return &this
}

// NewWorkflowLibraryTriggerWithDefaults instantiates a new WorkflowLibraryTrigger object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowLibraryTriggerWithDefaults() *WorkflowLibraryTrigger {
	this := WorkflowLibraryTrigger{}
	var isDynamicSchema bool = false
	this.IsDynamicSchema = &isDynamicSchema
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WorkflowLibraryTrigger) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowLibraryTrigger) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WorkflowLibraryTrigger) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WorkflowLibraryTrigger) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WorkflowLibraryTrigger) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowLibraryTrigger) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WorkflowLibraryTrigger) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *WorkflowLibraryTrigger) SetType(v string) {
	o.Type = &v
}

// GetDeprecated returns the Deprecated field value if set, zero value otherwise.
func (o *WorkflowLibraryTrigger) GetDeprecated() bool {
	if o == nil || IsNil(o.Deprecated) {
		var ret bool
		return ret
	}
	return *o.Deprecated
}

// GetDeprecatedOk returns a tuple with the Deprecated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowLibraryTrigger) GetDeprecatedOk() (*bool, bool) {
	if o == nil || IsNil(o.Deprecated) {
		return nil, false
	}
	return o.Deprecated, true
}

// HasDeprecated returns a boolean if a field has been set.
func (o *WorkflowLibraryTrigger) HasDeprecated() bool {
	if o != nil && !IsNil(o.Deprecated) {
		return true
	}

	return false
}

// SetDeprecated gets a reference to the given bool and assigns it to the Deprecated field.
func (o *WorkflowLibraryTrigger) SetDeprecated(v bool) {
	o.Deprecated = &v
}

// GetDeprecatedBy returns the DeprecatedBy field value if set, zero value otherwise.
func (o *WorkflowLibraryTrigger) GetDeprecatedBy() time.Time {
	if o == nil || IsNil(o.DeprecatedBy) {
		var ret time.Time
		return ret
	}
	return *o.DeprecatedBy
}

// GetDeprecatedByOk returns a tuple with the DeprecatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowLibraryTrigger) GetDeprecatedByOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DeprecatedBy) {
		return nil, false
	}
	return o.DeprecatedBy, true
}

// HasDeprecatedBy returns a boolean if a field has been set.
func (o *WorkflowLibraryTrigger) HasDeprecatedBy() bool {
	if o != nil && !IsNil(o.DeprecatedBy) {
		return true
	}

	return false
}

// SetDeprecatedBy gets a reference to the given time.Time and assigns it to the DeprecatedBy field.
func (o *WorkflowLibraryTrigger) SetDeprecatedBy(v time.Time) {
	o.DeprecatedBy = &v
}

// GetIsSimulationEnabled returns the IsSimulationEnabled field value if set, zero value otherwise.
func (o *WorkflowLibraryTrigger) GetIsSimulationEnabled() bool {
	if o == nil || IsNil(o.IsSimulationEnabled) {
		var ret bool
		return ret
	}
	return *o.IsSimulationEnabled
}

// GetIsSimulationEnabledOk returns a tuple with the IsSimulationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowLibraryTrigger) GetIsSimulationEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSimulationEnabled) {
		return nil, false
	}
	return o.IsSimulationEnabled, true
}

// HasIsSimulationEnabled returns a boolean if a field has been set.
func (o *WorkflowLibraryTrigger) HasIsSimulationEnabled() bool {
	if o != nil && !IsNil(o.IsSimulationEnabled) {
		return true
	}

	return false
}

// SetIsSimulationEnabled gets a reference to the given bool and assigns it to the IsSimulationEnabled field.
func (o *WorkflowLibraryTrigger) SetIsSimulationEnabled(v bool) {
	o.IsSimulationEnabled = &v
}

// GetOutputSchema returns the OutputSchema field value if set, zero value otherwise.
func (o *WorkflowLibraryTrigger) GetOutputSchema() map[string]interface{} {
	if o == nil || IsNil(o.OutputSchema) {
		var ret map[string]interface{}
		return ret
	}
	return o.OutputSchema
}

// GetOutputSchemaOk returns a tuple with the OutputSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowLibraryTrigger) GetOutputSchemaOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.OutputSchema) {
		return map[string]interface{}{}, false
	}
	return o.OutputSchema, true
}

// HasOutputSchema returns a boolean if a field has been set.
func (o *WorkflowLibraryTrigger) HasOutputSchema() bool {
	if o != nil && !IsNil(o.OutputSchema) {
		return true
	}

	return false
}

// SetOutputSchema gets a reference to the given map[string]interface{} and assigns it to the OutputSchema field.
func (o *WorkflowLibraryTrigger) SetOutputSchema(v map[string]interface{}) {
	o.OutputSchema = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowLibraryTrigger) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowLibraryTrigger) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowLibraryTrigger) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowLibraryTrigger) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkflowLibraryTrigger) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowLibraryTrigger) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkflowLibraryTrigger) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkflowLibraryTrigger) SetDescription(v string) {
	o.Description = &v
}

// GetIsDynamicSchema returns the IsDynamicSchema field value if set, zero value otherwise.
func (o *WorkflowLibraryTrigger) GetIsDynamicSchema() bool {
	if o == nil || IsNil(o.IsDynamicSchema) {
		var ret bool
		return ret
	}
	return *o.IsDynamicSchema
}

// GetIsDynamicSchemaOk returns a tuple with the IsDynamicSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowLibraryTrigger) GetIsDynamicSchemaOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDynamicSchema) {
		return nil, false
	}
	return o.IsDynamicSchema, true
}

// HasIsDynamicSchema returns a boolean if a field has been set.
func (o *WorkflowLibraryTrigger) HasIsDynamicSchema() bool {
	if o != nil && !IsNil(o.IsDynamicSchema) {
		return true
	}

	return false
}

// SetIsDynamicSchema gets a reference to the given bool and assigns it to the IsDynamicSchema field.
func (o *WorkflowLibraryTrigger) SetIsDynamicSchema(v bool) {
	o.IsDynamicSchema = &v
}

// GetInputExample returns the InputExample field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowLibraryTrigger) GetInputExample() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.InputExample
}

// GetInputExampleOk returns a tuple with the InputExample field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowLibraryTrigger) GetInputExampleOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.InputExample) {
		return map[string]interface{}{}, false
	}
	return o.InputExample, true
}

// HasInputExample returns a boolean if a field has been set.
func (o *WorkflowLibraryTrigger) HasInputExample() bool {
	if o != nil && !IsNil(o.InputExample) {
		return true
	}

	return false
}

// SetInputExample gets a reference to the given map[string]interface{} and assigns it to the InputExample field.
func (o *WorkflowLibraryTrigger) SetInputExample(v map[string]interface{}) {
	o.InputExample = v
}

// GetFormFields returns the FormFields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowLibraryTrigger) GetFormFields() []WorkflowLibraryFormFields {
	if o == nil {
		var ret []WorkflowLibraryFormFields
		return ret
	}
	return o.FormFields
}

// GetFormFieldsOk returns a tuple with the FormFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowLibraryTrigger) GetFormFieldsOk() ([]WorkflowLibraryFormFields, bool) {
	if o == nil || IsNil(o.FormFields) {
		return nil, false
	}
	return o.FormFields, true
}

// HasFormFields returns a boolean if a field has been set.
func (o *WorkflowLibraryTrigger) HasFormFields() bool {
	if o != nil && !IsNil(o.FormFields) {
		return true
	}

	return false
}

// SetFormFields gets a reference to the given []WorkflowLibraryFormFields and assigns it to the FormFields field.
func (o *WorkflowLibraryTrigger) SetFormFields(v []WorkflowLibraryFormFields) {
	o.FormFields = v
}

func (o WorkflowLibraryTrigger) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowLibraryTrigger) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Deprecated) {
		toSerialize["deprecated"] = o.Deprecated
	}
	if !IsNil(o.DeprecatedBy) {
		toSerialize["deprecatedBy"] = o.DeprecatedBy
	}
	if !IsNil(o.IsSimulationEnabled) {
		toSerialize["isSimulationEnabled"] = o.IsSimulationEnabled
	}
	if !IsNil(o.OutputSchema) {
		toSerialize["outputSchema"] = o.OutputSchema
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.IsDynamicSchema) {
		toSerialize["isDynamicSchema"] = o.IsDynamicSchema
	}
	if o.InputExample != nil {
		toSerialize["inputExample"] = o.InputExample
	}
	if o.FormFields != nil {
		toSerialize["formFields"] = o.FormFields
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WorkflowLibraryTrigger) UnmarshalJSON(data []byte) (err error) {
	varWorkflowLibraryTrigger := _WorkflowLibraryTrigger{}

	err = json.Unmarshal(data, &varWorkflowLibraryTrigger)

	if err != nil {
		return err
	}

	*o = WorkflowLibraryTrigger(varWorkflowLibraryTrigger)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "deprecated")
		delete(additionalProperties, "deprecatedBy")
		delete(additionalProperties, "isSimulationEnabled")
		delete(additionalProperties, "outputSchema")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "isDynamicSchema")
		delete(additionalProperties, "inputExample")
		delete(additionalProperties, "formFields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowLibraryTrigger struct {
	value *WorkflowLibraryTrigger
	isSet bool
}

func (v NullableWorkflowLibraryTrigger) Get() *WorkflowLibraryTrigger {
	return v.value
}

func (v *NullableWorkflowLibraryTrigger) Set(val *WorkflowLibraryTrigger) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowLibraryTrigger) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowLibraryTrigger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowLibraryTrigger(val *WorkflowLibraryTrigger) *NullableWorkflowLibraryTrigger {
	return &NullableWorkflowLibraryTrigger{value: val, isSet: true}
}

func (v NullableWorkflowLibraryTrigger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowLibraryTrigger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

