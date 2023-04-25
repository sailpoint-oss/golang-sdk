# ListApplications200ResponseInner

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

### NewListApplications200ResponseInner

`func NewListApplications200ResponseInner() *ListApplications200ResponseInner`

NewListApplications200ResponseInner instantiates a new ListApplications200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListApplications200ResponseInnerWithDefaults

`func NewListApplications200ResponseInnerWithDefaults() *ListApplications200ResponseInner`

NewListApplications200ResponseInnerWithDefaults instantiates a new ListApplications200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListApplications200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListApplications200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListApplications200ResponseInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ListApplications200ResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAppId

`func (o *ListApplications200ResponseInner) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *ListApplications200ResponseInner) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *ListApplications200ResponseInner) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *ListApplications200ResponseInner) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetServiceId

`func (o *ListApplications200ResponseInner) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ListApplications200ResponseInner) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ListApplications200ResponseInner) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *ListApplications200ResponseInner) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetServiceAppId

`func (o *ListApplications200ResponseInner) GetServiceAppId() string`

GetServiceAppId returns the ServiceAppId field if non-nil, zero value otherwise.

### GetServiceAppIdOk

`func (o *ListApplications200ResponseInner) GetServiceAppIdOk() (*string, bool)`

GetServiceAppIdOk returns a tuple with the ServiceAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAppId

`func (o *ListApplications200ResponseInner) SetServiceAppId(v string)`

SetServiceAppId sets ServiceAppId field to given value.

### HasServiceAppId

`func (o *ListApplications200ResponseInner) HasServiceAppId() bool`

HasServiceAppId returns a boolean if a field has been set.

### GetName

