package factory

import (
	"strconv"
	"strings"
	"time"
)

type (
	// XMLResponseString is the generic incoming response
	XMLResponseString struct {
		Value string `xml:",chardata"`
	}

	// Coursework is the incoming coursework representation
	Coursework struct {
		Courses  []Course  `json:"CurrentCourses"`
		Grades   []Grade   `json:"CourseWork"`
		Midterms []Midterm `json:"Midterm"`
	}

	// Course is the incoming course representation
	Course struct {
		ID   string `json:"sm_crs_id"`
		Name string `json:"course_short_name"`
	}

	// Grade is the incoming grade representation
	Grade struct {
		CourseID   string `json:"sm_crs_id"`
		ModuleName string `json:"eval_method_name"`
		Point      string `json:"grade"`
		MaxPoint   string `json:"max_point"`
	}

	// Midterm is the incoming midterm representation
	Midterm struct {
		CourseName string `json:"course_full_name"`
		Percentage string `json:"total_perc"`
	}

	// Absence is the incoming absence representation
	Absence struct {
		AbsenceReports []AbsenceReport `json:"AbsenceReport"`
	}

	// AbsenceReport is the incoming absence report representation
	AbsenceReport struct {
		CourseName   string `json:"Name"`
		AbsenceLevel string `json:"AbsenceLevel"`
	}

	// Exam is the incoming exam representation
	Exam struct {
		Course   string `json:"course_name"`
		DateTime string `json:"start_time"`
		Venue    string `json:"rsrc_code"`
		Seat     string `json:"seat_code"`
	}

	// Schedule is your schedule
	Schedule struct {
		Slot    string `json:"scd_col"`
		Course  string `json:"course"`
		Weekday string `json:"weekday"`
		Group   string `json:"group_name"`
	}

	// ResponseAPI is the generic outgoing response
	ResponseAPI struct {
		Error interface{} `json:"error"`
		Data  interface{} `json:"data"`
	}

	// AuthorizedAPI is the outgoing authorized representation
	AuthorizedAPI struct {
		IsAuthorized bool `json:"authorized"`
	}

	// CourseworkAPI is the outgoing coursework representation
	CourseworkAPI struct {
		ID     string     `json:"-"`
		Code   string     `json:"-"`
		Name   string     `json:"course"`
		Grades []GradeAPI `json:"grades"`
	}

	// GradeAPI is the outgoing grade representation
	GradeAPI struct {
		Module   string `json:"module"`
		Point    string `json:"point"`
		MaxPoint string `json:"maxPoint"`
	}

	// MidtermAPI is the outgoing midterm representation
	MidtermAPI struct {
		Name       string `json:"course"`
		Percentage string `json:"percentage"`
	}

	// AbsenceReportAPI is the outgoing absence report representation
	AbsenceReportAPI struct {
		CourseName string `json:"course"`
		Level      string `json:"level"`
	}

	// ExamAPI is the outgoing exam representation
	ExamAPI struct {
		Course   string    `json:"course"`
		DateTime time.Time `json:"dateTime"`
		Venue    string    `json:"venue"`
		Seat     string    `json:"seat"`
	}

	// ScheduleAPI is the outgoing schedule representation
	ScheduleAPI struct {
		Weekday string `json:"weekday"`
		Slot    int    `json:"slot"`
		Course  string `json:"course"`
		Group   string `json:"group_name"`
	}

	// StudentAPI is the encapsulating representation of all student's data
	StudentAPI struct {
		Username        string             `json:"-"`
		Password        string             `json:"-"`
		Authorized      bool               `json:"authorized"`
		Coursework      []CourseworkAPI    `json:"coursework"`
		MidtermsGrades  []MidtermAPI       `json:"midtermsGrades"`
		AbsenceLevels   []AbsenceReportAPI `json:"absenceLevels"`
		ExamsSchedule   []ExamAPI          `json:"examsSchedule"`
		StudentSchedule []ScheduleAPI      `json:"studentSchedule"`
	}
)

// NewAuthorizedAPI is the AuthorizedAPI constructor
func NewAuthorizedAPI(authorized string) AuthorizedAPI {
	authorizedAPI := AuthorizedAPI{}

	if strings.Compare(authorized, "True") == 0 {
		authorizedAPI.IsAuthorized = true
	} else {
		authorizedAPI.IsAuthorized = false
	}

	return authorizedAPI
}

// NewCourseworkAPI is the CourseworkAPI constructor
func NewCourseworkAPI(course Course) CourseworkAPI {
	courseAPI := CourseworkAPI{}

	courseAPI.ID = course.ID
	courseAPI.Grades = []GradeAPI{}

	courseNameSplit := strings.Split(course.Name, "(")
	courseAPI.Name = strings.TrimSpace(courseNameSplit[0])
	courseAPI.Code = courseNameSplit[1][0 : len(courseNameSplit[1])-1]

	return courseAPI
}

// NewGradeAPI is the GradeAPI constructor
func NewGradeAPI(grade Grade) GradeAPI {
	gradeAPI := GradeAPI{}

	gradeAPI.Module = grade.ModuleName
	gradeAPI.Point = grade.Point
	gradeAPI.MaxPoint = grade.MaxPoint

	return gradeAPI
}

// NewMidtermAPI is the MidtermAPI constructor
func NewMidtermAPI(midterm Midterm) MidtermAPI {
	midtermAPI := MidtermAPI{}

	nameAndCode := strings.TrimSpace(strings.Split(midterm.CourseName, "-")[1])
	midtermAPI.Name = strings.TrimSpace(nameAndCode[:strings.LastIndex(nameAndCode, " ")])

	midtermAPI.Percentage = midterm.Percentage

	return midtermAPI
}

// NewAbsenceReportAPI is the AbsenceReportAPI constructor
func NewAbsenceReportAPI(absenceReport AbsenceReport) AbsenceReportAPI {
	absenceReportAPI := AbsenceReportAPI{}

	absenceReportAPI.CourseName = absenceReport.CourseName
	absenceReportAPI.Level = absenceReport.AbsenceLevel

	return absenceReportAPI
}

// NewExamAPI is the ExamAPI constructor
func NewExamAPI(exam Exam) ExamAPI {
	examAPI := ExamAPI{}

	codeAndName := strings.TrimSpace(strings.Split(exam.Course, "-")[1])
	examAPI.Course = strings.TrimSpace(codeAndName[strings.Index(codeAndName, " "):])

	examAPI.DateTime, _ = time.Parse("Jan 2 2006  3:04PM", exam.DateTime)

	examAPI.Venue = exam.Venue
	examAPI.Seat = exam.Seat

	return examAPI
}

// NewScheduleAPI is the ScheduleApi constructor
func NewScheduleAPI(schedule Schedule) ScheduleAPI {
	scheduleAPI := ScheduleAPI{}

	slotNo, err := strconv.Atoi(schedule.Slot)
	if err != nil {

	}

	scheduleAPI.Course = schedule.Course
	scheduleAPI.Group = schedule.Group
	scheduleAPI.Slot = slotNo
	scheduleAPI.Weekday = schedule.Weekday

	return scheduleAPI
}

