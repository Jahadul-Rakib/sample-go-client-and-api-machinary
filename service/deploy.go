package service

import (
	"github.com/labstack/echo/v4"
	"go_client/common"
	appsV1 "k8s.io/api/apps/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var deployConfig = common.GetConfig()

func GetAllDeployment(context echo.Context) error {
	deploymentName := make(map[int]appsV1.Deployment)

	deploymentList, err := deployConfig.ClientSet.
		AppsV1().Deployments(metaV1.NamespaceDefault).List(deployConfig.Context, metaV1.ListOptions{})
	if err != nil {
		return common.ErrorResponse(context, err.Error(), "Error")
	}

	for i, deployment := range deploymentList.Items {
		deploymentName[i] = deployment
	}
	return common.SuccessResponse(context, "Get All Deployment.", deploymentName)
}

func GetDeployment(context echo.Context) error {
	name := context.Param("name")
	deployment, err := deployConfig.ClientSet.
		AppsV1().Deployments(metaV1.NamespaceDefault).Get(deployConfig.Context, name, metaV1.GetOptions{})
	if err != nil {
		return common.ErrorResponse(context, err.Error(), "Error")
	}
	return common.SuccessResponse(context, "Get Deployment", deployment)
}

func CreateDeployment(context echo.Context) error {

	deployPayload, err := common.FitDeployPayload(context)
	if err != nil {
		return common.ErrorResponse(context, err.Error(), "Binding Error Occur.")
	}
	deployment, err := deployConfig.ClientSet.
		AppsV1().Deployments(metaV1.NamespaceDefault).Create(deployConfig.Context, deployPayload, metaV1.CreateOptions{})
	if err != nil {
		return common.ErrorResponse(context, err.Error(), "Deployment Error Occur.")
	}
	return common.SuccessResponse(context, "Deployed Successful", deployment)
}

func UpdateDeployment(context echo.Context) error {
	deploy := new(common.DeployPayload)
	if err := context.Bind(deploy); err != nil {
		return common.ErrorResponse(context, err.Error(), "Data binding error!!!")
	}
	name := context.Param("name")
	deployment, err := deployConfig.ClientSet.
		AppsV1().Deployments(metaV1.NamespaceDefault).Get(deployConfig.Context, name, metaV1.GetOptions{})
	if err != nil {
		return common.ErrorResponse(context, err.Error(), "Error Not Found!!!")
	}

	deployment.Spec.Replicas = deploy.Replica
	deployment.Spec.Template.Spec.Containers[0].Image = deploy.Image

	upd, err := deployConfig.ClientSet.AppsV1().Deployments(metaV1.NamespaceDefault).
		Update(deployConfig.Context, deployment, metaV1.UpdateOptions{})
	if err != nil {
		return common.ErrorResponse(context, err.Error(), "Updated Error.")
	}
	return common.SuccessResponse(context, "Updated", upd)
}

func DeleteDeployment(context echo.Context) error {
	name := context.Param("name")
	_, err := deployConfig.ClientSet.
		AppsV1().Deployments(metaV1.NamespaceDefault).Get(deployConfig.Context, name, metaV1.GetOptions{})
	if err != nil {
		return common.ErrorResponse(context, err.Error(), "Data Not Found!!!")
	}
	err = deployConfig.ClientSet.
		AppsV1().Deployments(metaV1.NamespaceDefault).Delete(deployConfig.Context, name, metaV1.DeleteOptions{})
	if err != nil {
		return common.ErrorResponse(context, err.Error(), "Deployment delete Error Occur.")
	}
	return common.SuccessResponse(context, "Deployed Successful", "Done!!!")
}
