// This code is part of the touka
package Gist

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/fatih/color"
)

// GIT_IO_URL is the github's URL shortner
// API v3 is the current version of Github API
const (
	GITHUB_API_URL = "https://api.github.com/"
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

/* Color initialization */
var (
	yellow := color.New (color.FgYellow).SprintfFunc ()
	yellow_printf := color.New (color.FgYellow).PrintfFunc ()

	green := color.New (color.FgGreen).SprintfFunc ()
	green_printf := color.New (color.FgGreen).PrintfFunc ()
)

type GistFile struct {
	Content string `json:\"content\"`
}

type Gist struct {
	Description string                  `json:\"description\"`
	publicFile  bool                    `json:\"public\"`
	GistFile    map [string] GistFile   `json:\"files\"`
}

/* Load Token From ~/.gist */
func loadTokenFromFIle () (token string) {
	file := os.Getenv ("$HOME") + "/.gist"
	github_token, err := ioutil.ReadFile (file)
	if err != nil {
		log.Fatal (err)
	}
	return string (github_token)
}

/* Error handling */
func errhandler (handler err) {
	if handler != nil {
		log.Fatal (handler)
		os.Exit (1)
	}
}

/* Push the file on gist.github.com */
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

func List (url, token string) {

	if token != "" {
		url = url + "?access_token=" + token
	}

	res, err := http.Get (url)
	errhandler (err)
	
	defer res.Body.Close ()
	body, err := ioutil.ReadAll (res.Body)
	errhandler (err)

	var jsonRes []GistJSON.Response
	err= json.Unmarshal (body, &jsonRes)
	errhandler (err)

	for _, val := range jsonRes {
	    yellow_printf ("%sn", val.HtmlUrl)
		fmt.Printf ("(%s)\t%s\n", shortDate (val.CreatedAt), val.Description)
		fmt.Printf ("\n")
	}
	
}	
