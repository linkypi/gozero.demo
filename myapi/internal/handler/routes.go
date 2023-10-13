// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"github.com/zeromicro/go-zero/rest"
	"net/http"
)

type Routers struct {
	server     *rest.Server
	middleware []rest.Middleware
	ropts      []rest.RouteOption
}

func NewRouters(server *rest.Server) *Routers {
	return &Routers{
		server: server,
	}
}

func (r *Routers) Get(path string, handler http.HandlerFunc) {
	r.server.AddRoutes(
		rest.WithMiddlewares(
			r.middleware,
			rest.Route{
				Method:  http.MethodGet,
				Path:    path,
				Handler: handler,
			},
		), r.ropts...,
	)
}

func (r *Routers) Post(path string, handler http.HandlerFunc) {
	r.server.AddRoutes(
		rest.WithMiddlewares(
			r.middleware,
			rest.Route{
				Method:  http.MethodPost,
				Path:    path,
				Handler: handler,
			},
		), r.ropts...,
	)
}

func (r *Routers) Group(ops ...rest.RouteOption) *Routers {
	router := &Routers{
		server: r.server,
		ropts:  ops,
	}
	return router
}

func (r *Routers) AddOpt(ops ...rest.RouteOption) {
	r.ropts = append(r.ropts, ops...)
}
