package constant

const (
	ContextJwtToken  = "jwt_token"
	ContextRequestId = "request_id"

	EndpointPublic  = "Public"
	EndpointPrivate = "Private"
	EndpointRefresh = "Refresh"

	HeaderNamespace = "X-App-Namespace"
	HeaderRequestId = "X-Request-ID"
	HeaderTokenId   = "X-Token-ID"
	HeaderTokenKind = "X-Token-Kind"
	HeaderUserId    = "X-User-ID"

	TokenTypeAccess  = "access"
	TokenTypeRefresh = "refresh"
)
