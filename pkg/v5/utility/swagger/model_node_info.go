/*
 * Platform Utility function 
 *
 * Other functions needed by anyone using the platform
 *
 * API version: 5.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type NodeInfo struct {
	PublicKey string `json:"public_key"`
	MasterPublicKey string `json:"master_public_key"`
	NodeName string `json:"node_name"`
	Role string `json:"role"`
	// Available when node's role is IdP
	MaxIal float64 `json:"max_ial,omitempty"`
	// Available when node's role is IdP
	MaxAal float64 `json:"max_aal,omitempty"`
	// Available when node's role is IdP
	OnTheFlySupport bool `json:"on_the_fly_support,omitempty"`
	// Present only on IdP nodes. Supported request message MIME types.
	SupportedRequestMessageDataUrlTypeList []string `json:"supported_request_message_data_url_type_list,omitempty"`
	// Present only on IdP nodes
	Agent bool `json:"agent,omitempty"`
	// Present only on RP and IdP nodes
	NodeIdWhitelist []string `json:"node_id_whitelist,omitempty"`
	// Present only on RP and IdP nodes
	NodeIdWhitelistActive bool `json:"node_id_whitelist_active,omitempty"`
	Mq []NodeInfoMq `json:"mq,omitempty"`
	Proxy *NodeInfoProxy `json:"proxy,omitempty"`
	Active bool `json:"active"`
}
