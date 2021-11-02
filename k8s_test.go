package utils

import (
    "testing"
    "fmt"
)

func TestGetClientSet(t *testing.T) {
	cs, err := GetClientSet()
	Check(err)

	fmt.Printf("%+v", cs)
}

func TestGetPods(t *testing.T) {
	cs, err := GetClientSet()
	Check(err)

	pods, err := GetPods(cs, "")
	Check(err)

    fmt.Println("---- All Pods ----")
    for _, p := range pods.Items {
    	fmt.Printf("%+v\n", p.Name)
    }
}

func TestGetDeployments(t *testing.T) {
	cs, err := GetClientSet()
	Check(err)

	deployments, err := GetDeployments(cs, "")
	Check(err)

    fmt.Println("---- All Deployments ----")
    for _, d := range deployments.Items {
    	fmt.Printf("%+v\n", d.Name)
    }
}