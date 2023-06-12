# CampaignAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | Created time of the campaign | [optional] [readonly] 
**Modified** | Pointer to **time.Time** | Modified time of the campaign | [optional] [readonly] 
**CorrelatedStatus** | Pointer to **map[string]interface{}** | The correlatedStatus of the campaign. Only SOURCE_OWNER campaigns can be Uncorrelated. An Uncorrelated certification campaign only includes Uncorrelated identities (An identity is uncorrelated if it has no accounts on an authoritative source). | [optional] 
**Filter** | Pointer to [**CampaignAllOfFilter**](CampaignAllOfFilter.md) |  | [optional] 
**SunsetCommentsRequired** | Pointer to **bool** | Determines if comments on sunset date changes are required. | [optional] [default to true]
**SourceOwnerCampaignInfo** | Pointer to [**CampaignAllOfSourceOwnerCampaignInfo**](CampaignAllOfSourceOwnerCampaignInfo.md) |  | [optional] 
**SearchCampaignInfo** | Pointer to [**CampaignAllOfSearchCampaignInfo**](CampaignAllOfSearchCampaignInfo.md) |  | [optional] 
**RoleCompositionCampaignInfo** | Pointer to [**CampaignAllOfRoleCompositionCampaignInfo**](CampaignAllOfRoleCompositionCampaignInfo.md) |  | [optional] 
**Alerts** | Pointer to [**[]CampaignAlert**](CampaignAlert.md) | A list of errors and warnings that have accumulated. | [optional] [readonly] 
**TotalCertifications** | Pointer to **int32** | The total number of certifications in this campaign. | [optional] [readonly] 
**CompletedCertifications** | Pointer to **int32** | The number of completed certifications in this campaign. | [optional] [readonly] 
**SourcesWithOrphanEntitlements** | Pointer to [**[]CampaignAllOfSourcesWithOrphanEntitlements**](CampaignAllOfSourcesWithOrphanEntitlements.md) | A list of sources in the campaign that contain \\\&quot;orphan entitlements\\\&quot; (entitlements without a corresponding Managed Attribute). An empty list indicates the campaign has no orphan entitlements. Null indicates there may be unknown orphan entitlements in the campaign (the campaign was created before this feature was implemented). | [optional] [readonly] 
**MandatoryCommentRequirement** | Pointer to **string** | Determines whether comments are required for decisions during certification reviews. You can require comments for all decisions, revoke-only decisions, or no decisions. By default, comments are not required for decisions. | [optional] 

## Methods

### NewCampaignAllOf

`func NewCampaignAllOf() *CampaignAllOf`

NewCampaignAllOf instantiates a new CampaignAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignAllOfWithDefaults

`func NewCampaignAllOfWithDefaults() *CampaignAllOf`

NewCampaignAllOfWithDefaults instantiates a new CampaignAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *CampaignAllOf) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CampaignAllOf) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CampaignAllOf) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *CampaignAllOf) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetModified

`func (o *CampaignAllOf) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *CampaignAllOf) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *CampaignAllOf) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *CampaignAllOf) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetCorrelatedStatus

`func (o *CampaignAllOf) GetCorrelatedStatus() map[string]interface{}`

GetCorrelatedStatus returns the CorrelatedStatus field if non-nil, zero value otherwise.

### GetCorrelatedStatusOk

`func (o *CampaignAllOf) GetCorrelatedStatusOk() (*map[string]interface{}, bool)`

GetCorrelatedStatusOk returns a tuple with the CorrelatedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelatedStatus

`func (o *CampaignAllOf) SetCorrelatedStatus(v map[string]interface{})`

SetCorrelatedStatus sets CorrelatedStatus field to given value.

### HasCorrelatedStatus

`func (o *CampaignAllOf) HasCorrelatedStatus() bool`

HasCorrelatedStatus returns a boolean if a field has been set.

### GetFilter

`func (o *CampaignAllOf) GetFilter() CampaignAllOfFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *CampaignAllOf) GetFilterOk() (*CampaignAllOfFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *CampaignAllOf) SetFilter(v CampaignAllOfFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *CampaignAllOf) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetSunsetCommentsRequired

