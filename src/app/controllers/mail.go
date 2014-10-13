package controllers

import (
	"net/http"
	"appengine"
	"appengine/mail"
	"github.com/gorilla/mux"
	"bytes"
)

func SendMail(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	vars := mux.Vars(r)
	address := vars["email"]

	m := &mail.Message{
		Sender: "kyokomi-dev<organic-victory-708@appspot.gserviceaccount.com>",
//		ReplyTo: "",
		To: []string{address},
//		Cc: "",
//		Bcc: []string{},
		Subject: "Test Mail",
		Body:    "サンプルメールを送信。",
		HTMLBody: "",
//		Attachments: []Attachment{},
//		Headers: mail.Header{},
	}

	if err := mail.Send(c, m); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func ReceiveMail(_ http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	defer r.Body.Close()

	var b bytes.Buffer
	if _, err := b.ReadFrom(r.Body); err != nil {
		c.Errorf("Error reading body: %v", err)
		return
	}

	c.Infof("Received mail: \n%s", b.String())
}
