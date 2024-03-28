# ImportEntitlementsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CsvFile** | ***os.File** |  | 

## Methods

### NewImportEntitlementsRequest

`func NewImportEntitlementsRequest(csvFile *os.File, ) *ImportEntitlementsRequest`

NewImportEntitlementsRequest instantiates a new ImportEntitlementsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportEntitlementsRequestWithDefaults

`func NewImportEntitlementsRequestWithDefaults() *ImportEntitlementsRequest`

NewImportEntitlementsRequestWithDefaults instantiates a new ImportEntitlementsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCsvFile

`func (o *ImportEntitlementsRequest) GetCsvFile() *os.File`

GetCsvFile returns the CsvFile field if non-nil, zero value otherwise.

### GetCsvFileOk

`func (o *ImportEntitlementsRequest) GetCsvFileOk() (**os.File, bool)`

GetCsvFileOk returns a tuple with the CsvFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsvFile

`func (o *ImportEntitlementsRequest) SetCsvFile(v *os.File)`

SetCsvFile sets CsvFile field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


