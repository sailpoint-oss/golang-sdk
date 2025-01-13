# DateCompareFirstDate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceName** | **string** | A reference to the source to search for the account | 
**AttributeName** | **string** | The name of the attribute on the account to return. This should match the name of the account attribute name visible in the user interface, or on the source schema. | 
**AccountSortAttribute** | Pointer to **string** | The value of this configuration is a string name of the attribute to use when determining the ordering of returned accounts when there are multiple entries | [optional] [default to "created"]
**AccountSortDescending** | Pointer to **bool** | The value of this configuration is a boolean (true/false). Controls the order of the sort when there are multiple accounts. If not defined, the transform will default to false (ascending order) | [optional] [default to false]
**AccountReturnFirstLink** | Pointer to **bool** | The value of this configuration is a boolean (true/false). Controls which account to source a value from for an attribute.  If this flag is set to true, the transform returns the value from the first account in the list, even if it is null. If it is set to false, the transform returns the first non-null value. If not defined, the transform will default to false | [optional] [default to false]
**AccountFilter** | Pointer to **string** | This expression queries the database to narrow search results. The value of this configuration is a sailpoint.object.Filter expression and used when searching against the database.  The default filter will always include the source and identity, and any subsequent expressions will be combined in an AND operation to the existing search criteria. Only certain searchable attributes are available:  - &#x60;nativeIdentity&#x60; - the Account ID  - &#x60;displayName&#x60; - the Account Name  - &#x60;entitlements&#x60; - a boolean value to determine if the account has entitlements | [optional] 
**AccountPropertyFilter** | Pointer to **string** | This expression is used to search and filter accounts in memory. The value of this configuration is a sailpoint.object.Filter expression and used when searching against the returned resultset.  All account attributes are available for filtering as this operation is performed in memory. | [optional] 
**RequiresPeriodicRefresh** | Pointer to **bool** | A value that indicates whether the transform logic should be re-evaluated every evening as part of the identity refresh process | [optional] [default to false]
**Input** | Pointer to **map[string]interface{}** | This is an optional attribute that can explicitly define the input data which will be fed into the transform logic. If input is not provided, the transform will take its input from the source and attribute combination configured via the UI. | [optional] 
**InputFormat** | Pointer to [**DateFormatInputFormat**](DateFormatInputFormat.md) |  | [optional] 
**OutputFormat** | Pointer to [**DateFormatOutputFormat**](DateFormatOutputFormat.md) |  | [optional] 

## Methods

### NewDateCompareFirstDate

`func NewDateCompareFirstDate(sourceName string, attributeName string, ) *DateCompareFirstDate`

NewDateCompareFirstDate instantiates a new DateCompareFirstDate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDateCompareFirstDateWithDefaults

`func NewDateCompareFirstDateWithDefaults() *DateCompareFirstDate`

NewDateCompareFirstDateWithDefaults instantiates a new DateCompareFirstDate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceName

