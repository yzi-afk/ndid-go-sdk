/*
 * Platform Utility function 
 *
 * Other functions needed by anyone using the platform
 *
 * API version: 5.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Message struct {
	MessageId string `json:"message_id,omitempty"`
	RequesterNodeId string `json:"requester_node_id"`
	Message string `json:"message"`
	Purpose string `json:"purpose"`
	// <CHAIN_ID>:<BLOCK_HEIGHT>
	CreationBlockHeight string `json:"creation_block_height"`
}
