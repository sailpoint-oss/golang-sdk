# GetApplication200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**AppId** | Pointer to **string** |  | [optional] 
**ServiceId** | Pointer to **string** |  | [optional] 
**ServiceAppId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**AppCenterEnabled** | Pointer to **bool** |  | [optional] 
**ProvisionRequestEnabled** | Pointer to **bool** |  | [optional] 
**ControlType** | Pointer to **string** |  | [optional] 
**Mobile** | Pointer to **bool** |  | [optional] 
**PrivateApp** | Pointer to **bool** |  | [optional] 
**ScriptName** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Icon** | Pointer to **string** |  | [optional] 
**Health** | Pointer to [**ListApplications200ResponseInnerHealth**](ListApplications200ResponseInnerHealth.md) |  | [optional] 
**EnableSso** | Pointer to **bool** |  | [optional] 
**SsoMethod** | Pointer to **string** |  | [optional] 
**HasLinks** | Pointer to **bool** |  | [optional] 
**HasAutomations** | Pointer to **bool** |  | [optional] 
**StepUpAuthData** | Pointer to **map[string]interface{}** |  | [optional] 
**StepUpAuthType** | Pointer to **string** |  | [optional] 
**UsageAnalytics** | Pointer to **bool** |  | [optional] 
**UsageCertRequired** | Pointer to **bool** |  | [optional] 
**UsageCertText** | Pointer to **map[string]interface{}** |  | [optional] 
**LaunchpadEnabled** | Pointer to **bool** |  | [optional] 
**PasswordManaged** | Pointer to **bool** |  | [optional] 
**Owner** | Pointer to [**ListApplications200ResponseInnerOwner**](ListApplications200ResponseInnerOwner.md) |  | [optional] 
**DateCreated** | Pointer to **float32** |  | [optional] 
**LastUpdated** | Pointer to **float32** |  | [optional] 
**DefaultAccessProfile** | Pointer to **map[string]interface{}** |  | [optional] 
**Service** | Pointer to **string** |  | [optional] 
**SelectedSsoMethod** | Pointer to **string** |  | [optional] 
**SupportedSsoMethods** | Pointer to **float32** |  | [optional] 
**OffNetworkBlockedRoles** | Pointer to **map[string]interface{}** |  | [optional] 
**SupportedOffNetwork** | Pointer to **string** |  | [optional] 
**AccountServiceId** | Pointer to **float32** |  | [optional] 
**LauncherCount** | Pointer to **float32** |  | [optional] 
**AccountServiceName** | Pointer to **string** |  | [optional] 
**AccountServiceExternalId** | Pointer to **string** |  | [optional] 
**AccountServiceMatchAllAccounts** | Pointer to **bool** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**AccountServiceUseForPasswordManagement** | Pointer to **bool** |  | [optional] 
**AccountServicePolicyId** | Pointer to **string** |  | [optional] 
**AccountServicePolicyName** | Pointer to **string** |  | [optional] 
**RequireStrongAuthn** | Pointer to **bool** |  | [optional] 
**AccountServicePolicies** | Pointer to [**[]ListApplications200ResponseInnerAccountServicePoliciesInner**](ListApplications200ResponseInnerAccountServicePoliciesInner.md) |  | [optional] 
**XsdVersion** | Pointer to **string** |  | [optional] 
**AppProfiles** | Pointer to [**[]ListApplications200ResponseInnerAppProfilesInner**](ListApplications200ResponseInnerAppProfilesInner.md) |  | [optional] 
**PasswordServiceId** | Pointer to **float32** |  | [optional] 
**AccessProfileIds** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewGetApplication200Response

`func NewGetApplication200Response() *GetApplication200Response`

NewGetApplication200Response instantiates a new GetApplication200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetApplication200ResponseWithDefaults

`func NewGetApplication200ResponseWithDefaults() *GetApplication200Response`

NewGetApplication200ResponseWithDefaults instantiates a new GetApplication200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetApplication200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetApplication200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetApplication200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetApplication200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAppId

`func (o *GetApplication200Response) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *GetApplication200Response) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *GetApplication200Response) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *GetApplication200Response) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetServiceId

`func (o *GetApplication200Response) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *GetApplication200Response) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *GetApplication200Response) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *GetApplication200Response) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetServiceAppId

