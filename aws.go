package utils

import (
	"fmt"
	// "net/http"
	// "html/template"
	// "log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"k8s.io/client-go/kubernetes"
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

func AWSGetParameter() {
	mySession := session.Must(session.NewSession())

	// Create a SSM client from just a session.
	svc := ssm.New(mySession)

	// Create a SSM client with additional configuration
	svc := ssm.New(mySession, aws.NewConfig().WithRegion("us-east-2"))

	output, err := ssm.GetParameter(&GetParameterOutput{
		Name:           aws.String("/patreon/client-id"),
		WithDecryption: aws.Bool(true),
	})

	Check(err)

	fmt.Printf("%+v\n", output)

}