`func (o *CampaignAllOf) GetSunsetCommentsRequired() bool`

GetSunsetCommentsRequired returns the SunsetCommentsRequired field if non-nil, zero value otherwise.

### GetSunsetCommentsRequiredOk

`func (o *CampaignAllOf) GetSunsetCommentsRequiredOk() (*bool, bool)`

GetSunsetCommentsRequiredOk returns a tuple with the SunsetCommentsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSunsetCommentsRequired

`func (o *CampaignAllOf) SetSunsetCommentsRequired(v bool)`

SetSunsetCommentsRequired sets SunsetCommentsRequired field to given value.

### HasSunsetCommentsRequired

`func (o *CampaignAllOf) HasSunsetCommentsRequired() bool`

HasSunsetCommentsRequired returns a boolean if a field has been set.

### GetSourceOwnerCampaignInfo

`func (o *CampaignAllOf) GetSourceOwnerCampaignInfo() CampaignAllOfSourceOwnerCampaignInfo`

GetSourceOwnerCampaignInfo returns the SourceOwnerCampaignInfo field if non-nil, zero value otherwise.

### GetSourceOwnerCampaignInfoOk

`func (o *CampaignAllOf) GetSourceOwnerCampaignInfoOk() (*CampaignAllOfSourceOwnerCampaignInfo, bool)`

GetSourceOwnerCampaignInfoOk returns a tuple with the SourceOwnerCampaignInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceOwnerCampaignInfo

`func (o *CampaignAllOf) SetSourceOwnerCampaignInfo(v CampaignAllOfSourceOwnerCampaignInfo)`

SetSourceOwnerCampaignInfo sets SourceOwnerCampaignInfo field to given value.

### HasSourceOwnerCampaignInfo

`func (o *CampaignAllOf) HasSourceOwnerCampaignInfo() bool`

HasSourceOwnerCampaignInfo returns a boolean if a field has been set.

### GetSearchCampaignInfo

`func (o *CampaignAllOf) GetSearchCampaignInfo() CampaignAllOfSearchCampaignInfo`

GetSearchCampaignInfo returns the SearchCampaignInfo field if non-nil, zero value otherwise.

### GetSearchCampaignInfoOk

`func (o *CampaignAllOf) GetSearchCampaignInfoOk() (*CampaignAllOfSearchCampaignInfo, bool)`

GetSearchCampaignInfoOk returns a tuple with the SearchCampaignInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchCampaignInfo

`func (o *CampaignAllOf) SetSearchCampaignInfo(v CampaignAllOfSearchCampaignInfo)`

SetSearchCampaignInfo sets SearchCampaignInfo field to given value.

### HasSearchCampaignInfo

`func (o *CampaignAllOf) HasSearchCampaignInfo() bool`

HasSearchCampaignInfo returns a boolean if a field has been set.

### GetRoleCompositionCampaignInfo

`func (o *CampaignAllOf) GetRoleCompositionCampaignInfo() CampaignAllOfRoleCompositionCampaignInfo`

GetRoleCompositionCampaignInfo returns the RoleCompositionCampaignInfo field if non-nil, zero value otherwise.

### GetRoleCompositionCampaignInfoOk

`func (o *CampaignAllOf) GetRoleCompositionCampaignInfoOk() (*CampaignAllOfRoleCompositionCampaignInfo, bool)`

GetRoleCompositionCampaignInfoOk returns a tuple with the RoleCompositionCampaignInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleCompositionCampaignInfo

`func (o *CampaignAllOf) SetRoleCompositionCampaignInfo(v CampaignAllOfRoleCompositionCampaignInfo)`

SetRoleCompositionCampaignInfo sets RoleCompositionCampaignInfo field to given value.

### HasRoleCompositionCampaignInfo