`func (o *GetApplication200Response) GetServiceAppId() string`

GetServiceAppId returns the ServiceAppId field if non-nil, zero value otherwise.

### GetServiceAppIdOk

`func (o *GetApplication200Response) GetServiceAppIdOk() (*string, bool)`

GetServiceAppIdOk returns a tuple with the ServiceAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAppId

`func (o *GetApplication200Response) SetServiceAppId(v string)`

SetServiceAppId sets ServiceAppId field to given value.

### HasServiceAppId

`func (o *GetApplication200Response) HasServiceAppId() bool`

HasServiceAppId returns a boolean if a field has been set.

### GetName

`func (o *GetApplication200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetApplication200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetApplication200Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetApplication200Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetApplication200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetApplication200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetApplication200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetApplication200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAppCenterEnabled

`func (o *GetApplication200Response) GetAppCenterEnabled() bool`

GetAppCenterEnabled returns the AppCenterEnabled field if non-nil, zero value otherwise.

### GetAppCenterEnabledOk

`func (o *GetApplication200Response) GetAppCenterEnabledOk() (*bool, bool)`

GetAppCenterEnabledOk returns a tuple with the AppCenterEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppCenterEnabled

`func (o *GetApplication200Response) SetAppCenterEnabled(v bool)`

SetAppCenterEnabled sets AppCenterEnabled field to given value.

### HasAppCenterEnabled

`func (o *GetApplication200Response) HasAppCenterEnabled() bool`

HasAppCenterEnabled returns a boolean if a field has been set.

### GetProvisionRequestEnabled

`func (o *GetApplication200Response) GetProvisionRequestEnabled() bool`

GetProvisionRequestEnabled returns the ProvisionRequestEnabled field if non-nil, zero value otherwise.

### GetProvisionRequestEnabledOk

`func (o *GetApplication200Response) GetProvisionRequestEnabledOk() (*bool, bool)`

GetProvisionRequestEnabledOk returns a tuple with the ProvisionRequestEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionRequestEnabled

`func (o *GetApplication200Response) SetProvisionRequestEnabled(v bool)`

SetProvisionRequestEnabled sets ProvisionRequestEnabled field to given value.

### HasProvisionRequestEnabled

`func (o *GetApplication200Response) HasProvisionRequestEnabled() bool`

HasProvisionRequestEnabled returns a boolean if a field has been set.

### GetControlType

`func (o *GetApplication200Response) GetControlType() string`

GetControlType returns the ControlType field if non-nil, zero value otherwise.

### GetControlTypeOk

`func (o *GetApplication200Response) GetControlTypeOk() (*string, bool)`

GetControlTypeOk returns a tuple with the ControlType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlType

`func (o *GetApplication200Response) SetControlType(v string)`

SetControlType sets ControlType field to given value.

### HasControlType

`func (o *GetApplication200Response) HasControlType() bool`

HasControlType returns a boolean if a field has been set.

### GetMobile

`func (o *GetApplication200Response) GetMobile() bool`

GetMobile returns the Mobile field if non-nil, zero value otherwise.

### GetMobileOk

`func (o *GetApplication200Response) GetMobileOk() (*bool, bool)`

GetMobileOk returns a tuple with the Mobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobile

`func (o *GetApplication200Response) SetMobile(v bool)`

SetMobile sets Mobile field to given value.

### HasMobile

`func (o *GetApplication200Response) HasMobile() bool`

HasMobile returns a boolean if a field has been set.

### GetPrivateApp

`func (o *GetApplication200Response) GetPrivateApp() bool`

GetPrivateApp returns the PrivateApp field if non-nil, zero value otherwise.

### GetPrivateAppOk

`func (o *GetApplication200Response) GetPrivateAppOk() (*bool, bool)`

GetPrivateAppOk returns a tuple with the PrivateApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateApp

`func (o *GetApplication200Response) SetPrivateApp(v bool)`

SetPrivateApp sets PrivateApp field to given value.

### HasPrivateApp

`func (o *GetApplication200Response) HasPrivateApp() bool`

HasPrivateApp returns a boolean if a field has been set.

### GetScriptName

`func (o *GetApplication200Response) GetScriptName() string`

GetScriptName returns the ScriptName field if non-nil, zero value otherwise.

### GetScriptNameOk

`func (o *GetApplication200Response) GetScriptNameOk() (*string, bool)`

GetScriptNameOk returns a tuple with the ScriptName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptName

`func (o *GetApplication200Response) SetScriptName(v string)`

SetScriptName sets ScriptName field to given value.

### HasScriptName

`func (o *GetApplication200Response) HasScriptName() bool`

HasScriptName returns a boolean if a field has been set.

### GetStatus

`func (o *GetApplication200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetApplication200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetApplication200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetApplication200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetIcon

`func (o *GetApplication200Response) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *GetApplication200Response) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *GetApplication200Response) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *GetApplication200Response) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetHealth

`func (o *GetApplication200Response) GetHealth() ListApplications200ResponseInnerHealth`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *GetApplication200Response) GetHealthOk() (*ListApplications200ResponseInnerHealth, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *GetApplication200Response) SetHealth(v ListApplications200ResponseInnerHealth)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *GetApplication200Response) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetEnableSso

`func (o *GetApplication200Response) GetEnableSso() bool`

GetEnableSso returns the EnableSso field if non-nil, zero value otherwise.

### GetEnableSsoOk

`func (o *GetApplication200Response) GetEnableSsoOk() (*bool, bool)`

GetEnableSsoOk returns a tuple with the EnableSso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSso

`func (o *GetApplication200Response) SetEnableSso(v bool)`

SetEnableSso sets EnableSso field to given value.

### HasEnableSso

`func (o *GetApplication200Response) HasEnableSso() bool`

HasEnableSso returns a boolean if a field has been set.

### GetSsoMethod

`func (o *GetApplication200Response) GetSsoMethod() string`

GetSsoMethod returns the SsoMethod field if non-nil, zero value otherwise.

### GetSsoMethodOk

`func (o *GetApplication200Response) GetSsoMethodOk() (*string, bool)`

GetSsoMethodOk returns a tuple with the SsoMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoMethod

`func (o *GetApplication200Response) SetSsoMethod(v string)`

SetSsoMethod sets SsoMethod field to given value.

### HasSsoMethod

`func (o *GetApplication200Response) HasSsoMethod() bool`

HasSsoMethod returns a boolean if a field has been set.

### GetHasLinks

`func (o *GetApplication200Response) GetHasLinks() bool`

GetHasLinks returns the HasLinks field if non-nil, zero value otherwise.

### GetHasLinksOk

`func (o *GetApplication200Response) GetHasLinksOk() (*bool, bool)`

GetHasLinksOk returns a tuple with the HasLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasLinks

`func (o *GetApplication200Response) SetHasLinks(v bool)`

SetHasLinks sets HasLinks field to given value.

### HasHasLinks

`func (o *GetApplication200Response) HasHasLinks() bool`

HasHasLinks returns a boolean if a field has been set.

### GetHasAutomations

`func (o *GetApplication200Response) GetHasAutomations() bool`

GetHasAutomations returns the HasAutomations field if non-nil, zero value otherwise.

### GetHasAutomationsOk

`func (o *GetApplication200Response) GetHasAutomationsOk() (*bool, bool)`

GetHasAutomationsOk returns a tuple with the HasAutomations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAutomations

`func (o *GetApplication200Response) SetHasAutomations(v bool)`

SetHasAutomations sets HasAutomations field to given value.

### HasHasAutomations

`func (o *GetApplication200Response) HasHasAutomations() bool`

HasHasAutomations returns a boolean if a field has been set.

### GetStepUpAuthData

`func (o *GetApplication200Response) GetStepUpAuthData() map[string]interface{}`

GetStepUpAuthData returns the StepUpAuthData field if non-nil, zero value otherwise.

### GetStepUpAuthDataOk

`func (o *GetApplication200Response) GetStepUpAuthDataOk() (*map[string]interface{}, bool)`

GetStepUpAuthDataOk returns a tuple with the StepUpAuthData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepUpAuthData

`func (o *GetApplication200Response) SetStepUpAuthData(v map[string]interface{})`

SetStepUpAuthData sets StepUpAuthData field to given value.

### HasStepUpAuthData

`func (o *GetApplication200Response) HasStepUpAuthData() bool`

HasStepUpAuthData returns a boolean if a field has been set.

### GetStepUpAuthType

`func (o *GetApplication200Response) GetStepUpAuthType() string`

GetStepUpAuthType returns the StepUpAuthType field if non-nil, zero value otherwise.

### GetStepUpAuthTypeOk

`func (o *GetApplication200Response) GetStepUpAuthTypeOk() (*string, bool)`

GetStepUpAuthTypeOk returns a tuple with the StepUpAuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepUpAuthType

`func (o *GetApplication200Response) SetStepUpAuthType(v string)`

SetStepUpAuthType sets StepUpAuthType field to given value.

### HasStepUpAuthType

`func (o *GetApplication200Response) HasStepUpAuthType() bool`

HasStepUpAuthType returns a boolean if a field has been set.

### GetUsageAnalytics

`func (o *GetApplication200Response) GetUsageAnalytics() bool`

GetUsageAnalytics returns the UsageAnalytics field if non-nil, zero value otherwise.

### GetUsageAnalyticsOk

`func (o *GetApplication200Response) GetUsageAnalyticsOk() (*bool, bool)`

GetUsageAnalyticsOk returns a tuple with the UsageAnalytics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageAnalytics

`func (o *GetApplication200Response) SetUsageAnalytics(v bool)`

SetUsageAnalytics sets UsageAnalytics field to given value.

### HasUsageAnalytics

`func (o *GetApplication200Response) HasUsageAnalytics() bool`

HasUsageAnalytics returns a boolean if a field has been set.

### GetUsageCertRequired

`func (o *GetApplication200Response) GetUsageCertRequired() bool`

GetUsageCertRequired returns the UsageCertRequired field if non-nil, zero value otherwise.

### GetUsageCertRequiredOk

`func (o *GetApplication200Response) GetUsageCertRequiredOk() (*bool, bool)`

GetUsageCertRequiredOk returns a tuple with the UsageCertRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageCertRequired

`func (o *GetApplication200Response) SetUsageCertRequired(v bool)`

SetUsageCertRequired sets UsageCertRequired field to given value.

### HasUsageCertRequired

`func (o *GetApplication200Response) HasUsageCertRequired() bool`

HasUsageCertRequired returns a boolean if a field has been set.

### GetUsageCertText

`func (o *GetApplication200Response) GetUsageCertText() map[string]interface{}`

GetUsageCertText returns the UsageCertText field if non-nil, zero value otherwise.

### GetUsageCertTextOk

`func (o *GetApplication200Response) GetUsageCertTextOk() (*map[string]interface{}, bool)`

GetUsageCertTextOk returns a tuple with the UsageCertText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageCertText

`func (o *GetApplication200Response) SetUsageCertText(v map[string]interface{})`

SetUsageCertText sets UsageCertText field to given value.

### HasUsageCertText

`func (o *GetApplication200Response) HasUsageCertText() bool`

HasUsageCertText returns a boolean if a field has been set.

### GetLaunchpadEnabled

`func (o *GetApplication200Response) GetLaunchpadEnabled() bool`

GetLaunchpadEnabled returns the LaunchpadEnabled field if non-nil, zero value otherwise.

### GetLaunchpadEnabledOk

`func (o *GetApplication200Response) GetLaunchpadEnabledOk() (*bool, bool)`

GetLaunchpadEnabledOk returns a tuple with the LaunchpadEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchpadEnabled

`func (o *GetApplication200Response) SetLaunchpadEnabled(v bool)`

SetLaunchpadEnabled sets LaunchpadEnabled field to given value.

### HasLaunchpadEnabled

`func (o *GetApplication200Response) HasLaunchpadEnabled() bool`

HasLaunchpadEnabled returns a boolean if a field has been set.

### GetPasswordManaged

`func (o *GetApplication200Response) GetPasswordManaged() bool`

GetPasswordManaged returns the PasswordManaged field if non-nil, zero value otherwise.

### GetPasswordManagedOk

`func (o *GetApplication200Response) GetPasswordManagedOk() (*bool, bool)`

GetPasswordManagedOk returns a tuple with the PasswordManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordManaged

`func (o *GetApplication200Response) SetPasswordManaged(v bool)`

SetPasswordManaged sets PasswordManaged field to given value.

### HasPasswordManaged

`func (o *GetApplication200Response) HasPasswordManaged() bool`

HasPasswordManaged returns a boolean if a field has been set.

### GetOwner

`func (o *GetApplication200Response) GetOwner() ListApplications200ResponseInnerOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *GetApplication200Response) GetOwnerOk() (*ListApplications200ResponseInnerOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *GetApplication200Response) SetOwner(v ListApplications200ResponseInnerOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *GetApplication200Response) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetApplication200Response) GetDateCreated() float32`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetApplication200Response) GetDateCreatedOk() (*float32, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetApplication200Response) SetDateCreated(v float32)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetApplication200Response) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetApplication200Response) GetLastUpdated() float32`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetApplication200Response) GetLastUpdatedOk() (*float32, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetApplication200Response) SetLastUpdated(v float32)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetApplication200Response) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetDefaultAccessProfile

