# RefreshIdentitiesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to **string** | Defines the identity or identities which this refresh applies to. The filter must use searchable identity attributes. If the filter cannot be understood or parsed, all identities will be refreshed.  | [optional] 
**RefreshArgs** | Pointer to [**RefreshIdentitiesRequestRefreshArgs**](RefreshIdentitiesRequestRefreshArgs.md) |  | [optional] 

## Methods

### NewRefreshIdentitiesRequest

`func NewRefreshIdentitiesRequest() *RefreshIdentitiesRequest`

NewRefreshIdentitiesRequest instantiates a new RefreshIdentitiesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefreshIdentitiesRequestWithDefaults

`func NewRefreshIdentitiesRequestWithDefaults() *RefreshIdentitiesRequest`

NewRefreshIdentitiesRequestWithDefaults instantiates a new RefreshIdentitiesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *RefreshIdentitiesRequest) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *RefreshIdentitiesRequest) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *RefreshIdentitiesRequest) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *RefreshIdentitiesRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetRefreshArgs

`func (o *RefreshIdentitiesRequest) GetRefreshArgs() RefreshIdentitiesRequestRefreshArgs`

GetRefreshArgs returns the RefreshArgs field if non-nil, zero value otherwise.

### GetRefreshArgsOk

`func (o *RefreshIdentitiesRequest) GetRefreshArgsOk() (*RefreshIdentitiesRequestRefreshArgs, bool)`

GetRefreshArgsOk returns a tuple with the RefreshArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshArgs

`func (o *RefreshIdentitiesRequest) SetRefreshArgs(v RefreshIdentitiesRequestRefreshArgs)`

SetRefreshArgs sets RefreshArgs field to given value.

### HasRefreshArgs

`func (o *RefreshIdentitiesRequest) HasRefreshArgs() bool`

HasRefreshArgs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


