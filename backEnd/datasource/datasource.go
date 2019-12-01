package datasource

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"peipei2/conf"
	"peipei2/models"
	"strings"
	"time"
)

var (
	db       *gorm.DB
	dbConfig = conf.Config.DBConfig
)

func GetDB() *gorm.DB {
	return db
}

func init() {
	// user:password@(localhost)/dbname?charset=utf8&parseTime=True&loc=Local
	path := strings.Join([]string{dbConfig.DBuser, ":", dbConfig.DBpassword, "@(", dbConfig.DBip, ":", dbConfig.DBport, ")/", dbConfig.DBname, "?charset=utf8&parseTime=true"}, "")
	//log.Println(path)
	var err error
	db, err = gorm.Open("mysql", path)
	if err != nil {
		panic(err)
	}
	db.SingularTable(false) //表生成结尾不带s
	db.DB().SetConnMaxLifetime(2 * time.Second)
	db.DB().SetMaxIdleConns(20)                  //设置最大闲置个数
	db.DB().SetMaxOpenConns(2000)                //最大打开的连接数
	db.LogMode(true)                             // 启用Logger，显示详细日志
	db = db.Set("gorm:save_associations", false) //关闭自动更新关联字段
	db.LogMode(false)                            //关闭日志

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "peipei2_" + defaultTableName
	}

	if !db.HasTable(&models.School{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&models.School{}).Error; err != nil {
			panic(err)
		}
	} else {
		if err := db.AutoMigrate(&models.School{}).Error; err != nil {
			panic(err)
		}
	}

	if !db.HasTable(&models.Major{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&models.Major{}).Error; err != nil {
			panic(err)
		}
	} else {
		if err := db.AutoMigrate(&models.Major{}).Error; err != nil {
			panic(err)
		}
	}

	if !db.HasTable(&models.Student{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&models.Student{}).Error; err != nil {
			panic(err)
		}
	} else {
		if err := db.AutoMigrate(&models.Student{}).Error; err != nil {
			panic(err)
		}
	}

	if !db.HasTable(&models.Game{}) {
		if err := db.CreateTable(&models.Game{}).Error; err != nil {
			panic(err)
		}
	}

	if !db.HasTable(&models.Video{}) {
		if err := db.CreateTable(&models.Video{}).Error; err != nil {
			panic(err)
		}
	}

	if !db.HasTable(&models.Novel{}) {
		if err := db.CreateTable(&models.Novel{}).Error; err != nil {
			panic(err)
		}
	}

	if !db.HasTable(&models.Sport{}) {
		if err := db.CreateTable(&models.Sport{}).Error; err != nil {
			panic(err)
		}
	}

	//log.Println("Initializing database successfully")

}
