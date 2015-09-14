package server

import (
	home "modern-web-apps-in-go/server/home/routes"
	"modern-web-apps-in-go/server/router"
)

func Routes() router.Routes {
	routes := home.Routes()
	return routes
}
