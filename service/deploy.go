package service

import (
	"github.com/labstack/echo/v4"
	"go_client/common"
	appsV1 "k8s.io/api/apps/v1"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Payload struct {
	Port     int    `json:"port"`
	Image    string `json:"image"`
	Replica  *int32 `json:"replica"`
	AppsName string `json:"apps_name"`
}

var deployConfig = common.GetConfig()

func GetAllDeployment(context echo.Context) error {

	deploymentName := make(map[int]string)

	deploymentList, err := deployConfig.ClientSet.
		AppsV1().Deployments(metaV1.NamespaceDefault).List(deployConfig.Context, metaV1.ListOptions{})
	if err != nil {
		return common.ErrorResponse(context, err.Error(), "Error")
	}

	for i, deployment := range deploymentList.Items {
		deploymentName[i] = deployment.Name
	}
	return common.SuccessResponse(context, "Get All Deployment Name", deploymentName)
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

	deployPayload, err := fitDeployPayload(context)
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
	deploy := new(Payload)
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
		return common.ErrorResponse(context, err.Error(), "Error Not Found!!!")
	}
	err = deployConfig.ClientSet.
		AppsV1().Deployments(metaV1.NamespaceDefault).Delete(deployConfig.Context, name, metaV1.DeleteOptions{})
	if err != nil {
		return common.ErrorResponse(context, err.Error(), "Deployment delete Error Occur.")
	}
	return common.SuccessResponse(context, "Deployed Successful", "Done!!!")
}

func fitDeployPayload(context echo.Context) (*appsV1.Deployment, error) {

	data := new(Payload)
	if err := context.Bind(data); err != nil {
		return nil, err
	}
	deployment := &appsV1.Deployment{
		ObjectMeta: metaV1.ObjectMeta{
			Name: data.AppsName + "-deployment",
		},
		Spec: appsV1.DeploymentSpec{
			Replicas: data.Replica,
			Selector: &metaV1.LabelSelector{
				MatchLabels: map[string]string{
					"app": data.AppsName,
				},
			},
			Template: coreV1.PodTemplateSpec{
				ObjectMeta: metaV1.ObjectMeta{
					Labels: map[string]string{
						"app": data.AppsName,
					},
				},
				Spec: coreV1.PodSpec{
					Containers: []coreV1.Container{
						{
							Name:  data.AppsName,
							Image: data.Image,
							Ports: []coreV1.ContainerPort{
								{
									Name:          "http",
									Protocol:      coreV1.ProtocolTCP,
									ContainerPort: int32(data.Port),
								},
							},
						},
					},
				},
			},
		},
	}

	return deployment, nil
}
