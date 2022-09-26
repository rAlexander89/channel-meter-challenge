package exam

// I don't think exams will have negative numbers. Using uint.
// Specifying uint64 over 32 or leaving it platform is mostly arbitrary on my end.
// The needed to specify came func ParseUint needing a type to convert the Exam Number in GET /exam/{number}

type Exam struct {
	StudentID  string  `json:"studentId"`
	ExamNumber uint64  `json:"exam"`
	Score      float64 `json:"score"`
}

// I feel like I'm being a little "extra" by creating types such as these
// But if the data types of these values ever changes,
// changing the types below would change the types everywhere they are being used
// rather than having to hunt for every references the type directly.
//  eg: Data type of Exam.Score is changed to a discrete value.
//  Better to change it here than changing every float to a int or uint everywhere we a float was used

// Same rationale for the types in the Student model.
type ExamSlice = []Exam
type ExamScores = []float64
type ExamAverage = float64
type ExamNumber = uint64

// fake exam table
var ExamsMap map[uint64][]Exam = make(map[uint64]ExamSlice)

// Returns an average
func GetExamsAverageScore(exams ExamSlice) ExamAverage {
	var totalScore float64

	for _, exam := range exams {
		totalScore += exam.Score
	}

	return totalScore / float64(len(exams))
}

func GetAllExamResults(exams ExamSlice) ExamScores {
	var examScores ExamScores

	for _, exam := range exams {
		examScores = append(examScores, exam.Score)
	}

	return examScores

}
