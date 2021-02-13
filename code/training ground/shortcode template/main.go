package main

import (
	"net/http"
	"shortcodes/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)

}
