# IdentityProfileExportedObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **int32** | Version or object from the target service. | [optional] 
**Self** | Pointer to [**SelfImportExportDto**](SelfImportExportDto.md) |  | [optional] 
**Object** | Pointer to [**IdentityProfile1**](IdentityProfile1.md) |  | [optional] 

## Methods

### NewIdentityProfileExportedObject

`func NewIdentityProfileExportedObject() *IdentityProfileExportedObject`

NewIdentityProfileExportedObject instantiates a new IdentityProfileExportedObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityProfileExportedObjectWithDefaults

`func NewIdentityProfileExportedObjectWithDefaults() *IdentityProfileExportedObject`

NewIdentityProfileExportedObjectWithDefaults instantiates a new IdentityProfileExportedObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *IdentityProfileExportedObject) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *IdentityProfileExportedObject) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *IdentityProfileExportedObject) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *IdentityProfileExportedObject) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetSelf

`func (o *IdentityProfileExportedObject) GetSelf() SelfImportExportDto`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *IdentityProfileExportedObject) GetSelfOk() (*SelfImportExportDto, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *IdentityProfileExportedObject) SetSelf(v SelfImportExportDto)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *IdentityProfileExportedObject) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetObject

`func (o *IdentityProfileExportedObject) GetObject() IdentityProfile1`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *IdentityProfileExportedObject) GetObjectOk() (*IdentityProfile1, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *IdentityProfileExportedObject) SetObject(v IdentityProfile1)`

SetObject sets Object field to given value.

### HasObject

`func (o *IdentityProfileExportedObject) HasObject() bool`

HasObject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


