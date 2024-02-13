# \PasswordManagementAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPasswordChangeStatus**](PasswordManagementAPI.md#GetPasswordChangeStatus) | **Get** /password-change-status/{id} | Get Password Change Request Status
[**QueryPasswordInfo**](PasswordManagementAPI.md#QueryPasswordInfo) | **Post** /query-password-info | Query Password Info
[**SetPassword**](PasswordManagementAPI.md#SetPassword) | **Post** /set-password | Set Identity&#39;s Password



## Get Password Change Request Status

> PasswordStatus GetPasswordChangeStatus(ctx, id).Execute()

This API returns the status of a password change request. A token with identity owner or trusted API client application authority is required to call this API.

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
    id := "089899f13a8f4da7824996191587bab9" // string | Password change request ID

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.PasswordManagementAPI.GetPasswordChangeStatus(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordManagementAPI.GetPasswordChangeStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPasswordChangeStatus`: PasswordStatus
    fmt.Fprintf(os.Stdout, "Response from `PasswordManagementAPI.GetPasswordChangeStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Password change request ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPasswordChangeStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PasswordStatus**](PasswordStatus.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Query Password Info

> PasswordInfo QueryPasswordInfo(ctx).PasswordInfoQueryDTO(passwordInfoQueryDTO).Execute()

This API is used to query password related information. 

A token with [API authority](https://developer.sailpoint.com/idn/api/authentication#client-credentials-grant-flow) 
is required to call this API.  "API authority" refers to a token that only has the "client_credentials" 
grant type, and therefore no user context. A [personal access token](https://developer.sailpoint.com/idn/api/authentication#personal-access-tokens) 
or a token generated with the [authorization_code](https://developer.sailpoint.com/idn/api/authentication#authorization-code-grant-flow) 
grant type will **NOT** work on this endpoint, and a `403 Forbidden` response 
will be returned.


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
    passwordInfoQueryDTO := *sailpoint.NewPasswordInfoQueryDTO() // PasswordInfoQueryDTO | 

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.PasswordManagementAPI.QueryPasswordInfo(context.Background()).PasswordInfoQueryDTO(passwordInfoQueryDTO).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordManagementAPI.QueryPasswordInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryPasswordInfo`: PasswordInfo
    fmt.Fprintf(os.Stdout, "Response from `PasswordManagementAPI.QueryPasswordInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryPasswordInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **passwordInfoQueryDTO** | [**PasswordInfoQueryDTO**](PasswordInfoQueryDTO.md) |  | 

### Return type

[**PasswordInfo**](PasswordInfo.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Set Identity's Password

> PasswordChangeResponse SetPassword(ctx).PasswordChangeRequest(passwordChangeRequest).Execute()

This API is used to set a password for an identity. 

An identity can change their own password (as well as any of their accounts' passwords) if they use a token generated by their IDN user, such as a [personal access token](https://developer.sailpoint.com/idn/api/authentication#personal-access-tokens) or ["authorization_code" derived OAuth token](https://developer.sailpoint.com/idn/api/authentication#authorization-code-grant-flow).

A token with [API authority](https://developer.sailpoint.com/idn/api/authentication#client-credentials-grant-flow) can be used to change **any** identity's password or the password of any of the identity's accounts. 
"API authority" refers to a token that only has the "client_credentials" grant type.

You can use this endpoint to generate an `encryptedPassword` (RSA encrypted using publicKey). 
To do so, follow these steps:

1. Use [Query Password Info](https://developer.sailpoint.com/idn/api/v3/query-password-info) to get the following information: `identityId`, `sourceId`, `publicKeyId`, `publicKey`, `accounts`, and `policies`. 

2. Choose an account from the previous response that you will provide as an `accountId` in your request to set an encrypted password. 

3. Use [Set Identity's Password](https://developer.sailpoint.com/idn/api/v3/set-password) and provide the information you got from your earlier query. Then add this code to your request to get the encrypted password:

```java
import javax.crypto.Cipher;
import java.security.KeyFactory;
import java.security.PublicKey;
import java.security.spec.X509EncodedKeySpec;
import java util.Base64;

String encrypt(String publicKey, String toEncrypt) throws Exception {
  byte[] publicKeyBytes = Base64.getDecoder().decode(publicKey);
  byte[] encryptedBytes = encryptRsa(publicKeyBytes, toEncrypt.getBytes("UTF-8"));
  return Base64.getEncoder().encodeToString(encryptedBytes);
}

private byte[] encryptRsa(byte[] publicKeyBytes, byte[] toEncryptBytes) throws Exception {
  PublicKey key = KeyFactory.getInstance("RSA").generatePublic(new X509EncodedKeySpec(publicKeyBytes));
  String transformation = "RSA/ECB/PKCS1Padding";
  Cipher cipher = Cipher.getInstance(transformation);
  cipher.init(1, key);
  return cipher.doFinal(toEncryptBytes);
}
```    

In this example, `toEncrypt` refers to the plain text password you are setting and then encrypting, and the `publicKey` refers to the publicKey you got from the first request you sent. 

You can then use [Get Password Change Request Status](https://developer.sailpoint.com/idn/api/v3/get-password-change-status) to check the password change request status. To do so, you must provide the `requestId` from your earlier request to set the password. 


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
    passwordChangeRequest := *sailpoint.NewPasswordChangeRequest() // PasswordChangeRequest | 

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.PasswordManagementAPI.SetPassword(context.Background()).PasswordChangeRequest(passwordChangeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordManagementAPI.SetPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetPassword`: PasswordChangeResponse
    fmt.Fprintf(os.Stdout, "Response from `PasswordManagementAPI.SetPassword`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **passwordChangeRequest** | [**PasswordChangeRequest**](PasswordChangeRequest.md) |  | 

### Return type

[**PasswordChangeResponse**](PasswordChangeResponse.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