`func (o *ListApplications200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListApplications200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListApplications200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListApplications200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ListApplications200ResponseInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListApplications200ResponseInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListApplications200ResponseInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListApplications200ResponseInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAppCenterEnabled

`func (o *ListApplications200ResponseInner) GetAppCenterEnabled() bool`

GetAppCenterEnabled returns the AppCenterEnabled field if non-nil, zero value otherwise.

### GetAppCenterEnabledOk

`func (o *ListApplications200ResponseInner) GetAppCenterEnabledOk() (*bool, bool)`

GetAppCenterEnabledOk returns a tuple with the AppCenterEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppCenterEnabled

`func (o *ListApplications200ResponseInner) SetAppCenterEnabled(v bool)`

SetAppCenterEnabled sets AppCenterEnabled field to given value.

### HasAppCenterEnabled

`func (o *ListApplications200ResponseInner) HasAppCenterEnabled() bool`

HasAppCenterEnabled returns a boolean if a field has been set.

### GetProvisionRequestEnabled

`func (o *ListApplications200ResponseInner) GetProvisionRequestEnabled() bool`

GetProvisionRequestEnabled returns the ProvisionRequestEnabled field if non-nil, zero value otherwise.

### GetProvisionRequestEnabledOk

`func (o *ListApplications200ResponseInner) GetProvisionRequestEnabledOk() (*bool, bool)`

GetProvisionRequestEnabledOk returns a tuple with the ProvisionRequestEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionRequestEnabled

`func (o *ListApplications200ResponseInner) SetProvisionRequestEnabled(v bool)`

SetProvisionRequestEnabled sets ProvisionRequestEnabled field to given value.

### HasProvisionRequestEnabled

`func (o *ListApplications200ResponseInner) HasProvisionRequestEnabled() bool`

HasProvisionRequestEnabled returns a boolean if a field has been set.

### GetControlType

`func (o *ListApplications200ResponseInner) GetControlType() string`

GetControlType returns the ControlType field if non-nil, zero value otherwise.

### GetControlTypeOk

`func (o *ListApplications200ResponseInner) GetControlTypeOk() (*string, bool)`

GetControlTypeOk returns a tuple with the ControlType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlType

`func (o *ListApplications200ResponseInner) SetControlType(v string)`

SetControlType sets ControlType field to given value.

### HasControlType

`func (o *ListApplications200ResponseInner) HasControlType() bool`

HasControlType returns a boolean if a field has been set.

### GetMobile

`func (o *ListApplications200ResponseInner) GetMobile() bool`

GetMobile returns the Mobile field if non-nil, zero value otherwise.

### GetMobileOk

`func (o *ListApplications200ResponseInner) GetMobileOk() (*bool, bool)`

GetMobileOk returns a tuple with the Mobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobile

`func (o *ListApplications200ResponseInner) SetMobile(v bool)`

SetMobile sets Mobile field to given value.

### HasMobile

`func (o *ListApplications200ResponseInner) HasMobile() bool`

HasMobile returns a boolean if a field has been set.

### GetPrivateApp

`func (o *ListApplications200ResponseInner) GetPrivateApp() bool`

GetPrivateApp returns the PrivateApp field if non-nil, zero value otherwise.

### GetPrivateAppOk

`func (o *ListApplications200ResponseInner) GetPrivateAppOk() (*bool, bool)`

GetPrivateAppOk returns a tuple with the PrivateApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateApp

`func (o *ListApplications200ResponseInner) SetPrivateApp(v bool)`

SetPrivateApp sets PrivateApp field to given value.

### HasPrivateApp

`func (o *ListApplications200ResponseInner) HasPrivateApp() bool`

HasPrivateApp returns a boolean if a field has been set.

### GetScriptName

`func (o *ListApplications200ResponseInner) GetScriptName() string`

GetScriptName returns the ScriptName field if non-nil, zero value otherwise.

### GetScriptNameOk

`func (o *ListApplications200ResponseInner) GetScriptNameOk() (*string, bool)`

GetScriptNameOk returns a tuple with the ScriptName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptName

`func (o *ListApplications200ResponseInner) SetScriptName(v string)`

SetScriptName sets ScriptName field to given value.

### HasScriptName

`func (o *ListApplications200ResponseInner) HasScriptName() bool`

HasScriptName returns a boolean if a field has been set.

### GetStatus

`func (o *ListApplications200ResponseInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListApplications200ResponseInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListApplications200ResponseInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListApplications200ResponseInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetIcon

`func (o *ListApplications200ResponseInner) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *ListApplications200ResponseInner) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *ListApplications200ResponseInner) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *ListApplications200ResponseInner) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetHealth

`func (o *ListApplications200ResponseInner) GetHealth() ListApplications200ResponseInnerHealth`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *ListApplications200ResponseInner) GetHealthOk() (*ListApplications200ResponseInnerHealth, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *ListApplications200ResponseInner) SetHealth(v ListApplications200ResponseInnerHealth)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *ListApplications200ResponseInner) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetEnableSso

`func (o *ListApplications200ResponseInner) GetEnableSso() bool`

GetEnableSso returns the EnableSso field if non-nil, zero value otherwise.

### GetEnableSsoOk

`func (o *ListApplications200ResponseInner) GetEnableSsoOk() (*bool, bool)`

GetEnableSsoOk returns a tuple with the EnableSso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSso

`func (o *ListApplications200ResponseInner) SetEnableSso(v bool)`

SetEnableSso sets EnableSso field to given value.

### HasEnableSso

`func (o *ListApplications200ResponseInner) HasEnableSso() bool`

HasEnableSso returns a boolean if a field has been set.

### GetSsoMethod

`func (o *ListApplications200ResponseInner) GetSsoMethod() string`

GetSsoMethod returns the SsoMethod field if non-nil, zero value otherwise.

### GetSsoMethodOk

`func (o *ListApplications200ResponseInner) GetSsoMethodOk() (*string, bool)`

GetSsoMethodOk returns a tuple with the SsoMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoMethod

`func (o *ListApplications200ResponseInner) SetSsoMethod(v string)`

SetSsoMethod sets SsoMethod field to given value.

### HasSsoMethod

`func (o *ListApplications200ResponseInner) HasSsoMethod() bool`

HasSsoMethod returns a boolean if a field has been set.

### GetHasLinks

`func (o *ListApplications200ResponseInner) GetHasLinks() bool`

GetHasLinks returns the HasLinks field if non-nil, zero value otherwise.

### GetHasLinksOk

`func (o *ListApplications200ResponseInner) GetHasLinksOk() (*bool, bool)`

GetHasLinksOk returns a tuple with the HasLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasLinks

`func (o *ListApplications200ResponseInner) SetHasLinks(v bool)`

SetHasLinks sets HasLinks field to given value.

### HasHasLinks

`func (o *ListApplications200ResponseInner) HasHasLinks() bool`

HasHasLinks returns a boolean if a field has been set.

### GetHasAutomations

`func (o *ListApplications200ResponseInner) GetHasAutomations() bool`

GetHasAutomations returns the HasAutomations field if non-nil, zero value otherwise.

### GetHasAutomationsOk

`func (o *ListApplications200ResponseInner) GetHasAutomationsOk() (*bool, bool)`

GetHasAutomationsOk returns a tuple with the HasAutomations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAutomations

`func (o *ListApplications200ResponseInner) SetHasAutomations(v bool)`

SetHasAutomations sets HasAutomations field to given value.

### HasHasAutomations

`func (o *ListApplications200ResponseInner) HasHasAutomations() bool`

HasHasAutomations returns a boolean if a field has been set.

### GetStepUpAuthData

`func (o *ListApplications200ResponseInner) GetStepUpAuthData() map[string]interface{}`

GetStepUpAuthData returns the StepUpAuthData field if non-nil, zero value otherwise.

### GetStepUpAuthDataOk

`func (o *ListApplications200ResponseInner) GetStepUpAuthDataOk() (*map[string]interface{}, bool)`

GetStepUpAuthDataOk returns a tuple with the StepUpAuthData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepUpAuthData

`func (o *ListApplications200ResponseInner) SetStepUpAuthData(v map[string]interface{})`

SetStepUpAuthData sets StepUpAuthData field to given value.

### HasStepUpAuthData

`func (o *ListApplications200ResponseInner) HasStepUpAuthData() bool`

HasStepUpAuthData returns a boolean if a field has been set.

### GetStepUpAuthType

`func (o *ListApplications200ResponseInner) GetStepUpAuthType() string`

GetStepUpAuthType returns the StepUpAuthType field if non-nil, zero value otherwise.

### GetStepUpAuthTypeOk

`func (o *ListApplications200ResponseInner) GetStepUpAuthTypeOk() (*string, bool)`

GetStepUpAuthTypeOk returns a tuple with the StepUpAuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepUpAuthType

`func (o *ListApplications200ResponseInner) SetStepUpAuthType(v string)`

SetStepUpAuthType sets StepUpAuthType field to given value.

### HasStepUpAuthType

`func (o *ListApplications200ResponseInner) HasStepUpAuthType() bool`

HasStepUpAuthType returns a boolean if a field has been set.

### GetUsageAnalytics

`func (o *ListApplications200ResponseInner) GetUsageAnalytics() bool`

GetUsageAnalytics returns the UsageAnalytics field if non-nil, zero value otherwise.

### GetUsageAnalyticsOk

`func (o *ListApplications200ResponseInner) GetUsageAnalyticsOk() (*bool, bool)`

GetUsageAnalyticsOk returns a tuple with the UsageAnalytics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageAnalytics

`func (o *ListApplications200ResponseInner) SetUsageAnalytics(v bool)`

SetUsageAnalytics sets UsageAnalytics field to given value.

### HasUsageAnalytics

`func (o *ListApplications200ResponseInner) HasUsageAnalytics() bool`

HasUsageAnalytics returns a boolean if a field has been set.

### GetUsageCertRequired

`func (o *ListApplications200ResponseInner) GetUsageCertRequired() bool`

GetUsageCertRequired returns the UsageCertRequired field if non-nil, zero value otherwise.

### GetUsageCertRequiredOk

`func (o *ListApplications200ResponseInner) GetUsageCertRequiredOk() (*bool, bool)`

GetUsageCertRequiredOk returns a tuple with the UsageCertRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageCertRequired

`func (o *ListApplications200ResponseInner) SetUsageCertRequired(v bool)`

SetUsageCertRequired sets UsageCertRequired field to given value.

### HasUsageCertRequired

`func (o *ListApplications200ResponseInner) HasUsageCertRequired() bool`

HasUsageCertRequired returns a boolean if a field has been set.

### GetUsageCertText

`func (o *ListApplications200ResponseInner) GetUsageCertText() map[string]interface{}`

GetUsageCertText returns the UsageCertText field if non-nil, zero value otherwise.

### GetUsageCertTextOk

`func (o *ListApplications200ResponseInner) GetUsageCertTextOk() (*map[string]interface{}, bool)`

GetUsageCertTextOk returns a tuple with the UsageCertText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageCertText

`func (o *ListApplications200ResponseInner) SetUsageCertText(v map[string]interface{})`

SetUsageCertText sets UsageCertText field to given value.

### HasUsageCertText

`func (o *ListApplications200ResponseInner) HasUsageCertText() bool`

HasUsageCertText returns a boolean if a field has been set.

### GetLaunchpadEnabled

`func (o *ListApplications200ResponseInner) GetLaunchpadEnabled() bool`

GetLaunchpadEnabled returns the LaunchpadEnabled field if non-nil, zero value otherwise.

### GetLaunchpadEnabledOk

`func (o *ListApplications200ResponseInner) GetLaunchpadEnabledOk() (*bool, bool)`

GetLaunchpadEnabledOk returns a tuple with the LaunchpadEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchpadEnabled

`func (o *ListApplications200ResponseInner) SetLaunchpadEnabled(v bool)`

SetLaunchpadEnabled sets LaunchpadEnabled field to given value.

### HasLaunchpadEnabled

`func (o *ListApplications200ResponseInner) HasLaunchpadEnabled() bool`

HasLaunchpadEnabled returns a boolean if a field has been set.

### GetPasswordManaged

`func (o *ListApplications200ResponseInner) GetPasswordManaged() bool`

GetPasswordManaged returns the PasswordManaged field if non-nil, zero value otherwise.

### GetPasswordManagedOk

`func (o *ListApplications200ResponseInner) GetPasswordManagedOk() (*bool, bool)`

GetPasswordManagedOk returns a tuple with the PasswordManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordManaged

`func (o *ListApplications200ResponseInner) SetPasswordManaged(v bool)`

SetPasswordManaged sets PasswordManaged field to given value.

### HasPasswordManaged

`func (o *ListApplications200ResponseInner) HasPasswordManaged() bool`

HasPasswordManaged returns a boolean if a field has been set.

### GetOwner

`func (o *ListApplications200ResponseInner) GetOwner() ListApplications200ResponseInnerOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ListApplications200ResponseInner) GetOwnerOk() (*ListApplications200ResponseInnerOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ListApplications200ResponseInner) SetOwner(v ListApplications200ResponseInnerOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ListApplications200ResponseInner) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListApplications200ResponseInner) GetDateCreated() float32`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListApplications200ResponseInner) GetDateCreatedOk() (*float32, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListApplications200ResponseInner) SetDateCreated(v float32)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListApplications200ResponseInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListApplications200ResponseInner) GetLastUpdated() float32`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListApplications200ResponseInner) GetLastUpdatedOk() (*float32, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListApplications200ResponseInner) SetLastUpdated(v float32)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListApplications200ResponseInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetDefaultAccessProfile

`func (o *ListApplications200ResponseInner) GetDefaultAccessProfile() map[string]interface{}`

GetDefaultAccessProfile returns the DefaultAccessProfile field if non-nil, zero value otherwise.

### GetDefaultAccessProfileOk

`func (o *ListApplications200ResponseInner) GetDefaultAccessProfileOk() (*map[string]interface{}, bool)`

GetDefaultAccessProfileOk returns a tuple with the DefaultAccessProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAccessProfile

`func (o *ListApplications200ResponseInner) SetDefaultAccessProfile(v map[string]interface{})`

SetDefaultAccessProfile sets DefaultAccessProfile field to given value.

### HasDefaultAccessProfile

`func (o *ListApplications200ResponseInner) HasDefaultAccessProfile() bool`

HasDefaultAccessProfile returns a boolean if a field has been set.

### GetService

`func (o *ListApplications200ResponseInner) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ListApplications200ResponseInner) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *ListApplications200ResponseInner) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *ListApplications200ResponseInner) HasService() bool`

HasService returns a boolean if a field has been set.

### GetSelectedSsoMethod

`func (o *ListApplications200ResponseInner) GetSelectedSsoMethod() string`

GetSelectedSsoMethod returns the SelectedSsoMethod field if non-nil, zero value otherwise.

