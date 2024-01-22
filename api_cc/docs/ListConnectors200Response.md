# ListConnectors200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **float32** |  | [optional] 
**Items** | Pointer to [**[]ListConnectors200ResponseItemsInner**](ListConnectors200ResponseItemsInner.md) |  | [optional] 

## Methods

### NewListConnectors200Response

`func NewListConnectors200Response() *ListConnectors200Response`

NewListConnectors200Response instantiates a new ListConnectors200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListConnectors200ResponseWithDefaults

`func NewListConnectors200ResponseWithDefaults() *ListConnectors200Response`

NewListConnectors200ResponseWithDefaults instantiates a new ListConnectors200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *ListConnectors200Response) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ListConnectors200Response) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ListConnectors200Response) SetTotal(v float32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ListConnectors200Response) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetItems

`func (o *ListConnectors200Response) GetItems() []ListConnectors200ResponseItemsInner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListConnectors200Response) GetItemsOk() (*[]ListConnectors200ResponseItemsInner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListConnectors200Response) SetItems(v []ListConnectors200ResponseItemsInner)`

SetItems sets Items field to given value.

### HasItems

`func (o *ListConnectors200Response) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


