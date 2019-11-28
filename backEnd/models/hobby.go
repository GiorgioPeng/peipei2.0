package models

import "time"

type Sport struct {
	SportType  int `gorm:"type:int;primary_key;association_autoupdate:false"`
	CreatedAT time.Time `gorm:"association_autoupdate:false"`
}

type Novel struct {
	NovelType  int `gorm:"type:int;primary_key;association_autoupdate:false"`
	CreatedAT time.Time `gorm:"association_autoupdate:false"`
}

type Game struct {
	GameType  int `gorm:"type:int;primary_key;association_autoupdate:false"`
	CreatedAT time.Time `gorm:"association_autoupdate:false"`
}

type Video struct{
	VideoType int `gorm:"type:int;primary_key;association_autoupdate:false"`
	CreatedAT time.Time `gorm:"association_autoupdate:false"`
}
