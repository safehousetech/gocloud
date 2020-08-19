package google

import (
	bigquery "github.com/safehousetech/gocloud/analytics/bigquery"
	gce "github.com/safehousetech/gocloud/compute/gce"
	googlecontainer "github.com/safehousetech/gocloud/container/googlecontainer"
	bigtable "github.com/safehousetech/gocloud/database/bigtable"
	googledns "github.com/safehousetech/gocloud/dns/googledns"
	"github.com/safehousetech/gocloud/gocloudinterface"
	googleloadbalancer "github.com/safehousetech/gocloud/loadbalancer/googleloadbalancer"
	googlemachinelearning "github.com/safehousetech/gocloud/machinelearning/googlemachinelearning"
	googlenotification "github.com/safehousetech/gocloud/notification/googlenotification"
	googlecloudfunctions "github.com/safehousetech/gocloud/serverless/googlecloudfunctions"
	googlestorage "github.com/safehousetech/gocloud/storage/googlestorage"
	clouddataflow "github.com/safehousetech/gocloud/streamdataprocessing/clouddataflow"
)

// Google  struct represents Google Cloud provider.
type Google struct {
	gce.GCE
	googlestorage.GoogleStorage
	googleloadbalancer.Googleloadbalancer
	googlecontainer.Googlecontainer
	googledns.Googledns
	googlecloudfunctions.Googlecloudfunctions
	bigtable.Bigtable
	googlemachinelearning.Googlemachinelearning
	bigquery.Bigquery
	googlenotification.Googlenotification
	clouddataflow.Clouddataflow
}

func (*Google) Compute() gocloudinterface.Compute {
	return &gce.GCE{}
}

func (*Google) Storage() gocloudinterface.Storage {
	return &googlestorage.GoogleStorage{}
}

func (*Google) LoadBalancer() gocloudinterface.LoadBalancer {
	return &googleloadbalancer.Googleloadbalancer{}
}

func (*Google) Container() gocloudinterface.Container {
	return &googlecontainer.Googlecontainer{}
}

func (*Google) DNS() gocloudinterface.DNS {
	return &googledns.Googledns{}
}

func (*Google) Serverless() gocloudinterface.Serverless {
	return &googlecloudfunctions.Googlecloudfunctions{}
}

func (*Google) Database() gocloudinterface.Database {
	return &bigtable.Bigtable{}
}

func (*Google) Analytics() gocloudinterface.Analytics {
	return &bigquery.Bigquery{}
}

func (*Google) Notification() gocloudinterface.Notification {
	return &googlenotification.Googlenotification{}
}

func (*Google) MachineLearning() gocloudinterface.MachineLearning {
	return &googlemachinelearning.Googlemachinelearning{}
}

func (*Google) Streamdataprocessing() gocloudinterface.Streamdataprocessing {
	return &clouddataflow.Clouddataflow{}
}
