package main

import (
	"fmt"
	"github.com/ar-maged/guc-api/graphql"
	"github.com/ar-maged/guc-api/util"
	"github.com/ar-maged/guc-api/factory"
	"github.com/graphql-go/handler"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/api/login", loginHandler)
	http.HandleFunc("/api/coursework", courseworkHandler)
	http.HandleFunc("/api/midterms", midtermsHandler)
	http.HandleFunc("/api/attendance", attendanceHandler)
	http.HandleFunc("/api/exams", examsHandler)

	http.Handle("/graphql", handler.New(&handler.Config{
		Schema: &graphql.Schema,
		Pretty: true,
	}))

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	fmt.Printf("Server listening on port %s...\n", port)
	http.ListenAndServe(":"+port, nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	util.SendJsonResponse(w, factory.ResponseAPI{nil, factory.IsUserAuthorized(util.BasicAuthentication(r))})
}

func courseworkHandler(w http.ResponseWriter, r *http.Request) {
	if coursework, err := factory.GetUserCoursework(util.BasicAuthentication(r)); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		util.SendJsonResponse(w, factory.ResponseAPI{err.Error(), nil})
	} else {
		util.SendJsonResponse(w, factory.ResponseAPI{nil, coursework})
	}
}

func midtermsHandler(w http.ResponseWriter, r *http.Request) {
	if midterms, err := factory.GetUserMidterms(util.BasicAuthentication(r)); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		util.SendJsonResponse(w, factory.ResponseAPI{err.Error(), nil})
	} else {
		util.SendJsonResponse(w, factory.ResponseAPI{nil, midterms})
	}
}

func attendanceHandler(w http.ResponseWriter, r *http.Request) {
	if reports, err := factory.GetUserAbsenceReports(util.BasicAuthentication(r)); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		util.SendJsonResponse(w, factory.ResponseAPI{err.Error(), nil})
	} else {
		util.SendJsonResponse(w, factory.ResponseAPI{nil, reports})
	}
}

func examsHandler(w http.ResponseWriter, r *http.Request) {
	if exams, err := factory.GetUserExams(util.BasicAuthentication(r)); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		util.SendJsonResponse(w, factory.ResponseAPI{err.Error(), nil})
	} else {
		util.SendJsonResponse(w, factory.ResponseAPI{nil, exams})
	}
}
