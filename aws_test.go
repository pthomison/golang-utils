package utils

import (
	"fmt"
	"testing"
)

func TestAWSGetParameter(t *testing.T) {
	param, err := AWSGetParameter("/placeholder", "us-east-2")
	Check(err)

	fmt.Printf("%+v", param)
}
