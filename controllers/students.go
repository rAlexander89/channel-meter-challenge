package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rAlexander89/channel-meter-challenge/models/student"
)

// Returns:
//  1. a slice of all students that have recieved at least one score
// It appears to me that, by definition, all data sent over is of scored exams
// So, if a student exists in the map, they have a scored exam... Or am I missing something?
// {
//      students: []Student
// }
func IndexStudents(w http.ResponseWriter, r *http.Request) {
	var studentSlice student.StudentIDs

	for key, _ := range student.StudentsMap {
		studentSlice = append(studentSlice, key)
	}

	response := map[string][]string{
		"students": studentSlice,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// Returns:
// 1. A slice a slice of exams whose studentId matches the user provided ID
// 2. The average score of those exams
// {
//     "averageScore": exam.ExamAverage,
//     "scores": []exam.ExamScores
// }
func ShowStudent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	matchingStudent := student.StudentsMap[params["id"]]

	if matchingStudent.StudentID == "" {
		notFound := "Student not found"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(notFound)
		return
	}

	examResults := matchingStudent.GetStudentsExamResults()
	averageScore := student.GetStudentsAverageScore(examResults)

	response := map[string]interface{}{
		"averageScore": averageScore,
		"scores":       examResults,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
