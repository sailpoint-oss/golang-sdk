# Launcher

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Icon** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**UsageCertText** | Pointer to **string** |  | [optional] 
**UsageCertAttest** | Pointer to **time.Time** |  | [optional] 
**Account** | Pointer to [**LauncherAccount**](LauncherAccount.md) |  | [optional] 
**Links** | Pointer to [**[]LauncherLink**](LauncherLink.md) |  | [optional] 

## Methods

### NewLauncher

`func NewLauncher() *Launcher`

NewLauncher instantiates a new Launcher object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLauncherWithDefaults

`func NewLauncherWithDefaults() *Launcher`

NewLauncherWithDefaults instantiates a new Launcher object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Launcher) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Launcher) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Launcher) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Launcher) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Launcher) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Launcher) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Launcher) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Launcher) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Launcher) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Launcher) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Launcher) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Launcher) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDateCreated

`func (o *Launcher) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *Launcher) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *Launcher) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *Launcher) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *Launcher) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Launcher) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Launcher) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *Launcher) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetIcon

`func (o *Launcher) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *Launcher) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *Launcher) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *Launcher) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetState

`func (o *Launcher) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Launcher) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Launcher) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Launcher) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUsageCertText

`func (o *Launcher) GetUsageCertText() string`

GetUsageCertText returns the UsageCertText field if non-nil, zero value otherwise.

### GetUsageCertTextOk

`func (o *Launcher) GetUsageCertTextOk() (*string, bool)`

GetUsageCertTextOk returns a tuple with the UsageCertText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageCertText

`func (o *Launcher) SetUsageCertText(v string)`

SetUsageCertText sets UsageCertText field to given value.

### HasUsageCertText

`func (o *Launcher) HasUsageCertText() bool`

HasUsageCertText returns a boolean if a field has been set.

### GetUsageCertAttest

`func (o *Launcher) GetUsageCertAttest() time.Time`

GetUsageCertAttest returns the UsageCertAttest field if non-nil, zero value otherwise.

### GetUsageCertAttestOk

`func (o *Launcher) GetUsageCertAttestOk() (*time.Time, bool)`

GetUsageCertAttestOk returns a tuple with the UsageCertAttest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageCertAttest

`func (o *Launcher) SetUsageCertAttest(v time.Time)`

SetUsageCertAttest sets UsageCertAttest field to given value.

### HasUsageCertAttest

`func (o *Launcher) HasUsageCertAttest() bool`

HasUsageCertAttest returns a boolean if a field has been set.

### GetAccount

`func (o *Launcher) GetAccount() LauncherAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *Launcher) GetAccountOk() (*LauncherAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *Launcher) SetAccount(v LauncherAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *Launcher) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetLinks

`func (o *Launcher) GetLinks() []LauncherLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Launcher) GetLinksOk() (*[]LauncherLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Launcher) SetLinks(v []LauncherLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Launcher) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


