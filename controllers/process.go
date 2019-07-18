package controllers

import (
	"eaciit/example/connections"
	. "eaciit/example/models"
	"net/http"
	"strconv"

	db "github.com/eaciit/dbox"

	"gopkg.in/mgo.v2/bson"

	"log"

	_ "github.com/eaciit/dbox/dbc/mongo"
	"github.com/eaciit/knot/knot.v1"
	tk "github.com/eaciit/toolkit"
)

type Process struct {
}

func (p *Process) DoLogin(r *knot.WebContext) interface{} {
	r.Config.NoLog = true
	r.Config.OutputType = knot.OutputJson

	prm := struct {
		Username string
		Password string
	}{}

	err := r.GetPayload(&prm)
	if err != nil {
		log.Println(err.Error())
	}
	conn, err := connections.PrepareConnection()
	crs, err := conn.NewQuery().Select().From(new(UserModel).TableName()).Where(db.And(db.Eq("username", prm.Username), db.Eq("password", prm.Password))).Cursor(nil)
	defer crs.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
	result := make([]UserModel, 0)
	err = crs.Fetch(&result, 0, false)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer conn.Close()
	if len(result) > 0 {
		resUser := result[0]
		r.SetSession("ID", resUser.ID)
		return tk.M{"Status": true}
	} else {
		return tk.M{"Status": false}
	}
}

func (p *Process) Edit(r *knot.WebContext) interface{} {
	r.Config.OutputType = knot.OutputJson
	conn, _ := connections.PrepareConnection()
	prm := struct {
		ID string
	}{}
	e := r.GetPayload(&prm)
	if e != nil {
		log.Fatal(e.Error())
	}
	defer conn.Close()
	var ID = bson.ObjectIdHex(prm.ID)
	crs, err := conn.NewQuery().Select().From(new(StudentModel).TableName()).Where(db.Eq("_id", ID)).Cursor(nil)
	defer crs.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
	result := make([]StudentModel, 0)
	err = crs.Fetch(&result, 0, false)
	if err != nil {
		log.Fatal(err.Error())
	}
	return tk.M{"Data": result, "Status": "Success"}
}

func (p *Process) Update(r *knot.WebContext) interface{} {
	r.Config.OutputType = knot.OutputJson
	conn, err := connections.PrepareConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	prm := struct {
		ID   string
		Name string
		Age  string
	}{}
	e := r.GetPayload(&prm)
	if e != nil {
		log.Fatal(e.Error())
	}

	ageFormat, _ := strconv.Atoi(prm.Age)
	mdl := new(StudentModel)

	mdl.ID = bson.ObjectIdHex(prm.ID)
	mdl.Name = prm.Name
	mdl.Age = ageFormat
	crs := conn.NewQuery().From(new(StudentModel).TableName()).Save().Exec(tk.M{"data": mdl})
	if crs != nil {
		log.Fatal(crs.Error())
	}

	return tk.M{"Status": true, "Message": "Data has been Updated"}
}

func (p *Process) Add(r *knot.WebContext) interface{} {
	conn, _ := connections.PrepareConnection()
	if r.Request.Method != "POST" {
		tk.Println("This is not post")
	}
	prm := struct {
		Name string
		Age  string
	}{}
	e := r.GetPayload(&p)
	if e != nil {
		log.Fatal(e.Error())
	}
	ageFormat, _ := strconv.Atoi(prm.Age)
	mdl := new(StudentModel)

	mdl.ID = bson.NewObjectId()
	mdl.Name = prm.Name
	mdl.Age = ageFormat
	err := conn.NewQuery().From(new(StudentModel).TableName()).Insert().Exec(tk.M{"data": mdl})
	if err != nil {
		log.Fatal(err.Error())
	}

	http.Redirect(r.Writer, r.Request, "/views/dashboard", http.StatusTemporaryRedirect)
	return nil
}

func (p *Process) Delete(r *knot.WebContext) interface{} {
	r.Config.OutputType = knot.OutputJson
	var conn, err = connections.PrepareConnection()
	prm := struct {
		ID string
	}{}
	e := r.GetPayload(&prm)
	var idHex = bson.ObjectIdHex(prm.ID)
	if e != nil {
		log.Fatal(e.Error())
	}
	if err != nil {
		log.Fatal(err.Error())
	}
	crs := conn.NewQuery().Delete().From(new(StudentModel).TableName()).Where(db.Eq("_id", idHex)).Exec(nil)
	if crs != nil {
		log.Fatal(crs.Error())
	}

	return tk.M{"Status": true, "Message": "Data has been removed!"}
}

func (p *Process) GetData(r *knot.WebContext) interface{} {
	r.Config.OutputType = knot.OutputJson
	var conn, err = connections.PrepareConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	crs, err := conn.NewQuery().Select().From(new(StudentModel).TableName()).Cursor(nil)
	defer crs.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
	result := make([]StudentModel, 0)
	err = crs.Fetch(&result, 0, false)
	if err != nil {
		log.Fatal(err.Error())
	}
	return result
}
