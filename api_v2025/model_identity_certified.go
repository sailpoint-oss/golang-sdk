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

// checks if the IdentityCertified type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentityCertified{}

// IdentityCertified struct for IdentityCertified
type IdentityCertified struct {
	// the id of the certification item
	CertificationId *string `json:"certificationId,omitempty"`
	// the certification item name
	CertificationName *string `json:"certificationName,omitempty"`
	// the date ceritification was signed
	SignedDate *string `json:"signedDate,omitempty"`
	// this field is deprecated and may go away
	Certifiers []CertifierResponse `json:"certifiers,omitempty"`
	// The list of identities who review this certification
	Reviewers []CertifierResponse `json:"reviewers,omitempty"`
	Signer *CertifierResponse `json:"signer,omitempty"`
	// the event type
	EventType *string `json:"eventType,omitempty"`
	// the date of event
	Dt *string `json:"dt,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdentityCertified IdentityCertified

// NewIdentityCertified instantiates a new IdentityCertified object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityCertified() *IdentityCertified {
	this := IdentityCertified{}
	return &this
}

// NewIdentityCertifiedWithDefaults instantiates a new IdentityCertified object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityCertifiedWithDefaults() *IdentityCertified {
	this := IdentityCertified{}
	return &this
}

// GetCertificationId returns the CertificationId field value if set, zero value otherwise.
func (o *IdentityCertified) GetCertificationId() string {
	if o == nil || IsNil(o.CertificationId) {
		var ret string
		return ret
	}
	return *o.CertificationId
}

// GetCertificationIdOk returns a tuple with the CertificationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityCertified) GetCertificationIdOk() (*string, bool) {
	if o == nil || IsNil(o.CertificationId) {
		return nil, false
	}
	return o.CertificationId, true
}

// HasCertificationId returns a boolean if a field has been set.
func (o *IdentityCertified) HasCertificationId() bool {
	if o != nil && !IsNil(o.CertificationId) {
		return true
	}

	return false
}

// SetCertificationId gets a reference to the given string and assigns it to the CertificationId field.
func (o *IdentityCertified) SetCertificationId(v string) {
	o.CertificationId = &v
}

// GetCertificationName returns the CertificationName field value if set, zero value otherwise.
func (o *IdentityCertified) GetCertificationName() string {
	if o == nil || IsNil(o.CertificationName) {
		var ret string
		return ret
	}
	return *o.CertificationName
}

// GetCertificationNameOk returns a tuple with the CertificationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityCertified) GetCertificationNameOk() (*string, bool) {
	if o == nil || IsNil(o.CertificationName) {
		return nil, false
	}
	return o.CertificationName, true
}

// HasCertificationName returns a boolean if a field has been set.
func (o *IdentityCertified) HasCertificationName() bool {
	if o != nil && !IsNil(o.CertificationName) {
		return true
	}

	return false
}

// SetCertificationName gets a reference to the given string and assigns it to the CertificationName field.
func (o *IdentityCertified) SetCertificationName(v string) {
	o.CertificationName = &v
}

// GetSignedDate returns the SignedDate field value if set, zero value otherwise.
func (o *IdentityCertified) GetSignedDate() string {
	if o == nil || IsNil(o.SignedDate) {
		var ret string
		return ret
	}
	return *o.SignedDate
}

// GetSignedDateOk returns a tuple with the SignedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityCertified) GetSignedDateOk() (*string, bool) {
	if o == nil || IsNil(o.SignedDate) {
		return nil, false
	}
	return o.SignedDate, true
}

// HasSignedDate returns a boolean if a field has been set.
func (o *IdentityCertified) HasSignedDate() bool {
	if o != nil && !IsNil(o.SignedDate) {
		return true
	}

	return false
}

// SetSignedDate gets a reference to the given string and assigns it to the SignedDate field.
func (o *IdentityCertified) SetSignedDate(v string) {
	o.SignedDate = &v
}

// GetCertifiers returns the Certifiers field value if set, zero value otherwise.
func (o *IdentityCertified) GetCertifiers() []CertifierResponse {
	if o == nil || IsNil(o.Certifiers) {
		var ret []CertifierResponse
		return ret
	}
	return o.Certifiers
}

// GetCertifiersOk returns a tuple with the Certifiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityCertified) GetCertifiersOk() ([]CertifierResponse, bool) {
	if o == nil || IsNil(o.Certifiers) {
		return nil, false
	}
	return o.Certifiers, true
}

// HasCertifiers returns a boolean if a field has been set.
func (o *IdentityCertified) HasCertifiers() bool {
	if o != nil && !IsNil(o.Certifiers) {
		return true
	}

	return false
}

// SetCertifiers gets a reference to the given []CertifierResponse and assigns it to the Certifiers field.
func (o *IdentityCertified) SetCertifiers(v []CertifierResponse) {
	o.Certifiers = v
}

// GetReviewers returns the Reviewers field value if set, zero value otherwise.
func (o *IdentityCertified) GetReviewers() []CertifierResponse {
	if o == nil || IsNil(o.Reviewers) {
		var ret []CertifierResponse
		return ret
	}
	return o.Reviewers
}

// GetReviewersOk returns a tuple with the Reviewers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityCertified) GetReviewersOk() ([]CertifierResponse, bool) {
	if o == nil || IsNil(o.Reviewers) {
		return nil, false
	}
	return o.Reviewers, true
}

// HasReviewers returns a boolean if a field has been set.
func (o *IdentityCertified) HasReviewers() bool {
	if o != nil && !IsNil(o.Reviewers) {
		return true
	}

	return false
}

// SetReviewers gets a reference to the given []CertifierResponse and assigns it to the Reviewers field.
func (o *IdentityCertified) SetReviewers(v []CertifierResponse) {
	o.Reviewers = v
}

// GetSigner returns the Signer field value if set, zero value otherwise.
func (o *IdentityCertified) GetSigner() CertifierResponse {
	if o == nil || IsNil(o.Signer) {
		var ret CertifierResponse
		return ret
	}
	return *o.Signer
}

// GetSignerOk returns a tuple with the Signer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityCertified) GetSignerOk() (*CertifierResponse, bool) {
	if o == nil || IsNil(o.Signer) {
		return nil, false
	}
	return o.Signer, true
}

// HasSigner returns a boolean if a field has been set.
func (o *IdentityCertified) HasSigner() bool {
	if o != nil && !IsNil(o.Signer) {
		return true
	}

	return false
}

// SetSigner gets a reference to the given CertifierResponse and assigns it to the Signer field.
func (o *IdentityCertified) SetSigner(v CertifierResponse) {
	o.Signer = &v
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *IdentityCertified) GetEventType() string {
	if o == nil || IsNil(o.EventType) {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityCertified) GetEventTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EventType) {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *IdentityCertified) HasEventType() bool {
	if o != nil && !IsNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *IdentityCertified) SetEventType(v string) {
	o.EventType = &v
}

// GetDt returns the Dt field value if set, zero value otherwise.
func (o *IdentityCertified) GetDt() string {
	if o == nil || IsNil(o.Dt) {
		var ret string
		return ret
	}
	return *o.Dt
}

// GetDtOk returns a tuple with the Dt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityCertified) GetDtOk() (*string, bool) {
	if o == nil || IsNil(o.Dt) {
		return nil, false
	}
	return o.Dt, true
}

// HasDt returns a boolean if a field has been set.
func (o *IdentityCertified) HasDt() bool {
	if o != nil && !IsNil(o.Dt) {
		return true
	}

	return false
}

// SetDt gets a reference to the given string and assigns it to the Dt field.
func (o *IdentityCertified) SetDt(v string) {
	o.Dt = &v
}

func (o IdentityCertified) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentityCertified) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CertificationId) {
		toSerialize["certificationId"] = o.CertificationId
	}
	if !IsNil(o.CertificationName) {
		toSerialize["certificationName"] = o.CertificationName
	}
	if !IsNil(o.SignedDate) {
		toSerialize["signedDate"] = o.SignedDate
	}
	if !IsNil(o.Certifiers) {
		toSerialize["certifiers"] = o.Certifiers
	}
	if !IsNil(o.Reviewers) {
		toSerialize["reviewers"] = o.Reviewers
	}
	if !IsNil(o.Signer) {
		toSerialize["signer"] = o.Signer
	}
	if !IsNil(o.EventType) {
		toSerialize["eventType"] = o.EventType
	}
	if !IsNil(o.Dt) {
		toSerialize["dt"] = o.Dt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IdentityCertified) UnmarshalJSON(data []byte) (err error) {
	varIdentityCertified := _IdentityCertified{}

	err = json.Unmarshal(data, &varIdentityCertified)

	if err != nil {
		return err
	}

	*o = IdentityCertified(varIdentityCertified)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "certificationId")
		delete(additionalProperties, "certificationName")
		delete(additionalProperties, "signedDate")
		delete(additionalProperties, "certifiers")
		delete(additionalProperties, "reviewers")
		delete(additionalProperties, "signer")
		delete(additionalProperties, "eventType")
		delete(additionalProperties, "dt")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentityCertified struct {
	value *IdentityCertified
	isSet bool
}

func (v NullableIdentityCertified) Get() *IdentityCertified {
	return v.value
}

func (v *NullableIdentityCertified) Set(val *IdentityCertified) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityCertified) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityCertified) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityCertified(val *IdentityCertified) *NullableIdentityCertified {
	return &NullableIdentityCertified{value: val, isSet: true}
}

func (v NullableIdentityCertified) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityCertified) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


