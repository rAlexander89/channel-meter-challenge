package data

import (
	"encoding/json"

	"github.com/r3labs/sse/v2"
	"github.com/rAlexander89/channel-meter-challenge/models/exam"
	"github.com/rAlexander89/channel-meter-challenge/models/student"
)

// I decided to saving the incoming data to distinct maps: "Exams" and "Students".
// Though appending the incoming data to a single array would save space and a couple dozen lines of code,
// working with maps for this excersize would closer replicate working a DB than would an array.
// And since the API is only performing look-ups, searching a map would be faster than searching an array.
// var Exams exam.ExamMap = make(map[uint][]exam.Exam)
// var Exams exam.ExamsMap = make(map[uint]exam.ExamSlice)
// var Students student.StudentMap = make(map[string]student.Student)

func LoadData() {
	// I would typically API strings and such related data in an private env variable.
	// I figured that would be a bit much for this exercize.
	// Just wanted to point out what I would do in a live setting
	client := sse.NewClient("http://live-test-scores.herokuapp.com/scores")

	client.SubscribeRaw(func(msg *sse.Event) {
		var newExam exam.Exam
		var newStudent student.Student

		if err := json.Unmarshal(msg.Data, &newExam); err != nil {
			panic(err)
		}

		exam.ExamsMap[newExam.ExamNumber] = append(exam.ExamsMap[newExam.ExamNumber], newExam)

		if studentRecord, ok := student.StudentsMap[newExam.StudentID]; ok {
			studentRecord.Exams = append(studentRecord.Exams, newExam.ExamNumber)
			student.StudentsMap[newExam.StudentID] = studentRecord
		} else {
			newStudent.StudentID = newExam.StudentID
			newStudent.Exams = append(newStudent.Exams, newExam.ExamNumber)
			student.StudentsMap[newExam.StudentID] = newStudent
		}
	})
}
