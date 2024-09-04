package controller

import (
	"github.com/eolinker/apinto-dashboard/controller"
	"github.com/gin-gonic/gin"
)

func (m *Module) getUserId(ginCtx *gin.Context) int {
	userName := ginCtx.GetString(controller.UserName)
	if userName == "" {
		return 0
	}
	info, err := m.userInfoService.GetUserInfoByName(ginCtx, userName)
	if err != nil {
		return 0
	}
	return info.Id

}
