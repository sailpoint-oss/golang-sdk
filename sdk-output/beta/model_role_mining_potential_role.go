/*
IdentityNow Beta API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sailpointbetasdk

import (
	"encoding/json"
)

// RoleMiningPotentialRole struct for RoleMiningPotentialRole
type RoleMiningPotentialRole struct {
	// The creation date for a potential role
	CreateDate []string `json:"createDate,omitempty"`
	// The number of entitlements in a potential role.
	EntitlementCount *int32 `json:"entitlementCount,omitempty"`
	// The list of entitlement ids to be excluded.
	ExcludedEntitlements []string `json:"excludedEntitlements,omitempty"`
	// Id of the potential role
	Id *string `json:"id,omitempty"`
	// The number of identities in a potential role.
	IdentityCount *int32 `json:"identityCount,omitempty"`
	// Identity attribute distribution
	IdentityDistribution []RoleMiningIdentityDistribution `json:"identityDistribution,omitempty"`
	// The list of ids in a potential role.
	IdentityIds []string `json:"identityIds,omitempty"`
	// The modified date for a potential role
	ModifiedDate []string `json:"modifiedDate,omitempty"`
	// Name of the potential role
	Name *string `json:"name,omitempty"`
	Type *RoleMiningRoleType `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RoleMiningPotentialRole RoleMiningPotentialRole

// NewRoleMiningPotentialRole instantiates a new RoleMiningPotentialRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleMiningPotentialRole() *RoleMiningPotentialRole {
	this := RoleMiningPotentialRole{}
	return &this
}

// NewRoleMiningPotentialRoleWithDefaults instantiates a new RoleMiningPotentialRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleMiningPotentialRoleWithDefaults() *RoleMiningPotentialRole {
	this := RoleMiningPotentialRole{}
	return &this
}

// GetCreateDate returns the CreateDate field value if set, zero value otherwise.
func (o *RoleMiningPotentialRole) GetCreateDate() []string {
	if o == nil || isNil(o.CreateDate) {
		var ret []string
		return ret
	}
	return o.CreateDate
}

// GetCreateDateOk returns a tuple with the CreateDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningPotentialRole) GetCreateDateOk() ([]string, bool) {
	if o == nil || isNil(o.CreateDate) {
		return nil, false
	}
	return o.CreateDate, true
}

// HasCreateDate returns a boolean if a field has been set.
func (o *RoleMiningPotentialRole) HasCreateDate() bool {
	if o != nil && !isNil(o.CreateDate) {
		return true
	}

	return false
}

// SetCreateDate gets a reference to the given []string and assigns it to the CreateDate field.
func (o *RoleMiningPotentialRole) SetCreateDate(v []string) {
	o.CreateDate = v
}

// GetEntitlementCount returns the EntitlementCount field value if set, zero value otherwise.
func (o *RoleMiningPotentialRole) GetEntitlementCount() int32 {
	if o == nil || isNil(o.EntitlementCount) {
		var ret int32
		return ret
	}
	return *o.EntitlementCount
}

// GetEntitlementCountOk returns a tuple with the EntitlementCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningPotentialRole) GetEntitlementCountOk() (*int32, bool) {
	if o == nil || isNil(o.EntitlementCount) {
		return nil, false
	}
	return o.EntitlementCount, true
}

// HasEntitlementCount returns a boolean if a field has been set.
func (o *RoleMiningPotentialRole) HasEntitlementCount() bool {
	if o != nil && !isNil(o.EntitlementCount) {
		return true
	}

	return false
}

// SetEntitlementCount gets a reference to the given int32 and assigns it to the EntitlementCount field.
func (o *RoleMiningPotentialRole) SetEntitlementCount(v int32) {
	o.EntitlementCount = &v
}

// GetExcludedEntitlements returns the ExcludedEntitlements field value if set, zero value otherwise.
func (o *RoleMiningPotentialRole) GetExcludedEntitlements() []string {
	if o == nil || isNil(o.ExcludedEntitlements) {
		var ret []string
		return ret
	}
	return o.ExcludedEntitlements
}

// GetExcludedEntitlementsOk returns a tuple with the ExcludedEntitlements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningPotentialRole) GetExcludedEntitlementsOk() ([]string, bool) {
	if o == nil || isNil(o.ExcludedEntitlements) {
		return nil, false
	}
	return o.ExcludedEntitlements, true
}

// HasExcludedEntitlements returns a boolean if a field has been set.
func (o *RoleMiningPotentialRole) HasExcludedEntitlements() bool {
	if o != nil && !isNil(o.ExcludedEntitlements) {
		return true
	}

	return false
}

// SetExcludedEntitlements gets a reference to the given []string and assigns it to the ExcludedEntitlements field.
func (o *RoleMiningPotentialRole) SetExcludedEntitlements(v []string) {
	o.ExcludedEntitlements = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RoleMiningPotentialRole) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningPotentialRole) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RoleMiningPotentialRole) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RoleMiningPotentialRole) SetId(v string) {
	o.Id = &v
}

// GetIdentityCount returns the IdentityCount field value if set, zero value otherwise.
func (o *RoleMiningPotentialRole) GetIdentityCount() int32 {
	if o == nil || isNil(o.IdentityCount) {
		var ret int32
		return ret
	}
	return *o.IdentityCount
}

// GetIdentityCountOk returns a tuple with the IdentityCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningPotentialRole) GetIdentityCountOk() (*int32, bool) {
	if o == nil || isNil(o.IdentityCount) {
		return nil, false
	}
	return o.IdentityCount, true
}

// HasIdentityCount returns a boolean if a field has been set.
func (o *RoleMiningPotentialRole) HasIdentityCount() bool {
	if o != nil && !isNil(o.IdentityCount) {
		return true
	}

	return false
}

// SetIdentityCount gets a reference to the given int32 and assigns it to the IdentityCount field.
func (o *RoleMiningPotentialRole) SetIdentityCount(v int32) {
	o.IdentityCount = &v
}

// GetIdentityDistribution returns the IdentityDistribution field value if set, zero value otherwise.
func (o *RoleMiningPotentialRole) GetIdentityDistribution() []RoleMiningIdentityDistribution {
	if o == nil || isNil(o.IdentityDistribution) {
		var ret []RoleMiningIdentityDistribution
		return ret
	}
	return o.IdentityDistribution
}

// GetIdentityDistributionOk returns a tuple with the IdentityDistribution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningPotentialRole) GetIdentityDistributionOk() ([]RoleMiningIdentityDistribution, bool) {
	if o == nil || isNil(o.IdentityDistribution) {
		return nil, false
	}
	return o.IdentityDistribution, true
}

// HasIdentityDistribution returns a boolean if a field has been set.
func (o *RoleMiningPotentialRole) HasIdentityDistribution() bool {
	if o != nil && !isNil(o.IdentityDistribution) {
		return true
	}

	return false
}

// SetIdentityDistribution gets a reference to the given []RoleMiningIdentityDistribution and assigns it to the IdentityDistribution field.
func (o *RoleMiningPotentialRole) SetIdentityDistribution(v []RoleMiningIdentityDistribution) {
	o.IdentityDistribution = v
}

// GetIdentityIds returns the IdentityIds field value if set, zero value otherwise.
func (o *RoleMiningPotentialRole) GetIdentityIds() []string {
	if o == nil || isNil(o.IdentityIds) {
		var ret []string
		return ret
	}
	return o.IdentityIds
}

// GetIdentityIdsOk returns a tuple with the IdentityIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningPotentialRole) GetIdentityIdsOk() ([]string, bool) {
	if o == nil || isNil(o.IdentityIds) {
		return nil, false
	}
	return o.IdentityIds, true
}

// HasIdentityIds returns a boolean if a field has been set.
func (o *RoleMiningPotentialRole) HasIdentityIds() bool {
	if o != nil && !isNil(o.IdentityIds) {
		return true
	}

	return false
}

// SetIdentityIds gets a reference to the given []string and assigns it to the IdentityIds field.
func (o *RoleMiningPotentialRole) SetIdentityIds(v []string) {
	o.IdentityIds = v
}

// GetModifiedDate returns the ModifiedDate field value if set, zero value otherwise.
func (o *RoleMiningPotentialRole) GetModifiedDate() []string {
	if o == nil || isNil(o.ModifiedDate) {
		var ret []string
		return ret
	}
	return o.ModifiedDate
}

// GetModifiedDateOk returns a tuple with the ModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningPotentialRole) GetModifiedDateOk() ([]string, bool) {
	if o == nil || isNil(o.ModifiedDate) {
		return nil, false
	}
	return o.ModifiedDate, true
}

// HasModifiedDate returns a boolean if a field has been set.
func (o *RoleMiningPotentialRole) HasModifiedDate() bool {
	if o != nil && !isNil(o.ModifiedDate) {
		return true
	}

	return false
}

// SetModifiedDate gets a reference to the given []string and assigns it to the ModifiedDate field.
func (o *RoleMiningPotentialRole) SetModifiedDate(v []string) {
	o.ModifiedDate = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RoleMiningPotentialRole) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningPotentialRole) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RoleMiningPotentialRole) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RoleMiningPotentialRole) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RoleMiningPotentialRole) GetType() RoleMiningRoleType {
	if o == nil || isNil(o.Type) {
		var ret RoleMiningRoleType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningPotentialRole) GetTypeOk() (*RoleMiningRoleType, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RoleMiningPotentialRole) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given RoleMiningRoleType and assigns it to the Type field.
func (o *RoleMiningPotentialRole) SetType(v RoleMiningRoleType) {
	o.Type = &v
}

func (o RoleMiningPotentialRole) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CreateDate) {
		toSerialize["createDate"] = o.CreateDate
	}
	if !isNil(o.EntitlementCount) {
		toSerialize["entitlementCount"] = o.EntitlementCount
	}
	if !isNil(o.ExcludedEntitlements) {
		toSerialize["excludedEntitlements"] = o.ExcludedEntitlements
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.IdentityCount) {
		toSerialize["identityCount"] = o.IdentityCount
	}
	if !isNil(o.IdentityDistribution) {
		toSerialize["identityDistribution"] = o.IdentityDistribution
	}
	if !isNil(o.IdentityIds) {
		toSerialize["identityIds"] = o.IdentityIds
	}
	if !isNil(o.ModifiedDate) {
		toSerialize["modifiedDate"] = o.ModifiedDate
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RoleMiningPotentialRole) UnmarshalJSON(bytes []byte) (err error) {
	varRoleMiningPotentialRole := _RoleMiningPotentialRole{}

	if err = json.Unmarshal(bytes, &varRoleMiningPotentialRole); err == nil {
		*o = RoleMiningPotentialRole(varRoleMiningPotentialRole)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "createDate")
		delete(additionalProperties, "entitlementCount")
		delete(additionalProperties, "excludedEntitlements")
		delete(additionalProperties, "id")
		delete(additionalProperties, "identityCount")
		delete(additionalProperties, "identityDistribution")
		delete(additionalProperties, "identityIds")
		delete(additionalProperties, "modifiedDate")
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRoleMiningPotentialRole struct {
	value *RoleMiningPotentialRole
	isSet bool
}

func (v NullableRoleMiningPotentialRole) Get() *RoleMiningPotentialRole {
	return v.value
}

func (v *NullableRoleMiningPotentialRole) Set(val *RoleMiningPotentialRole) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleMiningPotentialRole) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleMiningPotentialRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleMiningPotentialRole(val *RoleMiningPotentialRole) *NullableRoleMiningPotentialRole {
	return &NullableRoleMiningPotentialRole{value: val, isSet: true}
}

func (v NullableRoleMiningPotentialRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleMiningPotentialRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


