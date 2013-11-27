package main

import (
	"github.com/mattn/go-xmpp"
	"log"
)

func AuthXMPPUser(username, password string) error {
	talk, err := xmpp.NewClientNoTLS(Config.XMPP_Server, username, password, false)
	if err != nil {
		return err
	}
	go func() {
		for {
			chat, err := talk.Recv()
			if err != nil {
				log.Fatal("recv error", err)
			}
			switch v := chat.(type) {
			case xmpp.Chat: 
				fmt.Println(v.Remote, v.Text)
			case xmpp.Presence:
				fmt.Println(v.From, v.Show)
			}
		}
	}()
	return nil
}