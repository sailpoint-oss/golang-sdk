# TaskResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ParentName** | Pointer to **string** |  | [optional] 
**Launcher** | Pointer to **string** |  | [optional] 
**Completed** | Pointer to **time.Time** |  | [optional] 
**Launched** | Pointer to **time.Time** |  | [optional] 
**CompletionStatus** | Pointer to **string** |  | [optional] 
**Messages** | Pointer to [**[]TaskResultMessagesInner**](TaskResultMessagesInner.md) |  | [optional] 
**Returns** | Pointer to [**[]TaskResultReturnsInner**](TaskResultReturnsInner.md) |  | [optional] 
**Attributes** | Pointer to [**DynamicSchemaEto**](DynamicSchemaEto.md) |  | [optional] 
**Progress** | Pointer to **string** |  | [optional] 

## Methods

### NewTaskResult

`func NewTaskResult() *TaskResult`

NewTaskResult instantiates a new TaskResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskResultWithDefaults

`func NewTaskResultWithDefaults() *TaskResult`

NewTaskResultWithDefaults instantiates a new TaskResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TaskResult) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TaskResult) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TaskResult) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TaskResult) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *TaskResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaskResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaskResult) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TaskResult) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *TaskResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TaskResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TaskResult) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TaskResult) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *TaskResult) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TaskResult) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TaskResult) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TaskResult) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetParentName

`func (o *TaskResult) GetParentName() string`

GetParentName returns the ParentName field if non-nil, zero value otherwise.

### GetParentNameOk

`func (o *TaskResult) GetParentNameOk() (*string, bool)`

GetParentNameOk returns a tuple with the ParentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentName

`func (o *TaskResult) SetParentName(v string)`

SetParentName sets ParentName field to given value.

### HasParentName

`func (o *TaskResult) HasParentName() bool`

HasParentName returns a boolean if a field has been set.

### GetLauncher

`func (o *TaskResult) GetLauncher() string`

GetLauncher returns the Launcher field if non-nil, zero value otherwise.

### GetLauncherOk

`func (o *TaskResult) GetLauncherOk() (*string, bool)`

GetLauncherOk returns a tuple with the Launcher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLauncher

`func (o *TaskResult) SetLauncher(v string)`

SetLauncher sets Launcher field to given value.

### HasLauncher

`func (o *TaskResult) HasLauncher() bool`

HasLauncher returns a boolean if a field has been set.

### GetCompleted

`func (o *TaskResult) GetCompleted() time.Time`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *TaskResult) GetCompletedOk() (*time.Time, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *TaskResult) SetCompleted(v time.Time)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *TaskResult) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.

### GetLaunched

`func (o *TaskResult) GetLaunched() time.Time`

GetLaunched returns the Launched field if non-nil, zero value otherwise.

### GetLaunchedOk

`func (o *TaskResult) GetLaunchedOk() (*time.Time, bool)`

GetLaunchedOk returns a tuple with the Launched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunched

`func (o *TaskResult) SetLaunched(v time.Time)`

SetLaunched sets Launched field to given value.

### HasLaunched

`func (o *TaskResult) HasLaunched() bool`

HasLaunched returns a boolean if a field has been set.

### GetCompletionStatus

`func (o *TaskResult) GetCompletionStatus() string`

GetCompletionStatus returns the CompletionStatus field if non-nil, zero value otherwise.

### GetCompletionStatusOk

`func (o *TaskResult) GetCompletionStatusOk() (*string, bool)`

GetCompletionStatusOk returns a tuple with the CompletionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionStatus

`func (o *TaskResult) SetCompletionStatus(v string)`

SetCompletionStatus sets CompletionStatus field to given value.

### HasCompletionStatus

`func (o *TaskResult) HasCompletionStatus() bool`

HasCompletionStatus returns a boolean if a field has been set.

### GetMessages

`func (o *TaskResult) GetMessages() []TaskResultMessagesInner`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *TaskResult) GetMessagesOk() (*[]TaskResultMessagesInner, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *TaskResult) SetMessages(v []TaskResultMessagesInner)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *TaskResult) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetReturns

`func (o *TaskResult) GetReturns() []TaskResultReturnsInner`

GetReturns returns the Returns field if non-nil, zero value otherwise.

### GetReturnsOk

`func (o *TaskResult) GetReturnsOk() (*[]TaskResultReturnsInner, bool)`

GetReturnsOk returns a tuple with the Returns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturns

`func (o *TaskResult) SetReturns(v []TaskResultReturnsInner)`

SetReturns sets Returns field to given value.

### HasReturns

`func (o *TaskResult) HasReturns() bool`

HasReturns returns a boolean if a field has been set.

### GetAttributes

`func (o *TaskResult) GetAttributes() DynamicSchemaEto`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TaskResult) GetAttributesOk() (*DynamicSchemaEto, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TaskResult) SetAttributes(v DynamicSchemaEto)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *TaskResult) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetProgress

`func (o *TaskResult) GetProgress() string`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *TaskResult) GetProgressOk() (*string, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *TaskResult) SetProgress(v string)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *TaskResult) HasProgress() bool`

HasProgress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


