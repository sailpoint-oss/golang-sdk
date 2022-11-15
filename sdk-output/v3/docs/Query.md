# Query

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | Pointer to **string** | The query using the Elasticsearch [Query String Query](https://www.elastic.co/guide/en/elasticsearch/reference/5.2/query-dsl-query-string-query.html#query-string) syntax from the Query DSL extended by SailPoint to support Nested queries. | [optional] 
**Fields** | Pointer to **[]string** | The fields to which the specified query will be applied.  The available fields are dependent on the indice(s) being searched on.  Please refer to the response schema of this API for a list of available fields. | [optional] 
**TimeZone** | Pointer to **string** | The time zone to be applied to any range query related to dates. | [optional] 
**InnerHit** | Pointer to [**InnerHit**](InnerHit.md) |  | [optional] 

## Methods

### NewQuery

`func NewQuery() *Query`

NewQuery instantiates a new Query object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryWithDefaults

`func NewQueryWithDefaults() *Query`

NewQueryWithDefaults instantiates a new Query object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *Query) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *Query) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *Query) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *Query) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetFields

`func (o *Query) GetFields() []string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *Query) GetFieldsOk() (*[]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *Query) SetFields(v []string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *Query) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetTimeZone

`func (o *Query) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *Query) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *Query) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *Query) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetInnerHit

`func (o *Query) GetInnerHit() InnerHit`

GetInnerHit returns the InnerHit field if non-nil, zero value otherwise.

### GetInnerHitOk

`func (o *Query) GetInnerHitOk() (*InnerHit, bool)`

GetInnerHitOk returns a tuple with the InnerHit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnerHit

`func (o *Query) SetInnerHit(v InnerHit)`

SetInnerHit sets InnerHit field to given value.

### HasInnerHit

`func (o *Query) HasInnerHit() bool`

HasInnerHit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


