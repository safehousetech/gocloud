package nova

// CreateServer WIP
type CreateServer struct {
	Server                  Server                 `json:"server"`
	OOSSCHHNTSchedulerHints OSSCHHNTSchedulerHints `json:"OS-SCH-HNT:scheduler_hints"`
}

//OSSCHHNTSchedulerHints .
type OSSCHHNTSchedulerHints struct {
	SameHost string `json:"same_host"`
}

//Server .
type Server struct {
	AccessIPv4 string `json:"accessIPv4"`

	AccessIPv6 string `json:"accessIPv6"`

	Name string `json:"name"`

	Networks string `json:"networks"`

	ImageRef string `json:"imageRef"`

	FlavorRef string `json:"flavorRef"`

	AvailabilityZone string `json:"availability_zone"`

	OSDCFDiskConfig string `json:"OS-DCF:diskConfig"`

	Personality []Personality `json:"personality"`

	UserData string `json:"user_data"`

	MetaData Metadata `json:"metadata"`

	SecurityGroups []SecurityGroups `json:"security_groups"`
}

//Personality .
type Personality struct {
	Path     string `json:"path"`
	Contents string `json:"contents"`
}

//Metadata .
type Metadata struct {
	MyServerName string `json:"My Server Name"`
}

//SecurityGroups .
type SecurityGroups struct {
	Name string `json:"name"`
}
