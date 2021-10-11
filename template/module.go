// Module      {{ Name }}
// Description {{ Description }}
// License     {{ License }}
// Copyright   {{ Year }} by {{ Author }}



package {{ lower Name }}

import (
	libservice "github.com/4thel00z/libservice/v1"
)

type {{ title Name }} struct{}

var (
	Module = {{ title Name }}{}
)

func (d {{ title Name }}) Version() string {
	return "v1"
}

func (d {{ title Name }}) Namespace() string {
	return "debug"
}

func (d {{ title Name }}) Routes() map[string]libservice.Route {
	// Add route definitions here
	return map[string]libservice.Route{
		"routes": {
			Path:        "routes",
			Method:      "GET",
			CurlExample: "curl http://<addr>/<version>/<namespace>/<path>",
			Service:     GetRoutesHandler,
		},
	}
}

func (d {{ title Name }}) LongPath(route libservice.Route) string {
	return libservice.DefaultLongPath(d, route)
}
