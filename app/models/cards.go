package models

import(
  "labix.org/v2/mgo/bson"
  "labix.org/v2/mgo"
  "flashback/db"
  //"fmt"
  //"reflect"
)

type Card struct {
  CollName      string
  UserId        bson.ObjectId
  TargetLang    string
  SourceLang    string
}

func NewCard() *Card {
  return &Card{CollName: "cards"}
}

func (m Card) Coll() *mgo.Collection {
  coll := db.Connect().C(NewCard().CollName)
  return coll
}

