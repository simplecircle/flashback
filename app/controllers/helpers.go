package controllers

import (
  "github.com/revel/revel"
  "labix.org/v2/mgo"
  "labix.org/v2/mgo/bson"
  "flashback/app/models"
)

type Helpers struct {
	*revel.Controller
}

func (c Helpers) CurrentUser() models.User {
  session, err := mgo.Dial("mongodb://elliottg:monkey75@kahana.mongohq.com:10026/flashbackDev")
  if err != nil {
          panic(err)
  }
  defer session.Close()
  cookie, _ := c.Request.Cookie("authToken")
  cookieAuthToken := cookie.Value
  currentUser := models.User{}
  coll := session.DB("flashbackDev").C("users")
  coll.Find(bson.M{"authtoken" : cookieAuthToken}).One(&currentUser)
  if err != nil {
          panic(err)
  }
  return currentUser
}
