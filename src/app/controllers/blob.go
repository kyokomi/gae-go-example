package controllers

import (
	"net/http"
	"appengine"
	"appengine/blobstore"
	"html/template"
)

var htmlTemplate = template.Must(template.New("html").Parse(htmlText))

const htmlText = `
<!DOCTYPE html>
<html><head><title>BlobStore API</title></head>
<body><form action="{{.}}" method="POST" enctype="multipart/form-data">
<lable>更新するファイル: <input type="file" name="file"></label>
<input type="submit" value="送信">
</form></body></html>
`

func BlobPage(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	uploadURL, err := blobstore.UploadURL(c, "/blob/upload", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := htmlTemplate.Execute(w, uploadURL); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func BlobUpload(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	blobs, _, err := blobstore.ParseUpload(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	file := blobs["file"]
	if len(file) == 0 {
		c.Errorf("no file upload")
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	http.Redirect(w, r, "/blob/store?blobkey="+string(file[0].BlobKey), http.StatusFound)
}

func BlobStore(w http.ResponseWriter, r *http.Request) {
	blobstore.Send(w, appengine.BlobKey(r.FormValue("blobkey")))
}
