package dynamodb

func preparedescribetables(params map[string]string, TableName string, Region string) {

	if TableName != "" {
		params["TableName"] = TableName
	}

	if Region != "" {
		params["Region"] = Region
	}

	params["amztarget"] = "DynamoDB_20120810.DescribeTable"
}

func preparedeletetables(params map[string]string, TableName string, Region string) {

	if TableName != "" {
		params["TableName"] = TableName
	}

	if Region != "" {
		params["Region"] = Region
	}

	params["amztarget"] = "DynamoDB_20120810.DeleteTable"
}

func preparelisttables(params map[string]string, Region string) {

	if Region != "" {
		params["Region"] = Region
	}
	params["amztarget"] = "DynamoDB_20120810.ListTables"
}

func preparelisttablesparamsdict(listtablesjsonmap map[string]interface{}, ExclusiveStartTableName string, Limit int) {

	if ExclusiveStartTableName != "" {
		listtablesjsonmap["ExclusiveStartTableName"] = ExclusiveStartTableName
	}

	if Limit != 0 {
		listtablesjsonmap["Limit"] = Limit
	}
}

func preparecreatetableStreamSpecificationparams(createtablejsonmap map[string]interface{}, createtable Createtable) {

	if (createtable.StreamSpecification != StreamSpecification{}) {

		streamSpecificationv := make(map[string]interface{})

		if createtable.StreamSpecification.StreamViewType != "" {
			streamSpecificationv["StreamViewType"] = createtable.StreamSpecification.StreamViewType
		}

		if createtable.StreamSpecification.StreamEnabled {
			streamSpecificationv["StreamEnabled"] = createtable.StreamSpecification.StreamEnabled
		}

		createtablejsonmap["StreamSpecification"] = streamSpecificationv
	}
}

func preparecreatetableSSESpecificationparams(createtablejsonmap map[string]interface{}, createtable Createtable) {
	if (createtable.SSESpecification != SSESpecification{}) {
		sSESpecificationv := make(map[string]interface{})
		sSESpecificationv["Enabled"] = createtable.SSESpecification.Enabled
		createtablejsonmap["SSESpecification"] = sSESpecificationv
	}
}

func preparecreatetableProvisionedThroughputparams(createtablejsonmap map[string]interface{}, createtable Createtable) {
	if (createtable.ProvisionedThroughput != ProvisionedThroughput{}) {
		provisionedThroughputv := make(map[string]interface{})
		if createtable.ProvisionedThroughput.ReadCapacityUnits != 0 {
			provisionedThroughputv["ReadCapacityUnits"] = createtable.ProvisionedThroughput.ReadCapacityUnits
		}
		if createtable.ProvisionedThroughput.WriteCapacityUnits != 0 {
			provisionedThroughputv["WriteCapacityUnits"] = createtable.ProvisionedThroughput.WriteCapacityUnits
		}
		createtablejsonmap["ProvisionedThroughput"] = provisionedThroughputv
	}
}

func prepareAttributeDefinitionsparams(createtablejsonmap map[string]interface{}, createtable Createtable) {
	if len(createtable.AttributeDefinitions) != 0 {
		attributeDefinitionvs := make([]map[string]interface{}, 0)

		for i := 0; i < len(createtable.KeySchema); i++ {
			attributeDefinitionv := make(map[string]interface{})

			if createtable.AttributeDefinitions[i].AttributeName != "" {
				attributeDefinitionv["AttributeName"] = createtable.AttributeDefinitions[i].AttributeName
			}

			if createtable.AttributeDefinitions[i].AttributeType != "" {
				attributeDefinitionv["AttributeType"] = createtable.AttributeDefinitions[i].AttributeType
			}

			attributeDefinitionvs = append(attributeDefinitionvs, attributeDefinitionv)
		}

		createtablejsonmap["AttributeDefinitions"] = attributeDefinitionvs
	}
}

func preparekeySchemaparams(createtablejsonmap map[string]interface{}, createtable Createtable) {
	if len(createtable.KeySchema) != 0 {
		keySchemavs := make([]map[string]interface{}, 0)

		for i := 0; i < len(createtable.KeySchema); i++ {
			keySchemav := make(map[string]interface{})

			if createtable.KeySchema[i].AttributeName != "" {
				keySchemav["AttributeName"] = createtable.KeySchema[i].AttributeName
			}

			if createtable.KeySchema[i].KeyType != "" {
				keySchemav["KeyType"] = createtable.KeySchema[i].KeyType
			}

			keySchemavs = append(keySchemavs, keySchemav)
		}

		createtablejsonmap["KeySchema"] = keySchemavs
	}
}

