package trufaas

const (
	//ExternalCompBaseURL = "https://trufaas-external.onrender.com"
	//ExternalCompBaseURL = "https://turfaas-external-comp-avishka.onrender.com"

	/*
		Get your https url for external component running on local host

			Method 1:
				Go to https://theboroer.github.io/localtunnel-www/ and install local tunnel
				Then run TruFaaS external component locally
				Get https url by running 'lt --port 8080'

			Method 2:
				Go to https://ngrok.com and setup ngrok with auth token
				Then run TruFaaS external component locally
				Get https url by running 'ngrok http 8080'
	*/

	ExternalCompBaseCreateURL = "http://localhost:8080"
	ExternalCompBaseExecURL   = "http://192.168.8.150:8080"

	CreateURL = ExternalCompBaseCreateURL + "/fn/create"
	VerifyURL = ExternalCompBaseExecURL + "/fn/verify"
)
