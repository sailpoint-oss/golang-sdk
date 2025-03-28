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

// checks if the TemplateDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplateDto{}

// TemplateDto struct for TemplateDto
type TemplateDto struct {
	// The key of the template
	Key string `json:"key"`
	// The name of the Task Manager Subscription
	Name *string `json:"name,omitempty"`
	// The message medium. More mediums may be added in the future.
	Medium string `json:"medium"`
	// The locale for the message text, a BCP 47 language tag.
	Locale string `json:"locale"`
	// The subject line in the template
	Subject *string `json:"subject,omitempty"`
	// The header value is now located within the body field. If included with non-null values, will result in a 400.
	// Deprecated
	Header NullableString `json:"header,omitempty"`
	// The body in the template
	Body *string `json:"body,omitempty"`
	// The footer value is now located within the body field. If included with non-null values, will result in a 400.
	// Deprecated
	Footer NullableString `json:"footer,omitempty"`
	// The \"From:\" address in the template
	From *string `json:"from,omitempty"`
	// The \"Reply To\" line in the template
	ReplyTo *string `json:"replyTo,omitempty"`
	// The description in the template
	Description *string `json:"description,omitempty"`
	// This is auto-generated.
	Id *string `json:"id,omitempty"`
	// The time when this template is created. This is auto-generated.
	Created *SailPointTime `json:"created,omitempty"`
	// The time when this template was last modified. This is auto-generated.
	Modified *SailPointTime `json:"modified,omitempty"`
	SlackTemplate NullableString `json:"slackTemplate,omitempty"`
	TeamsTemplate NullableString `json:"teamsTemplate,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TemplateDto TemplateDto

// NewTemplateDto instantiates a new TemplateDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateDto(key string, medium string, locale string) *TemplateDto {
	this := TemplateDto{}
	this.Key = key
	this.Medium = medium
	this.Locale = locale
	return &this
}

// NewTemplateDtoWithDefaults instantiates a new TemplateDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateDtoWithDefaults() *TemplateDto {
	this := TemplateDto{}
	return &this
}

// GetKey returns the Key field value
func (o *TemplateDto) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *TemplateDto) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *TemplateDto) SetKey(v string) {
	o.Key = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TemplateDto) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateDto) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TemplateDto) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TemplateDto) SetName(v string) {
	o.Name = &v
}

// GetMedium returns the Medium field value
func (o *TemplateDto) GetMedium() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Medium
}

// GetMediumOk returns a tuple with the Medium field value
// and a boolean to check if the value has been set.
func (o *TemplateDto) GetMediumOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Medium, true
}

// SetMedium sets field value
func (o *TemplateDto) SetMedium(v string) {
	o.Medium = v
}

// GetLocale returns the Locale field value
func (o *TemplateDto) GetLocale() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value
// and a boolean to check if the value has been set.
func (o *TemplateDto) GetLocaleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Locale, true
}

// SetLocale sets field value
func (o *TemplateDto) SetLocale(v string) {
	o.Locale = v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *TemplateDto) GetSubject() string {
	if o == nil || IsNil(o.Subject) {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateDto) GetSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *TemplateDto) HasSubject() bool {
	if o != nil && !IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *TemplateDto) SetSubject(v string) {
	o.Subject = &v
}

// GetHeader returns the Header field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *TemplateDto) GetHeader() string {
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
func (o *TemplateDto) GetHeaderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Header.Get(), o.Header.IsSet()
}

// HasHeader returns a boolean if a field has been set.
func (o *TemplateDto) HasHeader() bool {
	if o != nil && o.Header.IsSet() {
		return true
	}

	return false
}

// SetHeader gets a reference to the given NullableString and assigns it to the Header field.
// Deprecated
func (o *TemplateDto) SetHeader(v string) {
	o.Header.Set(&v)
}
// SetHeaderNil sets the value for Header to be an explicit nil
func (o *TemplateDto) SetHeaderNil() {
	o.Header.Set(nil)
}

// UnsetHeader ensures that no value is present for Header, not even an explicit nil
func (o *TemplateDto) UnsetHeader() {
	o.Header.Unset()
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *TemplateDto) GetBody() string {
	if o == nil || IsNil(o.Body) {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateDto) GetBodyOk() (*string, bool) {
	if o == nil || IsNil(o.Body) {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *TemplateDto) HasBody() bool {
	if o != nil && !IsNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *TemplateDto) SetBody(v string) {
	o.Body = &v
}

// GetFooter returns the Footer field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *TemplateDto) GetFooter() string {
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
func (o *TemplateDto) GetFooterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Footer.Get(), o.Footer.IsSet()
}

// HasFooter returns a boolean if a field has been set.
func (o *TemplateDto) HasFooter() bool {
	if o != nil && o.Footer.IsSet() {
		return true
	}

	return false
}

// SetFooter gets a reference to the given NullableString and assigns it to the Footer field.
// Deprecated
func (o *TemplateDto) SetFooter(v string) {
	o.Footer.Set(&v)
}
// SetFooterNil sets the value for Footer to be an explicit nil
func (o *TemplateDto) SetFooterNil() {
	o.Footer.Set(nil)
}

// UnsetFooter ensures that no value is present for Footer, not even an explicit nil
func (o *TemplateDto) UnsetFooter() {
	o.Footer.Unset()
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *TemplateDto) GetFrom() string {
	if o == nil || IsNil(o.From) {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateDto) GetFromOk() (*string, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *TemplateDto) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *TemplateDto) SetFrom(v string) {
	o.From = &v
}

// GetReplyTo returns the ReplyTo field value if set, zero value otherwise.
func (o *TemplateDto) GetReplyTo() string {
	if o == nil || IsNil(o.ReplyTo) {
		var ret string
		return ret
	}
	return *o.ReplyTo
}

// GetReplyToOk returns a tuple with the ReplyTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateDto) GetReplyToOk() (*string, bool) {
	if o == nil || IsNil(o.ReplyTo) {
		return nil, false
	}
	return o.ReplyTo, true
}

// HasReplyTo returns a boolean if a field has been set.
func (o *TemplateDto) HasReplyTo() bool {
	if o != nil && !IsNil(o.ReplyTo) {
		return true
	}

	return false
}

// SetReplyTo gets a reference to the given string and assigns it to the ReplyTo field.
func (o *TemplateDto) SetReplyTo(v string) {
	o.ReplyTo = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TemplateDto) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateDto) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TemplateDto) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TemplateDto) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TemplateDto) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateDto) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TemplateDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TemplateDto) SetId(v string) {
	o.Id = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *TemplateDto) GetCreated() SailPointTime {
	if o == nil || IsNil(o.Created) {
		var ret SailPointTime
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateDto) GetCreatedOk() (*SailPointTime, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *TemplateDto) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given SailPointTime and assigns it to the Created field.
func (o *TemplateDto) SetCreated(v SailPointTime) {
	o.Created = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *TemplateDto) GetModified() SailPointTime {
	if o == nil || IsNil(o.Modified) {
		var ret SailPointTime
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateDto) GetModifiedOk() (*SailPointTime, bool) {
	if o == nil || IsNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *TemplateDto) HasModified() bool {
	if o != nil && !IsNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given SailPointTime and assigns it to the Modified field.
func (o *TemplateDto) SetModified(v SailPointTime) {
	o.Modified = &v
}

// GetSlackTemplate returns the SlackTemplate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateDto) GetSlackTemplate() string {
	if o == nil || IsNil(o.SlackTemplate.Get()) {
		var ret string
		return ret
	}
	return *o.SlackTemplate.Get()
}

// GetSlackTemplateOk returns a tuple with the SlackTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateDto) GetSlackTemplateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SlackTemplate.Get(), o.SlackTemplate.IsSet()
}

// HasSlackTemplate returns a boolean if a field has been set.
func (o *TemplateDto) HasSlackTemplate() bool {
	if o != nil && o.SlackTemplate.IsSet() {
		return true
	}

	return false
}

// SetSlackTemplate gets a reference to the given NullableString and assigns it to the SlackTemplate field.
func (o *TemplateDto) SetSlackTemplate(v string) {
	o.SlackTemplate.Set(&v)
}
// SetSlackTemplateNil sets the value for SlackTemplate to be an explicit nil
func (o *TemplateDto) SetSlackTemplateNil() {
	o.SlackTemplate.Set(nil)
}

// UnsetSlackTemplate ensures that no value is present for SlackTemplate, not even an explicit nil
func (o *TemplateDto) UnsetSlackTemplate() {
	o.SlackTemplate.Unset()
}

// GetTeamsTemplate returns the TeamsTemplate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateDto) GetTeamsTemplate() string {
	if o == nil || IsNil(o.TeamsTemplate.Get()) {
		var ret string
		return ret
	}
	return *o.TeamsTemplate.Get()
}

// GetTeamsTemplateOk returns a tuple with the TeamsTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateDto) GetTeamsTemplateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TeamsTemplate.Get(), o.TeamsTemplate.IsSet()
}

// HasTeamsTemplate returns a boolean if a field has been set.
func (o *TemplateDto) HasTeamsTemplate() bool {
	if o != nil && o.TeamsTemplate.IsSet() {
		return true
	}

	return false
}

// SetTeamsTemplate gets a reference to the given NullableString and assigns it to the TeamsTemplate field.
func (o *TemplateDto) SetTeamsTemplate(v string) {
	o.TeamsTemplate.Set(&v)
}
// SetTeamsTemplateNil sets the value for TeamsTemplate to be an explicit nil
func (o *TemplateDto) SetTeamsTemplateNil() {
	o.TeamsTemplate.Set(nil)
}

// UnsetTeamsTemplate ensures that no value is present for TeamsTemplate, not even an explicit nil
func (o *TemplateDto) UnsetTeamsTemplate() {
	o.TeamsTemplate.Unset()
}

func (o TemplateDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplateDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["medium"] = o.Medium
	toSerialize["locale"] = o.Locale
	if !IsNil(o.Subject) {
		toSerialize["subject"] = o.Subject
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
	if !IsNil(o.From) {
		toSerialize["from"] = o.From
	}
	if !IsNil(o.ReplyTo) {
		toSerialize["replyTo"] = o.ReplyTo
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Modified) {
		toSerialize["modified"] = o.Modified
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

func (o *TemplateDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"key",
		"medium",
		"locale",
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

	varTemplateDto := _TemplateDto{}

	err = json.Unmarshal(data, &varTemplateDto)

	if err != nil {
		return err
	}

	*o = TemplateDto(varTemplateDto)

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
		delete(additionalProperties, "id")
		delete(additionalProperties, "created")
		delete(additionalProperties, "modified")
		delete(additionalProperties, "slackTemplate")
		delete(additionalProperties, "teamsTemplate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTemplateDto struct {
	value *TemplateDto
	isSet bool
}

func (v NullableTemplateDto) Get() *TemplateDto {
	return v.value
}

func (v *NullableTemplateDto) Set(val *TemplateDto) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateDto) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateDto(val *TemplateDto) *NullableTemplateDto {
	return &NullableTemplateDto{value: val, isSet: true}
}

func (v NullableTemplateDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


