package trufaas

import "net/http"

var protocolHeaders = map[string]string{
	ProtocolTrustVerificationStatus: "",
	ProtocolMsgAuthCode:             "",
	ProtocolVerifiedFnName:          "",
	ProtocolInvokerPubKey:           " ",
}

func GetTrustProtocolHeadersFromExComp(resp http.Response) {
	for _, headerName := range protocolHeaders {
		if headerValue := resp.Header.Get(headerName); headerValue != "" {
			protocolHeaders[headerName] = headerValue
		}
	}
}
func AddTrustProtocolHeadersToFnResponse(resp *http.Response) {

	for _, headerName := range protocolHeaders {
		if headerValue := protocolHeaders[headerName]; headerValue != "" {
			resp.Header.Set(headerName, headerValue)
		}
	}
}

func AddTrustProtocolHeadersToErrResponse(rw http.ResponseWriter) {

	for _, headerName := range protocolHeaders {
		if headerValue := protocolHeaders[headerName]; headerValue != "" {
			rw.Header().Set(headerName, headerValue)
		}
	}
}

func AddTrustProtocolHeadersToExCompReq(req *http.Request) {
	if protocolHeaders[ProtocolInvokerPubKey] != "" {
		req.Header.Set(ProtocolInvokerPubKey, protocolHeaders[ProtocolInvokerPubKey])
	}
}

func GetTrustProtocolHeadersFromInvoker(req *http.Request) {
	invokerPubKey := req.Header.Get(ProtocolInvokerPubKey)
	if invokerPubKey != "" {
		protocolHeaders[ProtocolInvokerPubKey] = invokerPubKey
	}
}
