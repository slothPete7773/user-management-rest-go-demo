package service

import "net/http"

func MockPublicEndpoint() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Public.\n"))
	}
}

func MockPrivateEndpoint() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Private.\n"))
	}
}

func MockHealthcheckEndpoint() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Health check.\n"))
	}
}
