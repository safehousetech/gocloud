package clouddataflow

//CloudDataFlow struct reperesnts google streaming service.
type CloudDataFlow struct {
}

//CreateStream struct reperesnts Cloud dataflow service.
type CreateStream struct {
	ID                   string               `json:"id"`
	ClientRequestID      string               `json:"clientRequestId"`
	CreateTime           string               `json:"createTime"`
	CurrentState         string               `json:"currentState"`
	CurrentStateTime     string               `json:"currentStateTime"`
	Location             string               `json:"location"`
	Name                 string               `json:"name"`
	ProjectID            string               `json:"projectId"`
	ReplaceJobID         string               `json:"replaceJobId"`
	ReplacedByJobID      string               `json:"replacedByJobId"`
	RequestedState       string               `json:"requestedState"`
	TempFiles            []string             `json:"tempFiles"`
	Type                 string               `json:"type"`
	Environment          Environment          `json:"environment"`
	JobMetadata          JobMetadata          `json:"jobMetadata"`
	Labels               Labels               `json:"labels"`
	PipelineDescription  PipelineDescription  `json:"pipelineDescription"`
	StageStates          []StageStates        `json:"stageStates"`
	Steps                []Steps              `json:"steps"`
	TransformNameMapping TransformNameMapping `json:"transformNameMapping"`
	ExecutionInfo        ExecutionInfo        `json:"executionInfo"`
}

//StageStates  struct reperesnts Clouddataflow internal params.
type StageStates struct {
	CurrentStateTime    string `json:"currentStateTime"`
	ExecutionStageName  string `json:"executionStageName"`
	ExecutionStageState string `json:"executionStageState"`
}

//Properties  struct reperesnts Clouddataflow internal params.
type Properties struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
}

//ExecutionInfo  struct reperesnts Clouddataflow internal params.
type ExecutionInfo struct {
	Stages Stages `json:"stages"`
}

//Stages  struct reperesnts Clouddataflow internal params.
type Stages struct {
}

//TransformNameMapping  struct reperesnts Clouddataflow internal params.
type TransformNameMapping struct {
}

//Steps  struct reperesnts Clouddataflow internal params.
type Steps struct {
	Kind       string     `json:"kind"`
	Name       string     `json:"name"`
	Properties Properties `json:"properties"`
}

//DisplayData  struct reperesnts Clouddataflow internal params.
type DisplayData struct {
}

//OriginalPipelineTransform  struct reperesnts Clouddataflow internal params.
type OriginalPipelineTransform struct {
}

//ExecutionPipelineStage ...
type ExecutionPipelineStage struct {
}

//PipelineDescription  struct reperesnts Clouddataflow internal params.
type PipelineDescription struct {
	DisplayData               []DisplayData               `json:"displayData"`
	ExecutionPipelineStage    []ExecutionPipelineStage    `json:"executionPipelineStage"`
	OriginalPipelineTransform []OriginalPipelineTransform `json:"originalPipelineTransform"`
}

//JobMetadata  struct reperesnts Clouddataflow internal params.
type JobMetadata struct {
}

//Labels  struct reperesnts Clouddataflow internal params.
type Labels struct {
}

//InternalExperiments  struct reperesnts Clouddataflow internal params.
type InternalExperiments struct {
}

//SdkPipelineOptions  struct reperesnts Clouddataflow internal params.
type SdkPipelineOptions struct {
}

//Environment  struct reperesnts Clouddataflow internal params.
type Environment struct {
	ClusterManagerAPIService string              `json:"clusterManagerApiService"`
	Dataset                  string              `json:"dataset"`
	Experiments              []string            `json:"experiments"`
	InternalExperiments      InternalExperiments `json:"internalExperiments"`
	SdkPipelineOptions       SdkPipelineOptions  `json:"sdkPipelineOptions"`
	ServiceAccountEmail      string              `json:"serviceAccountEmail"`
	TempStoragePrefix        string              `json:"tempStoragePrefix"`
	UserAgent                UserAgent           `json:"userAgent"`
	Version                  Version             `json:"version"`
	WorkerPools              WorkerPools         `json:"workerPools"`
}

//UserAgent  struct reperesnts Clouddataflow internal params.
type UserAgent struct {
	Name      string  `json:"name"`
	Support   Support `json:"support"`
	BuildDate string  `json:"build.date"`
	Version   string  `json:"version"`
}

