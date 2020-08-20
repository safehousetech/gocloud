package bigquery

//Bigquery struct reperesnts google analytics service.
type Bigquery struct {
}

//View  struct reperesnts CreateDatasets internal struct.
type View struct {
	DatasetID string `json:"datasetId"`
	ProjectID string `json:"projectId"`
	TableID   string `json:"tableId"`
}

//Access struct reperesnts CreateDatasets internal struct.
type Access struct {
	Domain       string `json:"domain"`
	GroupByEmail string `json:"groupByEmail"`
	Role         string `json:"role"`
	SpecialGroup string `json:"specialGroup"`
	UserByEmail  string `json:"userByEmail"`
	View         View   `json:"view"`
}

//DatasetReference struct reperesnts CreateDatasets internal struct.
type DatasetReference struct {
	DatasetID string `json:"datasetId"`
	ProjectID string `json:"projectId"`
}

//CreateDatasets struct reperesnts google bigquery CreateDatasets service struct.
type CreateDatasets struct {
	Access                   []Access         `json:"access"`
	DatasetReference         DatasetReference `json:"datasetReference"`
	DefaultTableExpirationMs string           `json:"defaultTableExpirationMs"`
	Description              string           `json:"description"`
	Etag                     string           `json:"etag"`
	FriendlyName             string           `json:"friendlyName"`
	ID                       string           `json:"id"`
	Kind                     string           `json:"kind"`
	LastModifiedTime         string           `json:"lastModifiedTime"`
	CreationTime             string           `json:"creationTime"`
	Location                 string           `json:"location"`
	SelfLink                 string           `json:"selfLink"`
}
