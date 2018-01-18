package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("authentication", func() {
	BasePath("/auth")
	NoSecurity()
	Action("login", func() {
		NoSecurity()
		Routing(
			POST("/login"),
		)
		Description("Sign a user in")
		Payload(LoginPayload)
		Response(OK, Token)
		Response(InternalServerError)
		Response(BadRequest, ErrorMedia)
	})

	Action("register", func() {
		NoSecurity()
		Routing(
			POST("/register"),
		)
		Description("Create a new user")
		Payload(RegisterPayload)
		Response(OK, Token)
		Response(InternalServerError)
		Response(BadRequest, ErrorMedia)
	})

})

var _ = Resource("swagger", func() {
	Files("/swagger.json", "swagger/swagger.json")
})
