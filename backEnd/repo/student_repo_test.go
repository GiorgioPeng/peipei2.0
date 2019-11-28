package repo

import (
	"encoding/json"
	"log"
	"testing"
)

func TestMap(t *testing.T) {
	m := map[string]interface{}{}

	log.Println(m["page"])

}

func TestStudentRepository_Get(t *testing.T) {
	studentRepo := NewStudentRepository()
	student, err := studentRepo.Get(17722027)
	if err != nil {
		t.Log(err)
		return
	}
	data, _ := json.Marshal(student)
	t.Log(string(data))
}

func TestStudentRepository_List(t *testing.T) {
	m := map[string]interface{}{}
	studentRepo := NewStudentRepository()
	total, students, _ := studentRepo.List(m)
	t.Log(total)
	data, _ := json.Marshal(students)
	t.Log(string(data))
}
