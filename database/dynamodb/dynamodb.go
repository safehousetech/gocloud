package dynamodb

//ListTables /List tables.
func (dynamodb *Dynamodb) ListTables(request interface{}) (resp interface{}, err error) {

	param := request.(map[string]interface{})

	var ExclusiveStartTableName, Region string
	var Limit int
	for key, value := range param {
		switch key {
		case "ExclusiveStartTableName":
			ExclusiveStartTableNameV, _ := value.(string)
			ExclusiveStartTableName = ExclusiveStartTableNameV

		case "Limit":
			LimitV, _ := value.(int)
			Limit = LimitV

		case "Region":
			RegionV, _ := value.(string)
			Region = RegionV
		}
	}

	params := make(map[string]string)

	preparelisttables(params, Region)

	listtablesjsonmap := make(map[string]interface{})

	preparelisttablesparamsdict(listtablesjsonmap, ExclusiveStartTableName, Limit)
	response := make(map[string]interface{})

	err = dynamodb.PrepareSignatureV4query(params, listtablesjsonmap, response)
	resp = response
	return resp, err
}

//DeleteTables Delete tables.
func (dynamodb *Dynamodb) DeleteTables(request interface{}) (resp interface{}, err error) {
	param := request.(map[string]interface{})

	var TableName, Region string

	for key, value := range param {
		switch key {
		case "TableName":
			TableNameV, _ := value.(string)
			TableName = TableNameV

		case "Region":
			RegionV, _ := value.(string)
			Region = RegionV
		}
	}

	params := make(map[string]string)

	preparedeletetables(params, TableName, Region)

	deletetablesjsonmap := map[string]interface{}{
		"TableName": TableName,
	}

	response := make(map[string]interface{})
	err = dynamodb.PrepareSignatureV4query(params, deletetablesjsonmap, response)
	resp = response
	return resp, err
}

