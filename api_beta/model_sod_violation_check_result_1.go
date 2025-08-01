/*
Identity Security Cloud Beta API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_beta

import (
	"encoding/json"
)

// checks if the SodViolationCheckResult1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SodViolationCheckResult1{}

// SodViolationCheckResult1 The inner object representing the completed SOD Violation check
type SodViolationCheckResult1 struct {
	Message *ErrorMessageDto `json:"message,omitempty"`
	// Arbitrary key-value pairs. They will never be processed by the IdentityNow system but will be returned on completion of the violation check.
	ClientMetadata *map[string]string `json:"clientMetadata,omitempty"`
	ViolationContexts []SodViolationContext1 `json:"violationContexts,omitempty"`
	// A list of the Policies that were violated.
	ViolatedPolicies []SodPolicyDto1 `json:"violatedPolicies,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SodViolationCheckResult1 SodViolationCheckResult1

// NewSodViolationCheckResult1 instantiates a new SodViolationCheckResult1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSodViolationCheckResult1() *SodViolationCheckResult1 {
	this := SodViolationCheckResult1{}
	return &this
}

// NewSodViolationCheckResult1WithDefaults instantiates a new SodViolationCheckResult1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSodViolationCheckResult1WithDefaults() *SodViolationCheckResult1 {
	this := SodViolationCheckResult1{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *SodViolationCheckResult1) GetMessage() ErrorMessageDto {
	if o == nil || IsNil(o.Message) {
		var ret ErrorMessageDto
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SodViolationCheckResult1) GetMessageOk() (*ErrorMessageDto, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *SodViolationCheckResult1) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given ErrorMessageDto and assigns it to the Message field.
func (o *SodViolationCheckResult1) SetMessage(v ErrorMessageDto) {
	o.Message = &v
}

// GetClientMetadata returns the ClientMetadata field value if set, zero value otherwise.
func (o *SodViolationCheckResult1) GetClientMetadata() map[string]string {
	if o == nil || IsNil(o.ClientMetadata) {
		var ret map[string]string
		return ret
	}
	return *o.ClientMetadata
}

// GetClientMetadataOk returns a tuple with the ClientMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SodViolationCheckResult1) GetClientMetadataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.ClientMetadata) {
		return nil, false
	}
	return o.ClientMetadata, true
}

// HasClientMetadata returns a boolean if a field has been set.
func (o *SodViolationCheckResult1) HasClientMetadata() bool {
	if o != nil && !IsNil(o.ClientMetadata) {
		return true
	}

	return false
}

// SetClientMetadata gets a reference to the given map[string]string and assigns it to the ClientMetadata field.
func (o *SodViolationCheckResult1) SetClientMetadata(v map[string]string) {
	o.ClientMetadata = &v
}

// GetViolationContexts returns the ViolationContexts field value if set, zero value otherwise.
func (o *SodViolationCheckResult1) GetViolationContexts() []SodViolationContext1 {
	if o == nil || IsNil(o.ViolationContexts) {
		var ret []SodViolationContext1
		return ret
	}
	return o.ViolationContexts
}

// GetViolationContextsOk returns a tuple with the ViolationContexts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SodViolationCheckResult1) GetViolationContextsOk() ([]SodViolationContext1, bool) {
	if o == nil || IsNil(o.ViolationContexts) {
		return nil, false
	}
	return o.ViolationContexts, true
}

// HasViolationContexts returns a boolean if a field has been set.
func (o *SodViolationCheckResult1) HasViolationContexts() bool {
	if o != nil && !IsNil(o.ViolationContexts) {
		return true
	}

	return false
}

// SetViolationContexts gets a reference to the given []SodViolationContext1 and assigns it to the ViolationContexts field.
func (o *SodViolationCheckResult1) SetViolationContexts(v []SodViolationContext1) {
	o.ViolationContexts = v
}

// GetViolatedPolicies returns the ViolatedPolicies field value if set, zero value otherwise.
func (o *SodViolationCheckResult1) GetViolatedPolicies() []SodPolicyDto1 {
	if o == nil || IsNil(o.ViolatedPolicies) {
		var ret []SodPolicyDto1
		return ret
	}
	return o.ViolatedPolicies
}

// GetViolatedPoliciesOk returns a tuple with the ViolatedPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SodViolationCheckResult1) GetViolatedPoliciesOk() ([]SodPolicyDto1, bool) {
	if o == nil || IsNil(o.ViolatedPolicies) {
		return nil, false
	}
	return o.ViolatedPolicies, true
}

// HasViolatedPolicies returns a boolean if a field has been set.
func (o *SodViolationCheckResult1) HasViolatedPolicies() bool {
	if o != nil && !IsNil(o.ViolatedPolicies) {
		return true
	}

	return false
}

// SetViolatedPolicies gets a reference to the given []SodPolicyDto1 and assigns it to the ViolatedPolicies field.
func (o *SodViolationCheckResult1) SetViolatedPolicies(v []SodPolicyDto1) {
	o.ViolatedPolicies = v
}

func (o SodViolationCheckResult1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SodViolationCheckResult1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.ClientMetadata) {
		toSerialize["clientMetadata"] = o.ClientMetadata
	}
	if !IsNil(o.ViolationContexts) {
		toSerialize["violationContexts"] = o.ViolationContexts
	}
	if !IsNil(o.ViolatedPolicies) {
		toSerialize["violatedPolicies"] = o.ViolatedPolicies
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SodViolationCheckResult1) UnmarshalJSON(data []byte) (err error) {
	varSodViolationCheckResult1 := _SodViolationCheckResult1{}

	err = json.Unmarshal(data, &varSodViolationCheckResult1)

	if err != nil {
		return err
	}

	*o = SodViolationCheckResult1(varSodViolationCheckResult1)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "message")
		delete(additionalProperties, "clientMetadata")
		delete(additionalProperties, "violationContexts")
		delete(additionalProperties, "violatedPolicies")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSodViolationCheckResult1 struct {
	value *SodViolationCheckResult1
	isSet bool
}

func (v NullableSodViolationCheckResult1) Get() *SodViolationCheckResult1 {
	return v.value
}

func (v *NullableSodViolationCheckResult1) Set(val *SodViolationCheckResult1) {
	v.value = val
	v.isSet = true
}

func (v NullableSodViolationCheckResult1) IsSet() bool {
	return v.isSet
}

func (v *NullableSodViolationCheckResult1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSodViolationCheckResult1(val *SodViolationCheckResult1) *NullableSodViolationCheckResult1 {
	return &NullableSodViolationCheckResult1{value: val, isSet: true}
}

func (v NullableSodViolationCheckResult1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSodViolationCheckResult1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


