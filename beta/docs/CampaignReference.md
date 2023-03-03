# CampaignReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique ID of the campaign. | 
**Name** | **string** | The name of the campaign. | 
**Type** | **string** | The type of object that is being referenced. | 
**CampaignType** | **string** | The type of the campaign. | 
**Description** | **NullableString** | The description of the campaign set by the admin who created it. | 

## Methods

### NewCampaignReference

`func NewCampaignReference(id string, name string, type_ string, campaignType string, description NullableString, ) *CampaignReference`

NewCampaignReference instantiates a new CampaignReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignReferenceWithDefaults

`func NewCampaignReferenceWithDefaults() *CampaignReference`

NewCampaignReferenceWithDefaults instantiates a new CampaignReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CampaignReference) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CampaignReference) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CampaignReference) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *CampaignReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CampaignReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CampaignReference) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *CampaignReference) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CampaignReference) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CampaignReference) SetType(v string)`

SetType sets Type field to given value.


### GetCampaignType

`func (o *CampaignReference) GetCampaignType() string`

GetCampaignType returns the CampaignType field if non-nil, zero value otherwise.

### GetCampaignTypeOk

`func (o *CampaignReference) GetCampaignTypeOk() (*string, bool)`

GetCampaignTypeOk returns a tuple with the CampaignType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignType

`func (o *CampaignReference) SetCampaignType(v string)`

SetCampaignType sets CampaignType field to given value.


### GetDescription

`func (o *CampaignReference) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CampaignReference) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CampaignReference) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *CampaignReference) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CampaignReference) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


