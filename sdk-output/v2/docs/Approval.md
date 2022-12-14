# Approval

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**AccessProfileName** | Pointer to **string** |  | [optional] 
**AccessProfileDescription** | Pointer to **string** |  | [optional] 
**AppName** | Pointer to **string** |  | [optional] 
**AppId** | Pointer to **string** |  | [optional] 
**AccessRequestId** | Pointer to **string** |  | [optional] 
**OwnerId** | Pointer to **string** |  | [optional] 
**OwnerName** | Pointer to **string** |  | [optional] 
**RequesterId** | Pointer to **string** |  | [optional] 
**RequesterName** | Pointer to **string** |  | [optional] 
**RequesteeId** | Pointer to **string** |  | [optional] 
**RequesteeName** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Completed** | Pointer to **bool** |  | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] 
**Modified** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewApproval

`func NewApproval() *Approval`

NewApproval instantiates a new Approval object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalWithDefaults

`func NewApprovalWithDefaults() *Approval`

NewApprovalWithDefaults instantiates a new Approval object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Approval) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Approval) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Approval) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Approval) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *Approval) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Approval) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Approval) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Approval) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAccessProfileName

`func (o *Approval) GetAccessProfileName() string`

GetAccessProfileName returns the AccessProfileName field if non-nil, zero value otherwise.

### GetAccessProfileNameOk

`func (o *Approval) GetAccessProfileNameOk() (*string, bool)`

GetAccessProfileNameOk returns a tuple with the AccessProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessProfileName

`func (o *Approval) SetAccessProfileName(v string)`

SetAccessProfileName sets AccessProfileName field to given value.

### HasAccessProfileName

`func (o *Approval) HasAccessProfileName() bool`

HasAccessProfileName returns a boolean if a field has been set.

### GetAccessProfileDescription

`func (o *Approval) GetAccessProfileDescription() string`

GetAccessProfileDescription returns the AccessProfileDescription field if non-nil, zero value otherwise.

### GetAccessProfileDescriptionOk

`func (o *Approval) GetAccessProfileDescriptionOk() (*string, bool)`

GetAccessProfileDescriptionOk returns a tuple with the AccessProfileDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessProfileDescription

`func (o *Approval) SetAccessProfileDescription(v string)`

SetAccessProfileDescription sets AccessProfileDescription field to given value.

### HasAccessProfileDescription

`func (o *Approval) HasAccessProfileDescription() bool`

HasAccessProfileDescription returns a boolean if a field has been set.

### GetAppName

`func (o *Approval) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *Approval) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *Approval) SetAppName(v string)`

SetAppName sets AppName field to given value.

### HasAppName

`func (o *Approval) HasAppName() bool`

HasAppName returns a boolean if a field has been set.

### GetAppId

`func (o *Approval) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *Approval) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *Approval) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *Approval) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetAccessRequestId

`func (o *Approval) GetAccessRequestId() string`

GetAccessRequestId returns the AccessRequestId field if non-nil, zero value otherwise.

### GetAccessRequestIdOk

`func (o *Approval) GetAccessRequestIdOk() (*string, bool)`

GetAccessRequestIdOk returns a tuple with the AccessRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessRequestId

`func (o *Approval) SetAccessRequestId(v string)`

SetAccessRequestId sets AccessRequestId field to given value.

### HasAccessRequestId

`func (o *Approval) HasAccessRequestId() bool`

HasAccessRequestId returns a boolean if a field has been set.

### GetOwnerId

`func (o *Approval) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *Approval) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *Approval) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *Approval) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetOwnerName

`func (o *Approval) GetOwnerName() string`

GetOwnerName returns the OwnerName field if non-nil, zero value otherwise.

### GetOwnerNameOk

`func (o *Approval) GetOwnerNameOk() (*string, bool)`

GetOwnerNameOk returns a tuple with the OwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerName

`func (o *Approval) SetOwnerName(v string)`

SetOwnerName sets OwnerName field to given value.

### HasOwnerName

`func (o *Approval) HasOwnerName() bool`

HasOwnerName returns a boolean if a field has been set.

### GetRequesterId

`func (o *Approval) GetRequesterId() string`

GetRequesterId returns the RequesterId field if non-nil, zero value otherwise.

### GetRequesterIdOk

`func (o *Approval) GetRequesterIdOk() (*string, bool)`

GetRequesterIdOk returns a tuple with the RequesterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterId

`func (o *Approval) SetRequesterId(v string)`

SetRequesterId sets RequesterId field to given value.

### HasRequesterId

`func (o *Approval) HasRequesterId() bool`

HasRequesterId returns a boolean if a field has been set.

### GetRequesterName

`func (o *Approval) GetRequesterName() string`

GetRequesterName returns the RequesterName field if non-nil, zero value otherwise.

### GetRequesterNameOk

`func (o *Approval) GetRequesterNameOk() (*string, bool)`

GetRequesterNameOk returns a tuple with the RequesterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterName

`func (o *Approval) SetRequesterName(v string)`

SetRequesterName sets RequesterName field to given value.

### HasRequesterName

`func (o *Approval) HasRequesterName() bool`

HasRequesterName returns a boolean if a field has been set.

### GetRequesteeId

`func (o *Approval) GetRequesteeId() string`

GetRequesteeId returns the RequesteeId field if non-nil, zero value otherwise.

### GetRequesteeIdOk

`func (o *Approval) GetRequesteeIdOk() (*string, bool)`

GetRequesteeIdOk returns a tuple with the RequesteeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesteeId

`func (o *Approval) SetRequesteeId(v string)`

SetRequesteeId sets RequesteeId field to given value.

### HasRequesteeId

`func (o *Approval) HasRequesteeId() bool`

HasRequesteeId returns a boolean if a field has been set.

### GetRequesteeName

`func (o *Approval) GetRequesteeName() string`

GetRequesteeName returns the RequesteeName field if non-nil, zero value otherwise.

### GetRequesteeNameOk

`func (o *Approval) GetRequesteeNameOk() (*string, bool)`

GetRequesteeNameOk returns a tuple with the RequesteeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesteeName

`func (o *Approval) SetRequesteeName(v string)`

SetRequesteeName sets RequesteeName field to given value.

### HasRequesteeName

`func (o *Approval) HasRequesteeName() bool`

HasRequesteeName returns a boolean if a field has been set.

### GetState

`func (o *Approval) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Approval) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Approval) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Approval) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCompleted

`func (o *Approval) GetCompleted() bool`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *Approval) GetCompletedOk() (*bool, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *Approval) SetCompleted(v bool)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *Approval) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.

### GetCreated

`func (o *Approval) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Approval) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Approval) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Approval) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetModified

`func (o *Approval) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *Approval) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *Approval) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *Approval) HasModified() bool`

HasModified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


