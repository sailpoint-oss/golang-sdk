/*
Identity Security Cloud V2024 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: v2024
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2024

import (
	"encoding/json"
)

// checks if the AppAccountDetailsSourceAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppAccountDetailsSourceAccount{}

// AppAccountDetailsSourceAccount struct for AppAccountDetailsSourceAccount
type AppAccountDetailsSourceAccount struct {
	// The account ID
	Id *string `json:"id,omitempty"`
	// The native identity of account
	NativeIdentity *string `json:"nativeIdentity,omitempty"`
	// The display name of account
	DisplayName *string `json:"displayName,omitempty"`
	// The source ID of account
	SourceId *string `json:"sourceId,omitempty"`
	// The source name of account
	SourceDisplayName *string `json:"sourceDisplayName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AppAccountDetailsSourceAccount AppAccountDetailsSourceAccount

// NewAppAccountDetailsSourceAccount instantiates a new AppAccountDetailsSourceAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppAccountDetailsSourceAccount() *AppAccountDetailsSourceAccount {
	this := AppAccountDetailsSourceAccount{}
	return &this
}

// NewAppAccountDetailsSourceAccountWithDefaults instantiates a new AppAccountDetailsSourceAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppAccountDetailsSourceAccountWithDefaults() *AppAccountDetailsSourceAccount {
	this := AppAccountDetailsSourceAccount{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AppAccountDetailsSourceAccount) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAccountDetailsSourceAccount) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AppAccountDetailsSourceAccount) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AppAccountDetailsSourceAccount) SetId(v string) {
	o.Id = &v
}

// GetNativeIdentity returns the NativeIdentity field value if set, zero value otherwise.
func (o *AppAccountDetailsSourceAccount) GetNativeIdentity() string {
	if o == nil || IsNil(o.NativeIdentity) {
		var ret string
		return ret
	}
	return *o.NativeIdentity
}

// GetNativeIdentityOk returns a tuple with the NativeIdentity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAccountDetailsSourceAccount) GetNativeIdentityOk() (*string, bool) {
	if o == nil || IsNil(o.NativeIdentity) {
		return nil, false
	}
	return o.NativeIdentity, true
}

// HasNativeIdentity returns a boolean if a field has been set.
func (o *AppAccountDetailsSourceAccount) HasNativeIdentity() bool {
	if o != nil && !IsNil(o.NativeIdentity) {
		return true
	}

	return false
}

// SetNativeIdentity gets a reference to the given string and assigns it to the NativeIdentity field.
func (o *AppAccountDetailsSourceAccount) SetNativeIdentity(v string) {
	o.NativeIdentity = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *AppAccountDetailsSourceAccount) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAccountDetailsSourceAccount) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *AppAccountDetailsSourceAccount) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *AppAccountDetailsSourceAccount) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetSourceId returns the SourceId field value if set, zero value otherwise.
func (o *AppAccountDetailsSourceAccount) GetSourceId() string {
	if o == nil || IsNil(o.SourceId) {
		var ret string
		return ret
	}
	return *o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAccountDetailsSourceAccount) GetSourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.SourceId) {
		return nil, false
	}
	return o.SourceId, true
}

// HasSourceId returns a boolean if a field has been set.
func (o *AppAccountDetailsSourceAccount) HasSourceId() bool {
	if o != nil && !IsNil(o.SourceId) {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given string and assigns it to the SourceId field.
func (o *AppAccountDetailsSourceAccount) SetSourceId(v string) {
	o.SourceId = &v
}

// GetSourceDisplayName returns the SourceDisplayName field value if set, zero value otherwise.
func (o *AppAccountDetailsSourceAccount) GetSourceDisplayName() string {
	if o == nil || IsNil(o.SourceDisplayName) {
		var ret string
		return ret
	}
	return *o.SourceDisplayName
}

// GetSourceDisplayNameOk returns a tuple with the SourceDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAccountDetailsSourceAccount) GetSourceDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.SourceDisplayName) {
		return nil, false
	}
	return o.SourceDisplayName, true
}

// HasSourceDisplayName returns a boolean if a field has been set.
func (o *AppAccountDetailsSourceAccount) HasSourceDisplayName() bool {
	if o != nil && !IsNil(o.SourceDisplayName) {
		return true
	}

	return false
}

// SetSourceDisplayName gets a reference to the given string and assigns it to the SourceDisplayName field.
func (o *AppAccountDetailsSourceAccount) SetSourceDisplayName(v string) {
	o.SourceDisplayName = &v
}

func (o AppAccountDetailsSourceAccount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppAccountDetailsSourceAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.NativeIdentity) {
		toSerialize["nativeIdentity"] = o.NativeIdentity
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.SourceId) {
		toSerialize["sourceId"] = o.SourceId
	}
	if !IsNil(o.SourceDisplayName) {
		toSerialize["sourceDisplayName"] = o.SourceDisplayName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AppAccountDetailsSourceAccount) UnmarshalJSON(data []byte) (err error) {
	varAppAccountDetailsSourceAccount := _AppAccountDetailsSourceAccount{}

	err = json.Unmarshal(data, &varAppAccountDetailsSourceAccount)

	if err != nil {
		return err
	}

	*o = AppAccountDetailsSourceAccount(varAppAccountDetailsSourceAccount)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "nativeIdentity")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "sourceId")
		delete(additionalProperties, "sourceDisplayName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppAccountDetailsSourceAccount struct {
	value *AppAccountDetailsSourceAccount
	isSet bool
}

func (v NullableAppAccountDetailsSourceAccount) Get() *AppAccountDetailsSourceAccount {
	return v.value
}

func (v *NullableAppAccountDetailsSourceAccount) Set(val *AppAccountDetailsSourceAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableAppAccountDetailsSourceAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableAppAccountDetailsSourceAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppAccountDetailsSourceAccount(val *AppAccountDetailsSourceAccount) *NullableAppAccountDetailsSourceAccount {
	return &NullableAppAccountDetailsSourceAccount{value: val, isSet: true}
}

func (v NullableAppAccountDetailsSourceAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppAccountDetailsSourceAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


