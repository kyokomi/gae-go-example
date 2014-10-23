package controllers

import (
	"appengine"
	"net/http"
	"appengine/capability"
	"fmt"
)

type ServiceStatus struct {
	Blob bool
	Read bool
	Write bool
	Mail bool
	Cache bool
	Task bool
	URLFetch bool
}

func Capabilities(w http.ResponseWriter, r *http.Request) {
	var stat ServiceStatus

	c := appengine.NewContext(r)

	stat.Blob     = capability.Enabled(c, "blobstore", "")
	stat.Read     = capability.Enabled(c, "datastore_v3", "")
	stat.Write    = capability.Enabled(c, "datastore_v3", "write")
	stat.Mail     = capability.Enabled(c, "mail", "")
	stat.Cache    = capability.Enabled(c, "memcache", "")
	stat.Task     = capability.Enabled(c, "taskqueue", "")
	stat.URLFetch = capability.Enabled(c, "urlfetch", "")

	fmt.Fprintf(w, "%v", stat)
}
