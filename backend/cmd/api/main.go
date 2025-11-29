package main

import (
	"hack-change-backend/internal/handlers/auth"
	"hack-change-backend/pkg/getenv"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(chiMiddleware.Logger)
	r.Use(chiMiddleware.Recoverer)

	// Add CORS middleware
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "http://"+getenv.GetValue("FRONT_HOST", "localhost")+getenv.GetValue("FRONT_PORT", ":3000"))
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		})
	})

	// r.Post("/login", users_api.Login)
	// r.Post("/register", users_api.Register)

	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", auth.Login)
		r.Post("/register", auth.Register)
		r.Post("/logout", auth.Logout)
	})

	// err := dBase.Connect()
	// if err != nil {
	// 	log.Fatal(err)
	// } else {
	// 	defer dBase.Db.Close()

	if err := http.ListenAndServe(getenv.GetValue("BACK_PORT", ":8080"), r); err != nil {
		log.Fatal(err)
		// fmt.Println(dBase.Db)
	}
	// }

}
