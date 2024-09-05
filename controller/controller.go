package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func (m *Module) getUserId(ginCtx *gin.Context) int {
	uId, _ := ginCtx.Cookie("uid")
	if uId == "" {
		return 0
	}
	userId, err := strconv.Atoi(uId)
	if err != nil {
		return 0
	}
	return userId
}
