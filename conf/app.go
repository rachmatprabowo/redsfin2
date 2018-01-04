package conf

import (
	"github.com/rachmatprabowo/redsfin2/core"
	"github.com/rachmatprabowo/redsfin2/modules/auth/controller"
)

// InitDB initialization database config
func InitDB() {
	masterDB := &core.MasterDB
	dbCfg := &core.Databases

	*dbCfg = map[string]core.DB{
		"master": core.DB{
			DBName:   "redsfin2",
			Username: "postgres",
			Password: "postgres",
		},
	}

	*masterDB = core.Databases["master"].Connect()
}

// InitRoute intialization route config
func InitRoute() {
	routeCfg := &core.Routes

	*routeCfg = []core.Route{
		core.Route{
			Path: "/user",
			Fn:   controller.UserHandler,
		},
		core.Route{
			Path: "/user/{id:[0-9]+}",
			Fn:   controller.UserHandler,
		},
	}
}

// RegisterDB regtering database for client's company
func RegisterDB(db core.DB, name string) {
	core.Databases[name] = db
}
