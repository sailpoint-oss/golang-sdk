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

// checks if the WorkgroupDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkgroupDto{}

// WorkgroupDto struct for WorkgroupDto
type WorkgroupDto struct {
	Owner *WorkgroupDtoOwner `json:"owner,omitempty"`
	// Governance group ID.
	Id *string `json:"id,omitempty"`
	// Governance group name.
	Name *string `json:"name,omitempty"`
	// Governance group description.
	Description *string `json:"description,omitempty"`
	// Number of members in the governance group.
	MemberCount *int64 `json:"memberCount,omitempty"`
	// Number of connections in the governance group.
	ConnectionCount *int64 `json:"connectionCount,omitempty"`
	Created *SailPointTime `json:"created,omitempty"`
	Modified *SailPointTime `json:"modified,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkgroupDto WorkgroupDto

// NewWorkgroupDto instantiates a new WorkgroupDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkgroupDto() *WorkgroupDto {
	this := WorkgroupDto{}
	return &this
}

// NewWorkgroupDtoWithDefaults instantiates a new WorkgroupDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkgroupDtoWithDefaults() *WorkgroupDto {
	this := WorkgroupDto{}
	return &this
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *WorkgroupDto) GetOwner() WorkgroupDtoOwner {
	if o == nil || IsNil(o.Owner) {
		var ret WorkgroupDtoOwner
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkgroupDto) GetOwnerOk() (*WorkgroupDtoOwner, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *WorkgroupDto) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given WorkgroupDtoOwner and assigns it to the Owner field.
func (o *WorkgroupDto) SetOwner(v WorkgroupDtoOwner) {
	o.Owner = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WorkgroupDto) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkgroupDto) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WorkgroupDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WorkgroupDto) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkgroupDto) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkgroupDto) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkgroupDto) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkgroupDto) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkgroupDto) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkgroupDto) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkgroupDto) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkgroupDto) SetDescription(v string) {
	o.Description = &v
}

// GetMemberCount returns the MemberCount field value if set, zero value otherwise.
func (o *WorkgroupDto) GetMemberCount() int64 {
	if o == nil || IsNil(o.MemberCount) {
		var ret int64
		return ret
	}
	return *o.MemberCount
}

// GetMemberCountOk returns a tuple with the MemberCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkgroupDto) GetMemberCountOk() (*int64, bool) {
	if o == nil || IsNil(o.MemberCount) {
		return nil, false
	}
	return o.MemberCount, true
}

// HasMemberCount returns a boolean if a field has been set.
func (o *WorkgroupDto) HasMemberCount() bool {
	if o != nil && !IsNil(o.MemberCount) {
		return true
	}

	return false
}

// SetMemberCount gets a reference to the given int64 and assigns it to the MemberCount field.
func (o *WorkgroupDto) SetMemberCount(v int64) {
	o.MemberCount = &v
}

// GetConnectionCount returns the ConnectionCount field value if set, zero value otherwise.
func (o *WorkgroupDto) GetConnectionCount() int64 {
	if o == nil || IsNil(o.ConnectionCount) {
		var ret int64
		return ret
	}
	return *o.ConnectionCount
}

// GetConnectionCountOk returns a tuple with the ConnectionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkgroupDto) GetConnectionCountOk() (*int64, bool) {
	if o == nil || IsNil(o.ConnectionCount) {
		return nil, false
	}
	return o.ConnectionCount, true
}

// HasConnectionCount returns a boolean if a field has been set.
func (o *WorkgroupDto) HasConnectionCount() bool {
	if o != nil && !IsNil(o.ConnectionCount) {
		return true
	}

	return false
}

// SetConnectionCount gets a reference to the given int64 and assigns it to the ConnectionCount field.
func (o *WorkgroupDto) SetConnectionCount(v int64) {
	o.ConnectionCount = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *WorkgroupDto) GetCreated() SailPointTime {
	if o == nil || IsNil(o.Created) {
		var ret SailPointTime
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkgroupDto) GetCreatedOk() (*SailPointTime, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *WorkgroupDto) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given SailPointTime and assigns it to the Created field.
func (o *WorkgroupDto) SetCreated(v SailPointTime) {
	o.Created = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *WorkgroupDto) GetModified() SailPointTime {
	if o == nil || IsNil(o.Modified) {
		var ret SailPointTime
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkgroupDto) GetModifiedOk() (*SailPointTime, bool) {
	if o == nil || IsNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *WorkgroupDto) HasModified() bool {
	if o != nil && !IsNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given SailPointTime and assigns it to the Modified field.
func (o *WorkgroupDto) SetModified(v SailPointTime) {
	o.Modified = &v
}

func (o WorkgroupDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkgroupDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.MemberCount) {
		toSerialize["memberCount"] = o.MemberCount
	}
	if !IsNil(o.ConnectionCount) {
		toSerialize["connectionCount"] = o.ConnectionCount
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Modified) {
		toSerialize["modified"] = o.Modified
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WorkgroupDto) UnmarshalJSON(data []byte) (err error) {
	varWorkgroupDto := _WorkgroupDto{}

	err = json.Unmarshal(data, &varWorkgroupDto)

	if err != nil {
		return err
	}

	*o = WorkgroupDto(varWorkgroupDto)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "owner")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "memberCount")
		delete(additionalProperties, "connectionCount")
		delete(additionalProperties, "created")
		delete(additionalProperties, "modified")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkgroupDto struct {
	value *WorkgroupDto
	isSet bool
}

func (v NullableWorkgroupDto) Get() *WorkgroupDto {
	return v.value
}

func (v *NullableWorkgroupDto) Set(val *WorkgroupDto) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkgroupDto) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkgroupDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkgroupDto(val *WorkgroupDto) *NullableWorkgroupDto {
	return &NullableWorkgroupDto{value: val, isSet: true}
}

func (v NullableWorkgroupDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkgroupDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


