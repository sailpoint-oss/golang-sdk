/*
Identity Security Cloud V2024 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: v2024
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2024

import (
	"encoding/json"
	"time"
)

// checks if the Product type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Product{}

// Product struct for Product
type Product struct {
	// Name of the Product
	ProductName *string `json:"productName,omitempty"`
	// URL of the Product
	Url *string `json:"url,omitempty"`
	// An identifier for a specific product-tenant combination
	ProductTenantId *string `json:"productTenantId,omitempty"`
	// Product region
	ProductRegion *string `json:"productRegion,omitempty"`
	// Right needed for the Product
	ProductRight *string `json:"productRight,omitempty"`
	// API URL of the Product
	ApiUrl NullableString `json:"apiUrl,omitempty"`
	Licenses []License `json:"licenses,omitempty"`
	// Additional attributes for a product
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	// Zone
	Zone *string `json:"zone,omitempty"`
	// Status of the product
	Status *string `json:"status,omitempty"`
	// Status datetime
	StatusDateTime *time.Time `json:"statusDateTime,omitempty"`
	// If there's a tenant provisioning failure then reason will have the description of error
	Reason *string `json:"reason,omitempty"`
	// Product could have additional notes added during tenant provisioning.
	Notes *string `json:"notes,omitempty"`
	// Date when the product was created
	DateCreated NullableTime `json:"dateCreated,omitempty"`
	// Date when the product was last updated
	LastUpdated NullableTime `json:"lastUpdated,omitempty"`
	// Type of org
	OrgType NullableString `json:"orgType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Product Product

// NewProduct instantiates a new Product object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProduct() *Product {
	this := Product{}
	return &this
}

// NewProductWithDefaults instantiates a new Product object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductWithDefaults() *Product {
	this := Product{}
	return &this
}

// GetProductName returns the ProductName field value if set, zero value otherwise.
func (o *Product) GetProductName() string {
	if o == nil || IsNil(o.ProductName) {
		var ret string
		return ret
	}
	return *o.ProductName
}

// GetProductNameOk returns a tuple with the ProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetProductNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProductName) {
		return nil, false
	}
	return o.ProductName, true
}

// HasProductName returns a boolean if a field has been set.
func (o *Product) HasProductName() bool {
	if o != nil && !IsNil(o.ProductName) {
		return true
	}

	return false
}

// SetProductName gets a reference to the given string and assigns it to the ProductName field.
func (o *Product) SetProductName(v string) {
	o.ProductName = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *Product) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *Product) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *Product) SetUrl(v string) {
	o.Url = &v
}

// GetProductTenantId returns the ProductTenantId field value if set, zero value otherwise.
func (o *Product) GetProductTenantId() string {
	if o == nil || IsNil(o.ProductTenantId) {
		var ret string
		return ret
	}
	return *o.ProductTenantId
}

// GetProductTenantIdOk returns a tuple with the ProductTenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetProductTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProductTenantId) {
		return nil, false
	}
	return o.ProductTenantId, true
}

// HasProductTenantId returns a boolean if a field has been set.
func (o *Product) HasProductTenantId() bool {
	if o != nil && !IsNil(o.ProductTenantId) {
		return true
	}

	return false
}

// SetProductTenantId gets a reference to the given string and assigns it to the ProductTenantId field.
func (o *Product) SetProductTenantId(v string) {
	o.ProductTenantId = &v
}

// GetProductRegion returns the ProductRegion field value if set, zero value otherwise.
func (o *Product) GetProductRegion() string {
	if o == nil || IsNil(o.ProductRegion) {
		var ret string
		return ret
	}
	return *o.ProductRegion
}

// GetProductRegionOk returns a tuple with the ProductRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetProductRegionOk() (*string, bool) {
	if o == nil || IsNil(o.ProductRegion) {
		return nil, false
	}
	return o.ProductRegion, true
}

// HasProductRegion returns a boolean if a field has been set.
func (o *Product) HasProductRegion() bool {
	if o != nil && !IsNil(o.ProductRegion) {
		return true
	}

	return false
}

// SetProductRegion gets a reference to the given string and assigns it to the ProductRegion field.
func (o *Product) SetProductRegion(v string) {
	o.ProductRegion = &v
}

// GetProductRight returns the ProductRight field value if set, zero value otherwise.
func (o *Product) GetProductRight() string {
	if o == nil || IsNil(o.ProductRight) {
		var ret string
		return ret
	}
	return *o.ProductRight
}

// GetProductRightOk returns a tuple with the ProductRight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetProductRightOk() (*string, bool) {
	if o == nil || IsNil(o.ProductRight) {
		return nil, false
	}
	return o.ProductRight, true
}

// HasProductRight returns a boolean if a field has been set.
func (o *Product) HasProductRight() bool {
	if o != nil && !IsNil(o.ProductRight) {
		return true
	}

	return false
}

// SetProductRight gets a reference to the given string and assigns it to the ProductRight field.
func (o *Product) SetProductRight(v string) {
	o.ProductRight = &v
}

// GetApiUrl returns the ApiUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Product) GetApiUrl() string {
	if o == nil || IsNil(o.ApiUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ApiUrl.Get()
}

// GetApiUrlOk returns a tuple with the ApiUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Product) GetApiUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApiUrl.Get(), o.ApiUrl.IsSet()
}

// HasApiUrl returns a boolean if a field has been set.
func (o *Product) HasApiUrl() bool {
	if o != nil && o.ApiUrl.IsSet() {
		return true
	}

	return false
}

// SetApiUrl gets a reference to the given NullableString and assigns it to the ApiUrl field.
func (o *Product) SetApiUrl(v string) {
	o.ApiUrl.Set(&v)
}
// SetApiUrlNil sets the value for ApiUrl to be an explicit nil
func (o *Product) SetApiUrlNil() {
	o.ApiUrl.Set(nil)
}

// UnsetApiUrl ensures that no value is present for ApiUrl, not even an explicit nil
func (o *Product) UnsetApiUrl() {
	o.ApiUrl.Unset()
}

// GetLicenses returns the Licenses field value if set, zero value otherwise.
func (o *Product) GetLicenses() []License {
	if o == nil || IsNil(o.Licenses) {
		var ret []License
		return ret
	}
	return o.Licenses
}

// GetLicensesOk returns a tuple with the Licenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetLicensesOk() ([]License, bool) {
	if o == nil || IsNil(o.Licenses) {
		return nil, false
	}
	return o.Licenses, true
}

// HasLicenses returns a boolean if a field has been set.
func (o *Product) HasLicenses() bool {
	if o != nil && !IsNil(o.Licenses) {
		return true
	}

	return false
}

// SetLicenses gets a reference to the given []License and assigns it to the Licenses field.
func (o *Product) SetLicenses(v []License) {
	o.Licenses = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *Product) GetAttributes() map[string]interface{} {
	if o == nil || IsNil(o.Attributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Attributes) {
		return map[string]interface{}{}, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *Product) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *Product) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetZone returns the Zone field value if set, zero value otherwise.
func (o *Product) GetZone() string {
	if o == nil || IsNil(o.Zone) {
		var ret string
		return ret
	}
	return *o.Zone
}

// GetZoneOk returns a tuple with the Zone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetZoneOk() (*string, bool) {
	if o == nil || IsNil(o.Zone) {
		return nil, false
	}
	return o.Zone, true
}

// HasZone returns a boolean if a field has been set.
func (o *Product) HasZone() bool {
	if o != nil && !IsNil(o.Zone) {
		return true
	}

	return false
}

// SetZone gets a reference to the given string and assigns it to the Zone field.
func (o *Product) SetZone(v string) {
	o.Zone = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Product) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Product) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Product) SetStatus(v string) {
	o.Status = &v
}

// GetStatusDateTime returns the StatusDateTime field value if set, zero value otherwise.
func (o *Product) GetStatusDateTime() time.Time {
	if o == nil || IsNil(o.StatusDateTime) {
		var ret time.Time
		return ret
	}
	return *o.StatusDateTime
}

// GetStatusDateTimeOk returns a tuple with the StatusDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetStatusDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StatusDateTime) {
		return nil, false
	}
	return o.StatusDateTime, true
}

// HasStatusDateTime returns a boolean if a field has been set.
func (o *Product) HasStatusDateTime() bool {
	if o != nil && !IsNil(o.StatusDateTime) {
		return true
	}

	return false
}

// SetStatusDateTime gets a reference to the given time.Time and assigns it to the StatusDateTime field.
func (o *Product) SetStatusDateTime(v time.Time) {
	o.StatusDateTime = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *Product) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *Product) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *Product) SetReason(v string) {
	o.Reason = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *Product) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *Product) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *Product) SetNotes(v string) {
	o.Notes = &v
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Product) GetDateCreated() time.Time {
	if o == nil || IsNil(o.DateCreated.Get()) {
		var ret time.Time
		return ret
	}
	return *o.DateCreated.Get()
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Product) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DateCreated.Get(), o.DateCreated.IsSet()
}

// HasDateCreated returns a boolean if a field has been set.
func (o *Product) HasDateCreated() bool {
	if o != nil && o.DateCreated.IsSet() {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given NullableTime and assigns it to the DateCreated field.
func (o *Product) SetDateCreated(v time.Time) {
	o.DateCreated.Set(&v)
}
// SetDateCreatedNil sets the value for DateCreated to be an explicit nil
func (o *Product) SetDateCreatedNil() {
	o.DateCreated.Set(nil)
}

// UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
func (o *Product) UnsetDateCreated() {
	o.DateCreated.Unset()
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Product) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Product) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *Product) HasLastUpdated() bool {
	if o != nil && o.LastUpdated.IsSet() {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given NullableTime and assigns it to the LastUpdated field.
func (o *Product) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}
// SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil
func (o *Product) SetLastUpdatedNil() {
	o.LastUpdated.Set(nil)
}

// UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
func (o *Product) UnsetLastUpdated() {
	o.LastUpdated.Unset()
}

// GetOrgType returns the OrgType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Product) GetOrgType() string {
	if o == nil || IsNil(o.OrgType.Get()) {
		var ret string
		return ret
	}
	return *o.OrgType.Get()
}

// GetOrgTypeOk returns a tuple with the OrgType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Product) GetOrgTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrgType.Get(), o.OrgType.IsSet()
}

// HasOrgType returns a boolean if a field has been set.
func (o *Product) HasOrgType() bool {
	if o != nil && o.OrgType.IsSet() {
		return true
	}

	return false
}

// SetOrgType gets a reference to the given NullableString and assigns it to the OrgType field.
func (o *Product) SetOrgType(v string) {
	o.OrgType.Set(&v)
}
// SetOrgTypeNil sets the value for OrgType to be an explicit nil
func (o *Product) SetOrgTypeNil() {
	o.OrgType.Set(nil)
}

// UnsetOrgType ensures that no value is present for OrgType, not even an explicit nil
func (o *Product) UnsetOrgType() {
	o.OrgType.Unset()
}

func (o Product) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Product) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProductName) {
		toSerialize["productName"] = o.ProductName
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.ProductTenantId) {
		toSerialize["productTenantId"] = o.ProductTenantId
	}
	if !IsNil(o.ProductRegion) {
		toSerialize["productRegion"] = o.ProductRegion
	}
	if !IsNil(o.ProductRight) {
		toSerialize["productRight"] = o.ProductRight
	}
	if o.ApiUrl.IsSet() {
		toSerialize["apiUrl"] = o.ApiUrl.Get()
	}
	if !IsNil(o.Licenses) {
		toSerialize["licenses"] = o.Licenses
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.Zone) {
		toSerialize["zone"] = o.Zone
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.StatusDateTime) {
		toSerialize["statusDateTime"] = o.StatusDateTime
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if o.DateCreated.IsSet() {
		toSerialize["dateCreated"] = o.DateCreated.Get()
	}
	if o.LastUpdated.IsSet() {
		toSerialize["lastUpdated"] = o.LastUpdated.Get()
	}
	if o.OrgType.IsSet() {
		toSerialize["orgType"] = o.OrgType.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Product) UnmarshalJSON(data []byte) (err error) {
	varProduct := _Product{}

	err = json.Unmarshal(data, &varProduct)

	if err != nil {
		return err
	}

	*o = Product(varProduct)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "productName")
		delete(additionalProperties, "url")
		delete(additionalProperties, "productTenantId")
		delete(additionalProperties, "productRegion")
		delete(additionalProperties, "productRight")
		delete(additionalProperties, "apiUrl")
		delete(additionalProperties, "licenses")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "zone")
		delete(additionalProperties, "status")
		delete(additionalProperties, "statusDateTime")
		delete(additionalProperties, "reason")
		delete(additionalProperties, "notes")
		delete(additionalProperties, "dateCreated")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "orgType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProduct struct {
	value *Product
	isSet bool
}

func (v NullableProduct) Get() *Product {
	return v.value
}

func (v *NullableProduct) Set(val *Product) {
	v.value = val
	v.isSet = true
}

func (v NullableProduct) IsSet() bool {
	return v.isSet
}

func (v *NullableProduct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProduct(val *Product) *NullableProduct {
	return &NullableProduct{value: val, isSet: true}
}

func (v NullableProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

