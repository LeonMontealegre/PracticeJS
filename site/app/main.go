package main

import (
    "html/template"
    "github.com/LeonMontealegre/PracticeJS/site/app/handlers"
    "github.com/gin-gonic/gin"
    "log"
)

func main() {
    router := gin.Default()

    router.SetFuncMap(template.FuncMap{
        "mod":  func(i, j int) bool { return i%j == 0 },
        "mod2": func(i, j int) bool { return i%j == 2 },
    })

    router.Use(gin.Logger())
    router.Use(gin.Recovery())

    router.Static("/css", "../public/css")
    router.Static("/img", "../public/img")
    router.Static("/js",  "../public/js")
    router.Static("/ace-builds",  "../public/ace-builds")

    router.LoadHTMLGlob("../templates/*")

    router.GET("/", handlers.IndexHandler)
    router.GET("/section/:section", handlers.SectionHandler)
    router.GET("/section/:section/:problem", handlers.ProblemHandler)

    for true {
        err := router.Run("127.0.0.1:8081")
        if err != nil {
            log.Printf("Web server crashed!\n %v", err)
        }
    }
}
