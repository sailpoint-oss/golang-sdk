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

// checks if the TemplateDtoDefault type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplateDtoDefault{}

// TemplateDtoDefault struct for TemplateDtoDefault
type TemplateDtoDefault struct {
	// The key of the default template
	Key *string `json:"key,omitempty"`
	// The name of the default template
	Name *string `json:"name,omitempty"`
	// The message medium. More mediums may be added in the future.
	Medium *string `json:"medium,omitempty"`
	// The locale for the message text, a BCP 47 language tag.
	Locale *string `json:"locale,omitempty"`
	// The subject of the default template
	Subject NullableString `json:"subject,omitempty"`
	// The header value is now located within the body field. If included with non-null values, will result in a 400.
	// Deprecated
	Header NullableString `json:"header,omitempty"`
	// The body of the default template
	Body *string `json:"body,omitempty"`
	// The footer value is now located within the body field. If included with non-null values, will result in a 400.
	// Deprecated
	Footer NullableString `json:"footer,omitempty"`
	// The \"From:\" address of the default template
	From NullableString `json:"from,omitempty"`
	// The \"Reply To\" field of the default template
	ReplyTo NullableString `json:"replyTo,omitempty"`
	// The description of the default template
	Description NullableString `json:"description,omitempty"`
	SlackTemplate NullableTemplateSlack `json:"slackTemplate,omitempty"`
	TeamsTemplate NullableTemplateTeams `json:"teamsTemplate,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TemplateDtoDefault TemplateDtoDefault

// NewTemplateDtoDefault instantiates a new TemplateDtoDefault object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateDtoDefault() *TemplateDtoDefault {
	this := TemplateDtoDefault{}
	return &this
}

// NewTemplateDtoDefaultWithDefaults instantiates a new TemplateDtoDefault object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateDtoDefaultWithDefaults() *TemplateDtoDefault {
	this := TemplateDtoDefault{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *TemplateDtoDefault) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateDtoDefault) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *TemplateDtoDefault) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *TemplateDtoDefault) SetKey(v string) {
	o.Key = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TemplateDtoDefault) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateDtoDefault) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TemplateDtoDefault) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TemplateDtoDefault) SetName(v string) {
	o.Name = &v
}

// GetMedium returns the Medium field value if set, zero value otherwise.
func (o *TemplateDtoDefault) GetMedium() string {
	if o == nil || IsNil(o.Medium) {
		var ret string
		return ret
	}
	return *o.Medium
}

// GetMediumOk returns a tuple with the Medium field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateDtoDefault) GetMediumOk() (*string, bool) {
	if o == nil || IsNil(o.Medium) {
		return nil, false
	}
	return o.Medium, true
}

// HasMedium returns a boolean if a field has been set.
func (o *TemplateDtoDefault) HasMedium() bool {
	if o != nil && !IsNil(o.Medium) {
		return true
	}

	return false
}

// SetMedium gets a reference to the given string and assigns it to the Medium field.
func (o *TemplateDtoDefault) SetMedium(v string) {
	o.Medium = &v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *TemplateDtoDefault) GetLocale() string {
	if o == nil || IsNil(o.Locale) {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateDtoDefault) GetLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.Locale) {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *TemplateDtoDefault) HasLocale() bool {
	if o != nil && !IsNil(o.Locale) {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *TemplateDtoDefault) SetLocale(v string) {
	o.Locale = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateDtoDefault) GetSubject() string {
	if o == nil || IsNil(o.Subject.Get()) {
		var ret string
		return ret
	}
	return *o.Subject.Get()
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateDtoDefault) GetSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Subject.Get(), o.Subject.IsSet()
}

// HasSubject returns a boolean if a field has been set.
func (o *TemplateDtoDefault) HasSubject() bool {
	if o != nil && o.Subject.IsSet() {
		return true
	}

	return false
}

// SetSubject gets a reference to the given NullableString and assigns it to the Subject field.
func (o *TemplateDtoDefault) SetSubject(v string) {
	o.Subject.Set(&v)
}
// SetSubjectNil sets the value for Subject to be an explicit nil
func (o *TemplateDtoDefault) SetSubjectNil() {
	o.Subject.Set(nil)
}

// UnsetSubject ensures that no value is present for Subject, not even an explicit nil
func (o *TemplateDtoDefault) UnsetSubject() {
	o.Subject.Unset()
}

// GetHeader returns the Header field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *TemplateDtoDefault) GetHeader() string {
	if o == nil || IsNil(o.Header.Get()) {
		var ret string
		return ret
	}
	return *o.Header.Get()
}

// GetHeaderOk returns a tuple with the Header field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
// Deprecated
func (o *TemplateDtoDefault) GetHeaderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Header.Get(), o.Header.IsSet()
}

// HasHeader returns a boolean if a field has been set.
func (o *TemplateDtoDefault) HasHeader() bool {
	if o != nil && o.Header.IsSet() {
		return true
	}

	return false
}

// SetHeader gets a reference to the given NullableString and assigns it to the Header field.
// Deprecated
func (o *TemplateDtoDefault) SetHeader(v string) {
	o.Header.Set(&v)
}
// SetHeaderNil sets the value for Header to be an explicit nil
func (o *TemplateDtoDefault) SetHeaderNil() {
	o.Header.Set(nil)
}

// UnsetHeader ensures that no value is present for Header, not even an explicit nil
func (o *TemplateDtoDefault) UnsetHeader() {
	o.Header.Unset()
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *TemplateDtoDefault) GetBody() string {
	if o == nil || IsNil(o.Body) {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateDtoDefault) GetBodyOk() (*string, bool) {
	if o == nil || IsNil(o.Body) {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *TemplateDtoDefault) HasBody() bool {
	if o != nil && !IsNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *TemplateDtoDefault) SetBody(v string) {
	o.Body = &v
}

// GetFooter returns the Footer field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *TemplateDtoDefault) GetFooter() string {
	if o == nil || IsNil(o.Footer.Get()) {
		var ret string
		return ret
	}
	return *o.Footer.Get()
}

// GetFooterOk returns a tuple with the Footer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
// Deprecated
func (o *TemplateDtoDefault) GetFooterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Footer.Get(), o.Footer.IsSet()
}

// HasFooter returns a boolean if a field has been set.
func (o *TemplateDtoDefault) HasFooter() bool {
	if o != nil && o.Footer.IsSet() {
		return true
	}

	return false
}

// SetFooter gets a reference to the given NullableString and assigns it to the Footer field.
// Deprecated
func (o *TemplateDtoDefault) SetFooter(v string) {
	o.Footer.Set(&v)
}
// SetFooterNil sets the value for Footer to be an explicit nil
func (o *TemplateDtoDefault) SetFooterNil() {
	o.Footer.Set(nil)
}

// UnsetFooter ensures that no value is present for Footer, not even an explicit nil
func (o *TemplateDtoDefault) UnsetFooter() {
	o.Footer.Unset()
}

// GetFrom returns the From field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateDtoDefault) GetFrom() string {
	if o == nil || IsNil(o.From.Get()) {
		var ret string
		return ret
	}
	return *o.From.Get()
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateDtoDefault) GetFromOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.From.Get(), o.From.IsSet()
}

// HasFrom returns a boolean if a field has been set.
func (o *TemplateDtoDefault) HasFrom() bool {
	if o != nil && o.From.IsSet() {
		return true
	}

	return false
}

// SetFrom gets a reference to the given NullableString and assigns it to the From field.
func (o *TemplateDtoDefault) SetFrom(v string) {
	o.From.Set(&v)
}
// SetFromNil sets the value for From to be an explicit nil
func (o *TemplateDtoDefault) SetFromNil() {
	o.From.Set(nil)
}

// UnsetFrom ensures that no value is present for From, not even an explicit nil
func (o *TemplateDtoDefault) UnsetFrom() {
	o.From.Unset()
}

// GetReplyTo returns the ReplyTo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateDtoDefault) GetReplyTo() string {
	if o == nil || IsNil(o.ReplyTo.Get()) {
		var ret string
		return ret
	}
	return *o.ReplyTo.Get()
}

// GetReplyToOk returns a tuple with the ReplyTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateDtoDefault) GetReplyToOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReplyTo.Get(), o.ReplyTo.IsSet()
}

// HasReplyTo returns a boolean if a field has been set.
func (o *TemplateDtoDefault) HasReplyTo() bool {
	if o != nil && o.ReplyTo.IsSet() {
		return true
	}

	return false
}

// SetReplyTo gets a reference to the given NullableString and assigns it to the ReplyTo field.
func (o *TemplateDtoDefault) SetReplyTo(v string) {
	o.ReplyTo.Set(&v)
}
// SetReplyToNil sets the value for ReplyTo to be an explicit nil
func (o *TemplateDtoDefault) SetReplyToNil() {
	o.ReplyTo.Set(nil)
}

// UnsetReplyTo ensures that no value is present for ReplyTo, not even an explicit nil
func (o *TemplateDtoDefault) UnsetReplyTo() {
	o.ReplyTo.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateDtoDefault) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateDtoDefault) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *TemplateDtoDefault) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *TemplateDtoDefault) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *TemplateDtoDefault) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *TemplateDtoDefault) UnsetDescription() {
	o.Description.Unset()
}

// GetSlackTemplate returns the SlackTemplate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateDtoDefault) GetSlackTemplate() TemplateSlack {
	if o == nil || IsNil(o.SlackTemplate.Get()) {
		var ret TemplateSlack
		return ret
	}
	return *o.SlackTemplate.Get()
}

// GetSlackTemplateOk returns a tuple with the SlackTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateDtoDefault) GetSlackTemplateOk() (*TemplateSlack, bool) {
	if o == nil {
		return nil, false
	}
	return o.SlackTemplate.Get(), o.SlackTemplate.IsSet()
}

// HasSlackTemplate returns a boolean if a field has been set.
func (o *TemplateDtoDefault) HasSlackTemplate() bool {
	if o != nil && o.SlackTemplate.IsSet() {
		return true
	}

	return false
}

// SetSlackTemplate gets a reference to the given NullableTemplateSlack and assigns it to the SlackTemplate field.
func (o *TemplateDtoDefault) SetSlackTemplate(v TemplateSlack) {
	o.SlackTemplate.Set(&v)
}
// SetSlackTemplateNil sets the value for SlackTemplate to be an explicit nil
func (o *TemplateDtoDefault) SetSlackTemplateNil() {
	o.SlackTemplate.Set(nil)
}

// UnsetSlackTemplate ensures that no value is present for SlackTemplate, not even an explicit nil
func (o *TemplateDtoDefault) UnsetSlackTemplate() {
	o.SlackTemplate.Unset()
}

// GetTeamsTemplate returns the TeamsTemplate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateDtoDefault) GetTeamsTemplate() TemplateTeams {
	if o == nil || IsNil(o.TeamsTemplate.Get()) {
		var ret TemplateTeams
		return ret
	}
	return *o.TeamsTemplate.Get()
}

// GetTeamsTemplateOk returns a tuple with the TeamsTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateDtoDefault) GetTeamsTemplateOk() (*TemplateTeams, bool) {
	if o == nil {
		return nil, false
	}
	return o.TeamsTemplate.Get(), o.TeamsTemplate.IsSet()
}

// HasTeamsTemplate returns a boolean if a field has been set.
func (o *TemplateDtoDefault) HasTeamsTemplate() bool {
	if o != nil && o.TeamsTemplate.IsSet() {
		return true
	}

	return false
}

// SetTeamsTemplate gets a reference to the given NullableTemplateTeams and assigns it to the TeamsTemplate field.
func (o *TemplateDtoDefault) SetTeamsTemplate(v TemplateTeams) {
	o.TeamsTemplate.Set(&v)
}
// SetTeamsTemplateNil sets the value for TeamsTemplate to be an explicit nil
func (o *TemplateDtoDefault) SetTeamsTemplateNil() {
	o.TeamsTemplate.Set(nil)
}

// UnsetTeamsTemplate ensures that no value is present for TeamsTemplate, not even an explicit nil
func (o *TemplateDtoDefault) UnsetTeamsTemplate() {
	o.TeamsTemplate.Unset()
}

func (o TemplateDtoDefault) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplateDtoDefault) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Medium) {
		toSerialize["medium"] = o.Medium
	}
	if !IsNil(o.Locale) {
		toSerialize["locale"] = o.Locale
	}
	if o.Subject.IsSet() {
		toSerialize["subject"] = o.Subject.Get()
	}
	if o.Header.IsSet() {
		toSerialize["header"] = o.Header.Get()
	}
	if !IsNil(o.Body) {
		toSerialize["body"] = o.Body
	}
	if o.Footer.IsSet() {
		toSerialize["footer"] = o.Footer.Get()
	}
	if o.From.IsSet() {
		toSerialize["from"] = o.From.Get()
	}
	if o.ReplyTo.IsSet() {
		toSerialize["replyTo"] = o.ReplyTo.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.SlackTemplate.IsSet() {
		toSerialize["slackTemplate"] = o.SlackTemplate.Get()
	}
	if o.TeamsTemplate.IsSet() {
		toSerialize["teamsTemplate"] = o.TeamsTemplate.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TemplateDtoDefault) UnmarshalJSON(data []byte) (err error) {
	varTemplateDtoDefault := _TemplateDtoDefault{}

	err = json.Unmarshal(data, &varTemplateDtoDefault)

	if err != nil {
		return err
	}

	*o = TemplateDtoDefault(varTemplateDtoDefault)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "key")
		delete(additionalProperties, "name")
		delete(additionalProperties, "medium")
		delete(additionalProperties, "locale")
		delete(additionalProperties, "subject")
		delete(additionalProperties, "header")
		delete(additionalProperties, "body")
		delete(additionalProperties, "footer")
		delete(additionalProperties, "from")
		delete(additionalProperties, "replyTo")
		delete(additionalProperties, "description")
		delete(additionalProperties, "slackTemplate")
		delete(additionalProperties, "teamsTemplate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTemplateDtoDefault struct {
	value *TemplateDtoDefault
	isSet bool
}

func (v NullableTemplateDtoDefault) Get() *TemplateDtoDefault {
	return v.value
}

func (v *NullableTemplateDtoDefault) Set(val *TemplateDtoDefault) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateDtoDefault) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateDtoDefault) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateDtoDefault(val *TemplateDtoDefault) *NullableTemplateDtoDefault {
	return &NullableTemplateDtoDefault{value: val, isSet: true}
}

func (v NullableTemplateDtoDefault) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateDtoDefault) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

