package route

import (
	"github.com/androzd/fingo/http/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func Initialize() *mux.Router {
	r := mux.NewRouter()
	r.Handle("/", http.FileServer(http.Dir("./views/")))
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	})
	r.Handle("/auth/authorize", handlers.AuthorizeHandler).Methods("POST")
	r.Handle("/auth/register", handlers.GetTokenHandler).Methods("POST")
	r.Handle("/auth/token-checker", handlers.JwtMiddleware.Handler(handlers.Handler)).Methods("GET")
	r.Handle("/status", handlers.NotImplemented).Methods("GET")

	return r
}
