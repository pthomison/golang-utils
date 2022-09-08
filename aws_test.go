package utils

import (
	"testing"
)

func TestAWSGetParameter(t *testing.T) {
	param, err := AWSGetParameter("/placeholder", "us-east-2")
	Check(err)

	if param != "placeholderValue" {
		t.Logf("Unexpected Parameter Value: %s\n", param)
		t.Fail()
	}
}
