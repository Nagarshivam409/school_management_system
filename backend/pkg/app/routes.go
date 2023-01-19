package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func IntialiseRouter() {
	r := mux.NewRouter()
	// Route Handlers
	r.HandleFunc("/api/school/student", createStudent).Methods("POST")
	r.HandleFunc("/api/school/teacher", createTeacher).Methods("POST")
	r.HandleFunc("/api/school/student/{id}/{month}/{year}", getStudentAttendance).Methods("GET")
	r.HandleFunc("/api/school/teacher/{id}/{month}/{year}", getTeacherAttendance).Methods("GET")
	r.HandleFunc("/api/school/student/punchin", PunchinStudent).Methods("POST")
	r.HandleFunc("/api/school/student/punchout", PunchoutStudent).Methods("POST")
	r.HandleFunc("/api/school/teacher/punchin", PunchinTeacher).Methods("POST")
	r.HandleFunc("/api/school/teacher/punchout", PunchoutTeacher).Methods("POST")
	r.HandleFunc("/api/school/{class}", getClassAttendance).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))

}
