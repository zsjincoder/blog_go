package util

import (
	"../setting"
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
)

func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = setting.PageSize * (page - 1)
	}
	return result
}
