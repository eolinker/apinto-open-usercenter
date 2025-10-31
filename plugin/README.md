## 基本介绍
开源版用户中心插件（单用户）。提供登录、注销、个人资料查看与修改、修改密码等基础能力。不支持创建新账号或多用户权限管理。

## 默认账号
- 账号：admin
- 默认密码：12345678（首次登录后请立即修改）

## 功能特性
- 登录、注销、会话校验（JWT + Cookie）
- 个人资料查看与更新
- 修改密码（校验旧密码并设新密码）
- 返回已启用插件模块的访问信息

## 接口
- `POST /sso/login` 登录
- `POST /sso/logout` 注销
- `POST /sso/login/check` 会话校验
- `GET /api/my/profile` 获取个人资料
- `PUT /api/my/profile` 更新个人资料
- `POST /api/module/user/my/password` 修改密码
- `GET /api/my/access` 获取模块访问信息

## 前端
- 路由：`login`（根路由）、`user`（常规）
- 暴露模块：`LoginModule`、`AppModule`
- 入口名称：`开源版用户中心`

## 安装与使用
- 插件标识：`user-open.apinto.com`
- 部署方式：本地驱动（`driver: local`），与 Apinto Dashboard 同机部署。
- 启动后在控制台登录页面访问，使用默认账号登录并修改密码。

## 版本
- 当前版本：v3.0.0

## 注意事项
- 插件为单用户实现，不支持多用户、角色或精细化权限。
- 不提供与外部用户中心对接或监控告警联动能力。