# TransformUpdateAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceName** | **string** | A reference to the source to search for the account | 
**AttributeName** | **string** | The name of the attribute on the account to return. This should match the name of the account attribute name visible in the user interface, or on the source schema. | 
**AccountSortAttribute** | Pointer to **string** | The value of this configuration is a string name of the attribute to use when determining the ordering of returned accounts when there are multiple entries | [optional] 
**AccountSortDescending** | Pointer to **bool** | The value of this configuration is a boolean (true/false). Controls the order of the sort when there are multiple accounts. If not defined, the transform will default to false (ascending order) | [optional] 
**AccountReturnFirstLink** | Pointer to **bool** | The value of this configuration is a boolean (true/false). Controls which account to source a value from for an attribute.  If this flag is set to true, the transform returns the value from the first account in the list, even if it is null. If it is set to false, the transform returns the first non-null value. If not defined, the transform will default to false | [optional] 
**AccountFilter** | Pointer to **string** | This expression queries the database to narrow search results. The value of this configuration is a sailpoint.object.Filter expression and used when searching against the database.  The default filter will always include the source and identity, and any subsequent expressions will be combined in an AND operation to the existing search criteria. Only certain searchable attributes are available:  - &#x60;nativeIdentity&#x60; - the Account ID  - &#x60;displayName&#x60; - the Account Name  - &#x60;entitlements&#x60; - a boolean value to determine if the account has entitlements | [optional] 
**AccountPropertyFilter** | Pointer to **string** | This expression is used to search and filter accounts in memory. The value of this configuration is a sailpoint.object.Filter expression and used when searching against the returned resultset.  All account attributes are available for filtering as this operation is performed in memory. | [optional] 
**RequiresPeriodicRefresh** | Pointer to **bool** | A value that indicates whether the transform logic should be re-evaluated every evening as part of the identity refresh process | [optional] [default to false]
**Input** | Pointer to **map[string]interface{}** | This is an optional attribute that can explicitly define the input data which will be fed into the transform logic. If input is not provided, the transform will take its input from the source and attribute combination configured via the UI. | [optional] 
**Values** | **string** | This must evaluate to a JSON string, either through a fixed value or through conditional logic using the Apache Velocity Template Language. | 
**Expression** | **string** | A string value of the date and time components to operation on, along with the math operations to execute.  | 
**PositiveCondition** | **string** | The output of the transform if the expression evalutes to true | 
**NegativeCondition** | **string** | The output of the transform if the expression evalutes to false | 
**FirstDate** | [**DateCompareFirstDate**](DateCompareFirstDate.md) |  | 
**SecondDate** | [**DateCompareSecondDate**](DateCompareSecondDate.md) |  | 
**Operator** | **string** | This is the comparison to perform. | Operation | Description | | --------- | ------- | | LT        | Strictly less than: firstDate &lt; secondDate | | LTE       | Less than or equal to: firstDate &lt;&#x3D; secondDate | | GT        | Strictly greater than: firstDate &gt; secondDate | | GTE       | Greater than or equal to: firstDate &gt;&#x3D; secondDate |  | 
**InputFormat** | Pointer to [**DateFormatInputFormat**](DateFormatInputFormat.md) |  | [optional] 
**OutputFormat** | Pointer to [**DateFormatOutputFormat**](DateFormatOutputFormat.md) |  | [optional] 
**RoundUp** | Pointer to **bool** | A boolean value to indicate whether the transform should round up or down when a rounding &#x60;/&#x60; operation is defined in the expression.    If not provided, the transform will default to &#x60;false&#x60;   &#x60;true&#x60; indicates the transform should round up (i.e., truncate the fractional date/time component indicated and then add one unit of that component)   &#x60;false&#x60; indicates the transform should round down (i.e., truncate the fractional date/time component indicated)  | [optional] 
**DefaultRegion** | Pointer to **string** | This is an optional attribute that can be used to define the region of the phone number to format into.   If defaultRegion is not provided, it will take US as the default country.   The format of the country code should be in [ISO 3166-1 alpha-2 format](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)  | [optional] 
**IgnoreErrors** | Pointer to **bool** | a true or false value representing to move on to the next option if an error (like an Null Pointer Exception) were to occur. | [optional] 
**Name** | **string** | The system (camel-cased) name of the identity attribute to bring in | 
**Operation** | **string** | The operation to perform &#x60;getReferenceIdentityAttribute&#x60; | 
**IncludeNumbers** | **bool** | This must be either \&quot;true\&quot; or \&quot;false\&quot; to indicate whether the generator logic should include numbers | 
**IncludeSpecialChars** | **bool** | This must be either \&quot;true\&quot; or \&quot;false\&quot; to indicate whether the generator logic should include special characters | 
**Length** | **string** | An integer value for the desired length of the final output string | 
**Uid** | **string** | This is the SailPoint User Name (uid) value of the identity whose attribute is desired  As a convenience feature, you can use the &#x60;manager&#x60; keyword to dynamically look up the user&#39;s manager and then get that manager&#39;s identity attribute.  | 
**Substring** | **string** | A substring to search for, searches the entire calling string, and returns the index of the first occurrence of the specified substring. | 
**Format** | Pointer to **string** | An optional value to denote which ISO 3166 format to return. Valid values are:   &#x60;alpha2&#x60; - Two-character country code (e.g., \&quot;US\&quot;); this is the default value if no format is supplied   &#x60;alpha3&#x60; - Three-character country code (e.g., \&quot;USA\&quot;)   &#x60;numeric&#x60; - The numeric country code (e.g., \&quot;840\&quot;)  | [optional] 
**Padding** | Pointer to **string** | A string value representing the character that the incoming data should be padded with to get to the desired length   If not provided, the transform will default to a single space (\&quot; \&quot;) character for padding  | [optional] 
**Table** | **map[string]interface{}** | An attribute of key-value pairs. Each pair identifies the pattern to search for as its key, and the replacement string as its value. | 
**Id** | **string** | This ID specifies the name of the pre-existing transform which you want to use within your current transform | 
**Regex** | **string** | This can be a string or a regex pattern in which you want to replace. | 
**Replacement** | **string** | This is the replacement string that should be substituded wherever the string or pattern is found. | 
**Delimiter** | **string** | This can be either a single character or a regex expression, and is used by the transform to identify the break point between two substrings in the incoming data | 
**Index** | **string** | An integer value for the desired array element after the incoming data has been split into a list; the array is a 0-based object, so the first array element would be index 0, the second element would be index 1, etc. | 
**Throws** | Pointer to **bool** | A boolean (true/false) value which indicates whether an exception should be thrown and returned as an output when an index is out of bounds with the resultant array (i.e., the provided index value is larger than the size of the array)   &#x60;true&#x60; - The transform should return \&quot;IndexOutOfBoundsException\&quot;   &#x60;false&#x60; - The transform should return null   If not provided, the transform will default to false and return a null  | [optional] 
**Begin** | **int32** | The index of the first character to include in the returned substring.   If &#x60;begin&#x60; is set to -1, the transform will begin at character 0 of the input data  | 
**BeginOffset** | Pointer to **int32** | This integer value is the number of characters to add to the begin attribute when returning a substring.   This attribute is only used if begin is not -1.  | [optional] 
**End** | Pointer to **int32** | The index of the first character to exclude from the returned substring.  If end is -1 or not provided at all, the substring transform will return everything up to the end of the input string.  | [optional] 
**EndOffset** | Pointer to **int32** | This integer value is the number of characters to add to the end attribute when returning a substring.   This attribute is only used if end is provided and is not -1.  | [optional] 

