package trufaas

import (
	"net/http"
)

/**
Following utility functions are used to send the required headers for the trufaas trust protocol through fission microservices

The header "x-invoker-public-key" is sent from invoker to external component through fission.

The headers "x-trufaas-public-key", "x-trufaas-mac", "x-trufaas-trust-verification" sent form external component back to invoker through fission

The path headers take,
	"invoker" <--> "fission router service" <--> "fission executor service" <--> "fission fetcher service" <--> "trufaas external component"
*/

var protocolHeadersMap = map[string]string{
	ProtocolTrustVerificationStatus: "",
	ProtocolMsgAuthCode:             "",
	ProtocolInvokerPubKey:           "",
	ProtocolExternalCompPubKey:      "",
}

func GetTrustProtocolHeadersFromReq(req *http.Request) {
	for headerName, _ := range protocolHeadersMap {
		headerValue := req.Header.Get(headerName)
		if headerValue != "" {
			protocolHeadersMap[headerName] = headerValue
		}
	}
}

func AddTrustProtocolHeadersToReq(req *http.Request) {
	for headerName, headerValue := range protocolHeadersMap {
		if headerValue != "" {
			req.Header.Set(headerName, headerValue)
			protocolHeadersMap[headerName] = ""
		}
	}
}

func GetTrustProtocolHeadersFromResp(resp *http.Response) {
	for headerName, _ := range protocolHeadersMap {
		headerValue := resp.Header.Get(headerName)
		if headerValue != "" {
			protocolHeadersMap[headerName] = headerValue
		}
	}
}

func AddTrustProtocolHeadersToResp(resp *http.Response) {
	for headerName, headerValue := range protocolHeadersMap {
		if headerValue != "" {
			resp.Header.Set(headerName, headerValue)
			protocolHeadersMap[headerName] = ""
		}
	}
}

func AddTrustProtocolHeadersToRespWriter(rw http.ResponseWriter) {
	for headerName, headerValue := range protocolHeadersMap {
		if headerValue != "" {
			rw.Header().Set(headerName, headerValue)
			protocolHeadersMap[headerName] = ""
		}
	}
}
