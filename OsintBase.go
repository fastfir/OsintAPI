package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLFiles("/templates/index.html")
	router.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{"placeholder": 1})
	})
	router.GET("/usersearch", func(Context *gin.Context) {
		querystring, ok := Context.GetQuery("user")
		if !ok {
			go UserSearch(querystring)
		}
	})
	gin.SetMode(gin.ReleaseMode)
	router.Run(":8080")
}
