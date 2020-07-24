package googlemachinelearning

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	googleauth "github.com/cloudlibz/gocloud/googleauth"
)

const (
	//UnixDate ...
	UnixDate = "Mon Jan _2 15:04:05 MST 2006"
	//RFC3339 ...
	RFC3339 = "2006-01-02T15:04:05Z07:00"
)

//CreateMLModel creates model.
func (googlemachinelearning *Googlemachinelearning) CreateMLModel(request interface{}) (resp interface{}, err error) {

	param := request.(map[string]interface{})
	var Parent string
	var option CreateMLModel

	for key, value := range param {
		switch key {

		case "Parent":
			ParentV, _ := value.(string)
			Parent = ParentV

		case "Name":
			NameV, _ := value.(string)
			option.name = NameV

		case "Description":
			DescriptionV, _ := value.(string)
			option.description = DescriptionV

		case "OnlinePredictionLogging":
			OnlinePredictionLoggingV, _ := value.(bool)
			option.onlinePredictionLogging = OnlinePredictionLoggingV

		case "Regions":
			RegionsV, _ := value.([]string)
			option.regions = RegionsV

		case "DefaultVersion":
			defaultVersionparam, _ := value.(map[string]interface{})

			for defaultVersionparamkey, defaultVersionparamvalue := range defaultVersionparam {

				switch defaultVersionparamkey {
				case "Name":
					namev, _ := defaultVersionparamvalue.(string)
					option.defaultVersion.name = namev

				case "Description":
					descriptionv, _ := defaultVersionparamvalue.(string)
					option.defaultVersion.description = descriptionv

				case "IsDefault":
					isDefaultv, _ := defaultVersionparamvalue.(bool)
					option.defaultVersion.isDefault = isDefaultv

				case "DeploymentUri":
					deploymentUriv, _ := defaultVersionparamvalue.(string)
					option.defaultVersion.deploymentURI = deploymentUriv

				case "CreateTime":
					createTimev := time.Now().UTC().Format(time.RFC3339)
					option.defaultVersion.createTime = createTimev

				case "LastUseTime":
					lastUseTimev, _ := defaultVersionparamvalue.(string)
					option.defaultVersion.lastUseTime = lastUseTimev

				case "RuntimeVersion":
					runtimeVersionv, _ := defaultVersionparamvalue.(string)
					option.defaultVersion.runtimeVersion = runtimeVersionv

				case "State":
					statev, _ := defaultVersionparamvalue.(string)
					option.defaultVersion.state = statev

				case "ErrorMessage":
					errorMessagev, _ := defaultVersionparamvalue.(string)
					option.defaultVersion.errorMessage = errorMessagev

				case "Framework":
					frameworkv, _ := defaultVersionparamvalue.(string)
					option.defaultVersion.framework = frameworkv

				case "PythonVersion":
					pythonVersionv, _ := defaultVersionparamvalue.(string)
					option.defaultVersion.pythonVersion = pythonVersionv

				case "AutoScaling":
					autoScalingparam, _ := defaultVersionparamvalue.(map[string]interface{})
					for autoScalingparamkey, autoScalingparamvalue := range autoScalingparam {
						switch autoScalingparamkey {
						case "MinNodes":
							minNodesv, _ := autoScalingparamvalue.(string)
							option.defaultVersion.autoScaling.minNodes = minNodesv
						}
					}

				case "ManualScaling":
					manualScalingparam, _ := defaultVersionparamvalue.(map[string]interface{})
					for manualScalingparamkey, manualScalingparamvalue := range manualScalingparam {
						switch manualScalingparamkey {
						case "Nodes":
							nodesv, _ := manualScalingparamvalue.(string)
							option.defaultVersion.manualScaling.nodes = nodesv
						}
					}
				}
			}
		}
	}

	createMLModeljsonmap := make(map[string]interface{})

	createMLModeldictnoaryconvert(option, createMLModeljsonmap)

	createMLModeljson, _ := json.Marshal(createMLModeljsonmap)

	createMLModeljsonstring := string(createMLModeljson)

	var createMLModeljsonstringbyte = []byte(createMLModeljsonstring)

	url := "https://ml.googleapis.com/v1/" + Parent + "/models"

	client := googleauth.SignJWT()

	createMLModelrequest, err := http.NewRequest("POST", url, bytes.NewBuffer(createMLModeljsonstringbyte))

	createMLModelrequest.Header.Set("Content-Type", "application/json")

	createMLModelresp, err := client.Do(createMLModelrequest)

	if err != nil {

	}

	defer createMLModelresp.Body.Close()

	body, err := ioutil.ReadAll(createMLModelresp.Body)

	createMLModelresponse := make(map[string]interface{})
	createMLModelresponse["status"] = createMLModelresp.StatusCode
	createMLModelresponse["body"] = string(body)
	resp = createMLModelresponse
	return resp, err
}

//DeleteMLModel delete model.
func (googlemachinelearning *Googlemachinelearning) DeleteMLModel(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)

	url := "https://ml.googleapis.com/v1/" + options["name"]

	client := googleauth.SignJWT()

	DeleteMLModelrequest, err := http.NewRequest("DELETE", url, nil)

	DeleteMLModelrequest.Header.Set("Content-Type", "application/json")

	DeleteMLModelresp, err := client.Do(DeleteMLModelrequest)

	if err != nil {

	}

	defer DeleteMLModelresp.Body.Close()

	body, err := ioutil.ReadAll(DeleteMLModelresp.Body)

	DeleteMLModelresponse := make(map[string]interface{})
	DeleteMLModelresponse["status"] = DeleteMLModelresp.StatusCode
	DeleteMLModelresponse["body"] = string(body)
	resp = DeleteMLModelresponse

	return resp, err

}

