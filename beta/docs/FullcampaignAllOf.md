# FullcampaignAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Modified** | Pointer to **time.Time** | Modified time of the campaign | [optional] [readonly] 
**CorrelatedStatus** | Pointer to **map[string]interface{}** | The correlatedStatus of the campaign. Only SOURCE_OWNER campaigns can be Uncorrelated. An Uncorrelated certification campaign only includes Uncorrelated identities (An identity is uncorrelated if it has no accounts on an authoritative source). | [optional] 
**Filter** | Pointer to [**FullcampaignAllOfFilter**](FullcampaignAllOfFilter.md) |  | [optional] 
**SunsetCommentsRequired** | Pointer to **bool** | Determines if comments on sunset date changes are required. | [optional] [default to true]
**SourceOwnerCampaignInfo** | Pointer to [**FullcampaignAllOfSourceOwnerCampaignInfo**](FullcampaignAllOfSourceOwnerCampaignInfo.md) |  | [optional] 
**SearchCampaignInfo** | Pointer to [**FullcampaignAllOfSearchCampaignInfo**](FullcampaignAllOfSearchCampaignInfo.md) |  | [optional] 
**RoleCompositionCampaignInfo** | Pointer to [**FullcampaignAllOfRoleCompositionCampaignInfo**](FullcampaignAllOfRoleCompositionCampaignInfo.md) |  | [optional] 
**SourcesWithOrphanEntitlements** | Pointer to [**[]FullcampaignAllOfSourcesWithOrphanEntitlements**](FullcampaignAllOfSourcesWithOrphanEntitlements.md) | A list of sources in the campaign that contain \\\&quot;orphan entitlements\\\&quot; (entitlements without a corresponding Managed Attribute). An empty list indicates the campaign has no orphan entitlements. Null indicates there may be unknown orphan entitlements in the campaign (the campaign was created before this feature was implemented). | [optional] [readonly] 
**MandatoryCommentRequirement** | Pointer to **string** | Determines whether comments are required for decisions during certification reviews. You can require comments for all decisions, revoke-only decisions, or no decisions. By default, comments are not required for decisions. | [optional] 

## Methods

### NewFullcampaignAllOf

`func NewFullcampaignAllOf() *FullcampaignAllOf`

NewFullcampaignAllOf instantiates a new FullcampaignAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFullcampaignAllOfWithDefaults

`func NewFullcampaignAllOfWithDefaults() *FullcampaignAllOf`

NewFullcampaignAllOfWithDefaults instantiates a new FullcampaignAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModified

`func (o *FullcampaignAllOf) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *FullcampaignAllOf) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *FullcampaignAllOf) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *FullcampaignAllOf) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetCorrelatedStatus

`func (o *FullcampaignAllOf) GetCorrelatedStatus() map[string]interface{}`

GetCorrelatedStatus returns the CorrelatedStatus field if non-nil, zero value otherwise.

### GetCorrelatedStatusOk

`func (o *FullcampaignAllOf) GetCorrelatedStatusOk() (*map[string]interface{}, bool)`

GetCorrelatedStatusOk returns a tuple with the CorrelatedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelatedStatus

`func (o *FullcampaignAllOf) SetCorrelatedStatus(v map[string]interface{})`

SetCorrelatedStatus sets CorrelatedStatus field to given value.

### HasCorrelatedStatus

`func (o *FullcampaignAllOf) HasCorrelatedStatus() bool`

HasCorrelatedStatus returns a boolean if a field has been set.

### GetFilter

`func (o *FullcampaignAllOf) GetFilter() FullcampaignAllOfFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *FullcampaignAllOf) GetFilterOk() (*FullcampaignAllOfFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *FullcampaignAllOf) SetFilter(v FullcampaignAllOfFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *FullcampaignAllOf) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetSunsetCommentsRequired

`func (o *FullcampaignAllOf) GetSunsetCommentsRequired() bool`

GetSunsetCommentsRequired returns the SunsetCommentsRequired field if non-nil, zero value otherwise.

### GetSunsetCommentsRequiredOk

`func (o *FullcampaignAllOf) GetSunsetCommentsRequiredOk() (*bool, bool)`

GetSunsetCommentsRequiredOk returns a tuple with the SunsetCommentsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSunsetCommentsRequired

`func (o *FullcampaignAllOf) SetSunsetCommentsRequired(v bool)`

SetSunsetCommentsRequired sets SunsetCommentsRequired field to given value.

### HasSunsetCommentsRequired

`func (o *FullcampaignAllOf) HasSunsetCommentsRequired() bool`

HasSunsetCommentsRequired returns a boolean if a field has been set.

### GetSourceOwnerCampaignInfo

`func (o *FullcampaignAllOf) GetSourceOwnerCampaignInfo() FullcampaignAllOfSourceOwnerCampaignInfo`

GetSourceOwnerCampaignInfo returns the SourceOwnerCampaignInfo field if non-nil, zero value otherwise.