## Methods

### NewTransformUpdateAttributes

`func NewTransformUpdateAttributes(sourceName string, attributeName string, values string, expression string, positiveCondition string, negativeCondition string, firstDate DateCompareFirstDate, secondDate DateCompareSecondDate, operator string, name string, operation string, includeNumbers bool, includeSpecialChars bool, length string, uid string, substring string, table map[string]interface{}, id string, regex string, replacement string, delimiter string, index string, begin int32, ) *TransformUpdateAttributes`

NewTransformUpdateAttributes instantiates a new TransformUpdateAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransformUpdateAttributesWithDefaults

`func NewTransformUpdateAttributesWithDefaults() *TransformUpdateAttributes`

NewTransformUpdateAttributesWithDefaults instantiates a new TransformUpdateAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceName

`func (o *TransformUpdateAttributes) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *TransformUpdateAttributes) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *TransformUpdateAttributes) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.


### GetAttributeName

`func (o *TransformUpdateAttributes) GetAttributeName() string`

GetAttributeName returns the AttributeName field if non-nil, zero value otherwise.

### GetAttributeNameOk

`func (o *TransformUpdateAttributes) GetAttributeNameOk() (*string, bool)`

GetAttributeNameOk returns a tuple with the AttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeName

`func (o *TransformUpdateAttributes) SetAttributeName(v string)`

SetAttributeName sets AttributeName field to given value.


### GetAccountSortAttribute

`func (o *TransformUpdateAttributes) GetAccountSortAttribute() string`

GetAccountSortAttribute returns the AccountSortAttribute field if non-nil, zero value otherwise.

### GetAccountSortAttributeOk

`func (o *TransformUpdateAttributes) GetAccountSortAttributeOk() (*string, bool)`

GetAccountSortAttributeOk returns a tuple with the AccountSortAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSortAttribute

`func (o *TransformUpdateAttributes) SetAccountSortAttribute(v string)`

SetAccountSortAttribute sets AccountSortAttribute field to given value.

### HasAccountSortAttribute

`func (o *TransformUpdateAttributes) HasAccountSortAttribute() bool`

HasAccountSortAttribute returns a boolean if a field has been set.

### GetAccountSortDescending

`func (o *TransformUpdateAttributes) GetAccountSortDescending() bool`

GetAccountSortDescending returns the AccountSortDescending field if non-nil, zero value otherwise.

### GetAccountSortDescendingOk

`func (o *TransformUpdateAttributes) GetAccountSortDescendingOk() (*bool, bool)`

GetAccountSortDescendingOk returns a tuple with the AccountSortDescending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSortDescending

`func (o *TransformUpdateAttributes) SetAccountSortDescending(v bool)`

SetAccountSortDescending sets AccountSortDescending field to given value.

### HasAccountSortDescending

`func (o *TransformUpdateAttributes) HasAccountSortDescending() bool`

HasAccountSortDescending returns a boolean if a field has been set.

### GetAccountReturnFirstLink

`func (o *TransformUpdateAttributes) GetAccountReturnFirstLink() bool`

GetAccountReturnFirstLink returns the AccountReturnFirstLink field if non-nil, zero value otherwise.

### GetAccountReturnFirstLinkOk

`func (o *TransformUpdateAttributes) GetAccountReturnFirstLinkOk() (*bool, bool)`

GetAccountReturnFirstLinkOk returns a tuple with the AccountReturnFirstLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountReturnFirstLink

`func (o *TransformUpdateAttributes) SetAccountReturnFirstLink(v bool)`

SetAccountReturnFirstLink sets AccountReturnFirstLink field to given value.

### HasAccountReturnFirstLink

`func (o *TransformUpdateAttributes) HasAccountReturnFirstLink() bool`

HasAccountReturnFirstLink returns a boolean if a field has been set.

### GetAccountFilter

`func (o *TransformUpdateAttributes) GetAccountFilter() string`

GetAccountFilter returns the AccountFilter field if non-nil, zero value otherwise.

### GetAccountFilterOk

`func (o *TransformUpdateAttributes) GetAccountFilterOk() (*string, bool)`

GetAccountFilterOk returns a tuple with the AccountFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountFilter

`func (o *TransformUpdateAttributes) SetAccountFilter(v string)`

SetAccountFilter sets AccountFilter field to given value.

### HasAccountFilter

`func (o *TransformUpdateAttributes) HasAccountFilter() bool`

HasAccountFilter returns a boolean if a field has been set.

### GetAccountPropertyFilter

`func (o *TransformUpdateAttributes) GetAccountPropertyFilter() string`

GetAccountPropertyFilter returns the AccountPropertyFilter field if non-nil, zero value otherwise.

### GetAccountPropertyFilterOk

`func (o *TransformUpdateAttributes) GetAccountPropertyFilterOk() (*string, bool)`

GetAccountPropertyFilterOk returns a tuple with the AccountPropertyFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountPropertyFilter

`func (o *TransformUpdateAttributes) SetAccountPropertyFilter(v string)`

SetAccountPropertyFilter sets AccountPropertyFilter field to given value.

### HasAccountPropertyFilter

`func (o *TransformUpdateAttributes) HasAccountPropertyFilter() bool`

HasAccountPropertyFilter returns a boolean if a field has been set.

### GetRequiresPeriodicRefresh

`func (o *TransformUpdateAttributes) GetRequiresPeriodicRefresh() bool`

GetRequiresPeriodicRefresh returns the RequiresPeriodicRefresh field if non-nil, zero value otherwise.

### GetRequiresPeriodicRefreshOk

`func (o *TransformUpdateAttributes) GetRequiresPeriodicRefreshOk() (*bool, bool)`

GetRequiresPeriodicRefreshOk returns a tuple with the RequiresPeriodicRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresPeriodicRefresh

`func (o *TransformUpdateAttributes) SetRequiresPeriodicRefresh(v bool)`

SetRequiresPeriodicRefresh sets RequiresPeriodicRefresh field to given value.

### HasRequiresPeriodicRefresh

`func (o *TransformUpdateAttributes) HasRequiresPeriodicRefresh() bool`

HasRequiresPeriodicRefresh returns a boolean if a field has been set.

### GetInput

`func (o *TransformUpdateAttributes) GetInput() map[string]interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *TransformUpdateAttributes) GetInputOk() (*map[string]interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *TransformUpdateAttributes) SetInput(v map[string]interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *TransformUpdateAttributes) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetValues

