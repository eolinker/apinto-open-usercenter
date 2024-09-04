package controller

//
//func (m *Module) getMyAccess(ginCtx *gin.Context) {
//	audit_model.LogOperateTypeNone.Handler(ginCtx)
//
//	userId := m.getUserId(ginCtx)
//	if userId == 0 {
//		ginCtx.JSON(http.StatusOK, loginError)
//		ginCtx.Abort()
//		return
//	}
//	accessForUser, err := m.userInfoService.GetAccessInfo(ginCtx, userId)
//	if err != nil {
//		ginCtx.JSON(http.StatusOK, loginError)
//		ginCtx.Abort()
//		return
//	}
//
//	uas := make(map[string]*dto.UserAccess)
//	for k := range accessForUser {
//		module, acc, err := access.ReadAccess(k)
//		if err != nil {
//			continue
//		}
//		if ov, has := uas[module]; has {
//			if ov.Access != acc && acc == access.AccessEdit {
//				ov.Access = acc
//			}
//
//		} else {
//			uas[module] = &dto.UserAccess{Name: module, Access: acc, Module: module}
//		}
//	}
//	ual := make([]*dto.UserAccess, 0, len(uas))
//	for _, v := range uas {
//		ual = append(ual, v)
//	}
//	ginCtx.JSON(http.StatusOK, controller.NewSuccessResult(map[string]interface{}{
//		"access": ual,
//	}))
//}
//func (m *Module) getAllAccess(ginCtx *gin.Context) {
//	modules, err := m.consoleClient.GetNavigationModules(ginCtx, &grpcservice.EmptyRequest{})
//	if err != nil {
//		ginCtx.JSON(http.StatusOK, controller.NewErrorResult(fmt.Sprintf("getAllAccess fail. err:%s", err.Error())))
//		return
//	}
//
//	data := make(map[string]interface{})
//	data["modules"] = getModules(modules)
//	data["depth"] = 2
//
//	ginCtx.JSON(http.StatusOK, controller.NewSuccessResult(data))
//
//}
//
//func getModules(modules *grpcservice.NavigationModulesResp) []*dto.SystemModuleItem {
//	all := make(map[string]*dto.SystemModuleItem)
//
//	for _, n := range modules.NavigationItems {
//		it := &dto.SystemModuleItem{
//
//			Title:    n.Cname,
//			Module:   n.Id,
//			Access:   nil,
//			Children: nil,
//		}
//
//		all[it.Module] = it
//	}
//	for _, n := range modules.ModulesItems {
//		it := &dto.SystemModuleItem{
//
//			Title:    n.Cname,
//			Module:   n.Name,
//			Access:   make([]*dto.AccessItem, 0, len(n.Access)),
//			Children: nil,
//		}
//		for _, a := range n.Access {
//			it.Access = append(it.Access, &dto.AccessItem{
//				Key:          a.Name,
//				Title:        a.Cname,
//				Dependencies: a.Depend,
//			})
//		}
//		all[it.Module] = it
//		if parent, has := all[n.NavigationId]; has {
//			parent.Children = append(parent.Children, it)
//		}
//	}
//	root := make([]*dto.SystemModuleItem, 0, len(modules.NavigationItems))
//	for _, n := range modules.NavigationItems {
//		if nv, has := all[n.Id]; has {
//			if len(nv.Children) > 0 {
//				root = append(root, nv)
//			}
//		}
//	}
//	return root
//}
//func accessItemForModule(name string) []*dto.AccessItem {
//	return []*dto.AccessItem{{
//		Key:   access.CreateAccess(name, access.AccessView)[0],
//		Title: access.AccessTitleView,
//	}, {
//		Key:          access.CreateAccess(name, access.AccessEdit)[0],
//		Title:        access.AccessTitleEdit,
//		Dependencies: access.CreateAccess(name, access.AccessView),
//	}}
//}
//func (m *Module) getMyProfile(ginCtx *gin.Context) {
//	audit_model.LogOperateTypeNone.Handler(ginCtx)
//
//	userId := m.getUserId(ginCtx)
//
//	if userId == 0 {
//		ginCtx.JSON(http.StatusOK, loginError)
//		ginCtx.Abort()
//		return
//	}
//	userInfo, err := m.userInfoService.GetUserInfo(ginCtx, userId)
//	if err != nil {
//		ginCtx.JSON(http.StatusOK, controller.NewLoginInvalidError(controller.CodeLoginUserNoExistent, fmt.Sprintf("getMyProfile fail. err:%s", err.Error())))
//		return
//	}
//
//	lastLogin := ""
//	if userInfo.LastLoginTime != nil {
//		lastLogin = common.TimeToStr(*userInfo.LastLoginTime)
//	}
//	resUserInfo := dto.UserInfo{
//		Id:           userInfo.Id,
//		Sex:          userInfo.Sex,
//		Avatar:       userInfo.Avatar,
//		Desc:         userInfo.Remark,
//		Email:        userInfo.Email,
//		Phone:        userInfo.Phone,
//		Status:       userInfo.Status,
//		UserName:     userInfo.UserName,
//		NickName:     userInfo.NickName,
//		NoticeUserId: userInfo.NoticeUserId,
//		LastLogin:    lastLogin,
//		CreateTime:   common.TimeToStr(userInfo.CreateTime),
//		UpdateTime:   common.TimeToStr(userInfo.UpdateTime),
//		RoleIds:      strings.Split(userInfo.RoleIds, ","),
//	}
//
//	data := make(map[string]interface{})
//	data["profile"] = resUserInfo
//	data["describe"] = userInfo.Remark
//
//	ginCtx.JSON(http.StatusOK, controller.NewSuccessResult(data))
//}
//
//func (m *Module) updateMyProfile(ginCtx *gin.Context) {
//	audit_model.LogOperateTypeNone.Handler(ginCtx)
//
//	userId := m.getUserId(ginCtx)
//
//	req := &dto.UpdateMyProfileReq{}
//	err := ginCtx.BindJSON(req)
//	if err != nil {
//		ginCtx.JSON(http.StatusOK, controller.NewErrorResult(fmt.Sprintf("updateMyProfile fail. err:%s", err.Error())))
//		return
//	}
//
//	if err = m.userInfoService.UpdateMyProfile(ginCtx, userId, req); err != nil {
//		ginCtx.JSON(http.StatusOK, controller.NewErrorResult(fmt.Sprintf("updateMyProfile fail. err:%s", err.Error())))
//		return
//	}
//
//	ginCtx.JSON(http.StatusOK, controller.NewSuccessResult(nil))
//}
//
//func (m *Module) updateMyPassword(ginCtx *gin.Context) {
//	audit_model.LogOperateTypeNone.Handler(ginCtx)
//
//	userId := m.getUserId(ginCtx)
//
//	req := &dto.UpdateMyPasswordReq{}
//	err := ginCtx.BindJSON(req)
//	if err != nil {
//		ginCtx.JSON(http.StatusOK, controller.NewErrorResult(fmt.Sprintf("updateMyPassword fail. err:%s", err.Error())))
//		return
//	}
//
//	if err = m.userInfoService.UpdateMyPassword(ginCtx, userId, req); err != nil {
//		ginCtx.JSON(http.StatusOK, controller.NewErrorResult(fmt.Sprintf("updateMyPassword fail. err:%s", err.Error())))
//		return
//	}
//
//	ginCtx.JSON(http.StatusOK, controller.NewSuccessResult(nil))
//}
//
//func (m *Module) getRoleList(ginCtx *gin.Context) {
//
//	roleList, totalUsers, err := m.userInfoService.GetRoleList(ginCtx)
//	if err != nil {
//		ginCtx.JSON(http.StatusOK, controller.NewErrorResult(fmt.Sprintf("GetRoleList fail. err:%s", err.Error())))
//		return
//	}
//	roles := make([]*dto.RoleListItem, 0, len(roleList))
//	for _, item := range roleList {
//		role := &dto.RoleListItem{
//			ID:             item.ID,
//			Title:          item.Title,
//			UserNum:        item.UserNum,
//			OperateDisable: item.OperateDisable,
//			Type:           item.Type,
//		}
//		roles = append(roles, role)
//	}
//
//	data := make(map[string]interface{})
//	data["roles"] = roles
//	data["total"] = totalUsers
//	ginCtx.JSON(http.StatusOK, controller.NewSuccessResult(data))
//}
//
//func (m *Module) getRoleInfo(ginCtx *gin.Context) {
//	roleID := ginCtx.Query("id")
//	if roleID == "" {
//		ginCtx.JSON(http.StatusOK, controller.NewErrorResult(fmt.Sprintf("GetRoleInfo fail. err: id can't be nil")))
//		return
//	}
//
//	info, err := m.userInfoService.GetRoleInfo(ginCtx, roleID)
//	if err != nil {
//		ginCtx.JSON(http.StatusOK, controller.NewErrorResult(fmt.Sprintf("GetRoleInfo fail. err:%s", err.Error())))
//		return
//	}
//	role := &dto.ProxyRoleInfo{
//		Title:  info.Title,
//		Desc:   info.Desc,
//		Access: info.Access,
//	}
//	data := make(map[string]interface{})
//	data["role"] = role
//	ginCtx.JSON(http.StatusOK, controller.NewSuccessResult(data))
//}
//
//func (m *Module) getRoleOptions(ginCtx *gin.Context) {
//	optionList, err := m.userInfoService.GetRoleOptions(ginCtx)
//	if err != nil {
//		ginCtx.JSON(http.StatusOK, controller.NewErrorResult(fmt.Sprintf("GetRoleOptions fail. err:%s", err.Error())))
//		return
//	}
//	options := make([]*dto.RoleOptionItem, 0, len(optionList))
//	for _, item := range optionList {
//		option := &dto.RoleOptionItem{
//			ID:             item.ID,
//			Title:          item.Title,
//			OperateDisable: item.OperateDisable,
//		}
//		options = append(options, option)
//	}
//
//	data := make(map[string]interface{})
//	data["roles"] = options
//	ginCtx.JSON(http.StatusOK, controller.NewSuccessResult(data))
//}
//
//func (m *Module) createRole(ginCtx *gin.Context) {
//	userID := m.getUserId(ginCtx)
//
//	input := new(dto.ProxyRoleInfo)
//	if err := ginCtx.BindJSON(input); err != nil {
//		ginCtx.JSON(http.StatusOK, controller.NewErrorResult(err.Error()))
//		return
//	}
//
//	//Check Input
//	if input.Title == "" {
//		ginCtx.JSON(http.StatusOK, controller.NewErrorResult(fmt.Sprintf("CreateRole fail. err: title can't be nil. ")))
//		return
//	}
//
//	err := m.userInfoService.CreateRole(ginCtx, userID, input)
//	if err != nil {
//		ginCtx.JSON(http.StatusOK, controller.NewErrorResult(fmt.Sprintf("CreateAPI fail. err:%s", err.Error())))
//		return
//	}
//
//	ginCtx.JSON(http.StatusOK, controller.NewSuccessResult(nil))
//}
//
//func (m *Module) updateRole(ginCtx *gin.Context) {
//	userID := m.getUserId(ginCtx)
//
//	roleID := ginCtx.Query("id")
//	if roleID == "" {
//		ginCtx.JSON(http.StatusOK, controller.NewErrorResult(fmt.Sprintf("UpdateRole fail. err: id can't be nil")))
//		return
//	}
//
//	input := new(dto.ProxyRoleInfo)
//	if err := ginCtx.BindJSON(input); err != nil {
//		ginCtx.JSON(http.StatusOK, controller.NewErrorResult(err.Error()))
//		return
//	}
//
//	//Check Input
//	if input.Title == "" {
//		ginCtx.JSON(http.StatusOK, controller.NewErrorResult(fmt.Sprintf("UpdateRole fail. err: title can't be nil. ")))
//		return
//	}
//
//	err := m.userInfoService.UpdateRole(ginCtx, userID, roleID, input)
//	if err != nil {
//		ginCtx.JSON(http.StatusOK, controller.NewErrorResult(fmt.Sprintf("UpdateRole fail. err:%s", err.Error())))
//		return
//	}
//
//	ginCtx.JSON(http.StatusOK, controller.NewSuccessResult(nil))
//}
//
//func (m *Module) deleteRole(ginCtx *gin.Context) {
//	userID := m.getUserId(ginCtx)
//	roleID := ginCtx.Query("id")
//	if roleID == "" {
//		ginCtx.JSON(http.StatusOK, controller.NewErrorResult(fmt.Sprintf("DeleteRole fail. err: id can't be nil")))
//		return
//	}
//
//	err := m.userInfoService.DeleteRole(ginCtx, userID, roleID)
//	if err != nil {
//		ginCtx.JSON(http.StatusOK, controller.NewErrorResult(fmt.Sprintf("DeleteRole fail. err:%s", err.Error())))
//		return
//	}
//
//	ginCtx.JSON(http.StatusOK, controller.NewSuccessResult(nil))
//}
//
//func (m *Module) roleBatchUpdate(ginCtx *gin.Context) {
//	input := new(dto.BatchUpdateRole)
//	if err := ginCtx.BindJSON(input); err != nil {
//		ginCtx.JSON(http.StatusOK, controller.NewErrorResult(err.Error()))
//		return
//	}
//
//	err := m.userInfoService.RoleBatchUpdate(ginCtx, input.Ids, input.RoleId)
//	if err != nil {
//		ginCtx.JSON(http.StatusOK, controller.NewErrorResult(fmt.Sprintf("roleBatchUpdate fail. err:%s", err.Error())))
//		return
//	}
//
//	ginCtx.JSON(http.StatusOK, controller.NewSuccessResult(nil))
//}
//
//func (m *Module) roleBatchRemove(ginCtx *gin.Context) {
//	input := new(dto.BatchRemoveRole)
//	if err := ginCtx.BindJSON(input); err != nil {
//		ginCtx.JSON(http.StatusOK, controller.NewErrorResult(err.Error()))
//		return
//	}
//
//	err := m.userInfoService.RoleBatchRemove(ginCtx, input.Ids, input.RoleId)
//	if err != nil {
//		ginCtx.JSON(http.StatusOK, controller.NewErrorResult(fmt.Sprintf("roleBatchRemove fail. err:%s", err.Error())))
//		return
//	}
//
//	ginCtx.JSON(http.StatusOK, controller.NewSuccessResult(nil))
//}
//
//func (m *Module) delUser(ginCtx *gin.Context) {
//	userID := m.getUserId(ginCtx)
//
//	req := &dto.DelUserReq{}
//	if err := ginCtx.BindJSON(req); err != nil {
//		ginCtx.JSON(http.StatusOK, controller.NewErrorResult(fmt.Sprintf("delUser fail. err:%s", err.Error())))
//		return
//	}
//
//	if len(req.UserIds) == 0 {
//		ginCtx.JSON(http.StatusOK, controller.NewErrorResult(fmt.Sprintf("delUser fail. err:%s", "参数错误")))
//		return
//	}
//
//	err := m.userInfoService.DelUser(ginCtx, userID, req.UserIds)
//	if err != nil {
//		ginCtx.JSON(http.StatusOK, controller.NewErrorResult(fmt.Sprintf("delUser fail. err:%s", err.Error())))
//		return
//	}
//	//info, err := m.userInfoService.GetUserInfo(ginCtx, userID)
//	//if err == nil {
//	//	apintomodule.AddEvent(ginCtx, "user-delete", info.UserName)
//	//}
//	ginCtx.JSON(http.StatusOK, controller.NewSuccessResult(nil))
//
//}
//
//func (m *Module) createUser(ginCtx *gin.Context) {
//	userID := m.getUserId(ginCtx)
//
//	req := &dto.SaveUserReq{}
//	if err := ginCtx.BindJSON(req); err != nil {
//		ginCtx.JSON(http.StatusOK, controller.NewErrorResult(fmt.Sprintf("createUser fail. err:%s", err.Error())))
//		return
//	}
//
//	if err := common.IsMatchString(common.EnglishOrNumber_, req.UserName); err != nil {
//		return
//	}
//
//	err := m.userInfoService.CreateUser(ginCtx, userID, req)
//	if err != nil {
//		ginCtx.JSON(http.StatusOK, controller.NewErrorResult(fmt.Sprintf("createUser fail. err:%s", err.Error())))
//		return
//	}
//
//	ginCtx.JSON(http.StatusOK, controller.NewSuccessResult(nil))
//
//}
//
//func (m *Module) patchUser(ginCtx *gin.Context) {
//	operator := m.getUserId(ginCtx)
//	userIDStr := ginCtx.Query("id")
//
//	userId, err := strconv.Atoi(userIDStr)
//	if err != nil {
//		ginCtx.JSON(http.StatusOK, controller.NewErrorResult(fmt.Sprintf("patchUser fail. err:%s", err.Error())))
//		return
//	}
//
//	req := &dto.PatchUserReq{}
//	if err = ginCtx.BindJSON(req); err != nil {
//		ginCtx.JSON(http.StatusOK, controller.NewErrorResult(fmt.Sprintf("patchUser fail. err:%s", err.Error())))
//		return
//	}
//
//	if err = m.userInfoService.PatchUser(ginCtx, operator, userId, req); err != nil {
//		ginCtx.JSON(http.StatusOK, controller.NewErrorResult(fmt.Sprintf("patchUser fail. err:%s", err.Error())))
//		return
//	}
//
//	ginCtx.JSON(http.StatusOK, controller.NewSuccessResult(nil))
//
//}
//
//func (m *Module) updateUser(ginCtx *gin.Context) {
//	operator := m.getUserId(ginCtx)
//
//	userIdStr := ginCtx.Query("id")
//	userId, err := strconv.Atoi(userIdStr)
//	if err != nil {
//		ginCtx.JSON(http.StatusOK, controller.NewErrorResult(fmt.Sprintf("updateUser fail. err:%s", err.Error())))
//		return
//	}
//
//	req := &dto.SaveUserReq{}
//	if err := ginCtx.BindJSON(req); err != nil {
//		ginCtx.JSON(http.StatusOK, controller.NewErrorResult(fmt.Sprintf("updateUser fail. err:%s", err.Error())))
//		return
//	}
//
//	if err = m.userInfoService.UpdateUser(ginCtx, operator, userId, req); err != nil {
//		ginCtx.JSON(http.StatusOK, controller.NewErrorResult(fmt.Sprintf("updateUser fail. err:%s", err.Error())))
//		return
//	}
//	//info, err := m.userInfoService.GetUserInfo(ginCtx, userId)
//	//if err == nil {
//	//	apintomodule.AddEvent(ginCtx, "user-update", info.Base())
//	//}
//	ginCtx.JSON(http.StatusOK, controller.NewSuccessResult(nil))
//
//}
//
//func (m *Module) getUserList(ginCtx *gin.Context) {
//
//	roleId := ginCtx.Query("role")
//	keyword := ginCtx.Query("keyword")
//
//	userInfoList, err := m.userInfoService.GetUserInfoList(ginCtx, roleId, keyword)
//	if err != nil {
//		ginCtx.JSON(http.StatusOK, controller.NewErrorResult(fmt.Sprintf("getUserList fail. err:%s", err.Error())))
//		return
//	}
//	userIds := common.SliceToSliceIds(userInfoList, func(t *model.UserInfo) int {
//		return t.Operator
//	})
//
//	userInfoMaps, _ := m.userInfoService.GetUserInfoMaps(ginCtx, userIds...)
//	resList := make([]dto.UserInfo, 0, len(userInfoList))
//
//	for _, userInfo := range userInfoList {
//		lastLogin := ""
//		if userInfo.LastLoginTime != nil {
//			lastLogin = common.TimeToStr(*userInfo.LastLoginTime)
//		}
//		operatorName := ""
//		if o, ok := userInfoMaps[userInfo.Operator]; ok {
//			operatorName = o.NickName
//		}
//		resUserInfo := dto.UserInfo{
//			Id:             userInfo.Id,
//			Sex:            userInfo.Sex,
//			Avatar:         userInfo.Avatar,
//			Email:          userInfo.Email,
//			Phone:          userInfo.Phone,
//			Status:         userInfo.Status,
//			UserName:       userInfo.UserName,
//			NickName:       userInfo.NickName,
//			NoticeUserId:   userInfo.NoticeUserId,
//			LastLogin:      lastLogin,
//			CreateTime:     common.TimeToStr(userInfo.CreateTime),
//			UpdateTime:     common.TimeToStr(userInfo.UpdateTime),
//			OperateDisable: userInfo.UserName == service.AdminName,
//			Desc:           userInfo.Remark,
//			Operator:       operatorName,
//			RoleIds:        strings.Split(userInfo.RoleIds, ","),
//		}
//		resList = append(resList, resUserInfo)
//	}
//
//	data := make(map[string]interface{})
//	data["users"] = resList
//
//	ginCtx.JSON(http.StatusOK, controller.NewSuccessResult(data))
//
//}
//
//func (m *Module) getUser(ginCtx *gin.Context) {
//
//	userIdStr := ginCtx.Query("id")
//	userId, err := strconv.Atoi(userIdStr)
//	if err != nil {
//		ginCtx.JSON(http.StatusOK, controller.NewErrorResult(fmt.Sprintf("getUser fail. err:%s", err.Error())))
//		return
//	}
//
//	userInfo, err := m.userInfoService.GetUserInfo(ginCtx, userId)
//	if err != nil {
//		ginCtx.JSON(http.StatusOK, controller.NewErrorResult(fmt.Sprintf("getUser fail. err:%s", err.Error())))
//		return
//	}
//
//	lastLogin := ""
//	if userInfo.LastLoginTime != nil {
//		lastLogin = common.TimeToStr(*userInfo.LastLoginTime)
//	}
//	resUserInfo := dto.UserInfo{
//		Id:           userInfo.Id,
//		Sex:          userInfo.Sex,
//		Avatar:       userInfo.Avatar,
//		Desc:         userInfo.Remark,
//		Email:        userInfo.Email,
//		NoticeUserId: userInfo.NoticeUserId,
//		Phone:        userInfo.Phone,
//		Status:       userInfo.Status,
//		UserName:     userInfo.UserName,
//		NickName:     userInfo.NickName,
//		LastLogin:    lastLogin,
//		CreateTime:   common.TimeToStr(userInfo.CreateTime),
//		UpdateTime:   common.TimeToStr(userInfo.UpdateTime),
//		RoleIds:      strings.Split(userInfo.RoleIds, ","),
//	}
//
//	data := make(map[string]interface{})
//	data["profile"] = resUserInfo
//	data["describe"] = userInfo.Remark
//
//	ginCtx.JSON(http.StatusOK, controller.NewSuccessResult(data))
//
//}
//
//func (m *Module) resetUserPwd(ginCtx *gin.Context) {
//
//	operator := m.getUserId(ginCtx)
//
//	resetPasswordReq := new(dto.ResetPasswordReq)
//	err := ginCtx.BindJSON(resetPasswordReq)
//	if err != nil {
//		ginCtx.JSON(http.StatusOK, controller.NewErrorResult(fmt.Sprintf("resetUserPwd fail. err:%s", err.Error())))
//		return
//	}
//
//	err = m.userInfoService.ResetUserPwd(ginCtx, operator, resetPasswordReq.Id, resetPasswordReq.Password)
//	if err != nil {
//		ginCtx.JSON(http.StatusOK, controller.NewErrorResult(fmt.Sprintf("resetUserPwd fail. err:%s", err.Error())))
//		return
//	}
//
//	ginCtx.JSON(http.StatusOK, controller.NewSuccessResult(nil))
//
//}
