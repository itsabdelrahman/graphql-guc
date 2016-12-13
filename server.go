package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	"net/http"
	"os"
	"strings"
)

type user struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var data map[string]user

var userType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type: userType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					idQuery, isOK := p.Args["id"].(string)
					if isOK {
						return data[idQuery], nil
					}
					return nil, nil
				},
			},
		},
	})

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: queryType,
	},
)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("Wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}

func main() {
	data = make(map[string]user)
	data["1"] = user{"1", "Dan"}

	http.HandleFunc("/api/login", loginHandler)
	http.HandleFunc("/api/coursework", courseworkHandler)
	http.HandleFunc("/api/midterms", midtermsHandler)
	http.HandleFunc("/api/attendance", attendanceHandler)
	http.HandleFunc("/api/exams", examsHandler)
	http.HandleFunc("/graphql", graphqlHandler)

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

func examsHandler(w http.ResponseWriter, r *http.Request) {
	if exams, err := GetUserExams(basicAuthentication(r)); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		sendJsonResponse(w, ResponseAPI{err.Error(), nil})
	} else {
		sendJsonResponse(w, ResponseAPI{nil, exams})
	}
}

func graphqlHandler(w http.ResponseWriter, r *http.Request) {
	result := executeQuery(r.URL.Query()["query"][0], schema)
	sendJsonResponse(w, result)
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
