package app

import (
	"net/http"

	"app/controllers"

	"github.com/gorilla/mux"

	"appengine/xmpp"
)

func init() {
	r := mux.NewRouter()

	r.HandleFunc("/", controllers.Index).Methods("GET")
	r.HandleFunc("/guest-book", controllers.Index).Methods("GET")
	r.HandleFunc("/guest-book/sign", controllers.GuestSign).Methods("POST")

	r.HandleFunc("/admin/show-runtime", controllers.ShowRuntime).Methods("GET")
	r.HandleFunc("/admin/counter", controllers.Counter).Methods("GET")

	r.HandleFunc("/user/login", controllers.LoginHandler).Methods("GET")
	r.HandleFunc("/user/logout", controllers.LogoutHandler).Methods("GET")

	r.HandleFunc("/memo", controllers.MemoAdd).Methods("POST")
	r.HandleFunc("/memo/{key}", controllers.MemoShow).Methods("GET")
	r.HandleFunc("/memo/{key}", controllers.MemoRemove).Methods("DELETE")

	r.HandleFunc("/mail-send/{email}", controllers.SendMail).Methods("GET")
	r.HandleFunc("/_ah/mail/{email}", controllers.ReceiveMail).Methods("POST")

	r.HandleFunc("/task/sign", controllers.TaskAddMemo).Methods("GET")
	r.HandleFunc("/task/sign/{name}", controllers.TaskRemove).Methods("DELETE")
	r.HandleFunc("/task/auto-sign", controllers.TaskBackend).Methods("POST")
	r.HandleFunc("/task/delay", controllers.DelayFunc).Methods("GET")

	r.HandleFunc("/xmpp/send", controllers.SendXMPP).Methods("GET")

	http.Handle("/", r)

	xmpp.Handle(controllers.ReceiveXMPP)
}
