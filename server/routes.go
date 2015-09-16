package server

import (
	examples "modern-web-apps-in-go/server/examples/routes"
	home "modern-web-apps-in-go/server/home/routes"
	"modern-web-apps-in-go/server/router"
)

func Routes() router.Routes {
	routes := home.Routes()
	routes = append(routes, examples.Routes()...)
	return routes
}
