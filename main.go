package main

import (
	"projects/GinFramework/gin-todolist/handlers"

	"github.com/gin-gonic/gin"
)

func main(){
	r:=gin.Default()
	r.LoadHTMLGlob("templates/*.html")

	r.GET("/",handlers.Index)
	r.GET("/todo",handlers.RenderHTML)
	r.POST("/add",handlers.AddItem)

	r.Run()
}