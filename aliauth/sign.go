package aliauth

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

const formatISO8601 = "2006-01-02T15:04:05Z"

//LoadBalancerSignAndDoRequest .
func LoadBalancerSignAndDoRequest(action string, params map[string]interface{}, response map[string]interface{}) error {
	// Add common params and action param
	params["Action"] = action
	params["Format"] = "json"
	params["Version"] = "2014-05-15"
	params["AccessKeyId"] = Config.AliAccessKeyID
	params["Timestamp"] = time.Now().UTC().Format(formatISO8601)
	params["SignatureMethod"] = "HMAC-SHA1"
	params["SignatureVersion"] = "1.0"
	params["SignatureNonce"] = createRandomString()

	var endpoint string
	if params["RegionID"] != nil {
		if params["RegionID"].(string) != "" {
			endpoint = getEndpointWithRegion(params["RegionID"].(string))
		}
	}

	if endpoint == "" {
		endpoint = "slb.aliyuncs.com"
	} else {
		endpoint = "slb." + endpoint + ".aliyuncs.com"
	}

	err := signAndDoRequest(endpoint, params, response)
	return err
}

//DNSSignAndDoRequest .
func DNSSignAndDoRequest(action string, params map[string]interface{}, response map[string]interface{}) error {
	// Add common params and action param
	params["Action"] = action
	params["Format"] = "json"
	params["Version"] = "2015-01-09"
	params["AccessKeyId"] = Config.AliAccessKeyID
	params["Timestamp"] = time.Now().UTC().Format(formatISO8601)
	params["SignatureMethod"] = "HMAC-SHA1"
	params["SignatureVersion"] = "1.0"
	params["SignatureNonce"] = createRandomString()

	err := signAndDoRequest("alidns.aliyuncs.com", params, response)
	return err
}

//ECSSignAndDoRequest .
func ECSSignAndDoRequest(action string, params map[string]interface{}, response map[string]interface{}) error {
	// Add common params and action param
	params["Action"] = action
	params["Format"] = "json"
	params["Version"] = "2014-05-26"
	params["AccessKeyId"] = Config.AliAccessKeyID
	params["TimeStamp"] = time.Now().UTC().Format(formatISO8601)
	params["SignatureMethod"] = "HMAC-SHA1"
	params["SignatureVersion"] = "1.0"
	params["SignatureNonce"] = createRandomString()

	var endpoint string
	if params["RegionID"] != nil {
		if params["RegionID"].(string) != "" {
			endpoint = getEndpointWithRegion(params["RegionID"].(string))
		}
	}

	if endpoint == "" {
		endpoint = "ecs.aliyuncs.com"
	} else {
		endpoint = "ecs." + endpoint + ".aliyuncs.com"
	}

	err := signAndDoRequest(endpoint, params, response)
	return err
}

// signAndDoRequest sign and do request by action parameter and specific parameters
func signAndDoRequest(endpoint string, params map[string]interface{}, response map[string]interface{}) error {
	// Sort the parameters by key
	keys := make([]string, len(params))
	i := 0
	for k := range params {
		keys[i] = k
		i++
	}
	sort.Strings(keys)

	// Generate the stringToSign
	var canonicalizedQueryString string
	for _, k := range keys {
		v := params[k]
		canonicalizedQueryString += "&"
		canonicalizedQueryString += percentEncode(k)
		canonicalizedQueryString += "="
		canonicalizedQueryString += percentEncode(getString(v))
	}
	stringToSign := "GET" + "&%2F&" + percentEncode(canonicalizedQueryString[1:])

	// Create signature
	base64Sign := createSignature(stringToSign, Config.AliAccessKeySecret+"&")

	// Init url query
	query := url.Values{}
	for key, value := range params {
		query.Add(key, getString(value))
	}

	// Generate the request URL
	requestURL := "https://" + endpoint + "/?" + query.Encode() + "&Signature=" + url.QueryEscape(base64Sign)

	httpReq, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
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

func percentEncode(str string) string {
	return percentReplace(url.QueryEscape(str))
}

func percentReplace(str string) string {
	str = strings.Replace(str, "+", "%20", -1)
	str = strings.Replace(str, "*", "%2A", -1)
	str = strings.Replace(str, "%7E", "~", -1)

	return str
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
