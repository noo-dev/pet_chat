package server

import (
	"strconv"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
	"github.com/noo-dev/pet_chat/internal/config"
)

type Server struct {
	config config.HTTPServerCfg
	dbConn *pgx.Conn
	engine *echo.Echo
}

func NewServer(cfg config.HTTPServerCfg, dbConn *pgx.Conn, engine *echo.Echo) *Server {
	return &Server{}
}

func (s *Server) Start() {
	baseApi := s.engine.Group("/api")
	InitRoutes(baseApi, s.dbConn)
	if err := s.engine.Start(":" + strconv.Itoa(s.config.Port)); err != nil {
		panic(err)
	}
}
