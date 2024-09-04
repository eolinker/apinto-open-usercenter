package controller

import (
	"fmt"
	"net/http"
	"time"
	
	user_model "github.com/eolinker/apinto-open-usercenter/model"
	
	"github.com/go-basic/uuid"
	
	"github.com/eolinker/apinto-dashboard/common"
	"github.com/eolinker/apinto-dashboard/controller"
	user_dto "github.com/eolinker/apinto-open-usercenter/dto"
	"github.com/gin-gonic/gin"
)

func (m *Module) myProfile(ginCtx *gin.Context) {
	userId := m.getUserId(ginCtx)

	userInfo, err := m.userInfoService.GetUserInfo(ginCtx, userId)
	if err != nil {
		controller.ErrorJson(ginCtx, http.StatusOK, fmt.Sprintf("getMyProfile fail. err:%s", err.Error()))

		return
	}

	lastLogin := ""
	if userInfo.LastLoginTime != nil {
		lastLogin = common.TimeToStr(*userInfo.LastLoginTime)
	}
	resUserInfo := user_dto.UserInfo{
		Id:     userInfo.Id,
		Sex:    userInfo.Sex,
		Avatar: userInfo.Avatar,

		Email: userInfo.Email,
		Phone: userInfo.Phone,

		UserName:     userInfo.UserName,
		NickName:     userInfo.NickName,
		NoticeUserId: userInfo.NoticeUserId,
		LastLogin:    lastLogin,
	}

	result := make(map[string]interface{})
	result["profile"] = resUserInfo
	result["describe"] = ""

	ginCtx.JSON(http.StatusOK, controller.NewSuccessResult(result))
}
func (m *Module) myProfileUpdate(ginCtx *gin.Context) {
	userId := m.getUserId(ginCtx)
	req := &user_dto.UpdateMyProfileReq{}
	err := ginCtx.BindJSON(req)
	if err != nil {
		controller.ErrorJson(ginCtx, http.StatusOK, fmt.Sprintf("updateMyProfile fail. err:%s", err.Error()))
		return
	}

	if err = m.userInfoService.UpdateMyProfile(ginCtx, userId, req); err != nil {
		controller.ErrorJson(ginCtx, http.StatusOK, fmt.Sprintf("updateMyProfile fail. err:%s", err.Error()))
		return
	}

	ginCtx.JSON(http.StatusOK, controller.NewSuccessResult(nil))
}
func (m *Module) setPassword(ginCtx *gin.Context) {
	userId := m.getUserId(ginCtx)
	req := &user_dto.UpdateMyPasswordReq{}
	err := ginCtx.BindJSON(req)
	if err != nil {
		controller.ErrorJson(ginCtx, http.StatusOK, fmt.Sprintf("updateMyPassword fail. err:%s", err.Error()))
		return
	}

	if err = m.userInfoService.UpdateMyPassword(ginCtx, userId, req); err != nil {
		controller.ErrorJson(ginCtx, http.StatusOK, fmt.Sprintf("updateMyPassword fail. err:%s", err.Error()))
		return

	}
	info, err := m.userInfoService.GetUserInfo(ginCtx, userId)
	if err != nil {
		controller.ErrorJson(ginCtx, http.StatusOK, fmt.Sprintf("get user info fail. err:%s", err.Error()))
		return
	}
	userJWT, err := common.JWTEncode(&controller.UserClaim{
		Id:        info.Id,
		Uname:     info.UserName,
		LoginTime: info.LastLoginTime.Format("2006-01-02 15:04:05"),
	}, jwtSecret)
	if err != nil {
		controller.ErrorJson(ginCtx, http.StatusOK, "登录失效")
		return
	}
	cookieValue := uuid.New()
	session := &user_model.Session{
		Jwt:  userJWT,
		RJwt: common.Md5(userJWT),
	}

	if err = m.sessionCache.Set(ginCtx, cookieValue, session); err != nil {
		controller.ErrorJson(ginCtx, http.StatusOK, err.Error())
		return
	}

	//每次登陆都把之前的token清除掉
	{
		userIdCookieKey := fmt.Sprintf("userId:%d", info.Id)
		oldCookie, _ := m.commonCache.Get(ginCtx, userIdCookieKey)
		_ = m.sessionCache.Delete(ginCtx, string(oldCookie))
		_ = m.commonCache.Set(ginCtx, userIdCookieKey, []byte(cookieValue), time.Hour*24*7)
	}

	ginCtx.SetCookie(controller.Session, cookieValue, 0, "", "", false, true)

	ginCtx.JSON(http.StatusOK, controller.NewSuccessResult(nil))

}
func (m *Module) userEnum(ginCtx *gin.Context) {
	userInfoList, err := m.userInfoService.GetAllUsers(ginCtx)
	if err != nil {
		controller.ErrorJson(ginCtx, http.StatusOK, fmt.Sprintf("获取用户列表失败. err:%s", err.Error()))
		return
	}

	resList := make([]user_dto.UserInfo, 0, len(userInfoList))

	for _, userInfo := range userInfoList {
		lastLogin := ""
		if userInfo.LastLoginTime != nil {
			lastLogin = common.TimeToStr(*userInfo.LastLoginTime)
		}
		resUserInfo := user_dto.UserInfo{
			Id:           userInfo.Id,
			Sex:          userInfo.Sex,
			Avatar:       userInfo.Avatar,
			Email:        userInfo.Email,
			Phone:        userInfo.Phone,
			UserName:     userInfo.UserName,
			NickName:     userInfo.NickName,
			NoticeUserId: userInfo.NoticeUserId,
			LastLogin:    lastLogin,
		}
		resList = append(resList, resUserInfo)
	}

	result := make(map[string]interface{})
	result["users"] = resList

	ginCtx.JSON(http.StatusOK, controller.NewSuccessResult(result))
}
func (m *Module) access(ginCtx *gin.Context) {
	modules, err := m.moduleService.GetEnabled(ginCtx)
	if err != nil {
		controller.ErrorJson(ginCtx, http.StatusOK, fmt.Sprintf("get module info fail. err:%s", err.Error()))
		return
	}
	access := make([]*user_dto.UserAccess, 0, len(modules))
	for _, mm := range modules {
		access = append(access, &user_dto.UserAccess{
			Name:   mm.UUID,
			Access: "edit",
		})
	}
	ginCtx.JSON(http.StatusOK, controller.NewSuccessResult(map[string]interface{}{
		"access": access,
	}))
}

