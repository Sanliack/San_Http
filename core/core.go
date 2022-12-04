package core

import (
	"net/http"
	"san_http/context"
)

type HandleFunc func(c *context.Context)

type Engine struct {
	route *Router
}

func (e *Engine) AddRoute(method string, pattern string, handle HandleFunc) {
	e.route.addRoute(method, pattern, handle)
}

func (e *Engine) Get(pattern string, handle HandleFunc) {
	e.route.addRoute("GET", pattern, handle)
}

func (e *Engine) Post(pattern string, handle HandleFunc) {
	e.route.addRoute("POST", pattern, handle)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := context.NewContext(w, req)
	e.route.handle(c)
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

func NewEngine() *Engine {
	return &Engine{
		route: NewRouter(),
	}
}
