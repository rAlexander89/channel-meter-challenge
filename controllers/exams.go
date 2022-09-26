package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rAlexander89/channel-meter-challenge/models/exam"
)

type ExamID uint

// Returns:
// 1.  slice of all exams from the Exams map.
// {
//      exams: exam.ExamsSlice
// }
func IndexExams(w http.ResponseWriter, r *http.Request) {
	var examsSlice exam.ExamSlice
	for _, value := range exam.ExamsMap {
		examsSlice = append(examsSlice, value...)
	}

	response := map[string]exam.ExamSlice{
		"exams": examsSlice,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// Returns:
// 1. A slice of exams that match the user provided ID
// 2. The average score of those exams
// {
//      "averageScore": float64,
//      "scores": []exam.ExamScores
// }
//
func ShowExam(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	// Is base10 a good idea?
	examNumber, err := strconv.ParseUint(params["number"], 10, 64)
	examNotFound := "Exam not found"
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(examNotFound)
		return
	}

	matchingExams := exam.ExamsMap[examNumber]
	if matchingExams == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(examNotFound)
		return
	}
	averageScore := exam.GetExamsAverageScore(matchingExams)
	scores := exam.GetAllExamResults(matchingExams)

	response := map[string]interface{}{
		"averageScore": averageScore,
		"scores":       scores,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}
