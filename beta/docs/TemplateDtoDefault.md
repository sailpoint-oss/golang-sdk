# TemplateDtoDefault

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Medium** | Pointer to **string** | The message medium. More mediums may be added in the future. | [optional] 
**Locale** | Pointer to **string** | The locale for the message text, a BCP 47 language tag. | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**Header** | Pointer to **string** |  | [optional] 
**Body** | Pointer to **string** |  | [optional] 
**Footer** | Pointer to **string** |  | [optional] 
**From** | Pointer to **string** |  | [optional] 
**ReplyTo** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewTemplateDtoDefault

`func NewTemplateDtoDefault() *TemplateDtoDefault`

NewTemplateDtoDefault instantiates a new TemplateDtoDefault object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateDtoDefaultWithDefaults

`func NewTemplateDtoDefaultWithDefaults() *TemplateDtoDefault`

NewTemplateDtoDefaultWithDefaults instantiates a new TemplateDtoDefault object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *TemplateDtoDefault) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TemplateDtoDefault) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TemplateDtoDefault) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *TemplateDtoDefault) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *TemplateDtoDefault) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TemplateDtoDefault) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TemplateDtoDefault) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TemplateDtoDefault) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMedium

`func (o *TemplateDtoDefault) GetMedium() string`

GetMedium returns the Medium field if non-nil, zero value otherwise.

### GetMediumOk

`func (o *TemplateDtoDefault) GetMediumOk() (*string, bool)`

GetMediumOk returns a tuple with the Medium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedium

`func (o *TemplateDtoDefault) SetMedium(v string)`

SetMedium sets Medium field to given value.

### HasMedium

`func (o *TemplateDtoDefault) HasMedium() bool`

HasMedium returns a boolean if a field has been set.

### GetLocale

`func (o *TemplateDtoDefault) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *TemplateDtoDefault) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *TemplateDtoDefault) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *TemplateDtoDefault) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetSubject

`func (o *TemplateDtoDefault) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *TemplateDtoDefault) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *TemplateDtoDefault) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *TemplateDtoDefault) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetHeader

`func (o *TemplateDtoDefault) GetHeader() string`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *TemplateDtoDefault) GetHeaderOk() (*string, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *TemplateDtoDefault) SetHeader(v string)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *TemplateDtoDefault) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetBody

`func (o *TemplateDtoDefault) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *TemplateDtoDefault) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *TemplateDtoDefault) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *TemplateDtoDefault) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetFooter

`func (o *TemplateDtoDefault) GetFooter() string`

GetFooter returns the Footer field if non-nil, zero value otherwise.

### GetFooterOk

`func (o *TemplateDtoDefault) GetFooterOk() (*string, bool)`

GetFooterOk returns a tuple with the Footer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooter

`func (o *TemplateDtoDefault) SetFooter(v string)`

SetFooter sets Footer field to given value.

### HasFooter

`func (o *TemplateDtoDefault) HasFooter() bool`

HasFooter returns a boolean if a field has been set.

### GetFrom

`func (o *TemplateDtoDefault) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *TemplateDtoDefault) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *TemplateDtoDefault) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *TemplateDtoDefault) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetReplyTo

`func (o *TemplateDtoDefault) GetReplyTo() string`

GetReplyTo returns the ReplyTo field if non-nil, zero value otherwise.

### GetReplyToOk

`func (o *TemplateDtoDefault) GetReplyToOk() (*string, bool)`

GetReplyToOk returns a tuple with the ReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyTo

`func (o *TemplateDtoDefault) SetReplyTo(v string)`

SetReplyTo sets ReplyTo field to given value.

### HasReplyTo

`func (o *TemplateDtoDefault) HasReplyTo() bool`

HasReplyTo returns a boolean if a field has been set.

### GetDescription

`func (o *TemplateDtoDefault) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TemplateDtoDefault) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TemplateDtoDefault) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TemplateDtoDefault) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


