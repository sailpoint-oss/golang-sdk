# LoadAccountsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisableOptimization** | Pointer to **bool** |  | [optional] 
**File** | Pointer to ***os.File** |  | [optional] 

## Methods

### NewLoadAccountsRequest

`func NewLoadAccountsRequest() *LoadAccountsRequest`

NewLoadAccountsRequest instantiates a new LoadAccountsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadAccountsRequestWithDefaults

`func NewLoadAccountsRequestWithDefaults() *LoadAccountsRequest`

NewLoadAccountsRequestWithDefaults instantiates a new LoadAccountsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisableOptimization

`func (o *LoadAccountsRequest) GetDisableOptimization() bool`

GetDisableOptimization returns the DisableOptimization field if non-nil, zero value otherwise.

### GetDisableOptimizationOk

`func (o *LoadAccountsRequest) GetDisableOptimizationOk() (*bool, bool)`

GetDisableOptimizationOk returns a tuple with the DisableOptimization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableOptimization

`func (o *LoadAccountsRequest) SetDisableOptimization(v bool)`

SetDisableOptimization sets DisableOptimization field to given value.

### HasDisableOptimization

`func (o *LoadAccountsRequest) HasDisableOptimization() bool`

HasDisableOptimization returns a boolean if a field has been set.

### GetFile

`func (o *LoadAccountsRequest) GetFile() *os.File`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *LoadAccountsRequest) GetFileOk() (**os.File, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *LoadAccountsRequest) SetFile(v *os.File)`

SetFile sets File field to given value.

### HasFile

`func (o *LoadAccountsRequest) HasFile() bool`

HasFile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


