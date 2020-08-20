package lambda

import "io/ioutil"

//GetFunction  describe lambda function.
func (lambda *Lambda) GetFunction(request interface{}) (resp interface{}, err error) {
	param := request.(map[string]interface{})
	var Region string
	var getfunction Getfunction

	for key, value := range param {
		switch key {
		case "FunctionName":
			FunctionNameV, _ := value.(string)
			getfunction.FunctionName = FunctionNameV

		case "Region":
			RegionV, _ := value.(string)
			Region = RegionV

		case "Qualifier":
			QualifierV, _ := value.(string)
			getfunction.Qualifier = QualifierV
		}
	}

	params := make(map[string]string)
	preparegetfunctionqueryString(params, getfunction)

	response := make(map[string]interface{})
	err = Preparegetfnrequest(params, Region, response)

	resp = response
	return resp, err
}

//CreateFunction  Create lambda function.
func (lambda *Lambda) CreateFunction(request interface{}) (resp interface{}, err error) {

	param := request.(map[string]interface{})
	var region string
	var createfunction CreateFunction

	for key, value := range param {
		switch key {

		case "Region":
			regionv, _ := value.(string)
			region = regionv

		case "FunctionName":
			functionNamev, _ := value.(string)
			createfunction.FunctionName = functionNamev

		case "Handler":
			handlerv, _ := value.(string)
			createfunction.Handler = handlerv

		case "KMSKeyArn":
			kMSKeyArnv, _ := value.(string)
			createfunction.KMSKeyArn = kMSKeyArnv

		case "MemorySize":
			memorySizev, _ := value.(int)
			createfunction.MemorySize = memorySizev

		case "Role":
			rolev, _ := value.(string)
			createfunction.Role = rolev

		case "Publish":
			publishv, _ := value.(bool)
			createfunction.Publish = publishv

		case "Runtime":
			runtimev, _ := value.(string)
			createfunction.Runtime = runtimev

		case "Tags":
			tagsv, _ := value.(string)
			createfunction.Tags.String = tagsv

		case "Description":
			descriptionv, _ := value.(string)
			createfunction.Description = descriptionv

		case "Timeout":
			timeoutv, _ := value.(int)
			createfunction.Timeout = timeoutv

		case "DeadLetterConfig":
			deadLetterConfigparam, _ := value.(map[string]string)

			for deadLetterConfigparamkey, deadLetterConfigparamvalue := range deadLetterConfigparam {

				switch deadLetterConfigparamkey {
				case "TargetArn":
					targetArnv := deadLetterConfigparamvalue
					createfunction.DeadLetterConfig.TargetArn = targetArnv
				}

			}

		case "TracingConfig":
			tracingConfigparam, _ := value.(map[string]string)
			for tracingConfigparamkey, tracingConfigparamvalue := range tracingConfigparam {

				switch tracingConfigparamkey {
				case "Mode":
					modev := tracingConfigparamvalue
					createfunction.TracingConfig.Mode = modev
				}

			}

		case "VpcConfig":
			vpcConfigparam, _ := value.(map[string][]string)
			for vpcConfigparamkey, vpcConfigparamvalue := range vpcConfigparam {

				switch vpcConfigparamkey {
				case "SubnetIds":
					subnetIdsv := vpcConfigparamvalue
					createfunction.VpcConfig.SubnetIds = subnetIdsv

				case "SecurityGroupIds":
					securityGroupIdsv := vpcConfigparamvalue
					createfunction.VpcConfig.SecurityGroupIds = securityGroupIdsv

				}
			}

		case "Code":
			codeparam, _ := value.(map[string]interface{})
			for codeparamkey, codeparamvalue := range codeparam {

				switch codeparamkey {
				case "S3Bucket":
					s3Bucketv, _ := codeparamvalue.(string)
					createfunction.Code.S3Bucket = s3Bucketv

				case "S3Key":
					s3Keyv, _ := codeparamvalue.(string)
					createfunction.Code.S3Key = s3Keyv

				case "S3ObjectVersion":
					s3ObjectVersionv, _ := codeparamvalue.(string)
					createfunction.Code.S3ObjectVersion = s3ObjectVersionv

				case "ZipFile":
					zipFilev, _ := codeparamvalue.(string)
					contents, _ := ioutil.ReadFile(zipFilev + ".zip")
					createfunction.Code.zipFile = contents

				case "Location":
					locationv, _ := codeparamvalue.(string)
					createfunction.Code.Location = locationv

				case "RepositoryType":
					repositoryTypev, _ := codeparamvalue.(string)
					createfunction.Code.RepositoryType = repositoryTypev
				}
			}

		}
	}

	params := make(map[string]interface{})
	preparecreatefunctiondict(params, createfunction)

	response := make(map[string]interface{})
	err = PreparePostrequest(params, region, response)

	resp = response
	return resp, err
}

