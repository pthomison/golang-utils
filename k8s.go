package utils

import (
    "fmt"
    // "net/http"
    // "html/template"
    // "log"
    "path/filepath"
    "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/api/core/v1"
	"context"

)

func GetClientSet() (*kubernetes.Clientset, error) {
	cs, err := internalClientSet()
	if err == nil {
		return cs, nil
	} else {
		fmt.Println(err)
		return externalClientSet()
	}
}

func externalClientSet() (*kubernetes.Clientset, error) {
	home := homedir.HomeDir()	
	kubeconfig := filepath.Join(home, ".kube", "config")

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, err
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return clientset, nil
}

func internalClientSet() (*kubernetes.Clientset, error) {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return clientset, nil
}

func GetPods(cs *kubernetes.Clientset, namespace string) (*v1.PodList, error) {
    listOptions:= metav1.ListOptions{}
    pods, err:=  cs.CoreV1().Pods(namespace).List(context.TODO(), listOptions)
    return pods, err
}