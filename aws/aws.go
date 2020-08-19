package aws

import (
	redshift "github.com/safehousetech/gocloud/analytics/redshift"
	ec2 "github.com/safehousetech/gocloud/compute/ec2"
	awscontainer "github.com/safehousetech/gocloud/container/awscontainer"
	dynamodb "github.com/safehousetech/gocloud/database/dynamodb"
	awsdns "github.com/safehousetech/gocloud/dns/awsdns"
	"github.com/safehousetech/gocloud/gocloudinterface"
	awsloadbalancer "github.com/safehousetech/gocloud/loadbalancer/awsloadbalancer"
	"github.com/safehousetech/gocloud/machinelearning/awsmachinelearning"
	amazonsimplenotification "github.com/safehousetech/gocloud/notification/amazonsimplenotification"
	lambda "github.com/safehousetech/gocloud/serverless/lambda"
	amazonstorage "github.com/safehousetech/gocloud/storage/amazonstorage"
	kinesis "github.com/safehousetech/gocloud/streamdataprocessing/kinesis"
)

//AWS struct reperents amazon cloud provider.
type AWS struct {
	ec2.EC2
	amazonstorage.Amazonstorage
	awsloadbalancer.Awsloadbalancer
	awscontainer.Ecscontainer
	awsdns.Awsdns
	lambda.Lambda
	dynamodb.Dynamodb
	awsmachinelearning.Awsmachinelearning
	amazonsimplenotification.Amazonsimplenotification
	redshift.Redshift
	kinesis.Kinesis
}

func (*AWS) Compute() gocloudinterface.Compute {
	return &ec2.EC2{}
}

func (*AWS) Storage() gocloudinterface.Storage {
	return &amazonstorage.Amazonstorage{}
}

func (*AWS) LoadBalancer() gocloudinterface.LoadBalancer {
	return &awsloadbalancer.Awsloadbalancer{}
}

func (*AWS) Container() gocloudinterface.Container {
	return &awscontainer.Ecscontainer{}
}

func (*AWS) DNS() gocloudinterface.DNS {
	return &awsdns.Awsdns{}
}

func (*AWS) Serverless() gocloudinterface.Serverless {
	return &lambda.Lambda{}
}

func (*AWS) Database() gocloudinterface.Database {
	return &dynamodb.Dynamodb{}
}

func (*AWS) MachineLearning() gocloudinterface.MachineLearning {
	return &awsmachinelearning.Awsmachinelearning{}
}

func (*AWS) Analytics() gocloudinterface.Analytics {
	return &redshift.Redshift{}
}

func (*AWS) Notification() gocloudinterface.Notification {
	return &amazonsimplenotification.Amazonsimplenotification{}
}

func (*AWS) Streamdataprocessing() gocloudinterface.Streamdataprocessing {
	return &kinesis.Kinesis{}
}
