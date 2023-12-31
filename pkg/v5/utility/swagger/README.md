# Go API client for swagger

Other functions needed by anyone using the platform

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 5.2
- Package version: 1.0.0
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./swagger"
```

## Documentation for API Endpoints

All URIs are relative to *https://virtserver.swaggerhub.com/ndid/utility/0.1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**GetAsIdList**](docs/DefaultApi.md#getasidlist) | **Get** /utility/as/{service_id} | Retrieve list of id of all available as given a service id
*DefaultApi* | [**GetIdpList**](docs/DefaultApi.md#getidplist) | **Get** /utility/idp | Retrieve all IdP nodes
*DefaultApi* | [**GetIdpListForIdentifier**](docs/DefaultApi.md#getidplistforidentifier) | **Get** /utility/idp/{namespace}/{identifier} | Retrieve all IdP nodes relevant to/authorized by/have onboarded this {namespace}/{identifier}.
*DefaultApi* | [**GetMessage**](docs/DefaultApi.md#getmessage) | **Get** /utility/messages/{message_id} | Fetch Message info
*DefaultApi* | [**GetRequestStatus**](docs/DefaultApi.md#getrequeststatus) | **Get** /utility/requests/{request_id} | Fetch Request info and status
*DefaultApi* | [**GetRequestTypeList**](docs/DefaultApi.md#getrequesttypelist) | **Get** /utility/request_types | Get all valid request types
*DefaultApi* | [**GetSuppressedIdentityModificationNotificationNodeIdList**](docs/DefaultApi.md#getsuppressedidentitymodificationnotificationnodeidlist) | **Get** /utility/suppressed_identity_modification_notification_node_ids | Get all suppressed identity modification notification node ID list
*DefaultApi* | [**GetSuppressedIdentityModificationNotificationStatus**](docs/DefaultApi.md#getsuppressedidentitymodificationnotificationstatus) | **Get** /utility/suppressed_identity_modification_notification_node_ids/{node_id} | Get suppressed identity modification notification status
*DefaultApi* | [**UtilityAsErrorCodesGet**](docs/DefaultApi.md#utilityaserrorcodesget) | **Get** /utility/as_error_codes | Retrieve list of all AS error codes
*DefaultApi* | [**UtilityAsServicePriceServiceIdGet**](docs/DefaultApi.md#utilityasservicepriceserviceidget) | **Get** /utility/as/service_price/{service_id} | Retrieve AS fee of a specific service_id
*DefaultApi* | [**UtilityIdpErrorCodesGet**](docs/DefaultApi.md#utilityidperrorcodesget) | **Get** /utility/idp_error_codes | Retrieve list of all IdP error codes
*DefaultApi* | [**UtilityNamespacesGet**](docs/DefaultApi.md#utilitynamespacesget) | **Get** /utility/namespaces | Retrieve list of all available namespaces
*DefaultApi* | [**UtilityNodesNodeIdGet**](docs/DefaultApi.md#utilitynodesnodeidget) | **Get** /utility/nodes/{node_id} | Retrieve information of node ID
*DefaultApi* | [**UtilityNodesNodeIdTokenGet**](docs/DefaultApi.md#utilitynodesnodeidtokenget) | **Get** /utility/nodes/{node_id}/token | Retrieve available token for node ID
*DefaultApi* | [**UtilityPrivateMessageRemovalPost**](docs/DefaultApi.md#utilityprivatemessageremovalpost) | **Post** /utility/private_message_removal | Remove all messages from MQ
*DefaultApi* | [**UtilityPrivateMessageRemovalRequestIdPost**](docs/DefaultApi.md#utilityprivatemessageremovalrequestidpost) | **Post** /utility/private_message_removal/{request_id} | Remove messages from MQ
*DefaultApi* | [**UtilityPrivateMessagesRequestIdGet**](docs/DefaultApi.md#utilityprivatemessagesrequestidget) | **Get** /utility/private_messages/{request_id} | Get messages from MQ
*DefaultApi* | [**UtilityServicePriceCeilingGet**](docs/DefaultApi.md#utilityservicepriceceilingget) | **Get** /utility/service_price_ceiling | Get maximum price allowed for AS pricing
*DefaultApi* | [**UtilityServicePriceMinEffectiveDatetimeDelayGet**](docs/DefaultApi.md#utilityservicepricemineffectivedatetimedelayget) | **Get** /utility/service_price_min_effective_datetime_delay | Get service price minimum effective datetime delay
*DefaultApi* | [**UtilityServicesGet**](docs/DefaultApi.md#utilityservicesget) | **Get** /utility/services | Retrieve list of all available services
*DefaultApi* | [**UtilityServicesServiceIdGet**](docs/DefaultApi.md#utilityservicesserviceidget) | **Get** /utility/services/{service_id} | Retrieve service details

## Documentation For Models

 - [As](docs/As.md)
 - [AsPriceInner](docs/AsPriceInner.md)
 - [ErrorError](docs/ErrorError.md)
 - [Idp](docs/Idp.md)
 - [IdpForIdentity](docs/IdpForIdentity.md)
 - [InboundPrivateMessageInner](docs/InboundPrivateMessageInner.md)
 - [InlineResponse200](docs/InlineResponse200.md)
 - [InlineResponse2001](docs/InlineResponse2001.md)
 - [InlineResponse2002](docs/InlineResponse2002.md)
 - [InlineResponse2003](docs/InlineResponse2003.md)
 - [InlineResponse2004](docs/InlineResponse2004.md)
 - [InlineResponse2005](docs/InlineResponse2005.md)
 - [InlineResponse2006](docs/InlineResponse2006.md)
 - [Message](docs/Message.md)
 - [ModelError](docs/ModelError.md)
 - [NodeInfo](docs/NodeInfo.md)
 - [NodeInfoMq](docs/NodeInfoMq.md)
 - [NodeInfoProxy](docs/NodeInfoProxy.md)
 - [OutboundPrivateMessageInner](docs/OutboundPrivateMessageInner.md)
 - [PriceCeiling](docs/PriceCeiling.md)
 - [PriceCeilingPriceCeilingByCurrencyList](docs/PriceCeilingPriceCeilingByCurrencyList.md)
 - [PriceMinEffectiveDatetimeDelay](docs/PriceMinEffectiveDatetimeDelay.md)
 - [RequestsStatus](docs/RequestsStatus.md)
 - [RequestsStatusDataRequestList](docs/RequestsStatusDataRequestList.md)
 - [RequestsStatusResponseList](docs/RequestsStatusResponseList.md)
 - [RequestsStatusResponseList1](docs/RequestsStatusResponseList1.md)

## Documentation For Authorization
 Endpoints do not require authorization.


## Author


