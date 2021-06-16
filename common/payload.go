package common

import (
	"github.com/labstack/echo/v4"
	appsV1 "k8s.io/api/apps/v1"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

type DeployPayload struct {
	Port     int    `json:"port"`
	Image    string `json:"image"`
	Replica  *int32 `json:"replica"`
	AppsName string `json:"apps_name"`
}

func FitDeployPayload(context echo.Context) (*appsV1.Deployment, error) {

	data := new(DeployPayload)
	if err := context.Bind(data); err != nil {
		return nil, err
	}
	deployment := &appsV1.Deployment{
		ObjectMeta: metaV1.ObjectMeta{
			Name: data.AppsName,
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

type ServicePayload struct {
	Selector    string `json:"selector"`
	Type        string `json:"type"`
	Port        int32  `json:"port"`
	TargetPort  int    `json:"target_port"`
	ServiceName string `json:"service_name"`
}

func FitServicePayload(context echo.Context) (*coreV1.Service, error) {
	data := new(ServicePayload)
	if err := context.Bind(data); err != nil {
		return nil, err
	}
	var types coreV1.ServiceType

	switch data.Type {
	case "inside":
		types = coreV1.ServiceTypeClusterIP
	case "outside":
		types = coreV1.ServiceTypeLoadBalancer
	default:
		types = coreV1.ServiceTypeLoadBalancer
	}

	service := &coreV1.Service{
		ObjectMeta: metaV1.ObjectMeta{
			Name: data.ServiceName,
		},
		Spec: coreV1.ServiceSpec{
			Selector: map[string]string{
				"app": data.Selector,
			},
			Type: types,
			Ports: []coreV1.ServicePort{
				{Name: data.ServiceName,
					Protocol:   coreV1.ProtocolTCP,
					Port:       data.Port,
					TargetPort: intstr.FromInt(data.TargetPort),
				},
			},
		},
	}

	return service, nil
}
