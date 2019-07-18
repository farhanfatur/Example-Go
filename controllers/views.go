package controllers

import (
	knot "github.com/eaciit/knot/knot.v1"
)

type Views struct {
}

func (l *Views) Login(r *knot.WebContext) interface{} {
	r.Config.OutputType = knot.OutputTemplate
	r.Config.ViewName = "_template.html"
	return nil
}

func (l *Views) Dashboard(r *knot.WebContext) interface{} {
	r.Config.OutputType = knot.OutputTemplate
	r.Config.ViewName = "dashboard.html"
	return nil
}

func (l *Views) Add(r *knot.WebContext) interface{} {
	r.Config.OutputType = knot.OutputTemplate
	r.Config.ViewName = "add.html"
	return nil
}
