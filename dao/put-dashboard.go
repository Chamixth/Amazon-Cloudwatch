package dao

import (
	awsConfig "amazon-cloudwatch/awsConfig"
	"amazon-cloudwatch/dto"

	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

func PutDashboard(application dto.Dashboard) (*cloudwatch.PutDashboardOutput, error) {
	client, err := awsConfig.AwsSession()

	if err != nil {
		return nil, err
	}

	result, err := client.PutDashboard(&cloudwatch.PutDashboardInput{
		DashboardName: application.DashboardName,
		DashboardBody: application.DashboardBody,
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}
