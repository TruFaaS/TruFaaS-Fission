package trufaas

const (
	ExternalCompBaseURL = "http://{your-ipv4-address}:8080"

	CreateURL = "http://localhost:8080/fn/create"
	VerifyURL = ExternalCompBaseURL + "/fn/verify"
)
