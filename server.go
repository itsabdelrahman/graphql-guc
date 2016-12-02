package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/api/login", loginHandler)
	http.HandleFunc("/api/coursework", courseworkHandler)
	http.HandleFunc("/api/midterms", midtermsHandler)
	http.HandleFunc("/api/attendance", attendanceHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	fmt.Printf("Server listening on port %s...\n", port)
	http.ListenAndServe(":"+port, nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	sendJsonResponse(w, IsUserAuthorized(basicAuthentication(r)))
}

func courseworkHandler(w http.ResponseWriter, r *http.Request) {
	sendJsonResponse(w, GetUserCoursework(basicAuthentication(r)))
}

func midtermsHandler(w http.ResponseWriter, r *http.Request) {
	sendJsonResponse(w, GetUserMidterms(basicAuthentication(r)))
}

func attendanceHandler(w http.ResponseWriter, r *http.Request) {
	sendJsonResponse(w, GetUserAbsenceReports(basicAuthentication(r)))
}

func basicAuthentication(r *http.Request) (string, string) {
	auth := strings.SplitN(r.Header["Authorization"][0], " ", 2)
	payload, _ := base64.StdEncoding.DecodeString(auth[1])
	pair := strings.SplitN(string(payload), ":", 2)

	return pair[0], pair[1]
}

func sendJsonResponse(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}
