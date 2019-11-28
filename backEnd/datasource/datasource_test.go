package datasource

import (
	"encoding/json"
	"peipei2/models"
	"testing"
	"time"
)

func TestInit(t *testing.T) {
	t.Log("testing finished")
}

func TestInitHobbies(t *testing.T){
	db := GetDB()
	for i := 1 ; i<10; i++{
		sport := models.Sport{SportType: i}
		db.Create(&sport)

	}
	for i := 1 ; i<12; i++{
		novel := models.Novel{NovelType: i, CreatedAT: time.Now()}
		db.Create(&novel)
	}
	for i := 1 ; i<14; i++{
		video := models.Video{VideoType: i, CreatedAT: time.Now()}
		db.Create(&video)
	}
	for i := 1 ; i<17; i++{
		game := models.Game{GameType: i, CreatedAT: time.Now()}
		db.Create(&game)
	}

}

func TestInitTestDB(t *testing.T) {
	db := GetDB()
	if !db.HasTable(&models.User{}) {
		if err := db.CreateTable(&models.User{}).Error; err != nil {
			panic(err)
		}
	}

	if !db.HasTable(&models.Profile{}) {
		if err := db.CreateTable(&models.Profile{}).Error; err != nil {
			panic(err)
		}
	}

	if !db.HasTable(&models.CustomizeAccount{}) {
		if err := db.CreateTable(&models.CustomizeAccount{}).Error; err != nil {
			panic(err)
		}
	}

	if !db.HasTable(&models.CustomizePerson{}) {
		if err := db.CreateTable(&models.CustomizePerson{}).Error; err != nil {
			panic(err)
		}
	}
}


func TestCreateForeign(t *testing.T){
	db := GetDB()
	profile := models.Profile{Name:"profile 3", User:models.User{Name:"小红"}}
	db.Create(&profile)
}

func TestCreateMany2Many(t *testing.T){
	db := GetDB()
	//person := models.CustomizePerson{IdPerson:"Person 1"}
	//account_1 := models.CustomizeAccount{IdAccount:"Account 1"}
	//account_2 := models.CustomizeAccount{IdAccount:"Account 2"}
	//db.Create(account_2)
	//db.Create(account_1)
	//db.Create(person)
	//db.Model(&person).Association("Accounts").Replace(account_1,account_2)
	account_3 := models.CustomizeAccount{IdAccount:"Account 3"}
	db.Create(&account_3)

}

func TestGetForeign(t *testing.T){
	db := GetDB()
	user := models.Profile{}
	db.Where("id = ?", "1").First(&user)
	data, _ := json.Marshal(&user)
	t.Logf("%+v",string(data))

}

func TestGetMany2Many(t *testing.T){
	db := GetDB()
	person := models.CustomizePerson{}
	db.Where("id_person = ?", "Person 1").First(&person)

	//accounts := []models.CustomizeAccount{}
	//Db.Model(&person).Related(&accounts,"Accounts")
	//accounts_id := []string{}
	//for _,v := range accounts{
	//	accounts_id =append(accounts_id,v.IdAccount)
	//}
	//person.AccountsID = accounts_id
	data, _ := json.Marshal(&person)
	t.Logf("%+v",string(data))
}

func TestCreate(t *testing.T){
	db := GetDB().New()
	student := models.Student{ID:17722027,Name:"王逸润",Gender:true,SchoolID:1,MajorID:1,IsSuper:true,SmokeWeight:4}
	db.Model(&student).Association("Games").Replace(&models.Game{GameType: 1},&models.Game{GameType: 3})
	db.Model(&student).Association("Videos").Replace(&models.Video{VideoType: 2},&models.Video{VideoType: 6})
	db.Model(&student).Association("Novels").Replace(&models.Novel{NovelType: 3},&models.Novel{NovelType: 4},&models.Novel{NovelType: 6})
	db.Model(&student).Association("Sports").Replace(&models.Sport{SportType: 7})
	db.Save(&student)

}

func TestGet(t *testing.T){
	db := GetDB()
	student := models.Student{}
	db.Where("id = ?", "17722027").First(&student)
	data, err := json.Marshal(student)
	if err!= nil {
		t.Error(err)
	}
	t.Logf("%v",string(data))
}