`func (o *GetApplication200Response) GetDefaultAccessProfile() map[string]interface{}`

GetDefaultAccessProfile returns the DefaultAccessProfile field if non-nil, zero value otherwise.

### GetDefaultAccessProfileOk

`func (o *GetApplication200Response) GetDefaultAccessProfileOk() (*map[string]interface{}, bool)`

GetDefaultAccessProfileOk returns a tuple with the DefaultAccessProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAccessProfile

`func (o *GetApplication200Response) SetDefaultAccessProfile(v map[string]interface{})`

SetDefaultAccessProfile sets DefaultAccessProfile field to given value.

### HasDefaultAccessProfile

`func (o *GetApplication200Response) HasDefaultAccessProfile() bool`

HasDefaultAccessProfile returns a boolean if a field has been set.

### GetService

`func (o *GetApplication200Response) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *GetApplication200Response) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *GetApplication200Response) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *GetApplication200Response) HasService() bool`

HasService returns a boolean if a field has been set.

### GetSelectedSsoMethod

`func (o *GetApplication200Response) GetSelectedSsoMethod() string`

GetSelectedSsoMethod returns the SelectedSsoMethod field if non-nil, zero value otherwise.

### GetSelectedSsoMethodOk

`func (o *GetApplication200Response) GetSelectedSsoMethodOk() (*string, bool)`

