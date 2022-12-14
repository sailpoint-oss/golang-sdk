# Go API client for sailpointv2sdk

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 2.0.0
- Package version: 0.1.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sailpointv2sdk "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sailpointv2sdk.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sailpointv2sdk.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), sailpointv2sdk.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sailpointv2sdk.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://sailpoint.api.identitynow.com/v2*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AccessProfilesApi* | [**CreateAccessProfile**](docs/AccessProfilesApi.md#createaccessprofile) | **Post** /access-profiles | Creates a new access profile.
*AccessProfilesApi* | [**DeleteAccessProfile**](docs/AccessProfilesApi.md#deleteaccessprofile) | **Delete** /access-profiles/{accessProfileId} | Deletes an access profile.
*AccessProfilesApi* | [**GetAccessProfile**](docs/AccessProfilesApi.md#getaccessprofile) | **Get** /access-profiles/{accessProfileId} | Retrieves an access profile.
*AccessProfilesApi* | [**ListAccessProfileEntitlements**](docs/AccessProfilesApi.md#listaccessprofileentitlements) | **Get** /access-profiles/{accessProfileId}/entitlements | Lists all entitlements of an access profile.
*AccessProfilesApi* | [**ListAccessProfiles**](docs/AccessProfilesApi.md#listaccessprofiles) | **Get** /access-profiles | Lists the access profiles.
*AccessProfilesApi* | [**PatchAccessProfile**](docs/AccessProfilesApi.md#patchaccessprofile) | **Patch** /access-profiles/{accessProfileId} | Updates one or more access profile attributes.
*AccessProfilesApi* | [**PutAccessProfile**](docs/AccessProfilesApi.md#putaccessprofile) | **Put** /access-profiles/{accessProfileId} | Updates an existing access profile.
*AccessRequestsApi* | [**ListAccessRequests**](docs/AccessRequestsApi.md#listaccessrequests) | **Get** /access-requests | Lists the access request for an identity.
*AccountsApi* | [**CreateAccount**](docs/AccountsApi.md#createaccount) | **Post** /accounts | Creates a new account on a flat-file source.
*AccountsApi* | [**DeleteAccount**](docs/AccountsApi.md#deleteaccount) | **Delete** /accounts/{id} | Deletes an existing account from a flat-file source.
*AccountsApi* | [**GetAccount**](docs/AccountsApi.md#getaccount) | **Get** /accounts/{id} | Retrieves the Account by Id.
*AccountsApi* | [**GetAccounts**](docs/AccountsApi.md#getaccounts) | **Get** /accounts | Retrieves list of Accounts for a given source.
*AccountsApi* | [**PatchAccount**](docs/AccountsApi.md#patchaccount) | **Patch** /accounts/{id} | Updates an existing account from a flat-file source.
*AccountsApi* | [**PutAccount**](docs/AccountsApi.md#putaccount) | **Put** /accounts/{id} | Updates an existing account from a flat-file source by replacing all values.
*ApprovalsApi* | [**ApproveAccessRequest**](docs/ApprovalsApi.md#approveaccessrequest) | **Post** /approvals/{approvalId}/approve-request | Approves an access request.
*ApprovalsApi* | [**ForwardAccessRequestApproval**](docs/ApprovalsApi.md#forwardaccessrequestapproval) | **Post** /approvals/{approvalId}/forward | Forwards an access request approval.
*ApprovalsApi* | [**ListAccessRequestApproval**](docs/ApprovalsApi.md#listaccessrequestapproval) | **Get** /approvals | Lists the approvals.
*ApprovalsApi* | [**RejectAccessRequest**](docs/ApprovalsApi.md#rejectaccessrequest) | **Post** /approvals/{approvalId}/reject-request | Rejects an access request.
*GovernanceGroupsApi* | [**BulkDeleteWorkGroups**](docs/GovernanceGroupsApi.md#bulkdeleteworkgroups) | **Post** /workgroups/bulk-delete | Bulk delete work groups
*GovernanceGroupsApi* | [**CreateWorkgroup**](docs/GovernanceGroupsApi.md#createworkgroup) | **Post** /workgroups | Create Work Group
*GovernanceGroupsApi* | [**DeleteWorkgroup**](docs/GovernanceGroupsApi.md#deleteworkgroup) | **Delete** /v2/workgroups/{workgroupId} | Delete Work Group By Id
*GovernanceGroupsApi* | [**GetWorkgroup**](docs/GovernanceGroupsApi.md#getworkgroup) | **Get** /v2/workgroups/{workgroupId} | Get Work Group By Id
*GovernanceGroupsApi* | [**ListWorkgroupConnections**](docs/GovernanceGroupsApi.md#listworkgroupconnections) | **Get** /workgroups/{workgroupId}/connections | List Work Group Connections
*GovernanceGroupsApi* | [**ListWorkgroupMembers**](docs/GovernanceGroupsApi.md#listworkgroupmembers) | **Get** /workgroups/{workgroupId}/members | List Work Group Members
*GovernanceGroupsApi* | [**ListWorkgroups**](docs/GovernanceGroupsApi.md#listworkgroups) | **Get** /workgroups | List Work Groups
*GovernanceGroupsApi* | [**ModifyWorkgroupMembers**](docs/GovernanceGroupsApi.md#modifyworkgroupmembers) | **Post** /workgroups/{workgroupId}/members | Modify Work Group Members
*GovernanceGroupsApi* | [**UpdateWorkgroup**](docs/GovernanceGroupsApi.md#updateworkgroup) | **Patch** /v2/workgroups/{workgroupId} | Update Work Group By Id
*IdentitiesApi* | [**CreateIdentity**](docs/IdentitiesApi.md#createidentity) | **Post** /identities | Creates a new identity.
*IdentitiesApi* | [**CreateLauncher**](docs/IdentitiesApi.md#createlauncher) | **Post** /identities/{identityIdOrAlias}/launchers | Creates a new launcher.
*IdentitiesApi* | [**DeleteIdentity**](docs/IdentitiesApi.md#deleteidentity) | **Delete** /identities/{identityIdOrAlias} | Deletes an identity.
*IdentitiesApi* | [**DeleteLauncher**](docs/IdentitiesApi.md#deletelauncher) | **Delete** /identities/{identityIdOrAlias}/launchers/{launcherId} | Deletes a launcher.
*IdentitiesApi* | [**GetIdentity**](docs/IdentitiesApi.md#getidentity) | **Get** /identities/{identityIdOrAlias} | Retrieves the identity by ID or alias.
*IdentitiesApi* | [**ListApprovals**](docs/IdentitiesApi.md#listapprovals) | **Get** /identities/{identityIdOrAlias}/approvals | Lists the approvals.
*IdentitiesApi* | [**ListApps**](docs/IdentitiesApi.md#listapps) | **Get** /identities/{identityIdOrAlias}/apps | Lists available apps.
*IdentitiesApi* | [**ListIdentities**](docs/IdentitiesApi.md#listidentities) | **Get** /identities | Retrieves the identities.
*IdentitiesApi* | [**ListLaunchers**](docs/IdentitiesApi.md#listlaunchers) | **Get** /identities/{identityIdOrAlias}/launchers | Lists the launchers.
*IdentitiesApi* | [**LockIdentities**](docs/IdentitiesApi.md#lockidentities) | **Post** /identities/bulk-lock | Locks one or more identities.
*IdentitiesApi* | [**UpdateIdentity**](docs/IdentitiesApi.md#updateidentity) | **Patch** /identities/{identityIdOrAlias} | Updates one or more identity attributes.
*LaunchersApi* | [**GetLauncher**](docs/LaunchersApi.md#getlauncher) | **Get** /launchers/{launcherId} | Retrieves the details of the launcher.
*LaunchersApi* | [**LogLauncherClick**](docs/LaunchersApi.md#loglauncherclick) | **Post** /launchers/{launcherId}/click | Records a launcher click.
*LaunchersApi* | [**UpdateLauncher**](docs/LaunchersApi.md#updatelauncher) | **Patch** /launchers/{launcherId} | Updates one or more attributes of a launcher.
*OrgApi* | [**GetOrgSettings**](docs/OrgApi.md#getorgsettings) | **Get** /org | Retrieves your org settings.
*OrgApi* | [**UpdateOrgSettings**](docs/OrgApi.md#updateorgsettings) | **Patch** /org | Updates one or more org attributes.
*ProvisioningActivitiesApi* | [**GetProvisioningActivity**](docs/ProvisioningActivitiesApi.md#getprovisioningactivity) | **Get** /provisioning-activities/{provisioningActivityId} | Retrieves a provisioning activity.
*ProvisioningActivitiesApi* | [**ListProvisioningActivities**](docs/ProvisioningActivitiesApi.md#listprovisioningactivities) | **Get** /provisioning-activities | Lists the provisioning activities.
*SearchApi* | [**EntitlementSearch**](docs/SearchApi.md#entitlementsearch) | **Get** /search/entitlements | Searches and retrieves the entitlements.
*SearchApi* | [**EntitlementSearchExport**](docs/SearchApi.md#entitlementsearchexport) | **Get** /search/entitlements/runExport | Runs csv results export job for a given search for entitlements.
*SearchApi* | [**EntitlementSearchExportPost**](docs/SearchApi.md#entitlementsearchexportpost) | **Post** /search/entitlements/runExport | Runs csv results export job for a given search for entitlements.
*SearchApi* | [**EntitlementSearchPost**](docs/SearchApi.md#entitlementsearchpost) | **Post** /search/entitlements | Searches and retrieves the entitlements.
*SearchApi* | [**ExportIdentitySearch**](docs/SearchApi.md#exportidentitysearch) | **Get** /search/identities/runExport | Runs csv results export job for a given search for identities.
*SearchApi* | [**ExportSearchCsv**](docs/SearchApi.md#exportsearchcsv) | **Get** /search/runExport | Runs csv results export job for a given search query.
*SearchApi* | [**ExportSearchCsvPost**](docs/SearchApi.md#exportsearchcsvpost) | **Post** /search/runExport | Runs csv results export job for a given search query.
*SearchApi* | [**GetSearchIndexMapping**](docs/SearchApi.md#getsearchindexmapping) | **Get** /search/{index}/mappings | Retrieves the mappings and operators for the search service for the given index path.
*SearchApi* | [**GetSearchMapping**](docs/SearchApi.md#getsearchmapping) | **Get** /search/mappings | Retrieves the mappings and operators for the search service.
*SearchApi* | [**RunEventSearch**](docs/SearchApi.md#runeventsearch) | **Get** /search/events | Searches and retrieves the events.
*SearchApi* | [**RunEventSearchPost**](docs/SearchApi.md#runeventsearchpost) | **Post** /search/events | Searches and retrieves the events.
*SearchApi* | [**RunIdentitySearch**](docs/SearchApi.md#runidentitysearch) | **Post** /search/identities/runExport | Runs csv results export job for a given search for identities.
*SearchApi* | [**RunSearch**](docs/SearchApi.md#runsearch) | **Get** /search | Searches and retrieves the types specified (currently identity, entitlement, and event).
*SearchApi* | [**RunSearchPost**](docs/SearchApi.md#runsearchpost) | **Post** /search | Searches and retrieves the types specified (current identity, entitlement, and event).
*SearchApi* | [**SearchIdentities**](docs/SearchApi.md#searchidentities) | **Get** /search/identities | Searches and retrieves the identities.
*SearchApi* | [**SearchIdentitiesPost**](docs/SearchApi.md#searchidentitiespost) | **Post** /search/identities | Searches and retrieves the identities.
*TaskResultsApi* | [**GetBackgroundTaskResults**](docs/TaskResultsApi.md#getbackgroundtaskresults) | **Get** /task-results/{taskResultIdOrName} | Retrieve Result of Background Task


## Documentation For Models

 - [AccessProfile](docs/AccessProfile.md)
 - [AccessProfileCreateEto](docs/AccessProfileCreateEto.md)
 - [AccessRequest](docs/AccessRequest.md)
 - [Account](docs/Account.md)
 - [AccountApplication](docs/AccountApplication.md)
 - [AccountAttributes](docs/AccountAttributes.md)
 - [AccountIdentity](docs/AccountIdentity.md)
 - [AccountMeta](docs/AccountMeta.md)
 - [App](docs/App.md)
 - [AppAccessMethodsInner](docs/AppAccessMethodsInner.md)
 - [Approval](docs/Approval.md)
 - [ApprovalConfigEto](docs/ApprovalConfigEto.md)
 - [BulkDeleteWorkGroups200Response](docs/BulkDeleteWorkGroups200Response.md)
 - [BulkDeleteWorkGroupsRequest](docs/BulkDeleteWorkGroupsRequest.md)
 - [CommentEto](docs/CommentEto.md)
 - [CreateWorkgroupRequest](docs/CreateWorkgroupRequest.md)
 - [CreateWorkgroupRequestOwner](docs/CreateWorkgroupRequestOwner.md)
 - [DynamicSchemaEto](docs/DynamicSchemaEto.md)
 - [Entitlement](docs/Entitlement.md)
 - [ExceptionObject](docs/ExceptionObject.md)
 - [ExceptionObjectErrorsInner](docs/ExceptionObjectErrorsInner.md)
 - [ExportSearchCsvPostRequest](docs/ExportSearchCsvPostRequest.md)
 - [ExportSearchCsvPostRequestMatch](docs/ExportSearchCsvPostRequestMatch.md)
 - [ForwardApprovalEto](docs/ForwardApprovalEto.md)
 - [IdentityEto](docs/IdentityEto.md)
 - [IdentityEtoIdentityFlags](docs/IdentityEtoIdentityFlags.md)
 - [IdentityV2](docs/IdentityV2.md)
 - [Launcher](docs/Launcher.md)
 - [LauncherAccount](docs/LauncherAccount.md)
 - [LauncherEto](docs/LauncherEto.md)
 - [LauncherLink](docs/LauncherLink.md)
 - [Mapping](docs/Mapping.md)
 - [ModifyWorkgroupMembersRequest](docs/ModifyWorkgroupMembersRequest.md)
 - [MultiStatusObject](docs/MultiStatusObject.md)
 - [MultiStatusObjectFailInner](docs/MultiStatusObjectFailInner.md)
 - [Notification](docs/Notification.md)
 - [NotificationThresholds](docs/NotificationThresholds.md)
 - [Org](docs/Org.md)
 - [OrgEto](docs/OrgEto.md)
 - [OrgSystemNotificationConfig](docs/OrgSystemNotificationConfig.md)
 - [ProvisioningActivity](docs/ProvisioningActivity.md)
 - [RequestSummary](docs/RequestSummary.md)
 - [RunSearch200Response](docs/RunSearch200Response.md)
 - [SearchEntitlement](docs/SearchEntitlement.md)
 - [SearchEvent](docs/SearchEvent.md)
 - [SearchEventActor](docs/SearchEventActor.md)
 - [SearchEventAttributes](docs/SearchEventAttributes.md)
 - [SearchEventAttributesReviewer](docs/SearchEventAttributesReviewer.md)
 - [SearchEventAttributesSSO](docs/SearchEventAttributesSSO.md)
 - [SearchIdentity](docs/SearchIdentity.md)
 - [SearchIdentityAccessInner](docs/SearchIdentityAccessInner.md)
 - [SearchIdentityAccessInnerOwner](docs/SearchIdentityAccessInnerOwner.md)
 - [SearchIdentityAccountsInner](docs/SearchIdentityAccountsInner.md)
 - [SearchIdentityAccountsInnerSource](docs/SearchIdentityAccountsInnerSource.md)
 - [SearchIdentityAppsInner](docs/SearchIdentityAppsInner.md)
 - [SearchIdentityAppsInnerAccount](docs/SearchIdentityAppsInnerAccount.md)
 - [SearchIdentityAppsInnerSource](docs/SearchIdentityAppsInnerSource.md)
 - [SearchIdentityAttributes](docs/SearchIdentityAttributes.md)
 - [SearchIdentityManager](docs/SearchIdentityManager.md)
 - [SearchIdentityProcessingDetails](docs/SearchIdentityProcessingDetails.md)
 - [TaskResult](docs/TaskResult.md)
 - [TaskResultMessagesInner](docs/TaskResultMessagesInner.md)
 - [TaskResultReturnsInner](docs/TaskResultReturnsInner.md)
 - [Workgroup](docs/Workgroup.md)
 - [WorkgroupConnection](docs/WorkgroupConnection.md)
 - [WorkgroupMember](docs/WorkgroupMember.md)
 - [WorkgroupOwner](docs/WorkgroupOwner.md)


## Documentation For Authorization



### bearerAuth


- **Type**: OAuth
- **Flow**: accessCode
- **Authorization URL**: https://tenant.identitynow.com/oauth/authorize
- **Scopes**: 
 - **idn:task-definition:read**: Task Definition read access
 - **idn:task-definition:write**: Task Definition write access
 - **idn:task-management:read**: Task Management read access (TaskStatus)
 - **idn:task-management:write**: Task Management write access (TaskInvocation)
 - **idn:service-desk-integration:read**: Service Desk integration read access
 - **idn:service-desk-integration:write**: Service Desk integration write access
 - **idn:managed-cluster:read**: ManagedCluster read access
 - **idn:managed-cluster-log-config:read**: ManagedCluster log configuration read access
 - **idn:managed-cluster-log-config:write**: ManagedCluster log configuration write access
 - **idn:managed-cluster:upgrade**: ManagedCluster client version upgrade access

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



