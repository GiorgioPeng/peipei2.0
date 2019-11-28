package service

import (
	"encoding/json"
	"peipei2/models"

	"testing"
)

func TestStudentService_List(t *testing.T) {
	studentService := NewStudentService()
	m := map[string]interface{}{}
	m["size"] = 1
	m["page"] = 1
	result := studentService.List(m)
	data, _ := json.Marshal(result)
	t.Log(string(data))
}

func TestStudentService_Get(t *testing.T) {
	studentService := NewStudentService()
	result := studentService.Get(17722027)
	data, _ := result.JsonEncode()
	t.Log(string(data))
}

func TestStudentService_Save(t *testing.T) {
	studentService := NewStudentService()
	student := models.Student{ID: 1, Name: "测试用户", IsSuper: false, SchoolID: 1, MajorID: 1, NovelsType: []int{1},
		GamesType: []int{1}, VideosType: []int{1}, SportsType: []int{1}}
	result := studentService.Save(student)
	data, _ := result.JsonEncode()
	t.Log(string(data))

}

func TestStudentService_Authenticate(t *testing.T) {
	studentService := NewStudentService()
	result := studentService.Authenticate(17722017, "王逸润")
	data, _ := result.JsonEncode()
	t.Log(string(data))
}
