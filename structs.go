package main

import (
	"strings"
)

type (
	XMLResponseString struct {
		Value string `xml:",chardata"`
	}

	Coursework struct {
		Courses  []Course  `json:"CurrentCourses"`
		Grades   []Grade   `json:"CourseWork"`
		Midterms []Midterm `json:"Midterm"`
	}

	Course struct {
		Id   string `json:"sm_crs_id"`
		Name string `json:"course_short_name"`
	}

	Grade struct {
		CourseId   string `json:"sm_crs_id"`
		ModuleName string `json:"eval_method_name"`
		Point      string `json:"grade"`
		MaxPoint   string `json:"max_point"`
	}

	Midterm struct {
		CourseName string `json:"course_full_name"`
		Percentage string `json:"total_perc"`
	}

	Absence struct {
		AbsenceReports []AbsenceReport `json:"AbsenceReport"`
	}

	AbsenceReport struct {
		CourseName   string `json:"Name"`
		AbsenceLevel string `json:"AbsenceLevel"`
	}

	ResponseAPI struct {
		Error interface{} `json:"error"`
		Data  interface{} `json:"data"`
	}

	AuthorizedAPI struct {
		IsAuthorized bool `json:"authorized"`
	}

	CourseworkAPI struct {
		Id     string     `json:"-"`
		Code   string     `json:"code"`
		Name   string     `json:"name"`
		Grades []GradeAPI `json:"grades"`
	}

	GradeAPI struct {
		Module   string `json:"module"`
		Point    string `json:"point"`
		MaxPoint string `json:"maxPoint"`
	}

	MidtermAPI struct {
		Name       string `json:"name"`
		Percentage string `json:"percentage"`
	}

	AbsenceReportAPI struct {
		CourseName string `json:"name"`
		Level      string `json:"level"`
	}
)

func NewAuthorizedAPI(authorized string) AuthorizedAPI {
	authorizedAPI := AuthorizedAPI{}

	if strings.Compare(authorized, "True") == 0 {
		authorizedAPI.IsAuthorized = true
	} else {
		authorizedAPI.IsAuthorized = false
	}

	return authorizedAPI
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

func NewAbsenceReportAPI(absenceReport AbsenceReport) AbsenceReportAPI {
	absenceReportAPI := AbsenceReportAPI{}

	absenceReportAPI.CourseName = absenceReport.CourseName
	absenceReportAPI.Level = absenceReport.AbsenceLevel

	return absenceReportAPI
}
