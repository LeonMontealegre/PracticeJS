package main

import (
    "github.com/LeonMontealegre/PracticeJS/site/app/handlers"
    "github.com/gin-gonic/gin"
    "log"
)

func main() {
    router := gin.Default()

    router.Use(gin.Logger())
    router.Use(gin.Recovery())

    router.Static("/css", "../public/css")
    router.Static("/img", "../public/img")
    router.Static("/js",  "../public/js")

    router.LoadHTMLGlob("../templates/*")

    router.GET("/", handlers.IndexHandler)

    for true {
        err := router.Run("127.0.0.1:8081")
        if err != nil {
            log.Printf("Web server crashed!\n %v", err)
        }
    }
}
