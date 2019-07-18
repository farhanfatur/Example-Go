package models

import (
	"github.com/eaciit/orm"
	"gopkg.in/mgo.v2/bson"
)

type UserModel struct {
	orm.ModelBase `bson:"-"`
	ID            bson.ObjectId `bson:"_id"`
	Username      string
	Password      string
}

func NewUserModel() *UserModel {
	m := new(UserModel)
	m.ID = bson.NewObjectId()
	return m
}

func (l *UserModel) TableName() string {
	return "user"
}

func (l *UserModel) RecordID() interface{} {
	return l.ID
}
