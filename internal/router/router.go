package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Router struct {
}

func NewRouter() Router {
	return Router{}
}

func (r *Router) GetHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
