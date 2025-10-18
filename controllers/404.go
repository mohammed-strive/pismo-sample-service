package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type NoRouteController interface {
	NoRouteHandler(c *gin.Context)
}

type noRouteController struct {
}

func NewNoRouteController() NoRouteController {
	return noRouteController{}
}

func (n noRouteController) NoRouteHandler(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"error": "invalid route"})
}
