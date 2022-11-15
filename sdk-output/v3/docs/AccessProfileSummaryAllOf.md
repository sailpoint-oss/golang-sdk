# AccessProfileSummaryAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | Pointer to [**Reference1**](Reference1.md) |  | [optional] 
**Owner** | Pointer to [**DisplayReference**](DisplayReference.md) |  | [optional] 
**Revocable** | Pointer to **bool** |  | [optional] 

## Methods

### NewAccessProfileSummaryAllOf

`func NewAccessProfileSummaryAllOf() *AccessProfileSummaryAllOf`

NewAccessProfileSummaryAllOf instantiates a new AccessProfileSummaryAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessProfileSummaryAllOfWithDefaults

`func NewAccessProfileSummaryAllOfWithDefaults() *AccessProfileSummaryAllOf`

NewAccessProfileSummaryAllOfWithDefaults instantiates a new AccessProfileSummaryAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *AccessProfileSummaryAllOf) GetSource() Reference1`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AccessProfileSummaryAllOf) GetSourceOk() (*Reference1, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AccessProfileSummaryAllOf) SetSource(v Reference1)`

SetSource sets Source field to given value.

### HasSource

`func (o *AccessProfileSummaryAllOf) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetOwner

`func (o *AccessProfileSummaryAllOf) GetOwner() DisplayReference`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *AccessProfileSummaryAllOf) GetOwnerOk() (*DisplayReference, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *AccessProfileSummaryAllOf) SetOwner(v DisplayReference)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *AccessProfileSummaryAllOf) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetRevocable

`func (o *AccessProfileSummaryAllOf) GetRevocable() bool`

GetRevocable returns the Revocable field if non-nil, zero value otherwise.

### GetRevocableOk

`func (o *AccessProfileSummaryAllOf) GetRevocableOk() (*bool, bool)`

GetRevocableOk returns a tuple with the Revocable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocable

`func (o *AccessProfileSummaryAllOf) SetRevocable(v bool)`

SetRevocable sets Revocable field to given value.

### HasRevocable

`func (o *AccessProfileSummaryAllOf) HasRevocable() bool`

HasRevocable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


