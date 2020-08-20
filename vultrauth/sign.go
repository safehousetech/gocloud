package vultrauth

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// SignAndDoRequest sign and do request by action parameter and specific parameters
func SignAndDoRequest(method string, path string, params map[string]interface{}, response map[string]interface{}) error {

	requestURL := "https://api.vultr.com" + path

	// Init url query
	query := url.Values{}
	for key, value := range params {
		query.Add(key, getString(value))
	}

	httpReq, err := http.NewRequest(method, requestURL, strings.NewReader(query.Encode()))
	if err != nil {
		return err
	}

	httpReq.Header.Set("API-Key", Config.VultrAPIKey)
	httpReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	response["body"] = string(body)
	response["status"] = resp.StatusCode
	return err
}

// getString return string from interface{}
func getString(v interface{}) string {
	switch v.(type) {
	case string:
		return v.(string)
	case int:
		return strconv.Itoa(v.(int))
	case bool:
		return strconv.FormatBool(v.(bool))
	}
	return ""
}
