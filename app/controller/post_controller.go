package controller

import (
	"app/model"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	books := model.GetAll()
	c.HTML(200, "index.html", gin.H{"books": books})
}
