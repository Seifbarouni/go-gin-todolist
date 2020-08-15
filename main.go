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

	r.GET("/delete/:id",handlers.DeleteItem)
	r.GET("/update/:id/:todo",handlers.RenderUpdateHTML)
	r.POST("/update",handlers.UpdateItem)
	
	r.Run()
}