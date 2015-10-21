package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"html/template"
)

func main() {
	router := gin.Default()
	html := template.Must(template.New("template").Delims("[[", "]]").ParseGlob("templates/*.tmpl"))
	router.SetHTMLTemplate(html)

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Tracker",
		})
	})

	router.Static("public/", "./public")
	router.Run(":8080")
}
