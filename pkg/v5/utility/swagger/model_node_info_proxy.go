/*
 * Platform Utility function 
 *
 * Other functions needed by anyone using the platform
 *
 * API version: 5.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type NodeInfoProxy struct {
	NodeId string `json:"node_id"`
	NodeName string `json:"node_name"`
	PublicKey string `json:"public_key"`
	MasterPublicKey string `json:"master_public_key"`
	Mq *NodeInfoMq `json:"mq,omitempty"`
	Config string `json:"config"`
}
