package handlers

import (
	"net/http"
	"projects/GinFramework/gin-todolist/models"

	"github.com/gin-gonic/gin"
)



func RenderHTML(c *gin.Context) {
	data:=models.GetTodos()
	c.HTML(200, "index.html", data)
}

func Index(c *gin.Context){
	c.Redirect(http.StatusFound,"/todo")
}

func AddItem(c *gin.Context){

	item:=models.Todo{}
	item.Todostr=c.PostForm("todo")
	models.AddTodo(&item)
	c.Redirect(http.StatusFound,"/todo")
	
}