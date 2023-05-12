# MailFromAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The identity or domain address | [optional] 
**MxRecord** | Pointer to **string** | MX record that is required in customer&#39;s DNS to allow the domain to receive bounce and complaint notifications that email providers send you | [optional] 
**TxtRecord** | Pointer to **string** | TXT record that is required in customer&#39;s DNS in order to prove that Amazon SES is authorized to send email from your domain | [optional] 
**MailFromDomainStatus** | Pointer to **string** | The current status of the MAIL FROM verification | [optional] 

## Methods

### NewMailFromAttributes

`func NewMailFromAttributes() *MailFromAttributes`

NewMailFromAttributes instantiates a new MailFromAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMailFromAttributesWithDefaults

`func NewMailFromAttributesWithDefaults() *MailFromAttributes`

NewMailFromAttributesWithDefaults instantiates a new MailFromAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MailFromAttributes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MailFromAttributes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MailFromAttributes) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MailFromAttributes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMxRecord

`func (o *MailFromAttributes) GetMxRecord() string`

GetMxRecord returns the MxRecord field if non-nil, zero value otherwise.

### GetMxRecordOk

`func (o *MailFromAttributes) GetMxRecordOk() (*string, bool)`

GetMxRecordOk returns a tuple with the MxRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMxRecord

`func (o *MailFromAttributes) SetMxRecord(v string)`

SetMxRecord sets MxRecord field to given value.

### HasMxRecord

`func (o *MailFromAttributes) HasMxRecord() bool`

HasMxRecord returns a boolean if a field has been set.

### GetTxtRecord

`func (o *MailFromAttributes) GetTxtRecord() string`

GetTxtRecord returns the TxtRecord field if non-nil, zero value otherwise.

### GetTxtRecordOk

`func (o *MailFromAttributes) GetTxtRecordOk() (*string, bool)`

GetTxtRecordOk returns a tuple with the TxtRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxtRecord

`func (o *MailFromAttributes) SetTxtRecord(v string)`

SetTxtRecord sets TxtRecord field to given value.

### HasTxtRecord

`func (o *MailFromAttributes) HasTxtRecord() bool`

HasTxtRecord returns a boolean if a field has been set.

### GetMailFromDomainStatus

`func (o *MailFromAttributes) GetMailFromDomainStatus() string`

GetMailFromDomainStatus returns the MailFromDomainStatus field if non-nil, zero value otherwise.

### GetMailFromDomainStatusOk

`func (o *MailFromAttributes) GetMailFromDomainStatusOk() (*string, bool)`

GetMailFromDomainStatusOk returns a tuple with the MailFromDomainStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailFromDomainStatus

`func (o *MailFromAttributes) SetMailFromDomainStatus(v string)`

SetMailFromDomainStatus sets MailFromDomainStatus field to given value.

### HasMailFromDomainStatus

`func (o *MailFromAttributes) HasMailFromDomainStatus() bool`

HasMailFromDomainStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


