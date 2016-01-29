// This code is part of the touka

package gist

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// GIT_IO_URL is the github's URL shortner
// API v3 is the current version of Github API
const (
	GITHUB_API_URL = "https://api.github.com/gists"
	GIT_IO_URL = "http://git.io"
	GIT_BASE_PATH = "/api/v3"
)

var (
	USER_AGENT = "gist/#/touka"
	token = os.Getenv ("GITHUB_TOKEN")
)

var (
	publicFlag bool
	description string
	anonymous bool
	responseObj map [string]interface{}
)

type GistFile struct {
	Content string `json:\"content\"`
}

type Gist struct {
	Description string                  `json:\"description\"`
	publicFile  bool                    `json:\"public\"`
	GistFile    map [string] GistFile   `json:\"files\"`
}

func loadTokenFromFIle () (token string) {
	file := os.Getenv ("$HOME") + "/.gist"
	github_token, err := ioutil.ReadFile (file)
	if err != nil {
		log.Fatal (err)
	}
	return string (github_token)
}

func Push (file Gist) {
	pfile, err := json.Marshal (file)

	if err {
		log.Fatal ("[ !! ] Cannot marshal json: ", err)
	}

	b := bytes.NewBuffer (pfile)
	fmt.Println ("[ ** ] Uploading...")

	req, err := http.NewRequest ("POST", GITHUB_API_URL, b)

	if !anonymous {
		req.Header.Set ("Content-Type", "application/json")
		req.Headet.Set ("Accept", "application/json")
		req.SetBasicAuth (token, "x-oauth-basic")
	}

}	
