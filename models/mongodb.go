package models

import (
	"api-tutorial01/adapter"

	"github.com/astaxie/beego"

	mgo "gopkg.in/mgo.v2"
)

// ConnectMongo ...
func ConnectMongo() (*mgo.Session, error) {
	dbURL := adapter.CallAdapter("mongodb")
	session, err := mgo.Dial(dbURL)
	CheckError(err, "error connect mongo db")

	mgo.SetDebug(true)
	session.SetMode(mgo.Monotonic, true)

	return session, err
}

// NameDBTutorial01 ...
func NameDBTutorial01() string {
	return "api_tutorial_01"
}

// NameCollObject ...
func NameCollObject() string {
	return "object"
}

// GetCollObject ...
func GetCollObject() (*mgo.Session, *mgo.Collection) {
	conn, err := ConnectMongo()
	CheckError(err, "")

	coll := conn.DB(NameDBTutorial01()).C(NameCollObject())

	return conn, coll
}

// CheckError ...
func CheckError(err error, msg string) {
	if err != nil {
		beego.Warning(msg)
		beego.Warning(err)
	}
}
