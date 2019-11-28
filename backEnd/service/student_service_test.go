package service

import (
	"encoding/json"

	"testing"
)

func TestStudentService_List(t *testing.T) {
	studentService := NewStudentService()
	result := studentService.List(nil)
	data, _ := json.Marshal(result)
	t.Log(string(data))
}

func TestStudentService_Get(t *testing.T) {
	studentService := NewStudentService()
	result := studentService.Get(17722027)
	data, _ := result.JsonEncode()
	t.Log(string(data))
}
