package router

import (
	"context"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"github.com/kataras/iris/v12/sessions/sessiondb/redis"
	"peipei2/conf"
	"peipei2/controllers"
	"time"
)

func NewApp() *iris.Application {
	cookieName := "studentCookie"
	sess := sessions.New(sessions.Config{Cookie: cookieName, Expires: 45 * time.Minute, AllowReclaim: true})
	redisConfig := conf.Config.RedisConfig
	sysConfig := conf.Config.IrisConfig
	//log.Println(redisConfig)
	db := redis.New(redisConfig)

	sess.UseDatabase(db)

	app := iris.New()
	app.Configure(iris.WithOptimizations, iris.WithConfiguration(sysConfig))
	//crs := cors.New(cors.Options{
	//	AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
	//	AllowCredentials: true,
	//	AllowedHeaders:   []string{"*"},
	//})
	//app.Use(crs)
	studentApp := mvc.New(app.Party("/students"))
	iris.RegisterOnInterrupt(func() {
		_ = db.Close()
		timeout := 3 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		// 关闭所有主机
		_ = app.Shutdown(ctx)
	})

	studentApp.Register(
		sess.Start,
	)
	studentApp.Handle(controllers.NewStudentController())
	return app
}
