package controllers

import(
  "fmt"
  "github.com/revel/revel"
  "labix.org/v2/mgo"
  "labix.org/v2/mgo/bson"
  "code.google.com/p/go.crypto/bcrypt"
  "reflect"
  "github.com/dchest/uniuri"
  //"time"
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
  fmt.Println(authToken)
  user := &User{Id: bson.NewObjectId(), Email: email, Password: bcryptPassword, AuthToken: authToken}
  err = coll.Insert(user)


  fmt.Println(bcryptPassword)
  fmt.Println(user.Password)
  fmt.Println(reflect.TypeOf(user.Password).Kind())
  fmt.Println(reflect.TypeOf(bcryptPassword).Kind())
  err = bcrypt.CompareHashAndPassword(user.Password, []byte(password))
   if err != nil {
      panic(err)
  }

  var updatedUser User
  coll.FindId(user.Id).One(&updatedUser)
  fmt.Println(updatedUser.Password)
  //fmt.Println(user.Id)
  revel.TRACE.Printf(email)
  revel.TRACE.Printf(password)
  if err != nil {
          panic(err)
  }


	return c.Redirect(Cards.Index)
}
