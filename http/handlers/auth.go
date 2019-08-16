package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/androzd/fingo/http/response"
	"github.com/androzd/fingo/model"
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
)

var Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	user := User(r)
	log.Println(user.Username)
	//GetUserById(userId)
	w.Write([]byte("Implemented"))
})

var mySigningKey = []byte("agh4aeseeb4Chaepharit6aDuo5cohJe")

var AuthorizeHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	var user model.User
	err := user.Find(username)

	if err != nil {
		errorData := make(map[string]string)
		errorData["username"] = "User with given credentials not fond"
		response, _ := json.Marshal(response.UserNotFoundOrPasswordIsWrong{ErrorResponse: response.ErrorResponse{Errors: errorData}})
		w.Write(response)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		errorData := make(map[string]string)
		errorData["username"] = "User with given credentials not fond"
		response, _ := json.Marshal(response.UserNotFoundOrPasswordIsWrong{ErrorResponse: response.ErrorResponse{Errors: errorData}})
		w.Write(response)
		return
	}

	fmt.Println(user)
	token := jwt.New(jwt.SigningMethodHS256)

	// Устанавливаем набор параметров для токена
	mapClaims := make(jwt.MapClaims)
	mapClaims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	mapClaims["user"] = user

	token.Claims = mapClaims

	// Подписываем токен нашим секретным ключем
	tokenString, _ := token.SignedString(mySigningKey)

	response, _ := json.Marshal(response.UserLoggedIn{
		SuccessResponse: response.SuccessResponse{
			Status: "OK",
		},
		Data: response.UserData{
			User:  user,
			Token: tokenString,
		},
	})
	w.Write(response)
})

//var RegisterHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//	username := r.PostFormValue("username")
//	password := r.PostFormValue("password")
//
//	var user model.User
//	err := user.Find(username)
//
//	if err != nil {
//		errorData := make(map[string]string)
//		errorData["username"] = "Username already taken"
//		response, _ := json.Marshal(response.UserNotFoundOrPasswordIsWrong{ErrorResponse: response.ErrorResponse{Errors: errorData}})
//		w.Write(response)
//		return
//	}
//
//	pwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	user.Username = username
//	user.Password = string(pwd)
//	user.Roles = []string{"admin", "user"}
//
//	err = user.Create()
//	if err != nil {
//		errorData := make(map[string]string)
//		errorData["username"] = "Registration temporary unavailable, please try again later"
//		response, _ := json.Marshal(response.UserNotFoundOrPasswordIsWrong{ErrorResponse: response.ErrorResponse{Errors: errorData}})
//		w.Write(response)
//		return
//	}
//
//	token := jwt.New(jwt.SigningMethodHS256)
//
//	// Устанавливаем набор параметров для токена
//	mapClaims := make(jwt.MapClaims)
//	mapClaims["exp"] = time.Now().Add(time.Minute * 30).Unix()
//	mapClaims["user"] = user
//
//	token.Claims = mapClaims
//
//	// Подписываем токен нашим секретным ключем
//	tokenString, _ := token.SignedString(mySigningKey)
//
//	// Отдаем токен клиенту
//	w.Write()
//})

var GetTokenHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	var user model.User
	err := user.Find(username)

	if err != nil {
		pwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
		if err != nil {
			log.Fatal(err)
		}

		user.Username = username
		user.Password = string(pwd)
		user.Roles = []string{"admin", "user"}

		err = user.Create()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("user created")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(user)
	token := jwt.New(jwt.SigningMethodHS256)

	// Устанавливаем набор параметров для токена
	mapClaims := make(jwt.MapClaims)
	mapClaims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	mapClaims["user"] = user

	token.Claims = mapClaims

	// Подписываем токен нашим секретным ключем
	tokenString, _ := token.SignedString(mySigningKey)

	// Отдаем токен клиенту
	w.Write([]byte(tokenString))
})

var JwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})

func User(r *http.Request) (user model.User) {
	userMap := r.Context().Value("user").(*jwt.Token)
	userId := userMap.Claims.(jwt.MapClaims)["user"].(map[string]interface{})["_id"]
	user.Find(fmt.Sprintf("%v", userId))

	return
}