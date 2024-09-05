package controller

import (
	"net/http"
	"strings"

	"github.com/eolinker/apinto-dashboard/cache"
	grpcservice "github.com/eolinker/apinto-dashboard/grpc-service"
	"github.com/eolinker/apinto-dashboard/modules/mpm3"
	"github.com/eolinker/apinto-dashboard/plugin/go-plugin/plugin"
	"github.com/eolinker/apinto-dashboard/plugin/go-plugin/shared"
	"github.com/eolinker/apinto-dashboard/pm3"
	"github.com/eolinker/apinto-open-usercenter/frontend"
	"github.com/eolinker/apinto-open-usercenter/service"
	"github.com/eolinker/eosc/common/bean"
)

type Driver struct {
	apis        []pm3.Api
	frontend    []pm3.FrontendAsset
	middlewares []shared.Middleware
}

func NewPlugin() *Driver {
	m := newModule()

	return &Driver{
		apis:        m.apis(),
		middlewares: m.middlewares(),
		frontend:    frontend.FrontendAssets(),
	}
}

func (m *Driver) Frontend() []pm3.FrontendAsset {
	return m.frontend
}

func (m *Driver) Apis() []pm3.Api {
	return m.apis
}

func (m *Driver) Middleware() []shared.Middleware {
	return m.middlewares
}

type Module struct {
	userInfoService service.IUserInfoService
	sessionCache    service.ISessionCache
	commonCache     cache.ICommonCache
	consoleClient   grpcservice.GetConsoleInfoClient
	moduleService   mpm3.IPluginService
}

func newModule() *Module {
	l := &Module{}
	bean.Autowired(&l.userInfoService)
	bean.Autowired(&l.sessionCache)
	bean.Autowired(&l.commonCache)
	bean.Autowired(&l.consoleClient)
	return l
}

func (m *Module) apis() []pm3.Api {

	return []pm3.Api{
		{Method: http.MethodGet, Authority: pm3.Public, Path: "/api/my/access", HandlerFunc: m.access},
		{Method: http.MethodPut, Authority: pm3.Public, Path: "/api/my/profile", HandlerFunc: m.myProfileUpdate},
		{Method: http.MethodGet, Authority: pm3.Public, Path: "/api/my/profile", HandlerFunc: m.myProfile},
		{Method: http.MethodPost, Authority: pm3.Public, Path: "/api/module/user/my/password", HandlerFunc: m.setPassword},
		//{Method: http.MethodGet, Authority: pm3.Public, Path: "/api/user/enum", HandlerFunc: m.userEnum},
		{Method: http.MethodPost, Path: "/sso/login", HandlerFunc: m.ssoLogin, Authority: pm3.Anonymous},
		{Method: http.MethodPost, Path: "/sso/logout", HandlerFunc: m.ssoLogout, Authority: pm3.Anonymous},
		{Method: http.MethodPost, Path: "/sso/login/check", HandlerFunc: m.ssoLoginCheck, Authority: pm3.Anonymous},
	}

}

func (m *Module) middlewares() []shared.Middleware {
	mds := make([]shared.Middleware, 0, 2)
	mds = append(mds, plugin.NewMiddleware(func(info pm3.ApiInfo) bool {
		return false
	}, plugin.ProcessRequestBy(m.ModuleLogin)))

	mds = append(mds, plugin.NewMiddleware(func(info pm3.ApiInfo) bool {
		if info.Authority == pm3.Anonymous {
			return false
		}

		return strings.HasPrefix(info.Path, "/api/")
	}, plugin.ProcessRequestBy(m.ApiLogin)))

	return mds
}
