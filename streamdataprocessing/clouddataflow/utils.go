package clouddataflow

func createstreamdictnoaryconvert(option CreateStream, createstreamjsonmap map[string]interface{}) {

	if option.ID != "" {
		createstreamjsonmap["id"] = option.ID
	}

	if option.ProjectID != "" {
		createstreamjsonmap["projectId"] = option.ProjectID
	}

	if option.Name != "" {
		createstreamjsonmap["name"] = option.Name
	}

	if option.Type != "" {
		createstreamjsonmap["type"] = option.Type
	}

	if option.CurrentState != "" {
		createstreamjsonmap["currentState"] = option.CurrentState
	}

	if option.CurrentStateTime != "" {
		createstreamjsonmap["currentStateTime"] = option.CurrentStateTime
	}

	if option.RequestedState != "" {
		createstreamjsonmap["requestedState"] = option.RequestedState
	}

	if option.CreateTime != "" {
		createstreamjsonmap["createTime"] = option.CreateTime
	}

	if option.ReplaceJobID != "" {
		createstreamjsonmap["replaceJobId"] = option.ReplaceJobID
	}

	if option.Location != "" {
		createstreamjsonmap["location"] = option.Location
	}

	prepareStageStates(option, createstreamjsonmap)
	prepareEnvironment(option, createstreamjsonmap)
}

func prepareEnvironment(option CreateStream, createstreamjsonmap map[string]interface{}) {

	environmentv := make(map[string]interface{})

	versionv := make(map[string]interface{})

	if option.Environment.Version.Major != "" {
		versionv["major"] = option.Environment.Version.Major
	}

	if option.Environment.Version.JobType != "" {
		versionv["job_type"] = option.Environment.Version.JobType
	}

	environmentv["version"] = versionv

	userAgentv := make(map[string]interface{})

	if option.Environment.UserAgent.Name != "" {
		userAgentv["name"] = option.Environment.UserAgent.Name
	}

	if option.Environment.UserAgent.BuildDate != "" {
		userAgentv["build.date"] = option.Environment.UserAgent.BuildDate
	}

	if option.Environment.UserAgent.Version != "" {
		userAgentv["version"] = option.Environment.UserAgent.Version
	}

	supportv := make(map[string]interface{})

	if option.Environment.UserAgent.Support.Status != "" {
		supportv["status"] = option.Environment.UserAgent.Support.Status
	}

	if option.Environment.UserAgent.Support.URL != "" {
		supportv["url"] = option.Environment.UserAgent.Support.URL
	}

	userAgentv["support"] = supportv

	environmentv["userAgent"] = userAgentv

	createstreamjsonmap["environment"] = environmentv
}

func prepareStageStates(option CreateStream, createstreamjsonmap map[string]interface{}) {

	if len(option.StageStates) > 0 {

		stageStatesarray := make([]map[string]interface{}, 0)

		for i := 0; i < len(option.StageStates); i++ {

			stageState := make(map[string]interface{})
			stageState["currentStateTime"] = option.StageStates[i].CurrentStateTime
			stageState["executionStageName"] = option.StageStates[i].ExecutionStageName
			stageState["executionStageName"] = option.StageStates[i].ExecutionStageName

			stageStatesarray = append(stageStatesarray, stageState)

		}

		createstreamjsonmap["stageStates"] = stageStatesarray
	}
}
