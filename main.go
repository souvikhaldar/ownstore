package main

import (
	"fmt"
	"net/http"
)

func isValidUser(h http.Handler) http.HandlerFunc {
	validToken := "souvikhaldar"
	return func(w http.ResponseWriter, r *http.Request) {
		rcvdToken := r.Header.Get("token")
		fmt.Println("Received token:", rcvdToken)
		if rcvdToken == validToken {
			h.ServeHTTP(w, r)
		}
	}
}

func writeMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("-----Runing writeMessage------ ")
	w.Write([]byte("Hello world!"))
}
func main() {
	//r := chi.NewRouter()
	//r.Get("/", isValidUser(writeMessage))
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":8192", nil)
}
