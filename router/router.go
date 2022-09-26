package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rAlexander89/channel-meter-challenge/controllers"
)

func Router() {
	router := mux.NewRouter()

	// Though the prompt does not require a PUT or DELETE path,
	// specifying the method type makes route handling obvious
	// since those methods tend to share paths with GET
	router.HandleFunc("/students", controllers.IndexStudents).Methods("GET")
	router.HandleFunc("/students/{id}", controllers.ShowStudent).Methods("GET")
	router.HandleFunc("/exams", controllers.IndexExams).Methods("GET")
	router.HandleFunc("/exams/{number}", controllers.ShowExam).Methods("GET")

	println("Server running on port 3000!")
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		panic(err)
	}
}
