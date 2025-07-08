package routers

type EnterUserRouter struct {
	UserRouter
}

var UserRouterInstance = new(EnterUserRouter)
