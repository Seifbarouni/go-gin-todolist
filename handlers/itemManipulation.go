package handlers

import (
	"net/http"
	"projects/GinFramework/gin-todolist/models"
	"strconv"

	"github.com/gin-gonic/gin"
)



func AddItem(c *gin.Context){

	item:=models.Todo{}
	todo:=c.PostForm("todo")
	if todo!=""{
		item.Todostr=todo
		models.AddTodo(&item)
		c.Redirect(http.StatusFound,"/todo")
	}else{
		c.JSON(http.StatusBadRequest,gin.H{
			"Error":"Empty todo",
		})
	}
	
}

func DeleteItem(c *gin.Context){
	idstr:=c.Param("id")
	id,err:=strconv.Atoi(idstr)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"Error":"ID invalid",
		})
		return
	}
	models.DeleteTodo(id)
	c.Redirect(http.StatusFound,"/todo")
	

}
func UpdateItem(c *gin.Context){
	todo:=c.PostForm("todo")
	models.UpdateTodo(id,todo)
	c.Redirect(http.StatusFound,"/todo")
}


