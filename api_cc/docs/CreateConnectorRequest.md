# CreateConnectorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ClassName** | Pointer to **string** |  | [optional] 
**DirectConnect** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateConnectorRequest

`func NewCreateConnectorRequest() *CreateConnectorRequest`

NewCreateConnectorRequest instantiates a new CreateConnectorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateConnectorRequestWithDefaults

`func NewCreateConnectorRequestWithDefaults() *CreateConnectorRequest`

NewCreateConnectorRequestWithDefaults instantiates a new CreateConnectorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateConnectorRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateConnectorRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateConnectorRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateConnectorRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *CreateConnectorRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateConnectorRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateConnectorRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateConnectorRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetClassName

`func (o *CreateConnectorRequest) GetClassName() string`

GetClassName returns the ClassName field if non-nil, zero value otherwise.

### GetClassNameOk

`func (o *CreateConnectorRequest) GetClassNameOk() (*string, bool)`

GetClassNameOk returns a tuple with the ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassName

`func (o *CreateConnectorRequest) SetClassName(v string)`

SetClassName sets ClassName field to given value.

### HasClassName

`func (o *CreateConnectorRequest) HasClassName() bool`

HasClassName returns a boolean if a field has been set.

### GetDirectConnect

`func (o *CreateConnectorRequest) GetDirectConnect() bool`

GetDirectConnect returns the DirectConnect field if non-nil, zero value otherwise.

### GetDirectConnectOk

`func (o *CreateConnectorRequest) GetDirectConnectOk() (*bool, bool)`

GetDirectConnectOk returns a tuple with the DirectConnect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectConnect

`func (o *CreateConnectorRequest) SetDirectConnect(v bool)`

SetDirectConnect sets DirectConnect field to given value.

### HasDirectConnect

`func (o *CreateConnectorRequest) HasDirectConnect() bool`

HasDirectConnect returns a boolean if a field has been set.

### GetStatus

`func (o *CreateConnectorRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateConnectorRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateConnectorRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateConnectorRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