`func (o *TransformUpdateAttributes) GetValues() string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *TransformUpdateAttributes) GetValuesOk() (*string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *TransformUpdateAttributes) SetValues(v string)`

SetValues sets Values field to given value.


### GetExpression

`func (o *TransformUpdateAttributes) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *TransformUpdateAttributes) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *TransformUpdateAttributes) SetExpression(v string)`

SetExpression sets Expression field to given value.


### GetPositiveCondition

`func (o *TransformUpdateAttributes) GetPositiveCondition() string`

GetPositiveCondition returns the PositiveCondition field if non-nil, zero value otherwise.

### GetPositiveConditionOk

`func (o *TransformUpdateAttributes) GetPositiveConditionOk() (*string, bool)`

GetPositiveConditionOk returns a tuple with the PositiveCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositiveCondition

`func (o *TransformUpdateAttributes) SetPositiveCondition(v string)`

SetPositiveCondition sets PositiveCondition field to given value.


### GetNegativeCondition

`func (o *TransformUpdateAttributes) GetNegativeCondition() string`

GetNegativeCondition returns the NegativeCondition field if non-nil, zero value otherwise.

### GetNegativeConditionOk

`func (o *TransformUpdateAttributes) GetNegativeConditionOk() (*string, bool)`

GetNegativeConditionOk returns a tuple with the NegativeCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegativeCondition

