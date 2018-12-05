# \DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteWorkspace**](DefaultApi.md#DeleteWorkspace) | **Delete** /workspaces/{workspaceName} | 
[**DeleteWorkspaces**](DefaultApi.md#DeleteWorkspaces) | **Delete** /workspaces | 
[**GetWorkspace**](DefaultApi.md#GetWorkspace) | **Get** /workspaces/{workspaceName} | Retrieve a layer group
[**GetWorkspaces**](DefaultApi.md#GetWorkspaces) | **Get** /workspaces | Get a list of workspaces
[**PostWorkspace**](DefaultApi.md#PostWorkspace) | **Post** /workspaces/{workspaceName} | 
[**PostWorkspaces**](DefaultApi.md#PostWorkspaces) | **Post** /workspaces | add a new workspace to GeoServer
[**PutWorkspace**](DefaultApi.md#PutWorkspace) | **Put** /workspaces/{workspaceName} | Update a workspace
[**PutWorkspaces**](DefaultApi.md#PutWorkspaces) | **Put** /workspaces | 


# **DeleteWorkspace**
> DeleteWorkspace(ctx, workspaceName, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspaceName** | **string**| name of workspace | 
 **optional** | ***DeleteWorkspaceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteWorkspaceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **recurse** | **optional.Bool**| delete workspace contents (default false) | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteWorkspaces**
> DeleteWorkspaces(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWorkspace**
> WorkspaceResponse GetWorkspace(ctx, workspaceName)
Retrieve a layer group

Retrieves a single workspace definition. Use the \"Accept:\" header to specify format or append an extension to the endpoint (example \"/workspaces/{workspace}.xml\" for XML).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspaceName** | **string**| the name of the workspace to fetch | 

### Return type

[**WorkspaceResponse**](WorkspaceResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWorkspaces**
> WorkspacesResponse GetWorkspaces(ctx, )
Get a list of workspaces

Displays a list of all workspaces on the server. Use the \"Accept:\" header to specify format or append an extension to the endpoint (example \"/workspaces.xml\" for XML)

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**WorkspacesResponse**](WorkspacesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/html, application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostWorkspace**
> PostWorkspace(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostWorkspaces**
> string PostWorkspaces(ctx, workspace, optional)
add a new workspace to GeoServer

Adds a new workspace to the server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | [**Workspace**](Workspace.md)| The layer group body information to upload. | 
 **optional** | ***PostWorkspacesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PostWorkspacesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **default_** | **optional.Bool**| New workspace will be the used as the default. Allowed values are true or false,  The default value is false. | [default to false]

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: text/html, application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutWorkspace**
> PutWorkspace(ctx, workspaceName, workspace)
Update a workspace

takes the body of the post and modifies the workspace from it.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspaceName** | **string**| name of workspace | 
  **workspace** | [**Workspace**](Workspace.md)| The layer group body information to upload. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutWorkspaces**
> PutWorkspaces(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

