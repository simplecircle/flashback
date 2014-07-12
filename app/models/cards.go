package models

import(
  "labix.org/v2/mgo/bson"
)

type Card struct {
  UserId        bson.ObjectId
  TargetLang    string
  SourceLang    string
}

