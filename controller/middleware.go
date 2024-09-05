package controller

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	redis "github.com/redis/go-redis/v9"

	"github.com/eolinker/apinto-dashboard/common"
	"github.com/eolinker/apinto-dashboard/controller"
	apinto_module "github.com/eolinker/apinto-dashboard/module"
	"github.com/eolinker/apinto-dashboard/pm3"
	jwt "github.com/golang-jwt/jwt/v5"
)

var loginError = "登录已失效，请重新登录"

var jwtSecret = []byte("apintp-dashboard")

var secret = []byte("1P&dG^5MceRb0T#7QDu6OtF%)$Nh@q")

func tryAbort(w apinto_module.MiddlewareResponseWriter, data interface{}, abort bool) {
	if abort {
		w.JSON(http.StatusOK, data)
		w.SetAbort(true)
	}
}

func (m *Module) ApiLogin(ctx context.Context, request *apinto_module.MiddlewareRequest, w apinto_module.MiddlewareResponseWriter) {

	_, authority := apinto_module.ReadApiInfo(request)
	isNeedLogin := authority != pm3.Anonymous

	session, _ := request.GetCookie(controller.Session)
	if session == "" {

		tryAbort(w, loginError, isNeedLogin)

		return
	}

	tokens, err := m.sessionCache.Get(ctx, session)
	if err == redis.Nil || tokens == nil {
		tryAbort(w, loginError, isNeedLogin)
		return
	}

	//1.从ginCtx的header中拿到token，没拿到报错提醒用户重新登录
	verifyToken, err := common.VerifyToken(tokens.Jwt, secret)
	if err != nil {
		tryAbort(w, loginError, isNeedLogin)
		return
	}
	//1.1拿到用户ID和过期时间 过期了重新登录

	if !verifyToken.Valid {

		tryAbort(w, loginError, isNeedLogin)
		return
	}

	w.SetHeader(controller.Authorization, tokens.Jwt)
	w.SetHeader(controller.RAuthorization, tokens.RJwt)
	claims := verifyToken.Claims.(jwt.MapClaims)
	userId, _ := strconv.Atoi(claims["userId"].(string))

	info, err := m.userInfoService.GetUserInfo(ctx, userId)
	if err != nil {

		tryAbort(w, err, isNeedLogin)
		return
	}
	_ = m.sessionCache.Set(ctx, session, tokens)
	w.Set(controller.UserName, info.UserName)

	if authority == pm3.Public {
		return
	}
	//m.AccessCheck(ctx, request, w, info.UserName)
}

func (m *Module) ModuleLogin(ctx context.Context, request *apinto_module.MiddlewareRequest, w apinto_module.MiddlewareResponseWriter) {

	session, _ := request.GetCookie(controller.Session)
	if session == "" {
		DoRedirect(w, request.Url)
		return
	}

	tokens, err := m.sessionCache.Get(ctx, session)
	if err == redis.Nil || tokens == nil {
		DoRedirect(w, request.Url)
		return
	}

	//1.从ginCtx的header中拿到token，没拿到报错提醒用户重新登录
	verifyToken, err := common.VerifyToken(tokens.Jwt, secret)
	if err != nil {

		DoRedirect(w, request.Url)
		return
	}
	//1.1拿到用户ID和过期时间 过期了重新登录
	if !verifyToken.Valid {

		DoRedirect(w, request.Url)
		return
	}

	w.SetHeader(controller.Authorization, tokens.Jwt)
	w.SetHeader(controller.RAuthorization, tokens.RJwt)
	claims := verifyToken.Claims.(jwt.MapClaims)
	userId, _ := strconv.Atoi(claims["userId"].(string))

	info, err := m.userInfoService.GetUserInfo(ctx, userId)
	if err != nil {
		DoRedirect(w, request.Url)
		return
	}
	_ = m.sessionCache.Set(ctx, session, tokens)
	w.Set(controller.UserName, info.UserName)

}
func DoRedirect(w apinto_module.MiddlewareResponseWriter, callback string) {
	redirect := "/login"

	redirect = fmt.Sprint(redirect, "?callback=", callback)

	w.SetHeader("Cache-Control", "no-store, no-cache, max-age=0, must-revalidate, proxy-revalidate")
	w.Redirect(http.StatusFound, redirect)
	w.SetAbort(true)
}

//
//func (m *Module) AccessCheck(ctx context.Context, request *apinto_module.MiddlewareRequest, w apinto_module.MiddlewareResponseWriter, userName string) {
//
//	access, _ := apinto_module.ReadApiInfo(request)
//	if access == "" {
//		return
//	}
//	if strings.HasPrefix(request.FulPath, "/api/my/") {
//		return
//	}
//
//	var err error
//	userInfo, err := m.userInfoService.GetUserInfoByName(ctx, userName)
//	if err != nil {
//
//		w.JSON(http.StatusOK, controller.NewNoAccessError(err.Error()))
//		w.SetAbort(true)
//
//		return
//	}
//	accessInfo, err := m.userInfoService.GetAccessInfo(ctx, userInfo.Id)
//	if err != nil {
//		w.JSON(http.StatusOK, controller.NewNoAccessError(err.Error()))
//		w.SetAbort(true)
//		return
//	}
//
//	//actionAccess := access.access(request.Module, selectAction(request.Method))
//
//	if accessInfo.IsAccess(access) {
//		return
//	}
//
//	w.JSON(http.StatusOK, controller.NewNoAccessError("权限不足"))
//	w.SetAbort(true)
//}

//func selectAction(method string) string {
//	switch method {
//	case http.MethodGet:
//		return access.AccessView
//	}
//	return access.AccessEdit
//}
