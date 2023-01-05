package service

import (
	. "goa.design/goa/v3/dsl"
)

// JWTAuth defines a security scheme that uses JWT tokens.
var JWTAuth = JWTSecurity("jwt", func() {
	Description(`Secures endpoint by requiring a valid JWT token retrieved via the signin endpoint`)
	Scope("api:read")  // Enforce presence of both "api:read"
	Scope("api:write") // and "api:write" scopes in OAuth2 claims.
})

// APIKeyAuth defines a security scheme that uses API keys.
var APIKeyAuth = APIKeySecurity("api_key", func() {
	Description("Secures endpoint by requiring an API key.")
})

// BasicAuth defines a security scheme using basic authentication. The scheme
// protects the "signin" action used to create JWTs.
var BasicAuth = BasicAuthSecurity("basic", func() {
	Description("Basic authentication used to authenticate security principal during signin")
})

var OAuth2 = OAuth2Security("googauth", func() {
	ImplicitFlow("/authorization", "/index")
	Scope("api:write", "Write acess")
	Scope("api:read", "Read access")
})