`func (o *TransformUpdateAttributes) SetNegativeCondition(v string)`

SetNegativeCondition sets NegativeCondition field to given value.


### GetFirstDate

`func (o *TransformUpdateAttributes) GetFirstDate() DateCompareFirstDate`

GetFirstDate returns the FirstDate field if non-nil, zero value otherwise.

### GetFirstDateOk

`func (o *TransformUpdateAttributes) GetFirstDateOk() (*DateCompareFirstDate, bool)`

GetFirstDateOk returns a tuple with the FirstDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstDate

`func (o *TransformUpdateAttributes) SetFirstDate(v DateCompareFirstDate)`

SetFirstDate sets FirstDate field to given value.


### GetSecondDate

`func (o *TransformUpdateAttributes) GetSecondDate() DateCompareSecondDate`

GetSecondDate returns the SecondDate field if non-nil, zero value otherwise.

### GetSecondDateOk

`func (o *TransformUpdateAttributes) GetSecondDateOk() (*DateCompareSecondDate, bool)`

GetSecondDateOk returns a tuple with the SecondDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondDate

`func (o *TransformUpdateAttributes) SetSecondDate(v DateCompareSecondDate)`

SetSecondDate sets SecondDate field to given value.


### GetOperator

`func (o *TransformUpdateAttributes) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *TransformUpdateAttributes) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *TransformUpdateAttributes) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetInputFormat

`func (o *TransformUpdateAttributes) GetInputFormat() DateFormatInputFormat`

GetInputFormat returns the InputFormat field if non-nil, zero value otherwise.

### GetInputFormatOk

`func (o *TransformUpdateAttributes) GetInputFormatOk() (*DateFormatInputFormat, bool)`

GetInputFormatOk returns a tuple with the InputFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputFormat

`func (o *TransformUpdateAttributes) SetInputFormat(v DateFormatInputFormat)`

SetInputFormat sets InputFormat field to given value.

### HasInputFormat

`func (o *TransformUpdateAttributes) HasInputFormat() bool`

HasInputFormat returns a boolean if a field has been set.

### GetOutputFormat

`func (o *TransformUpdateAttributes) GetOutputFormat() DateFormatOutputFormat`

GetOutputFormat returns the OutputFormat field if non-nil, zero value otherwise.

### GetOutputFormatOk

`func (o *TransformUpdateAttributes) GetOutputFormatOk() (*DateFormatOutputFormat, bool)`

GetOutputFormatOk returns a tuple with the OutputFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFormat

`func (o *TransformUpdateAttributes) SetOutputFormat(v DateFormatOutputFormat)`

SetOutputFormat sets OutputFormat field to given value.

### HasOutputFormat

`func (o *TransformUpdateAttributes) HasOutputFormat() bool`

HasOutputFormat returns a boolean if a field has been set.

### GetRoundUp

`func (o *TransformUpdateAttributes) GetRoundUp() bool`

GetRoundUp returns the RoundUp field if non-nil, zero value otherwise.

### GetRoundUpOk

`func (o *TransformUpdateAttributes) GetRoundUpOk() (*bool, bool)`

GetRoundUpOk returns a tuple with the RoundUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoundUp

`func (o *TransformUpdateAttributes) SetRoundUp(v bool)`

SetRoundUp sets RoundUp field to given value.

### HasRoundUp

`func (o *TransformUpdateAttributes) HasRoundUp() bool`

HasRoundUp returns a boolean if a field has been set.

### GetDefaultRegion

`func (o *TransformUpdateAttributes) GetDefaultRegion() string`

GetDefaultRegion returns the DefaultRegion field if non-nil, zero value otherwise.

### GetDefaultRegionOk

`func (o *TransformUpdateAttributes) GetDefaultRegionOk() (*string, bool)`

GetDefaultRegionOk returns a tuple with the DefaultRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRegion

`func (o *TransformUpdateAttributes) SetDefaultRegion(v string)`

SetDefaultRegion sets DefaultRegion field to given value.

### HasDefaultRegion

`func (o *TransformUpdateAttributes) HasDefaultRegion() bool`

HasDefaultRegion returns a boolean if a field has been set.

### GetIgnoreErrors

`func (o *TransformUpdateAttributes) GetIgnoreErrors() bool`

GetIgnoreErrors returns the IgnoreErrors field if non-nil, zero value otherwise.

### GetIgnoreErrorsOk

`func (o *TransformUpdateAttributes) GetIgnoreErrorsOk() (*bool, bool)`

GetIgnoreErrorsOk returns a tuple with the IgnoreErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreErrors

`func (o *TransformUpdateAttributes) SetIgnoreErrors(v bool)`

SetIgnoreErrors sets IgnoreErrors field to given value.

### HasIgnoreErrors

`func (o *TransformUpdateAttributes) HasIgnoreErrors() bool`

HasIgnoreErrors returns a boolean if a field has been set.

### GetName

`func (o *TransformUpdateAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TransformUpdateAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TransformUpdateAttributes) SetName(v string)`

SetName sets Name field to given value.


### GetOperation

`func (o *TransformUpdateAttributes) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *TransformUpdateAttributes) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *TransformUpdateAttributes) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetIncludeNumbers

