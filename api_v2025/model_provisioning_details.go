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

// checks if the ProvisioningDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProvisioningDetails{}

// ProvisioningDetails Provides additional details about provisioning for this request.
type ProvisioningDetails struct {
	// Ordered CSV of sub phase references to objects that contain more information about provisioning. For example, this can contain \"manualWorkItemDetails\" which indicate that there is further information in that object for this phase.
	OrderedSubPhaseReferences *string `json:"orderedSubPhaseReferences,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProvisioningDetails ProvisioningDetails

// NewProvisioningDetails instantiates a new ProvisioningDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisioningDetails() *ProvisioningDetails {
	this := ProvisioningDetails{}
	return &this
}

// NewProvisioningDetailsWithDefaults instantiates a new ProvisioningDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisioningDetailsWithDefaults() *ProvisioningDetails {
	this := ProvisioningDetails{}
	return &this
}

// GetOrderedSubPhaseReferences returns the OrderedSubPhaseReferences field value if set, zero value otherwise.
func (o *ProvisioningDetails) GetOrderedSubPhaseReferences() string {
	if o == nil || IsNil(o.OrderedSubPhaseReferences) {
		var ret string
		return ret
	}
	return *o.OrderedSubPhaseReferences
}

// GetOrderedSubPhaseReferencesOk returns a tuple with the OrderedSubPhaseReferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningDetails) GetOrderedSubPhaseReferencesOk() (*string, bool) {
	if o == nil || IsNil(o.OrderedSubPhaseReferences) {
		return nil, false
	}
	return o.OrderedSubPhaseReferences, true
}

// HasOrderedSubPhaseReferences returns a boolean if a field has been set.
func (o *ProvisioningDetails) HasOrderedSubPhaseReferences() bool {
	if o != nil && !IsNil(o.OrderedSubPhaseReferences) {
		return true
	}

	return false
}

// SetOrderedSubPhaseReferences gets a reference to the given string and assigns it to the OrderedSubPhaseReferences field.
func (o *ProvisioningDetails) SetOrderedSubPhaseReferences(v string) {
	o.OrderedSubPhaseReferences = &v
}

func (o ProvisioningDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProvisioningDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OrderedSubPhaseReferences) {
		toSerialize["orderedSubPhaseReferences"] = o.OrderedSubPhaseReferences
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProvisioningDetails) UnmarshalJSON(data []byte) (err error) {
	varProvisioningDetails := _ProvisioningDetails{}

	err = json.Unmarshal(data, &varProvisioningDetails)

	if err != nil {
		return err
	}

	*o = ProvisioningDetails(varProvisioningDetails)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "orderedSubPhaseReferences")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProvisioningDetails struct {
	value *ProvisioningDetails
	isSet bool
}

func (v NullableProvisioningDetails) Get() *ProvisioningDetails {
	return v.value
}

func (v *NullableProvisioningDetails) Set(val *ProvisioningDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisioningDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisioningDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisioningDetails(val *ProvisioningDetails) *NullableProvisioningDetails {
	return &NullableProvisioningDetails{value: val, isSet: true}
}

func (v NullableProvisioningDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisioningDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


