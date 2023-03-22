package controller

import (
	"app/model"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	books := model.GetAll()
	c.HTML(200, "index.html", gin.H{"books": books})
}

func Show(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	book := model.GetOne(id)
	c.HTML(200, "show.html", gin.H{"book": book})
}

func GetCreate(c *gin.Context) {
	c.HTML(200, "create.html", gin.H{})
}

func PostCreate(c *gin.Context) {
	title := c.PostForm("title")
	body := c.PostForm("body")
	book := model.Book{Title: title, Body: body}
	book.Create()
	c.Redirect(301, "/")
}

func GetEdit(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	book := model.GetOne(id)
	c.HTML(200, "edit.html", gin.H{"book": book})
}

func PostEdit(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	book := model.GetOne(id)
	title := c.PostForm("title")
	book.Title = title
	body := c.PostForm("body")
	book.Body = body
	book.Update()
	c.Redirect(301, "/")
}

func GetDelete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	book := model.GetOne(id)
	c.HTML(200, "delete.html", gin.H{"book": book})
}

func PostDelete(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	fmt.Printf("ここに削除用のidを入れる%d", id)
	book := model.GetOne(id)
	book.Delete()
	c.Redirect(301, "/")
}
