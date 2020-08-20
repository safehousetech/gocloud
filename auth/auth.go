package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// AWSConfiguration struct for representing AWS credentials...
type AWSConfiguration struct {
	AWSAccessKeyID string
	AWSSecretKey   string
}

// Config is variable of type AWSConfiguration.
var Config AWSConfiguration

// LoadConfig loads the AWS credentials...
func LoadConfig() {

	// Read from file first.
	var home string = os.Getenv("HOME")
	file, err := os.Open(home + "/.gocloud" + "/amazoncloudconfig.json")

	if err != nil {
		fmt.Println(err)
		return
	}

	// Defer the closing of our jsonFile so that we can parse it later on.
	defer file.Close()

	// We initialize AWSConfiguration struct.
	decoder := json.NewDecoder(file)
	Config = AWSConfiguration{}
	_ = decoder.Decode(&Config)

	if Config.AWSAccessKeyID == "" || Config.AWSSecretKey == "" {
		// If amazoncloudconfig.json doesn't exist, look for credentials as
		// environment variables.

		Config.AWSAccessKeyID = os.Getenv("AWSAccessKeyID")
		Config.AWSSecretKey = os.Getenv("AWSSecretKey")

		if Config.AWSAccessKeyID == "" || Config.AWSSecretKey == "" {
			log.Fatalln("Cannot get access key and secret key.")
		}
	}
}
