# TriggerInputCampaignEndedCampaign

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique ID for the campaign. | 
**Name** | **string** | The human friendly name of the campaign. | 
**Description** | **string** | Extended description of the campaign. | 
**Created** | **time.Time** | The date and time the campaign was created. | 
**Modified** | Pointer to **NullableTime** | The date and time the campaign was last modified. | [optional] 
**Deadline** | **time.Time** | The date and time the campaign is due. | 
**Type** | **map[string]interface{}** | The type of campaign. | 
**CampaignOwner** | [**TriggerInputCampaignActivatedCampaignCampaignOwner**](TriggerInputCampaignActivatedCampaignCampaignOwner.md) |  | 
**Status** | **map[string]interface{}** | The current status of the campaign. | 

## Methods

### NewTriggerInputCampaignEndedCampaign

`func NewTriggerInputCampaignEndedCampaign(id string, name string, description string, created time.Time, deadline time.Time, type_ map[string]interface{}, campaignOwner TriggerInputCampaignActivatedCampaignCampaignOwner, status map[string]interface{}, ) *TriggerInputCampaignEndedCampaign`

NewTriggerInputCampaignEndedCampaign instantiates a new TriggerInputCampaignEndedCampaign object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerInputCampaignEndedCampaignWithDefaults

`func NewTriggerInputCampaignEndedCampaignWithDefaults() *TriggerInputCampaignEndedCampaign`

NewTriggerInputCampaignEndedCampaignWithDefaults instantiates a new TriggerInputCampaignEndedCampaign object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TriggerInputCampaignEndedCampaign) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TriggerInputCampaignEndedCampaign) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TriggerInputCampaignEndedCampaign) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *TriggerInputCampaignEndedCampaign) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TriggerInputCampaignEndedCampaign) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TriggerInputCampaignEndedCampaign) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *TriggerInputCampaignEndedCampaign) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TriggerInputCampaignEndedCampaign) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TriggerInputCampaignEndedCampaign) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCreated

`func (o *TriggerInputCampaignEndedCampaign) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *TriggerInputCampaignEndedCampaign) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *TriggerInputCampaignEndedCampaign) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetModified

`func (o *TriggerInputCampaignEndedCampaign) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *TriggerInputCampaignEndedCampaign) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *TriggerInputCampaignEndedCampaign) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *TriggerInputCampaignEndedCampaign) HasModified() bool`

HasModified returns a boolean if a field has been set.

### SetModifiedNil

`func (o *TriggerInputCampaignEndedCampaign) SetModifiedNil(b bool)`

 SetModifiedNil sets the value for Modified to be an explicit nil

### UnsetModified
`func (o *TriggerInputCampaignEndedCampaign) UnsetModified()`

UnsetModified ensures that no value is present for Modified, not even an explicit nil
### GetDeadline

`func (o *TriggerInputCampaignEndedCampaign) GetDeadline() time.Time`

GetDeadline returns the Deadline field if non-nil, zero value otherwise.

### GetDeadlineOk

`func (o *TriggerInputCampaignEndedCampaign) GetDeadlineOk() (*time.Time, bool)`

GetDeadlineOk returns a tuple with the Deadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadline

`func (o *TriggerInputCampaignEndedCampaign) SetDeadline(v time.Time)`

SetDeadline sets Deadline field to given value.


### GetType

`func (o *TriggerInputCampaignEndedCampaign) GetType() map[string]interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TriggerInputCampaignEndedCampaign) GetTypeOk() (*map[string]interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TriggerInputCampaignEndedCampaign) SetType(v map[string]interface{})`

SetType sets Type field to given value.


### GetCampaignOwner

`func (o *TriggerInputCampaignEndedCampaign) GetCampaignOwner() TriggerInputCampaignActivatedCampaignCampaignOwner`

GetCampaignOwner returns the CampaignOwner field if non-nil, zero value otherwise.

### GetCampaignOwnerOk

`func (o *TriggerInputCampaignEndedCampaign) GetCampaignOwnerOk() (*TriggerInputCampaignActivatedCampaignCampaignOwner, bool)`

GetCampaignOwnerOk returns a tuple with the CampaignOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignOwner

`func (o *TriggerInputCampaignEndedCampaign) SetCampaignOwner(v TriggerInputCampaignActivatedCampaignCampaignOwner)`

SetCampaignOwner sets CampaignOwner field to given value.


### GetStatus

`func (o *TriggerInputCampaignEndedCampaign) GetStatus() map[string]interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TriggerInputCampaignEndedCampaign) GetStatusOk() (*map[string]interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TriggerInputCampaignEndedCampaign) SetStatus(v map[string]interface{})`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


