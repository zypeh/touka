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
  touka list
  touka -h | --help
  touka --version

Option:
  -h --help    Show this screen.
  --version    Show version
`
	arguments, _ := docopt.Parse (usage, nil, true, "touka", false)
	fmt.Println (arguments)
	
}
