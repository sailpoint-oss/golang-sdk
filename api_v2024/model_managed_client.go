/*
Identity Security Cloud V2024 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: v2024
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2024

import (
	"encoding/json"
	
	"fmt"
)

// checks if the ManagedClient type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagedClient{}

// ManagedClient Managed Client
type ManagedClient struct {
	// ManagedClient ID
	Id NullableString `json:"id,omitempty"`
	// ManagedClient alert key
	AlertKey NullableString `json:"alertKey,omitempty"`
	// apiGatewayBaseUrl for the Managed client
	ApiGatewayBaseUrl NullableString `json:"apiGatewayBaseUrl,omitempty"`
	// cookbook id for the Managed client
	Cookbook NullableString `json:"cookbook,omitempty"`
	// Previous CC ID to be used in data migration. (This field will be deleted after CC migration!)
	CcId NullableInt64 `json:"ccId,omitempty"`
	// The client ID used in API management
	ClientId string `json:"clientId"`
	// Cluster ID that the ManagedClient is linked to
	ClusterId string `json:"clusterId"`
	// ManagedClient description
	Description string `json:"description"`
	// The public IP address of the ManagedClient
	IpAddress NullableString `json:"ipAddress,omitempty"`
	// When the ManagedClient was last seen by the server
	LastSeen NullableTime `json:"lastSeen,omitempty"`
	// ManagedClient name
	Name NullableString `json:"name,omitempty"`
	// Milliseconds since the ManagedClient has polled the server
	SinceLastSeen NullableString `json:"sinceLastSeen,omitempty"`
	// Status of the ManagedClient
	Status NullableString `json:"status,omitempty"`
	// Type of the ManagedClient (VA, CCG)
	Type string `json:"type"`
	// Cluster Type of the ManagedClient
	ClusterType NullableString `json:"clusterType,omitempty"`
	// ManagedClient VA download URL
	VaDownloadUrl NullableString `json:"vaDownloadUrl,omitempty"`
	// Version that the ManagedClient's VA is running
	VaVersion NullableString `json:"vaVersion,omitempty"`
	// Client's apiKey
	Secret NullableString `json:"secret,omitempty"`
	// The date/time this ManagedClient was created
	CreatedAt NullableTime `json:"createdAt,omitempty"`
	// The date/time this ManagedClient was last updated
	UpdatedAt NullableTime `json:"updatedAt,omitempty"`
	// The provisioning status of the ManagedClient
	ProvisionStatus NullableString `json:"provisionStatus,omitempty"`
	// The health indicators of the ManagedClient
	HealthIndicators map[string]interface{} `json:"healthIndicators,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ManagedClient ManagedClient

// NewManagedClient instantiates a new ManagedClient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagedClient(clientId string, clusterId string, description string, type_ string) *ManagedClient {
	this := ManagedClient{}
	this.ClientId = clientId
	this.ClusterId = clusterId
	this.Description = description
	var name string = "VA-$clientId"
	this.Name = *NewNullableString(&name)
	this.Type = type_
	return &this
}

// NewManagedClientWithDefaults instantiates a new ManagedClient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagedClientWithDefaults() *ManagedClient {
	this := ManagedClient{}
	var description string = ""
	this.Description = description
	var name string = "VA-$clientId"
	this.Name = *NewNullableString(&name)
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManagedClient) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManagedClient) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *ManagedClient) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *ManagedClient) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *ManagedClient) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *ManagedClient) UnsetId() {
	o.Id.Unset()
}

// GetAlertKey returns the AlertKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManagedClient) GetAlertKey() string {
	if o == nil || IsNil(o.AlertKey.Get()) {
		var ret string
		return ret
	}
	return *o.AlertKey.Get()
}

// GetAlertKeyOk returns a tuple with the AlertKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManagedClient) GetAlertKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AlertKey.Get(), o.AlertKey.IsSet()
}

// HasAlertKey returns a boolean if a field has been set.
func (o *ManagedClient) HasAlertKey() bool {
	if o != nil && o.AlertKey.IsSet() {
		return true
	}

	return false
}

// SetAlertKey gets a reference to the given NullableString and assigns it to the AlertKey field.
func (o *ManagedClient) SetAlertKey(v string) {
	o.AlertKey.Set(&v)
}
// SetAlertKeyNil sets the value for AlertKey to be an explicit nil
func (o *ManagedClient) SetAlertKeyNil() {
	o.AlertKey.Set(nil)
}

// UnsetAlertKey ensures that no value is present for AlertKey, not even an explicit nil
func (o *ManagedClient) UnsetAlertKey() {
	o.AlertKey.Unset()
}

// GetApiGatewayBaseUrl returns the ApiGatewayBaseUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManagedClient) GetApiGatewayBaseUrl() string {
	if o == nil || IsNil(o.ApiGatewayBaseUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ApiGatewayBaseUrl.Get()
}

// GetApiGatewayBaseUrlOk returns a tuple with the ApiGatewayBaseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManagedClient) GetApiGatewayBaseUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApiGatewayBaseUrl.Get(), o.ApiGatewayBaseUrl.IsSet()
}

// HasApiGatewayBaseUrl returns a boolean if a field has been set.
func (o *ManagedClient) HasApiGatewayBaseUrl() bool {
	if o != nil && o.ApiGatewayBaseUrl.IsSet() {
		return true
	}

	return false
}

// SetApiGatewayBaseUrl gets a reference to the given NullableString and assigns it to the ApiGatewayBaseUrl field.
func (o *ManagedClient) SetApiGatewayBaseUrl(v string) {
	o.ApiGatewayBaseUrl.Set(&v)
}
// SetApiGatewayBaseUrlNil sets the value for ApiGatewayBaseUrl to be an explicit nil
func (o *ManagedClient) SetApiGatewayBaseUrlNil() {
	o.ApiGatewayBaseUrl.Set(nil)
}

// UnsetApiGatewayBaseUrl ensures that no value is present for ApiGatewayBaseUrl, not even an explicit nil
func (o *ManagedClient) UnsetApiGatewayBaseUrl() {
	o.ApiGatewayBaseUrl.Unset()
}

// GetCookbook returns the Cookbook field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManagedClient) GetCookbook() string {
	if o == nil || IsNil(o.Cookbook.Get()) {
		var ret string
		return ret
	}
	return *o.Cookbook.Get()
}

// GetCookbookOk returns a tuple with the Cookbook field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManagedClient) GetCookbookOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cookbook.Get(), o.Cookbook.IsSet()
}

// HasCookbook returns a boolean if a field has been set.
func (o *ManagedClient) HasCookbook() bool {
	if o != nil && o.Cookbook.IsSet() {
		return true
	}

	return false
}

// SetCookbook gets a reference to the given NullableString and assigns it to the Cookbook field.
func (o *ManagedClient) SetCookbook(v string) {
	o.Cookbook.Set(&v)
}
// SetCookbookNil sets the value for Cookbook to be an explicit nil
func (o *ManagedClient) SetCookbookNil() {
	o.Cookbook.Set(nil)
}

// UnsetCookbook ensures that no value is present for Cookbook, not even an explicit nil
func (o *ManagedClient) UnsetCookbook() {
	o.Cookbook.Unset()
}

// GetCcId returns the CcId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManagedClient) GetCcId() int64 {
	if o == nil || IsNil(o.CcId.Get()) {
		var ret int64
		return ret
	}
	return *o.CcId.Get()
}

// GetCcIdOk returns a tuple with the CcId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManagedClient) GetCcIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CcId.Get(), o.CcId.IsSet()
}

// HasCcId returns a boolean if a field has been set.
func (o *ManagedClient) HasCcId() bool {
	if o != nil && o.CcId.IsSet() {
		return true
	}

	return false
}

// SetCcId gets a reference to the given NullableInt64 and assigns it to the CcId field.
func (o *ManagedClient) SetCcId(v int64) {
	o.CcId.Set(&v)
}
// SetCcIdNil sets the value for CcId to be an explicit nil
func (o *ManagedClient) SetCcIdNil() {
	o.CcId.Set(nil)
}

// UnsetCcId ensures that no value is present for CcId, not even an explicit nil
func (o *ManagedClient) UnsetCcId() {
	o.CcId.Unset()
}

// GetClientId returns the ClientId field value
func (o *ManagedClient) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *ManagedClient) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *ManagedClient) SetClientId(v string) {
	o.ClientId = v
}

// GetClusterId returns the ClusterId field value
func (o *ManagedClient) GetClusterId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value
// and a boolean to check if the value has been set.
func (o *ManagedClient) GetClusterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterId, true
}

// SetClusterId sets field value
func (o *ManagedClient) SetClusterId(v string) {
	o.ClusterId = v
}

// GetDescription returns the Description field value
func (o *ManagedClient) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ManagedClient) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ManagedClient) SetDescription(v string) {
	o.Description = v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManagedClient) GetIpAddress() string {
	if o == nil || IsNil(o.IpAddress.Get()) {
		var ret string
		return ret
	}
	return *o.IpAddress.Get()
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManagedClient) GetIpAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IpAddress.Get(), o.IpAddress.IsSet()
}

// HasIpAddress returns a boolean if a field has been set.
func (o *ManagedClient) HasIpAddress() bool {
	if o != nil && o.IpAddress.IsSet() {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given NullableString and assigns it to the IpAddress field.
func (o *ManagedClient) SetIpAddress(v string) {
	o.IpAddress.Set(&v)
}
// SetIpAddressNil sets the value for IpAddress to be an explicit nil
func (o *ManagedClient) SetIpAddressNil() {
	o.IpAddress.Set(nil)
}

// UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
func (o *ManagedClient) UnsetIpAddress() {
	o.IpAddress.Unset()
}

// GetLastSeen returns the LastSeen field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManagedClient) GetLastSeen() SailPointTime {
	if o == nil || IsNil(o.LastSeen.Get()) {
		var ret SailPointTime
		return ret
	}
	return *o.LastSeen.Get()
}

// GetLastSeenOk returns a tuple with the LastSeen field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManagedClient) GetLastSeenOk() (*SailPointTime, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastSeen.Get(), o.LastSeen.IsSet()
}

// HasLastSeen returns a boolean if a field has been set.
func (o *ManagedClient) HasLastSeen() bool {
	if o != nil && o.LastSeen.IsSet() {
		return true
	}

	return false
}

// SetLastSeen gets a reference to the given NullableTime and assigns it to the LastSeen field.
func (o *ManagedClient) SetLastSeen(v SailPointTime) {
	o.LastSeen.Set(&v)
}
// SetLastSeenNil sets the value for LastSeen to be an explicit nil
func (o *ManagedClient) SetLastSeenNil() {
	o.LastSeen.Set(nil)
}

// UnsetLastSeen ensures that no value is present for LastSeen, not even an explicit nil
func (o *ManagedClient) UnsetLastSeen() {
	o.LastSeen.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManagedClient) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManagedClient) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ManagedClient) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ManagedClient) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ManagedClient) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ManagedClient) UnsetName() {
	o.Name.Unset()
}

// GetSinceLastSeen returns the SinceLastSeen field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManagedClient) GetSinceLastSeen() string {
	if o == nil || IsNil(o.SinceLastSeen.Get()) {
		var ret string
		return ret
	}
	return *o.SinceLastSeen.Get()
}

// GetSinceLastSeenOk returns a tuple with the SinceLastSeen field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManagedClient) GetSinceLastSeenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SinceLastSeen.Get(), o.SinceLastSeen.IsSet()
}

// HasSinceLastSeen returns a boolean if a field has been set.
func (o *ManagedClient) HasSinceLastSeen() bool {
	if o != nil && o.SinceLastSeen.IsSet() {
		return true
	}

	return false
}

// SetSinceLastSeen gets a reference to the given NullableString and assigns it to the SinceLastSeen field.
func (o *ManagedClient) SetSinceLastSeen(v string) {
	o.SinceLastSeen.Set(&v)
}
// SetSinceLastSeenNil sets the value for SinceLastSeen to be an explicit nil
func (o *ManagedClient) SetSinceLastSeenNil() {
	o.SinceLastSeen.Set(nil)
}

// UnsetSinceLastSeen ensures that no value is present for SinceLastSeen, not even an explicit nil
func (o *ManagedClient) UnsetSinceLastSeen() {
	o.SinceLastSeen.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManagedClient) GetStatus() string {
	if o == nil || IsNil(o.Status.Get()) {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManagedClient) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *ManagedClient) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *ManagedClient) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *ManagedClient) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *ManagedClient) UnsetStatus() {
	o.Status.Unset()
}

// GetType returns the Type field value
func (o *ManagedClient) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ManagedClient) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ManagedClient) SetType(v string) {
	o.Type = v
}

// GetClusterType returns the ClusterType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManagedClient) GetClusterType() string {
	if o == nil || IsNil(o.ClusterType.Get()) {
		var ret string
		return ret
	}
	return *o.ClusterType.Get()
}

// GetClusterTypeOk returns a tuple with the ClusterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManagedClient) GetClusterTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClusterType.Get(), o.ClusterType.IsSet()
}

// HasClusterType returns a boolean if a field has been set.
func (o *ManagedClient) HasClusterType() bool {
	if o != nil && o.ClusterType.IsSet() {
		return true
	}

	return false
}

// SetClusterType gets a reference to the given NullableString and assigns it to the ClusterType field.
func (o *ManagedClient) SetClusterType(v string) {
	o.ClusterType.Set(&v)
}
// SetClusterTypeNil sets the value for ClusterType to be an explicit nil
func (o *ManagedClient) SetClusterTypeNil() {
	o.ClusterType.Set(nil)
}

// UnsetClusterType ensures that no value is present for ClusterType, not even an explicit nil
func (o *ManagedClient) UnsetClusterType() {
	o.ClusterType.Unset()
}

// GetVaDownloadUrl returns the VaDownloadUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManagedClient) GetVaDownloadUrl() string {
	if o == nil || IsNil(o.VaDownloadUrl.Get()) {
		var ret string
		return ret
	}
	return *o.VaDownloadUrl.Get()
}

// GetVaDownloadUrlOk returns a tuple with the VaDownloadUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManagedClient) GetVaDownloadUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VaDownloadUrl.Get(), o.VaDownloadUrl.IsSet()
}

// HasVaDownloadUrl returns a boolean if a field has been set.
func (o *ManagedClient) HasVaDownloadUrl() bool {
	if o != nil && o.VaDownloadUrl.IsSet() {
		return true
	}

	return false
}

// SetVaDownloadUrl gets a reference to the given NullableString and assigns it to the VaDownloadUrl field.
func (o *ManagedClient) SetVaDownloadUrl(v string) {
	o.VaDownloadUrl.Set(&v)
}
// SetVaDownloadUrlNil sets the value for VaDownloadUrl to be an explicit nil
func (o *ManagedClient) SetVaDownloadUrlNil() {
	o.VaDownloadUrl.Set(nil)
}

// UnsetVaDownloadUrl ensures that no value is present for VaDownloadUrl, not even an explicit nil
func (o *ManagedClient) UnsetVaDownloadUrl() {
	o.VaDownloadUrl.Unset()
}

// GetVaVersion returns the VaVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManagedClient) GetVaVersion() string {
	if o == nil || IsNil(o.VaVersion.Get()) {
		var ret string
		return ret
	}
	return *o.VaVersion.Get()
}

// GetVaVersionOk returns a tuple with the VaVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManagedClient) GetVaVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VaVersion.Get(), o.VaVersion.IsSet()
}

// HasVaVersion returns a boolean if a field has been set.
func (o *ManagedClient) HasVaVersion() bool {
	if o != nil && o.VaVersion.IsSet() {
		return true
	}

	return false
}

// SetVaVersion gets a reference to the given NullableString and assigns it to the VaVersion field.
func (o *ManagedClient) SetVaVersion(v string) {
	o.VaVersion.Set(&v)
}
// SetVaVersionNil sets the value for VaVersion to be an explicit nil
func (o *ManagedClient) SetVaVersionNil() {
	o.VaVersion.Set(nil)
}

// UnsetVaVersion ensures that no value is present for VaVersion, not even an explicit nil
func (o *ManagedClient) UnsetVaVersion() {
	o.VaVersion.Unset()
}

// GetSecret returns the Secret field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManagedClient) GetSecret() string {
	if o == nil || IsNil(o.Secret.Get()) {
		var ret string
		return ret
	}
	return *o.Secret.Get()
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManagedClient) GetSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Secret.Get(), o.Secret.IsSet()
}

// HasSecret returns a boolean if a field has been set.
func (o *ManagedClient) HasSecret() bool {
	if o != nil && o.Secret.IsSet() {
		return true
	}

	return false
}

// SetSecret gets a reference to the given NullableString and assigns it to the Secret field.
func (o *ManagedClient) SetSecret(v string) {
	o.Secret.Set(&v)
}
// SetSecretNil sets the value for Secret to be an explicit nil
func (o *ManagedClient) SetSecretNil() {
	o.Secret.Set(nil)
}

// UnsetSecret ensures that no value is present for Secret, not even an explicit nil
func (o *ManagedClient) UnsetSecret() {
	o.Secret.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManagedClient) GetCreatedAt() SailPointTime {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret SailPointTime
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManagedClient) GetCreatedAtOk() (*SailPointTime, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ManagedClient) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableTime and assigns it to the CreatedAt field.
func (o *ManagedClient) SetCreatedAt(v SailPointTime) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *ManagedClient) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *ManagedClient) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManagedClient) GetUpdatedAt() SailPointTime {
	if o == nil || IsNil(o.UpdatedAt.Get()) {
		var ret SailPointTime
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManagedClient) GetUpdatedAtOk() (*SailPointTime, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ManagedClient) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableTime and assigns it to the UpdatedAt field.
func (o *ManagedClient) SetUpdatedAt(v SailPointTime) {
	o.UpdatedAt.Set(&v)
}
// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *ManagedClient) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *ManagedClient) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetProvisionStatus returns the ProvisionStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManagedClient) GetProvisionStatus() string {
	if o == nil || IsNil(o.ProvisionStatus.Get()) {
		var ret string
		return ret
	}
	return *o.ProvisionStatus.Get()
}

// GetProvisionStatusOk returns a tuple with the ProvisionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManagedClient) GetProvisionStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProvisionStatus.Get(), o.ProvisionStatus.IsSet()
}

// HasProvisionStatus returns a boolean if a field has been set.
func (o *ManagedClient) HasProvisionStatus() bool {
	if o != nil && o.ProvisionStatus.IsSet() {
		return true
	}

	return false
}

// SetProvisionStatus gets a reference to the given NullableString and assigns it to the ProvisionStatus field.
func (o *ManagedClient) SetProvisionStatus(v string) {
	o.ProvisionStatus.Set(&v)
}
// SetProvisionStatusNil sets the value for ProvisionStatus to be an explicit nil
func (o *ManagedClient) SetProvisionStatusNil() {
	o.ProvisionStatus.Set(nil)
}

// UnsetProvisionStatus ensures that no value is present for ProvisionStatus, not even an explicit nil
func (o *ManagedClient) UnsetProvisionStatus() {
	o.ProvisionStatus.Unset()
}

// GetHealthIndicators returns the HealthIndicators field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManagedClient) GetHealthIndicators() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.HealthIndicators
}

// GetHealthIndicatorsOk returns a tuple with the HealthIndicators field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManagedClient) GetHealthIndicatorsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.HealthIndicators) {
		return map[string]interface{}{}, false
	}
	return o.HealthIndicators, true
}

// HasHealthIndicators returns a boolean if a field has been set.
func (o *ManagedClient) HasHealthIndicators() bool {
	if o != nil && !IsNil(o.HealthIndicators) {
		return true
	}

	return false
}

// SetHealthIndicators gets a reference to the given map[string]interface{} and assigns it to the HealthIndicators field.
func (o *ManagedClient) SetHealthIndicators(v map[string]interface{}) {
	o.HealthIndicators = v
}

func (o ManagedClient) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagedClient) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.AlertKey.IsSet() {
		toSerialize["alertKey"] = o.AlertKey.Get()
	}
	if o.ApiGatewayBaseUrl.IsSet() {
		toSerialize["apiGatewayBaseUrl"] = o.ApiGatewayBaseUrl.Get()
	}
	if o.Cookbook.IsSet() {
		toSerialize["cookbook"] = o.Cookbook.Get()
	}
	if o.CcId.IsSet() {
		toSerialize["ccId"] = o.CcId.Get()
	}
	toSerialize["clientId"] = o.ClientId
	toSerialize["clusterId"] = o.ClusterId
	toSerialize["description"] = o.Description
	if o.IpAddress.IsSet() {
		toSerialize["ipAddress"] = o.IpAddress.Get()
	}
	if o.LastSeen.IsSet() {
		toSerialize["lastSeen"] = o.LastSeen.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.SinceLastSeen.IsSet() {
		toSerialize["sinceLastSeen"] = o.SinceLastSeen.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	toSerialize["type"] = o.Type
	if o.ClusterType.IsSet() {
		toSerialize["clusterType"] = o.ClusterType.Get()
	}
	if o.VaDownloadUrl.IsSet() {
		toSerialize["vaDownloadUrl"] = o.VaDownloadUrl.Get()
	}
	if o.VaVersion.IsSet() {
		toSerialize["vaVersion"] = o.VaVersion.Get()
	}
	if o.Secret.IsSet() {
		toSerialize["secret"] = o.Secret.Get()
	}
	if o.CreatedAt.IsSet() {
		toSerialize["createdAt"] = o.CreatedAt.Get()
	}
	if o.UpdatedAt.IsSet() {
		toSerialize["updatedAt"] = o.UpdatedAt.Get()
	}
	if o.ProvisionStatus.IsSet() {
		toSerialize["provisionStatus"] = o.ProvisionStatus.Get()
	}
	if o.HealthIndicators != nil {
		toSerialize["healthIndicators"] = o.HealthIndicators
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ManagedClient) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"clientId",
		"clusterId",
		"description",
		"type",
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

	varManagedClient := _ManagedClient{}

	err = json.Unmarshal(data, &varManagedClient)

	if err != nil {
		return err
	}

	*o = ManagedClient(varManagedClient)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "alertKey")
		delete(additionalProperties, "apiGatewayBaseUrl")
		delete(additionalProperties, "cookbook")
		delete(additionalProperties, "ccId")
		delete(additionalProperties, "clientId")
		delete(additionalProperties, "clusterId")
		delete(additionalProperties, "description")
		delete(additionalProperties, "ipAddress")
		delete(additionalProperties, "lastSeen")
		delete(additionalProperties, "name")
		delete(additionalProperties, "sinceLastSeen")
		delete(additionalProperties, "status")
		delete(additionalProperties, "type")
		delete(additionalProperties, "clusterType")
		delete(additionalProperties, "vaDownloadUrl")
		delete(additionalProperties, "vaVersion")
		delete(additionalProperties, "secret")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "updatedAt")
		delete(additionalProperties, "provisionStatus")
		delete(additionalProperties, "healthIndicators")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableManagedClient struct {
	value *ManagedClient
	isSet bool
}

func (v NullableManagedClient) Get() *ManagedClient {
	return v.value
}

func (v *NullableManagedClient) Set(val *ManagedClient) {
	v.value = val
	v.isSet = true
}

func (v NullableManagedClient) IsSet() bool {
	return v.isSet
}

func (v *NullableManagedClient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagedClient(val *ManagedClient) *NullableManagedClient {
	return &NullableManagedClient{value: val, isSet: true}
}

func (v NullableManagedClient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagedClient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


