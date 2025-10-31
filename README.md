# Apinto 开源版用户中心插件（apinto-open-usercenter）

开源版用户中心是 Apinto 控制台的本地插件，提供基础的账号登录、注销、个人资料查看与修改、以及修改密码功能。该插件不支持创建新账户，系统默认仅有一个用户。

## 功能特性
- 登录与注销：支持账号登录、退出登录
- 个人资料：查看并修改昵称、邮箱、通知用户 ID 等
- 修改密码：校验旧密码后更新新密码
- 前端内嵌：Angular 前端资源打包并嵌入后端二进制，安装即用

## 使用限制
- 单用户模式：不支持用户注册/创建，仅保留一个用户
- 默认账号：`admin`
- 默认密码：`12345678`（首次登录后请尽快修改密码）

## 目录结构（简要）
- `app/user-center/` 插件入口与运行逻辑（gRPC、DB/Redis 初始化）
- `controller/` 登录、注销、个人资料、密码修改等路由与中间件
- `service/` 用户信息与会话缓存逻辑
- `store/` 用户信息持久化与默认管理员初始化
- `frontend/` Angular 前端工程，构建后嵌入二进制（`embed`）
- `plugin/` 插件元数据（`plugin.yml`、`logo.png`）
- `scripts/` 构建与打包脚本

## 构建与打包
环境要求：
- Go ≥ 1.18（推荐 1.19+）
- Node.js ≥ 16，建议安装 `yarn`
- 具备可用的 Apinto 控制台数据库与 Redis 配置

构建命令：
```
./scripts/build.sh [输出目录] [BUILD_MODE] [VERSION] [ARCH]
```
说明：
- 若不传 `输出目录`，脚本会创建并使用默认目录 `apinto-build`
- `BUILD_MODE=all` 时强制重新构建前端；不传则前端有 dist 时跳过
- `VERSION` 默认读取 `scripts/VERSION`
- `ARCH` 默认为 `amd64`（可选 `arm64` 等）

产物：
- 二进制：`user-open.apinto.com`（Linux）
- 插件资源包：`plugin-user.zip`（包含 `plugin.yml`、图标等）
- 打包输出：`user_${VERSION}_linux_amd64.tar.gz`（内含上述二者）

## 安装与启用
插件为本地驱动（`driver: local`），默认连接本机 `127.0.0.1:<Dashboard 端口>` 的 gRPC 服务，建议与 Apinto 控制台同机部署。

通用安装流程：
1. 将 `user_${VERSION}_linux_amd64.tar.gz` 拷贝到 Apinto 控制台所在主机并解压，得到目录 `user_${VERSION}/`
2. 确认包含：
   - 二进制：`user-open.apinto.com`（需赋予执行权限）
   - 插件资源包：`plugin-user.zip`
3. 安装方式 A（本地插件目录）：
   - 在控制台插件目录（例如 `plugins/user-open.apinto.com/`）放置二进制并确保可执行
   - 将 `plugin-user.zip` 解压出的 `plugin.yml`、`logo.png` 置于该目录
   - 重启/刷新控制台以生效
4. 安装方式 B（控制台 UI）：
   - 在控制台的插件管理页面，选择安装本地插件并上传 `plugin-user.zip`
   - 确保后端二进制 `user-open.apinto.com` 就绪，名称需与 `plugin.yml` 中 `define.cmd` 一致

安装完成后，控制台导航的“系统”分组下将显示“开源版用户中心”。

## 使用说明
- 登录：在登录页输入账号 `admin` 与密码 `12345678`
- 注销：点击右上角用户菜单，选择“退出登录”
- 修改密码：右上角用户菜单选择“修改密码”，输入旧密码与新密码提交
- 用户设置：右上角用户菜单选择“用户设置”，可编辑昵称、邮箱、通知用户 ID 等

## 相关接口（供参考）
- `POST /sso/login` 登录
- `POST /sso/logout` 注销
- `POST /sso/login/check` 登录状态检查
- `GET /api/my/profile` 查看个人资料
- `PUT /api/my/profile` 更新个人资料
- `POST /api/module/user/my/password` 修改密码

## 开发与调试
- 前端构建：在 `frontend/` 目录执行 `yarn install` 与 `yarn build`
- 本地调试：可使用 `debugPlugin` 构建标签启动本地 HTTP 服务（端口 `:9900`）
  - 例如：`go run -tags debugPlugin ./app/user-center`

## 插件标识
- 插件 ID：`user-open.apinto.com`
- 命令（binary 名）：`user-open.apinto.com`（来自 `plugin.yml` 的 `define.cmd`）