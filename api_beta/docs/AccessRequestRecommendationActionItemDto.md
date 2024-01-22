# AccessRequestRecommendationActionItemDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdentityId** | **string** | The identity ID taking the action. | 
**Access** | [**AccessRequestRecommendationItem**](AccessRequestRecommendationItem.md) |  | 

## Methods

### NewAccessRequestRecommendationActionItemDto

`func NewAccessRequestRecommendationActionItemDto(identityId string, access AccessRequestRecommendationItem, ) *AccessRequestRecommendationActionItemDto`

NewAccessRequestRecommendationActionItemDto instantiates a new AccessRequestRecommendationActionItemDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessRequestRecommendationActionItemDtoWithDefaults

`func NewAccessRequestRecommendationActionItemDtoWithDefaults() *AccessRequestRecommendationActionItemDto`

NewAccessRequestRecommendationActionItemDtoWithDefaults instantiates a new AccessRequestRecommendationActionItemDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentityId

`func (o *AccessRequestRecommendationActionItemDto) GetIdentityId() string`

GetIdentityId returns the IdentityId field if non-nil, zero value otherwise.

### GetIdentityIdOk

`func (o *AccessRequestRecommendationActionItemDto) GetIdentityIdOk() (*string, bool)`

GetIdentityIdOk returns a tuple with the IdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityId

`func (o *AccessRequestRecommendationActionItemDto) SetIdentityId(v string)`

SetIdentityId sets IdentityId field to given value.


### GetAccess

`func (o *AccessRequestRecommendationActionItemDto) GetAccess() AccessRequestRecommendationItem`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *AccessRequestRecommendationActionItemDto) GetAccessOk() (*AccessRequestRecommendationItem, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *AccessRequestRecommendationActionItemDto) SetAccess(v AccessRequestRecommendationItem)`

SetAccess sets Access field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


