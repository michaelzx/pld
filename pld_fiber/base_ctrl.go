package pld_fiber

import "github.com/gofiber/fiber/v2"

type BaseCtrl struct {
	baseRouterIsSet bool
	baseRouter      fiber.Router
	middlewareList  []fiber.Handler
}

func (bc *BaseCtrl) checkRoot() {
	if !bc.baseRouterIsSet {
		panic("baseRouter 未设置")
	}
}
func (bc *BaseCtrl) AddMiddleware(middleware fiber.Handler) {
	bc.middlewareList = append(bc.middlewareList, middleware)
}

func (bc *BaseCtrl) SetBaseRouter(parentRoute fiber.Router) {
	bc.baseRouterIsSet = true
	bc.baseRouter = parentRoute
}

func (bc *BaseCtrl) Get(path string, handlers ...fiber.Handler) {
	bc.checkRoot()
	var finalHandlers []fiber.Handler
	finalHandlers = append(finalHandlers, bc.middlewareList...)
	finalHandlers = append(finalHandlers, handlers...)
	bc.baseRouter.Get(path, finalHandlers...)
}
func (bc *BaseCtrl) Post(path string, handlers ...fiber.Handler) {
	bc.checkRoot()
	var finalHandlers []fiber.Handler
	finalHandlers = append(finalHandlers, bc.middlewareList...)
	finalHandlers = append(finalHandlers, handlers...)
	bc.baseRouter.Post(path, finalHandlers...)
}
