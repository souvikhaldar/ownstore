package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func isValidUser(h http.Handler) http.Handler {
	validToken := "souvikhaldar"
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rcvdToken := r.Header.Get("token")
		fmt.Println("Received token:", rcvdToken)
		if rcvdToken == validToken {
			h.ServeHTTP(w, r)
		}
	})
}

func writeMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("-----Runing writeMessage------ ")
	granularity := r.URL.Query().Get("granularity")
	date := r.URL.Query().Get("date")
	fmt.Println(granularity, date)
	w.Write([]byte(fmt.Sprintf("Hello world!")))
}

func main() {
	router := chi.NewRouter()
	router.Handle("/", http.FileServer(http.Dir(".")))
	router.Get("/hello", writeMessage)
	http.ListenAndServe(":8192", router)
}
