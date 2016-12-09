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
	sendJsonResponse(w, ResponseAPI{nil, IsUserAuthorized(basicAuthentication(r))})
}

func courseworkHandler(w http.ResponseWriter, r *http.Request) {
	if coursework, err := GetUserCoursework(basicAuthentication(r)); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		sendJsonResponse(w, ResponseAPI{err.Error(), nil})
	} else {
		sendJsonResponse(w, ResponseAPI{nil, coursework})
	}
}

func midtermsHandler(w http.ResponseWriter, r *http.Request) {
	if midterms, err := GetUserMidterms(basicAuthentication(r)); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		sendJsonResponse(w, ResponseAPI{err.Error(), nil})
	} else {
		sendJsonResponse(w, ResponseAPI{nil, midterms})
	}
}

func attendanceHandler(w http.ResponseWriter, r *http.Request) {
	if reports, err := GetUserAbsenceReports(basicAuthentication(r)); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		sendJsonResponse(w, ResponseAPI{err.Error(), nil})
	} else {
		sendJsonResponse(w, ResponseAPI{nil, reports})
	}
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
