package app

import (
	"fmt"
	"net/http"
)

type app struct {
	rs [][]*Route

	server *http.Server
}

// new app
func New() *app {
	return &app{
		rs: make([][]*Route,len(methodConst)),
	}
}

func (app *app) Use(args ...interface{}) Routes {
	panic("implement me")
}

func (app *app) Get(relative string, handles ...HandlerFunc) Routes {
	return app.register(MethodGet,relative,handles...)
}

func (app *app) Post(relative string, handles ...HandlerFunc) Routes {
	return app.register(MethodPost,relative,handles...)
}

func (app *app) Delete(relative string, handles ...HandlerFunc) Routes {
	return app.register(MethodDelete,relative,handles...)
}

func (app *app) Put(relative string, handles ...HandlerFunc) Routes {
	return app.register(MethodPut,relative,handles...)
}

func (app *app) Patch(relative string, handles ...HandlerFunc) Routes {
	return app.register(MethodPatch,relative,handles...)
}

func (app *app) Options(relative string, handles ...HandlerFunc) Routes {
	return app.register(MethodOptions,relative,handles...)
}

func (app *app) Head(relative string, handles ...HandlerFunc) Routes {
	return app.register(MethodHead,relative,handles...)
}

func (app *app) Static(relative string, handles ...HandlerFunc) Routes {
	return app.register(MethodGet,relative,handles...)
}

func (app *app) Group(prefix string, handlers ...HandlerFunc) Routes {
	if len(handlers) > 0 {
		app.register(methodUse,prefix,handlers...)
	}
	return &Group{app: app,prefix: prefix}
}

func (app *app) ServeHTTP(http.ResponseWriter, *http.Request) {}

// route register
func (app *app) register(method, relative string, handles ...HandlerFunc) Routes {
	r := Route{
		Method: method,
		Path: relative,
		Handles: handles,
	}
	app.addRoute(method,&r)
	return app
}

func (app *app) addRoute(method string, handle *Route) {


}


func (app *app) Run(addr string) error {
	app.server.Addr = fmt.Sprintf(":%v",addr)
	return app.server.ListenAndServe()
}

