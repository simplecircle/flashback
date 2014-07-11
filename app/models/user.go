package models

import(
  //"fmt"
  //"github.com/revel/revel"
  //"labix.org/v2/mgo"
  "labix.org/v2/mgo/bson"
)

type User struct {
  Id    bson.ObjectId `bson:"_id"`
  Email        string
  Password     []byte
  AuthToken    string
}

