# GetReferenceIdentityAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | This must always be set to \&quot;Cloud Services Deployment Utility\&quot; | 
**Operation** | **string** | The operation to perform &#x60;getReferenceIdentityAttribute&#x60; | 
**Uid** | **string** | This is the SailPoint User Name (uid) value of the identity whose attribute is desired  As a convenience feature, you can use the &#x60;manager&#x60; keyword to dynamically look up the user&#39;s manager and then get that manager&#39;s identity attribute.  | 
**RequiresPeriodicRefresh** | Pointer to **bool** | A value that indicates whether the transform logic should be re-evaluated every evening as part of the identity refresh process | [optional] 

## Methods

### NewGetReferenceIdentityAttribute

`func NewGetReferenceIdentityAttribute(name string, operation string, uid string, ) *GetReferenceIdentityAttribute`

NewGetReferenceIdentityAttribute instantiates a new GetReferenceIdentityAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetReferenceIdentityAttributeWithDefaults

`func NewGetReferenceIdentityAttributeWithDefaults() *GetReferenceIdentityAttribute`

NewGetReferenceIdentityAttributeWithDefaults instantiates a new GetReferenceIdentityAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GetReferenceIdentityAttribute) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetReferenceIdentityAttribute) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetReferenceIdentityAttribute) SetName(v string)`

SetName sets Name field to given value.


### GetOperation

`func (o *GetReferenceIdentityAttribute) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *GetReferenceIdentityAttribute) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *GetReferenceIdentityAttribute) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetUid

`func (o *GetReferenceIdentityAttribute) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *GetReferenceIdentityAttribute) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *GetReferenceIdentityAttribute) SetUid(v string)`

SetUid sets Uid field to given value.


### GetRequiresPeriodicRefresh

`func (o *GetReferenceIdentityAttribute) GetRequiresPeriodicRefresh() bool`

GetRequiresPeriodicRefresh returns the RequiresPeriodicRefresh field if non-nil, zero value otherwise.

### GetRequiresPeriodicRefreshOk

`func (o *GetReferenceIdentityAttribute) GetRequiresPeriodicRefreshOk() (*bool, bool)`

GetRequiresPeriodicRefreshOk returns a tuple with the RequiresPeriodicRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresPeriodicRefresh

`func (o *GetReferenceIdentityAttribute) SetRequiresPeriodicRefresh(v bool)`

SetRequiresPeriodicRefresh sets RequiresPeriodicRefresh field to given value.

### HasRequiresPeriodicRefresh

`func (o *GetReferenceIdentityAttribute) HasRequiresPeriodicRefresh() bool`

HasRequiresPeriodicRefresh returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


