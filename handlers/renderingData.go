package handlers

import (
	"net/http"
	"projects/GinFramework/gin-todolist/models"
	"strconv"

	"github.com/gin-gonic/gin"
)



func RenderHTML(c *gin.Context) {
	data:=models.GetTodos()
	c.HTML(200, "index.html", data)
}

func Index(c *gin.Context){
	c.Redirect(http.StatusFound,"/todo")
}

var id int
var err error
func RenderUpdateHTML(c *gin.Context){
	idstr:=c.Param("id")
	todo:=c.Param("todo")
	id,err=strconv.Atoi(idstr)
	

	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"Error":"ID invalid",
		})
		return
	}
	item:=models.Todo{
		Todostr: todo,
	}
	c.HTML(200,"update.html",item)

}

