package trufaas

const (

	//Response Messages
	TrustValuesCreationSuccessMsg = "[TruFaaS] Function Trust Value Generated."
	TrustValuesCreationFailMsg    = "[TruFaaS] Function Trust Value Generation Failed.Delete The Function and Try Again"
	TrustVerificationFailedMsg    = "[TruFaaS] Function Invocation Stopped as Function Trust Verification Failed."

	// Trust Protocol Headers
	InvokerPubKey           = "Invoker-Public-Key"
	MsgAuthCode             = "MAC"
	TrustVerificationStatus = "Trust-Verification-Status"
	VerifiedFnName          = "Function-Name"
)
