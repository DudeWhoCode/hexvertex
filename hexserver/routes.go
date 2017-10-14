package hexserver

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"ping",
		"GET",
		"/ping",
		Ping,
	},
	Route{
		"redirectUser",
		"GET",
		"/api/twitter-login",
		RedirectUserToTwitter,
	},
	Route{
		"getToken",
		"GET",
		"/maketoken",
		GetTwitterToken,
	},
	Route{
		"login",
		"GET",
		"/",
		Login,
	},
}