package main

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"github.com/kataras/iris/v12/sessions/sessiondb/redis"
	"peipei2/conf"
	"time"
)

var(
	cookieName = "studentCookie"
	sess = sessions.New(sessions.Config{Cookie:cookieName,Expires:60* time.Minute,AllowReclaim:true})
	sysConfig = conf.Config.IrisConfig
	redisConfig = conf.Config.RedisConfig
)



func main()  {
	db := redis.New(redisConfig)
	iris.RegisterOnInterrupt(func() {
		_ = db.Close()
	})
	defer db.Close()
	sess.UseDatabase(db)

	app := newApp()
	app.Configure()

	err := app.Run(iris.Addr(":8000"), iris.WithoutServerError(iris.ErrServerClosed))
	if err!=nil{
		panic(err)
	}

}


func newApp() *iris.Application {
	app := iris.New()
	app.Configure(iris.WithOptimizations,iris.WithConfiguration(sysConfig))
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
	})
	app.Use(crs)
	app.AllowMethods(iris.MethodGet,iris.MethodPut,iris.MethodConnect,iris.MethodPost)
	return app
}