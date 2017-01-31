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
	if username, password, err := util.BasicAuthentication(r); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		util.SendJSONResponse(w, factory.ResponseAPI{Error: err.Error(), Data: nil})
	} else {
		util.SendJSONResponse(w, factory.ResponseAPI{Error: nil, Data: factory.IsUserAuthorized(username, password)})
	}
}

func courseworkHandler(w http.ResponseWriter, r *http.Request) {
	username, password, err := util.BasicAuthentication(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		util.SendJSONResponse(w, factory.ResponseAPI{Error: err.Error(), Data: nil})
		return
	}

	coursework, err := factory.GetUserCoursework(username, password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		util.SendJSONResponse(w, factory.ResponseAPI{Error: err.Error(), Data: nil})
		return
	}

	util.SendJSONResponse(w, factory.ResponseAPI{Error: nil, Data: coursework})
}

func midtermsHandler(w http.ResponseWriter, r *http.Request) {
	username, password, err := util.BasicAuthentication(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		util.SendJSONResponse(w, factory.ResponseAPI{Error: err.Error(), Data: nil})
		return
	}

	midterms, err := factory.GetUserMidterms(username, password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		util.SendJSONResponse(w, factory.ResponseAPI{Error: err.Error(), Data: nil})
		return
	}

	util.SendJSONResponse(w, factory.ResponseAPI{Error: nil, Data: midterms})
}

func attendanceHandler(w http.ResponseWriter, r *http.Request) {
	username, password, err := util.BasicAuthentication(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		util.SendJSONResponse(w, factory.ResponseAPI{Error: err.Error(), Data: nil})
		return
	}

	reports, err := factory.GetUserAbsenceReports(username, password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		util.SendJSONResponse(w, factory.ResponseAPI{Error: err.Error(), Data: nil})
		return
	}

	util.SendJSONResponse(w, factory.ResponseAPI{Error: nil, Data: reports})
}

func examsHandler(w http.ResponseWriter, r *http.Request) {
	username, password, err := util.BasicAuthentication(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		util.SendJSONResponse(w, factory.ResponseAPI{Error: err.Error(), Data: nil})
		return
	}

	exams, err := factory.GetUserExams(username, password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		util.SendJSONResponse(w, factory.ResponseAPI{Error: err.Error(), Data: nil})
		return
	}

	util.SendJSONResponse(w, factory.ResponseAPI{Error: nil, Data: exams})
}
