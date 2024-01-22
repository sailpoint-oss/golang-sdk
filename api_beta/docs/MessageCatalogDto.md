# MessageCatalogDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locale** | Pointer to **string** | The language in which the messages are returned | [optional] 
**Messages** | Pointer to [**[]ResourceBundleMessage**](ResourceBundleMessage.md) | The list of message with their keys and formats | [optional] 

## Methods

### NewMessageCatalogDto

`func NewMessageCatalogDto() *MessageCatalogDto`

NewMessageCatalogDto instantiates a new MessageCatalogDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageCatalogDtoWithDefaults

`func NewMessageCatalogDtoWithDefaults() *MessageCatalogDto`

NewMessageCatalogDtoWithDefaults instantiates a new MessageCatalogDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocale

`func (o *MessageCatalogDto) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *MessageCatalogDto) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *MessageCatalogDto) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *MessageCatalogDto) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetMessages

`func (o *MessageCatalogDto) GetMessages() []ResourceBundleMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *MessageCatalogDto) GetMessagesOk() (*[]ResourceBundleMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *MessageCatalogDto) SetMessages(v []ResourceBundleMessage)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *MessageCatalogDto) HasMessages() bool`

HasMessages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


