package main

import (
	"bitbucket.org/kardianos/osext"
	"encoding/json"
	"fmt"
	"flag"
	"path"
	"io/ioutil"
)

func main() {
	var conffile string
	current, _ := osext.Executable()
	current = path.Dir(current)

	flag.StringVar(&conffile, "config", current+"/kiwi.json", "Configuration File")
	flag.Parse()

	c, err := ioutil.ReadFile(conffile)
	if err != nil {
		fmt.Println("Error Reading Config File ", err)
		return 1
	}
}