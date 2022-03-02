package server

import "net/http"

type Country struct {
	Name     string `json:"name"`
	Language string `json:"language"`
}

var countries []*Country

func New(addr string) *http.Server {
	initRoutes()

	return &http.Server{
		Addr: addr,
	}
}
