package router

import (
	"github.com/darianJmy/fortis-services/api/server/router/cmdb"
	"github.com/darianJmy/fortis-services/api/server/router/info"
	"github.com/darianJmy/fortis-services/cmd/app/options"
)

type registerFunc func(o *options.ServerRunOptions)

func InstallRouters(o *options.ServerRunOptions) {

	fs := []registerFunc{
		//middleware.NewMiddlewares,

		info.NewRouter,
		cmdb.NewRouter,
	}

	install(o, fs...)
}

func install(o *options.ServerRunOptions, fs ...registerFunc) {
	for _, f := range fs {
		f(o)
	}
}
