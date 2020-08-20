package awsauth

import (
	"encoding/hex"
	"net/http"
	"strings"
	"time"

	auth "github.com/safehousetech/gocloud/auth"
)

//Getrequestsign4 .
func Getrequestsign4(request *http.Request, region string, service string) *http.Request {

	var algorithm string
	var credentialScope string
	var signedHeaders string
	var date string

	AccessKeyID := auth.Config.AWSAccessKeyID
	SecretAccessKey := auth.Config.AWSSecretKey

	host := service + "." + region + ".amazonaws.com"

	t := time.Now().UTC()

	xamzdate := t.Format("20060102T150405Z")

	dateStamp := t.Format("20060102")

	defaultscontenttype := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded; charset=utf-8",
		"X-Amz-Date":   xamzdate,
	}

	for header, value := range defaultscontenttype {
		if request.Header.Get(header) == "" {
			request.Header.Set(header, value)
		}
	}

	if request.URL.Path == "" {
		request.URL.Path += "/"
	}

	// TASK 1. http://docs.aws.amazon.com/general/latest/gr/sigv4-create-canonical-request.html
	payload := preparepayload(request)
	payloadHash := sha256Hasher(payload)

	request.Header.Set("X-Amz-Content-Sha256", payloadHash)
	request.Header.Set("Host", request.Host)

	signedHeaders = "content-type;host;x-amz-content-sha256;x-amz-date"
	ContentType := "application/x-www-form-urlencoded; charset=utf-8"
	xAmzDate := request.Header.Get("X-Amz-Date")

	canonicalHeaders := "content-type:" + ContentType + "\n" + "host:" + host + "\n" + "x-amz-content-sha256:" + payloadHash + "\n" + "x-amz-date:" + xAmzDate + "\n"

	queryString := request.URL.Query().Encode()
	canonicalQueryString := strings.Replace(queryString, "+", "%20", -1)

	canonicalURI := "/"

	canonicalRequest := request.Method + "\n" + canonicalURI + "\n" + canonicalQueryString + "\n" + canonicalHeaders + "\n" + signedHeaders + "\n" + payloadHash
	hashedCanonReq := sha256Hasher([]byte(canonicalRequest))

	// TASK 2. http://docs.aws.amazon.com/general/latest/gr/sigv4-create-string-to-sign.html

	requestTs := request.Header.Get("X-Amz-Date")
	algorithm = "AWS4-HMAC-SHA256"
	date = dateStamp

	credentialScope = date + "/" + region + "/" + service + "/" + "aws4_request"
	stringToSign := algorithm + "\n" + requestTs + "\n" + credentialScope + "\n" + hashedCanonReq

	// Task 3

	kDate := hmacSHA256([]byte("AWS4"+SecretAccessKey), date)
	kRegion := hmacSHA256(kDate, region)
	kService := hmacSHA256(kRegion, service)
	signingKey := hmacSHA256(kService, "aws4_request")

	signature := hex.EncodeToString(hmacSHA256(signingKey, stringToSign))
	credential := AccessKeyID + "/" + credentialScope

	header := algorithm + " Credential=" + credential + ", SignedHeaders=" + signedHeaders + ", Signature=" + signature
	request.Header.Set("Authorization", header)
	return request
}
