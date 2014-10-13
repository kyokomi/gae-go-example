package controllers

import (
	"net/http"
	"appengine/taskqueue"
	"net/url"
	"appengine"
	"appengine/datastore"
	"app/models"
	"time"
	"fmt"
	"github.com/gorilla/mux"
)


func TaskAddMemo(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	param := url.Values{}
	param.Set("content", "auto: message")

	for i := 0; i < 1; i++ {
		t := taskqueue.NewPOSTTask("/task/auto-sign", param)
		// 2秒後に実行
		t.Delay = time.Second*2
		if _, err := taskqueue.Add(c, t, "test"); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func TaskBackend(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	content := r.FormValue("content")
	c.Debugf("content = %s", content)

	for i := 0; i < 1; i++ {
		g := models.Greeting{
			Content: content + fmt.Sprintf(" inc = %d", i),
			Date:    time.Now(),
		}

		key := datastore.NewIncompleteKey(c, "Greeting", models.GuestBookKey(c))
		_, err := datastore.Put(c, key, &g)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	fmt.Fprintf(w, "task exec end")
}

func TaskRemove(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	name := mux.Vars(r)["name"]
	c.Debugf("remove task name = %s", name)

	if err := taskqueue.Purge(c, name); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "remove task = %s", name)
}
