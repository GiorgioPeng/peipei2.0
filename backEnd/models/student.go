package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm"
	"time"
)

type School struct {
	gorm.Model
	Name     string  `gorm:"type:varchar(20);unique;not null"`
	Province string  `gorm:"type:varchar(20)"`
	City     string  `gorm:"type:varchar(20)"`
	Majors   []Major `gorm:"many2many:school_major"`
}

type Major struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);unique;not null"`
}

type Student struct {
	ID                uint    `gorm:"primary_key" json:"id"`
	Name              string  `gorm:"type:varchar(20);not null" json:"name"`
	Gender            bool    `json:"gender"`
	IsSuper           bool    `gorm:"default:false" json:"is_super"`
	Dormitory         int     `gorm:"type:int" json:"dormitory"`
	School            School  `gorm:"foreignkey:ID;association_foreignkey:SchoolID" json:"-"`
	SchoolID          uint    `json:"school_id"`
	Major             Major   `gorm:"foreignkey:ID;association_foreignkey:MajorID" json:"-"`
	MajorID           uint    `json:"major_id"`
	Province          string  `gorm:"type:varchar(10)" json:"province"`
	SouthNorth        bool    `json:"south_north"`
	SouthNorthWeight  int     `gorm:"type:int" json:"south_north_weight"`
	Sports            []Sport `gorm:"many2many:student_sport;association_foreignkey:sport_type" json:"-"`
	SportsType        []int   `gorm:"-" json:"sports_type"`
	SportsTypeWeight  uint    `gorm:"type:int" json:"sports_type_weight"`
	Novels            []Novel `gorm:"many2many:student_novel;association_foreignkey:novel_type" json:"-"`
	NovelsType        []int   `gorm:"-" json:"novels_type"`
	NovelsTypeWeight  int     `gorm:"type:int" json:"novels_type_weight"`
	Videos            []Video `gorm:"many2many:student_video;association_foreignkey:video_type" json:"-"`
	VideosType        []int   `gorm:"-" json:"videos_type"`
	VideosTypeWeight  int     `gorm:"type:int" json:"videos_type_weight"`
	Games             []Game  `gorm:"many2many:student_game;association_foreignkey:game_type" json:"-"`
	GamesType         []int   `gorm:"-" json:"games_type"`
	GamesTypeWeight   int     `gorm:"type:int" json:"games_type_weight"`
	WinterTemperature int     `gorm:"type:int" json:"winter_temperature"`
	SummerTemperature int     `gorm:"type:int" json:"summer_temperature"`
	TemperatureWeight int     `gorm:"type:int" json:"temperature_weight"`
	SleepTime         int     `gorm:"type:int" json:"sleep_time"`
	WakeTime          int     `gorm:"type:int" json:"wake_time"`
	SleepWeight       int     `gorm:"type:int" json:"sleep_weight"`
	Smoke             bool    `json:"smoke"`
	SmokeWeight       int     `gorm:"type:int" json:"smoke_weight"`
	CreatedAt         time.Time `json:"-"`
	UpdatedAt         time.Time `json:"-"`
	DeletedAt         *time.Time `sql:"index" json:"-"`
}

func (student *Student) AfterFind(scope *gorm.Scope) error {
	var (
		sportIDs []int
		novelIDs []int
		videoIDs []int
		gameIDs  []int
		sports   []Sport
		novels   []Novel
		videos   []Video
		games    []Game
	)
	db := scope.NewDB()
	db.Model(&student).Related(&sports, "Sports")
	db.Model(&student).Related(&novels, "Novels")
	db.Model(&student).Related(&videos, "Videos")
	db.Model(&student).Related(&games, "Games")
	for _, v := range sports {
		sportIDs = append(sportIDs, v.SportType)
	}
	for _, v := range novels {
		novelIDs = append(novelIDs, v.NovelType)
	}
	for _, v := range videos {
		videoIDs = append(videoIDs, v.VideoType)
	}
	for _, v := range games {
		gameIDs = append(gameIDs, v.GameType)
	}
	student.SportsType = sportIDs
	student.NovelsType = novelIDs
	student.VideosType = videoIDs
	student.GamesType = gameIDs
	return nil
}

func (student *Student) BeforeSave(scope *gorm.Scope) error {
	return nil

}