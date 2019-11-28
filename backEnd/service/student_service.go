package service

import (
	"peipei2/models"
	"peipei2/repo"
)

type StudentService interface {
	List(m map[string]interface{}) models.Result
	Save(student models.Student) models.Result
	Create(student models.Student) models.Result
	Get(id uint) models.Result
	Del(student models.Student) models.Result
	Authenticate(id int, name string) models.Result
}

type studentService struct {
	repo repo.StudentRepository
}

func NewStudentService() StudentService {
	return &studentService{repo: repo.NewStudentRepository()}
}

func (u studentService) List(m map[string]interface{}) models.Result {
	total, students, err := u.repo.List(m)
	result := models.Result{}
	if err != nil {
		result.Data = nil
		result.Code = -1
		result.Msg = "FAILURE: " + err.Error()
	}
	maps := make(map[string]interface{}, 2)

	maps["total"] = total
	maps["students"] = students
	result.Data = maps
	result.Code = 200
	result.Msg = "SUCCESS"
	return result
}
func (n studentService) Save(student models.Student) models.Result {
	err := n.repo.Save(student)
	result := models.Result{}
	if err != nil {
		result.Data = nil
		result.Code = -1
		result.Msg = "FAILURE: " + err.Error()
	} else {
		result.Data = nil
		result.Code = 200
		result.Msg = "SUCCESS"
	}
	return result
}

func (n studentService) Create(student models.Student) models.Result {
	err := n.repo.Create(student)
	result := models.Result{}
	if err != nil {
		result.Data = nil
		result.Code = -1
		result.Msg = "FAILURE: " + err.Error()
	} else {
		result.Data = nil
		result.Code = 200
		result.Msg = "SUCCESS"
	}
	return result
}

func (n studentService) Get(id uint) models.Result {
	student, err := n.repo.Get(id)
	var result models.Result
	if err != nil {
		result.Data = nil
		result.Code = -1
		result.Msg = err.Error()
	} else {
		result.Data = student
		result.Code = 200
		result.Msg = "SUCCESS"
	}
	return result
}
func (n studentService) Del(student models.Student) models.Result {
	err := n.repo.Del(student)
	var result models.Result
	if err != nil {
		result.Data = nil
		result.Code = -1
		result.Msg = "FAILURE: " + err.Error()
	} else {
		result.Data = nil
		result.Code = 200
		result.Msg = "SUCCESS"
	}
	return result
}

func (n studentService) Authenticate(id int, name string) models.Result {
	err, isSuper := n.repo.Authenticate(id, name)
	var result models.Result
	if err != nil {
		result.Data = nil
		result.Code = -1
		result.Msg = "FAILURE: " + err.Error()
	} else {
		if isSuper {
			result.Data = nil
			result.Code = 1000
			result.Msg = "SUPER LOGIN SUCCESS"
		} else {
			result.Data = nil
			result.Code = 200
			result.Msg = "SUCCESS"
		}
	}
	return result
}
