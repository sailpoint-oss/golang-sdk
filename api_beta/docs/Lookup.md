# Lookup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Table** | **map[string]interface{}** | This is a JSON object of key-value pairs. The key is the string that will attempt to be matched to the input, and the value is the output string that should be returned if the key is matched   &gt;**Note** the use of the optional default key value here; if none of the three countries in the above example match the input string, the transform will return \&quot;Unknown Region\&quot; for the attribute that is mapped to this transform.  | 
**RequiresPeriodicRefresh** | Pointer to **bool** | A value that indicates whether the transform logic should be re-evaluated every evening as part of the identity refresh process | [optional] [default to false]
**Input** | Pointer to **map[string]interface{}** | This is an optional attribute that can explicitly define the input data which will be fed into the transform logic. If input is not provided, the transform will take its input from the source and attribute combination configured via the UI. | [optional] 

## Methods

### NewLookup

`func NewLookup(table map[string]interface{}, ) *Lookup`

NewLookup instantiates a new Lookup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLookupWithDefaults

`func NewLookupWithDefaults() *Lookup`

NewLookupWithDefaults instantiates a new Lookup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTable

`func (o *Lookup) GetTable() map[string]interface{}`

GetTable returns the Table field if non-nil, zero value otherwise.

### GetTableOk

`func (o *Lookup) GetTableOk() (*map[string]interface{}, bool)`

GetTableOk returns a tuple with the Table field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTable

`func (o *Lookup) SetTable(v map[string]interface{})`

SetTable sets Table field to given value.


### GetRequiresPeriodicRefresh

`func (o *Lookup) GetRequiresPeriodicRefresh() bool`

GetRequiresPeriodicRefresh returns the RequiresPeriodicRefresh field if non-nil, zero value otherwise.

### GetRequiresPeriodicRefreshOk

`func (o *Lookup) GetRequiresPeriodicRefreshOk() (*bool, bool)`

GetRequiresPeriodicRefreshOk returns a tuple with the RequiresPeriodicRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresPeriodicRefresh

`func (o *Lookup) SetRequiresPeriodicRefresh(v bool)`

SetRequiresPeriodicRefresh sets RequiresPeriodicRefresh field to given value.

### HasRequiresPeriodicRefresh

`func (o *Lookup) HasRequiresPeriodicRefresh() bool`

HasRequiresPeriodicRefresh returns a boolean if a field has been set.

### GetInput

`func (o *Lookup) GetInput() map[string]interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *Lookup) GetInputOk() (*map[string]interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *Lookup) SetInput(v map[string]interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *Lookup) HasInput() bool`

HasInput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


