package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"teacher/service"
	"fmt"
	"encoding/json"
)

type respFormat struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func HandleSuccess(message string, w http.ResponseWriter, httpStatus int, payload interface{}) {
	w.WriteHeader(httpStatus)
	rf := respFormat{}
	rf.Message = message
	rf.Data = payload
	buff, _ := json.Marshal(rf)
	w.Write(buff)
}
func HandleError(err error, message string, w http.ResponseWriter, httpStatus int) {
	w.WriteHeader(httpStatus)
	w.Write([]byte(fmt.Sprintf(message, err)))
}

func teachersHandler(w http.ResponseWriter, r *http.Request) {
	teach, err := teacher.GetAll()
	if err != nil {
		HandleError(err,"no content of teachers %v", w, http.StatusNoContent)
		return
	}
	HandleSuccess("teachers:", w, http.StatusOK, teach)
	return
}

func teacherHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	teach, err := teacher.GetOne(id)
	if err != nil {
		HandleError(err, "not found", w, http.StatusNotFound)
		return
	}
	rs, err := json.Marshal(teach)
	if err != nil {
		HandleError(err, "not found", w, http.StatusBadGateway)
		return
	}
	w.Write(rs)
}


func createTeacherHandler(w http.ResponseWriter, r *http.Request) {
	tea := new(teacher.Teacher)
	if err := json.NewDecoder(r.Body).Decode(&tea); err != nil {
		HandleError(err, "not decoded", w, http.StatusBadGateway)
	}
	if err := tea.Insert(); err != nil {
		HandleError(err, "not inserted", w, http.StatusInternalServerError)
		return
	}
	HandleSuccess("teacher created", w, http.StatusCreated, tea)
}

func deleteTeacherHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	err := teacher.Delete(id)
	if err != nil {
		HandleError(err, "no content", w, http.StatusNoContent)
		return
	}
	w.Write([]byte("DELETED!"))
}


func main() {
	router := mux.NewRouter()
	router.HandleFunc("/teachers", teachersHandler).Methods("GET")
	router.HandleFunc("/teacher/{id}", teacherHandler).Methods("GET")
	router.HandleFunc("/teacher", createTeacherHandler).Methods("POST")
	router.HandleFunc("/teacher/{id}", deleteTeacherHandler).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
