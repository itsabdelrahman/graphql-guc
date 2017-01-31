package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ar-maged/guc-api/factory"
	"github.com/ar-maged/guc-api/graphql"
	"github.com/ar-maged/guc-api/util"
	"github.com/gorilla/mux"
	"github.com/graphql-go/handler"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/login", loginHandler).Methods("GET")
	r.HandleFunc("/api/coursework", courseworkHandler).Methods("GET")
	r.HandleFunc("/api/midterms", midtermsHandler).Methods("GET")
	r.HandleFunc("/api/attendance", attendanceHandler).Methods("GET")
	r.HandleFunc("/api/exams", examsHandler).Methods("GET")

	r.Handle("/graphql", handler.New(&handler.Config{
		Schema: &graphql.Schema,
		Pretty: true,
	}))

	http.Handle("/", r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	fmt.Printf("Server listening on port %s...\n", port)
	http.ListenAndServe(":"+port, nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	util.SendJSONResponse(w, factory.ResponseAPI{Error: nil, Data: factory.IsUserAuthorized(util.BasicAuthentication(r))})
}

func courseworkHandler(w http.ResponseWriter, r *http.Request) {
	if coursework, err := factory.GetUserCoursework(util.BasicAuthentication(r)); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		util.SendJSONResponse(w, factory.ResponseAPI{Error: err.Error(), Data: nil})
	} else {
		util.SendJSONResponse(w, factory.ResponseAPI{Error: nil, Data: coursework})
	}
}

func midtermsHandler(w http.ResponseWriter, r *http.Request) {
	if midterms, err := factory.GetUserMidterms(util.BasicAuthentication(r)); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		util.SendJSONResponse(w, factory.ResponseAPI{Error: err.Error(), Data: nil})
	} else {
		util.SendJSONResponse(w, factory.ResponseAPI{Error: nil, Data: midterms})
	}
}

func attendanceHandler(w http.ResponseWriter, r *http.Request) {
	if reports, err := factory.GetUserAbsenceReports(util.BasicAuthentication(r)); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		util.SendJSONResponse(w, factory.ResponseAPI{Error: err.Error(), Data: nil})
	} else {
		util.SendJSONResponse(w, factory.ResponseAPI{Error: nil, Data: reports})
	}
}

func examsHandler(w http.ResponseWriter, r *http.Request) {
	if exams, err := factory.GetUserExams(util.BasicAuthentication(r)); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		util.SendJSONResponse(w, factory.ResponseAPI{Error: err.Error(), Data: nil})
	} else {
		util.SendJSONResponse(w, factory.ResponseAPI{Error: nil, Data: exams})
	}
}
