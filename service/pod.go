package service

import (
	"github.com/labstack/echo/v4"
	"go_client/common"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var podConfig = common.GetConfig()

func GetAllPods(context echo.Context) error {

	podName := make(map[int]string)

	pods, ERROR := podConfig.ClientSet.CoreV1().Pods("default").List(podConfig.Context, v1.ListOptions{})
	if ERROR != nil {
		return common.ErrorResponse(context, ERROR.Error(), "Cluster Error")
	}

	for i, pod := range pods.Items {
		podName[i] = pod.Name
	}
	return common.SuccessResponse(context, "Get All Pod Name", podName)
}
