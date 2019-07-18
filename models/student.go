package models

import (
	"github.com/eaciit/orm"
	"gopkg.in/mgo.v2/bson"
)

type StudentModel struct {
	orm.ModelBase `bson:"-"`
	ID            bson.ObjectId `bson:"_id"`
	Name          string        `bson:"name"`
	Age           int           `bson:"age"`
}

func NewStudentModel() *StudentModel {
	m := new(StudentModel)
	m.ID = bson.NewObjectId()
	return m
}

func (u *StudentModel) TableName() string {
	return "student"
}

func (u *StudentModel) RecordID() interface{} {
	return u.ID
}
