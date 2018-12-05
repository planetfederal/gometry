# \DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDataStoreUpload**](DefaultApi.md#DeleteDataStoreUpload) | **Delete** /workspaces/{workspaceName}/datastores/{storeName}/{method}.{format} | 
[**DeleteDatastore**](DefaultApi.md#DeleteDatastore) | **Delete** /workspaces/{workspaceName}/datastores/{storeName} | Delete data store
[**Deletedatastores**](DefaultApi.md#Deletedatastores) | **Delete** /workspaces/{workspaceName}/datastores | 
[**GetDataStore**](DefaultApi.md#GetDataStore) | **Get** /workspaces/{workspaceName}/datastores/{storeName} | Retrieve a particular data store from a workspace
[**GetDataStoreUpload**](DefaultApi.md#GetDataStoreUpload) | **Get** /workspaces/{workspaceName}/datastores/{storeName}/{method}.{format} | 
[**GetDatastores**](DefaultApi.md#GetDatastores) | **Get** /workspaces/{workspaceName}/datastores | Get a list of data stores
[**PostDataStoreUpload**](DefaultApi.md#PostDataStoreUpload) | **Post** /workspaces/{workspaceName}/datastores/{storeName}/{method}.{format} | 
[**PostDatastore**](DefaultApi.md#PostDatastore) | **Post** /workspaces/{workspaceName}/datastores/{storeName} | 
[**PostDatastores**](DefaultApi.md#PostDatastores) | **Post** /workspaces/{workspaceName}/datastores | Create a new data store
[**PutDataStoreUpload**](DefaultApi.md#PutDataStoreUpload) | **Put** /workspaces/{workspaceName}/datastores/{storeName}/{method}.{format} | Uploads files to the data store, creating it if necessary
[**PutDatastore**](DefaultApi.md#PutDatastore) | **Put** /workspaces/{workspaceName}/datastores/{storeName} | Modify a data store.
[**Putdatastores**](DefaultApi.md#Putdatastores) | **Put** /workspaces/{workspaceName}/datastores | 


# **DeleteDataStoreUpload**
> DeleteDataStoreUpload(ctx, )


Invalid, only used for uploads

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

# **DeleteDatastore**
> DeleteDatastore(ctx, workspaceName, storeName, optional)
Delete data store

Deletes a data store from the server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspaceName** | **string**| The name of the worskpace containing the data store. | 
  **storeName** | **string**| The name of the data store to delete. | 
 **optional** | ***DeleteDatastoreOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteDatastoreOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **recurse** | **optional.Bool**| The recurse controls recursive deletion. When set to true all resources contained in the store are also removed. The default value is \&quot;false\&quot;. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Deletedatastores**
> Deletedatastores(ctx, )


Invalid. Use /datastores/{datastore} instead.

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

# **GetDataStore**
> Datastore GetDataStore(ctx, workspaceName, storeName, optional)
Retrieve a particular data store from a workspace

Controls a particular data store in a given workspace. Use the \"Accept:\" header to specify format or append an extension to the endpoint (example \"/datastores/{datastore}.xml\" for XML).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspaceName** | **string**| The name of the worskpace containing the data store. | 
  **storeName** | **string**| The name of the data store to retrieve. | 
 **optional** | ***GetDataStoreOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetDataStoreOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **quietOnNotFound** | **optional.Bool**| The quietOnNotFound parameter avoids logging an exception when the data store is not present. Note that 404 status code will still be returned. | 

### Return type

[**Datastore**](datastore.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDataStoreUpload**
> GetDataStoreUpload(ctx, workspaceName, storeName, method, format)


Deprecated. Retrieve the underlying files for the data store as a zip file with MIME type application/zip

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspaceName** | **string**| The name of the worskpace containing the data store. | 
  **storeName** | **string**| The name of the store to be retrieved | 
  **method** | **string**| The upload method. Can be \&quot;url\&quot;, \&quot;file\&quot;, \&quot;external\&quot;. Unused for GET | 
  **format** | **string**| The type of source data store (e.g., \&quot;shp\&quot;). Unused for GET | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDatastores**
> []map[string]interface{} GetDatastores(ctx, workspaceName)
Get a list of data stores

List all data stores in workspace ws. Use the \"Accept:\" header to specify format or append an extension to the endpoint (example \"/datastores.xml\" for XML)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspaceName** | **string**| The name of the worskpace containing the data stores. | 

### Return type

[**[]map[string]interface{}**](map[string]interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostDataStoreUpload**
> PostDataStoreUpload(ctx, )


Invalid, use PUT for uploads

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

# **PostDatastore**
> PostDatastore(ctx, )


Invalid. Use PUT to edit a data store definition, or POST with /datastore to add a new definition.

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

# **PostDatastores**
> string PostDatastores(ctx, workspaceName, datastore)
Create a new data store

Adds a new data store to the workspace.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspaceName** | **string**| The name of the worskpace containing the data stores. | 
  **datastore** | [**Datastore**](Datastore.md)| The data store body information to upload.  The contents of the connection parameters will differ depending on the type of data store being added.  - GeoPackage    Examples:   - application/xml:      &#x60;&#x60;&#x60;     &lt;dataStore&gt;       &lt;name&gt;nyc&lt;/name&gt;       &lt;connectionParameters&gt;         &lt;database&gt;file:///path/to/nyc.gpkg&lt;/database&gt;         &lt;dbtype&gt;geopkg&lt;/dbtype&gt;       &lt;/connectionParameters&gt;     &lt;/dataStore&gt;     &#x60;&#x60;&#x60;    - application/json:      &#x60;&#x60;&#x60;     {       \&quot;dataStore\&quot;: {         \&quot;name\&quot;: \&quot;nyc\&quot;,         \&quot;connectionParameters\&quot;: {           \&quot;entry\&quot;: [             {\&quot;@key\&quot;:\&quot;database\&quot;,\&quot;$\&quot;:\&quot;file:///path/to/nyc.gpkg\&quot;},             {\&quot;@key\&quot;:\&quot;dbtype\&quot;,\&quot;$\&quot;:\&quot;geopkg\&quot;}           ]         }       }     }     &#x60;&#x60;&#x60;    Connection Parameters:    | key | description | level | type | required | default |   | --- | ----------- | ----- | ---- | -------- | ------- |   | Primary key metadata table | The optional table containing primary key structure and sequence associations. Can be expressed as &#39;schema.name&#39; or just &#39;name&#39; | user | String | False | &#x60; &#x60; |   | Callback factory | Name of JDBCReaderCallbackFactory to enable on the data store | user | String | False | &#x60; &#x60; |   | Evictor tests per run | number of connections checked by the idle connection evictor for each of its runs (defaults to 3) | user | Integer | False | &#x60;3&#x60; |   | database | Database | user | File | True | &#x60; &#x60; |   | Batch insert size | Number of records inserted in the same batch (default, 1). For optimal performance, set to 100. | user | Integer | False | &#x60;1&#x60; |   | fetch size | number of records read with each iteraction with the dbms | user | Integer | False | &#x60;1000&#x60; |   | Connection timeout | number of seconds the connection pool will wait before timing out attempting to get a new connection (default, 20 seconds) | user | Integer | False | &#x60;20&#x60; |   | namespace | Namespace prefix | user | String | False | &#x60; &#x60; |   | max connections | maximum number of open connections | user | Integer | False | &#x60;10&#x60; |   | Test while idle | Periodically test the connections are still valid also while idle in the pool | user | Boolean | False | &#x60;True&#x60; |   | Max connection idle time | number of seconds a connection needs to stay idle for the evictor to consider closing it | user | Integer | False | &#x60;300&#x60; |   | Session startup SQL | SQL statement executed when the connection is grabbed from the pool | user | String | False | &#x60; &#x60; |   | validate connections | check connection is alive before using it | user | Boolean | False | &#x60;True&#x60; |   | dbtype | Type | program | String | True | &#x60;geopkg&#x60; |   | passwd | password used to login | user | String | False | &#x60; &#x60; |   | Expose primary keys | Expose primary key columns as attributes of the feature type | user | Boolean | False | &#x60;False&#x60; |   | min connections | minimum number of pooled connection | user | Integer | False | &#x60;1&#x60; |   | Evictor run periodicity | number of seconds between idle object evitor runs (default, 300 seconds) | user | Integer | False | &#x60;300&#x60; |   | Session close-up SQL | SQL statement executed when the connection is released to the pool | user | String | False | &#x60; &#x60; |   | user | user name to login as | user | String | False | &#x60; &#x60; |  - PostGIS    Examples:   - application/xml:      &#x60;&#x60;&#x60;     &lt;dataStore&gt;       &lt;name&gt;nyc&lt;/name&gt;       &lt;connectionParameters&gt;         &lt;host&gt;localhost&lt;/host&gt;         &lt;port&gt;5432&lt;/port&gt;         &lt;database&gt;nyc&lt;/database&gt;         &lt;user&gt;bob&lt;/user&gt;         &lt;passwd&gt;postgres&lt;/passwd&gt;         &lt;dbtype&gt;postgis&lt;/dbtype&gt;       &lt;/connectionParameters&gt;     &lt;/dataStore&gt;     &#x60;&#x60;&#x60;    - application/json:      &#x60;&#x60;&#x60;     {       \&quot;dataStore\&quot;: {         \&quot;name\&quot;: \&quot;nyc\&quot;,         \&quot;connectionParameters\&quot;: {           \&quot;entry\&quot;: [             {\&quot;@key\&quot;:\&quot;host\&quot;,\&quot;$\&quot;:\&quot;localhost\&quot;},             {\&quot;@key\&quot;:\&quot;port\&quot;,\&quot;$\&quot;:\&quot;5432\&quot;},             {\&quot;@key\&quot;:\&quot;database\&quot;,\&quot;$\&quot;:\&quot;nyc\&quot;},             {\&quot;@key\&quot;:\&quot;user\&quot;,\&quot;$\&quot;:\&quot;bob\&quot;},             {\&quot;@key\&quot;:\&quot;passwd\&quot;,\&quot;$\&quot;:\&quot;postgres\&quot;},             {\&quot;@key\&quot;:\&quot;dbtype\&quot;,\&quot;$\&quot;:\&quot;postgis\&quot;}           ]         }       }     }     &#x60;&#x60;&#x60;    Connection Parameters:    | key | description | level | type | required | default |   | --- | ----------- | ----- | ---- | -------- | ------- |   | Connection timeout | number of seconds the connection pool will wait before timing out attempting to get a new connection (default, 20 seconds) | user | Integer | False | &#x60;20&#x60; |   | validate connections | check connection is alive before using it | user | Boolean | False | &#x60;True&#x60; |   | port | Port | user | Integer | True | &#x60;5432&#x60; |   | Primary key metadata table | The optional table containing primary key structure and sequence associations. Can be expressed as &#39;schema.name&#39; or just &#39;name&#39; | user | String | False | &#x60; &#x60; |   | Support on the fly geometry simplification | When enabled, operations such as map rendering will pass a hint that will enable the usage of ST_Simplify | user | Boolean | False | &#x60;True&#x60; |   | create database | Creates the database if it does not exist yet | advanced | Boolean | False | &#x60;False&#x60; |   | create database params | Extra specifications appeneded to the CREATE DATABASE command | advanced | String | False | &#x60;&#x60; |   | dbtype | Type | program | String | True | &#x60;postgis&#x60; |   | Batch insert size | Number of records inserted in the same batch (default, 1). For optimal performance, set to 100. | user | Integer | False | &#x60;1&#x60; |   | namespace | Namespace prefix | user | String | False | &#x60; &#x60; |   | Max connection idle time | number of seconds a connection needs to stay idle for the evictor to consider closing it | user | Integer | False | &#x60;300&#x60; |   | Session startup SQL | SQL statement executed when the connection is grabbed from the pool | user | String | False | &#x60; &#x60; |   | Expose primary keys | Expose primary key columns as attributes of the feature type | user | Boolean | False | &#x60;False&#x60; |   | min connections | minimum number of pooled connection | user | Integer | False | &#x60;1&#x60; |   | Max open prepared statements | Maximum number of prepared statements kept open and cached for each connection in the pool. Set to 0 to have unbounded caching, to -1 to disable caching | user | Integer | False | &#x60;50&#x60; |   | Callback factory | Name of JDBCReaderCallbackFactory to enable on the data store | user | String | False | &#x60; &#x60; |   | passwd | password used to login | user | String | False | &#x60; &#x60; |   | encode functions | set to true to have a set of filter functions be translated directly in SQL. Due to differences in the type systems the result might not be the same as evaluating them in memory, including the SQL failing with errors while the in memory version works fine. However this allows to push more of the filter into the database, increasing performance.the postgis table. | advanced | Boolean | False | &#x60;False&#x60; |   | host | Host | user | String | True | &#x60;localhost&#x60; |   | Evictor tests per run | number of connections checked by the idle connection evictor for each of its runs (defaults to 3) | user | Integer | False | &#x60;3&#x60; |   | Loose bbox | Perform only primary filter on bbox | user | Boolean | False | &#x60;True&#x60; |   | Evictor run periodicity | number of seconds between idle object evitor runs (default, 300 seconds) | user | Integer | False | &#x60;300&#x60; |   | Estimated extends | Use the spatial index information to quickly get an estimate of the data bounds | user | Boolean | False | &#x60;True&#x60; |   | database | Database | user | String | False | &#x60; &#x60; |   | fetch size | number of records read with each iteraction with the dbms | user | Integer | False | &#x60;1000&#x60; |   | Test while idle | Periodically test the connections are still valid also while idle in the pool | user | Boolean | False | &#x60;True&#x60; |   | max connections | maximum number of open connections | user | Integer | False | &#x60;10&#x60; |   | preparedStatements | Use prepared statements | user | Boolean | False | &#x60;False&#x60; |   | Session close-up SQL | SQL statement executed when the connection is released to the pool | user | String | False | &#x60; &#x60; |   | schema | Schema | user | String | False | &#x60;public&#x60; |   | user | user name to login as | user | String | True | &#x60; &#x60; |  - Shapefile    Examples:   - application/xml:      &#x60;&#x60;&#x60;     &lt;dataStore&gt;       &lt;name&gt;nyc&lt;/name&gt;       &lt;connectionParameters&gt;         &lt;url&gt;file:/path/to/nyc.shp&lt;/database&gt;       &lt;/connectionParameters&gt;     &lt;/dataStore&gt;     &#x60;&#x60;&#x60;    - application/json:      &#x60;&#x60;&#x60;     {       \&quot;dataStore\&quot;: {         \&quot;name\&quot;: \&quot;nyc\&quot;,         \&quot;connectionParameters\&quot;: {           \&quot;entry\&quot;: [             {\&quot;@key\&quot;:\&quot;url\&quot;,\&quot;$\&quot;:\&quot;file:/path/to/nyc.shp\&quot;}           ]         }       }     }     &#x60;&#x60;&#x60;    Connection Parameters:    | key | description | level | type | required | default |   | --- | ----------- | ----- | ---- | -------- | ------- |   | cache and reuse memory maps | only memory map a file one, then cache and reuse the map | advanced | Boolean | False | &#x60;True&#x60; |   | namespace | uri to a the namespace | advanced | URI | False | &#x60; &#x60; |   | filetype | Discriminator for directory stores | program | String | False | &#x60;shapefile&#x60; |   | charset | character used to decode strings from the DBF file | advanced | Charset | False | &#x60;ISO-8859-1&#x60; |   | create spatial index | enable/disable the automatic creation of spatial index | advanced | Boolean | False | &#x60;True&#x60; |   | fstype | Enable using a setting of &#39;shape&#39;. | advanced | String | False | &#x60;shape&#x60; |   | url | url to a .shp file | user | URL | True | &#x60; &#x60; |   | enable spatial index | enable/disable the use of spatial index for local shapefiles | advanced | Boolean | False | &#x60;True&#x60; |   | memory mapped buffer | enable/disable the use of memory-mapped io | advanced | Boolean | False | &#x60;False&#x60; |   | timezone | time zone used to read dates from the DBF file | advanced | TimeZone | False | &#x60;Pacific Standard Time&#x60; |  - Directory of spatial files (shapefiles)    Examples:   - application/xml:      &#x60;&#x60;&#x60;     &lt;dataStore&gt;       &lt;name&gt;nyc&lt;/name&gt;       &lt;connectionParameters&gt;         &lt;url&gt;file:/path/to/directory&lt;/database&gt;       &lt;/connectionParameters&gt;     &lt;/dataStore&gt;     &#x60;&#x60;&#x60;    - application/json:      &#x60;&#x60;&#x60;     {       \&quot;dataStore\&quot;: {         \&quot;name\&quot;: \&quot;nyc\&quot;,         \&quot;connectionParameters\&quot;: {           \&quot;entry\&quot;: [             {\&quot;@key\&quot;:\&quot;url\&quot;,\&quot;$\&quot;:\&quot;file:/path/to/directory\&quot;}           ]         }       }     }     &#x60;&#x60;&#x60;    Connection Parameters:    | key | description | level | type | required | default |   | --- | ----------- | ----- | ---- | -------- | ------- |   | cache and reuse memory maps | only memory map a file one, then cache and reuse the map | advanced | Boolean | False | &#x60;True&#x60; |   | namespace | uri to a the namespace | advanced | URI | False | &#x60; &#x60; |   | filetype | Discriminator for directory stores | program | String | False | &#x60;shapefile&#x60; |   | charset | character used to decode strings from the DBF file | advanced | Charset | False | &#x60;ISO-8859-1&#x60; |   | create spatial index | enable/disable the automatic creation of spatial index | advanced | Boolean | False | &#x60;True&#x60; |   | fstype | Enable using a setting of &#39;shape&#39;. | advanced | String | False | &#x60;shape&#x60; |   | url | url to a .shp file | user | URL | True | &#x60; &#x60; |   | enable spatial index | enable/disable the use of spatial index for local shapefiles | advanced | Boolean | False | &#x60;True&#x60; |   | memory mapped buffer | enable/disable the use of memory-mapped io | advanced | Boolean | False | &#x60;False&#x60; |   | timezone | time zone used to read dates from the DBF file | advanced | TimeZone | False | &#x60;Pacific Standard Time&#x60; |   - Web Feature Service    Examples:   - application/xml:      &#x60;&#x60;&#x60;     &lt;dataStore&gt;       &lt;name&gt;nyc&lt;/name&gt;       &lt;connectionParameters&gt;         &lt;GET_CAPABILITIES_URL&gt;http://localhost:8080/geoserver/wfs?request&#x3D;GetCapabilities&lt;/GET_CAPABILITIES_URL&gt;       &lt;/connectionParameters&gt;     &lt;/dataStore&gt;     &#x60;&#x60;&#x60;    - application/json:      &#x60;&#x60;&#x60;     {       \&quot;dataStore\&quot;: {         \&quot;name\&quot;: \&quot;nyc\&quot;,         \&quot;connectionParameters\&quot;: {           \&quot;entry\&quot;: [             {\&quot;@key\&quot;:\&quot;GET_CAPABILITIES_URL\&quot;,\&quot;$\&quot;:\&quot;http://localhost:8080/geoserver/wfs?request&#x3D;GetCapabilities\&quot;}           ]         }       }     }     &#x60;&#x60;&#x60;    Connection Parameters:    | key | description | level | type | required | default |   | --- | ----------- | ----- | ---- | -------- | ------- |   | Protocol | Sets a preference for the HTTP protocol to use when requesting WFS functionality. Set this value to Boolean.TRUE for POST, Boolean.FALSE for GET or NULL for AUTO | user | Boolean | False | &#x60; &#x60; |   | WFS GetCapabilities URL | Represents a URL to the getCapabilities document or a server instance. | user | URL | False | &#x60; &#x60; |   | Buffer Size | This allows the user to specify a buffer size in features. This param has a default value of 10 features. | user | Integer | False | &#x60;10&#x60; |   | Filter compliance | Level of compliance to WFS specification (0-low,1-medium,2-high) | user | Integer | False | &#x60; &#x60; |   | EntityResolver | Sets the entity resolver used to expand XML entities | program | EntityResolver | False | &#x60;org.geotools.xml.PreventLocalEntityResolver@75e98519&#x60; |   | Time-out | This allows the user to specify a timeout in milliseconds. This param has a default value of 3000ms. | user | Integer | False | &#x60;3000&#x60; |   | GmlComplianceLevel | Optional OGC GML compliance level required. | user | Integer | False | &#x60;0&#x60; |   | Lenient | Indicates that datastore should do its best to create features from the provided data even if it does not accurately match the schema.  Errors will be logged but the parsing will continue if this is true.  Default is false | user | Boolean | False | &#x60;False&#x60; |   | Password | This allows the user to specify a username. This param should not be used without the USERNAME param. | user | String | False | &#x60; &#x60; |   | Use Default SRS | Use always the declared DefaultSRS for requests and reproject locally if necessary | advanced | Boolean | False | &#x60;False&#x60; |   | Namespace | Override the original WFS type name namespaces | advanced | String | False | &#x60; &#x60; |   | Username | This allows the user to specify a username. This param should not be used without the PASSWORD param. | user | String | False | &#x60; &#x60; |   | Axis Order Filter | Indicates axis order used by the remote WFS server for filters. It applies only to WFS 1.x.0 servers. Default is the same as AXIS_ORDER | advanced | String | False | &#x60; &#x60; |   | GmlCompatibleTypeNames | Use Gml Compatible TypeNames (replace : by _). | user | Boolean | False | &#x60;False&#x60; |   | Maximum features | Positive integer used as a hard limit for the amount of Features to retrieve for each FeatureType. A value of zero or not providing this parameter means no limit. | user | Integer | False | &#x60;0&#x60; |   | Axis Order | Indicates axis order used by the remote WFS server in result coordinates. It applies only to WFS 1.x.0 servers. Default is Compliant | advanced | String | False | &#x60;Compliant&#x60; |   | WFS Strategy | Override wfs stragegy with either cubwerx, ionic, mapserver, geoserver, strict, nonstrict or arcgis strategy. | user | String | False | &#x60;auto&#x60; |   | Try GZIP | Indicates that datastore should use gzip to transfer data if the server supports it. Default is true | user | Boolean | False | &#x60;True&#x60; |   | Encoding | This allows the user to specify the character encoding of the XML-Requests sent to the Server. Defaults to UTF-8 | user | String | False | &#x60;UTF-8&#x60; |   | Outputformat | This allows the user to specify an outputFormat, different from the default one. | advanced | String | False | &#x60; &#x60; |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutDataStoreUpload**
> PutDataStoreUpload(ctx, workspaceName, storeName, method, format, optional)
Uploads files to the data store, creating it if necessary

Creates or modifies a single data store by uploading spatial data or mapping configuration (in case an app-schema data store is targeted) files.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspaceName** | **string**| The name of the worskpace containing the coverage stores. | 
  **storeName** | **string**| The name of the store to be retrieved | 
  **method** | **string**| The upload method. Can be \&quot;url\&quot;, \&quot;file\&quot;, \&quot;external\&quot;. \&quot;file\&quot; uploads a file from a local source. The body of the request is the file itself. \&quot;url\&quot; uploads a file from an remote source. The body of the request is a URL pointing to the file to upload. This URL must be visible from the server. \&quot;external\&quot; uses an existing file on the server. The body of the request is the absolute path to the existing file. | 
  **format** | **string**| The type of source data store (e.g., \&quot;shp\&quot;). | 
 **optional** | ***PutDataStoreUploadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PutDataStoreUploadOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **configure** | **optional.String**| The configure parameter controls if a coverage/layer are configured upon file upload, in addition to creating the store. It can have a value of \&quot;none\&quot; to avoid configuring coverages. | 
 **target** | **optional.String**| The type of target data store (e.g., \&quot;shp\&quot;). Same as format if not provided. | 
 **update** | **optional.String**| The update mode. If \&quot;overwrite\&quot;, will overwrite existing data. Otherwise, will append to existing data. | 
 **charset** | **optional.String**| The character set of the data. | 
 **filename** | **optional.String**| The filename parameter specifies the target file name for the file to be uploaded. This is important to avoid clashes with existing files. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutDatastore**
> PutDatastore(ctx, workspaceName, storeName, datastore)
Modify a data store.

Modify data store ds. Use the \"Accept:\" header to specify format or append an extension to the endpoint (example \"/datastores/{ds}.xml\" for XML).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspaceName** | **string**| The name of the worskpace containing the data store. | 
  **storeName** | **string**| The name of the data store to modify. | 
  **datastore** | [**Datastore**](Datastore.md)| The updated data store definition. For a PUT, only values which should be changed need to be included. The connectionParameters map counts as a single value,  so if you change it all preexisting connection parameters will be overwritten.  The contents of the connection parameters will differ depending on the type of data store being added.  - GeoPackage    Examples:   - application/xml:      &#x60;&#x60;&#x60;     &lt;dataStore&gt;       &lt;description&gt;A data store&lt;/description&gt;       &lt;enabled&gt;true&lt;/enabled&gt;       &lt;__default&gt;true&lt;/__default&gt;       &lt;connectionParameters&gt;         &lt;database&gt;file:///path/to/nyc.gpkg&lt;/database&gt;       &lt;/connectionParameters&gt;     &lt;/dataStore&gt;     &#x60;&#x60;&#x60;    - application/json:      &#x60;&#x60;&#x60;     {       \&quot;dataStore\&quot;: {         \&quot;description\&quot;: \&quot;A data store\&quot;,         \&quot;enabled\&quot;: \&quot;true\&quot;,         \&quot;_default\&quot;: \&quot;true\&quot;,         \&quot;connectionParameters\&quot;: {           \&quot;entry\&quot;: [             {\&quot;@key\&quot;:\&quot;database\&quot;,\&quot;$\&quot;:\&quot;file:///path/to/nyc.gpkg\&quot;},           ]         }       }     }     &#x60;&#x60;&#x60;    Connection Parameters:    | key | description | level | type | required | default |   | --- | ----------- | ----- | ---- | -------- | ------- |   | Primary key metadata table | The optional table containing primary key structure and sequence associations. Can be expressed as &#39;schema.name&#39; or just &#39;name&#39; | user | String | False | &#x60; &#x60; |   | Callback factory | Name of JDBCReaderCallbackFactory to enable on the data store | user | String | False | &#x60; &#x60; |   | Evictor tests per run | number of connections checked by the idle connection evictor for each of its runs (defaults to 3) | user | Integer | False | &#x60;3&#x60; |   | database | Database | user | File | True | &#x60; &#x60; |   | Batch insert size | Number of records inserted in the same batch (default, 1). For optimal performance, set to 100. | user | Integer | False | &#x60;1&#x60; |   | fetch size | number of records read with each iteraction with the dbms | user | Integer | False | &#x60;1000&#x60; |   | Connection timeout | number of seconds the connection pool will wait before timing out attempting to get a new connection (default, 20 seconds) | user | Integer | False | &#x60;20&#x60; |   | namespace | Namespace prefix | user | String | False | &#x60; &#x60; |   | max connections | maximum number of open connections | user | Integer | False | &#x60;10&#x60; |   | Test while idle | Periodically test the connections are still valid also while idle in the pool | user | Boolean | False | &#x60;True&#x60; |   | Max connection idle time | number of seconds a connection needs to stay idle for the evictor to consider closing it | user | Integer | False | &#x60;300&#x60; |   | Session startup SQL | SQL statement executed when the connection is grabbed from the pool | user | String | False | &#x60; &#x60; |   | validate connections | check connection is alive before using it | user | Boolean | False | &#x60;True&#x60; |   | dbtype | Type | program | String | True | &#x60;geopkg&#x60; |   | passwd | password used to login | user | String | False | &#x60; &#x60; |   | Expose primary keys | Expose primary key columns as attributes of the feature type | user | Boolean | False | &#x60;False&#x60; |   | min connections | minimum number of pooled connection | user | Integer | False | &#x60;1&#x60; |   | Evictor run periodicity | number of seconds between idle object evitor runs (default, 300 seconds) | user | Integer | False | &#x60;300&#x60; |   | Session close-up SQL | SQL statement executed when the connection is released to the pool | user | String | False | &#x60; &#x60; |   | user | user name to login as | user | String | False | &#x60; &#x60; |  - PostGIS    Examples:   - application/xml:      &#x60;&#x60;&#x60;     &lt;dataStore&gt;       &lt;description&gt;A data store&lt;/description&gt;       &lt;enabled&gt;true&lt;/enabled&gt;       &lt;__default&gt;true&lt;/__default&gt;       &lt;connectionParameters&gt;         &lt;host&gt;localhost&lt;/host&gt;         &lt;port&gt;5432&lt;/port&gt;         &lt;database&gt;nyc&lt;/database&gt;         &lt;user&gt;bob&lt;/user&gt;         &lt;passwd&gt;postgres&lt;/passwd&gt;       &lt;/connectionParameters&gt;     &lt;/dataStore&gt;     &#x60;&#x60;&#x60;    - application/json:      &#x60;&#x60;&#x60;     {       \&quot;dataStore\&quot;: {         \&quot;description\&quot;: \&quot;A data store\&quot;,         \&quot;enabled\&quot;: \&quot;true\&quot;,         \&quot;_default\&quot;: \&quot;true\&quot;,         \&quot;connectionParameters\&quot;: {           \&quot;entry\&quot;: [             {\&quot;@key\&quot;:\&quot;host\&quot;,\&quot;$\&quot;:\&quot;localhost\&quot;},             {\&quot;@key\&quot;:\&quot;port\&quot;,\&quot;$\&quot;:\&quot;5432\&quot;},             {\&quot;@key\&quot;:\&quot;database\&quot;,\&quot;$\&quot;:\&quot;nyc\&quot;},             {\&quot;@key\&quot;:\&quot;user\&quot;,\&quot;$\&quot;:\&quot;bob\&quot;},             {\&quot;@key\&quot;:\&quot;passwd\&quot;,\&quot;$\&quot;:\&quot;postgres\&quot;},           ]         }       }     }     &#x60;&#x60;&#x60;    Connection Parameters:    | key | description | level | type | required | default |   | --- | ----------- | ----- | ---- | -------- | ------- |   | Connection timeout | number of seconds the connection pool will wait before timing out attempting to get a new connection (default, 20 seconds) | user | Integer | False | &#x60;20&#x60; |   | validate connections | check connection is alive before using it | user | Boolean | False | &#x60;True&#x60; |   | port | Port | user | Integer | True | &#x60;5432&#x60; |   | Primary key metadata table | The optional table containing primary key structure and sequence associations. Can be expressed as &#39;schema.name&#39; or just &#39;name&#39; | user | String | False | &#x60; &#x60; |   | Support on the fly geometry simplification | When enabled, operations such as map rendering will pass a hint that will enable the usage of ST_Simplify | user | Boolean | False | &#x60;True&#x60; |   | create database | Creates the database if it does not exist yet | advanced | Boolean | False | &#x60;False&#x60; |   | create database params | Extra specifications appeneded to the CREATE DATABASE command | advanced | String | False | &#x60;&#x60; |   | dbtype | Type | program | String | True | &#x60;postgis&#x60; |   | Batch insert size | Number of records inserted in the same batch (default, 1). For optimal performance, set to 100. | user | Integer | False | &#x60;1&#x60; |   | namespace | Namespace prefix | user | String | False | &#x60; &#x60; |   | Max connection idle time | number of seconds a connection needs to stay idle for the evictor to consider closing it | user | Integer | False | &#x60;300&#x60; |   | Session startup SQL | SQL statement executed when the connection is grabbed from the pool | user | String | False | &#x60; &#x60; |   | Expose primary keys | Expose primary key columns as attributes of the feature type | user | Boolean | False | &#x60;False&#x60; |   | min connections | minimum number of pooled connection | user | Integer | False | &#x60;1&#x60; |   | Max open prepared statements | Maximum number of prepared statements kept open and cached for each connection in the pool. Set to 0 to have unbounded caching, to -1 to disable caching | user | Integer | False | &#x60;50&#x60; |   | Callback factory | Name of JDBCReaderCallbackFactory to enable on the data store | user | String | False | &#x60; &#x60; |   | passwd | password used to login | user | String | False | &#x60; &#x60; |   | encode functions | set to true to have a set of filter functions be translated directly in SQL. Due to differences in the type systems the result might not be the same as evaluating them in memory, including the SQL failing with errors while the in memory version works fine. However this allows to push more of the filter into the database, increasing performance.the postgis table. | advanced | Boolean | False | &#x60;False&#x60; |   | host | Host | user | String | True | &#x60;localhost&#x60; |   | Evictor tests per run | number of connections checked by the idle connection evictor for each of its runs (defaults to 3) | user | Integer | False | &#x60;3&#x60; |   | Loose bbox | Perform only primary filter on bbox | user | Boolean | False | &#x60;True&#x60; |   | Evictor run periodicity | number of seconds between idle object evitor runs (default, 300 seconds) | user | Integer | False | &#x60;300&#x60; |   | Estimated extends | Use the spatial index information to quickly get an estimate of the data bounds | user | Boolean | False | &#x60;True&#x60; |   | database | Database | user | String | False | &#x60; &#x60; |   | fetch size | number of records read with each iteraction with the dbms | user | Integer | False | &#x60;1000&#x60; |   | Test while idle | Periodically test the connections are still valid also while idle in the pool | user | Boolean | False | &#x60;True&#x60; |   | max connections | maximum number of open connections | user | Integer | False | &#x60;10&#x60; |   | preparedStatements | Use prepared statements | user | Boolean | False | &#x60;False&#x60; |   | Session close-up SQL | SQL statement executed when the connection is released to the pool | user | String | False | &#x60; &#x60; |   | schema | Schema | user | String | False | &#x60;public&#x60; |   | user | user name to login as | user | String | True | &#x60; &#x60; |  - Shapefile    Examples:   - application/xml:      &#x60;&#x60;&#x60;     &lt;dataStore&gt;       &lt;description&gt;A data store&lt;/description&gt;       &lt;enabled&gt;true&lt;/enabled&gt;       &lt;__default&gt;true&lt;/__default&gt;       &lt;connectionParameters&gt;         &lt;url&gt;file:/path/to/nyc.shp&lt;/database&gt;       &lt;/connectionParameters&gt;     &lt;/dataStore&gt;     &#x60;&#x60;&#x60;    - application/json:      &#x60;&#x60;&#x60;     {       \&quot;dataStore\&quot;: {         \&quot;description\&quot;: \&quot;A data store\&quot;,         \&quot;enabled\&quot;: \&quot;true\&quot;,         \&quot;_default\&quot;: \&quot;true\&quot;,         \&quot;connectionParameters\&quot;: {           \&quot;entry\&quot;: [             {\&quot;@key\&quot;:\&quot;url\&quot;,\&quot;$\&quot;:\&quot;file:/path/to/nyc.shp\&quot;}           ]         }       }     }     &#x60;&#x60;&#x60;    Connection Parameters:    | key | description | level | type | required | default |   | --- | ----------- | ----- | ---- | -------- | ------- |   | cache and reuse memory maps | only memory map a file one, then cache and reuse the map | advanced | Boolean | False | &#x60;True&#x60; |   | namespace | uri to a the namespace | advanced | URI | False | &#x60; &#x60; |   | filetype | Discriminator for directory stores | program | String | False | &#x60;shapefile&#x60; |   | charset | character used to decode strings from the DBF file | advanced | Charset | False | &#x60;ISO-8859-1&#x60; |   | create spatial index | enable/disable the automatic creation of spatial index | advanced | Boolean | False | &#x60;True&#x60; |   | fstype | Enable using a setting of &#39;shape&#39;. | advanced | String | False | &#x60;shape&#x60; |   | url | url to a .shp file | user | URL | True | &#x60; &#x60; |   | enable spatial index | enable/disable the use of spatial index for local shapefiles | advanced | Boolean | False | &#x60;True&#x60; |   | memory mapped buffer | enable/disable the use of memory-mapped io | advanced | Boolean | False | &#x60;False&#x60; |   | timezone | time zone used to read dates from the DBF file | advanced | TimeZone | False | &#x60;Pacific Standard Time&#x60; |  - Directory of spatial files (shapefiles)    Examples:   - application/xml:      &#x60;&#x60;&#x60;     &lt;dataStore&gt;       &lt;description&gt;A data store&lt;/description&gt;       &lt;enabled&gt;true&lt;/enabled&gt;       &lt;__default&gt;true&lt;/__default&gt;       &lt;connectionParameters&gt;         &lt;url&gt;file:/path/to/directory&lt;/database&gt;       &lt;/connectionParameters&gt;     &lt;/dataStore&gt;     &#x60;&#x60;&#x60;    - application/json:      &#x60;&#x60;&#x60;     {       \&quot;dataStore\&quot;: {         \&quot;description\&quot;: \&quot;A data store\&quot;,         \&quot;enabled\&quot;: \&quot;true\&quot;,         \&quot;_default\&quot;: \&quot;true\&quot;,         \&quot;connectionParameters\&quot;: {           \&quot;entry\&quot;: [             {\&quot;@key\&quot;:\&quot;url\&quot;,\&quot;$\&quot;:\&quot;file:/path/to/directory\&quot;}           ]         }       }     }     &#x60;&#x60;&#x60;    Connection Parameters:    | key | description | level | type | required | default |   | --- | ----------- | ----- | ---- | -------- | ------- |   | cache and reuse memory maps | only memory map a file one, then cache and reuse the map | advanced | Boolean | False | &#x60;True&#x60; |   | namespace | uri to a the namespace | advanced | URI | False | &#x60; &#x60; |   | filetype | Discriminator for directory stores | program | String | False | &#x60;shapefile&#x60; |   | charset | character used to decode strings from the DBF file | advanced | Charset | False | &#x60;ISO-8859-1&#x60; |   | create spatial index | enable/disable the automatic creation of spatial index | advanced | Boolean | False | &#x60;True&#x60; |   | fstype | Enable using a setting of &#39;shape&#39;. | advanced | String | False | &#x60;shape&#x60; |   | url | url to a .shp file | user | URL | True | &#x60; &#x60; |   | enable spatial index | enable/disable the use of spatial index for local shapefiles | advanced | Boolean | False | &#x60;True&#x60; |   | memory mapped buffer | enable/disable the use of memory-mapped io | advanced | Boolean | False | &#x60;False&#x60; |   | timezone | time zone used to read dates from the DBF file | advanced | TimeZone | False | &#x60;Pacific Standard Time&#x60; |   - Web Feature Service    Examples:   - application/xml:      &#x60;&#x60;&#x60;     &lt;dataStore&gt;       &lt;description&gt;A data store&lt;/description&gt;       &lt;enabled&gt;true&lt;/enabled&gt;       &lt;__default&gt;true&lt;/__default&gt;       &lt;connectionParameters&gt;         &lt;GET_CAPABILITIES_URL&gt;http://localhost:8080/geoserver/wfs?request&#x3D;GetCapabilities&lt;/GET_CAPABILITIES_URL&gt;       &lt;/connectionParameters&gt;     &lt;/dataStore&gt;     &#x60;&#x60;&#x60;    - application/json:      &#x60;&#x60;&#x60;     {       \&quot;dataStore\&quot;: {         \&quot;description\&quot;: \&quot;A data store\&quot;,         \&quot;enabled\&quot;: \&quot;true\&quot;,         \&quot;_default\&quot;: \&quot;true\&quot;,         \&quot;connectionParameters\&quot;: {           \&quot;entry\&quot;: [             {\&quot;@key\&quot;:\&quot;GET_CAPABILITIES_URL\&quot;,\&quot;$\&quot;:\&quot;http://localhost:8080/geoserver/wfs?request&#x3D;GetCapabilities\&quot;}           ]         }       }     }     &#x60;&#x60;&#x60;    Connection Parameters:    | key | description | level | type | required | default |   | --- | ----------- | ----- | ---- | -------- | ------- |   | Protocol | Sets a preference for the HTTP protocol to use when requesting WFS functionality. Set this value to Boolean.TRUE for POST, Boolean.FALSE for GET or NULL for AUTO | user | Boolean | False | &#x60; &#x60; |   | WFS GetCapabilities URL | Represents a URL to the getCapabilities document or a server instance. | user | URL | False | &#x60; &#x60; |   | Buffer Size | This allows the user to specify a buffer size in features. This param has a default value of 10 features. | user | Integer | False | &#x60;10&#x60; |   | Filter compliance | Level of compliance to WFS specification (0-low,1-medium,2-high) | user | Integer | False | &#x60; &#x60; |   | EntityResolver | Sets the entity resolver used to expand XML entities | program | EntityResolver | False | &#x60;org.geotools.xml.PreventLocalEntityResolver@75e98519&#x60; |   | Time-out | This allows the user to specify a timeout in milliseconds. This param has a default value of 3000ms. | user | Integer | False | &#x60;3000&#x60; |   | GmlComplianceLevel | Optional OGC GML compliance level required. | user | Integer | False | &#x60;0&#x60; |   | Lenient | Indicates that datastore should do its best to create features from the provided data even if it does not accurately match the schema.  Errors will be logged but the parsing will continue if this is true.  Default is false | user | Boolean | False | &#x60;False&#x60; |   | Password | This allows the user to specify a username. This param should not be used without the USERNAME param. | user | String | False | &#x60; &#x60; |   | Use Default SRS | Use always the declared DefaultSRS for requests and reproject locally if necessary | advanced | Boolean | False | &#x60;False&#x60; |   | Namespace | Override the original WFS type name namespaces | advanced | String | False | &#x60; &#x60; |   | Username | This allows the user to specify a username. This param should not be used without the PASSWORD param. | user | String | False | &#x60; &#x60; |   | Axis Order Filter | Indicates axis order used by the remote WFS server for filters. It applies only to WFS 1.x.0 servers. Default is the same as AXIS_ORDER | advanced | String | False | &#x60; &#x60; |   | GmlCompatibleTypeNames | Use Gml Compatible TypeNames (replace : by _). | user | Boolean | False | &#x60;False&#x60; |   | Maximum features | Positive integer used as a hard limit for the amount of Features to retrieve for each FeatureType. A value of zero or not providing this parameter means no limit. | user | Integer | False | &#x60;0&#x60; |   | Axis Order | Indicates axis order used by the remote WFS server in result coordinates. It applies only to WFS 1.x.0 servers. Default is Compliant | advanced | String | False | &#x60;Compliant&#x60; |   | WFS Strategy | Override wfs stragegy with either cubwerx, ionic, mapserver, geoserver, strict, nonstrict or arcgis strategy. | user | String | False | &#x60;auto&#x60; |   | Try GZIP | Indicates that datastore should use gzip to transfer data if the server supports it. Default is true | user | Boolean | False | &#x60;True&#x60; |   | Encoding | This allows the user to specify the character encoding of the XML-Requests sent to the Server. Defaults to UTF-8 | user | String | False | &#x60;UTF-8&#x60; |   | Outputformat | This allows the user to specify an outputFormat, different from the default one. | advanced | String | False | &#x60; &#x60; |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Putdatastores**
> Putdatastores(ctx, )


Invalid. Use POST for adding a new data store, or PUT on /datastores/{datastore} to edit an existing data store.

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