GetSelectedSsoMethodOk returns a tuple with the SelectedSsoMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedSsoMethod

`func (o *GetApplication200Response) SetSelectedSsoMethod(v string)`

SetSelectedSsoMethod sets SelectedSsoMethod field to given value.

### HasSelectedSsoMethod

`func (o *GetApplication200Response) HasSelectedSsoMethod() bool`

HasSelectedSsoMethod returns a boolean if a field has been set.

### GetSupportedSsoMethods

`func (o *GetApplication200Response) GetSupportedSsoMethods() float32`

GetSupportedSsoMethods returns the SupportedSsoMethods field if non-nil, zero value otherwise.

### GetSupportedSsoMethodsOk

`func (o *GetApplication200Response) GetSupportedSsoMethodsOk() (*float32, bool)`

GetSupportedSsoMethodsOk returns a tuple with the SupportedSsoMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedSsoMethods

`func (o *GetApplication200Response) SetSupportedSsoMethods(v float32)`

SetSupportedSsoMethods sets SupportedSsoMethods field to given value.

### HasSupportedSsoMethods

`func (o *GetApplication200Response) HasSupportedSsoMethods() bool`

HasSupportedSsoMethods returns a boolean if a field has been set.

### GetOffNetworkBlockedRoles

`func (o *GetApplication200Response) GetOffNetworkBlockedRoles() map[string]interface{}`

