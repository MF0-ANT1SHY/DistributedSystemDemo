package grades

import (
	"fmt"
	"sync"
)

type Student struct {
	ID        int
	FirstName string
	LastName  string
	Grades    []Grade
}

//计算学生平均分数
func (s Student) Average() float32 {
	var result float32
	for _, grade := range s.Grades {
		result += grade.Score
	}
	return result / float32(len(s.Grades))
}

type Students []Student

var (
	students      Students
	studentsMutex sync.Mutex
)

//根据id寻找指定学生
func (ss Students) GetStudentById(id int) (*Student, error) {
	for i := range ss {
		if ss[i].ID == id {
			return &ss[i], nil
		}
	}
	return nil, fmt.Errorf("student with id %s not found", id)
}

type GradeType string

const (
	GradeQuiz = GradeType("Quiz")
	GradeTest = GradeType("Test")
	GradeExam = GradeType("Exam")
)

type Grade struct {
	Title string
	Type  GradeType
	Score float32
}
