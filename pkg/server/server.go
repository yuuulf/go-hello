package server

import (
	"github.com/labstack/echo"
	"github.com/yuuulf/go-hello/pkg/handler"
)

type Server struct {
	mux *echo.Echo
}

func (s *Server) Start() {
	s.mux.Logger.Fatal(s.mux.Start(":1323"))
}

func (s *Server) route() {
	s.mux.GET("/users", handler.GetUsers)
	s.mux.GET("/users/:id", handler.GetUser)
}

func New() *Server {
	s := &Server{
		mux: echo.New(),
	}
	s.route()
	return s
}
