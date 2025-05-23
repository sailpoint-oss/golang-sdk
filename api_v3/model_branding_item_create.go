/*
Identity Security Cloud V3 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v3

import (
	"encoding/json"
	"os"
	"fmt"
)

// checks if the BrandingItemCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BrandingItemCreate{}

// BrandingItemCreate struct for BrandingItemCreate
type BrandingItemCreate struct {
	// name of branding item
	Name string `json:"name"`
	// product name
	ProductName NullableString `json:"productName"`
	// hex value of color for action button
	ActionButtonColor *string `json:"actionButtonColor,omitempty"`
	// hex value of color for link
	ActiveLinkColor *string `json:"activeLinkColor,omitempty"`
	// hex value of color for navigation bar
	NavigationColor *string `json:"navigationColor,omitempty"`
	// email from address
	EmailFromAddress *string `json:"emailFromAddress,omitempty"`
	// login information message
	LoginInformationalMessage *string `json:"loginInformationalMessage,omitempty"`
	// png file with logo
	FileStandard **os.File `json:"fileStandard,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BrandingItemCreate BrandingItemCreate

// NewBrandingItemCreate instantiates a new BrandingItemCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrandingItemCreate(name string, productName NullableString) *BrandingItemCreate {
	this := BrandingItemCreate{}
	this.Name = name
	this.ProductName = productName
	return &this
}

// NewBrandingItemCreateWithDefaults instantiates a new BrandingItemCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrandingItemCreateWithDefaults() *BrandingItemCreate {
	this := BrandingItemCreate{}
	return &this
}

// GetName returns the Name field value
func (o *BrandingItemCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BrandingItemCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BrandingItemCreate) SetName(v string) {
	o.Name = v
}

// GetProductName returns the ProductName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *BrandingItemCreate) GetProductName() string {
	if o == nil || o.ProductName.Get() == nil {
		var ret string
		return ret
	}

	return *o.ProductName.Get()
}

// GetProductNameOk returns a tuple with the ProductName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BrandingItemCreate) GetProductNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProductName.Get(), o.ProductName.IsSet()
}

// SetProductName sets field value
func (o *BrandingItemCreate) SetProductName(v string) {
	o.ProductName.Set(&v)
}

// GetActionButtonColor returns the ActionButtonColor field value if set, zero value otherwise.
func (o *BrandingItemCreate) GetActionButtonColor() string {
	if o == nil || IsNil(o.ActionButtonColor) {
		var ret string
		return ret
	}
	return *o.ActionButtonColor
}

// GetActionButtonColorOk returns a tuple with the ActionButtonColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandingItemCreate) GetActionButtonColorOk() (*string, bool) {
	if o == nil || IsNil(o.ActionButtonColor) {
		return nil, false
	}
	return o.ActionButtonColor, true
}

// HasActionButtonColor returns a boolean if a field has been set.
func (o *BrandingItemCreate) HasActionButtonColor() bool {
	if o != nil && !IsNil(o.ActionButtonColor) {
		return true
	}

	return false
}

// SetActionButtonColor gets a reference to the given string and assigns it to the ActionButtonColor field.
func (o *BrandingItemCreate) SetActionButtonColor(v string) {
	o.ActionButtonColor = &v
}

// GetActiveLinkColor returns the ActiveLinkColor field value if set, zero value otherwise.
func (o *BrandingItemCreate) GetActiveLinkColor() string {
	if o == nil || IsNil(o.ActiveLinkColor) {
		var ret string
		return ret
	}
	return *o.ActiveLinkColor
}

// GetActiveLinkColorOk returns a tuple with the ActiveLinkColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandingItemCreate) GetActiveLinkColorOk() (*string, bool) {
	if o == nil || IsNil(o.ActiveLinkColor) {
		return nil, false
	}
	return o.ActiveLinkColor, true
}

// HasActiveLinkColor returns a boolean if a field has been set.
func (o *BrandingItemCreate) HasActiveLinkColor() bool {
	if o != nil && !IsNil(o.ActiveLinkColor) {
		return true
	}

	return false
}

// SetActiveLinkColor gets a reference to the given string and assigns it to the ActiveLinkColor field.
func (o *BrandingItemCreate) SetActiveLinkColor(v string) {
	o.ActiveLinkColor = &v
}

// GetNavigationColor returns the NavigationColor field value if set, zero value otherwise.
func (o *BrandingItemCreate) GetNavigationColor() string {
	if o == nil || IsNil(o.NavigationColor) {
		var ret string
		return ret
	}
	return *o.NavigationColor
}

// GetNavigationColorOk returns a tuple with the NavigationColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandingItemCreate) GetNavigationColorOk() (*string, bool) {
	if o == nil || IsNil(o.NavigationColor) {
		return nil, false
	}
	return o.NavigationColor, true
}

// HasNavigationColor returns a boolean if a field has been set.
func (o *BrandingItemCreate) HasNavigationColor() bool {
	if o != nil && !IsNil(o.NavigationColor) {
		return true
	}

	return false
}

// SetNavigationColor gets a reference to the given string and assigns it to the NavigationColor field.
func (o *BrandingItemCreate) SetNavigationColor(v string) {
	o.NavigationColor = &v
}

// GetEmailFromAddress returns the EmailFromAddress field value if set, zero value otherwise.
func (o *BrandingItemCreate) GetEmailFromAddress() string {
	if o == nil || IsNil(o.EmailFromAddress) {
		var ret string
		return ret
	}
	return *o.EmailFromAddress
}

// GetEmailFromAddressOk returns a tuple with the EmailFromAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandingItemCreate) GetEmailFromAddressOk() (*string, bool) {
	if o == nil || IsNil(o.EmailFromAddress) {
		return nil, false
	}
	return o.EmailFromAddress, true
}

// HasEmailFromAddress returns a boolean if a field has been set.
func (o *BrandingItemCreate) HasEmailFromAddress() bool {
	if o != nil && !IsNil(o.EmailFromAddress) {
		return true
	}

	return false
}

// SetEmailFromAddress gets a reference to the given string and assigns it to the EmailFromAddress field.
func (o *BrandingItemCreate) SetEmailFromAddress(v string) {
	o.EmailFromAddress = &v
}

// GetLoginInformationalMessage returns the LoginInformationalMessage field value if set, zero value otherwise.
func (o *BrandingItemCreate) GetLoginInformationalMessage() string {
	if o == nil || IsNil(o.LoginInformationalMessage) {
		var ret string
		return ret
	}
	return *o.LoginInformationalMessage
}

// GetLoginInformationalMessageOk returns a tuple with the LoginInformationalMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandingItemCreate) GetLoginInformationalMessageOk() (*string, bool) {
	if o == nil || IsNil(o.LoginInformationalMessage) {
		return nil, false
	}
	return o.LoginInformationalMessage, true
}

// HasLoginInformationalMessage returns a boolean if a field has been set.
func (o *BrandingItemCreate) HasLoginInformationalMessage() bool {
	if o != nil && !IsNil(o.LoginInformationalMessage) {
		return true
	}

	return false
}

// SetLoginInformationalMessage gets a reference to the given string and assigns it to the LoginInformationalMessage field.
func (o *BrandingItemCreate) SetLoginInformationalMessage(v string) {
	o.LoginInformationalMessage = &v
}

// GetFileStandard returns the FileStandard field value if set, zero value otherwise.
func (o *BrandingItemCreate) GetFileStandard() *os.File {
	if o == nil || IsNil(o.FileStandard) {
		var ret *os.File
		return ret
	}
	return *o.FileStandard
}

// GetFileStandardOk returns a tuple with the FileStandard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandingItemCreate) GetFileStandardOk() (**os.File, bool) {
	if o == nil || IsNil(o.FileStandard) {
		return nil, false
	}
	return o.FileStandard, true
}

// HasFileStandard returns a boolean if a field has been set.
func (o *BrandingItemCreate) HasFileStandard() bool {
	if o != nil && !IsNil(o.FileStandard) {
		return true
	}

	return false
}

// SetFileStandard gets a reference to the given *os.File and assigns it to the FileStandard field.
func (o *BrandingItemCreate) SetFileStandard(v *os.File) {
	o.FileStandard = &v
}

func (o BrandingItemCreate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BrandingItemCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["productName"] = o.ProductName.Get()
	if !IsNil(o.ActionButtonColor) {
		toSerialize["actionButtonColor"] = o.ActionButtonColor
	}
	if !IsNil(o.ActiveLinkColor) {
		toSerialize["activeLinkColor"] = o.ActiveLinkColor
	}
	if !IsNil(o.NavigationColor) {
		toSerialize["navigationColor"] = o.NavigationColor
	}
	if !IsNil(o.EmailFromAddress) {
		toSerialize["emailFromAddress"] = o.EmailFromAddress
	}
	if !IsNil(o.LoginInformationalMessage) {
		toSerialize["loginInformationalMessage"] = o.LoginInformationalMessage
	}
	if !IsNil(o.FileStandard) {
		toSerialize["fileStandard"] = o.FileStandard
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BrandingItemCreate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"productName",
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

	varBrandingItemCreate := _BrandingItemCreate{}

	err = json.Unmarshal(data, &varBrandingItemCreate)

	if err != nil {
		return err
	}

	*o = BrandingItemCreate(varBrandingItemCreate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "productName")
		delete(additionalProperties, "actionButtonColor")
		delete(additionalProperties, "activeLinkColor")
		delete(additionalProperties, "navigationColor")
		delete(additionalProperties, "emailFromAddress")
		delete(additionalProperties, "loginInformationalMessage")
		delete(additionalProperties, "fileStandard")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBrandingItemCreate struct {
	value *BrandingItemCreate
	isSet bool
}

func (v NullableBrandingItemCreate) Get() *BrandingItemCreate {
	return v.value
}

func (v *NullableBrandingItemCreate) Set(val *BrandingItemCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableBrandingItemCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableBrandingItemCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrandingItemCreate(val *BrandingItemCreate) *NullableBrandingItemCreate {
	return &NullableBrandingItemCreate{value: val, isSet: true}
}

func (v NullableBrandingItemCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBrandingItemCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


