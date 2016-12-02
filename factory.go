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
	"strings"
)

func GetUserCoursework(username, password string) []CourseworkAPI {
	api := "https://m.guc.edu.eg"
	resource := "/StudentServices.asmx/GetCourseWork"

	response := httpPostWithFormDataCredentials(api, resource, username, password, "1.3", nil, nil)
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
	api := "https://m.guc.edu.eg"
	resource := "/StudentServices.asmx/GetCourseWork"

	response := httpPostWithFormDataCredentials(api, resource, username, password, "1.3", nil, nil)
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

func httpPostWithFormDataCredentials(api, resource, username, password, clientVersion, appOS, osVersion string) *http.Response {
	data := url.Values{}
	data.Set("username", username)
	data.Add("password", password)
	data.Add("clientVersion", clientVersion)

	if appOS != nil && osVersion != nil {
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

type XMLResponseString struct {
	Value string `xml:",chardata"`
}

type Coursework struct {
	Courses  []Course  `json:"CurrentCourses"`
	Grades   []Grade   `json:"CourseWork"`
	Midterms []Midterm `json:"Midterm"`
}

type Course struct {
	Id   string `json:"sm_crs_id"`
	Name string `json:"course_short_name"`
}

type Grade struct {
	CourseId   string `json:"sm_crs_id"`
	ModuleName string `json:"eval_method_name"`
	Point      string `json:"grade"`
	MaxPoint   string `json:"max_point"`
}

type Midterm struct {
	CourseName string `json:"course_full_name"`
	Percentage string `json:"total_perc"`
}

type CourseworkAPI struct {
	Id     string     `json:"-"`
	Code   string     `json:"code"`
	Name   string     `json:"name"`
	Grades []GradeAPI `json:"grades"`
}

type GradeAPI struct {
	Module   string `json:"module"`
	Point    string `json:"point"`
	MaxPoint string `json:"maxPoint"`
}

type MidtermAPI struct {
	Name       string `json:"name"`
	Percentage string `json:"percentage"`
}

func NewCourseworkAPI(course Course) CourseworkAPI {
	courseAPI := CourseworkAPI{}

	courseAPI.Id = course.Id
	courseAPI.Grades = []GradeAPI{}

	courseNameSplit := strings.Split(course.Name, "(")
	courseAPI.Name = strings.TrimSpace(courseNameSplit[0])
	courseAPI.Code = courseNameSplit[1][0 : len(courseNameSplit[1])-1]

	return courseAPI
}

func NewGradeAPI(grade Grade) GradeAPI {
	gradeAPI := GradeAPI{}

	gradeAPI.Module = grade.ModuleName
	gradeAPI.Point = grade.Point
	gradeAPI.MaxPoint = grade.MaxPoint

	return gradeAPI
}

func NewMidtermAPI(midterm Midterm) MidtermAPI {
	midtermAPI := MidtermAPI{}

	midtermAPI.Name = midterm.CourseName
	midtermAPI.Percentage = midterm.Percentage

	return midtermAPI
}
