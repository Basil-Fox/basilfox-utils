package constant

const (
	ContextTokenClaims    = "token_claims"
	ContextLogger         = "logger"
	ContextNamespace      = "namespace"
	ContextRequestId      = "request_id"
	ContextUserId         = "user_id"
	ContextFirebaseUserId = "firebase_user_id"

	EndpointPublic  = "Public"
	EndpointPrivate = "Private"

	HeaderNamespace      = "X-App-Namespace"
	HeaderRequestId      = "X-Request-ID"
	HeaderUserId         = "X-User-ID"
	HeaderFirebaseUserId = "X-Firebase-User-ID"

	TokenTypeAccess  = "access"
	TokenTypeRefresh = "refresh"
)
