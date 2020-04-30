package design

import (
	. "goa.design/goa/v3/dsl"
)

var Owner = Type("StationOwner", func() {
	Attribute("id", Int32)
	Attribute("name", String)
	Required("id", "name")
})

var StationFullImageRef = Type("ImageRef", func() {
	Attribute("url", String)
	Required("url")
})

var StationFullPhotos = Type("StationPhotos", func() {
	Attribute("small", String)
	Required("small")
})

var StationUpload = Type("StationUpload", func() {
	Attribute("id", Int64)
	Attribute("time", Int64)
	Attribute("upload_id", String)
	Attribute("size", Int64)
	Attribute("url", String)
	Attribute("type", String)
	Attribute("blocks", ArrayOf(Int64))
	Required("id", "time", "upload_id", "size", "url", "type", "blocks")
})

var StationSensor = Type("StationSensor", func() {
	Attribute("name", String)
	Attribute("unit_of_measure", String)
	Required("name", "unit_of_measure")
})

var StationModule = Type("StationModule", func() {
	Attribute("id", String)
	Attribute("name", String)
	Attribute("position", Int32)
	Attribute("sensors", ArrayOf(StationSensor))
	Required("id", "name", "position", "sensors")
})

var StationFull = ResultType("application/vnd.app.station.full", func() {
	TypeName("StationFull")
	Attributes(func() {
		Attribute("id", Int32)
		Attribute("name")
		Attribute("owner", Owner)
		Attribute("device_id", String)
		Attribute("uploads", ArrayOf(StationUpload))
		Attribute("images", ArrayOf(StationFullImageRef))
		Attribute("photos", StationFullPhotos)
		Attribute("read_only", Boolean)
		Required("id", "name", "owner", "device_id", "uploads", "images", "photos", "read_only")

		Attribute("battery", Float32)
		Attribute("recording_started_at", Int64)
		Attribute("memory_used", Int32)
		Attribute("memory_available", Int32)
		Attribute("firmware_number", Int32)
		Attribute("firmware_time", Int32)
		Attribute("modules", ArrayOf(StationModule))
		Required("battery", "recording_started_at", "memory_used", "memory_available", "firmware_number", "firmware_time", "modules")
	})
	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("owner")
		Attribute("device_id")
		Attribute("uploads")
		Attribute("images")
		Attribute("photos")
		Attribute("read_only")

		Attribute("battery")
		Attribute("recording_started_at")
		Attribute("memory_used")
		Attribute("memory_available")
		Attribute("firmware_number")
		Attribute("firmware_time")
		Attribute("modules")
	})
})

var _ = Service("station", func() {
	Method("station", func() {
		Security(JWTAuth, func() {
			Scope("api:access")
		})

		Payload(func() {
			Token("auth")
			Required("auth")
			Attribute("id", Int32)
			Required("id")
		})

		Result(StationFull)

		HTTP(func() {
			GET("stations/@/{id}")

			httpAuthentication()
		})
	})

	Method("update", func() {
		Security(JWTAuth, func() {
			Scope("api:access")
		})

		Payload(func() {
			Token("auth")
			Required("auth")
			Attribute("id", Int32)
			Required("id")
			Attribute("name", String)
			Attribute("status_json", MapOf(String, Any))
			Required("name", "status_json")
		})

		Result(StationFull)

		HTTP(func() {
			POST("stations/{id}")

			httpAuthentication()
		})
	})

	Error("unauthorized", String, "credentials are invalid")
	Error("not-found", String, "not found")

	HTTP(func() {
		Response("unauthorized", StatusUnauthorized)
		Response("not-found", StatusNotFound)
	})

	commonOptions()
})
