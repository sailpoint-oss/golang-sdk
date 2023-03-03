# SectionDetailsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** | Label of the section | [optional] 
**FormItems** | Pointer to **[]map[string]interface{}** | List of FormItems. FormItems can be SectionDetails and/or FieldDetails | [optional] 

## Methods

### NewSectionDetailsAllOf

`func NewSectionDetailsAllOf() *SectionDetailsAllOf`

NewSectionDetailsAllOf instantiates a new SectionDetailsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSectionDetailsAllOfWithDefaults

`func NewSectionDetailsAllOfWithDefaults() *SectionDetailsAllOf`

NewSectionDetailsAllOfWithDefaults instantiates a new SectionDetailsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *SectionDetailsAllOf) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *SectionDetailsAllOf) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *SectionDetailsAllOf) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *SectionDetailsAllOf) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetFormItems

`func (o *SectionDetailsAllOf) GetFormItems() []map[string]interface{}`

GetFormItems returns the FormItems field if non-nil, zero value otherwise.

### GetFormItemsOk

`func (o *SectionDetailsAllOf) GetFormItemsOk() (*[]map[string]interface{}, bool)`

GetFormItemsOk returns a tuple with the FormItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormItems

`func (o *SectionDetailsAllOf) SetFormItems(v []map[string]interface{})`

SetFormItems sets FormItems field to given value.

### HasFormItems

`func (o *SectionDetailsAllOf) HasFormItems() bool`

HasFormItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


