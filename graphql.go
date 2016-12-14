package main

import (
	"fmt"
	"github.com/graphql-go/graphql"
)

var studentType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Student",
		Fields: graphql.Fields{
			"authorized": &graphql.Field{
				Type: graphql.Boolean,
			},
			"coursework": &graphql.Field{
				Type: graphql.NewList(courseworkType),
			},
			"midtermsGrades": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"absenceLevels": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"examsSchedule": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
		},
	},
)

var courseworkType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "coursework",
		Fields: graphql.Fields{
			"course": &graphql.Field{
				Type: graphql.String,
			},
			"grades": &graphql.Field{
				Type: graphql.NewList(gradeType),
			},
		},
	},
)

var gradeType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "grade",
		Fields: graphql.Fields{
			"module": &graphql.Field{
				Type: graphql.String,
			},
			"point": &graphql.Field{
				Type: graphql.Float,
			},
			"maxPoint": &graphql.Field{
				Type: graphql.Float,
			},
		},
	},
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"student": &graphql.Field{
				Type: studentType,
				Args: graphql.FieldConfigArgument{
					"username": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					username, isUsernameOK := p.Args["username"].(string)
					password, isPasswordOK := p.Args["password"].(string)

					if isUsernameOK && isPasswordOK {
						authorized := IsUserAuthorized(username, password)
						coursework, _ := GetUserCoursework(username, password)
						midtermsGrades, _ := GetUserMidterms(username, password)
						absenceLevels, _ := GetUserAbsenceReports(username, password)
						examsSchedule, _ := GetUserExams(username, password)

						return NewStudentAPI(authorized, coursework, midtermsGrades, absenceLevels, examsSchedule), nil
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
