package server

import (
	"api_service/internal/routes"
	"net/http"
)

func NewServer() http.Handler {
	mux := http.NewServeMux()
	routes.AddRoutes(mux)
	var handler http.Handler = mux
	return handler
}
