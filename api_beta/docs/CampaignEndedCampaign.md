# CampaignEndedCampaign

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Campaign&#39;s unique ID for the campaign. | 
**Name** | **string** | Campaign&#39;s unique ID. | 
**Description** | **string** | Campaign&#39;s extended description. | 
**Created** | **time.Time** | Date and time when the campaign was created. | 
**Modified** | Pointer to **NullableTime** | Date and time when the campaign was last modified. | [optional] 
**Deadline** | **time.Time** | Date and time when the campaign is due. | 
**Type** | **map[string]interface{}** | Campaign&#39;s type. | 
**CampaignOwner** | [**CampaignActivatedCampaignCampaignOwner**](CampaignActivatedCampaignCampaignOwner.md) |  | 
**Status** | **map[string]interface{}** | Campaign&#39;s current status. | 

## Methods

### NewCampaignEndedCampaign

`func NewCampaignEndedCampaign(id string, name string, description string, created time.Time, deadline time.Time, type_ map[string]interface{}, campaignOwner CampaignActivatedCampaignCampaignOwner, status map[string]interface{}, ) *CampaignEndedCampaign`

NewCampaignEndedCampaign instantiates a new CampaignEndedCampaign object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignEndedCampaignWithDefaults

`func NewCampaignEndedCampaignWithDefaults() *CampaignEndedCampaign`

NewCampaignEndedCampaignWithDefaults instantiates a new CampaignEndedCampaign object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CampaignEndedCampaign) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CampaignEndedCampaign) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CampaignEndedCampaign) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *CampaignEndedCampaign) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CampaignEndedCampaign) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CampaignEndedCampaign) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CampaignEndedCampaign) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CampaignEndedCampaign) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CampaignEndedCampaign) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCreated

`func (o *CampaignEndedCampaign) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CampaignEndedCampaign) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CampaignEndedCampaign) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetModified

`func (o *CampaignEndedCampaign) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *CampaignEndedCampaign) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *CampaignEndedCampaign) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *CampaignEndedCampaign) HasModified() bool`

HasModified returns a boolean if a field has been set.

### SetModifiedNil

`func (o *CampaignEndedCampaign) SetModifiedNil(b bool)`

 SetModifiedNil sets the value for Modified to be an explicit nil

### UnsetModified
`func (o *CampaignEndedCampaign) UnsetModified()`

UnsetModified ensures that no value is present for Modified, not even an explicit nil
### GetDeadline

`func (o *CampaignEndedCampaign) GetDeadline() time.Time`

GetDeadline returns the Deadline field if non-nil, zero value otherwise.

### GetDeadlineOk

`func (o *CampaignEndedCampaign) GetDeadlineOk() (*time.Time, bool)`

GetDeadlineOk returns a tuple with the Deadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadline

`func (o *CampaignEndedCampaign) SetDeadline(v time.Time)`

SetDeadline sets Deadline field to given value.


### GetType

`func (o *CampaignEndedCampaign) GetType() map[string]interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CampaignEndedCampaign) GetTypeOk() (*map[string]interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CampaignEndedCampaign) SetType(v map[string]interface{})`

SetType sets Type field to given value.


### GetCampaignOwner

`func (o *CampaignEndedCampaign) GetCampaignOwner() CampaignActivatedCampaignCampaignOwner`

GetCampaignOwner returns the CampaignOwner field if non-nil, zero value otherwise.

### GetCampaignOwnerOk

`func (o *CampaignEndedCampaign) GetCampaignOwnerOk() (*CampaignActivatedCampaignCampaignOwner, bool)`

GetCampaignOwnerOk returns a tuple with the CampaignOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignOwner

`func (o *CampaignEndedCampaign) SetCampaignOwner(v CampaignActivatedCampaignCampaignOwner)`

SetCampaignOwner sets CampaignOwner field to given value.


### GetStatus

`func (o *CampaignEndedCampaign) GetStatus() map[string]interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CampaignEndedCampaign) GetStatusOk() (*map[string]interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CampaignEndedCampaign) SetStatus(v map[string]interface{})`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