GetOffNetworkBlockedRoles returns the OffNetworkBlockedRoles field if non-nil, zero value otherwise.

### GetOffNetworkBlockedRolesOk

`func (o *GetApplication200Response) GetOffNetworkBlockedRolesOk() (*map[string]interface{}, bool)`

GetOffNetworkBlockedRolesOk returns a tuple with the OffNetworkBlockedRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffNetworkBlockedRoles

`func (o *GetApplication200Response) SetOffNetworkBlockedRoles(v map[string]interface{})`

SetOffNetworkBlockedRoles sets OffNetworkBlockedRoles field to given value.

### HasOffNetworkBlockedRoles

`func (o *GetApplication200Response) HasOffNetworkBlockedRoles() bool`

HasOffNetworkBlockedRoles returns a boolean if a field has been set.

### GetSupportedOffNetwork

`func (o *GetApplication200Response) GetSupportedOffNetwork() string`

GetSupportedOffNetwork returns the SupportedOffNetwork field if non-nil, zero value otherwise.

### GetSupportedOffNetworkOk

`func (o *GetApplication200Response) GetSupportedOffNetworkOk() (*string, bool)`

GetSupportedOffNetworkOk returns a tuple with the SupportedOffNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedOffNetwork

`func (o *GetApplication200Response) SetSupportedOffNetwork(v string)`

