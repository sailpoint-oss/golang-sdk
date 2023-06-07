# AccessProfileUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessProfileId** | Pointer to **string** | ID of the Access Profile that is in use | [optional] 
**UsedBy** | Pointer to [**[]BaseReferenceDto1**](BaseReferenceDto1.md) | List of references to objects which are using the indicated Access Profile | [optional] 

## Methods

### NewAccessProfileUsage

`func NewAccessProfileUsage() *AccessProfileUsage`

NewAccessProfileUsage instantiates a new AccessProfileUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessProfileUsageWithDefaults

`func NewAccessProfileUsageWithDefaults() *AccessProfileUsage`

NewAccessProfileUsageWithDefaults instantiates a new AccessProfileUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessProfileId

`func (o *AccessProfileUsage) GetAccessProfileId() string`

GetAccessProfileId returns the AccessProfileId field if non-nil, zero value otherwise.

### GetAccessProfileIdOk

`func (o *AccessProfileUsage) GetAccessProfileIdOk() (*string, bool)`

GetAccessProfileIdOk returns a tuple with the AccessProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessProfileId

`func (o *AccessProfileUsage) SetAccessProfileId(v string)`

SetAccessProfileId sets AccessProfileId field to given value.

### HasAccessProfileId

`func (o *AccessProfileUsage) HasAccessProfileId() bool`

HasAccessProfileId returns a boolean if a field has been set.

### GetUsedBy

`func (o *AccessProfileUsage) GetUsedBy() []BaseReferenceDto1`

GetUsedBy returns the UsedBy field if non-nil, zero value otherwise.

### GetUsedByOk

`func (o *AccessProfileUsage) GetUsedByOk() (*[]BaseReferenceDto1, bool)`

GetUsedByOk returns a tuple with the UsedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedBy

`func (o *AccessProfileUsage) SetUsedBy(v []BaseReferenceDto1)`

SetUsedBy sets UsedBy field to given value.

### HasUsedBy

`func (o *AccessProfileUsage) HasUsedBy() bool`

HasUsedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


