package models

import (
	"api-tutorial01/adapter"
	"testing"

	mgo "gopkg.in/mgo.v2"

	. "github.com/smartystreets/goconvey/convey"
)

// TestMongoConnection ...
func TestMongoConnection(t *testing.T) {
	dbURL := adapter.CallAdapter("mongodb")
	_, err := mgo.Dial(dbURL)
	session, errModel := ConnectMongo()

	defer func() {
		session.Close()
	}()

	Convey("Subject : Mongo Connection\n", t, func() {
		Convey("mongo connected", func() {
			So(err, ShouldBeNil)
		})
		Convey("mongo connection for models", func() {
			So(errModel, ShouldBeNil)
		})
		Convey("mongo session mongotonic mode", func() {
			So(session.Mode(), ShouldEqual, mgo.Monotonic)
		})
	})
}
