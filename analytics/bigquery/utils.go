package bigquery

const (
	//UnixDate ...
	UnixDate = "Mon Jan _2 15:04:05 MST 2006"
	//RFC3339 ...
	RFC3339 = "2006-01-02T15:04:05Z07:00"
)

func createdatasetsdictnoaryconvert(option CreateDatasets, createdatasetsjsonmap map[string]interface{}) {

	if option.DefaultTableExpirationMs != "" {
		createdatasetsjsonmap["defaultTableExpirationMs"] = option.DefaultTableExpirationMs
	}

	if option.DefaultTableExpirationMs != "" {
		createdatasetsjsonmap["defaultTableExpirationMs"] = option.DefaultTableExpirationMs
	}

	if option.DefaultTableExpirationMs != "" {
		createdatasetsjsonmap["description"] = option.Description
	}

	if option.Etag != "" {
		createdatasetsjsonmap["etag"] = option.Etag
	}

	if option.ID != "" {
		createdatasetsjsonmap["id"] = option.ID
	}

	if option.FriendlyName != "" {
		createdatasetsjsonmap["friendlyName"] = option.FriendlyName
	}

	if option.Kind != "" {
		createdatasetsjsonmap["kind"] = option.Kind
	}

	if option.LastModifiedTime != "" {
		createdatasetsjsonmap["lastModifiedTime"] = option.LastModifiedTime
	}

	if option.Location != "" {
		createdatasetsjsonmap["location"] = option.Location
	}

	if option.SelfLink != "" {
		createdatasetsjsonmap["selfLink"] = option.SelfLink
	}

	preparedatasetReferenceparam(option, createdatasetsjsonmap)
	prepareAccessparam(option, createdatasetsjsonmap)
}

func preparedatasetReferenceparam(option CreateDatasets, createdatasetsjsonmap map[string]interface{}) {

	datasetReferencejsonmap := make(map[string]interface{})

	if option.DatasetReference.ProjectID != "" {
		datasetReferencejsonmap["projectId"] = option.DatasetReference.ProjectID
	}

	if option.DatasetReference.DatasetID != "" {
		datasetReferencejsonmap["datasetId"] = option.DatasetReference.DatasetID
	}

	createdatasetsjsonmap["datasetReference"] = datasetReferencejsonmap
}

func prepareAccessparam(option CreateDatasets, createdatasetsjsonmap map[string]interface{}) {

	if len(option.Access) != 0 {

		accessarrayjsonmap := make([]map[string]interface{}, 0)

		for i := 0; i < len(option.Access); i++ {

			accessjsonmap := make(map[string]interface{})

			if option.Access[i].Domain != "" {
				accessjsonmap["domain"] = option.Access[i].Domain
			}

			if option.Access[i].GroupByEmail != "" {
				accessjsonmap["groupByEmail"] = option.Access[i].GroupByEmail
			}

			if option.Access[i].Role != "" {
				accessjsonmap["role"] = option.Access[i].Role
			}

			if option.Access[i].SpecialGroup != "" {
				accessjsonmap["specialGroup"] = option.Access[i].SpecialGroup
			}

			if option.Access[i].UserByEmail != "" {
				accessjsonmap["userByEmail"] = option.Access[i].UserByEmail
			}

			v := View{}

			if option.Access[i].View != v {

				viewjsonmap := make(map[string]interface{})

				if option.Access[i].View.DatasetID != "" {
					viewjsonmap["datasetID"] = option.Access[i].View.DatasetID
				}

				if option.Access[i].View.ProjectID != "" {
					viewjsonmap["projectID"] = option.Access[i].View.ProjectID
				}

				if option.Access[i].View.TableID != "" {
					viewjsonmap["tableID"] = option.Access[i].View.TableID
				}

				accessjsonmap["view"] = viewjsonmap
			}

			accessarrayjsonmap = append(accessarrayjsonmap, accessjsonmap)
		}
		createdatasetsjsonmap["access"] = accessarrayjsonmap
	}
}
