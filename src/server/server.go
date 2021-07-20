package server

import (
	"fmt"

	"github.com/ThalesLoreto/ogavirt-backend/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type WebServer struct {
	port   uint16
	server *echo.Echo
}

func NewServer(port uint16) *WebServer {
	// Intance of a WebServer
	instance := WebServer{
		port:   port,
		server: echo.New(),
	}

	return &instance
}

func (ws *WebServer) Start() {
	// Logger
	ws.server.Use(middleware.Logger())
	// Recover
	ws.server.Use(middleware.Recover())
	// CORS
	ws.server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.DELETE},
	}))
	// Endpoints
	ws.server.GET("/hotels", controller.GetAllHotel)
	ws.server.POST("/hotels", controller.CreateHotel)
	// Start server
	ws.server.Logger.Fatal(ws.server.Start(fmt.Sprint(":", ws.port)))
}
