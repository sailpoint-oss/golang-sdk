# Aggregation1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nested** | Pointer to [**NestedAggregation**](NestedAggregation.md) |  | [optional] 
**Metric** | Pointer to [**MetricAggregation**](MetricAggregation.md) |  | [optional] 
**Filter** | Pointer to [**FilterAggregation**](FilterAggregation.md) |  | [optional] 
**Bucket** | Pointer to [**BucketAggregation**](BucketAggregation.md) |  | [optional] 
**SubAggregation** | Pointer to [**Aggregation2**](Aggregation2.md) |  | [optional] 

## Methods

### NewAggregation1

`func NewAggregation1() *Aggregation1`

NewAggregation1 instantiates a new Aggregation1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregation1WithDefaults

`func NewAggregation1WithDefaults() *Aggregation1`

NewAggregation1WithDefaults instantiates a new Aggregation1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNested

`func (o *Aggregation1) GetNested() NestedAggregation`

GetNested returns the Nested field if non-nil, zero value otherwise.

### GetNestedOk

`func (o *Aggregation1) GetNestedOk() (*NestedAggregation, bool)`

GetNestedOk returns a tuple with the Nested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNested

`func (o *Aggregation1) SetNested(v NestedAggregation)`

SetNested sets Nested field to given value.

### HasNested

`func (o *Aggregation1) HasNested() bool`

HasNested returns a boolean if a field has been set.

### GetMetric

`func (o *Aggregation1) GetMetric() MetricAggregation`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *Aggregation1) GetMetricOk() (*MetricAggregation, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *Aggregation1) SetMetric(v MetricAggregation)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *Aggregation1) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetFilter

`func (o *Aggregation1) GetFilter() FilterAggregation`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *Aggregation1) GetFilterOk() (*FilterAggregation, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *Aggregation1) SetFilter(v FilterAggregation)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *Aggregation1) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetBucket

`func (o *Aggregation1) GetBucket() BucketAggregation`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *Aggregation1) GetBucketOk() (*BucketAggregation, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *Aggregation1) SetBucket(v BucketAggregation)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *Aggregation1) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetSubAggregation

`func (o *Aggregation1) GetSubAggregation() Aggregation2`

GetSubAggregation returns the SubAggregation field if non-nil, zero value otherwise.

### GetSubAggregationOk

`func (o *Aggregation1) GetSubAggregationOk() (*Aggregation2, bool)`

GetSubAggregationOk returns a tuple with the SubAggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubAggregation

`func (o *Aggregation1) SetSubAggregation(v Aggregation2)`

SetSubAggregation sets SubAggregation field to given value.

### HasSubAggregation

`func (o *Aggregation1) HasSubAggregation() bool`

HasSubAggregation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