### GetSelectedSsoMethodOk

`func (o *ListApplications200ResponseInner) GetSelectedSsoMethodOk() (*string, bool)`

GetSelectedSsoMethodOk returns a tuple with the SelectedSsoMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedSsoMethod

`func (o *ListApplications200ResponseInner) SetSelectedSsoMethod(v string)`

SetSelectedSsoMethod sets SelectedSsoMethod field to given value.

### HasSelectedSsoMethod

`func (o *ListApplications200ResponseInner) HasSelectedSsoMethod() bool`

HasSelectedSsoMethod returns a boolean if a field has been set.

### GetSupportedSsoMethods

`func (o *ListApplications200ResponseInner) GetSupportedSsoMethods() float32`

GetSupportedSsoMethods returns the SupportedSsoMethods field if non-nil, zero value otherwise.

### GetSupportedSsoMethodsOk

`func (o *ListApplications200ResponseInner) GetSupportedSsoMethodsOk() (*float32, bool)`

GetSupportedSsoMethodsOk returns a tuple with the SupportedSsoMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedSsoMethods

`func (o *ListApplications200ResponseInner) SetSupportedSsoMethods(v float32)`

SetSupportedSsoMethods sets SupportedSsoMethods field to given value.

### HasSupportedSsoMethods

`func (o *ListApplications200ResponseInner) HasSupportedSsoMethods() bool`

