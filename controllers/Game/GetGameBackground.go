package Game

import (
	"github.com/gin-gonic/gin"
	"github.com/yacw-team/yacw/models"
	"github.com/yacw-team/yacw/utils"
	"net/http"
)

func GetGameBackground(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "2007"})
			// 进行适当的处理
		}
	}()
	var gameArray []models.Game

	err := utils.DB.Table("game").Find(&gameArray).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "3009"})
		return
	}
	c.JSON(http.StatusOK, gameArray)
}
