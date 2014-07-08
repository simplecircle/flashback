package controllers

import(
  "fmt"
  "github.com/revel/revel"
  "labix.org/v2/mgo"
  "labix.org/v2/mgo/bson"
  "code.google.com/p/go.crypto/bcrypt"
  "github.com/dchest/uniuri"
  "net/http"
  //"reflect"
  //"strings"
)

type Users struct {
	*revel.Controller
}
type User struct {
  Id    bson.ObjectId `bson:"_id"`
  Email        string
  Password     []byte
  AuthToken    string
}


func (c Users) New() revel.Result {
  return c.Render()
}

func (c Users) CurrentUser() User {
  session, err := mgo.Dial("mongodb://elliottg:monkey75@kahana.mongohq.com:10026/flashbackDev")
  if err != nil {
          panic(err)
  }
  defer session.Close()
  cookie, _ := c.Request.Cookie("authToken")
  cookieAuthToken := cookie.Value
  currentUser := User{}
  coll := session.DB("flashbackDev").C("users")
  coll.Find(bson.M{"authtoken" : cookieAuthToken}).One(&currentUser)
  if err != nil {
          panic(err)
  }
  //fmt.Println(currentUser.Email)
  return currentUser
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
  user := &User{Id: bson.NewObjectId(), Email: email, Password: bcryptPassword, AuthToken: authToken}
  err = coll.Insert(user)
  if err != nil {
    panic(err)
  }

  c.SetCookie(&http.Cookie{
		Name:     "authToken",
		Value:    authToken,
		Path:     "/",
	})
  //fmt.Println(http.ReadSetCookies(c.Request.Header))

  //fmt.Println(bcryptPassword)
  //fmt.Println(user.Password)
  //fmt.Println(reflect.TypeOf(user.Password).Kind())
  //fmt.Println(reflect.TypeOf(bcryptPassword).Kind())
  err = bcrypt.CompareHashAndPassword(user.Password, []byte(password))
   if err != nil {
      panic(err)
  }
  var updatedUser User
  coll.FindId(user.Id).One(&updatedUser)
  //fmt.Println(updatedUser.Password)
  ////fmt.Println(user.Id)
  //revel.TRACE.Printf(email)
  //revel.TRACE.Printf(password)
  if err != nil {
          panic(err)
  }


  fmt.Println(c.CurrentUser().Email)
	return c.Redirect(Cards.Index)
}