//CreateTables Create tables.
func (dynamodb *Dynamodb) CreateTables(request interface{}) (resp interface{}, err error) {
	param := request.(map[string]interface{})
	var createtable Createtable
	var Region string

	for key, value := range param {
		switch key {

		case "Region":
			RegionV, _ := value.(string)
			Region = RegionV

		case "TableName":
			TableNameV, _ := value.(string)
			createtable.TableName = TableNameV

		case "StreamSpecification":
			streamSpecificationparam, _ := value.(map[string]interface{})
			for streamSpecificationparamkey, streamSpecificationparamvalue := range streamSpecificationparam {
				switch streamSpecificationparamkey {
				case "StreamViewType":
					createtable.StreamSpecification.StreamViewType = streamSpecificationparamvalue.(string)

				case "StreamEnabled":
					createtable.StreamSpecification.StreamEnabled = streamSpecificationparamvalue.(bool)
				}
			}

		case "SSESpecification":
			sSESpecificationparam, _ := value.(map[string]interface{})
			for sSESpecificationparamkey, sSESpecificationparamparamvalue := range sSESpecificationparam {
				switch sSESpecificationparamkey {
				case "Enabled":
					createtable.SSESpecification.Enabled = sSESpecificationparamparamvalue.(bool)
				}
			}

		case "ProvisionedThroughput":
			provisionedThroughputparam, _ := value.(map[string]interface{})

			for provisionedThroughputparamkey, provisionedThroughputparamvalue := range provisionedThroughputparam {
				switch provisionedThroughputparamkey {
				case "ReadCapacityUnits":
					ReadCapacityUnitsv, _ := provisionedThroughputparamvalue.(int)
					createtable.ProvisionedThroughput.ReadCapacityUnits = ReadCapacityUnitsv

				case "WriteCapacityUnits":
					createtable.ProvisionedThroughput.WriteCapacityUnits = provisionedThroughputparamvalue.(int)

				}
			}

		case "KeySchema":
			keySchemaparam, _ := value.([]map[string]interface{})
			for i := 0; i < len(keySchemaparam); i++ {
				var keySchema KeySchema
				for keySchemaparamkey, keySchemaparamvalue := range keySchemaparam[i] {
					switch keySchemaparamkey {
					case "AttributeName":
						keySchema.AttributeName = keySchemaparamvalue.(string)

					case "KeyType":
						keySchema.KeyType = keySchemaparamvalue.(string)
					}
				}
				createtable.KeySchema = append(createtable.KeySchema, keySchema)
			}
		case "LocalSecondaryIndexes":
			localSecondaryIndexesparam, _ := value.([]map[string]interface{})

			for i := 0; i < len(localSecondaryIndexesparam); i++ {
				var localSecondaryIndexes LocalSecondaryIndexes
				for localSecondaryIndexesparamkey, localSecondaryIndexesparamvalue := range localSecondaryIndexesparam[i] {
					switch localSecondaryIndexesparamkey {
					case "IndexName":
						localSecondaryIndexes.IndexName = localSecondaryIndexesparamvalue.(string)

					case "KeySchema":
						keySchemaparam, _ := localSecondaryIndexesparamvalue.([]map[string]interface{})
						for i := 0; i < len(keySchemaparam); i++ {
							var keySchema KeySchema
							for keySchemaparamkey, keySchemaparamvalue := range keySchemaparam[i] {
								switch keySchemaparamkey {
								case "AttributeName":
									keySchema.AttributeName = keySchemaparamvalue.(string)

								case "KeyType":
									keySchema.KeyType = keySchemaparamvalue.(string)
								}
							}

							localSecondaryIndexes.KeySchema = append(localSecondaryIndexes.KeySchema, keySchema)
						}

					case "Projection":
						projectionparam, _ := localSecondaryIndexesparamvalue.(map[string]interface{})
						for projectionparamkey, projectionparamvalue := range projectionparam {
							switch projectionparamkey {
							case "NonKeyAttributes":
								localSecondaryIndexes.Projection.NonKeyAttributes = projectionparamvalue.([]string)

							case "ProjectionType":
								localSecondaryIndexes.Projection.ProjectionType = projectionparamvalue.(string)
							}
						}
					}
				}
				createtable.LocalSecondaryIndexes = append(createtable.LocalSecondaryIndexes, localSecondaryIndexes)
			}

		case "globalSecondaryIndexes":
			globalSecondaryIndexesparam, _ := value.([]map[string]interface{})
			for i := 0; i < len(globalSecondaryIndexesparam); i++ {
				var globalSecondaryIndexes GlobalSecondaryIndexes
				for globalSecondaryIndexesparamkey, globalSecondaryIndexesparamvalue := range globalSecondaryIndexesparam[i] {
					switch globalSecondaryIndexesparamkey {
					case "IndexName":
						globalSecondaryIndexes.IndexName = globalSecondaryIndexesparamvalue.(string)

					case "keySchema":
						keySchemaparam, _ := globalSecondaryIndexesparamvalue.([]map[string]interface{})
						for i := 0; i < len(keySchemaparam); i++ {
							var keySchema KeySchema
							for keySchemaparamkey, keySchemaparamvalue := range keySchemaparam[i] {
								switch keySchemaparamkey {
								case "AttributeName":
									keySchema.AttributeName = keySchemaparamvalue.(string)

								case "KeyType":
									keySchema.KeyType = keySchemaparamvalue.(string)
								}
							}

							globalSecondaryIndexes.KeySchema = append(globalSecondaryIndexes.KeySchema, keySchema)
						}

					case "Projection":
						projectionparam, _ := globalSecondaryIndexesparamvalue.(map[string]interface{})
						for projectionparamkey, projectionparamvalue := range projectionparam {
							switch projectionparamkey {
							case "NonKeyAttributes":
								globalSecondaryIndexes.Projection.NonKeyAttributes = projectionparamvalue.([]string)

							case "ProjectionType":
								globalSecondaryIndexes.Projection.ProjectionType = projectionparamvalue.(string)
							}
						}

					case "ProvisionedThroughput":
						provisionedThroughputparam, _ := globalSecondaryIndexesparamvalue.(map[string]interface{})
						for provisionedThroughputparamkey, provisionedThroughputparamvalue := range provisionedThroughputparam {
							switch provisionedThroughputparamkey {

							case "ReadCapacityUnits":
								createtable.ProvisionedThroughput.ReadCapacityUnits = provisionedThroughputparamvalue.(int)

							case "WriteCapacityUnits":
								createtable.ProvisionedThroughput.WriteCapacityUnits = provisionedThroughputparamvalue.(int)
							}
						}

					}
				}
				createtable.GlobalSecondaryIndexes = append(createtable.GlobalSecondaryIndexes, globalSecondaryIndexes)
			}

		case "AttributeDefinitions":
			attributeDefinitionsparam, _ := value.([]map[string]interface{})
			for i := 0; i < len(attributeDefinitionsparam); i++ {
				var attributeDefinitions AttributeDefinitions
				for attributeDefinitionsparamkey, attributeDefinitionsparamvalue := range attributeDefinitionsparam[i] {
					switch attributeDefinitionsparamkey {
					case "AttributeName":
						attributeDefinitions.AttributeName = attributeDefinitionsparamvalue.(string)

					case "AttributeType":
						attributeDefinitions.AttributeType = attributeDefinitionsparamvalue.(string)
					}
				}
				createtable.AttributeDefinitions = append(createtable.AttributeDefinitions, attributeDefinitions)
			}
		}
	}
	params := make(map[string]string)

	preparecreatetable(params, Region)

	createtablejsonmap := map[string]interface{}{}

	preparecreatetablejsonmap(createtablejsonmap, createtable)

	response := make(map[string]interface{})
	err = dynamodb.PrepareSignatureV4query(params, createtablejsonmap, response)
	resp = response
	return resp, err
}

//DescribeTables Describe tables.
func (dynamodb *Dynamodb) DescribeTables(request interface{}) (resp interface{}, err error) {

	param := request.(map[string]interface{})

	var TableName, Region string

	for key, value := range param {
		switch key {
		case "TableName":
			TableNameV, _ := value.(string)
			TableName = TableNameV

		case "Region":
			RegionV, _ := value.(string)
			Region = RegionV
		}
	}

	params := make(map[string]string)

	preparedescribetables(params, TableName, Region)

	deletetablesjsonmap := map[string]interface{}{
		"TableName": TableName,
	}

	response := make(map[string]interface{})
	err = dynamodb.PrepareSignatureV4query(params, deletetablesjsonmap, response)
	resp = response
	return resp, err
}
