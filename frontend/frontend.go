package frontend

import (
	"embed"
	_ "embed"
	apinto_module "github.com/eolinker/apinto-dashboard/module"
	"github.com/eolinker/apinto-dashboard/pm3"
	"io/fs"
	"net/http"
)

//go:embed dist
var dist embed.FS

func getFileSystem(dir string) http.FileSystem {
	sub, err := fs.Sub(dist, dir)
	if err != nil {
		panic(err)
	}

	return http.FS(sub)
}

func FrontendAssets() []pm3.FrontendAsset {
	return []pm3.FrontendAsset{

		apinto_module.StaticFileDir("/plugin-frontend/user/", getFileSystem("dist")),
	}
}
