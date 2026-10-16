package main

import (
	"log/slog"
	"net/http"
	"os"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

func newRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", healthHandler)
	return mux
}

func main() {
	r := newRouter()

	if err := http.ListenAndServe("127.0.0.1:8080", r); err != nil {
		slog.Error("error to start server:" + err.Error())
		os.Exit(1)
	}

}
