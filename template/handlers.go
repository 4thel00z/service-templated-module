package {{ lower Name }}

import (
	"github.com/4thel00z/libhttp"
	libservice "github.com/4thel00z/libservice/v1"
)

func GetHandlerExample(app libservice.App) libhttp.Service {
	return func(req libhttp.Request) libhttp.Response {

		response := req.Response(map[string]string{})
		response.StatusCode = 200
		return response
	}
}
