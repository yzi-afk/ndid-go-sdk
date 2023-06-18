# {{classname}}

All URIs are relative to *https://virtserver.swaggerhub.com/ndid/utility/0.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAsIdList**](DefaultApi.md#GetAsIdList) | **Get** /utility/as/{service_id} | Retrieve list of id of all available as given a service id
[**GetIdpList**](DefaultApi.md#GetIdpList) | **Get** /utility/idp | Retrieve all IdP nodes
[**GetIdpListForIdentifier**](DefaultApi.md#GetIdpListForIdentifier) | **Get** /utility/idp/{namespace}/{identifier} | Retrieve all IdP nodes relevant to/authorized by/have onboarded this {namespace}/{identifier}.
[**GetMessage**](DefaultApi.md#GetMessage) | **Get** /utility/messages/{message_id} | Fetch Message info
[**GetRequestStatus**](DefaultApi.md#GetRequestStatus) | **Get** /utility/requests/{request_id} | Fetch Request info and status
[**GetRequestTypeList**](DefaultApi.md#GetRequestTypeList) | **Get** /utility/request_types | Get all valid request types
[**GetSuppressedIdentityModificationNotificationNodeIdList**](DefaultApi.md#GetSuppressedIdentityModificationNotificationNodeIdList) | **Get** /utility/suppressed_identity_modification_notification_node_ids | Get all suppressed identity modification notification node ID list
[**GetSuppressedIdentityModificationNotificationStatus**](DefaultApi.md#GetSuppressedIdentityModificationNotificationStatus) | **Get** /utility/suppressed_identity_modification_notification_node_ids/{node_id} | Get suppressed identity modification notification status
[**UtilityAsErrorCodesGet**](DefaultApi.md#UtilityAsErrorCodesGet) | **Get** /utility/as_error_codes | Retrieve list of all AS error codes
[**UtilityAsServicePriceServiceIdGet**](DefaultApi.md#UtilityAsServicePriceServiceIdGet) | **Get** /utility/as/service_price/{service_id} | Retrieve AS fee of a specific service_id
[**UtilityIdpErrorCodesGet**](DefaultApi.md#UtilityIdpErrorCodesGet) | **Get** /utility/idp_error_codes | Retrieve list of all IdP error codes
[**UtilityNamespacesGet**](DefaultApi.md#UtilityNamespacesGet) | **Get** /utility/namespaces | Retrieve list of all available namespaces
[**UtilityNodesNodeIdGet**](DefaultApi.md#UtilityNodesNodeIdGet) | **Get** /utility/nodes/{node_id} | Retrieve information of node ID
[**UtilityNodesNodeIdTokenGet**](DefaultApi.md#UtilityNodesNodeIdTokenGet) | **Get** /utility/nodes/{node_id}/token | Retrieve available token for node ID
[**UtilityPrivateMessageRemovalPost**](DefaultApi.md#UtilityPrivateMessageRemovalPost) | **Post** /utility/private_message_removal | Remove all messages from MQ
[**UtilityPrivateMessageRemovalRequestIdPost**](DefaultApi.md#UtilityPrivateMessageRemovalRequestIdPost) | **Post** /utility/private_message_removal/{request_id} | Remove messages from MQ
[**UtilityPrivateMessagesRequestIdGet**](DefaultApi.md#UtilityPrivateMessagesRequestIdGet) | **Get** /utility/private_messages/{request_id} | Get messages from MQ
[**UtilityServicePriceCeilingGet**](DefaultApi.md#UtilityServicePriceCeilingGet) | **Get** /utility/service_price_ceiling | Get maximum price allowed for AS pricing
[**UtilityServicePriceMinEffectiveDatetimeDelayGet**](DefaultApi.md#UtilityServicePriceMinEffectiveDatetimeDelayGet) | **Get** /utility/service_price_min_effective_datetime_delay | Get service price minimum effective datetime delay
[**UtilityServicesGet**](DefaultApi.md#UtilityServicesGet) | **Get** /utility/services | Retrieve list of all available services
[**UtilityServicesServiceIdGet**](DefaultApi.md#UtilityServicesServiceIdGet) | **Get** /utility/services/{service_id} | Retrieve service details

# **GetAsIdList**
> []As GetAsIdList(ctx, serviceId)
Retrieve list of id of all available as given a service id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **string**| Service ID | 

### Return type

[**[]As**](AS.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIdpList**
> []Idp GetIdpList(ctx, optional)
Retrieve all IdP nodes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiGetIdpListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiGetIdpListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **minIal** | **optional.Float64**| minimum ial supported by IDP looking for | 
 **minAal** | **optional.Float64**| minimum aal supported by IDP looking for | 
 **agent** | **optional.Bool**| filter IdP agent type | 
 **onTheFlySupport** | **optional.Bool**| filter only IdP that support on-the-fly onboarding | 
 **filterForNodeId** | **optional.String**| result IdP must be included in node&#x27;s whitelist | 

### Return type

[**[]Idp**](IDP.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIdpListForIdentifier**
> []IdpForIdentity GetIdpListForIdentifier(ctx, namespace, identifier, optional)
Retrieve all IdP nodes relevant to/authorized by/have onboarded this {namespace}/{identifier}.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**| ID Namespace | 
  **identifier** | **string**| Identifier for the ID | 
 **optional** | ***DefaultApiGetIdpListForIdentifierOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiGetIdpListForIdentifierOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **minIal** | **optional.Float64**| minimum ial supported by IDP looking for | 
 **minAal** | **optional.Float64**| minimum aal supported by IDP looking for | 
 **onTheFlySupport** | **optional.Bool**| filter only IdP that support on-the-fly onboarding | 
 **mode** | **optional.Float64**| mode at the time of identity creation | 
 **filterForNodeId** | **optional.String**| result IdP must be included in node&#x27;s whitelist | 

### Return type

[**[]IdpForIdentity**](IDPForIdentity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMessage**
> Message GetMessage(ctx, messageId)
Fetch Message info

Fetch Message info

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **messageId** | **string**| message ID | 

### Return type

[**Message**](Message.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRequestStatus**
> RequestsStatus GetRequestStatus(ctx, requestId)
Fetch Request info and status

Fetch Request info and status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **requestId** | **string**| request ID | 

### Return type

[**RequestsStatus**](RequestsStatus.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRequestTypeList**
> []string GetRequestTypeList(ctx, )
Get all valid request types

Get all valid request types

### Required Parameters
This endpoint does not need any parameter.

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSuppressedIdentityModificationNotificationNodeIdList**
> []string GetSuppressedIdentityModificationNotificationNodeIdList(ctx, optional)
Get all suppressed identity modification notification node ID list

Get all suppressed identity modification notification node ID list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiGetSuppressedIdentityModificationNotificationNodeIdListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiGetSuppressedIdentityModificationNotificationNodeIdListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **prefix** | **optional.String**| Node ID prefix | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSuppressedIdentityModificationNotificationStatus**
> InlineResponse2006 GetSuppressedIdentityModificationNotificationStatus(ctx, nodeId)
Get suppressed identity modification notification status

Get suppressed identity modification notification status by node ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeId** | **string**| Node ID | 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UtilityAsErrorCodesGet**
> []InlineResponse2004 UtilityAsErrorCodesGet(ctx, )
Retrieve list of all AS error codes

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]InlineResponse2004**](inline_response_200_4.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UtilityAsServicePriceServiceIdGet**
> []AsPriceInner UtilityAsServicePriceServiceIdGet(ctx, serviceId, optional)
Retrieve AS fee of a specific service_id

Retrieve AS fee history list sorted by latest first of a specific service_id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **string**| Service ID | 
 **optional** | ***DefaultApiUtilityAsServicePriceServiceIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiUtilityAsServicePriceServiceIdGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nodeId** | **optional.String**| Filter by node ID | 

### Return type

[**[]AsPriceInner**](array.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UtilityIdpErrorCodesGet**
> []InlineResponse2004 UtilityIdpErrorCodesGet(ctx, )
Retrieve list of all IdP error codes

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]InlineResponse2004**](inline_response_200_4.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UtilityNamespacesGet**
> []InlineResponse2001 UtilityNamespacesGet(ctx, )
Retrieve list of all available namespaces

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]InlineResponse2001**](inline_response_200_1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UtilityNodesNodeIdGet**
> NodeInfo UtilityNodesNodeIdGet(ctx, nodeId)
Retrieve information of node ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeId** | **string**| Node ID | 

### Return type

[**NodeInfo**](NodeInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UtilityNodesNodeIdTokenGet**
> InlineResponse200 UtilityNodesNodeIdTokenGet(ctx, nodeId)
Retrieve available token for node ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeId** | **string**| Node ID | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UtilityPrivateMessageRemovalPost**
> UtilityPrivateMessageRemovalPost(ctx, optional)
Remove all messages from MQ

Remove all messages received from message queue

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiUtilityPrivateMessageRemovalPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiUtilityPrivateMessageRemovalPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of interface{}**](interface{}.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UtilityPrivateMessageRemovalRequestIdPost**
> UtilityPrivateMessageRemovalRequestIdPost(ctx, requestId, optional)
Remove messages from MQ

Remove messages from MQ with specified request ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **requestId** | **string**| Request ID | 
 **optional** | ***DefaultApiUtilityPrivateMessageRemovalRequestIdPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiUtilityPrivateMessageRemovalRequestIdPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of interface{}**](interface{}.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UtilityPrivateMessagesRequestIdGet**
> InlineResponse2005 UtilityPrivateMessagesRequestIdGet(ctx, requestId, optional)
Get messages from MQ

Get messages received from message queue with specified request ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **requestId** | **string**| Request ID | 
 **optional** | ***DefaultApiUtilityPrivateMessagesRequestIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiUtilityPrivateMessagesRequestIdGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nodeId** | **optional.String**| Call this API as node ID | 

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UtilityServicePriceCeilingGet**
> PriceCeiling UtilityServicePriceCeilingGet(ctx, serviceId)
Get maximum price allowed for AS pricing

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **string**| Service ID | 

### Return type

[**PriceCeiling**](PriceCeiling.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UtilityServicePriceMinEffectiveDatetimeDelayGet**
> PriceMinEffectiveDatetimeDelay UtilityServicePriceMinEffectiveDatetimeDelayGet(ctx, optional)
Get service price minimum effective datetime delay

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiUtilityServicePriceMinEffectiveDatetimeDelayGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiUtilityServicePriceMinEffectiveDatetimeDelayGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceId** | **optional.String**| Service ID | 

### Return type

[**PriceMinEffectiveDatetimeDelay**](PriceMinEffectiveDatetimeDelay.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UtilityServicesGet**
> []InlineResponse2002 UtilityServicesGet(ctx, )
Retrieve list of all available services

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]InlineResponse2002**](inline_response_200_2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UtilityServicesServiceIdGet**
> InlineResponse2003 UtilityServicesServiceIdGet(ctx, serviceId)
Retrieve service details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **string**| Service ID | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

