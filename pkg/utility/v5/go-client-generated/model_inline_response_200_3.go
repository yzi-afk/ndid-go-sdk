/*
 * Platform Utility function 
 *
 * Other functions needed by anyone using the platform
 *
 * API version: 5.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type InlineResponse2003 struct {
	ServiceId string `json:"service_id"`
	ServiceName string `json:"service_name"`
	DataSchema string `json:"data_schema,omitempty"`
	DataSchemaVersion string `json:"data_schema_version,omitempty"`
	Active bool `json:"active"`
}
