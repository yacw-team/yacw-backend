package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/yacw-team/yacw/models"
	"github.com/yacw-team/yacw/utils"
	"net/http"
)

func GetMyPersonality(c *gin.Context) {
	var personality []models.Personality
	utils.DB.Table("personality").Where("designer = ?", 1).Find(&personality)
	c.JSON(http.StatusOK, personality)
}