`func (o *TransformUpdateAttributes) GetIncludeNumbers() bool`

GetIncludeNumbers returns the IncludeNumbers field if non-nil, zero value otherwise.

### GetIncludeNumbersOk

`func (o *TransformUpdateAttributes) GetIncludeNumbersOk() (*bool, bool)`

GetIncludeNumbersOk returns a tuple with the IncludeNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeNumbers

`func (o *TransformUpdateAttributes) SetIncludeNumbers(v bool)`

SetIncludeNumbers sets IncludeNumbers field to given value.


### GetIncludeSpecialChars

`func (o *TransformUpdateAttributes) GetIncludeSpecialChars() bool`

GetIncludeSpecialChars returns the IncludeSpecialChars field if non-nil, zero value otherwise.

### GetIncludeSpecialCharsOk

`func (o *TransformUpdateAttributes) GetIncludeSpecialCharsOk() (*bool, bool)`

GetIncludeSpecialCharsOk returns a tuple with the IncludeSpecialChars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSpecialChars

`func (o *TransformUpdateAttributes) SetIncludeSpecialChars(v bool)`

SetIncludeSpecialChars sets IncludeSpecialChars field to given value.


### GetLength

`func (o *TransformUpdateAttributes) GetLength() string`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *TransformUpdateAttributes) GetLengthOk() (*string, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *TransformUpdateAttributes) SetLength(v string)`

SetLength sets Length field to given value.


### GetUid

`func (o *TransformUpdateAttributes) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *TransformUpdateAttributes) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *TransformUpdateAttributes) SetUid(v string)`

SetUid sets Uid field to given value.


### GetSubstring

`func (o *TransformUpdateAttributes) GetSubstring() string`

GetSubstring returns the Substring field if non-nil, zero value otherwise.

### GetSubstringOk

`func (o *TransformUpdateAttributes) GetSubstringOk() (*string, bool)`

GetSubstringOk returns a tuple with the Substring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubstring

`func (o *TransformUpdateAttributes) SetSubstring(v string)`

SetSubstring sets Substring field to given value.


### GetFormat

`func (o *TransformUpdateAttributes) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *TransformUpdateAttributes) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *TransformUpdateAttributes) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *TransformUpdateAttributes) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetPadding

