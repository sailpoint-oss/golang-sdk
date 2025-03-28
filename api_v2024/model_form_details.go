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

// checks if the FormDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FormDetails{}

// FormDetails struct for FormDetails
type FormDetails struct {
	// ID of the form
	Id NullableString `json:"id,omitempty"`
	// Name of the form
	Name NullableString `json:"name,omitempty"`
	// The form title
	Title NullableString `json:"title,omitempty"`
	// The form subtitle.
	Subtitle NullableString `json:"subtitle,omitempty"`
	// The name of the user that should be shown this form
	TargetUser *string `json:"targetUser,omitempty"`
	// Sections of the form
	Sections []SectionDetails `json:"sections,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FormDetails FormDetails

// NewFormDetails instantiates a new FormDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormDetails() *FormDetails {
	this := FormDetails{}
	return &this
}

// NewFormDetailsWithDefaults instantiates a new FormDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormDetailsWithDefaults() *FormDetails {
	this := FormDetails{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FormDetails) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FormDetails) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *FormDetails) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *FormDetails) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *FormDetails) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *FormDetails) UnsetId() {
	o.Id.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FormDetails) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FormDetails) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *FormDetails) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *FormDetails) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *FormDetails) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *FormDetails) UnsetName() {
	o.Name.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FormDetails) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FormDetails) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *FormDetails) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *FormDetails) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *FormDetails) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *FormDetails) UnsetTitle() {
	o.Title.Unset()
}

// GetSubtitle returns the Subtitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FormDetails) GetSubtitle() string {
	if o == nil || IsNil(o.Subtitle.Get()) {
		var ret string
		return ret
	}
	return *o.Subtitle.Get()
}

// GetSubtitleOk returns a tuple with the Subtitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FormDetails) GetSubtitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Subtitle.Get(), o.Subtitle.IsSet()
}

// HasSubtitle returns a boolean if a field has been set.
func (o *FormDetails) HasSubtitle() bool {
	if o != nil && o.Subtitle.IsSet() {
		return true
	}

	return false
}

// SetSubtitle gets a reference to the given NullableString and assigns it to the Subtitle field.
func (o *FormDetails) SetSubtitle(v string) {
	o.Subtitle.Set(&v)
}
// SetSubtitleNil sets the value for Subtitle to be an explicit nil
func (o *FormDetails) SetSubtitleNil() {
	o.Subtitle.Set(nil)
}

// UnsetSubtitle ensures that no value is present for Subtitle, not even an explicit nil
func (o *FormDetails) UnsetSubtitle() {
	o.Subtitle.Unset()
}

// GetTargetUser returns the TargetUser field value if set, zero value otherwise.
func (o *FormDetails) GetTargetUser() string {
	if o == nil || IsNil(o.TargetUser) {
		var ret string
		return ret
	}
	return *o.TargetUser
}

// GetTargetUserOk returns a tuple with the TargetUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDetails) GetTargetUserOk() (*string, bool) {
	if o == nil || IsNil(o.TargetUser) {
		return nil, false
	}
	return o.TargetUser, true
}

// HasTargetUser returns a boolean if a field has been set.
func (o *FormDetails) HasTargetUser() bool {
	if o != nil && !IsNil(o.TargetUser) {
		return true
	}

	return false
}

// SetTargetUser gets a reference to the given string and assigns it to the TargetUser field.
func (o *FormDetails) SetTargetUser(v string) {
	o.TargetUser = &v
}

// GetSections returns the Sections field value if set, zero value otherwise.
func (o *FormDetails) GetSections() []SectionDetails {
	if o == nil || IsNil(o.Sections) {
		var ret []SectionDetails
		return ret
	}
	return o.Sections
}

// GetSectionsOk returns a tuple with the Sections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDetails) GetSectionsOk() ([]SectionDetails, bool) {
	if o == nil || IsNil(o.Sections) {
		return nil, false
	}
	return o.Sections, true
}

// HasSections returns a boolean if a field has been set.
func (o *FormDetails) HasSections() bool {
	if o != nil && !IsNil(o.Sections) {
		return true
	}

	return false
}

// SetSections gets a reference to the given []SectionDetails and assigns it to the Sections field.
func (o *FormDetails) SetSections(v []SectionDetails) {
	o.Sections = v
}

func (o FormDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FormDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.Subtitle.IsSet() {
		toSerialize["subtitle"] = o.Subtitle.Get()
	}
	if !IsNil(o.TargetUser) {
		toSerialize["targetUser"] = o.TargetUser
	}
	if !IsNil(o.Sections) {
		toSerialize["sections"] = o.Sections
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FormDetails) UnmarshalJSON(data []byte) (err error) {
	varFormDetails := _FormDetails{}

	err = json.Unmarshal(data, &varFormDetails)

	if err != nil {
		return err
	}

	*o = FormDetails(varFormDetails)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "title")
		delete(additionalProperties, "subtitle")
		delete(additionalProperties, "targetUser")
		delete(additionalProperties, "sections")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFormDetails struct {
	value *FormDetails
	isSet bool
}

func (v NullableFormDetails) Get() *FormDetails {
	return v.value
}

func (v *NullableFormDetails) Set(val *FormDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableFormDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableFormDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormDetails(val *FormDetails) *NullableFormDetails {
	return &NullableFormDetails{value: val, isSet: true}
}

func (v NullableFormDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


