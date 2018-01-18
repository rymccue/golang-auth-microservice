package design

import (
	. "github.com/goadesign/goa/design/apidsl"
)

var JWT = JWTSecurity("jwt", func() {
	Description("Use JWT to authenticate")
	Header("Authorization")
})
