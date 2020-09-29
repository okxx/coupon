package app

import (
	"fmt"
	"reflect"
)

type Group struct {
	app    *app
	prefix string
}

func (g *Group) Use(args ...interface{}) Routes {
	var prefix = ""
	var handles []HandlerFunc
	for i:=0;i < len(args);i++ {
		switch arg := args[i].(type) {
		case string:
			prefix = arg
		case HandlerFunc:
			handles = append(handles,arg)
		default:
			panic(fmt.Sprintf("invalid handleFunc %v\n",reflect.TypeOf(arg)))
		}
	}
	return g.app.register(methodUse,prefix,handles...)
}

func (g *Group) Get(relative string, handles ...HandlerFunc) Routes {
	return g.app.register(MethodGet,relative,handles...)
}

func (g *Group) Post(relative string, handles ...HandlerFunc) Routes {
	return g.app.register(MethodPost,relative,handles...)
}

func (g *Group) Delete(relative string, handles ...HandlerFunc) Routes {
	return g.app.register(MethodDelete,relative,handles...)
}

func (g *Group) Put(relative string, handles ...HandlerFunc) Routes {
	return g.app.register(MethodPut,relative,handles...)
}

func (g *Group) Patch(relative string, handles ...HandlerFunc) Routes {
	return g.app.register(MethodPatch,relative,handles...)
}

func (g *Group) Options(relative string, handles ...HandlerFunc) Routes {
	return g.app.register(MethodOptions,relative,handles...)
}

func (g *Group) Head(relative string, handles ...HandlerFunc) Routes {
	return g.app.register(MethodHead,relative,handles...)
}

func (g *Group) Static(relative string, handles ...HandlerFunc) Routes {
	return g.app.register(MethodGet,relative,handles...)
}

func (g *Group) Group(prefix string, handles ...HandlerFunc) Routes {
	if len(handles) > 0 {
		g.app.register(methodUse,prefix,handles...)
	}
	return g
}

