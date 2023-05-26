package Personality

import (
	"github.com/gin-gonic/gin"
	"github.com/yacw-team/yacw/models"
	"github.com/yacw-team/yacw/utils"
	"net/http"
	"strings"
)

// CreatePersonality 用户创建personality
func CreatePersonality(c *gin.Context) {

	var err error
	var reqBody map[string]interface{}
	reqTemp, ok := c.Get("reqBody")
	if !ok {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "2006"})
		return
	}
	reqBody = reqTemp.(map[string]interface{})

	uid := reqBody["apiKey"].(string)
	name := reqBody["name"].(string)
	description := reqBody["description"].(string)
	prompts := reqBody["prompt"].(string)

	//检测utf-8编码
	slice := []string{uid, name, description, prompts}
	if !utils.Utf8Check(slice) {
		c.JSON(http.StatusBadRequest, models.ErrCode{ErrCode: "1011"})
		return
	}

	uid, err = utils.Encrypt(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "3006"})
		return
	}
	if len(strings.TrimSpace(name)) == 0 {
		c.JSON(http.StatusBadRequest, models.ErrCode{ErrCode: "1007"})
		return
	}

	if len(strings.TrimSpace(prompts)) == 0 {
		c.JSON(http.StatusBadRequest, models.ErrCode{ErrCode: "1007"})
		return
	}

	personality := models.Personality{
		Uid:         uid,
		ModelName:   name,
		Description: description,
		Prompts:     prompts,
	}

	err = utils.DB.Table("personality").Create(&personality).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrCode{ErrCode: "3009"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id":          personality.Id,
		"name":        personality.ModelName,
		"description": personality.Description,
		"prompts":     personality.Prompts,
	})
}
