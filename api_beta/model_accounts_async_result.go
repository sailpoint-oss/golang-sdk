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

// checks if the AccountsAsyncResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountsAsyncResult{}

// AccountsAsyncResult Accounts async response containing details on started async process
type AccountsAsyncResult struct {
	// id of the task
	Id string `json:"id"`
	AdditionalProperties map[string]interface{}
}

type _AccountsAsyncResult AccountsAsyncResult

// NewAccountsAsyncResult instantiates a new AccountsAsyncResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountsAsyncResult(id string) *AccountsAsyncResult {
	this := AccountsAsyncResult{}
	this.Id = id
	return &this
}

// NewAccountsAsyncResultWithDefaults instantiates a new AccountsAsyncResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountsAsyncResultWithDefaults() *AccountsAsyncResult {
	this := AccountsAsyncResult{}
	return &this
}

// GetId returns the Id field value
func (o *AccountsAsyncResult) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AccountsAsyncResult) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AccountsAsyncResult) SetId(v string) {
	o.Id = v
}

func (o AccountsAsyncResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountsAsyncResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountsAsyncResult) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varAccountsAsyncResult := _AccountsAsyncResult{}

	if err = json.Unmarshal(bytes, &varAccountsAsyncResult); err == nil {
	*o = AccountsAsyncResult(varAccountsAsyncResult)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountsAsyncResult struct {
	value *AccountsAsyncResult
	isSet bool
}

func (v NullableAccountsAsyncResult) Get() *AccountsAsyncResult {
	return v.value
}

func (v *NullableAccountsAsyncResult) Set(val *AccountsAsyncResult) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountsAsyncResult) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountsAsyncResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountsAsyncResult(val *AccountsAsyncResult) *NullableAccountsAsyncResult {
	return &NullableAccountsAsyncResult{value: val, isSet: true}
}

func (v NullableAccountsAsyncResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountsAsyncResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