HasSupportedSsoMethods returns a boolean if a field has been set.

### GetOffNetworkBlockedRoles

`func (o *ListApplications200ResponseInner) GetOffNetworkBlockedRoles() map[string]interface{}`

GetOffNetworkBlockedRoles returns the OffNetworkBlockedRoles field if non-nil, zero value otherwise.

### GetOffNetworkBlockedRolesOk

`func (o *ListApplications200ResponseInner) GetOffNetworkBlockedRolesOk() (*map[string]interface{}, bool)`

GetOffNetworkBlockedRolesOk returns a tuple with the OffNetworkBlockedRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffNetworkBlockedRoles

`func (o *ListApplications200ResponseInner) SetOffNetworkBlockedRoles(v map[string]interface{})`

SetOffNetworkBlockedRoles sets OffNetworkBlockedRoles field to given value.

### HasOffNetworkBlockedRoles

`func (o *ListApplications200ResponseInner) HasOffNetworkBlockedRoles() bool`

HasOffNetworkBlockedRoles returns a boolean if a field has been set.

### GetSupportedOffNetwork

`func (o *ListApplications200ResponseInner) GetSupportedOffNetwork() string`

GetSupportedOffNetwork returns the SupportedOffNetwork field if non-nil, zero value otherwise.

### GetSupportedOffNetworkOk

`func (o *ListApplications200ResponseInner) GetSupportedOffNetworkOk() (*string, bool)`

GetSupportedOffNetworkOk returns a tuple with the SupportedOffNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedOffNetwork

`func (o *ListApplications200ResponseInner) SetSupportedOffNetwork(v string)`

SetSupportedOffNetwork sets SupportedOffNetwork field to given value.

