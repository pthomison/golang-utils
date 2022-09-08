package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

func AWSGetParameter(parameterName string, region string) (string, error) {
	sesh := session.Must(session.NewSession())
	svc := ssm.New(sesh, aws.NewConfig().WithRegion(region))

	output, err := svc.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String(parameterName),
		WithDecryption: aws.Bool(true),
	})

	if err != nil {
		return "", err
	}

	return *output.Parameter.Value, nil

}
