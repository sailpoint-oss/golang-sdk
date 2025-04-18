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

// checks if the AccessProfileDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessProfileDetails{}

// AccessProfileDetails struct for AccessProfileDetails
type AccessProfileDetails struct {
	// The ID of the Access Profile
	Id *string `json:"id,omitempty"`
	// Name of the Access Profile
	Name *string `json:"name,omitempty"`
	// Information about the Access Profile
	Description NullableString `json:"description,omitempty"`
	// Date the Access Profile was created
	Created *SailPointTime `json:"created,omitempty"`
	// Date the Access Profile was last modified.
	Modified *SailPointTime `json:"modified,omitempty"`
	// Whether the Access Profile is enabled.
	Disabled *bool `json:"disabled,omitempty"`
	// Whether the Access Profile is requestable via access request.
	Requestable *bool `json:"requestable,omitempty"`
	// Whether the Access Profile is protected.
	Protected *bool `json:"protected,omitempty"`
	// The owner ID of the Access Profile
	OwnerId *string `json:"ownerId,omitempty"`
	// The source ID of the Access Profile
	SourceId NullableInt64 `json:"sourceId,omitempty"`
	// The source name of the Access Profile
	SourceName *string `json:"sourceName,omitempty"`
	// The source app ID of the Access Profile
	AppId NullableInt64 `json:"appId,omitempty"`
	// The source app name of the Access Profile
	AppName NullableString `json:"appName,omitempty"`
	// The id of the application
	ApplicationId *string `json:"applicationId,omitempty"`
	// The type of the access profile
	Type *string `json:"type,omitempty"`
	// List of IDs of entitlements
	Entitlements []string `json:"entitlements,omitempty"`
	// The number of entitlements in the access profile
	EntitlementCount *int32 `json:"entitlementCount,omitempty"`
	// List of IDs of segments, if any, to which this Access Profile is assigned.
	Segments []string `json:"segments,omitempty"`
	// Comma-separated list of approval schemes. Each approval scheme is one of - manager - appOwner - sourceOwner - accessProfileOwner - workgroup:&lt;workgroupId&gt; 
	ApprovalSchemes *string `json:"approvalSchemes,omitempty"`
	// Comma-separated list of revoke request approval schemes. Each approval scheme is one of - manager - sourceOwner - accessProfileOwner - workgroup:&lt;workgroupId&gt; 
	RevokeRequestApprovalSchemes *string `json:"revokeRequestApprovalSchemes,omitempty"`
	// Whether the access profile require request comment for access request.
	RequestCommentsRequired *bool `json:"requestCommentsRequired,omitempty"`
	// Whether denied comment is required when access request is denied.
	DeniedCommentsRequired *bool `json:"deniedCommentsRequired,omitempty"`
	AccountSelector *AccessProfileDetailsAccountSelector `json:"accountSelector,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccessProfileDetails AccessProfileDetails

// NewAccessProfileDetails instantiates a new AccessProfileDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessProfileDetails() *AccessProfileDetails {
	this := AccessProfileDetails{}
	var disabled bool = true
	this.Disabled = &disabled
	var requestable bool = false
	this.Requestable = &requestable
	var protected bool = false
	this.Protected = &protected
	var requestCommentsRequired bool = false
	this.RequestCommentsRequired = &requestCommentsRequired
	var deniedCommentsRequired bool = false
	this.DeniedCommentsRequired = &deniedCommentsRequired
	return &this
}

// NewAccessProfileDetailsWithDefaults instantiates a new AccessProfileDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessProfileDetailsWithDefaults() *AccessProfileDetails {
	this := AccessProfileDetails{}
	var disabled bool = true
	this.Disabled = &disabled
	var requestable bool = false
	this.Requestable = &requestable
	var protected bool = false
	this.Protected = &protected
	var requestCommentsRequired bool = false
	this.RequestCommentsRequired = &requestCommentsRequired
	var deniedCommentsRequired bool = false
	this.DeniedCommentsRequired = &deniedCommentsRequired
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AccessProfileDetails) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessProfileDetails) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AccessProfileDetails) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AccessProfileDetails) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AccessProfileDetails) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessProfileDetails) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AccessProfileDetails) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AccessProfileDetails) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessProfileDetails) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessProfileDetails) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *AccessProfileDetails) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *AccessProfileDetails) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *AccessProfileDetails) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *AccessProfileDetails) UnsetDescription() {
	o.Description.Unset()
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *AccessProfileDetails) GetCreated() SailPointTime {
	if o == nil || IsNil(o.Created) {
		var ret SailPointTime
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessProfileDetails) GetCreatedOk() (*SailPointTime, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *AccessProfileDetails) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given SailPointTime and assigns it to the Created field.
func (o *AccessProfileDetails) SetCreated(v SailPointTime) {
	o.Created = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *AccessProfileDetails) GetModified() SailPointTime {
	if o == nil || IsNil(o.Modified) {
		var ret SailPointTime
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessProfileDetails) GetModifiedOk() (*SailPointTime, bool) {
	if o == nil || IsNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *AccessProfileDetails) HasModified() bool {
	if o != nil && !IsNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given SailPointTime and assigns it to the Modified field.
func (o *AccessProfileDetails) SetModified(v SailPointTime) {
	o.Modified = &v
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *AccessProfileDetails) GetDisabled() bool {
	if o == nil || IsNil(o.Disabled) {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessProfileDetails) GetDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Disabled) {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *AccessProfileDetails) HasDisabled() bool {
	if o != nil && !IsNil(o.Disabled) {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *AccessProfileDetails) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetRequestable returns the Requestable field value if set, zero value otherwise.
func (o *AccessProfileDetails) GetRequestable() bool {
	if o == nil || IsNil(o.Requestable) {
		var ret bool
		return ret
	}
	return *o.Requestable
}

// GetRequestableOk returns a tuple with the Requestable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessProfileDetails) GetRequestableOk() (*bool, bool) {
	if o == nil || IsNil(o.Requestable) {
		return nil, false
	}
	return o.Requestable, true
}

// HasRequestable returns a boolean if a field has been set.
func (o *AccessProfileDetails) HasRequestable() bool {
	if o != nil && !IsNil(o.Requestable) {
		return true
	}

	return false
}

// SetRequestable gets a reference to the given bool and assigns it to the Requestable field.
func (o *AccessProfileDetails) SetRequestable(v bool) {
	o.Requestable = &v
}

// GetProtected returns the Protected field value if set, zero value otherwise.
func (o *AccessProfileDetails) GetProtected() bool {
	if o == nil || IsNil(o.Protected) {
		var ret bool
		return ret
	}
	return *o.Protected
}

// GetProtectedOk returns a tuple with the Protected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessProfileDetails) GetProtectedOk() (*bool, bool) {
	if o == nil || IsNil(o.Protected) {
		return nil, false
	}
	return o.Protected, true
}

// HasProtected returns a boolean if a field has been set.
func (o *AccessProfileDetails) HasProtected() bool {
	if o != nil && !IsNil(o.Protected) {
		return true
	}

	return false
}

// SetProtected gets a reference to the given bool and assigns it to the Protected field.
func (o *AccessProfileDetails) SetProtected(v bool) {
	o.Protected = &v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *AccessProfileDetails) GetOwnerId() string {
	if o == nil || IsNil(o.OwnerId) {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessProfileDetails) GetOwnerIdOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerId) {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *AccessProfileDetails) HasOwnerId() bool {
	if o != nil && !IsNil(o.OwnerId) {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *AccessProfileDetails) SetOwnerId(v string) {
	o.OwnerId = &v
}

// GetSourceId returns the SourceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessProfileDetails) GetSourceId() int64 {
	if o == nil || IsNil(o.SourceId.Get()) {
		var ret int64
		return ret
	}
	return *o.SourceId.Get()
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessProfileDetails) GetSourceIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SourceId.Get(), o.SourceId.IsSet()
}

// HasSourceId returns a boolean if a field has been set.
func (o *AccessProfileDetails) HasSourceId() bool {
	if o != nil && o.SourceId.IsSet() {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given NullableInt64 and assigns it to the SourceId field.
func (o *AccessProfileDetails) SetSourceId(v int64) {
	o.SourceId.Set(&v)
}
// SetSourceIdNil sets the value for SourceId to be an explicit nil
func (o *AccessProfileDetails) SetSourceIdNil() {
	o.SourceId.Set(nil)
}

// UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
func (o *AccessProfileDetails) UnsetSourceId() {
	o.SourceId.Unset()
}

// GetSourceName returns the SourceName field value if set, zero value otherwise.
func (o *AccessProfileDetails) GetSourceName() string {
	if o == nil || IsNil(o.SourceName) {
		var ret string
		return ret
	}
	return *o.SourceName
}

// GetSourceNameOk returns a tuple with the SourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessProfileDetails) GetSourceNameOk() (*string, bool) {
	if o == nil || IsNil(o.SourceName) {
		return nil, false
	}
	return o.SourceName, true
}

// HasSourceName returns a boolean if a field has been set.
func (o *AccessProfileDetails) HasSourceName() bool {
	if o != nil && !IsNil(o.SourceName) {
		return true
	}

	return false
}

// SetSourceName gets a reference to the given string and assigns it to the SourceName field.
func (o *AccessProfileDetails) SetSourceName(v string) {
	o.SourceName = &v
}

// GetAppId returns the AppId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessProfileDetails) GetAppId() int64 {
	if o == nil || IsNil(o.AppId.Get()) {
		var ret int64
		return ret
	}
	return *o.AppId.Get()
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessProfileDetails) GetAppIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AppId.Get(), o.AppId.IsSet()
}

// HasAppId returns a boolean if a field has been set.
func (o *AccessProfileDetails) HasAppId() bool {
	if o != nil && o.AppId.IsSet() {
		return true
	}

	return false
}

// SetAppId gets a reference to the given NullableInt64 and assigns it to the AppId field.
func (o *AccessProfileDetails) SetAppId(v int64) {
	o.AppId.Set(&v)
}
// SetAppIdNil sets the value for AppId to be an explicit nil
func (o *AccessProfileDetails) SetAppIdNil() {
	o.AppId.Set(nil)
}

// UnsetAppId ensures that no value is present for AppId, not even an explicit nil
func (o *AccessProfileDetails) UnsetAppId() {
	o.AppId.Unset()
}

// GetAppName returns the AppName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessProfileDetails) GetAppName() string {
	if o == nil || IsNil(o.AppName.Get()) {
		var ret string
		return ret
	}
	return *o.AppName.Get()
}

// GetAppNameOk returns a tuple with the AppName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessProfileDetails) GetAppNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AppName.Get(), o.AppName.IsSet()
}

// HasAppName returns a boolean if a field has been set.
func (o *AccessProfileDetails) HasAppName() bool {
	if o != nil && o.AppName.IsSet() {
		return true
	}

	return false
}

// SetAppName gets a reference to the given NullableString and assigns it to the AppName field.
func (o *AccessProfileDetails) SetAppName(v string) {
	o.AppName.Set(&v)
}
// SetAppNameNil sets the value for AppName to be an explicit nil
func (o *AccessProfileDetails) SetAppNameNil() {
	o.AppName.Set(nil)
}

// UnsetAppName ensures that no value is present for AppName, not even an explicit nil
func (o *AccessProfileDetails) UnsetAppName() {
	o.AppName.Unset()
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise.
func (o *AccessProfileDetails) GetApplicationId() string {
	if o == nil || IsNil(o.ApplicationId) {
		var ret string
		return ret
	}
	return *o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessProfileDetails) GetApplicationIdOk() (*string, bool) {
	if o == nil || IsNil(o.ApplicationId) {
		return nil, false
	}
	return o.ApplicationId, true
}

// HasApplicationId returns a boolean if a field has been set.
func (o *AccessProfileDetails) HasApplicationId() bool {
	if o != nil && !IsNil(o.ApplicationId) {
		return true
	}

	return false
}

// SetApplicationId gets a reference to the given string and assigns it to the ApplicationId field.
func (o *AccessProfileDetails) SetApplicationId(v string) {
	o.ApplicationId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AccessProfileDetails) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessProfileDetails) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AccessProfileDetails) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AccessProfileDetails) SetType(v string) {
	o.Type = &v
}

// GetEntitlements returns the Entitlements field value if set, zero value otherwise.
func (o *AccessProfileDetails) GetEntitlements() []string {
	if o == nil || IsNil(o.Entitlements) {
		var ret []string
		return ret
	}
	return o.Entitlements
}

// GetEntitlementsOk returns a tuple with the Entitlements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessProfileDetails) GetEntitlementsOk() ([]string, bool) {
	if o == nil || IsNil(o.Entitlements) {
		return nil, false
	}
	return o.Entitlements, true
}

// HasEntitlements returns a boolean if a field has been set.
func (o *AccessProfileDetails) HasEntitlements() bool {
	if o != nil && !IsNil(o.Entitlements) {
		return true
	}

	return false
}

// SetEntitlements gets a reference to the given []string and assigns it to the Entitlements field.
func (o *AccessProfileDetails) SetEntitlements(v []string) {
	o.Entitlements = v
}

// GetEntitlementCount returns the EntitlementCount field value if set, zero value otherwise.
func (o *AccessProfileDetails) GetEntitlementCount() int32 {
	if o == nil || IsNil(o.EntitlementCount) {
		var ret int32
		return ret
	}
	return *o.EntitlementCount
}

// GetEntitlementCountOk returns a tuple with the EntitlementCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessProfileDetails) GetEntitlementCountOk() (*int32, bool) {
	if o == nil || IsNil(o.EntitlementCount) {
		return nil, false
	}
	return o.EntitlementCount, true
}

// HasEntitlementCount returns a boolean if a field has been set.
func (o *AccessProfileDetails) HasEntitlementCount() bool {
	if o != nil && !IsNil(o.EntitlementCount) {
		return true
	}

	return false
}

// SetEntitlementCount gets a reference to the given int32 and assigns it to the EntitlementCount field.
func (o *AccessProfileDetails) SetEntitlementCount(v int32) {
	o.EntitlementCount = &v
}

// GetSegments returns the Segments field value if set, zero value otherwise.
func (o *AccessProfileDetails) GetSegments() []string {
	if o == nil || IsNil(o.Segments) {
		var ret []string
		return ret
	}
	return o.Segments
}

// GetSegmentsOk returns a tuple with the Segments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessProfileDetails) GetSegmentsOk() ([]string, bool) {
	if o == nil || IsNil(o.Segments) {
		return nil, false
	}
	return o.Segments, true
}

// HasSegments returns a boolean if a field has been set.
func (o *AccessProfileDetails) HasSegments() bool {
	if o != nil && !IsNil(o.Segments) {
		return true
	}

	return false
}

// SetSegments gets a reference to the given []string and assigns it to the Segments field.
func (o *AccessProfileDetails) SetSegments(v []string) {
	o.Segments = v
}

// GetApprovalSchemes returns the ApprovalSchemes field value if set, zero value otherwise.
func (o *AccessProfileDetails) GetApprovalSchemes() string {
	if o == nil || IsNil(o.ApprovalSchemes) {
		var ret string
		return ret
	}
	return *o.ApprovalSchemes
}

// GetApprovalSchemesOk returns a tuple with the ApprovalSchemes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessProfileDetails) GetApprovalSchemesOk() (*string, bool) {
	if o == nil || IsNil(o.ApprovalSchemes) {
		return nil, false
	}
	return o.ApprovalSchemes, true
}

// HasApprovalSchemes returns a boolean if a field has been set.
func (o *AccessProfileDetails) HasApprovalSchemes() bool {
	if o != nil && !IsNil(o.ApprovalSchemes) {
		return true
	}

	return false
}

// SetApprovalSchemes gets a reference to the given string and assigns it to the ApprovalSchemes field.
func (o *AccessProfileDetails) SetApprovalSchemes(v string) {
	o.ApprovalSchemes = &v
}

// GetRevokeRequestApprovalSchemes returns the RevokeRequestApprovalSchemes field value if set, zero value otherwise.
func (o *AccessProfileDetails) GetRevokeRequestApprovalSchemes() string {
	if o == nil || IsNil(o.RevokeRequestApprovalSchemes) {
		var ret string
		return ret
	}
	return *o.RevokeRequestApprovalSchemes
}

// GetRevokeRequestApprovalSchemesOk returns a tuple with the RevokeRequestApprovalSchemes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessProfileDetails) GetRevokeRequestApprovalSchemesOk() (*string, bool) {
	if o == nil || IsNil(o.RevokeRequestApprovalSchemes) {
		return nil, false
	}
	return o.RevokeRequestApprovalSchemes, true
}

// HasRevokeRequestApprovalSchemes returns a boolean if a field has been set.
func (o *AccessProfileDetails) HasRevokeRequestApprovalSchemes() bool {
	if o != nil && !IsNil(o.RevokeRequestApprovalSchemes) {
		return true
	}

	return false
}

// SetRevokeRequestApprovalSchemes gets a reference to the given string and assigns it to the RevokeRequestApprovalSchemes field.
func (o *AccessProfileDetails) SetRevokeRequestApprovalSchemes(v string) {
	o.RevokeRequestApprovalSchemes = &v
}

// GetRequestCommentsRequired returns the RequestCommentsRequired field value if set, zero value otherwise.
func (o *AccessProfileDetails) GetRequestCommentsRequired() bool {
	if o == nil || IsNil(o.RequestCommentsRequired) {
		var ret bool
		return ret
	}
	return *o.RequestCommentsRequired
}

// GetRequestCommentsRequiredOk returns a tuple with the RequestCommentsRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessProfileDetails) GetRequestCommentsRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.RequestCommentsRequired) {
		return nil, false
	}
	return o.RequestCommentsRequired, true
}

// HasRequestCommentsRequired returns a boolean if a field has been set.
func (o *AccessProfileDetails) HasRequestCommentsRequired() bool {
	if o != nil && !IsNil(o.RequestCommentsRequired) {
		return true
	}

	return false
}

// SetRequestCommentsRequired gets a reference to the given bool and assigns it to the RequestCommentsRequired field.
func (o *AccessProfileDetails) SetRequestCommentsRequired(v bool) {
	o.RequestCommentsRequired = &v
}

// GetDeniedCommentsRequired returns the DeniedCommentsRequired field value if set, zero value otherwise.
func (o *AccessProfileDetails) GetDeniedCommentsRequired() bool {
	if o == nil || IsNil(o.DeniedCommentsRequired) {
		var ret bool
		return ret
	}
	return *o.DeniedCommentsRequired
}

// GetDeniedCommentsRequiredOk returns a tuple with the DeniedCommentsRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessProfileDetails) GetDeniedCommentsRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.DeniedCommentsRequired) {
		return nil, false
	}
	return o.DeniedCommentsRequired, true
}

// HasDeniedCommentsRequired returns a boolean if a field has been set.
func (o *AccessProfileDetails) HasDeniedCommentsRequired() bool {
	if o != nil && !IsNil(o.DeniedCommentsRequired) {
		return true
	}

	return false
}

// SetDeniedCommentsRequired gets a reference to the given bool and assigns it to the DeniedCommentsRequired field.
func (o *AccessProfileDetails) SetDeniedCommentsRequired(v bool) {
	o.DeniedCommentsRequired = &v
}

// GetAccountSelector returns the AccountSelector field value if set, zero value otherwise.
func (o *AccessProfileDetails) GetAccountSelector() AccessProfileDetailsAccountSelector {
	if o == nil || IsNil(o.AccountSelector) {
		var ret AccessProfileDetailsAccountSelector
		return ret
	}
	return *o.AccountSelector
}

// GetAccountSelectorOk returns a tuple with the AccountSelector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessProfileDetails) GetAccountSelectorOk() (*AccessProfileDetailsAccountSelector, bool) {
	if o == nil || IsNil(o.AccountSelector) {
		return nil, false
	}
	return o.AccountSelector, true
}

// HasAccountSelector returns a boolean if a field has been set.
func (o *AccessProfileDetails) HasAccountSelector() bool {
	if o != nil && !IsNil(o.AccountSelector) {
		return true
	}

	return false
}

// SetAccountSelector gets a reference to the given AccessProfileDetailsAccountSelector and assigns it to the AccountSelector field.
func (o *AccessProfileDetails) SetAccountSelector(v AccessProfileDetailsAccountSelector) {
	o.AccountSelector = &v
}

func (o AccessProfileDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessProfileDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Modified) {
		toSerialize["modified"] = o.Modified
	}
	if !IsNil(o.Disabled) {
		toSerialize["disabled"] = o.Disabled
	}
	if !IsNil(o.Requestable) {
		toSerialize["requestable"] = o.Requestable
	}
	if !IsNil(o.Protected) {
		toSerialize["protected"] = o.Protected
	}
	if !IsNil(o.OwnerId) {
		toSerialize["ownerId"] = o.OwnerId
	}
	if o.SourceId.IsSet() {
		toSerialize["sourceId"] = o.SourceId.Get()
	}
	if !IsNil(o.SourceName) {
		toSerialize["sourceName"] = o.SourceName
	}
	if o.AppId.IsSet() {
		toSerialize["appId"] = o.AppId.Get()
	}
	if o.AppName.IsSet() {
		toSerialize["appName"] = o.AppName.Get()
	}
	if !IsNil(o.ApplicationId) {
		toSerialize["applicationId"] = o.ApplicationId
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Entitlements) {
		toSerialize["entitlements"] = o.Entitlements
	}
	if !IsNil(o.EntitlementCount) {
		toSerialize["entitlementCount"] = o.EntitlementCount
	}
	if !IsNil(o.Segments) {
		toSerialize["segments"] = o.Segments
	}
	if !IsNil(o.ApprovalSchemes) {
		toSerialize["approvalSchemes"] = o.ApprovalSchemes
	}
	if !IsNil(o.RevokeRequestApprovalSchemes) {
		toSerialize["revokeRequestApprovalSchemes"] = o.RevokeRequestApprovalSchemes
	}
	if !IsNil(o.RequestCommentsRequired) {
		toSerialize["requestCommentsRequired"] = o.RequestCommentsRequired
	}
	if !IsNil(o.DeniedCommentsRequired) {
		toSerialize["deniedCommentsRequired"] = o.DeniedCommentsRequired
	}
	if !IsNil(o.AccountSelector) {
		toSerialize["accountSelector"] = o.AccountSelector
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccessProfileDetails) UnmarshalJSON(data []byte) (err error) {
	varAccessProfileDetails := _AccessProfileDetails{}

	err = json.Unmarshal(data, &varAccessProfileDetails)

	if err != nil {
		return err
	}

	*o = AccessProfileDetails(varAccessProfileDetails)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "created")
		delete(additionalProperties, "modified")
		delete(additionalProperties, "disabled")
		delete(additionalProperties, "requestable")
		delete(additionalProperties, "protected")
		delete(additionalProperties, "ownerId")
		delete(additionalProperties, "sourceId")
		delete(additionalProperties, "sourceName")
		delete(additionalProperties, "appId")
		delete(additionalProperties, "appName")
		delete(additionalProperties, "applicationId")
		delete(additionalProperties, "type")
		delete(additionalProperties, "entitlements")
		delete(additionalProperties, "entitlementCount")
		delete(additionalProperties, "segments")
		delete(additionalProperties, "approvalSchemes")
		delete(additionalProperties, "revokeRequestApprovalSchemes")
		delete(additionalProperties, "requestCommentsRequired")
		delete(additionalProperties, "deniedCommentsRequired")
		delete(additionalProperties, "accountSelector")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccessProfileDetails struct {
	value *AccessProfileDetails
	isSet bool
}

func (v NullableAccessProfileDetails) Get() *AccessProfileDetails {
	return v.value
}

func (v *NullableAccessProfileDetails) Set(val *AccessProfileDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessProfileDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessProfileDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessProfileDetails(val *AccessProfileDetails) *NullableAccessProfileDetails {
	return &NullableAccessProfileDetails{value: val, isSet: true}
}

func (v NullableAccessProfileDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessProfileDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


