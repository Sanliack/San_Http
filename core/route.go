package core

import (
	"log"
	"net/http"
	"san_http/context"
)

type Router struct {
	handlers map[string]HandleFunc
}

func NewRouter() *Router {
	return &Router{
		handlers: map[string]HandleFunc{},
	}
}

func (r *Router) addRoute(method string, pattern string, handle HandleFunc) {
	log.Printf("Route %4s - %s\n", method, pattern)
	key := method + "-" + pattern
	r.handlers[key] = handle
}

func (r *Router) handle(c *context.Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT Found: %s\n", c.Path)
	}
}
