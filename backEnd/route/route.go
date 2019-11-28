package route

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"net/http"
	"peipei2/controllers"
)

func InitRouter(app *iris.Application) {
	//app.Use(CrossAccess)
	bathPath := "/api/v1"
	mvc.New(app.Party(bathPath+"/book")).Handle(controllers.NewBookController())
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



//app.Get("/", func(ctx iris.Context) {
//	_, _ = ctx.Writef("You should navigate to the /set, /get, /delete, /clear,/destroy instead")
//})
//app.Get("/set", func(ctx iris.Context) {
//	s := sess.Start(ctx)
//	// set session values
//	s.Set("name", "iris")
//
//	// test if set here
//	_, _ = ctx.Writef("All ok session value of the 'name' is: %s", s.GetString("name"))
//})
//
//app.Get("/set/{key}/{value}", func(ctx iris.Context) {
//	key, value := ctx.Params().Get("key"), ctx.Params().Get("value")
//	s := sess.Start(ctx)
//	// set session values
//	s.Set(key, value)
//
//	// test if set here
//	_, _ = ctx.Writef("All ok session value of the '%s' is: %s", key, s.GetString(key))
//})
//
//app.Get("/set/int/{key}/{value}", func(ctx iris.Context) {
//	key := ctx.Params().Get("key")
//	value, _ := ctx.Params().GetInt("value")
//	s := sess.Start(ctx)
//	// set session values
//	s.Set(key, value)
//	valueSet := s.Get(key)
//	// test if set here
//	_, _ = ctx.Writef("All ok session value of the '%s' is: %v", key, valueSet)
//})
//
//app.Get("/get/{key}", func(ctx iris.Context) {
//	key := ctx.Params().Get("key")
//	value := sess.Start(ctx).Get(key)
//
//	_, _ = ctx.Writef("The '%s' on the /set was: %v", key, value)
//})
//
//app.Get("/get", func(ctx iris.Context) {
//	// get a specific key, as string, if no found returns just an empty string
//	name := sess.Start(ctx).GetString("name")
//
//	_, _ = ctx.Writef("The 'name' on the /set was: %s", name)
//})
//
//app.Get("/get/{key}", func(ctx iris.Context) {
//	// get a specific key, as string, if no found returns just an empty string
//	name := sess.Start(ctx).GetString(ctx.Params().Get("key"))
//
//	_, _ = ctx.Writef("The name on the /set was: %s", name)
//})
//
//app.Get("/delete", func(ctx iris.Context) {
//	// delete a specific key
//	sess.Start(ctx).Delete("name")
//})
//
//app.Get("/clear", func(ctx iris.Context) {
//	// removes all entries
//	sess.Start(ctx).Clear()
//})
//
//app.Get("/destroy", func(ctx iris.Context) {
//	// destroy, removes the entire session data and cookie
//	sess.Destroy(ctx)
//})
//
//app.Get("/update", func(ctx iris.Context) {
//	// updates resets the expiration based on the session's `Expires` field.
//	if err := sess.ShiftExpiration(ctx); err != nil {
//		if errors.Is(err, sessions.ErrNotFound) {
//			ctx.StatusCode(iris.StatusNotFound)
//		} else if errors.Is(err, sessions.ErrNotImplemented) {
//			ctx.StatusCode(iris.StatusNotImplemented)
//		} else {
//			ctx.StatusCode(iris.StatusNotModified)
//		}
//
//		_, _ = ctx.Writef("%v", err)
//		ctx.Application().Logger().Error(err)
//	}
//})