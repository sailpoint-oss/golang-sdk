# ListCampaignFilters200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]CampaignFilterDetails**](CampaignFilterDetails.md) | List of campaign filters. | [optional] 
**Count** | Pointer to **int32** | Number of filters returned. | [optional] 

## Methods

### NewListCampaignFilters200Response

`func NewListCampaignFilters200Response() *ListCampaignFilters200Response`

NewListCampaignFilters200Response instantiates a new ListCampaignFilters200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCampaignFilters200ResponseWithDefaults

`func NewListCampaignFilters200ResponseWithDefaults() *ListCampaignFilters200Response`

NewListCampaignFilters200ResponseWithDefaults instantiates a new ListCampaignFilters200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ListCampaignFilters200Response) GetItems() []CampaignFilterDetails`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListCampaignFilters200Response) GetItemsOk() (*[]CampaignFilterDetails, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListCampaignFilters200Response) SetItems(v []CampaignFilterDetails)`

SetItems sets Items field to given value.

### HasItems

`func (o *ListCampaignFilters200Response) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetCount

`func (o *ListCampaignFilters200Response) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListCampaignFilters200Response) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListCampaignFilters200Response) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ListCampaignFilters200Response) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


