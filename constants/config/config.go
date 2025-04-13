package config

type RouteType string

const (
	ContextTokenClaims = "token_claims"
	ContextUserID      = "userID"
	ContextFirebaseUID = "firebaseUID"
	ContextGroupID     = "groupID"

	FirebaseCustomUID = "backend_uid"

	RoutePublic   RouteType = "Public"
	RouteInternal RouteType = "Internal"
	RoutePrivate  RouteType = "Private"
	RouteRegister RouteType = "Register"

	GuestEmailSuffix = "@billsplit_guest"
)
