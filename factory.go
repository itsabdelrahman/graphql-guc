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

func GetUserCoursework(username, password string) []Grade {
	api := "https://m.guc.edu.eg"
	resource := "/StudentServices.asmx/GetCourseWork"

	response := httpPostWithFormDataCredentials(api, resource, username, password, "1.3")
	responseBodyString := httpResponseBodyToString(response.Body)

	responseString := XMLResponseString{}
	xmlToStruct(responseBodyString, &responseString)

	courseWork := Coursework{}
	jsonToStruct(responseString.Value, &courseWork)

	for i := range courseWork.Grades {
		for j := range courseWork.Courses {
			if courseWork.Grades[i].CourseId == courseWork.Courses[j].Id {
				courseWork.Grades[i].CourseName = courseWork.Courses[j].Name
			}
		}
	}

	return courseWork.Grades
}

func httpPostWithFormDataCredentials(api, resource, username, password, clientVersion string) *http.Response {
	data := url.Values{}
	data.Set("username", username)
	data.Add("password", password)
	data.Add("clientVersion", clientVersion)

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
	Courses []Course `json:"CurrentCourses"`
	Grades  []Grade  `json:"CourseWork"`
}

type Course struct {
	Id   string `json:"sm_crs_id"`
	Name string `json:"course_short_name"`
}

type Grade struct {
	CourseId   string `json:"sm_crs_id"`
	CourseName string
	ModuleName string `json:"eval_method_name"`
	Point      string `json:"grade"`
	MaxPoint   string `json:"max_point"`
}

type CourseworkAPI struct {
	Id     string
	Code   string     `json:"code"`
	Name   string     `json:"name"`
	Grades []GradeAPI `json:"grades"`
}

type GradeAPI struct {
	Module   string `json:"module"`
	Point    string `json:"point"`
	MaxPoint string `json:"maxPoint"`
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