SetSupportedOffNetwork sets SupportedOffNetwork field to given value.

### HasSupportedOffNetwork

`func (o *GetApplication200Response) HasSupportedOffNetwork() bool`

HasSupportedOffNetwork returns a boolean if a field has been set.

### GetAccountServiceId

`func (o *GetApplication200Response) GetAccountServiceId() float32`

GetAccountServiceId returns the AccountServiceId field if non-nil, zero value otherwise.

### GetAccountServiceIdOk

`func (o *GetApplication200Response) GetAccountServiceIdOk() (*float32, bool)`

GetAccountServiceIdOk returns a tuple with the AccountServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountServiceId

`func (o *GetApplication200Response) SetAccountServiceId(v float32)`

SetAccountServiceId sets AccountServiceId field to given value.

### HasAccountServiceId

`func (o *GetApplication200Response) HasAccountServiceId() bool`

HasAccountServiceId returns a boolean if a field has been set.

### GetLauncherCount

`func (o *GetApplication200Response) GetLauncherCount() float32`

GetLauncherCount returns the LauncherCount field if non-nil, zero value otherwise.

### GetLauncherCountOk

`func (o *GetApplication200Response) GetLauncherCountOk() (*float32, bool)`

GetLauncherCountOk returns a tuple with the LauncherCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLauncherCount

`func (o *GetApplication200Response) SetLauncherCount(v float32)`

SetLauncherCount sets LauncherCount field to given value.

### HasLauncherCount

`func (o *GetApplication200Response) HasLauncherCount() bool`

HasLauncherCount returns a boolean if a field has been set.

### GetAccountServiceName

`func (o *GetApplication200Response) GetAccountServiceName() string`

GetAccountServiceName returns the AccountServiceName field if non-nil, zero value otherwise.

### GetAccountServiceNameOk

`func (o *GetApplication200Response) GetAccountServiceNameOk() (*string, bool)`

GetAccountServiceNameOk returns a tuple with the AccountServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountServiceName

`func (o *GetApplication200Response) SetAccountServiceName(v string)`

SetAccountServiceName sets AccountServiceName field to given value.

### HasAccountServiceName

`func (o *GetApplication200Response) HasAccountServiceName() bool`

HasAccountServiceName returns a boolean if a field has been set.

### GetAccountServiceExternalId

`func (o *GetApplication200Response) GetAccountServiceExternalId() string`

GetAccountServiceExternalId returns the AccountServiceExternalId field if non-nil, zero value otherwise.

### GetAccountServiceExternalIdOk

`func (o *GetApplication200Response) GetAccountServiceExternalIdOk() (*string, bool)`

GetAccountServiceExternalIdOk returns a tuple with the AccountServiceExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountServiceExternalId

`func (o *GetApplication200Response) SetAccountServiceExternalId(v string)`

SetAccountServiceExternalId sets AccountServiceExternalId field to given value.

### HasAccountServiceExternalId

`func (o *GetApplication200Response) HasAccountServiceExternalId() bool`

HasAccountServiceExternalId returns a boolean if a field has been set.

### GetAccountServiceMatchAllAccounts

