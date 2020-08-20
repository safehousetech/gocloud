package awsmachinelearning

//CreateMLModel creates model.
func (awsmachinelearning *Awsmachinelearning) CreateMLModel(request interface{}) (resp interface{}, err error) {

	param := request.(map[string]interface{})

	var Region string
	var createMLModel CreateMLModel

	for key, value := range param {
		switch key {
		case "MLModelId":
			MLModelIDV, _ := value.(string)
			createMLModel.MLModelID = MLModelIDV

		case "Region":
			RegionV, _ := value.(string)
			Region = RegionV

		case "MLModelName":
			MLModelNameV, _ := value.(string)
			createMLModel.MLModelName = MLModelNameV

		case "MLModelType":
			MLModelTypeV, _ := value.(string)
			createMLModel.MLModelType = MLModelTypeV

		case "Recipe":
			RecipeV, _ := value.(string)
			createMLModel.Recipe = RecipeV

		case "RecipeUri":
			RecipeURIV, _ := value.(string)
			createMLModel.RecipeURI = RecipeURIV

		case "TrainingDataSourceId":
			TrainingDataSourceIDV, _ := value.(string)
			createMLModel.TrainingDataSourceID = TrainingDataSourceIDV

		case "String":
			StringV, _ := value.(string)
			createMLModel.Parameters.String = StringV
		}
	}

	params := make(map[string]string)

	preparecreateMLModel(params, createMLModel, Region)

	createMLModeljsonmap := make(map[string]interface{})

	preparecreateMLModelparamsdict(createMLModeljsonmap, createMLModel)

	response := make(map[string]interface{})
	err = awsmachinelearning.PrepareSignatureV4query(params, createMLModeljsonmap, response)
	resp = response

	return resp, err
}

//DeleteMLModel delete model.
func (awsmachinelearning *Awsmachinelearning) DeleteMLModel(request interface{}) (resp interface{}, err error) {

	param := request.(map[string]interface{})

	var MLModelID, Region string

	for key, value := range param {
		switch key {
		case "MLModelId":
			MLModelIDV, _ := value.(string)
			MLModelID = MLModelIDV

		case "Region":
			RegionV, _ := value.(string)
			Region = RegionV
		}
	}

	params := make(map[string]string)

	preparedeletemodel(params, MLModelID, Region)

	deletemodeljsonmap := map[string]interface{}{
		"MLModelId": MLModelID,
	}

	response := make(map[string]interface{})
	err = awsmachinelearning.PrepareSignatureV4query(params, deletemodeljsonmap, response)
	resp = response
	return resp, err
}

//GetMLModel describe model.
func (awsmachinelearning *Awsmachinelearning) GetMLModel(request interface{}) (resp interface{}, err error) {

	param := request.(map[string]interface{})

	var MLModelID, Verbose, Region string

	for key, value := range param {
		switch key {
		case "MLModelId":
			MLModelIDV, _ := value.(string)
			MLModelID = MLModelIDV

		case "Verbose":
			VerboseV, _ := value.(string)
			Verbose = VerboseV

		case "Region":
			RegionV, _ := value.(string)
			Region = RegionV
		}
	}

	params := make(map[string]string)

	preparedeletemodel(params, MLModelID, Region)

	deletemodeljsonmap := map[string]interface{}{
		"MLModelId": MLModelID,
		"Verbose":   Verbose,
	}

	response := make(map[string]interface{})
	err = awsmachinelearning.PrepareSignatureV4query(params, deletemodeljsonmap, response)
	resp = response
	return resp, err
}

//UpdateMLModel update model.
func (awsmachinelearning *Awsmachinelearning) UpdateMLModel(request interface{}) (resp interface{}, err error) {

	param := request.(map[string]interface{})

	var MLModelID, MLModelName, Region string
	var ScoreThreshold int
	for key, value := range param {
		switch key {
		case "MLModelId":
			MLModelIDV, _ := value.(string)
			MLModelID = MLModelIDV

		case "MLModelName":
			MLModelNameV, _ := value.(string)
			MLModelName = MLModelNameV

		case "ScoreThreshold":
			ScoreThresholdV, _ := value.(int)
			ScoreThreshold = ScoreThresholdV

		case "Region":
			RegionV, _ := value.(string)
			Region = RegionV
		}
	}

	params := make(map[string]string)

	prepareupdatemodel(params, Region)

	updatemodeljsonmap := make(map[string]interface{})

	prepareupdatemodelparamsdict(updatemodeljsonmap, MLModelID, ScoreThreshold, MLModelName)

	response := make(map[string]interface{})
	err = awsmachinelearning.PrepareSignatureV4query(params, updatemodeljsonmap, response)
	resp = response
	return resp, err
}
