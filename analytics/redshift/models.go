package redshift

//Redshift struct reperesnts aws alianalytics service.
type Redshift struct {
}

//Describecluster .
type Describecluster struct {
	clusterIdentifier string
	marker            string
	maxRecords        int
	tagKeys           []string
	tagValues         []string
}

//DeleteCluster .
type DeleteCluster struct {
	clusterIdentifier              string
	finalClusterSnapshotIdentifier string
	skipFinalClusterSnapshot       bool
}

//CreateCluster .
type CreateCluster struct {
	clusterIdentifier                string
	masterUsername                   string
	masterUserPassword               string
	nodeType                         string
	additionalInfo                   string
	allowVersionUpgrade              bool
	automatedSnapshotRetentionPeriod string
	availabilityZone                 string
	clusterParameterGroupName        string
	clusterType                      string
	clusterVersion                   string
	dBName                           string
	elasticIP                        string
	encrypted                        bool
	enhancedVpcRouting               bool
	hsmClientCertificateIdentifier   string
	hsmConfigurationIdentifier       string
	kmsKeyID                         string
	numberOfNodes                    int
	port                             int
	preferredMaintenanceWindow       string
	publiclyAccessible               bool
	tagKeys                          []string
	tagValues                        []string
	iamRoles                         []string
	clusterSecurityGroups            []string
	vpcSecurityGroupIds              []string
}
