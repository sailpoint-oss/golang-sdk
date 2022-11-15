# BaseDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Type** | [**DocumentType**](DocumentType.md) |  | 

## Methods

### NewBaseDocument

`func NewBaseDocument(id string, name string, type_ DocumentType, ) *BaseDocument`

NewBaseDocument instantiates a new BaseDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseDocumentWithDefaults

`func NewBaseDocumentWithDefaults() *BaseDocument`

NewBaseDocumentWithDefaults instantiates a new BaseDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BaseDocument) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BaseDocument) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BaseDocument) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BaseDocument) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BaseDocument) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BaseDocument) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *BaseDocument) GetType() DocumentType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BaseDocument) GetTypeOk() (*DocumentType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BaseDocument) SetType(v DocumentType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


