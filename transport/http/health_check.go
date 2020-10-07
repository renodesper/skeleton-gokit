package http

import (
	"context"
	"net/http"
)

// Request Decoder for health check service
func decodeHealthCheckRequest(_ context.Context, r *http.Request) (interface{}, error) {
	/*
		What we usually do in here:
		- Get query params
		- Create an instance of struct to be return to endpoint

		Example:

		id := r.URL.Query().Get("id")
		name := r.URL.Query().Get("name")

		userRequest := endpoint.UserRequest{
			ID: id,
			Name: name,
		}

		return userRequest, nil
	*/

	return nil, nil
}
