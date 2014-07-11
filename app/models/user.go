package models

import(
  //"fmt"
  //"github.com/revel/revel"
  //"labix.org/v2/mgo"
  "labix.org/v2/mgo/bson"
)

type User struct {
  Id    bson.ObjectId `bson:"_id"`
  Email        string
  Password     []byte
  AuthToken    string
}


//func (c Users) CurrentUser() User {
  //session, err := mgo.Dial("mongodb://elliottg:monkey75@kahana.mongohq.com:10026/flashbackDev")
  //if err != nil {
          //panic(err)
  //}
  //defer session.Close()
  //cookie, _ := c.Request.Cookie("authToken")
  //cookieAuthToken := cookie.Value
  //currentUser := User{}
  //coll := session.DB("flashbackDev").C("users")
  //ifcoll.Find(bson.M{"authtoken" : cookieAuthToken}).One(&currentUser)
  //if err != nil {
          //panic(err)
  //}
  ////fmt.Println(currentUser.Email)
  //return currentUser
//}
