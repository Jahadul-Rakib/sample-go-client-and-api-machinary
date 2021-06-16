package main

import (
	"github.com/labstack/echo/v4"
	"go_client/service"
)

func Router(e *echo.Echo) {

	kubeApiDeployRouter := e.Group("api/v1/deployment")
	DeployRouter(kubeApiDeployRouter)

	kubeApiServiceRouter := e.Group("api/v1/service")
	ServiceRouter(kubeApiServiceRouter)
}

func DeployRouter(baseRouter *echo.Group) {
	baseRouter.GET("", service.GetAllDeployment)
	baseRouter.GET("/:name", service.GetDeployment)
	baseRouter.POST("", service.CreateDeployment)
	baseRouter.PUT("/:name", service.UpdateDeployment)
	baseRouter.DELETE("/:name", service.DeleteDeployment)
}

func ServiceRouter(baseRouter *echo.Group) {
	baseRouter.GET("", service.GetAllService)
	baseRouter.GET("/:name", service.GetService)
	baseRouter.DELETE("/:name", service.DeleteService)
	baseRouter.POST("", service.CreateService)
	baseRouter.PUT("/:name", service.UpdateService)
}
