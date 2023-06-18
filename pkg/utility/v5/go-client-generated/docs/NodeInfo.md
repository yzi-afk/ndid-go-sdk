# NodeInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicKey** | **string** |  | [default to null]
**MasterPublicKey** | **string** |  | [default to null]
**NodeName** | **string** |  | [default to null]
**Role** | **string** |  | [default to null]
**MaxIal** | **float64** | Available when node&#x27;s role is IdP | [optional] [default to null]
**MaxAal** | **float64** | Available when node&#x27;s role is IdP | [optional] [default to null]
**OnTheFlySupport** | **bool** | Available when node&#x27;s role is IdP | [optional] [default to null]
**SupportedRequestMessageDataUrlTypeList** | **[]string** | Present only on IdP nodes. Supported request message MIME types. | [optional] [default to null]
**Agent** | **bool** | Present only on IdP nodes | [optional] [default to null]
**NodeIdWhitelist** | **[]string** | Present only on RP and IdP nodes | [optional] [default to null]
**NodeIdWhitelistActive** | **bool** | Present only on RP and IdP nodes | [optional] [default to null]
**Mq** | [**[]NodeInfoMq**](NodeInfo_mq.md) |  | [optional] [default to null]
**Proxy** | [***NodeInfoProxy**](NodeInfo_proxy.md) |  | [optional] [default to null]
**Active** | **bool** |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

