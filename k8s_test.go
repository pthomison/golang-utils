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

	pods, err := GetAllPods(cs)
	Check(err)

	fmt.Printf("%+v", pods)
}