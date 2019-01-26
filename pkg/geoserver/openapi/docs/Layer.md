# Layer

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the layer | [optional] 
**Path** | **string** | Location of the layer in the WMS capabilities layer tree | [optional] 
**Type** | **string** | Type of published layer. Can be VECTOR, RASTER, REMOTE, WMS or GROUP. Must be consistent with resource definition. | [optional] 
**DefaultStyle** | [**StyleReference**](StyleReference.md) |  | [optional] 
**Styles** | [**LayerStyles**](Layer_styles.md) |  | [optional] 
**Resource** | [**LayerResource**](Layer_resource.md) |  | [optional] 
**Opaque** | **bool** | Controls layer transparency (whether the layer is opaque or transparent). | [optional] 
**Metadata** | [**[]MetadataEntry**](MetadataEntry.md) |  | [optional] 
**Attribution** | [**LayerAttribution**](Layer_attribution.md) |  | [optional] 
**AuthorityURLs** | [**[]AuthorityUrl**](AuthorityURL.md) |  | [optional] 
**Identifiers** | [**[]Identifier**](Identifier.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


