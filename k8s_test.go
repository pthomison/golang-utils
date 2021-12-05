package utils

import (
	"fmt"
	"testing"
)

func TestK8SGetClientSet(t *testing.T) {
	cs, err := GetClientSet()
	Check(err)

	fmt.Printf("%+v", cs)
}

func TestK8SGetPods(t *testing.T) {
	cs, err := GetClientSet()
	Check(err)

	pods, err := GetPods(cs, "")
	Check(err)

	fmt.Println("---- All Pods ----")
	for _, p := range pods.Items {
		fmt.Printf("%+v\n", p.Name)
	}
}

func TestK8SGetDeployments(t *testing.T) {
	cs, err := GetClientSet()
	Check(err)

	deployments, err := GetDeployments(cs, "")
	Check(err)

	fmt.Println("---- All Deployments ----")
	for _, d := range deployments.Items {
		fmt.Printf("%+v\n", d.Name)
	}
}

func TestK8SGetSecrets(t *testing.T) {
	cs, err := GetClientSet()
	Check(err)

	secrets, err := GetSecrets(cs, "")
	Check(err)

	fmt.Println("---- All Secrets ----")
	for _, d := range secrets.Items {
		fmt.Printf("%+v\n", d.Name)
	}
}

func TestK8SGetSecret(t *testing.T) {
	cs, err := GetClientSet()
	Check(err)

	secret, err := GetSecret(cs, "flux-system", "flux-system")
	Check(err)

	fmt.Println("---- Flux System Secret ----")
	for k, v := range secret.Data {
		fmt.Printf("%+v %+s\n", k, v)
		fmt.Printf("%T %T\n", k, v)
	}
}

func TestK8SApplySecret(t *testing.T) {
	cs, err := GetClientSet()
	Check(err)

	data := make(Secret)

	data["test_key"] = []byte("test_value")

	fmt.Printf("%+v\n", data)

	result, err := ApplySecret(cs, "test-apply-secret", "default", data)
	Check(err)

	fmt.Printf("%+v\n", result)

	err = DeleteSecret(cs, "test-apply-secret", "default")
	Check(err)
}

func TestK8SUpdateSecret(t *testing.T) {
	cs, err := GetClientSet()
	Check(err)

	// Update Start

	emptyData := make(Secret)

	result, err := ApplySecret(cs, "test-update-secret", "default", emptyData)
	Check(err)

	fmt.Printf("Update Start: %+v\n", result)

	// Update A

	aData := make(Secret)
	aData["key_a"] = []byte("value_a")

	result, err = UpdateSecret(cs, "test-update-secret", "default", aData)
	Check(err)

	fmt.Printf("Update A: %+v\n", result)

	// Update B

	bData := make(Secret)
	bData["key_b"] = []byte("value_b")

	result, err = UpdateSecret(cs, "test-update-secret", "default", bData)
	Check(err)

	fmt.Printf("Update B: %+v\n", result)

	err = DeleteSecret(cs, "test-update-secret", "default")
	Check(err)

}
