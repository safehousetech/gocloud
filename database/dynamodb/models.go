package dynamodb

//Dynamodb struct reperesnts  aws Dynamodb service.
type Dynamodb struct {
}

//AttributeDefinitions struct reperesnts  Createtable AttributeDefinitions.
type AttributeDefinitions struct {
	AttributeName string `json:"AttributeName"`
	AttributeType string `json:"AttributeType"`
}

//KeySchema struct reperesnts  Createtable KeySchema.
type KeySchema struct {
	AttributeName string `json:"AttributeName"`
	KeyType       string `json:"KeyType"`
}

//Projection struct reperesnts  Createtable Projection.
type Projection struct {
	NonKeyAttributes []string `json:"NonKeyAttributes"`
	ProjectionType   string   `json:"ProjectionType"`
}

//GlobalSecondaryIndexes struct reperesnts  Createtable GlobalSecondaryIndexes.
type GlobalSecondaryIndexes struct {
	IndexName             string                `json:"IndexName"`
	KeySchema             []KeySchema           `json:"KeySchema"`
	Projection            Projection            `json:"Projection"`
	ProvisionedThroughput ProvisionedThroughput `json:"ProvisionedThroughput"`
}

//Createtable struct reperesnts  Createtable.
type Createtable struct {
	AttributeDefinitions   []AttributeDefinitions   `json:"AttributeDefinitions"`
	GlobalSecondaryIndexes []GlobalSecondaryIndexes `json:"GlobalSecondaryIndexes"`
	LocalSecondaryIndexes  []LocalSecondaryIndexes  `json:"LocalSecondaryIndexes"`
	KeySchema              []KeySchema              `json:"KeySchema"`
	ProvisionedThroughput  ProvisionedThroughput    `json:"ProvisionedThroughput"`
	SSESpecification       SSESpecification         `json:"SSESpecification"`
	StreamSpecification    StreamSpecification      `json:"StreamSpecification"`
	TableName              string                   `json:"TableName"`
}

//LocalSecondaryIndexes struct reperesnts Createtable LocalSecondaryIndexes.
type LocalSecondaryIndexes struct {
	IndexName  string      `json:"IndexName"`
	KeySchema  []KeySchema `json:"KeySchema"`
	Projection Projection  `json:"Projection"`
}

//ProvisionedThroughput struct reperesnts  Createtable ProvisionedThroughput.
type ProvisionedThroughput struct {
	ReadCapacityUnits  int `json:"ReadCapacityUnits"`
	WriteCapacityUnits int `json:"WriteCapacityUnits"`
}

//SSESpecification struct reperesnts  Createtable SSESpecification.
type SSESpecification struct {
	Enabled bool `json:"Enabled"`
}

//StreamSpecification struct reperesnts Createtable StreamSpecification.
type StreamSpecification struct {
	StreamViewType string `json:"StreamViewType"`
	StreamEnabled  bool   `json:"StreamEnabled"`
}
