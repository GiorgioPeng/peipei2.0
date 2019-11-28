package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"github.com/spf13/cast"
	"peipei2/constants"
	"peipei2/models"
	"peipei2/service"
)

type StudentController struct {
	Ctx     iris.Context
	Service service.StudentService
	Sess    sessions.Session
}

func NewStudentController() *StudentController {
	return &StudentController{Service: service.NewStudentService()}
}

func (g *StudentController) Get() (result models.Result) {
	if g.Sess.GetBooleanDefault("IsSuper", false) == false {
		return models.Result{Data: nil, Code: -1, Msg: "you are not super manager"}
	}
	r := g.Ctx
	m := make(map[string]interface{})
	page := r.URLParam("page")
	size := r.URLParam("size")
	m["page"] = page
	m["size"] = size
	return g.Service.List(m)
}

func (g *StudentController) Post() models.Result {
	r := g.Ctx
	if r.URLParam("action") == constants.HTTP_ACTION_REGISTER {
		if g.Sess.GetBooleanDefault("IsSuper", false) == false {
			return models.Result{Data: nil, Code: -1, Msg: "you are not super manager"}
		}
		student := models.Student{}
		student.ID = cast.ToUint(r.PostValue("id"))
		student.Name = cast.ToString(r.PostValue("name"))
		student.Gender = cast.ToBool(r.PostValue("gender"))
		student.SchoolID = cast.ToUint(r.PostValue("school_id"))
		student.MajorID = cast.ToUint(r.PostValue("major_id"))
		return g.Service.Create(student)
	} else {
		name := r.PostValue("name")
		id, _ := r.PostValueInt("id")
		result := g.Service.Authenticate(id, name)
		if result.Code == 200 {
			g.Sess.Set("Authenticated", true)
			g.Sess.Set("ID", id)
		} else if result.Code == 100 {
			g.Sess.Set("Authenticated", true)
			g.Sess.Set("IsSuper", true)
			g.Sess.Set("ID", id)
		} else {
			g.Sess.Set("Authenticated", false)
		}
		return result
	}

}

func (g *StudentController) PutSave() models.Result {
	r := g.Ctx
	student := models.Student{}
	student.ID = cast.ToUint(r.PostValue("id"))
	student.Name = cast.ToString(r.PostValue("name"))
	student.Gender = cast.ToBool(r.PostValue("gender"))
	student.SchoolID = cast.ToUint(r.PostValue("school_id"))
	student.MajorID = cast.ToUint(r.PostValue("major_id"))
	student.Province = cast.ToString(r.PostValue("province"))
	student.SouthNorth = cast.ToBool(r.PostValue("south_north"))
	student.SouthNorthWeight = cast.ToInt(r.PostValue("south_north_weight"))
	student.NovelsType = cast.ToIntSlice(r.PostValue("novels_type"))
	student.NovelsTypeWeight = cast.ToInt(r.PostValue("novels_type_weight"))
	student.SportsType = cast.ToIntSlice(r.PostValue("sports_type"))
	student.SportsTypeWeight = cast.ToInt(r.PostValue("sports_type_weight"))
	student.VideosType = cast.ToIntSlice(r.PostValue("videos_type"))
	student.VideosTypeWeight = cast.ToInt(r.PostValue("videos_type_weight"))
	student.GamesType = cast.ToIntSlice(r.PostValue("games_type"))
	student.GamesTypeWeight = cast.ToInt(r.PostValue("games_type_weight"))
	student.WinterTemperature = cast.ToInt(r.PostValue("winter_temperature"))
	student.SummerTemperature = cast.ToInt(r.PostValue("summer_temperature"))
	student.TemperatureWeight = cast.ToInt(r.PostValue("temperature_weight"))
	student.SleepTime = cast.ToInt(r.PostValue("sleep_time"))
	student.WakeTime = cast.ToInt(r.PostValue("wake_time"))
	student.SleepWeight = cast.ToInt(r.PostValue("sleep_weight"))
	student.Smoke = cast.ToBool(r.PostValue("smoke"))
	student.SmokeWeight = cast.ToInt(r.PostValue("smoke_weight"))
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

func (g *StudentController) DeleteBy() models.Result {
	if g.Sess.GetBooleanDefault("IsSuper", false) == false {
		return models.Result{Data: nil, Code: -1, Msg: "you are not super manager"}
	}
	r := g.Ctx
	id := cast.ToUint(r.PostValue("id"))
	student := models.Student{ID: id}
	return g.Service.Del(student)
}
