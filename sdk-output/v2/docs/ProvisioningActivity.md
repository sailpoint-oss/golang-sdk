# ProvisioningActivity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Operation** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**SourceId** | Pointer to **string** |  | [optional] 
**SourceName** | Pointer to **string** |  | [optional] 
**AccountName** | Pointer to **string** |  | [optional] 
**OwnerName** | Pointer to **string** |  | [optional] 
**ApproverName** | Pointer to **string** |  | [optional] 
**Warnings** | Pointer to **[]string** |  | [optional] 
**Errors** | Pointer to **[]string** |  | [optional] 
**ProvisioningPlan** | Pointer to **string** |  | [optional] 

## Methods

### NewProvisioningActivity

`func NewProvisioningActivity() *ProvisioningActivity`

NewProvisioningActivity instantiates a new ProvisioningActivity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisioningActivityWithDefaults

`func NewProvisioningActivityWithDefaults() *ProvisioningActivity`

NewProvisioningActivityWithDefaults instantiates a new ProvisioningActivity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProvisioningActivity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProvisioningActivity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProvisioningActivity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProvisioningActivity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ProvisioningActivity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProvisioningActivity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProvisioningActivity) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProvisioningActivity) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDateCreated

`func (o *ProvisioningActivity) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ProvisioningActivity) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ProvisioningActivity) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ProvisioningActivity) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ProvisioningActivity) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ProvisioningActivity) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ProvisioningActivity) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ProvisioningActivity) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetOperation

`func (o *ProvisioningActivity) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *ProvisioningActivity) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *ProvisioningActivity) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *ProvisioningActivity) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetStatus

`func (o *ProvisioningActivity) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProvisioningActivity) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProvisioningActivity) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProvisioningActivity) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSourceId

`func (o *ProvisioningActivity) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *ProvisioningActivity) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *ProvisioningActivity) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *ProvisioningActivity) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetSourceName

`func (o *ProvisioningActivity) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *ProvisioningActivity) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *ProvisioningActivity) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *ProvisioningActivity) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### GetAccountName

`func (o *ProvisioningActivity) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *ProvisioningActivity) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *ProvisioningActivity) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *ProvisioningActivity) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetOwnerName

`func (o *ProvisioningActivity) GetOwnerName() string`

GetOwnerName returns the OwnerName field if non-nil, zero value otherwise.

### GetOwnerNameOk

`func (o *ProvisioningActivity) GetOwnerNameOk() (*string, bool)`

GetOwnerNameOk returns a tuple with the OwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerName

`func (o *ProvisioningActivity) SetOwnerName(v string)`

SetOwnerName sets OwnerName field to given value.

### HasOwnerName

`func (o *ProvisioningActivity) HasOwnerName() bool`

HasOwnerName returns a boolean if a field has been set.

### GetApproverName

`func (o *ProvisioningActivity) GetApproverName() string`

GetApproverName returns the ApproverName field if non-nil, zero value otherwise.

### GetApproverNameOk

`func (o *ProvisioningActivity) GetApproverNameOk() (*string, bool)`

GetApproverNameOk returns a tuple with the ApproverName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverName

`func (o *ProvisioningActivity) SetApproverName(v string)`

SetApproverName sets ApproverName field to given value.

### HasApproverName

`func (o *ProvisioningActivity) HasApproverName() bool`

HasApproverName returns a boolean if a field has been set.

### GetWarnings

`func (o *ProvisioningActivity) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *ProvisioningActivity) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *ProvisioningActivity) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *ProvisioningActivity) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### GetErrors

`func (o *ProvisioningActivity) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ProvisioningActivity) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ProvisioningActivity) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ProvisioningActivity) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetProvisioningPlan

`func (o *ProvisioningActivity) GetProvisioningPlan() string`

GetProvisioningPlan returns the ProvisioningPlan field if non-nil, zero value otherwise.

### GetProvisioningPlanOk

`func (o *ProvisioningActivity) GetProvisioningPlanOk() (*string, bool)`

GetProvisioningPlanOk returns a tuple with the ProvisioningPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningPlan

`func (o *ProvisioningActivity) SetProvisioningPlan(v string)`

SetProvisioningPlan sets ProvisioningPlan field to given value.

### HasProvisioningPlan

`func (o *ProvisioningActivity) HasProvisioningPlan() bool`

HasProvisioningPlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