`func (o *DateCompareFirstDate) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *DateCompareFirstDate) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *DateCompareFirstDate) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.


### GetAttributeName

`func (o *DateCompareFirstDate) GetAttributeName() string`

GetAttributeName returns the AttributeName field if non-nil, zero value otherwise.

### GetAttributeNameOk

`func (o *DateCompareFirstDate) GetAttributeNameOk() (*string, bool)`

GetAttributeNameOk returns a tuple with the AttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeName

`func (o *DateCompareFirstDate) SetAttributeName(v string)`

SetAttributeName sets AttributeName field to given value.


### GetAccountSortAttribute

`func (o *DateCompareFirstDate) GetAccountSortAttribute() string`

GetAccountSortAttribute returns the AccountSortAttribute field if non-nil, zero value otherwise.

### GetAccountSortAttributeOk

`func (o *DateCompareFirstDate) GetAccountSortAttributeOk() (*string, bool)`

GetAccountSortAttributeOk returns a tuple with the AccountSortAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSortAttribute

`func (o *DateCompareFirstDate) SetAccountSortAttribute(v string)`

SetAccountSortAttribute sets AccountSortAttribute field to given value.

### HasAccountSortAttribute

`func (o *DateCompareFirstDate) HasAccountSortAttribute() bool`

HasAccountSortAttribute returns a boolean if a field has been set.

### GetAccountSortDescending

`func (o *DateCompareFirstDate) GetAccountSortDescending() bool`

GetAccountSortDescending returns the AccountSortDescending field if non-nil, zero value otherwise.

### GetAccountSortDescendingOk

`func (o *DateCompareFirstDate) GetAccountSortDescendingOk() (*bool, bool)`

GetAccountSortDescendingOk returns a tuple with the AccountSortDescending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSortDescending

`func (o *DateCompareFirstDate) SetAccountSortDescending(v bool)`

SetAccountSortDescending sets AccountSortDescending field to given value.

### HasAccountSortDescending

`func (o *DateCompareFirstDate) HasAccountSortDescending() bool`

HasAccountSortDescending returns a boolean if a field has been set.

### GetAccountReturnFirstLink

`func (o *DateCompareFirstDate) GetAccountReturnFirstLink() bool`

GetAccountReturnFirstLink returns the AccountReturnFirstLink field if non-nil, zero value otherwise.

### GetAccountReturnFirstLinkOk

`func (o *DateCompareFirstDate) GetAccountReturnFirstLinkOk() (*bool, bool)`

GetAccountReturnFirstLinkOk returns a tuple with the AccountReturnFirstLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountReturnFirstLink

`func (o *DateCompareFirstDate) SetAccountReturnFirstLink(v bool)`

SetAccountReturnFirstLink sets AccountReturnFirstLink field to given value.

### HasAccountReturnFirstLink

`func (o *DateCompareFirstDate) HasAccountReturnFirstLink() bool`

HasAccountReturnFirstLink returns a boolean if a field has been set.

### GetAccountFilter

`func (o *DateCompareFirstDate) GetAccountFilter() string`

GetAccountFilter returns the AccountFilter field if non-nil, zero value otherwise.

### GetAccountFilterOk

`func (o *DateCompareFirstDate) GetAccountFilterOk() (*string, bool)`

GetAccountFilterOk returns a tuple with the AccountFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountFilter

`func (o *DateCompareFirstDate) SetAccountFilter(v string)`

SetAccountFilter sets AccountFilter field to given value.

### HasAccountFilter

`func (o *DateCompareFirstDate) HasAccountFilter() bool`

HasAccountFilter returns a boolean if a field has been set.

### GetAccountPropertyFilter

`func (o *DateCompareFirstDate) GetAccountPropertyFilter() string`

GetAccountPropertyFilter returns the AccountPropertyFilter field if non-nil, zero value otherwise.

### GetAccountPropertyFilterOk

`func (o *DateCompareFirstDate) GetAccountPropertyFilterOk() (*string, bool)`

GetAccountPropertyFilterOk returns a tuple with the AccountPropertyFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountPropertyFilter

`func (o *DateCompareFirstDate) SetAccountPropertyFilter(v string)`

SetAccountPropertyFilter sets AccountPropertyFilter field to given value.

### HasAccountPropertyFilter

`func (o *DateCompareFirstDate) HasAccountPropertyFilter() bool`

HasAccountPropertyFilter returns a boolean if a field has been set.

### GetRequiresPeriodicRefresh

`func (o *DateCompareFirstDate) GetRequiresPeriodicRefresh() bool`

GetRequiresPeriodicRefresh returns the RequiresPeriodicRefresh field if non-nil, zero value otherwise.

### GetRequiresPeriodicRefreshOk

`func (o *DateCompareFirstDate) GetRequiresPeriodicRefreshOk() (*bool, bool)`

GetRequiresPeriodicRefreshOk returns a tuple with the RequiresPeriodicRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresPeriodicRefresh

`func (o *DateCompareFirstDate) SetRequiresPeriodicRefresh(v bool)`

SetRequiresPeriodicRefresh sets RequiresPeriodicRefresh field to given value.

### HasRequiresPeriodicRefresh

`func (o *DateCompareFirstDate) HasRequiresPeriodicRefresh() bool`

HasRequiresPeriodicRefresh returns a boolean if a field has been set.

### GetInput

`func (o *DateCompareFirstDate) GetInput() map[string]interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *DateCompareFirstDate) GetInputOk() (*map[string]interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *DateCompareFirstDate) SetInput(v map[string]interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *DateCompareFirstDate) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetInputFormat

`func (o *DateCompareFirstDate) GetInputFormat() DateFormatInputFormat`

GetInputFormat returns the InputFormat field if non-nil, zero value otherwise.

### GetInputFormatOk

`func (o *DateCompareFirstDate) GetInputFormatOk() (*DateFormatInputFormat, bool)`

GetInputFormatOk returns a tuple with the InputFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputFormat

`func (o *DateCompareFirstDate) SetInputFormat(v DateFormatInputFormat)`

SetInputFormat sets InputFormat field to given value.

### HasInputFormat

`func (o *DateCompareFirstDate) HasInputFormat() bool`

HasInputFormat returns a boolean if a field has been set.

### GetOutputFormat

`func (o *DateCompareFirstDate) GetOutputFormat() DateFormatOutputFormat`

GetOutputFormat returns the OutputFormat field if non-nil, zero value otherwise.

### GetOutputFormatOk

`func (o *DateCompareFirstDate) GetOutputFormatOk() (*DateFormatOutputFormat, bool)`

GetOutputFormatOk returns a tuple with the OutputFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFormat

`func (o *DateCompareFirstDate) SetOutputFormat(v DateFormatOutputFormat)`

SetOutputFormat sets OutputFormat field to given value.

### HasOutputFormat

`func (o *DateCompareFirstDate) HasOutputFormat() bool`

HasOutputFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


