package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
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
	response := httpPostWithFormData(API, LOGIN_ENDPOINT, username, password, CLIENT_VERSION, APP_OS, OS_VERSION)
	responseBodyString := httpResponseBodyToString(response.Body)

	responseString := XMLResponseString{}
	xmlToStruct(responseBodyString, &responseString)

	return NewAuthorizedAPI(responseString.Value)
}

func GetUserCoursework(username, password string) ([]CourseworkAPI, error) {
	response := httpPostWithFormData(API, COURSEWORK_ENDPOINT, username, password, CLIENT_VERSION, "", "")
	responseBodyString := httpResponseBodyToString(response.Body)

	responseString := XMLResponseString{}
	xmlToStruct(responseBodyString, &responseString)

	if strings.Compare(responseString.Value, "[{\"error\":\"Unauthorized\"}]") == 0 {
		return nil, errors.New("Unauthorized")
	}

	courseWork := Coursework{}
	jsonToStruct(responseString.Value, &courseWork)

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
	response := httpPostWithFormData(API, COURSEWORK_ENDPOINT, username, password, CLIENT_VERSION, "", "")
	responseBodyString := httpResponseBodyToString(response.Body)

	responseString := XMLResponseString{}
	xmlToStruct(responseBodyString, &responseString)

	if strings.Compare(responseString.Value, "[{\"error\":\"Unauthorized\"}]") == 0 {
		return nil, errors.New("Unauthorized")
	}

	courseWork := Coursework{}
	jsonToStruct(responseString.Value, &courseWork)

	midtermsAPI := []MidtermAPI{}

	for _, midterm := range courseWork.Midterms {
		midtermsAPI = append(midtermsAPI, NewMidtermAPI(midterm))
	}

	return midtermsAPI, nil
}

func GetUserAbsenceReports(username, password string) ([]AbsenceReportAPI, error) {
	response := httpPostWithFormData(API, ATTENDANCE_ENDPOINT, username, password, CLIENT_VERSION, "", "")
	responseBodyString := httpResponseBodyToString(response.Body)

	responseString := XMLResponseString{}
	xmlToStruct(responseBodyString, &responseString)

	if strings.Compare(responseString.Value, "[{\"error\":\"Unauthorized\"}]") == 0 {
		return nil, errors.New("Unauthorized")
	}

	absence := Absence{}
	jsonToStruct(responseString.Value, &absence)

	absenceReportsAPI := []AbsenceReportAPI{}

	for _, report := range absence.AbsenceReports {
		absenceReportsAPI = append(absenceReportsAPI, NewAbsenceReportAPI(report))
	}

	return absenceReportsAPI, nil
}

func GetUserExams(username, password string) ([]ExamAPI, error) {
	response := httpPostWithFormData(API, EXAMS_ENDPOINT, username, password, CLIENT_VERSION, "", "")
	responseBodyString := httpResponseBodyToString(response.Body)

	responseString := XMLResponseString{}
	xmlToStruct(responseBodyString, &responseString)

	if strings.Compare(responseString.Value, "[{\"error\":\"Unauthorized\"}]") == 0 {
		return nil, errors.New("Unauthorized")
	}

	exams := []Exam{}
	jsonToStruct(responseString.Value, &exams)

	examsAPI := []ExamAPI{}

	for _, exam := range exams {
		examsAPI = append(examsAPI, NewExamAPI(exam))
	}

	return examsAPI, nil
}

func httpPostWithFormData(api, resource, username, password, clientVersion, appOS, osVersion string) *http.Response {
	data := url.Values{}
	data.Set("username", username)
	data.Add("password", password)
	data.Add("clientVersion", clientVersion)

	if appOS != "" && osVersion != "" {
		data.Add("app_os", appOS)
		data.Add("os_version", osVersion)
	}

	uri, _ := url.ParseRequestURI(api)
	uri.Path = resource
	uriString := fmt.Sprintf("%v", uri)

	client := &http.Client{}
	request, _ := http.NewRequest("POST", uriString, bytes.NewBufferString(data.Encode()))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	response, _ := client.Do(request)
	return response
}

func httpResponseBodyToString(responseBody io.ReadCloser) string {
	responseBodyRead, _ := ioutil.ReadAll(responseBody)
	return string(responseBodyRead)
}

func jsonToStruct(j string, v interface{}) error {
	return json.Unmarshal([]byte(j), v)
}

func xmlToStruct(x string, v interface{}) error {
	return xml.Unmarshal([]byte(x), v)
}
