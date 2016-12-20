package factory

import (
	"../util"
	"errors"
	"strings"
)

const (
	API                 = "https://m.guc.edu.eg"
	LOGIN_ENDPOINT      = "/StudentServices.asmx/Login"
	COURSEWORK_ENDPOINT = "/StudentServices.asmx/GetCourseWork"
	ATTENDANCE_ENDPOINT = "/StudentServices.asmx/GetAttendance"
	EXAMS_ENDPOINT      = "/StudentServices.asmx/GetExamsSchedule"
	CLIENT_VERSION      = "1.3"
	APP_OS              = "0"
	OS_VERSION          = "6.0.1"
)

func IsUserAuthorized(username, password string) AuthorizedAPI {
	responseBodyString := util.HttpPostWithFormData(API, LOGIN_ENDPOINT, username, password, CLIENT_VERSION, APP_OS, OS_VERSION)

	responseString := XMLResponseString{}
	util.XmlToStruct(responseBodyString, &responseString)

	return NewAuthorizedAPI(responseString.Value)
}

func GetUserCoursework(username, password string) ([]CourseworkAPI, error) {
	responseBodyString := util.HttpPostWithFormData(API, COURSEWORK_ENDPOINT, username, password, CLIENT_VERSION, "", "")

	responseString := XMLResponseString{}
	util.XmlToStruct(responseBodyString, &responseString)

	if strings.Compare(responseString.Value, "[{\"error\":\"Unauthorized\"}]") == 0 {
		return nil, errors.New("Unauthorized")
	}

	courseWork := Coursework{}
	util.JsonToStruct(responseString.Value, &courseWork)

	allCoursework := []CourseworkAPI{}

	for _, course := range courseWork.Courses {
		courseAPI := NewCourseworkAPI(course)

		for _, grade := range courseWork.Grades {
			if grade.CourseId == courseAPI.Id {
				if len(grade.Point) > 0 {
					courseAPI.Grades = append(courseAPI.Grades, NewGradeAPI(grade))
				}
			}
		}

		allCoursework = append(allCoursework, courseAPI)
	}

	return allCoursework, nil
}

func GetUserMidterms(username, password string) ([]MidtermAPI, error) {
	responseBodyString := util.HttpPostWithFormData(API, COURSEWORK_ENDPOINT, username, password, CLIENT_VERSION, "", "")

	responseString := XMLResponseString{}
	util.XmlToStruct(responseBodyString, &responseString)

	if strings.Compare(responseString.Value, "[{\"error\":\"Unauthorized\"}]") == 0 {
		return nil, errors.New("Unauthorized")
	}

	courseWork := Coursework{}
	util.JsonToStruct(responseString.Value, &courseWork)

	midtermsAPI := []MidtermAPI{}

	for _, midterm := range courseWork.Midterms {
		midtermsAPI = append(midtermsAPI, NewMidtermAPI(midterm))
	}

	return midtermsAPI, nil
}

func GetUserAbsenceReports(username, password string) ([]AbsenceReportAPI, error) {
	responseBodyString := util.HttpPostWithFormData(API, ATTENDANCE_ENDPOINT, username, password, CLIENT_VERSION, "", "")

	responseString := XMLResponseString{}
	util.XmlToStruct(responseBodyString, &responseString)

	if strings.Compare(responseString.Value, "[{\"error\":\"Unauthorized\"}]") == 0 {
		return nil, errors.New("Unauthorized")
	}

	absence := Absence{}
	util.JsonToStruct(responseString.Value, &absence)

	absenceReportsAPI := []AbsenceReportAPI{}

	for _, report := range absence.AbsenceReports {
		absenceReportsAPI = append(absenceReportsAPI, NewAbsenceReportAPI(report))
	}

	return absenceReportsAPI, nil
}

func GetUserExams(username, password string) ([]ExamAPI, error) {
	responseBodyString := util.HttpPostWithFormData(API, EXAMS_ENDPOINT, username, password, CLIENT_VERSION, "", "")

	responseString := XMLResponseString{}
	util.XmlToStruct(responseBodyString, &responseString)

	if strings.Compare(responseString.Value, "[{\"error\":\"Unauthorized\"}]") == 0 {
		return nil, errors.New("Unauthorized")
	}

	exams := []Exam{}
	util.JsonToStruct(responseString.Value, &exams)

	examsAPI := []ExamAPI{}

	for _, exam := range exams {
		examsAPI = append(examsAPI, NewExamAPI(exam))
	}

	return examsAPI, nil
}
