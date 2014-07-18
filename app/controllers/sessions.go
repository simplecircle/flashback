package controllers

import(
  "github.com/revel/revel"
  "labix.org/v2/mgo/bson"
  "code.google.com/p/go.crypto/bcrypt"
  "flashback/app/models"
  "time"
  "net/http"
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
  c.SetCookie(&http.Cookie{
		Name:     "authToken",
		Value:    user.AuthToken,
		Path:     "/",
    Expires:  time.Now().Add(365*24*time.Hour), // One year cookie
	})
  return c.Redirect(Cards.Index)
}

func (c Sessions) Destroy() revel.Result {
  t := time.Date(1975, time.February, 2, 15, 0, 0, 0, time.Local)
  c.SetCookie(&http.Cookie{
		Name:     "authToken",
		Value:    "dead",
		Path:     "/",
    Expires:  t,
	})
  return c.Redirect(Sessions.New)
}
