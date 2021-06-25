package main

import (
	"fmt"
	"log"
	"net/http"
)
func handleFacebookCallbackT(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")

	oauthStateString := "thisshouldberandom"
	if state != oauthStateString {
		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	code := r.FormValue("code")
	fmt.Print(code)
}

func main() {
	http.HandleFunc("/login", handleFacebookCallbackT)
	fmt.Print("Started running on http://localhost:8094\n")
	log.Fatal(http.ListenAndServe(":8094", nil))
}