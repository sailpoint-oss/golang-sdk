# SearchEventAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 
**Reviewer** | Pointer to [**SearchEventAttributesReviewer**](SearchEventAttributesReviewer.md) |  | [optional] 
**Requester** | Pointer to [**SearchEventAttributesReviewer**](SearchEventAttributesReviewer.md) |  | [optional] 
**OldValue** | Pointer to **string** |  | [optional] 
**NewValue** | Pointer to **string** |  | [optional] 
**AccountName** | Pointer to **string** |  | [optional] 
**SourceName** | Pointer to **string** |  | [optional] 
**SSO** | Pointer to [**SearchEventAttributesSSO**](SearchEventAttributesSSO.md) |  | [optional] 

## Methods

### NewSearchEventAttributes

`func NewSearchEventAttributes() *SearchEventAttributes`

NewSearchEventAttributes instantiates a new SearchEventAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchEventAttributesWithDefaults

`func NewSearchEventAttributesWithDefaults() *SearchEventAttributes`

NewSearchEventAttributesWithDefaults instantiates a new SearchEventAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *SearchEventAttributes) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SearchEventAttributes) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SearchEventAttributes) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *SearchEventAttributes) HasError() bool`

HasError returns a boolean if a field has been set.

### GetReviewer

`func (o *SearchEventAttributes) GetReviewer() SearchEventAttributesReviewer`

GetReviewer returns the Reviewer field if non-nil, zero value otherwise.

### GetReviewerOk

`func (o *SearchEventAttributes) GetReviewerOk() (*SearchEventAttributesReviewer, bool)`

GetReviewerOk returns a tuple with the Reviewer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewer

`func (o *SearchEventAttributes) SetReviewer(v SearchEventAttributesReviewer)`

SetReviewer sets Reviewer field to given value.

### HasReviewer

`func (o *SearchEventAttributes) HasReviewer() bool`

HasReviewer returns a boolean if a field has been set.

### GetRequester

`func (o *SearchEventAttributes) GetRequester() SearchEventAttributesReviewer`

GetRequester returns the Requester field if non-nil, zero value otherwise.

### GetRequesterOk

`func (o *SearchEventAttributes) GetRequesterOk() (*SearchEventAttributesReviewer, bool)`

GetRequesterOk returns a tuple with the Requester field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequester

`func (o *SearchEventAttributes) SetRequester(v SearchEventAttributesReviewer)`

SetRequester sets Requester field to given value.

### HasRequester

`func (o *SearchEventAttributes) HasRequester() bool`

HasRequester returns a boolean if a field has been set.

### GetOldValue

`func (o *SearchEventAttributes) GetOldValue() string`

GetOldValue returns the OldValue field if non-nil, zero value otherwise.

### GetOldValueOk

`func (o *SearchEventAttributes) GetOldValueOk() (*string, bool)`

GetOldValueOk returns a tuple with the OldValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldValue

`func (o *SearchEventAttributes) SetOldValue(v string)`

SetOldValue sets OldValue field to given value.

### HasOldValue

`func (o *SearchEventAttributes) HasOldValue() bool`

HasOldValue returns a boolean if a field has been set.

### GetNewValue

`func (o *SearchEventAttributes) GetNewValue() string`

GetNewValue returns the NewValue field if non-nil, zero value otherwise.

### GetNewValueOk

`func (o *SearchEventAttributes) GetNewValueOk() (*string, bool)`

GetNewValueOk returns a tuple with the NewValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewValue

`func (o *SearchEventAttributes) SetNewValue(v string)`

SetNewValue sets NewValue field to given value.

### HasNewValue

`func (o *SearchEventAttributes) HasNewValue() bool`

HasNewValue returns a boolean if a field has been set.

### GetAccountName

`func (o *SearchEventAttributes) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *SearchEventAttributes) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *SearchEventAttributes) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *SearchEventAttributes) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetSourceName

`func (o *SearchEventAttributes) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *SearchEventAttributes) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *SearchEventAttributes) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *SearchEventAttributes) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### GetSSO

`func (o *SearchEventAttributes) GetSSO() SearchEventAttributesSSO`

GetSSO returns the SSO field if non-nil, zero value otherwise.

### GetSSOOk

`func (o *SearchEventAttributes) GetSSOOk() (*SearchEventAttributesSSO, bool)`

GetSSOOk returns a tuple with the SSO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSSO

`func (o *SearchEventAttributes) SetSSO(v SearchEventAttributesSSO)`

SetSSO sets SSO field to given value.

### HasSSO

`func (o *SearchEventAttributes) HasSSO() bool`

HasSSO returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


