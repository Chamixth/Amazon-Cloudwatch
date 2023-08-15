package dao

import (
	awsConfig "amazon-cloudwatch/awsConfig"
	"amazon-cloudwatch/dto"
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

func PutDashboard(application dto.Dashboard) (*cloudwatch.PutDashboardOutput, error) {
	client, err := awsConfig.AwsSession()

	if err != nil {
		return nil, err
	}

	// Convert DashboardBody struct to JSON string
	dashboardBodyJSON, err := json.Marshal(application.DashboardBody)
	if err != nil {
		return nil, err
	}

	result, err := client.PutDashboard(&cloudwatch.PutDashboardInput{
		DashboardName: &application.DashboardName,
		DashboardBody:aws.String(string(dashboardBodyJSON)) ,
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}
