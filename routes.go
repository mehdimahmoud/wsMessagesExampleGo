package main

import "net/http"

/*
 * Route structure
 * @field Name string, name of the route
 * @field Method string, type of the CRUD operation
 * @field Pattern string, route pattern from URL ("/", ...)
 * @field HandlerFunc http.HandlerFunc, name of the handler function
 */
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

/*
 * Routes slice with elements of type Route'structure
 */
type Routes []Route

/*
 *
 */
var routes = Routes{
	Route{
		"IndexHandler",
		"GET",
		"/",
		IndexHandler,
	},
	Route{
		"IndexMessagesHandler",
		"GET",
		"/messages/test",
		IndexMessagesHandler,
	},
	Route{
		"MessagesHandler",
		"GET",
		"/messages",
		MessagesHandler,
	},
	Route{
		"GetMessageHandler",
		"GET",
		"/messages/{msgID}",
		GetMessageHandler,
	},
	Route{
		"StoreMessageHandler",
		"POST",
		"/messages/",
		StoreMessageHandler,
	},
}
