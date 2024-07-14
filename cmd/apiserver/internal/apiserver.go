package internal

import (
	"fmt"
	"log"
	"net/http"
)

func RunApiServer() {
	router := middleware(router())

	fmt.Println("Listening on 0.0.0.0:8000")
	log.Fatal(http.ListenAndServe(":8000", router).Error())
}

func router() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("/v1/user/register", userRegister)
	router.HandleFunc("/v1/user/get", withAuth(userGet))

	return router
}

func middleware(router http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Handle CORS preflight access
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Session")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		if r.Method == http.MethodGet {
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		router.ServeHTTP(w, r)
	})
}

func withAuth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// handle auth here
		// ...
		next(w, r)
	})
}
