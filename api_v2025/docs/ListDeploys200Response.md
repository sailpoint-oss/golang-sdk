# ListDeploys200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]DeployResponse**](DeployResponse.md) | list of deployments | [optional] 

## Methods

### NewListDeploys200Response

`func NewListDeploys200Response() *ListDeploys200Response`

NewListDeploys200Response instantiates a new ListDeploys200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDeploys200ResponseWithDefaults

`func NewListDeploys200ResponseWithDefaults() *ListDeploys200Response`

NewListDeploys200ResponseWithDefaults instantiates a new ListDeploys200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ListDeploys200Response) GetItems() []DeployResponse`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListDeploys200Response) GetItemsOk() (*[]DeployResponse, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListDeploys200Response) SetItems(v []DeployResponse)`

SetItems sets Items field to given value.

### HasItems

`func (o *ListDeploys200Response) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