### HasSupportedOffNetwork

`func (o *ListApplications200ResponseInner) HasSupportedOffNetwork() bool`

HasSupportedOffNetwork returns a boolean if a field has been set.

### GetAccountServiceId

`func (o *ListApplications200ResponseInner) GetAccountServiceId() float32`

GetAccountServiceId returns the AccountServiceId field if non-nil, zero value otherwise.

### GetAccountServiceIdOk

`func (o *ListApplications200ResponseInner) GetAccountServiceIdOk() (*float32, bool)`

GetAccountServiceIdOk returns a tuple with the AccountServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountServiceId

`func (o *ListApplications200ResponseInner) SetAccountServiceId(v float32)`

SetAccountServiceId sets AccountServiceId field to given value.

### HasAccountServiceId

`func (o *ListApplications200ResponseInner) HasAccountServiceId() bool`

HasAccountServiceId returns a boolean if a field has been set.

### GetLauncherCount

`func (o *ListApplications200ResponseInner) GetLauncherCount() float32`

GetLauncherCount returns the LauncherCount field if non-nil, zero value otherwise.

### GetLauncherCountOk

`func (o *ListApplications200ResponseInner) GetLauncherCountOk() (*float32, bool)`

GetLauncherCountOk returns a tuple with the LauncherCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLauncherCount

`func (o *ListApplications200ResponseInner) SetLauncherCount(v float32)`

SetLauncherCount sets LauncherCount field to given value.

### HasLauncherCount

`func (o *ListApplications200ResponseInner) HasLauncherCount() bool`

HasLauncherCount returns a boolean if a field has been set.

### GetAccountServiceName

`func (o *ListApplications200ResponseInner) GetAccountServiceName() string`

GetAccountServiceName returns the AccountServiceName field if non-nil, zero value otherwise.

### GetAccountServiceNameOk

`func (o *ListApplications200ResponseInner) GetAccountServiceNameOk() (*string, bool)`

GetAccountServiceNameOk returns a tuple with the AccountServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountServiceName

`func (o *ListApplications200ResponseInner) SetAccountServiceName(v string)`

SetAccountServiceName sets AccountServiceName field to given value.

### HasAccountServiceName

`func (o *ListApplications200ResponseInner) HasAccountServiceName() bool`

HasAccountServiceName returns a boolean if a field has been set.

### GetAccountServiceExternalId

`func (o *ListApplications200ResponseInner) GetAccountServiceExternalId() string`

GetAccountServiceExternalId returns the AccountServiceExternalId field if non-nil, zero value otherwise.

### GetAccountServiceExternalIdOk

`func (o *ListApplications200ResponseInner) GetAccountServiceExternalIdOk() (*string, bool)`

GetAccountServiceExternalIdOk returns a tuple with the AccountServiceExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountServiceExternalId

`func (o *ListApplications200ResponseInner) SetAccountServiceExternalId(v string)`

SetAccountServiceExternalId sets AccountServiceExternalId field to given value.

### HasAccountServiceExternalId

`func (o *ListApplications200ResponseInner) HasAccountServiceExternalId() bool`

HasAccountServiceExternalId returns a boolean if a field has been set.

### GetAccountServiceMatchAllAccounts

`func (o *ListApplications200ResponseInner) GetAccountServiceMatchAllAccounts() bool`

GetAccountServiceMatchAllAccounts returns the AccountServiceMatchAllAccounts field if non-nil, zero value otherwise.

### GetAccountServiceMatchAllAccountsOk

`func (o *ListApplications200ResponseInner) GetAccountServiceMatchAllAccountsOk() (*bool, bool)`

GetAccountServiceMatchAllAccountsOk returns a tuple with the AccountServiceMatchAllAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountServiceMatchAllAccounts

`func (o *ListApplications200ResponseInner) SetAccountServiceMatchAllAccounts(v bool)`

SetAccountServiceMatchAllAccounts sets AccountServiceMatchAllAccounts field to given value.

