package awsauth

import (
	"net/http"
	"time"

	auth "github.com/safehousetech/gocloud/auth"
)

//SignatureV4 .
func SignatureV4(request *http.Request, requestParameters []byte, amztarget string, method string, region string, service string, host string, ContentType string, signedheaders string) (signrequest *http.Request) {
	AccessKeyID := auth.Config.AWSAccessKeyID
	SecretAccessKey := auth.Config.AWSSecretKey

	t := time.Now().UTC()
	XAmzDate := t.Format("20060102T150405Z")
	dateStamp := t.Format("20060102")
	canonicalURI := "/"
	canonicalQueryString := ""
	canonicalHeaders := "content-type:" + ContentType + "\n" + "host:" + host + "\n" + "x-amz-date:" + XAmzDate + "\n" + "x-amz-target:" + amztarget + "\n"
	payloadHashString := sha256Hasher(requestParameters)
	canonicalRequest := method + "\n" + canonicalURI + "\n" + canonicalQueryString + "\n" + canonicalHeaders + "\n" + signedheaders + "\n" + payloadHashString

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
	request.Header.Add("X-Amz-Date", XAmzDate)
	request.Header.Add("X-Amz-Target", amztarget)
	request.Header.Add("Authorization", authorizationHeader)

	return request
}
