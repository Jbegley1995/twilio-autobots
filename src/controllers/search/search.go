package search

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jbegley1995/twilio-autobots/src/services"
)

//Build builds out all of the API controllers for search.
func Build(router *gin.RouterGroup) {
	searchController := router.Group("/search")

	searchController.GET("/", GetStarsByOrigin)
}

//GetStarsByOrigin searches github for a repo and returns the stars.
func GetStarsByOrigin(c *gin.Context) {
	var (
		query = c.Request.URL.Query()
	)

	service := services.Github()
	var stars = map[string]int{}
	for _, origin := range query["origin"] {
		repo, err := service.GetByOrigin(origin)
		if err != nil {
			fmt.Println(err)
			return
		}

		stars[*repo.Name] = *repo.StargazersCount
	}

	c.JSON(http.StatusOK, gin.H{"data": stars})
}
