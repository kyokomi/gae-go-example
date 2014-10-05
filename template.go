package gaehoge

import "html/template"

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
