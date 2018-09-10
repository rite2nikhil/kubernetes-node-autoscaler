package metrics

import (
	"os"

	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type KubeNodeClient interface {
	ListNodes() (objs *v1.NodeList, err error)
}

// KubeClient k8s client
type KubeClient struct {
	// Main Kubernetes client
	ClientSet *kubernetes.Clientset
}

// NewKubeClient new kubernetes api client
func NewKubeClient() (*KubeClient, error) {
	clientset, err := getkubeclient()
	if err != nil {
		return nil, err
	}
	return &KubeClient{ClientSet: clientset}, nil
}

func (k8s *KubeClient) ListNodes() (objs *v1.NodeList, err error) {
	return k8s.ClientSet.CoreV1().Nodes().List(metav1.ListOptions{})
}

func getkubeclient() (*kubernetes.Clientset, error) {
	config, err := buildConfig()
	if err != nil {
		return nil, err
	}
	// creates the clientset
	kubeClient, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return kubeClient, err
}

// Create the client config. Use kubeconfig if given, otherwise assume in-cluster.
func buildConfig() (*rest.Config, error) {
	kubeconfigPath := os.Getenv("KUBECONFIG")
	if kubeconfigPath != "" {
		return clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	}
	return rest.InClusterConfig()
}
