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
	clouddataflow.CloudDataFlow
}

//Compute ...
func (*Google) Compute() gocloudinterface.Compute {
	return &gce.GCE{}
}

//Storage ...
func (*Google) Storage() gocloudinterface.Storage {
	return &googlestorage.GoogleStorage{}
}

//LoadBalancer ...
func (*Google) LoadBalancer() gocloudinterface.LoadBalancer {
	return &googleloadbalancer.Googleloadbalancer{}
}

//Container ...
func (*Google) Container() gocloudinterface.Container {
	return &googlecontainer.Googlecontainer{}
}

//DNS ...
func (*Google) DNS() gocloudinterface.DNS {
	return &googledns.Googledns{}
}

//Serverless ...
func (*Google) Serverless() gocloudinterface.Serverless {
	return &googlecloudfunctions.Googlecloudfunctions{}
}

//Database ...
func (*Google) Database() gocloudinterface.Database {
	return &bigtable.Bigtable{}
}

//Analytics ...
func (*Google) Analytics() gocloudinterface.Analytics {
	return &bigquery.Bigquery{}
}

//Notification ...
func (*Google) Notification() gocloudinterface.Notification {
	return &googlenotification.Googlenotification{}
}

//MachineLearning ...
func (*Google) MachineLearning() gocloudinterface.MachineLearning {
	return &googlemachinelearning.Googlemachinelearning{}
}

//Streamdataprocessing ...
func (*Google) Streamdataprocessing() gocloudinterface.Streamdataprocessing {
	return &clouddataflow.CloudDataFlow{}
}
