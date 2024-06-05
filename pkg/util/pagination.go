package util

import (
	"gin-test/pkg/setting"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.PostForm("page")).Int()
	if page > 0 {
		result = (page - 1) * setting.AppSetting.PageSize
	}
	return result
}