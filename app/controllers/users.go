package controllers

import(
  "github.com/revel/revel"
  "labix.org/v2/mgo/bson"
  "code.google.com/p/go.crypto/bcrypt"
  "github.com/dchest/uniuri"
  "net/http"
  "time"
  "flashback/app/models"
  //"fmt"
)

type Users struct {
	*revel.Controller
  Helpers
}

func (c Users) New() revel.Result {
  return c.Render()
}

func (c Users) Create(email, password string) revel.Result {
  coll := models.User{}.Coll()
  bcryptPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
  authToken := uniuri.NewLen(22)

  user := models.NewUser()
  user.Id = bson.NewObjectId()
  user.Email = email
  user.Password = bcryptPassword
  user.AuthToken = authToken

  err := coll.Insert(user)
  if err != nil {
    panic(err)
  }

  c.SetCookie(&http.Cookie{
		Name:     "authToken",
		Value:    authToken,
		Path:     "/",
    Expires:  time.Now().Add(365*24*time.Hour), // One year cookie
	})
	return c.Redirect(Cards.Index)
}
