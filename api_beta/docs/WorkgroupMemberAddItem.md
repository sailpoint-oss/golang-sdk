# WorkgroupMemberAddItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Identifier of identity in bulk member add request. | 
**Status** | **string** |  The HTTP response status code returned for an individual member that is requested for addition during a bulk add operation.   The HTTP response status code returned for an individual Governance Group is requested for deletion.   &gt; 201   - Identity is added into Governance Group members list.  &gt; 409   - Identity is already member of  Governance Group.  | 
**Description** | Pointer to **string** | Human readable status description and containing additional context information about success or failures etc.  | [optional] 

## Methods

### NewWorkgroupMemberAddItem

`func NewWorkgroupMemberAddItem(id string, status string, ) *WorkgroupMemberAddItem`

NewWorkgroupMemberAddItem instantiates a new WorkgroupMemberAddItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkgroupMemberAddItemWithDefaults

`func NewWorkgroupMemberAddItemWithDefaults() *WorkgroupMemberAddItem`

NewWorkgroupMemberAddItemWithDefaults instantiates a new WorkgroupMemberAddItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkgroupMemberAddItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkgroupMemberAddItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkgroupMemberAddItem) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *WorkgroupMemberAddItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkgroupMemberAddItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkgroupMemberAddItem) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetDescription

`func (o *WorkgroupMemberAddItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkgroupMemberAddItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkgroupMemberAddItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkgroupMemberAddItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


