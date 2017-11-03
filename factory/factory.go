package factory

import (
	"errors"
	"strings"

	"github.com/ar-maged/guc-api/util"
)

const (
	api                = "https://m.guc.edu.eg"
	loginEndpoint      = "/StudentServices.asmx/Login"
	courseworkEndpoint = "/StudentServices.asmx/GetCourseWork"
	attendanceEndpoint = "/StudentServices.asmx/GetAttendance"
	examsEndpoint      = "/StudentServices.asmx/GetExamsSchedule"
	scheduleEndpoint   = "/StudentServices.asmx/GetSchedule"
	clientVersion      = "1.3"
	appOs              = "0"
	osVersion          = "6.0.1"
)

// IsUserAuthorized returns validity of student's credentials
func IsUserAuthorized(username, password string) AuthorizedAPI {
	responseBodyString := util.HTTPPostWithFormData(api, loginEndpoint, username, password, clientVersion, appOs, osVersion)

	responseString := XMLResponseString{}
	util.XMLToStruct(responseBodyString, &responseString)

	return NewAuthorizedAPI(responseString.Value)
}

// GetUserCoursework returns student's coursework
func GetUserCoursework(username, password string) ([]CourseworkAPI, error) {
	responseBodyString := util.HTTPPostWithFormData(api, courseworkEndpoint, username, password, clientVersion, "", "")

	responseString := XMLResponseString{}
	util.XMLToStruct(responseBodyString, &responseString)

	if strings.Compare(responseString.Value, "[{\"error\":\"Unauthorized\"}]") == 0 {
		return nil, errors.New("Unauthorized")
	}

	courseWork := Coursework{}
	util.JSONToStruct(responseString.Value, &courseWork)

	allCoursework := []CourseworkAPI{}

	for _, course := range courseWork.Courses {
		courseAPI := NewCourseworkAPI(course)

		for _, grade := range courseWork.Grades {
			if grade.CourseID == courseAPI.ID {
				if len(grade.Point) > 0 {
					courseAPI.Grades = append(courseAPI.Grades, NewGradeAPI(grade))
				}
			}
		}

		allCoursework = append(allCoursework, courseAPI)
	}

	return allCoursework, nil
}

// GetUserMidterms returns student's midterms grades
func GetUserMidterms(username, password string) ([]MidtermAPI, error) {
	responseBodyString := util.HTTPPostWithFormData(api, courseworkEndpoint, username, password, clientVersion, "", "")

	responseString := XMLResponseString{}
	util.XMLToStruct(responseBodyString, &responseString)

	if strings.Compare(responseString.Value, "[{\"error\":\"Unauthorized\"}]") == 0 {
		return nil, errors.New("Unauthorized")
	}

	courseWork := Coursework{}
	util.JSONToStruct(responseString.Value, &courseWork)

	midtermsAPI := []MidtermAPI{}

	for _, midterm := range courseWork.Midterms {
		midtermsAPI = append(midtermsAPI, NewMidtermAPI(midterm))
	}

	return midtermsAPI, nil
}

// GetUserAbsenceReports returns student's absence levels
func GetUserAbsenceReports(username, password string) ([]AbsenceReportAPI, error) {
	responseBodyString := util.HTTPPostWithFormData(api, attendanceEndpoint, username, password, clientVersion, "", "")
	
	responseString := XMLResponseString{}
	util.XMLToStruct(responseBodyString, &responseString)

	if strings.Compare(responseString.Value, "[{\"error\":\"Unauthorized\"}]") == 0 {
		return nil, errors.New("Unauthorized")
	}

	absence := Absence{}
	util.JSONToStruct(responseString.Value, &absence)

	absenceReportsAPI := []AbsenceReportAPI{}

	for _, report := range absence.AbsenceReports {
		absenceReportsAPI = append(absenceReportsAPI, NewAbsenceReportAPI(report))
	}

	return absenceReportsAPI, nil
}

// GetUserExams return student's exams schedule
func GetUserExams(username, password string) ([]ExamAPI, error) {
	responseBodyString := util.HTTPPostWithFormData(api, examsEndpoint, username, password, clientVersion, "", "")

	responseString := XMLResponseString{}
	util.XMLToStruct(responseBodyString, &responseString)

	if strings.Compare(responseString.Value, "[{\"error\":\"Unauthorized\"}]") == 0 {
		return nil, errors.New("Unauthorized")
	}

	exams := []Exam{}
	util.JSONToStruct(responseString.Value, &exams)

	examsAPI := []ExamAPI{}

	for _, exam := range exams {
		examsAPI = append(examsAPI, NewExamAPI(exam))
	}

	return examsAPI, nil
}

// GetUserSchedule return student's schedule
func GetUserSchedule(username, password string) ([]ScheduleAPI, error) {
	responseBodyString := util.HTTPPostWithFormData(api, scheduleEndpoint, username, password, clientVersion, "", "")

	responseString := XMLResponseString{}
	util.XMLToStruct(responseBodyString, &responseString)

	if strings.Compare(responseString.Value, "[{\"error\":\"Unauthorized\"}]") == 0 {
		return nil, errors.New("Unauthorized")
	}

	schedules := []Schedule{}
	util.JSONToStruct(responseString.Value, &schedules)

	scheduleAPI := []ScheduleAPI{}

	for _, schedule := range schedules {
		scheduleAPI = append(scheduleAPI, NewScheduleAPI(schedule))
	}

	return scheduleAPI, nil
}

