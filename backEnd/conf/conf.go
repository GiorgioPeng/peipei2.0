package conf

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions/sessiondb/redis"
	"io/ioutil"
	"log"
	"time"
)

type DBConfig struct {
	DBuser     string `json:"db_user"`
	DBpassword string `json:"db_password"`
	DBip       string `json:"db_ip"`
	DBport     string `json:"db_port"`
	DBname     string `json:"db_name"`
}

type Sysconfig struct {
	DBConfig    DBConfig           `json:"db_config"`
	IrisConfig  iris.Configuration `json:"iris_config"`
	RedisConfig redis.Config       `json:"redis_config"`
}

var (
	Config  = &Sysconfig{}
	hasInit = false
	//pwd, _  = os.Getwd()
	//path    = filepath.Dir(pwd)
)

func init() {
	if !hasInit {
		configPath := "config.json"
		//configPath := filepath.Join(path,"config.json")
		data, err := ioutil.ReadFile(configPath)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = json.Unmarshal(data, Config)
		if err != nil {
			fmt.Println(err)
			return
		}
		log.Println("Start loading config")
		Config.RedisConfig.Clusters = nil
		Config.RedisConfig.Timeout = time.Duration(30) * time.Second
		Config.RedisConfig.Driver = redis.Redigo()
		hasInit = true
	}
}
