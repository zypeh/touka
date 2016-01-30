package main

import (
	"fmt"
	"github.com/docopt/docopt-go"
	"github.com/zypeh/touka/gist"
)

func main () {
	usage := `Touka

Usage:
  touka push [private] <filename>...
  touka list <username>
  touka -h | --help
  touka --version

Option:
  -h --help    Show this screen.
  --version    Show version
`
	// func Parse (doc string, argv []string, help bool, version string,
	// optionsFirst bool, exit ...bool) (map[string]interface{}, error)
	arguments, _ := docopt.Parse (usage, nil, true, "0.1.0", false)
	fmt.Println (arguments)

	if arguments ["push"] == true {
		if arguments ["private"] == true {
			fmt.Println ("touka private push")
		} else {
			fmt.Println ("touka public push")
		}
	}	

	if arguments ["list"] == true {
		username := arguments ["<username>"]
		Gist.List (username, token)
	}
   
}
