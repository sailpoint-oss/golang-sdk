# SpConfigImportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | **string** | Name of JSON file containing the objects to be imported. | 
**Options** | Pointer to [**ImportOptions**](ImportOptions.md) |  | [optional] 

## Methods

### NewSpConfigImportRequest

`func NewSpConfigImportRequest(data string, ) *SpConfigImportRequest`

NewSpConfigImportRequest instantiates a new SpConfigImportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpConfigImportRequestWithDefaults

`func NewSpConfigImportRequestWithDefaults() *SpConfigImportRequest`

NewSpConfigImportRequestWithDefaults instantiates a new SpConfigImportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SpConfigImportRequest) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SpConfigImportRequest) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SpConfigImportRequest) SetData(v string)`

SetData sets Data field to given value.


### GetOptions

`func (o *SpConfigImportRequest) GetOptions() ImportOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *SpConfigImportRequest) GetOptionsOk() (*ImportOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *SpConfigImportRequest) SetOptions(v ImportOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *SpConfigImportRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


