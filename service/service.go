package service

import (
	"github.com/labstack/echo/v4"
	"go_client/common"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var serviceConfig = common.GetConfig()

func GetAllService(context echo.Context) error {

	serviceName := make(map[int]string)

	serviceList, ERROR := serviceConfig.ClientSet.CoreV1().
		Services("default").List(serviceConfig.Context, metaV1.ListOptions{})
	if ERROR != nil {
		return common.ErrorResponse(context, ERROR.Error(), "Error")
	}

	for i, service := range serviceList.Items {
		serviceName[i] = service.Name
	}
	return common.SuccessResponse(context, "Get All Service Name", serviceName)
}

func GetService(context echo.Context) error {
	name := context.Param("name")
	service, err := serviceConfig.ClientSet.
		CoreV1().Services(metaV1.NamespaceDefault).Get(serviceConfig.Context, name, metaV1.GetOptions{})
	if err != nil {
		return common.ErrorResponse(context, err.Error(), "Error")
	}
	return common.SuccessResponse(context, "Get Deployment", service)
}

func DeleteService(context echo.Context) error {
	name := context.Param("name")
	err := serviceConfig.ClientSet.
		CoreV1().Services(metaV1.NamespaceDefault).Delete(serviceConfig.Context, name, metaV1.DeleteOptions{})
	if err != nil {
		return common.ErrorResponse(context, err.Error(), "Service delete Error Occur.")
	}
	return common.SuccessResponse(context, "Service Successful", "Done!!!")
}

func CreateService(context echo.Context) error {
	payload, err := common.FitServicePayload(context)
	if err != nil {
		return common.ErrorResponse(context, err.Error(), "Binding Error Occur.")
	}
	service, err := serviceConfig.ClientSet.CoreV1().Services(metaV1.NamespaceDefault).
		Create(serviceConfig.Context, payload, metaV1.CreateOptions{})
	if err != nil {
		return common.ErrorResponse(context, err.Error(), "Creation Service Error Occur.")
	}
	return common.SuccessResponse(context, "Service Successfully created", service)
}

//func UpdateService(context echo.Context) error {
//	payload, err := common.FitServicePayload(context)
//	if err != nil {
//		return common.ErrorResponse(context, err.Error(), "Binding Error Occur.")
//	}
//
//
//}
