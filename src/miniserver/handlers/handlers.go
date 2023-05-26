package handlers

import (
	"fmt"
	"net/http"
	"time"
)

func greetingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello from the rver!")
}
func lgreetingHandler(w http.ResponseWriter, r *http.Request) {
	<-time.After(5 * time.Second)
	fmt.Fprint(w, "Hello from the server delayed for 5 seconds")
}

func Handlers() http.Handler {
	handlers := http.NewServeMux()
	handlers.HandleFunc("/greeting", greetingHandler)
	handlers.HandleFunc("/lgreeting", lgreetingHandler)
	return handlers
}
