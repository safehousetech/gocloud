package lambda

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	auth "github.com/safehousetech/gocloud/auth"
)

func hmacsignatureV4(signingKey []byte, stringToSign string) string {
	return hex.EncodeToString(hmacSHA256(signingKey, stringToSign))
}

func hmacSHA256(key []byte, content string) []byte {
	mac := hmac.New(sha256.New, key)
	mac.Write([]byte(content))
	return mac.Sum(nil)
}

func sha256Hasher(payload []byte) string {
	return fmt.Sprintf("%x", sha256.Sum256(payload))
}

func preparepayload(request *http.Request) []byte {

	if request.Body == nil {
		return []byte{}
	}
	payload, _ := ioutil.ReadAll(request.Body)
	request.Body = ioutil.NopCloser(bytes.NewReader(payload))
	return payload
}

//Preparedeletefnrequest ..
func Preparedeletefnrequest(params map[string]string, region string, response map[string]interface{}) (err error) {
	service := "lambda"
	method := "DELETE"
	host := "lambda" + "." + region + ".amazonaws.com"
	signedheaders := "host;x-amz-date"

	endpoint := "https://lambda.us-east-1.amazonaws.com"

	AccessKeyID := auth.Config.AWSAccessKeyID
	SecretAccessKey := auth.Config.AWSSecretKey

	t := time.Now().UTC()

	XAmzDate := t.Format("20060102T150405Z")
	dateStamp := t.Format("20060102")

	canonicalURI := "/2015-03-31/functions/" + params["FunctionName"]

	request, _ := http.NewRequest("DELETE	", endpoint+canonicalURI, nil)

	payload := preparepayload(request)
	payloadHash := sha256Hasher(payload)

	queryString := ""

	// Go encodes a space as '+' but Amazon requires '%20'. Luckily any '+' in the
	// original query string has been percent escaped so all '+' chars that are left
	// were originally spaces.
	canonicalQuerystring := strings.Replace(queryString, "+", "%20", -1)
	canonicalHeaders := "host:" + host + "\n" + "x-amz-date:" + XAmzDate + "\n"
	canonicalRequest := method + "\n" + canonicalURI + "\n" + canonicalQuerystring + "\n" + canonicalHeaders + "\n" + signedheaders + "\n" + payloadHash

	algorithm := "AWS4-HMAC-SHA256"
	credentialScope := dateStamp + "/" + region + "/" + service + "/" + "aws4_request"
	strToSign := algorithm + "\n" + XAmzDate + "\n" + credentialScope + "\n" + sha256Hasher([]byte(canonicalRequest))

	kDate := hmacSHA256([]byte("AWS4"+SecretAccessKey), dateStamp)
	kRegion := hmacSHA256(kDate, region)
	kService := hmacSHA256(kRegion, service)
	kSigning := hmacSHA256(kService, "aws4_request")

	signature := hmacsignatureV4(kSigning, strToSign)
	authorizationHeader := algorithm + " " + "Credential=" + AccessKeyID + "/" + credentialScope + ", " + "SignedHeaders=" + signedheaders + ", " + "Signature=" + signature

	request.Header.Add("X-Amz-Date", XAmzDate)
	request.Header.Add("Authorization", authorizationHeader)

	client := new(http.Client)

	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

	response["body"] = string(body)
	response["status"] = resp.StatusCode

	return
}

//Preparegetfnrequest ...
func Preparegetfnrequest(params map[string]string, region string, response map[string]interface{}) (err error) {
	service := "lambda"
	method := "GET"
	host := "lambda" + "." + region + ".amazonaws.com"
	signedheaders := "host;x-amz-date"

	endpoint := "https://lambda.us-east-1.amazonaws.com"

	AccessKeyID := auth.Config.AWSAccessKeyID
	SecretAccessKey := auth.Config.AWSSecretKey

	t := time.Now().UTC()

	XAmzDate := t.Format("20060102T150405Z")
	dateStamp := t.Format("20060102")

	canonicalURI := "/2015-03-31/functions/" + params["FunctionName"]

	request, _ := http.NewRequest("GET", endpoint+canonicalURI, nil)

	payload := preparepayload(request)
	payloadHash := sha256Hasher(payload)

	queryString := ""

	// Go encodes a space as '+' but Amazon requires '%20'. Luckily any '+' in the
	// original query string has been percent escaped so all '+' chars that are left
	// were originally spaces.
	canonicalQuerystring := strings.Replace(queryString, "+", "%20", -1)
	canonicalHeaders := "host:" + host + "\n" + "x-amz-date:" + XAmzDate + "\n"
	canonicalRequest := method + "\n" + canonicalURI + "\n" + canonicalQuerystring + "\n" + canonicalHeaders + "\n" + signedheaders + "\n" + payloadHash

	algorithm := "AWS4-HMAC-SHA256"
	credentialScope := dateStamp + "/" + region + "/" + service + "/" + "aws4_request"
	strToSign := algorithm + "\n" + XAmzDate + "\n" + credentialScope + "\n" + sha256Hasher([]byte(canonicalRequest))

	kDate := hmacSHA256([]byte("AWS4"+SecretAccessKey), dateStamp)
	kRegion := hmacSHA256(kDate, region)
	kService := hmacSHA256(kRegion, service)
	kSigning := hmacSHA256(kService, "aws4_request")

	signature := hmacsignatureV4(kSigning, strToSign)
	authorizationHeader := algorithm + " " + "Credential=" + AccessKeyID + "/" + credentialScope + ", " + "SignedHeaders=" + signedheaders + ", " + "Signature=" + signature

	request.Header.Add("X-Amz-Date", XAmzDate)
	request.Header.Add("Authorization", authorizationHeader)

	client := new(http.Client)

	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

	response["body"] = string(body)
	response["status"] = resp.StatusCode

	return
}

