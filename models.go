package main

type Configuration struct {
	Port		int		`json: "port"`
	StaticDir	string	`json: "static_dir"`
	TmplDir		string	`json: "template_dir"`
	XMPP_Server	string	`json: "xmpp_server"`
	XMPP_Port	int		`json: "xmpp_port"`
}