---
id: nerm-workflow-actions
title: WorkflowActions
pagination_label: WorkflowActions
sidebar_label: WorkflowActions
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'WorkflowActions', 'NERMWorkflowActions'] 
slug: /tools/sdk/go/nerm/methods/workflow-actions
tags: ['SDK', 'Software Development Kit', 'WorkflowActions', 'NERMWorkflowActions']
---

# WorkflowActionsAPI
   
All URIs are relative to *https://acmeco.nonemployee.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create-approval-action**](#create-approval-action) | **Post** `/workflow_actions/approval_actions` | Create an approval action
[**create-ask-security-question-action**](#create-ask-security-question-action) | **Post** `/workflow_actions/ask_security_question_actions` | Create ask security question action
[**create-auto-assign-action**](#create-auto-assign-action) | **Post** `/workflow_actions/auto_assign_actions` | Create an auto assign action
[**create-batch-update-action**](#create-batch-update-action) | **Post** `/workflow_actions/batch_update_actions` | Create a batch update action
[**create-close-session-action**](#create-close-session-action) | **Post** `/workflow_actions/close_session_actions` | Create a close session action
[**create-contributors-action**](#create-contributors-action) | **Post** `/workflow_actions/contributors_actions` | Create a contributors action
[**create-create-profile-action**](#create-create-profile-action) | **Post** `/workflow_actions/create_profile_actions` | Create a create profile action
[**create-duplicate-prevention-action**](#create-duplicate-prevention-action) | **Post** `/workflow_actions/duplicate_prevention_actions` | Create a duplicate prevention action
[**create-email-verification-action**](#create-email-verification-action) | **Post** `/workflow_actions/email_verification_actions` | Create an email verification action
[**create-fulfillment-action**](#create-fulfillment-action) | **Post** `/workflow_actions/fulfillment_actions` | Create a fulfillment action
[**create-identity-proofing-action**](#create-identity-proofing-action) | **Post** `/workflow_actions/identity_proofing_actions` | Create an identity proofing action
[**create-invitation-action**](#create-invitation-action) | **Post** `/workflow_actions/invitation_actions` | Create an invitation action
[**create-ldap-action**](#create-ldap-action) | **Post** `/workflow_actions/ldap_actions` | Create a ldap action
[**create-notification-action**](#create-notification-action) | **Post** `/workflow_actions/notification_actions` | Create a notification action
[**create-password-reset-action**](#create-password-reset-action) | **Post** `/workflow_actions/password_reset_actions` | Create a password reset action
[**create-profile-check-action**](#create-profile-check-action) | **Post** `/workflow_actions/profile_check_actions` | Create a profile check action
[**create-profile-select-action**](#create-profile-select-action) | **Post** `/workflow_actions/profile_select_actions` | Create a profile select action
[**create-request-action**](#create-request-action) | **Post** `/workflow_actions/request_actions` | Create a request action
[**create-rest-api-action**](#create-rest-api-action) | **Post** `/workflow_actions/rest_api_actions` | Create a REST API action
[**create-review-action**](#create-review-action) | **Post** `/workflow_actions/review_actions` | Create a review action
[**create-run-workflow-action**](#create-run-workflow-action) | **Post** `/workflow_actions/run_workflow_actions` | Create a run workflow action
[**create-set-attributes-action**](#create-set-attributes-action) | **Post** `/workflow_actions/set_attributes_actions` | Create a set attributes action
[**create-set-security-question-action**](#create-set-security-question-action) | **Post** `/workflow_actions/set_security_question_actions` | Create set security question action
[**create-soap-api-action**](#create-soap-api-action) | **Post** `/workflow_actions/soap_api_actions` | Create a SOAP API action
[**create-status-change-action**](#create-status-change-action) | **Post** `/workflow_actions/status_change_actions` | Create a status change action
[**create-unassign-action**](#create-unassign-action) | **Post** `/workflow_actions/unassign_actions` | Create an unassign action
[**create-update-profile-action**](#create-update-profile-action) | **Post** `/workflow_actions/update_profile_actions` | Create an update profile action
[**create-username-password-action**](#create-username-password-action) | **Post** `/workflow_actions/username_password_actions` | Create a username password action
[**get-workflow-actions**](#get-workflow-actions) | **Get** `/workflow_actions` | Get Workflow Actions


## create-approval-action
Create an approval action
Create an approval action. The following supporting objects will need to be created after this action is created (which are tied together via workflow_action_id) - workflow_action_performers, workflow_action_roles, workflow_action_performer_notification_email, workflow_action_approval_email, workflow_action_rejection_email. These supporting objects must be created for this action to be complete (APIs for these supporting objects not yet implemented, use UI).

[API Spec](https://developer.sailpoint.com/docs/api/NERM/create-approval-action)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApprovalActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createApprovalActionRequest** | [**CreateApprovalActionRequest**](../models/create-approval-action-request) |  | 

### Return type

[**CreateApprovalAction200Response**](../models/create-approval-action200-response)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
  "encoding/json"
    NERM "github.com/sailpoint-oss/golang-sdk/v2/api_nerm"
	sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    createapprovalactionrequest := []byte(``) // CreateApprovalActionRequest | 

    var createApprovalActionRequest NERM.CreateApprovalActionRequest
    if err := json.Unmarshal(createapprovalactionrequest, &createApprovalActionRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.WorkflowActionsAPI.CreateApprovalAction(context.Background()).CreateApprovalActionRequest(createApprovalActionRequest).Execute()
	  //resp, r, err := apiClient.NERM.WorkflowActionsAPI.CreateApprovalAction(context.Background()).CreateApprovalActionRequest(createApprovalActionRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `WorkflowActionsAPI.CreateApprovalAction``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApprovalAction`: CreateApprovalAction200Response
    fmt.Fprintf(os.Stdout, "Response from `WorkflowActionsAPI.CreateApprovalAction`: %v\n", resp)
}
```

[[Back to top]](#)

## create-ask-security-question-action
Create ask security question action
Create an ask security question action

[API Spec](https://developer.sailpoint.com/docs/api/NERM/create-ask-security-question-action)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAskSecurityQuestionActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAskSecurityQuestionActionRequest** | [**CreateAskSecurityQuestionActionRequest**](../models/create-ask-security-question-action-request) |  | 

### Return type

[**CreateApprovalAction200Response**](../models/create-approval-action200-response)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
  "encoding/json"
    NERM "github.com/sailpoint-oss/golang-sdk/v2/api_nerm"
	sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    createasksecurityquestionactionrequest := []byte(``) // CreateAskSecurityQuestionActionRequest | 

    var createAskSecurityQuestionActionRequest NERM.CreateAskSecurityQuestionActionRequest
    if err := json.Unmarshal(createasksecurityquestionactionrequest, &createAskSecurityQuestionActionRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.WorkflowActionsAPI.CreateAskSecurityQuestionAction(context.Background()).CreateAskSecurityQuestionActionRequest(createAskSecurityQuestionActionRequest).Execute()
	  //resp, r, err := apiClient.NERM.WorkflowActionsAPI.CreateAskSecurityQuestionAction(context.Background()).CreateAskSecurityQuestionActionRequest(createAskSecurityQuestionActionRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `WorkflowActionsAPI.CreateAskSecurityQuestionAction``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAskSecurityQuestionAction`: CreateApprovalAction200Response
    fmt.Fprintf(os.Stdout, "Response from `WorkflowActionsAPI.CreateAskSecurityQuestionAction`: %v\n", resp)
}
```

[[Back to top]](#)

## create-auto-assign-action
Create an auto assign action
Create an auto assign action. The following supporting objects will need to be created after this action is created (which are tied together via workflow_action_id) - workflow_action_roles. These supporting objects must be created for this action to be complete (APIs for these supporting objects not yet implemented, use UI).

[API Spec](https://developer.sailpoint.com/docs/api/NERM/create-auto-assign-action)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAutoAssignActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAutoAssignActionRequest** | [**CreateAutoAssignActionRequest**](../models/create-auto-assign-action-request) |  | 

### Return type

[**CreateApprovalAction200Response**](../models/create-approval-action200-response)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
  "encoding/json"
    NERM "github.com/sailpoint-oss/golang-sdk/v2/api_nerm"
	sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    createautoassignactionrequest := []byte(``) // CreateAutoAssignActionRequest | 

    var createAutoAssignActionRequest NERM.CreateAutoAssignActionRequest
    if err := json.Unmarshal(createautoassignactionrequest, &createAutoAssignActionRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.WorkflowActionsAPI.CreateAutoAssignAction(context.Background()).CreateAutoAssignActionRequest(createAutoAssignActionRequest).Execute()
	  //resp, r, err := apiClient.NERM.WorkflowActionsAPI.CreateAutoAssignAction(context.Background()).CreateAutoAssignActionRequest(createAutoAssignActionRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `WorkflowActionsAPI.CreateAutoAssignAction``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAutoAssignAction`: CreateApprovalAction200Response
    fmt.Fprintf(os.Stdout, "Response from `WorkflowActionsAPI.CreateAutoAssignAction`: %v\n", resp)
}
```

[[Back to top]](#)

## create-batch-update-action
Create a batch update action
Create a batch update action

[API Spec](https://developer.sailpoint.com/docs/api/NERM/create-batch-update-action)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBatchUpdateActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBatchUpdateActionRequest** | [**CreateBatchUpdateActionRequest**](../models/create-batch-update-action-request) |  | 

### Return type

[**CreateApprovalAction200Response**](../models/create-approval-action200-response)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
  "encoding/json"
    NERM "github.com/sailpoint-oss/golang-sdk/v2/api_nerm"
	sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    createbatchupdateactionrequest := []byte(``) // CreateBatchUpdateActionRequest | 

    var createBatchUpdateActionRequest NERM.CreateBatchUpdateActionRequest
    if err := json.Unmarshal(createbatchupdateactionrequest, &createBatchUpdateActionRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.WorkflowActionsAPI.CreateBatchUpdateAction(context.Background()).CreateBatchUpdateActionRequest(createBatchUpdateActionRequest).Execute()
	  //resp, r, err := apiClient.NERM.WorkflowActionsAPI.CreateBatchUpdateAction(context.Background()).CreateBatchUpdateActionRequest(createBatchUpdateActionRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `WorkflowActionsAPI.CreateBatchUpdateAction``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBatchUpdateAction`: CreateApprovalAction200Response
    fmt.Fprintf(os.Stdout, "Response from `WorkflowActionsAPI.CreateBatchUpdateAction`: %v\n", resp)
}
```

[[Back to top]](#)

## create-close-session-action
Create a close session action
Create a close session action

[API Spec](https://developer.sailpoint.com/docs/api/NERM/create-close-session-action)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCloseSessionActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCloseSessionActionRequest** | [**CreateCloseSessionActionRequest**](../models/create-close-session-action-request) |  | 

### Return type

[**CreateApprovalAction200Response**](../models/create-approval-action200-response)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
  "encoding/json"
    NERM "github.com/sailpoint-oss/golang-sdk/v2/api_nerm"
	sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    createclosesessionactionrequest := []byte(``) // CreateCloseSessionActionRequest | 

    var createCloseSessionActionRequest NERM.CreateCloseSessionActionRequest
    if err := json.Unmarshal(createclosesessionactionrequest, &createCloseSessionActionRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.WorkflowActionsAPI.CreateCloseSessionAction(context.Background()).CreateCloseSessionActionRequest(createCloseSessionActionRequest).Execute()
	  //resp, r, err := apiClient.NERM.WorkflowActionsAPI.CreateCloseSessionAction(context.Background()).CreateCloseSessionActionRequest(createCloseSessionActionRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `WorkflowActionsAPI.CreateCloseSessionAction``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCloseSessionAction`: CreateApprovalAction200Response
    fmt.Fprintf(os.Stdout, "Response from `WorkflowActionsAPI.CreateCloseSessionAction`: %v\n", resp)
}
```

[[Back to top]](#)

## create-contributors-action
Create a contributors action
Create a contributors action. The following supporting objects will need to be created after this action is created (which are tied together via workflow_action_id) - workflow_action_performers, workflow_action_roles, workflow_action_performer_notification_email. These supporting objects must be created for this action to be complete (APIs for these supporting objects not yet implemented, use UI).

[API Spec](https://developer.sailpoint.com/docs/api/NERM/create-contributors-action)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateContributorsActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createContributorsActionRequest** | [**CreateContributorsActionRequest**](../models/create-contributors-action-request) |  | 

### Return type

[**CreateApprovalAction200Response**](../models/create-approval-action200-response)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
  "encoding/json"
    NERM "github.com/sailpoint-oss/golang-sdk/v2/api_nerm"
	sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    createcontributorsactionrequest := []byte(``) // CreateContributorsActionRequest | 

    var createContributorsActionRequest NERM.CreateContributorsActionRequest
    if err := json.Unmarshal(createcontributorsactionrequest, &createContributorsActionRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.WorkflowActionsAPI.CreateContributorsAction(context.Background()).CreateContributorsActionRequest(createContributorsActionRequest).Execute()
	  //resp, r, err := apiClient.NERM.WorkflowActionsAPI.CreateContributorsAction(context.Background()).CreateContributorsActionRequest(createContributorsActionRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `WorkflowActionsAPI.CreateContributorsAction``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateContributorsAction`: CreateApprovalAction200Response
    fmt.Fprintf(os.Stdout, "Response from `WorkflowActionsAPI.CreateContributorsAction`: %v\n", resp)
}
```

[[Back to top]](#)

## create-create-profile-action
Create a create profile action
Create a create profile action

[API Spec](https://developer.sailpoint.com/docs/api/NERM/create-create-profile-action)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCreateProfileActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCreateProfileActionRequest** | [**CreateCreateProfileActionRequest**](../models/create-create-profile-action-request) |  | 

### Return type

[**CreateApprovalAction200Response**](../models/create-approval-action200-response)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
  "encoding/json"
    NERM "github.com/sailpoint-oss/golang-sdk/v2/api_nerm"
	sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    createcreateprofileactionrequest := []byte(``) // CreateCreateProfileActionRequest | 

    var createCreateProfileActionRequest NERM.CreateCreateProfileActionRequest
    if err := json.Unmarshal(createcreateprofileactionrequest, &createCreateProfileActionRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.WorkflowActionsAPI.CreateCreateProfileAction(context.Background()).CreateCreateProfileActionRequest(createCreateProfileActionRequest).Execute()
	  //resp, r, err := apiClient.NERM.WorkflowActionsAPI.CreateCreateProfileAction(context.Background()).CreateCreateProfileActionRequest(createCreateProfileActionRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `WorkflowActionsAPI.CreateCreateProfileAction``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCreateProfileAction`: CreateApprovalAction200Response
    fmt.Fprintf(os.Stdout, "Response from `WorkflowActionsAPI.CreateCreateProfileAction`: %v\n", resp)
}
```

[[Back to top]](#)

## create-duplicate-prevention-action
Create a duplicate prevention action
Create a duplicate prevention action. The following supporting objects will need to be created after this action is created (which are tied together via workflow_action_id) - duplicatation_prevention_attributes, workflow_action_performers, workflow_action_roles, workflow_action_performer_notification_email. These supporting objects must be created for this action to be complete (APIs for these supporting objects not yet implemented, use UI).

[API Spec](https://developer.sailpoint.com/docs/api/NERM/create-duplicate-prevention-action)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDuplicatePreventionActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDuplicatePreventionActionRequest** | [**CreateDuplicatePreventionActionRequest**](../models/create-duplicate-prevention-action-request) |  | 

### Return type

[**CreateApprovalAction200Response**](../models/create-approval-action200-response)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
  "encoding/json"
    NERM "github.com/sailpoint-oss/golang-sdk/v2/api_nerm"
	sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    createduplicatepreventionactionrequest := []byte(``) // CreateDuplicatePreventionActionRequest | 

    var createDuplicatePreventionActionRequest NERM.CreateDuplicatePreventionActionRequest
    if err := json.Unmarshal(createduplicatepreventionactionrequest, &createDuplicatePreventionActionRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.WorkflowActionsAPI.CreateDuplicatePreventionAction(context.Background()).CreateDuplicatePreventionActionRequest(createDuplicatePreventionActionRequest).Execute()
	  //resp, r, err := apiClient.NERM.WorkflowActionsAPI.CreateDuplicatePreventionAction(context.Background()).CreateDuplicatePreventionActionRequest(createDuplicatePreventionActionRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `WorkflowActionsAPI.CreateDuplicatePreventionAction``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDuplicatePreventionAction`: CreateApprovalAction200Response
    fmt.Fprintf(os.Stdout, "Response from `WorkflowActionsAPI.CreateDuplicatePreventionAction`: %v\n", resp)
}
```

[[Back to top]](#)

## create-email-verification-action
Create an email verification action
Create an email verification action

[API Spec](https://developer.sailpoint.com/docs/api/NERM/create-email-verification-action)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEmailVerificationActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createEmailVerificationActionRequest** | [**CreateEmailVerificationActionRequest**](../models/create-email-verification-action-request) |  | 

### Return type

[**CreateApprovalAction200Response**](../models/create-approval-action200-response)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
  "encoding/json"
    NERM "github.com/sailpoint-oss/golang-sdk/v2/api_nerm"
	sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    createemailverificationactionrequest := []byte(``) // CreateEmailVerificationActionRequest | 

    var createEmailVerificationActionRequest NERM.CreateEmailVerificationActionRequest
    if err := json.Unmarshal(createemailverificationactionrequest, &createEmailVerificationActionRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.WorkflowActionsAPI.CreateEmailVerificationAction(context.Background()).CreateEmailVerificationActionRequest(createEmailVerificationActionRequest).Execute()
	  //resp, r, err := apiClient.NERM.WorkflowActionsAPI.CreateEmailVerificationAction(context.Background()).CreateEmailVerificationActionRequest(createEmailVerificationActionRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `WorkflowActionsAPI.CreateEmailVerificationAction``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEmailVerificationAction`: CreateApprovalAction200Response
    fmt.Fprintf(os.Stdout, "Response from `WorkflowActionsAPI.CreateEmailVerificationAction`: %v\n", resp)
}
```

[[Back to top]](#)

## create-fulfillment-action
Create a fulfillment action
Create a fulfillment action. The following supporting objects will need to be created after this action is created (which are tied together via workflow_action_id) - workflow_action_performers, workflow_action_roles, workflow_action_performer_notification_email. These supporting objects must be created for this action to be complete (APIs for these supporting objects not yet implemented, use UI).

[API Spec](https://developer.sailpoint.com/docs/api/NERM/create-fulfillment-action)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFulfillmentActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFulfillmentActionRequest** | [**CreateFulfillmentActionRequest**](../models/create-fulfillment-action-request) |  | 

### Return type

[**CreateApprovalAction200Response**](../models/create-approval-action200-response)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
  "encoding/json"
    NERM "github.com/sailpoint-oss/golang-sdk/v2/api_nerm"
	sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    createfulfillmentactionrequest := []byte(``) // CreateFulfillmentActionRequest | 

    var createFulfillmentActionRequest NERM.CreateFulfillmentActionRequest
    if err := json.Unmarshal(createfulfillmentactionrequest, &createFulfillmentActionRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.WorkflowActionsAPI.CreateFulfillmentAction(context.Background()).CreateFulfillmentActionRequest(createFulfillmentActionRequest).Execute()
	  //resp, r, err := apiClient.NERM.WorkflowActionsAPI.CreateFulfillmentAction(context.Background()).CreateFulfillmentActionRequest(createFulfillmentActionRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `WorkflowActionsAPI.CreateFulfillmentAction``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFulfillmentAction`: CreateApprovalAction200Response
    fmt.Fprintf(os.Stdout, "Response from `WorkflowActionsAPI.CreateFulfillmentAction`: %v\n", resp)
}
```

[[Back to top]](#)

## create-identity-proofing-action
Create an identity proofing action
Create an identity proofing action. The following supporting objects will need to be created after this action is created (which are tied together via workflow_action_id) - identity_proofing_action_configuration, identity_proofing_action_mappings. These supporting objects must be created for this action to be complete (APIs for these supporting objects not yet implemented, use UI).

[API Spec](https://developer.sailpoint.com/docs/api/NERM/create-identity-proofing-action)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIdentityProofingActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createIdentityProofingActionRequest** | [**CreateIdentityProofingActionRequest**](../models/create-identity-proofing-action-request) |  | 

### Return type

[**CreateApprovalAction200Response**](../models/create-approval-action200-response)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
  "encoding/json"
    NERM "github.com/sailpoint-oss/golang-sdk/v2/api_nerm"
	sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    createidentityproofingactionrequest := []byte(``) // CreateIdentityProofingActionRequest | 

    var createIdentityProofingActionRequest NERM.CreateIdentityProofingActionRequest
    if err := json.Unmarshal(createidentityproofingactionrequest, &createIdentityProofingActionRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.WorkflowActionsAPI.CreateIdentityProofingAction(context.Background()).CreateIdentityProofingActionRequest(createIdentityProofingActionRequest).Execute()
	  //resp, r, err := apiClient.NERM.WorkflowActionsAPI.CreateIdentityProofingAction(context.Background()).CreateIdentityProofingActionRequest(createIdentityProofingActionRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `WorkflowActionsAPI.CreateIdentityProofingAction``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIdentityProofingAction`: CreateApprovalAction200Response
    fmt.Fprintf(os.Stdout, "Response from `WorkflowActionsAPI.CreateIdentityProofingAction`: %v\n", resp)
}
```

[[Back to top]](#)

## create-invitation-action
Create an invitation action
Create an invitation action. The following supporting objects will need to be created after this action is created (which are tied together via workflow_action_id) - workflow_action_pause_action. These supporting objects must be created for this action to be complete (APIs for these supporting objects not yet implemented, use UI).

[API Spec](https://developer.sailpoint.com/docs/api/NERM/create-invitation-action)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInvitationActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createInvitationActionRequest** | [**CreateInvitationActionRequest**](../models/create-invitation-action-request) |  | 

### Return type

[**CreateApprovalAction200Response**](../models/create-approval-action200-response)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
  "encoding/json"
    NERM "github.com/sailpoint-oss/golang-sdk/v2/api_nerm"
	sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    createinvitationactionrequest := []byte(``) // CreateInvitationActionRequest | 

    var createInvitationActionRequest NERM.CreateInvitationActionRequest
    if err := json.Unmarshal(createinvitationactionrequest, &createInvitationActionRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.WorkflowActionsAPI.CreateInvitationAction(context.Background()).CreateInvitationActionRequest(createInvitationActionRequest).Execute()
	  //resp, r, err := apiClient.NERM.WorkflowActionsAPI.CreateInvitationAction(context.Background()).CreateInvitationActionRequest(createInvitationActionRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `WorkflowActionsAPI.CreateInvitationAction``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateInvitationAction`: CreateApprovalAction200Response
    fmt.Fprintf(os.Stdout, "Response from `WorkflowActionsAPI.CreateInvitationAction`: %v\n", resp)
}
```

[[Back to top]](#)

## create-ldap-action
Create a ldap action
Create a ldap action. The following supporting objects will need to be created after this action is created (which are tied together via workflow_action_id) - workflow_action_performers, workflow_action_roles. These supporting objects must be created for this action to be complete (APIs for these supporting objects not yet implemented, use UI).

[API Spec](https://developer.sailpoint.com/docs/api/NERM/create-ldap-action)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLdapActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createLdapActionRequest** | [**CreateLdapActionRequest**](../models/create-ldap-action-request) |  | 

### Return type

[**CreateApprovalAction200Response**](../models/create-approval-action200-response)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
  "encoding/json"
    NERM "github.com/sailpoint-oss/golang-sdk/v2/api_nerm"
	sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    createldapactionrequest := []byte(``) // CreateLdapActionRequest | 

    var createLdapActionRequest NERM.CreateLdapActionRequest
    if err := json.Unmarshal(createldapactionrequest, &createLdapActionRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.WorkflowActionsAPI.CreateLdapAction(context.Background()).CreateLdapActionRequest(createLdapActionRequest).Execute()
	  //resp, r, err := apiClient.NERM.WorkflowActionsAPI.CreateLdapAction(context.Background()).CreateLdapActionRequest(createLdapActionRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `WorkflowActionsAPI.CreateLdapAction``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLdapAction`: CreateApprovalAction200Response
    fmt.Fprintf(os.Stdout, "Response from `WorkflowActionsAPI.CreateLdapAction`: %v\n", resp)
}
```

[[Back to top]](#)

## create-notification-action
Create a notification action
Create a notification action

[API Spec](https://developer.sailpoint.com/docs/api/NERM/create-notification-action)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNotificationActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNotificationActionRequest** | [**CreateNotificationActionRequest**](../models/create-notification-action-request) |  | 

### Return type

[**CreateApprovalAction200Response**](../models/create-approval-action200-response)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
  "encoding/json"
    NERM "github.com/sailpoint-oss/golang-sdk/v2/api_nerm"
	sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    createnotificationactionrequest := []byte(``) // CreateNotificationActionRequest | 

    var createNotificationActionRequest NERM.CreateNotificationActionRequest
    if err := json.Unmarshal(createnotificationactionrequest, &createNotificationActionRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.WorkflowActionsAPI.CreateNotificationAction(context.Background()).CreateNotificationActionRequest(createNotificationActionRequest).Execute()
	  //resp, r, err := apiClient.NERM.WorkflowActionsAPI.CreateNotificationAction(context.Background()).CreateNotificationActionRequest(createNotificationActionRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `WorkflowActionsAPI.CreateNotificationAction``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNotificationAction`: CreateApprovalAction200Response
    fmt.Fprintf(os.Stdout, "Response from `WorkflowActionsAPI.CreateNotificationAction`: %v\n", resp)
}
```

[[Back to top]](#)

## create-password-reset-action
Create a password reset action
Create a password reset action

[API Spec](https://developer.sailpoint.com/docs/api/NERM/create-password-reset-action)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePasswordResetActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPasswordResetActionRequest** | [**CreatePasswordResetActionRequest**](../models/create-password-reset-action-request) |  | 

### Return type

[**CreateApprovalAction200Response**](../models/create-approval-action200-response)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
  "encoding/json"
    NERM "github.com/sailpoint-oss/golang-sdk/v2/api_nerm"
	sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    createpasswordresetactionrequest := []byte(``) // CreatePasswordResetActionRequest | 

    var createPasswordResetActionRequest NERM.CreatePasswordResetActionRequest
    if err := json.Unmarshal(createpasswordresetactionrequest, &createPasswordResetActionRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.WorkflowActionsAPI.CreatePasswordResetAction(context.Background()).CreatePasswordResetActionRequest(createPasswordResetActionRequest).Execute()
	  //resp, r, err := apiClient.NERM.WorkflowActionsAPI.CreatePasswordResetAction(context.Background()).CreatePasswordResetActionRequest(createPasswordResetActionRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `WorkflowActionsAPI.CreatePasswordResetAction``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePasswordResetAction`: CreateApprovalAction200Response
    fmt.Fprintf(os.Stdout, "Response from `WorkflowActionsAPI.CreatePasswordResetAction`: %v\n", resp)
}
```

[[Back to top]](#)

## create-profile-check-action
Create a profile check action
Create a profile check action

[API Spec](https://developer.sailpoint.com/docs/api/NERM/create-profile-check-action)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProfileCheckActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createProfileCheckActionRequest** | [**CreateProfileCheckActionRequest**](../models/create-profile-check-action-request) |  | 

### Return type

[**CreateApprovalAction200Response**](../models/create-approval-action200-response)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
  "encoding/json"
    NERM "github.com/sailpoint-oss/golang-sdk/v2/api_nerm"
	sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    createprofilecheckactionrequest := []byte(``) // CreateProfileCheckActionRequest | 

    var createProfileCheckActionRequest NERM.CreateProfileCheckActionRequest
    if err := json.Unmarshal(createprofilecheckactionrequest, &createProfileCheckActionRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.WorkflowActionsAPI.CreateProfileCheckAction(context.Background()).CreateProfileCheckActionRequest(createProfileCheckActionRequest).Execute()
	  //resp, r, err := apiClient.NERM.WorkflowActionsAPI.CreateProfileCheckAction(context.Background()).CreateProfileCheckActionRequest(createProfileCheckActionRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `WorkflowActionsAPI.CreateProfileCheckAction``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProfileCheckAction`: CreateApprovalAction200Response
    fmt.Fprintf(os.Stdout, "Response from `WorkflowActionsAPI.CreateProfileCheckAction`: %v\n", resp)
}
```

[[Back to top]](#)

## create-profile-select-action
Create a profile select action
Create a profile select action

[API Spec](https://developer.sailpoint.com/docs/api/NERM/create-profile-select-action)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProfileSelectActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createProfileSelectActionRequest** | [**CreateProfileSelectActionRequest**](../models/create-profile-select-action-request) |  | 

### Return type

[**CreateApprovalAction200Response**](../models/create-approval-action200-response)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
  "encoding/json"
    NERM "github.com/sailpoint-oss/golang-sdk/v2/api_nerm"
	sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    createprofileselectactionrequest := []byte(``) // CreateProfileSelectActionRequest | 

    var createProfileSelectActionRequest NERM.CreateProfileSelectActionRequest
    if err := json.Unmarshal(createprofileselectactionrequest, &createProfileSelectActionRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.WorkflowActionsAPI.CreateProfileSelectAction(context.Background()).CreateProfileSelectActionRequest(createProfileSelectActionRequest).Execute()
	  //resp, r, err := apiClient.NERM.WorkflowActionsAPI.CreateProfileSelectAction(context.Background()).CreateProfileSelectActionRequest(createProfileSelectActionRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `WorkflowActionsAPI.CreateProfileSelectAction``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProfileSelectAction`: CreateApprovalAction200Response
    fmt.Fprintf(os.Stdout, "Response from `WorkflowActionsAPI.CreateProfileSelectAction`: %v\n", resp)
}
```

[[Back to top]](#)

## create-request-action
Create a request action
Create a request action

[API Spec](https://developer.sailpoint.com/docs/api/NERM/create-request-action)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequestActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRequestActionRequest** | [**CreateRequestActionRequest**](../models/create-request-action-request) |  | 

### Return type

[**CreateApprovalAction200Response**](../models/create-approval-action200-response)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
  "encoding/json"
    NERM "github.com/sailpoint-oss/golang-sdk/v2/api_nerm"
	sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    createrequestactionrequest := []byte(``) // CreateRequestActionRequest | 

    var createRequestActionRequest NERM.CreateRequestActionRequest
    if err := json.Unmarshal(createrequestactionrequest, &createRequestActionRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.WorkflowActionsAPI.CreateRequestAction(context.Background()).CreateRequestActionRequest(createRequestActionRequest).Execute()
	  //resp, r, err := apiClient.NERM.WorkflowActionsAPI.CreateRequestAction(context.Background()).CreateRequestActionRequest(createRequestActionRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `WorkflowActionsAPI.CreateRequestAction``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRequestAction`: CreateApprovalAction200Response
    fmt.Fprintf(os.Stdout, "Response from `WorkflowActionsAPI.CreateRequestAction`: %v\n", resp)
}
```

[[Back to top]](#)

## create-rest-api-action
Create a REST API action
Create a REST API action. The following supporting objects will need to be created after this action is created (which are tied together via workflow_action_id) - rest_api_action_configuration, api_configuration_attributes. These supporting objects must be created for this action to be complete (APIs for these supporting objects not yet implemented, use UI).

[API Spec](https://developer.sailpoint.com/docs/api/NERM/create-rest-api-action)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRestApiActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRestApiActionRequest** | [**CreateRestApiActionRequest**](../models/create-rest-api-action-request) |  | 

### Return type

[**CreateApprovalAction200Response**](../models/create-approval-action200-response)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
  "encoding/json"
    NERM "github.com/sailpoint-oss/golang-sdk/v2/api_nerm"
	sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    createrestapiactionrequest := []byte(``) // CreateRestApiActionRequest | 

    var createRestApiActionRequest NERM.CreateRestApiActionRequest
    if err := json.Unmarshal(createrestapiactionrequest, &createRestApiActionRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.WorkflowActionsAPI.CreateRestApiAction(context.Background()).CreateRestApiActionRequest(createRestApiActionRequest).Execute()
	  //resp, r, err := apiClient.NERM.WorkflowActionsAPI.CreateRestApiAction(context.Background()).CreateRestApiActionRequest(createRestApiActionRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `WorkflowActionsAPI.CreateRestApiAction``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRestApiAction`: CreateApprovalAction200Response
    fmt.Fprintf(os.Stdout, "Response from `WorkflowActionsAPI.CreateRestApiAction`: %v\n", resp)
}
```

[[Back to top]](#)

## create-review-action
Create a review action
Create a review action. The following supporting objects will need to be created after this action is created (which are tied together via workflow_action_id) - workflow_action_performer_notification_email. These supporting objects must be created for this action to be complete (APIs for these supporting objects not yet implemented, use UI).

[API Spec](https://developer.sailpoint.com/docs/api/NERM/create-review-action)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateReviewActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createReviewActionRequest** | [**CreateReviewActionRequest**](../models/create-review-action-request) |  | 

### Return type

[**CreateApprovalAction200Response**](../models/create-approval-action200-response)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
  "encoding/json"
    NERM "github.com/sailpoint-oss/golang-sdk/v2/api_nerm"
	sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    createreviewactionrequest := []byte(``) // CreateReviewActionRequest | 

    var createReviewActionRequest NERM.CreateReviewActionRequest
    if err := json.Unmarshal(createreviewactionrequest, &createReviewActionRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.WorkflowActionsAPI.CreateReviewAction(context.Background()).CreateReviewActionRequest(createReviewActionRequest).Execute()
	  //resp, r, err := apiClient.NERM.WorkflowActionsAPI.CreateReviewAction(context.Background()).CreateReviewActionRequest(createReviewActionRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `WorkflowActionsAPI.CreateReviewAction``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateReviewAction`: CreateApprovalAction200Response
    fmt.Fprintf(os.Stdout, "Response from `WorkflowActionsAPI.CreateReviewAction`: %v\n", resp)
}
```

[[Back to top]](#)

## create-run-workflow-action
Create a run workflow action
Create a run workflow action. The following supporting objects will need to be created after this action is created (which are tied together via workflow_action_id) - configuration_profile_attribute. These supporting objects must be created for this action to be complete (APIs for these supporting objects not yet implemented, use UI).

[API Spec](https://developer.sailpoint.com/docs/api/NERM/create-run-workflow-action)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRunWorkflowActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRunWorkflowActionRequest** | [**CreateRunWorkflowActionRequest**](../models/create-run-workflow-action-request) |  | 

### Return type

[**CreateApprovalAction200Response**](../models/create-approval-action200-response)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
  "encoding/json"
    NERM "github.com/sailpoint-oss/golang-sdk/v2/api_nerm"
	sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    createrunworkflowactionrequest := []byte(``) // CreateRunWorkflowActionRequest | 

    var createRunWorkflowActionRequest NERM.CreateRunWorkflowActionRequest
    if err := json.Unmarshal(createrunworkflowactionrequest, &createRunWorkflowActionRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.WorkflowActionsAPI.CreateRunWorkflowAction(context.Background()).CreateRunWorkflowActionRequest(createRunWorkflowActionRequest).Execute()
	  //resp, r, err := apiClient.NERM.WorkflowActionsAPI.CreateRunWorkflowAction(context.Background()).CreateRunWorkflowActionRequest(createRunWorkflowActionRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `WorkflowActionsAPI.CreateRunWorkflowAction``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRunWorkflowAction`: CreateApprovalAction200Response
    fmt.Fprintf(os.Stdout, "Response from `WorkflowActionsAPI.CreateRunWorkflowAction`: %v\n", resp)
}
```

[[Back to top]](#)

## create-set-attributes-action
Create a set attributes action
Create a set attributes action. The following supporting objects will need to be created after this action is created (which are tied together via workflow_action_id) - workflow_action_set_attributes. These supporting objects must be created for this action to be complete (APIs for these supporting objects not yet implemented, use UI).

[API Spec](https://developer.sailpoint.com/docs/api/NERM/create-set-attributes-action)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSetAttributesActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSetAttributesActionRequest** | [**CreateSetAttributesActionRequest**](../models/create-set-attributes-action-request) |  | 

### Return type

[**CreateApprovalAction200Response**](../models/create-approval-action200-response)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
  "encoding/json"
    NERM "github.com/sailpoint-oss/golang-sdk/v2/api_nerm"
	sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    createsetattributesactionrequest := []byte(``) // CreateSetAttributesActionRequest | 

    var createSetAttributesActionRequest NERM.CreateSetAttributesActionRequest
    if err := json.Unmarshal(createsetattributesactionrequest, &createSetAttributesActionRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.WorkflowActionsAPI.CreateSetAttributesAction(context.Background()).CreateSetAttributesActionRequest(createSetAttributesActionRequest).Execute()
	  //resp, r, err := apiClient.NERM.WorkflowActionsAPI.CreateSetAttributesAction(context.Background()).CreateSetAttributesActionRequest(createSetAttributesActionRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `WorkflowActionsAPI.CreateSetAttributesAction``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSetAttributesAction`: CreateApprovalAction200Response
    fmt.Fprintf(os.Stdout, "Response from `WorkflowActionsAPI.CreateSetAttributesAction`: %v\n", resp)
}
```

[[Back to top]](#)

## create-set-security-question-action
Create set security question action
Create a set security question action

[API Spec](https://developer.sailpoint.com/docs/api/NERM/create-set-security-question-action)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSetSecurityQuestionActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSetSecurityQuestionActionRequest** | [**CreateSetSecurityQuestionActionRequest**](../models/create-set-security-question-action-request) |  | 

### Return type

[**CreateApprovalAction200Response**](../models/create-approval-action200-response)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
  "encoding/json"
    NERM "github.com/sailpoint-oss/golang-sdk/v2/api_nerm"
	sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    createsetsecurityquestionactionrequest := []byte(``) // CreateSetSecurityQuestionActionRequest | 

    var createSetSecurityQuestionActionRequest NERM.CreateSetSecurityQuestionActionRequest
    if err := json.Unmarshal(createsetsecurityquestionactionrequest, &createSetSecurityQuestionActionRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.WorkflowActionsAPI.CreateSetSecurityQuestionAction(context.Background()).CreateSetSecurityQuestionActionRequest(createSetSecurityQuestionActionRequest).Execute()
	  //resp, r, err := apiClient.NERM.WorkflowActionsAPI.CreateSetSecurityQuestionAction(context.Background()).CreateSetSecurityQuestionActionRequest(createSetSecurityQuestionActionRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `WorkflowActionsAPI.CreateSetSecurityQuestionAction``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSetSecurityQuestionAction`: CreateApprovalAction200Response
    fmt.Fprintf(os.Stdout, "Response from `WorkflowActionsAPI.CreateSetSecurityQuestionAction`: %v\n", resp)
}
```

[[Back to top]](#)

## create-soap-api-action
Create a SOAP API action
Create a SOAP API action. The following supporting objects will need to be created after this action is created (which are tied together via workflow_action_id) - soap_api_action_configuration, api_configuration_attributes. These supporting objects must be created for this action to be complete (APIs for these supporting objects not yet implemented, use UI).

[API Spec](https://developer.sailpoint.com/docs/api/NERM/create-soap-api-action)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSoapApiActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSoapApiActionRequest** | [**CreateSoapApiActionRequest**](../models/create-soap-api-action-request) |  | 

### Return type

[**CreateApprovalAction200Response**](../models/create-approval-action200-response)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
  "encoding/json"
    NERM "github.com/sailpoint-oss/golang-sdk/v2/api_nerm"
	sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    createsoapapiactionrequest := []byte(``) // CreateSoapApiActionRequest | 

    var createSoapApiActionRequest NERM.CreateSoapApiActionRequest
    if err := json.Unmarshal(createsoapapiactionrequest, &createSoapApiActionRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.WorkflowActionsAPI.CreateSoapApiAction(context.Background()).CreateSoapApiActionRequest(createSoapApiActionRequest).Execute()
	  //resp, r, err := apiClient.NERM.WorkflowActionsAPI.CreateSoapApiAction(context.Background()).CreateSoapApiActionRequest(createSoapApiActionRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `WorkflowActionsAPI.CreateSoapApiAction``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSoapApiAction`: CreateApprovalAction200Response
    fmt.Fprintf(os.Stdout, "Response from `WorkflowActionsAPI.CreateSoapApiAction`: %v\n", resp)
}
```

[[Back to top]](#)

## create-status-change-action
Create a status change action
Create a status change action

[API Spec](https://developer.sailpoint.com/docs/api/NERM/create-status-change-action)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateStatusChangeActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createStatusChangeActionRequest** | [**CreateStatusChangeActionRequest**](../models/create-status-change-action-request) |  | 

### Return type

[**CreateApprovalAction200Response**](../models/create-approval-action200-response)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
  "encoding/json"
    NERM "github.com/sailpoint-oss/golang-sdk/v2/api_nerm"
	sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    createstatuschangeactionrequest := []byte(``) // CreateStatusChangeActionRequest | 

    var createStatusChangeActionRequest NERM.CreateStatusChangeActionRequest
    if err := json.Unmarshal(createstatuschangeactionrequest, &createStatusChangeActionRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.WorkflowActionsAPI.CreateStatusChangeAction(context.Background()).CreateStatusChangeActionRequest(createStatusChangeActionRequest).Execute()
	  //resp, r, err := apiClient.NERM.WorkflowActionsAPI.CreateStatusChangeAction(context.Background()).CreateStatusChangeActionRequest(createStatusChangeActionRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `WorkflowActionsAPI.CreateStatusChangeAction``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateStatusChangeAction`: CreateApprovalAction200Response
    fmt.Fprintf(os.Stdout, "Response from `WorkflowActionsAPI.CreateStatusChangeAction`: %v\n", resp)
}
```

[[Back to top]](#)

## create-unassign-action
Create an unassign action
Create an unassign action. The following supporting objects will need to be created after this action is created (which are tied together via workflow_action_id) - workflow_action_roles. These supporting objects must be created for this action to be complete (APIs for these supporting objects not yet implemented, use UI).

[API Spec](https://developer.sailpoint.com/docs/api/NERM/create-unassign-action)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUnassignActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createUnassignActionRequest** | [**CreateUnassignActionRequest**](../models/create-unassign-action-request) |  | 

### Return type

[**CreateApprovalAction200Response**](../models/create-approval-action200-response)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
  "encoding/json"
    NERM "github.com/sailpoint-oss/golang-sdk/v2/api_nerm"
	sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    createunassignactionrequest := []byte(``) // CreateUnassignActionRequest | 

    var createUnassignActionRequest NERM.CreateUnassignActionRequest
    if err := json.Unmarshal(createunassignactionrequest, &createUnassignActionRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.WorkflowActionsAPI.CreateUnassignAction(context.Background()).CreateUnassignActionRequest(createUnassignActionRequest).Execute()
	  //resp, r, err := apiClient.NERM.WorkflowActionsAPI.CreateUnassignAction(context.Background()).CreateUnassignActionRequest(createUnassignActionRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `WorkflowActionsAPI.CreateUnassignAction``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUnassignAction`: CreateApprovalAction200Response
    fmt.Fprintf(os.Stdout, "Response from `WorkflowActionsAPI.CreateUnassignAction`: %v\n", resp)
}
```

[[Back to top]](#)

## create-update-profile-action
Create an update profile action
Create an update profile action

[API Spec](https://developer.sailpoint.com/docs/api/NERM/create-update-profile-action)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUpdateProfileActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createUpdateProfileActionRequest** | [**CreateUpdateProfileActionRequest**](../models/create-update-profile-action-request) |  | 

### Return type

[**CreateApprovalAction200Response**](../models/create-approval-action200-response)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
  "encoding/json"
    NERM "github.com/sailpoint-oss/golang-sdk/v2/api_nerm"
	sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    createupdateprofileactionrequest := []byte(``) // CreateUpdateProfileActionRequest | 

    var createUpdateProfileActionRequest NERM.CreateUpdateProfileActionRequest
    if err := json.Unmarshal(createupdateprofileactionrequest, &createUpdateProfileActionRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.WorkflowActionsAPI.CreateUpdateProfileAction(context.Background()).CreateUpdateProfileActionRequest(createUpdateProfileActionRequest).Execute()
	  //resp, r, err := apiClient.NERM.WorkflowActionsAPI.CreateUpdateProfileAction(context.Background()).CreateUpdateProfileActionRequest(createUpdateProfileActionRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `WorkflowActionsAPI.CreateUpdateProfileAction``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUpdateProfileAction`: CreateApprovalAction200Response
    fmt.Fprintf(os.Stdout, "Response from `WorkflowActionsAPI.CreateUpdateProfileAction`: %v\n", resp)
}
```

[[Back to top]](#)

## create-username-password-action
Create a username password action
Create a username password action

[API Spec](https://developer.sailpoint.com/docs/api/NERM/create-username-password-action)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUsernamePasswordActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createUsernamePasswordActionRequest** | [**CreateUsernamePasswordActionRequest**](../models/create-username-password-action-request) |  | 

### Return type

[**CreateApprovalAction200Response**](../models/create-approval-action200-response)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
  "encoding/json"
    NERM "github.com/sailpoint-oss/golang-sdk/v2/api_nerm"
	sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    createusernamepasswordactionrequest := []byte(``) // CreateUsernamePasswordActionRequest | 

    var createUsernamePasswordActionRequest NERM.CreateUsernamePasswordActionRequest
    if err := json.Unmarshal(createusernamepasswordactionrequest, &createUsernamePasswordActionRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.WorkflowActionsAPI.CreateUsernamePasswordAction(context.Background()).CreateUsernamePasswordActionRequest(createUsernamePasswordActionRequest).Execute()
	  //resp, r, err := apiClient.NERM.WorkflowActionsAPI.CreateUsernamePasswordAction(context.Background()).CreateUsernamePasswordActionRequest(createUsernamePasswordActionRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `WorkflowActionsAPI.CreateUsernamePasswordAction``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUsernamePasswordAction`: CreateApprovalAction200Response
    fmt.Fprintf(os.Stdout, "Response from `WorkflowActionsAPI.CreateUsernamePasswordAction`: %v\n", resp)
}
```

[[Back to top]](#)

## get-workflow-actions
Get Workflow Actions
This endpoint can retrieve workflow actions

[API Spec](https://developer.sailpoint.com/docs/api/NERM/get-workflow-actions)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowActionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowId** | **string** | Workflow ID for filtering | 

### Return type

[**GetWorkflowActions200Response**](../models/get-workflow-actions200-response)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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
    workflowId := `bba9cfb2-96c1-4acb-ac79-a21732527265` // string | Workflow ID for filtering (optional) # string | Workflow ID for filtering (optional)

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.WorkflowActionsAPI.GetWorkflowActions(context.Background()).Execute()
	  //resp, r, err := apiClient.NERM.WorkflowActionsAPI.GetWorkflowActions(context.Background()).WorkflowId(workflowId).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `WorkflowActionsAPI.GetWorkflowActions``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkflowActions`: GetWorkflowActions200Response
    fmt.Fprintf(os.Stdout, "Response from `WorkflowActionsAPI.GetWorkflowActions`: %v\n", resp)
}
```

[[Back to top]](#)

