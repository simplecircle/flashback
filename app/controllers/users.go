package controllers

import(
  //"fmt"
  "github.com/revel/revel"
  //"labix.org/v2/mgo"
  //"labix.org/v2/mgo/bson"
)

type Users struct {
	*revel.Controller
}


func (c Users) New() revel.Result {
	return c.Render()
}

func (c Users) Create(email, password string) revel.Result {
  revel.TRACE.Printf(email)
  revel.TRACE.Printf(password)

	return c.Redirect(Cards.Index)
}
