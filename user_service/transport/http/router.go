package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Config struct {
	Host string
	Port string
}

type Router struct {
	config Config
	Router *mux.Router
	Handler Handler
}

func NewRouter(cfg Config, h *Handler) *Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/test", h.HealthCheck).Methods("POST")
	return &Router{
		config: cfg,
		Router: r,
	}
}

func (r *Router) Run() {
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", r.config.Host, r.config.Port), nil))
}