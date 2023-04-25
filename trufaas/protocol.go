package trufaas

import "net/http"

var protocolHeaders = map[string]string{
	TrustVerificationStatus: "",
	MsgAuthCode:             "",
	VerifiedFnName:          "",
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

}

func GetTrustProtocolHeadersFromInvoker(req *http.Request) {

}
