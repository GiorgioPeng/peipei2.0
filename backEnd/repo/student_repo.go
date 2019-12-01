package repo

import (
	"errors"
	"github.com/spf13/cast"
	"peipei2/datasource"
	"peipei2/models"
)

type StudentRepository interface {
	List(m map[string]int) (int, []models.Student, error)
	Save(book models.Student) error
	Create(book models.Student) error
	Get(id uint) (models.Student, error)
	Del(book models.Student) error
	Authenticate(id int, name string) (error, bool)
}

func NewStudentRepository() StudentRepository {
	return &studentRepository{}
}

var db = datasource.GetDB()

type studentRepository struct{}

func (n studentRepository) List(m map[string]int) (int, []models.Student, error) {
	var (
		students []models.Student
		total    int
	)
	size := m["size"]
	page := m["page"]
	if size == 0 && page == 0 {
		err := db.Find(&students).Error
		total = len(students)
		return total, students, err
	} else {
		size := cast.ToInt(m["size"])
		page := cast.ToInt(m["page"])
		if page < 1 {
			page = 1
		}
		if size < 1 {
			size = 1
		}
		err := db.Limit(size).Offset((page - 1) * cast.ToInt(size)).Find(&students).Error
		total = len(students)
		return total, students, err
	}

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
	err := db.Where("id = ?", cast.ToString(id)).First(&student).Error
	return student, err
}

func (n studentRepository) Del(student models.Student) error {
	err := db.Delete(&student).Error //软删除
	return err
}

func (n studentRepository) Authenticate(id int, name string) (error, bool) {
	var student models.Student
	err := db.Where("id = ?", cast.ToString(id)).First(&student).Error
	if err != nil {
		return errors.New("ID does not exist"), false
	}
	if name != student.Name {
		return errors.New("wrong information"), false
	}
	return nil, student.IsSuper
}
