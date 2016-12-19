package main

import (
	"strings"

	"github.com/graphql-go/graphql"
)

var (
	schema, _ = graphql.NewSchema(
		graphql.SchemaConfig{
			Query: queryType,
		},
	)

	queryType = graphql.NewObject(
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
							return StudentAPI{Username: username, Password: password, Authorized: IsUserAuthorized(username, password).IsAuthorized}, nil
						}

						return nil, nil
					},
				},
			},
		})

	studentType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Student",
			Fields: graphql.Fields{
				"authorized": &graphql.Field{
					Type: graphql.Boolean,
				},
				"coursework": &graphql.Field{
					Type: graphql.NewList(courseworkType),
					Args: graphql.FieldConfigArgument{
						"course": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						courseName, isCourseNameOK := p.Args["course"].(string)

						student := p.Source.(StudentAPI)
						allCoursework, _ := GetUserCoursework(student.Username, student.Password)

						if isCourseNameOK {
							for _, coursework := range allCoursework {
								if strings.Contains(coursework.Name, courseName) {
									return []CourseworkAPI{coursework}, nil
								}
							}
						}

						return allCoursework, nil
					},
				},
				"midtermsGrades": &graphql.Field{
					Type: graphql.NewList(midtermType),
					Args: graphql.FieldConfigArgument{
						"course": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						courseName, isCourseNameOK := p.Args["course"].(string)

						student := p.Source.(StudentAPI)
						allMidterms, _ := GetUserMidterms(student.Username, student.Password)

						if isCourseNameOK {
							for _, midterm := range allMidterms {
								if strings.Contains(midterm.Name, courseName) {
									return []MidtermAPI{midterm}, nil
								}
							}
						}

						return allMidterms, nil
					},
				},
				"absenceLevels": &graphql.Field{
					Type: graphql.NewList(absenceType),
					Args: graphql.FieldConfigArgument{
						"course": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						courseName, isCourseNameOK := p.Args["course"].(string)

						student := p.Source.(StudentAPI)
						allAbsenceLevels, _ := GetUserAbsenceReports(student.Username, student.Password)

						if isCourseNameOK {
							for _, absenceLevel := range allAbsenceLevels {
								if strings.Contains(absenceLevel.CourseName, courseName) {
									return []AbsenceReportAPI{absenceLevel}, nil
								}
							}
						}

						return allAbsenceLevels, nil
					},
				},
				"examsSchedule": &graphql.Field{
					Type: graphql.NewList(examType),
					Args: graphql.FieldConfigArgument{
						"course": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						courseName, isCourseNameOK := p.Args["course"].(string)

						student := p.Source.(StudentAPI)
						allExams, _ := GetUserExams(student.Username, student.Password)

						if isCourseNameOK {
							for _, exam := range allExams {
								if strings.Contains(exam.Course, courseName) {
									return []ExamAPI{exam}, nil
								}
							}
						}

						return allExams, nil
					},
				},
			},
		},
	)

	courseworkType = graphql.NewObject(
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

	gradeType = graphql.NewObject(
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

	midtermType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "midtermsGrades",
			Fields: graphql.Fields{
				"course": &graphql.Field{
					Type: graphql.String,
				},
				"percentage": &graphql.Field{
					Type: graphql.Float,
				},
			},
		},
	)

	absenceType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "absenceLevels",
			Fields: graphql.Fields{
				"course": &graphql.Field{
					Type: graphql.String,
				},
				"level": &graphql.Field{
					Type: graphql.Int,
				},
			},
		},
	)

	examType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "examsSchedule",
			Fields: graphql.Fields{
				"course": &graphql.Field{
					Type: graphql.String,
				},
				"dateTime": &graphql.Field{
					Type: graphql.String,
				},
				"venue": &graphql.Field{
					Type: graphql.String,
				},
				"seat": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)
)
