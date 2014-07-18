package controllers

import(
  "github.com/revel/revel"
  "labix.org/v2/mgo/bson"
  "code.google.com/p/go.crypto/bcrypt"
  "flashback/app/models"
  //"github.com/dchest/uniuri"
  //"net/http"
  //"fmt"
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

func (c Sessions) Create(email, password string) revel.Result {
  var user models.User
  coll := models.User{}.Coll()
  coll.Find(bson.M{"email": email}).One(&user)

  err := bcrypt.CompareHashAndPassword(user.Password, []byte(password))
  if err != nil {
    return c.Redirect(Sessions.New)
  }
  return c.Redirect(Cards.Index)
}
