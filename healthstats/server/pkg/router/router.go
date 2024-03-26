package router

import (
	"healthstats/pkg/handler"
	"healthstats/pkg/service"
	"net/http"
)

type router struct {
	router  *http.ServeMux
	service *service.Service
}

func NewRouter(service *service.Service) *http.ServeMux {
	r := &router{router: http.NewServeMux(), service: service}

	r.setupRoutes()

	return r.router
}

func (r *router) setupRoutes() {
	fileHandler := handler.NewFileHandler(r.service)
	fileHandler.InitRoutes(r.router)
}
