package bigtable

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	googleauth "github.com/safehousetech/gocloud/googleauth"
)

// ListTables ...
func (bigtable *Bigtable) ListTables(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)

	url := "https://bigtableadmin.googleapis.com/v2/" + options["parent"] + "/tables"

	client := googleauth.SignJWT()

	listbigtablerequest, err := http.NewRequest("GET", url, nil)

	listbigtablerequestparam := listbigtablerequest.URL.Query()

	listbigtablerequestparam.Add("pageToken", options["pageToken"])

	if options["view"] != "" {
		listbigtablerequestparam.Add("view", options["view"])
	}

	listbigtablerequest.URL.RawQuery = listbigtablerequestparam.Encode()

	listbigtablerequest.Header.Set("Content-Type", "application/json")

	if err != nil {
		return nil, err
	}

	listbigtableresp, err := client.Do(listbigtablerequest)

	if err != nil {
		return nil, err
	}

	defer listbigtableresp.Body.Close()

	body, err := ioutil.ReadAll(listbigtableresp.Body)

	listbigtableresponse := make(map[string]interface{})
	listbigtableresponse["status"] = listbigtableresp.StatusCode
	listbigtableresponse["body"] = string(body)
	resp = listbigtableresponse
	return resp, err
}

// DeleteTables ...
func (bigtable *Bigtable) DeleteTables(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)

	url := "https://bigtableadmin.googleapis.com/v2/" + options["name"]

	client := googleauth.SignJWT()

	Deletebigtablerequest, err := http.NewRequest("DELETE", url, nil)

	Deletebigtablerequest.Header.Set("Content-Type", "application/json")

	Deletebigtableresp, err := client.Do(Deletebigtablerequest)

	if err != nil {
		return nil, err
	}

	defer Deletebigtableresp.Body.Close()

	body, err := ioutil.ReadAll(Deletebigtableresp.Body)

	Deletebigtableresponse := make(map[string]interface{})
	Deletebigtableresponse["status"] = Deletebigtableresp.StatusCode
	Deletebigtableresponse["body"] = string(body)
	resp = Deletebigtableresponse
	return resp, err
}

// DescribeTables ...
func (bigtable *Bigtable) DescribeTables(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)

	url := "https://bigtableadmin.googleapis.com/v2/" + options["name"]

	client := googleauth.SignJWT()

	Describebigtablerequest, err := http.NewRequest("GET", url, nil)

	Describebigtablerequest.Header.Set("Content-Type", "application/json")

	Describebigtableresp, err := client.Do(Describebigtablerequest)

	if err != nil {
		return nil, err
	}

	defer Describebigtableresp.Body.Close()

	body, err := ioutil.ReadAll(Describebigtableresp.Body)

	Describebigtableresponse := make(map[string]interface{})
	Describebigtableresponse["status"] = Describebigtableresp.StatusCode
	Describebigtableresponse["body"] = string(body)
	resp = Describebigtableresponse
	return resp, err
}

// CreateTables ...
func (bigtable *Bigtable) CreateTables(request interface{}) (resp interface{}, err error) {

	param := request.(map[string]interface{})

	var parent string

	var option Createbigtable

	for key, value := range param {
		switch key {
		case "parent":
			parentv, _ := value.(string)
			parent = parentv

		case "tableId":
			tableIdv, _ := value.(string)
			option.tableID = tableIdv

		case "table":
			tableparam, _ := value.(map[string]interface{})

			for tablekey, tablevalue := range tableparam {

				switch tablekey {

				case "name":
					namev := tablevalue.(string)
					option.table.name = namev

				case "granularity":
					granularityv := tablevalue.(string)
					option.table.granularity = granularityv

				}
			}

		case "initialSplits":
			initialSplitsparam, _ := value.([]map[string]interface{})
			for i := 0; i < len(initialSplitsparam); i++ {
				var initialSplits InitialSplits
				for initialSplitsparamkey, initialSplitsparamvalue := range initialSplitsparam[i] {
					switch initialSplitsparamkey {
					case "key":
						initialSplits.key = initialSplitsparamvalue.(string)
					}
				}
				option.initialSplits = append(option.initialSplits, initialSplits)
			}
		}
	}

	Createbigtablejsonmap := make(map[string]interface{})

	Createbigtabledictnoaryconvert(option, Createbigtablejsonmap)

	Createbigtablejson, _ := json.Marshal(Createbigtablejsonmap)

	Createbigtablejsonstring := string(Createbigtablejson)

	var Createbigtablejsonstringbyte = []byte(Createbigtablejsonstring)

	url := "https://bigtableadmin.googleapis.com/v2/" + parent + "/tables"

	client := googleauth.SignJWT()

	Createbigtablerequest, err := http.NewRequest("POST", url, bytes.NewBuffer(Createbigtablejsonstringbyte))

	Createbigtablerequest.Header.Set("Content-Type", "application/json")

	Createbigtableresp, err := client.Do(Createbigtablerequest)

	if err != nil {
		return nil, err
	}

	defer Createbigtableresp.Body.Close()

	body, err := ioutil.ReadAll(Createbigtableresp.Body)

	Createbigtableresponse := make(map[string]interface{})
	Createbigtableresponse["status"] = Createbigtableresp.StatusCode
	Createbigtableresponse["body"] = string(body)
	resp = Createbigtableresponse
	return resp, err
}