### HasAccountServiceMatchAllAccounts

`func (o *ListApplications200ResponseInner) HasAccountServiceMatchAllAccounts() bool`

HasAccountServiceMatchAllAccounts returns a boolean if a field has been set.

### GetExternalId

`func (o *ListApplications200ResponseInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ListApplications200ResponseInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ListApplications200ResponseInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ListApplications200ResponseInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetAccountServiceUseForPasswordManagement

`func (o *ListApplications200ResponseInner) GetAccountServiceUseForPasswordManagement() bool`

GetAccountServiceUseForPasswordManagement returns the AccountServiceUseForPasswordManagement field if non-nil, zero value otherwise.

### GetAccountServiceUseForPasswordManagementOk

`func (o *ListApplications200ResponseInner) GetAccountServiceUseForPasswordManagementOk() (*bool, bool)`

GetAccountServiceUseForPasswordManagementOk returns a tuple with the AccountServiceUseForPasswordManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountServiceUseForPasswordManagement

`func (o *ListApplications200ResponseInner) SetAccountServiceUseForPasswordManagement(v bool)`

SetAccountServiceUseForPasswordManagement sets AccountServiceUseForPasswordManagement field to given value.

### HasAccountServiceUseForPasswordManagement

`func (o *ListApplications200ResponseInner) HasAccountServiceUseForPasswordManagement() bool`

HasAccountServiceUseForPasswordManagement returns a boolean if a field has been set.

### GetAccountServicePolicyId

`func (o *ListApplications200ResponseInner) GetAccountServicePolicyId() string`

GetAccountServicePolicyId returns the AccountServicePolicyId field if non-nil, zero value otherwise.

### GetAccountServicePolicyIdOk

`func (o *ListApplications200ResponseInner) GetAccountServicePolicyIdOk() (*string, bool)`

GetAccountServicePolicyIdOk returns a tuple with the AccountServicePolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountServicePolicyId

`func (o *ListApplications200ResponseInner) SetAccountServicePolicyId(v string)`

SetAccountServicePolicyId sets AccountServicePolicyId field to given value.

### HasAccountServicePolicyId

`func (o *ListApplications200ResponseInner) HasAccountServicePolicyId() bool`

HasAccountServicePolicyId returns a boolean if a field has been set.

### GetAccountServicePolicyName

`func (o *ListApplications200ResponseInner) GetAccountServicePolicyName() string`

GetAccountServicePolicyName returns the AccountServicePolicyName field if non-nil, zero value otherwise.

### GetAccountServicePolicyNameOk

`func (o *ListApplications200ResponseInner) GetAccountServicePolicyNameOk() (*string, bool)`

GetAccountServicePolicyNameOk returns a tuple with the AccountServicePolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountServicePolicyName

`func (o *ListApplications200ResponseInner) SetAccountServicePolicyName(v string)`

SetAccountServicePolicyName sets AccountServicePolicyName field to given value.

### HasAccountServicePolicyName

`func (o *ListApplications200ResponseInner) HasAccountServicePolicyName() bool`

HasAccountServicePolicyName returns a boolean if a field has been set.

### GetRequireStrongAuthn

`func (o *ListApplications200ResponseInner) GetRequireStrongAuthn() bool`

GetRequireStrongAuthn returns the RequireStrongAuthn field if non-nil, zero value otherwise.

### GetRequireStrongAuthnOk

`func (o *ListApplications200ResponseInner) GetRequireStrongAuthnOk() (*bool, bool)`

GetRequireStrongAuthnOk returns a tuple with the RequireStrongAuthn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireStrongAuthn

`func (o *ListApplications200ResponseInner) SetRequireStrongAuthn(v bool)`

SetRequireStrongAuthn sets RequireStrongAuthn field to given value.

### HasRequireStrongAuthn

`func (o *ListApplications200ResponseInner) HasRequireStrongAuthn() bool`

HasRequireStrongAuthn returns a boolean if a field has been set.

### GetAccountServicePolicies

`func (o *ListApplications200ResponseInner) GetAccountServicePolicies() []ListApplications200ResponseInnerAccountServicePoliciesInner`

GetAccountServicePolicies returns the AccountServicePolicies field if non-nil, zero value otherwise.

### GetAccountServicePoliciesOk

`func (o *ListApplications200ResponseInner) GetAccountServicePoliciesOk() (*[]ListApplications200ResponseInnerAccountServicePoliciesInner, bool)`

GetAccountServicePoliciesOk returns a tuple with the AccountServicePolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountServicePolicies

`func (o *ListApplications200ResponseInner) SetAccountServicePolicies(v []ListApplications200ResponseInnerAccountServicePoliciesInner)`

SetAccountServicePolicies sets AccountServicePolicies field to given value.

### HasAccountServicePolicies

`func (o *ListApplications200ResponseInner) HasAccountServicePolicies() bool`

HasAccountServicePolicies returns a boolean if a field has been set.

### GetXsdVersion

`func (o *ListApplications200ResponseInner) GetXsdVersion() string`

GetXsdVersion returns the XsdVersion field if non-nil, zero value otherwise.

### GetXsdVersionOk

`func (o *ListApplications200ResponseInner) GetXsdVersionOk() (*string, bool)`

GetXsdVersionOk returns a tuple with the XsdVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXsdVersion

`func (o *ListApplications200ResponseInner) SetXsdVersion(v string)`

SetXsdVersion sets XsdVersion field to given value.

### HasXsdVersion

`func (o *ListApplications200ResponseInner) HasXsdVersion() bool`

HasXsdVersion returns a boolean if a field has been set.

### GetAppProfiles

`func (o *ListApplications200ResponseInner) GetAppProfiles() []ListApplications200ResponseInnerAppProfilesInner`

GetAppProfiles returns the AppProfiles field if non-nil, zero value otherwise.

### GetAppProfilesOk

`func (o *ListApplications200ResponseInner) GetAppProfilesOk() (*[]ListApplications200ResponseInnerAppProfilesInner, bool)`

GetAppProfilesOk returns a tuple with the AppProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppProfiles

`func (o *ListApplications200ResponseInner) SetAppProfiles(v []ListApplications200ResponseInnerAppProfilesInner)`

SetAppProfiles sets AppProfiles field to given value.

### HasAppProfiles

`func (o *ListApplications200ResponseInner) HasAppProfiles() bool`

HasAppProfiles returns a boolean if a field has been set.

### GetPasswordServiceId

`func (o *ListApplications200ResponseInner) GetPasswordServiceId() float32`

GetPasswordServiceId returns the PasswordServiceId field if non-nil, zero value otherwise.

### GetPasswordServiceIdOk

`func (o *ListApplications200ResponseInner) GetPasswordServiceIdOk() (*float32, bool)`

GetPasswordServiceIdOk returns a tuple with the PasswordServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordServiceId

`func (o *ListApplications200ResponseInner) SetPasswordServiceId(v float32)`

SetPasswordServiceId sets PasswordServiceId field to given value.

### HasPasswordServiceId

`func (o *ListApplications200ResponseInner) HasPasswordServiceId() bool`

HasPasswordServiceId returns a boolean if a field has been set.

### GetAccessProfileIds

`func (o *ListApplications200ResponseInner) GetAccessProfileIds() map[string]interface{}`

GetAccessProfileIds returns the AccessProfileIds field if non-nil, zero value otherwise.

### GetAccessProfileIdsOk

`func (o *ListApplications200ResponseInner) GetAccessProfileIdsOk() (*map[string]interface{}, bool)`

GetAccessProfileIdsOk returns a tuple with the AccessProfileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessProfileIds

`func (o *ListApplications200ResponseInner) SetAccessProfileIds(v map[string]interface{})`

SetAccessProfileIds sets AccessProfileIds field to given value.

### HasAccessProfileIds

`func (o *ListApplications200ResponseInner) HasAccessProfileIds() bool`

HasAccessProfileIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


