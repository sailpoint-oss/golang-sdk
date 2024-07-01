# CampaignFilterDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Campaign filter name. | 
**Description** | Pointer to **string** | Campaign filter description. | [optional] 
**Owner** | **NullableString** | Owner of the filter. This field automatically populates at creation time with the current user. | 
**Mode** | **map[string]interface{}** | Mode/type of filter, either the INCLUSION or EXCLUSION type. The INCLUSION type includes the data in generated campaigns  as per specified in the criteria, whereas the EXCLUSION type excludes the data in generated campaigns as per specified in criteria. | 
**CriteriaList** | Pointer to [**[]CampaignFilterDetailsCriteriaListInner**](CampaignFilterDetailsCriteriaListInner.md) | List of criteria. | [optional] 

## Methods

### NewCampaignFilterDetails

`func NewCampaignFilterDetails(name string, owner NullableString, mode map[string]interface{}, ) *CampaignFilterDetails`

NewCampaignFilterDetails instantiates a new CampaignFilterDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignFilterDetailsWithDefaults

`func NewCampaignFilterDetailsWithDefaults() *CampaignFilterDetails`

NewCampaignFilterDetailsWithDefaults instantiates a new CampaignFilterDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CampaignFilterDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CampaignFilterDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CampaignFilterDetails) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CampaignFilterDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CampaignFilterDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CampaignFilterDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CampaignFilterDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwner

`func (o *CampaignFilterDetails) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CampaignFilterDetails) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CampaignFilterDetails) SetOwner(v string)`

SetOwner sets Owner field to given value.


### SetOwnerNil

`func (o *CampaignFilterDetails) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *CampaignFilterDetails) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetMode

`func (o *CampaignFilterDetails) GetMode() map[string]interface{}`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *CampaignFilterDetails) GetModeOk() (*map[string]interface{}, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *CampaignFilterDetails) SetMode(v map[string]interface{})`

SetMode sets Mode field to given value.


### GetCriteriaList

`func (o *CampaignFilterDetails) GetCriteriaList() []CampaignFilterDetailsCriteriaListInner`

GetCriteriaList returns the CriteriaList field if non-nil, zero value otherwise.

### GetCriteriaListOk

`func (o *CampaignFilterDetails) GetCriteriaListOk() (*[]CampaignFilterDetailsCriteriaListInner, bool)`

GetCriteriaListOk returns a tuple with the CriteriaList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteriaList

`func (o *CampaignFilterDetails) SetCriteriaList(v []CampaignFilterDetailsCriteriaListInner)`

SetCriteriaList sets CriteriaList field to given value.

### HasCriteriaList

`func (o *CampaignFilterDetails) HasCriteriaList() bool`

HasCriteriaList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


