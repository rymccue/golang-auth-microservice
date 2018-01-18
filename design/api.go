package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("Authentication API", func() {
	Title("The Authentication API")
	Description("An API which an Authentication API")
	Contact(func() {
		Name("Tester Test")
		Email("test@tester.com")
	})
	Host("localhost:8080")
	Scheme("http")
	BasePath("/api/")
	Origin("*", func() {
		Headers("Content-Type")
		Methods("GET", "POST", "PATCH", "DELETE", "PUT", "OPTION")
	})
	Consumes("application/json")
	Produces("application/json")
})
