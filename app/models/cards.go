package models

import(
  "labix.org/v2/mgo/bson"
  "labix.org/v2/mgo"
  "flashback/conf"
  "fmt"
  "reflect"
)

type Card struct {
  UserId        bson.ObjectId
  TargetLang    string
  SourceLang    string
}

func Collection(collection string) *mgo.Collection {
  //coll := conf.Coll("cards")

  coll := conf.Db().C(collection)
  fmt.Println("+++++++++++++++++++")
  fmt.Println(reflect.TypeOf(coll))
  //fmt.Println(coll)
  return coll
}
