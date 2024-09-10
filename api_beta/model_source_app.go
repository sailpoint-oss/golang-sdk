/*
Identity Security Cloud Beta API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_beta

import (
	"encoding/json"
	"time"
)

// checks if the SourceApp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SourceApp{}

// SourceApp struct for SourceApp
type SourceApp struct {
	// The source app id
	Id *string `json:"id,omitempty"`
	// The deprecated source app id
	CloudAppId *string `json:"cloudAppId,omitempty"`
	// The source app name
	Name *string `json:"name,omitempty"`
	// Time when the source app was created
	Created *time.Time `json:"created,omitempty"`
	// Time when the source app was last modified
	Modified *time.Time `json:"modified,omitempty"`
	// True if the source app is enabled
	Enabled *bool `json:"enabled,omitempty"`
	// True if the source app is provision request enabled
	ProvisionRequestEnabled *bool `json:"provisionRequestEnabled,omitempty"`
	// The description of the source app
	Description *string `json:"description,omitempty"`
	// True if the source app match all accounts
	MatchAllAccounts *bool `json:"matchAllAccounts,omitempty"`
	// True if the source app is shown in the app center
	AppCenterEnabled *bool `json:"appCenterEnabled,omitempty"`
	AccountSource NullableSourceAppAccountSource `json:"accountSource,omitempty"`
	// The owner of source app
	Owner NullableBaseReferenceDto1 `json:"owner,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SourceApp SourceApp

// NewSourceApp instantiates a new SourceApp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceApp() *SourceApp {
	this := SourceApp{}
	var enabled bool = false
	this.Enabled = &enabled
	var provisionRequestEnabled bool = false
	this.ProvisionRequestEnabled = &provisionRequestEnabled
	var matchAllAccounts bool = false
	this.MatchAllAccounts = &matchAllAccounts
	var appCenterEnabled bool = true
	this.AppCenterEnabled = &appCenterEnabled
	return &this
}

// NewSourceAppWithDefaults instantiates a new SourceApp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceAppWithDefaults() *SourceApp {
	this := SourceApp{}
	var enabled bool = false
	this.Enabled = &enabled
	var provisionRequestEnabled bool = false
	this.ProvisionRequestEnabled = &provisionRequestEnabled
	var matchAllAccounts bool = false
	this.MatchAllAccounts = &matchAllAccounts
	var appCenterEnabled bool = true
	this.AppCenterEnabled = &appCenterEnabled
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SourceApp) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceApp) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SourceApp) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SourceApp) SetId(v string) {
	o.Id = &v
}

// GetCloudAppId returns the CloudAppId field value if set, zero value otherwise.
func (o *SourceApp) GetCloudAppId() string {
	if o == nil || IsNil(o.CloudAppId) {
		var ret string
		return ret
	}
	return *o.CloudAppId
}

// GetCloudAppIdOk returns a tuple with the CloudAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceApp) GetCloudAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.CloudAppId) {
		return nil, false
	}
	return o.CloudAppId, true
}

// HasCloudAppId returns a boolean if a field has been set.
func (o *SourceApp) HasCloudAppId() bool {
	if o != nil && !IsNil(o.CloudAppId) {
		return true
	}

	return false
}

// SetCloudAppId gets a reference to the given string and assigns it to the CloudAppId field.
func (o *SourceApp) SetCloudAppId(v string) {
	o.CloudAppId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SourceApp) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceApp) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SourceApp) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SourceApp) SetName(v string) {
	o.Name = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *SourceApp) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceApp) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *SourceApp) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *SourceApp) SetCreated(v time.Time) {
	o.Created = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *SourceApp) GetModified() time.Time {
	if o == nil || IsNil(o.Modified) {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceApp) GetModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *SourceApp) HasModified() bool {
	if o != nil && !IsNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *SourceApp) SetModified(v time.Time) {
	o.Modified = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SourceApp) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceApp) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SourceApp) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SourceApp) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetProvisionRequestEnabled returns the ProvisionRequestEnabled field value if set, zero value otherwise.
func (o *SourceApp) GetProvisionRequestEnabled() bool {
	if o == nil || IsNil(o.ProvisionRequestEnabled) {
		var ret bool
		return ret
	}
	return *o.ProvisionRequestEnabled
}

// GetProvisionRequestEnabledOk returns a tuple with the ProvisionRequestEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceApp) GetProvisionRequestEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ProvisionRequestEnabled) {
		return nil, false
	}
	return o.ProvisionRequestEnabled, true
}

// HasProvisionRequestEnabled returns a boolean if a field has been set.
func (o *SourceApp) HasProvisionRequestEnabled() bool {
	if o != nil && !IsNil(o.ProvisionRequestEnabled) {
		return true
	}

	return false
}

// SetProvisionRequestEnabled gets a reference to the given bool and assigns it to the ProvisionRequestEnabled field.
func (o *SourceApp) SetProvisionRequestEnabled(v bool) {
	o.ProvisionRequestEnabled = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SourceApp) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceApp) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SourceApp) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SourceApp) SetDescription(v string) {
	o.Description = &v
}

// GetMatchAllAccounts returns the MatchAllAccounts field value if set, zero value otherwise.
func (o *SourceApp) GetMatchAllAccounts() bool {
	if o == nil || IsNil(o.MatchAllAccounts) {
		var ret bool
		return ret
	}
	return *o.MatchAllAccounts
}

// GetMatchAllAccountsOk returns a tuple with the MatchAllAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceApp) GetMatchAllAccountsOk() (*bool, bool) {
	if o == nil || IsNil(o.MatchAllAccounts) {
		return nil, false
	}
	return o.MatchAllAccounts, true
}

// HasMatchAllAccounts returns a boolean if a field has been set.
func (o *SourceApp) HasMatchAllAccounts() bool {
	if o != nil && !IsNil(o.MatchAllAccounts) {
		return true
	}

	return false
}

// SetMatchAllAccounts gets a reference to the given bool and assigns it to the MatchAllAccounts field.
func (o *SourceApp) SetMatchAllAccounts(v bool) {
	o.MatchAllAccounts = &v
}

// GetAppCenterEnabled returns the AppCenterEnabled field value if set, zero value otherwise.
func (o *SourceApp) GetAppCenterEnabled() bool {
	if o == nil || IsNil(o.AppCenterEnabled) {
		var ret bool
		return ret
	}
	return *o.AppCenterEnabled
}

// GetAppCenterEnabledOk returns a tuple with the AppCenterEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceApp) GetAppCenterEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.AppCenterEnabled) {
		return nil, false
	}
	return o.AppCenterEnabled, true
}

// HasAppCenterEnabled returns a boolean if a field has been set.
func (o *SourceApp) HasAppCenterEnabled() bool {
	if o != nil && !IsNil(o.AppCenterEnabled) {
		return true
	}

	return false
}

// SetAppCenterEnabled gets a reference to the given bool and assigns it to the AppCenterEnabled field.
func (o *SourceApp) SetAppCenterEnabled(v bool) {
	o.AppCenterEnabled = &v
}

// GetAccountSource returns the AccountSource field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourceApp) GetAccountSource() SourceAppAccountSource {
	if o == nil || IsNil(o.AccountSource.Get()) {
		var ret SourceAppAccountSource
		return ret
	}
	return *o.AccountSource.Get()
}

// GetAccountSourceOk returns a tuple with the AccountSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SourceApp) GetAccountSourceOk() (*SourceAppAccountSource, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountSource.Get(), o.AccountSource.IsSet()
}

// HasAccountSource returns a boolean if a field has been set.
func (o *SourceApp) HasAccountSource() bool {
	if o != nil && o.AccountSource.IsSet() {
		return true
	}

	return false
}

// SetAccountSource gets a reference to the given NullableSourceAppAccountSource and assigns it to the AccountSource field.
func (o *SourceApp) SetAccountSource(v SourceAppAccountSource) {
	o.AccountSource.Set(&v)
}
// SetAccountSourceNil sets the value for AccountSource to be an explicit nil
func (o *SourceApp) SetAccountSourceNil() {
	o.AccountSource.Set(nil)
}

// UnsetAccountSource ensures that no value is present for AccountSource, not even an explicit nil
func (o *SourceApp) UnsetAccountSource() {
	o.AccountSource.Unset()
}

// GetOwner returns the Owner field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourceApp) GetOwner() BaseReferenceDto1 {
	if o == nil || IsNil(o.Owner.Get()) {
		var ret BaseReferenceDto1
		return ret
	}
	return *o.Owner.Get()
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SourceApp) GetOwnerOk() (*BaseReferenceDto1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Owner.Get(), o.Owner.IsSet()
}

// HasOwner returns a boolean if a field has been set.
func (o *SourceApp) HasOwner() bool {
	if o != nil && o.Owner.IsSet() {
		return true
	}

	return false
}

// SetOwner gets a reference to the given NullableBaseReferenceDto1 and assigns it to the Owner field.
func (o *SourceApp) SetOwner(v BaseReferenceDto1) {
	o.Owner.Set(&v)
}
// SetOwnerNil sets the value for Owner to be an explicit nil
func (o *SourceApp) SetOwnerNil() {
	o.Owner.Set(nil)
}

// UnsetOwner ensures that no value is present for Owner, not even an explicit nil
func (o *SourceApp) UnsetOwner() {
	o.Owner.Unset()
}

func (o SourceApp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SourceApp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.CloudAppId) {
		toSerialize["cloudAppId"] = o.CloudAppId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Modified) {
		toSerialize["modified"] = o.Modified
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.ProvisionRequestEnabled) {
		toSerialize["provisionRequestEnabled"] = o.ProvisionRequestEnabled
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.MatchAllAccounts) {
		toSerialize["matchAllAccounts"] = o.MatchAllAccounts
	}
	if !IsNil(o.AppCenterEnabled) {
		toSerialize["appCenterEnabled"] = o.AppCenterEnabled
	}
	if o.AccountSource.IsSet() {
		toSerialize["accountSource"] = o.AccountSource.Get()
	}
	if o.Owner.IsSet() {
		toSerialize["owner"] = o.Owner.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SourceApp) UnmarshalJSON(data []byte) (err error) {
	varSourceApp := _SourceApp{}

	err = json.Unmarshal(data, &varSourceApp)

	if err != nil {
		return err
	}

	*o = SourceApp(varSourceApp)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "cloudAppId")
		delete(additionalProperties, "name")
		delete(additionalProperties, "created")
		delete(additionalProperties, "modified")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "provisionRequestEnabled")
		delete(additionalProperties, "description")
		delete(additionalProperties, "matchAllAccounts")
		delete(additionalProperties, "appCenterEnabled")
		delete(additionalProperties, "accountSource")
		delete(additionalProperties, "owner")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSourceApp struct {
	value *SourceApp
	isSet bool
}

func (v NullableSourceApp) Get() *SourceApp {
	return v.value
}

func (v *NullableSourceApp) Set(val *SourceApp) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceApp) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceApp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceApp(val *SourceApp) *NullableSourceApp {
	return &NullableSourceApp{value: val, isSet: true}
}

func (v NullableSourceApp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceApp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

