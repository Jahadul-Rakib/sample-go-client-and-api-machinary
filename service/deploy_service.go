package service

import (
	"github.com/labstack/echo/v4"
	"go_client/common"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var deployConfig = common.GetConfig()

func GetAllDeployment(context echo.Context) error {

	deploymentName := make(map[int]string)

	deploymentList, ERROR := deployConfig.ClientSet.AppsV1().Deployments("default").List(deployConfig.Context, v1.ListOptions{})
	if ERROR != nil {
		return common.ErrorResponse(context, ERROR.Error(), "Error")
	}

	for i, deployment := range deploymentList.Items {
		deploymentName[i] = deployment.Name
	}
	return common.SuccessResponse(context, "Get All Deployment Name", deploymentName)
}