func preparecreatefunctiondict(params map[string]interface{}, createfunction CreateFunction) {

	if createfunction.FunctionName != "" {
		params["FunctionName"] = createfunction.FunctionName
	}

	if createfunction.Handler != "" {
		params["Handler"] = createfunction.Handler
	}

	if createfunction.KMSKeyArn != "" {
		params["KMSKeyArn"] = createfunction.KMSKeyArn
	}

	if createfunction.MemorySize > 0 {
		params["MemorySize"] = createfunction.MemorySize
	}

	params["Publish"] = createfunction.Publish

	if createfunction.Role != "" {
		params["Role"] = createfunction.Role
	}

	if createfunction.Runtime != "" {
		params["Runtime"] = createfunction.Runtime
	}

	if createfunction.Tags.String != "" {
		params["Tags"] = createfunction.Tags.String
	}

	if createfunction.Description != "" {
		params["Description"] = createfunction.Description
	}

	if createfunction.Timeout > 0 {
		params["Timeout"] = createfunction.Timeout
	}

	if createfunction.DeadLetterConfig.TargetArn != "" {
		param := make(map[string]interface{})
		param["TargetArn"] = createfunction.DeadLetterConfig.TargetArn
		params["DeadLetterConfig"] = param
	}

	if createfunction.TracingConfig.Mode != "" {
		param := make(map[string]interface{})
		param["Mode"] = createfunction.TracingConfig.Mode
		params["TracingConfig"] = param
	}

	if createfunction.TracingConfig.Mode != "" {
		param := make(map[string]interface{})
		param["Mode"] = createfunction.TracingConfig.Mode
		params["TracingConfig"] = param
	}

	if (len(createfunction.VpcConfig.SecurityGroupIds) > 0) || (len(createfunction.VpcConfig.SubnetIds) > 0) {
		param := make(map[string]interface{})

		if len(createfunction.VpcConfig.SecurityGroupIds) > 0 {
			param["SecurityGroupIds"] = createfunction.VpcConfig.SecurityGroupIds
		}

		if len(createfunction.VpcConfig.SubnetIds) > 0 {
			param["SubnetIds"] = createfunction.VpcConfig.SubnetIds
		}

		params["VpcConfig"] = param
	}

	preparecode(params, createfunction)

}

func preparecode(params map[string]interface{}, createfunction CreateFunction) {

	param := make(map[string]interface{})

	if createfunction.Code.S3Bucket != "" {
		param["s3Bucket"] = createfunction.Code.S3Bucket
	}

	if createfunction.Code.S3Key != "" {
		param["S3Key"] = createfunction.Code.S3Key
	}

	if createfunction.Code.S3ObjectVersion != "" {
		param["S3ObjectVersion"] = createfunction.Code.S3ObjectVersion
	}

	if len(createfunction.Code.zipFile) > 0 {
		param["ZipFile"] = createfunction.Code.zipFile
	}

	if createfunction.Code.RepositoryType != "" {
		param["RepositoryType"] = createfunction.Code.RepositoryType
	}

	if createfunction.Code.Location != "" {
		param["Location"] = createfunction.Code.Location
	}

	params["Code"] = param

}

//CallFunction  call lambda function.
func (lambda *Lambda) CallFunction(request interface{}) (resp interface{}, err error) {

	return resp, err
}

//ListFunction  list lambda function.
func (lambda *Lambda) ListFunction(request interface{}) (resp interface{}, err error) {
	param := request.(map[string]interface{})
	var Region string
	var listfunction Listfunction

	for key, value := range param {
		switch key {

		case "FunctionVersion":
			functionVersionv, _ := value.(string)
			listfunction.functionVersion = functionVersionv

		case "Marker":
			markerv, _ := value.(string)
			listfunction.marker = markerv

		case "MasterRegion":
			masterRegionv, _ := value.(string)
			listfunction.masterRegion = masterRegionv

		case "MaxItems":
			maxItemsv, _ := value.(string)
			listfunction.maxItems = maxItemsv

		case "Region":
			RegionV, _ := value.(string)
			Region = RegionV
		}
	}

	params := make(map[string]string)
	preparelistfunctionqueryString(params, listfunction)

	response := make(map[string]interface{})
	err = Preparegetrequest(params, Region, response)

	resp = response
	return resp, err

}

//DeleteFunction  delete lambda function.
func (lambda *Lambda) DeleteFunction(request interface{}) (resp interface{}, err error) {
	param := request.(map[string]interface{})
	var Region string
	var deletefunction Deletefunction

	for key, value := range param {
		switch key {
		case "FunctionName":
			FunctionNameV, _ := value.(string)
			deletefunction.FunctionName = FunctionNameV

		case "Region":
			RegionV, _ := value.(string)
			Region = RegionV

		case "Qualifier":
			QualifierV, _ := value.(string)
			deletefunction.Qualifier = QualifierV
		}
	}

	params := make(map[string]string)
	preparedeleteservicequeryString(params, deletefunction)

	response := make(map[string]interface{})
	err = Preparedeletefnrequest(params, Region, response)

	resp = response
	return resp, err
}
