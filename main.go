package main

import (
	"flag"
	"fmt"
	"github.com/androzd/fingo/model"
	"github.com/androzd/fingo/route"
	"github.com/gorilla/handlers"
	"net/http"
	"os"
	"time"
)

func main() {
	var startType string
	flag.StringVar(&startType, "type", "server", "Type of loading [server, seeder]")
	flag.Parse()

	model.Initialize()
	switch startType {
	case "server":
		server()
	case "seeder":
		seeder()
	default:
		fmt.Println("-type must be one of [server, seeder]")
	}
}

func seeder() {
	var user model.User
	err := user.Find("andrey")
	if err == nil {
		fmt.Println("Seeder prevented, because user fond")
		fmt.Println(user)
	}
	return
	//pwd, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.MinCost)
	//_ = model.User{
	//	Username: "andrey",
	//	Password: string(pwd),
	//	Roles:    []string{"admin", "user"},
	//}.Create()
}
func server() {
	r := route.Initialize()

	go http.ListenAndServe(":3000", handlers.LoggingHandler(os.Stdout, r))

	for {
		time.Sleep(1 * time.Minute)
	}
}
