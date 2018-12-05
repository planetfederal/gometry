# \DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LayersDelete**](DefaultApi.md#LayersDelete) | **Delete** /layers | 
[**LayersGet**](DefaultApi.md#LayersGet) | **Get** /layers | Get a list of layers
[**LayersNameDelete**](DefaultApi.md#LayersNameDelete) | **Delete** /layers/{layerName} | Delete layer
[**LayersNameGet**](DefaultApi.md#LayersNameGet) | **Get** /layers/{layerName} | Retrieve a layer
[**LayersNamePost**](DefaultApi.md#LayersNamePost) | **Post** /layers/{layerName} | 
[**LayersNamePut**](DefaultApi.md#LayersNamePut) | **Put** /layers/{layerName} | Modify a layer.
[**LayersNameWorkspaceDelete**](DefaultApi.md#LayersNameWorkspaceDelete) | **Delete** /workspaces/{workspaceName}/layers/{layerName} | Delete layer
[**LayersNameWorkspaceGet**](DefaultApi.md#LayersNameWorkspaceGet) | **Get** /workspaces/{workspaceName}/layers/{layerName} | Retrieve a layer
[**LayersNameWorkspacePost**](DefaultApi.md#LayersNameWorkspacePost) | **Post** /workspaces/{workspaceName}/layers/{layerName} | 
[**LayersNameWorkspacePut**](DefaultApi.md#LayersNameWorkspacePut) | **Put** /workspaces/{workspaceName}/layers/{layerName} | Modify a layer.
[**LayersPost**](DefaultApi.md#LayersPost) | **Post** /layers | 
[**LayersPut**](DefaultApi.md#LayersPut) | **Put** /layers | 
[**LayersWorkspaceDelete**](DefaultApi.md#LayersWorkspaceDelete) | **Delete** /workspaces/{workspaceName}/layers | 
[**LayersWorkspaceGet**](DefaultApi.md#LayersWorkspaceGet) | **Get** /workspaces/{workspaceName}/layers | Get a list of layers in a workspace.
[**LayersWorkspacePost**](DefaultApi.md#LayersWorkspacePost) | **Post** /workspaces/{workspaceName}/layers | 
[**LayersWorkspacePut**](DefaultApi.md#LayersWorkspacePut) | **Put** /workspaces/{workspaceName}/layers | 


# **LayersDelete**
> LayersDelete(ctx, )


Invalid.

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

# **LayersGet**
> Layers LayersGet(ctx, )
Get a list of layers

Displays a list of all layers on the server. You must use the \"Accept:\" header to specify format or append an extension to the endpoint (example \"/layers.xml\" for XML)

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Layers**](Layers.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/html, application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LayersNameDelete**
> LayersNameDelete(ctx, layerName, optional)
Delete layer

Deletes a layer from the server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **layerName** | **string**| The name of the layer to delete. | 
 **optional** | ***LayersNameDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LayersNameDeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **recurse** | **optional.Bool**| Recursively removes the layer from all layer groups which reference it. If this results in an empty layer group, also delete the layer group. Allowed values for this parameter are true or false. The default value is false. A request with &#39;recurse&#x3D;false&#39; will fail if any layer groups reference the layer. | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LayersNameGet**
> Layer LayersNameGet(ctx, layerName)
Retrieve a layer

Retrieves a single layer definition. Use the \"Accept:\" header to specify format or append an extension to the endpoint (example \"/layers/{layer}.xml\" for XML).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **layerName** | **string**| The name of the layer to retrieve. | 

### Return type

[**Layer**](Layer.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LayersNamePost**
> LayersNamePost(ctx, )


Invalid. To create a new layer, instead POST to one of `/workspaces/{workspaceName}/coveragestores/{coveragestoreName}/coverages`, `/workspaces/{workspaceName}/datastores/{datastoreName}/featuretypes`, `/workspaces/{workspaceName}/wmsstores/{wmsstoreName}/wmslayers`, or `/workspaces/{workspaceName}/wmtsstores/{wmststoreName}/wmtslayers`

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

# **LayersNamePut**
> LayersNamePut(ctx, layerName, layer)
Modify a layer.

Modifies an existing layer on the server. Use the \"Accept:\" header to specify format or append an extension to the endpoint (example \"/layers/{layer}.xml\" for XML).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **layerName** | **string**| The name of the layer to modify. | 
  **layer** | [**Layer**](Layer.md)| The updated layer definition. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LayersNameWorkspaceDelete**
> LayersNameWorkspaceDelete(ctx, workspaceName, layerName, optional)
Delete layer

Deletes a layer from the server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspaceName** | [**map[string]interface{}**](.md)| The name of the workspace the layer is in. | 
  **layerName** | **string**| The name of the layer to delete. | 
 **optional** | ***LayersNameWorkspaceDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LayersNameWorkspaceDeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **recurse** | **optional.Bool**| Recursively removes the layer from all layer groups which reference it. If this results in an empty layer group, also delete the layer group. Allowed values for this parameter are true or false. The default value is false. A request with &#39;recurse&#x3D;false&#39; will fail if any layer groups reference the layer. | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LayersNameWorkspaceGet**
> Layer LayersNameWorkspaceGet(ctx, workspaceName, layerName)
Retrieve a layer

Retrieves a single layer definition. Use the \"Accept:\" header to specify format or append an extension to the endpoint (example \"/layers/{layer}.xml\" for XML).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspaceName** | **string**| The name of the workspace the layer is in. | 
  **layerName** | **string**| The name of the layer to retrieve. | 

### Return type

[**Layer**](Layer.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LayersNameWorkspacePost**
> LayersNameWorkspacePost(ctx, )


Invalid. To create a new layer, instead POST to one of `/workspaces/{workspaceName}/coveragestores/{coveragestoreName}/coverages`, `/workspaces/{workspaceName}/datastores/{datastoreName}/featuretypes`, `/workspaces/{workspaceName}/wmsstores/{wmsstoreName}/wmslayers`, or `/workspaces/{workspaceName}/wmtsstores/{wmststoreName}/wmtslayers`

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

# **LayersNameWorkspacePut**
> LayersNameWorkspacePut(ctx, workspaceName, layerName, layer)
Modify a layer.

Modifies an existing layer on the server. Use the \"Accept:\" header to specify format or append an extension to the endpoint (example \"/layers/{layer}.xml\" for XML).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspaceName** | [**map[string]interface{}**](.md)| The name of the workspace the layer is in. | 
  **layerName** | **string**| The name of the layer to modify. | 
  **layer** | [**Layer**](Layer.md)| The updated layer definition. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LayersPost**
> LayersPost(ctx, )


Invalid. To create a new layer, instead POST to one of `/workspaces/{workspaceName}/coveragestores/{coveragestoreName}/coverages`, `/workspaces/{workspaceName}/datastores/{datastoreName}/featuretypes`, `/workspaces/{workspaceName}/wmsstores/{wmsstoreName}/wmslayers`, or `/workspaces/{workspaceName}/wmtsstores/{wmststoreName}/wmtslayers`

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

# **LayersPut**
> LayersPut(ctx, )


Invalid. To edit a layer, use PUT on an individual layer instead.

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

# **LayersWorkspaceDelete**
> LayersWorkspaceDelete(ctx, )


Invalid.

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

# **LayersWorkspaceGet**
> Layers LayersWorkspaceGet(ctx, workspaceName)
Get a list of layers in a workspace.

Displays a list of all layers in the provided workspace. You must use the \"Accept:\" header to specify format or append an extension to the endpoint (example \"/layers.xml\" for XML)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspaceName** | **string**| The name of the workspace to list layers in | 

### Return type

[**Layers**](Layers.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/html, application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LayersWorkspacePost**
> LayersWorkspacePost(ctx, )


Invalid. To create a new layer, instead POST to one of `/workspaces/{workspaceName}/coveragestores/{coveragestoreName}/coverages`, `/workspaces/{workspaceName}/datastores/{datastoreName}/featuretypes`, `/workspaces/{workspaceName}/wmsstores/{wmsstoreName}/wmslayers`, or `/workspaces/{workspaceName}/wmtsstores/{wmststoreName}/wmtslayers`

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

# **LayersWorkspacePut**
> LayersWorkspacePut(ctx, )


Invalid. To edit a layer, use PUT on an individual layer instead.

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

