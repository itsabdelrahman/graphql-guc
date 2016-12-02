package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	API                 = "https://m.guc.edu.eg"
	LOGIN_ENDPOINT      = "/StudentServices.asmx/Login"
	COURSEWORK_ENDPOINT = "/StudentServices.asmx/GetCourseWork"
	ATTENDANCE_ENDPOINT = "/StudentServices.asmx/GetAttendance"
	CLIENT_VERSION      = "1.3"
	APP_OS              = "0"
	OS_VERSION          = "6.0.1"
)

func IsUserAuthorized(username, password string) AuthorizedAPI {
	response := httpPostWithFormDataCredentials(API, LOGIN_ENDPOINT, username, password, CLIENT_VERSION, APP_OS, OS_VERSION)
	responseBodyString := httpResponseBodyToString(response.Body)

	responseString := XMLResponseString{}
	xmlToStruct(responseBodyString, &responseString)

	return NewAuthorizedAPI(responseString.Value)
}

func GetUserCoursework(username, password string) []CourseworkAPI {
	response := httpPostWithFormDataCredentials(API, COURSEWORK_ENDPOINT, username, password, CLIENT_VERSION, "", "")
	responseBodyString := httpResponseBodyToString(response.Body)

	responseString := XMLResponseString{}
	xmlToStruct(responseBodyString, &responseString)

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

	return allCoursework
}

func GetUserMidterms(username, password string) []MidtermAPI {
	response := httpPostWithFormDataCredentials(API, COURSEWORK_ENDPOINT, username, password, CLIENT_VERSION, "", "")
	responseBodyString := httpResponseBodyToString(response.Body)

	responseString := XMLResponseString{}
	xmlToStruct(responseBodyString, &responseString)

	courseWork := Coursework{}
	jsonToStruct(responseString.Value, &courseWork)

	midtermsAPI := []MidtermAPI{}

	for _, midterm := range courseWork.Midterms {
		midtermsAPI = append(midtermsAPI, NewMidtermAPI(midterm))
	}

	return midtermsAPI
}

func GetUserAbsenceReports(username, password string) []AbsenceReportAPI {
	response := httpPostWithFormDataCredentials(API, ATTENDANCE_ENDPOINT, username, password, CLIENT_VERSION, "", "")
	responseBodyString := httpResponseBodyToString(response.Body)

	responseString := XMLResponseString{}
	xmlToStruct(responseBodyString, &responseString)

	absence := Absence{}
	jsonToStruct(responseString.Value, &absence)

	absenceReportsAPI := []AbsenceReportAPI{}

	for _, report := range absence.AbsenceReports {
		absenceReportsAPI = append(absenceReportsAPI, NewAbsenceReportAPI(report))
	}

	return absenceReportsAPI
}

func httpPostWithFormDataCredentials(api, resource, username, password, clientVersion, appOS, osVersion string) *http.Response {
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

func jsonToStruct(j string, v interface{}) {
	json.Unmarshal([]byte(j), v)
}

func xmlToStruct(x string, v interface{}) {
	xml.Unmarshal([]byte(x), v)
}
