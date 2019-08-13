package main

import (
	"fmt"
	"github.com/androzd/finance/components/auth"
	"github.com/androzd/finance/model"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {
	model.Initialize()

	accounts, err := model.Account{}.All()
	if err != nil {
		panic(err)
	}
	for _, a := range accounts {
		c, _ := a.Currency()
		fmt.Println(a, c)
	}
	return
	//u := model.User{}
	//err := u.Find("andrey")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(u)
	//return
	r := mux.NewRouter()

	r.Handle("/", http.FileServer(http.Dir("./views/")))

	r.Handle("/get-token", auth.GetTokenHandler).Methods("POST")
	r.Handle("/token-checker", auth.JwtMiddleware.Handler(auth.Handler)).Methods("GET")
	r.Handle("/status", NotImplemented).Methods("GET")

	http.ListenAndServe(":3000", handlers.LoggingHandler(os.Stdout, r))
}

var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not Implemented"))
})
