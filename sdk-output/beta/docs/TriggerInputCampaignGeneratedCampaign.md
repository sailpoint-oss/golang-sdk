# TriggerInputCampaignGeneratedCampaign

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique ID of the campaign. | 
**Name** | **string** | Human friendly name of the campaign. | 
**Description** | **string** | Extended description of the campaign. | 
**Created** | **time.Time** | The date and time the campaign was created. | 
**Modified** | Pointer to **NullableString** | The date and time the campaign was last modified. | [optional] 
**Deadline** | Pointer to **NullableString** | The date and time when the campaign must be finished by. | [optional] 
**Type** | **map[string]interface{}** | The type of campaign that was generated. | 
**CampaignOwner** | [**TriggerInputCampaignGeneratedCampaignCampaignOwner**](TriggerInputCampaignGeneratedCampaignCampaignOwner.md) |  | 
**Status** | **map[string]interface{}** | The current status of the campaign. | 

## Methods

### NewTriggerInputCampaignGeneratedCampaign

`func NewTriggerInputCampaignGeneratedCampaign(id string, name string, description string, created time.Time, type_ map[string]interface{}, campaignOwner TriggerInputCampaignGeneratedCampaignCampaignOwner, status map[string]interface{}, ) *TriggerInputCampaignGeneratedCampaign`

NewTriggerInputCampaignGeneratedCampaign instantiates a new TriggerInputCampaignGeneratedCampaign object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerInputCampaignGeneratedCampaignWithDefaults

`func NewTriggerInputCampaignGeneratedCampaignWithDefaults() *TriggerInputCampaignGeneratedCampaign`

NewTriggerInputCampaignGeneratedCampaignWithDefaults instantiates a new TriggerInputCampaignGeneratedCampaign object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TriggerInputCampaignGeneratedCampaign) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TriggerInputCampaignGeneratedCampaign) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TriggerInputCampaignGeneratedCampaign) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *TriggerInputCampaignGeneratedCampaign) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TriggerInputCampaignGeneratedCampaign) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TriggerInputCampaignGeneratedCampaign) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *TriggerInputCampaignGeneratedCampaign) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TriggerInputCampaignGeneratedCampaign) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TriggerInputCampaignGeneratedCampaign) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCreated

`func (o *TriggerInputCampaignGeneratedCampaign) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *TriggerInputCampaignGeneratedCampaign) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *TriggerInputCampaignGeneratedCampaign) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetModified

`func (o *TriggerInputCampaignGeneratedCampaign) GetModified() string`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *TriggerInputCampaignGeneratedCampaign) GetModifiedOk() (*string, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *TriggerInputCampaignGeneratedCampaign) SetModified(v string)`

SetModified sets Modified field to given value.

### HasModified

`func (o *TriggerInputCampaignGeneratedCampaign) HasModified() bool`

HasModified returns a boolean if a field has been set.

### SetModifiedNil

`func (o *TriggerInputCampaignGeneratedCampaign) SetModifiedNil(b bool)`

 SetModifiedNil sets the value for Modified to be an explicit nil

### UnsetModified
`func (o *TriggerInputCampaignGeneratedCampaign) UnsetModified()`

UnsetModified ensures that no value is present for Modified, not even an explicit nil
### GetDeadline

`func (o *TriggerInputCampaignGeneratedCampaign) GetDeadline() string`

GetDeadline returns the Deadline field if non-nil, zero value otherwise.

### GetDeadlineOk

`func (o *TriggerInputCampaignGeneratedCampaign) GetDeadlineOk() (*string, bool)`

GetDeadlineOk returns a tuple with the Deadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadline

`func (o *TriggerInputCampaignGeneratedCampaign) SetDeadline(v string)`

SetDeadline sets Deadline field to given value.

### HasDeadline

`func (o *TriggerInputCampaignGeneratedCampaign) HasDeadline() bool`

HasDeadline returns a boolean if a field has been set.

### SetDeadlineNil

`func (o *TriggerInputCampaignGeneratedCampaign) SetDeadlineNil(b bool)`

 SetDeadlineNil sets the value for Deadline to be an explicit nil

### UnsetDeadline
`func (o *TriggerInputCampaignGeneratedCampaign) UnsetDeadline()`

UnsetDeadline ensures that no value is present for Deadline, not even an explicit nil
### GetType

`func (o *TriggerInputCampaignGeneratedCampaign) GetType() map[string]interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TriggerInputCampaignGeneratedCampaign) GetTypeOk() (*map[string]interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TriggerInputCampaignGeneratedCampaign) SetType(v map[string]interface{})`

SetType sets Type field to given value.


### GetCampaignOwner

`func (o *TriggerInputCampaignGeneratedCampaign) GetCampaignOwner() TriggerInputCampaignGeneratedCampaignCampaignOwner`

GetCampaignOwner returns the CampaignOwner field if non-nil, zero value otherwise.

### GetCampaignOwnerOk

`func (o *TriggerInputCampaignGeneratedCampaign) GetCampaignOwnerOk() (*TriggerInputCampaignGeneratedCampaignCampaignOwner, bool)`

GetCampaignOwnerOk returns a tuple with the CampaignOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignOwner

`func (o *TriggerInputCampaignGeneratedCampaign) SetCampaignOwner(v TriggerInputCampaignGeneratedCampaignCampaignOwner)`

SetCampaignOwner sets CampaignOwner field to given value.


### GetStatus

`func (o *TriggerInputCampaignGeneratedCampaign) GetStatus() map[string]interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TriggerInputCampaignGeneratedCampaign) GetStatusOk() (*map[string]interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TriggerInputCampaignGeneratedCampaign) SetStatus(v map[string]interface{})`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


