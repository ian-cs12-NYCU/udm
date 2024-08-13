package sbi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) getUEIDRoutes() []Route {
	return []Route{
		{
			"Index",
			"GET",
			"/",
			s.HandleIndex,
		},

		{
			"Deconceal",
			"Post",
			"/deconceal",
			s.HandleDeconceal,
		},
	}
}

func (s *Server) HandleDeconceal(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{})
}
