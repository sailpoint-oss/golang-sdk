# ImportUploadedBackupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | ***os.File** | JSON file containing the objects to be imported. | 
**Name** | **string** | Name that will be assigned to the uploaded file. | 

## Methods

### NewImportUploadedBackupRequest

`func NewImportUploadedBackupRequest(data *os.File, name string, ) *ImportUploadedBackupRequest`

NewImportUploadedBackupRequest instantiates a new ImportUploadedBackupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportUploadedBackupRequestWithDefaults

`func NewImportUploadedBackupRequestWithDefaults() *ImportUploadedBackupRequest`

NewImportUploadedBackupRequestWithDefaults instantiates a new ImportUploadedBackupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ImportUploadedBackupRequest) GetData() *os.File`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ImportUploadedBackupRequest) GetDataOk() (**os.File, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ImportUploadedBackupRequest) SetData(v *os.File)`

SetData sets Data field to given value.


### GetName

`func (o *ImportUploadedBackupRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ImportUploadedBackupRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ImportUploadedBackupRequest) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


