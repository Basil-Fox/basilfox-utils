package constant

const (
	ContextJwtToken  = "jwt_token"
	ContextLogger    = "logger"
	ContextNamespace = "namespace"
	ContextRequestId = "request_id"
	ContextTokenId   = "token_id"
	ContextTokenKind = "token_kind"
	ContextUserId    = "user_id"

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
