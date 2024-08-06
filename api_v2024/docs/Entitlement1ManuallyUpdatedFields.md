# Entitlement1ManuallyUpdatedFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DISPLAY_NAME** | Pointer to **bool** | True if the entitlements name was updated manually via entitlement import csv or patch endpoint.  False means that property value has not been change after first entitlement aggregation. Field refers to [Entitlement response schema](https://developer.sailpoint.com/idn/api/beta/get-entitlement) &gt; &#x60;name&#x60; property. | [optional] [default to false]
**DESCRIPTION** | Pointer to **bool** | True if the entitlement description was updated manually via entitlement import csv or patch endpoint.  False means that property value has not been change after first entitlement aggregation. Field refers to [Entitlement response schema](https://developer.sailpoint.com/idn/api/beta/get-entitlement) &gt; &#x60;description&#x60; property. | [optional] [default to false]

## Methods

### NewEntitlement1ManuallyUpdatedFields

`func NewEntitlement1ManuallyUpdatedFields() *Entitlement1ManuallyUpdatedFields`

NewEntitlement1ManuallyUpdatedFields instantiates a new Entitlement1ManuallyUpdatedFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlement1ManuallyUpdatedFieldsWithDefaults

`func NewEntitlement1ManuallyUpdatedFieldsWithDefaults() *Entitlement1ManuallyUpdatedFields`

NewEntitlement1ManuallyUpdatedFieldsWithDefaults instantiates a new Entitlement1ManuallyUpdatedFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDISPLAY_NAME

`func (o *Entitlement1ManuallyUpdatedFields) GetDISPLAY_NAME() bool`

GetDISPLAY_NAME returns the DISPLAY_NAME field if non-nil, zero value otherwise.

### GetDISPLAY_NAMEOk

`func (o *Entitlement1ManuallyUpdatedFields) GetDISPLAY_NAMEOk() (*bool, bool)`

GetDISPLAY_NAMEOk returns a tuple with the DISPLAY_NAME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDISPLAY_NAME

`func (o *Entitlement1ManuallyUpdatedFields) SetDISPLAY_NAME(v bool)`

SetDISPLAY_NAME sets DISPLAY_NAME field to given value.

### HasDISPLAY_NAME

`func (o *Entitlement1ManuallyUpdatedFields) HasDISPLAY_NAME() bool`

HasDISPLAY_NAME returns a boolean if a field has been set.

### GetDESCRIPTION

`func (o *Entitlement1ManuallyUpdatedFields) GetDESCRIPTION() bool`

GetDESCRIPTION returns the DESCRIPTION field if non-nil, zero value otherwise.

### GetDESCRIPTIONOk

`func (o *Entitlement1ManuallyUpdatedFields) GetDESCRIPTIONOk() (*bool, bool)`

GetDESCRIPTIONOk returns a tuple with the DESCRIPTION field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDESCRIPTION

`func (o *Entitlement1ManuallyUpdatedFields) SetDESCRIPTION(v bool)`

SetDESCRIPTION sets DESCRIPTION field to given value.

### HasDESCRIPTION

`func (o *Entitlement1ManuallyUpdatedFields) HasDESCRIPTION() bool`

HasDESCRIPTION returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


