package main

import (
	"healthstats/pkg/router"
	"net/http"
)

func main() {

	r := router.NewRouter()
	http.ListenAndServe(":9000", r)
}
