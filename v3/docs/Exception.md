# Exception

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Id of a SOD exception. | [optional] 
**Created** | Pointer to **NullableTime** | The time when this SOD exception is created. | [optional] 
**Modified** | Pointer to **NullableTime** | The time when this SOD exception is modified. | [optional] 
**SodPolicy** | Pointer to [**BaseReferenceDto**](BaseReferenceDto.md) |  | [optional] 
**Identity** | Pointer to [**BaseReferenceDto**](BaseReferenceDto.md) |  | [optional] 
**Start** | Pointer to **time.Time** | The earliest date-time when this SOD exception is applicable. | [optional] 
**End** | Pointer to **time.Time** | The last date-time when this SOD exception is applicable. | [optional] 
**BusinessJustification** | Pointer to **string** | The business justification for the exception. | [optional] 
**MitigatingControl** | Pointer to **string** | The mitigating control for the exception. | [optional] 
**AccessCriteria** | Pointer to [**ExceptionAccessCriteria**](ExceptionAccessCriteria.md) |  | [optional] 
**Origin** | Pointer to [**BaseReferenceDto**](BaseReferenceDto.md) |  | [optional] 

## Methods

### NewException

`func NewException() *Exception`

NewException instantiates a new Exception object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExceptionWithDefaults

`func NewExceptionWithDefaults() *Exception`

NewExceptionWithDefaults instantiates a new Exception object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Exception) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Exception) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Exception) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Exception) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *Exception) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Exception) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetCreated

`func (o *Exception) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Exception) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Exception) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Exception) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreatedNil

`func (o *Exception) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *Exception) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetModified

`func (o *Exception) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *Exception) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *Exception) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *Exception) HasModified() bool`

HasModified returns a boolean if a field has been set.

### SetModifiedNil

`func (o *Exception) SetModifiedNil(b bool)`

 SetModifiedNil sets the value for Modified to be an explicit nil

### UnsetModified
`func (o *Exception) UnsetModified()`

UnsetModified ensures that no value is present for Modified, not even an explicit nil
### GetSodPolicy

`func (o *Exception) GetSodPolicy() BaseReferenceDto`

GetSodPolicy returns the SodPolicy field if non-nil, zero value otherwise.

### GetSodPolicyOk

`func (o *Exception) GetSodPolicyOk() (*BaseReferenceDto, bool)`

GetSodPolicyOk returns a tuple with the SodPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSodPolicy

`func (o *Exception) SetSodPolicy(v BaseReferenceDto)`

SetSodPolicy sets SodPolicy field to given value.

### HasSodPolicy

`func (o *Exception) HasSodPolicy() bool`

HasSodPolicy returns a boolean if a field has been set.

### GetIdentity

`func (o *Exception) GetIdentity() BaseReferenceDto`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *Exception) GetIdentityOk() (*BaseReferenceDto, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *Exception) SetIdentity(v BaseReferenceDto)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *Exception) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetStart

`func (o *Exception) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *Exception) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *Exception) SetStart(v time.Time)`

SetStart sets Start field to given value.

### HasStart

`func (o *Exception) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetEnd

`func (o *Exception) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *Exception) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *Exception) SetEnd(v time.Time)`

SetEnd sets End field to given value.

### HasEnd

`func (o *Exception) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetBusinessJustification

`func (o *Exception) GetBusinessJustification() string`

GetBusinessJustification returns the BusinessJustification field if non-nil, zero value otherwise.

### GetBusinessJustificationOk

`func (o *Exception) GetBusinessJustificationOk() (*string, bool)`

GetBusinessJustificationOk returns a tuple with the BusinessJustification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessJustification

`func (o *Exception) SetBusinessJustification(v string)`

SetBusinessJustification sets BusinessJustification field to given value.

### HasBusinessJustification

`func (o *Exception) HasBusinessJustification() bool`

HasBusinessJustification returns a boolean if a field has been set.

### GetMitigatingControl

`func (o *Exception) GetMitigatingControl() string`

GetMitigatingControl returns the MitigatingControl field if non-nil, zero value otherwise.

### GetMitigatingControlOk

`func (o *Exception) GetMitigatingControlOk() (*string, bool)`

GetMitigatingControlOk returns a tuple with the MitigatingControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMitigatingControl

`func (o *Exception) SetMitigatingControl(v string)`

SetMitigatingControl sets MitigatingControl field to given value.

### HasMitigatingControl

`func (o *Exception) HasMitigatingControl() bool`

HasMitigatingControl returns a boolean if a field has been set.

### GetAccessCriteria

`func (o *Exception) GetAccessCriteria() ExceptionAccessCriteria`

GetAccessCriteria returns the AccessCriteria field if non-nil, zero value otherwise.

### GetAccessCriteriaOk

`func (o *Exception) GetAccessCriteriaOk() (*ExceptionAccessCriteria, bool)`

GetAccessCriteriaOk returns a tuple with the AccessCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessCriteria

`func (o *Exception) SetAccessCriteria(v ExceptionAccessCriteria)`

SetAccessCriteria sets AccessCriteria field to given value.

### HasAccessCriteria

`func (o *Exception) HasAccessCriteria() bool`

HasAccessCriteria returns a boolean if a field has been set.

### GetOrigin

`func (o *Exception) GetOrigin() BaseReferenceDto`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *Exception) GetOriginOk() (*BaseReferenceDto, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *Exception) SetOrigin(v BaseReferenceDto)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *Exception) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


