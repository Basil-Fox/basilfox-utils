package constant

type EndpointType string

const (
	ContextTokenClaims = "token_claims"

	EndpointPublic  EndpointType = "Public"
	EndpointPrivate EndpointType = "Private"

	HeaderNamespace = "X-App-Namespace"
	HeaderRequestId = "X-Request-ID"
	HeaderUserId    = "X-User-ID"

	TokenClaimUserID = "backend_uid"
)
