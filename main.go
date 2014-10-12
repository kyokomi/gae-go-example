package gaehoge

import (
	"net/http"

	"github.com/gorilla/mux"
	"app/controllers"
)

func init() {
	r := mux.NewRouter()

	r.HandleFunc("/", controllers.Index)
	r.HandleFunc("/guest-book", controllers.Index)
	r.HandleFunc("/guest-book/sign", controllers.GuestSign)

	r.HandleFunc("/admin/show-runtime", controllers.ShowRuntime)

	r.HandleFunc("/user/login", controllers.LoginHandler)
	r.HandleFunc("/user/logout", controllers.LogoutHandler)

	http.Handle("/", r)
}
