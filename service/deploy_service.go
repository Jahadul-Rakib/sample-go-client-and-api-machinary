package service

import (
	"github.com/labstack/echo/v4"
	"go_client/common"
	appsV1 "k8s.io/api/apps/v1"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

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

	deployment := &appsV1.Deployment{
		ObjectMeta: metaV1.ObjectMeta{
			Name: "demo-deployment",
		},
		Spec: appsV1.DeploymentSpec{
			Replicas: int32Cnv(1),
			Selector: &metaV1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "demo",
				},
			},
			Template: coreV1.PodTemplateSpec{
				ObjectMeta: metaV1.ObjectMeta{
					Labels: map[string]string{
						"app": "demo",
					},
				},
				Spec: coreV1.PodSpec{
					Containers: []coreV1.Container{
						{
							Name:  "web",
							Image: "nginx:1.12",
							Ports: []coreV1.ContainerPort{
								{
									Name:          "http",
									Protocol:      coreV1.ProtocolTCP,
									ContainerPort: 8081,
								},
							},
						},
					},
				},
			},
		},
	}

	deployment, err := deployConfig.ClientSet.
		AppsV1().Deployments(metaV1.NamespaceDefault).Create(deployConfig.Context, deployment, metaV1.CreateOptions{})
	if err != nil {
		return common.ErrorResponse(context, err.Error(), "Deployment Error Occure.")
	}
	return common.SuccessResponse(context, "Deployed Successful", deployment)
}

func int32Cnv(i int32) *int32 { return &i }
