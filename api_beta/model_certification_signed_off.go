/*
IdentityNow Beta API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_beta

import (
	"encoding/json"
	"fmt"
)

// checks if the CertificationSignedOff type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CertificationSignedOff{}

// CertificationSignedOff struct for CertificationSignedOff
type CertificationSignedOff struct {
	Certification CertificationSignedOffCertification `json:"certification"`
	AdditionalProperties map[string]interface{}
}

type _CertificationSignedOff CertificationSignedOff

// NewCertificationSignedOff instantiates a new CertificationSignedOff object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificationSignedOff(certification CertificationSignedOffCertification) *CertificationSignedOff {
	this := CertificationSignedOff{}
	this.Certification = certification
	return &this
}

// NewCertificationSignedOffWithDefaults instantiates a new CertificationSignedOff object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificationSignedOffWithDefaults() *CertificationSignedOff {
	this := CertificationSignedOff{}
	return &this
}

// GetCertification returns the Certification field value
func (o *CertificationSignedOff) GetCertification() CertificationSignedOffCertification {
	if o == nil {
		var ret CertificationSignedOffCertification
		return ret
	}

	return o.Certification
}

// GetCertificationOk returns a tuple with the Certification field value
// and a boolean to check if the value has been set.
func (o *CertificationSignedOff) GetCertificationOk() (*CertificationSignedOffCertification, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Certification, true
}

// SetCertification sets field value
func (o *CertificationSignedOff) SetCertification(v CertificationSignedOffCertification) {
	o.Certification = v
}

func (o CertificationSignedOff) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CertificationSignedOff) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["certification"] = o.Certification

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CertificationSignedOff) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"certification",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCertificationSignedOff := _CertificationSignedOff{}

	if err = json.Unmarshal(bytes, &varCertificationSignedOff); err == nil {
	*o = CertificationSignedOff(varCertificationSignedOff)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "certification")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCertificationSignedOff struct {
	value *CertificationSignedOff
	isSet bool
}

func (v NullableCertificationSignedOff) Get() *CertificationSignedOff {
	return v.value
}

func (v *NullableCertificationSignedOff) Set(val *CertificationSignedOff) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificationSignedOff) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificationSignedOff) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificationSignedOff(val *CertificationSignedOff) *NullableCertificationSignedOff {
	return &NullableCertificationSignedOff{value: val, isSet: true}
}

func (v NullableCertificationSignedOff) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificationSignedOff) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

