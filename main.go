package gaehoge

import (
	"net/http"

	"github.com/gorilla/mux"
	"app/controllers"
)

func init() {
	r := mux.NewRouter()

	r.HandleFunc("/", controllers.Index).Methods("GET")
	r.HandleFunc("/guest-book", controllers.Index).Methods("GET")
	r.HandleFunc("/guest-book/sign", controllers.GuestSign).Methods("POST")

	r.HandleFunc("/admin/show-runtime", controllers.ShowRuntime).Methods("GET")

	r.HandleFunc("/user/login", controllers.LoginHandler).Methods("GET")
	r.HandleFunc("/user/logout", controllers.LogoutHandler).Methods("GET")

	http.Handle("/", r)
}