func prepareGlobalSecondaryIndexesparams(createtablejsonmap map[string]interface{}, createtable Createtable) {

	if len(createtable.GlobalSecondaryIndexes) != 0 {

		globalSecondaryIndexesvarrayjsonmap := make([]map[string]interface{}, 0)

		for i := 0; i < len(createtable.GlobalSecondaryIndexes); i++ {

			globalSecondaryIndexessvjsonmap := make(map[string]interface{})

			if createtable.GlobalSecondaryIndexes[i].IndexName != "" {
				globalSecondaryIndexessvjsonmap["IndexName"] = createtable.GlobalSecondaryIndexes[i].IndexName
			}

			p := Projection{}

			if createtable.GlobalSecondaryIndexes[i].Projection.ProjectionType == p.ProjectionType && len(createtable.GlobalSecondaryIndexes[i].Projection.NonKeyAttributes) == len(p.NonKeyAttributes) {

				projectionv := make(map[string]interface{})
				projectionv["ProjectionType"] = createtable.GlobalSecondaryIndexes[i].Projection.ProjectionType
				projectionv["NonKeyAttributes"] = createtable.GlobalSecondaryIndexes[i].Projection.NonKeyAttributes
				globalSecondaryIndexessvjsonmap["Projection"] = projectionv
			}

			if (createtable.GlobalSecondaryIndexes[i].ProvisionedThroughput != ProvisionedThroughput{}) {

				provisionedThroughputv := make(map[string]interface{})
				provisionedThroughputv["ReadCapacityUnits"] = createtable.GlobalSecondaryIndexes[i].ProvisionedThroughput.ReadCapacityUnits
				provisionedThroughputv["WriteCapacityUnits"] = createtable.GlobalSecondaryIndexes[i].ProvisionedThroughput.WriteCapacityUnits
				globalSecondaryIndexessvjsonmap["ProvisionedThroughput"] = provisionedThroughputv
			}

			if len(createtable.GlobalSecondaryIndexes[i].KeySchema) != 0 {

				keySchemavs := make([]map[string]interface{}, 0)

				for i := 0; i < len(createtable.GlobalSecondaryIndexes[i].KeySchema); i++ {

					keySchemav := make(map[string]interface{})

					if createtable.GlobalSecondaryIndexes[i].KeySchema[i].AttributeName != "" {
						keySchemav["AttributeName"] = createtable.GlobalSecondaryIndexes[i].KeySchema[i].AttributeName
					}

					if createtable.LocalSecondaryIndexes[i].KeySchema[i].KeyType != "" {
						keySchemav["KeyType"] = createtable.GlobalSecondaryIndexes[i].KeySchema[i].KeyType
					}

					keySchemavs = append(keySchemavs, keySchemav)
				}

				globalSecondaryIndexessvjsonmap["KeySchema"] = keySchemavs
			}

			globalSecondaryIndexesvarrayjsonmap = append(globalSecondaryIndexesvarrayjsonmap, globalSecondaryIndexessvjsonmap)
		}

		createtablejsonmap["GlobalSecondaryIndexes"] = globalSecondaryIndexesvarrayjsonmap
	}

}

func prepareLocalSecondaryIndexesparams(createtablejsonmap map[string]interface{}, createtable Createtable) {

	if len(createtable.LocalSecondaryIndexes) != 0 {

		localSecondaryIndexesvarrayjsonmap := make([]map[string]interface{}, 0)

		for i := 0; i < len(createtable.LocalSecondaryIndexes); i++ {

			localSecondaryIndexesvjsonmap := make(map[string]interface{})

			if createtable.LocalSecondaryIndexes[i].IndexName != "" {
				localSecondaryIndexesvjsonmap["IndexName"] = createtable.LocalSecondaryIndexes[i].IndexName
			}

			p := Projection{}

			if createtable.LocalSecondaryIndexes[i].Projection.ProjectionType == p.ProjectionType && len(createtable.LocalSecondaryIndexes[i].Projection.NonKeyAttributes) == len(p.NonKeyAttributes) {
				projectionv := make(map[string]interface{})
				projectionv["ProjectionType"] = createtable.LocalSecondaryIndexes[i].Projection.ProjectionType
				projectionv["NonKeyAttributes"] = createtable.LocalSecondaryIndexes[i].Projection.NonKeyAttributes
				localSecondaryIndexesvjsonmap["Projection"] = projectionv
			}

			if len(createtable.LocalSecondaryIndexes[i].KeySchema) != 0 {

				lenv := len(createtable.LocalSecondaryIndexes[i].KeySchema)

				keySchemavs := make([]map[string]interface{}, 0)

				for j := 0; j < lenv; i++ {

					keySchemav := make(map[string]interface{})
					keySchemav["AttributeName"] = createtable.LocalSecondaryIndexes[i].KeySchema[j].AttributeName

					if createtable.LocalSecondaryIndexes[i].KeySchema[j].KeyType != "" {
						keySchemav["KeyType"] = createtable.LocalSecondaryIndexes[i].KeySchema[j].KeyType
					}

					keySchemavs = append(keySchemavs, keySchemav)
				}

				localSecondaryIndexesvjsonmap["KeySchema"] = keySchemavs
			}

			localSecondaryIndexesvarrayjsonmap = append(localSecondaryIndexesvarrayjsonmap, localSecondaryIndexesvjsonmap)
		}
		createtablejsonmap["LocalSecondaryIndexes"] = localSecondaryIndexesvarrayjsonmap
	}

}

func preparecreatetable(params map[string]string, Region string) {

	if Region != "" {
		params["Region"] = Region
	}

	params["amztarget"] = "DynamoDB_20120810.CreateTable"
}

func preparecreatetablejsonmap(createtablejsonmap map[string]interface{}, createtable Createtable) {

	if createtable.TableName != "" {
		createtablejsonmap["TableName"] = createtable.TableName
	}

	preparecreatetableStreamSpecificationparams(createtablejsonmap, createtable)
	preparecreatetableSSESpecificationparams(createtablejsonmap, createtable)
	preparecreatetableProvisionedThroughputparams(createtablejsonmap, createtable)
	preparekeySchemaparams(createtablejsonmap, createtable)
	prepareAttributeDefinitionsparams(createtablejsonmap, createtable)
	prepareGlobalSecondaryIndexesparams(createtablejsonmap, createtable)
}
