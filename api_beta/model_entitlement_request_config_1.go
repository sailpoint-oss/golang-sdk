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

// checks if the EntitlementRequestConfig1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitlementRequestConfig1{}

// EntitlementRequestConfig1 struct for EntitlementRequestConfig1
type EntitlementRequestConfig1 struct {
	// If this is true, entitlement requests are allowed.
	AllowEntitlementRequest *bool `json:"allowEntitlementRequest,omitempty"`
	// If this is true, comments are required to submit entitlement requests.
	RequestCommentsRequired *bool `json:"requestCommentsRequired,omitempty"`
	// If this is true, comments are required to reject entitlement requests.
	DeniedCommentsRequired *bool `json:"deniedCommentsRequired,omitempty"`
	// Approval schemes for granting entitlement request. This can be empty if no approval is needed. Multiple schemes must be comma-separated. The valid schemes are \"entitlementOwner\", \"sourceOwner\", \"manager\" and \"`workgroup:{id}`\". You can use multiple governance groups (workgroups). 
	GrantRequestApprovalSchemes NullableString `json:"grantRequestApprovalSchemes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EntitlementRequestConfig1 EntitlementRequestConfig1

// NewEntitlementRequestConfig1 instantiates a new EntitlementRequestConfig1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementRequestConfig1() *EntitlementRequestConfig1 {
	this := EntitlementRequestConfig1{}
	var allowEntitlementRequest bool = false
	this.AllowEntitlementRequest = &allowEntitlementRequest
	var requestCommentsRequired bool = false
	this.RequestCommentsRequired = &requestCommentsRequired
	var deniedCommentsRequired bool = false
	this.DeniedCommentsRequired = &deniedCommentsRequired
	var grantRequestApprovalSchemes string = "sourceOwner"
	this.GrantRequestApprovalSchemes = *NewNullableString(&grantRequestApprovalSchemes)
	return &this
}

// NewEntitlementRequestConfig1WithDefaults instantiates a new EntitlementRequestConfig1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementRequestConfig1WithDefaults() *EntitlementRequestConfig1 {
	this := EntitlementRequestConfig1{}
	var allowEntitlementRequest bool = false
	this.AllowEntitlementRequest = &allowEntitlementRequest
	var requestCommentsRequired bool = false
	this.RequestCommentsRequired = &requestCommentsRequired
	var deniedCommentsRequired bool = false
	this.DeniedCommentsRequired = &deniedCommentsRequired
	var grantRequestApprovalSchemes string = "sourceOwner"
	this.GrantRequestApprovalSchemes = *NewNullableString(&grantRequestApprovalSchemes)
	return &this
}

// GetAllowEntitlementRequest returns the AllowEntitlementRequest field value if set, zero value otherwise.
func (o *EntitlementRequestConfig1) GetAllowEntitlementRequest() bool {
	if o == nil || IsNil(o.AllowEntitlementRequest) {
		var ret bool
		return ret
	}
	return *o.AllowEntitlementRequest
}

// GetAllowEntitlementRequestOk returns a tuple with the AllowEntitlementRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementRequestConfig1) GetAllowEntitlementRequestOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowEntitlementRequest) {
		return nil, false
	}
	return o.AllowEntitlementRequest, true
}

// HasAllowEntitlementRequest returns a boolean if a field has been set.
func (o *EntitlementRequestConfig1) HasAllowEntitlementRequest() bool {
	if o != nil && !IsNil(o.AllowEntitlementRequest) {
		return true
	}

	return false
}

// SetAllowEntitlementRequest gets a reference to the given bool and assigns it to the AllowEntitlementRequest field.
func (o *EntitlementRequestConfig1) SetAllowEntitlementRequest(v bool) {
	o.AllowEntitlementRequest = &v
}

// GetRequestCommentsRequired returns the RequestCommentsRequired field value if set, zero value otherwise.
func (o *EntitlementRequestConfig1) GetRequestCommentsRequired() bool {
	if o == nil || IsNil(o.RequestCommentsRequired) {
		var ret bool
		return ret
	}
	return *o.RequestCommentsRequired
}

// GetRequestCommentsRequiredOk returns a tuple with the RequestCommentsRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementRequestConfig1) GetRequestCommentsRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.RequestCommentsRequired) {
		return nil, false
	}
	return o.RequestCommentsRequired, true
}

// HasRequestCommentsRequired returns a boolean if a field has been set.
func (o *EntitlementRequestConfig1) HasRequestCommentsRequired() bool {
	if o != nil && !IsNil(o.RequestCommentsRequired) {
		return true
	}

	return false
}

// SetRequestCommentsRequired gets a reference to the given bool and assigns it to the RequestCommentsRequired field.
func (o *EntitlementRequestConfig1) SetRequestCommentsRequired(v bool) {
	o.RequestCommentsRequired = &v
}

// GetDeniedCommentsRequired returns the DeniedCommentsRequired field value if set, zero value otherwise.
func (o *EntitlementRequestConfig1) GetDeniedCommentsRequired() bool {
	if o == nil || IsNil(o.DeniedCommentsRequired) {
		var ret bool
		return ret
	}
	return *o.DeniedCommentsRequired
}

// GetDeniedCommentsRequiredOk returns a tuple with the DeniedCommentsRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementRequestConfig1) GetDeniedCommentsRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.DeniedCommentsRequired) {
		return nil, false
	}
	return o.DeniedCommentsRequired, true
}

// HasDeniedCommentsRequired returns a boolean if a field has been set.
func (o *EntitlementRequestConfig1) HasDeniedCommentsRequired() bool {
	if o != nil && !IsNil(o.DeniedCommentsRequired) {
		return true
	}

	return false
}

// SetDeniedCommentsRequired gets a reference to the given bool and assigns it to the DeniedCommentsRequired field.
func (o *EntitlementRequestConfig1) SetDeniedCommentsRequired(v bool) {
	o.DeniedCommentsRequired = &v
}

// GetGrantRequestApprovalSchemes returns the GrantRequestApprovalSchemes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EntitlementRequestConfig1) GetGrantRequestApprovalSchemes() string {
	if o == nil || IsNil(o.GrantRequestApprovalSchemes.Get()) {
		var ret string
		return ret
	}
	return *o.GrantRequestApprovalSchemes.Get()
}

// GetGrantRequestApprovalSchemesOk returns a tuple with the GrantRequestApprovalSchemes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EntitlementRequestConfig1) GetGrantRequestApprovalSchemesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GrantRequestApprovalSchemes.Get(), o.GrantRequestApprovalSchemes.IsSet()
}

// HasGrantRequestApprovalSchemes returns a boolean if a field has been set.
func (o *EntitlementRequestConfig1) HasGrantRequestApprovalSchemes() bool {
	if o != nil && o.GrantRequestApprovalSchemes.IsSet() {
		return true
	}

	return false
}

// SetGrantRequestApprovalSchemes gets a reference to the given NullableString and assigns it to the GrantRequestApprovalSchemes field.
func (o *EntitlementRequestConfig1) SetGrantRequestApprovalSchemes(v string) {
	o.GrantRequestApprovalSchemes.Set(&v)
}
// SetGrantRequestApprovalSchemesNil sets the value for GrantRequestApprovalSchemes to be an explicit nil
func (o *EntitlementRequestConfig1) SetGrantRequestApprovalSchemesNil() {
	o.GrantRequestApprovalSchemes.Set(nil)
}

// UnsetGrantRequestApprovalSchemes ensures that no value is present for GrantRequestApprovalSchemes, not even an explicit nil
func (o *EntitlementRequestConfig1) UnsetGrantRequestApprovalSchemes() {
	o.GrantRequestApprovalSchemes.Unset()
}

func (o EntitlementRequestConfig1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitlementRequestConfig1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllowEntitlementRequest) {
		toSerialize["allowEntitlementRequest"] = o.AllowEntitlementRequest
	}
	if !IsNil(o.RequestCommentsRequired) {
		toSerialize["requestCommentsRequired"] = o.RequestCommentsRequired
	}
	if !IsNil(o.DeniedCommentsRequired) {
		toSerialize["deniedCommentsRequired"] = o.DeniedCommentsRequired
	}
	if o.GrantRequestApprovalSchemes.IsSet() {
		toSerialize["grantRequestApprovalSchemes"] = o.GrantRequestApprovalSchemes.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EntitlementRequestConfig1) UnmarshalJSON(data []byte) (err error) {
	varEntitlementRequestConfig1 := _EntitlementRequestConfig1{}

	err = json.Unmarshal(data, &varEntitlementRequestConfig1)

	if err != nil {
		return err
	}

	*o = EntitlementRequestConfig1(varEntitlementRequestConfig1)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "allowEntitlementRequest")
		delete(additionalProperties, "requestCommentsRequired")
		delete(additionalProperties, "deniedCommentsRequired")
		delete(additionalProperties, "grantRequestApprovalSchemes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntitlementRequestConfig1 struct {
	value *EntitlementRequestConfig1
	isSet bool
}

func (v NullableEntitlementRequestConfig1) Get() *EntitlementRequestConfig1 {
	return v.value
}

func (v *NullableEntitlementRequestConfig1) Set(val *EntitlementRequestConfig1) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementRequestConfig1) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementRequestConfig1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementRequestConfig1(val *EntitlementRequestConfig1) *NullableEntitlementRequestConfig1 {
	return &NullableEntitlementRequestConfig1{value: val, isSet: true}
}

func (v NullableEntitlementRequestConfig1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementRequestConfig1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


