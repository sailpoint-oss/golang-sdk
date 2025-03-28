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

// checks if the Form type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Form{}

// Form struct for Form
type Form struct {
	// ID of the form
	Id NullableString `json:"id,omitempty"`
	// Name of the form
	Name NullableString `json:"name,omitempty"`
	// The form title
	Title *string `json:"title,omitempty"`
	// The form subtitle.
	Subtitle *string `json:"subtitle,omitempty"`
	// The name of the user that should be shown this form
	TargetUser *string `json:"targetUser,omitempty"`
	Sections []SectionDetails `json:"sections,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Form Form

// NewForm instantiates a new Form object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewForm() *Form {
	this := Form{}
	return &this
}

// NewFormWithDefaults instantiates a new Form object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormWithDefaults() *Form {
	this := Form{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Form) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Form) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *Form) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *Form) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *Form) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *Form) UnsetId() {
	o.Id.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Form) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Form) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *Form) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *Form) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *Form) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *Form) UnsetName() {
	o.Name.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *Form) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Form) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *Form) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *Form) SetTitle(v string) {
	o.Title = &v
}

// GetSubtitle returns the Subtitle field value if set, zero value otherwise.
func (o *Form) GetSubtitle() string {
	if o == nil || IsNil(o.Subtitle) {
		var ret string
		return ret
	}
	return *o.Subtitle
}

// GetSubtitleOk returns a tuple with the Subtitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Form) GetSubtitleOk() (*string, bool) {
	if o == nil || IsNil(o.Subtitle) {
		return nil, false
	}
	return o.Subtitle, true
}

// HasSubtitle returns a boolean if a field has been set.
func (o *Form) HasSubtitle() bool {
	if o != nil && !IsNil(o.Subtitle) {
		return true
	}

	return false
}

// SetSubtitle gets a reference to the given string and assigns it to the Subtitle field.
func (o *Form) SetSubtitle(v string) {
	o.Subtitle = &v
}

// GetTargetUser returns the TargetUser field value if set, zero value otherwise.
func (o *Form) GetTargetUser() string {
	if o == nil || IsNil(o.TargetUser) {
		var ret string
		return ret
	}
	return *o.TargetUser
}

// GetTargetUserOk returns a tuple with the TargetUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Form) GetTargetUserOk() (*string, bool) {
	if o == nil || IsNil(o.TargetUser) {
		return nil, false
	}
	return o.TargetUser, true
}

// HasTargetUser returns a boolean if a field has been set.
func (o *Form) HasTargetUser() bool {
	if o != nil && !IsNil(o.TargetUser) {
		return true
	}

	return false
}

// SetTargetUser gets a reference to the given string and assigns it to the TargetUser field.
func (o *Form) SetTargetUser(v string) {
	o.TargetUser = &v
}

// GetSections returns the Sections field value if set, zero value otherwise.
func (o *Form) GetSections() []SectionDetails {
	if o == nil || IsNil(o.Sections) {
		var ret []SectionDetails
		return ret
	}
	return o.Sections
}

// GetSectionsOk returns a tuple with the Sections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Form) GetSectionsOk() ([]SectionDetails, bool) {
	if o == nil || IsNil(o.Sections) {
		return nil, false
	}
	return o.Sections, true
}

// HasSections returns a boolean if a field has been set.
func (o *Form) HasSections() bool {
	if o != nil && !IsNil(o.Sections) {
		return true
	}

	return false
}

// SetSections gets a reference to the given []SectionDetails and assigns it to the Sections field.
func (o *Form) SetSections(v []SectionDetails) {
	o.Sections = v
}

func (o Form) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Form) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Subtitle) {
		toSerialize["subtitle"] = o.Subtitle
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

func (o *Form) UnmarshalJSON(data []byte) (err error) {
	varForm := _Form{}

	err = json.Unmarshal(data, &varForm)

	if err != nil {
		return err
	}

	*o = Form(varForm)

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

type NullableForm struct {
	value *Form
	isSet bool
}

func (v NullableForm) Get() *Form {
	return v.value
}

func (v *NullableForm) Set(val *Form) {
	v.value = val
	v.isSet = true
}

func (v NullableForm) IsSet() bool {
	return v.isSet
}

func (v *NullableForm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableForm(val *Form) *NullableForm {
	return &NullableForm{value: val, isSet: true}
}

func (v NullableForm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableForm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


