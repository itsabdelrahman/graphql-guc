package graphql

import (
	"strings"

	"guc-api/factory"

	"github.com/graphql-go/graphql"
)

var (
	Schema, _ = graphql.NewSchema(
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
					Resolve: resolveLogin,
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
					Resolve: resolveCoursework,
				},
				"midtermsGrades": &graphql.Field{
					Type: graphql.NewList(midtermType),
					Args: graphql.FieldConfigArgument{
						"course": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Resolve: resolveMidterms,
				},
				"absenceLevels": &graphql.Field{
					Type: graphql.NewList(absenceType),
					Args: graphql.FieldConfigArgument{
						"course": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Resolve: resolveAbsence,
				},
				"examsSchedule": &graphql.Field{
					Type: graphql.NewList(examType),
					Args: graphql.FieldConfigArgument{
						"course": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Resolve: resolveExams,
				},
				"schedule": &graphql.Field{
					Type:    graphql.NewList(scheduleEntryType),
					Resolve: resolveSchedule,
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

	scheduleEntryType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "scheduleEntry",
			Fields: graphql.Fields{
				"course": &graphql.Field{
					Type: graphql.String,
				},
				"weekday": &graphql.Field{
					Type: graphql.NewEnum(graphql.EnumConfig{
						Name: "Weekday",
						Values: graphql.EnumValueConfigMap{
							"SATURDAY": &graphql.EnumValueConfig{
								Value: "SATURDAY",
							},
							"SUNDAY": &graphql.EnumValueConfig{
								Value: "SUNDAY",
							},
							"MONDAY": &graphql.EnumValueConfig{
								Value: "MONDAY",
							},
							"TUESDAY": &graphql.EnumValueConfig{
								Value: "TUESDAY",
							},
							"WEDNESDAY": &graphql.EnumValueConfig{
								Value: "WEDNESDAY",
							},
							"THURSDAY": &graphql.EnumValueConfig{
								Value: "THURSDAY",
							},
						},
					}),
				},
				"type": &graphql.Field{
					Type: graphql.NewEnum(graphql.EnumConfig{
						Name: "Type",
						Values: graphql.EnumValueConfigMap{
							"LECTURE": &graphql.EnumValueConfig{
								Value: "LECTURE",
							},
							"TUTORIAL": &graphql.EnumValueConfig{
								Value: "TUTORIAL",
							},
							"LAB": &graphql.EnumValueConfig{
								Value: "LAB",
							},
						},
					}),
				},
				"slot": &graphql.Field{
					Type: graphql.Int,
				},
				"group": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)
)

func resolveLogin(p graphql.ResolveParams) (interface{}, error) {
	username, isUsernameOK := p.Args["username"].(string)
	password, isPasswordOK := p.Args["password"].(string)

	if isUsernameOK && isPasswordOK {
		username, password = strings.TrimSpace(username), strings.TrimSpace(password)
		return factory.StudentAPI{Username: username, Password: password, Authorized: factory.IsUserAuthorized(username, password).IsAuthorized}, nil
	}

	return nil, nil
}

func resolveCoursework(p graphql.ResolveParams) (interface{}, error) {
	courseName, isCourseNameOK := p.Args["course"].(string)

	student := p.Source.(factory.StudentAPI)
	allCoursework, _ := factory.GetUserCoursework(student.Username, student.Password)

	if isCourseNameOK {
		for _, coursework := range allCoursework {
			if strings.Contains(strings.ToUpper(coursework.Name), strings.ToUpper(courseName)) {
				return []factory.CourseworkAPI{coursework}, nil
			}
		}
	}

	return allCoursework, nil
}

func resolveMidterms(p graphql.ResolveParams) (interface{}, error) {
	courseName, isCourseNameOK := p.Args["course"].(string)

	student := p.Source.(factory.StudentAPI)
	allMidterms, _ := factory.GetUserMidterms(student.Username, student.Password)

	if isCourseNameOK {
		for _, midterm := range allMidterms {
			if strings.Contains(midterm.Name, courseName) {
				return []factory.MidtermAPI{midterm}, nil
			}
		}
	}

	return allMidterms, nil
}

func resolveAbsence(p graphql.ResolveParams) (interface{}, error) {
	courseName, isCourseNameOK := p.Args["course"].(string)

	student := p.Source.(factory.StudentAPI)
	allAbsenceLevels, _ := factory.GetUserAbsenceReports(student.Username, student.Password)

	if isCourseNameOK {
		for _, absenceLevel := range allAbsenceLevels {
			if strings.Contains(absenceLevel.CourseName, courseName) {
				return []factory.AbsenceReportAPI{absenceLevel}, nil
			}
		}
	}

	return allAbsenceLevels, nil
}

func resolveExams(p graphql.ResolveParams) (interface{}, error) {
	courseName, isCourseNameOK := p.Args["course"].(string)

	student := p.Source.(factory.StudentAPI)
	allExams, _ := factory.GetUserExams(student.Username, student.Password)

	if isCourseNameOK {
		for _, exam := range allExams {
			if strings.Contains(exam.Course, courseName) {
				return []factory.ExamAPI{exam}, nil
			}
		}
	}

	return allExams, nil
}

func resolveSchedule(p graphql.ResolveParams) (interface{}, error) {
	student := p.Source.(factory.StudentAPI)
	schedule, _ := factory.GetUserSchedule(student.Username, student.Password)

	return schedule, nil
}