`func (o *TransformUpdateAttributes) GetPadding() string`

GetPadding returns the Padding field if non-nil, zero value otherwise.

### GetPaddingOk

`func (o *TransformUpdateAttributes) GetPaddingOk() (*string, bool)`

GetPaddingOk returns a tuple with the Padding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPadding

`func (o *TransformUpdateAttributes) SetPadding(v string)`

SetPadding sets Padding field to given value.

### HasPadding

`func (o *TransformUpdateAttributes) HasPadding() bool`

HasPadding returns a boolean if a field has been set.

### GetTable

`func (o *TransformUpdateAttributes) GetTable() map[string]interface{}`

GetTable returns the Table field if non-nil, zero value otherwise.

### GetTableOk

`func (o *TransformUpdateAttributes) GetTableOk() (*map[string]interface{}, bool)`

GetTableOk returns a tuple with the Table field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTable

`func (o *TransformUpdateAttributes) SetTable(v map[string]interface{})`

SetTable sets Table field to given value.


### GetId

`func (o *TransformUpdateAttributes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransformUpdateAttributes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransformUpdateAttributes) SetId(v string)`

SetId sets Id field to given value.


### GetRegex

`func (o *TransformUpdateAttributes) GetRegex() string`

GetRegex returns the Regex field if non-nil, zero value otherwise.

