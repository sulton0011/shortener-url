package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/teris-io/shortid"
)

func main() {
	var count int
	r := gin.New()
	r.GET("/", func(c *gin.Context) {
		count++
		c.Redirect(http.StatusPermanentRedirect, "https://www.youtube.com/watch?v=YwnwsxE36LY&list=WL&index=6")
	})
	r.GET("/count", func(c *gin.Context) {
		s, _ := shortid.Generate()
		c.JSON(200, s)
	})
	r.Run()
}
