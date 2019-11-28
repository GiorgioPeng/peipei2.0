package repo

import (
	"github.com/spf13/cast"
	"peipei2/datasource"
	"peipei2/models"
)

type StudentRepository interface {
	List(m map[string]interface{})  (int, []models.Student)
	Save(book models.Student) error
	Create(book models.Student)  error
	Get(id uint)  (models.Student,  error)
	Del(book models.Student)  error
}

func NewStudentRepository() StudentRepository {
	return &studentRepository{}
}

var db = datasource.GetDB()

type studentRepository struct{}

func (n studentRepository) List(m map[string]interface{}) (total int, students []models.Student) {
	db.Table("peipei2_student").Count(&total)
	err := db.Limit(cast.ToInt(m["size"])).Offset((cast.ToInt(m["page"]) - 1) * cast.ToInt(m["size"])).Find(&students).Error
	if err != nil {
		panic("select Error")
	}
	return total, students
}

func (n studentRepository) Save(student models.Student) error {
	err := db.Save(&student).Error
	return err

}

func (n studentRepository) Create(student models.Student) error {
	err := db.Create(&student).Error
	return err

}

func (n studentRepository) Get(id uint) (models.Student, error) {
	var student models.Student
	err := db.Where("id = ?", string(id)).First(&student).Error
	return student, err
}

func (n studentRepository) Del(student models.Student) (err error) {
	err = db.Unscoped().Delete(&student).Error
	return err
}
