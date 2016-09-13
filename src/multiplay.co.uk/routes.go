package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"Regions",
		"GET",
		"/regions",
		AllRegions,
	},
	Route{
		"Region",
		"GET",
		"/region/{id:[0-9]+}",
		RegionShow,
	},
	Route{
		"Providers",
		"GET",
		"/region/{id:[0-9]+}/providers",
		AllProviders,
	},
	Route{
		"DataCenters",
		"GET",
		"/region/{id:[0-9]+}/datacenters",
		AllDataCenters,
	},
	Route{
		"Machines",
		"GET",
		"/region/{id:[0-9]+}/machines",
		AllMachines,
	},
	Route{
		"Machines",
		"GET",
		"/region/{id:[0-9]+}/machines",
		AllMachines,
	},
	Route{
		"Crashes",
		"GET",
		"/region/{id:[0-9]+}/crashes",
		AllCrashes,
	},
	Route{
		"PSUs",
		"GET",
		"/region/{id:[0-9]+}/psu",
		AllPSUs,
	},
	Route{
		"Auth",
		"GET",
		"/auth",
		CallbackHandler,
	},
}
