package controllers

import (
	"fmt"
	"net/http"

	"appengine"
	"appengine/xmpp"
)

func SendXMPP(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	m := &xmpp.Message{
		To:   []string{"organic-victory-708@appspot.com"},
		Body: `Hi! How's the carrot?`,
	}
	if err := m.Send(c); err != nil {
		c.Errorf(err.Error())
	}

	fmt.Fprintf(w, "ok")
}

func ReceiveXMPP(c appengine.Context, m *xmpp.Message) {
	c.Debugf("receive!!!!!!")
	c.Debugf("%v", m)
}
