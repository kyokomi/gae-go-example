package controllers

import (
	"net/http"
	"appengine"
	"appengine/urlfetch"
	"bytes"
	"io"
	"fmt"
	"github.com/gorilla/mux"
)

func GitHubFetch(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	client := urlfetch.Client(c)

	userName := mux.Vars(r)["user"]

	res, err := client.Get("https://github.com/" + userName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if res.StatusCode != http.StatusOK {
		http.Error(w, res.Status, res.StatusCode)
		return
	}

	var body = make([]byte, res.ContentLength)
	buffer := bytes.NewBuffer(body)

	defer res.Body.Close()

	for {
		var buf = make([]byte, res.ContentLength)
		_, err := res.Body.Read(buf)
		if err != nil && err != io.EOF {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} else if err == io.EOF {
			break
		}

		buffer.Write(buf)
	}

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "%s", buffer.String())
}