`func (o *GetApplication200Response) GetAccountServiceMatchAllAccounts() bool`

GetAccountServiceMatchAllAccounts returns the AccountServiceMatchAllAccounts field if non-nil, zero value otherwise.

### GetAccountServiceMatchAllAccountsOk

`func (o *GetApplication200Response) GetAccountServiceMatchAllAccountsOk() (*bool, bool)`

GetAccountServiceMatchAllAccountsOk returns a tuple with the AccountServiceMatchAllAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountServiceMatchAllAccounts

`func (o *GetApplication200Response) SetAccountServiceMatchAllAccounts(v bool)`

SetAccountServiceMatchAllAccounts sets AccountServiceMatchAllAccounts field to given value.

### HasAccountServiceMatchAllAccounts

`func (o *GetApplication200Response) HasAccountServiceMatchAllAccounts() bool`

HasAccountServiceMatchAllAccounts returns a boolean if a field has been set.

### GetExternalId

`func (o *GetApplication200Response) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetApplication200Response) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetApplication200Response) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GetApplication200Response) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetAccountServiceUseForPasswordManagement

`func (o *GetApplication200Response) GetAccountServiceUseForPasswordManagement() bool`

GetAccountServiceUseForPasswordManagement returns the AccountServiceUseForPasswordManagement field if non-nil, zero value otherwise.

### GetAccountServiceUseForPasswordManagementOk

`func (o *GetApplication200Response) GetAccountServiceUseForPasswordManagementOk() (*bool, bool)`

GetAccountServiceUseForPasswordManagementOk returns a tuple with the AccountServiceUseForPasswordManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountServiceUseForPasswordManagement

`func (o *GetApplication200Response) SetAccountServiceUseForPasswordManagement(v bool)`

SetAccountServiceUseForPasswordManagement sets AccountServiceUseForPasswordManagement field to given value.

### HasAccountServiceUseForPasswordManagement

`func (o *GetApplication200Response) HasAccountServiceUseForPasswordManagement() bool`

HasAccountServiceUseForPasswordManagement returns a boolean if a field has been set.

### GetAccountServicePolicyId

`func (o *GetApplication200Response) GetAccountServicePolicyId() string`

GetAccountServicePolicyId returns the AccountServicePolicyId field if non-nil, zero value otherwise.

### GetAccountServicePolicyIdOk

`func (o *GetApplication200Response) GetAccountServicePolicyIdOk() (*string, bool)`

GetAccountServicePolicyIdOk returns a tuple with the AccountServicePolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountServicePolicyId

`func (o *GetApplication200Response) SetAccountServicePolicyId(v string)`

SetAccountServicePolicyId sets AccountServicePolicyId field to given value.

### HasAccountServicePolicyId

`func (o *GetApplication200Response) HasAccountServicePolicyId() bool`

HasAccountServicePolicyId returns a boolean if a field has been set.

### GetAccountServicePolicyName

`func (o *GetApplication200Response) GetAccountServicePolicyName() string`

GetAccountServicePolicyName returns the AccountServicePolicyName field if non-nil, zero value otherwise.

### GetAccountServicePolicyNameOk

`func (o *GetApplication200Response) GetAccountServicePolicyNameOk() (*string, bool)`

GetAccountServicePolicyNameOk returns a tuple with the AccountServicePolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountServicePolicyName

`func (o *GetApplication200Response) SetAccountServicePolicyName(v string)`

SetAccountServicePolicyName sets AccountServicePolicyName field to given value.

### HasAccountServicePolicyName

`func (o *GetApplication200Response) HasAccountServicePolicyName() bool`

HasAccountServicePolicyName returns a boolean if a field has been set.

### GetRequireStrongAuthn

`func (o *GetApplication200Response) GetRequireStrongAuthn() bool`

GetRequireStrongAuthn returns the RequireStrongAuthn field if non-nil, zero value otherwise.

### GetRequireStrongAuthnOk

`func (o *GetApplication200Response) GetRequireStrongAuthnOk() (*bool, bool)`

GetRequireStrongAuthnOk returns a tuple with the RequireStrongAuthn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireStrongAuthn

`func (o *GetApplication200Response) SetRequireStrongAuthn(v bool)`

SetRequireStrongAuthn sets RequireStrongAuthn field to given value.

### HasRequireStrongAuthn

`func (o *GetApplication200Response) HasRequireStrongAuthn() bool`

HasRequireStrongAuthn returns a boolean if a field has been set.

### GetAccountServicePolicies

`func (o *GetApplication200Response) GetAccountServicePolicies() []ListApplications200ResponseInnerAccountServicePoliciesInner`

GetAccountServicePolicies returns the AccountServicePolicies field if non-nil, zero value otherwise.

### GetAccountServicePoliciesOk

`func (o *GetApplication200Response) GetAccountServicePoliciesOk() (*[]ListApplications200ResponseInnerAccountServicePoliciesInner, bool)`

GetAccountServicePoliciesOk returns a tuple with the AccountServicePolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountServicePolicies

`func (o *GetApplication200Response) SetAccountServicePolicies(v []ListApplications200ResponseInnerAccountServicePoliciesInner)`

SetAccountServicePolicies sets AccountServicePolicies field to given value.

### HasAccountServicePolicies

`func (o *GetApplication200Response) HasAccountServicePolicies() bool`

HasAccountServicePolicies returns a boolean if a field has been set.

### GetXsdVersion

`func (o *GetApplication200Response) GetXsdVersion() string`

GetXsdVersion returns the XsdVersion field if non-nil, zero value otherwise.

### GetXsdVersionOk

`func (o *GetApplication200Response) GetXsdVersionOk() (*string, bool)`

GetXsdVersionOk returns a tuple with the XsdVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXsdVersion

`func (o *GetApplication200Response) SetXsdVersion(v string)`

SetXsdVersion sets XsdVersion field to given value.

### HasXsdVersion

`func (o *GetApplication200Response) HasXsdVersion() bool`

HasXsdVersion returns a boolean if a field has been set.

### GetAppProfiles

`func (o *GetApplication200Response) GetAppProfiles() []ListApplications200ResponseInnerAppProfilesInner`

GetAppProfiles returns the AppProfiles field if non-nil, zero value otherwise.

### GetAppProfilesOk

`func (o *GetApplication200Response) GetAppProfilesOk() (*[]ListApplications200ResponseInnerAppProfilesInner, bool)`

GetAppProfilesOk returns a tuple with the AppProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppProfiles

`func (o *GetApplication200Response) SetAppProfiles(v []ListApplications200ResponseInnerAppProfilesInner)`

SetAppProfiles sets AppProfiles field to given value.

### HasAppProfiles

`func (o *GetApplication200Response) HasAppProfiles() bool`

HasAppProfiles returns a boolean if a field has been set.

### GetPasswordServiceId

`func (o *GetApplication200Response) GetPasswordServiceId() float32`

GetPasswordServiceId returns the PasswordServiceId field if non-nil, zero value otherwise.

### GetPasswordServiceIdOk

`func (o *GetApplication200Response) GetPasswordServiceIdOk() (*float32, bool)`

GetPasswordServiceIdOk returns a tuple with the PasswordServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordServiceId

`func (o *GetApplication200Response) SetPasswordServiceId(v float32)`

SetPasswordServiceId sets PasswordServiceId field to given value.

### HasPasswordServiceId

`func (o *GetApplication200Response) HasPasswordServiceId() bool`

HasPasswordServiceId returns a boolean if a field has been set.

### GetAccessProfileIds

`func (o *GetApplication200Response) GetAccessProfileIds() map[string]interface{}`

GetAccessProfileIds returns the AccessProfileIds field if non-nil, zero value otherwise.

### GetAccessProfileIdsOk

`func (o *GetApplication200Response) GetAccessProfileIdsOk() (*map[string]interface{}, bool)`

GetAccessProfileIdsOk returns a tuple with the AccessProfileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessProfileIds

`func (o *GetApplication200Response) SetAccessProfileIds(v map[string]interface{})`

SetAccessProfileIds sets AccessProfileIds field to given value.

### HasAccessProfileIds

`func (o *GetApplication200Response) HasAccessProfileIds() bool`

HasAccessProfileIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


