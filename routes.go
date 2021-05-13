package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func initializeRoutes() {

	// Handle the index route
	router.GET("/", showIndexPage)

	  // Handle GET requests at /article/view/some_article_id
	router.GET("/article/view/:article_id", getArticle)
}
// Render one of HTML, JSON or CSV based on the 'Accept' header of the request
// If the header doesn't specify this, HTML is rendered, provided that
// the template name is present
func render(c *gin.Context, data gin.H, templateName string) {

	switch c.Request.Header.Get("Accept") {
	case "application/json":
	  // Respond with JSON
	  c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
	  // Respond with XML
	  c.XML(http.StatusOK, data["payload"])
	default:
	  // Respond with HTML
	  c.HTML(http.StatusOK, templateName, data)
	}
  
  }