### GetRegexOk

`func (o *TransformUpdateAttributes) GetRegexOk() (*string, bool)`

GetRegexOk returns a tuple with the Regex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegex

`func (o *TransformUpdateAttributes) SetRegex(v string)`

SetRegex sets Regex field to given value.


### GetReplacement

`func (o *TransformUpdateAttributes) GetReplacement() string`

GetReplacement returns the Replacement field if non-nil, zero value otherwise.

### GetReplacementOk

`func (o *TransformUpdateAttributes) GetReplacementOk() (*string, bool)`

GetReplacementOk returns a tuple with the Replacement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacement

`func (o *TransformUpdateAttributes) SetReplacement(v string)`

SetReplacement sets Replacement field to given value.


### GetDelimiter

`func (o *TransformUpdateAttributes) GetDelimiter() string`

GetDelimiter returns the Delimiter field if non-nil, zero value otherwise.

### GetDelimiterOk

`func (o *TransformUpdateAttributes) GetDelimiterOk() (*string, bool)`

GetDelimiterOk returns a tuple with the Delimiter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelimiter

`func (o *TransformUpdateAttributes) SetDelimiter(v string)`

SetDelimiter sets Delimiter field to given value.


### GetIndex

`func (o *TransformUpdateAttributes) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *TransformUpdateAttributes) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *TransformUpdateAttributes) SetIndex(v string)`

SetIndex sets Index field to given value.


### GetThrows

`func (o *TransformUpdateAttributes) GetThrows() bool`

GetThrows returns the Throws field if non-nil, zero value otherwise.

### GetThrowsOk

`func (o *TransformUpdateAttributes) GetThrowsOk() (*bool, bool)`

GetThrowsOk returns a tuple with the Throws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrows

`func (o *TransformUpdateAttributes) SetThrows(v bool)`

SetThrows sets Throws field to given value.

### HasThrows

`func (o *TransformUpdateAttributes) HasThrows() bool`

HasThrows returns a boolean if a field has been set.

### GetBegin

`func (o *TransformUpdateAttributes) GetBegin() int32`

GetBegin returns the Begin field if non-nil, zero value otherwise.

### GetBeginOk

`func (o *TransformUpdateAttributes) GetBeginOk() (*int32, bool)`

GetBeginOk returns a tuple with the Begin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBegin

`func (o *TransformUpdateAttributes) SetBegin(v int32)`

SetBegin sets Begin field to given value.


### GetBeginOffset

`func (o *TransformUpdateAttributes) GetBeginOffset() int32`

GetBeginOffset returns the BeginOffset field if non-nil, zero value otherwise.

### GetBeginOffsetOk

`func (o *TransformUpdateAttributes) GetBeginOffsetOk() (*int32, bool)`

GetBeginOffsetOk returns a tuple with the BeginOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeginOffset

`func (o *TransformUpdateAttributes) SetBeginOffset(v int32)`

SetBeginOffset sets BeginOffset field to given value.

### HasBeginOffset

`func (o *TransformUpdateAttributes) HasBeginOffset() bool`

HasBeginOffset returns a boolean if a field has been set.

### GetEnd

`func (o *TransformUpdateAttributes) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *TransformUpdateAttributes) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *TransformUpdateAttributes) SetEnd(v int32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *TransformUpdateAttributes) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetEndOffset

`func (o *TransformUpdateAttributes) GetEndOffset() int32`

GetEndOffset returns the EndOffset field if non-nil, zero value otherwise.

### GetEndOffsetOk

`func (o *TransformUpdateAttributes) GetEndOffsetOk() (*int32, bool)`

GetEndOffsetOk returns a tuple with the EndOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndOffset

`func (o *TransformUpdateAttributes) SetEndOffset(v int32)`

SetEndOffset sets EndOffset field to given value.

### HasEndOffset

`func (o *TransformUpdateAttributes) HasEndOffset() bool`

HasEndOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


