package main

import "net/http"

//Route a route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	IsSecured   bool
	HandlerFunc http.HandlerFunc
}

//Routes route list
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		false,
		Index,
	},
	Route{
		"UserList",
		"GET",
		"/users",
		true,
		UserList,
	},
	Route{
		"User",
		"GET",
		"/users/{userId}",
		true,
		User,
	},
	Route{
		"UserCreate",
		"POST",
		"/users",
		true,
		UserCreate,
	},
	Route{
		"Login",
		"GET",
		"/login/{login}/{password}",
		false,
		LoginUser,
	},
}
