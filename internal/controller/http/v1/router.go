package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter(handler *gin.Engine) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	h := handler.Group("/v1")
	{
		h.GET("/", index)
	}
}

type serverInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func index(c *gin.Context) {
	info := serverInfo{
		Name:    "API",
		Version: "1.0",
	}

	c.JSON(http.StatusOK, info)
}
