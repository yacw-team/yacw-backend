package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/yacw-team/yacw/models"
	"github.com/yacw-team/yacw/utils"
	"net/http"
)

// DeletePrompt 删除用户创建的prompt
func DeletePrompt(c *gin.Context) {

	uid := utils.HashAndSalt(c.PostForm("apiKey")) //用户id
	id := c.PostForm("promptsId")

	err := utils.DB.Table("prompt").Where("id = ? AND uid = ?", id, uid).Delete(models.Prompt{}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, "NO")
		return
	}
	c.JSON(http.StatusOK, "OK")
}
