/*
 * GeoServer Layers
 *
 * A layer is a published resource (feature type or coverage).
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type Layer struct {
	// Name of the layer
	Name string `json:"name,omitempty"`
	// Location of the layer in the WMS capabilities layer tree
	Path string `json:"path,omitempty"`
	// Type of published layer. Can be VECTOR, RASTER, REMOTE, WMS or GROUP. Must be consistent with resource definition.
	Type string `json:"type,omitempty"`
	DefaultStyle StyleReference `json:"defaultStyle,omitempty"`
	Styles LayerStyles `json:"styles,omitempty"`
	Resource LayerResource `json:"resource,omitempty"`
	// Controls layer transparency (whether the layer is opaque or transparent).
	Opaque bool `json:"opaque,omitempty"`
	Metadata []MetadataEntry `json:"metadata,omitempty"`
	Attribution LayerAttribution `json:"attribution,omitempty"`
	AuthorityURLs []AuthorityUrl `json:"authorityURLs,omitempty"`
	Identifiers []Identifier `json:"identifiers,omitempty"`
}