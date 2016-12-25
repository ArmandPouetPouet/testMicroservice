package main

import "net/http"

//Route a route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes route list
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"UserList",
		"GET",
		"/users",
		UserList,
	},
	Route{
		"User",
		"GET",
		"/users/{userId}",
		User,
	},
	Route{
		"UserCreate",
		"POST",
		"/users",
		UserCreate,
	},
}
