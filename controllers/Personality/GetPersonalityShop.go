package Personality

import (
	"github.com/gin-gonic/gin"
	"github.com/yacw-team/yacw/models"
	"github.com/yacw-team/yacw/utils"
	"net/http"
)

// GetPersonalityShop 获取人格商店的内容
func GetPersonalityShop(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "2007"})
			// 进行适当的处理
		}
	}()
	var personality []models.Personality
	err := utils.DB.Table("personality").Where("designer = ?", 0).Find(&personality).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "3009"})
		return
	}

	c.JSON(http.StatusOK, personality)
}
