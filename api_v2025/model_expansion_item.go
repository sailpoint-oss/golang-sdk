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

// checks if the ExpansionItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExpansionItem{}

// ExpansionItem struct for ExpansionItem
type ExpansionItem struct {
	// The ID of the account
	AccountId *string `json:"accountId,omitempty"`
	// Cause of the expansion item.
	Cause *string `json:"cause,omitempty"`
	// The name of the item
	Name *string `json:"name,omitempty"`
	AttributeRequest *AttributeRequest `json:"attributeRequest,omitempty"`
	Source *AccountSource `json:"source,omitempty"`
	// ID of the expansion item
	Id *string `json:"id,omitempty"`
	// State of the expansion item
	State *string `json:"state,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ExpansionItem ExpansionItem

// NewExpansionItem instantiates a new ExpansionItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExpansionItem() *ExpansionItem {
	this := ExpansionItem{}
	return &this
}

// NewExpansionItemWithDefaults instantiates a new ExpansionItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExpansionItemWithDefaults() *ExpansionItem {
	this := ExpansionItem{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *ExpansionItem) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpansionItem) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *ExpansionItem) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *ExpansionItem) SetAccountId(v string) {
	o.AccountId = &v
}

// GetCause returns the Cause field value if set, zero value otherwise.
func (o *ExpansionItem) GetCause() string {
	if o == nil || IsNil(o.Cause) {
		var ret string
		return ret
	}
	return *o.Cause
}

// GetCauseOk returns a tuple with the Cause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpansionItem) GetCauseOk() (*string, bool) {
	if o == nil || IsNil(o.Cause) {
		return nil, false
	}
	return o.Cause, true
}

// HasCause returns a boolean if a field has been set.
func (o *ExpansionItem) HasCause() bool {
	if o != nil && !IsNil(o.Cause) {
		return true
	}

	return false
}

// SetCause gets a reference to the given string and assigns it to the Cause field.
func (o *ExpansionItem) SetCause(v string) {
	o.Cause = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ExpansionItem) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpansionItem) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ExpansionItem) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ExpansionItem) SetName(v string) {
	o.Name = &v
}

// GetAttributeRequest returns the AttributeRequest field value if set, zero value otherwise.
func (o *ExpansionItem) GetAttributeRequest() AttributeRequest {
	if o == nil || IsNil(o.AttributeRequest) {
		var ret AttributeRequest
		return ret
	}
	return *o.AttributeRequest
}

// GetAttributeRequestOk returns a tuple with the AttributeRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpansionItem) GetAttributeRequestOk() (*AttributeRequest, bool) {
	if o == nil || IsNil(o.AttributeRequest) {
		return nil, false
	}
	return o.AttributeRequest, true
}

// HasAttributeRequest returns a boolean if a field has been set.
func (o *ExpansionItem) HasAttributeRequest() bool {
	if o != nil && !IsNil(o.AttributeRequest) {
		return true
	}

	return false
}

// SetAttributeRequest gets a reference to the given AttributeRequest and assigns it to the AttributeRequest field.
func (o *ExpansionItem) SetAttributeRequest(v AttributeRequest) {
	o.AttributeRequest = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *ExpansionItem) GetSource() AccountSource {
	if o == nil || IsNil(o.Source) {
		var ret AccountSource
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpansionItem) GetSourceOk() (*AccountSource, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *ExpansionItem) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given AccountSource and assigns it to the Source field.
func (o *ExpansionItem) SetSource(v AccountSource) {
	o.Source = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ExpansionItem) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpansionItem) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ExpansionItem) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ExpansionItem) SetId(v string) {
	o.Id = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ExpansionItem) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpansionItem) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ExpansionItem) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *ExpansionItem) SetState(v string) {
	o.State = &v
}

func (o ExpansionItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExpansionItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountId) {
		toSerialize["accountId"] = o.AccountId
	}
	if !IsNil(o.Cause) {
		toSerialize["cause"] = o.Cause
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.AttributeRequest) {
		toSerialize["attributeRequest"] = o.AttributeRequest
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ExpansionItem) UnmarshalJSON(data []byte) (err error) {
	varExpansionItem := _ExpansionItem{}

	err = json.Unmarshal(data, &varExpansionItem)

	if err != nil {
		return err
	}

	*o = ExpansionItem(varExpansionItem)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "accountId")
		delete(additionalProperties, "cause")
		delete(additionalProperties, "name")
		delete(additionalProperties, "attributeRequest")
		delete(additionalProperties, "source")
		delete(additionalProperties, "id")
		delete(additionalProperties, "state")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExpansionItem struct {
	value *ExpansionItem
	isSet bool
}

func (v NullableExpansionItem) Get() *ExpansionItem {
	return v.value
}

func (v *NullableExpansionItem) Set(val *ExpansionItem) {
	v.value = val
	v.isSet = true
}

func (v NullableExpansionItem) IsSet() bool {
	return v.isSet
}

func (v *NullableExpansionItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExpansionItem(val *ExpansionItem) *NullableExpansionItem {
	return &NullableExpansionItem{value: val, isSet: true}
}

func (v NullableExpansionItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExpansionItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


