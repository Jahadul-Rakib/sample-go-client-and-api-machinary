package service

import (
	"github.com/labstack/echo/v4"
	"go_client/common"
	coreV1 "k8s.io/api/core/v1"
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

func UpdateService(context echo.Context) error {
	service := new(common.ServicePayload)
	if err := context.Bind(service); err != nil {
		return common.ErrorResponse(context, err.Error(), "Data binding error!!!")
	}
	name := context.Param("name")

	var getService, err = serviceConfig.ClientSet.CoreV1().Services(metaV1.NamespaceDefault).
		Get(serviceConfig.Context, name, metaV1.GetOptions{})
	if err != nil {
		return common.ErrorResponse(context, err.Error(), "Data getting error!!!")
	}
	//if &service.ServiceName != nil {
	//	get_service.Name = service.ServiceName
	//}
	if &service.Selector != nil {
		getService.Spec.Selector = map[string]string{
			"app": service.Selector,
		}
	}
	if &service.Type != nil {
		switch service.Type {
		case "inside":
			getService.Spec.Type = coreV1.ServiceTypeClusterIP
		case "outside":
			getService.Spec.Type = coreV1.ServiceTypeLoadBalancer
		default:
			getService.Spec.Type = coreV1.ServiceTypeLoadBalancer
		}
	}
	//if &service.Port != nil {
	//	get_service.Spec.Ports.
	//}
	updated, err := serviceConfig.ClientSet.CoreV1().Services(metaV1.NamespaceDefault).
		Update(serviceConfig.Context, getService, metaV1.UpdateOptions{})
	if err != nil {
		return common.ErrorResponse(context, err.Error(), "Data Updated error!!!")
	}
	return common.SuccessResponse(context, "Updated successfully", updated)
}
