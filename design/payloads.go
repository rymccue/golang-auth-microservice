package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var LoginPayload = Type("LoginPayload", func() {
	Attribute("email", String, func() {
		MinLength(6)
		MaxLength(400)
		Format("email")
		Example("jamesbond@gmail.com")
	})

	Attribute("password", String, func() {
		MinLength(5)
		MaxLength(100)
		Example("abcd1234")
	})
	Required("email", "password")
})

var RegisterPayload = Type("RegisterPayload", func() {
	Attribute("email", String, func() {
		MinLength(6)
		MaxLength(150)
		Format("email")
		Example("jamesbond@gmail.com")
	})

	Attribute("first_name", String, func() {
		MinLength(1)
		MaxLength(200)
		Example("John")
	})

	Attribute("last_name", String, func() {
		MinLength(1)
		MaxLength(200)
		Example("Doe")
	})

	Attribute("password", String, func() {
		MinLength(5)
		MaxLength(100)
		Example("abcd1234")
	})

	Required("email", "password", "first_name", "last_name")
})
