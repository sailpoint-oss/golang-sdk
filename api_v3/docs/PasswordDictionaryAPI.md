# \PasswordDictionaryAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPasswordDictionary**](PasswordDictionaryAPI.md#GetPasswordDictionary) | **Get** /password-dictionary | Get Password Dictionary
[**PutPasswordDictionary**](PasswordDictionaryAPI.md#PutPasswordDictionary) | **Put** /password-dictionary | Update Password Dictionary



## Get Password Dictionary

> string GetPasswordDictionary(ctx).Execute()

This gets password dictionary for the organization.
A token with ORG_ADMIN authority is required to call this API.
The password dictionary file can contain lines that are:
1. comment lines - the first character is '#', can be 128 Unicode codepoints in length, and are ignored during processing
2. empty lines
3. locale line - the first line that starts with "locale=" is considered to be locale line, the rest are treated as normal content lines
4. line containing the password dictionary word - it must start with non-whitespace character and only non-whitespace characters are allowed;
        maximum length of the line is 128 Unicode codepoints


Password dictionary file may not contain more than 2,500 lines (not counting whitespace lines, comment lines and locale line).
  Password dict file must contain UTF-8 characters only.

# Sample password text file

```

# Password dictionary small test file

locale=en_US

# Password dictionary prohibited words

qwerty
abcd
aaaaa
password
qazxsws

```

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.PasswordDictionaryAPI.GetPasswordDictionary(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordDictionaryAPI.GetPasswordDictionary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPasswordDictionary`: string
    fmt.Fprintf(os.Stdout, "Response from `PasswordDictionaryAPI.GetPasswordDictionary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPasswordDictionaryRequest struct via the builder pattern


### Return type

**string**

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update Password Dictionary

> PutPasswordDictionary(ctx).File(file).Execute()

This updates password dictionary for the organization.
A token with ORG_ADMIN authority is required to call this API.
The password dictionary file can contain lines that are:
1. comment lines - the first character is '#', can be 128 Unicode codepoints in length, and are ignored during processing
2. empty lines
3. locale line - the first line that starts with "locale=" is considered to be locale line, the rest are treated as normal content lines
4. line containing the password dictionary word - it must start with non-whitespace character and only non-whitespace characters are allowed;
        maximum length of the line is 128 Unicode codepoints


Password dictionary file may not contain more than 2,500 lines (not counting whitespace lines, comment lines and locale line).
  Password dict file must contain UTF-8 characters only.

# Sample password text file

```

# Password dictionary small test file

locale=en_US

# Password dictionary prohibited words

qwerty
abcd
aaaaa
password
qazxsws

```

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    file := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    r, err := apiClient.V3.PasswordDictionaryAPI.PutPasswordDictionary(context.Background()).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordDictionaryAPI.PutPasswordDictionary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutPasswordDictionaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** |  | 

### Return type

 (empty response body)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

