// main.go

package main

import "github.com/gin-gonic/gin"

var router *gin.Engine

func main() {

    router := gin.Default()
    router.LoadHTMLGlob("templates/*")

    // Handle Index
    router.GET("/", showIndexPage)
    // Handle GET requests at /article/view/some_article_id
    router.GET("/article/view/:article_id", getArticle)

    router.Run()

}

