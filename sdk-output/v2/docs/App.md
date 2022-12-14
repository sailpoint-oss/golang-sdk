# App

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Icon** | Pointer to **string** |  | [optional] 
**AccessMethods** | Pointer to [**[]AppAccessMethodsInner**](AppAccessMethodsInner.md) |  | [optional] 

## Methods

### NewApp

`func NewApp() *App`

NewApp instantiates a new App object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppWithDefaults

`func NewAppWithDefaults() *App`

NewAppWithDefaults instantiates a new App object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *App) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *App) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *App) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *App) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *App) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *App) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *App) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *App) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *App) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *App) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *App) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *App) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDateCreated

`func (o *App) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *App) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *App) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *App) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *App) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *App) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *App) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *App) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetIcon

`func (o *App) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *App) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *App) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *App) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetAccessMethods

`func (o *App) GetAccessMethods() []AppAccessMethodsInner`

GetAccessMethods returns the AccessMethods field if non-nil, zero value otherwise.

### GetAccessMethodsOk

`func (o *App) GetAccessMethodsOk() (*[]AppAccessMethodsInner, bool)`

GetAccessMethodsOk returns a tuple with the AccessMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessMethods

`func (o *App) SetAccessMethods(v []AppAccessMethodsInner)`

SetAccessMethods sets AccessMethods field to given value.

### HasAccessMethods

`func (o *App) HasAccessMethods() bool`

HasAccessMethods returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


