package student

import "github.com/rAlexander89/channel-meter-challenge/models/exam"

type Student struct {
	StudentID string   `json:"studentId"`
	Exams     []uint64 `json:"exams"`
}

// fake student table
var StudentsMap map[string]Student = make(map[string]Student)

type StudentIDs []string

// Returns:
// 1. Slice of a particular student's exam scores
func (student *Student) GetStudentsExamResults() exam.ExamScores {
	var examsToFiler exam.ExamSlice

	for _, examNumber := range student.Exams {
		examsToFiler = append(examsToFiler, exam.ExamsMap[examNumber]...)
	}

	var examScores exam.ExamScores

	for _, exam := range examsToFiler {
		if exam.StudentID == student.StudentID {
			examScores = append(examScores, exam.Score)
		}
	}

	return examScores
}

// Returns:
// 1. An average of a student's test scores.
func GetStudentsAverageScore(examScores exam.ExamScores) exam.ExamAverage {
	var totalScore float64

	for _, score := range examScores {
		totalScore += score
	}

	return totalScore / float64(len(examScores))
}
