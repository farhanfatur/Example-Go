package main

import (
	"eaciit/example/controllers"
	"net/http"
	"os"

	knot "github.com/eaciit/knot/knot.v1"
)

var (
	appViewsPath = (func(dir string, _ error) string { return dir + "/views/" }(os.Getwd()))
)

func main() {
	var routes = map[string]knot.FnContent{
		"/": func(r *knot.WebContext) interface{} {
			http.Redirect(r.Writer, r.Request, "/views/login", http.StatusTemporaryRedirect)
			return true
		},
	}
	app := knot.NewApp("")
	app.ViewsPath = appViewsPath
	app.Static("cssstatic", "asset/css")
	app.Static("jsstatics", "asset/js")
	app.Statics()
	app.Register(new(controllers.Views))
	app.Register(new(controllers.Process))
	// app.LayoutTemplate = "_template.html"
	app.DefaultOutputType = knot.OutputTemplate

	knot.RegisterApp(app)
	knot.StartAppWithFn(app, "localhost:1234", routes)
}
