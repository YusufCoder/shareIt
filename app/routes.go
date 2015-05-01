package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	Secure		bool
}

type Routes []Route

var routes = Routes{
	Route{
		"GetIndex",
		"GET",
		"/",
		GetIndex,
		false,
	},
	Route{
		"GetAdmin",
		"GET",
		"/admin",
		GetAdmin,
		true,
	},
	Route{
		"GetRoom",
		"GET",
		"/room",
		GetRoom,
		false,
	},
	Route{
		"GetCSS",
		"GET",
		"/css/{filename}",
		GetCSS,
		false,
	},
	Route{
		"GetJS",
		"GET",
		"/js/{filename}",
		GetJS,
		false,
	},
	Route{
		"GetFavicon",
		"GET",
		"/favicon.ico",
		GetFavicon,
		false,
	},
	Route{
		"GetApiRoomId",
		"GET",
		"/api/room/{id}",
		GetApiRoomId,
		false,
	},
	Route{
		"PostApiRoom",
		"POST",
		"/api/room",
		PostApiRoom,
		true,
	},
	Route{
		"GetApiFileName",
		"GET",
		"/api/room/{roomid}/file/{filename}",
		GetApiFileName,
		false,
	},
	Route{
		"DeleteApiFileName",
		"DELETE",
		"/api/room/{roomid}/file/{filename}",
		DeleteApiFileName,
		false,
	},
	Route{
		"PostApiFile",
		"POST",
		"/api/room/{roomid}/file",
		PostApiFile,
		false,
	},
}
