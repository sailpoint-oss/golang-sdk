# WorkgroupDeleteItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Id of the Governance Group. | 
**Status** | **string** |  The HTTP response status code returned for an individual Governance Group that is requested for deletion during a bulk delete operation.  &gt; 204   -  Governance Group deleted successfully.  &gt; 409   - Governance Group is in use,hence can not be deleted.  &gt; 404   - Governance Group not found.  | 
**Description** | Pointer to **string** | Human readable status description and containing additional context information about success or failures etc.  | [optional] 

## Methods

### NewWorkgroupDeleteItem

`func NewWorkgroupDeleteItem(id string, status string, ) *WorkgroupDeleteItem`

NewWorkgroupDeleteItem instantiates a new WorkgroupDeleteItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkgroupDeleteItemWithDefaults

`func NewWorkgroupDeleteItemWithDefaults() *WorkgroupDeleteItem`

NewWorkgroupDeleteItemWithDefaults instantiates a new WorkgroupDeleteItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkgroupDeleteItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkgroupDeleteItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkgroupDeleteItem) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *WorkgroupDeleteItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkgroupDeleteItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkgroupDeleteItem) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetDescription

`func (o *WorkgroupDeleteItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkgroupDeleteItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkgroupDeleteItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkgroupDeleteItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


