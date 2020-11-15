package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jbegley1995/twilio-autobots/server/controllers/search"
)

//Build builds out all of the API controllers for the application without cluttering up main.
func Build(router *gin.Engine) {
	v1 := router.Group("/api/v1")

	// Build out all of the controllers for the API
	search.Build(v1)
}
