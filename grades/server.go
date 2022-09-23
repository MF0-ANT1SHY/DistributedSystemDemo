package grades

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

//设置特定路径的Handler
func RegisterHandlers() {
	handler := new(studentsHandler)
	http.Handle("/students", handler)
	http.Handle("/students/", handler)
}

type studentsHandler struct{}

//三种访问类型: 1. /students 2. /students/{id} 3. /students/{id}/grades
func (sh studentsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pathSegments := strings.Split(r.URL.Path, "/")
	switch len(pathSegments) {
	case 2:
		sh.getAll(w, r)
	case 3:
		id, err := strconv.Atoi(pathSegments[2])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		sh.getOne(w, r, id)
	case 4:
		id, err := strconv.Atoi(pathSegments[2])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		sh.addGrade(w, r, id)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

//获取所有学生的信息
func (sh studentsHandler) getAll(w http.ResponseWriter, r *http.Request) {
	studentsMutex.Lock()
	defer studentsMutex.Unlock()

	data, err := sh.toJSON(students)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Failed to serialize student: %q", err)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}

//获取一个学生的信息
func (sh studentsHandler) getOne(w http.ResponseWriter, r *http.Request, id int) {
	studentsMutex.Lock()
	defer studentsMutex.Unlock()

	student, err := students.GetStudentById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Println(err)
		return
	}

	data, err := sh.toJSON(student)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Failed to serialize student: %q", err)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}

//为特定的学生新增成绩
func (sh studentsHandler) addGrade(w http.ResponseWriter, r *http.Request, id int) {
	studentsMutex.Lock()
	defer studentsMutex.Unlock()

	student, err := students.GetStudentById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Println(err)
		return
	}
	var g Grade
	dec := json.NewDecoder(r.Body)
	err = dec.Decode(&g)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}
	student.Grades = append(student.Grades, g)
	w.WriteHeader(http.StatusCreated)
	data, err := sh.toJSON(g)
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}

func (sh studentsHandler) toJSON(obj interface{}) ([]byte, error) {
	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	err := enc.Encode(obj)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize students: %q", err)
	}
	return b.Bytes(), nil
}
