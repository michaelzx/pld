package pld_gin

import (
	"github.com/gin-gonic/gin"
)

type HttpMethod string

const (
	HttpGet     HttpMethod = "GET"
	HttpHead    HttpMethod = "HEAD"
	HttpPost    HttpMethod = "POST"
	HttpPut     HttpMethod = "PUT"
	HttpPatch   HttpMethod = "PATCH" // RFC 5789
	HttpDelete  HttpMethod = "DELETE"
	HttpConnect HttpMethod = "CONNECT"
	HttpOptions HttpMethod = "OPTIONS"
	HttpTrace   HttpMethod = "TRACE"
)

type Action struct {
	Method   HttpMethod
	Path     string
	Handlers []gin.HandlerFunc
}

func NewAction(method HttpMethod, path string, handlers ...gin.HandlerFunc) *Action {
	return &Action{Method: method, Path: path, Handlers: handlers}
}
func NewActionGet(path string, handlers ...gin.HandlerFunc) *Action {
	return &Action{Method: HttpGet, Path: path, Handlers: handlers}
}
func NewActionPost(path string, handlers ...gin.HandlerFunc) *Action {
	return &Action{Method: HttpPost, Path: path, Handlers: handlers}
}

type Module struct {
	Path        string
	MiddleWares []gin.HandlerFunc
	Actions     []*Action
	SubModules  []*Module
}

func NewModule(path string, middleWares []gin.HandlerFunc, actions []*Action, subModules []*Module) *Module {
	return &Module{Path: path, MiddleWares: middleWares, Actions: actions, SubModules: subModules}
}

func (m *Module) SetSubModules(modules ...*Module) {
	m.SubModules = modules
}
func (m *Module) Router(router *gin.RouterGroup) *gin.RouterGroup {
	moduleRouter := router.Group(m.Path)
	for _, ware := range m.MiddleWares {
		moduleRouter.Use(ware)
	}
	for _, action := range m.Actions {
		switch action.Method {
		case HttpGet:
			moduleRouter.GET(action.Path, action.Handlers...)
		case HttpPost:
			moduleRouter.POST(action.Path, action.Handlers...)
		case HttpOptions:
			moduleRouter.OPTIONS(action.Path, action.Handlers...)
		}
	}
	for _, sub := range m.SubModules {
		sub.Router(moduleRouter)
	}
	return moduleRouter
}
