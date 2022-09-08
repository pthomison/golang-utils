package utils

import (
	"bytes"
	"errors"
	"testing"
)

func TestK8SGetClientSet(t *testing.T) {
	_, err := GetClientSet()
	CheckTest(err, t)
}

func TestK8SGetPods(t *testing.T) {
	cs, err := GetClientSet()
	CheckTest(err, t)

	pods, err := GetPods(cs, "")
	CheckTest(err, t)

	t.Log("---- All Pods ----")
	for _, p := range pods.Items {
		t.Logf("%+v\n", p.Name)
	}
	t.Log()

}

func TestK8SGetDeployments(t *testing.T) {
	cs, err := GetClientSet()
	CheckTest(err, t)

	deployments, err := GetDeployments(cs, "")
	CheckTest(err, t)

	t.Log("---- All Deployments ----")
	for _, d := range deployments.Items {
		t.Logf("%+v\n", d.Name)
	}
	t.Log()

}

func TestK8SGetSecrets(t *testing.T) {
	cs, err := GetClientSet()
	CheckTest(err, t)

	secrets, err := GetSecrets(cs, "")
	CheckTest(err, t)

	t.Log("---- All Secrets ----")
	for _, d := range secrets.Items {
		t.Logf("%+v\n", d.Name)
	}
	t.Log()
}

func TestK8SGetSecret(t *testing.T) {
	cs, err := GetClientSet()
	CheckTest(err, t)

	secret, err := GetSecret(cs, "k3s-serving", "kube-system")
	CheckTest(err, t)

	t.Log("---- k3s-serving/kube-system Secret ----")
	for k, v := range secret.Data {
		t.Logf("%+v %+s\n", k, v)
	}
	t.Log()
}

func TestK8SApplySecret(t *testing.T) {
	cs, err := GetClientSet()
	CheckTest(err, t)

	data := make(Secret)
	data["test_key"] = []byte("test_value")

	_, err = ApplySecret(cs, "test-apply-secret", "default", data)
	CheckTest(err, t)

	err = DeleteSecret(cs, "test-apply-secret", "default")
	CheckTest(err, t)
}

func TestK8SUpdateSecret(t *testing.T) {
	secretName := "test-update-secret"
	secretNamespace := "default"

	keyA := "key_a"
	valueA := "value_a"

	keyB := "key_b"
	valueB := "value_b"

	cs, err := GetClientSet()
	CheckTest(err, t)

	// Create An Empty Secret
	emptyData := make(Secret)
	_, err = ApplySecret(cs, secretName, secretNamespace, emptyData)
	CheckTest(err, t)

	// Add one piece of data
	aData := make(Secret)
	aData[keyA] = []byte(valueA)
	_, err = UpdateSecret(cs, secretName, secretNamespace, aData)
	CheckTest(err, t)

	// Add a second piece of data
	bData := make(Secret)
	bData[keyB] = []byte(valueB)
	_, err = UpdateSecret(cs, secretName, secretNamespace, bData)
	CheckTest(err, t)

	// Ensure that both pieces of data are present
	secret, err := GetSecret(cs, secretName, secretNamespace)
	CheckTest(err, t)

	if (bytes.Compare(secret.Data[keyA], []byte(valueA)) != 0) || (bytes.Compare(secret.Data[keyB], []byte(valueB)) != 0) {
		CheckTest(errors.New("Retrieved Data Does Not Match Injected Data"), t)
	}

	// Clean Up Secret
	err = DeleteSecret(cs, secretName, secretNamespace)
	CheckTest(err, t)

	// Ensure Clean Up
	_, err = GetSecret(cs, secretName, secretNamespace)
	if err == nil {
		CheckTest(errors.New("Secret Not Correctly Deleted"), t)
	}

}
