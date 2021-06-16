package service

import (
	"github.com/labstack/echo/v4"
	"go_client/common"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var serviceConfig = common.GetConfig()

func GetAllService(context echo.Context) error {

	serviceName := make(map[int]string)

	serviceList, ERROR := serviceConfig.ClientSet.CoreV1().Services("default").List(serviceConfig.Context, v1.ListOptions{})
	if ERROR != nil {
		return common.ErrorResponse(context, ERROR.Error(), "Error")
	}

	for i, service := range serviceList.Items {
		serviceName[i] = service.Name
	}
	return common.SuccessResponse(context, "Get All Service Name", serviceName)
}
