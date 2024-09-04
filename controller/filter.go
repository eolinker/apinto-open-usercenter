package controller

import (
	apinto_module "github.com/eolinker/apinto-dashboard/module"
	"github.com/gin-gonic/gin"
)

func HandleAddKey(key string, value any) gin.HandlerFunc {
	return func(ginCtx *gin.Context) {
		apinto_module.AddKey(ginCtx, key, value)
	}
}
