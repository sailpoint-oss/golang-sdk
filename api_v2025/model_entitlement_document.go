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

// checks if the EntitlementDocument type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitlementDocument{}

// EntitlementDocument Entitlement
type EntitlementDocument struct {
	// ID of the referenced object.
	Id string `json:"id"`
	// The human readable name of the referenced object.
	Name string `json:"name"`
	// ISO-8601 date-time referring to the time when the object was last modified.
	Modified NullableTime `json:"modified,omitempty"`
	// ISO-8601 date-time referring to the date-time when object was queued to be synced into search database for use in the search API.   This date-time changes anytime there is an update to the object, which triggers a synchronization event being sent to the search database.  There may be some delay between the `synced` time and the time when the updated data is actually available in the search API. 
	Synced *string `json:"synced,omitempty"`
	// Entitlement's display name.
	DisplayName *string `json:"displayName,omitempty"`
	Source *EntitlementDocumentAllOfSource `json:"source,omitempty"`
	// Segments with the entitlement.
	Segments []BaseSegment `json:"segments,omitempty"`
	// Number of segments with the role.
	SegmentCount *int32 `json:"segmentCount,omitempty"`
	// Indicates whether the entitlement is requestable.
	Requestable *bool `json:"requestable,omitempty"`
	// Indicates whether the entitlement is cloud governed.
	CloudGoverned *bool `json:"cloudGoverned,omitempty"`
	// ISO-8601 date-time referring to the time when the object was created.
	Created NullableTime `json:"created,omitempty"`
	// Indicates whether the entitlement is privileged.
	Privileged *bool `json:"privileged,omitempty"`
	// Tags that have been applied to the object.
	Tags []string `json:"tags,omitempty"`
	// Attribute information for the entitlement.
	Attribute *string `json:"attribute,omitempty"`
	// Value of the entitlement.
	Value *string `json:"value,omitempty"`
	// Source schema object type of the entitlement.
	SourceSchemaObjectType *string `json:"sourceSchemaObjectType,omitempty"`
	// Schema type of the entitlement.
	Schema *string `json:"schema,omitempty"`
	// Read-only calculated hash value of an entitlement.
	Hash *string `json:"hash,omitempty"`
	// Attributes of the entitlement.
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	// Truncated attributes of the entitlement.
	TruncatedAttributes []string `json:"truncatedAttributes,omitempty"`
	// Indicates whether the entitlement contains data access.
	ContainsDataAccess *bool `json:"containsDataAccess,omitempty"`
	ManuallyUpdatedFields NullableEntitlementDocumentAllOfManuallyUpdatedFields `json:"manuallyUpdatedFields,omitempty"`
	Permissions []EntitlementDocumentAllOfPermissions `json:"permissions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EntitlementDocument EntitlementDocument

// NewEntitlementDocument instantiates a new EntitlementDocument object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementDocument(id string, name string) *EntitlementDocument {
	this := EntitlementDocument{}
	this.Id = id
	this.Name = name
	var requestable bool = false
	this.Requestable = &requestable
	var cloudGoverned bool = false
	this.CloudGoverned = &cloudGoverned
	var privileged bool = false
	this.Privileged = &privileged
	var containsDataAccess bool = false
	this.ContainsDataAccess = &containsDataAccess
	return &this
}

// NewEntitlementDocumentWithDefaults instantiates a new EntitlementDocument object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementDocumentWithDefaults() *EntitlementDocument {
	this := EntitlementDocument{}
	var requestable bool = false
	this.Requestable = &requestable
	var cloudGoverned bool = false
	this.CloudGoverned = &cloudGoverned
	var privileged bool = false
	this.Privileged = &privileged
	var containsDataAccess bool = false
	this.ContainsDataAccess = &containsDataAccess
	return &this
}

// GetId returns the Id field value
func (o *EntitlementDocument) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EntitlementDocument) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EntitlementDocument) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *EntitlementDocument) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EntitlementDocument) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EntitlementDocument) SetName(v string) {
	o.Name = v
}

// GetModified returns the Modified field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EntitlementDocument) GetModified() SailPointTime {
	if o == nil || IsNil(o.Modified.Get()) {
		var ret SailPointTime
		return ret
	}
	return *o.Modified.Get()
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EntitlementDocument) GetModifiedOk() (*SailPointTime, bool) {
	if o == nil {
		return nil, false
	}
	return o.Modified.Get(), o.Modified.IsSet()
}

// HasModified returns a boolean if a field has been set.
func (o *EntitlementDocument) HasModified() bool {
	if o != nil && o.Modified.IsSet() {
		return true
	}

	return false
}

// SetModified gets a reference to the given NullableTime and assigns it to the Modified field.
func (o *EntitlementDocument) SetModified(v SailPointTime) {
	o.Modified.Set(&v)
}
// SetModifiedNil sets the value for Modified to be an explicit nil
func (o *EntitlementDocument) SetModifiedNil() {
	o.Modified.Set(nil)
}

// UnsetModified ensures that no value is present for Modified, not even an explicit nil
func (o *EntitlementDocument) UnsetModified() {
	o.Modified.Unset()
}

// GetSynced returns the Synced field value if set, zero value otherwise.
func (o *EntitlementDocument) GetSynced() string {
	if o == nil || IsNil(o.Synced) {
		var ret string
		return ret
	}
	return *o.Synced
}

// GetSyncedOk returns a tuple with the Synced field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDocument) GetSyncedOk() (*string, bool) {
	if o == nil || IsNil(o.Synced) {
		return nil, false
	}
	return o.Synced, true
}

// HasSynced returns a boolean if a field has been set.
func (o *EntitlementDocument) HasSynced() bool {
	if o != nil && !IsNil(o.Synced) {
		return true
	}

	return false
}

// SetSynced gets a reference to the given string and assigns it to the Synced field.
func (o *EntitlementDocument) SetSynced(v string) {
	o.Synced = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *EntitlementDocument) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDocument) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *EntitlementDocument) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *EntitlementDocument) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *EntitlementDocument) GetSource() EntitlementDocumentAllOfSource {
	if o == nil || IsNil(o.Source) {
		var ret EntitlementDocumentAllOfSource
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDocument) GetSourceOk() (*EntitlementDocumentAllOfSource, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *EntitlementDocument) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given EntitlementDocumentAllOfSource and assigns it to the Source field.
func (o *EntitlementDocument) SetSource(v EntitlementDocumentAllOfSource) {
	o.Source = &v
}

// GetSegments returns the Segments field value if set, zero value otherwise.
func (o *EntitlementDocument) GetSegments() []BaseSegment {
	if o == nil || IsNil(o.Segments) {
		var ret []BaseSegment
		return ret
	}
	return o.Segments
}

// GetSegmentsOk returns a tuple with the Segments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDocument) GetSegmentsOk() ([]BaseSegment, bool) {
	if o == nil || IsNil(o.Segments) {
		return nil, false
	}
	return o.Segments, true
}

// HasSegments returns a boolean if a field has been set.
func (o *EntitlementDocument) HasSegments() bool {
	if o != nil && !IsNil(o.Segments) {
		return true
	}

	return false
}

// SetSegments gets a reference to the given []BaseSegment and assigns it to the Segments field.
func (o *EntitlementDocument) SetSegments(v []BaseSegment) {
	o.Segments = v
}

// GetSegmentCount returns the SegmentCount field value if set, zero value otherwise.
func (o *EntitlementDocument) GetSegmentCount() int32 {
	if o == nil || IsNil(o.SegmentCount) {
		var ret int32
		return ret
	}
	return *o.SegmentCount
}

// GetSegmentCountOk returns a tuple with the SegmentCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDocument) GetSegmentCountOk() (*int32, bool) {
	if o == nil || IsNil(o.SegmentCount) {
		return nil, false
	}
	return o.SegmentCount, true
}

// HasSegmentCount returns a boolean if a field has been set.
func (o *EntitlementDocument) HasSegmentCount() bool {
	if o != nil && !IsNil(o.SegmentCount) {
		return true
	}

	return false
}

// SetSegmentCount gets a reference to the given int32 and assigns it to the SegmentCount field.
func (o *EntitlementDocument) SetSegmentCount(v int32) {
	o.SegmentCount = &v
}

// GetRequestable returns the Requestable field value if set, zero value otherwise.
func (o *EntitlementDocument) GetRequestable() bool {
	if o == nil || IsNil(o.Requestable) {
		var ret bool
		return ret
	}
	return *o.Requestable
}

// GetRequestableOk returns a tuple with the Requestable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDocument) GetRequestableOk() (*bool, bool) {
	if o == nil || IsNil(o.Requestable) {
		return nil, false
	}
	return o.Requestable, true
}

// HasRequestable returns a boolean if a field has been set.
func (o *EntitlementDocument) HasRequestable() bool {
	if o != nil && !IsNil(o.Requestable) {
		return true
	}

	return false
}

// SetRequestable gets a reference to the given bool and assigns it to the Requestable field.
func (o *EntitlementDocument) SetRequestable(v bool) {
	o.Requestable = &v
}

// GetCloudGoverned returns the CloudGoverned field value if set, zero value otherwise.
func (o *EntitlementDocument) GetCloudGoverned() bool {
	if o == nil || IsNil(o.CloudGoverned) {
		var ret bool
		return ret
	}
	return *o.CloudGoverned
}

// GetCloudGovernedOk returns a tuple with the CloudGoverned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDocument) GetCloudGovernedOk() (*bool, bool) {
	if o == nil || IsNil(o.CloudGoverned) {
		return nil, false
	}
	return o.CloudGoverned, true
}

// HasCloudGoverned returns a boolean if a field has been set.
func (o *EntitlementDocument) HasCloudGoverned() bool {
	if o != nil && !IsNil(o.CloudGoverned) {
		return true
	}

	return false
}

// SetCloudGoverned gets a reference to the given bool and assigns it to the CloudGoverned field.
func (o *EntitlementDocument) SetCloudGoverned(v bool) {
	o.CloudGoverned = &v
}

// GetCreated returns the Created field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EntitlementDocument) GetCreated() SailPointTime {
	if o == nil || IsNil(o.Created.Get()) {
		var ret SailPointTime
		return ret
	}
	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EntitlementDocument) GetCreatedOk() (*SailPointTime, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// HasCreated returns a boolean if a field has been set.
func (o *EntitlementDocument) HasCreated() bool {
	if o != nil && o.Created.IsSet() {
		return true
	}

	return false
}

// SetCreated gets a reference to the given NullableTime and assigns it to the Created field.
func (o *EntitlementDocument) SetCreated(v SailPointTime) {
	o.Created.Set(&v)
}
// SetCreatedNil sets the value for Created to be an explicit nil
func (o *EntitlementDocument) SetCreatedNil() {
	o.Created.Set(nil)
}

// UnsetCreated ensures that no value is present for Created, not even an explicit nil
func (o *EntitlementDocument) UnsetCreated() {
	o.Created.Unset()
}

// GetPrivileged returns the Privileged field value if set, zero value otherwise.
func (o *EntitlementDocument) GetPrivileged() bool {
	if o == nil || IsNil(o.Privileged) {
		var ret bool
		return ret
	}
	return *o.Privileged
}

// GetPrivilegedOk returns a tuple with the Privileged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDocument) GetPrivilegedOk() (*bool, bool) {
	if o == nil || IsNil(o.Privileged) {
		return nil, false
	}
	return o.Privileged, true
}

// HasPrivileged returns a boolean if a field has been set.
func (o *EntitlementDocument) HasPrivileged() bool {
	if o != nil && !IsNil(o.Privileged) {
		return true
	}

	return false
}

// SetPrivileged gets a reference to the given bool and assigns it to the Privileged field.
func (o *EntitlementDocument) SetPrivileged(v bool) {
	o.Privileged = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *EntitlementDocument) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDocument) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *EntitlementDocument) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *EntitlementDocument) SetTags(v []string) {
	o.Tags = v
}

// GetAttribute returns the Attribute field value if set, zero value otherwise.
func (o *EntitlementDocument) GetAttribute() string {
	if o == nil || IsNil(o.Attribute) {
		var ret string
		return ret
	}
	return *o.Attribute
}

// GetAttributeOk returns a tuple with the Attribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDocument) GetAttributeOk() (*string, bool) {
	if o == nil || IsNil(o.Attribute) {
		return nil, false
	}
	return o.Attribute, true
}

// HasAttribute returns a boolean if a field has been set.
func (o *EntitlementDocument) HasAttribute() bool {
	if o != nil && !IsNil(o.Attribute) {
		return true
	}

	return false
}

// SetAttribute gets a reference to the given string and assigns it to the Attribute field.
func (o *EntitlementDocument) SetAttribute(v string) {
	o.Attribute = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *EntitlementDocument) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDocument) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *EntitlementDocument) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *EntitlementDocument) SetValue(v string) {
	o.Value = &v
}

// GetSourceSchemaObjectType returns the SourceSchemaObjectType field value if set, zero value otherwise.
func (o *EntitlementDocument) GetSourceSchemaObjectType() string {
	if o == nil || IsNil(o.SourceSchemaObjectType) {
		var ret string
		return ret
	}
	return *o.SourceSchemaObjectType
}

// GetSourceSchemaObjectTypeOk returns a tuple with the SourceSchemaObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDocument) GetSourceSchemaObjectTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SourceSchemaObjectType) {
		return nil, false
	}
	return o.SourceSchemaObjectType, true
}

// HasSourceSchemaObjectType returns a boolean if a field has been set.
func (o *EntitlementDocument) HasSourceSchemaObjectType() bool {
	if o != nil && !IsNil(o.SourceSchemaObjectType) {
		return true
	}

	return false
}

// SetSourceSchemaObjectType gets a reference to the given string and assigns it to the SourceSchemaObjectType field.
func (o *EntitlementDocument) SetSourceSchemaObjectType(v string) {
	o.SourceSchemaObjectType = &v
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *EntitlementDocument) GetSchema() string {
	if o == nil || IsNil(o.Schema) {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDocument) GetSchemaOk() (*string, bool) {
	if o == nil || IsNil(o.Schema) {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *EntitlementDocument) HasSchema() bool {
	if o != nil && !IsNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *EntitlementDocument) SetSchema(v string) {
	o.Schema = &v
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *EntitlementDocument) GetHash() string {
	if o == nil || IsNil(o.Hash) {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDocument) GetHashOk() (*string, bool) {
	if o == nil || IsNil(o.Hash) {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *EntitlementDocument) HasHash() bool {
	if o != nil && !IsNil(o.Hash) {
		return true
	}

	return false
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *EntitlementDocument) SetHash(v string) {
	o.Hash = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *EntitlementDocument) GetAttributes() map[string]interface{} {
	if o == nil || IsNil(o.Attributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDocument) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Attributes) {
		return map[string]interface{}{}, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *EntitlementDocument) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *EntitlementDocument) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetTruncatedAttributes returns the TruncatedAttributes field value if set, zero value otherwise.
func (o *EntitlementDocument) GetTruncatedAttributes() []string {
	if o == nil || IsNil(o.TruncatedAttributes) {
		var ret []string
		return ret
	}
	return o.TruncatedAttributes
}

// GetTruncatedAttributesOk returns a tuple with the TruncatedAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDocument) GetTruncatedAttributesOk() ([]string, bool) {
	if o == nil || IsNil(o.TruncatedAttributes) {
		return nil, false
	}
	return o.TruncatedAttributes, true
}

// HasTruncatedAttributes returns a boolean if a field has been set.
func (o *EntitlementDocument) HasTruncatedAttributes() bool {
	if o != nil && !IsNil(o.TruncatedAttributes) {
		return true
	}

	return false
}

// SetTruncatedAttributes gets a reference to the given []string and assigns it to the TruncatedAttributes field.
func (o *EntitlementDocument) SetTruncatedAttributes(v []string) {
	o.TruncatedAttributes = v
}

// GetContainsDataAccess returns the ContainsDataAccess field value if set, zero value otherwise.
func (o *EntitlementDocument) GetContainsDataAccess() bool {
	if o == nil || IsNil(o.ContainsDataAccess) {
		var ret bool
		return ret
	}
	return *o.ContainsDataAccess
}

// GetContainsDataAccessOk returns a tuple with the ContainsDataAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDocument) GetContainsDataAccessOk() (*bool, bool) {
	if o == nil || IsNil(o.ContainsDataAccess) {
		return nil, false
	}
	return o.ContainsDataAccess, true
}

// HasContainsDataAccess returns a boolean if a field has been set.
func (o *EntitlementDocument) HasContainsDataAccess() bool {
	if o != nil && !IsNil(o.ContainsDataAccess) {
		return true
	}

	return false
}

// SetContainsDataAccess gets a reference to the given bool and assigns it to the ContainsDataAccess field.
func (o *EntitlementDocument) SetContainsDataAccess(v bool) {
	o.ContainsDataAccess = &v
}

// GetManuallyUpdatedFields returns the ManuallyUpdatedFields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EntitlementDocument) GetManuallyUpdatedFields() EntitlementDocumentAllOfManuallyUpdatedFields {
	if o == nil || IsNil(o.ManuallyUpdatedFields.Get()) {
		var ret EntitlementDocumentAllOfManuallyUpdatedFields
		return ret
	}
	return *o.ManuallyUpdatedFields.Get()
}

// GetManuallyUpdatedFieldsOk returns a tuple with the ManuallyUpdatedFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EntitlementDocument) GetManuallyUpdatedFieldsOk() (*EntitlementDocumentAllOfManuallyUpdatedFields, bool) {
	if o == nil {
		return nil, false
	}
	return o.ManuallyUpdatedFields.Get(), o.ManuallyUpdatedFields.IsSet()
}

// HasManuallyUpdatedFields returns a boolean if a field has been set.
func (o *EntitlementDocument) HasManuallyUpdatedFields() bool {
	if o != nil && o.ManuallyUpdatedFields.IsSet() {
		return true
	}

	return false
}

// SetManuallyUpdatedFields gets a reference to the given NullableEntitlementDocumentAllOfManuallyUpdatedFields and assigns it to the ManuallyUpdatedFields field.
func (o *EntitlementDocument) SetManuallyUpdatedFields(v EntitlementDocumentAllOfManuallyUpdatedFields) {
	o.ManuallyUpdatedFields.Set(&v)
}
// SetManuallyUpdatedFieldsNil sets the value for ManuallyUpdatedFields to be an explicit nil
func (o *EntitlementDocument) SetManuallyUpdatedFieldsNil() {
	o.ManuallyUpdatedFields.Set(nil)
}

// UnsetManuallyUpdatedFields ensures that no value is present for ManuallyUpdatedFields, not even an explicit nil
func (o *EntitlementDocument) UnsetManuallyUpdatedFields() {
	o.ManuallyUpdatedFields.Unset()
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *EntitlementDocument) GetPermissions() []EntitlementDocumentAllOfPermissions {
	if o == nil || IsNil(o.Permissions) {
		var ret []EntitlementDocumentAllOfPermissions
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDocument) GetPermissionsOk() ([]EntitlementDocumentAllOfPermissions, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *EntitlementDocument) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []EntitlementDocumentAllOfPermissions and assigns it to the Permissions field.
func (o *EntitlementDocument) SetPermissions(v []EntitlementDocumentAllOfPermissions) {
	o.Permissions = v
}

func (o EntitlementDocument) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitlementDocument) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if o.Modified.IsSet() {
		toSerialize["modified"] = o.Modified.Get()
	}
	if !IsNil(o.Synced) {
		toSerialize["synced"] = o.Synced
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.Segments) {
		toSerialize["segments"] = o.Segments
	}
	if !IsNil(o.SegmentCount) {
		toSerialize["segmentCount"] = o.SegmentCount
	}
	if !IsNil(o.Requestable) {
		toSerialize["requestable"] = o.Requestable
	}
	if !IsNil(o.CloudGoverned) {
		toSerialize["cloudGoverned"] = o.CloudGoverned
	}
	if o.Created.IsSet() {
		toSerialize["created"] = o.Created.Get()
	}
	if !IsNil(o.Privileged) {
		toSerialize["privileged"] = o.Privileged
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Attribute) {
		toSerialize["attribute"] = o.Attribute
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.SourceSchemaObjectType) {
		toSerialize["sourceSchemaObjectType"] = o.SourceSchemaObjectType
	}
	if !IsNil(o.Schema) {
		toSerialize["schema"] = o.Schema
	}
	if !IsNil(o.Hash) {
		toSerialize["hash"] = o.Hash
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.TruncatedAttributes) {
		toSerialize["truncatedAttributes"] = o.TruncatedAttributes
	}
	if !IsNil(o.ContainsDataAccess) {
		toSerialize["containsDataAccess"] = o.ContainsDataAccess
	}
	if o.ManuallyUpdatedFields.IsSet() {
		toSerialize["manuallyUpdatedFields"] = o.ManuallyUpdatedFields.Get()
	}
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EntitlementDocument) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
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

	varEntitlementDocument := _EntitlementDocument{}

	err = json.Unmarshal(data, &varEntitlementDocument)

	if err != nil {
		return err
	}

	*o = EntitlementDocument(varEntitlementDocument)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "modified")
		delete(additionalProperties, "synced")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "source")
		delete(additionalProperties, "segments")
		delete(additionalProperties, "segmentCount")
		delete(additionalProperties, "requestable")
		delete(additionalProperties, "cloudGoverned")
		delete(additionalProperties, "created")
		delete(additionalProperties, "privileged")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "attribute")
		delete(additionalProperties, "value")
		delete(additionalProperties, "sourceSchemaObjectType")
		delete(additionalProperties, "schema")
		delete(additionalProperties, "hash")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "truncatedAttributes")
		delete(additionalProperties, "containsDataAccess")
		delete(additionalProperties, "manuallyUpdatedFields")
		delete(additionalProperties, "permissions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntitlementDocument struct {
	value *EntitlementDocument
	isSet bool
}

func (v NullableEntitlementDocument) Get() *EntitlementDocument {
	return v.value
}

func (v *NullableEntitlementDocument) Set(val *EntitlementDocument) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementDocument) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementDocument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementDocument(val *EntitlementDocument) *NullableEntitlementDocument {
	return &NullableEntitlementDocument{value: val, isSet: true}
}

func (v NullableEntitlementDocument) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementDocument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


