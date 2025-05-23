/*
Identity Security Cloud V2025 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: v2025
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2025

import (
	"encoding/json"
	"fmt"
)

// checks if the AccessItemEntitlementResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessItemEntitlementResponse{}

// AccessItemEntitlementResponse struct for AccessItemEntitlementResponse
type AccessItemEntitlementResponse struct {
	// the access item type. entitlement in this case
	AccessType *string `json:"accessType,omitempty"`
	// the access item id
	Id *string `json:"id,omitempty"`
	// the entitlement attribute
	Attribute *string `json:"attribute,omitempty"`
	// the associated value
	Value *string `json:"value,omitempty"`
	// the type of entitlement
	EntitlementType *string `json:"entitlementType,omitempty"`
	// the name of the source
	SourceName *string `json:"sourceName,omitempty"`
	// the id of the source
	SourceId *string `json:"sourceId,omitempty"`
	// the description for the entitlment
	Description *string `json:"description,omitempty"`
	// the display name of the identity
	DisplayName *string `json:"displayName,omitempty"`
	// indicates whether the entitlement is standalone
	Standalone bool `json:"standalone"`
	// indicates whether the entitlement is privileged
	Privileged bool `json:"privileged"`
	// indicates whether the entitlement is cloud governed
	CloudGoverned bool `json:"cloudGoverned"`
	AdditionalProperties map[string]interface{}
}

type _AccessItemEntitlementResponse AccessItemEntitlementResponse

// NewAccessItemEntitlementResponse instantiates a new AccessItemEntitlementResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessItemEntitlementResponse(standalone bool, privileged bool, cloudGoverned bool) *AccessItemEntitlementResponse {
	this := AccessItemEntitlementResponse{}
	this.Standalone = standalone
	this.Privileged = privileged
	this.CloudGoverned = cloudGoverned
	return &this
}

// NewAccessItemEntitlementResponseWithDefaults instantiates a new AccessItemEntitlementResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessItemEntitlementResponseWithDefaults() *AccessItemEntitlementResponse {
	this := AccessItemEntitlementResponse{}
	return &this
}

// GetAccessType returns the AccessType field value if set, zero value otherwise.
func (o *AccessItemEntitlementResponse) GetAccessType() string {
	if o == nil || IsNil(o.AccessType) {
		var ret string
		return ret
	}
	return *o.AccessType
}

// GetAccessTypeOk returns a tuple with the AccessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessItemEntitlementResponse) GetAccessTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AccessType) {
		return nil, false
	}
	return o.AccessType, true
}

// HasAccessType returns a boolean if a field has been set.
func (o *AccessItemEntitlementResponse) HasAccessType() bool {
	if o != nil && !IsNil(o.AccessType) {
		return true
	}

	return false
}

// SetAccessType gets a reference to the given string and assigns it to the AccessType field.
func (o *AccessItemEntitlementResponse) SetAccessType(v string) {
	o.AccessType = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AccessItemEntitlementResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessItemEntitlementResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AccessItemEntitlementResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AccessItemEntitlementResponse) SetId(v string) {
	o.Id = &v
}

// GetAttribute returns the Attribute field value if set, zero value otherwise.
func (o *AccessItemEntitlementResponse) GetAttribute() string {
	if o == nil || IsNil(o.Attribute) {
		var ret string
		return ret
	}
	return *o.Attribute
}

// GetAttributeOk returns a tuple with the Attribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessItemEntitlementResponse) GetAttributeOk() (*string, bool) {
	if o == nil || IsNil(o.Attribute) {
		return nil, false
	}
	return o.Attribute, true
}

// HasAttribute returns a boolean if a field has been set.
func (o *AccessItemEntitlementResponse) HasAttribute() bool {
	if o != nil && !IsNil(o.Attribute) {
		return true
	}

	return false
}

// SetAttribute gets a reference to the given string and assigns it to the Attribute field.
func (o *AccessItemEntitlementResponse) SetAttribute(v string) {
	o.Attribute = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *AccessItemEntitlementResponse) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessItemEntitlementResponse) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *AccessItemEntitlementResponse) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *AccessItemEntitlementResponse) SetValue(v string) {
	o.Value = &v
}

// GetEntitlementType returns the EntitlementType field value if set, zero value otherwise.
func (o *AccessItemEntitlementResponse) GetEntitlementType() string {
	if o == nil || IsNil(o.EntitlementType) {
		var ret string
		return ret
	}
	return *o.EntitlementType
}

// GetEntitlementTypeOk returns a tuple with the EntitlementType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessItemEntitlementResponse) GetEntitlementTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EntitlementType) {
		return nil, false
	}
	return o.EntitlementType, true
}

// HasEntitlementType returns a boolean if a field has been set.
func (o *AccessItemEntitlementResponse) HasEntitlementType() bool {
	if o != nil && !IsNil(o.EntitlementType) {
		return true
	}

	return false
}

// SetEntitlementType gets a reference to the given string and assigns it to the EntitlementType field.
func (o *AccessItemEntitlementResponse) SetEntitlementType(v string) {
	o.EntitlementType = &v
}

// GetSourceName returns the SourceName field value if set, zero value otherwise.
func (o *AccessItemEntitlementResponse) GetSourceName() string {
	if o == nil || IsNil(o.SourceName) {
		var ret string
		return ret
	}
	return *o.SourceName
}

// GetSourceNameOk returns a tuple with the SourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessItemEntitlementResponse) GetSourceNameOk() (*string, bool) {
	if o == nil || IsNil(o.SourceName) {
		return nil, false
	}
	return o.SourceName, true
}

// HasSourceName returns a boolean if a field has been set.
func (o *AccessItemEntitlementResponse) HasSourceName() bool {
	if o != nil && !IsNil(o.SourceName) {
		return true
	}

	return false
}

// SetSourceName gets a reference to the given string and assigns it to the SourceName field.
func (o *AccessItemEntitlementResponse) SetSourceName(v string) {
	o.SourceName = &v
}

// GetSourceId returns the SourceId field value if set, zero value otherwise.
func (o *AccessItemEntitlementResponse) GetSourceId() string {
	if o == nil || IsNil(o.SourceId) {
		var ret string
		return ret
	}
	return *o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessItemEntitlementResponse) GetSourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.SourceId) {
		return nil, false
	}
	return o.SourceId, true
}

// HasSourceId returns a boolean if a field has been set.
func (o *AccessItemEntitlementResponse) HasSourceId() bool {
	if o != nil && !IsNil(o.SourceId) {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given string and assigns it to the SourceId field.
func (o *AccessItemEntitlementResponse) SetSourceId(v string) {
	o.SourceId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AccessItemEntitlementResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessItemEntitlementResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AccessItemEntitlementResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AccessItemEntitlementResponse) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *AccessItemEntitlementResponse) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessItemEntitlementResponse) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *AccessItemEntitlementResponse) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *AccessItemEntitlementResponse) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetStandalone returns the Standalone field value
func (o *AccessItemEntitlementResponse) GetStandalone() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Standalone
}

// GetStandaloneOk returns a tuple with the Standalone field value
// and a boolean to check if the value has been set.
func (o *AccessItemEntitlementResponse) GetStandaloneOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Standalone, true
}

// SetStandalone sets field value
func (o *AccessItemEntitlementResponse) SetStandalone(v bool) {
	o.Standalone = v
}

// GetPrivileged returns the Privileged field value
func (o *AccessItemEntitlementResponse) GetPrivileged() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Privileged
}

// GetPrivilegedOk returns a tuple with the Privileged field value
// and a boolean to check if the value has been set.
func (o *AccessItemEntitlementResponse) GetPrivilegedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Privileged, true
}

// SetPrivileged sets field value
func (o *AccessItemEntitlementResponse) SetPrivileged(v bool) {
	o.Privileged = v
}

// GetCloudGoverned returns the CloudGoverned field value
func (o *AccessItemEntitlementResponse) GetCloudGoverned() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CloudGoverned
}

// GetCloudGovernedOk returns a tuple with the CloudGoverned field value
// and a boolean to check if the value has been set.
func (o *AccessItemEntitlementResponse) GetCloudGovernedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudGoverned, true
}

// SetCloudGoverned sets field value
func (o *AccessItemEntitlementResponse) SetCloudGoverned(v bool) {
	o.CloudGoverned = v
}

func (o AccessItemEntitlementResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessItemEntitlementResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessType) {
		toSerialize["accessType"] = o.AccessType
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Attribute) {
		toSerialize["attribute"] = o.Attribute
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.EntitlementType) {
		toSerialize["entitlementType"] = o.EntitlementType
	}
	if !IsNil(o.SourceName) {
		toSerialize["sourceName"] = o.SourceName
	}
	if !IsNil(o.SourceId) {
		toSerialize["sourceId"] = o.SourceId
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	toSerialize["standalone"] = o.Standalone
	toSerialize["privileged"] = o.Privileged
	toSerialize["cloudGoverned"] = o.CloudGoverned

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccessItemEntitlementResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"standalone",
		"privileged",
		"cloudGoverned",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAccessItemEntitlementResponse := _AccessItemEntitlementResponse{}

	err = json.Unmarshal(data, &varAccessItemEntitlementResponse)

	if err != nil {
		return err
	}

	*o = AccessItemEntitlementResponse(varAccessItemEntitlementResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "accessType")
		delete(additionalProperties, "id")
		delete(additionalProperties, "attribute")
		delete(additionalProperties, "value")
		delete(additionalProperties, "entitlementType")
		delete(additionalProperties, "sourceName")
		delete(additionalProperties, "sourceId")
		delete(additionalProperties, "description")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "standalone")
		delete(additionalProperties, "privileged")
		delete(additionalProperties, "cloudGoverned")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccessItemEntitlementResponse struct {
	value *AccessItemEntitlementResponse
	isSet bool
}

func (v NullableAccessItemEntitlementResponse) Get() *AccessItemEntitlementResponse {
	return v.value
}

func (v *NullableAccessItemEntitlementResponse) Set(val *AccessItemEntitlementResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessItemEntitlementResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessItemEntitlementResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessItemEntitlementResponse(val *AccessItemEntitlementResponse) *NullableAccessItemEntitlementResponse {
	return &NullableAccessItemEntitlementResponse{value: val, isSet: true}
}

func (v NullableAccessItemEntitlementResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessItemEntitlementResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