### GetSourceOwnerCampaignInfoOk

`func (o *FullcampaignAllOf) GetSourceOwnerCampaignInfoOk() (*FullcampaignAllOfSourceOwnerCampaignInfo, bool)`

GetSourceOwnerCampaignInfoOk returns a tuple with the SourceOwnerCampaignInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceOwnerCampaignInfo

`func (o *FullcampaignAllOf) SetSourceOwnerCampaignInfo(v FullcampaignAllOfSourceOwnerCampaignInfo)`

SetSourceOwnerCampaignInfo sets SourceOwnerCampaignInfo field to given value.

### HasSourceOwnerCampaignInfo

`func (o *FullcampaignAllOf) HasSourceOwnerCampaignInfo() bool`

HasSourceOwnerCampaignInfo returns a boolean if a field has been set.

### GetSearchCampaignInfo

`func (o *FullcampaignAllOf) GetSearchCampaignInfo() FullcampaignAllOfSearchCampaignInfo`

GetSearchCampaignInfo returns the SearchCampaignInfo field if non-nil, zero value otherwise.

### GetSearchCampaignInfoOk

`func (o *FullcampaignAllOf) GetSearchCampaignInfoOk() (*FullcampaignAllOfSearchCampaignInfo, bool)`

GetSearchCampaignInfoOk returns a tuple with the SearchCampaignInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchCampaignInfo

`func (o *FullcampaignAllOf) SetSearchCampaignInfo(v FullcampaignAllOfSearchCampaignInfo)`

SetSearchCampaignInfo sets SearchCampaignInfo field to given value.

### HasSearchCampaignInfo

`func (o *FullcampaignAllOf) HasSearchCampaignInfo() bool`

HasSearchCampaignInfo returns a boolean if a field has been set.

### GetRoleCompositionCampaignInfo

`func (o *FullcampaignAllOf) GetRoleCompositionCampaignInfo() FullcampaignAllOfRoleCompositionCampaignInfo`

GetRoleCompositionCampaignInfo returns the RoleCompositionCampaignInfo field if non-nil, zero value otherwise.

### GetRoleCompositionCampaignInfoOk

`func (o *FullcampaignAllOf) GetRoleCompositionCampaignInfoOk() (*FullcampaignAllOfRoleCompositionCampaignInfo, bool)`

GetRoleCompositionCampaignInfoOk returns a tuple with the RoleCompositionCampaignInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleCompositionCampaignInfo

`func (o *FullcampaignAllOf) SetRoleCompositionCampaignInfo(v FullcampaignAllOfRoleCompositionCampaignInfo)`

SetRoleCompositionCampaignInfo sets RoleCompositionCampaignInfo field to given value.

### HasRoleCompositionCampaignInfo

`func (o *FullcampaignAllOf) HasRoleCompositionCampaignInfo() bool`

HasRoleCompositionCampaignInfo returns a boolean if a field has been set.

### GetSourcesWithOrphanEntitlements

`func (o *FullcampaignAllOf) GetSourcesWithOrphanEntitlements() []FullcampaignAllOfSourcesWithOrphanEntitlements`

GetSourcesWithOrphanEntitlements returns the SourcesWithOrphanEntitlements field if non-nil, zero value otherwise.

### GetSourcesWithOrphanEntitlementsOk

`func (o *FullcampaignAllOf) GetSourcesWithOrphanEntitlementsOk() (*[]FullcampaignAllOfSourcesWithOrphanEntitlements, bool)`

GetSourcesWithOrphanEntitlementsOk returns a tuple with the SourcesWithOrphanEntitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcesWithOrphanEntitlements

`func (o *FullcampaignAllOf) SetSourcesWithOrphanEntitlements(v []FullcampaignAllOfSourcesWithOrphanEntitlements)`

SetSourcesWithOrphanEntitlements sets SourcesWithOrphanEntitlements field to given value.

### HasSourcesWithOrphanEntitlements

`func (o *FullcampaignAllOf) HasSourcesWithOrphanEntitlements() bool`

HasSourcesWithOrphanEntitlements returns a boolean if a field has been set.

### GetMandatoryCommentRequirement

`func (o *FullcampaignAllOf) GetMandatoryCommentRequirement() string`

GetMandatoryCommentRequirement returns the MandatoryCommentRequirement field if non-nil, zero value otherwise.

### GetMandatoryCommentRequirementOk

`func (o *FullcampaignAllOf) GetMandatoryCommentRequirementOk() (*string, bool)`

GetMandatoryCommentRequirementOk returns a tuple with the MandatoryCommentRequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatoryCommentRequirement

`func (o *FullcampaignAllOf) SetMandatoryCommentRequirement(v string)`

SetMandatoryCommentRequirement sets MandatoryCommentRequirement field to given value.

### HasMandatoryCommentRequirement

`func (o *FullcampaignAllOf) HasMandatoryCommentRequirement() bool`

HasMandatoryCommentRequirement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


