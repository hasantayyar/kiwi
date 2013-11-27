package main

import (
	"bitbucket.org/kardianos/osext"
	"encoding/json"
	"flag"
	"path"
	"io/ioutil"
)

func ImportConf() error {
	var conffile string
	current, _ := osext.Executable()
	current = path.Dir(current)

	flag.StringVar(&conffile, "config", current+"/"+defaultConfigFile, "Configuration File")
	flag.Parse()

	c, err := ioutil.ReadFile(conffile)
	if err != nil {
		return err
	}

	json.Unmarshal(c, &Config)

	// Sanitize
	if Config.Port == 0 {
		Config.Port = 8042
	}
	if Config.StaticDir == "" {
		Config.StaticDir = current + "/static"
	}
	if Config.TmplDir == "" {
		Config.TmplDir = current + "/templates"
	}
	if Config.XMPP_Server == "" {
		Config.XMPP_Server = "localhost"
	}
	if Config.XMPP_Port == 0 {
		Config.XMPP_Port = 5222
	}

	return nil
}