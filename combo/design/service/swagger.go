package service

import . "goa.design/goa/v3/dsl"

var _ = Service(
	"swagger", func() {
		Description("The swagger service servers the API definition.")

		Files("/openapi.json", "gen/http/openapi.json", func() {
			Description("JSON document containing the API definition.")
		})
	},
)
