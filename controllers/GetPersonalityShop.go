package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/yacw-team/yacw/models"
	"github.com/yacw-team/yacw/utils"
	"net/http"
)

// GetPersonalityShop 获取人格商店的内容
func GetPersonalityShop(c *gin.Context) {
	var personality []models.Personality
	utils.DB.Table("personality").Where("designer = ?", 0).Find(&personality)

	c.JSON(http.StatusOK, personality)
}
