package repo

import (
	"encoding/json"
	"log"
	"peipei2/models"
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

func TestStudentRepository_Save(t *testing.T) {
	n := NewStudentRepository()
	student := models.Student{ID: 1, Name: "测试用户", IsSuper: false, SchoolID: 1, MajorID: 1, NovelsType: []int{1},
		GamesType: []int{1}, VideosType: []int{2}, SportsType: []int{1}}
	err := n.Save(student)
	log.Println(err)

}

func TestStudentRepository_Authenticate(t *testing.T) {
	n := NewStudentRepository()
	err, isSuper := n.Authenticate(17722026, "王逸润")
	log.Println(err)
	log.Println(isSuper)
}
