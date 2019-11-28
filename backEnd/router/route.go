package router

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"github.com/kataras/iris/v12/sessions/sessiondb/redis"
	"net/http"
	"peipei2/conf"
	"peipei2/controllers"
	"time"
)

func NewApp() *iris.Application {
	cookieName := "studentCookie"
	sess := sessions.New(sessions.Config{Cookie: cookieName, Expires: 60 * time.Minute, AllowReclaim: true})
	redisConfig := conf.Config.RedisConfig
	sysConfig := conf.Config.IrisConfig
	db := redis.New(redisConfig)
	iris.RegisterOnInterrupt(func() {
		_ = db.Close()
	})
	defer db.Close()
	sess.UseDatabase(db)

	app := iris.New()
	app.Configure(iris.WithOptimizations, iris.WithConfiguration(sysConfig))
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
	})
	app.Use(crs)
	app.AllowMethods(iris.MethodGet, iris.MethodPut, iris.MethodConnect, iris.MethodPost, iris.MethodDelete)
	studentApp := mvc.New(app.Party("/students"))

	studentApp.Register(
		sess.Start,
	)
	studentApp.Handle(controllers.NewStudentController())
	return app
}

func CrossAccess11(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}

func CrossAccess(ctx iris.Context) {
	ctx.ResponseWriter().Header().Add("Access-Control-Allow-Origin", "*")
}
