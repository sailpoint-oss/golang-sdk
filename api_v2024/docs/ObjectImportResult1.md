# ObjectImportResult1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Infos** | [**[]SpConfigMessage1**](SpConfigMessage1.md) | Informational messages returned from the target service on import. | 
**Warnings** | [**[]SpConfigMessage1**](SpConfigMessage1.md) | Warning messages returned from the target service on import. | 
**Errors** | [**[]SpConfigMessage1**](SpConfigMessage1.md) | Error messages returned from the target service on import. | 
**ImportedObjects** | [**[]ImportObject**](ImportObject.md) | References to objects that were created or updated by the import. | 

## Methods

### NewObjectImportResult1

`func NewObjectImportResult1(infos []SpConfigMessage1, warnings []SpConfigMessage1, errors []SpConfigMessage1, importedObjects []ImportObject, ) *ObjectImportResult1`

NewObjectImportResult1 instantiates a new ObjectImportResult1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectImportResult1WithDefaults

`func NewObjectImportResult1WithDefaults() *ObjectImportResult1`

NewObjectImportResult1WithDefaults instantiates a new ObjectImportResult1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInfos

`func (o *ObjectImportResult1) GetInfos() []SpConfigMessage1`

GetInfos returns the Infos field if non-nil, zero value otherwise.

### GetInfosOk

`func (o *ObjectImportResult1) GetInfosOk() (*[]SpConfigMessage1, bool)`

GetInfosOk returns a tuple with the Infos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfos

`func (o *ObjectImportResult1) SetInfos(v []SpConfigMessage1)`

SetInfos sets Infos field to given value.


### GetWarnings

`func (o *ObjectImportResult1) GetWarnings() []SpConfigMessage1`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *ObjectImportResult1) GetWarningsOk() (*[]SpConfigMessage1, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *ObjectImportResult1) SetWarnings(v []SpConfigMessage1)`

SetWarnings sets Warnings field to given value.


### GetErrors

`func (o *ObjectImportResult1) GetErrors() []SpConfigMessage1`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ObjectImportResult1) GetErrorsOk() (*[]SpConfigMessage1, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ObjectImportResult1) SetErrors(v []SpConfigMessage1)`

SetErrors sets Errors field to given value.


### GetImportedObjects

`func (o *ObjectImportResult1) GetImportedObjects() []ImportObject`

GetImportedObjects returns the ImportedObjects field if non-nil, zero value otherwise.

### GetImportedObjectsOk

`func (o *ObjectImportResult1) GetImportedObjectsOk() (*[]ImportObject, bool)`

GetImportedObjectsOk returns a tuple with the ImportedObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportedObjects

`func (o *ObjectImportResult1) SetImportedObjects(v []ImportObject)`

SetImportedObjects sets ImportedObjects field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


