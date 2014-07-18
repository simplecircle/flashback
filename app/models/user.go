package models

import(
  "labix.org/v2/mgo/bson"
  "labix.org/v2/mgo"
  "flashback/db"
)

type User struct {
  CollName     string
  Id           bson.ObjectId `bson:"_id"`
  Email        string
  Password     []byte
  AuthToken    string
}

func NewUser() *User {
  return &User{CollName: "users"}
}

func (m User) Coll() *mgo.Collection {
  coll := db.Connect().C(NewUser().CollName)
  return coll
}
