package gaehoge

import (
	"datastores"
	"templates"

	"net/http"
	"time"

	"appengine"
	"appengine/datastore"
	"appengine/user"
)

func init() {
	http.HandleFunc("/", doIndex)

	http.HandleFunc("/guest-book/sign", doSign)
}

func doIndex(w http.ResponseWriter, r *http.Request) {

	c := appengine.NewContext(r)

	q := datastore.NewQuery("Greeting").Ancestor(datastores.GuestBookKey(c)).Order("-Date").Limit(10)
	greetings := make([]datastores.Greeting, 0, 10)
	if _, err := q.GetAll(c, &greetings); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	type exec struct {
		Author    string
		Greetings []datastores.Greeting
	}
	e := exec{
		Greetings: greetings,
	}
	if u := user.Current(c); u != nil {
		e.Author = u.String()
	}

	if err := templates.GuestBookTemplate.Execute(w, e); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func doSign(w http.ResponseWriter, r *http.Request) {

	c := appengine.NewContext(r)

	g := datastores.Greeting{
		Content: r.FormValue("content"),
		Date:    time.Now(),
	}

	if u := user.Current(c); u != nil {
		g.Author = u.String()
	}

	key := datastore.NewIncompleteKey(c, "Greeting", datastores.GuestBookKey(c))
	_, err := datastore.Put(c, key, &g)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
