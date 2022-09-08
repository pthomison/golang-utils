package aws

import (
	"testing"

	utils "github.com/pthomison/golang-utils"
)

func TestAWSGetParameter(t *testing.T) {
	param, err := AWSGetParameter("/placeholder", "us-east-2")
	utils.CheckTest(err, t)

	if param != "placeholderValue" {
		t.Logf("Unexpected Parameter Value: %s\n", param)
		t.Fail()
	}
}
