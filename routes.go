package main

import "github.com/labstack/echo/v4"

func routes(e *echo.Echo) {
	e.POST("/", StartSimulationSession)
}
