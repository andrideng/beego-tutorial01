package models

import (
	"errors"
	"strconv"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/astaxie/beego"
)

var (
	// Objects ...
	Objects map[string]*Object
)

// Object ...
type Object struct {
	ObjectID   string `json:"object_id" bson:"object_id"`
	Score      int64  `json:"score" bson:"score"`
	PlayerName string `json:"player_name" bson:"player_name"`
}

func init() {

	// conn, coll := GetCollObject()
	// defer conn.Close()
	// // var Objects = make(map[string]*Object)
	// // Objects["hjkhsbnmn123"] = &Object{"hjkhsbnmn123", 100, "astaxie"}
	// // Objects["mjjkxsxsaa23"] = &Object{"mjjkxsxsaa23", 101, "someone"}
	// var obj []Object
	// obj = append(obj, Object{"hjkhsbnmn123", 100, "astaxie"})
	// obj = append(obj, Object{"mjjkxsxsaa23", 101, "someone"})
	// // beego.Debug(obj)
	// for _, val := range obj {
	// 	count, err := coll.Find(bson.M{"object_id": val.ObjectID}).Count()
	// 	CheckError(err, "Object ID already exist")
	// 	if count == 0 {
	// 		err := coll.Insert(val)
	// 		CheckError(err, "Error insert data")
	// 	}
	//
	// }

}

// AddOne ...
func AddOne(object Object) (ObjectID string) {
	object.ObjectID = "astaxie" + strconv.FormatInt(time.Now().UnixNano(), 10)
	Objects[object.ObjectID] = &object
	return object.ObjectID
}

// GetOne ...
func GetOne(ObjectID string) (object *Object, err error) {
	// if v, ok := Objects[ObjectID]; ok {
	// 	return v, nil
	// }
	conn, coll := GetCollObject()
	defer conn.Close()

	// get ObjectController
	obj := &Object{}
	err = coll.Find(bson.M{"object_id": ObjectID}).One(obj)
	// beego.Debug(a)
	if err != nil {
		return nil, errors.New("ObjectId Not Exist")
	}
	return obj, nil
	// return nil, errors.New("ObjectId Not Exist")
}

// GetAll ...
func GetAll() *[]Object {
	conn, coll := GetCollObject()
	defer conn.Close()

	obj := &[]Object{}
	// objID := "mjjkxsxsaa23"
	coll.Find(nil).All(obj)
	beego.Debug(obj)
	return obj
}

// Update ...
func Update(ObjectID string, Score int64) (err error) {
	if v, ok := Objects[ObjectID]; ok {
		v.Score = Score
		return nil
	}
	return errors.New("ObjectId Not Exist")
}

// Delete ...
func Delete(ObjectID string) {
	delete(Objects, ObjectID)
}
