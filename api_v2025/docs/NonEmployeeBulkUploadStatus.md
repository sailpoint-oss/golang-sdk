# NonEmployeeBulkUploadStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Returns the following values indicating the progress or result of the bulk upload job. \&quot;PENDING\&quot; means the job is queued and waiting to be processed. \&quot;IN_PROGRESS\&quot; means the job is currently being processed. \&quot;COMPLETED\&quot; means the job has been completed without any errors. \&quot;ERROR\&quot; means the job failed to process with errors. null means job has been submitted to the source.  | [optional] 

## Methods

### NewNonEmployeeBulkUploadStatus

`func NewNonEmployeeBulkUploadStatus() *NonEmployeeBulkUploadStatus`

NewNonEmployeeBulkUploadStatus instantiates a new NonEmployeeBulkUploadStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNonEmployeeBulkUploadStatusWithDefaults

`func NewNonEmployeeBulkUploadStatusWithDefaults() *NonEmployeeBulkUploadStatus`

NewNonEmployeeBulkUploadStatusWithDefaults instantiates a new NonEmployeeBulkUploadStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *NonEmployeeBulkUploadStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NonEmployeeBulkUploadStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NonEmployeeBulkUploadStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NonEmployeeBulkUploadStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


