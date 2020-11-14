package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jbegley1995/twilio-autobots/src/controllers"
)

func main() {
	r := gin.Default()

	controllers.Build(r)

	r.Run()
}
