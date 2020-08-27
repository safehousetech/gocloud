package googleauth

import (
	"log"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googleClient *http.Client

// LoadConfig ...
func LoadConfig() {

	c, err := google.DefaultClient(oauth2.NoContext)

	if err != nil {
		log.Fatal(err)
	}

	googleClient = c
}

// Client ...
func Client() *http.Client {
	return googleClient
}