//GetMLModel get model.
func (googlemachinelearning *Googlemachinelearning) GetMLModel(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)

	url := "https://ml.googleapis.com/v1/" + options["name"]

	client := googleauth.SignJWT()

	GetMLModelrequest, err := http.NewRequest("GET", url, nil)

	GetMLModelrequest.Header.Set("Content-Type", "application/json")

	GetMLModelresp, err := client.Do(GetMLModelrequest)

	if err != nil {

	}

	defer GetMLModelresp.Body.Close()

	body, err := ioutil.ReadAll(GetMLModelresp.Body)

	GetMLModelresponse := make(map[string]interface{})
	GetMLModelresponse["status"] = GetMLModelresp.StatusCode
	GetMLModelresponse["body"] = string(body)
	resp = GetMLModelresponse

	return resp, err
}

//UpdateMLModel  Update model.
func (googlemachinelearning *Googlemachinelearning) UpdateMLModel(request interface{}) (resp interface{}, err error) {

	param := request.(map[string]interface{})
	var Parent, UpdateMask string
	var option CreateMLModel

	for key, value := range param {
		switch key {

		case "Parent":
			ParentV, _ := value.(string)
			Parent = ParentV

		case "UpdateMask":
			updateMaskV, _ := value.(string)
			UpdateMask = updateMaskV

		case "Name":
			NameV, _ := value.(string)
			option.name = NameV

		case "Description":
			DescriptionV, _ := value.(string)
			option.description = DescriptionV

		case "OnlinePredictionLogging":
			OnlinePredictionLoggingV, _ := value.(bool)
			option.onlinePredictionLogging = OnlinePredictionLoggingV

		case "Regions":
			RegionsV, _ := value.([]string)
			option.regions = RegionsV

		case "DefaultVersion":
			defaultVersionparam, _ := value.(map[string]interface{})

			for defaultVersionparamkey, defaultVersionparamvalue := range defaultVersionparam {

				switch defaultVersionparamkey {
				case "Name":
					namev, _ := defaultVersionparamvalue.(string)
					option.defaultVersion.name = namev

				case "Description":
					descriptionv, _ := defaultVersionparamvalue.(string)
					option.defaultVersion.description = descriptionv

				case "IsDefault":
					isDefaultv, _ := defaultVersionparamvalue.(bool)
					option.defaultVersion.isDefault = isDefaultv

				case "DeploymentUri":
					deploymentUriv, _ := defaultVersionparamvalue.(string)
					option.defaultVersion.deploymentURI = deploymentUriv

				case "CreateTime":
					createTimev := time.Now().UTC().Format(time.RFC3339)
					option.defaultVersion.createTime = createTimev

				case "LastUseTime":
					lastUseTimev, _ := defaultVersionparamvalue.(string)
					option.defaultVersion.lastUseTime = lastUseTimev

				case "RuntimeVersion":
					runtimeVersionv, _ := defaultVersionparamvalue.(string)
					option.defaultVersion.runtimeVersion = runtimeVersionv

				case "State":
					statev, _ := defaultVersionparamvalue.(string)
					option.defaultVersion.state = statev

				case "ErrorMessage":
					errorMessagev, _ := defaultVersionparamvalue.(string)
					option.defaultVersion.errorMessage = errorMessagev

				case "Framework":
					frameworkv, _ := defaultVersionparamvalue.(string)
					option.defaultVersion.framework = frameworkv

				case "PythonVersion":
					pythonVersionv, _ := defaultVersionparamvalue.(string)
					option.defaultVersion.pythonVersion = pythonVersionv

				case "AutoScaling":
					autoScalingparam, _ := defaultVersionparamvalue.(map[string]interface{})
					for autoScalingparamkey, autoScalingparamvalue := range autoScalingparam {
						switch autoScalingparamkey {
						case "MinNodes":
							minNodesv, _ := autoScalingparamvalue.(string)
							option.defaultVersion.autoScaling.minNodes = minNodesv
						}
					}

				case "ManualScaling":
					manualScalingparam, _ := defaultVersionparamvalue.(map[string]interface{})
					for manualScalingparamkey, manualScalingparamvalue := range manualScalingparam {
						switch manualScalingparamkey {
						case "Nodes":
							nodesv, _ := manualScalingparamvalue.(string)
							option.defaultVersion.manualScaling.nodes = nodesv
						}
					}
				}
			}
		}
	}

	updateMLModeljsonmap := make(map[string]interface{})

	createMLModeldictnoaryconvert(option, updateMLModeljsonmap)

	updateMLModeljson, _ := json.Marshal(updateMLModeljsonmap)

	updateMLModeljsonstring := string(updateMLModeljson)

	var updateMLModeljsonstringbyte = []byte(updateMLModeljsonstring)

	url := "https://ml.googleapis.com/v1/" + Parent

	client := googleauth.SignJWT()

	updateMLModelrequest, err := http.NewRequest("PATCH", url, bytes.NewBuffer(updateMLModeljsonstringbyte))

	updateMLModelparam := updateMLModelrequest.URL.Query()

	if UpdateMask != "" {
		updateMLModelparam.Add("updateMask", UpdateMask)
	}

	updateMLModelrequest.URL.RawQuery = updateMLModelparam.Encode()

	updateMLModelrequest.Header.Set("Content-Type", "application/json")

	updateMLModelresp, err := client.Do(updateMLModelrequest)

	if err != nil {

	}

	defer updateMLModelresp.Body.Close()

	body, err := ioutil.ReadAll(updateMLModelresp.Body)

	updateMLModelresponse := make(map[string]interface{})
	updateMLModelresponse["status"] = updateMLModelresp.StatusCode
	updateMLModelresponse["body"] = string(body)
	resp = updateMLModelresponse
	return resp, err
}
