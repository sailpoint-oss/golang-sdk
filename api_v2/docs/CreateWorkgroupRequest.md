# CreateWorkgroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**CreateWorkgroupRequestOwner**](CreateWorkgroupRequestOwner.md) |  | [optional] 

## Methods

### NewCreateWorkgroupRequest

`func NewCreateWorkgroupRequest() *CreateWorkgroupRequest`

NewCreateWorkgroupRequest instantiates a new CreateWorkgroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWorkgroupRequestWithDefaults

`func NewCreateWorkgroupRequestWithDefaults() *CreateWorkgroupRequest`

NewCreateWorkgroupRequestWithDefaults instantiates a new CreateWorkgroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateWorkgroupRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateWorkgroupRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateWorkgroupRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateWorkgroupRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *CreateWorkgroupRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateWorkgroupRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateWorkgroupRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateWorkgroupRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwner

`func (o *CreateWorkgroupRequest) GetOwner() CreateWorkgroupRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CreateWorkgroupRequest) GetOwnerOk() (*CreateWorkgroupRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CreateWorkgroupRequest) SetOwner(v CreateWorkgroupRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *CreateWorkgroupRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


