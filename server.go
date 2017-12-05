package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ar-maged/guc-api/factory"
	"github.com/ar-maged/guc-api/graphql"
	"github.com/ar-maged/guc-api/util"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/graphql-go/handler"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/login", loginHandler).Methods("GET")
	router.HandleFunc("/api/coursework", courseworkHandler).Methods("GET")
	router.HandleFunc("/api/midterms", midtermsHandler).Methods("GET")
	router.HandleFunc("/api/attendance", attendanceHandler).Methods("GET")
	router.HandleFunc("/api/exams", examsHandler).Methods("GET")
	router.HandleFunc("/api/schedule", scheduleHandler).Methods("GET")

	router.Handle("/graphql", handler.New(&handler.Config{
		Schema:   &graphql.Schema,
		Pretty:   true,
		GraphiQL: true,
	}))

	http.Handle("/", router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	allowedHeaders := handlers.AllowedHeaders([]string{"Authorization"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})

	fmt.Printf("Server listening on port %s...\n", port)
	http.ListenAndServe(":"+port, handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(router))
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if username, password, err := util.BasicAuthentication(r); err != nil {
		sendUnauthorizedJSONResponse(w, err)
	} else {
		sendDataJSONResponse(w, factory.IsUserAuthorized(username, password))
	}
}

func courseworkHandler(w http.ResponseWriter, r *http.Request) {
	if username, password, err := util.BasicAuthentication(r); err != nil {
		sendUnauthorizedJSONResponse(w, err)
	} else {
		if coursework, err := factory.GetUserCoursework(username, password); err != nil {
			sendUnauthorizedJSONResponse(w, err)
		} else {
			sendDataJSONResponse(w, coursework)
		}
	}
}

func midtermsHandler(w http.ResponseWriter, r *http.Request) {
	if username, password, err := util.BasicAuthentication(r); err != nil {
		sendUnauthorizedJSONResponse(w, err)
	} else {
		if midterms, err := factory.GetUserMidterms(username, password); err != nil {
			sendUnauthorizedJSONResponse(w, err)
		} else {
			sendDataJSONResponse(w, midterms)
		}
	}
}

func attendanceHandler(w http.ResponseWriter, r *http.Request) {
	if username, password, err := util.BasicAuthentication(r); err != nil {
		sendUnauthorizedJSONResponse(w, err)
	} else {
		if reports, err := factory.GetUserAbsenceReports(username, password); err != nil {
			sendUnauthorizedJSONResponse(w, err)
		} else {
			sendDataJSONResponse(w, reports)
		}
	}
}

func examsHandler(w http.ResponseWriter, r *http.Request) {
	if username, password, err := util.BasicAuthentication(r); err != nil {
		sendUnauthorizedJSONResponse(w, err)
	} else {
		if exams, err := factory.GetUserExams(username, password); err != nil {
			sendUnauthorizedJSONResponse(w, err)
		} else {
			sendDataJSONResponse(w, exams)
		}
	}
}

func scheduleHandler(w http.ResponseWriter, r *http.Request) {
	if username, password, err := util.BasicAuthentication(r); err != nil {
		sendUnauthorizedJSONResponse(w, err)
	} else {
		if schedules, err := factory.GetUserSchedule(username, password); err != nil {
			sendUnauthorizedJSONResponse(w, err)
		} else {
			sendDataJSONResponse(w, schedules)
		}
	}
}

func sendUnauthorizedJSONResponse(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusUnauthorized)
	util.SendJSONResponse(w, factory.ResponseAPI{Error: err.Error(), Data: nil})
}

func sendDataJSONResponse(w http.ResponseWriter, data interface{}) {
	util.SendJSONResponse(w, factory.ResponseAPI{Error: nil, Data: data})
}
