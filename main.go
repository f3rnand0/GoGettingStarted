package main

import (
	"net/http"

	"github.com/f3rnand0/GoGettingStarted/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
