/*
SailPoint SaaS API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sailpointv2sdk

import (
	"encoding/json"
)

// SearchIdentityAppsInnerAccount struct for SearchIdentityAppsInnerAccount
type SearchIdentityAppsInnerAccount struct {
	Id *string `json:"id,omitempty"`
	AccountId *string `json:"accountId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SearchIdentityAppsInnerAccount SearchIdentityAppsInnerAccount

// NewSearchIdentityAppsInnerAccount instantiates a new SearchIdentityAppsInnerAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchIdentityAppsInnerAccount() *SearchIdentityAppsInnerAccount {
	this := SearchIdentityAppsInnerAccount{}
	return &this
}

// NewSearchIdentityAppsInnerAccountWithDefaults instantiates a new SearchIdentityAppsInnerAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchIdentityAppsInnerAccountWithDefaults() *SearchIdentityAppsInnerAccount {
	this := SearchIdentityAppsInnerAccount{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SearchIdentityAppsInnerAccount) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchIdentityAppsInnerAccount) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SearchIdentityAppsInnerAccount) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SearchIdentityAppsInnerAccount) SetId(v string) {
	o.Id = &v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *SearchIdentityAppsInnerAccount) GetAccountId() string {
	if o == nil || isNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchIdentityAppsInnerAccount) GetAccountIdOk() (*string, bool) {
	if o == nil || isNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *SearchIdentityAppsInnerAccount) HasAccountId() bool {
	if o != nil && !isNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *SearchIdentityAppsInnerAccount) SetAccountId(v string) {
	o.AccountId = &v
}

func (o SearchIdentityAppsInnerAccount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.AccountId) {
		toSerialize["accountId"] = o.AccountId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SearchIdentityAppsInnerAccount) UnmarshalJSON(bytes []byte) (err error) {
	varSearchIdentityAppsInnerAccount := _SearchIdentityAppsInnerAccount{}

	if err = json.Unmarshal(bytes, &varSearchIdentityAppsInnerAccount); err == nil {
		*o = SearchIdentityAppsInnerAccount(varSearchIdentityAppsInnerAccount)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "accountId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSearchIdentityAppsInnerAccount struct {
	value *SearchIdentityAppsInnerAccount
	isSet bool
}

func (v NullableSearchIdentityAppsInnerAccount) Get() *SearchIdentityAppsInnerAccount {
	return v.value
}

func (v *NullableSearchIdentityAppsInnerAccount) Set(val *SearchIdentityAppsInnerAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchIdentityAppsInnerAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchIdentityAppsInnerAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchIdentityAppsInnerAccount(val *SearchIdentityAppsInnerAccount) *NullableSearchIdentityAppsInnerAccount {
	return &NullableSearchIdentityAppsInnerAccount{value: val, isSet: true}
}

func (v NullableSearchIdentityAppsInnerAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchIdentityAppsInnerAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


