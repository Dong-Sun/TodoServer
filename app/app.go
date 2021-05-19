package app

import (
	"log"
	"net/http"
	"project/TodoServer/model"
	"strconv"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {

	c.HTML(http.StatusOK, "todo.html", nil)
	//	c.Redirect(http.StatusTemporaryRedirect, "/todo.html")

}

func GetHandler(c *gin.Context) {
	temp := model.Read()
	c.JSON(http.StatusOK, temp)
}

func PostHandler(c *gin.Context) {
	c.Request.ParseForm()
	msg := c.PostForm("sText")
	model.Create(msg)
}

func PutHandler(c *gin.Context) {
	c.Request.ParseForm()
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		log.Println("PutHandler Atoi(c.PostForm()) err:", err)
	}
	model.Update(id)
}

func DeleteHandler(c *gin.Context) {
	c.Request.ParseForm()
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		log.Println("DeleteHandler Atoi(c.PostForm()) err:", err)
	}
	model.Delete(id)
}

func Run() {
	r := gin.Default()

	r.LoadHTMLGlob("./templates/*")
	r.Use(static.Serve("/", static.LocalFile("./templates/", false)))

	r.GET("/", IndexHandler)

	r.GET("/todo", GetHandler)
	r.POST("/todo", PostHandler)
	r.PUT("/todo", PutHandler)
	r.DELETE("/todo", DeleteHandler)

	r.Run(":8080")
}