`func (o *CampaignAllOf) HasRoleCompositionCampaignInfo() bool`

HasRoleCompositionCampaignInfo returns a boolean if a field has been set.

### GetAlerts

`func (o *CampaignAllOf) GetAlerts() []CampaignAlert`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *CampaignAllOf) GetAlertsOk() (*[]CampaignAlert, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *CampaignAllOf) SetAlerts(v []CampaignAlert)`

SetAlerts sets Alerts field to given value.

### HasAlerts

`func (o *CampaignAllOf) HasAlerts() bool`

HasAlerts returns a boolean if a field has been set.

### GetTotalCertifications

`func (o *CampaignAllOf) GetTotalCertifications() int32`

GetTotalCertifications returns the TotalCertifications field if non-nil, zero value otherwise.

### GetTotalCertificationsOk

`func (o *CampaignAllOf) GetTotalCertificationsOk() (*int32, bool)`

GetTotalCertificationsOk returns a tuple with the TotalCertifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCertifications

`func (o *CampaignAllOf) SetTotalCertifications(v int32)`

SetTotalCertifications sets TotalCertifications field to given value.

### HasTotalCertifications

`func (o *CampaignAllOf) HasTotalCertifications() bool`

HasTotalCertifications returns a boolean if a field has been set.

### GetCompletedCertifications

`func (o *CampaignAllOf) GetCompletedCertifications() int32`

GetCompletedCertifications returns the CompletedCertifications field if non-nil, zero value otherwise.

### GetCompletedCertificationsOk

`func (o *CampaignAllOf) GetCompletedCertificationsOk() (*int32, bool)`

GetCompletedCertificationsOk returns a tuple with the CompletedCertifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedCertifications

`func (o *CampaignAllOf) SetCompletedCertifications(v int32)`

SetCompletedCertifications sets CompletedCertifications field to given value.

### HasCompletedCertifications

`func (o *CampaignAllOf) HasCompletedCertifications() bool`

HasCompletedCertifications returns a boolean if a field has been set.

### GetSourcesWithOrphanEntitlements

`func (o *CampaignAllOf) GetSourcesWithOrphanEntitlements() []CampaignAllOfSourcesWithOrphanEntitlements`

GetSourcesWithOrphanEntitlements returns the SourcesWithOrphanEntitlements field if non-nil, zero value otherwise.

### GetSourcesWithOrphanEntitlementsOk

`func (o *CampaignAllOf) GetSourcesWithOrphanEntitlementsOk() (*[]CampaignAllOfSourcesWithOrphanEntitlements, bool)`

GetSourcesWithOrphanEntitlementsOk returns a tuple with the SourcesWithOrphanEntitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcesWithOrphanEntitlements

`func (o *CampaignAllOf) SetSourcesWithOrphanEntitlements(v []CampaignAllOfSourcesWithOrphanEntitlements)`

SetSourcesWithOrphanEntitlements sets SourcesWithOrphanEntitlements field to given value.

### HasSourcesWithOrphanEntitlements

`func (o *CampaignAllOf) HasSourcesWithOrphanEntitlements() bool`

HasSourcesWithOrphanEntitlements returns a boolean if a field has been set.

### GetMandatoryCommentRequirement

`func (o *CampaignAllOf) GetMandatoryCommentRequirement() string`

GetMandatoryCommentRequirement returns the MandatoryCommentRequirement field if non-nil, zero value otherwise.

### GetMandatoryCommentRequirementOk

`func (o *CampaignAllOf) GetMandatoryCommentRequirementOk() (*string, bool)`

GetMandatoryCommentRequirementOk returns a tuple with the MandatoryCommentRequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatoryCommentRequirement

`func (o *CampaignAllOf) SetMandatoryCommentRequirement(v string)`

SetMandatoryCommentRequirement sets MandatoryCommentRequirement field to given value.

### HasMandatoryCommentRequirement

`func (o *CampaignAllOf) HasMandatoryCommentRequirement() bool`

HasMandatoryCommentRequirement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


