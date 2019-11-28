package service

import (
	"peipei2/models"
	"peipei2/repo"
)

type StudentService interface {
	List (m map[string]interface{})  models.Result
	Save(student models.Student)  models.Result
	Create (student models.Student)
	Get(id uint)  models.Result
	Del(student models.Student)  models.Result
}

type studentService struct {}

func NewStudentService() StudentService{
	return &studentService{}
}

var studentRepo = repo.NewStudentRepository()

func (u studentService)List (m map[string]interface{})  models.Result{
	total, students := studentRepo.List(m)
	maps := make(map[string]interface{},2)
	maps["Total"] = total
	maps["List"] = students
	result := models.Result{}
	result.Data = maps
	result.Code = 200
	result.Msg ="SUCCESS"
	return result
}
func (n studentService) Save(student models.Student) models.Result{
	err :=  studentRepo.Save(student)
	result := models.Result{}
	if err != nil{
		result.Data = nil
		result.Code = -1
		result.Msg ="保存失败: "+err.Error()
	}else{
		result.Data = nil
		result.Code = 1
		result.Msg ="保存成功"
	}
	return result
}

func (n studentService) Create(student models.Student) models.Result{
	err :=  studentRepo.Create(student)
	result := models.Result{}
	if err != nil{
		result.Data = nil
		result.Code = -1
		result.Msg ="创建失败: " + err.Error()
	}else{
		result.Data = nil
		result.Code = 1
		result.Msg ="创建成功"
	}
	return result
}


func (n bookService) Get(id uint)(result models.Result){
	book,err := bookRepo.Get(id)
	if err!= nil{
		result.Data = nil
		result.Code = -1
		result.Msg = err.Error()
	}else{
		result.Data = book
		result.Code = 0
		result.Msg ="SUCCESS"
	}
	return
}
func (n bookService) Del(book models.Book)(result models.Result){
	err := bookRepo.Del(book)
	if err!= nil{
		result.Data = nil
		result.Code = -1
		result.Msg = err.Error()
	}else{
		result.Data = nil
		result.Code = 0
		result.Msg ="SUCCESS"
	}
	return
}

