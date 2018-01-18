package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var Token = MediaType("application/vnd.token+json", func() {
	Description("A token")
	Attributes(func() {
		Attribute("token", String, "A JWT token")
	})

	View("default", func() {
		Attribute("token")
	})
})

var User = MediaType("application/vnd.user+json", func() {
	Description("A user")
	Attributes(func() {
		Attribute("id", Integer, "ID of account", func() {
			Example(1)
		})
		Attribute("email", String, "Email of the user", func() {
			Format("email")
			Example("bob@gmail.com")
		})
		Attribute("first_name", String, "First name of the user", func() {
			Example("John")
		})
		Attribute("last_name", String, "Last name of the user", func() {
			Example("Snow")
		})
		Attribute("password", String, "Avatar of user", func() {
			Example("password")
		})
		Attribute("salt", String, "Phone number of the user", func() {
			Example("salt")
		})
	})

	View("default", func() {
		Attribute("id")
		Attribute("email")
		Attribute("first_name")
		Attribute("last_name")
		Attribute("password")
		Attribute("salt")
	})
})
