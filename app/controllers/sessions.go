package controllers

import(
  //"fmt"
  "github.com/revel/revel"
  //"labix.org/v2/mgo"
  //"labix.org/v2/mgo/bson"
  //"code.google.com/p/go.crypto/bcrypt"
  //"github.com/dchest/uniuri"
  //"net/http"
  //"flashback/app/models"
  //"reflect"
  //"strings"
)

type Sessions struct {
	*revel.Controller
  Helpers
}

func (c Sessions) New() revel.Result {
  return c.Render()
}
func (c Sessions) Create() revel.Result {
  return c.Render()
}
