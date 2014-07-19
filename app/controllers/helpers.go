package controllers

import (
  "github.com/revel/revel"
  "labix.org/v2/mgo/bson"
  "flashback/app/models"
  //"fmt"
)

type Helpers struct {
	*revel.Controller
}

func (c Helpers) CurrentUser() models.User {
  currentUser := models.User{}
  cookie, _ := c.Request.Cookie("authToken")

  if cookie != nil {
    coll := models.User{}.Coll()
    err := coll.Find(bson.M{"authtoken" : cookie.Value}).One(&currentUser)
    if err != nil {
      panic(err)
    }
    return currentUser
  }
  return currentUser
}
