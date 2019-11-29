package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"github.com/spf13/cast"
	"log"
	"peipei2/constants"
	"peipei2/models"
	"peipei2/service"
)

type StudentController struct {
	Sess    *sessions.Session
	Service service.StudentService
	Ctx     iris.Context
}

func NewStudentController() *StudentController {
	return &StudentController{Service: service.NewStudentService()}
}

func (g *StudentController) Get() (result models.Result) {
	r := g.Ctx
	page := r.URLParam("page")
	size := r.URLParam("size")
	if g.Sess.GetBooleanDefault("IsSuper", false) == false {
		return models.Result{Data: nil, Code: -1, Msg: "you are not super manager"}
	}
	m := make(map[string]int)

	pageInt := cast.ToInt(page)
	sizeInt := cast.ToInt(size)
	m["page"] = pageInt
	m["size"] = sizeInt
	return g.Service.List(m)
}

func (g *StudentController) Post() models.Result {

	r := g.Ctx
	action := r.URLParam("action")
	log.Println(action)
	if action == constants.HTTP_ACTION_REGISTER {
		if g.Sess.GetBooleanDefault("IsSuper", false) == false {
			return models.Result{Data: nil, Code: -1, Msg: "you are not super manager"}
		}
		student := models.Student{}
		err := r.ReadJSON(&student)
		if err != nil {
			return models.Result{Code: -1, Msg: "FAILURE: " + err.Error(), Data: nil}
		}
		return g.Service.Create(student)
	} else if action == constants.HTTP_ACTION_LOGIN {

		id, _ := r.PostValueInt("id")
		name := r.PostValue("name")
		result := g.Service.Authenticate(int(id), name)

		if result.Code == 200 {
			g.Sess.Set("Authenticated", true)
			g.Sess.Set("IsSuper", false)
			g.Sess.Set("ID", id)
		} else if result.Code == 1000 {
			g.Sess.Set("Authenticated", true)
			g.Sess.Set("IsSuper", true)
			g.Sess.Set("ID", id)
		} else {
			g.Sess.Set("Authenticated", false)
		}
		return result
	} else {
		return models.Result{Data: nil, Code: -1, Msg: "invalid path"}
	}

}

func (g *StudentController) PutBy(id int) models.Result {

	if g.Sess.GetIntDefault("ID", 0) != id && g.Sess.GetBooleanDefault("IsSuper", false) == false {
		return models.Result{Data: nil, Code: -1, Msg: "you can't modify this account"}
	}
	r := g.Ctx
	student := models.Student{}
	err := r.ReadJSON(&student)
	if err != nil {
		return models.Result{Code: -1, Msg: "FAILURE: " + err.Error(), Data: nil}
	}
	return g.Service.Save(student)
}

func (g *StudentController) GetBy(id uint) models.Result {

	if g.Sess.GetBooleanDefault("Authenticated", false) == false {
		return models.Result{Data: nil, Code: -1, Msg: "please login in first"}
	}
	isSuper := g.Sess.GetBooleanDefault("IsSuper", false)
	if isSuper {
		return g.Service.Get(id)
	}
	ID, err := g.Sess.GetInt("ID")
	if ID != cast.ToInt(id) || err != nil {
		return models.Result{Data: nil, Code: -1, Msg: "you can't get info for " + cast.ToString(id)}
	}
	return g.Service.Get(id)
}

func (g *StudentController) GetMe() models.Result {

	if g.Sess.GetBooleanDefault("Authenticated", false) == false {
		return models.Result{Data: nil, Code: -1, Msg: "please login in first"}
	}

	ID, err := g.Sess.GetInt("ID")
	if err != nil {
		return models.Result{Code: -1, Msg: "FAILURE: " + err.Error(), Data: nil}
	}
	return g.Service.Get(uint(ID))
}

func (g *StudentController) PutMe() models.Result {
	if g.Sess.GetBooleanDefault("Authenticated", false) == false {
		return models.Result{Data: nil, Code: -1, Msg: "please login in first"}
	}

	ID, err := g.Sess.GetInt("ID")
	if err != nil {
		return models.Result{Code: -1, Msg: "FAILURE: " + err.Error(), Data: nil}
	}
	r := g.Ctx
	student := models.Student{}
	err = r.ReadJSON(&student)
	if err != nil {
		return models.Result{Code: -1, Msg: "FAILURE: " + err.Error(), Data: nil}
	}
	studentID := student.ID
	if studentID != uint(ID) {
		return models.Result{Code: -1, Msg: "FAILURE: " + "information mismatch", Data: nil}
	}
	return g.Service.Save(student)

}

func (g *StudentController) DeleteBy(id uint) models.Result {

	if g.Sess.GetBooleanDefault("IsSuper", false) == false {
		return models.Result{Data: nil, Code: -1, Msg: "you are not super manager"}
	}

	student := models.Student{ID: id}
	return g.Service.Del(student)
}
