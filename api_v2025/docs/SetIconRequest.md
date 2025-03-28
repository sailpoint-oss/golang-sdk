# SetIconRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Image** | ***os.File** | file with icon. Allowed mime-types [&#39;image/png&#39;, &#39;image/jpeg&#39;] | 

## Methods

### NewSetIconRequest

`func NewSetIconRequest(image *os.File, ) *SetIconRequest`

NewSetIconRequest instantiates a new SetIconRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetIconRequestWithDefaults

`func NewSetIconRequestWithDefaults() *SetIconRequest`

NewSetIconRequestWithDefaults instantiates a new SetIconRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImage

`func (o *SetIconRequest) GetImage() *os.File`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *SetIconRequest) GetImageOk() (**os.File, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *SetIconRequest) SetImage(v *os.File)`

SetImage sets Image field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