//Support  struct reperesnts Clouddataflow internal params.
type Support struct {
	Status string `json:"status"`
	URL    string `json:"url"`
}

//Version  struct reperesnts Clouddataflow internal params.
type Version struct {
	Major   string `json:"major"`
	JobType string `json:"job_type"`
}

//Metadata  struct reperesnts Clouddataflow internal params.
type Metadata struct {
}

//TaskrunnerSettings  struct reperesnts Clouddataflow internal params.
type TaskrunnerSettings struct {
	Alsologtostderr          bool                   `json:"alsologtostderr"`
	BaseTaskDir              string                 `json:"baseTaskDir"`
	BaseURL                  string                 `json:"baseUrl"`
	CommandlinesFileName     string                 `json:"commandlinesFileName"`
	ContinueOnException      bool                   `json:"continueOnException"`
	DataflowAPIVersion       string                 `json:"dataflowApiVersion"`
	HarnessCommand           string                 `json:"harnessCommand"`
	LanguageHint             string                 `json:"languageHint"`
	LogDir                   string                 `json:"logDir"`
	LogToSerialconsole       bool                   `json:"logToSerialconsole"`
	LogUploadLocation        string                 `json:"logUploadLocation"`
	OauthScopes              []string               `json:"oauthScopes"`
	ParallelWorkerSettings   ParallelWorkerSettings `json:"parallelWorkerSettings"`
	StreamingWorkerMainClass string                 `json:"streamingWorkerMainClass"`
	TaskGroup                string                 `json:"taskGroup"`
	TaskUser                 string                 `json:"taskUser"`
	TempStoragePrefix        string                 `json:"tempStoragePrefix"`
	VMID                     string                 `json:"vmId"`
	WorkflowFileName         string                 `json:"workflowFileName"`
}

//ParallelWorkerSettings  struct reperesnts Clouddataflow internal params.
type ParallelWorkerSettings struct {
	BaseURL            string `json:"baseUrl"`
	ReportingEnabled   bool   `json:"reportingEnabled"`
	ServicePath        string `json:"servicePath"`
	ShuffleServicePath string `json:"shuffleServicePath"`
	TempStoragePrefix  string `json:"tempStoragePrefix"`
	WorkerID           string `json:"workerId"`
}

//PoolArgs  struct reperesnts Clouddataflow internal params.
type PoolArgs struct {
}

//Packages  struct reperesnts Clouddataflow internal params.
type Packages struct {
	Location string `json:"location"`
	Name     string `json:"name"`
}

//DataDisks  struct reperesnts Clouddataflow internal params.
type DataDisks struct {
	DiskType   string `json:"diskType"`
	MountPoint string `json:"mountPoint"`
	SizeGb     int    `json:"sizeGb"`
}

//AutoscalingSettings  struct reperesnts Clouddataflow internal params.
type AutoscalingSettings struct {
	Algorithm     string `json:"algorithm"`
	MaxNumWorkers int    `json:"maxNumWorkers"`
}

//WorkerPools  struct reperesnts Clouddataflow internal params.
type WorkerPools struct {
	AutoscalingSettings         AutoscalingSettings `json:"autoscalingSettings"`
	DataDisks                   []DataDisks         `json:"dataDisks"`
	DefaultPackageSet           string              `json:"defaultPackageSet"`
	DiskSizeGb                  int                 `json:"diskSizeGb"`
	DiskSourceImage             string              `json:"diskSourceImage"`
	DiskType                    string              `json:"diskType"`
	IPConfiguration             string              `json:"ipConfiguration"`
	Kind                        string              `json:"kind"`
	MachineType                 string              `json:"machineType"`
	Metadata                    Metadata            `json:"metadata"`
	Network                     string              `json:"network"`
	NumThreadsPerWorker         int                 `json:"numThreadsPerWorker"`
	NumWorkers                  int                 `json:"numWorkers"`
	OnHostMaintenance           string              `json:"onHostMaintenance"`
	Packages                    []Packages          `json:"packages"`
	PoolArgs                    PoolArgs            `json:"poolArgs"`
	Subnetwork                  string              `json:"subnetwork"`
	TaskrunnerSettings          TaskrunnerSettings  `json:"taskrunnerSettings"`
	TeardownPolicy              string              `json:"teardownPolicy"`
	WorkerHarnessContainerImage string              `json:"workerHarnessContainerImage"`
	Zone                        string              `json:"zone"`
}
