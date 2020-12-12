package design

import (
	. "goa.design/goa/v3/dsl"
)

var AuthenticateDiscourseFields = Type("AuthenticateDiscourseFields", func() {
	Attribute("sso", String)
	Attribute("sig", String)
	Required("sso", "sig")

	Attribute("email", String)
	Attribute("password", String, func() {
		MinLength(10)
	})
})

var _ = Service("discourse", func() {
	Error("user-unverified", func() {
	})

	Method("authenticate", func() {
		Payload(func() {
			Attribute("token", String)
			Attribute("login", AuthenticateDiscourseFields)
			Required("login")
		})

		Result(func() {
			Attribute("location", String)
			Required("location")
		})

		HTTP(func() {
			POST("discourse/auth")

			Header("token:Authorization", String, "authentication token", func() {
				Pattern("^Bearer [^ ]+$")
			})

			Body("login")

			Response(StatusTemporaryRedirect, func() {
				Headers(func() {
					Header("location:Location")
				})
			})

			Response("user-unverified", StatusForbidden)
		})
	})

	commonOptions()
})
