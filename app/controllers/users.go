package controllers

import(
  //"fmt"
  "github.com/revel/revel"
  "labix.org/v2/mgo"
  "labix.org/v2/mgo/bson"
  "code.google.com/p/go.crypto/bcrypt"
  "github.com/dchest/uniuri"
  "net/http"
  "time"
  "flashback/app/models"
  //"reflect"
  //"strings"
)

type Users struct {
	*revel.Controller
  Helpers
}

func (c Users) New() revel.Result {
  return c.Render()
}

func (c Users) Create(email, password string) revel.Result {
  session, err := mgo.Dial("mongodb://elliottg:monkey75@kahana.mongohq.com:10026/flashbackDev")
  if err != nil {
          panic(err)
  }
  defer session.Close()
	session.SetMode(mgo.Monotonic, true)

  coll := session.DB("flashbackDev").C("users")
  bcryptPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
  authToken := uniuri.NewLen(22)
  user := &models.User{Id: bson.NewObjectId(), Email: email, Password: bcryptPassword, AuthToken: authToken}
  err = coll.Insert(user)
  if err != nil {
    panic(err)
  }

  c.SetCookie(&http.Cookie{
		Name:     "authToken",
		Value:    authToken,
		Path:     "/",
    Expires:  time.Now().Add(365*24*time.Hour), // One year cookie
	})

  //err = bcrypt.CompareHashAndPassword(user.Password, []byte(password))
   //if err != nil {
      //panic(err)
  //}
  //var updatedUser models.User
  //coll.FindId(user.Id).One(&updatedUser)
  //if err != nil {
          //panic(err)
  //}
  //fmt.Println(c.CurrentUser().Email)
	return c.Redirect(Cards.Index)
}
