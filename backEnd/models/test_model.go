package models

import (
	"github.com/jinzhu/gorm"
	"log"
)


type User struct {
	gorm.Model
	Name string `json:"name"`
}



type Profile struct {
	gorm.Model `json:"-"`
	Name      string `json:"name"`
	User      User `gorm:"foreignkey:UserRefer" json:"-"` // 使用 UserRefer 作为外键
	UserRefer uint `json:"user_refer"`
}


type CustomizePerson struct {
	IdPerson string             `gorm:"primary_key:true" json:"id_person"`
	Accounts []CustomizeAccount `gorm:"many2many:PersonAccount;association_foreignkey:idAccount;foreignkey:idPerson" json:"-"`
	AccountsID []string 			`gorm:"-" json:"accounts_id"`
}

type CustomizeAccount struct {
	IdAccount string `gorm:"primary_key:true"`
	Name      string
}


func (account *CustomizeAccount) BeforeCreate(scope *gorm.Scope) error {
	_ = scope.SetColumn("Name", "name for " +account.IdAccount)
	return nil
}

func (person *CustomizePerson) AfterFind(scope *gorm.Scope) error {
	log.Println("%+v", *person)
	accounts := []CustomizeAccount{}
	db := scope.NewDB()

	db.Model(&person).Related(&accounts,"Accounts")
	accountIDs:= []string{}
	for _,v := range accounts{
		accountIDs = append(accountIDs, v.IdAccount)
	}
	person.AccountsID = accountIDs
	log.Println("After Find called")
	return nil
}