func (m *Module) ssoLogin(ginCtx *gin.Context) {
	var loginInfo user_dto.UserLogin
	err := ginCtx.BindJSON(&loginInfo)
	if err != nil {
		controller.ErrorJson(ginCtx, http.StatusOK, fmt.Sprintf("bind login param fail. err:%s", err.Error()))
		return
	}

	id, success := m.userInfoService.CheckPassword(ginCtx, loginInfo.Username, loginInfo.Password)
	if !success {
		ginCtx.JSON(http.StatusOK, controller.NewLoginInvalidError(controller.CodeLoginPwdErr, "登录失败，用户名或密码错误"))
		return
	}

	now := time.Now()
	// 成功登录，更新登录时间
	err = m.userInfoService.UpdateLastLoginTime(ginCtx, id, &now)
	if err != nil {
		controller.ErrorJson(ginCtx, http.StatusOK, "登录失败")
		return
	}

	userJWT, err := common.JWTEncode(&controller.UserClaim{
		Id:        id,
		Uname:     loginInfo.Username,
		LoginTime: now.Format("2006-01-02 15:04:05"),
	}, jwtSecret)
	if err != nil {
		controller.ErrorJson(ginCtx, http.StatusOK, "登录失败")
		return
	}
	cookieValue := uuid.New()
	session := &user_model.Session{
		Jwt:  userJWT,
		RJwt: common.Md5(userJWT),
	}

	if err = m.sessionCache.Set(ginCtx, cookieValue, session); err != nil {
		controller.ErrorJson(ginCtx, http.StatusOK, err.Error())
		return
	}

	//每次登陆都把之前的token清除掉
	{
		userIdCookieKey := fmt.Sprintf("userId:%d", id)
		oldCookie, _ := m.commonCache.Get(ginCtx, userIdCookieKey)
		_ = m.sessionCache.Delete(ginCtx, string(oldCookie))
		_ = m.commonCache.Set(ginCtx, userIdCookieKey, []byte(cookieValue), time.Hour*24*7)
	}

	ginCtx.SetCookie(controller.Session, cookieValue, 0, "", "", false, true)
	ginCtx.JSON(http.StatusOK, controller.NewSuccessResult(nil))
}
func (m *Module) ssoLogout(ginCtx *gin.Context) {
	cookie, err := ginCtx.Cookie(controller.Session)
	if err != nil {
		controller.ErrorJson(ginCtx, http.StatusOK, err.Error())
		return
	}

	m.sessionCache.Delete(ginCtx, cookie)
	defer func() {
		ginCtx.SetCookie(controller.Session, cookie, -1, "", "", false, true)
	}()

	ginCtx.JSON(http.StatusOK, controller.NewSuccessResult(nil))
}
func (m *Module) ssoLoginCheck(ginCtx *gin.Context) {

	cookie, err := ginCtx.Cookie(controller.Session)
	if err != nil {
		controller.ErrorJsonWithCode(ginCtx, http.StatusOK, controller.CodeLoginInvalid, loginError)
		return
	}

	session, _ := m.sessionCache.Get(ginCtx, cookie)
	if session == nil {
		controller.ErrorJsonWithCode(ginCtx, http.StatusOK, controller.CodeLoginInvalid, loginError)
		return
	}
	uc, err := common.JWTDecode(session.Jwt, jwtSecret)
	if err != nil {
		controller.ErrorJsonWithCode(ginCtx, http.StatusOK, controller.CodeLoginInvalid, loginError)
		return
	}

	info, err := m.userInfoService.GetUserInfo(ginCtx, uc.Id)
	if err != nil {
		controller.ErrorJsonWithCode(ginCtx, http.StatusOK, controller.CodeLoginInvalid, loginError)
		return
	}
	if info.LastLoginTime.Format("2006-01-02 15:04:05") != uc.LoginTime || info.UserName != uc.Uname {
		controller.ErrorJsonWithCode(ginCtx, http.StatusOK, controller.CodeLoginInvalid, loginError)
		return
	}

	ginCtx.JSON(http.StatusOK, controller.NewSuccessResult(true))
}

//func randomRouters() apinto_module.RoutersInfo {
//	r := random_controller.NewRandomController()
//	return apinto_module.RoutersInfo{
//		{
//			Method:      http.MethodGet,
//			Path:        "/api/random/:template/id",
//			Handler:     "core.random.id",
//			Labels:      apinto_module.RouterLabelApi,
//			HandlerFunc: []apinto_module.HandlerFunc{r.GET},
//			Replaceable: false,
//		}}
//}