//Preparegetrequest ...
func Preparegetrequest(params map[string]string, region string, response map[string]interface{}) (err error) {
	service := "lambda"
	method := "GET"
	host := "lambda" + "." + region + ".amazonaws.com"
	signedheaders := "host;x-amz-date"

	endpoint := "https://lambda.us-east-1.amazonaws.com"

	AccessKeyID := auth.Config.AWSAccessKeyID
	SecretAccessKey := auth.Config.AWSSecretKey

	t := time.Now().UTC()

	XAmzDate := t.Format("20060102T150405Z")
	dateStamp := t.Format("20060102")

	canonicalURI := "/2015-03-31/functions/"

	request, _ := http.NewRequest("GET", endpoint+canonicalURI, nil)

	payload := preparepayload(request)
	payloadHash := sha256Hasher(payload)

	requestparam := request.URL.Query()

	for key, value := range params {
		requestparam.Add(key, value)
	}

	request.URL.RawQuery = requestparam.Encode()

	queryString := request.URL.RawQuery

	// Go encodes a space as '+' but Amazon requires '%20'. Luckily any '+' in the
	// original query string has been percent escaped so all '+' chars that are left
	// were originally spaces.
	canonicalQuerystring := strings.Replace(queryString, "+", "%20", -1)
	canonicalHeaders := "host:" + host + "\n" + "x-amz-date:" + XAmzDate + "\n"
	canonicalRequest := method + "\n" + canonicalURI + "\n" + canonicalQuerystring + "\n" + canonicalHeaders + "\n" + signedheaders + "\n" + payloadHash

	algorithm := "AWS4-HMAC-SHA256"
	credentialScope := dateStamp + "/" + region + "/" + service + "/" + "aws4_request"
	strToSign := algorithm + "\n" + XAmzDate + "\n" + credentialScope + "\n" + sha256Hasher([]byte(canonicalRequest))

	kDate := hmacSHA256([]byte("AWS4"+SecretAccessKey), dateStamp)
	kRegion := hmacSHA256(kDate, region)
	kService := hmacSHA256(kRegion, service)
	kSigning := hmacSHA256(kService, "aws4_request")

	signature := hmacsignatureV4(kSigning, strToSign)
	authorizationHeader := algorithm + " " + "Credential=" + AccessKeyID + "/" + credentialScope + ", " + "SignedHeaders=" + signedheaders + ", " + "Signature=" + signature

	request.Header.Add("X-Amz-Date", XAmzDate)
	request.Header.Add("Authorization", authorizationHeader)

	client := new(http.Client)

	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

	response["body"] = string(body)
	response["status"] = resp.StatusCode

	return
}

//PreparePostrequest ...
func PreparePostrequest(params map[string]interface{}, region string, response map[string]interface{}) (err error) {

	service := "lambda"
	method := "POST"
	host := "lambda" + "." + region + ".amazonaws.com"

	ContentType := "application/x-amz-json-1.0"
	signedheaders := "content-type;host;x-amz-date"

	endpoint := "https://lambda.us-east-1.amazonaws.com"

	AccessKeyID := auth.Config.AWSAccessKeyID
	SecretAccessKey := auth.Config.AWSSecretKey

	t := time.Now().UTC()

	XAmzDate := t.Format("20060102T150405Z")
	dateStamp := t.Format("20060102")

	canonicalURI := "/2015-03-31/functions/"

	requestparametersjson, _ := json.Marshal(params)
	requestparametersjsonstring := string(requestparametersjson)
	requestparametersjsonstringbyte := []byte(requestparametersjsonstring)

	request, _ := http.NewRequest("POST", endpoint+canonicalURI, bytes.NewBuffer(requestparametersjsonstringbyte))

	payload := preparepayload(request)
	payloadHash := sha256Hasher(payload)

	// Go encodes a space as '+' but Amazon requires '%20'. Luckily any '+' in the
	// original query string has been percent escaped so all '+' chars that are left
	// were originally spaces.
	canonicalQuerystring := ""

	canonicalHeaders := "content-type:" + ContentType + "\n" + "host:" + host + "\n" + "x-amz-date:" + XAmzDate + "\n"

	//payload_hashstring := sha256Hasher(request_parameters)

	canonicalRequest := method + "\n" + canonicalURI + "\n" + canonicalQuerystring + "\n" + canonicalHeaders + "\n" + signedheaders + "\n" + payloadHash

	algorithm := "AWS4-HMAC-SHA256"
	credentialScope := dateStamp + "/" + region + "/" + service + "/" + "aws4_request"
	strToSign := algorithm + "\n" + XAmzDate + "\n" + credentialScope + "\n" + sha256Hasher([]byte(canonicalRequest))

	kDate := hmacSHA256([]byte("AWS4"+SecretAccessKey), dateStamp)
	kRegion := hmacSHA256(kDate, region)
	kService := hmacSHA256(kRegion, service)
	kSigning := hmacSHA256(kService, "aws4_request")

	signature := hmacsignatureV4(kSigning, strToSign)
	authorizationHeader := algorithm + " " + "Credential=" + AccessKeyID + "/" + credentialScope + ", " + "SignedHeaders=" + signedheaders + ", " + "Signature=" + signature

	request.Header.Set("Content-Type", ContentType)
	request.Header.Set("Host", "lambda.us-east-1.amazonaws.com")
	request.Header.Add("X-Amz-Date", XAmzDate)
	request.Header.Add("Authorization", authorizationHeader)

	client := new(http.Client)

	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

	response["body"] = string(body)
	response["status"] = resp.StatusCode

	return
}
