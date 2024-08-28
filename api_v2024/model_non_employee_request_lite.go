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

// checks if the NonEmployeeRequestLite type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NonEmployeeRequestLite{}

// NonEmployeeRequestLite struct for NonEmployeeRequestLite
type NonEmployeeRequestLite struct {
	// Non-Employee request id.
	Id *string `json:"id,omitempty"`
	Requester *NonEmployeeIdentityReferenceWithId `json:"requester,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NonEmployeeRequestLite NonEmployeeRequestLite

// NewNonEmployeeRequestLite instantiates a new NonEmployeeRequestLite object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNonEmployeeRequestLite() *NonEmployeeRequestLite {
	this := NonEmployeeRequestLite{}
	return &this
}

// NewNonEmployeeRequestLiteWithDefaults instantiates a new NonEmployeeRequestLite object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNonEmployeeRequestLiteWithDefaults() *NonEmployeeRequestLite {
	this := NonEmployeeRequestLite{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NonEmployeeRequestLite) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeRequestLite) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NonEmployeeRequestLite) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NonEmployeeRequestLite) SetId(v string) {
	o.Id = &v
}

// GetRequester returns the Requester field value if set, zero value otherwise.
func (o *NonEmployeeRequestLite) GetRequester() NonEmployeeIdentityReferenceWithId {
	if o == nil || IsNil(o.Requester) {
		var ret NonEmployeeIdentityReferenceWithId
		return ret
	}
	return *o.Requester
}

// GetRequesterOk returns a tuple with the Requester field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeRequestLite) GetRequesterOk() (*NonEmployeeIdentityReferenceWithId, bool) {
	if o == nil || IsNil(o.Requester) {
		return nil, false
	}
	return o.Requester, true
}

// HasRequester returns a boolean if a field has been set.
func (o *NonEmployeeRequestLite) HasRequester() bool {
	if o != nil && !IsNil(o.Requester) {
		return true
	}

	return false
}

// SetRequester gets a reference to the given NonEmployeeIdentityReferenceWithId and assigns it to the Requester field.
func (o *NonEmployeeRequestLite) SetRequester(v NonEmployeeIdentityReferenceWithId) {
	o.Requester = &v
}

func (o NonEmployeeRequestLite) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NonEmployeeRequestLite) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Requester) {
		toSerialize["requester"] = o.Requester
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NonEmployeeRequestLite) UnmarshalJSON(data []byte) (err error) {
	varNonEmployeeRequestLite := _NonEmployeeRequestLite{}

	err = json.Unmarshal(data, &varNonEmployeeRequestLite)

	if err != nil {
		return err
	}

	*o = NonEmployeeRequestLite(varNonEmployeeRequestLite)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "requester")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNonEmployeeRequestLite struct {
	value *NonEmployeeRequestLite
	isSet bool
}

func (v NullableNonEmployeeRequestLite) Get() *NonEmployeeRequestLite {
	return v.value
}

func (v *NullableNonEmployeeRequestLite) Set(val *NonEmployeeRequestLite) {
	v.value = val
	v.isSet = true
}

func (v NullableNonEmployeeRequestLite) IsSet() bool {
	return v.isSet
}

func (v *NullableNonEmployeeRequestLite) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNonEmployeeRequestLite(val *NonEmployeeRequestLite) *NullableNonEmployeeRequestLite {
	return &NullableNonEmployeeRequestLite{value: val, isSet: true}
}

func (v NullableNonEmployeeRequestLite) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNonEmployeeRequestLite) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

