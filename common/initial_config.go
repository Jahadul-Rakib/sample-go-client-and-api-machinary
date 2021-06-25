package common

import (
	"context"
	"flag"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"sync"
)

type Config struct {
	ClientSet *kubernetes.Clientset
	Context   context.Context
}

var configManager *Config
var runOnce sync.Once

func GetConfig() *Config {
	runOnce.Do(func() {
		configManager = &Config{}
		configManager.ConnectionKubernetesCluster()
	})
	return configManager
}

func (c *Config) ConnectionKubernetesCluster() {
	kubeconfig := flag.String("kubeconfig", "~/.kube/config", "kubeconfig file location")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		log.Println("Connection Error: ", err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Println("Connection Error: ", err.Error())
	}
	c.ClientSet = clientset
	c.Context = context.Background()
	log.Println("Connected to Cluster")
}
