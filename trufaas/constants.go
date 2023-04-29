package trufaas

const (

	// Response Messages
	TrustValuesCreationSuccessMsg = "[TruFaaS] Function Trust Value Generated."
	TrustValuesCreationFailMsg    = "[TruFaaS] Function Trust Value Generation Failed.Delete The Function and Try Again"
	TrustVerificationFailedMsg    = "[TruFaaS] Function Invocation Stopped as Function Trust Verification Failed."

	// Trust Protocol Headers
	ProtocolInvokerPubKey           = "x-invoker-public-key"
	ProtocolExternalCompPubKey      = "x-trufaas-public-key"
	ProtocolMsgAuthCode             = "x-trufaas-mac"
	ProtocolTrustVerificationStatus = "x-trufaas-trust-verification"
)
