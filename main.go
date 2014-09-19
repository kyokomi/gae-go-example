package gaehoge

import (
	"html/template"
	"net/http"
	"time"

	"appengine"
	"appengine/datastore"
	"appengine/user"
)

type Greeting struct {
	Author  string
	Content string
	Date    time.Time
}

var guestBookTemplate = template.Must(template.New("book").Parse(`
<html>
  <head>
    <title>Go Guestbook</title>
  </head>
  <body>
    {{range .Greetings}}
      {{with .Author}}
        <p><b>{{.}}</b> wrote:</p>
      {{else}}
        <p>An anonymous person wrote:</p>
      {{end}}
      <pre>{{.Content}}</pre>
    {{end}}
    <form action="/sign" method="post">
      <div><textarea name="content" rows="3" cols="60"></textarea></div>
      <div><input type="submit" value="Sign Guestbook"></div>
    </form>

    Author: {{.Author}} <br />
    {{with .Author}}
    <a href="/logout"><button>logout</button></a>
    {{else}}
    <a href="/login"><button>login</button></a>
    {{end}}
  </body>
</html>
`))

func init() {
	http.HandleFunc("/", doIndex)
	http.HandleFunc("/sign", doSign)

	http.HandleFunc("/login", doLoginHandler)
	http.HandleFunc("/logout", doLogoutHandler)
}

func guestBookKey(c appengine.Context) *datastore.Key {
	return datastore.NewKey(c, "Guestbook", "defualt_guestbook", 0, nil)
}

func doIndex(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	index(w, r, c)
}

func index(w http.ResponseWriter, _ *http.Request, c appengine.Context) {
	q := datastore.NewQuery("Greeting").Ancestor(guestBookKey(c)).Order("-Date").Limit(10)
	greetings := make([]Greeting, 0, 10)
	if _, err := q.GetAll(c, &greetings); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	type exec struct {
		Author    string
		Greetings []Greeting
	}
	e := exec{
		Greetings: greetings,
	}
	if u := user.Current(c); u != nil {
		e.Author = u.String()
	}

	if err := guestBookTemplate.Execute(w, e); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func doSign(w http.ResponseWriter, r *http.Request) {

	c := appengine.NewContext(r)
	g := Greeting{
		Content: r.FormValue("content"),
		Date:    time.Now(),
	}

	if u := user.Current(c); u != nil {
		g.Author = u.String()
	}

	key := datastore.NewIncompleteKey(c, "Greeting", guestBookKey(c))
	_, err := datastore.Put(c, key, &g)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func doLoginHandler(w http.ResponseWriter, r *http.Request) {

	c := appengine.NewContext(r)
	u := user.Current(c)
	if u == nil {
		url, err := user.LoginURL(c, r.URL.String())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Location", url)
		w.WriteHeader(http.StatusFound)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func doLogoutHandler(w http.ResponseWriter, r *http.Request) {

	c := appengine.NewContext(r)
	u := user.Current(c)
	if u != nil {
		url, err := user.LogoutURL(c, r.URL.String())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Location", url)
		w.WriteHeader(http.StatusFound)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}
