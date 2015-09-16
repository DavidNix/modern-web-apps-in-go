package routes

import (
	"modern-web-apps-in-go/server/home"
	"modern-web-apps-in-go/server/router"
)

func Routes() router.Routes {
	return router.Routes{
		{
			"home_root",
			"GET", "/", home.Hello(),
		}, {
			"home_goodbye",
			"GET", "/goodbye", home.Goodbye(),
		},
	}
}
